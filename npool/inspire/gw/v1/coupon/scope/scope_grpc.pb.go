// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/coupon/scope/scope.proto

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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 9589c2455 (update scope)
	Gateway_CreateScope_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName    = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
	Gateway_GetAppScopes_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> feba7689e (delete get app scopes)
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 67de434f0 (delete get app scopes)
=======
>>>>>>> 9589c2455 (update scope)
=======
	Gateway_CreateScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName     = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
	Gateway_GetAppScopes_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes"
	Gateway_GetNAppScopes_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/GetNAppScopes"
>>>>>>> 960f133ce (update scope pb)
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error)
	DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error)
	GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 1fb4844cc (update scope)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 9589c2455 (update scope)
=======
	GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 960f133ce (update scope pb)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error) {
	out := new(CreateScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error) {
	out := new(DeleteScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error) {
	out := new(GetScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 9589c2455 (update scope)
func (c *gatewayClient) GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
>>>>>>> 9589c2455 (update scope)
=======
func (c *gatewayClient) GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetNAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

>>>>>>> 960f133ce (update scope pb)
// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error)
	DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error)
	GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 1fb4844cc (update scope)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 9589c2455 (update scope)
=======
	GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 960f133ce (update scope pb)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScope not implemented")
}
func (UnimplementedGatewayServer) DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScope not implemented")
}
func (UnimplementedGatewayServer) GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScopes not implemented")
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 1fb4844cc (update scope)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 9589c2455 (update scope)
=======
func (UnimplementedGatewayServer) GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppScopes not implemented")
}
>>>>>>> 960f133ce (update scope pb)
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

func _Gateway_CreateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateScope(ctx, req.(*CreateScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteScope(ctx, req.(*DeleteScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetScopes(ctx, req.(*GetScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 9589c2455 (update scope)
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
>>>>>>> 9589c2455 (update scope)
=======
func _Gateway_GetNAppScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppScopes(ctx, req.(*GetAppScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

>>>>>>> 960f133ce (update scope pb)
// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coupon.scope.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScope",
			Handler:    _Gateway_CreateScope_Handler,
		},
		{
			MethodName: "DeleteScope",
			Handler:    _Gateway_DeleteScope_Handler,
		},
		{
			MethodName: "GetScopes",
			Handler:    _Gateway_GetScopes_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 9589c2455 (update scope)
		{
			MethodName: "GetAppScopes",
			Handler:    _Gateway_GetAppScopes_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
=======
>>>>>>> 67de434f0 (delete get app scopes)
=======
>>>>>>> 9589c2455 (update scope)
=======
		{
			MethodName: "GetNAppScopes",
			Handler:    _Gateway_GetNAppScopes_Handler,
		},
>>>>>>> 960f133ce (update scope pb)
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coupon/scope/scope.proto",
}
