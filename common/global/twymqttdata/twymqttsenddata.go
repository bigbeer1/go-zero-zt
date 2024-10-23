package twymqttdata

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"sync"
	"time"
	"tpmt/common/eval"
	"tpmt/common/global"
	"tpmt/common/jsonx"
	archivemodel "tpmt/service/archive/model"
	"tpmt/service/mqttsend/mqttsend"
	"tpmt/service/mqttsend/mqttsendclient"
	"tpmt/service/tpmt/model"
)

type SchedulerTaskNumber2Data struct {
	GwSn      string `json:"gw_sn"`      // 网关序列号
	MsgInfo   string `json:"msg_info"`   // 消息内容介绍
	MqttHost  string `json:"mqtt_host"`  // mqtt目标服务器加端口 例子 10.192.38.21:1883
	MqttUser  string `json:"mqtt_user"`  // mqtt账号
	MqttPass  string `json:"mqtt_pass"`  // mqtt密码
	SendTopic string `json:"send_topic"` //  mqtt发送主题
}

//type SchedulerTaskNumber2Monitor struct {
//	Id            int64   `json:"id"`             // id
//	SerialNumber  string  `json:"serial_number"`  // 编号
//	Name          string  `json:"name"`           // 监测点名称
//	PointCategory int64   `json:"point_category"` // 类别：1:遥信/2:遥测/3:遥脉
//	Coefficient   float64 `json:"coefficient"`    // 系数
//}

type TwyData struct {
	GwSn      string            `json:"gw_sn"`
	MsgType   int               `json:"msg_type"`
	MsgInfo   string            `json:"msg_info"`
	TimeStamp int64             `json:"time_stamp,omitempty"`
	Data      map[string]string `json:"data"`
}

func (s SchedulerTaskNumber2Data) ConvertGetRealTimeValueAndSend(ctx context.Context, mqttSendRpc mqttsend.MqttSend, tpmtMonitorPointModel model.TpmtMonitorPointModel, taos *sql.DB, scheduledTasksId string, scheduledTasksLimit int64) {

	// 讲uuid中- 全部替换为_  确保插入和查询成功
	id := strings.Replace(scheduledTasksId, "-", "_", -10)

	tddb := &archivemodel.TdDb{
		DbName:    "scheduled_tasks_log.d1" + id,             // 普通表名称
		TableName: "scheduled_tasks_log.ptm_scheduled_tasks", // 超级表名称
	}

	var twyData TwyData
	twyData.GwSn = s.GwSn
	twyData.MsgInfo = s.MsgInfo
	twyData.TimeStamp = time.Now().UnixMilli()
	twyData.MsgType = 1001

	// 根据采集器ID 查询所有对应的监测点  删除
	whereBuilder := tpmtMonitorPointModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"sys_programme_id ": global.SysProgrammeId,
	})

	//  查询全部 当前任务下所有点位信息
	all, err := tpmtMonitorPointModel.FindAll(ctx, whereBuilder)
	if err != nil {
		// 写入时序数据库日志
		scheduledTasksLog := &archivemodel.ScheduledTasksLog{
			Ts:               time.Now(),
			ScheduledTasksId: scheduledTasksId,
			IsRequest:        2, // 失败
			RequestData:      "",
			ResponseData:     err.Error(),
		}
		scheduledTasksLog.Insert(ctx, taos, tddb)
		return
	}

	/*
		协程并发查询提升速度
	*/
	// 开启协程池

	var wg sync.WaitGroup
	// 协程获取返回数据
	chanData := make(chan map[string]string, len(all))
	// 协程数控制
	var limit = make(chan struct{}, scheduledTasksLimit)

	newMap := make(map[string]string)

	for _, item := range all {
		limit <- struct{}{}
		wg.Add(1)
		go func(item *model.TpmtMonitorPoint) {
			defer func() {
				wg.Done()
				<-limit
			}()
			var redisKey, resultValue string

			redisKey = fmt.Sprintf("action_time:id:%v",
				item.Id)
			// 根据key获取redis中的实时数据
			temp := tpmtMonitorPointModel.GetRealData(ctx, redisKey)

			// 获取实时数据和时间
			resultValue, _ = eval.MonitorValueEval(temp, item.Coefficient)

			if resultValue == "-" {
				resultValue = "0"
			}

			var datax = make(map[string]string)
			datax[item.SerialNumber] = resultValue
			data := datax
			chanData <- data
		}(item)

	}

	wg.Wait()
	close(chanData)
	for i := range chanData {
		for key, value := range i {
			newMap[key] = value
		}
	}

	twyData.Data = newMap

	s.MqttSendDo(ctx, scheduledTasksId, mqttSendRpc, taos, tddb, twyData)

}

func (s SchedulerTaskNumber2Data) MqttSendDo(ctx context.Context, scheduledTasksId string, mqttSendRpc mqttsend.MqttSend, taos *sql.DB, tddb *archivemodel.TdDb, twyData TwyData) {

	twyDataBytes, err := jsonx.Marshal(twyData)

	_, err = mqttSendRpc.TwyMqttSend(ctx, &mqttsendclient.TwyMqttSendReq{
		MqttClientId: scheduledTasksId,
		MqttHost:     s.MqttHost,
		MqttUser:     s.MqttUser,
		MqttPass:     s.MqttPass,
		SendTopic:    s.SendTopic,
		Data:         twyDataBytes,
	})
	if err != nil {
		// 写入时序数据库日志
		scheduledTasksLog := &archivemodel.ScheduledTasksLog{
			Ts:               time.Now(),
			ScheduledTasksId: scheduledTasksId,
			IsRequest:        2, // 失败
			RequestData:      string(twyDataBytes),
			ResponseData:     err.Error(),
		}
		scheduledTasksLog.Insert(ctx, taos, tddb)
		return
	}

	// 写入时序数据库日志
	scheduledTasksLog := &archivemodel.ScheduledTasksLog{
		Ts:               time.Now(),
		ScheduledTasksId: scheduledTasksId,
		IsRequest:        1, // 成功
		RequestData:      string(twyDataBytes),
		ResponseData:     "",
	}
	scheduledTasksLog.Insert(ctx, taos, tddb)
	return
}
