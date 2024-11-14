package logic

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
	"tpmt-zt/common/global/twymqttdata"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksAddLogic {
	return &TpmtScheduledTasksAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自定义定时任务
func (l *TpmtScheduledTasksAddLogic) TpmtScheduledTasksAdd(in *tpmtclient.TpmtScheduledTasksAddReq) (*tpmtclient.CommonResp, error) {

	switch in.SchedulerTaskNumber {

	case 1:
		// 转换成已接任务结构
		var schedulerTaskNumber2Data twymqttdata.SchedulerTaskNumber2Data
		err := jsonx.Str2Struct(in.SchedulerData, &schedulerTaskNumber2Data)
		if err != nil {
			return nil, fmt.Errorf("数据格式错误:" + err.Error())
		}

	default:
		// 暂不支持其他任务
		return nil, fmt.Errorf("暂不支持其他任务")

	}

	// 添加定时任务
	_, err := l.svcCtx.TpmtScheduledTasksModel.Insert(l.ctx, &model.TpmtScheduledTasks{
		Id:                  uuid.NewV4().String(),  // ID
		CreatedAt:           time.Now(),             // 创建时间
		CreatedName:         in.CreatedName,         // 创建人
		SchedulerName:       in.SchedulerName,       // 名称
		SchedulerCategory:   in.SchedulerCategory,   // 类别 1:已接入任务, 2:自定义任务
		SchedulerTaskNumber: in.SchedulerTaskNumber, // 已接入任务号
		SchedulerType:       in.SchedulerType,       // 类型 1:Http任务,2:Webservices任务
		IntervalTime:        in.IntervalTime,        // 间隔时间按秒
		ErrorOrder:          in.ErrorOrder,          // 失败重新发送次数1-10次 不可超过10次
		FailIntervalTime:    in.FailIntervalTime,    // 失败间隔时间按秒
		State:               in.State,               // 状态 1:启动  2:暂停
		SchedulerData:       in.SchedulerData,       // 内容
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
