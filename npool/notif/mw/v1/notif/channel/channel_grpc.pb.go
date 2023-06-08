// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/notif/channel/channel.proto

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

const (
	Middleware_GetChannels_FullMethodName       = "/notif.middleware.channel.v1.Middleware/GetChannels"
	Middleware_GetChannelOnly_FullMethodName    = "/notif.middleware.channel.v1.Middleware/GetChannelOnly"
	Middleware_CreateChannel_FullMethodName     = "/notif.middleware.channel.v1.Middleware/CreateChannel"
	Middleware_CreateChannels_FullMethodName    = "/notif.middleware.channel.v1.Middleware/CreateChannels"
	Middleware_UpdateChannel_FullMethodName     = "/notif.middleware.channel.v1.Middleware/UpdateChannel"
	Middleware_GetChannel_FullMethodName        = "/notif.middleware.channel.v1.Middleware/GetChannel"
	Middleware_ExistChannel_FullMethodName      = "/notif.middleware.channel.v1.Middleware/ExistChannel"
	Middleware_ExistChannelConds_FullMethodName = "/notif.middleware.channel.v1.Middleware/ExistChannelConds"
	Middleware_DeleteChannel_FullMethodName     = "/notif.middleware.channel.v1.Middleware/DeleteChannel"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error)
	GetChannelOnly(ctx context.Context, in *GetChannelOnlyRequest, opts ...grpc.CallOption) (*GetChannelOnlyResponse, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	CreateChannels(ctx context.Context, in *CreateChannelsRequest, opts ...grpc.CallOption) (*CreateChannelsResponse, error)
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error)
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error)
	ExistChannel(ctx context.Context, in *ExistChannelRequest, opts ...grpc.CallOption) (*ExistChannelResponse, error)
	ExistChannelConds(ctx context.Context, in *ExistChannelCondsRequest, opts ...grpc.CallOption) (*ExistChannelCondsResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetChannels(ctx context.Context, in *GetChannelsRequest, opts ...grpc.CallOption) (*GetChannelsResponse, error) {
	out := new(GetChannelsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetChannels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetChannelOnly(ctx context.Context, in *GetChannelOnlyRequest, opts ...grpc.CallOption) (*GetChannelOnlyResponse, error) {
	out := new(GetChannelOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetChannelOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateChannels(ctx context.Context, in *CreateChannelsRequest, opts ...grpc.CallOption) (*CreateChannelsResponse, error) {
	out := new(CreateChannelsResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateChannels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error) {
	out := new(UpdateChannelResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*GetChannelResponse, error) {
	out := new(GetChannelResponse)
	err := c.cc.Invoke(ctx, Middleware_GetChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistChannel(ctx context.Context, in *ExistChannelRequest, opts ...grpc.CallOption) (*ExistChannelResponse, error) {
	out := new(ExistChannelResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistChannelConds(ctx context.Context, in *ExistChannelCondsRequest, opts ...grpc.CallOption) (*ExistChannelCondsResponse, error) {
	out := new(ExistChannelCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistChannelConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error)
	GetChannelOnly(context.Context, *GetChannelOnlyRequest) (*GetChannelOnlyResponse, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	CreateChannels(context.Context, *CreateChannelsRequest) (*CreateChannelsResponse, error)
	UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error)
	GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error)
	ExistChannel(context.Context, *ExistChannelRequest) (*ExistChannelResponse, error)
	ExistChannelConds(context.Context, *ExistChannelCondsRequest) (*ExistChannelCondsResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetChannels(context.Context, *GetChannelsRequest) (*GetChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannels not implemented")
}
func (UnimplementedMiddlewareServer) GetChannelOnly(context.Context, *GetChannelOnlyRequest) (*GetChannelOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelOnly not implemented")
}
func (UnimplementedMiddlewareServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedMiddlewareServer) CreateChannels(context.Context, *CreateChannelsRequest) (*CreateChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannels not implemented")
}
func (UnimplementedMiddlewareServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedMiddlewareServer) GetChannel(context.Context, *GetChannelRequest) (*GetChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (UnimplementedMiddlewareServer) ExistChannel(context.Context, *ExistChannelRequest) (*ExistChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistChannel not implemented")
}
func (UnimplementedMiddlewareServer) ExistChannelConds(context.Context, *ExistChannelCondsRequest) (*ExistChannelCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistChannelConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_GetChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetChannels(ctx, req.(*GetChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetChannelOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetChannelOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetChannelOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetChannelOnly(ctx, req.(*GetChannelOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateChannels(ctx, req.(*CreateChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistChannel(ctx, req.(*ExistChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistChannelConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistChannelCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistChannelConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistChannelConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistChannelConds(ctx, req.(*ExistChannelCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.channel.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChannels",
			Handler:    _Middleware_GetChannels_Handler,
		},
		{
			MethodName: "GetChannelOnly",
			Handler:    _Middleware_GetChannelOnly_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _Middleware_CreateChannel_Handler,
		},
		{
			MethodName: "CreateChannels",
			Handler:    _Middleware_CreateChannels_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _Middleware_UpdateChannel_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _Middleware_GetChannel_Handler,
		},
		{
			MethodName: "ExistChannel",
			Handler:    _Middleware_ExistChannel_Handler,
		},
		{
			MethodName: "ExistChannelConds",
			Handler:    _Middleware_ExistChannelConds_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _Middleware_DeleteChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/notif/channel/channel.proto",
}
