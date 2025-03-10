// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error)
	DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error)
	GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error)
	GetAppGoodScope(ctx context.Context, in *GetAppGoodScopeRequest, opts ...grpc.CallOption) (*GetAppGoodScopeResponse, error)
	ExistAppGoodScopeConds(ctx context.Context, in *ExistAppGoodScopeCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodScopeCondsResponse, error)
	VerifyCouponScopes(ctx context.Context, in *VerifyCouponScopesRequest, opts ...grpc.CallOption) (*VerifyCouponScopesResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateAppGoodScope(ctx context.Context, in *CreateAppGoodScopeRequest, opts ...grpc.CallOption) (*CreateAppGoodScopeResponse, error) {
	out := new(CreateAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/CreateAppGoodScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAppGoodScope(ctx context.Context, in *DeleteAppGoodScopeRequest, opts ...grpc.CallOption) (*DeleteAppGoodScopeResponse, error) {
	out := new(DeleteAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/DeleteAppGoodScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppGoodScopes(ctx context.Context, in *GetAppGoodScopesRequest, opts ...grpc.CallOption) (*GetAppGoodScopesResponse, error) {
	out := new(GetAppGoodScopesResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/GetAppGoodScopes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppGoodScope(ctx context.Context, in *GetAppGoodScopeRequest, opts ...grpc.CallOption) (*GetAppGoodScopeResponse, error) {
	out := new(GetAppGoodScopeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/GetAppGoodScope", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAppGoodScopeConds(ctx context.Context, in *ExistAppGoodScopeCondsRequest, opts ...grpc.CallOption) (*ExistAppGoodScopeCondsResponse, error) {
	out := new(ExistAppGoodScopeCondsResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/ExistAppGoodScopeConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) VerifyCouponScopes(ctx context.Context, in *VerifyCouponScopesRequest, opts ...grpc.CallOption) (*VerifyCouponScopesResponse, error) {
	out := new(VerifyCouponScopesResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.coupon.app.scope.v1.Middleware/VerifyCouponScopes", in, out, opts...)
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
	GetAppGoodScope(context.Context, *GetAppGoodScopeRequest) (*GetAppGoodScopeResponse, error)
	ExistAppGoodScopeConds(context.Context, *ExistAppGoodScopeCondsRequest) (*ExistAppGoodScopeCondsResponse, error)
	VerifyCouponScopes(context.Context, *VerifyCouponScopesRequest) (*VerifyCouponScopesResponse, error)
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
func (UnimplementedMiddlewareServer) GetAppGoodScope(context.Context, *GetAppGoodScopeRequest) (*GetAppGoodScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodScope not implemented")
}
func (UnimplementedMiddlewareServer) ExistAppGoodScopeConds(context.Context, *ExistAppGoodScopeCondsRequest) (*ExistAppGoodScopeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppGoodScopeConds not implemented")
}
func (UnimplementedMiddlewareServer) VerifyCouponScopes(context.Context, *VerifyCouponScopesRequest) (*VerifyCouponScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCouponScopes not implemented")
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
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/CreateAppGoodScope",
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
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/DeleteAppGoodScope",
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
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/GetAppGoodScopes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppGoodScopes(ctx, req.(*GetAppGoodScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppGoodScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppGoodScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/GetAppGoodScope",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppGoodScope(ctx, req.(*GetAppGoodScopeRequest))
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
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/ExistAppGoodScopeConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAppGoodScopeConds(ctx, req.(*ExistAppGoodScopeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_VerifyCouponScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCouponScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).VerifyCouponScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.coupon.app.scope.v1.Middleware/VerifyCouponScopes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).VerifyCouponScopes(ctx, req.(*VerifyCouponScopesRequest))
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
			MethodName: "GetAppGoodScope",
			Handler:    _Middleware_GetAppGoodScope_Handler,
		},
		{
			MethodName: "ExistAppGoodScopeConds",
			Handler:    _Middleware_ExistAppGoodScopeConds_Handler,
		},
		{
			MethodName: "VerifyCouponScopes",
			Handler:    _Middleware_VerifyCouponScopes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coupon/app/scope/scope.proto",
}
