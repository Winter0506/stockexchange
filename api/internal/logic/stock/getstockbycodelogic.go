package stock

import (
	"context"
	"fmt"
	"net/http"
	"stockexchange/rpc/stock/stock"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetStockByCodeLogic {
	return GetStockByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockByCodeLogic) GetStockByCode(req types.ReqStockByCode) (*types.RespStockDetail, error) {
	rsp, err := l.svcCtx.Stock.GetStockByCode(l.ctx, &stock.CodeRequest{
		StockCode: req.StockCode,
	})
	fmt.Println(rsp, err)
	if err != nil {
		meta := types.DetailMeta{
			Msg:    "查询股票信息失败",
			Status: http.StatusInternalServerError,
		}
		logx.Errorf("查询股票信息失败", err)
		return &types.RespStockDetail{
			DetailMeta: meta,
		}, err
	}
	return &types.RespStockDetail{
		DetailMessage: types.DetailMessage{
			Id:               rsp.Id,
			StockName:        rsp.StockName,
			StockCode:        rsp.StockCode,
			TodayOpenPrice:   rsp.BaseInfo.TodayOpenPrice,
			LastClosePrice:   rsp.BaseInfo.LastClosePrice,
			PresentPrice:     rsp.BaseInfo.PresentPrice,
			HighPrice:        rsp.BaseInfo.HighPrice,
			LowPrice:         rsp.BaseInfo.LowPrice,
			CurrentBuyPrice:  rsp.BaseInfo.CurrentBuyPrice,
			CurrentSellPrice: rsp.BaseInfo.CurrentSellPrice,
			TransCount:       rsp.BaseInfo.TransCount,
			TransAmount:      rsp.BaseInfo.TransAmount,
			BuyOneCount:      rsp.FiveBuyInfo.BuyOneCount,
			BuyOnePrice:      rsp.FiveBuyInfo.BuyOnePrice,
			BuyTwoCount:      rsp.FiveBuyInfo.BuyTwoCount,
			BuyTwoPrice:      rsp.FiveBuyInfo.BuyTwoPrice,
			BuyThreeCount:    rsp.FiveBuyInfo.BuyThreeCount,
			BuyThreePrice:    rsp.FiveBuyInfo.BuyThreePrice,
			BuyFourCount:     rsp.FiveBuyInfo.BuyFourCount,
			BuyFourPrice:     rsp.FiveBuyInfo.BuyFourPrice,
			BuyFiveCount:     rsp.FiveBuyInfo.BuyFiveCount,
			BuyFivePrice:     rsp.FiveBuyInfo.BuyFivePrice,
			SellOneCount:     rsp.FiveSellInfo.SellOneCount,
			SellOnePrice:     rsp.FiveSellInfo.SellOnePrice,
			SellTwoCount:     rsp.FiveSellInfo.SellTwoCount,
			SellTwoPrice:     rsp.FiveSellInfo.SellTwoPrice,
			SellThreeCount:   rsp.FiveSellInfo.SellThreeCount,
			SellThreePrice:   rsp.FiveSellInfo.SellThreePrice,
			SellFourCount:    rsp.FiveSellInfo.SellFourCount,
			SellFourPrice:    rsp.FiveSellInfo.SellFourPrice,
			SellFiveCount:    rsp.FiveSellInfo.SellFiveCount,
			SellFivePrice:    rsp.FiveSellInfo.SellFivePrice,
			// 这个local()必须写
			CurrentTime: rsp.Time.AsTime().Local().Format("2006-01-02 15:04:05"),
		},
		DetailMeta: types.DetailMeta{
			Msg:    "请求股票信息成功",
			Status: http.StatusOK,
		},
	}, nil
}
