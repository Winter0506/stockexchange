package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	// zrpc.RpcServerConf 表明继承了 rpc 服务端的配置
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}
