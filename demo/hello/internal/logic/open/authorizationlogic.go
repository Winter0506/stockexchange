package logic

import (
	"context"

	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthorizationLogic struct {
	ctx context.Context
	logx.Logger
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorizationLogic {
	return AuthorizationLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthorizationLogic) Authorization(req types.UserOptReq) (*types.UserOptResp, error) {
	return &types.UserOptResp{}, nil
}
