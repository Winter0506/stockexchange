package logic

import (
	"context"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.PageInfo) (*user.UserListResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserListResponse{}, nil
}
