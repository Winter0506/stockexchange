package main

import (
	"flag"
	"fmt"

	"stockexchange/rpc/operation/internal/config"
	"stockexchange/rpc/operation/internal/server"
	"stockexchange/rpc/operation/internal/svc"
	"stockexchange/rpc/operation/operation"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/operation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOperationServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		operation.RegisterOperationServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
