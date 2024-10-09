package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleInterfaceModel = (*customSysRoleInterfaceModel)(nil)

type (
	// SysRoleInterfaceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleInterfaceModel.

	SysRoleInterfaceModel interface {
		sysRoleInterfaceModel
	}

	customSysRoleInterfaceModel struct {
		*defaultSysRoleInterfaceModel
	}
)

// NewSysRoleInterfaceModel returns a model for the database table.
func NewSysRoleInterfaceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleInterfaceModel {
	return &customSysRoleInterfaceModel{
		defaultSysRoleInterfaceModel: newSysRoleInterfaceModel(conn, c, opts...),
	}
}
