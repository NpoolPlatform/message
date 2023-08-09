// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/good/default/default.proto

package _default

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
	Middleware_CreateDefault_FullMethodName  = "/good.middleware.app.good.default1.v1.Middleware/CreateDefault"
	Middleware_UpdateDefault_FullMethodName  = "/good.middleware.app.good.default1.v1.Middleware/UpdateDefault"
	Middleware_GetDefault_FullMethodName     = "/good.middleware.app.good.default1.v1.Middleware/GetDefault"
	Middleware_GetDefaultOnly_FullMethodName = "/good.middleware.app.good.default1.v1.Middleware/GetDefaultOnly"
	Middleware_GetDefaults_FullMethodName    = "/good.middleware.app.good.default1.v1.Middleware/GetDefaults"
	Middleware_DeleteDefault_FullMethodName  = "/good.middleware.app.good.default1.v1.Middleware/DeleteDefault"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateDefault(ctx context.Context, in *CreateDefaultRequest, opts ...grpc.CallOption) (*CreateDefaultResponse, error)
	UpdateDefault(ctx context.Context, in *UpdateDefaultRequest, opts ...grpc.CallOption) (*UpdateDefaultResponse, error)
	GetDefault(ctx context.Context, in *GetDefaultRequest, opts ...grpc.CallOption) (*GetDefaultResponse, error)
	GetDefaultOnly(ctx context.Context, in *GetDefaultOnlyRequest, opts ...grpc.CallOption) (*GetDefaultOnlyResponse, error)
	GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...grpc.CallOption) (*GetDefaultsResponse, error)
	DeleteDefault(ctx context.Context, in *DeleteDefaultRequest, opts ...grpc.CallOption) (*DeleteDefaultResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateDefault(ctx context.Context, in *CreateDefaultRequest, opts ...grpc.CallOption) (*CreateDefaultResponse, error) {
	out := new(CreateDefaultResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateDefault(ctx context.Context, in *UpdateDefaultRequest, opts ...grpc.CallOption) (*UpdateDefaultResponse, error) {
	out := new(UpdateDefaultResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDefault(ctx context.Context, in *GetDefaultRequest, opts ...grpc.CallOption) (*GetDefaultResponse, error) {
	out := new(GetDefaultResponse)
	err := c.cc.Invoke(ctx, Middleware_GetDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDefaultOnly(ctx context.Context, in *GetDefaultOnlyRequest, opts ...grpc.CallOption) (*GetDefaultOnlyResponse, error) {
	out := new(GetDefaultOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetDefaultOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...grpc.CallOption) (*GetDefaultsResponse, error) {
	out := new(GetDefaultsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetDefaults_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteDefault(ctx context.Context, in *DeleteDefaultRequest, opts ...grpc.CallOption) (*DeleteDefaultResponse, error) {
	out := new(DeleteDefaultResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateDefault(context.Context, *CreateDefaultRequest) (*CreateDefaultResponse, error)
	UpdateDefault(context.Context, *UpdateDefaultRequest) (*UpdateDefaultResponse, error)
	GetDefault(context.Context, *GetDefaultRequest) (*GetDefaultResponse, error)
	GetDefaultOnly(context.Context, *GetDefaultOnlyRequest) (*GetDefaultOnlyResponse, error)
	GetDefaults(context.Context, *GetDefaultsRequest) (*GetDefaultsResponse, error)
	DeleteDefault(context.Context, *DeleteDefaultRequest) (*DeleteDefaultResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateDefault(context.Context, *CreateDefaultRequest) (*CreateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefault not implemented")
}
func (UnimplementedMiddlewareServer) UpdateDefault(context.Context, *UpdateDefaultRequest) (*UpdateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDefault not implemented")
}
func (UnimplementedMiddlewareServer) GetDefault(context.Context, *GetDefaultRequest) (*GetDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefault not implemented")
}
func (UnimplementedMiddlewareServer) GetDefaultOnly(context.Context, *GetDefaultOnlyRequest) (*GetDefaultOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetDefaults(context.Context, *GetDefaultsRequest) (*GetDefaultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaults not implemented")
}
func (UnimplementedMiddlewareServer) DeleteDefault(context.Context, *DeleteDefaultRequest) (*DeleteDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefault not implemented")
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

func _Middleware_CreateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateDefault(ctx, req.(*CreateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateDefault(ctx, req.(*UpdateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDefault(ctx, req.(*GetDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDefaultOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDefaultOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetDefaultOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDefaultOnly(ctx, req.(*GetDefaultOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetDefaults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDefaults(ctx, req.(*GetDefaultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteDefault(ctx, req.(*DeleteDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.good.default1.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDefault",
			Handler:    _Middleware_CreateDefault_Handler,
		},
		{
			MethodName: "UpdateDefault",
			Handler:    _Middleware_UpdateDefault_Handler,
		},
		{
			MethodName: "GetDefault",
			Handler:    _Middleware_GetDefault_Handler,
		},
		{
			MethodName: "GetDefaultOnly",
			Handler:    _Middleware_GetDefaultOnly_Handler,
		},
		{
			MethodName: "GetDefaults",
			Handler:    _Middleware_GetDefaults_Handler,
		},
		{
			MethodName: "DeleteDefault",
			Handler:    _Middleware_DeleteDefault_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/good/default/default.proto",
}
