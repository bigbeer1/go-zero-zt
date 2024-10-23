package global

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fasthttp/websocket"
	"sync"
	"time"
)

// 遥性解译
type YxInterpreting struct {
	Key   string `json:"key"`   // 键值
	Value string `json:"value"` // 对应显示
}

var (
	AlarmSConnection      = make(map[string]*SConnectionData)  // socket列表 key为ip value为 会话conn
	AlarmOnlineList       = make(map[string][]*OnlineData)     // key为uid  value 为[]ip
	MonitorSConnection    = make(map[string]*SConnectionData)  // socket列表 key为ip value为 会话conn
	MonitorOnlineList     = make(map[string][]*OnlineData)     // key为uid  value 为[]ip
	ShangHaiTime, _       = time.LoadLocation("Asia/Shanghai") //上海
	Ulock                 sync.Mutex
	AlarmUlock            sync.Mutex
	MonitorUlock          sync.Mutex
	SysProgrammeId              = "" // 方案ID
	ComCollectionInterval int64 = 3  // 从串口获取数据时间延时3秒
	ComMemoryCycleTime    int64 = 10 // 数据存储时序数据库时间
	AlarmAddTime          int64 = 24 // 告警阻塞小时
)

// tpmt-com websocket
var (
	TpmtSConnection  = make(map[string]*TpmtSConnectionData) // socket列表 key为ip value为 会话conn
	SubscribeGateWay = make(map[string][]string)             // key为gatewayId  value 为[]ip
)

// tpmt-send
var (
	MqttConnection = make(map[string]mqtt.Client)
)

var WebSocketUserPrefix = "websocketUser:"

type SConnectionData struct {
	Uid  string          `json:"uid"`  // uid
	Conn *websocket.Conn `json:"conn"` // IP
}

type TpmtSConnectionData struct {
	Uid       string          `json:"uid"`         // uid
	Conn      *websocket.Conn `json:"conn"`        // IP
	GateWayId string          `json:"gate_way_ID"` // IP
}

type OnlineData struct {
	Ip           string `json:"ip"`            // IP
	Device       string `json:"device"`        // 设备
	RegisterTime int64  `json:"register_time"` // 注册时间
}

type SocketMonitor struct {
	Id            int64  `json:"id"`
	ResultValue   string `json:"result_value"`   // 监测值
	UpdateTime    int64  `json:"update_time"`    // 更新时间
	Level         int64  `json:"level"`          // 告警等级  0正常
	RuleType      int64  `json:"rule_type"`      // 告警状态  1
	PointCategory int64  `json:"point_category"` // 类别：1:遥信/2:遥测/3:遥脉
	Unit          string `json:"unit"`           // 单位
}
