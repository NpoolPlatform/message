// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/appuser/appuser.proto

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

// AppUserMgrClient is the client API for AppUserMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserMgrClient interface {
	CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error)
	CreateAppUsers(ctx context.Context, in *CreateAppUsersRequest, opts ...grpc.CallOption) (*CreateAppUsersResponse, error)
	UpdateAppUser(ctx context.Context, in *UpdateAppUserRequest, opts ...grpc.CallOption) (*UpdateAppUserResponse, error)
	GetAppUser(ctx context.Context, in *GetAppUserRequest, opts ...grpc.CallOption) (*GetAppUserResponse, error)
	GetAppUserOnly(ctx context.Context, in *GetAppUserOnlyRequest, opts ...grpc.CallOption) (*GetAppUserOnlyResponse, error)
	GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error)
	ExistAppUser(ctx context.Context, in *ExistAppUserRequest, opts ...grpc.CallOption) (*ExistAppUserResponse, error)
	ExistAppUserConds(ctx context.Context, in *ExistAppUserCondsRequest, opts ...grpc.CallOption) (*ExistAppUserCondsResponse, error)
	CountAppUsers(ctx context.Context, in *CountAppUsersRequest, opts ...grpc.CallOption) (*CountAppUsersResponse, error)
	DeleteAppUser(ctx context.Context, in *DeleteAppUserRequest, opts ...grpc.CallOption) (*DeleteAppUserResponse, error)
}

type appUserMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserMgrClient(cc grpc.ClientConnInterface) AppUserMgrClient {
	return &appUserMgrClient{cc}
}

