// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/coupon/app/scope/scope.proto

package scope

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
	Gateway_CreateAppScope_FullMethodName     = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope"
	Gateway_GetAppScopes_FullMethodName       = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppScopes"
	Gateway_CreateAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope"
	Gateway_DeleteAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppScope(ctx context.Context, in *CreateAppScopeRequest, opts ...grpc.CallOption) (*CreateAppScopeResponse, error)
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppScope(ctx context.Context, in *CreateAppScopeRequest, opts ...grpc.CallOption) (*CreateAppScopeResponse, error) {
	out := new(CreateAppScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error) {
	out := new(CreateAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppGoodScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error) {
	out := new(DeleteAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteAppGoodScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppScope(context.Context, *CreateAppScopeRequest) (*CreateAppScopeResponse, error)
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAppScope(context.Context, *CreateAppScopeRequest) (*CreateAppScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppScope not implemented")
}
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
func (UnimplementedGatewayServer) CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodScope not implemented")
}
func (UnimplementedGatewayServer) DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppGoodScope not implemented")
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

func _Gateway_CreateAppScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppScope(ctx, req.(*CreateAppScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppScopes(ctx, req.(*GetAppScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppGoodScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppGoodScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppGoodScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppGoodScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppGoodScope(ctx, req.(*CreateAppGoodScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAppGoodScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppGoodScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAppGoodScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteAppGoodScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAppGoodScope(ctx, req.(*DeleteAppGoodScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coupon.app.scope.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppScope",
			Handler:    _Gateway_CreateAppScope_Handler,
		},
		{
			MethodName: "GetAppScopes",
			Handler:    _Gateway_GetAppScopes_Handler,
		},
		{
			MethodName: "CreateAppGoodScope",
			Handler:    _Gateway_CreateAppGoodScope_Handler,
		},
		{
			MethodName: "DeleteAppGoodScope",
			Handler:    _Gateway_DeleteAppGoodScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coupon/app/scope/scope.proto",
}
