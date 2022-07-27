// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/appusersecret/appusersecret.proto

package appusersecret

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

// AppUserSecretMgrClient is the client API for AppUserSecretMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserSecretMgrClient interface {
	CreateAppUserSecret(ctx context.Context, in *CreateAppUserSecretRequest, opts ...grpc.CallOption) (*CreateAppUserSecretResponse, error)
	CreateAppUserSecrets(ctx context.Context, in *CreateAppUserSecretsRequest, opts ...grpc.CallOption) (*CreateAppUserSecretsResponse, error)
	UpdateAppUserSecret(ctx context.Context, in *UpdateAppUserSecretRequest, opts ...grpc.CallOption) (*UpdateAppUserSecretResponse, error)
	GetAppUserSecret(ctx context.Context, in *GetAppUserSecretRequest, opts ...grpc.CallOption) (*GetAppUserSecretResponse, error)
	GetAppUserSecretOnly(ctx context.Context, in *GetAppUserSecretOnlyRequest, opts ...grpc.CallOption) (*GetAppUserSecretOnlyResponse, error)
	GetAppUserSecrets(ctx context.Context, in *GetAppUserSecretsRequest, opts ...grpc.CallOption) (*GetAppUserSecretsResponse, error)
	ExistAppUserSecret(ctx context.Context, in *ExistAppUserSecretRequest, opts ...grpc.CallOption) (*ExistAppUserSecretResponse, error)
	ExistAppUserSecretConds(ctx context.Context, in *ExistAppUserSecretCondsRequest, opts ...grpc.CallOption) (*ExistAppUserSecretCondsResponse, error)
	CountAppUserSecrets(ctx context.Context, in *CountAppUserSecretsRequest, opts ...grpc.CallOption) (*CountAppUserSecretsResponse, error)
	DeleteAppUserSecret(ctx context.Context, in *DeleteAppUserSecretRequest, opts ...grpc.CallOption) (*DeleteAppUserSecretResponse, error)
}

type appUserSecretMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserSecretMgrClient(cc grpc.ClientConnInterface) AppUserSecretMgrClient {
	return &appUserSecretMgrClient{cc}
}

