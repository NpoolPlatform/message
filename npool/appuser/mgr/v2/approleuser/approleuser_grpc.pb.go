// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/approleuser/approleuser.proto

package approleuser

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

// AppRoleUserMgrClient is the client API for AppRoleUserMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppRoleUserMgrClient interface {
	CreateAppRoleUser(ctx context.Context, in *CreateAppRoleUserRequest, opts ...grpc.CallOption) (*CreateAppRoleUserResponse, error)
	CreateAppRoleUsers(ctx context.Context, in *CreateAppRoleUsersRequest, opts ...grpc.CallOption) (*CreateAppRoleUsersResponse, error)
	UpdateAppRoleUser(ctx context.Context, in *UpdateAppRoleUserRequest, opts ...grpc.CallOption) (*UpdateAppRoleUserResponse, error)
	GetAppRoleUser(ctx context.Context, in *GetAppRoleUserRequest, opts ...grpc.CallOption) (*GetAppRoleUserResponse, error)
	GetAppRoleUserOnly(ctx context.Context, in *GetAppRoleUserOnlyRequest, opts ...grpc.CallOption) (*GetAppRoleUserOnlyResponse, error)
	GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error)
	ExistAppRoleUser(ctx context.Context, in *ExistAppRoleUserRequest, opts ...grpc.CallOption) (*ExistAppRoleUserResponse, error)
	ExistAppRoleUserConds(ctx context.Context, in *ExistAppRoleUserCondsRequest, opts ...grpc.CallOption) (*ExistAppRoleUserCondsResponse, error)
	CountAppRoleUsers(ctx context.Context, in *CountAppRoleUsersRequest, opts ...grpc.CallOption) (*CountAppRoleUsersResponse, error)
	DeleteAppRoleUser(ctx context.Context, in *DeleteAppRoleUserRequest, opts ...grpc.CallOption) (*DeleteAppRoleUserResponse, error)
}

type appRoleUserMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewAppRoleUserMgrClient(cc grpc.ClientConnInterface) AppRoleUserMgrClient {
	return &appRoleUserMgrClient{cc}
}

