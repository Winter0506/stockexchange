package logic

import (
	"context"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
	"stockexchange/rpc/user/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.ReqUserLogin) (*types.RespUserLogin, error) {
	// todo: add your logic here and delete this line
	// 登录不用实现验证码逻辑, 注册的时候才需要 邮件验证码 和 图片验证码
	// fan返回的时候需要携带token回去

	// 判断用户有没有被删除
	resp, err := l.svcCtx.User.GetUserByEmail(l.ctx, &user.EmailRequest{Email: req.Email})
	logx.Infof("email: %v", l.ctx.Value("email")) // 这里的key和生成jwt token时传入的key一致
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				loginMessage := types.LoginMessage{}
				loginStatus := types.LoginMeta{
					Msg:    "邮件用户不存在",
					Status: http.StatusBadRequest,
				}
				return &types.RespUserLogin{
					LoginMessage: loginMessage,
					LoginMeta:    loginStatus,
				}, nil
			default:
				loginMessage := types.LoginMessage{}
				loginStatus := types.LoginMeta{
					Msg:    "登陆失败",
					Status: http.StatusInternalServerError,
				}
				return &types.RespUserLogin{
					LoginMessage: loginMessage,
					LoginMeta:    loginStatus,
				}, nil
			}
		}
	}

	if resp.IsDeleted == 1 {
		loginMessage := types.LoginMessage{}
		loginStatus := types.LoginMeta{
			Msg:    "无效用户",
			Status: http.StatusOK,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}

	// 还需要验证密码
	passwordResp, passwordErr := l.svcCtx.User.CheckPassWord(l.ctx, &user.PasswordCheckInfo{Password: req.Password,
		EncryptedPassword: resp.PassWord})
	if passwordErr != nil {
		loginMessage := types.LoginMessage{}
		loginStatus := types.LoginMeta{
			Msg:    "登陆失败",
			Status: http.StatusInternalServerError,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}
	if passwordResp.Success {
		// 生成token
		now := time.Now().Unix()
		// accessExpire := l.svcCtx.Config.Auth.AccessExpire
		jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, resp.Id)
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
			Msg:    "登陆成功",
			Status: http.StatusOK,
		}
		return &types.RespUserLogin{
			LoginMessage: loginMessage,
			LoginMeta:    loginStatus,
		}, nil
	}

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
