package logic

import (
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mssola/user_agent"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"strings"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/websocket/internal/util"
)

func (l *Websocket) AlarmSubscribe(ctx *fasthttp.RequestCtx) {
	// 获取token
	tokenByte := ctx.Request.Header.Peek("Sec-WebSocket-Protocol")
	// 返回必须同样返回 否则websocket 错误
	ctx.Response.Header.Set("Sec-WebSocket-Protocol", string(tokenByte))

	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		// ip和端口
		remoteAddr := ws.RemoteAddr().String()
		// 设备
		var device string = "未知"

		defer func() {
			global.AlarmUlock.Lock()
			util.AlarmDelSocket(ctx, remoteAddr, l.svcCtx.Redis)
			ws.Close()
			global.AlarmUlock.Unlock()
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

		// 获取设备信息
		userAgent := ctx.Request.Header.Peek("User-Agent")

		if len(userAgent) != 0 {
			ua := user_agent.New(string(userAgent))
			uaSee, uaVersion := ua.Browser()
			device = fmt.Sprintf("%s/%s", ua.OS(), uaSee+" "+uaVersion)
		}

		// 注册上线
		l.registerForAlarm(ctx, ws, tokenData, device)

		//给客户端发送信息
		err = util.AlarmSendSocket(ws, fmt.Sprintf("register|success|"))
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
				err = util.AlarmSendSocket(ws, fmt.Sprintf("error|服务器任务数超过限制请稍后再试"))
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
				err = util.AlarmSendSocket(ws, fmt.Sprintf("error|发送格式错误请以`功能|内容`格式发送|"))
				if err != nil {
					logx.Errorf(err.Error())
				}
				continue
			}

			switch data[0] {
			case "ping":
				err = util.AlarmSendSocket(ws, fmt.Sprintf("pong|"))
				if err != nil {
					logx.Errorf(err.Error())
				}
			default:
				err = util.AlarmSendSocket(ws, fmt.Sprintf("error|不支持指令|"))
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
