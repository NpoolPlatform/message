// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/appuser/appuser.proto

package appuser

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

// AppUserGwClient is the client API for AppUserGw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserGwClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error)
}

type appUserGwClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserGwClient(cc grpc.ClientConnInterface) AppUserGwClient {
	return &appUserGwClient{cc}
}

func (c *appUserGwClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.appuser1.v1.AppUserGw/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGwClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.appuser1.v1.AppUserGw/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGwClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.appuser1.v1.AppUserGw/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGwClient) GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error) {
	out := new(GetAppUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.appuser1.v1.AppUserGw/GetAppUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserGwServer is the server API for AppUserGw service.
// All implementations must embed UnimplementedAppUserGwServer
// for forward compatibility
type AppUserGwServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error)
	mustEmbedUnimplementedAppUserGwServer()
}

// UnimplementedAppUserGwServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserGwServer struct {
}

func (UnimplementedAppUserGwServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAppUserGwServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAppUserGwServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAppUserGwServer) GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUsers not implemented")
}
func (UnimplementedAppUserGwServer) mustEmbedUnimplementedAppUserGwServer() {}

// UnsafeAppUserGwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserGwServer will
// result in compilation errors.
type UnsafeAppUserGwServer interface {
	mustEmbedUnimplementedAppUserGwServer()
}

func RegisterAppUserGwServer(s grpc.ServiceRegistrar, srv AppUserGwServer) {
	s.RegisterService(&AppUserGw_ServiceDesc, srv)
}

func _AppUserGw_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGwServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.appuser1.v1.AppUserGw/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGwServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGw_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGwServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.appuser1.v1.AppUserGw/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGwServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGw_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGwServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.appuser1.v1.AppUserGw/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGwServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGw_GetAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGwServer).GetAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.appuser1.v1.AppUserGw/GetAppUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGwServer).GetAppUsers(ctx, req.(*GetAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserGw_ServiceDesc is the grpc.ServiceDesc for AppUserGw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserGw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.appuser1.v1.AppUserGw",
	HandlerType: (*AppUserGwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AppUserGw_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AppUserGw_UpdateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _AppUserGw_GetUsers_Handler,
		},
		{
			MethodName: "GetAppUsers",
			Handler:    _AppUserGw_GetAppUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/appuser/appuser.proto",
}
