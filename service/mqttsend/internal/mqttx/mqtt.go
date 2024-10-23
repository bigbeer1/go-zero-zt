package mqttx

import (
	"fmt"
	mq "github.com/eclipse/paho.mqtt.golang"
	"time"
	"tpmt-zt/common/global"
)

type Mqtt struct {
	Host     string
	User     string
	Pass     string
	ClientId string
}

func (m Mqtt) NewMqttManager() (mq.Client, error) {

	// 判断是否已经连接
	// 已经连接则返回
	// 未连接则创建连接
	if mqttClient, ok := global.MqttConnection[m.Host]; ok {
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {

		} else {
			return mqttClient, nil
		}

	}

	opts := mq.NewClientOptions()
	opts.AddBroker(m.Host).
		SetUsername(m.User).
		SetPassword(m.Pass).
		SetClientID(m.ClientId).
		SetAutoReconnect(false).
		SetConnectRetry(false).
		SetOrderMatters(false).
		SetKeepAlive(60*time.Second).
		SetPingTimeout(30*time.Second).
		SetWill("server_will", "lose_connect", 2, false)

	mqttClient := mq.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("初始化Mqtt客户端失败" + token.Error().Error())
	}
	mqttClient.Publish("server_will", 2, true, "server_start")

	// 将连接添加到全局变量中
	global.MqttConnection[m.Host] = mqttClient

	return mqttClient, nil
}

func Producer(client mq.Client, topic string, message []byte) error {
	if token := client.Publish(topic, 0, true, message); token.Wait() && token.Error() != nil {
		fmt.Errorf("发布消息失败" + token.Error().Error())
	}
	return nil
}
