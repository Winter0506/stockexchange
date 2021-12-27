package logic

import (
	"context"
	"database/sql"
	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/model"
	"stockexchange/rpc/user/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserInfo) (*user.Empty, error) {
	rsp, _ := l.svcCtx.Model.FindOne(in.Id)
	createdTime := rsp.CreatedAt
	updatedTime := rsp.UpdatedAt
	// 因为传进来的都为空
	if in.IsDeleted != 0 {
		err := l.svcCtx.Model.Update(&model.User{
			Id:        in.Id,
			Username:  in.UserName,
			Password:  in.PassWord,
			Email:     in.Email,
			Gender:    in.Gender,
			Role:      1, // TODO 更改用户为管理员应该有特别方法来执行
			CreatedAt: createdTime,
			UpdatedAt: updatedTime,
			DeletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			IsDeleted: 1,
		})
		if err != nil {
			return nil, err
		}
	} else {
		err := l.svcCtx.Model.Update(&model.User{
			Id:        in.Id,
			Username:  in.UserName,
			Password:  in.PassWord,
			Email:     in.Email,
			Gender:    in.Gender,
			Role:      1, // TODO 更改用户为管理员应该有特别方法来执行
			CreatedAt: createdTime,
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	return &user.Empty{}, nil
}
