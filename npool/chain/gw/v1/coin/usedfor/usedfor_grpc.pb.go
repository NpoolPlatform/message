// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/gw/v1/coin/usedfor/usedfor.proto

package usedfor

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
	Gateway_CreateCoinUsedFor_FullMethodName = "/chain.gateway.coin.usedfor.v1.Gateway/CreateCoinUsedFor"
	Gateway_GetCoinUsedFors_FullMethodName   = "/chain.gateway.coin.usedfor.v1.Gateway/GetCoinUsedFors"
	Gateway_DeleteCoinUsedFor_FullMethodName = "/chain.gateway.coin.usedfor.v1.Gateway/DeleteCoinUsedFor"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateCoinUsedFor(ctx context.Context, in *CreateCoinUsedForRequest, opts ...grpc.CallOption) (*CreateCoinUsedForResponse, error)
	GetCoinUsedFors(ctx context.Context, in *GetCoinUsedForsRequest, opts ...grpc.CallOption) (*GetCoinUsedForsResponse, error)
	DeleteCoinUsedFor(ctx context.Context, in *DeleteCoinUsedForRequest, opts ...grpc.CallOption) (*DeleteCoinUsedForResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateCoinUsedFor(ctx context.Context, in *CreateCoinUsedForRequest, opts ...grpc.CallOption) (*CreateCoinUsedForResponse, error) {
	out := new(CreateCoinUsedForResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateCoinUsedFor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetCoinUsedFors(ctx context.Context, in *GetCoinUsedForsRequest, opts ...grpc.CallOption) (*GetCoinUsedForsResponse, error) {
	out := new(GetCoinUsedForsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCoinUsedFors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteCoinUsedFor(ctx context.Context, in *DeleteCoinUsedForRequest, opts ...grpc.CallOption) (*DeleteCoinUsedForResponse, error) {
	out := new(DeleteCoinUsedForResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteCoinUsedFor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateCoinUsedFor(context.Context, *CreateCoinUsedForRequest) (*CreateCoinUsedForResponse, error)
	GetCoinUsedFors(context.Context, *GetCoinUsedForsRequest) (*GetCoinUsedForsResponse, error)
	DeleteCoinUsedFor(context.Context, *DeleteCoinUsedForRequest) (*DeleteCoinUsedForResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateCoinUsedFor(context.Context, *CreateCoinUsedForRequest) (*CreateCoinUsedForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinUsedFor not implemented")
}
func (UnimplementedGatewayServer) GetCoinUsedFors(context.Context, *GetCoinUsedForsRequest) (*GetCoinUsedForsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinUsedFors not implemented")
}
func (UnimplementedGatewayServer) DeleteCoinUsedFor(context.Context, *DeleteCoinUsedForRequest) (*DeleteCoinUsedForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoinUsedFor not implemented")
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

func _Gateway_CreateCoinUsedFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinUsedForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateCoinUsedFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateCoinUsedFor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateCoinUsedFor(ctx, req.(*CreateCoinUsedForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetCoinUsedFors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinUsedForsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCoinUsedFors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetCoinUsedFors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCoinUsedFors(ctx, req.(*GetCoinUsedForsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteCoinUsedFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinUsedForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteCoinUsedFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteCoinUsedFor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteCoinUsedFor(ctx, req.(*DeleteCoinUsedForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.coin.usedfor.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinUsedFor",
			Handler:    _Gateway_CreateCoinUsedFor_Handler,
		},
		{
			MethodName: "GetCoinUsedFors",
			Handler:    _Gateway_GetCoinUsedFors_Handler,
		},
		{
			MethodName: "DeleteCoinUsedFor",
			Handler:    _Gateway_DeleteCoinUsedFor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/gw/v1/coin/usedfor/usedfor.proto",
}
