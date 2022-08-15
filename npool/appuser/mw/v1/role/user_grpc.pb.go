// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/role/user.proto

package role

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
	CreateRoleUser(ctx context.Context, in *CreateRoleUserRequest, opts ...grpc.CallOption) (*CreateRoleUserResponse, error)
	DeleteRoleUser(ctx context.Context, in *DeleteRoleUserRequest, opts ...grpc.CallOption) (*DeleteRoleUserResponse, error)
	GetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error)
	GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateRoleUser(ctx context.Context, in *CreateRoleUserRequest, opts ...grpc.CallOption) (*CreateRoleUserResponse, error) {
	out := new(CreateRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/CreateRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteRoleUser(ctx context.Context, in *DeleteRoleUserRequest, opts ...grpc.CallOption) (*DeleteRoleUserResponse, error) {
	out := new(DeleteRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/DeleteRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error) {
	out := new(GetRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error) {
	out := new(GetAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateRoleUser(context.Context, *CreateRoleUserRequest) (*CreateRoleUserResponse, error)
	DeleteRoleUser(context.Context, *DeleteRoleUserRequest) (*DeleteRoleUserResponse, error)
	GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error)
	GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateRoleUser(context.Context, *CreateRoleUserRequest) (*CreateRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleUser not implemented")
}
func (UnimplementedMiddlewareServer) DeleteRoleUser(context.Context, *DeleteRoleUserRequest) (*DeleteRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleUser not implemented")
}
func (UnimplementedMiddlewareServer) GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUsers not implemented")
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

func _Middleware_CreateRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/CreateRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateRoleUser(ctx, req.(*CreateRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/DeleteRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteRoleUser(ctx, req.(*DeleteRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRoleUsers(ctx, req.(*GetRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppRoleUsers(ctx, req.(*GetAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.role.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoleUser",
			Handler:    _Middleware_CreateRoleUser_Handler,
		},
		{
			MethodName: "DeleteRoleUser",
			Handler:    _Middleware_DeleteRoleUser_Handler,
		},
		{
			MethodName: "GetRoleUsers",
			Handler:    _Middleware_GetRoleUsers_Handler,
		},
		{
			MethodName: "GetAppRoleUsers",
			Handler:    _Middleware_GetAppRoleUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/role/user.proto",
}
