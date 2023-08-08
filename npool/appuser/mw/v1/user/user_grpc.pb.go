// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/user/user.proto

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
	Middleware_CreateUser_FullMethodName      = "/appuser.middleware.user.v1.Middleware/CreateUser"
	Middleware_CreateThirdUser_FullMethodName = "/appuser.middleware.user.v1.Middleware/CreateThirdUser"
	Middleware_UpdateUser_FullMethodName      = "/appuser.middleware.user.v1.Middleware/UpdateUser"
	Middleware_GetUsers_FullMethodName        = "/appuser.middleware.user.v1.Middleware/GetUsers"
	Middleware_GetUser_FullMethodName         = "/appuser.middleware.user.v1.Middleware/GetUser"
	Middleware_GetThirdUsers_FullMethodName   = "/appuser.middleware.user.v1.Middleware/GetThirdUsers"
	Middleware_VerifyAccount_FullMethodName   = "/appuser.middleware.user.v1.Middleware/VerifyAccount"
	Middleware_VerifyUser_FullMethodName      = "/appuser.middleware.user.v1.Middleware/VerifyUser"
	Middleware_ExistUser_FullMethodName       = "/appuser.middleware.user.v1.Middleware/ExistUser"
	Middleware_ExistUserConds_FullMethodName  = "/appuser.middleware.user.v1.Middleware/ExistUserConds"
	Middleware_DeleteUser_FullMethodName      = "/appuser.middleware.user.v1.Middleware/DeleteUser"
	Middleware_DeleteThirdUser_FullMethodName = "/appuser.middleware.user.v1.Middleware/DeleteThirdUser"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateThirdUser(ctx context.Context, in *CreateThirdUserRequest, opts ...grpc.CallOption) (*CreateThirdUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetThirdUsers(ctx context.Context, in *GetThirdUsersRequest, opts ...grpc.CallOption) (*GetThirdUsersResponse, error)
	VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error)
	VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error)
	ExistUser(ctx context.Context, in *ExistUserRequest, opts ...grpc.CallOption) (*ExistUserResponse, error)
	ExistUserConds(ctx context.Context, in *ExistUserCondsRequest, opts ...grpc.CallOption) (*ExistUserCondsResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	DeleteThirdUser(ctx context.Context, in *DeleteThirdUserRequest, opts ...grpc.CallOption) (*DeleteThirdUserResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateThirdUser(ctx context.Context, in *CreateThirdUserRequest, opts ...grpc.CallOption) (*CreateThirdUserResponse, error) {
	out := new(CreateThirdUserResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateThirdUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, Middleware_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetThirdUsers(ctx context.Context, in *GetThirdUsersRequest, opts ...grpc.CallOption) (*GetThirdUsersResponse, error) {
	out := new(GetThirdUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetThirdUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error) {
	out := new(VerifyAccountResponse)
	err := c.cc.Invoke(ctx, Middleware_VerifyAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) VerifyUser(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error) {
	out := new(VerifyUserResponse)
	err := c.cc.Invoke(ctx, Middleware_VerifyUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistUser(ctx context.Context, in *ExistUserRequest, opts ...grpc.CallOption) (*ExistUserResponse, error) {
	out := new(ExistUserResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistUserConds(ctx context.Context, in *ExistUserCondsRequest, opts ...grpc.CallOption) (*ExistUserCondsResponse, error) {
	out := new(ExistUserCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistUserConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteThirdUser(ctx context.Context, in *DeleteThirdUserRequest, opts ...grpc.CallOption) (*DeleteThirdUserResponse, error) {
	out := new(DeleteThirdUserResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteThirdUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateThirdUser(context.Context, *CreateThirdUserRequest) (*CreateThirdUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetThirdUsers(context.Context, *GetThirdUsersRequest) (*GetThirdUsersResponse, error)
	VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error)
	VerifyUser(context.Context, *VerifyUserRequest) (*VerifyUserResponse, error)
	ExistUser(context.Context, *ExistUserRequest) (*ExistUserResponse, error)
	ExistUserConds(context.Context, *ExistUserCondsRequest) (*ExistUserCondsResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	DeleteThirdUser(context.Context, *DeleteThirdUserRequest) (*DeleteThirdUserResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMiddlewareServer) CreateThirdUser(context.Context, *CreateThirdUserRequest) (*CreateThirdUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThirdUser not implemented")
}
func (UnimplementedMiddlewareServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMiddlewareServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedMiddlewareServer) GetThirdUsers(context.Context, *GetThirdUsersRequest) (*GetThirdUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThirdUsers not implemented")
}
func (UnimplementedMiddlewareServer) VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
func (UnimplementedMiddlewareServer) VerifyUser(context.Context, *VerifyUserRequest) (*VerifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedMiddlewareServer) ExistUser(context.Context, *ExistUserRequest) (*ExistUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUser not implemented")
}
func (UnimplementedMiddlewareServer) ExistUserConds(context.Context, *ExistUserCondsRequest) (*ExistUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUserConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMiddlewareServer) DeleteThirdUser(context.Context, *DeleteThirdUserRequest) (*DeleteThirdUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteThirdUser not implemented")
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

func _Middleware_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateThirdUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThirdUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateThirdUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateThirdUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateThirdUser(ctx, req.(*CreateThirdUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetThirdUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThirdUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetThirdUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetThirdUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetThirdUsers(ctx, req.(*GetThirdUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_VerifyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).VerifyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_VerifyAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).VerifyAccount(ctx, req.(*VerifyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_VerifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).VerifyUser(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistUser(ctx, req.(*ExistUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistUserConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistUserConds(ctx, req.(*ExistUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteThirdUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteThirdUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteThirdUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteThirdUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteThirdUser(ctx, req.(*DeleteThirdUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.user.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Middleware_CreateUser_Handler,
		},
		{
			MethodName: "CreateThirdUser",
			Handler:    _Middleware_CreateThirdUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Middleware_UpdateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Middleware_GetUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Middleware_GetUser_Handler,
		},
		{
			MethodName: "GetThirdUsers",
			Handler:    _Middleware_GetThirdUsers_Handler,
		},
		{
			MethodName: "VerifyAccount",
			Handler:    _Middleware_VerifyAccount_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _Middleware_VerifyUser_Handler,
		},
		{
			MethodName: "ExistUser",
			Handler:    _Middleware_ExistUser_Handler,
		},
		{
			MethodName: "ExistUserConds",
			Handler:    _Middleware_ExistUserConds_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Middleware_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteThirdUser",
			Handler:    _Middleware_DeleteThirdUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/user/user.proto",
}
