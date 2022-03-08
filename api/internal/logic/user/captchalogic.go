package user

import (
	"context"
	"net/http"
	"stockexchange/api/internal/utils"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (*types.RespCaptha, error) {
	// todo: add your logic here and delete this line
	id, b64s, err := utils.GetCaptcha()
	if err != nil {
		logx.Errorf("生成图片验证码错误: ", err.Error())
		return &types.RespCaptha{
			Msg:    "生成图片验证码错误",
			Status: http.StatusInternalServerError,
		}, err
	}
	return &types.RespCaptha{
		CaptchaId: id,
		Captcha:   b64s,
		Msg:       "生成验证码成功",
		Status:    http.StatusOK,
	}, nil
}
