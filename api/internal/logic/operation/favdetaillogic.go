package operation

import (
	"context"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FavDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) FavDetailLogic {
	return FavDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavDetailLogic) FavDetail(req types.ReqUserFav) (*types.RespFavDetail, error) {
	// todo: add your logic here and delete this line

	return &types.RespFavDetail{}, nil
}
