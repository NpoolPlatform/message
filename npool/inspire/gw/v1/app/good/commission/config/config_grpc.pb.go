// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/app/good/commission/config/config.proto

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
	Gateway_CreateAppGoodCommissionConfig_FullMethodName  = "/inspire.gateway.app.good.commission.config.v1.Gateway/CreateAppGoodCommissionConfig"
	Gateway_CreateNAppGoodCommissionConfig_FullMethodName = "/inspire.gateway.app.good.commission.config.v1.Gateway/CreateNAppGoodCommissionConfig"
	Gateway_UpdateAppGoodCommissionConfig_FullMethodName  = "/inspire.gateway.app.good.commission.config.v1.Gateway/UpdateAppGoodCommissionConfig"
	Gateway_UpdateNAppGoodCommissionConfig_FullMethodName = "/inspire.gateway.app.good.commission.config.v1.Gateway/UpdateNAppGoodCommissionConfig"
	Gateway_GetAppGoodCommissionConfigs_FullMethodName    = "/inspire.gateway.app.good.commission.config.v1.Gateway/GetAppGoodCommissionConfigs"
	Gateway_GetNAppGoodCommissionConfigs_FullMethodName   = "/inspire.gateway.app.good.commission.config.v1.Gateway/GetNAppGoodCommissionConfigs"
	Gateway_CloneAppGoodCommissionConfigs_FullMethodName  = "/inspire.gateway.app.good.commission.config.v1.Gateway/CloneAppGoodCommissionConfigs"
	Gateway_CloneNAppGoodCommissionConfigs_FullMethodName = "/inspire.gateway.app.good.commission.config.v1.Gateway/CloneNAppGoodCommissionConfigs"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error)
	CreateNAppGoodCommissionConfig(ctx context.Context, in *CreateNAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateNAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error)
	UpdateNAppGoodCommissionConfig(ctx context.Context, in *UpdateNAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateNAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error)
	GetNAppGoodCommissionConfigs(ctx context.Context, in *GetNAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetNAppGoodCommissionConfigsResponse, error)
	CloneAppGoodCommissionConfigs(ctx context.Context, in *CloneAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneAppGoodCommissionConfigsResponse, error)
	CloneNAppGoodCommissionConfigs(ctx context.Context, in *CloneNAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneNAppGoodCommissionConfigsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error) {
	out := new(CreateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateNAppGoodCommissionConfig(ctx context.Context, in *CreateNAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateNAppGoodCommissionConfigResponse, error) {
	out := new(CreateNAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateNAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error) {
	out := new(UpdateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateNAppGoodCommissionConfig(ctx context.Context, in *UpdateNAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateNAppGoodCommissionConfigResponse, error) {
	out := new(UpdateNAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateNAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error) {
	out := new(GetAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppGoodCommissionConfigs(ctx context.Context, in *GetNAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetNAppGoodCommissionConfigsResponse, error) {
	out := new(GetNAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetNAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CloneAppGoodCommissionConfigs(ctx context.Context, in *CloneAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneAppGoodCommissionConfigsResponse, error) {
	out := new(CloneAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_CloneAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CloneNAppGoodCommissionConfigs(ctx context.Context, in *CloneNAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*CloneNAppGoodCommissionConfigsResponse, error) {
	out := new(CloneNAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_CloneNAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error)
	CreateNAppGoodCommissionConfig(context.Context, *CreateNAppGoodCommissionConfigRequest) (*CreateNAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error)
	UpdateNAppGoodCommissionConfig(context.Context, *UpdateNAppGoodCommissionConfigRequest) (*UpdateNAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error)
	GetNAppGoodCommissionConfigs(context.Context, *GetNAppGoodCommissionConfigsRequest) (*GetNAppGoodCommissionConfigsResponse, error)
	CloneAppGoodCommissionConfigs(context.Context, *CloneAppGoodCommissionConfigsRequest) (*CloneAppGoodCommissionConfigsResponse, error)
	CloneNAppGoodCommissionConfigs(context.Context, *CloneNAppGoodCommissionConfigsRequest) (*CloneNAppGoodCommissionConfigsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodCommissionConfig not implemented")
}
func (UnimplementedGatewayServer) CreateNAppGoodCommissionConfig(context.Context, *CreateNAppGoodCommissionConfigRequest) (*CreateNAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNAppGoodCommissionConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppGoodCommissionConfig not implemented")
}
func (UnimplementedGatewayServer) UpdateNAppGoodCommissionConfig(context.Context, *UpdateNAppGoodCommissionConfigRequest) (*UpdateNAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNAppGoodCommissionConfig not implemented")
}
func (UnimplementedGatewayServer) GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodCommissionConfigs not implemented")
}
func (UnimplementedGatewayServer) GetNAppGoodCommissionConfigs(context.Context, *GetNAppGoodCommissionConfigsRequest) (*GetNAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppGoodCommissionConfigs not implemented")
}
func (UnimplementedGatewayServer) CloneAppGoodCommissionConfigs(context.Context, *CloneAppGoodCommissionConfigsRequest) (*CloneAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneAppGoodCommissionConfigs not implemented")
}
func (UnimplementedGatewayServer) CloneNAppGoodCommissionConfigs(context.Context, *CloneNAppGoodCommissionConfigsRequest) (*CloneNAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneNAppGoodCommissionConfigs not implemented")
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

func _Gateway_CreateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppGoodCommissionConfig(ctx, req.(*CreateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateNAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateNAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNAppGoodCommissionConfig(ctx, req.(*CreateNAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppGoodCommissionConfig(ctx, req.(*UpdateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateNAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateNAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateNAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateNAppGoodCommissionConfig(ctx, req.(*UpdateNAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGoodCommissionConfigs(ctx, req.(*GetAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppGoodCommissionConfigs(ctx, req.(*GetNAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CloneAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CloneAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CloneAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CloneAppGoodCommissionConfigs(ctx, req.(*CloneAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CloneNAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneNAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CloneNAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CloneNAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CloneNAppGoodCommissionConfigs(ctx, req.(*CloneNAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.app.good.commission.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppGoodCommissionConfig",
			Handler:    _Gateway_CreateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "CreateNAppGoodCommissionConfig",
			Handler:    _Gateway_CreateNAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "UpdateAppGoodCommissionConfig",
			Handler:    _Gateway_UpdateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "UpdateNAppGoodCommissionConfig",
			Handler:    _Gateway_UpdateNAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "GetAppGoodCommissionConfigs",
			Handler:    _Gateway_GetAppGoodCommissionConfigs_Handler,
		},
		{
			MethodName: "GetNAppGoodCommissionConfigs",
			Handler:    _Gateway_GetNAppGoodCommissionConfigs_Handler,
		},
		{
			MethodName: "CloneAppGoodCommissionConfigs",
			Handler:    _Gateway_CloneAppGoodCommissionConfigs_Handler,
		},
		{
			MethodName: "CloneNAppGoodCommissionConfigs",
			Handler:    _Gateway_CloneNAppGoodCommissionConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/app/good/commission/config/config.proto",
}
