// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/gw/v1/notif/channel/channel.proto

package channel

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateChannels(ctx context.Context, in *CreateChannelsRequest, opts ...grpc.CallOption) (*CreateChannelsResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	GetAppChannels(ctx context.Context, in *GetAppChannelsRequest, opts ...grpc.CallOption) (*GetAppChannelsResponse, error)
	GetNAppChannels(ctx context.Context, in *GetNAppChannelsRequest, opts ...grpc.CallOption) (*GetNAppChannelsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateChannels(ctx context.Context, in *CreateChannelsRequest, opts ...grpc.CallOption) (*CreateChannelsResponse, error) {
	out := new(CreateChannelsResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif2.channel.v1.Gateway/CreateChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif2.channel.v1.Gateway/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppChannels(ctx context.Context, in *GetAppChannelsRequest, opts ...grpc.CallOption) (*GetAppChannelsResponse, error) {
	out := new(GetAppChannelsResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif2.channel.v1.Gateway/GetAppChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppChannels(ctx context.Context, in *GetNAppChannelsRequest, opts ...grpc.CallOption) (*GetNAppChannelsResponse, error) {
	out := new(GetNAppChannelsResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif2.channel.v1.Gateway/GetNAppChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateChannels(context.Context, *CreateChannelsRequest) (*CreateChannelsResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	GetAppChannels(context.Context, *GetAppChannelsRequest) (*GetAppChannelsResponse, error)
	GetNAppChannels(context.Context, *GetNAppChannelsRequest) (*GetNAppChannelsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateChannels(context.Context, *CreateChannelsRequest) (*CreateChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannels not implemented")
}
func (UnimplementedGatewayServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedGatewayServer) GetAppChannels(context.Context, *GetAppChannelsRequest) (*GetAppChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppChannels not implemented")
}
func (UnimplementedGatewayServer) GetNAppChannels(context.Context, *GetNAppChannelsRequest) (*GetNAppChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppChannels not implemented")
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

func _Gateway_CreateChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif2.channel.v1.Gateway/CreateChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateChannels(ctx, req.(*CreateChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif2.channel.v1.Gateway/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif2.channel.v1.Gateway/GetAppChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppChannels(ctx, req.(*GetAppChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif2.channel.v1.Gateway/GetNAppChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppChannels(ctx, req.(*GetNAppChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.notif2.channel.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannels",
			Handler:    _Gateway_CreateChannels_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _Gateway_DeleteChannel_Handler,
		},
		{
			MethodName: "GetAppChannels",
			Handler:    _Gateway_GetAppChannels_Handler,
		},
		{
			MethodName: "GetNAppChannels",
			Handler:    _Gateway_GetNAppChannels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/gw/v1/notif/channel/channel.proto",
}
