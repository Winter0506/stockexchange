package logic

import (
	"context"
	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO gorm中的分页实现
/*
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
*/

func (l *GetUserListLogic) GetUserList(in *user.PageInfo) (*user.UserListResponse, error) {
	// 自己逻辑上实现分页
	all, err := l.svcCtx.Model.FindAll()

	if err != nil {
		return nil, err
	}

	rsp := &user.UserListResponse{
		Total: 0,
		Data:  nil,
	}

	rsp.Total = int32(len(*all))
	allValue := *all
	page, pageSize := in.Pn, in.PSize
	tmpAll := allValue[(page-1)*pageSize : (page-1)*pageSize+pageSize]

	for _, eveRet := range tmpAll {
		userInfoRsp := &user.UserInfoResponse{
			Id:        eveRet.Id,
			UserName:  eveRet.Username,
			PassWord:  eveRet.Password,
			Email:     eveRet.Email,
			Gender:    eveRet.Gender,
			Role:      int32(eveRet.Role), // role 就是1 和 2 二者之一
			IsDeleted: int32(eveRet.IsDeleted),
		}
		rsp.Data = append(rsp.Data, userInfoRsp)
	}

	return rsp, nil
}
