package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksFailureRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksFailureRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksFailureRecordListLogic {
	return &TpmtScheduledTasksFailureRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 失败任务且还在重试的任务
func (l *TpmtScheduledTasksFailureRecordListLogic) TpmtScheduledTasksFailureRecordList(in *tpmtclient.TpmtScheduledTasksFailureRecordListReq) (*tpmtclient.TpmtScheduledTasksFailureRecordListResp, error) {

	whereBuilder := l.svcCtx.TpmtScheduledTasksFailureRecordModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 任务ID
	if len(in.ScheduledTasksId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"scheduled_tasks_id ": in.ScheduledTasksId,
		})
	}

	all, err := l.svcCtx.TpmtScheduledTasksFailureRecordModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TpmtScheduledTasksFailureRecordModel.CountBuilder("id")

	// 任务ID
	if len(in.ScheduledTasksId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"scheduled_tasks_id ": in.ScheduledTasksId,
		})
	}
	count, err := l.svcCtx.TpmtScheduledTasksFailureRecordModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.TpmtScheduledTasksFailureRecordListData
	for _, item := range all {
		list = append(list, &tpmtclient.TpmtScheduledTasksFailureRecordListData{
			Id:                  item.Id,                         //失败记录ID
			ScheduledTasksId:    item.ScheduledTasksId,           //任务ID
			CreatedAt:           item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:           item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:         item.CreatedName,                //创建人
			UpdatedName:         item.UpdatedName.String,         //更新人
			SchedulerName:       item.SchedulerName,              //名称
			SchedulerCategory:   item.SchedulerCategory,          //类别 1:已接入任务, 2:自定义任务
			SchedulerTaskNumber: item.SchedulerTaskNumber,        //已接入任务号
			SchedulerType:       item.SchedulerType,              //类型 1:Http任务,2:Webservices任务
			ErrorOrder:          item.ErrorOrder,                 //失败重新发送次数1-10次 不可超过10次
			FailIntervalTime:    item.FailIntervalTime,           //失败间隔时间按秒
			FailOrder:           item.FailOrder,                  //失败次数
			SchedulerData:       item.SchedulerData,              //内容
			RequestData:         item.RequestData,                //请求内容
		})
	}

	return &tpmtclient.TpmtScheduledTasksFailureRecordListResp{
		Total: count,
		List:  list,
	}, nil
}
