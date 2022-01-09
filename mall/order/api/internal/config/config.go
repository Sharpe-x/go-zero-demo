package config

import "github.com/tal-tech/go-zero/rest"
import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
}
