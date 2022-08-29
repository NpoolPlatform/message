// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/role/role.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// Create role should ignore genesis roles
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error)
	// Update role should ignore genesis roles
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error)
	// Admin apis
	CreateAppRole(ctx context.Context, in *CreateAppRoleRequest, opts ...grpc.CallOption) (*CreateAppRoleResponse, error)
	GetAppRoles(ctx context.Context, in *GetAppRolesRequest, opts ...grpc.CallOption) (*GetAppRolesResponse, error)
	CreateRoleUser(ctx context.Context, in *CreateRoleUserRequest, opts ...grpc.CallOption) (*CreateRoleUserResponse, error)
	GetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error)
	DeleteRoleUser(ctx context.Context, in *DeleteRoleUserRequest, opts ...grpc.CallOption) (*DeleteRoleUserResponse, error)
	DeleteAppRoleUser(ctx context.Context, in *DeleteAppRoleUserRequest, opts ...grpc.CallOption) (*DeleteAppRoleUserResponse, error)
	CreateAppRoleUser(ctx context.Context, in *CreateAppRoleUserRequest, opts ...grpc.CallOption) (*CreateAppRoleUserResponse, error)
	GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesResponse, error) {
	out := new(GetRolesResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*UpdateRoleResponse, error) {
	out := new(UpdateRoleResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppRole(ctx context.Context, in *CreateAppRoleRequest, opts ...grpc.CallOption) (*CreateAppRoleResponse, error) {
	out := new(CreateAppRoleResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/CreateAppRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppRoles(ctx context.Context, in *GetAppRolesRequest, opts ...grpc.CallOption) (*GetAppRolesResponse, error) {
	out := new(GetAppRolesResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/GetAppRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateRoleUser(ctx context.Context, in *CreateRoleUserRequest, opts ...grpc.CallOption) (*CreateRoleUserResponse, error) {
	out := new(CreateRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/CreateRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRoleUsers(ctx context.Context, in *GetRoleUsersRequest, opts ...grpc.CallOption) (*GetRoleUsersResponse, error) {
	out := new(GetRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/GetRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteRoleUser(ctx context.Context, in *DeleteRoleUserRequest, opts ...grpc.CallOption) (*DeleteRoleUserResponse, error) {
	out := new(DeleteRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/DeleteRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAppRoleUser(ctx context.Context, in *DeleteAppRoleUserRequest, opts ...grpc.CallOption) (*DeleteAppRoleUserResponse, error) {
	out := new(DeleteAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/DeleteAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppRoleUser(ctx context.Context, in *CreateAppRoleUserRequest, opts ...grpc.CallOption) (*CreateAppRoleUserResponse, error) {
	out := new(CreateAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/CreateAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error) {
	out := new(GetAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.role.v1.Gateway/GetAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// Create role should ignore genesis roles
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error)
	// Update role should ignore genesis roles
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
	// Admin apis
	CreateAppRole(context.Context, *CreateAppRoleRequest) (*CreateAppRoleResponse, error)
	GetAppRoles(context.Context, *GetAppRolesRequest) (*GetAppRolesResponse, error)
	CreateRoleUser(context.Context, *CreateRoleUserRequest) (*CreateRoleUserResponse, error)
	GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error)
	DeleteRoleUser(context.Context, *DeleteRoleUserRequest) (*DeleteRoleUserResponse, error)
	DeleteAppRoleUser(context.Context, *DeleteAppRoleUserRequest) (*DeleteAppRoleUserResponse, error)
	CreateAppRoleUser(context.Context, *CreateAppRoleUserRequest) (*CreateAppRoleUserResponse, error)
	GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedGatewayServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedGatewayServer) UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedGatewayServer) CreateAppRole(context.Context, *CreateAppRoleRequest) (*CreateAppRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppRole not implemented")
}
func (UnimplementedGatewayServer) GetAppRoles(context.Context, *GetAppRolesRequest) (*GetAppRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoles not implemented")
}
func (UnimplementedGatewayServer) CreateRoleUser(context.Context, *CreateRoleUserRequest) (*CreateRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleUser not implemented")
}
func (UnimplementedGatewayServer) GetRoleUsers(context.Context, *GetRoleUsersRequest) (*GetRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleUsers not implemented")
}
func (UnimplementedGatewayServer) DeleteRoleUser(context.Context, *DeleteRoleUserRequest) (*DeleteRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleUser not implemented")
}
func (UnimplementedGatewayServer) DeleteAppRoleUser(context.Context, *DeleteAppRoleUserRequest) (*DeleteAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppRoleUser not implemented")
}
func (UnimplementedGatewayServer) CreateAppRoleUser(context.Context, *CreateAppRoleUserRequest) (*CreateAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppRoleUser not implemented")
}
func (UnimplementedGatewayServer) GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUsers not implemented")
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

func _Gateway_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/CreateAppRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppRole(ctx, req.(*CreateAppRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/GetAppRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppRoles(ctx, req.(*GetAppRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/CreateRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateRoleUser(ctx, req.(*CreateRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/GetRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRoleUsers(ctx, req.(*GetRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/DeleteRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteRoleUser(ctx, req.(*DeleteRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/DeleteAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAppRoleUser(ctx, req.(*DeleteAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/CreateAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppRoleUser(ctx, req.(*CreateAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.role.v1.Gateway/GetAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppRoleUsers(ctx, req.(*GetAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.role.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _Gateway_CreateRole_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _Gateway_GetRoles_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Gateway_UpdateRole_Handler,
		},
		{
			MethodName: "CreateAppRole",
			Handler:    _Gateway_CreateAppRole_Handler,
		},
		{
			MethodName: "GetAppRoles",
			Handler:    _Gateway_GetAppRoles_Handler,
		},
		{
			MethodName: "CreateRoleUser",
			Handler:    _Gateway_CreateRoleUser_Handler,
		},
		{
			MethodName: "GetRoleUsers",
			Handler:    _Gateway_GetRoleUsers_Handler,
		},
		{
			MethodName: "DeleteRoleUser",
			Handler:    _Gateway_DeleteRoleUser_Handler,
		},
		{
			MethodName: "DeleteAppRoleUser",
			Handler:    _Gateway_DeleteAppRoleUser_Handler,
		},
		{
			MethodName: "CreateAppRoleUser",
			Handler:    _Gateway_CreateAppRoleUser_Handler,
		},
		{
			MethodName: "GetAppRoleUsers",
			Handler:    _Gateway_GetAppRoleUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/role/role.proto",
}
