package logic

import (
	"context"
	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.NameRequest) (*user.UserInfoResponse, error) {
	ret, err := l.svcCtx.Model.FindOneByUsername(in.Name)
	if err != nil {
		return nil, err
	}
	return &user.UserInfoResponse{
		Id:        ret.Id,
		UserName:  ret.Username,
		PassWord:  ret.Password,
		Email:     ret.Email,
		Gender:    ret.Gender,
		Role:      int32(ret.Role), // role 就是1 和 2 二者之一
		IsDeleted: int32(ret.IsDeleted),
	}, nil
}
