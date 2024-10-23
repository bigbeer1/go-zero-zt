package logic

import (
	"context"
	"fmt"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/service/websocket/internal/svc"
	"tpmt-zt/service/websocket/internal/util"
	"tpmt-zt/service/websocket/websocketclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmMessageLogic {
	return &AlarmMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送告警消息
func (l *AlarmMessageLogic) AlarmMessage(in *websocketclient.AlarmMessageReq) (*websocketclient.CommonResp, error) {

	var data = &AlarmMessage{
		Ts:           in.Ts,
		Id:           in.Id,
		Mid:          in.Mid,
		Name:         in.Name,
		AlarmType:    in.AlarmType,
		AlarmGrade:   in.AlarmGrade,
		AlarmContent: in.AlarmContent,
		AssetCode:    in.AssetCode,
	}

	reData, _ := jsonx.ToJSONStr(data)

	// 循环所有订阅告警信息的websocket
	for _, item := range global.AlarmSConnection {
		util.AlarmSendSocket(item.Conn, fmt.Sprintf("alarm|%s|", reData))
	}

	return &websocketclient.CommonResp{}, nil
}

type AlarmMessage struct {
	Ts           int64  `json:"ts"`            // 告警事件
	Id           string `json:"id"`            // 告警ID
	Mid          string `json:"mid"`           // 告警设备ID
	Name         string `json:"name"`          // 告警设备名称
	AlarmType    int64  `json:"alarm_type"`    // 告警类型
	AlarmGrade   int64  `json:"alarm_grade"`   // 告警等级
	AlarmContent string `json:"alarm_content"` // 告警内容
	AssetCode    string `json:"asset_code"`    // 柜号
}
