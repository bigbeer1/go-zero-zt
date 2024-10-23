package logic

import (
	"github.com/fasthttp/websocket"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jwtx"
)

/*
*
注册
*/
func (l *Websocket) registerUser(conn *websocket.Conn, tokenData *jwtx.TokenData) {
	global.Ulock.Lock()
	defer global.Ulock.Unlock()

	remoteAddr := (*conn).RemoteAddr().String()
	// 根据会话ip+端口存储 会话
	var tpmtSConnectionData = &global.TpmtSConnectionData{
		Uid:       tokenData.Uid,
		Conn:      conn,
		GateWayId: "",
	}
	global.TpmtSConnection[remoteAddr] = tpmtSConnectionData

}
