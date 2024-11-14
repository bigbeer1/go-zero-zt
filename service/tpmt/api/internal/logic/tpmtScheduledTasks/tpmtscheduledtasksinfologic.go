package tpmtScheduledTasks

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

type TpmtScheduledTasksInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtScheduledTasksInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksInfoLogic {
	return &TpmtScheduledTasksInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtScheduledTasksInfoLogic) TpmtScheduledTasksInfo(req *types.TpmtScheduledTasksInfoRequest) (resp *types.Response, err error) {
	// 用户登录信息

	res, err := l.svcCtx.TpmtRpc.TpmtScheduledTasksFindOne(l.ctx, &tpmtclient.TpmtScheduledTasksFindOneReq{
		Id: req.Id, // 定时任务ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtScheduledTasksFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtScheduledTasksFindOneResp struct {
	Id                  string `json:"id"`                    // 定时任务ID,
	CreatedAt           int64  `json:"created_at"`            // 创建时间,
	UpdatedAt           int64  `json:"updated_at"`            // 更新时间,
	CreatedName         string `json:"created_name"`          // 创建人,
	UpdatedName         string `json:"updated_name"`          // 更新人,
	SchedulerName       string `json:"scheduler_name"`        // 名称,
	SchedulerCategory   int64  `json:"scheduler_category"`    // 类别 1:已接入任务, 2:自定义任务,
	SchedulerTaskNumber int64  `json:"scheduler_task_number"` // 已接入任务号,
	SchedulerType       int64  `json:"scheduler_type"`        // 类型 1:Http任务,2:Webservices任务,
	IntervalTime        int64  `json:"interval_time"`         // 间隔时间按秒,
	ErrorOrder          int64  `json:"error_order"`           // 失败重新发送次数1-10次 不可超过10次,
	FailIntervalTime    int64  `json:"fail_interval_time"`    // 失败间隔时间按秒,
	State               int64  `json:"state"`                 // 状态 1:启动  2:暂停,
	SchedulerData       string `json:"scheduler_data"`        // 内容,
}
