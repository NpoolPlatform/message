// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/account/gw/v1/user/user.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetDepositAccount(ctx context.Context, in *GetDepositAccountRequest, opts ...grpc.CallOption) (*GetDepositAccountResponse, error)
	GetAppDepositAccounts(ctx context.Context, in *GetAppDepositAccountsRequest, opts ...grpc.CallOption) (*GetAppDepositAccountsResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	UpdateAppUserAccount(ctx context.Context, in *UpdateAppUserAccountRequest, opts ...grpc.CallOption) (*UpdateAppUserAccountResponse, error)
	GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
	GetAppAccounts(ctx context.Context, in *GetAppAccountsRequest, opts ...grpc.CallOption) (*GetAppAccountsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetDepositAccount(ctx context.Context, in *GetDepositAccountRequest, opts ...grpc.CallOption) (*GetDepositAccountResponse, error) {
	out := new(GetDepositAccountResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/GetDepositAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppDepositAccounts(ctx context.Context, in *GetAppDepositAccountsRequest, opts ...grpc.CallOption) (*GetAppDepositAccountsResponse, error) {
	out := new(GetAppDepositAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/GetAppDepositAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppUserAccount(ctx context.Context, in *UpdateAppUserAccountRequest, opts ...grpc.CallOption) (*UpdateAppUserAccountResponse, error) {
	out := new(UpdateAppUserAccountResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/UpdateAppUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppAccounts(ctx context.Context, in *GetAppAccountsRequest, opts ...grpc.CallOption) (*GetAppAccountsResponse, error) {
	out := new(GetAppAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.gateway.user.v1.Gateway/GetAppAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetDepositAccount(context.Context, *GetDepositAccountRequest) (*GetDepositAccountResponse, error)
	GetAppDepositAccounts(context.Context, *GetAppDepositAccountsRequest) (*GetAppDepositAccountsResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	UpdateAppUserAccount(context.Context, *UpdateAppUserAccountRequest) (*UpdateAppUserAccountResponse, error)
	GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
	GetAppAccounts(context.Context, *GetAppAccountsRequest) (*GetAppAccountsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetDepositAccount(context.Context, *GetDepositAccountRequest) (*GetDepositAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositAccount not implemented")
}
func (UnimplementedGatewayServer) GetAppDepositAccounts(context.Context, *GetAppDepositAccountsRequest) (*GetAppDepositAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDepositAccounts not implemented")
}
func (UnimplementedGatewayServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedGatewayServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedGatewayServer) UpdateAppUserAccount(context.Context, *UpdateAppUserAccountRequest) (*UpdateAppUserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppUserAccount not implemented")
}
func (UnimplementedGatewayServer) GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedGatewayServer) GetAppAccounts(context.Context, *GetAppAccountsRequest) (*GetAppAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppAccounts not implemented")
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

func _Gateway_GetDepositAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepositAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetDepositAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/GetDepositAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetDepositAccount(ctx, req.(*GetDepositAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppDepositAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDepositAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppDepositAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/GetAppDepositAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppDepositAccounts(ctx, req.(*GetAppDepositAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/UpdateAppUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppUserAccount(ctx, req.(*UpdateAppUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.gateway.user.v1.Gateway/GetAppAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppAccounts(ctx, req.(*GetAppAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.gateway.user.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDepositAccount",
			Handler:    _Gateway_GetDepositAccount_Handler,
		},
		{
			MethodName: "GetAppDepositAccounts",
			Handler:    _Gateway_GetAppDepositAccounts_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Gateway_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Gateway_UpdateAccount_Handler,
		},
		{
			MethodName: "UpdateAppUserAccount",
			Handler:    _Gateway_UpdateAppUserAccount_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _Gateway_GetAccounts_Handler,
		},
		{
			MethodName: "GetAppAccounts",
			Handler:    _Gateway_GetAppAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/account/gw/v1/user/user.proto",
}
