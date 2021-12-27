package logic

import (
	"context"
	"errors"
	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/model"
	"stockexchange/rpc/stock/stock"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockLogic {
	return &CreateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockLogic) CreateStock(in *stock.CreateStockInfo) (*stock.StockInfoResponse, error) {
	// 先查找信息  再在数据库中查询
	stockSlice, err := requestApi(in.StockCode)
	if err == errors.New("错误股票代码") {
		return nil, err
	}
	if err != nil {
		return nil, errors.New("创建股票错误")
	}

	ret, err := l.svcCtx.Model.Insert(&model.Stock{
		Stockname: in.StockName,
		Stockcode: in.StockCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: 0,
	})
	if err != nil {
		return nil, errors.New("创建股票错误")
	}

	baseInfo, fiveBuyInfo, fiveSellInfo := buildStockStruct(stockSlice)
	priceTime, err := TimestampProto(stockSlice[31], stockSlice[32])
	if err != nil {
		return nil, err
	}
	stockId, _ := ret.LastInsertId()
	return &stock.StockInfoResponse{
		Id:           int32(stockId),
		StockName:    in.StockName,
		StockCode:    in.StockCode,
		BaseInfo:     baseInfo,
		FiveBuyInfo:  fiveBuyInfo,
		FiveSellInfo: fiveSellInfo,
		Time:         priceTime,
	}, nil
}
