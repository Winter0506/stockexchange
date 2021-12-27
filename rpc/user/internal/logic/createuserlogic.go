package logic

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/model"
	"stockexchange/rpc/user/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserInfo) (*user.UserInfoResponse, error) {
	// 1.先查询用户id或者email是否已经被使用
	hasUserNameInfo, err := l.svcCtx.Model.FindOneByUsername(in.UserName)
	fmt.Println(hasUserNameInfo)
	if hasUserNameInfo != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户名已存在")
	}
	hasUserEmailInfo, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if hasUserEmailInfo != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户邮箱已存在")
	}

	// UserName PassWord Email Gender
	// 2.密码加密操作
	options := &password.Options{SaltLen: 8, Iterations: 10, KeyLen: 16, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(in.PassWord, options)
	// pbkdf2 是密钥算法
	passWord := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	ret, err := l.svcCtx.Model.Insert(&model.User{
		Username: in.UserName,
		Password: passWord,
		Email:    in.Email,
		Gender:   in.Gender,
		Role:     1,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	insertId, _ := ret.LastInsertId()
	return &user.UserInfoResponse{
		Id:       insertId,
		UserName: in.UserName,
		PassWord: in.PassWord,
		Email:    in.Email,
		Gender:   in.Gender,
		Role:     1,
	}, nil
}
