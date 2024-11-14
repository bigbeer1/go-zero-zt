package logic

import (
	"context"
	"tpmt-zt/service/mqttsend/internal/mqttx"
	"tpmt-zt/service/mqttsend/internal/svc"
	"tpmt-zt/service/mqttsend/mqttsendclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TwyMqttSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTwyMqttSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TwyMqttSendLogic {
	return &TwyMqttSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给泰无忧发送信息
func (l *TwyMqttSendLogic) TwyMqttSend(in *mqttsendclient.TwyMqttSendReq) (*mqttsendclient.CommonResp, error) {

	mqtt := &mqttx.Mqtt{
		Host:     in.MqttHost,
		User:     in.MqttUser,
		Pass:     in.MqttPass,
		ClientId: in.MqttClientId,
	}

	// 创建mqtt客户端
	mqttClient, err := mqtt.NewMqttManager()
	if err != nil {
		return nil, err
	}

	// 给对应的主体发送mqtt内容
	err = mqttx.Producer(mqttClient, in.SendTopic, in.Data)

	if err != nil {
		return nil, err
	}

	return &mqttsendclient.CommonResp{}, nil
}
