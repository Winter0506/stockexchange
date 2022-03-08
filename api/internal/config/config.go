package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// RpcClientConf 是 rpc 客户端的配置, 用来解析在 stockexchange.yaml 中的配置
	User       zrpc.RpcClientConf
	Stock      zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}
