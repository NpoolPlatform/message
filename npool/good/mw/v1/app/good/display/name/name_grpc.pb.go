// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/good/display/name/name.proto

package name

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
	Middleware_CreateDisplayName_FullMethodName     = "/good.middleware.app.good1.display.name.v1.Middleware/CreateDisplayName"
	Middleware_UpdateDisplayName_FullMethodName     = "/good.middleware.app.good1.display.name.v1.Middleware/UpdateDisplayName"
	Middleware_GetDisplayName_FullMethodName        = "/good.middleware.app.good1.display.name.v1.Middleware/GetDisplayName"
	Middleware_GetDisplayNames_FullMethodName       = "/good.middleware.app.good1.display.name.v1.Middleware/GetDisplayNames"
	Middleware_ExistDisplayNameConds_FullMethodName = "/good.middleware.app.good1.display.name.v1.Middleware/ExistDisplayNameConds"
	Middleware_DeleteDisplayName_FullMethodName     = "/good.middleware.app.good1.display.name.v1.Middleware/DeleteDisplayName"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateDisplayName(ctx context.Context, in *CreateDisplayNameRequest, opts ...grpc.CallOption) (*CreateDisplayNameResponse, error)
	UpdateDisplayName(ctx context.Context, in *UpdateDisplayNameRequest, opts ...grpc.CallOption) (*UpdateDisplayNameResponse, error)
	GetDisplayName(ctx context.Context, in *GetDisplayNameRequest, opts ...grpc.CallOption) (*GetDisplayNameResponse, error)
	GetDisplayNames(ctx context.Context, in *GetDisplayNamesRequest, opts ...grpc.CallOption) (*GetDisplayNamesResponse, error)
	ExistDisplayNameConds(ctx context.Context, in *ExistDisplayNameCondsRequest, opts ...grpc.CallOption) (*ExistDisplayNameCondsResponse, error)
	DeleteDisplayName(ctx context.Context, in *DeleteDisplayNameRequest, opts ...grpc.CallOption) (*DeleteDisplayNameResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateDisplayName(ctx context.Context, in *CreateDisplayNameRequest, opts ...grpc.CallOption) (*CreateDisplayNameResponse, error) {
	out := new(CreateDisplayNameResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateDisplayName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateDisplayName(ctx context.Context, in *UpdateDisplayNameRequest, opts ...grpc.CallOption) (*UpdateDisplayNameResponse, error) {
	out := new(UpdateDisplayNameResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateDisplayName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDisplayName(ctx context.Context, in *GetDisplayNameRequest, opts ...grpc.CallOption) (*GetDisplayNameResponse, error) {
	out := new(GetDisplayNameResponse)
	err := c.cc.Invoke(ctx, Middleware_GetDisplayName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDisplayNames(ctx context.Context, in *GetDisplayNamesRequest, opts ...grpc.CallOption) (*GetDisplayNamesResponse, error) {
	out := new(GetDisplayNamesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetDisplayNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistDisplayNameConds(ctx context.Context, in *ExistDisplayNameCondsRequest, opts ...grpc.CallOption) (*ExistDisplayNameCondsResponse, error) {
	out := new(ExistDisplayNameCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistDisplayNameConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteDisplayName(ctx context.Context, in *DeleteDisplayNameRequest, opts ...grpc.CallOption) (*DeleteDisplayNameResponse, error) {
	out := new(DeleteDisplayNameResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteDisplayName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateDisplayName(context.Context, *CreateDisplayNameRequest) (*CreateDisplayNameResponse, error)
	UpdateDisplayName(context.Context, *UpdateDisplayNameRequest) (*UpdateDisplayNameResponse, error)
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)
	GetDisplayNames(context.Context, *GetDisplayNamesRequest) (*GetDisplayNamesResponse, error)
	ExistDisplayNameConds(context.Context, *ExistDisplayNameCondsRequest) (*ExistDisplayNameCondsResponse, error)
	DeleteDisplayName(context.Context, *DeleteDisplayNameRequest) (*DeleteDisplayNameResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateDisplayName(context.Context, *CreateDisplayNameRequest) (*CreateDisplayNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDisplayName not implemented")
}
func (UnimplementedMiddlewareServer) UpdateDisplayName(context.Context, *UpdateDisplayNameRequest) (*UpdateDisplayNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDisplayName not implemented")
}
func (UnimplementedMiddlewareServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisplayName not implemented")
}
func (UnimplementedMiddlewareServer) GetDisplayNames(context.Context, *GetDisplayNamesRequest) (*GetDisplayNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisplayNames not implemented")
}
func (UnimplementedMiddlewareServer) ExistDisplayNameConds(context.Context, *ExistDisplayNameCondsRequest) (*ExistDisplayNameCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDisplayNameConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteDisplayName(context.Context, *DeleteDisplayNameRequest) (*DeleteDisplayNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDisplayName not implemented")
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

func _Middleware_CreateDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateDisplayName(ctx, req.(*CreateDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateDisplayName(ctx, req.(*UpdateDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDisplayName(ctx, req.(*GetDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDisplayNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisplayNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDisplayNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetDisplayNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDisplayNames(ctx, req.(*GetDisplayNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistDisplayNameConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDisplayNameCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistDisplayNameConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistDisplayNameConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistDisplayNameConds(ctx, req.(*ExistDisplayNameCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteDisplayName(ctx, req.(*DeleteDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.good1.display.name.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDisplayName",
			Handler:    _Middleware_CreateDisplayName_Handler,
		},
		{
			MethodName: "UpdateDisplayName",
			Handler:    _Middleware_UpdateDisplayName_Handler,
		},
		{
			MethodName: "GetDisplayName",
			Handler:    _Middleware_GetDisplayName_Handler,
		},
		{
			MethodName: "GetDisplayNames",
			Handler:    _Middleware_GetDisplayNames_Handler,
		},
		{
			MethodName: "ExistDisplayNameConds",
			Handler:    _Middleware_ExistDisplayNameConds_Handler,
		},
		{
			MethodName: "DeleteDisplayName",
			Handler:    _Middleware_DeleteDisplayName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/good/display/name/name.proto",
}
