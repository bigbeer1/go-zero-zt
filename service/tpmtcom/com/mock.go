package com

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"strconv"
	"sync"
	"time"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/global"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmtcom/util"
)

func (l *ComScheduler) mockTask(ctx context.Context) {

	// 根据字典查询查询com_port 有多少类型
	whereBuilderA := l.svcCtx.SysDictModel.RowBuilder()

	whereBuilderA = whereBuilderA.Where("deleted_at is null")
	whereBuilderA = whereBuilderA.OrderBy("created_at DESC, id DESC")

	// 字典类型
	whereBuilderA = whereBuilderA.Where(squirrel.Eq{
		"dict_type ": "com_port",
	})

	dictAll, err := l.svcCtx.SysDictModel.FindList(ctx, whereBuilderA, 1, 100)
	if err != nil {
		return
	}

	var wg sync.WaitGroup

	// 注册协程池方法
	p, _ := ants.NewPoolWithFunc(300, func(req interface{}) {
		data, _ := req.(findGatewayByComPortReq)
		l.mockFindGatewayByComPort(data.ctx, data.dictValue)
		wg.Done()
	})

	defer p.Release()

	// 循环com_port 查询网关信息 并启动协程池
	for _, dict := range dictAll {
		wg.Add(1)
		req := findGatewayByComPortReq{
			ctx:       ctx,
			dictValue: dict.DictValue,
		}
		_ = p.Invoke(req)
	}
	wg.Wait()

}

// 根据字典com_port 查询网关
func (l *ComScheduler) mockFindGatewayByComPort(ctx context.Context, dictValue string) {

	whereBuilder := l.svcCtx.TpmtGatewayModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"com_port": dictValue,
	})

	all, err := l.svcCtx.TpmtGatewayModel.FindList(ctx, whereBuilder, 1, 5000)
	if err != nil {
		return
	}

	// 循环网关去找监测点
	for _, gateway := range all {
		l.mockConnectRTUStore(ctx, gateway)
	}

}

/*
查询网关下监测点
*/
func (l *ComScheduler) mockConnectRTUStore(ctx context.Context, gateway *model.TpmtGateway) {

	whereBuilder := l.svcCtx.TpmtMonitorPointModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	// 采集点ID
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tpmt_gateway_id ": gateway.Id,
	})

	monitorPoints, err := l.svcCtx.TpmtMonitorPointModel.FindList(ctx, whereBuilder, 1, 50000)
	if err != nil {
		l.GateWayLoggerHookSend(gateway.Id, "error|mock查询数据库对应监测点失败"+err.Error())
		return
	}

	// 传到websocket里的信息
	l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("mock对应监测点数量为:%v", len(monitorPoints)))

	// 存储网关在线状态标识  默认网关设备连接成功  1
	gateWayData := &global.GateWayOnlineUpData{
		Id:           gateway.Id,
		Name:         gateway.GatewayName,
		OnlineStatus: "1",
	}

	// 流转网关状态
	l.GateWayOnlineDataFlowReq(gateWayData)

	l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("mock连接串口成功:%v", gateway.ComPort))

	for _, monitorPoint := range monitorPoints {

		// 发送指令
		l.mockGateWayLoggerHookSendWithData(gateway.Id, gateway.AddressCode, monitorPoint.Id, monitorPoint.Name, monitorPoint.RegisterAddress, monitorPoint.DataBits, monitorPoint.PointCollectorInstruction)

		if monitorPoint.PointCollectorInstruction != 1 && monitorPoint.PointCollectorInstruction != 2 && monitorPoint.PointCollectorInstruction != 3 {
			l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,不存在除了 02 03 04 以外的指令可以发送数据存在问题当前PointCollectorInstruction为:%v", monitorPoint.Id, monitorPoint.Name, monitorPoint.PointCollectorInstruction))
			return
		}

		// 解析返回指令
		var resultData *global.MonitorCacheUpData

		if monitorPoint.PointCategory == 1 {
			// 模拟数据存储
			resultData = &global.MonitorCacheUpData{
				Id:                    datax.ToString(monitorPoint.Id),
				Name:                  monitorPoint.Name,
				PointCategory:         monitorPoint.PointCategory,
				IsDisplacementWarning: monitorPoint.IsDisplacementWarning,
				Ts:                    time.Now(),
				AssetId:               monitorPoint.AssetId,
				Data:                  datax.ToString(RandomDecimal("0", 1)),
			}
		} else {
			// 模拟数据存储
			resultData = &global.MonitorCacheUpData{
				Id:                    datax.ToString(monitorPoint.Id),
				Name:                  monitorPoint.Name,
				PointCategory:         monitorPoint.PointCategory,
				IsDisplacementWarning: monitorPoint.IsDisplacementWarning,
				Ts:                    time.Now(),
				AssetId:               monitorPoint.AssetId,
				Data:                  datax.ToString(RandomDecimal("0", 300)),
			}
		}

		// 数据流转
		l.DataFlowRedisReq(resultData)

		l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("监测点ID:%v,编号:%s,名称:%s,mock解析结果:数值:%v,", monitorPoint.Id, monitorPoint.SerialNumber, monitorPoint.Name, resultData.Data))

	}

}
func RandomDecimal(bit string, multiple float64) float64 {
	bit = "%." + bit + "f"
	data := rand.Float64() * multiple
	data, _ = strconv.ParseFloat(fmt.Sprintf(bit, data), 64)
	return data
}

func (l *ComScheduler) mockGateWayLoggerHookSendWithData(gateWayId string, gateWayAddressCode, monitorPointId int64, monitorPointName, address string, dataBits int64, collectorInstruction int64) {

	if all, ok := global.SubscribeGateWay[gateWayId]; ok {

		addressCompletion := Hex16Completion(address, 4)
		dataBitsCompletion := fmt.Sprintf("00,0%v", dataBits)

		var message string
		switch collectorInstruction {
		case 1:
			message = fmt.Sprintf("监测点ID:%v,监测点名称:%s,发送指令:0%v,02,%s,%s", monitorPointId, monitorPointName, gateWayAddressCode, addressCompletion, dataBitsCompletion)
		case 2:
			message = fmt.Sprintf("监测点ID:%v,监测点名称:%s,发送指令:0%v,03,%s,%s", monitorPointId, monitorPointName, gateWayAddressCode, addressCompletion, dataBitsCompletion)
		case 3:
			message = fmt.Sprintf("监测点ID:%v,监测点名称:%s,发送指令:0%v,04,%s,%s", monitorPointId, monitorPointName, gateWayAddressCode, addressCompletion, dataBitsCompletion)
		default:
			message = fmt.Sprintf("监测点ID:%v,监测点名称:%s,指令存在错误collectorInstruction不是123", monitorPointId, monitorPointName)

		}

		for _, v := range all {
			if connData, oka := global.TpmtSConnection[v]; oka {
				req := util.SendMessage{
					Ws:      connData.Conn,
					Message: message,
				}
				_ = l.svcCtx.SendSocketAntsPool.Invoke(req)
			}
		}

	}

}
