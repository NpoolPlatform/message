// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/billing/mw/v1/subscription/subscription.proto

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
	Middleware_GetSubscriptions_FullMethodName       = "/billing.middleware.subscription.v1.Middleware/GetSubscriptions"
	Middleware_GetSubscriptionsCount_FullMethodName  = "/billing.middleware.subscription.v1.Middleware/GetSubscriptionsCount"
	Middleware_GetSubscription_FullMethodName        = "/billing.middleware.subscription.v1.Middleware/GetSubscription"
	Middleware_ExistSubscriptionConds_FullMethodName = "/billing.middleware.subscription.v1.Middleware/ExistSubscriptionConds"
	Middleware_DeleteSubscriptions_FullMethodName    = "/billing.middleware.subscription.v1.Middleware/DeleteSubscriptions"
	Middleware_CreateSubscription_FullMethodName     = "/billing.middleware.subscription.v1.Middleware/CreateSubscription"
	Middleware_DeleteSubscription_FullMethodName     = "/billing.middleware.subscription.v1.Middleware/DeleteSubscription"
	Middleware_UpdateSubscription_FullMethodName     = "/billing.middleware.subscription.v1.Middleware/UpdateSubscription"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error)
	GetSubscriptionsCount(ctx context.Context, in *GetSubscriptionsCountRequest, opts ...grpc.CallOption) (*GetSubscriptionsCountResponse, error)
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	ExistSubscriptionConds(ctx context.Context, in *ExistSubscriptionCondsRequest, opts ...grpc.CallOption) (*ExistSubscriptionCondsResponse, error)
	DeleteSubscriptions(ctx context.Context, in *DeleteSubscriptionsRequest, opts ...grpc.CallOption) (*DeleteSubscriptionsResponse, error)
	CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error)
	DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetSubscriptions(ctx context.Context, in *GetSubscriptionsRequest, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error) {
	out := new(GetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubscriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSubscriptionsCount(ctx context.Context, in *GetSubscriptionsCountRequest, opts ...grpc.CallOption) (*GetSubscriptionsCountResponse, error) {
	out := new(GetSubscriptionsCountResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubscriptionsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistSubscriptionConds(ctx context.Context, in *ExistSubscriptionCondsRequest, opts ...grpc.CallOption) (*ExistSubscriptionCondsResponse, error) {
	out := new(ExistSubscriptionCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistSubscriptionConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteSubscriptions(ctx context.Context, in *DeleteSubscriptionsRequest, opts ...grpc.CallOption) (*DeleteSubscriptionsResponse, error) {
	out := new(DeleteSubscriptionsResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteSubscriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateSubscription(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*CreateSubscriptionResponse, error) {
	out := new(CreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*DeleteSubscriptionResponse, error) {
	out := new(DeleteSubscriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateSubscription(ctx context.Context, in *UpdateSubscriptionRequest, opts ...grpc.CallOption) (*UpdateSubscriptionResponse, error) {
	out := new(UpdateSubscriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error)
	GetSubscriptionsCount(context.Context, *GetSubscriptionsCountRequest) (*GetSubscriptionsCountResponse, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	ExistSubscriptionConds(context.Context, *ExistSubscriptionCondsRequest) (*ExistSubscriptionCondsResponse, error)
	DeleteSubscriptions(context.Context, *DeleteSubscriptionsRequest) (*DeleteSubscriptionsResponse, error)
	CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error)
	DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error)
	UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*UpdateSubscriptionResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetSubscriptions(context.Context, *GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedMiddlewareServer) GetSubscriptionsCount(context.Context, *GetSubscriptionsCountRequest) (*GetSubscriptionsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionsCount not implemented")
}
func (UnimplementedMiddlewareServer) GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedMiddlewareServer) ExistSubscriptionConds(context.Context, *ExistSubscriptionCondsRequest) (*ExistSubscriptionCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSubscriptionConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteSubscriptions(context.Context, *DeleteSubscriptionsRequest) (*DeleteSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscriptions not implemented")
}
func (UnimplementedMiddlewareServer) CreateSubscription(context.Context, *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedMiddlewareServer) DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedMiddlewareServer) UpdateSubscription(context.Context, *UpdateSubscriptionRequest) (*UpdateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscription not implemented")
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

func _Middleware_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubscriptions(ctx, req.(*GetSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSubscriptionsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubscriptionsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubscriptionsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubscriptionsCount(ctx, req.(*GetSubscriptionsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistSubscriptionConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSubscriptionCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistSubscriptionConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistSubscriptionConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistSubscriptionConds(ctx, req.(*ExistSubscriptionCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteSubscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteSubscriptions(ctx, req.(*DeleteSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateSubscription(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteSubscription(ctx, req.(*DeleteSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateSubscription(ctx, req.(*UpdateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billing.middleware.subscription.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriptions",
			Handler:    _Middleware_GetSubscriptions_Handler,
		},
		{
			MethodName: "GetSubscriptionsCount",
			Handler:    _Middleware_GetSubscriptionsCount_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _Middleware_GetSubscription_Handler,
		},
		{
			MethodName: "ExistSubscriptionConds",
			Handler:    _Middleware_ExistSubscriptionConds_Handler,
		},
		{
			MethodName: "DeleteSubscriptions",
			Handler:    _Middleware_DeleteSubscriptions_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _Middleware_CreateSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _Middleware_DeleteSubscription_Handler,
		},
		{
			MethodName: "UpdateSubscription",
			Handler:    _Middleware_UpdateSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/billing/mw/v1/subscription/subscription.proto",
}
