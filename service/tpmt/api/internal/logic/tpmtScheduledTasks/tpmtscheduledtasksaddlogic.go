package tpmtScheduledTasks

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtScheduledTasksAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksAddLogic {
	return &TpmtScheduledTasksAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtScheduledTasksAddLogic) TpmtScheduledTasksAdd(req *types.TpmtScheduledTasksAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtScheduledTasksAdd(l.ctx, &tpmtclient.TpmtScheduledTasksAddReq{
		CreatedName:         tokenData.NickName,      // 创建人
		SchedulerName:       req.SchedulerName,       // 名称
		SchedulerCategory:   req.SchedulerCategory,   // 类别 1:已接入任务, 2:自定义任务
		SchedulerTaskNumber: req.SchedulerTaskNumber, // 已接入任务号
		SchedulerType:       req.SchedulerType,       // 类型 1:Http任务,2:Webservices任务
		IntervalTime:        req.IntervalTime,        // 间隔时间按秒
		ErrorOrder:          req.ErrorOrder,          // 失败重新发送次数1-10次 不可超过10次
		FailIntervalTime:    req.FailIntervalTime,    // 失败间隔时间按秒
		State:               req.State,               // 状态 1:启动  2:暂停
		SchedulerData:       req.SchedulerData,       // 内容
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
