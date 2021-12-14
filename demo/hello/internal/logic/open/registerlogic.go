package logic

import (
	"context"

	"hello/internal/models"
	"hello/internal/svc"
	"hello/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserOptReq) (*types.UserOptResp, error) {
	user := models.User{
		Email:  req.Email,
		Passwd: req.Passwd,
	}
	result := l.svcCtx.DbEngin.Create(&user)
	return &types.UserOptResp{
		Id: user.ID,
	}, result.Error
}
