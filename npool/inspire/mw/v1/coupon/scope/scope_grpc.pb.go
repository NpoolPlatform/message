// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/coupon/scope/scope.proto

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
=======
>>>>>>> fdf74a356 (update scope)
	Middleware_CreateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/CreateScope"
=======
	Middleware_CreateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/CreateScope"
<<<<<<< HEAD
	Middleware_UpdateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/UpdateScope"
>>>>>>> 0c59f0169 (add ExistScopeConds)
=======
>>>>>>> 5f58823bb (delete update)
=======
	Middleware_CreateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/CreateScope"
=======
	Middleware_CreateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/CreateScope"
<<<<<<< HEAD
	Middleware_UpdateScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/UpdateScope"
>>>>>>> 7b4ce1d7e (add ExistScopeConds)
<<<<<<< HEAD
>>>>>>> 894f7fc93 (add ExistScopeConds)
=======
=======
>>>>>>> 9de0afdf4 (delete update)
>>>>>>> 34d2dcf6a (delete update)
	Middleware_ExistScopeConds_FullMethodName = "/inspire.middleware.coupon.scope.v1.Middleware/ExistScopeConds"
	Middleware_GetScope_FullMethodName        = "/inspire.middleware.coupon.scope.v1.Middleware/GetScope"
	Middleware_GetScopes_FullMethodName       = "/inspire.middleware.coupon.scope.v1.Middleware/GetScopes"
	Middleware_DeleteScope_FullMethodName     = "/inspire.middleware.coupon.scope.v1.Middleware/DeleteScope"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> fdf74a356 (update scope)
=======
>>>>>>> 894f7fc93 (add ExistScopeConds)
=======
	Middleware_CreateScope_FullMethodName = "/inspire.middleware.coupon.scope.v1.Middleware/CreateScope"
	Middleware_UpdateScope_FullMethodName = "/inspire.middleware.coupon.scope.v1.Middleware/UpdateScope"
	Middleware_GetScope_FullMethodName    = "/inspire.middleware.coupon.scope.v1.Middleware/GetScope"
	Middleware_GetScopes_FullMethodName   = "/inspire.middleware.coupon.scope.v1.Middleware/GetScopes"
	Middleware_DeleteScope_FullMethodName = "/inspire.middleware.coupon.scope.v1.Middleware/DeleteScope"
<<<<<<< HEAD
>>>>>>> 517e40d6c (update scope)
=======
>>>>>>> 0c59f0169 (add ExistScopeConds)
=======
>>>>>>> e2ba5c456 (update scope)
<<<<<<< HEAD
>>>>>>> fdf74a356 (update scope)
=======
=======
>>>>>>> 7b4ce1d7e (add ExistScopeConds)
>>>>>>> 894f7fc93 (add ExistScopeConds)
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5f58823bb (delete update)
	ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error)
=======
	UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error)
<<<<<<< HEAD
>>>>>>> 517e40d6c (update scope)
=======
	ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error)
>>>>>>> 0c59f0169 (add ExistScopeConds)
=======
=======
=======
>>>>>>> 9de0afdf4 (delete update)
>>>>>>> 34d2dcf6a (delete update)
	ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error)
=======
	UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error)
<<<<<<< HEAD
>>>>>>> e2ba5c456 (update scope)
<<<<<<< HEAD
>>>>>>> fdf74a356 (update scope)
=======
=======
	ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error)
>>>>>>> 7b4ce1d7e (add ExistScopeConds)
>>>>>>> 894f7fc93 (add ExistScopeConds)
	GetScope(ctx context.Context, in *GetScopeRequest, opts ...grpc.CallOption) (*GetScopeResponse, error)
	GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error)
	DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error) {
	out := new(CreateScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *middlewareClient) ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error) {
	out := new(ExistScopeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistScopeConds_FullMethodName, in, out, opts...)
=======
func (c *middlewareClient) UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error) {
	out := new(UpdateScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateScope_FullMethodName, in, out, opts...)
>>>>>>> 517e40d6c (update scope)
	if err != nil {
		return nil, err
	}
	return out, nil
}

