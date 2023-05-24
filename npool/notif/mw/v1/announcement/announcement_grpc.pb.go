// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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

const (
	Middleware_GetAnnouncementStates_FullMethodName = "/notif.middleware.announcement.announcement.v1.Middleware/GetAnnouncementStates"
	Middleware_GetAnnouncements_FullMethodName      = "/notif.middleware.announcement.announcement.v1.Middleware/GetAnnouncements"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetAnnouncementStates(ctx context.Context, in *GetAnnouncementStatesRequest, opts ...grpc.CallOption) (*GetAnnouncementStatesResponse, error)
	GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetAnnouncementStates(ctx context.Context, in *GetAnnouncementStatesRequest, opts ...grpc.CallOption) (*GetAnnouncementStatesResponse, error) {
	out := new(GetAnnouncementStatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAnnouncementStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error) {
	out := new(GetAnnouncementsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAnnouncements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetAnnouncementStates(context.Context, *GetAnnouncementStatesRequest) (*GetAnnouncementStatesResponse, error)
	GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetAnnouncementStates(context.Context, *GetAnnouncementStatesRequest) (*GetAnnouncementStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncementStates not implemented")
}
func (UnimplementedMiddlewareServer) GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncements not implemented")
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

func _Middleware_GetAnnouncementStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncementStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAnnouncementStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAnnouncementStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAnnouncementStates(ctx, req.(*GetAnnouncementStatesRequest))
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
		FullMethod: Middleware_GetAnnouncements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAnnouncements(ctx, req.(*GetAnnouncementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.announcement.announcement.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnnouncementStates",
			Handler:    _Middleware_GetAnnouncementStates_Handler,
		},
		{
			MethodName: "GetAnnouncements",
			Handler:    _Middleware_GetAnnouncements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/announcement/announcement.proto",
}
