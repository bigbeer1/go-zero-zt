package logic

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/websocket/internal/util"
)

type SendMessage struct {
	SUserId string `json:"s_user_id"` // 发送人
	RUserId string `json:"r_user_id"` // 接收人
	Content string `json:"content"`   // 内容
	Type    int    `json:"type"`      // 类型  1文字聊天
}

func (l *Websocket) Chat(ctx context.Context, conn *websocket.Conn, dataString string, tokenData *jwtx.TokenData) {
	var in *SendMessage

	// 解析json
	err := jsonx.Str2Struct(dataString, &in)
	if err != nil {
		err = util.MonitorSendSocket(conn, fmt.Sprintf("error|json数据格式错误|%s|", dataString))
		if err != nil {
			logx.Errorf(err.Error())
		}
		return
	}

	in.SUserId = tokenData.Uid

	rUserId := fmt.Sprintf("%v|%v", common.UserTokenType, in.RUserId)
	// 找对应人的conn
	// 删除
	if v, ok := global.MonitorOnlineList[rUserId]; ok {
		for _, item := range v {
			if data, okA := global.MonitorSConnection[item.Ip]; okA {
				if data.Uid == rUserId {
					// 给用户回复信息
					content, _ := jsonx.ToJSONStr(in)
					util.MonitorSendSocket(data.Conn, fmt.Sprintf("sendMessage|%s|", content))
				}

			}
		}
	} else {
		// 人不在线
		util.MonitorSendSocket(conn, fmt.Sprintf("sendMessageError|%s|", "人不在线"))
	}

}
