package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"stockexchange/rpc/stock/internal/config"
	"stockexchange/rpc/stock/model"
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
