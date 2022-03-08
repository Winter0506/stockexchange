package main

import (
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"stockexchange/rpc/inventory/internal/logic"

	"stockexchange/rpc/inventory/internal/config"
	"stockexchange/rpc/inventory/internal/server"
	"stockexchange/rpc/inventory/internal/svc"
	"stockexchange/rpc/inventory/inventory"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/inventory.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewInventoryServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		inventory.RegisterInventoryServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//监听库存归还topic
	rocketmqConsumer, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("stockexchange-inventory"),
	)

	if err := rocketmqConsumer.Subscribe("order_reback", consumer.MessageSelector{}, logic.AutoReback); err != nil {
		fmt.Println("读取消息失败")
	}
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
	_ = rocketmqConsumer.Shutdown()
}
