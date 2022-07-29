// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/app/app.proto
<<<<<<< HEAD
=======
// source: npool/appusermw/app/app.proto
>>>>>>> basically completed new apis:npool/appusermw/app/app_grpc.pb.go
=======
>>>>>>> Add enum ERR

package app

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

// AppMwClient is the client API for AppMw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppMwClient interface {
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error)
	GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	GetUserApps(ctx context.Context, in *GetUserAppsRequest, opts ...grpc.CallOption) (*GetUserAppsResponse, error)
	BanApp(ctx context.Context, in *BanAppRequest, opts ...grpc.CallOption) (*BanAppResponse, error)
	GetSignMethods(ctx context.Context, in *GetSignMethodsRequest, opts ...grpc.CallOption) (*GetSignMethodsResponse, error)
	SetSignMethods(ctx context.Context, in *SetSignMethodsRequest, opts ...grpc.CallOption) (*SetSignMethodsResponse, error)
	GetRecaptchas(ctx context.Context, in *GetRecaptchasRequest, opts ...grpc.CallOption) (*GetRecaptchasResponse, error)
	SetRecaptcha(ctx context.Context, in *SetRecaptchaRequest, opts ...grpc.CallOption) (*SetRecaptchaResponse, error)
	SetKyc(ctx context.Context, in *SetKycRequest, opts ...grpc.CallOption) (*SetKycResponse, error)
	SetSigninVerify(ctx context.Context, in *SetSigninVerifyRequest, opts ...grpc.CallOption) (*SetSigninVerifyResponse, error)
	SetInvitationCodeMust(ctx context.Context, in *SetInvitationCodeMustRequest, opts ...grpc.CallOption) (*SetInvitationCodeMustResponse, error)
}

type appMwClient struct {
	cc grpc.ClientConnInterface
}

func NewAppMwClient(cc grpc.ClientConnInterface) AppMwClient {
	return &appMwClient{cc}
}

