package logic

import (
	"context"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPassWordLogic {
	return &CheckPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckPassWordLogic) CheckPassWord(in *user.PasswordCheckInfo) (*user.CheckResponse, error) {
	// todo: add your logic here and delete this line
	// 业务代码  1213学习以后再写
	return &user.CheckResponse{}, nil
}
