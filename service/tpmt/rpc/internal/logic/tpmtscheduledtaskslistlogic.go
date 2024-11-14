package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksListLogic {
	return &TpmtScheduledTasksListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtScheduledTasksListLogic) TpmtScheduledTasksList(in *tpmtclient.TpmtScheduledTasksListReq) (*tpmtclient.TpmtScheduledTasksListResp, error) {

	whereBuilder := l.svcCtx.TpmtScheduledTasksModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 名称
	if len(in.SchedulerName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"scheduler_name ": in.SchedulerName,
		})
	}
	// 类别 1:已接入任务, 2:自定义任务
	if in.SchedulerCategory != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"scheduler_category ": in.SchedulerCategory,
		})
	}
	// 已接入任务号
	if in.SchedulerTaskNumber != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"scheduler_task_number ": in.SchedulerTaskNumber,
		})
	}
	// 类型 1:Http任务,2:Webservices任务
	if in.SchedulerType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"scheduler_type ": in.SchedulerType,
		})
	}
	// 状态 1:启动  2:暂停
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.TpmtScheduledTasksModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TpmtScheduledTasksModel.CountBuilder("id")

	// 名称
	if len(in.SchedulerName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"scheduler_name ": in.SchedulerName,
		})
	}
	// 类别 1:已接入任务, 2:自定义任务
	if in.SchedulerCategory != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"scheduler_category ": in.SchedulerCategory,
		})
	}
	// 已接入任务号
	if in.SchedulerTaskNumber != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"scheduler_task_number ": in.SchedulerTaskNumber,
		})
	}
	// 类型 1:Http任务,2:Webservices任务
	if in.SchedulerType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"scheduler_type ": in.SchedulerType,
		})
	}
	// 状态 1:启动  2:暂停
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	count, err := l.svcCtx.TpmtScheduledTasksModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.TpmtScheduledTasksListData
	for _, item := range all {
		list = append(list, &tpmtclient.TpmtScheduledTasksListData{
			Id:                  item.Id,                         //定时任务ID
			CreatedAt:           item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:           item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:         item.CreatedName,                //创建人
			UpdatedName:         item.UpdatedName.String,         //更新人
			SchedulerName:       item.SchedulerName,              //名称
			SchedulerCategory:   item.SchedulerCategory,          //类别 1:已接入任务, 2:自定义任务
			SchedulerTaskNumber: item.SchedulerTaskNumber,        //已接入任务号
			SchedulerType:       item.SchedulerType,              //类型 1:Http任务,2:Webservices任务
			IntervalTime:        item.IntervalTime,               //间隔时间按秒
			ErrorOrder:          item.ErrorOrder,                 //失败重新发送次数1-10次 不可超过10次
			FailIntervalTime:    item.FailIntervalTime,           //失败间隔时间按秒
			State:               item.State,                      //状态 1:启动  2:暂停
			SchedulerData:       item.SchedulerData,              //内容
		})
	}

	return &tpmtclient.TpmtScheduledTasksListResp{
		Total: count,
		List:  list,
	}, nil
}
