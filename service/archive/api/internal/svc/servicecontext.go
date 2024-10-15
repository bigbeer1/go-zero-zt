package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/service/archive/api/internal/config"
	"tpmt-zt/service/archive/rpc/archive"
)

type ServiceContext struct {
	Config config.Config

	// 日志服务
	ArchiveRpc archive.Archive
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArchiveRpc: archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)),
	}
}
