package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByNameLogic {
	return &GetStockByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByNameLogic) GetStockByName(in *stock.NameRequest) (*stock.StockInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &stock.StockInfoResponse{}, nil
}
