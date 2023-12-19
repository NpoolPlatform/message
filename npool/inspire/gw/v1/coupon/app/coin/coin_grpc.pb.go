// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/coupon/app/coin/coin.proto

package CouponCoin

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
	Gateway_GetCouponCoins_FullMethodName   = "/inspire.gateway.coupon.app.CouponCoin.v1.Gateway/GetCouponCoins"
	Gateway_CreateCouponCoin_FullMethodName = "/inspire.gateway.coupon.app.CouponCoin.v1.Gateway/CreateCouponCoin"
	Gateway_DeleteCouponCoin_FullMethodName = "/inspire.gateway.coupon.app.CouponCoin.v1.Gateway/DeleteCouponCoin"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetCouponCoins(ctx context.Context, in *GetCouponCoinsRequest, opts ...grpc.CallOption) (*GetCouponCoinsResponse, error)
	CreateCouponCoin(ctx context.Context, in *CreateCouponCoinRequest, opts ...grpc.CallOption) (*CreateCouponCoinResponse, error)
	DeleteCouponCoin(ctx context.Context, in *DeleteCouponCoinRequest, opts ...grpc.CallOption) (*DeleteCouponCoinResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetCouponCoins(ctx context.Context, in *GetCouponCoinsRequest, opts ...grpc.CallOption) (*GetCouponCoinsResponse, error) {
	out := new(GetCouponCoinsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCouponCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateCouponCoin(ctx context.Context, in *CreateCouponCoinRequest, opts ...grpc.CallOption) (*CreateCouponCoinResponse, error) {
	out := new(CreateCouponCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateCouponCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteCouponCoin(ctx context.Context, in *DeleteCouponCoinRequest, opts ...grpc.CallOption) (*DeleteCouponCoinResponse, error) {
	out := new(DeleteCouponCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteCouponCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetCouponCoins(context.Context, *GetCouponCoinsRequest) (*GetCouponCoinsResponse, error)
	CreateCouponCoin(context.Context, *CreateCouponCoinRequest) (*CreateCouponCoinResponse, error)
	DeleteCouponCoin(context.Context, *DeleteCouponCoinRequest) (*DeleteCouponCoinResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetCouponCoins(context.Context, *GetCouponCoinsRequest) (*GetCouponCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponCoins not implemented")
}
func (UnimplementedGatewayServer) CreateCouponCoin(context.Context, *CreateCouponCoinRequest) (*CreateCouponCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCouponCoin not implemented")
}
func (UnimplementedGatewayServer) DeleteCouponCoin(context.Context, *DeleteCouponCoinRequest) (*DeleteCouponCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCouponCoin not implemented")
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

func _Gateway_GetCouponCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCouponCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetCouponCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCouponCoins(ctx, req.(*GetCouponCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateCouponCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCouponCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateCouponCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateCouponCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateCouponCoin(ctx, req.(*CreateCouponCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteCouponCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCouponCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteCouponCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteCouponCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteCouponCoin(ctx, req.(*DeleteCouponCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coupon.app.CouponCoin.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCouponCoins",
			Handler:    _Gateway_GetCouponCoins_Handler,
		},
		{
			MethodName: "CreateCouponCoin",
			Handler:    _Gateway_CreateCouponCoin_Handler,
		},
		{
			MethodName: "DeleteCouponCoin",
			Handler:    _Gateway_DeleteCouponCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coupon/app/coin/coin.proto",
}
