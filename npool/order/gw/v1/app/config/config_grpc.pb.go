// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/gw/v1/app/config/config.proto

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
	Gateway_CreateAppConfig_FullMethodName    = "/order.gateway.app.config.v1.Gateway/CreateAppConfig"
	Gateway_UpdateAppConfig_FullMethodName    = "/order.gateway.app.config.v1.Gateway/UpdateAppConfig"
	Gateway_GetAppConfig_FullMethodName       = "/order.gateway.app.config.v1.Gateway/GetAppConfig"
	Gateway_GetAppConfigs_FullMethodName      = "/order.gateway.app.config.v1.Gateway/GetAppConfigs"
	Gateway_CreateAppAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/CreateAppAppConfig"
	Gateway_UpdateAppAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/UpdateAppAppConfig"
	Gateway_GetAppAppConfigs_FullMethodName   = "/order.gateway.app.config.v1.Gateway/GetAppAppConfigs"
	Gateway_DeleteAppAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/DeleteAppAppConfig"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppConfig(ctx context.Context, in *CreateAppConfigRequest, opts ...grpc.CallOption) (*CreateAppConfigResponse, error)
	UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppConfigResponse, error)
	GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigResponse, error)
	GetAppConfigs(ctx context.Context, in *GetAppConfigsRequest, opts ...grpc.CallOption) (*GetAppConfigsResponse, error)
	// Admin apis
	CreateAppAppConfig(ctx context.Context, in *CreateAppAppConfigRequest, opts ...grpc.CallOption) (*CreateAppAppConfigResponse, error)
	UpdateAppAppConfig(ctx context.Context, in *UpdateAppAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppAppConfigResponse, error)
	GetAppAppConfigs(ctx context.Context, in *GetAppAppConfigsRequest, opts ...grpc.CallOption) (*GetAppAppConfigsResponse, error)
	DeleteAppAppConfig(ctx context.Context, in *DeleteAppAppConfigRequest, opts ...grpc.CallOption) (*DeleteAppAppConfigResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppConfig(ctx context.Context, in *CreateAppConfigRequest, opts ...grpc.CallOption) (*CreateAppConfigResponse, error) {
	out := new(CreateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppConfigResponse, error) {
	out := new(UpdateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigResponse, error) {
	out := new(GetAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppConfigs(ctx context.Context, in *GetAppConfigsRequest, opts ...grpc.CallOption) (*GetAppConfigsResponse, error) {
	out := new(GetAppConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppAppConfig(ctx context.Context, in *CreateAppAppConfigRequest, opts ...grpc.CallOption) (*CreateAppAppConfigResponse, error) {
	out := new(CreateAppAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppAppConfig(ctx context.Context, in *UpdateAppAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppAppConfigResponse, error) {
	out := new(UpdateAppAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppAppConfigs(ctx context.Context, in *GetAppAppConfigsRequest, opts ...grpc.CallOption) (*GetAppAppConfigsResponse, error) {
	out := new(GetAppAppConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppAppConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAppAppConfig(ctx context.Context, in *DeleteAppAppConfigRequest, opts ...grpc.CallOption) (*DeleteAppAppConfigResponse, error) {
	out := new(DeleteAppAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteAppAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppConfig(context.Context, *CreateAppConfigRequest) (*CreateAppConfigResponse, error)
	UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error)
	GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigResponse, error)
	GetAppConfigs(context.Context, *GetAppConfigsRequest) (*GetAppConfigsResponse, error)
	// Admin apis
	CreateAppAppConfig(context.Context, *CreateAppAppConfigRequest) (*CreateAppAppConfigResponse, error)
	UpdateAppAppConfig(context.Context, *UpdateAppAppConfigRequest) (*UpdateAppAppConfigResponse, error)
	GetAppAppConfigs(context.Context, *GetAppAppConfigsRequest) (*GetAppAppConfigsResponse, error)
	DeleteAppAppConfig(context.Context, *DeleteAppAppConfigRequest) (*DeleteAppAppConfigResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAppConfig(context.Context, *CreateAppConfigRequest) (*CreateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppConfigs(context.Context, *GetAppConfigsRequest) (*GetAppConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfigs not implemented")
}
func (UnimplementedGatewayServer) CreateAppAppConfig(context.Context, *CreateAppAppConfigRequest) (*CreateAppAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppAppConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateAppAppConfig(context.Context, *UpdateAppAppConfigRequest) (*UpdateAppAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppAppConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppAppConfigs(context.Context, *GetAppAppConfigsRequest) (*GetAppAppConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppAppConfigs not implemented")
}
func (UnimplementedGatewayServer) DeleteAppAppConfig(context.Context, *DeleteAppAppConfigRequest) (*DeleteAppAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppAppConfig not implemented")
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

func _Gateway_CreateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppConfig(ctx, req.(*CreateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppConfig(ctx, req.(*UpdateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppConfig(ctx, req.(*GetAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppConfigs(ctx, req.(*GetAppConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppAppConfig(ctx, req.(*CreateAppAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppAppConfig(ctx, req.(*UpdateAppAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppAppConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppAppConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppAppConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppAppConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppAppConfigs(ctx, req.(*GetAppAppConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAppAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAppAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteAppAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAppAppConfig(ctx, req.(*DeleteAppAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.app.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppConfig",
			Handler:    _Gateway_CreateAppConfig_Handler,
		},
		{
			MethodName: "UpdateAppConfig",
			Handler:    _Gateway_UpdateAppConfig_Handler,
		},
		{
			MethodName: "GetAppConfig",
			Handler:    _Gateway_GetAppConfig_Handler,
		},
		{
			MethodName: "GetAppConfigs",
			Handler:    _Gateway_GetAppConfigs_Handler,
		},
		{
			MethodName: "CreateAppAppConfig",
			Handler:    _Gateway_CreateAppAppConfig_Handler,
		},
		{
			MethodName: "UpdateAppAppConfig",
			Handler:    _Gateway_UpdateAppAppConfig_Handler,
		},
		{
			MethodName: "GetAppAppConfigs",
			Handler:    _Gateway_GetAppAppConfigs_Handler,
		},
		{
			MethodName: "DeleteAppAppConfig",
			Handler:    _Gateway_DeleteAppAppConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/app/config/config.proto",
}
