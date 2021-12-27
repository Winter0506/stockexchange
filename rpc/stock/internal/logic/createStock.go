package logic

import (
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"stockexchange/rpc/stock/model"
	"time"
)

func createStock(l *GetStockByCodeLogic, stockName, stockCode string) error {
	_, err := l.svcCtx.Model.Insert(&model.Stock{
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
