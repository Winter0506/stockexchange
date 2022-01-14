package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// RpcClientConf 是 rpc 客户端的配置, 用来解析在 stockexchange.yaml 中的配置
	User       zrpc.RpcClientConf
	Stock      zrpc.RpcClientConf
	Operation  zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
