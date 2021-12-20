package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockListLogic {
	return &GetStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockListLogic) GetStockList(in *user.PageInfo) (*user.StockListResponse, error) {
	// todo: add your logic here and delete this line

	return &user.StockListResponse{}, nil
}
