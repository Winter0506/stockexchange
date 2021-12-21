package logic

import (
	"context"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *user.EmailRequest) (*user.UserInfoResponse, error) {
	ret, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if err != nil {
		return nil, err
	}
	// 不用去 判断用户是否被软删除 因为这些方法都是 管理员调用 也可以后面再写
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
