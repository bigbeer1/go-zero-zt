package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/global/twymqttdata"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/tpmt/model"
)

type scheduledTasksHandler struct {
	svcCtx *svc.ServiceContext
}

func NewScheduledTasksHandler(svcCtx *svc.ServiceContext) *scheduledTasksHandler {
	return &scheduledTasksHandler{
		svcCtx: svcCtx,
	}
}

type ScheduledTasksReq struct {
	TaskType  int64       `json:"task_type"`  // 任务类型 1：定时执行任务 2：重试任务
	BeginTime time.Time   `json:"begin_time"` //  启动时间
	Data      interface{} `json:"data"`       // 任务参数
}

func (l *scheduledTasksHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var in ScheduledTasksReq
	if err := jsonx.Unmarshal(t.Payload(), &in); err != nil {
		logx.Errorf("自定义任务json ScheduledTasksReq 转换失败 : ", t.Payload())
		return nil
	}

	// 判断任务类型
	switch in.TaskType {
	case 1:
		dataByte, _ := jsonx.Marshal(in.Data)
		var data model.TpmtScheduledTasks
		err := jsonx.Unmarshal(dataByte, &data)
		if err != nil {
			logx.Errorf("自定义任务 TwyScheduledTasks 转换失败 : ", t.Payload())
			return nil
		}
		err = l.TwyScheduledTasksDo(ctx, data, in.BeginTime)
		if err != nil {
			logx.Errorf("自定义任务 TwyScheduledTasksDo 执行失败 : ", err)
			return nil
		}
	case 2:
		dataByte, _ := jsonx.Marshal(in.Data)
		var data model.TpmtScheduledTasksFailureRecord
		err := jsonx.Unmarshal(dataByte, &data)
		if err != nil {
			logx.Errorf("重试任务 TwyScheduledTasksFailureRecord 转换失败 : ", t.Payload())
			return nil
		}
		err = l.TwyScheduledTasksFailureRecordDo(ctx, data, in.BeginTime)
		if err != nil {
			logx.Errorf("重试任务 TwyScheduledTasksFailureRecordDo 执行失败 : ", err)
			return nil
		}
	default:
		logx.Errorf("自定义任务 in.TaskType未知数据不正确: ", in)
		return nil

	}

	return nil

}

func (l *scheduledTasksHandler) TwyScheduledTasksDo(ctx context.Context, req model.TpmtScheduledTasks, now time.Time) error {

	switch req.SchedulerCategory {
	case 1: // 已接入任务
		switch req.SchedulerTaskNumber {
		case 1:
			// 转换成已接任务结构
			var schedulerTaskNumber2Data twymqttdata.SchedulerTaskNumber2Data
			err := jsonx.Str2Struct(req.SchedulerData, &schedulerTaskNumber2Data)
			if err != nil {
				return fmt.Errorf("数据格式错误:" + err.Error())
			}

			// 转换成实时数据去执行传输到mqttsend服务
			err = schedulerTaskNumber2Data.ConvertGetRealTimeValueAndSend(ctx, l.svcCtx.MqttSendRpc, l.svcCtx.TpmtMonitorPointModel, l.svcCtx.Taos, req.Id, l.svcCtx.Config.ScheduledTasksLimit, false)
			if err != nil {
				// 写入重试任务
				_, err = l.svcCtx.TpmtScheduledTasksFailureRecordModel.Insert(ctx, &model.TpmtScheduledTasksFailureRecord{
					Id:                  uuid.NewV4().String(), // ID
					ScheduledTasksId:    req.Id,
					CreatedAt:           now,                     // 创建时间
					CreatedName:         req.CreatedName,         // 创建人
					SchedulerName:       req.SchedulerName,       // 名称
					SchedulerCategory:   req.SchedulerCategory,   // 类别 1:已接入任务, 2:自定义任务
					SchedulerTaskNumber: req.SchedulerTaskNumber, // 已接入任务号
					SchedulerType:       req.SchedulerType,       // 类型 1:Http任务,2:Webservices任务
					ErrorOrder:          req.ErrorOrder,          // 失败重新发送次数1-10次 不可超过10次
					FailIntervalTime:    req.FailIntervalTime,    // 失败间隔时间按秒
					FailOrder:           0,
					SchedulerData:       req.SchedulerData, // 内容
				})

			}

		default:
			return errors.New(fmt.Sprintf("暂无已接入任务Id:%v", req.SchedulerTaskNumber))

		}

	case 2: // 自定义任务
		return errors.New(fmt.Sprintf("暂不支持自定义任务"))
	default:
		return errors.New(fmt.Sprintf("暂不支持除已接入任务和自定义任务以外任务"))
	}

	return nil

}

func (l *scheduledTasksHandler) TwyScheduledTasksFailureRecordDo(ctx context.Context, req model.TpmtScheduledTasksFailureRecord, now time.Time) error {

	switch req.SchedulerCategory {
	case 1: // 已接入任务
		switch req.SchedulerTaskNumber {
		case 1:
			// 转换成已接任务结构
			var schedulerTaskNumber2Data twymqttdata.SchedulerTaskNumber2Data
			err := jsonx.Str2Struct(req.SchedulerData, &schedulerTaskNumber2Data)
			if err != nil {
				return fmt.Errorf("数据格式错误:" + err.Error())
			}

			// 转换成实时数据去执行传输到mqttsend服务
			err = schedulerTaskNumber2Data.ConvertGetRealTimeValueAndSend(ctx, l.svcCtx.MqttSendRpc, l.svcCtx.TpmtMonitorPointModel, l.svcCtx.Taos, req.ScheduledTasksId, l.svcCtx.Config.ScheduledTasksLimit, true)
			if err == nil {
				// 写入重试任务
				err = l.svcCtx.TpmtScheduledTasksFailureRecordModel.Delete(ctx, req.Id)
			}

		default:
			return errors.New(fmt.Sprintf("暂无已接入任务Id:%v", req.SchedulerTaskNumber))

		}

	case 2: // 自定义任务
		return errors.New(fmt.Sprintf("暂不支持自定义任务"))
	default:
		return errors.New(fmt.Sprintf("暂不支持除已接入任务和自定义任务以外任务"))
	}

	return nil

}
