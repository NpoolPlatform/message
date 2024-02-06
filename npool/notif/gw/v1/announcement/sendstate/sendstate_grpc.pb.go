// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/gw/v1/announcement/sendstate/sendstate.proto

package sendstate

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

const (
	Gateway_GetSendStates_FullMethodName        = "/notif.gateway.announcement.sendstate.v1.Gateway/GetSendStates"
	Gateway_GetAppUserSendStates_FullMethodName = "/notif.gateway.announcement.sendstate.v1.Gateway/GetAppUserSendStates"
	Gateway_GetAppSendStates_FullMethodName     = "/notif.gateway.announcement.sendstate.v1.Gateway/GetAppSendStates"
	Gateway_GetNAppSendStates_FullMethodName    = "/notif.gateway.announcement.sendstate.v1.Gateway/GetNAppSendStates"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetSendStates(ctx context.Context, in *GetSendStatesRequest, opts ...grpc.CallOption) (*GetSendStatesResponse, error)
	GetAppUserSendStates(ctx context.Context, in *GetAppUserSendStatesRequest, opts ...grpc.CallOption) (*GetAppUserSendStatesResponse, error)
	GetAppSendStates(ctx context.Context, in *GetAppSendStatesRequest, opts ...grpc.CallOption) (*GetAppSendStatesResponse, error)
	GetNAppSendStates(ctx context.Context, in *GetNAppSendStatesRequest, opts ...grpc.CallOption) (*GetNAppSendStatesResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetSendStates(ctx context.Context, in *GetSendStatesRequest, opts ...grpc.CallOption) (*GetSendStatesResponse, error) {
	out := new(GetSendStatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSendStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppUserSendStates(ctx context.Context, in *GetAppUserSendStatesRequest, opts ...grpc.CallOption) (*GetAppUserSendStatesResponse, error) {
	out := new(GetAppUserSendStatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppUserSendStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppSendStates(ctx context.Context, in *GetAppSendStatesRequest, opts ...grpc.CallOption) (*GetAppSendStatesResponse, error) {
	out := new(GetAppSendStatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppSendStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppSendStates(ctx context.Context, in *GetNAppSendStatesRequest, opts ...grpc.CallOption) (*GetNAppSendStatesResponse, error) {
	out := new(GetNAppSendStatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetNAppSendStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetSendStates(context.Context, *GetSendStatesRequest) (*GetSendStatesResponse, error)
	GetAppUserSendStates(context.Context, *GetAppUserSendStatesRequest) (*GetAppUserSendStatesResponse, error)
	GetAppSendStates(context.Context, *GetAppSendStatesRequest) (*GetAppSendStatesResponse, error)
	GetNAppSendStates(context.Context, *GetNAppSendStatesRequest) (*GetNAppSendStatesResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetSendStates(context.Context, *GetSendStatesRequest) (*GetSendStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendStates not implemented")
}
func (UnimplementedGatewayServer) GetAppUserSendStates(context.Context, *GetAppUserSendStatesRequest) (*GetAppUserSendStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserSendStates not implemented")
}
func (UnimplementedGatewayServer) GetAppSendStates(context.Context, *GetAppSendStatesRequest) (*GetAppSendStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSendStates not implemented")
}
func (UnimplementedGatewayServer) GetNAppSendStates(context.Context, *GetNAppSendStatesRequest) (*GetNAppSendStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppSendStates not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_GetSendStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSendStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSendStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSendStates(ctx, req.(*GetSendStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppUserSendStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserSendStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppUserSendStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppUserSendStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppUserSendStates(ctx, req.(*GetAppUserSendStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppSendStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSendStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppSendStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppSendStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppSendStates(ctx, req.(*GetAppSendStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppSendStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppSendStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppSendStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppSendStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppSendStates(ctx, req.(*GetNAppSendStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.announcement.sendstate.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSendStates",
			Handler:    _Gateway_GetSendStates_Handler,
		},
		{
			MethodName: "GetAppUserSendStates",
			Handler:    _Gateway_GetAppUserSendStates_Handler,
		},
		{
			MethodName: "GetAppSendStates",
			Handler:    _Gateway_GetAppSendStates_Handler,
		},
		{
			MethodName: "GetNAppSendStates",
			Handler:    _Gateway_GetNAppSendStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/gw/v1/announcement/sendstate/sendstate.proto",
}
