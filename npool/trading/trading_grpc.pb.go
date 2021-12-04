// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trading

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TradingClient is the client API for Trading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	GetWalletBalance(ctx context.Context, in *GetWalletBalanceRequest, opts ...grpc.CallOption) (*GetWalletBalanceResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type tradingClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingClient(cc grpc.ClientConnInterface) TradingClient {
	return &tradingClient{cc}
}

func (c *tradingClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) GetWalletBalance(ctx context.Context, in *GetWalletBalanceRequest, opts ...grpc.CallOption) (*GetWalletBalanceResponse, error) {
	out := new(GetWalletBalanceResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/GetWalletBalance", in, out, opts...)
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

func (c *tradingClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/sphinx.v1.Trading/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServer is the server API for Trading service.
// All implementations must embed UnimplementedTradingServer
// for forward compatibility
type TradingServer interface {
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWalletBalance(context.Context, *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedTradingServer()
}

// UnimplementedTradingServer must be embedded to have forward compatible implementations.
type UnimplementedTradingServer struct {
}

func (UnimplementedTradingServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedTradingServer) GetWalletBalance(context.Context, *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletBalance not implemented")
}
func (UnimplementedTradingServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTradingServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedTradingServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
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

func _Trading_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_GetWalletBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).GetWalletBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/GetWalletBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).GetWalletBalance(ctx, req.(*GetWalletBalanceRequest))
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

func _Trading_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trading_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.v1.Trading/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServer).Version(ctx, req.(*emptypb.Empty))
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
			MethodName: "CreateWallet",
			Handler:    _Trading_CreateWallet_Handler,
		},
		{
			MethodName: "GetWalletBalance",
			Handler:    _Trading_GetWalletBalance_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Trading_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Trading_GetTransaction_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Trading_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/trading/trading.proto",
}
