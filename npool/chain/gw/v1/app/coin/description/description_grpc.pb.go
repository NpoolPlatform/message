// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/gw/v1/app/coin/description/description.proto

package description

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
	Gateway_CreateCoinDescription_FullMethodName    = "/chain.gateway.app.coin.description.v1.Gateway/CreateCoinDescription"
	Gateway_CreateAppCoinDescription_FullMethodName = "/chain.gateway.app.coin.description.v1.Gateway/CreateAppCoinDescription"
	Gateway_GetCoinDescriptions_FullMethodName      = "/chain.gateway.app.coin.description.v1.Gateway/GetCoinDescriptions"
	Gateway_GetAppCoinDescriptions_FullMethodName   = "/chain.gateway.app.coin.description.v1.Gateway/GetAppCoinDescriptions"
	Gateway_UpdateCoinDescription_FullMethodName    = "/chain.gateway.app.coin.description.v1.Gateway/UpdateCoinDescription"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error)
	CreateAppCoinDescription(ctx context.Context, in *CreateAppCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateAppCoinDescriptionResponse, error)
	GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error)
	GetAppCoinDescriptions(ctx context.Context, in *GetAppCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetAppCoinDescriptionsResponse, error)
	UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error) {
	out := new(CreateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppCoinDescription(ctx context.Context, in *CreateAppCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateAppCoinDescriptionResponse, error) {
	out := new(CreateAppCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error) {
	out := new(GetCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCoinDescriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppCoinDescriptions(ctx context.Context, in *GetAppCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetAppCoinDescriptionsResponse, error) {
	out := new(GetAppCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppCoinDescriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error) {
	out := new(UpdateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error)
	CreateAppCoinDescription(context.Context, *CreateAppCoinDescriptionRequest) (*CreateAppCoinDescriptionResponse, error)
	GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error)
	GetAppCoinDescriptions(context.Context, *GetAppCoinDescriptionsRequest) (*GetAppCoinDescriptionsResponse, error)
	UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinDescription not implemented")
}
func (UnimplementedGatewayServer) CreateAppCoinDescription(context.Context, *CreateAppCoinDescriptionRequest) (*CreateAppCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppCoinDescription not implemented")
}
func (UnimplementedGatewayServer) GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescriptions not implemented")
}
func (UnimplementedGatewayServer) GetAppCoinDescriptions(context.Context, *GetAppCoinDescriptionsRequest) (*GetAppCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppCoinDescriptions not implemented")
}
func (UnimplementedGatewayServer) UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinDescription not implemented")
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

func _Gateway_CreateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateCoinDescription(ctx, req.(*CreateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppCoinDescription(ctx, req.(*CreateAppCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetCoinDescriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCoinDescriptions(ctx, req.(*GetCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppCoinDescriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppCoinDescriptions(ctx, req.(*GetAppCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateCoinDescription(ctx, req.(*UpdateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.app.coin.description.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinDescription",
			Handler:    _Gateway_CreateCoinDescription_Handler,
		},
		{
			MethodName: "CreateAppCoinDescription",
			Handler:    _Gateway_CreateAppCoinDescription_Handler,
		},
		{
			MethodName: "GetCoinDescriptions",
			Handler:    _Gateway_GetCoinDescriptions_Handler,
		},
		{
			MethodName: "GetAppCoinDescriptions",
			Handler:    _Gateway_GetAppCoinDescriptions_Handler,
		},
		{
			MethodName: "UpdateCoinDescription",
			Handler:    _Gateway_UpdateCoinDescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/gw/v1/app/coin/description/description.proto",
}
