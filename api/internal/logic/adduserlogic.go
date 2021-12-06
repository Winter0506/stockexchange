package logic

import (
	"context"
	"stockexchange/rpc/user/users"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.ReqUser) (*types.CommResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.User.Create(l.ctx, &users.ReqUser{
		Username: req.Username,
		Password: req.Password,

	})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
