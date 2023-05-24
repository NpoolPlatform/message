// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/gw/v1/coin/currency/currency.proto

package currency

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
<<<<<<< HEAD
	Gateway_GetCurrency_FullMethodName   = "/chain.gateway.coin.currency.v1.Gateway/GetCurrency"
	Gateway_GetCurrencies_FullMethodName = "/chain.gateway.coin.currency.v1.Gateway/GetCurrencies"
=======
	Gateway_GetCurrencies_FullMethodName = "/chain.gateway.coin.currency.v1.Gateway/GetCurrencies"
	Gateway_GetHistories_FullMethodName  = "/chain.gateway.coin.currency.v1.Gateway/GetHistories"
>>>>>>> Regenerate
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
	GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

<<<<<<< HEAD
func (c *gatewayClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCurrency_FullMethodName, in, out, opts...)
=======
func (c *gatewayClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCurrencies_FullMethodName, in, out, opts...)
>>>>>>> Regenerate
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *gatewayClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCurrencies_FullMethodName, in, out, opts...)
=======
func (c *gatewayClient) GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error) {
	out := new(GetHistoriesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetHistories_FullMethodName, in, out, opts...)
>>>>>>> Regenerate
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (UnimplementedGatewayServer) GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
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

func _Gateway_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Gateway_GetCurrency_FullMethodName,
=======
		FullMethod: Gateway_GetCurrencies_FullMethodName,
>>>>>>> Regenerate
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Gateway_GetCurrencies_FullMethodName,
=======
		FullMethod: Gateway_GetHistories_FullMethodName,
>>>>>>> Regenerate
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCurrencies(ctx, req.(*GetCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.coin.currency.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrency",
			Handler:    _Gateway_GetCurrency_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _Gateway_GetCurrencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/gw/v1/coin/currency/currency.proto",
}
