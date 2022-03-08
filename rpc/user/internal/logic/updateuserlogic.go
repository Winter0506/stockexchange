package logic

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
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
			Username:  rsp.Username,
			Password:  rsp.Password,
			Email:     rsp.Email,
			Gender:    rsp.Gender,
			Role:      rsp.Role,
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
		var role int
		if in.Role == 0 {
			role = rsp.Role
		} else {
			role = 2
		}
		// 更新的时候也要加密密码
		options := &password.Options{SaltLen: 8, Iterations: 10, KeyLen: 16, HashFunction: sha512.New}
		salt, encodedPwd := password.Encode(in.PassWord, options)
		// pbkdf2 是密钥算法
		passWord := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
		err := l.svcCtx.Model.Update(&model.User{
			Id:        in.Id,
			Username:  in.UserName,
			Password:  passWord,
			Email:     in.Email,
			Gender:    in.Gender,
			Role:      role,
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
