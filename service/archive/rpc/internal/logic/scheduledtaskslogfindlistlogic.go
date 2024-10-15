package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/archive/model"
	"tpmt-zt/service/archive/rpc/archiveclient"
	"tpmt-zt/service/archive/rpc/internal/svc"
)

type ScheduledTasksLogFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScheduledTasksLogFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScheduledTasksLogFindListLogic {
	return &ScheduledTasksLogFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自定义任务日志
func (l *ScheduledTasksLogFindListLogic) ScheduledTasksLogFindList(in *archiveclient.ScheduledTasksLogFindListReq) (*archiveclient.ScheduledTasksLogFindListResp, error) {
	scheduledTasksLog := model.ScheduledTasksLog{
		ScheduledTasksId: in.ScheduledTasksId,
		IsRequest:        in.IsRequest,
	}

	// 讲uuid中- 全部替换为_  确保插入和查询成功
	id := strings.Replace(in.ScheduledTasksId, "-", "_", -10)

	tdDb := &model.TdDb{
		DbName:    "scheduled_tasks_log.d1" + id,
		TableName: "",
	}

	all, err := scheduledTasksLog.FindList(l.ctx, l.svcCtx.Taos, tdDb, in.Current, in.PageSize, in.StartTime, in.EndTime)
	if err != nil {
		if err.Error() != tdenginex.ErrNotFoundTable {
			return nil, err
		}
	}

	total := scheduledTasksLog.Count(l.ctx, l.svcCtx.Taos, tdDb, in.StartTime, in.EndTime)

	var datas []*archiveclient.ScheduledTasksLogData
	for _, item := range all {
		datas = append(datas, &archiveclient.ScheduledTasksLogData{
			Ts:               item.Ts.UnixMilli(),
			ScheduledTasksId: item.ScheduledTasksId,
			IsRequest:        item.IsRequest,
			RequestData:      item.RequestData,
			ResponseData:     item.ResponseData,
			TenantId:         item.TenantId,
		})

	}

	return &archiveclient.ScheduledTasksLogFindListResp{
		Total: total,
		List:  datas,
	}, nil
}
