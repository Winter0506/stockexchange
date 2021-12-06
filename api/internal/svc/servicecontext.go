package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"stockexchange/api/internal/config"
	"stockexchange/rpc/user/users"
)

type ServiceContext struct {
	Config config.Config
	// users.Users 是 user rpc 服务对外暴露的接口
	User   users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//  zrpc.MustNewClient(c.User) 创建了一个 grpc 客户端
		User:   users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
