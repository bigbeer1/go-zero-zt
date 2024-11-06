package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	TpmtAssetModel model.TpmtAssetModel // 资产

	TpmtGatewayModel model.TpmtGatewayModel // 网关

	TpmtMonitorPointModel model.TpmtMonitorPointModel // 监测点

	// 时序数据库连接
	Taos *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库的连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                c,
		Taos:                  tdenginex.NewTDengineManager(c.Tdengine),
		TpmtAssetModel:        model.NewTpmtAssetModel(conn, c.CacheRedis),
		TpmtGatewayModel:      model.NewTpmtGatewayModel(conn, c.CacheRedis),
		TpmtMonitorPointModel: model.NewTpmtMonitorPointModel(conn, c.CacheRedis),
	}
}
