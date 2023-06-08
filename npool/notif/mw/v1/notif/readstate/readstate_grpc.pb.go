// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/notif/readstate/readstate.proto

package readstate

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
	Middleware_CreateReadState_FullMethodName     = "/notif.middleware.notif.readstate.v1.Middleware/CreateReadState"
	Middleware_CreateReadStates_FullMethodName    = "/notif.middleware.notif.readstate.v1.Middleware/CreateReadStates"
	Middleware_UpdateReadState_FullMethodName     = "/notif.middleware.notif.readstate.v1.Middleware/UpdateReadState"
	Middleware_GetReadState_FullMethodName        = "/notif.middleware.notif.readstate.v1.Middleware/GetReadState"
	Middleware_GetReadStateOnly_FullMethodName    = "/notif.middleware.notif.readstate.v1.Middleware/GetReadStateOnly"
	Middleware_GetReadStates_FullMethodName       = "/notif.middleware.notif.readstate.v1.Middleware/GetReadStates"
	Middleware_ExistReadState_FullMethodName      = "/notif.middleware.notif.readstate.v1.Middleware/ExistReadState"
	Middleware_ExistReadStateConds_FullMethodName = "/notif.middleware.notif.readstate.v1.Middleware/ExistReadStateConds"
	Middleware_CountReadStates_FullMethodName     = "/notif.middleware.notif.readstate.v1.Middleware/CountReadStates"
	Middleware_DeleteReadState_FullMethodName     = "/notif.middleware.notif.readstate.v1.Middleware/DeleteReadState"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateReadState(ctx context.Context, in *CreateReadStateRequest, opts ...grpc.CallOption) (*CreateReadStateResponse, error)
	CreateReadStates(ctx context.Context, in *CreateReadStatesRequest, opts ...grpc.CallOption) (*CreateReadStatesResponse, error)
	UpdateReadState(ctx context.Context, in *UpdateReadStateRequest, opts ...grpc.CallOption) (*UpdateReadStateResponse, error)
	GetReadState(ctx context.Context, in *GetReadStateRequest, opts ...grpc.CallOption) (*GetReadStateResponse, error)
	GetReadStateOnly(ctx context.Context, in *GetReadStateOnlyRequest, opts ...grpc.CallOption) (*GetReadStateOnlyResponse, error)
	GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error)
	ExistReadState(ctx context.Context, in *ExistReadStateRequest, opts ...grpc.CallOption) (*ExistReadStateResponse, error)
	ExistReadStateConds(ctx context.Context, in *ExistReadStateCondsRequest, opts ...grpc.CallOption) (*ExistReadStateCondsResponse, error)
	CountReadStates(ctx context.Context, in *CountReadStatesRequest, opts ...grpc.CallOption) (*CountReadStatesResponse, error)
	DeleteReadState(ctx context.Context, in *DeleteReadStateRequest, opts ...grpc.CallOption) (*DeleteReadStateResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateReadState(ctx context.Context, in *CreateReadStateRequest, opts ...grpc.CallOption) (*CreateReadStateResponse, error) {
	out := new(CreateReadStateResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateReadState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateReadStates(ctx context.Context, in *CreateReadStatesRequest, opts ...grpc.CallOption) (*CreateReadStatesResponse, error) {
	out := new(CreateReadStatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateReadStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateReadState(ctx context.Context, in *UpdateReadStateRequest, opts ...grpc.CallOption) (*UpdateReadStateResponse, error) {
	out := new(UpdateReadStateResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateReadState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetReadState(ctx context.Context, in *GetReadStateRequest, opts ...grpc.CallOption) (*GetReadStateResponse, error) {
	out := new(GetReadStateResponse)
	err := c.cc.Invoke(ctx, Middleware_GetReadState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetReadStateOnly(ctx context.Context, in *GetReadStateOnlyRequest, opts ...grpc.CallOption) (*GetReadStateOnlyResponse, error) {
	out := new(GetReadStateOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetReadStateOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error) {
	out := new(GetReadStatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetReadStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistReadState(ctx context.Context, in *ExistReadStateRequest, opts ...grpc.CallOption) (*ExistReadStateResponse, error) {
	out := new(ExistReadStateResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistReadState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistReadStateConds(ctx context.Context, in *ExistReadStateCondsRequest, opts ...grpc.CallOption) (*ExistReadStateCondsResponse, error) {
	out := new(ExistReadStateCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistReadStateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountReadStates(ctx context.Context, in *CountReadStatesRequest, opts ...grpc.CallOption) (*CountReadStatesResponse, error) {
	out := new(CountReadStatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CountReadStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteReadState(ctx context.Context, in *DeleteReadStateRequest, opts ...grpc.CallOption) (*DeleteReadStateResponse, error) {
	out := new(DeleteReadStateResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteReadState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateReadState(context.Context, *CreateReadStateRequest) (*CreateReadStateResponse, error)
	CreateReadStates(context.Context, *CreateReadStatesRequest) (*CreateReadStatesResponse, error)
	UpdateReadState(context.Context, *UpdateReadStateRequest) (*UpdateReadStateResponse, error)
	GetReadState(context.Context, *GetReadStateRequest) (*GetReadStateResponse, error)
	GetReadStateOnly(context.Context, *GetReadStateOnlyRequest) (*GetReadStateOnlyResponse, error)
	GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error)
	ExistReadState(context.Context, *ExistReadStateRequest) (*ExistReadStateResponse, error)
	ExistReadStateConds(context.Context, *ExistReadStateCondsRequest) (*ExistReadStateCondsResponse, error)
	CountReadStates(context.Context, *CountReadStatesRequest) (*CountReadStatesResponse, error)
	DeleteReadState(context.Context, *DeleteReadStateRequest) (*DeleteReadStateResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateReadState(context.Context, *CreateReadStateRequest) (*CreateReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadState not implemented")
}
func (UnimplementedMiddlewareServer) CreateReadStates(context.Context, *CreateReadStatesRequest) (*CreateReadStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadStates not implemented")
}
func (UnimplementedMiddlewareServer) UpdateReadState(context.Context, *UpdateReadStateRequest) (*UpdateReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReadState not implemented")
}
func (UnimplementedMiddlewareServer) GetReadState(context.Context, *GetReadStateRequest) (*GetReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadState not implemented")
}
func (UnimplementedMiddlewareServer) GetReadStateOnly(context.Context, *GetReadStateOnlyRequest) (*GetReadStateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadStateOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadStates not implemented")
}
func (UnimplementedMiddlewareServer) ExistReadState(context.Context, *ExistReadStateRequest) (*ExistReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistReadState not implemented")
}
func (UnimplementedMiddlewareServer) ExistReadStateConds(context.Context, *ExistReadStateCondsRequest) (*ExistReadStateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistReadStateConds not implemented")
}
func (UnimplementedMiddlewareServer) CountReadStates(context.Context, *CountReadStatesRequest) (*CountReadStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountReadStates not implemented")
}
func (UnimplementedMiddlewareServer) DeleteReadState(context.Context, *DeleteReadStateRequest) (*DeleteReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReadState not implemented")
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

func _Middleware_CreateReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateReadState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateReadState(ctx, req.(*CreateReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateReadStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateReadStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateReadStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateReadStates(ctx, req.(*CreateReadStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateReadState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateReadState(ctx, req.(*UpdateReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetReadState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReadState(ctx, req.(*GetReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetReadStateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReadStateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetReadStateOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReadStateOnly(ctx, req.(*GetReadStateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetReadStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReadStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetReadStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReadStates(ctx, req.(*GetReadStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistReadState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistReadState(ctx, req.(*ExistReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistReadStateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistReadStateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistReadStateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistReadStateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistReadStateConds(ctx, req.(*ExistReadStateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountReadStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountReadStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountReadStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountReadStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountReadStates(ctx, req.(*CountReadStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteReadState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteReadState(ctx, req.(*DeleteReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.notif.readstate.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReadState",
			Handler:    _Middleware_CreateReadState_Handler,
		},
		{
			MethodName: "CreateReadStates",
			Handler:    _Middleware_CreateReadStates_Handler,
		},
		{
			MethodName: "UpdateReadState",
			Handler:    _Middleware_UpdateReadState_Handler,
		},
		{
			MethodName: "GetReadState",
			Handler:    _Middleware_GetReadState_Handler,
		},
		{
			MethodName: "GetReadStateOnly",
			Handler:    _Middleware_GetReadStateOnly_Handler,
		},
		{
			MethodName: "GetReadStates",
			Handler:    _Middleware_GetReadStates_Handler,
		},
		{
			MethodName: "ExistReadState",
			Handler:    _Middleware_ExistReadState_Handler,
		},
		{
			MethodName: "ExistReadStateConds",
			Handler:    _Middleware_ExistReadStateConds_Handler,
		},
		{
			MethodName: "CountReadStates",
			Handler:    _Middleware_CountReadStates_Handler,
		},
		{
			MethodName: "DeleteReadState",
			Handler:    _Middleware_DeleteReadState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/notif/readstate/readstate.proto",
}
