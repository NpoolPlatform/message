// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/good/coin/coin.proto

package coin

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
	Gateway_AdminCreateGoodCoin_FullMethodName = "/good.gateway.good1.coin.v1.Gateway/AdminCreateGoodCoin"
	Gateway_AdminUpdateGoodCoin_FullMethodName = "/good.gateway.good1.coin.v1.Gateway/AdminUpdateGoodCoin"
	Gateway_GetGoodCoins_FullMethodName        = "/good.gateway.good1.coin.v1.Gateway/GetGoodCoins"
	Gateway_AdminDeleteGoodCoin_FullMethodName = "/good.gateway.good1.coin.v1.Gateway/AdminDeleteGoodCoin"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateGoodCoin(ctx context.Context, in *AdminCreateGoodCoinRequest, opts ...grpc.CallOption) (*AdminCreateGoodCoinResponse, error)
	AdminUpdateGoodCoin(ctx context.Context, in *AdminUpdateGoodCoinRequest, opts ...grpc.CallOption) (*AdminUpdateGoodCoinResponse, error)
	GetGoodCoins(ctx context.Context, in *GetGoodCoinsRequest, opts ...grpc.CallOption) (*GetGoodCoinsResponse, error)
	AdminDeleteGoodCoin(ctx context.Context, in *AdminDeleteGoodCoinRequest, opts ...grpc.CallOption) (*AdminDeleteGoodCoinResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateGoodCoin(ctx context.Context, in *AdminCreateGoodCoinRequest, opts ...grpc.CallOption) (*AdminCreateGoodCoinResponse, error) {
	out := new(AdminCreateGoodCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateGoodCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateGoodCoin(ctx context.Context, in *AdminUpdateGoodCoinRequest, opts ...grpc.CallOption) (*AdminUpdateGoodCoinResponse, error) {
	out := new(AdminUpdateGoodCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateGoodCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGoodCoins(ctx context.Context, in *GetGoodCoinsRequest, opts ...grpc.CallOption) (*GetGoodCoinsResponse, error) {
	out := new(GetGoodCoinsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetGoodCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteGoodCoin(ctx context.Context, in *AdminDeleteGoodCoinRequest, opts ...grpc.CallOption) (*AdminDeleteGoodCoinResponse, error) {
	out := new(AdminDeleteGoodCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteGoodCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateGoodCoin(context.Context, *AdminCreateGoodCoinRequest) (*AdminCreateGoodCoinResponse, error)
	AdminUpdateGoodCoin(context.Context, *AdminUpdateGoodCoinRequest) (*AdminUpdateGoodCoinResponse, error)
	GetGoodCoins(context.Context, *GetGoodCoinsRequest) (*GetGoodCoinsResponse, error)
	AdminDeleteGoodCoin(context.Context, *AdminDeleteGoodCoinRequest) (*AdminDeleteGoodCoinResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreateGoodCoin(context.Context, *AdminCreateGoodCoinRequest) (*AdminCreateGoodCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateGoodCoin not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateGoodCoin(context.Context, *AdminUpdateGoodCoinRequest) (*AdminUpdateGoodCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateGoodCoin not implemented")
}
func (UnimplementedGatewayServer) GetGoodCoins(context.Context, *GetGoodCoinsRequest) (*GetGoodCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodCoins not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteGoodCoin(context.Context, *AdminDeleteGoodCoinRequest) (*AdminDeleteGoodCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteGoodCoin not implemented")
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

func _Gateway_AdminCreateGoodCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateGoodCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateGoodCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateGoodCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateGoodCoin(ctx, req.(*AdminCreateGoodCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateGoodCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateGoodCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateGoodCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateGoodCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateGoodCoin(ctx, req.(*AdminUpdateGoodCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGoodCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGoodCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetGoodCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGoodCoins(ctx, req.(*GetGoodCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteGoodCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteGoodCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteGoodCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteGoodCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteGoodCoin(ctx, req.(*AdminDeleteGoodCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.good1.coin.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateGoodCoin",
			Handler:    _Gateway_AdminCreateGoodCoin_Handler,
		},
		{
			MethodName: "AdminUpdateGoodCoin",
			Handler:    _Gateway_AdminUpdateGoodCoin_Handler,
		},
		{
			MethodName: "GetGoodCoins",
			Handler:    _Gateway_GetGoodCoins_Handler,
		},
		{
			MethodName: "AdminDeleteGoodCoin",
			Handler:    _Gateway_AdminDeleteGoodCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/good/coin/coin.proto",
}