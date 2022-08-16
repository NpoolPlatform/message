// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/role/role.proto

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
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error)
	GetAppRoles(ctx context.Context, in *GetAppRolesRequest, opts ...grpc.CallOption) (*GetAppRolesResponse, error)
	GetManyRoles(ctx context.Context, in *GetManyRolesRequest, opts ...grpc.CallOption) (*GetManyRolesResponse, error)
	GetRoleUser(ctx context.Context, in *GetRoleUserRequest, opts ...grpc.CallOption) (*GetRoleUserResponse, error)
	GetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error)
	GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error)
	GetManyRoleUsers(ctx context.Context, in *GetManyRoleUsersRequest, opts ...grpc.CallOption) (*GetManyRoleUsersResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	out := new(GetRoleResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error) {
	out := new(GetRolesResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppRoles(ctx context.Context, in *GetAppRolesRequest, opts ...grpc.CallOption) (*GetAppRolesResponse, error) {
	out := new(GetAppRolesResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetAppRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetManyRoles(ctx context.Context, in *GetManyRolesRequest, opts ...grpc.CallOption) (*GetManyRolesResponse, error) {
	out := new(GetManyRolesResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetManyRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRoleUser(ctx context.Context, in *GetRoleUserRequest, opts ...grpc.CallOption) (*GetRoleUserResponse, error) {
	out := new(GetRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetRoleUser", in, out, opts...)
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

func (c *middlewareClient) GetManyRoleUsers(ctx context.Context, in *GetManyRoleUsersRequest, opts ...grpc.CallOption) (*GetManyRoleUsersResponse, error) {
	out := new(GetManyRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.role.v1.Middleware/GetManyRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error)
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error)
	GetAppRoles(context.Context, *GetAppRolesRequest) (*GetAppRolesResponse, error)
	GetManyRoles(context.Context, *GetManyRolesRequest) (*GetManyRolesResponse, error)
	GetRoleUser(context.Context, *GetRoleUserRequest) (*GetRoleUserResponse, error)
	GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error)
	GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error)
	GetManyRoleUsers(context.Context, *GetManyRoleUsersRequest) (*GetManyRoleUsersResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedMiddlewareServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedMiddlewareServer) GetAppRoles(context.Context, *GetAppRolesRequest) (*GetAppRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoles not implemented")
}
func (UnimplementedMiddlewareServer) GetManyRoles(context.Context, *GetManyRolesRequest) (*GetManyRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyRoles not implemented")
}
func (UnimplementedMiddlewareServer) GetRoleUser(context.Context, *GetRoleUserRequest) (*GetRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleUser not implemented")
}
func (UnimplementedMiddlewareServer) GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetManyRoleUsers(context.Context, *GetManyRoleUsersRequest) (*GetManyRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyRoleUsers not implemented")
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

func _Middleware_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetAppRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppRoles(ctx, req.(*GetAppRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetManyRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetManyRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetManyRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetManyRoles(ctx, req.(*GetManyRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRoleUser(ctx, req.(*GetRoleUserRequest))
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

func _Middleware_GetManyRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetManyRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.role.v1.Middleware/GetManyRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetManyRoleUsers(ctx, req.(*GetManyRoleUsersRequest))
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
			MethodName: "GetRole",
			Handler:    _Middleware_GetRole_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _Middleware_GetRoles_Handler,
		},
		{
			MethodName: "GetAppRoles",
			Handler:    _Middleware_GetAppRoles_Handler,
		},
		{
			MethodName: "GetManyRoles",
			Handler:    _Middleware_GetManyRoles_Handler,
		},
		{
			MethodName: "GetRoleUser",
			Handler:    _Middleware_GetRoleUser_Handler,
		},
		{
			MethodName: "GetRoleUsers",
			Handler:    _Middleware_GetRoleUsers_Handler,
		},
		{
			MethodName: "GetAppRoleUsers",
			Handler:    _Middleware_GetAppRoleUsers_Handler,
		},
		{
			MethodName: "GetManyRoleUsers",
			Handler:    _Middleware_GetManyRoleUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/role/role.proto",
}
