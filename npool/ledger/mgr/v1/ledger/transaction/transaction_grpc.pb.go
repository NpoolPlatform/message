// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ledger/mgr/v1/ledger/transaction/transaction.proto

package transaction

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
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	CreateTransactions(ctx context.Context, in *CreateTransactionsRequest, opts ...grpc.CallOption) (*CreateTransactionsResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	GetTransactionOnly(ctx context.Context, in *GetTransactionOnlyRequest, opts ...grpc.CallOption) (*GetTransactionOnlyResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	ExistTransaction(ctx context.Context, in *ExistTransactionRequest, opts ...grpc.CallOption) (*ExistTransactionResponse, error)
	ExistTransactionConds(ctx context.Context, in *ExistTransactionCondsRequest, opts ...grpc.CallOption) (*ExistTransactionCondsResponse, error)
	CountTransactions(ctx context.Context, in *CountTransactionsRequest, opts ...grpc.CallOption) (*CountTransactionsResponse, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateTransactions(ctx context.Context, in *CreateTransactionsRequest, opts ...grpc.CallOption) (*CreateTransactionsResponse, error) {
	out := new(CreateTransactionsResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/CreateTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error) {
	out := new(UpdateTransactionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTransactionOnly(ctx context.Context, in *GetTransactionOnlyRequest, opts ...grpc.CallOption) (*GetTransactionOnlyResponse, error) {
	out := new(GetTransactionOnlyResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/GetTransactionOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistTransaction(ctx context.Context, in *ExistTransactionRequest, opts ...grpc.CallOption) (*ExistTransactionResponse, error) {
	out := new(ExistTransactionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/ExistTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistTransactionConds(ctx context.Context, in *ExistTransactionCondsRequest, opts ...grpc.CallOption) (*ExistTransactionCondsResponse, error) {
	out := new(ExistTransactionCondsResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/ExistTransactionConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountTransactions(ctx context.Context, in *CountTransactionsRequest, opts ...grpc.CallOption) (*CountTransactionsResponse, error) {
	out := new(CountTransactionsResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/CountTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*DeleteTransactionResponse, error) {
	out := new(DeleteTransactionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.ledger.transaction.v1.Manager/DeleteTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	CreateTransactions(context.Context, *CreateTransactionsRequest) (*CreateTransactionsResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	GetTransactionOnly(context.Context, *GetTransactionOnlyRequest) (*GetTransactionOnlyResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	ExistTransaction(context.Context, *ExistTransactionRequest) (*ExistTransactionResponse, error)
	ExistTransactionConds(context.Context, *ExistTransactionCondsRequest) (*ExistTransactionCondsResponse, error)
	CountTransactions(context.Context, *CountTransactionsRequest) (*CountTransactionsResponse, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedManagerServer) CreateTransactions(context.Context, *CreateTransactionsRequest) (*CreateTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransactions not implemented")
}
func (UnimplementedManagerServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedManagerServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedManagerServer) GetTransactionOnly(context.Context, *GetTransactionOnlyRequest) (*GetTransactionOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionOnly not implemented")
}
func (UnimplementedManagerServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedManagerServer) ExistTransaction(context.Context, *ExistTransactionRequest) (*ExistTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTransaction not implemented")
}
func (UnimplementedManagerServer) ExistTransactionConds(context.Context, *ExistTransactionCondsRequest) (*ExistTransactionCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTransactionConds not implemented")
}
func (UnimplementedManagerServer) CountTransactions(context.Context, *CountTransactionsRequest) (*CountTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTransactions not implemented")
}
func (UnimplementedManagerServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*DeleteTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
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

func _Manager_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/CreateTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateTransactions(ctx, req.(*CreateTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTransactionOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransactionOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/GetTransactionOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransactionOnly(ctx, req.(*GetTransactionOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/ExistTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistTransaction(ctx, req.(*ExistTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistTransactionConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTransactionCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistTransactionConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/ExistTransactionConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistTransactionConds(ctx, req.(*ExistTransactionCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/CountTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountTransactions(ctx, req.(*CountTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.ledger.transaction.v1.Manager/DeleteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.manager.ledger.transaction.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _Manager_CreateTransaction_Handler,
		},
		{
			MethodName: "CreateTransactions",
			Handler:    _Manager_CreateTransactions_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _Manager_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Manager_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionOnly",
			Handler:    _Manager_GetTransactionOnly_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _Manager_GetTransactions_Handler,
		},
		{
			MethodName: "ExistTransaction",
			Handler:    _Manager_ExistTransaction_Handler,
		},
		{
			MethodName: "ExistTransactionConds",
			Handler:    _Manager_ExistTransactionConds_Handler,
		},
		{
			MethodName: "CountTransactions",
			Handler:    _Manager_CountTransactions_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _Manager_DeleteTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/mgr/v1/ledger/transaction/transaction.proto",
}
