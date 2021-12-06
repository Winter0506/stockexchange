package logic

import (
	"context"
	"stockexchange/rpc/user/users"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req types.ReqUpdateUser) (*types.CommResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.User.Update(l.ctx, &users.User{Id: req.Id, Password: req.Password, Username: req.Username})
	if err != nil {
		return nil, err
	}
	return &types.CommResp{Ok: resp.Ok}, nil
}