func (c *appUserSecretMgrClient) CreateAppUserSecret(ctx context.Context, in *CreateAppUserSecretRequest, opts ...grpc.CallOption) (*CreateAppUserSecretResponse, error) {
	out := new(CreateAppUserSecretResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CreateAppUserSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) CreateAppUserSecrets(ctx context.Context, in *CreateAppUserSecretsRequest, opts ...grpc.CallOption) (*CreateAppUserSecretsResponse, error) {
	out := new(CreateAppUserSecretsResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CreateAppUserSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) UpdateAppUserSecret(ctx context.Context, in *UpdateAppUserSecretRequest, opts ...grpc.CallOption) (*UpdateAppUserSecretResponse, error) {
	out := new(UpdateAppUserSecretResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/UpdateAppUserSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) GetAppUserSecret(ctx context.Context, in *GetAppUserSecretRequest, opts ...grpc.CallOption) (*GetAppUserSecretResponse, error) {
	out := new(GetAppUserSecretResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) GetAppUserSecretOnly(ctx context.Context, in *GetAppUserSecretOnlyRequest, opts ...grpc.CallOption) (*GetAppUserSecretOnlyResponse, error) {
	out := new(GetAppUserSecretOnlyResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecretOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) GetAppUserSecrets(ctx context.Context, in *GetAppUserSecretsRequest, opts ...grpc.CallOption) (*GetAppUserSecretsResponse, error) {
	out := new(GetAppUserSecretsResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) ExistAppUserSecret(ctx context.Context, in *ExistAppUserSecretRequest, opts ...grpc.CallOption) (*ExistAppUserSecretResponse, error) {
	out := new(ExistAppUserSecretResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/ExistAppUserSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) ExistAppUserSecretConds(ctx context.Context, in *ExistAppUserSecretCondsRequest, opts ...grpc.CallOption) (*ExistAppUserSecretCondsResponse, error) {
	out := new(ExistAppUserSecretCondsResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/ExistAppUserSecretConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) CountAppUserSecrets(ctx context.Context, in *CountAppUserSecretsRequest, opts ...grpc.CallOption) (*CountAppUserSecretsResponse, error) {
	out := new(CountAppUserSecretsResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CountAppUserSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserSecretMgrClient) DeleteAppUserSecret(ctx context.Context, in *DeleteAppUserSecretRequest, opts ...grpc.CallOption) (*DeleteAppUserSecretResponse, error) {
	out := new(DeleteAppUserSecretResponse)
	err := c.cc.Invoke(ctx, "/app.user.manager.appusersecret.v2.AppUserSecretMgr/DeleteAppUserSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserSecretMgrServer is the server API for AppUserSecretMgr service.
// All implementations must embed UnimplementedAppUserSecretMgrServer
// for forward compatibility
type AppUserSecretMgrServer interface {
	CreateAppUserSecret(context.Context, *CreateAppUserSecretRequest) (*CreateAppUserSecretResponse, error)
	CreateAppUserSecrets(context.Context, *CreateAppUserSecretsRequest) (*CreateAppUserSecretsResponse, error)
	UpdateAppUserSecret(context.Context, *UpdateAppUserSecretRequest) (*UpdateAppUserSecretResponse, error)
	GetAppUserSecret(context.Context, *GetAppUserSecretRequest) (*GetAppUserSecretResponse, error)
	GetAppUserSecretOnly(context.Context, *GetAppUserSecretOnlyRequest) (*GetAppUserSecretOnlyResponse, error)
	GetAppUserSecrets(context.Context, *GetAppUserSecretsRequest) (*GetAppUserSecretsResponse, error)
	ExistAppUserSecret(context.Context, *ExistAppUserSecretRequest) (*ExistAppUserSecretResponse, error)
	ExistAppUserSecretConds(context.Context, *ExistAppUserSecretCondsRequest) (*ExistAppUserSecretCondsResponse, error)
	CountAppUserSecrets(context.Context, *CountAppUserSecretsRequest) (*CountAppUserSecretsResponse, error)
	DeleteAppUserSecret(context.Context, *DeleteAppUserSecretRequest) (*DeleteAppUserSecretResponse, error)
	mustEmbedUnimplementedAppUserSecretMgrServer()
}

// UnimplementedAppUserSecretMgrServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserSecretMgrServer struct {
}

func (UnimplementedAppUserSecretMgrServer) CreateAppUserSecret(context.Context, *CreateAppUserSecretRequest) (*CreateAppUserSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUserSecret not implemented")
}
func (UnimplementedAppUserSecretMgrServer) CreateAppUserSecrets(context.Context, *CreateAppUserSecretsRequest) (*CreateAppUserSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUserSecrets not implemented")
}
func (UnimplementedAppUserSecretMgrServer) UpdateAppUserSecret(context.Context, *UpdateAppUserSecretRequest) (*UpdateAppUserSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppUserSecret not implemented")
}
func (UnimplementedAppUserSecretMgrServer) GetAppUserSecret(context.Context, *GetAppUserSecretRequest) (*GetAppUserSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserSecret not implemented")
}
func (UnimplementedAppUserSecretMgrServer) GetAppUserSecretOnly(context.Context, *GetAppUserSecretOnlyRequest) (*GetAppUserSecretOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserSecretOnly not implemented")
}
func (UnimplementedAppUserSecretMgrServer) GetAppUserSecrets(context.Context, *GetAppUserSecretsRequest) (*GetAppUserSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserSecrets not implemented")
}
func (UnimplementedAppUserSecretMgrServer) ExistAppUserSecret(context.Context, *ExistAppUserSecretRequest) (*ExistAppUserSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppUserSecret not implemented")
}
func (UnimplementedAppUserSecretMgrServer) ExistAppUserSecretConds(context.Context, *ExistAppUserSecretCondsRequest) (*ExistAppUserSecretCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAppUserSecretConds not implemented")
}
func (UnimplementedAppUserSecretMgrServer) CountAppUserSecrets(context.Context, *CountAppUserSecretsRequest) (*CountAppUserSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAppUserSecrets not implemented")
}
func (UnimplementedAppUserSecretMgrServer) DeleteAppUserSecret(context.Context, *DeleteAppUserSecretRequest) (*DeleteAppUserSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppUserSecret not implemented")
}
func (UnimplementedAppUserSecretMgrServer) mustEmbedUnimplementedAppUserSecretMgrServer() {}

// UnsafeAppUserSecretMgrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserSecretMgrServer will
// result in compilation errors.
type UnsafeAppUserSecretMgrServer interface {
	mustEmbedUnimplementedAppUserSecretMgrServer()
}

func RegisterAppUserSecretMgrServer(s grpc.ServiceRegistrar, srv AppUserSecretMgrServer) {
	s.RegisterService(&AppUserSecretMgr_ServiceDesc, srv)
}

func _AppUserSecretMgr_CreateAppUserSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).CreateAppUserSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CreateAppUserSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).CreateAppUserSecret(ctx, req.(*CreateAppUserSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_CreateAppUserSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).CreateAppUserSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CreateAppUserSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).CreateAppUserSecrets(ctx, req.(*CreateAppUserSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_UpdateAppUserSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppUserSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).UpdateAppUserSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/UpdateAppUserSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).UpdateAppUserSecret(ctx, req.(*UpdateAppUserSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_GetAppUserSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).GetAppUserSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).GetAppUserSecret(ctx, req.(*GetAppUserSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_GetAppUserSecretOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserSecretOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).GetAppUserSecretOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecretOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).GetAppUserSecretOnly(ctx, req.(*GetAppUserSecretOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_GetAppUserSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).GetAppUserSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/GetAppUserSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).GetAppUserSecrets(ctx, req.(*GetAppUserSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_ExistAppUserSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppUserSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).ExistAppUserSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/ExistAppUserSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).ExistAppUserSecret(ctx, req.(*ExistAppUserSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_ExistAppUserSecretConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAppUserSecretCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).ExistAppUserSecretConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/ExistAppUserSecretConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).ExistAppUserSecretConds(ctx, req.(*ExistAppUserSecretCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_CountAppUserSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAppUserSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).CountAppUserSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/CountAppUserSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).CountAppUserSecrets(ctx, req.(*CountAppUserSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserSecretMgr_DeleteAppUserSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppUserSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserSecretMgrServer).DeleteAppUserSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.manager.appusersecret.v2.AppUserSecretMgr/DeleteAppUserSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserSecretMgrServer).DeleteAppUserSecret(ctx, req.(*DeleteAppUserSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserSecretMgr_ServiceDesc is the grpc.ServiceDesc for AppUserSecretMgr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserSecretMgr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.manager.appusersecret.v2.AppUserSecretMgr",
	HandlerType: (*AppUserSecretMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppUserSecret",
			Handler:    _AppUserSecretMgr_CreateAppUserSecret_Handler,
		},
		{
			MethodName: "CreateAppUserSecrets",
			Handler:    _AppUserSecretMgr_CreateAppUserSecrets_Handler,
		},
		{
			MethodName: "UpdateAppUserSecret",
			Handler:    _AppUserSecretMgr_UpdateAppUserSecret_Handler,
		},
		{
			MethodName: "GetAppUserSecret",
			Handler:    _AppUserSecretMgr_GetAppUserSecret_Handler,
		},
		{
			MethodName: "GetAppUserSecretOnly",
			Handler:    _AppUserSecretMgr_GetAppUserSecretOnly_Handler,
		},
		{
			MethodName: "GetAppUserSecrets",
			Handler:    _AppUserSecretMgr_GetAppUserSecrets_Handler,
		},
		{
			MethodName: "ExistAppUserSecret",
			Handler:    _AppUserSecretMgr_ExistAppUserSecret_Handler,
		},
		{
			MethodName: "ExistAppUserSecretConds",
			Handler:    _AppUserSecretMgr_ExistAppUserSecretConds_Handler,
		},
		{
			MethodName: "CountAppUserSecrets",
			Handler:    _AppUserSecretMgr_CountAppUserSecrets_Handler,
		},
		{
			MethodName: "DeleteAppUserSecret",
			Handler:    _AppUserSecretMgr_DeleteAppUserSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/appusersecret/appusersecret.proto",
}
