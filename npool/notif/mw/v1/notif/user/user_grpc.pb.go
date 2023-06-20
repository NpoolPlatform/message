// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/notif/user/user.proto

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

const (
	Middleware_CreateNotifUser_FullMethodName     = "/notif.middleware.notif.user.v1.Middleware/CreateNotifUser"
	Middleware_CreateNotifUsers_FullMethodName    = "/notif.middleware.notif.user.v1.Middleware/CreateNotifUsers"
	Middleware_UpdateNotifUser_FullMethodName     = "/notif.middleware.notif.user.v1.Middleware/UpdateNotifUser"
	Middleware_GetNotifUser_FullMethodName        = "/notif.middleware.notif.user.v1.Middleware/GetNotifUser"
	Middleware_GetNotifUserOnly_FullMethodName    = "/notif.middleware.notif.user.v1.Middleware/GetNotifUserOnly"
	Middleware_GetNotifUsers_FullMethodName       = "/notif.middleware.notif.user.v1.Middleware/GetNotifUsers"
	Middleware_ExistNotifUser_FullMethodName      = "/notif.middleware.notif.user.v1.Middleware/ExistNotifUser"
	Middleware_ExistNotifUserConds_FullMethodName = "/notif.middleware.notif.user.v1.Middleware/ExistNotifUserConds"
	Middleware_CountNotifUsers_FullMethodName     = "/notif.middleware.notif.user.v1.Middleware/CountNotifUsers"
	Middleware_DeleteNotifUser_FullMethodName     = "/notif.middleware.notif.user.v1.Middleware/DeleteNotifUser"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateNotifUser(ctx context.Context, in *CreateNotifUserRequest, opts ...grpc.CallOption) (*CreateNotifUserResponse, error)
	CreateNotifUsers(ctx context.Context, in *CreateNotifUsersRequest, opts ...grpc.CallOption) (*CreateNotifUsersResponse, error)
	UpdateNotifUser(ctx context.Context, in *UpdateNotifUserRequest, opts ...grpc.CallOption) (*UpdateNotifUserResponse, error)
	GetNotifUser(ctx context.Context, in *GetNotifUserRequest, opts ...grpc.CallOption) (*GetNotifUserResponse, error)
	GetNotifUserOnly(ctx context.Context, in *GetNotifUserOnlyRequest, opts ...grpc.CallOption) (*GetNotifUserOnlyResponse, error)
	GetNotifUsers(ctx context.Context, in *GetNotifUsersRequest, opts ...grpc.CallOption) (*GetNotifUsersResponse, error)
	ExistNotifUser(ctx context.Context, in *ExistNotifUserRequest, opts ...grpc.CallOption) (*ExistNotifUserResponse, error)
	ExistNotifUserConds(ctx context.Context, in *ExistNotifUserCondsRequest, opts ...grpc.CallOption) (*ExistNotifUserCondsResponse, error)
	CountNotifUsers(ctx context.Context, in *CountNotifUsersRequest, opts ...grpc.CallOption) (*CountNotifUsersResponse, error)
	DeleteNotifUser(ctx context.Context, in *DeleteNotifUserRequest, opts ...grpc.CallOption) (*DeleteNotifUserResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateNotifUser(ctx context.Context, in *CreateNotifUserRequest, opts ...grpc.CallOption) (*CreateNotifUserResponse, error) {
	out := new(CreateNotifUserResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateNotifUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateNotifUsers(ctx context.Context, in *CreateNotifUsersRequest, opts ...grpc.CallOption) (*CreateNotifUsersResponse, error) {
	out := new(CreateNotifUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateNotifUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateNotifUser(ctx context.Context, in *UpdateNotifUserRequest, opts ...grpc.CallOption) (*UpdateNotifUserResponse, error) {
	out := new(UpdateNotifUserResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateNotifUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetNotifUser(ctx context.Context, in *GetNotifUserRequest, opts ...grpc.CallOption) (*GetNotifUserResponse, error) {
	out := new(GetNotifUserResponse)
	err := c.cc.Invoke(ctx, Middleware_GetNotifUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetNotifUserOnly(ctx context.Context, in *GetNotifUserOnlyRequest, opts ...grpc.CallOption) (*GetNotifUserOnlyResponse, error) {
	out := new(GetNotifUserOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetNotifUserOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetNotifUsers(ctx context.Context, in *GetNotifUsersRequest, opts ...grpc.CallOption) (*GetNotifUsersResponse, error) {
	out := new(GetNotifUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetNotifUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistNotifUser(ctx context.Context, in *ExistNotifUserRequest, opts ...grpc.CallOption) (*ExistNotifUserResponse, error) {
	out := new(ExistNotifUserResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistNotifUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistNotifUserConds(ctx context.Context, in *ExistNotifUserCondsRequest, opts ...grpc.CallOption) (*ExistNotifUserCondsResponse, error) {
	out := new(ExistNotifUserCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistNotifUserConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountNotifUsers(ctx context.Context, in *CountNotifUsersRequest, opts ...grpc.CallOption) (*CountNotifUsersResponse, error) {
	out := new(CountNotifUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_CountNotifUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteNotifUser(ctx context.Context, in *DeleteNotifUserRequest, opts ...grpc.CallOption) (*DeleteNotifUserResponse, error) {
	out := new(DeleteNotifUserResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteNotifUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateNotifUser(context.Context, *CreateNotifUserRequest) (*CreateNotifUserResponse, error)
	CreateNotifUsers(context.Context, *CreateNotifUsersRequest) (*CreateNotifUsersResponse, error)
	UpdateNotifUser(context.Context, *UpdateNotifUserRequest) (*UpdateNotifUserResponse, error)
	GetNotifUser(context.Context, *GetNotifUserRequest) (*GetNotifUserResponse, error)
	GetNotifUserOnly(context.Context, *GetNotifUserOnlyRequest) (*GetNotifUserOnlyResponse, error)
	GetNotifUsers(context.Context, *GetNotifUsersRequest) (*GetNotifUsersResponse, error)
	ExistNotifUser(context.Context, *ExistNotifUserRequest) (*ExistNotifUserResponse, error)
	ExistNotifUserConds(context.Context, *ExistNotifUserCondsRequest) (*ExistNotifUserCondsResponse, error)
	CountNotifUsers(context.Context, *CountNotifUsersRequest) (*CountNotifUsersResponse, error)
	DeleteNotifUser(context.Context, *DeleteNotifUserRequest) (*DeleteNotifUserResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateNotifUser(context.Context, *CreateNotifUserRequest) (*CreateNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifUser not implemented")
}
func (UnimplementedMiddlewareServer) CreateNotifUsers(context.Context, *CreateNotifUsersRequest) (*CreateNotifUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifUsers not implemented")
}
func (UnimplementedMiddlewareServer) UpdateNotifUser(context.Context, *UpdateNotifUserRequest) (*UpdateNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifUser not implemented")
}
func (UnimplementedMiddlewareServer) GetNotifUser(context.Context, *GetNotifUserRequest) (*GetNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifUser not implemented")
}
func (UnimplementedMiddlewareServer) GetNotifUserOnly(context.Context, *GetNotifUserOnlyRequest) (*GetNotifUserOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifUserOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetNotifUsers(context.Context, *GetNotifUsersRequest) (*GetNotifUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifUsers not implemented")
}
func (UnimplementedMiddlewareServer) ExistNotifUser(context.Context, *ExistNotifUserRequest) (*ExistNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistNotifUser not implemented")
}
func (UnimplementedMiddlewareServer) ExistNotifUserConds(context.Context, *ExistNotifUserCondsRequest) (*ExistNotifUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistNotifUserConds not implemented")
}
func (UnimplementedMiddlewareServer) CountNotifUsers(context.Context, *CountNotifUsersRequest) (*CountNotifUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountNotifUsers not implemented")
}
func (UnimplementedMiddlewareServer) DeleteNotifUser(context.Context, *DeleteNotifUserRequest) (*DeleteNotifUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifUser not implemented")
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

func _Middleware_CreateNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateNotifUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateNotifUser(ctx, req.(*CreateNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateNotifUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateNotifUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateNotifUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateNotifUsers(ctx, req.(*CreateNotifUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateNotifUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateNotifUser(ctx, req.(*UpdateNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetNotifUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetNotifUser(ctx, req.(*GetNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetNotifUserOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifUserOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetNotifUserOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetNotifUserOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetNotifUserOnly(ctx, req.(*GetNotifUserOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetNotifUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetNotifUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetNotifUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetNotifUsers(ctx, req.(*GetNotifUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistNotifUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistNotifUser(ctx, req.(*ExistNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistNotifUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistNotifUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistNotifUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistNotifUserConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistNotifUserConds(ctx, req.(*ExistNotifUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountNotifUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountNotifUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountNotifUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountNotifUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountNotifUsers(ctx, req.(*CountNotifUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteNotifUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotifUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteNotifUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteNotifUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteNotifUser(ctx, req.(*DeleteNotifUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.notif.user.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotifUser",
			Handler:    _Middleware_CreateNotifUser_Handler,
		},
		{
			MethodName: "CreateNotifUsers",
			Handler:    _Middleware_CreateNotifUsers_Handler,
		},
		{
			MethodName: "UpdateNotifUser",
			Handler:    _Middleware_UpdateNotifUser_Handler,
		},
		{
			MethodName: "GetNotifUser",
			Handler:    _Middleware_GetNotifUser_Handler,
		},
		{
			MethodName: "GetNotifUserOnly",
			Handler:    _Middleware_GetNotifUserOnly_Handler,
		},
		{
			MethodName: "GetNotifUsers",
			Handler:    _Middleware_GetNotifUsers_Handler,
		},
		{
			MethodName: "ExistNotifUser",
			Handler:    _Middleware_ExistNotifUser_Handler,
		},
		{
			MethodName: "ExistNotifUserConds",
			Handler:    _Middleware_ExistNotifUserConds_Handler,
		},
		{
			MethodName: "CountNotifUsers",
			Handler:    _Middleware_CountNotifUsers_Handler,
		},
		{
			MethodName: "DeleteNotifUser",
			Handler:    _Middleware_DeleteNotifUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/notif/user/user.proto",
}
