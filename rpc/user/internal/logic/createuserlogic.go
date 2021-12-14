package logic

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/model"
	"time"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

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
	if hasUserNameInfo != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户名已存在")
	} else if err == sqlx.ErrNotFound {
		return nil, status.Errorf(codes.Internal, "查询数据库错误")
	}

	hasUserEmailInfo, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if hasUserEmailInfo != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户邮箱已存在")
	} else if err == sqlx.ErrNotFound {
		return nil, status.Errorf(codes.Internal, "查询数据库错误")
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
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: false,
		},
		// TODO 检查一下数据库中的时间 有无插入
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
