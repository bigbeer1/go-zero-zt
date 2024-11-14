package tpmtScheduledTasksFailureRecord

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksFailureRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtScheduledTasksFailureRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksFailureRecordListLogic {
	return &TpmtScheduledTasksFailureRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtScheduledTasksFailureRecordListLogic) TpmtScheduledTasksFailureRecordList(req *types.TpmtScheduledTasksFailureRecordListRequest) (resp *types.Response, err error) {
	// 用户登录信息
	all, err := l.svcCtx.TpmtRpc.TpmtScheduledTasksFailureRecordList(l.ctx, &tpmtclient.TpmtScheduledTasksFailureRecordListReq{
		Current:          req.Current,          // 页码
		PageSize:         req.PageSize,         // 页数
		ScheduledTasksId: req.ScheduledTasksId, // 任务ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtScheduledTasksFailureRecordListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtScheduledTasksFailureRecordListResp struct {
	Total int64                                      `json:"total"`
	List  []*TpmtScheduledTasksFailureRecordDataList `json:"list"`
}

type TpmtScheduledTasksFailureRecordDataList struct {
	Id                  string `json:"id"`                    // 失败记录ID,
	ScheduledTasksId    string `json:"scheduled_tasks_id"`    // 任务ID,
	CreatedAt           int64  `json:"created_at"`            // 创建时间,
	UpdatedAt           int64  `json:"updated_at"`            // 更新时间,
	CreatedName         string `json:"created_name"`          // 创建人,
	UpdatedName         string `json:"updated_name"`          // 更新人,
	SchedulerName       string `json:"scheduler_name"`        // 名称,
	SchedulerCategory   int64  `json:"scheduler_category"`    // 类别 1:已接入任务, 2:自定义任务,
	SchedulerTaskNumber int64  `json:"scheduler_task_number"` // 已接入任务号,
	SchedulerType       int64  `json:"scheduler_type"`        // 类型 1:Http任务,2:Webservices任务,
	ErrorOrder          int64  `json:"error_order"`           // 失败重新发送次数1-10次 不可超过10次,
	FailIntervalTime    int64  `json:"fail_interval_time"`    // 失败间隔时间按秒,
	FailOrder           int64  `json:"fail_order"`            // 失败次数
	SchedulerData       string `json:"scheduler_data"`        // 内容,
	RequestData         string `json:"request_data"`          // 请求内容,
}
