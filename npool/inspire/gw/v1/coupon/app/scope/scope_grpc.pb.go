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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
	Gateway_CreateAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope"
	Gateway_DeleteAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope"
=======
	Gateway_CreateAppScope_FullMethodName     = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope"
	Gateway_GetAppScopes_FullMethodName       = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppScopes"
<<<<<<< HEAD
<<<<<<< HEAD
=======
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
>>>>>>> ce3f1742c (delete createappscope)
	Gateway_CreateAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope"
	Gateway_DeleteAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope"
<<<<<<< HEAD
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
=======
=======
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
	Gateway_CreateAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope"
	Gateway_DeleteAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope"
<<<<<<< HEAD
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
	Gateway_CreateAppScope_FullMethodName     = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppScope"
	Gateway_GetAppScopes_FullMethodName       = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppScopes"
	Gateway_CreateAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/CreateAppGoodScope"
	Gateway_DeleteAppGoodScope_FullMethodName = "/inspire.gateway.coupon.app.scope.v1.Gateway/DeleteAppGoodScope"
	Gateway_GetAppGoodScopes_FullMethodName   = "/inspire.gateway.coupon.app.scope.v1.Gateway/GetAppGoodScopes"
>>>>>>> b20bec5db (update scope)
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 808e18b40 (update req)
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
=======
=======
>>>>>>> f7e76ad3e (update req)
>>>>>>> b695b1f4b (update req)
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
=======
	CreateAppScope(ctx context.Context, in *CreateAppScopeRequest, opts ...grpc.CallOption) (*CreateAppScopeResponse, error)
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppGoodScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> ce3f1742c (delete createappscope)
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
<<<<<<< HEAD
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
=======
=======
	GetAppGoodScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
<<<<<<< HEAD
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
	CreateAppScope(ctx context.Context, in *CreateAppScopeRequest, opts ...grpc.CallOption) (*CreateAppScopeResponse, error)
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
>>>>>>> b20bec5db (update scope)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
=======
>>>>>>> b695b1f4b (update req)
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodScopes_FullMethodName, in, out, opts...)
=======
=======
>>>>>>> b20bec5db (update scope)
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
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
=======
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
>>>>>>> 808e18b40 (update req)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodScopes_FullMethodName, in, out, opts...)
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
=======
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
>>>>>>> f7e76ad3e (update req)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodScopes_FullMethodName, in, out, opts...)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
=======
>>>>>>> b20bec5db (update scope)
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
=======
>>>>>>> b20bec5db (update scope)
func (c *gatewayClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
>>>>>>> b20bec5db (update scope)
// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 808e18b40 (update req)
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
=======
=======
>>>>>>> f7e76ad3e (update req)
>>>>>>> b695b1f4b (update req)
	GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error)
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
=======
	CreateAppScope(context.Context, *CreateAppScopeRequest) (*CreateAppScopeResponse, error)
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppGoodScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> ce3f1742c (delete createappscope)
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
<<<<<<< HEAD
	GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error)
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
=======
=======
	GetAppGoodScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
<<<<<<< HEAD
	GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error)
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
	CreateAppScope(context.Context, *CreateAppScopeRequest) (*CreateAppScopeResponse, error)
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
	GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error)
>>>>>>> b20bec5db (update scope)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
=======
>>>>>>> b695b1f4b (update req)
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
=======
=======
>>>>>>> b20bec5db (update scope)
func (UnimplementedGatewayServer) CreateAppScope(context.Context, *CreateAppScopeRequest) (*CreateAppScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppScope not implemented")
}
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
>>>>>>> 808e18b40 (update req)
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
>>>>>>> f7e76ad3e (update req)
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
=======
>>>>>>> b20bec5db (update scope)
}
func (UnimplementedGatewayServer) CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodScope not implemented")
}
func (UnimplementedGatewayServer) DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppGoodScope not implemented")
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
}
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
func (UnimplementedGatewayServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
}
>>>>>>> b20bec5db (update scope)
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e6bdbe852 (delete createappscope)
func _Gateway_GetAppGoodScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodScopesRequest)
<<<<<<< HEAD
=======
func _Gateway_CreateAppScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppScopeRequest)
>>>>>>> 789287843 (update scope)
=======
func _Gateway_CreateAppScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppScopeRequest)
>>>>>>> b20bec5db (update scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
<<<<<<< HEAD
		return srv.(GatewayServer).GetAppGoodScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
=======
=======
>>>>>>> b20bec5db (update scope)
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
<<<<<<< HEAD
=======
func _Gateway_GetAppGoodScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
<<<<<<< HEAD
>>>>>>> ce3f1742c (delete createappscope)
	in := new(GetAppScopesRequest)
=======
>>>>>>> 808e18b40 (update req)
=======
func _Gateway_GetAppGoodScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodScopesRequest)
=======
func _Gateway_CreateAppScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppScopeRequest)
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
	in := new(GetAppScopesRequest)
