// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/billing/gw/v1/user/subscription/subscription.proto

package subscription

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
	Gateway_CreateSubscription_FullMethodName      = "/billing.gateway.user.subscription.v1.Gateway/CreateSubscription"
	Gateway_UpdateSubscription_FullMethodName      = "/billing.gateway.user.subscription.v1.Gateway/UpdateSubscription"
	Gateway_GetSubscription_FullMethodName         = "/billing.gateway.user.subscription.v1.Gateway/GetSubscription"
	Gateway_GetSubscriptions_FullMethodName        = "/billing.gateway.user.subscription.v1.Gateway/GetSubscriptions"
	Gateway_GetSubscriptionsCount_FullMethodName   = "/billing.gateway.user.subscription.v1.Gateway/GetSubscriptionsCount"
	Gateway_AdminDeleteSubscription_FullMethodName = "/billing.gateway.user.subscription.v1.Gateway/AdminDeleteSubscription"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error)
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error)
	GetSubscriptionsCount(ctx context.Context, in *GetSubscriptionsCountRequest, opts ...grpc.CallOption) (*GetSubscriptionsCountResponse, error)
	AdminDeleteSubscription(ctx context.Context, in *AdminDeleteSubscriptionRequest, opts ...grpc.CallOption) (*AdminDeleteSubscriptionResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error) {
	out := new(CreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error) {
	out := new(UpdateSubscriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error) {
	out := new(GetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSubscriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSubscriptionsCount(ctx context.Context, in *GetSubscriptionsCountRequest, opts ...grpc.CallOption) (*GetSubscriptionsCountResponse, error) {
	out := new(GetSubscriptionsCountResponse)
	err := c.cc.Invoke(ctx, Gateway_GetSubscriptionsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteSubscription(ctx context.Context, in *AdminDeleteSubscriptionRequest, opts ...grpc.CallOption) (*AdminDeleteSubscriptionResponse, error) {
	out := new(AdminDeleteSubscriptionResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error)
	UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*UpdateSubscriptionResponse, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error)
	GetSubscriptionsCount(context.Context, *GetSubscriptionsCountRequest) (*GetSubscriptionsCountResponse, error)
	AdminDeleteSubscription(context.Context, *AdminDeleteSubscriptionRequest) (*AdminDeleteSubscriptionResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedGatewayServer) UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*UpdateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscription not implemented")
}
func (UnimplementedGatewayServer) GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedGatewayServer) GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedGatewayServer) GetSubscriptionsCount(context.Context, *GetSubscriptionsCountRequest) (*GetSubscriptionsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionsCount not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteSubscription(context.Context, *AdminDeleteSubscriptionRequest) (*AdminDeleteSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteSubscription not implemented")
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

func _Gateway_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateSubscription(ctx, req.(*UpdateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSubscriptions(ctx, req.(*GetSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSubscriptionsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSubscriptionsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetSubscriptionsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSubscriptionsCount(ctx, req.(*GetSubscriptionsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteSubscription(ctx, req.(*AdminDeleteSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billing.gateway.user.subscription.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscription",
			Handler:    _Gateway_CreateSubscription_Handler,
		},
		{
			MethodName: "UpdateSubscription",
			Handler:    _Gateway_UpdateSubscription_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _Gateway_GetSubscription_Handler,
		},
		{
			MethodName: "GetSubscriptions",
			Handler:    _Gateway_GetSubscriptions_Handler,
		},
		{
			MethodName: "GetSubscriptionsCount",
			Handler:    _Gateway_GetSubscriptionsCount_Handler,
		},
		{
			MethodName: "AdminDeleteSubscription",
			Handler:    _Gateway_AdminDeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/billing/gw/v1/user/subscription/subscription.proto",
}
