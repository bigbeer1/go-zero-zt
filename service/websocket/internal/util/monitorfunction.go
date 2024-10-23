package util

import (
	"errors"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common"
	"tpmt-zt/common/global"
)

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
func MonitorDelSocket(remoteAddr string) {
	if !common.IsEmpty(remoteAddr) {
		//关闭客户端socket连接
		logx.Infof("socket delete conn [%s]", remoteAddr)
		// 删除
		if v, ok := global.MonitorSConnection[remoteAddr]; ok {
			// 删除OnlineList中在线数据
			var onlineDatas []*global.OnlineData
			for _, item := range global.MonitorOnlineList[v.Uid] {
				if item.Ip != remoteAddr {
					onlineDatas = append(onlineDatas, item)
				}
			}
			if len(onlineDatas) == 0 {
				delete(global.MonitorOnlineList, v.Uid)
			} else {
				global.MonitorOnlineList[v.Uid] = onlineDatas
			}
			// 删除SConnection中在线数据
			delete(global.MonitorSConnection, remoteAddr)
		}
	}
}

/*
*
发送数据
*/
func MonitorSendSocket(ws *websocket.Conn, message string) error {
	defer global.MonitorUlock.Unlock()
	global.MonitorUlock.Lock()
	if _, ok := global.MonitorSConnection[ws.RemoteAddr().String()]; !ok {
		return errors.New("已关闭")
	}
	err := ws.WriteMessage(1, []byte(message))
	logx.Infof(fmt.Sprintf("socket send [%s] %v", ws.RemoteAddr().String(), message))
	return err
}

/*
*
发送数据
*/
func MonitorSendSocketBytes(ws *websocket.Conn, data []byte) error {
	defer global.MonitorUlock.Unlock()
	global.MonitorUlock.Lock()
	if _, ok := global.MonitorSConnection[ws.RemoteAddr().String()]; !ok {
		return errors.New("已关闭")
	}
	err := ws.WriteMessage(2, data)
	logx.Infof(fmt.Sprintf("socket send [%s] %v", ws.RemoteAddr().String(), data))
	return err
}
