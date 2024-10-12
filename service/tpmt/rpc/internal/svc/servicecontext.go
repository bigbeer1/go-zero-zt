package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	SysUserModel     model.SysUserModel     // 用户
	SysUserRoleModel model.SysUserRoleModel // 用户和角色中间表

	SysAuthModel model.SysAuthModel // 第三方用户

	SysRoleModel model.SysRoleModel // 角色

	SysMenuModel     model.SysMenuModel     // 菜单
	SysRoleMenuModel model.SysRoleMenuModel // 菜单和角色中间表

	SysInterfaceModel     model.SysInterfaceModel     // 接口
	SysRoleInterfaceModel model.SysRoleInterfaceModel // 接口和角色中间表

	SysDictTypeModel model.SysDictTypeModel // 字典类型
	SysDictModel     model.SysDictModel     // 字典

	TpmtAssetModel model.TpmtAssetModel // 资产

	TpmtGatewayModel model.TpmtGatewayModel // 网关

	TpmtMonitorPointModel model.TpmtMonitorPointModel // 监测点
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库的连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                c,
		SysUserModel:          model.NewSysUserModel(conn, c.CacheRedis),
		SysUserRoleModel:      model.NewSysUserRoleModel(conn, c.CacheRedis),
		SysAuthModel:          model.NewSysAuthModel(conn, c.CacheRedis),
		SysRoleModel:          model.NewSysRoleModel(conn, c.CacheRedis),
		SysMenuModel:          model.NewSysMenuModel(conn, c.CacheRedis),
		SysRoleMenuModel:      model.NewSysRoleMenuModel(conn, c.CacheRedis),
		SysInterfaceModel:     model.NewSysInterfaceModel(conn, c.CacheRedis),
		SysRoleInterfaceModel: model.NewSysRoleInterfaceModel(conn, c.CacheRedis),
		SysDictTypeModel:      model.NewSysDictTypeModel(conn, c.CacheRedis),
		SysDictModel:          model.NewSysDictModel(conn, c.CacheRedis),
		TpmtAssetModel:        model.NewTpmtAssetModel(conn, c.CacheRedis),
		TpmtGatewayModel:      model.NewTpmtGatewayModel(conn, c.CacheRedis),
		TpmtMonitorPointModel: model.NewTpmtMonitorPointModel(conn, c.CacheRedis),
	}
}
