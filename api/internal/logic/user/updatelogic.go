package user

import (
	"context"
	"net/http"
	"stockexchange/api/internal/utils"
	"stockexchange/rpc/user/user"
	"time"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.ReqUserUpdate) (*types.RespUserLogin, error) {
	// 更新用户要求填入所有信息  这个地方做的有瑕疵
	// 其他好说 但是密码涉及到加密 在api层代入得是加密的 在rpc层又得加密
	// TODO:解决方法 1.每次修改时都需要输入原始密码 这样就顺势一起更新了  // 2.把更新密码 更新权限这两个功能拆分成不同方法
	_, err := l.svcCtx.User.UpdateUser(l.ctx, &user.UpdateUserInfo{
		Id:        int64(req.Id),
		UserName:  req.UserName,
		PassWord:  req.Password,
		Email:     req.Email,
		Gender:    req.Gender,
		IsDeleted: 0,
	})
	if err != nil {
		meta := types.LoginMeta{
			Msg:    "更新用户信息失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("更新信息失败", err)
		return &types.RespUserLogin{
			LoginMeta: meta,
		}, err
	} else {
		meta := types.LoginMeta{
			Msg:    "更新用户信息成功",
			Status: http.StatusOK,
		}
		// 生成token
		now := time.Now().Unix()
		jwtToken, _ := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(req.Id), 1)
		message := types.LoginMessage{
			// 信息用于返回的只是需要更新的字段
			Id:          int64(req.Id),
			Username:    req.UserName,
			Password:    req.Password,
			Email:       req.Email,
			Gender:      req.Gender,
			Role:        1,
			AccessToken: jwtToken,
		}
		return &types.RespUserLogin{
			LoginMessage: message,
			LoginMeta:    meta,
		}, nil
	}
}
