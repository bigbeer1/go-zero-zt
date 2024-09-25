package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/service/tpmt/api/internal/config"
	"tpmt-zt/service/tpmt/rpc/tpmt"
)

type ServiceContext struct {
	Config config.Config

	TpmtRpc tpmt.Tpmt
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TpmtRpc: tpmt.NewTpmt(zrpc.MustNewClient(c.TpmtRpc)),
	}
}
