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

type ScheduledTasksFailureRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScheduledTasksFailureRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScheduledTasksFailureRecordLogic {
	return &ScheduledTasksFailureRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScheduledTasksFailureRecordLogic) ScheduledTasksFailureRecord(req *types.ScheduledTasksLogReqest) (resp *types.Response, err error) {
	if req.CreatedEndTime <= req.CreatedStartTime && req.CreatedEndTime != 0 && req.CreatedStartTime != 0 {
		return nil, common.NewDefaultError("结束时间不能小于等于开始时间")
	}

	res, err := l.svcCtx.ArchiveRpc.ScheduledTasksFailureRecordLogFindList(l.ctx, &archiveclient.ScheduledTasksLogFindListReq{
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
