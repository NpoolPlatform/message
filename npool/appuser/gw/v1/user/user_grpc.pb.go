// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/user/user.proto

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
	Gateway_Signup_FullMethodName            = "/appuser.gateway.user.v1.Gateway/Signup"
	Gateway_CreateUser_FullMethodName        = "/appuser.gateway.user.v1.Gateway/CreateUser"
	Gateway_DeleteUser_FullMethodName        = "/appuser.gateway.user.v1.Gateway/DeleteUser"
	Gateway_CreateAppUser_FullMethodName     = "/appuser.gateway.user.v1.Gateway/CreateAppUser"
	Gateway_UpdateUser_FullMethodName        = "/appuser.gateway.user.v1.Gateway/UpdateUser"
	Gateway_UpdateUserKol_FullMethodName     = "/appuser.gateway.user.v1.Gateway/UpdateUserKol"
	Gateway_UpdateAppUser_FullMethodName     = "/appuser.gateway.user.v1.Gateway/UpdateAppUser"
	Gateway_ResetUser_FullMethodName         = "/appuser.gateway.user.v1.Gateway/ResetUser"
	Gateway_GetUsers_FullMethodName          = "/appuser.gateway.user.v1.Gateway/GetUsers"
	Gateway_GetAppUsers_FullMethodName       = "/appuser.gateway.user.v1.Gateway/GetAppUsers"
	Gateway_Login_FullMethodName             = "/appuser.gateway.user.v1.Gateway/Login"
	Gateway_LoginVerify_FullMethodName       = "/appuser.gateway.user.v1.Gateway/LoginVerify"
	Gateway_Logined_FullMethodName           = "/appuser.gateway.user.v1.Gateway/Logined"
	Gateway_Logout_FullMethodName            = "/appuser.gateway.user.v1.Gateway/Logout"
	Gateway_GetLoginHistories_FullMethodName = "/appuser.gateway.user.v1.Gateway/GetLoginHistories"
	Gateway_BanUser_FullMethodName           = "/appuser.gateway.user.v1.Gateway/BanUser"
	Gateway_BanAppUser_FullMethodName        = "/appuser.gateway.user.v1.Gateway/BanAppUser"
	Gateway_BindUser_FullMethodName          = "/appuser.gateway.user.v1.Gateway/BindUser"
	Gateway_UnbindOAuth_FullMethodName       = "/appuser.gateway.user.v1.Gateway/UnbindOAuth"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	UpdateUserKol(ctx context.Context, in *UpdateUserKolRequest, opts ...grpc.CallOption) (*UpdateUserKolResponse, error)
	UpdateAppUser(ctx context.Context, in *UpdateAppUserRequest, opts ...grpc.CallOption) (*UpdateAppUserResponse, error)
	ResetUser(ctx context.Context, in *ResetUserRequest, opts ...grpc.CallOption) (*ResetUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LoginVerify(ctx context.Context, in *LoginVerifyRequest, opts ...grpc.CallOption) (*LoginVerifyResponse, error)
	Logined(ctx context.Context, in *LoginedRequest, opts ...grpc.CallOption) (*LoginedResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetLoginHistories(ctx context.Context, in *GetLoginHistoriesRequest, opts ...grpc.CallOption) (*GetLoginHistoriesResponse, error)
	BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error)
	BanAppUser(ctx context.Context, in *BanAppUserRequest, opts ...grpc.CallOption) (*BanAppUserResponse, error)
	BindUser(ctx context.Context, in *BindUserRequest, opts ...grpc.CallOption) (*BindUserResponse, error)
	UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, opts ...grpc.CallOption) (*UnbindOAuthResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, Gateway_Signup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error) {
	out := new(CreateAppUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateUserKol(ctx context.Context, in *UpdateUserKolRequest, opts ...grpc.CallOption) (*UpdateUserKolResponse, error) {
	out := new(UpdateUserKolResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateUserKol_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppUser(ctx context.Context, in *UpdateAppUserRequest, opts ...grpc.CallOption) (*UpdateAppUserResponse, error) {
	out := new(UpdateAppUserResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) ResetUser(ctx context.Context, in *ResetUserRequest, opts ...grpc.CallOption) (*ResetUserResponse, error) {
	out := new(ResetUserResponse)
	err := c.cc.Invoke(ctx, Gateway_ResetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error) {
	out := new(GetAppUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Gateway_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) LoginVerify(ctx context.Context, in *LoginVerifyRequest, opts ...grpc.CallOption) (*LoginVerifyResponse, error) {
	out := new(LoginVerifyResponse)
	err := c.cc.Invoke(ctx, Gateway_LoginVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Logined(ctx context.Context, in *LoginedRequest, opts ...grpc.CallOption) (*LoginedResponse, error) {
	out := new(LoginedResponse)
	err := c.cc.Invoke(ctx, Gateway_Logined_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, Gateway_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetLoginHistories(ctx context.Context, in *GetLoginHistoriesRequest, opts ...grpc.CallOption) (*GetLoginHistoriesResponse, error) {
	out := new(GetLoginHistoriesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetLoginHistories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error) {
	out := new(BanUserResponse)
	err := c.cc.Invoke(ctx, Gateway_BanUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) BanAppUser(ctx context.Context, in *BanAppUserRequest, opts ...grpc.CallOption) (*BanAppUserResponse, error) {
	out := new(BanAppUserResponse)
	err := c.cc.Invoke(ctx, Gateway_BanAppUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) BindUser(ctx context.Context, in *BindUserRequest, opts ...grpc.CallOption) (*BindUserResponse, error) {
	out := new(BindUserResponse)
	err := c.cc.Invoke(ctx, Gateway_BindUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UnbindOAuth(ctx context.Context, in *UnbindOAuthRequest, opts ...grpc.CallOption) (*UnbindOAuthResponse, error) {
	out := new(UnbindOAuthResponse)
	err := c.cc.Invoke(ctx, Gateway_UnbindOAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	UpdateUserKol(context.Context, *UpdateUserKolRequest) (*UpdateUserKolResponse, error)
	UpdateAppUser(context.Context, *UpdateAppUserRequest) (*UpdateAppUserResponse, error)
	ResetUser(context.Context, *ResetUserRequest) (*ResetUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	LoginVerify(context.Context, *LoginVerifyRequest) (*LoginVerifyResponse, error)
	Logined(context.Context, *LoginedRequest) (*LoginedResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetLoginHistories(context.Context, *GetLoginHistoriesRequest) (*GetLoginHistoriesResponse, error)
	BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error)
	BanAppUser(context.Context, *BanAppUserRequest) (*BanAppUserResponse, error)
	BindUser(context.Context, *BindUserRequest) (*BindUserResponse, error)
	UnbindOAuth(context.Context, *UnbindOAuthRequest) (*UnbindOAuthResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedGatewayServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGatewayServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedGatewayServer) CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUser not implemented")
}
func (UnimplementedGatewayServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedGatewayServer) UpdateUserKol(context.Context, *UpdateUserKolRequest) (*UpdateUserKolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserKol not implemented")
}
func (UnimplementedGatewayServer) UpdateAppUser(context.Context, *UpdateAppUserRequest) (*UpdateAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppUser not implemented")
}
func (UnimplementedGatewayServer) ResetUser(context.Context, *ResetUserRequest) (*ResetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetUser not implemented")
}
func (UnimplementedGatewayServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedGatewayServer) GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUsers not implemented")
}
func (UnimplementedGatewayServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGatewayServer) LoginVerify(context.Context, *LoginVerifyRequest) (*LoginVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginVerify not implemented")
}
func (UnimplementedGatewayServer) Logined(context.Context, *LoginedRequest) (*LoginedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logined not implemented")
}
func (UnimplementedGatewayServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedGatewayServer) GetLoginHistories(context.Context, *GetLoginHistoriesRequest) (*GetLoginHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginHistories not implemented")
}
func (UnimplementedGatewayServer) BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanUser not implemented")
}
func (UnimplementedGatewayServer) BanAppUser(context.Context, *BanAppUserRequest) (*BanAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanAppUser not implemented")
}
func (UnimplementedGatewayServer) BindUser(context.Context, *BindUserRequest) (*BindUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUser not implemented")
}
func (UnimplementedGatewayServer) UnbindOAuth(context.Context, *UnbindOAuthRequest) (*UnbindOAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindOAuth not implemented")
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

func _Gateway_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppUser(ctx, req.(*CreateAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateUserKol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserKolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateUserKol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateUserKol_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateUserKol(ctx, req.(*UpdateUserKolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppUser(ctx, req.(*UpdateAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_ResetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).ResetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_ResetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).ResetUser(ctx, req.(*ResetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppUsers(ctx, req.(*GetAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_LoginVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).LoginVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_LoginVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).LoginVerify(ctx, req.(*LoginVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Logined_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Logined(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_Logined_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Logined(ctx, req.(*LoginedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetLoginHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginHistoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetLoginHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetLoginHistories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetLoginHistories(ctx, req.(*GetLoginHistoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_BanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).BanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_BanUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).BanUser(ctx, req.(*BanUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_BanAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).BanAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_BanAppUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).BanAppUser(ctx, req.(*BanAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_BindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).BindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_BindUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).BindUser(ctx, req.(*BindUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UnbindOAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindOAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UnbindOAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UnbindOAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UnbindOAuth(ctx, req.(*UnbindOAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.user.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Gateway_Signup_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Gateway_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Gateway_DeleteUser_Handler,
		},
		{
			MethodName: "CreateAppUser",
			Handler:    _Gateway_CreateAppUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Gateway_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserKol",
			Handler:    _Gateway_UpdateUserKol_Handler,
		},
		{
			MethodName: "UpdateAppUser",
			Handler:    _Gateway_UpdateAppUser_Handler,
		},
		{
			MethodName: "ResetUser",
			Handler:    _Gateway_ResetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Gateway_GetUsers_Handler,
		},
		{
			MethodName: "GetAppUsers",
			Handler:    _Gateway_GetAppUsers_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Gateway_Login_Handler,
		},
		{
			MethodName: "LoginVerify",
			Handler:    _Gateway_LoginVerify_Handler,
		},
		{
			MethodName: "Logined",
			Handler:    _Gateway_Logined_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Gateway_Logout_Handler,
		},
		{
			MethodName: "GetLoginHistories",
			Handler:    _Gateway_GetLoginHistories_Handler,
		},
		{
			MethodName: "BanUser",
			Handler:    _Gateway_BanUser_Handler,
		},
		{
			MethodName: "BanAppUser",
			Handler:    _Gateway_BanAppUser_Handler,
		},
		{
			MethodName: "BindUser",
			Handler:    _Gateway_BindUser_Handler,
		},
		{
			MethodName: "UnbindOAuth",
			Handler:    _Gateway_UnbindOAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/user/user.proto",
}
