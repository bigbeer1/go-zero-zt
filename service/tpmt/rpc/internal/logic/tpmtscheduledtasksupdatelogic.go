package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksUpdateLogic {
	return &TpmtScheduledTasksUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtScheduledTasksUpdateLogic) TpmtScheduledTasksUpdate(in *tpmtclient.TpmtScheduledTasksUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtScheduledTasksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, errors.New("TpmtScheduledTasks没有该ID：" + in.Id)
		}
		return nil, err
	}

	// 名称
	if len(in.SchedulerName) > 0 {
		res.SchedulerName = in.SchedulerName
	}
	// 类别 1:已接入任务, 2:自定义任务
	if in.SchedulerCategory != 0 {
		res.SchedulerCategory = in.SchedulerCategory
	}
	// 已接入任务号
	if in.SchedulerTaskNumber != 0 {
		res.SchedulerTaskNumber = in.SchedulerTaskNumber
	}
	// 类型 1:Http任务,2:Webservices任务
	if in.SchedulerType != 0 {
		res.SchedulerType = in.SchedulerType
	}
	// 间隔时间按秒
	if in.IntervalTime != 0 {
		res.IntervalTime = in.IntervalTime
	}
	// 失败重新发送次数1-10次 不可超过10次
	if in.ErrorOrder != 0 {
		res.ErrorOrder = in.ErrorOrder
	}
	// 失败间隔时间按秒
	if in.FailIntervalTime != 0 {
		res.FailIntervalTime = in.FailIntervalTime
	}
	// 状态 1:启动  2:暂停
	if in.State != 0 {
		res.State = in.State
	}
	// 内容
	if len(in.SchedulerData) > 0 {
		res.SchedulerData = in.SchedulerData
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.TpmtScheduledTasksModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil
}
