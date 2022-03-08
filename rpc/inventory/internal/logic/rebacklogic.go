package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/inventory/model"

	"stockexchange/rpc/inventory/internal/svc"
	"stockexchange/rpc/inventory/inventory"

	"github.com/tal-tech/go-zero/core/logx"
)

type RebackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRebackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RebackLogic {
	return &RebackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RebackLogic) Reback(in *inventory.SellInfo) (*inventory.Empty, error) {
	// 库存归还 三种情况
	// 1.订单超时归还 2.订单创建失败 归还之前扣减的库存 3.手动归还
	tx := l.svcCtx.DbEngine.Begin()
	var inv model.Inventory
	if result := l.svcCtx.DbEngine.Where(&model.Inventory{Stock: in.StockId}).First(&inv); result.RowsAffected == 0 {
		tx.Rollback() // 回滚操作
		return nil, status.Errorf(codes.InvalidArgument, "无该只股票库存信息")

		inv.Total += in.Num
		tx.Save(&inv)
	}
	tx.Commit() // 提交事务
	return &inventory.Empty{}, nil
}