func (c *appMwClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error) {
	out := new(UpdateAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/GetApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) GetUserApps(ctx context.Context, in *GetUserAppsRequest, opts ...grpc.CallOption) (*GetUserAppsResponse, error) {
	out := new(GetUserAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/GetUserApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) BanApp(ctx context.Context, in *BanAppRequest, opts ...grpc.CallOption) (*BanAppResponse, error) {
	out := new(BanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/BanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) GetSignMethods(ctx context.Context, in *GetSignMethodsRequest, opts ...grpc.CallOption) (*GetSignMethodsResponse, error) {
	out := new(GetSignMethodsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/GetSignMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) SetSignMethods(ctx context.Context, in *SetSignMethodsRequest, opts ...grpc.CallOption) (*SetSignMethodsResponse, error) {
	out := new(SetSignMethodsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/SetSignMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) GetRecaptchas(ctx context.Context, in *GetRecaptchasRequest, opts ...grpc.CallOption) (*GetRecaptchasResponse, error) {
	out := new(GetRecaptchasResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/GetRecaptchas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) SetRecaptcha(ctx context.Context, in *SetRecaptchaRequest, opts ...grpc.CallOption) (*SetRecaptchaResponse, error) {
	out := new(SetRecaptchaResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/SetRecaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) SetKyc(ctx context.Context, in *SetKycRequest, opts ...grpc.CallOption) (*SetKycResponse, error) {
	out := new(SetKycResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/SetKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) SetSigninVerify(ctx context.Context, in *SetSigninVerifyRequest, opts ...grpc.CallOption) (*SetSigninVerifyResponse, error) {
	out := new(SetSigninVerifyResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/SetSigninVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMwClient) SetInvitationCodeMust(ctx context.Context, in *SetInvitationCodeMustRequest, opts ...grpc.CallOption) (*SetInvitationCodeMustResponse, error) {
	out := new(SetInvitationCodeMustResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.app.v1.AppMw/SetInvitationCodeMust", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppMwServer is the server API for AppMw service.
// All implementations must embed UnimplementedAppMwServer
// for forward compatibility
type AppMwServer interface {
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error)
	GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error)
	GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	GetUserApps(context.Context, *GetUserAppsRequest) (*GetUserAppsResponse, error)
	BanApp(context.Context, *BanAppRequest) (*BanAppResponse, error)
	GetSignMethods(context.Context, *GetSignMethodsRequest) (*GetSignMethodsResponse, error)
	SetSignMethods(context.Context, *SetSignMethodsRequest) (*SetSignMethodsResponse, error)
	GetRecaptchas(context.Context, *GetRecaptchasRequest) (*GetRecaptchasResponse, error)
	SetRecaptcha(context.Context, *SetRecaptchaRequest) (*SetRecaptchaResponse, error)
	SetKyc(context.Context, *SetKycRequest) (*SetKycResponse, error)
	SetSigninVerify(context.Context, *SetSigninVerifyRequest) (*SetSigninVerifyResponse, error)
	SetInvitationCodeMust(context.Context, *SetInvitationCodeMustRequest) (*SetInvitationCodeMustResponse, error)
	mustEmbedUnimplementedAppMwServer()
}

// UnimplementedAppMwServer must be embedded to have forward compatible implementations.
type UnimplementedAppMwServer struct {
}

func (UnimplementedAppMwServer) CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedAppMwServer) UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (UnimplementedAppMwServer) GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (UnimplementedAppMwServer) GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedAppMwServer) GetUserApps(context.Context, *GetUserAppsRequest) (*GetUserAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserApps not implemented")
}
func (UnimplementedAppMwServer) BanApp(context.Context, *BanAppRequest) (*BanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanApp not implemented")
}
func (UnimplementedAppMwServer) GetSignMethods(context.Context, *GetSignMethodsRequest) (*GetSignMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignMethods not implemented")
}
func (UnimplementedAppMwServer) SetSignMethods(context.Context, *SetSignMethodsRequest) (*SetSignMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignMethods not implemented")
}
func (UnimplementedAppMwServer) GetRecaptchas(context.Context, *GetRecaptchasRequest) (*GetRecaptchasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecaptchas not implemented")
}
func (UnimplementedAppMwServer) SetRecaptcha(context.Context, *SetRecaptchaRequest) (*SetRecaptchaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRecaptcha not implemented")
}
func (UnimplementedAppMwServer) SetKyc(context.Context, *SetKycRequest) (*SetKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKyc not implemented")
}
func (UnimplementedAppMwServer) SetSigninVerify(context.Context, *SetSigninVerifyRequest) (*SetSigninVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSigninVerify not implemented")
}
func (UnimplementedAppMwServer) SetInvitationCodeMust(context.Context, *SetInvitationCodeMustRequest) (*SetInvitationCodeMustResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInvitationCodeMust not implemented")
}
func (UnimplementedAppMwServer) mustEmbedUnimplementedAppMwServer() {}

// UnsafeAppMwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppMwServer will
// result in compilation errors.
type UnsafeAppMwServer interface {
	mustEmbedUnimplementedAppMwServer()
}

func RegisterAppMwServer(s grpc.ServiceRegistrar, srv AppMwServer) {
	s.RegisterService(&AppMw_ServiceDesc, srv)
}

func _AppMw_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).UpdateApp(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).GetApp(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/GetApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).GetApps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_GetUserApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).GetUserApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/GetUserApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).GetUserApps(ctx, req.(*GetUserAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_BanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).BanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/BanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).BanApp(ctx, req.(*BanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_GetSignMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).GetSignMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/GetSignMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).GetSignMethods(ctx, req.(*GetSignMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_SetSignMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSignMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).SetSignMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/SetSignMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).SetSignMethods(ctx, req.(*SetSignMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_GetRecaptchas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecaptchasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).GetRecaptchas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/GetRecaptchas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).GetRecaptchas(ctx, req.(*GetRecaptchasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_SetRecaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRecaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).SetRecaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/SetRecaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).SetRecaptcha(ctx, req.(*SetRecaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_SetKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).SetKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/SetKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).SetKyc(ctx, req.(*SetKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_SetSigninVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSigninVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).SetSigninVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/SetSigninVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).SetSigninVerify(ctx, req.(*SetSigninVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMw_SetInvitationCodeMust_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInvitationCodeMustRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMwServer).SetInvitationCodeMust(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.app.v1.AppMw/SetInvitationCodeMust",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMwServer).SetInvitationCodeMust(ctx, req.(*SetInvitationCodeMustRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppMw_ServiceDesc is the grpc.ServiceDesc for AppMw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppMw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.app.v1.AppMw",
	HandlerType: (*AppMwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _AppMw_CreateApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _AppMw_UpdateApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _AppMw_GetApp_Handler,
		},
		{
			MethodName: "GetApps",
			Handler:    _AppMw_GetApps_Handler,
		},
		{
			MethodName: "GetUserApps",
			Handler:    _AppMw_GetUserApps_Handler,
		},
		{
			MethodName: "BanApp",
			Handler:    _AppMw_BanApp_Handler,
		},
		{
			MethodName: "GetSignMethods",
			Handler:    _AppMw_GetSignMethods_Handler,
		},
		{
			MethodName: "SetSignMethods",
			Handler:    _AppMw_SetSignMethods_Handler,
		},
		{
			MethodName: "GetRecaptchas",
			Handler:    _AppMw_GetRecaptchas_Handler,
		},
		{
			MethodName: "SetRecaptcha",
			Handler:    _AppMw_SetRecaptcha_Handler,
		},
		{
			MethodName: "SetKyc",
			Handler:    _AppMw_SetKyc_Handler,
		},
		{
			MethodName: "SetSigninVerify",
			Handler:    _AppMw_SetSigninVerify_Handler,
		},
		{
			MethodName: "SetInvitationCodeMust",
			Handler:    _AppMw_SetInvitationCodeMust_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/app/app.proto",
<<<<<<< HEAD
=======
	Metadata: "npool/appusermw/app/app.proto",
>>>>>>> basically completed new apis:npool/appusermw/app/app_grpc.pb.go
=======
>>>>>>> Add enum ERR
}
