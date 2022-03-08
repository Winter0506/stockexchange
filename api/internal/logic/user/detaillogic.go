package user

import (
	"context"
	"net/http"
	"stockexchange/rpc/user/user"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req types.ReqUserId) (*types.RespUserLogin, error) {
	rsp, err := l.svcCtx.User.GetUserById(l.ctx, &user.IdRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		meta := types.LoginMeta{
			Msg:    "获取用户详细信息失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("获取用户详细信息失败", err)
		return &types.RespUserLogin{
			LoginMeta: meta,
		}, err
	} else {
		meta := types.LoginMeta{
			Msg:    "获取用户详细信息成功",
			Status: http.StatusOK,
		}
		message := types.LoginMessage{
			Id:       rsp.Id,
			Username: rsp.UserName,
			Password: rsp.PassWord,
			Email:    rsp.Email,
			Gender:   rsp.Gender,
			Role:     rsp.Role,
		}
		return &types.RespUserLogin{
			LoginMessage: message,
			LoginMeta:    meta,
		}, nil
	}
}
