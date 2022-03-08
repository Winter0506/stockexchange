package logic

import (
	"context"
	"errors"
	"fmt"
	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByIdLogic {
	return &GetStockByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByIdLogic) GetStockById(in *stock.IdRequest) (*stock.StockInfoResponse, error) {
	ret, err := l.svcCtx.Model.FindOne(in.Id)
	if err != nil {
		return nil, err
	}
	stockSlice, err := requestApi(ret.Stockcode)
	fmt.Println(stockSlice)
	//if err == errors.New("错误股票代码") {
	//	return nil, err
	//}
	if err != nil {
		return nil, errors.New("查询股票行情错误，请重试")
	}

	baseInfo, fiveBuyInfo, fiveSellInfo := buildStockStruct(stockSlice)
	priceTime, err := TimestampProto(stockSlice[31], stockSlice[32])
	if err != nil {
		return nil, err
	}

	return &stock.StockInfoResponse{
		Id:           ret.Id,
		StockName:    ret.Stockname,
		StockCode:    ret.Stockcode,
		BaseInfo:     baseInfo,
		FiveBuyInfo:  fiveBuyInfo,
		FiveSellInfo: fiveSellInfo,
		Time:         priceTime,
	}, nil
}
