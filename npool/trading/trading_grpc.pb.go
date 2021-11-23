// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trading

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

// TradingClient is the client API for Trading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingClient interface {
	// 创建账户
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	// 余额查询
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// 转账 / 提现
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	// TODO: 账户交易查询
	GetTxJSON(ctx context.Context, in *GetTxJSONRequest, opts ...grpc.CallOption) (*AccountTxJSON, error)
	// 交易状态查询
	GetInsiteTxStatus(ctx context.Context, in *GetInsiteTxStatusRequest, opts ...grpc.CallOption) (*GetInsiteTxStatusResponse, error)
	// 接收ack
	ACK(ctx context.Context, in *ACKRequest, opts ...grpc.CallOption) (*ACKResponse, error)
}

type tradingClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingClient(cc grpc.ClientConnInterface) TradingClient {
	return &tradingClient{cc}
}

func (c *tradingClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) GetTxJSON(ctx context.Context, in *GetTxJSONRequest, opts ...grpc.CallOption) (*AccountTxJSON, error) {
	out := new(AccountTxJSON)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/GetTxJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) GetInsiteTxStatus(ctx context.Context, in *GetInsiteTxStatusRequest, opts ...grpc.CallOption) (*GetInsiteTxStatusResponse, error) {
	out := new(GetInsiteTxStatusResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/GetInsiteTxStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) ACK(ctx context.Context, in *ACKRequest, opts ...grpc.CallOption) (*ACKResponse, error) {
	out := new(ACKResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/ACK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServer is the server API for Trading service.
// All implementations must embed UnimplementedTradingServer
// for forward compatibility
type TradingServer interface {
	// 创建账户
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	// 余额查询
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// 转账 / 提现
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	// TODO: 账户交易查询
	GetTxJSON(context.Context, *GetTxJSONRequest) (*AccountTxJSON, error)
	// 交易状态查询
	GetInsiteTxStatus(context.Context, *GetInsiteTxStatusRequest) (*GetInsiteTxStatusResponse, error)
	// 接收ack
	ACK(context.Context, *ACKRequest) (*ACKResponse, error)
	mustEmbedUnimplementedTradingServer()
}

// UnimplementedTradingServer must be embedded to have forward compatible implementations.
type UnimplementedTradingServer struct {
}

func (UnimplementedTradingServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedTradingServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedTradingServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTradingServer) GetTxJSON(context.Context, *GetTxJSONRequest) (*AccountTxJSON, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxJSON not implemented")
}
func (UnimplementedTradingServer) GetInsiteTxStatus(context.Context, *GetInsiteTxStatusRequest) (*GetInsiteTxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInsiteTxStatus not implemented")
}
func (UnimplementedTradingServer) ACK(context.Context, *ACKRequest) (*ACKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ACK not implemented")
}
func (UnimplementedTradingServer) mustEmbedUnimplementedTradingServer() {}

// UnsafeTradingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradingServer will
// result in compilation errors.
type UnsafeTradingServer interface {
	mustEmbedUnimplementedTradingServer()
}

func RegisterTradingServer(s grpc.ServiceRegistrar, srv TradingServer) {
	s.RegisterService(&Trading_ServiceDesc, srv)
}

func _Trading_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_GetTxJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxJSONRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).GetTxJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/GetTxJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).GetTxJSON(ctx, req.(*GetTxJSONRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_GetInsiteTxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInsiteTxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).GetInsiteTxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/GetInsiteTxStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).GetInsiteTxStatus(ctx, req.(*GetInsiteTxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_ACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).ACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/ACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).ACK(ctx, req.(*ACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Trading_ServiceDesc is the grpc.ServiceDesc for Trading service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Trading_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.v1.Trading",
	HandlerType: (*TradingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Trading_CreateAccount_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Trading_GetBalance_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Trading_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTxJSON",
			Handler:    _Trading_GetTxJSON_Handler,
		},
		{
			MethodName: "GetInsiteTxStatus",
			Handler:    _Trading_GetInsiteTxStatus_Handler,
		},
		{
			MethodName: "ACK",
			Handler:    _Trading_ACK_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/trading/trading.proto",
}
