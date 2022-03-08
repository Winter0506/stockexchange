package logic

import (
	"context"
	"stockexchange/rpc/inventory/model"

	"stockexchange/rpc/inventory/internal/svc"
	"stockexchange/rpc/inventory/inventory"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetInvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInvLogic {
	return &SetInvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetInvLogic) SetInv(in *inventory.StockInfo) (*inventory.Empty, error) {
	// 设置库存
	var inv model.Inventory
	l.svcCtx.DbEngine.Where(&model.Inventory{Stock: in.StockId}).First(&inv)
	inv.Stock = in.StockId
	inv.Total = in.Num
	l.svcCtx.DbEngine.Save(&inv)
	return &inventory.Empty{}, nil
}
