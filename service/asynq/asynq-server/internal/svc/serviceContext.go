package svc

import (
	"database/sql"
	"errors"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
	asynqx "tpmt-zt/common/asynq"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/asynq/asynq-server/internal/config"
	"tpmt-zt/service/mqttsend/mqttsend"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/websocket/websocket"
)

type ServiceContext struct {
	Config       config.Config
	AsynqServer  *asynq.Server
	AsynqClient  *asynq.Client
	RealDataTime time.Duration // 实时数据保持时间

	//TPMT监测点
	TpmtMonitorPointModel model.TpmtMonitorPointModel

	Redis cache.Cache

	// 全局缓存
	Cache *collection.Cache

	// websocket Rpc
	WebsocketRpc websocket.Websocket

	// mqtt-send Rpc
	MqttSendRpc mqttsend.MqttSend

	// Td连接
	Taos *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	c.WebsocketRpc.Timeout = 10000
	c.MqttSendRpc.Timeout = 30000

	return &ServiceContext{
		Config:                c,
		Taos:                  tdenginex.NewTDengineManager(c.Tdengine),
		Redis:                 cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("any"), errors.New("placeholder")),
		AsynqServer:           asynqx.NewAsynqServer(c.CacheRedis, c.Concurrency),
		TpmtMonitorPointModel: model.NewTpmtMonitorPointModel(conn, c.CacheRedis),
		AsynqClient:           asynqx.NewAsynqClient(c.CacheRedis),
		WebsocketRpc:          websocket.NewWebsocket(zrpc.MustNewClient(c.WebsocketRpc)),
		MqttSendRpc:           mqttsend.NewMqttSend(zrpc.MustNewClient(c.MqttSendRpc)),
	}
}
