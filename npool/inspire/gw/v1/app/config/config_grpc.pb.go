// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/app/config/config.proto

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
	Gateway_CreateAppConfig_FullMethodName  = "/inspire.gateway.app.config.v1.Gateway/CreateAppConfig"
	Gateway_CreateNAppConfig_FullMethodName = "/inspire.gateway.app.config.v1.Gateway/CreateNAppConfig"
	Gateway_UpdateAppConfig_FullMethodName  = "/inspire.gateway.app.config.v1.Gateway/UpdateAppConfig"
	Gateway_GetAppConfigs_FullMethodName    = "/inspire.gateway.app.config.v1.Gateway/GetAppConfigs"
	Gateway_GetNAppConfigs_FullMethodName   = "/inspire.gateway.app.config.v1.Gateway/GetNAppConfigs"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppConfig(ctx context.Context, in *CreateAppConfigRequest, opts ...grpc.CallOption) (*CreateAppConfigResponse, error)
	CreateNAppConfig(ctx context.Context, in *CreateNAppConfigRequest, opts ...grpc.CallOption) (*CreateNAppConfigResponse, error)
	UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppConfigResponse, error)
	GetAppConfigs(ctx context.Context, in *GetAppConfigsRequest, opts ...grpc.CallOption) (*GetAppConfigsResponse, error)
	GetNAppConfigs(ctx context.Context, in *GetNAppConfigsRequest, opts ...grpc.CallOption) (*GetNAppConfigsResponse, error)
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

func (c *gatewayClient) CreateNAppConfig(ctx context.Context, in *CreateNAppConfigRequest, opts ...grpc.CallOption) (*CreateNAppConfigResponse, error) {
	out := new(CreateNAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateNAppConfig_FullMethodName, in, out, opts...)
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

func (c *gatewayClient) GetAppConfigs(ctx context.Context, in *GetAppConfigsRequest, opts ...grpc.CallOption) (*GetAppConfigsResponse, error) {
	out := new(GetAppConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppConfigs(ctx context.Context, in *GetNAppConfigsRequest, opts ...grpc.CallOption) (*GetNAppConfigsResponse, error) {
	out := new(GetNAppConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetNAppConfigs_FullMethodName, in, out, opts...)
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
	CreateNAppConfig(context.Context, *CreateNAppConfigRequest) (*CreateNAppConfigResponse, error)
	UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error)
	GetAppConfigs(context.Context, *GetAppConfigsRequest) (*GetAppConfigsResponse, error)
	GetNAppConfigs(context.Context, *GetNAppConfigsRequest) (*GetNAppConfigsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAppConfig(context.Context, *CreateAppConfigRequest) (*CreateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppConfig not implemented")
}
func (UnimplementedGatewayServer) CreateNAppConfig(context.Context, *CreateNAppConfigRequest) (*CreateNAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNAppConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppConfigs(context.Context, *GetAppConfigsRequest) (*GetAppConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfigs not implemented")
}
func (UnimplementedGatewayServer) GetNAppConfigs(context.Context, *GetNAppConfigsRequest) (*GetNAppConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppConfigs not implemented")
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

func _Gateway_CreateNAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateNAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNAppConfig(ctx, req.(*CreateNAppConfigRequest))
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

func _Gateway_GetNAppConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppConfigs(ctx, req.(*GetNAppConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.app.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppConfig",
			Handler:    _Gateway_CreateAppConfig_Handler,
		},
		{
			MethodName: "CreateNAppConfig",
			Handler:    _Gateway_CreateNAppConfig_Handler,
		},
		{
			MethodName: "UpdateAppConfig",
			Handler:    _Gateway_UpdateAppConfig_Handler,
		},
		{
			MethodName: "GetAppConfigs",
			Handler:    _Gateway_GetAppConfigs_Handler,
		},
		{
			MethodName: "GetNAppConfigs",
			Handler:    _Gateway_GetNAppConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/app/config/config.proto",
}
