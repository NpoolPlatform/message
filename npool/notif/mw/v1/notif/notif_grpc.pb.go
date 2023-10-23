// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/notif/notif.proto

package notif

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
	Middleware_CreateNotif_FullMethodName     = "/notif.middleware.notif.v1.Middleware/CreateNotif"
	Middleware_CreateNotifs_FullMethodName    = "/notif.middleware.notif.v1.Middleware/CreateNotifs"
	Middleware_UpdateNotif_FullMethodName     = "/notif.middleware.notif.v1.Middleware/UpdateNotif"
	Middleware_UpdateNotifs_FullMethodName    = "/notif.middleware.notif.v1.Middleware/UpdateNotifs"
	Middleware_GetNotif_FullMethodName        = "/notif.middleware.notif.v1.Middleware/GetNotif"
	Middleware_GetNotifs_FullMethodName       = "/notif.middleware.notif.v1.Middleware/GetNotifs"
	Middleware_GenerateNotifs_FullMethodName  = "/notif.middleware.notif.v1.Middleware/GenerateNotifs"
	Middleware_ExistNotifConds_FullMethodName = "/notif.middleware.notif.v1.Middleware/ExistNotifConds"
	Middleware_DeleteNotif_FullMethodName     = "/notif.middleware.notif.v1.Middleware/DeleteNotif"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateNotif(ctx context.Context, in *CreateNotifRequest, opts ...grpc.CallOption) (*CreateNotifResponse, error)
	CreateNotifs(ctx context.Context, in *CreateNotifsRequest, opts ...grpc.CallOption) (*CreateNotifsResponse, error)
	UpdateNotif(ctx context.Context, in *UpdateNotifRequest, opts ...grpc.CallOption) (*UpdateNotifResponse, error)
	UpdateNotifs(ctx context.Context, in *UpdateNotifsRequest, opts ...grpc.CallOption) (*UpdateNotifsResponse, error)
	GetNotif(ctx context.Context, in *GetNotifRequest, opts ...grpc.CallOption) (*GetNotifResponse, error)
	GetNotifs(ctx context.Context, in *GetNotifsRequest, opts ...grpc.CallOption) (*GetNotifsResponse, error)
	GenerateNotifs(ctx context.Context, in *GenerateNotifsRequest, opts ...grpc.CallOption) (*GenerateNotifsResponse, error)
	ExistNotifConds(ctx context.Context, in *ExistNotifCondsRequest, opts ...grpc.CallOption) (*ExistNotifCondsResponse, error)
	DeleteNotif(ctx context.Context, in *DeleteNotifRequest, opts ...grpc.CallOption) (*DeleteNotifResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateNotif(ctx context.Context, in *CreateNotifRequest, opts ...grpc.CallOption) (*CreateNotifResponse, error) {
	out := new(CreateNotifResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateNotif_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateNotifs(ctx context.Context, in *CreateNotifsRequest, opts ...grpc.CallOption) (*CreateNotifsResponse, error) {
	out := new(CreateNotifsResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateNotifs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateNotif(ctx context.Context, in *UpdateNotifRequest, opts ...grpc.CallOption) (*UpdateNotifResponse, error) {
	out := new(UpdateNotifResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateNotif_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateNotifs(ctx context.Context, in *UpdateNotifsRequest, opts ...grpc.CallOption) (*UpdateNotifsResponse, error) {
	out := new(UpdateNotifsResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateNotifs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetNotif(ctx context.Context, in *GetNotifRequest, opts ...grpc.CallOption) (*GetNotifResponse, error) {
	out := new(GetNotifResponse)
	err := c.cc.Invoke(ctx, Middleware_GetNotif_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetNotifs(ctx context.Context, in *GetNotifsRequest, opts ...grpc.CallOption) (*GetNotifsResponse, error) {
	out := new(GetNotifsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetNotifs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GenerateNotifs(ctx context.Context, in *GenerateNotifsRequest, opts ...grpc.CallOption) (*GenerateNotifsResponse, error) {
	out := new(GenerateNotifsResponse)
	err := c.cc.Invoke(ctx, Middleware_GenerateNotifs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistNotifConds(ctx context.Context, in *ExistNotifCondsRequest, opts ...grpc.CallOption) (*ExistNotifCondsResponse, error) {
	out := new(ExistNotifCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistNotifConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteNotif(ctx context.Context, in *DeleteNotifRequest, opts ...grpc.CallOption) (*DeleteNotifResponse, error) {
	out := new(DeleteNotifResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteNotif_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateNotif(context.Context, *CreateNotifRequest) (*CreateNotifResponse, error)
	CreateNotifs(context.Context, *CreateNotifsRequest) (*CreateNotifsResponse, error)
	UpdateNotif(context.Context, *UpdateNotifRequest) (*UpdateNotifResponse, error)
	UpdateNotifs(context.Context, *UpdateNotifsRequest) (*UpdateNotifsResponse, error)
	GetNotif(context.Context, *GetNotifRequest) (*GetNotifResponse, error)
	GetNotifs(context.Context, *GetNotifsRequest) (*GetNotifsResponse, error)
	GenerateNotifs(context.Context, *GenerateNotifsRequest) (*GenerateNotifsResponse, error)
	ExistNotifConds(context.Context, *ExistNotifCondsRequest) (*ExistNotifCondsResponse, error)
	DeleteNotif(context.Context, *DeleteNotifRequest) (*DeleteNotifResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateNotif(context.Context, *CreateNotifRequest) (*CreateNotifResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotif not implemented")
}
func (UnimplementedMiddlewareServer) CreateNotifs(context.Context, *CreateNotifsRequest) (*CreateNotifsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifs not implemented")
}
func (UnimplementedMiddlewareServer) UpdateNotif(context.Context, *UpdateNotifRequest) (*UpdateNotifResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotif not implemented")
}
func (UnimplementedMiddlewareServer) UpdateNotifs(context.Context, *UpdateNotifsRequest) (*UpdateNotifsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifs not implemented")
}
func (UnimplementedMiddlewareServer) GetNotif(context.Context, *GetNotifRequest) (*GetNotifResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotif not implemented")
}
func (UnimplementedMiddlewareServer) GetNotifs(context.Context, *GetNotifsRequest) (*GetNotifsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifs not implemented")
}
func (UnimplementedMiddlewareServer) GenerateNotifs(context.Context, *GenerateNotifsRequest) (*GenerateNotifsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateNotifs not implemented")
}
func (UnimplementedMiddlewareServer) ExistNotifConds(context.Context, *ExistNotifCondsRequest) (*ExistNotifCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistNotifConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteNotif(context.Context, *DeleteNotifRequest) (*DeleteNotifResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotif not implemented")
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

func _Middleware_CreateNotif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateNotif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateNotif_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateNotif(ctx, req.(*CreateNotifRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateNotifs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateNotifs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateNotifs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateNotifs(ctx, req.(*CreateNotifsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateNotif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotifRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateNotif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateNotif_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateNotif(ctx, req.(*UpdateNotifRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateNotifs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotifsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateNotifs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateNotifs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateNotifs(ctx, req.(*UpdateNotifsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetNotif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetNotif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetNotif_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetNotif(ctx, req.(*GetNotifRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetNotifs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetNotifs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetNotifs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetNotifs(ctx, req.(*GetNotifsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GenerateNotifs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateNotifsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GenerateNotifs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GenerateNotifs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GenerateNotifs(ctx, req.(*GenerateNotifsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistNotifConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistNotifCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistNotifConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistNotifConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistNotifConds(ctx, req.(*ExistNotifCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteNotif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotifRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteNotif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteNotif_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteNotif(ctx, req.(*DeleteNotifRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.notif.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotif",
			Handler:    _Middleware_CreateNotif_Handler,
		},
		{
			MethodName: "CreateNotifs",
			Handler:    _Middleware_CreateNotifs_Handler,
		},
		{
			MethodName: "UpdateNotif",
			Handler:    _Middleware_UpdateNotif_Handler,
		},
		{
			MethodName: "UpdateNotifs",
			Handler:    _Middleware_UpdateNotifs_Handler,
		},
		{
			MethodName: "GetNotif",
			Handler:    _Middleware_GetNotif_Handler,
		},
		{
			MethodName: "GetNotifs",
			Handler:    _Middleware_GetNotifs_Handler,
		},
		{
			MethodName: "GenerateNotifs",
			Handler:    _Middleware_GenerateNotifs_Handler,
		},
		{
			MethodName: "ExistNotifConds",
			Handler:    _Middleware_ExistNotifConds_Handler,
		},
		{
			MethodName: "DeleteNotif",
			Handler:    _Middleware_DeleteNotif_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/notif/notif.proto",
}
