package store

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/eval"
	"tpmt-zt/common/global"
	archivemodel "tpmt-zt/service/archive/model"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/websocket/websocketclient"
)

type dataRedisSet struct {
	svcCtx *svc.ServiceContext
}

func NewDataRedisSetHandler(svcCtx *svc.ServiceContext) *dataRedisSet {
	return &dataRedisSet{
		svcCtx: svcCtx,
	}
}

func (l *dataRedisSet) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var data global.MonitorCacheUpData
	if err := jsonx.Unmarshal(t.Payload(), &data); err != nil {
		logx.Errorf("报警json CacheUpData 转换失败 : ", t.Payload())
		return nil
	}

	redisKey := fmt.Sprintf("action_time:id:%v",
		data.Id)

	// 判断数据类型是否yx 并开启了变位提醒
	if data.PointCategory == 1 && data.IsDisplacementWarning == 1 {
		// 根据key获取redis中的实时数据
		temp := l.svcCtx.TpmtMonitorPointModel.GetRealData(ctx, redisKey)
		// 获取实时数据和时间
		resultValue, _ := eval.MonitorValueEval(temp, 1)
		if resultValue != "-" && resultValue != data.Data {
			// 添加变位提醒
			alarmId := uuid.NewV4().String()
			alarmLog := archivemodel.AlarmLog{
				Ts:           time.Now(),
				Id:           alarmId,
				Mid:          data.Id,
				Name:         data.Name,
				AlarmType:    3,
				AlarmGrade:   3,
				AlarmContent: fmt.Sprintf("变位值:%v,当前值:%v", resultValue, data.Data),
				AssetId:      data.AssetId,
				AlarmState:   2,
			}

			alarmTddb := &archivemodel.TdDb{
				DbName:    "alarm_log." + "Yx",
				TableName: "alarm_log.ptm_alarm_log",
			}

			// 添加到时序数据库中去
			err := alarmLog.Insert(ctx, l.svcCtx.Taos, alarmTddb)
			if err != nil {
				logx.Errorf("时序数据库存储失败！" + err.Error())
			}

			// 创建60秒协程通知websocket
			ctxA, _ := context.WithTimeout(context.Background(), 60*time.Second)
			// 通知websocket
			go l.svcCtx.WebsocketRpc.AlarmMessage(ctxA, &websocketclient.AlarmMessageReq{
				Ts:           time.Now().UnixMilli(),
				Id:           alarmId,
				Mid:          data.Id,
				Name:         data.Name,
				AlarmType:    3,
				AlarmGrade:   3,
				AlarmContent: fmt.Sprintf("变位值:%v,当前值:%v", resultValue, data.Data),
			})
		}

	}

	actionData := fmt.Sprintf("%v|%v", data.Data, data.Ts.UnixMilli())
	err := l.svcCtx.Redis.SetWithExpireCtx(ctx, redisKey, actionData, l.svcCtx.RealDataTime)
	if err != nil {
		logx.Errorf("获取redis存入失败！")
	}

	return nil
}
