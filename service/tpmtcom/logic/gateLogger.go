package logic

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/tpmtcom/util"
)

func (l *Websocket) GateLogger(conn *websocket.Conn, gateWayId string, tokenData *jwtx.TokenData) {

	// 根据gateWayId 查询网关ID
	// 创建30秒ctx
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	_, err := l.svcCtx.TpmtGatewayModel.FindOne(ctx, gateWayId)

	global.Ulock.Lock()
	defer global.Ulock.Unlock()

	remoteAddr := (*conn).RemoteAddr().String()
	util.DelGateWaySubscribe(remoteAddr)

	if err != nil {
		if err == sqlc.ErrNotFound {
			l.SendSocketNoLockReq(conn, fmt.Sprintf("error|TpmtGateway没有该ID:%v", gateWayId))
			return
		}
		l.SendSocketNoLockReq(conn, fmt.Sprintf("error|查询gateway错误:%s", err.Error()))
		return
	}

	// 根据会话ip+端口存储 会话
	tpmtSConnectionData := &global.TpmtSConnectionData{
		Uid:       tokenData.Uid,
		Conn:      conn,
		GateWayId: gateWayId,
	}
	global.TpmtSConnection[remoteAddr] = tpmtSConnectionData

	global.SubscribeGateWay[gateWayId] = append(global.SubscribeGateWay[gateWayId], remoteAddr)

	return

}

func (l *Websocket) GateLoggerClose(conn *websocket.Conn) {

	global.Ulock.Lock()
	defer global.Ulock.Unlock()

	remoteAddr := (*conn).RemoteAddr().String()
	util.DelGateWaySubscribe(remoteAddr)

	l.SendSocketNoLockReq(conn, fmt.Sprintf("关闭成功"))

	return

}
