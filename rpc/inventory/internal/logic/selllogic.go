package logic

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/inventory/model"

	"stockexchange/rpc/inventory/internal/svc"
	"stockexchange/rpc/inventory/inventory"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/tal-tech/go-zero/core/logx"
)

type SellLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellLogic {
	return &SellLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SellLogic) Sell(in *inventory.SellInfo) (*inventory.Empty, error) {
	// 扣减库存功能
	// 并发场景下，可能出现超卖
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "127.0.0.1:6379",
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	tx := l.svcCtx.DbEngine.Begin()
	sellDetail := model.StockSellDetail{
		TrustSn: in.TrustSn,
		Status:  1, // 设置为1表示已经扣减
	}
	var detail model.StockDetail
	detail = model.StockDetail{
		Stock: in.StockId,
		Num:   in.Num,
	}

	var inv model.Inventory
	mutex := rs.NewMutex(fmt.Sprintf("stock_%s", in.StockId))
	if err := mutex.Lock(); err != nil {
		return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
	}

	if result := l.svcCtx.DbEngine.Where(&model.Inventory{Stock: in.StockId}).First(&inv); result.RowsAffected == 0 {
		tx.Rollback() //回滚之前的操作
		return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
	}
	// 判断库存是否充足
	if inv.Total < in.Num {
		tx.Rollback() // 回滚操作
		return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
	}
	// 扣减库存
	inv.Total -= in.Num
	tx.Save(&inv)

	if ok, err := mutex.Unlock(); !ok || err != nil {
		return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
	}
	sellDetail.Detail = detail
	// 写sellDetail表
	if result := tx.Create(&sellDetail); result.RowsAffected == 0 {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
	}
	tx.Commit()
	return &inventory.Empty{}, nil
}
