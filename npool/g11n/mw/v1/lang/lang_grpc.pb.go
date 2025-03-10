// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/g11n/mw/v1/lang/lang.proto

package lang

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
	CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error)
	CreateLangs(ctx context.Context, in *CreateLangsRequest, opts ...grpc.CallOption) (*CreateLangsResponse, error)
	UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error)
	GetLang(ctx context.Context, in *GetLangRequest, opts ...grpc.CallOption) (*GetLangResponse, error)
	GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error)
	ExistLangConds(ctx context.Context, in *ExistLangCondsRequest, opts ...grpc.CallOption) (*ExistLangCondsResponse, error)
	DeleteLang(ctx context.Context, in *DeleteLangRequest, opts ...grpc.CallOption) (*DeleteLangResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error) {
	out := new(CreateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/CreateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateLangs(ctx context.Context, in *CreateLangsRequest, opts ...grpc.CallOption) (*CreateLangsResponse, error) {
	out := new(CreateLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/CreateLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error) {
	out := new(UpdateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/UpdateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLang(ctx context.Context, in *GetLangRequest, opts ...grpc.CallOption) (*GetLangResponse, error) {
	out := new(GetLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/GetLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error) {
	out := new(GetLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/GetLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistLangConds(ctx context.Context, in *ExistLangCondsRequest, opts ...grpc.CallOption) (*ExistLangCondsResponse, error) {
	out := new(ExistLangCondsResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/ExistLangConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteLang(ctx context.Context, in *DeleteLangRequest, opts ...grpc.CallOption) (*DeleteLangResponse, error) {
	out := new(DeleteLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.middleware.lang.v1.Middleware/DeleteLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error)
	CreateLangs(context.Context, *CreateLangsRequest) (*CreateLangsResponse, error)
	UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error)
	GetLang(context.Context, *GetLangRequest) (*GetLangResponse, error)
	GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error)
	ExistLangConds(context.Context, *ExistLangCondsRequest) (*ExistLangCondsResponse, error)
	DeleteLang(context.Context, *DeleteLangRequest) (*DeleteLangResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLang not implemented")
}
func (UnimplementedMiddlewareServer) CreateLangs(context.Context, *CreateLangsRequest) (*CreateLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLangs not implemented")
}
func (UnimplementedMiddlewareServer) UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLang not implemented")
}
func (UnimplementedMiddlewareServer) GetLang(context.Context, *GetLangRequest) (*GetLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLang not implemented")
}
func (UnimplementedMiddlewareServer) GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLangs not implemented")
}
func (UnimplementedMiddlewareServer) ExistLangConds(context.Context, *ExistLangCondsRequest) (*ExistLangCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistLangConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteLang(context.Context, *DeleteLangRequest) (*DeleteLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLang not implemented")
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

func _Middleware_CreateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/CreateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateLang(ctx, req.(*CreateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/CreateLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateLangs(ctx, req.(*CreateLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/UpdateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateLang(ctx, req.(*UpdateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/GetLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLang(ctx, req.(*GetLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/GetLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLangs(ctx, req.(*GetLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistLangConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistLangCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistLangConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/ExistLangConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistLangConds(ctx, req.(*ExistLangCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.middleware.lang.v1.Middleware/DeleteLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteLang(ctx, req.(*DeleteLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.middleware.lang.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLang",
			Handler:    _Middleware_CreateLang_Handler,
		},
		{
			MethodName: "CreateLangs",
			Handler:    _Middleware_CreateLangs_Handler,
		},
		{
			MethodName: "UpdateLang",
			Handler:    _Middleware_UpdateLang_Handler,
		},
		{
			MethodName: "GetLang",
			Handler:    _Middleware_GetLang_Handler,
		},
		{
			MethodName: "GetLangs",
			Handler:    _Middleware_GetLangs_Handler,
		},
		{
			MethodName: "ExistLangConds",
			Handler:    _Middleware_ExistLangConds_Handler,
		},
		{
			MethodName: "DeleteLang",
			Handler:    _Middleware_DeleteLang_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/mw/v1/lang/lang.proto",
}
