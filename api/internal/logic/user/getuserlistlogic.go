package user

import (
	"context"
	"encoding/json"
	"net/http"
	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserListLogic {
	return GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req types.ReqUserList) (*types.RespUserList, error) {
	// JS里面 JSON.parse("{\"id\":7,\"username\":\"王其超11\",\"password\":\"$pbkdf2-sha512$y5M2LXhY$6b5\",
	// \"email\":\"wangqichao11@gmail.com\",\"gender\":\"female\",\"role\":1,\"token\":\"\"}");
	rsp, err := l.svcCtx.User.GetUserList(l.ctx, &user.PageInfo{
		Pn:    uint32(req.Pn),
		PSize: uint32(req.PSize),
	})
	if err != nil {
		meta := types.LoginMeta{
			Msg:    "获取用户列表失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("获取用户列表失败", err)
		return &types.RespUserList{
			LoginMeta: meta,
		}, err
	}
	meta := types.LoginMeta{
		Msg:    "获取用户列表成功",
		Status: http.StatusOK,
	}
	userList := make([]string, 0)
	for _, user := range rsp.Data {
		user := types.LoginMessage{
			Id:          user.Id,
			Username:    user.UserName,
			Password:    user.PassWord,
			Email:       user.Email,
			Gender:      user.Gender,
			Role:        user.Role,
			AccessToken: "",
		}
		tempUser, _ := json.Marshal(user)
		userList = append(userList, string(tempUser))
	}
	return &types.RespUserList{
		UserList:  userList,
		LoginMeta: meta,
	}, nil
}
