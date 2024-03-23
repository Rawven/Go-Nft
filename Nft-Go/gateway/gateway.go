package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/gateway/internal/config"
	"Nft-Go/gateway/internal/handler"
	"Nft-Go/gateway/internal/svc"
	"context"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\CodeProjects\\Nft-Project\\Nft-Go\\gateway\\etc\\gateway.yaml", "the config file")

func main() {
	flag.Parse()
	util.InitConfig("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	api.InitGatewayService()
	db.InitRedis()
	util.InitLimiter(context.Background())
	var c config.Config
	conf.MustLoad(*configFile, &c)
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logger.Info("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
