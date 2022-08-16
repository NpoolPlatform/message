// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/account/mgr/v1/deposit/deposit.proto

package deposit

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateAccounts(ctx context.Context, in *CreateAccountsRequest, opts ...grpc.CallOption) (*CreateAccountsResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetAccountOnly(ctx context.Context, in *GetAccountOnlyRequest, opts ...grpc.CallOption) (*GetAccountOnlyResponse, error)
	GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error)
	ExistAccount(ctx context.Context, in *ExistAccountRequest, opts ...grpc.CallOption) (*ExistAccountResponse, error)
	ExistAccountConds(ctx context.Context, in *ExistAccountCondsRequest, opts ...grpc.CallOption) (*ExistAccountCondsResponse, error)
	CountAccounts(ctx context.Context, in *CountAccountsRequest, opts ...grpc.CallOption) (*CountAccountsResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateAccounts(ctx context.Context, in *CreateAccountsRequest, opts ...grpc.CallOption) (*CreateAccountsResponse, error) {
	out := new(CreateAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/CreateAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AddAccountResponse, error) {
	out := new(AddAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetAccountOnly(ctx context.Context, in *GetAccountOnlyRequest, opts ...grpc.CallOption) (*GetAccountOnlyResponse, error) {
	out := new(GetAccountOnlyResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/GetAccountOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*GetAccountsResponse, error) {
	out := new(GetAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistAccount(ctx context.Context, in *ExistAccountRequest, opts ...grpc.CallOption) (*ExistAccountResponse, error) {
	out := new(ExistAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/ExistAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistAccountConds(ctx context.Context, in *ExistAccountCondsRequest, opts ...grpc.CallOption) (*ExistAccountCondsResponse, error) {
	out := new(ExistAccountCondsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/ExistAccountConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountAccounts(ctx context.Context, in *CountAccountsRequest, opts ...grpc.CallOption) (*CountAccountsResponse, error) {
	out := new(CountAccountsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/CountAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/account.manager.deposit.v1.Manager/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	CreateAccounts(context.Context, *CreateAccountsRequest) (*CreateAccountsResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetAccountOnly(context.Context, *GetAccountOnlyRequest) (*GetAccountOnlyResponse, error)
	GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error)
	ExistAccount(context.Context, *ExistAccountRequest) (*ExistAccountResponse, error)
	ExistAccountConds(context.Context, *ExistAccountCondsRequest) (*ExistAccountCondsResponse, error)
	CountAccounts(context.Context, *CountAccountsRequest) (*CountAccountsResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedManagerServer) CreateAccounts(context.Context, *CreateAccountsRequest) (*CreateAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccounts not implemented")
}
func (UnimplementedManagerServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedManagerServer) AddAccount(context.Context, *AddAccountRequest) (*AddAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedManagerServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedManagerServer) GetAccountOnly(context.Context, *GetAccountOnlyRequest) (*GetAccountOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountOnly not implemented")
}
func (UnimplementedManagerServer) GetAccounts(context.Context, *GetAccountsRequest) (*GetAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedManagerServer) ExistAccount(context.Context, *ExistAccountRequest) (*ExistAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAccount not implemented")
}
func (UnimplementedManagerServer) ExistAccountConds(context.Context, *ExistAccountCondsRequest) (*ExistAccountCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAccountConds not implemented")
}
func (UnimplementedManagerServer) CountAccounts(context.Context, *CountAccountsRequest) (*CountAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAccounts not implemented")
}
func (UnimplementedManagerServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/CreateAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateAccounts(ctx, req.(*CreateAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).AddAccount(ctx, req.(*AddAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetAccountOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetAccountOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/GetAccountOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetAccountOnly(ctx, req.(*GetAccountOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/ExistAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistAccount(ctx, req.(*ExistAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistAccountConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAccountCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistAccountConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/ExistAccountConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistAccountConds(ctx, req.(*ExistAccountCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/CountAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountAccounts(ctx, req.(*CountAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.deposit.v1.Manager/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.manager.deposit.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Manager_CreateAccount_Handler,
		},
		{
			MethodName: "CreateAccounts",
			Handler:    _Manager_CreateAccounts_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Manager_UpdateAccount_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _Manager_AddAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Manager_GetAccount_Handler,
		},
		{
			MethodName: "GetAccountOnly",
			Handler:    _Manager_GetAccountOnly_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _Manager_GetAccounts_Handler,
		},
		{
			MethodName: "ExistAccount",
			Handler:    _Manager_ExistAccount_Handler,
		},
		{
			MethodName: "ExistAccountConds",
			Handler:    _Manager_ExistAccountConds_Handler,
		},
		{
			MethodName: "CountAccounts",
			Handler:    _Manager_CountAccounts_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Manager_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/account/mgr/v1/deposit/deposit.proto",
}
