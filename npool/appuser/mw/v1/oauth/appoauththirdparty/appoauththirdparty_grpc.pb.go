// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/oauth/appoauththirdparty/appoauththirdparty.proto

package appoauththirdparty

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
	Middleware_CreateOAuthThirdParty_FullMethodName     = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/CreateOAuthThirdParty"
	Middleware_UpdateOAuthThirdParty_FullMethodName     = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/UpdateOAuthThirdParty"
	Middleware_GetOAuthThirdParties_FullMethodName      = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/GetOAuthThirdParties"
	Middleware_GetOAuthThirdParty_FullMethodName        = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/GetOAuthThirdParty"
	Middleware_ExistOAuthThirdParty_FullMethodName      = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/ExistOAuthThirdParty"
	Middleware_ExistOAuthThirdPartyConds_FullMethodName = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/ExistOAuthThirdPartyConds"
	Middleware_DeleteOAuthThirdParty_FullMethodName     = "/appuser.middleware.oauth.appoauththirdparty.v1.Middleware/DeleteOAuthThirdParty"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateOAuthThirdParty(ctx context.Context, in *CreateOAuthThirdPartyRequest, opts ...grpc.CallOption) (*CreateOAuthThirdPartyResponse, error)
	UpdateOAuthThirdParty(ctx context.Context, in *UpdateOAuthThirdPartyRequest, opts ...grpc.CallOption) (*UpdateOAuthThirdPartyResponse, error)
	GetOAuthThirdParties(ctx context.Context, in *GetOAuthThirdPartiesRequest, opts ...grpc.CallOption) (*GetOAuthThirdPartiesResponse, error)
	GetOAuthThirdParty(ctx context.Context, in *GetOAuthThirdPartyRequest, opts ...grpc.CallOption) (*GetOAuthThirdPartyResponse, error)
	ExistOAuthThirdParty(ctx context.Context, in *ExistOAuthThirdPartyRequest, opts ...grpc.CallOption) (*ExistOAuthThirdPartyResponse, error)
	ExistOAuthThirdPartyConds(ctx context.Context, in *ExistOAuthThirdPartyCondsRequest, opts ...grpc.CallOption) (*ExistOAuthThirdPartyCondsResponse, error)
	DeleteOAuthThirdParty(ctx context.Context, in *DeleteOAuthThirdPartyRequest, opts ...grpc.CallOption) (*DeleteOAuthThirdPartyResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateOAuthThirdParty(ctx context.Context, in *CreateOAuthThirdPartyRequest, opts ...grpc.CallOption) (*CreateOAuthThirdPartyResponse, error) {
	out := new(CreateOAuthThirdPartyResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOAuthThirdParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateOAuthThirdParty(ctx context.Context, in *UpdateOAuthThirdPartyRequest, opts ...grpc.CallOption) (*UpdateOAuthThirdPartyResponse, error) {
	out := new(UpdateOAuthThirdPartyResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateOAuthThirdParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOAuthThirdParties(ctx context.Context, in *GetOAuthThirdPartiesRequest, opts ...grpc.CallOption) (*GetOAuthThirdPartiesResponse, error) {
	out := new(GetOAuthThirdPartiesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOAuthThirdParties_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOAuthThirdParty(ctx context.Context, in *GetOAuthThirdPartyRequest, opts ...grpc.CallOption) (*GetOAuthThirdPartyResponse, error) {
	out := new(GetOAuthThirdPartyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOAuthThirdParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOAuthThirdParty(ctx context.Context, in *ExistOAuthThirdPartyRequest, opts ...grpc.CallOption) (*ExistOAuthThirdPartyResponse, error) {
	out := new(ExistOAuthThirdPartyResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOAuthThirdParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOAuthThirdPartyConds(ctx context.Context, in *ExistOAuthThirdPartyCondsRequest, opts ...grpc.CallOption) (*ExistOAuthThirdPartyCondsResponse, error) {
	out := new(ExistOAuthThirdPartyCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOAuthThirdPartyConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteOAuthThirdParty(ctx context.Context, in *DeleteOAuthThirdPartyRequest, opts ...grpc.CallOption) (*DeleteOAuthThirdPartyResponse, error) {
	out := new(DeleteOAuthThirdPartyResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteOAuthThirdParty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateOAuthThirdParty(context.Context, *CreateOAuthThirdPartyRequest) (*CreateOAuthThirdPartyResponse, error)
	UpdateOAuthThirdParty(context.Context, *UpdateOAuthThirdPartyRequest) (*UpdateOAuthThirdPartyResponse, error)
	GetOAuthThirdParties(context.Context, *GetOAuthThirdPartiesRequest) (*GetOAuthThirdPartiesResponse, error)
	GetOAuthThirdParty(context.Context, *GetOAuthThirdPartyRequest) (*GetOAuthThirdPartyResponse, error)
	ExistOAuthThirdParty(context.Context, *ExistOAuthThirdPartyRequest) (*ExistOAuthThirdPartyResponse, error)
	ExistOAuthThirdPartyConds(context.Context, *ExistOAuthThirdPartyCondsRequest) (*ExistOAuthThirdPartyCondsResponse, error)
	DeleteOAuthThirdParty(context.Context, *DeleteOAuthThirdPartyRequest) (*DeleteOAuthThirdPartyResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateOAuthThirdParty(context.Context, *CreateOAuthThirdPartyRequest) (*CreateOAuthThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOAuthThirdParty not implemented")
}
func (UnimplementedMiddlewareServer) UpdateOAuthThirdParty(context.Context, *UpdateOAuthThirdPartyRequest) (*UpdateOAuthThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOAuthThirdParty not implemented")
}
func (UnimplementedMiddlewareServer) GetOAuthThirdParties(context.Context, *GetOAuthThirdPartiesRequest) (*GetOAuthThirdPartiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuthThirdParties not implemented")
}
func (UnimplementedMiddlewareServer) GetOAuthThirdParty(context.Context, *GetOAuthThirdPartyRequest) (*GetOAuthThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOAuthThirdParty not implemented")
}
func (UnimplementedMiddlewareServer) ExistOAuthThirdParty(context.Context, *ExistOAuthThirdPartyRequest) (*ExistOAuthThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOAuthThirdParty not implemented")
}
func (UnimplementedMiddlewareServer) ExistOAuthThirdPartyConds(context.Context, *ExistOAuthThirdPartyCondsRequest) (*ExistOAuthThirdPartyCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOAuthThirdPartyConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteOAuthThirdParty(context.Context, *DeleteOAuthThirdPartyRequest) (*DeleteOAuthThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOAuthThirdParty not implemented")
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

func _Middleware_CreateOAuthThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOAuthThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOAuthThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOAuthThirdParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOAuthThirdParty(ctx, req.(*CreateOAuthThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateOAuthThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOAuthThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateOAuthThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateOAuthThirdParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateOAuthThirdParty(ctx, req.(*UpdateOAuthThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOAuthThirdParties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuthThirdPartiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOAuthThirdParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOAuthThirdParties_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOAuthThirdParties(ctx, req.(*GetOAuthThirdPartiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOAuthThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOAuthThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOAuthThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOAuthThirdParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOAuthThirdParty(ctx, req.(*GetOAuthThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOAuthThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOAuthThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOAuthThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOAuthThirdParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOAuthThirdParty(ctx, req.(*ExistOAuthThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOAuthThirdPartyConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOAuthThirdPartyCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOAuthThirdPartyConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOAuthThirdPartyConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOAuthThirdPartyConds(ctx, req.(*ExistOAuthThirdPartyCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteOAuthThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOAuthThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteOAuthThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteOAuthThirdParty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteOAuthThirdParty(ctx, req.(*DeleteOAuthThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.oauth.appoauththirdparty.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOAuthThirdParty",
			Handler:    _Middleware_CreateOAuthThirdParty_Handler,
		},
		{
			MethodName: "UpdateOAuthThirdParty",
			Handler:    _Middleware_UpdateOAuthThirdParty_Handler,
		},
		{
			MethodName: "GetOAuthThirdParties",
			Handler:    _Middleware_GetOAuthThirdParties_Handler,
		},
		{
			MethodName: "GetOAuthThirdParty",
			Handler:    _Middleware_GetOAuthThirdParty_Handler,
		},
		{
			MethodName: "ExistOAuthThirdParty",
			Handler:    _Middleware_ExistOAuthThirdParty_Handler,
		},
		{
			MethodName: "ExistOAuthThirdPartyConds",
			Handler:    _Middleware_ExistOAuthThirdPartyConds_Handler,
		},
		{
			MethodName: "DeleteOAuthThirdParty",
			Handler:    _Middleware_DeleteOAuthThirdParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/oauth/appoauththirdparty/appoauththirdparty.proto",
}
