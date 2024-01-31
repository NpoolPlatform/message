// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/gw/v1/simulate/config/config.proto

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
	Gateway_CreateSimulateConfig_FullMethodName    = "/order.gateway.simulate.config.v1.Gateway/CreateSimulateConfig"
	Gateway_UpdateSimulateConfig_FullMethodName    = "/order.gateway.simulate.config.v1.Gateway/UpdateSimulateConfig"
	Gateway_GetSimulateConfig_FullMethodName       = "/order.gateway.simulate.config.v1.Gateway/GetSimulateConfig"
	Gateway_GetSimulateConfigs_FullMethodName      = "/order.gateway.simulate.config.v1.Gateway/GetSimulateConfigs"
	Gateway_CreateAppSimulateConfig_FullMethodName = "/order.gateway.simulate.config.v1.Gateway/CreateAppSimulateConfig"
	Gateway_UpdateAppSimulateConfig_FullMethodName = "/order.gateway.simulate.config.v1.Gateway/UpdateAppSimulateConfig"
	Gateway_GetAppSimulateConfigs_FullMethodName   = "/order.gateway.simulate.config.v1.Gateway/GetAppSimulateConfigs"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateSimulateConfig(ctx context.Context, in *CreateSimulateConfigRequest, opts ...grpc.CallOption) (*CreateSimulateConfigResponse, error)
	UpdateSimulateConfig(ctx context.Context, in *UpdateSimulateConfigRequest, opts ...grpc.CallOption) (*UpdateSimulateConfigResponse, error)
	GetSimulateConfig(ctx context.Context, in *GetSimulateConfigRequest, opts ...grpc.CallOption) (*GetSimulateConfigResponse, error)
	GetSimulateConfigs(ctx context.Context, in *GetSimulateConfigsRequest, opts ...grpc.CallOption) (*GetSimulateConfigsResponse, error)
	// Admin apis
	CreateAppSimulateConfig(ctx context.Context, in *CreateAppSimulateConfigRequest, opts ...grpc.CallOption) (*CreateAppSimulateConfigResponse, error)
	UpdateAppSimulateConfig(ctx context.Context, in *UpdateAppSimulateConfigRequest, opts ...grpc.CallOption) (*UpdateAppSimulateConfigResponse, error)
	GetAppSimulateConfigs(ctx context.Context, in *GetAppSimulateConfigsRequest, opts ...grpc.CallOption) (*GetAppSimulateConfigsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateSimulateConfig(ctx context.Context, in *CreateSimulateConfigRequest, opts ...grpc.CallOption) (*CreateSimulateConfigResponse, error) {
	out := new(CreateSimulateConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateSimulateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateSimulateConfig(ctx context.Context, in *UpdateSimulateConfigRequest, opts ...grpc.CallOption) (*UpdateSimulateConfigResponse, error) {
	out := new(UpdateSimulateConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateSimulateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSimulateConfig(ctx context.Context, in *GetSimulateConfigRequest, opts ...grpc.CallOption) (*GetSimulateConfigResponse, error) {
	out := new(GetSimulateConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSimulateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSimulateConfigs(ctx context.Context, in *GetSimulateConfigsRequest, opts ...grpc.CallOption) (*GetSimulateConfigsResponse, error) {
	out := new(GetSimulateConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSimulateConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppSimulateConfig(ctx context.Context, in *CreateAppSimulateConfigRequest, opts ...grpc.CallOption) (*CreateAppSimulateConfigResponse, error) {
	out := new(CreateAppSimulateConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppSimulateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppSimulateConfig(ctx context.Context, in *UpdateAppSimulateConfigRequest, opts ...grpc.CallOption) (*UpdateAppSimulateConfigResponse, error) {
	out := new(UpdateAppSimulateConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppSimulateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppSimulateConfigs(ctx context.Context, in *GetAppSimulateConfigsRequest, opts ...grpc.CallOption) (*GetAppSimulateConfigsResponse, error) {
	out := new(GetAppSimulateConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppSimulateConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateSimulateConfig(context.Context, *CreateSimulateConfigRequest) (*CreateSimulateConfigResponse, error)
	UpdateSimulateConfig(context.Context, *UpdateSimulateConfigRequest) (*UpdateSimulateConfigResponse, error)
	GetSimulateConfig(context.Context, *GetSimulateConfigRequest) (*GetSimulateConfigResponse, error)
	GetSimulateConfigs(context.Context, *GetSimulateConfigsRequest) (*GetSimulateConfigsResponse, error)
	// Admin apis
	CreateAppSimulateConfig(context.Context, *CreateAppSimulateConfigRequest) (*CreateAppSimulateConfigResponse, error)
	UpdateAppSimulateConfig(context.Context, *UpdateAppSimulateConfigRequest) (*UpdateAppSimulateConfigResponse, error)
	GetAppSimulateConfigs(context.Context, *GetAppSimulateConfigsRequest) (*GetAppSimulateConfigsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateSimulateConfig(context.Context, *CreateSimulateConfigRequest) (*CreateSimulateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSimulateConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateSimulateConfig(context.Context, *UpdateSimulateConfigRequest) (*UpdateSimulateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSimulateConfig not implemented")
}
func (UnimplementedGatewayServer) GetSimulateConfig(context.Context, *GetSimulateConfigRequest) (*GetSimulateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulateConfig not implemented")
}
func (UnimplementedGatewayServer) GetSimulateConfigs(context.Context, *GetSimulateConfigsRequest) (*GetSimulateConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulateConfigs not implemented")
}
func (UnimplementedGatewayServer) CreateAppSimulateConfig(context.Context, *CreateAppSimulateConfigRequest) (*CreateAppSimulateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppSimulateConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateAppSimulateConfig(context.Context, *UpdateAppSimulateConfigRequest) (*UpdateAppSimulateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppSimulateConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppSimulateConfigs(context.Context, *GetAppSimulateConfigsRequest) (*GetAppSimulateConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSimulateConfigs not implemented")
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

func _Gateway_CreateSimulateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSimulateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateSimulateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateSimulateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateSimulateConfig(ctx, req.(*CreateSimulateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateSimulateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSimulateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateSimulateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateSimulateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateSimulateConfig(ctx, req.(*UpdateSimulateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSimulateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimulateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSimulateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSimulateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSimulateConfig(ctx, req.(*GetSimulateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSimulateConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimulateConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSimulateConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSimulateConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSimulateConfigs(ctx, req.(*GetSimulateConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppSimulateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppSimulateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppSimulateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppSimulateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppSimulateConfig(ctx, req.(*CreateAppSimulateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppSimulateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppSimulateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppSimulateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppSimulateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppSimulateConfig(ctx, req.(*UpdateAppSimulateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppSimulateConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSimulateConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppSimulateConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppSimulateConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppSimulateConfigs(ctx, req.(*GetAppSimulateConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.simulate.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSimulateConfig",
			Handler:    _Gateway_CreateSimulateConfig_Handler,
		},
		{
			MethodName: "UpdateSimulateConfig",
			Handler:    _Gateway_UpdateSimulateConfig_Handler,
		},
		{
			MethodName: "GetSimulateConfig",
			Handler:    _Gateway_GetSimulateConfig_Handler,
		},
		{
			MethodName: "GetSimulateConfigs",
			Handler:    _Gateway_GetSimulateConfigs_Handler,
		},
		{
			MethodName: "CreateAppSimulateConfig",
			Handler:    _Gateway_CreateAppSimulateConfig_Handler,
		},
		{
			MethodName: "UpdateAppSimulateConfig",
			Handler:    _Gateway_UpdateAppSimulateConfig_Handler,
		},
		{
			MethodName: "GetAppSimulateConfigs",
			Handler:    _Gateway_GetAppSimulateConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/simulate/config/config.proto",
}
