package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string          // 手动添加
	Cache      cache.CacheConf // 手动添加
	Stock      zrpc.RpcClientConf
	Inventory  zrpc.RpcClientConf
}
