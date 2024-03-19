// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/app/good/commission/config/config.proto

package config

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
	Middleware_CreateAppGoodCommissionConfig_FullMethodName     = "/inspire.middleware.app.good.commission.config.v1.Middleware/CreateAppGoodCommissionConfig"
	Middleware_UpdateAppGoodCommissionConfig_FullMethodName     = "/inspire.middleware.app.good.commission.config.v1.Middleware/UpdateAppGoodCommissionConfig"
	Middleware_GetAppGoodCommissionConfig_FullMethodName        = "/inspire.middleware.app.good.commission.config.v1.Middleware/GetAppGoodCommissionConfig"
	Middleware_GetAppGoodCommissionConfigs_FullMethodName       = "/inspire.middleware.app.good.commission.config.v1.Middleware/GetAppGoodCommissionConfigs"
	Middleware_ExistAppGoodCommissionConfigConds_FullMethodName = "/inspire.middleware.app.good.commission.config.v1.Middleware/ExistAppGoodCommissionConfigConds"
	Middleware_DeleteAppGoodCommissionConfig_FullMethodName     = "/inspire.middleware.app.good.commission.config.v1.Middleware/DeleteAppGoodCommissionConfig"
	Middleware_CloneAppGoodCommissionConfigs_FullMethodName     = "/inspire.middleware.app.good.commission.config.v1.Middleware/CloneAppGoodCommissionConfigs"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfig(ctx context.Context, in *GetAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error)
	ExistAppGoodCommissionConfigConds(ctx context.Context, in *ExistAppGoodCommissionConfigCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodCommissionConfigCondsResponse, error)
	DeleteAppGoodCommissionConfig(ctx context.Context, in *DeleteAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*DeleteAppGoodCommissionConfigResponse, error)
	CloneAppGoodCommissionConfigs(ctx context.Context, in *CloneAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneAppGoodCommissionConfigsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error) {
	out := new(CreateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error) {
	out := new(UpdateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppGoodCommissionConfig(ctx context.Context, in *GetAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigResponse, error) {
	out := new(GetAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error) {
	out := new(GetAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAppGoodCommissionConfigConds(ctx context.Context, in *ExistAppGoodCommissionConfigCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodCommissionConfigCondsResponse, error) {
	out := new(ExistAppGoodCommissionConfigCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistAppGoodCommissionConfigConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAppGoodCommissionConfig(ctx context.Context, in *DeleteAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*DeleteAppGoodCommissionConfigResponse, error) {
	out := new(DeleteAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CloneAppGoodCommissionConfigs(ctx context.Context, in *CloneAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneAppGoodCommissionConfigsResponse, error) {
	out := new(CloneAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Middleware_CloneAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfig(context.Context, *GetAppGoodCommissionConfigRequest) (*GetAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error)
	ExistAppGoodCommissionConfigConds(context.Context, *ExistAppGoodCommissionConfigCondsRequest) (*ExistAppGoodCommissionConfigCondsResponse, error)
	DeleteAppGoodCommissionConfig(context.Context, *DeleteAppGoodCommissionConfigRequest) (*DeleteAppGoodCommissionConfigResponse, error)
	CloneAppGoodCommissionConfigs(context.Context, *CloneAppGoodCommissionConfigsRequest) (*CloneAppGoodCommissionConfigsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodCommissionConfig not implemented")
}
func (UnimplementedMiddlewareServer) UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppGoodCommissionConfig not implemented")
}
func (UnimplementedMiddlewareServer) GetAppGoodCommissionConfig(context.Context, *GetAppGoodCommissionConfigRequest) (*GetAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodCommissionConfig not implemented")
}
func (UnimplementedMiddlewareServer) GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodCommissionConfigs not implemented")
}
func (UnimplementedMiddlewareServer) ExistAppGoodCommissionConfigConds(context.Context, *ExistAppGoodCommissionConfigCondsRequest) (*ExistAppGoodCommissionConfigCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppGoodCommissionConfigConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAppGoodCommissionConfig(context.Context, *DeleteAppGoodCommissionConfigRequest) (*DeleteAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppGoodCommissionConfig not implemented")
}
func (UnimplementedMiddlewareServer) CloneAppGoodCommissionConfigs(context.Context, *CloneAppGoodCommissionConfigsRequest) (*CloneAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneAppGoodCommissionConfigs not implemented")
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

func _Middleware_CreateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAppGoodCommissionConfig(ctx, req.(*CreateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateAppGoodCommissionConfig(ctx, req.(*UpdateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppGoodCommissionConfig(ctx, req.(*GetAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppGoodCommissionConfigs(ctx, req.(*GetAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAppGoodCommissionConfigConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppGoodCommissionConfigCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAppGoodCommissionConfigConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistAppGoodCommissionConfigConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAppGoodCommissionConfigConds(ctx, req.(*ExistAppGoodCommissionConfigCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAppGoodCommissionConfig(ctx, req.(*DeleteAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CloneAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CloneAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CloneAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CloneAppGoodCommissionConfigs(ctx, req.(*CloneAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.app.good.commission.config.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppGoodCommissionConfig",
			Handler:    _Middleware_CreateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "UpdateAppGoodCommissionConfig",
			Handler:    _Middleware_UpdateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "GetAppGoodCommissionConfig",
			Handler:    _Middleware_GetAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "GetAppGoodCommissionConfigs",
			Handler:    _Middleware_GetAppGoodCommissionConfigs_Handler,
		},
		{
			MethodName: "ExistAppGoodCommissionConfigConds",
			Handler:    _Middleware_ExistAppGoodCommissionConfigConds_Handler,
		},
		{
			MethodName: "DeleteAppGoodCommissionConfig",
			Handler:    _Middleware_DeleteAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "CloneAppGoodCommissionConfigs",
			Handler:    _Middleware_CloneAppGoodCommissionConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/app/good/commission/config/config.proto",
}
