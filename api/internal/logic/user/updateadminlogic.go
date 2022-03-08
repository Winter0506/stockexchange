package user

import (
	"context"
	"net/http"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
	"stockexchange/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateAdminLogic {
	return UpdateAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(req types.ReqUserUpdateAdmin) (*types.RespUserLogin, error) {
	_, err := l.svcCtx.User.UpdateUser(l.ctx, &user.UpdateUserInfo{
		Id:        int64(req.Id),
		UserName:  req.UserName,
		PassWord:  req.Password,
		Email:     req.Email,
		Gender:    req.Gender,
		Role:      req.Role,
		IsDeleted: 0,
	})
	if err != nil {
		meta := types.LoginMeta{
			Msg:    "更新用户信息及权限成功",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("更新用户信息及权限成功", err)
		return &types.RespUserLogin{
			LoginMeta: meta,
		}, err
	} else {
		meta := types.LoginMeta{
			Msg:    "更新用户信息及权限成功",
			Status: http.StatusOK,
		}
		// 生成token
		//now := time.Now().Unix()
		//jwtToken, _ := utils.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(req.Id), 1)
		message := types.LoginMessage{
			// 信息用于返回的只是需要更新的字段
			Id:       int64(req.Id),
			Username: req.UserName,
			// TODO 这个地方 在rpc层也该分更新不更新密码  现在都是我自己做 所以我更新其他用户时候 是也把他们密码更新的
			Password: req.Password,
			Email:    req.Email,
			Gender:   req.Gender,
			Role:     req.Role,
			// AccessToken: jwtToken, 不需要
		}
		return &types.RespUserLogin{
			LoginMessage: message,
			LoginMeta:    meta,
		}, nil
	}
}
