package logic

import (
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"stockexchange/rpc/stock/model"
	"time"
)

func createStock(l *GetStockByCodeLogic, stockName, stockCode string) error {
	hasStockNameInfo, err := l.svcCtx.Model.FindOneByStockname(stockName)
	if hasStockNameInfo != nil {
		return errors.New("股票信息已存在")
	}
	hasStockCodeInfo, err := l.svcCtx.Model.FindOneByStockcode(stockCode)
	if hasStockCodeInfo != nil {
		return errors.New("股票信息已存在")
	}
	_, err = l.svcCtx.Model.Insert(&model.Stock{
		Stockname: stockName,
		Stockcode: stockCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: 0,
	})

	if err != nil {
		logx.Errorf("创建股票信息失败: ", err.Error())
		return errors.New("创建股票信息失败")
	}

	return nil
}
