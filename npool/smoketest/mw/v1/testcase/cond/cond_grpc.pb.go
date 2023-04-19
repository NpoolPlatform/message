// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: npool/smoketest/mw/v1/testcase/cond/cond.proto

package cond

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCond(ctx context.Context, in *CreateCondRequest, opts ...grpc.CallOption) (*CreateCondResponse, error)
	UpdateCond(ctx context.Context, in *UpdateCondRequest, opts ...grpc.CallOption) (*UpdateCondResponse, error)
	GetCond(ctx context.Context, in *GetCondRequest, opts ...grpc.CallOption) (*GetCondResponse, error)
	GetConds(ctx context.Context, in *GetCondsRequest, opts ...grpc.CallOption) (*GetCondsResponse, error)
	GetCondOnly(ctx context.Context, in *GetCondOnlyRequest, opts ...grpc.CallOption) (*GetCondOnlyResponse, error)
	ExistCond(ctx context.Context, in *ExistCondRequest, opts ...grpc.CallOption) (*ExistCondResponse, error)
	DeleteCond(ctx context.Context, in *DeleteCondRequest, opts ...grpc.CallOption) (*DeleteCondResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCond(ctx context.Context, in *CreateCondRequest, opts ...grpc.CallOption) (*CreateCondResponse, error) {
	out := new(CreateCondResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/CreateCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCond(ctx context.Context, in *UpdateCondRequest, opts ...grpc.CallOption) (*UpdateCondResponse, error) {
	out := new(UpdateCondResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/UpdateCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCond(ctx context.Context, in *GetCondRequest, opts ...grpc.CallOption) (*GetCondResponse, error) {
	out := new(GetCondResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/GetCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetConds(ctx context.Context, in *GetCondsRequest, opts ...grpc.CallOption) (*GetCondsResponse, error) {
	out := new(GetCondsResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/GetConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCondOnly(ctx context.Context, in *GetCondOnlyRequest, opts ...grpc.CallOption) (*GetCondOnlyResponse, error) {
	out := new(GetCondOnlyResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/GetCondOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCond(ctx context.Context, in *ExistCondRequest, opts ...grpc.CallOption) (*ExistCondResponse, error) {
	out := new(ExistCondResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/ExistCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCond(ctx context.Context, in *DeleteCondRequest, opts ...grpc.CallOption) (*DeleteCondResponse, error) {
	out := new(DeleteCondResponse)
	err := c.cc.Invoke(ctx, "/smoketest.middleware.testcase.cond.v1.Middleware/DeleteCond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCond(context.Context, *CreateCondRequest) (*CreateCondResponse, error)
	UpdateCond(context.Context, *UpdateCondRequest) (*UpdateCondResponse, error)
	GetCond(context.Context, *GetCondRequest) (*GetCondResponse, error)
	GetConds(context.Context, *GetCondsRequest) (*GetCondsResponse, error)
	GetCondOnly(context.Context, *GetCondOnlyRequest) (*GetCondOnlyResponse, error)
	ExistCond(context.Context, *ExistCondRequest) (*ExistCondResponse, error)
	DeleteCond(context.Context, *DeleteCondRequest) (*DeleteCondResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCond(context.Context, *CreateCondRequest) (*CreateCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCond not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCond(context.Context, *UpdateCondRequest) (*UpdateCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCond not implemented")
}
func (UnimplementedMiddlewareServer) GetCond(context.Context, *GetCondRequest) (*GetCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCond not implemented")
}
func (UnimplementedMiddlewareServer) GetConds(context.Context, *GetCondsRequest) (*GetCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConds not implemented")
}
func (UnimplementedMiddlewareServer) GetCondOnly(context.Context, *GetCondOnlyRequest) (*GetCondOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCondOnly not implemented")
}
func (UnimplementedMiddlewareServer) ExistCond(context.Context, *ExistCondRequest) (*ExistCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCond not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCond(context.Context, *DeleteCondRequest) (*DeleteCondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCond not implemented")
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

func _Middleware_CreateCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/CreateCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCond(ctx, req.(*CreateCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/UpdateCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCond(ctx, req.(*UpdateCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/GetCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCond(ctx, req.(*GetCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/GetConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetConds(ctx, req.(*GetCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCondOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCondOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCondOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/GetCondOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCondOnly(ctx, req.(*GetCondOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/ExistCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCond(ctx, req.(*ExistCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.middleware.testcase.cond.v1.Middleware/DeleteCond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCond(ctx, req.(*DeleteCondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.middleware.testcase.cond.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCond",
			Handler:    _Middleware_CreateCond_Handler,
		},
		{
			MethodName: "UpdateCond",
			Handler:    _Middleware_UpdateCond_Handler,
		},
		{
			MethodName: "GetCond",
			Handler:    _Middleware_GetCond_Handler,
		},
		{
			MethodName: "GetConds",
			Handler:    _Middleware_GetConds_Handler,
		},
		{
			MethodName: "GetCondOnly",
			Handler:    _Middleware_GetCondOnly_Handler,
		},
		{
			MethodName: "ExistCond",
			Handler:    _Middleware_ExistCond_Handler,
		},
		{
			MethodName: "DeleteCond",
			Handler:    _Middleware_DeleteCond_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/mw/v1/testcase/cond/cond.proto",
}
