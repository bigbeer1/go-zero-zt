package com

import (
	"context"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/service/tpmtcom/svc"
)

type ComScheduler struct {
	svcCtx *svc.ServiceContext
}

type findGatewayByComPortReq struct {
	ctx       context.Context
	dictValue string
}

func NewComScheduler(svcCtx *svc.ServiceContext) *ComScheduler {
	return &ComScheduler{
		svcCtx: svcCtx,
	}
}

func (l *ComScheduler) Start() {

	go func() {
		for {
			// 创建120秒ctx
			ctx, _ := context.WithTimeout(context.Background(), time.Second*120)

			// 判断是否模拟数据
			if l.svcCtx.Config.Mock {
				l.mockTask(ctx)
			} else {
				l.task(ctx)
			}

			time.Sleep(time.Duration(global.ComCollectionInterval) * time.Second)
		}
	}()

}
