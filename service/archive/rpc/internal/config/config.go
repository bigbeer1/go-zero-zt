package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/common/tdenginex"
)

type Config struct {
	zrpc.RpcServerConf

	Tdengine tdenginex.TDengineConfig
}
