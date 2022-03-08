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

type StockFavLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStockFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) StockFavLogic {
	return StockFavLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StockFavLogic) StockFav(req types.ReqStockFavList) (*types.RespFavList, error) {
	// 这个需要管理员权限
	resp, err := l.svcCtx.Operation.GetFavList(l.ctx, &operation.UserFavRequest{
		StockId: req.StockId,
	})
	if err != nil {
		logx.Errorf("股票收藏列表查询失败: ", err.Error())
		detailStatus := types.DetailMeta{
			Msg:    "股票收藏列表查询失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespFavList{
			DetailMeta: detailStatus,
		}, errors.New("股票收藏列表查询失败")
	}
	detailStatus := types.DetailMeta{
		Msg:    "股票收藏列表查询成功",
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
