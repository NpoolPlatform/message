// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/coupon/app/scope/scope.proto

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
	Middleware_CreateAppGoodScope_FullMethodName     = "/inspire.middleware.coupon.app.scope.v1.Middleware/CreateAppGoodScope"
	Middleware_DeleteAppGoodScope_FullMethodName     = "/inspire.middleware.coupon.app.scope.v1.Middleware/DeleteAppGoodScope"
	Middleware_GetAppGoodScopes_FullMethodName       = "/inspire.middleware.coupon.app.scope.v1.Middleware/GetAppGoodScopes"
	Middleware_ExistAppGoodScopeConds_FullMethodName = "/inspire.middleware.coupon.app.scope.v1.Middleware/ExistAppGoodScopeConds"
	Middleware_VerifyCouponScope_FullMethodName      = "/inspire.middleware.coupon.app.scope.v1.Middleware/VerifyCouponScope"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
	ExistAppGoodScopeConds(ctx context.Context, in *ExistAppGoodScopeCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodScopeCondsResponse, error)
	VerifyCouponScope(ctx context.Context, in *VerifyCouponScopeRequest, opts ...grpc.CallOption) (*VerifyCouponScopeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error) {
	out := new(CreateAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateAppGoodScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error) {
	out := new(DeleteAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteAppGoodScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAppGoodScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAppGoodScopeConds(ctx context.Context, in *ExistAppGoodScopeCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodScopeCondsResponse, error) {
	out := new(ExistAppGoodScopeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistAppGoodScopeConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) VerifyCouponScope(ctx context.Context, in *VerifyCouponScopeRequest, opts ...grpc.CallOption) (*VerifyCouponScopeResponse, error) {
	out := new(VerifyCouponScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_VerifyCouponScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error)
	GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error)
	ExistAppGoodScopeConds(context.Context, *ExistAppGoodScopeCondsRequest) (*ExistAppGoodScopeCondsResponse, error)
	VerifyCouponScope(context.Context, *VerifyCouponScopeRequest) (*VerifyCouponScopeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateAppGoodScope(context.Context, *CreateAppGoodScopeRequest) (*CreateAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodScope not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAppGoodScope(context.Context, *DeleteAppGoodScopeRequest) (*DeleteAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppGoodScope not implemented")
}
func (UnimplementedMiddlewareServer) GetAppGoodScopes(context.Context, *GetAppGoodScopesRequest) (*GetAppGoodScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScopes not implemented")
}
func (UnimplementedMiddlewareServer) ExistAppGoodScopeConds(context.Context, *ExistAppGoodScopeCondsRequest) (*ExistAppGoodScopeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppGoodScopeConds not implemented")
}
func (UnimplementedMiddlewareServer) VerifyCouponScope(context.Context, *VerifyCouponScopeRequest) (*VerifyCouponScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCouponScope not implemented")
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

func _Middleware_CreateAppGoodScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppGoodScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAppGoodScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateAppGoodScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAppGoodScope(ctx, req.(*CreateAppGoodScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAppGoodScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppGoodScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAppGoodScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteAppGoodScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAppGoodScope(ctx, req.(*DeleteAppGoodScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppGoodScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppGoodScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAppGoodScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAppGoodScopeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppGoodScopeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAppGoodScopeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistAppGoodScopeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAppGoodScopeConds(ctx, req.(*ExistAppGoodScopeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_VerifyCouponScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCouponScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).VerifyCouponScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_VerifyCouponScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).VerifyCouponScope(ctx, req.(*VerifyCouponScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.coupon.app.scope.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppGoodScope",
			Handler:    _Middleware_CreateAppGoodScope_Handler,
		},
		{
			MethodName: "DeleteAppGoodScope",
			Handler:    _Middleware_DeleteAppGoodScope_Handler,
		},
		{
			MethodName: "GetAppGoodScopes",
			Handler:    _Middleware_GetAppGoodScopes_Handler,
		},
		{
			MethodName: "ExistAppGoodScopeConds",
			Handler:    _Middleware_ExistAppGoodScopeConds_Handler,
		},
		{
			MethodName: "VerifyCouponScope",
			Handler:    _Middleware_VerifyCouponScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coupon/app/scope/scope.proto",
}
