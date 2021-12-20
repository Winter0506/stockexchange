package logic

import (
	"context"

	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByIdLogic {
	return &GetStockByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByIdLogic) GetStockById(in *user.IdRequest) (*user.StockInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.StockInfoResponse{}, nil
}
