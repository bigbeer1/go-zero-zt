package alarm

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"
	"strconv"
	"time"
	"tpmt-zt/common/alarm"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/global"
	archivemodel "tpmt-zt/service/archive/model"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/websocket/websocketclient"
)

type DeviceAlarmHandler struct {
	svcCtx *svc.ServiceContext
}

func NewDeviceAlarmHandler(svcCtx *svc.ServiceContext) *DeviceAlarmHandler {
	return &DeviceAlarmHandler{
		svcCtx: svcCtx,
	}
}

func (l *DeviceAlarmHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var data model.TpmtMonitorPoint
	if err := jsonx.Unmarshal(t.Payload(), &data); err != nil {
		logx.Errorf("报警json TpmtMonitorPoint 转换失败 : ", t.Payload())
		return nil
	}

	// 查询时序数据库
	dataKey := fmt.Sprintf("sequential_id_%v",
		data.Id)

	tddb := &archivemodel.TdDb{
		DbName:    "monitor_point." + dataKey,
		TableName: "monitor_point.tpmt_monitor_point",
	}

	// 存储时序数据库记录检测值
	tdMonitor := archivemodel.TdMonitor{}

	startTime := time.Now().UnixMilli() - (data.AlarmDuration * 1000)

	count := tdMonitor.Count(ctx, l.svcCtx.Taos, tddb, startTime, time.Now().UnixMilli())

	if count < 2 {
		logx.Info("数据太少无法判断是否告警")
		return nil
	}

	monitorData, err := tdMonitor.FindAvgByTime(ctx, l.svcCtx.Taos, tddb, startTime, time.Now().UnixMilli())
	if err != nil {
		logx.Errorf("时序数据库查询监测点历史平均值用于告警失败！" + err.Error())
		return nil
	}

	if len(monitorData) == 0 {
		logx.Errorf("时序数据库查询监测点历史平均值用于告警失败！" + err.Error())
		return nil
	}

	// 乘以系数
	monitorValueFloatBig := big.NewFloat(monitorData[0].Data)
	monitorValueFloatBig.Mul(monitorValueFloatBig, big.NewFloat(data.Coefficient))
	bit := "%." + "2" + "f"
	monitorValue, _ := strconv.ParseFloat(fmt.Sprintf(bit, monitorValueFloatBig), 64)

	// 查询告警规则
	resultValueFloat64, err := datax.ToFloat64(monitorValue)
	if err != nil {
		logx.Errorf("resultValueFloat64转换失败！" + err.Error())
		return nil
	}

	info := alarm.CheckAlarmRuleYc(&data, resultValueFloat64)
	if info == nil {
		return nil
	}

	// 触发添加告警
	if info.Level != 0 {
		// 存储时序数据库
		alarmKey := fmt.Sprintf("alarm_monitor_id:%v",
			data.Id)

		var level int64
		// 添加redis 由于阻塞数据添加到 时序数据库 相当于锁
		l.svcCtx.Redis.GetCtx(ctx, alarmKey, &level)

		if level == 0 || info.Level > level {
			// 添加redis 由于阻塞数据添加到 时序数据库 相当于锁
			err := l.svcCtx.Redis.SetWithExpireCtx(ctx, alarmKey, info.Level, time.Duration(global.AlarmAddTime)*time.Hour)
			if err != nil {
				logx.Errorf("添加告警到redis失败！" + err.Error())
				return nil
			}

			alarmId := uuid.NewV4().String()
			// 添加到时序数据库中去
			alarmLog := archivemodel.AlarmLog{
				Ts:           time.Now(),
				Id:           alarmId,
				Mid:          datax.ToString(data.Id),
				Name:         data.Name,
				AlarmType:    info.RuleType,
				AlarmGrade:   info.Level,
				AlarmContent: info.RuleData,
				AssetId:      data.AssetId,
				AlarmState:   0,
			}

			alarmTddb := &archivemodel.TdDb{
				DbName:    "alarm_log." + "YcYm",
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
				Mid:          datax.ToString(data.Id),
				Name:         data.Name,
				AlarmType:    info.RuleType,
				AlarmGrade:   info.Level,
				AlarmContent: info.RuleData,
			})
		}

	}

	return nil
}
