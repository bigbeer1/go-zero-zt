package util

import (
	"errors"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common"
	"tpmt-zt/common/global"
)

type SendMessage struct {
	Ws      *websocket.Conn
	Message string
}

///**
//数据无效
//*/
//func invalidData(conn *net.Conn, msg string) {
//	remoteAddr := (*conn).RemoteAddr().String()
//	fmt.Println(fmt.Sprintf("socket invalid data [%s] %s", remoteAddr, msg))
//	SendSocket(conn, "error|invalid data")
//}

/*
*
清理连接map
*/
func DelSocket(remoteAddr string) {
	if !common.IsEmpty(remoteAddr) {
		//关闭客户端socket连接
		logx.Infof("socket delete conn [%s]", remoteAddr)
		// 删除
		if v, ok := global.TpmtSConnection[remoteAddr]; ok {
			if v.GateWayId != "" {
				// 删除Gateway中订阅的用户
				var ips []string
				for _, item := range global.SubscribeGateWay[v.GateWayId] {
					if item != remoteAddr {
						ips = append(ips, item)
					}
				}
				if len(ips) == 0 {
					delete(global.SubscribeGateWay, v.GateWayId)
				} else {
					global.SubscribeGateWay[v.GateWayId] = ips
				}
			}
			// 删除SConnection中在线数据
			delete(global.TpmtSConnection, remoteAddr)
		}
	}
}

/*
*
删除订阅的GateWay
*/
func DelGateWaySubscribe(remoteAddr string) {
	if !common.IsEmpty(remoteAddr) {
		//关闭客户端GateWay订阅
		logx.Infof("gateWay subscribe delete conn [%s]", remoteAddr)
		// 删除
		if v, ok := global.TpmtSConnection[remoteAddr]; ok {
			if v.GateWayId != "" {
				// 删除Gateway中订阅的用户
				var ips []string
				for _, item := range global.SubscribeGateWay[v.GateWayId] {
					if item != remoteAddr {
						ips = append(ips, item)
					}
				}
				if len(ips) == 0 {
					delete(global.SubscribeGateWay, v.GateWayId)
				} else {
					global.SubscribeGateWay[v.GateWayId] = ips
				}
			}
		}
	}
}

/*
*
发送数据
*/
func SendSocket(sendMessage SendMessage) error {
	defer global.Ulock.Unlock()
	global.Ulock.Lock()
	if _, ok := global.TpmtSConnection[sendMessage.Ws.RemoteAddr().String()]; !ok {
		return errors.New("已关闭")
	}
	err := sendMessage.Ws.WriteMessage(1, []byte(sendMessage.Message))
	logx.Infof(fmt.Sprintf("socket send [%s] %v", sendMessage.Ws.RemoteAddr().String(), sendMessage.Message))
	return err
}

/*
*
发送数据
*/
func SendSocketNoLock(sendMessage SendMessage) error {
	if _, ok := global.TpmtSConnection[sendMessage.Ws.RemoteAddr().String()]; !ok {
		return errors.New("已关闭")
	}
	err := sendMessage.Ws.WriteMessage(1, []byte(sendMessage.Message))
	logx.Infof(fmt.Sprintf("socket send [%s] %v", sendMessage.Ws.RemoteAddr().String(), sendMessage.Message))
	return err
}
