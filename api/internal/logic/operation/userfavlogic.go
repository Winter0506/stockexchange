package operation

import (
	"context"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserFavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserFavLogic {
	return UserFavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavLogic) UserFav(req types.ReqUserFav) (*types.RespFavList, error) {
	// todo: add your logic here and delete this line

	return &types.RespFavList{}, nil
}
