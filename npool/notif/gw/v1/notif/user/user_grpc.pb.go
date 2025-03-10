// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/gw/v1/notif/user/user.proto

package user

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
	CreateNotifUser(ctx context.Context, in *CreateNotifUserRequest, opts ...grpc.CallOption) (*CreateNotifUserResponse, error)
	DeleteNotifUser(ctx context.Context, in *DeleteNotifUserRequest, opts ...grpc.CallOption) (*DeleteNotifUserResponse, error)
	GetNotifUsers(ctx context.Context, in *GetNotifUsersRequest, opts ...grpc.CallOption) (*GetNotifUsersResponse, error)
	GetAppNotifUsers(ctx context.Context, in *GetAppNotifUsersRequest, opts ...grpc.CallOption) (*GetAppNotifUsersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateNotifUser(ctx context.Context, in *CreateNotifUserRequest, opts ...grpc.CallOption) (*CreateNotifUserResponse, error) {
	out := new(CreateNotifUserResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif.user.v1.Gateway/CreateNotifUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteNotifUser(ctx context.Context, in *DeleteNotifUserRequest, opts ...grpc.CallOption) (*DeleteNotifUserResponse, error) {
	out := new(DeleteNotifUserResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif.user.v1.Gateway/DeleteNotifUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNotifUsers(ctx context.Context, in *GetNotifUsersRequest, opts ...grpc.CallOption) (*GetNotifUsersResponse, error) {
	out := new(GetNotifUsersResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif.user.v1.Gateway/GetNotifUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppNotifUsers(ctx context.Context, in *GetAppNotifUsersRequest, opts ...grpc.CallOption) (*GetAppNotifUsersResponse, error) {
	out := new(GetAppNotifUsersResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.notif.user.v1.Gateway/GetAppNotifUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateNotifUser(context.Context, *CreateNotifUserRequest) (*CreateNotifUserResponse, error)
	DeleteNotifUser(context.Context, *DeleteNotifUserRequest) (*DeleteNotifUserResponse, error)
	GetNotifUsers(context.Context, *GetNotifUsersRequest) (*GetNotifUsersResponse, error)
	GetAppNotifUsers(context.Context, *GetAppNotifUsersRequest) (*GetAppNotifUsersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateNotifUser(context.Context, *CreateNotifUserRequest) (*CreateNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifUser not implemented")
}
func (UnimplementedGatewayServer) DeleteNotifUser(context.Context, *DeleteNotifUserRequest) (*DeleteNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifUser not implemented")
}
func (UnimplementedGatewayServer) GetNotifUsers(context.Context, *GetNotifUsersRequest) (*GetNotifUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifUsers not implemented")
}
func (UnimplementedGatewayServer) GetAppNotifUsers(context.Context, *GetAppNotifUsersRequest) (*GetAppNotifUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppNotifUsers not implemented")
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

func _Gateway_CreateNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif.user.v1.Gateway/CreateNotifUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNotifUser(ctx, req.(*CreateNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif.user.v1.Gateway/DeleteNotifUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteNotifUser(ctx, req.(*DeleteNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNotifUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNotifUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif.user.v1.Gateway/GetNotifUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNotifUsers(ctx, req.(*GetNotifUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppNotifUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppNotifUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppNotifUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.notif.user.v1.Gateway/GetAppNotifUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppNotifUsers(ctx, req.(*GetAppNotifUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.notif.user.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotifUser",
			Handler:    _Gateway_CreateNotifUser_Handler,
		},
		{
			MethodName: "DeleteNotifUser",
			Handler:    _Gateway_DeleteNotifUser_Handler,
		},
		{
			MethodName: "GetNotifUsers",
			Handler:    _Gateway_GetNotifUsers_Handler,
		},
		{
			MethodName: "GetAppNotifUsers",
			Handler:    _Gateway_GetAppNotifUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/gw/v1/notif/user/user.proto",
}
