package com

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/service/asynq/jobtype"
)

// 数据流转到redis
func (l *ComScheduler) DataFlowRedis(data *global.MonitorCacheUpData) {
	payloadData, err := jsonx.Marshal(data)
	if err != nil {
		logx.Errorf("报警数据转换失败", payloadData)
	}
	// 创建30秒Ctx
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	l.svcCtx.AsynqClient.EnqueueContext(ctx, asynq.NewTask(jobtype.DataRedisSet, payloadData))

	defer cancel()
}

// 数据流转到wocker
func (l *ComScheduler) DataFlowRedisReq(data *global.MonitorCacheUpData) {
	req := global.MonitorCacheUpData{
		Id:                    data.Id,
		Name:                  data.Name,
		PointCategory:         data.PointCategory,
		IsDisplacementWarning: data.IsDisplacementWarning,
		Ts:                    data.Ts,
		Data:                  data.Data,
	}
	_ = l.svcCtx.DataAntsPool.Invoke(req)
}

// 数据流转
func (l *ComScheduler) GateWayOnlineDataFlow(data *global.GateWayOnlineUpData) {
	payloadData, err := jsonx.Marshal(data)
	if err != nil {
		logx.Errorf("报警数据转换失败", payloadData)
	}
	// 创建30秒Ctx
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	l.svcCtx.AsynqClient.EnqueueContext(ctx, asynq.NewTask(jobtype.GatewayStateOnline, payloadData))

	defer cancel()
}

// 数据流转
func (l *ComScheduler) GateWayOnlineDataFlowReq(data *global.GateWayOnlineUpData) {
	req := global.GateWayOnlineUpData{
		Id:           data.Id,
		Name:         data.Name,
		OnlineStatus: data.OnlineStatus,
	}
	_ = l.svcCtx.GateWayAntsPool.Invoke(req)
}
