// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/authing/auth/auth.proto

package auth

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
	CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error)
	UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error)
	ExistAuth(ctx context.Context, in *ExistAuthRequest, opts ...grpc.CallOption) (*ExistAuthResponse, error)
	GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error)
	GetAuths(ctx context.Context, in *GetAuthsRequest, opts ...grpc.CallOption) (*GetAuthsResponse, error)
	DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error) {
	out := new(CreateAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/CreateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error) {
	out := new(UpdateAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/UpdateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAuth(ctx context.Context, in *ExistAuthRequest, opts ...grpc.CallOption) (*ExistAuthResponse, error) {
	out := new(ExistAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/ExistAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error) {
	out := new(GetAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/GetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAuths(ctx context.Context, in *GetAuthsRequest, opts ...grpc.CallOption) (*GetAuthsResponse, error) {
	out := new(GetAuthsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/GetAuths", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	out := new(DeleteAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error)
	UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error)
	ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error)
	GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error)
	GetAuths(context.Context, *GetAuthsRequest) (*GetAuthsResponse, error)
	DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuth not implemented")
}
func (UnimplementedMiddlewareServer) UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuth not implemented")
}
func (UnimplementedMiddlewareServer) ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAuth not implemented")
}
func (UnimplementedMiddlewareServer) GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedMiddlewareServer) GetAuths(context.Context, *GetAuthsRequest) (*GetAuthsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuths not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
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

func _Middleware_CreateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/CreateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAuth(ctx, req.(*CreateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/UpdateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateAuth(ctx, req.(*UpdateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/ExistAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAuth(ctx, req.(*ExistAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAuth(ctx, req.(*GetAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAuths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAuths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAuths(ctx, req.(*GetAuthsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAuth(ctx, req.(*DeleteAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.authing.auth.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuth",
			Handler:    _Middleware_CreateAuth_Handler,
		},
		{
			MethodName: "UpdateAuth",
			Handler:    _Middleware_UpdateAuth_Handler,
		},
		{
			MethodName: "ExistAuth",
			Handler:    _Middleware_ExistAuth_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _Middleware_GetAuth_Handler,
		},
		{
			MethodName: "GetAuths",
			Handler:    _Middleware_GetAuths_Handler,
		},
		{
			MethodName: "DeleteAuth",
			Handler:    _Middleware_DeleteAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/authing/auth/auth.proto",
}
