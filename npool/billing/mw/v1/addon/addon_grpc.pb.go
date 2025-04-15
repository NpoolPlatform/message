// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/billing/mw/v1/addon/addon.proto

package addon

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
	Middleware_GetAddons_FullMethodName       = "/billing.middleware.addon.v1.Middleware/GetAddons"
	Middleware_GetAddonsCount_FullMethodName  = "/billing.middleware.addon.v1.Middleware/GetAddonsCount"
	Middleware_GetAddon_FullMethodName        = "/billing.middleware.addon.v1.Middleware/GetAddon"
	Middleware_ExistAddonConds_FullMethodName = "/billing.middleware.addon.v1.Middleware/ExistAddonConds"
	Middleware_CreateAddon_FullMethodName     = "/billing.middleware.addon.v1.Middleware/CreateAddon"
	Middleware_DeleteAddon_FullMethodName     = "/billing.middleware.addon.v1.Middleware/DeleteAddon"
	Middleware_UpdateAddon_FullMethodName     = "/billing.middleware.addon.v1.Middleware/UpdateAddon"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetAddons(ctx context.Context, in *GetAddonsRequest, opts ...grpc.CallOption) (*GetAddonsResponse, error)
	GetAddonsCount(ctx context.Context, in *GetAddonsCountRequest, opts ...grpc.CallOption) (*GetAddonsCountResponse, error)
	GetAddon(ctx context.Context, in *GetAddonRequest, opts ...grpc.CallOption) (*GetAddonResponse, error)
	ExistAddonConds(ctx context.Context, in *ExistAddonCondsRequest, opts ...grpc.CallOption) (*ExistAddonCondsResponse, error)
	CreateAddon(ctx context.Context, in *CreateAddonRequest, opts ...grpc.CallOption) (*CreateAddonResponse, error)
	DeleteAddon(ctx context.Context, in *DeleteAddonRequest, opts ...grpc.CallOption) (*DeleteAddonResponse, error)
	UpdateAddon(ctx context.Context, in *UpdateAddonRequest, opts ...grpc.CallOption) (*UpdateAddonResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetAddons(ctx context.Context, in *GetAddonsRequest, opts ...grpc.CallOption) (*GetAddonsResponse, error) {
	out := new(GetAddonsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAddons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAddonsCount(ctx context.Context, in *GetAddonsCountRequest, opts ...grpc.CallOption) (*GetAddonsCountResponse, error) {
	out := new(GetAddonsCountResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAddonsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAddon(ctx context.Context, in *GetAddonRequest, opts ...grpc.CallOption) (*GetAddonResponse, error) {
	out := new(GetAddonResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAddon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAddonConds(ctx context.Context, in *ExistAddonCondsRequest, opts ...grpc.CallOption) (*ExistAddonCondsResponse, error) {
	out := new(ExistAddonCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistAddonConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateAddon(ctx context.Context, in *CreateAddonRequest, opts ...grpc.CallOption) (*CreateAddonResponse, error) {
	out := new(CreateAddonResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateAddon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAddon(ctx context.Context, in *DeleteAddonRequest, opts ...grpc.CallOption) (*DeleteAddonResponse, error) {
	out := new(DeleteAddonResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteAddon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateAddon(ctx context.Context, in *UpdateAddonRequest, opts ...grpc.CallOption) (*UpdateAddonResponse, error) {
	out := new(UpdateAddonResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateAddon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetAddons(context.Context, *GetAddonsRequest) (*GetAddonsResponse, error)
	GetAddonsCount(context.Context, *GetAddonsCountRequest) (*GetAddonsCountResponse, error)
	GetAddon(context.Context, *GetAddonRequest) (*GetAddonResponse, error)
	ExistAddonConds(context.Context, *ExistAddonCondsRequest) (*ExistAddonCondsResponse, error)
	CreateAddon(context.Context, *CreateAddonRequest) (*CreateAddonResponse, error)
	DeleteAddon(context.Context, *DeleteAddonRequest) (*DeleteAddonResponse, error)
	UpdateAddon(context.Context, *UpdateAddonRequest) (*UpdateAddonResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetAddons(context.Context, *GetAddonsRequest) (*GetAddonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddons not implemented")
}
func (UnimplementedMiddlewareServer) GetAddonsCount(context.Context, *GetAddonsCountRequest) (*GetAddonsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddonsCount not implemented")
}
func (UnimplementedMiddlewareServer) GetAddon(context.Context, *GetAddonRequest) (*GetAddonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddon not implemented")
}
func (UnimplementedMiddlewareServer) ExistAddonConds(context.Context, *ExistAddonCondsRequest) (*ExistAddonCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAddonConds not implemented")
}
func (UnimplementedMiddlewareServer) CreateAddon(context.Context, *CreateAddonRequest) (*CreateAddonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddon not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAddon(context.Context, *DeleteAddonRequest) (*DeleteAddonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddon not implemented")
}
func (UnimplementedMiddlewareServer) UpdateAddon(context.Context, *UpdateAddonRequest) (*UpdateAddonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddon not implemented")
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

func _Middleware_GetAddons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAddons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAddons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAddons(ctx, req.(*GetAddonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAddonsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddonsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAddonsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAddonsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAddonsCount(ctx, req.(*GetAddonsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAddon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAddon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAddon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAddon(ctx, req.(*GetAddonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAddonConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAddonCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAddonConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistAddonConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAddonConds(ctx, req.(*ExistAddonCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateAddon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAddon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateAddon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAddon(ctx, req.(*CreateAddonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAddon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAddon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteAddon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAddon(ctx, req.(*DeleteAddonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateAddon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateAddon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateAddon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateAddon(ctx, req.(*UpdateAddonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billing.middleware.addon.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddons",
			Handler:    _Middleware_GetAddons_Handler,
		},
		{
			MethodName: "GetAddonsCount",
			Handler:    _Middleware_GetAddonsCount_Handler,
		},
		{
			MethodName: "GetAddon",
			Handler:    _Middleware_GetAddon_Handler,
		},
		{
			MethodName: "ExistAddonConds",
			Handler:    _Middleware_ExistAddonConds_Handler,
		},
		{
			MethodName: "CreateAddon",
			Handler:    _Middleware_CreateAddon_Handler,
		},
		{
			MethodName: "DeleteAddon",
			Handler:    _Middleware_DeleteAddon_Handler,
		},
		{
			MethodName: "UpdateAddon",
			Handler:    _Middleware_UpdateAddon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/billing/mw/v1/addon/addon.proto",
}
