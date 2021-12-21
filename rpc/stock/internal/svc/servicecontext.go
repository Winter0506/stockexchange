package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"stockexchange/rpc/model"
	"stockexchange/rpc/stock/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.StockModel // 手动添加
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewStockModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动添加
	}
}
