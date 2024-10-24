// Code generated by goctl. DO NOT EDIT.
// Source: websocket.proto

package websocket

import (
	"context"
	"tpmt-zt/service/websocket/websocketclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlarmMessageReq = websocketclient.AlarmMessageReq
	CommonResp      = websocketclient.CommonResp

	Websocket interface {
		// 发送告警消息
		AlarmMessage(ctx context.Context, in *AlarmMessageReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultWebsocket struct {
		cli zrpc.Client
	}
)

func NewWebsocket(cli zrpc.Client) Websocket {
	return &defaultWebsocket{
		cli: cli,
	}
}

// 发送告警消息
func (m *defaultWebsocket) AlarmMessage(ctx context.Context, in *AlarmMessageReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := websocketclient.NewWebsocketClient(m.cli.Conn())
	return client.AlarmMessage(ctx, in, opts...)
}