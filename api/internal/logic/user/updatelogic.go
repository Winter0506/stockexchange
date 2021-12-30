package user

import (
	"context"
	"net/http"
	"stockexchange/rpc/user/user"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
			Msg:    "获取用户详细信息成功",
			Status: http.StatusOK,
		}
		message := types.LoginMessage{
			// 信息用于返回的只是需要更新的字段
			Username: req.UserName,
			Password: req.Password,
			Email:    req.Email,
			Gender:   req.Gender,
			Role:     1,
		}
		return &types.RespUserLogin{
			LoginMessage: message,
			LoginMeta:    meta,
		}, nil
	}
}
