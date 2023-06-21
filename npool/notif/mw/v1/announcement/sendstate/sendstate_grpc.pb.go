// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/announcement/sendstate/sendstate.proto

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
	Middleware_CreateSendState_FullMethodName     = "/notif.middleware.announcement.sendstate.v1.Middleware/CreateSendState"
	Middleware_GetSendStates_FullMethodName       = "/notif.middleware.announcement.sendstate.v1.Middleware/GetSendStates"
	Middleware_GetSendState_FullMethodName        = "/notif.middleware.announcement.sendstate.v1.Middleware/GetSendState"
	Middleware_GetSendStateOnly_FullMethodName    = "/notif.middleware.announcement.sendstate.v1.Middleware/GetSendStateOnly"
	Middleware_ExistSendState_FullMethodName      = "/notif.middleware.announcement.sendstate.v1.Middleware/ExistSendState"
	Middleware_ExistSendStateConds_FullMethodName = "/notif.middleware.announcement.sendstate.v1.Middleware/ExistSendStateConds"
	Middleware_DeleteSendState_FullMethodName     = "/notif.middleware.announcement.sendstate.v1.Middleware/DeleteSendState"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateSendState(ctx context.Context, in *CreateSendStateRequest, opts ...grpc.CallOption) (*CreateSendStateResponse, error)
	GetSendStates(ctx context.Context, in *GetSendStatesRequest, opts ...grpc.CallOption) (*GetSendStatesResponse, error)
	GetSendState(ctx context.Context, in *GetSendStateRequest, opts ...grpc.CallOption) (*GetSendStateResponse, error)
	GetSendStateOnly(ctx context.Context, in *GetSendStateOnlyRequest, opts ...grpc.CallOption) (*GetSendStateOnlyResponse, error)
	ExistSendState(ctx context.Context, in *ExistSendStateRequest, opts ...grpc.CallOption) (*ExistSendStateResponse, error)
	ExistSendStateConds(ctx context.Context, in *ExistSendStateCondsRequest, opts ...grpc.CallOption) (*ExistSendStateCondsResponse, error)
	DeleteSendState(ctx context.Context, in *DeleteSendStateRequest, opts ...grpc.CallOption) (*DeleteSendStateResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateSendState(ctx context.Context, in *CreateSendStateRequest, opts ...grpc.CallOption) (*CreateSendStateResponse, error) {
	out := new(CreateSendStateResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateSendState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSendStates(ctx context.Context, in *GetSendStatesRequest, opts ...grpc.CallOption) (*GetSendStatesResponse, error) {
	out := new(GetSendStatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSendStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSendState(ctx context.Context, in *GetSendStateRequest, opts ...grpc.CallOption) (*GetSendStateResponse, error) {
	out := new(GetSendStateResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSendState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSendStateOnly(ctx context.Context, in *GetSendStateOnlyRequest, opts ...grpc.CallOption) (*GetSendStateOnlyResponse, error) {
	out := new(GetSendStateOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSendStateOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistSendState(ctx context.Context, in *ExistSendStateRequest, opts ...grpc.CallOption) (*ExistSendStateResponse, error) {
	out := new(ExistSendStateResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistSendState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistSendStateConds(ctx context.Context, in *ExistSendStateCondsRequest, opts ...grpc.CallOption) (*ExistSendStateCondsResponse, error) {
	out := new(ExistSendStateCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistSendStateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteSendState(ctx context.Context, in *DeleteSendStateRequest, opts ...grpc.CallOption) (*DeleteSendStateResponse, error) {
	out := new(DeleteSendStateResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteSendState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateSendState(context.Context, *CreateSendStateRequest) (*CreateSendStateResponse, error)
	GetSendStates(context.Context, *GetSendStatesRequest) (*GetSendStatesResponse, error)
	GetSendState(context.Context, *GetSendStateRequest) (*GetSendStateResponse, error)
	GetSendStateOnly(context.Context, *GetSendStateOnlyRequest) (*GetSendStateOnlyResponse, error)
	ExistSendState(context.Context, *ExistSendStateRequest) (*ExistSendStateResponse, error)
	ExistSendStateConds(context.Context, *ExistSendStateCondsRequest) (*ExistSendStateCondsResponse, error)
	DeleteSendState(context.Context, *DeleteSendStateRequest) (*DeleteSendStateResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateSendState(context.Context, *CreateSendStateRequest) (*CreateSendStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSendState not implemented")
}
func (UnimplementedMiddlewareServer) GetSendStates(context.Context, *GetSendStatesRequest) (*GetSendStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendStates not implemented")
}
func (UnimplementedMiddlewareServer) GetSendState(context.Context, *GetSendStateRequest) (*GetSendStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendState not implemented")
}
func (UnimplementedMiddlewareServer) GetSendStateOnly(context.Context, *GetSendStateOnlyRequest) (*GetSendStateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendStateOnly not implemented")
}
func (UnimplementedMiddlewareServer) ExistSendState(context.Context, *ExistSendStateRequest) (*ExistSendStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSendState not implemented")
}
func (UnimplementedMiddlewareServer) ExistSendStateConds(context.Context, *ExistSendStateCondsRequest) (*ExistSendStateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSendStateConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteSendState(context.Context, *DeleteSendStateRequest) (*DeleteSendStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSendState not implemented")
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

func _Middleware_CreateSendState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSendStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateSendState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateSendState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateSendState(ctx, req.(*CreateSendStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSendStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSendStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSendStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSendStates(ctx, req.(*GetSendStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSendState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSendState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSendState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSendState(ctx, req.(*GetSendStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSendStateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendStateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSendStateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSendStateOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSendStateOnly(ctx, req.(*GetSendStateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistSendState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSendStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistSendState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistSendState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistSendState(ctx, req.(*ExistSendStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistSendStateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSendStateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistSendStateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistSendStateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistSendStateConds(ctx, req.(*ExistSendStateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteSendState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSendStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteSendState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteSendState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteSendState(ctx, req.(*DeleteSendStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.announcement.sendstate.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSendState",
			Handler:    _Middleware_CreateSendState_Handler,
		},
		{
			MethodName: "GetSendStates",
			Handler:    _Middleware_GetSendStates_Handler,
		},
		{
			MethodName: "GetSendState",
			Handler:    _Middleware_GetSendState_Handler,
		},
		{
			MethodName: "GetSendStateOnly",
			Handler:    _Middleware_GetSendStateOnly_Handler,
		},
		{
			MethodName: "ExistSendState",
			Handler:    _Middleware_ExistSendState_Handler,
		},
		{
			MethodName: "ExistSendStateConds",
			Handler:    _Middleware_ExistSendStateConds_Handler,
		},
		{
			MethodName: "DeleteSendState",
			Handler:    _Middleware_DeleteSendState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/announcement/sendstate/sendstate.proto",
}
