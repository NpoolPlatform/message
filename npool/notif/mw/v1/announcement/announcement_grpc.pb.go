// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/announcement/announcement.proto

package announcement

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
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error)
	GetAnnouncement(ctx context.Context, in *GetAnnouncementRequest, opts ...grpc.CallOption) (*GetAnnouncementResponse, error)
	GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error)
	ExistAnnouncement(ctx context.Context, in *ExistAnnouncementRequest, opts ...grpc.CallOption) (*ExistAnnouncementResponse, error)
	ExistAnnouncementConds(ctx context.Context, in *ExistAnnouncementCondsRequest, opts ...grpc.CallOption) (*ExistAnnouncementCondsResponse, error)
	DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/CreateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error) {
	out := new(UpdateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/UpdateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAnnouncement(ctx context.Context, in *GetAnnouncementRequest, opts ...grpc.CallOption) (*GetAnnouncementResponse, error) {
	out := new(GetAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/GetAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error) {
	out := new(GetAnnouncementsResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/GetAnnouncements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAnnouncement(ctx context.Context, in *ExistAnnouncementRequest, opts ...grpc.CallOption) (*ExistAnnouncementResponse, error) {
	out := new(ExistAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/ExistAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAnnouncementConds(ctx context.Context, in *ExistAnnouncementCondsRequest, opts ...grpc.CallOption) (*ExistAnnouncementCondsResponse, error) {
	out := new(ExistAnnouncementCondsResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/ExistAnnouncementConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error) {
	out := new(DeleteAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.v1.Middleware/DeleteAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error)
	GetAnnouncement(context.Context, *GetAnnouncementRequest) (*GetAnnouncementResponse, error)
	GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error)
	ExistAnnouncement(context.Context, *ExistAnnouncementRequest) (*ExistAnnouncementResponse, error)
	ExistAnnouncementConds(context.Context, *ExistAnnouncementCondsRequest) (*ExistAnnouncementCondsResponse, error)
	DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedMiddlewareServer) UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnnouncement not implemented")
}
func (UnimplementedMiddlewareServer) GetAnnouncement(context.Context, *GetAnnouncementRequest) (*GetAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncement not implemented")
}
func (UnimplementedMiddlewareServer) GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncements not implemented")
}
func (UnimplementedMiddlewareServer) ExistAnnouncement(context.Context, *ExistAnnouncementRequest) (*ExistAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAnnouncement not implemented")
}
func (UnimplementedMiddlewareServer) ExistAnnouncementConds(context.Context, *ExistAnnouncementCondsRequest) (*ExistAnnouncementCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAnnouncementConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncement not implemented")
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

func _Middleware_CreateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/CreateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAnnouncement(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/UpdateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateAnnouncement(ctx, req.(*UpdateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/GetAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAnnouncement(ctx, req.(*GetAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAnnouncements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAnnouncements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/GetAnnouncements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAnnouncements(ctx, req.(*GetAnnouncementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/ExistAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAnnouncement(ctx, req.(*ExistAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAnnouncementConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAnnouncementCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAnnouncementConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/ExistAnnouncementConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAnnouncementConds(ctx, req.(*ExistAnnouncementCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.v1.Middleware/DeleteAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAnnouncement(ctx, req.(*DeleteAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.announcement.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnnouncement",
			Handler:    _Middleware_CreateAnnouncement_Handler,
		},
		{
			MethodName: "UpdateAnnouncement",
			Handler:    _Middleware_UpdateAnnouncement_Handler,
		},
		{
			MethodName: "GetAnnouncement",
			Handler:    _Middleware_GetAnnouncement_Handler,
		},
		{
			MethodName: "GetAnnouncements",
			Handler:    _Middleware_GetAnnouncements_Handler,
		},
		{
			MethodName: "ExistAnnouncement",
			Handler:    _Middleware_ExistAnnouncement_Handler,
		},
		{
			MethodName: "ExistAnnouncementConds",
			Handler:    _Middleware_ExistAnnouncementConds_Handler,
		},
		{
			MethodName: "DeleteAnnouncement",
			Handler:    _Middleware_DeleteAnnouncement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/announcement/announcement.proto",
}
