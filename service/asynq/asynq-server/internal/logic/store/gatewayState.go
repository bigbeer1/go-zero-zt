package store

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jsonx"
	archivemodel "tpmt-zt/service/archive/model"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/websocket/websocketclient"
)

type gatewayState struct {
	svcCtx *svc.ServiceContext
}

func NewGatewayStateHandler(svcCtx *svc.ServiceContext) *gatewayState {
	return &gatewayState{
		svcCtx: svcCtx,
	}
}

func (l *gatewayState) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var in global.GateWayOnlineUpData
	if err := jsonx.Unmarshal(t.Payload(), &in); err != nil {
		logx.Errorf("报警 json GateWayOnlineUpData 转换失败 : ", t.Payload())
		return nil
	}

	redisKey := fmt.Sprintf("gateway_state:gateway_id:%s",
		in.Id)

	err := l.svcCtx.Redis.SetWithExpireCtx(ctx, redisKey, in.OnlineStatus, l.svcCtx.RealDataTime)
	if err != nil {
		logx.Errorf("获取redis存入失败！")
	}

	if in.OnlineStatus == "1" {
		return nil
	}

	// 存储时序数据库 redis key
	alarmGateWayKey := fmt.Sprintf("alarm_gateway_id:%v",
		in.Id)

	var gatewayAlarm int64
	// 添加redis 由于阻塞数据添加到 时序数据库 相当于锁
	l.svcCtx.Redis.GetCtx(ctx, alarmGateWayKey, &gatewayAlarm)

	if gatewayAlarm == 0 {
		// 添加redis 由于阻塞数据添加到 时序数据库 相当于锁
		err := l.svcCtx.Redis.SetWithExpireCtx(ctx, alarmGateWayKey, 1, time.Duration(global.AlarmAddTime)*time.Hour)
		if err != nil {
			logx.Errorf("添加告警到redis失败！" + err.Error())
			return nil
		}

		alarmId := uuid.NewV4().String()
		// 添加到时序数据库中去
		alarmLog := archivemodel.AlarmLog{
			Ts:           time.Now(),
			Id:           alarmId,
			Mid:          in.Id,
			Name:         in.Name,
			AlarmType:    3,
			AlarmGrade:   3,
			AlarmContent: "",
			AlarmState:   2,
		}

		alarmTddb := &archivemodel.TdDb{
			DbName:    "alarm_log." + "GateWay",
			TableName: "alarm_log.ptm_alarm_log",
		}

		// 添加到时序数据库中去
		err = alarmLog.Insert(ctx, l.svcCtx.Taos, alarmTddb)
		if err != nil {
			logx.Errorf("时序数据库存储失败！" + err.Error())
		}

		// 创建60秒协程通知websocket
		ctxA, _ := context.WithTimeout(context.Background(), 60*time.Second)
		// 通知websocket
		go l.svcCtx.WebsocketRpc.AlarmMessage(ctxA, &websocketclient.AlarmMessageReq{
			Ts:           time.Now().UnixMilli(),
			Id:           alarmId,
			Mid:          in.Id,
			Name:         in.Name,
			AlarmType:    3,
			AlarmGrade:   3,
			AlarmContent: "",
			AssetCode:    "",
		})
	}

	return nil
}
