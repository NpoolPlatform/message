// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/coin/config/config.proto

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
	Middleware_CreateCoinConfig_FullMethodName     = "/inspire.middleware.coin.config.v1.Middleware/CreateCoinConfig"
	Middleware_UpdateCoinConfig_FullMethodName     = "/inspire.middleware.coin.config.v1.Middleware/UpdateCoinConfig"
	Middleware_GetCoinConfig_FullMethodName        = "/inspire.middleware.coin.config.v1.Middleware/GetCoinConfig"
	Middleware_GetCoinConfigs_FullMethodName       = "/inspire.middleware.coin.config.v1.Middleware/GetCoinConfigs"
	Middleware_ExistCoinConfigConds_FullMethodName = "/inspire.middleware.coin.config.v1.Middleware/ExistCoinConfigConds"
	Middleware_DeleteCoinConfig_FullMethodName     = "/inspire.middleware.coin.config.v1.Middleware/DeleteCoinConfig"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCoinConfig(ctx context.Context, in *CreateCoinConfigRequest, opts ...grpc.CallOption) (*CreateCoinConfigResponse, error)
	UpdateCoinConfig(ctx context.Context, in *UpdateCoinConfigRequest, opts ...grpc.CallOption) (*UpdateCoinConfigResponse, error)
	GetCoinConfig(ctx context.Context, in *GetCoinConfigRequest, opts ...grpc.CallOption) (*GetCoinConfigResponse, error)
	GetCoinConfigs(ctx context.Context, in *GetCoinConfigsRequest, opts ...grpc.CallOption) (*GetCoinConfigsResponse, error)
	ExistCoinConfigConds(ctx context.Context, in *ExistCoinConfigCondsRequest, opts ...grpc.CallOption) (*ExistCoinConfigCondsResponse, error)
	DeleteCoinConfig(ctx context.Context, in *DeleteCoinConfigRequest, opts ...grpc.CallOption) (*DeleteCoinConfigResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCoinConfig(ctx context.Context, in *CreateCoinConfigRequest, opts ...grpc.CallOption) (*CreateCoinConfigResponse, error) {
	out := new(CreateCoinConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCoinConfig(ctx context.Context, in *UpdateCoinConfigRequest, opts ...grpc.CallOption) (*UpdateCoinConfigResponse, error) {
	out := new(UpdateCoinConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinConfig(ctx context.Context, in *GetCoinConfigRequest, opts ...grpc.CallOption) (*GetCoinConfigResponse, error) {
	out := new(GetCoinConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinConfigs(ctx context.Context, in *GetCoinConfigsRequest, opts ...grpc.CallOption) (*GetCoinConfigsResponse, error) {
	out := new(GetCoinConfigsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCoinConfigConds(ctx context.Context, in *ExistCoinConfigCondsRequest, opts ...grpc.CallOption) (*ExistCoinConfigCondsResponse, error) {
	out := new(ExistCoinConfigCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCoinConfigConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCoinConfig(ctx context.Context, in *DeleteCoinConfigRequest, opts ...grpc.CallOption) (*DeleteCoinConfigResponse, error) {
	out := new(DeleteCoinConfigResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCoinConfig(context.Context, *CreateCoinConfigRequest) (*CreateCoinConfigResponse, error)
	UpdateCoinConfig(context.Context, *UpdateCoinConfigRequest) (*UpdateCoinConfigResponse, error)
	GetCoinConfig(context.Context, *GetCoinConfigRequest) (*GetCoinConfigResponse, error)
	GetCoinConfigs(context.Context, *GetCoinConfigsRequest) (*GetCoinConfigsResponse, error)
	ExistCoinConfigConds(context.Context, *ExistCoinConfigCondsRequest) (*ExistCoinConfigCondsResponse, error)
	DeleteCoinConfig(context.Context, *DeleteCoinConfigRequest) (*DeleteCoinConfigResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCoinConfig(context.Context, *CreateCoinConfigRequest) (*CreateCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinConfig not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCoinConfig(context.Context, *UpdateCoinConfigRequest) (*UpdateCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinConfig not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinConfig(context.Context, *GetCoinConfigRequest) (*GetCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinConfig not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinConfigs(context.Context, *GetCoinConfigsRequest) (*GetCoinConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinConfigs not implemented")
}
func (UnimplementedMiddlewareServer) ExistCoinConfigConds(context.Context, *ExistCoinConfigCondsRequest) (*ExistCoinConfigCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoinConfigConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCoinConfig(context.Context, *DeleteCoinConfigRequest) (*DeleteCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoinConfig not implemented")
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

func _Middleware_CreateCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCoinConfig(ctx, req.(*CreateCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCoinConfig(ctx, req.(*UpdateCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinConfig(ctx, req.(*GetCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinConfigs(ctx, req.(*GetCoinConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCoinConfigConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinConfigCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCoinConfigConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCoinConfigConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCoinConfigConds(ctx, req.(*ExistCoinConfigCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCoinConfig(ctx, req.(*DeleteCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.coin.config.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinConfig",
			Handler:    _Middleware_CreateCoinConfig_Handler,
		},
		{
			MethodName: "UpdateCoinConfig",
			Handler:    _Middleware_UpdateCoinConfig_Handler,
		},
		{
			MethodName: "GetCoinConfig",
			Handler:    _Middleware_GetCoinConfig_Handler,
		},
		{
			MethodName: "GetCoinConfigs",
			Handler:    _Middleware_GetCoinConfigs_Handler,
		},
		{
			MethodName: "ExistCoinConfigConds",
			Handler:    _Middleware_ExistCoinConfigConds_Handler,
		},
		{
			MethodName: "DeleteCoinConfig",
			Handler:    _Middleware_DeleteCoinConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coin/config/config.proto",
}