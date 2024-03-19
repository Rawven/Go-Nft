package api

import (
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/user"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

// 引入生成的接口结构
var grpcBlcImpl = new(blc.BlcRpcServiceClientImpl)
var userRpc user.UserClient

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func InitDubbo() {
	config.SetConsumerService(grpcBlcImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	logger.Info("dubbo init success")
}

func GetBlcDubbo() *blc.BlcRpcServiceClientImpl {
	return grpcBlcImpl
}

func InitUser() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	userRpc = user.NewUserClient(conn.Conn())
	logger.Info("user rpc load success")
}

func GetUserClient() user.UserClient {
	return userRpc
}
