package logic

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jwtx"
)

/*
*
告警获取注册
*/
func (l *Websocket) registerForAlarm(ctx context.Context, conn *websocket.Conn, tokenData *jwtx.TokenData, device string) {
	global.AlarmUlock.Lock()
	defer global.AlarmUlock.Unlock()

	remoteAddr := (*conn).RemoteAddr().String()
	// 根据会话ip+端口存储 会话
	sConnectionData := &global.SConnectionData{
		Uid:  tokenData.TokenType + "|" + tokenData.Uid,
		Conn: conn,
	}
	global.AlarmSConnection[remoteAddr] = sConnectionData

	// 根据用户ID 存储对应ip和端口
	onlineData := &global.OnlineData{
		Ip:           remoteAddr,
		Device:       device,
		RegisterTime: time.Now().UnixMilli(),
	}
	global.AlarmOnlineList[tokenData.TokenType+"|"+tokenData.Uid] = append(global.AlarmOnlineList[tokenData.TokenType+"|"+tokenData.Uid], onlineData)
	// 存储redis信息  key为 userId +ip和端口  内容为   ip:端口
	redisKey := fmt.Sprintf("%s%s:%s", global.WebSocketUserPrefix, tokenData.TokenType+"|"+tokenData.Uid, remoteAddr)

	l.svcCtx.Redis.SetWithExpireCtx(ctx, redisKey, onlineData, l.svcCtx.OnlineTime)
}

/*
*
监测点获取注册
*/
func (l *Websocket) registerForMonitor(ctx context.Context, conn *websocket.Conn, tokenData *jwtx.TokenData, device string) {
	global.MonitorUlock.Lock()
	defer global.MonitorUlock.Unlock()

	remoteAddr := (*conn).RemoteAddr().String()
	// 根据会话ip+端口存储 会话
	sConnectionData := &global.SConnectionData{
		Uid:  tokenData.TokenType + "|" + tokenData.Uid,
		Conn: conn,
	}
	global.MonitorSConnection[remoteAddr] = sConnectionData

	// 根据用户ID 存储对应ip和端口
	onlineData := &global.OnlineData{
		Ip:           remoteAddr,
		Device:       device,
		RegisterTime: time.Now().UnixMilli(),
	}
	global.MonitorOnlineList[tokenData.TokenType+"|"+tokenData.Uid] = append(global.MonitorOnlineList[tokenData.TokenType+"|"+tokenData.Uid], onlineData)
	// 存储redis信息  key为 userId +ip和端口  内容为   ip:端口
	redisKey := fmt.Sprintf("%s%s:%s", global.WebSocketUserPrefix, tokenData.TokenType+"|"+tokenData.Uid, remoteAddr)

	l.svcCtx.Redis.SetWithExpireCtx(ctx, redisKey, onlineData, l.svcCtx.OnlineTime)
}
