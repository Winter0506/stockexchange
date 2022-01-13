package logic

import (
	"context"
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"strings"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

// 封装一层 user业务在这里编写
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
	// 校验密码
	options := &password.Options{SaltLen: 8, Iterations: 10, KeyLen: 16, HashFunction: sha512.New}
	passwordInfo := strings.Split(in.EncryptedPassword, "$")
	check := password.Verify(in.Password, passwordInfo[2], passwordInfo[3], options)
	return &user.CheckResponse{Success: check}, nil
}
