package tpmtScheduledTasks

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtScheduledTasksDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksDelLogic {
	return &TpmtScheduledTasksDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtScheduledTasksDelLogic) TpmtScheduledTasksDel(req *types.TpmtScheduledTasksDelRequest) (resp *types.Response, err error) {
	// 用户登录信息

	_, err = l.svcCtx.TpmtRpc.TpmtScheduledTasksDelete(l.ctx, &tpmtclient.TpmtScheduledTasksDeleteReq{
		Id: req.Id, // 定时任务ID
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
