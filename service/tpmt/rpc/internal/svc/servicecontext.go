package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	SysUserModel model.SysUserModel // 用户

	SysRoleModel model.SysRoleModel // 角色

	SysMenuModel     model.SysMenuModel     // 菜单
	SysRoleMenuModel model.SysRoleMenuModel // 菜单和角色中间表

	SysInterfaceModel     model.SysInterfaceModel     // 接口
	SysRoleInterfaceModel model.SysRoleInterfaceModel // 接口和角色中间表
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库的连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                c,
		SysUserModel:          model.NewSysUserModel(conn, c.CacheRedis),
		SysRoleModel:          model.NewSysRoleModel(conn, c.CacheRedis),
		SysMenuModel:          model.NewSysMenuModel(conn, c.CacheRedis),
		SysRoleMenuModel:      model.NewSysRoleMenuModel(conn, c.CacheRedis),
		SysInterfaceModel:     model.NewSysInterfaceModel(conn, c.CacheRedis),
		SysRoleInterfaceModel: model.NewSysRoleInterfaceModel(conn, c.CacheRedis),
	}
}
