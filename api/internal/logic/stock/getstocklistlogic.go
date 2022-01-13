package stock

import (
	"context"
	"encoding/json"
	"net/http"
	"stockexchange/rpc/stock/stock"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserListLogic {
	return GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetStockList(req types.ReqStockList) (*types.RespStockList, error) {
	rsp, err := l.svcCtx.Stock.GetStockList(l.ctx, &stock.PageInfo{
		Pn:    uint32(req.Pn),
		PSize: uint32(req.PSize),
	})
	if err != nil {
		meta := types.ListMeta{
			Msg:    "获取股票列表失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("获取股票列表失败", err)
		return &types.RespStockList{
			StockList: nil,
			ListMeta:  meta,
		}, err
	}
	meta := types.ListMeta{
		Msg:    "获取股票列表成功",
		Status: http.StatusOK,
	}
	stockList := make([]string, 0)
	for _, stock := range rsp.Data {
		user := types.StockMessage{
			Id:        stock.Id,
			StockName: stock.StockName,
			StockCode: stock.StockCode,
		}
		tempStock, _ := json.Marshal(user)
		stockList = append(stockList, string(tempStock))
	}
	return &types.RespStockList{
		StockList: stockList,
		ListMeta:  meta,
	}, nil
}
