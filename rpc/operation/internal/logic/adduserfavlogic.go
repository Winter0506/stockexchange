package logic

import (
	"context"
	"stockexchange/rpc/demo/global"
	"stockexchange/rpc/operation/model"

	"stockexchange/rpc/operation/internal/svc"
	"stockexchange/rpc/operation/operation"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserFavLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserFavLogic {
	return &AddUserFavLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserFavLogic) AddUserFav(in *operation.UserFavRequest) (*operation.Empty, error) {
	var userFav model.UserFav

	userFav.User = in.UserId
	userFav.Stock = in.StockId

	// 这里只是用一下gorm 为了方便不将缓存放入redis
	global.DB.Save(&userFav)

	return &operation.Empty{}, nil
}