func (c *appUserMgrClient) CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error) {
	out := new(CreateAppUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/CreateAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) CreateAppUsers(ctx context.Context, in *CreateAppUsersRequest, opts ...grpc.CallOption) (*CreateAppUsersResponse, error) {
	out := new(CreateAppUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/CreateAppUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) UpdateAppUser(ctx context.Context, in *UpdateAppUserRequest, opts ...grpc.CallOption) (*UpdateAppUserResponse, error) {
	out := new(UpdateAppUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/UpdateAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) GetAppUser(ctx context.Context, in *GetAppUserRequest, opts ...grpc.CallOption) (*GetAppUserResponse, error) {
	out := new(GetAppUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/GetAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) GetAppUserOnly(ctx context.Context, in *GetAppUserOnlyRequest, opts ...grpc.CallOption) (*GetAppUserOnlyResponse, error) {
	out := new(GetAppUserOnlyResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/GetAppUserOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error) {
	out := new(GetAppUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/GetAppUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) ExistAppUser(ctx context.Context, in *ExistAppUserRequest, opts ...grpc.CallOption) (*ExistAppUserResponse, error) {
	out := new(ExistAppUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/ExistAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) ExistAppUserConds(ctx context.Context, in *ExistAppUserCondsRequest, opts ...grpc.CallOption) (*ExistAppUserCondsResponse, error) {
	out := new(ExistAppUserCondsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/ExistAppUserConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) CountAppUsers(ctx context.Context, in *CountAppUsersRequest, opts ...grpc.CallOption) (*CountAppUsersResponse, error) {
	out := new(CountAppUsersResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/CountAppUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMgrClient) DeleteAppUser(ctx context.Context, in *DeleteAppUserRequest, opts ...grpc.CallOption) (*DeleteAppUserResponse, error) {
	out := new(DeleteAppUserResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.appuser.v2.AppUserMgr/DeleteAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserMgrServer is the server API for AppUserMgr service.
// All implementations must embed UnimplementedAppUserMgrServer
// for forward compatibility
type AppUserMgrServer interface {
	CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error)
	CreateAppUsers(context.Context, *CreateAppUsersRequest) (*CreateAppUsersResponse, error)
	UpdateAppUser(context.Context, *UpdateAppUserRequest) (*UpdateAppUserResponse, error)
	GetAppUser(context.Context, *GetAppUserRequest) (*GetAppUserResponse, error)
	GetAppUserOnly(context.Context, *GetAppUserOnlyRequest) (*GetAppUserOnlyResponse, error)
	GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error)
	ExistAppUser(context.Context, *ExistAppUserRequest) (*ExistAppUserResponse, error)
	ExistAppUserConds(context.Context, *ExistAppUserCondsRequest) (*ExistAppUserCondsResponse, error)
	CountAppUsers(context.Context, *CountAppUsersRequest) (*CountAppUsersResponse, error)
	DeleteAppUser(context.Context, *DeleteAppUserRequest) (*DeleteAppUserResponse, error)
	mustEmbedUnimplementedAppUserMgrServer()
}

// UnimplementedAppUserMgrServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserMgrServer struct {
}

func (UnimplementedAppUserMgrServer) CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUser not implemented")
}
func (UnimplementedAppUserMgrServer) CreateAppUsers(context.Context, *CreateAppUsersRequest) (*CreateAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUsers not implemented")
}
func (UnimplementedAppUserMgrServer) UpdateAppUser(context.Context, *UpdateAppUserRequest) (*UpdateAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppUser not implemented")
}
func (UnimplementedAppUserMgrServer) GetAppUser(context.Context, *GetAppUserRequest) (*GetAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUser not implemented")
}
func (UnimplementedAppUserMgrServer) GetAppUserOnly(context.Context, *GetAppUserOnlyRequest) (*GetAppUserOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserOnly not implemented")
}
func (UnimplementedAppUserMgrServer) GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUsers not implemented")
}
func (UnimplementedAppUserMgrServer) ExistAppUser(context.Context, *ExistAppUserRequest) (*ExistAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppUser not implemented")
}
func (UnimplementedAppUserMgrServer) ExistAppUserConds(context.Context, *ExistAppUserCondsRequest) (*ExistAppUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppUserConds not implemented")
}
func (UnimplementedAppUserMgrServer) CountAppUsers(context.Context, *CountAppUsersRequest) (*CountAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAppUsers not implemented")
}
func (UnimplementedAppUserMgrServer) DeleteAppUser(context.Context, *DeleteAppUserRequest) (*DeleteAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppUser not implemented")
}
func (UnimplementedAppUserMgrServer) mustEmbedUnimplementedAppUserMgrServer() {}

// UnsafeAppUserMgrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserMgrServer will
// result in compilation errors.
type UnsafeAppUserMgrServer interface {
	mustEmbedUnimplementedAppUserMgrServer()
}

func RegisterAppUserMgrServer(s grpc.ServiceRegistrar, srv AppUserMgrServer) {
	s.RegisterService(&AppUserMgr_ServiceDesc, srv)
}

func _AppUserMgr_CreateAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).CreateAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/CreateAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).CreateAppUser(ctx, req.(*CreateAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_CreateAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).CreateAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/CreateAppUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).CreateAppUsers(ctx, req.(*CreateAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_UpdateAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).UpdateAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/UpdateAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).UpdateAppUser(ctx, req.(*UpdateAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_GetAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).GetAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/GetAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).GetAppUser(ctx, req.(*GetAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_GetAppUserOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).GetAppUserOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/GetAppUserOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).GetAppUserOnly(ctx, req.(*GetAppUserOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_GetAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).GetAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/GetAppUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).GetAppUsers(ctx, req.(*GetAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_ExistAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).ExistAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/ExistAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).ExistAppUser(ctx, req.(*ExistAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_ExistAppUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).ExistAppUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/ExistAppUserConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).ExistAppUserConds(ctx, req.(*ExistAppUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_CountAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).CountAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/CountAppUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).CountAppUsers(ctx, req.(*CountAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMgr_DeleteAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMgrServer).DeleteAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.appuser.v2.AppUserMgr/DeleteAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMgrServer).DeleteAppUser(ctx, req.(*DeleteAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserMgr_ServiceDesc is the grpc.ServiceDesc for AppUserMgr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserMgr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.manager.appuser.v2.AppUserMgr",
	HandlerType: (*AppUserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppUser",
			Handler:    _AppUserMgr_CreateAppUser_Handler,
		},
		{
			MethodName: "CreateAppUsers",
			Handler:    _AppUserMgr_CreateAppUsers_Handler,
		},
		{
			MethodName: "UpdateAppUser",
			Handler:    _AppUserMgr_UpdateAppUser_Handler,
		},
		{
			MethodName: "GetAppUser",
			Handler:    _AppUserMgr_GetAppUser_Handler,
		},
		{
			MethodName: "GetAppUserOnly",
			Handler:    _AppUserMgr_GetAppUserOnly_Handler,
		},
		{
			MethodName: "GetAppUsers",
			Handler:    _AppUserMgr_GetAppUsers_Handler,
		},
		{
			MethodName: "ExistAppUser",
			Handler:    _AppUserMgr_ExistAppUser_Handler,
		},
		{
			MethodName: "ExistAppUserConds",
			Handler:    _AppUserMgr_ExistAppUserConds_Handler,
		},
		{
			MethodName: "CountAppUsers",
			Handler:    _AppUserMgr_CountAppUsers_Handler,
		},
		{
			MethodName: "DeleteAppUser",
			Handler:    _AppUserMgr_DeleteAppUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/appuser/appuser.proto",
}
