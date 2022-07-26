// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appusergw/admin/admin.proto

package admin

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

// AppUserGatewayAdminClient is the client API for AppUserGatewayAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserGatewayAdminClient interface {
	CreateAdminApps(ctx context.Context, in *CreateAdminAppsRequest, opts ...grpc.CallOption) (*CreateAdminAppsResponse, error)
	GetAdminApps(ctx context.Context, in *GetAdminAppsRequest, opts ...grpc.CallOption) (*GetAdminAppsResponse, error)
	CreateGenesisRole(ctx context.Context, in *CreateGenesisRoleRequest, opts ...grpc.CallOption) (*CreateGenesisRoleResponse, error)
	GetGenesisRole(ctx context.Context, in *GetGenesisRoleRequest, opts ...grpc.CallOption) (*GetGenesisRoleResponse, error)
	GetAppGenesisAppRoleUsers(ctx context.Context, in *GetAppGenesisAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppGenesisAppRoleUsersResponse, error)
}

type appUserGatewayAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserGatewayAdminClient(cc grpc.ClientConnInterface) AppUserGatewayAdminClient {
	return &appUserGatewayAdminClient{cc}
}

func (c *appUserGatewayAdminClient) CreateAdminApps(ctx context.Context, in *CreateAdminAppsRequest, opts ...grpc.CallOption) (*CreateAdminAppsResponse, error) {
	out := new(CreateAdminAppsResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.admin.v1.AppUserGatewayAdmin/CreateAdminApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGatewayAdminClient) GetAdminApps(ctx context.Context, in *GetAdminAppsRequest, opts ...grpc.CallOption) (*GetAdminAppsResponse, error) {
	out := new(GetAdminAppsResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetAdminApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGatewayAdminClient) CreateGenesisRole(ctx context.Context, in *CreateGenesisRoleRequest, opts ...grpc.CallOption) (*CreateGenesisRoleResponse, error) {
	out := new(CreateGenesisRoleResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.admin.v1.AppUserGatewayAdmin/CreateGenesisRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGatewayAdminClient) GetGenesisRole(ctx context.Context, in *GetGenesisRoleRequest, opts ...grpc.CallOption) (*GetGenesisRoleResponse, error) {
	out := new(GetGenesisRoleResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetGenesisRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserGatewayAdminClient) GetAppGenesisAppRoleUsers(ctx context.Context, in *GetAppGenesisAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppGenesisAppRoleUsersResponse, error) {
	out := new(GetAppGenesisAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetAppGenesisAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserGatewayAdminServer is the server API for AppUserGatewayAdmin service.
// All implementations must embed UnimplementedAppUserGatewayAdminServer
// for forward compatibility
type AppUserGatewayAdminServer interface {
	CreateAdminApps(context.Context, *CreateAdminAppsRequest) (*CreateAdminAppsResponse, error)
	GetAdminApps(context.Context, *GetAdminAppsRequest) (*GetAdminAppsResponse, error)
	CreateGenesisRole(context.Context, *CreateGenesisRoleRequest) (*CreateGenesisRoleResponse, error)
	GetGenesisRole(context.Context, *GetGenesisRoleRequest) (*GetGenesisRoleResponse, error)
	GetAppGenesisAppRoleUsers(context.Context, *GetAppGenesisAppRoleUsersRequest) (*GetAppGenesisAppRoleUsersResponse, error)
	mustEmbedUnimplementedAppUserGatewayAdminServer()
}

// UnimplementedAppUserGatewayAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserGatewayAdminServer struct {
}

func (UnimplementedAppUserGatewayAdminServer) CreateAdminApps(context.Context, *CreateAdminAppsRequest) (*CreateAdminAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminApps not implemented")
}
func (UnimplementedAppUserGatewayAdminServer) GetAdminApps(context.Context, *GetAdminAppsRequest) (*GetAdminAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminApps not implemented")
}
func (UnimplementedAppUserGatewayAdminServer) CreateGenesisRole(context.Context, *CreateGenesisRoleRequest) (*CreateGenesisRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenesisRole not implemented")
}
func (UnimplementedAppUserGatewayAdminServer) GetGenesisRole(context.Context, *GetGenesisRoleRequest) (*GetGenesisRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenesisRole not implemented")
}
func (UnimplementedAppUserGatewayAdminServer) GetAppGenesisAppRoleUsers(context.Context, *GetAppGenesisAppRoleUsersRequest) (*GetAppGenesisAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGenesisAppRoleUsers not implemented")
}
func (UnimplementedAppUserGatewayAdminServer) mustEmbedUnimplementedAppUserGatewayAdminServer() {}

// UnsafeAppUserGatewayAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserGatewayAdminServer will
// result in compilation errors.
type UnsafeAppUserGatewayAdminServer interface {
	mustEmbedUnimplementedAppUserGatewayAdminServer()
}

func RegisterAppUserGatewayAdminServer(s grpc.ServiceRegistrar, srv AppUserGatewayAdminServer) {
	s.RegisterService(&AppUserGatewayAdmin_ServiceDesc, srv)
}

func _AppUserGatewayAdmin_CreateAdminApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGatewayAdminServer).CreateAdminApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.admin.v1.AppUserGatewayAdmin/CreateAdminApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGatewayAdminServer).CreateAdminApps(ctx, req.(*CreateAdminAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGatewayAdmin_GetAdminApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGatewayAdminServer).GetAdminApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetAdminApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGatewayAdminServer).GetAdminApps(ctx, req.(*GetAdminAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGatewayAdmin_CreateGenesisRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGenesisRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGatewayAdminServer).CreateGenesisRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.admin.v1.AppUserGatewayAdmin/CreateGenesisRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGatewayAdminServer).CreateGenesisRole(ctx, req.(*CreateGenesisRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGatewayAdmin_GetGenesisRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenesisRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGatewayAdminServer).GetGenesisRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetGenesisRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGatewayAdminServer).GetGenesisRole(ctx, req.(*GetGenesisRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserGatewayAdmin_GetAppGenesisAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGenesisAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserGatewayAdminServer).GetAppGenesisAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.admin.v1.AppUserGatewayAdmin/GetAppGenesisAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserGatewayAdminServer).GetAppGenesisAppRoleUsers(ctx, req.(*GetAppGenesisAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserGatewayAdmin_ServiceDesc is the grpc.ServiceDesc for AppUserGatewayAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserGatewayAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.gateway.admin.v1.AppUserGatewayAdmin",
	HandlerType: (*AppUserGatewayAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdminApps",
			Handler:    _AppUserGatewayAdmin_CreateAdminApps_Handler,
		},
		{
			MethodName: "GetAdminApps",
			Handler:    _AppUserGatewayAdmin_GetAdminApps_Handler,
		},
		{
			MethodName: "CreateGenesisRole",
			Handler:    _AppUserGatewayAdmin_CreateGenesisRole_Handler,
		},
		{
			MethodName: "GetGenesisRole",
			Handler:    _AppUserGatewayAdmin_GetGenesisRole_Handler,
		},
		{
			MethodName: "GetAppGenesisAppRoleUsers",
			Handler:    _AppUserGatewayAdmin_GetAppGenesisAppRoleUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appusergw/admin/admin.proto",
}
