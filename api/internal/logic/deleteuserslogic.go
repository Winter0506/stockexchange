package logic

import (
	"context"
	"stockexchange/rpc/user/users"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUsersLogic {
	return DeleteUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUsersLogic) DeleteUsers(req types.ReqUserId) (*types.CommResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.User.Delete(l.ctx, &users.ReqUserId{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.CommResp{Ok: resp.Ok}, nil
}
