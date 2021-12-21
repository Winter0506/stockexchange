package logic

import (
	"context"
	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockLogic {
	return &CreateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockLogic) CreateStock(in *stock.CreateStockInfo) (*stock.StockInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 第一次 输入代码时候才 创建 这个创建先不写
	return &stock.StockInfoResponse{}, nil
}