=======
>>>>>>> 5f58823bb (delete update)
=======
>>>>>>> fdf74a356 (update scope)
=======
>>>>>>> 34d2dcf6a (delete update)
func (c *middlewareClient) ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error) {
	out := new(ExistScopeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistScopeConds_FullMethodName, in, out, opts...)
=======
func (c *middlewareClient) UpdateScope(ctx context.Context, in *UpdateScopeRequest, opts ...grpc.CallOption) (*UpdateScopeResponse, error) {
	out := new(UpdateScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateScope_FullMethodName, in, out, opts...)
>>>>>>> e2ba5c456 (update scope)
	if err != nil {
		return nil, err
	}
	return out, nil
}

=======
>>>>>>> 9de0afdf4 (delete update)
func (c *middlewareClient) ExistScopeConds(ctx context.Context, in *ExistScopeCondsRequest, opts ...grpc.CallOption) (*ExistScopeCondsResponse, error) {
	out := new(ExistScopeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistScopeConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetScope(ctx context.Context, in *GetScopeRequest, opts ...grpc.CallOption) (*GetScopeResponse, error) {
	out := new(GetScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_GetScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error) {
	out := new(GetScopesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error) {
	out := new(DeleteScopeResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5f58823bb (delete update)
	ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error)
=======
	UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error)
<<<<<<< HEAD
>>>>>>> 517e40d6c (update scope)
=======
	ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error)
>>>>>>> 0c59f0169 (add ExistScopeConds)
=======
=======
=======
>>>>>>> 9de0afdf4 (delete update)
>>>>>>> 34d2dcf6a (delete update)
	ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error)
=======
	UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error)
<<<<<<< HEAD
>>>>>>> e2ba5c456 (update scope)
<<<<<<< HEAD
>>>>>>> fdf74a356 (update scope)
=======
=======
	ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error)
>>>>>>> 7b4ce1d7e (add ExistScopeConds)
>>>>>>> 894f7fc93 (add ExistScopeConds)
	GetScope(context.Context, *GetScopeRequest) (*GetScopeResponse, error)
	GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error)
	DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScope not implemented")
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (UnimplementedMiddlewareServer) ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistScopeConds not implemented")
=======
func (UnimplementedMiddlewareServer) UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScope not implemented")
>>>>>>> 517e40d6c (update scope)
}
=======
>>>>>>> 5f58823bb (delete update)
=======
>>>>>>> fdf74a356 (update scope)
=======
>>>>>>> 34d2dcf6a (delete update)
func (UnimplementedMiddlewareServer) ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistScopeConds not implemented")
=======
func (UnimplementedMiddlewareServer) UpdateScope(context.Context, *UpdateScopeRequest) (*UpdateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScope not implemented")
>>>>>>> e2ba5c456 (update scope)
}
=======
>>>>>>> 9de0afdf4 (delete update)
func (UnimplementedMiddlewareServer) ExistScopeConds(context.Context, *ExistScopeCondsRequest) (*ExistScopeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistScopeConds not implemented")
}
func (UnimplementedMiddlewareServer) GetScope(context.Context, *GetScopeRequest) (*GetScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScope not implemented")
}
func (UnimplementedMiddlewareServer) GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScopes not implemented")
}
func (UnimplementedMiddlewareServer) DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScope not implemented")
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

func _Middleware_CreateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateScope(ctx, req.(*CreateScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func _Middleware_ExistScopeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistScopeCondsRequest)
=======
func _Middleware_UpdateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScopeRequest)
>>>>>>> 517e40d6c (update scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).ExistScopeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistScopeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistScopeConds(ctx, req.(*ExistScopeCondsRequest))
=======
		return srv.(MiddlewareServer).UpdateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateScope(ctx, req.(*UpdateScopeRequest))
>>>>>>> 517e40d6c (update scope)
	}
	return interceptor(ctx, in, info, handler)
}

=======
>>>>>>> 5f58823bb (delete update)
=======
>>>>>>> fdf74a356 (update scope)
=======
>>>>>>> 34d2dcf6a (delete update)
func _Middleware_ExistScopeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistScopeCondsRequest)
=======
func _Middleware_UpdateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScopeRequest)
>>>>>>> e2ba5c456 (update scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).ExistScopeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistScopeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistScopeConds(ctx, req.(*ExistScopeCondsRequest))
=======
		return srv.(MiddlewareServer).UpdateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateScope(ctx, req.(*UpdateScopeRequest))
>>>>>>> e2ba5c456 (update scope)
	}
	return interceptor(ctx, in, info, handler)
}

=======
>>>>>>> 9de0afdf4 (delete update)
func _Middleware_ExistScopeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistScopeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistScopeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistScopeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistScopeConds(ctx, req.(*ExistScopeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetScope(ctx, req.(*GetScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetScopes(ctx, req.(*GetScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteScope(ctx, req.(*DeleteScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.coupon.scope.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScope",
			Handler:    _Middleware_CreateScope_Handler,
		},
		{
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			MethodName: "ExistScopeConds",
			Handler:    _Middleware_ExistScopeConds_Handler,
=======
			MethodName: "UpdateScope",
			Handler:    _Middleware_UpdateScope_Handler,
>>>>>>> 517e40d6c (update scope)
		},
		{
=======
>>>>>>> 5f58823bb (delete update)
=======
>>>>>>> fdf74a356 (update scope)
=======
>>>>>>> 34d2dcf6a (delete update)
			MethodName: "ExistScopeConds",
			Handler:    _Middleware_ExistScopeConds_Handler,
=======
			MethodName: "UpdateScope",
			Handler:    _Middleware_UpdateScope_Handler,
>>>>>>> e2ba5c456 (update scope)
		},
		{
=======
>>>>>>> 9de0afdf4 (delete update)
			MethodName: "ExistScopeConds",
			Handler:    _Middleware_ExistScopeConds_Handler,
		},
		{
			MethodName: "GetScope",
			Handler:    _Middleware_GetScope_Handler,
		},
		{
			MethodName: "GetScopes",
			Handler:    _Middleware_GetScopes_Handler,
		},
		{
			MethodName: "DeleteScope",
			Handler:    _Middleware_DeleteScope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coupon/scope/scope.proto",
}
