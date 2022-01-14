package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/operation/model"

	"stockexchange/rpc/operation/internal/svc"
	"stockexchange/rpc/operation/operation"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserFavLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFavLogic {
	return &DeleteUserFavLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserFavLogic) DeleteUserFav(in *operation.UserFavRequest) (*operation.Empty, error) {
	// 这里使用到了硬删除，因为如果使用软删除，下一次用户再收藏这个商品会报唯一索引冲突
	if result := l.svcCtx.DbEngine.Unscoped().Where("user=? and stock=?", in.UserId, in.StockId).Delete(&model.UserFav{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}

	return &operation.Empty{}, nil
}
