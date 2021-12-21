package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	// 先查找信息  再在数据库中查询
	stockSlice, err := requestApi(in.Code)
	if err == errors.New("错误股票代码") {
		return nil, err
	}
	if err != nil {
		return nil, errors.New("查询股票行情错误")
	}

	ret, err := l.svcCtx.Model.FindOneByStockcode(in.Code)
	// TODO 在上层 如果没有 就创建 并返回
	/*if ret == nil && err == sqlx.ErrNotFound {
		// 说明数据库里没有 我们需要去添加
		// 调用 创建方法 在上一层中做
	}*/
	if err != nil {
		return nil, err
	}

	baseInfo, fiveBuyInfo, fiveSellInfo := buildStockStruct(stockSlice)
	priceTime, err := TimestampProto(stockSlice[32])
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
