package logic

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"

	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

const secretKey = "stockexchange"

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.ReqUser) (*user.RespLogin, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.Model.FindByName(in.Username)
	if err != nil {
		return nil, errors.Wrapf(err, "FindUser %s", in.Username)
	}

	if one.Password != in.Password {
		return nil, fmt.Errorf("user or password is invalid")
	}

	token := GenTokenByHmac(one.Username, secretKey)
	return &user.RespLogin{Token: token}, nil
	return &user.RespLogin{}, nil
}

func HmacCrypto(s, key string) string {
	hc := hmac.New(sha1.New, []byte(key))
	token := hc.Sum([]byte(s))
	return hex.EncodeToString(token)
}

func GenTokenByHmac(s, key string) string {
	return fmt.Sprintf("%s,%s", s, HmacCrypto(s, key))
}
