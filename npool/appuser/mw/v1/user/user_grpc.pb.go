// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// AppUserMiddlewareUserClient is the client API for AppUserMiddlewareUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserMiddlewareUserClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetAppUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserInfos(ctx context.Context, in *GetUserInfosRequest, opts ...grpc.CallOption) (*GetUserInfosResponse, error)
	CreateUserWithSecret(ctx context.Context, in *CreateUserWithSecretRequest, opts ...grpc.CallOption) (*CreateUserWithSecretResponse, error)
	CreateUserWithThirdParty(ctx context.Context, in *CreateUserWithThirdPartyRequest, opts ...grpc.CallOption) (*CreateUserWithThirdPartyResponse, error)
}

type appUserMiddlewareUserClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserMiddlewareUserClient(cc grpc.ClientConnInterface) AppUserMiddlewareUserClient {
	return &appUserMiddlewareUserClient{cc}
}

func (c *appUserMiddlewareUserClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/appusermw.user.v2.AppUserMiddlewareUser/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareUserClient) GetAppUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/appusermw.user.v2.AppUserMiddlewareUser/GetAppUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareUserClient) GetUserInfos(ctx context.Context, in *GetUserInfosRequest, opts ...grpc.CallOption) (*GetUserInfosResponse, error) {
	out := new(GetUserInfosResponse)
	err := c.cc.Invoke(ctx, "/appusermw.user.v2.AppUserMiddlewareUser/GetUserInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareUserClient) CreateUserWithSecret(ctx context.Context, in *CreateUserWithSecretRequest, opts ...grpc.CallOption) (*CreateUserWithSecretResponse, error) {
	out := new(CreateUserWithSecretResponse)
	err := c.cc.Invoke(ctx, "/appusermw.user.v2.AppUserMiddlewareUser/CreateUserWithSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareUserClient) CreateUserWithThirdParty(ctx context.Context, in *CreateUserWithThirdPartyRequest, opts ...grpc.CallOption) (*CreateUserWithThirdPartyResponse, error) {
	out := new(CreateUserWithThirdPartyResponse)
	err := c.cc.Invoke(ctx, "/appusermw.user.v2.AppUserMiddlewareUser/CreateUserWithThirdParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserMiddlewareUserServer is the server API for AppUserMiddlewareUser service.
// All implementations must embed UnimplementedAppUserMiddlewareUserServer
// for forward compatibility
type AppUserMiddlewareUserServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetAppUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserInfos(context.Context, *GetUserInfosRequest) (*GetUserInfosResponse, error)
	CreateUserWithSecret(context.Context, *CreateUserWithSecretRequest) (*CreateUserWithSecretResponse, error)
	CreateUserWithThirdParty(context.Context, *CreateUserWithThirdPartyRequest) (*CreateUserWithThirdPartyResponse, error)
	mustEmbedUnimplementedAppUserMiddlewareUserServer()
}

// UnimplementedAppUserMiddlewareUserServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserMiddlewareUserServer struct {
}

func (UnimplementedAppUserMiddlewareUserServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedAppUserMiddlewareUserServer) GetAppUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserInfo not implemented")
}
func (UnimplementedAppUserMiddlewareUserServer) GetUserInfos(context.Context, *GetUserInfosRequest) (*GetUserInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfos not implemented")
}
func (UnimplementedAppUserMiddlewareUserServer) CreateUserWithSecret(context.Context, *CreateUserWithSecretRequest) (*CreateUserWithSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserWithSecret not implemented")
}
func (UnimplementedAppUserMiddlewareUserServer) CreateUserWithThirdParty(context.Context, *CreateUserWithThirdPartyRequest) (*CreateUserWithThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserWithThirdParty not implemented")
}
func (UnimplementedAppUserMiddlewareUserServer) mustEmbedUnimplementedAppUserMiddlewareUserServer() {}

// UnsafeAppUserMiddlewareUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserMiddlewareUserServer will
// result in compilation errors.
type UnsafeAppUserMiddlewareUserServer interface {
	mustEmbedUnimplementedAppUserMiddlewareUserServer()
}

func RegisterAppUserMiddlewareUserServer(s grpc.ServiceRegistrar, srv AppUserMiddlewareUserServer) {
	s.RegisterService(&AppUserMiddlewareUser_ServiceDesc, srv)
}

func _AppUserMiddlewareUser_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareUserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appusermw.user.v2.AppUserMiddlewareUser/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareUserServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareUser_GetAppUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareUserServer).GetAppUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appusermw.user.v2.AppUserMiddlewareUser/GetAppUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareUserServer).GetAppUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareUser_GetUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareUserServer).GetUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appusermw.user.v2.AppUserMiddlewareUser/GetUserInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareUserServer).GetUserInfos(ctx, req.(*GetUserInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareUser_CreateUserWithSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserWithSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareUserServer).CreateUserWithSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appusermw.user.v2.AppUserMiddlewareUser/CreateUserWithSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareUserServer).CreateUserWithSecret(ctx, req.(*CreateUserWithSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareUser_CreateUserWithThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserWithThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareUserServer).CreateUserWithThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appusermw.user.v2.AppUserMiddlewareUser/CreateUserWithThirdParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareUserServer).CreateUserWithThirdParty(ctx, req.(*CreateUserWithThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserMiddlewareUser_ServiceDesc is the grpc.ServiceDesc for AppUserMiddlewareUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserMiddlewareUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appusermw.user.v2.AppUserMiddlewareUser",
	HandlerType: (*AppUserMiddlewareUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _AppUserMiddlewareUser_GetUserInfo_Handler,
		},
		{
			MethodName: "GetAppUserInfo",
			Handler:    _AppUserMiddlewareUser_GetAppUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfos",
			Handler:    _AppUserMiddlewareUser_GetUserInfos_Handler,
		},
		{
			MethodName: "CreateUserWithSecret",
			Handler:    _AppUserMiddlewareUser_CreateUserWithSecret_Handler,
		},
		{
			MethodName: "CreateUserWithThirdParty",
			Handler:    _AppUserMiddlewareUser_CreateUserWithThirdParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/user/user.proto",
}
