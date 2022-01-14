package operation

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"stockexchange/rpc/operation/operation"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserFavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserFavLogic {
	return UserFavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavLogic) UserFav(req types.ReqUserFavList) (*types.RespFavList, error) {
	resp, err := l.svcCtx.Operation.GetFavList(l.ctx, &operation.UserFavRequest{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("用户收藏列表查询失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "用户收藏列表查询失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespFavList{
			DetailMeta: detailStatus,
		}, errors.New("用户收藏列表查询失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "用户收藏列表查询成功",
		Status: http.StatusOK,
	}

	favList := make([]string, 0)
	for _, fav := range resp.Data {
		tempFav, _ := json.Marshal(fav)
		favList = append(favList, string(tempFav))
	}
	favMessage := types.FavMessage{
		Total:   int(resp.Total),
		FavList: favList,
	}
	return &types.RespFavList{
		FavMessage: favMessage,
		DetailMeta: detailStatus,
	}, nil
}
