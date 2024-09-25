package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	SysUserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库的连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		SysUserModel: model.NewSysUserModel(conn, c.CacheRedis), // ARole这张表数据库操作权限
	}
}
