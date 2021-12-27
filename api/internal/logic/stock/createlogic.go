package stock

import (
	"context"
	"errors"
	"net/http"
	"stockexchange/rpc/stock/stock"

	"stockexchange/api/internal/svc"
	"stockexchange/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.ReqStockCreate) (*types.RespStockDetail, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Stock.CreateStock(l.ctx, &stock.CreateStockInfo{
		StockName: req.StockName,
		StockCode: req.StockCode,
	})
	if err != nil {
		logx.Errorf("创建股票信息失败: ", err.Error())
		// 一开始先建立结构体?
		detailMessage := types.DetailMessage{}
		detailStatus := types.DetailMeta{
			Msg:    "请求股票详情失败",
			Status: http.StatusInternalServerError,
		}
		// 不能把本地错误暴露给外部用户
		return &types.RespStockDetail{
			DetailMessage: detailMessage,
			DetailMeta:    detailStatus,
		}, errors.New("create fail")
	}
	return &types.RespStockDetail{
		DetailMessage: types.DetailMessage{
			Id:               resp.Id,
			StockName:        resp.StockName,
			StockCode:        resp.StockCode,
			TodayOpenPrice:   resp.BaseInfo.TodayOpenPrice,
			LastClosePrice:   resp.BaseInfo.LastClosePrice,
			PresentPrice:     resp.BaseInfo.PresentPrice,
			HighPrice:        resp.BaseInfo.HighPrice,
			LowPrice:         resp.BaseInfo.LowPrice,
			CurrentBuyPrice:  resp.BaseInfo.CurrentBuyPrice,
			CurrentSellPrice: resp.BaseInfo.CurrentSellPrice,
			TransCount:       resp.BaseInfo.TransCount,
			TransAmount:      resp.BaseInfo.TransAmount,
			BuyOneCount:      resp.FiveBuyInfo.BuyOneCount,
			BuyOnePrice:      resp.FiveBuyInfo.BuyOnePrice,
			BuyTwoCount:      resp.FiveBuyInfo.BuyTwoCount,
			BuyTwoPrice:      resp.FiveBuyInfo.BuyTwoPrice,
			BuyThreeCount:    resp.FiveBuyInfo.BuyThreeCount,
			BuyThreePrice:    resp.FiveBuyInfo.BuyThreePrice,
			BuyFourCount:     resp.FiveBuyInfo.BuyFourCount,
			BuyFourPrice:     resp.FiveBuyInfo.BuyFourPrice,
			BuyFiveCount:     resp.FiveBuyInfo.BuyFiveCount,
			BuyFivePrice:     resp.FiveBuyInfo.BuyFivePrice,
			SellOneCount:     resp.FiveSellInfo.SellOneCount,
			SellOnePrice:     resp.FiveSellInfo.SellOnePrice,
			SellTwoCount:     resp.FiveSellInfo.SellTwoCount,
			SellTwoPrice:     resp.FiveSellInfo.SellTwoPrice,
			SellThreeCount:   resp.FiveSellInfo.SellThreeCount,
			SellThreePrice:   resp.FiveSellInfo.SellThreePrice,
			SellFourCount:    resp.FiveSellInfo.SellFourCount,
			SellFourPrice:    resp.FiveSellInfo.SellFourPrice,
			SellFiveCount:    resp.FiveSellInfo.SellFiveCount,
			SellFivePrice:    resp.FiveSellInfo.SellFivePrice,
			CurrentTime:      resp.Time.AsTime().Format("2006-01-02 15:04:05"),
		},
		DetailMeta: types.DetailMeta{
			Msg:    "创建股票成功",
			Status: http.StatusOK,
		},
	}, nil
}
