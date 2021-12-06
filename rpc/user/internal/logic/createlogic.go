package logic

import (
	"context"
	"stockexchange/rpc/model"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.ReqUser) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.Model.Insert(&model.User{
		Password: in.Password,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}

	return &user.CommResp{Ok:true}, nil
}