>>>>>>> b20bec5db (update scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
<<<<<<< HEAD
		return srv.(GatewayServer).GetAppGoodScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
<<<<<<< HEAD
<<<<<<< HEAD
		return srv.(GatewayServer).GetAppScopes(ctx, req.(*GetAppScopesRequest))
>>>>>>> 789287843 (update scope)
=======
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppScopesRequest))
>>>>>>> ce3f1742c (delete createappscope)
=======
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
<<<<<<< HEAD
>>>>>>> 808e18b40 (update req)
=======
=======
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
=======
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
	in := new(GetAppScopesRequest)
=======
>>>>>>> f7e76ad3e (update req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGoodScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
<<<<<<< HEAD
<<<<<<< HEAD
		return srv.(GatewayServer).GetAppScopes(ctx, req.(*GetAppScopesRequest))
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppScopesRequest))
>>>>>>> 80bc5ee31 (delete createappscope)
<<<<<<< HEAD
>>>>>>> e6bdbe852 (delete createappscope)
=======
=======
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
>>>>>>> f7e76ad3e (update req)
>>>>>>> b695b1f4b (update req)
=======
		return srv.(GatewayServer).GetAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppScopes(ctx, req.(*GetAppScopesRequest))
>>>>>>> b20bec5db (update scope)
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
=======
>>>>>>> b20bec5db (update scope)
func _Gateway_GetAppGoodScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGoodScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
>>>>>>> b20bec5db (update scope)
// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coupon.app.scope.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> e6bdbe852 (delete createappscope)
			MethodName: "GetAppGoodScopes",
			Handler:    _Gateway_GetAppGoodScopes_Handler,
=======
=======
>>>>>>> b20bec5db (update scope)
			MethodName: "CreateAppScope",
			Handler:    _Gateway_CreateAppScope_Handler,
		},
		{
			MethodName: "GetAppScopes",
			Handler:    _Gateway_GetAppScopes_Handler,
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
			MethodName: "GetAppGoodScopes",
			Handler:    _Gateway_GetAppGoodScopes_Handler,
>>>>>>> ce3f1742c (delete createappscope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
			MethodName: "GetAppGoodScopes",
			Handler:    _Gateway_GetAppGoodScopes_Handler,
>>>>>>> 80bc5ee31 (delete createappscope)
>>>>>>> e6bdbe852 (delete createappscope)
=======
>>>>>>> b20bec5db (update scope)
		},
		{
			MethodName: "CreateAppGoodScope",
			Handler:    _Gateway_CreateAppGoodScope_Handler,
		},
		{
			MethodName: "DeleteAppGoodScope",
			Handler:    _Gateway_DeleteAppGoodScope_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6712e122f (update scope)
=======
=======
>>>>>>> b20bec5db (update scope)
		{
			MethodName: "GetAppGoodScopes",
			Handler:    _Gateway_GetAppGoodScopes_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> 789287843 (update scope)
=======
>>>>>>> 60afcdcd5 (update scope)
=======
>>>>>>> 569d52611 (update scope)
<<<<<<< HEAD
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 909af8f64 (update scope)
>>>>>>> 6712e122f (update scope)
=======
>>>>>>> b20bec5db (update scope)
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coupon/app/scope/scope.proto",
}
