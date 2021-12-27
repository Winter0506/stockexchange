package logic

import (
	"context"
	"errors"
	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetStockByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByNameLogic {
	return &GetStockByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByNameLogic) GetStockByName(in *stock.NameRequest) (*stock.StockInfoResponse, error) {

	ret, err := l.svcCtx.Model.FindOneByStockname(in.StockName)
	// 只要查询出错, 就把责任推给 使用代码查询
	if err != nil {
		return nil, errors.New("请输入股票代码进行查询")
	}
	// 先查找信息  再在数据库中查询
	stockSlice, err := requestApi(ret.Stockcode)
	//if err == errors.New("错误股票代码") {
	//	return nil, err
	//}

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
	return &stock.StockInfoResponse{}, nil
}
