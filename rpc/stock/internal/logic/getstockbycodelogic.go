package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByCodeLogic {
	return &GetStockByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByCodeLogic) GetStockByCode(in *user.CodeRequest) (*user.StockInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.StockInfoResponse{}, nil
}
