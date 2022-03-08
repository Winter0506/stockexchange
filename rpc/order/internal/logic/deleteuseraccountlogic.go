package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAccountLogic {
	return &DeleteUserAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserAccountLogic) DeleteUserAccount(in *order.DeleteUserAccountInfo) (*order.DeleteUserAccountResponse, error) {
	// 这里进行硬删除以便于用户下次创建新的账户 新的金额
	_, err := l.svcCtx.UserAccountModel.FindOne(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "用户账户信息不存在")
	}
	err = l.svcCtx.UserAccountModel.Delete(in.UserId)
	if err != nil {
		return nil, errors.New("删除用户账户错误")
	}
	return &order.DeleteUserAccountResponse{
		Success: true,
	}, nil
}
