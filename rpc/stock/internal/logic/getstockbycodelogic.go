package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"stockexchange/rpc/stock/internal/svc"
	"stockexchange/rpc/stock/stock"
)

type GetStockByCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByCodeLogic {
	return &GetStockByCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByCodeLogic) GetStockByCode(in *stock.CodeRequest) (*stock.StockInfoResponse, error) {
	// 先查找信息  再在数据库中查询
	stockSlice, err := requestApi(in.StockCode)
	//if err == errors.New("错误股票代码") {
	//	return nil, err
	//}
	if err != nil {
		return nil, errors.New("查询股票行情错误")
	}

	ret, err := l.svcCtx.Model.FindOneByStockcode(in.StockCode)

	// 只有这种情况才回去创建股票
	if ret == nil && err == sqlx.ErrNotFound {
		if err := createStock(l, stockSlice[0], in.StockCode); err != nil {
			return nil, errors.New("查询股票行情错误") // 实际上是创建错误 把信息不暴露给用户
		}
	}
	if err != nil {
		return nil, err
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
