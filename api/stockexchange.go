package main

import (
	"flag"
	"fmt"

	"stockexchange/api/internal/config"
	"stockexchange/api/internal/handler"
	"stockexchange/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/stockexchange.yaml", "the config file")

/*
除了 goctl 常用的命令需要熟练掌握，
go-zero 文件命名也是有规律可循的。
配置文件是放在 etc 目录下的 yaml 文件，该 yaml 文件对应的结构体在 internal/config/config.go 中。
依赖管理一般会在 internal/svc/servicecontext.go 中进行封装。
需要我们填充业务逻辑的地方是 internal/logic 目录下的文件
运行 go run stockexchange.go -f etc/stockexchange.yaml
*/
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
