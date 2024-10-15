package appLog

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"
	"tpmt-zt/service/archive/rpc/archiveclient"
)

type ScheduledTasksLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScheduledTasksLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScheduledTasksLogLogic {
	return &ScheduledTasksLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScheduledTasksLogLogic) ScheduledTasksLog(req *types.ScheduledTasksLogReqest) (resp *types.Response, err error) {
	if req.CreatedEndTime <= req.CreatedStartTime && req.CreatedEndTime != 0 && req.CreatedStartTime != 0 {
		return nil, common.NewDefaultError("结束时间不能小于等于开始时间")
	}

	res, err := l.svcCtx.ArchiveRpc.ScheduledTasksLogFindList(l.ctx, &archiveclient.ScheduledTasksLogFindListReq{
		Current:          req.Current,
		PageSize:         req.PageSize,
		StartTime:        req.CreatedStartTime,
		EndTime:          req.CreatedEndTime,
		ScheduledTasksId: req.ScheduledTasksId,
		IsRequest:        req.IsRequest,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result ScheduledTasksLogFindListResp

	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type ScheduledTasksLogFindListResp struct {
	Total int64                    `json:"total"` //总数据量
	List  []*ScheduledTasksLogData `json:"list"`  //数据
}

type ScheduledTasksLogData struct {
	Ts               int64  `json:"ts"`                 // 创建时间
	ScheduledTasksId string `json:"scheduled_tasks_id"` // 自定义任务ID
	IsRequest        int64  `json:"is_request"`         // 状态  1:成功  其他失败
	RequestData      string `json:"request_data"`       // 数据  发送内容
	ResponseData     string `json:"response_data"`      //  数据  返回内容
}
