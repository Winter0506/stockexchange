package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/demo/global"
	"stockexchange/rpc/operation/model"

	"stockexchange/rpc/operation/internal/svc"
	"stockexchange/rpc/operation/operation"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserFavDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavDetailLogic {
	return &GetUserFavDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavDetailLogic) GetUserFavDetail(in *operation.UserFavRequest) (*operation.Empty, error) {
	var userFav model.UserFav

	if result := global.DB.Where("goods=? and user=?", in.UserId, in.StockId).Find(&userFav); result.RowsAffected == 0{
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}

	return &operation.Empty{}, nil
}
