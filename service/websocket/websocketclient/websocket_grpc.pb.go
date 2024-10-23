// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package websocketclient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WebsocketClient is the client API for Websocket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebsocketClient interface {
	// 发送告警消息
	AlarmMessage(ctx context.Context, in *AlarmMessageReq, opts ...grpc.CallOption) (*CommonResp, error)
}

type websocketClient struct {
	cc grpc.ClientConnInterface
}

func NewWebsocketClient(cc grpc.ClientConnInterface) WebsocketClient {
	return &websocketClient{cc}
}

func (c *websocketClient) AlarmMessage(ctx context.Context, in *AlarmMessageReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/websocketclient.websocket/AlarmMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebsocketServer is the server API for Websocket service.
// All implementations must embed UnimplementedWebsocketServer
// for forward compatibility
type WebsocketServer interface {
	// 发送告警消息
	AlarmMessage(context.Context, *AlarmMessageReq) (*CommonResp, error)
	mustEmbedUnimplementedWebsocketServer()
}

// UnimplementedWebsocketServer must be embedded to have forward compatible implementations.
type UnimplementedWebsocketServer struct {
}

func (UnimplementedWebsocketServer) AlarmMessage(context.Context, *AlarmMessageReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlarmMessage not implemented")
}
func (UnimplementedWebsocketServer) mustEmbedUnimplementedWebsocketServer() {}

// UnsafeWebsocketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebsocketServer will
// result in compilation errors.
type UnsafeWebsocketServer interface {
	mustEmbedUnimplementedWebsocketServer()
}

func RegisterWebsocketServer(s grpc.ServiceRegistrar, srv WebsocketServer) {
	s.RegisterService(&Websocket_ServiceDesc, srv)
}

func _Websocket_AlarmMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlarmMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebsocketServer).AlarmMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/websocketclient.websocket/AlarmMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebsocketServer).AlarmMessage(ctx, req.(*AlarmMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Websocket_ServiceDesc is the grpc.ServiceDesc for Websocket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Websocket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "websocketclient.websocket",
	HandlerType: (*WebsocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AlarmMessage",
			Handler:    _Websocket_AlarmMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "websocket.proto",
}
