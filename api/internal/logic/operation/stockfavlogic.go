package operation

import (
	"context"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StockFavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStockFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) StockFavLogic {
	return StockFavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StockFavLogic) StockFav(req types.ReqStockFav) (*types.RespFavList, error) {
	// todo: add your logic here and delete this line

	return &types.RespFavList{}, nil
}
