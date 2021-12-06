package logic

import (
	"context"
	"stockexchange/rpc/user/users"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.ReqUserId) (*types.User, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.User.Get(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.User{
		Id:       resp.Id,
		Username: resp.Username,
		Password: resp.Password,
	}, nil
}