func (c *appRoleUserMgrClient) CreateAppRoleUser(ctx context.Context, in *CreateAppRoleUserRequest, opts ...grpc.CallOption) (*CreateAppRoleUserResponse, error) {
	out := new(CreateAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/CreateAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) CreateAppRoleUsers(ctx context.Context, in *CreateAppRoleUsersRequest, opts ...grpc.CallOption) (*CreateAppRoleUsersResponse, error) {
	out := new(CreateAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/CreateAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) UpdateAppRoleUser(ctx context.Context, in *UpdateAppRoleUserRequest, opts ...grpc.CallOption) (*UpdateAppRoleUserResponse, error) {
	out := new(UpdateAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/UpdateAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) GetAppRoleUser(ctx context.Context, in *GetAppRoleUserRequest, opts ...grpc.CallOption) (*GetAppRoleUserResponse, error) {
	out := new(GetAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) GetAppRoleUserOnly(ctx context.Context, in *GetAppRoleUserOnlyRequest, opts ...grpc.CallOption) (*GetAppRoleUserOnlyResponse, error) {
	out := new(GetAppRoleUserOnlyResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUserOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) GetAppRoleUsers(ctx context.Context, in *GetAppRoleUsersRequest, opts ...grpc.CallOption) (*GetAppRoleUsersResponse, error) {
	out := new(GetAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) ExistAppRoleUser(ctx context.Context, in *ExistAppRoleUserRequest, opts ...grpc.CallOption) (*ExistAppRoleUserResponse, error) {
	out := new(ExistAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/ExistAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) ExistAppRoleUserConds(ctx context.Context, in *ExistAppRoleUserCondsRequest, opts ...grpc.CallOption) (*ExistAppRoleUserCondsResponse, error) {
	out := new(ExistAppRoleUserCondsResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/ExistAppRoleUserConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) CountAppRoleUsers(ctx context.Context, in *CountAppRoleUsersRequest, opts ...grpc.CallOption) (*CountAppRoleUsersResponse, error) {
	out := new(CountAppRoleUsersResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/CountAppRoleUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRoleUserMgrClient) DeleteAppRoleUser(ctx context.Context, in *DeleteAppRoleUserRequest, opts ...grpc.CallOption) (*DeleteAppRoleUserResponse, error) {
	out := new(DeleteAppRoleUserResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.approleuser.v2.AppRoleUserMgr/DeleteAppRoleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppRoleUserMgrServer is the server API for AppRoleUserMgr service.
// All implementations must embed UnimplementedAppRoleUserMgrServer
// for forward compatibility
type AppRoleUserMgrServer interface {
	CreateAppRoleUser(context.Context, *CreateAppRoleUserRequest) (*CreateAppRoleUserResponse, error)
	CreateAppRoleUsers(context.Context, *CreateAppRoleUsersRequest) (*CreateAppRoleUsersResponse, error)
	UpdateAppRoleUser(context.Context, *UpdateAppRoleUserRequest) (*UpdateAppRoleUserResponse, error)
	GetAppRoleUser(context.Context, *GetAppRoleUserRequest) (*GetAppRoleUserResponse, error)
	GetAppRoleUserOnly(context.Context, *GetAppRoleUserOnlyRequest) (*GetAppRoleUserOnlyResponse, error)
	GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error)
	ExistAppRoleUser(context.Context, *ExistAppRoleUserRequest) (*ExistAppRoleUserResponse, error)
	ExistAppRoleUserConds(context.Context, *ExistAppRoleUserCondsRequest) (*ExistAppRoleUserCondsResponse, error)
	CountAppRoleUsers(context.Context, *CountAppRoleUsersRequest) (*CountAppRoleUsersResponse, error)
	DeleteAppRoleUser(context.Context, *DeleteAppRoleUserRequest) (*DeleteAppRoleUserResponse, error)
	mustEmbedUnimplementedAppRoleUserMgrServer()
}

// UnimplementedAppRoleUserMgrServer must be embedded to have forward compatible implementations.
type UnimplementedAppRoleUserMgrServer struct {
}

func (UnimplementedAppRoleUserMgrServer) CreateAppRoleUser(context.Context, *CreateAppRoleUserRequest) (*CreateAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppRoleUser not implemented")
}
func (UnimplementedAppRoleUserMgrServer) CreateAppRoleUsers(context.Context, *CreateAppRoleUsersRequest) (*CreateAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppRoleUsers not implemented")
}
func (UnimplementedAppRoleUserMgrServer) UpdateAppRoleUser(context.Context, *UpdateAppRoleUserRequest) (*UpdateAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppRoleUser not implemented")
}
func (UnimplementedAppRoleUserMgrServer) GetAppRoleUser(context.Context, *GetAppRoleUserRequest) (*GetAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUser not implemented")
}
func (UnimplementedAppRoleUserMgrServer) GetAppRoleUserOnly(context.Context, *GetAppRoleUserOnlyRequest) (*GetAppRoleUserOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUserOnly not implemented")
}
func (UnimplementedAppRoleUserMgrServer) GetAppRoleUsers(context.Context, *GetAppRoleUsersRequest) (*GetAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRoleUsers not implemented")
}
func (UnimplementedAppRoleUserMgrServer) ExistAppRoleUser(context.Context, *ExistAppRoleUserRequest) (*ExistAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppRoleUser not implemented")
}
func (UnimplementedAppRoleUserMgrServer) ExistAppRoleUserConds(context.Context, *ExistAppRoleUserCondsRequest) (*ExistAppRoleUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppRoleUserConds not implemented")
}
func (UnimplementedAppRoleUserMgrServer) CountAppRoleUsers(context.Context, *CountAppRoleUsersRequest) (*CountAppRoleUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAppRoleUsers not implemented")
}
func (UnimplementedAppRoleUserMgrServer) DeleteAppRoleUser(context.Context, *DeleteAppRoleUserRequest) (*DeleteAppRoleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppRoleUser not implemented")
}
func (UnimplementedAppRoleUserMgrServer) mustEmbedUnimplementedAppRoleUserMgrServer() {}

// UnsafeAppRoleUserMgrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppRoleUserMgrServer will
// result in compilation errors.
type UnsafeAppRoleUserMgrServer interface {
	mustEmbedUnimplementedAppRoleUserMgrServer()
}

func RegisterAppRoleUserMgrServer(s grpc.ServiceRegistrar, srv AppRoleUserMgrServer) {
	s.RegisterService(&AppRoleUserMgr_ServiceDesc, srv)
}

func _AppRoleUserMgr_CreateAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).CreateAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/CreateAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).CreateAppRoleUser(ctx, req.(*CreateAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_CreateAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).CreateAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/CreateAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).CreateAppRoleUsers(ctx, req.(*CreateAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_UpdateAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).UpdateAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/UpdateAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).UpdateAppRoleUser(ctx, req.(*UpdateAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_GetAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).GetAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).GetAppRoleUser(ctx, req.(*GetAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_GetAppRoleUserOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRoleUserOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).GetAppRoleUserOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUserOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).GetAppRoleUserOnly(ctx, req.(*GetAppRoleUserOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_GetAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).GetAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/GetAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).GetAppRoleUsers(ctx, req.(*GetAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_ExistAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).ExistAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/ExistAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).ExistAppRoleUser(ctx, req.(*ExistAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_ExistAppRoleUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppRoleUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).ExistAppRoleUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/ExistAppRoleUserConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).ExistAppRoleUserConds(ctx, req.(*ExistAppRoleUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_CountAppRoleUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAppRoleUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).CountAppRoleUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/CountAppRoleUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).CountAppRoleUsers(ctx, req.(*CountAppRoleUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRoleUserMgr_DeleteAppRoleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRoleUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRoleUserMgrServer).DeleteAppRoleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.approleuser.v2.AppRoleUserMgr/DeleteAppRoleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRoleUserMgrServer).DeleteAppRoleUser(ctx, req.(*DeleteAppRoleUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppRoleUserMgr_ServiceDesc is the grpc.ServiceDesc for AppRoleUserMgr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppRoleUserMgr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.manager.approleuser.v2.AppRoleUserMgr",
	HandlerType: (*AppRoleUserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppRoleUser",
			Handler:    _AppRoleUserMgr_CreateAppRoleUser_Handler,
		},
		{
			MethodName: "CreateAppRoleUsers",
			Handler:    _AppRoleUserMgr_CreateAppRoleUsers_Handler,
		},
		{
			MethodName: "UpdateAppRoleUser",
			Handler:    _AppRoleUserMgr_UpdateAppRoleUser_Handler,
		},
		{
			MethodName: "GetAppRoleUser",
			Handler:    _AppRoleUserMgr_GetAppRoleUser_Handler,
		},
		{
			MethodName: "GetAppRoleUserOnly",
			Handler:    _AppRoleUserMgr_GetAppRoleUserOnly_Handler,
		},
		{
			MethodName: "GetAppRoleUsers",
			Handler:    _AppRoleUserMgr_GetAppRoleUsers_Handler,
		},
		{
			MethodName: "ExistAppRoleUser",
			Handler:    _AppRoleUserMgr_ExistAppRoleUser_Handler,
		},
		{
			MethodName: "ExistAppRoleUserConds",
			Handler:    _AppRoleUserMgr_ExistAppRoleUserConds_Handler,
		},
		{
			MethodName: "CountAppRoleUsers",
			Handler:    _AppRoleUserMgr_CountAppRoleUsers_Handler,
		},
		{
			MethodName: "DeleteAppRoleUser",
			Handler:    _AppRoleUserMgr_DeleteAppRoleUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/approleuser/approleuser.proto",
}
