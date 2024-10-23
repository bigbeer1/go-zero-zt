package com

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/goburrow/modbus"
	"github.com/panjf2000/ants/v2"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/global"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmtcom/util"
)

func (l *ComScheduler) task(ctx context.Context) {

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

	p, _ := ants.NewPoolWithFunc(300, func(req interface{}) {
		data, _ := req.(findGatewayByComPortReq)
		l.findGatewayByComPort(data.ctx, data.dictValue)
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
func (l *ComScheduler) findGatewayByComPort(ctx context.Context, dictValue string) {

	whereBuilder := l.svcCtx.TpmtGatewayModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"sys_programme_id ": global.SysProgrammeId,
	})
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"com_port": dictValue,
	})

	all, err := l.svcCtx.TpmtGatewayModel.FindList(ctx, whereBuilder, 1, 5000)
	if err != nil {
		return
	}

	for _, gateway := range all {
		l.connectRTUStore(ctx, gateway)
	}

}

func (l *ComScheduler) connectRTUStore(ctx context.Context, gateway *model.TpmtGateway) {

	whereBuilder := l.svcCtx.TpmtMonitorPointModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	// 采集点ID
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tpmt_gateway_id ": gateway.Id,
	})

	monitorPoints, err := l.svcCtx.TpmtMonitorPointModel.FindList(ctx, whereBuilder, 1, 50000)
	if err != nil {
		l.GateWayLoggerHookSend(gateway.Id, "error|查询数据库对应监测点失败"+err.Error())
		return
	}

	l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("对应监测点数量为:%v", len(monitorPoints)))

	handler := modbus.NewRTUClientHandler(gateway.ComPort)
	handler.BaudRate = int(gateway.BaudRate)
	handler.DataBits = int(gateway.DataBits)
	handler.Parity = gateway.Parity
	handler.StopBits = int(gateway.StopBits)
	handler.SlaveId = byte(gateway.AddressCode)
	handler.Timeout = time.Duration(l.svcCtx.Config.RtuTimeout) * time.Millisecond

	err = handler.Connect()

	var gateWayData *global.GateWayOnlineUpData
	if err != nil {
		// 存储网关离线状态标识
		gateWayData = &global.GateWayOnlineUpData{
			Id:           gateway.Id,
			Name:         gateway.GatewayName,
			OnlineStatus: "0",
		}

		// 流转网关状态
		l.GateWayOnlineDataFlowReq(gateWayData)

		l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|连接串口失败:%v", gateway.ComPort))
		return
	}

	// 存储网关在线状态标识
	gateWayData = &global.GateWayOnlineUpData{
		Id:           gateway.Id,
		Name:         gateway.GatewayName,
		OnlineStatus: "1",
	}

	// 流转网关状态
	l.GateWayOnlineDataFlowReq(gateWayData)

	l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("连接串口成功:%v", gateway.ComPort))
	defer handler.Close()

	client := modbus.NewClient(handler)

	for _, monitorPoint := range monitorPoints {

		address, err := strconv.ParseUint(monitorPoint.RegisterAddress, 0, 16) // 进行转换并指定基数为0表示自动判断，位数限制为16位
		if err != nil {
			l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|转换RegisterAddress错误:%v", err.Error()))
			return
		}

		value := uint16(monitorPoint.DataBits)

		// 发送指令
		l.GateWayLoggerHookSendWithData(gateway.Id, gateway.AddressCode, monitorPoint.Id, monitorPoint.Name, monitorPoint.RegisterAddress, monitorPoint.DataBits, monitorPoint.PointCollectorInstruction)

		var results []byte
		switch monitorPoint.PointCollectorInstruction {
		case 1:
			results, err = client.ReadDiscreteInputs(uint16(address), value)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,发送失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
		case 2:
			results, err = client.ReadHoldingRegisters(uint16(address), value)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,发送失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
		case 3:
			results, err = client.ReadInputRegisters(uint16(address), value)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,发送失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
		default:
			l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,不存在除了 02 03 04 以外的指令可以发送数据存在问题当前PointCollectorInstruction为:%v",
				monitorPoint.Id, monitorPoint.Name, monitorPoint.PointCollectorInstruction))
			return
		}

		hexStr := bytesToHex(results)

		var resultDataString string
		switch monitorPoint.PointAnalysisRule {
		case 1:
			n, err := strconv.ParseUint(hexStr, 16, 32)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,数据解析失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
			resultDataString = datax.ToString(n)
		case 2:
			n, err := strconv.ParseUint(hexStr, 16, 32)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,数据解析失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
			resultDataString = datax.ToString(math.Float32frombits(uint32(n)))
		case 3:
			n, err := strconv.ParseUint(hexStr, 16, 32)
			if err != nil {
				l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,数据解析失败:%s", monitorPoint.Id, monitorPoint.Name, err.Error()))
				return
			}
			resultDataString = datax.ToString(math.Float64frombits(n))
		default:
			l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("error|监测点ID:%v,监测点名称:%s,数据解析失败吗数据解析格式不是1,2", monitorPoint.Id, monitorPoint.Name))
			return
		}

		// 返回内容
		l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("监测点ID:%v,监测点名称:%s,返回指令:%v", monitorPoint.Id, monitorPoint.Name, results))

		// 数据存储
		var resultData *global.MonitorCacheUpData
		resultData = &global.MonitorCacheUpData{
			Id:                    datax.ToString(monitorPoint.Id),
			Name:                  monitorPoint.Name,
			PointCategory:         monitorPoint.PointCategory,
			IsDisplacementWarning: monitorPoint.IsDisplacementWarning,
			Ts:                    time.Now(),
			AssetId:               monitorPoint.AssetId,
			Data:                  resultDataString,
		}

		// 数据流转
		l.DataFlowRedisReq(resultData)

		l.GateWayLoggerHookSend(gateway.Id, fmt.Sprintf("监测点ID:%v,编号:%s,名称:%s,mock解析结果:数值:%v,", monitorPoint.Id, monitorPoint.SerialNumber, monitorPoint.Name, resultData.Data))

	}

}

func Hex16Completion(addressA string, completionLen int) string {

	addressA = strings.Replace(addressA, "0x", "", -1)

	if len(addressA) > 2 {

		addressAs := strings.Split(addressA, "")
		number := completionLen - len(addressAs)
		addressA = ""
		for i := 0; i < number; i++ {
			addressA = "0" + addressA
		}

		for i := 0; i < len(addressAs); i++ {
			if len(addressA)%5 == 2 {
				addressA = addressA + ","
			}
			addressA = addressA + addressAs[i]
		}

	} else {
		number := completionLen - len(addressA)
		for i := 0; i < number; i++ {
			if len(addressA)%5 == 2 {
				addressA = "," + addressA
			}
			addressA = "0" + addressA
		}
	}

	return addressA

}

func (l *ComScheduler) GateWayLoggerHookSendWithData(gateWayId string, gateWayAddressCode, monitorPointId int64, monitorPointName, address string, dataBits int64, collectorInstruction int64) {

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

func (l *ComScheduler) GateWayLoggerHookSend(gateWayId string, message string) {

	if all, ok := global.SubscribeGateWay[gateWayId]; ok {
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

func bytesToHex(byteArray []byte) string {
	hexStr := ""
	for _, b := range byteArray {
		hexStr += fmt.Sprintf("%02x", b)
	}
	return hexStr
}
