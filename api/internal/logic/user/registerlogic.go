package user

import (
	"context"
	"net/http"
	"stockexchange/api/internal/utils"
	"stockexchange/rpc/user/user"
	"time"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.ReqUserRegister) (*types.RespUserLogin, error) {
	if utils.Store.Verify(req.CaptchaId, req.Captcha, false) {
		loginMessage := types.LoginMessage{}
		loginStatus := types.LoginMeta{
			Msg:    "图片验证码错误",
			Status: http.StatusBadRequest,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}
	// 这个地方上边的逻辑应该是验证 短信/邮箱 验证码
	// 一开始用户 只能是普通用户  我不允许你是 管理员
	resp, err := l.svcCtx.User.CreateUser(l.ctx, &user.CreateUserInfo{
		UserName: req.UserName,
		PassWord: req.Password,
		Email:    req.Email,
		Gender:   req.Gender,
	})
	if err != nil {
		logx.Errorf("注册用户失败: ", err.Error())
		// 一开始先建立结构体?
		loginMessage := types.LoginMessage{}
		loginStatus := types.LoginMeta{
			Msg:    "注册用户失败",
			Status: http.StatusInternalServerError,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}

	// 生成token  准备返回用户信息
	now := time.Now().Unix()
	// accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, resp.Id)
	if err != nil {
		loginMessage := types.LoginMessage{}
		loginStatus := types.LoginMeta{
			Msg:    "生成token失败",
			Status: http.StatusInternalServerError,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}
	loginMessage := types.LoginMessage{
		Id:          resp.Id,
		Username:    resp.UserName,
		Password:    resp.PassWord,
		Email:       resp.Email,
		Gender:      resp.Gender,
		Role:        resp.Role,
		AccessToken: jwtToken,
	}
	loginStatus := types.LoginMeta{
		Msg:    "注册成功",
		Status: http.StatusOK,
	}
	return &types.RespUserLogin{
		LoginMessage: loginMessage,
		LoginMeta:    loginStatus,
	}, nil
}
