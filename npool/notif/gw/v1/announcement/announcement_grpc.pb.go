// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/gw/v1/announcement/announcement.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error)
	DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error)
	GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...grpc.CallOption) (*CreateAnnouncementResponse, error) {
	out := new(CreateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/CreateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...grpc.CallOption) (*UpdateAnnouncementResponse, error) {
	out := new(UpdateAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/UpdateAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...grpc.CallOption) (*DeleteAnnouncementResponse, error) {
	out := new(DeleteAnnouncementResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/DeleteAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...grpc.CallOption) (*GetAnnouncementsResponse, error) {
	out := new(GetAnnouncementsResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/GetAnnouncements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error)
	UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error)
	DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error)
	GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAnnouncement(context.Context, *CreateAnnouncementRequest) (*CreateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnnouncement not implemented")
}
func (UnimplementedGatewayServer) UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest) (*UpdateAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnnouncement not implemented")
}
func (UnimplementedGatewayServer) DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest) (*DeleteAnnouncementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnouncement not implemented")
}
func (UnimplementedGatewayServer) GetAnnouncements(context.Context, *GetAnnouncementsRequest) (*GetAnnouncementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncements not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/CreateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAnnouncement(ctx, req.(*CreateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/UpdateAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAnnouncement(ctx, req.(*UpdateAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnnouncementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/DeleteAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAnnouncement(ctx, req.(*DeleteAnnouncementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAnnouncements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAnnouncements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/GetAnnouncements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAnnouncements(ctx, req.(*GetAnnouncementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.announcement.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnnouncement",
			Handler:    _Gateway_CreateAnnouncement_Handler,
		},
		{
			MethodName: "UpdateAnnouncement",
			Handler:    _Gateway_UpdateAnnouncement_Handler,
		},
		{
			MethodName: "DeleteAnnouncement",
			Handler:    _Gateway_DeleteAnnouncement_Handler,
		},
		{
			MethodName: "GetAnnouncements",
			Handler:    _Gateway_GetAnnouncements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/gw/v1/announcement/announcement.proto",
}
