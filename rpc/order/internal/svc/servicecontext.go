package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"stockexchange/rpc/inventory/inventoryclient"
	"stockexchange/rpc/order/internal/config"
	"stockexchange/rpc/order/model"
	"stockexchange/rpc/stock/stockclient"
)

type ServiceContext struct {
	Config            config.Config
	UserAccountModel  model.UseraccountModel
	HoldPositionModel model.HoldpositionModel
	TrustModel        model.TrustModel
	OrderModel        model.OrderModel // 手动添加
	// 股票微服务
	Stock stockclient.Stock
	// 库存微服务
	Inventory inventoryclient.Inventory
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserAccountModel:  model.NewUseraccountModel(sqlx.NewMysql(c.DataSource), c.Cache),
		HoldPositionModel: model.NewHoldpositionModel(sqlx.NewMysql(c.DataSource), c.Cache),
		TrustModel:        model.NewTrustModel(sqlx.NewMysql(c.DataSource), c.Cache),
		OrderModel:        model.NewOrderModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动添加
		Stock:             stockclient.NewStock(zrpc.MustNewClient(c.Stock)),
		Inventory:         inventoryclient.NewInventory(zrpc.MustNewClient(c.Inventory)),
	}
}
