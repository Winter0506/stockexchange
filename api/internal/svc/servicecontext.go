package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"stockexchange/api/internal/config"
	"stockexchange/api/internal/middleware"
	"stockexchange/rpc/stock/stockclient"
	"stockexchange/rpc/user/userclient"
)

type ServiceContext struct {
	Config config.Config
	// user rpc 服务对外暴露的接口  要知道从哪里调用而来的
	User  userclient.User
	Stock stockclient.Stock
	Admin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//创建了一个 grpc 客户端
		User:  userclient.NewUser(zrpc.MustNewClient(c.User)),
		Stock: stockclient.NewStock(zrpc.MustNewClient(c.Stock)),
		Admin: middleware.NewAdminMiddleware().Handle,
	}
}
