package logic

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/golang-jwt/jwt/v4"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"strings"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/tpmtcom/util"
)

var upgrader = websocket.FastHTTPUpgrader{
	CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
		return true
	},
}

func (l *Websocket) EchoView(ctx *fasthttp.RequestCtx) {
	// 获取token
	tokenByte := ctx.Request.Header.Peek("Sec-WebSocket-Protocol")
	// 返回必须同样返回 否则websocket 错误
	ctx.Response.Header.Set("Sec-WebSocket-Protocol", string(tokenByte))

	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		// ip和端口
		remoteAddr := ws.RemoteAddr().String()
		// 设备

		defer func() {
			global.Ulock.Lock()
			util.DelSocket(remoteAddr)
			ws.Close()
			global.Ulock.Unlock()
		}()

		if len(tokenByte) == 0 {
			ws.WriteMessage(1, []byte("error|没有获取到令牌|"))
			return
		}

		// 解析token
		var mapClaims jwt.MapClaims
		_, err := jwt.ParseWithClaims(strings.TrimSpace(string(tokenByte)), &mapClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(l.svcCtx.Config.CAuth.AccessSecret), nil
		})
		if err != nil {
			ws.WriteMessage(1, []byte("error|令牌错误|"))
			return
		}
		// jwt转换
		tokenData := jwtx.ParseTokenMap(mapClaims)

		// todo 鉴权微服务
		//err = authx.WebsocketAuthx(tokenData)
		//if err != nil {
		//	message := fmt.Sprintf("error|%s", err.Error())
		//	logx.Errorf(message)
		//	ws.WriteMessage(1, []byte(message))
		//	return
		//}

		// 注册上线
		l.registerUser(ws, tokenData)

		//给客户端发送信息
		err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("register|success|")})
		if err != nil {
			logx.Errorf(err.Error())
		}

		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			// 限流器
			if l.svcCtx.TaskMaxLimit.Allow() == false {
				err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("error|服务器任务数超过限制请稍后再试")})
				if err != nil {
					logx.Errorf(err.Error())
				}
				continue
			}

			// 解析内容
			msg := strings.Trim(string(message), " ")
			logx.Infof("socket receive [%s] %s", remoteAddr, msg)
			data := strings.Split(strings.Trim(msg, " "), "|") // 按竖线分割

			if len(data) == 0 {
				err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("error|发送格式错误请以`功能|内容`格式发送|")})
				if err != nil {
					logx.Errorf(err.Error())
				}
				continue
			}

			switch data[0] {
			case "ping":
				err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("pong|")})
				if err != nil {
					logx.Errorf(err.Error())
				}
			case "subscribe":
				if len(data) > 1 {
					go l.GateLogger(ws, data[1], tokenData)
				} else {
					err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("error|发送格式错误请以`功能|内容`格式发送|")})
					if err != nil {
						logx.Errorf(err.Error())
					}
				}
			case "close":
				go l.GateLoggerClose(ws)

			default:
				err = util.SendSocket(util.SendMessage{ws, fmt.Sprintf("error|不支持指令|")})
				if err != nil {
					logx.Errorf(err.Error())
				}
			}
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}
