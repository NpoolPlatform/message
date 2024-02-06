// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/gw/v1/coin/currency/feed/feed.proto

package feed

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
	Gateway_CreateFeed_FullMethodName = "/chain.gateway.coin.currency.feed.v1.Gateway/CreateFeed"
	Gateway_UpdateFeed_FullMethodName = "/chain.gateway.coin.currency.feed.v1.Gateway/UpdateFeed"
	Gateway_GetFeeds_FullMethodName   = "/chain.gateway.coin.currency.feed.v1.Gateway/GetFeeds"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedResponse, error)
	UpdateFeed(ctx context.Context, in *UpdateFeedRequest, opts ...grpc.CallOption) (*UpdateFeedResponse, error)
	GetFeeds(ctx context.Context, in *GetFeedsRequest, opts ...grpc.CallOption) (*GetFeedsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateFeed(ctx context.Context, in *CreateFeedRequest, opts ...grpc.CallOption) (*CreateFeedResponse, error) {
	out := new(CreateFeedResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateFeed(ctx context.Context, in *UpdateFeedRequest, opts ...grpc.CallOption) (*UpdateFeedResponse, error) {
	out := new(UpdateFeedResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetFeeds(ctx context.Context, in *GetFeedsRequest, opts ...grpc.CallOption) (*GetFeedsResponse, error) {
	out := new(GetFeedsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetFeeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error)
	UpdateFeed(context.Context, *UpdateFeedRequest) (*UpdateFeedResponse, error)
	GetFeeds(context.Context, *GetFeedsRequest) (*GetFeedsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateFeed(context.Context, *CreateFeedRequest) (*CreateFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedGatewayServer) UpdateFeed(context.Context, *UpdateFeedRequest) (*UpdateFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeed not implemented")
}
func (UnimplementedGatewayServer) GetFeeds(context.Context, *GetFeedsRequest) (*GetFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
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

func _Gateway_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateFeed(ctx, req.(*CreateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateFeed(ctx, req.(*UpdateFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetFeeds(ctx, req.(*GetFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.coin.currency.feed.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeed",
			Handler:    _Gateway_CreateFeed_Handler,
		},
		{
			MethodName: "UpdateFeed",
			Handler:    _Gateway_UpdateFeed_Handler,
		},
		{
			MethodName: "GetFeeds",
			Handler:    _Gateway_GetFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/gw/v1/coin/currency/feed/feed.proto",
}
