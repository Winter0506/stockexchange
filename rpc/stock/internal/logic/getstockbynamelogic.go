package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/user"

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

func (l *GetStockByNameLogic) GetStockByName(in *user.NameRequest) (*user.StockInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.StockInfoResponse{}, nil
}
