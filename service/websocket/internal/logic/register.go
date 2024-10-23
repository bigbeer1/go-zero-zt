package logic

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/service/websocket/internal/svc"
)

type Websocket struct {
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(svcCtx *svc.ServiceContext) *Websocket {
	return &Websocket{
		svcCtx: svcCtx,
	}
}

func (l *Websocket) Register() (requestHandler fasthttp.RequestHandler) {

	go l.Xun()

	requestHandler = func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			l.EchoView(ctx)
		case "/home":
			l.MonitorHome(ctx)
		case "/alarm":
			l.AlarmSubscribe(ctx)
		case "/alarmHome":
			l.AlarmHome(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	return requestHandler
}

// 临时方法循环保持redis 值内容
func (l *Websocket) Xun() {

	for {
		time.Sleep(time.Minute * 3)
		for k, v := range global.AlarmOnlineList {
			for _, item := range v {
				global.AlarmUlock.Lock()
				redisKey := fmt.Sprintf("%s%s:%s", global.WebSocketUserPrefix, k, item.Ip)
				// redis 重置时间
				ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
				l.svcCtx.Redis.SetWithExpireCtx(ctx, redisKey, item, l.svcCtx.OnlineTime)
				global.AlarmUlock.Unlock()
			}
		}
	}

}
