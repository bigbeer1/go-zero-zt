package logic

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"tpmt-zt/service/tpmtcom/svc"
	"tpmt-zt/service/tpmtcom/util"
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

	requestHandler = func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			l.EchoView(ctx)
		case "/home":
			l.Home(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	return requestHandler
}

/*
*
发送数据
*/
func (l *Websocket) SendSocketReq(Ws *websocket.Conn, Message string) {
	req := util.SendMessage{
		Ws:      Ws,
		Message: Message,
	}
	_ = l.svcCtx.SendSocketAntsPool.Invoke(req)
}

/*
*
发送数据
*/
func (l *Websocket) SendSocketNoLockReq(Ws *websocket.Conn, Message string) {
	req := util.SendMessage{
		Ws:      Ws,
		Message: Message,
	}
	_ = l.svcCtx.SendSocketNoLockAntsPool.Invoke(req)
}
