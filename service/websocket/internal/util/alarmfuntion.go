package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
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
func AlarmDelSocket(ctx context.Context, remoteAddr string, redisx cache.Cache) {
	if !common.IsEmpty(remoteAddr) {
		//关闭客户端socket连接
		logx.Infof("socket delete conn [%s]", remoteAddr)
		// 删除
		if v, ok := global.AlarmSConnection[remoteAddr]; ok {
			// 删除OnlineList中在线数据
			var onlineDatas []*global.OnlineData
			for _, item := range global.AlarmOnlineList[v.Uid] {
				if item.Ip != remoteAddr {
					onlineDatas = append(onlineDatas, item)
				}
				redisKey := fmt.Sprintf("%s%s:%s", global.WebSocketUserPrefix, v.Uid, remoteAddr)
				// 删除redis 中值
				redisx.DelCtx(ctx, redisKey)
			}
			if len(onlineDatas) == 0 {
				delete(global.AlarmOnlineList, v.Uid)
			} else {
				global.AlarmOnlineList[v.Uid] = onlineDatas
			}
			// 删除SConnection中在线数据
			delete(global.AlarmSConnection, remoteAddr)
		}
	}
}

/*
*
发送数据
*/
func AlarmSendSocket(ws *websocket.Conn, message string) error {
	defer global.AlarmUlock.Unlock()
	global.AlarmUlock.Lock()
	if _, ok := global.AlarmSConnection[ws.RemoteAddr().String()]; !ok {
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
func AlarmSendSocketBytes(ws *websocket.Conn, data []byte) error {
	defer global.AlarmUlock.Unlock()
	global.AlarmUlock.Lock()
	if _, ok := global.AlarmSConnection[ws.RemoteAddr().String()]; !ok {
		return errors.New("已关闭")
	}
	err := ws.WriteMessage(2, data)
	logx.Infof(fmt.Sprintf("socket send [%s] %v", ws.RemoteAddr().String(), data))
	return err
}
