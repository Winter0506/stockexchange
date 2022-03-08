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

type InvDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvDetailLogic {
	return &InvDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvDetailLogic) InvDetail(in *inventory.StockInfo) (*inventory.StockInfo, error) {
	// 查询库存信息
	var inv model.Inventory
	if result := l.svcCtx.DbEngine.Where(&model.Inventory{Stock: in.StockId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该只股票没有库存信息")
	}
	return &inventory.StockInfo{
		StockId: inv.Stock,
		Num:     inv.Total,
	}, nil
}
