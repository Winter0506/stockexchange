package main

import (
	"flag"
	"fmt"
	"stockexchange/rpc/user/internal/config"
	"stockexchange/rpc/user/internal/server"
	"stockexchange/rpc/user/internal/svc"
	"stockexchange/rpc/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

// 使用命令 goctl rpc template -o user.proto, 生成 user.proto 文件
// 使用命令 goctl rpc proto -src user.proto -dir . 生成 user rpc 服务的代码。
// 启动 go run user.go -f etc/user.yaml
// etcd 客户端 查询 etcdctl get user.rpc --prefix
// 比如proto文件修改之后  原来的代码再生成不会被更改?
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
