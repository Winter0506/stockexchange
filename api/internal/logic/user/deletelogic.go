package user

import (
	"context"
	"net/http"
	"stockexchange/rpc/user/user"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLogic {
	return DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req types.ReqUserId) (*types.RespUserDelete, error) {
	_, err := l.svcCtx.User.UpdateUser(l.ctx, &user.UpdateUserInfo{
		Id:        int64(req.Id),
		IsDeleted: 1,
	})
	if err != nil {
		meta := types.LoginMeta{
			Msg:    "删除用户失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("删除用户失败", err)
		return &types.RespUserDelete{
			LoginMeta: meta,
		}, err
	} else {
		meta := types.LoginMeta{
			Msg:    "删除用户成功",
			Status: http.StatusOK,
		}
		return &types.RespUserDelete{
			LoginMeta: meta,
		}, nil
	}

}
