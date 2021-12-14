package logic

import (
	"context"
	"database/sql"
	"stockexchange/rpc/model"
	"stockexchange/rpc/user/internal/svc"
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
	// todo: add your logic here and delete this line
	if in.IsDeleted == 1 {
		err := l.svcCtx.Model.Update(&model.User{
			DeletedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
			// TODO 检查一下数据库中的时间 有无插入
		})
		if err != nil {
			return nil, err
		}
	} else {
		err := l.svcCtx.Model.Update(&model.User{
			Username: in.UserName,
			Password: in.PassWord,
			Email:    in.Email,
			Gender:   in.Gender, // TODO 更改用户为管理员应该有特别方法来执行
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: false,
			},
			// TODO 检查一下数据库中的时间 有无插入
		})
		if err != nil {
			return nil, err
		}
	}
	return &user.Empty{}, nil
}
