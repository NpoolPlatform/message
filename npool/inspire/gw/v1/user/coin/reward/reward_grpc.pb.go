// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/user/coin/reward/reward.proto

package reward

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
	Gateway_AdminGetUserCoinRewards_FullMethodName = "/inspire.gateway.user.coin.reward.v1.Gateway/AdminGetUserCoinRewards"
	Gateway_UserGetUserCoinRewards_FullMethodName  = "/inspire.gateway.user.coin.reward.v1.Gateway/UserGetUserCoinRewards"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminGetUserCoinRewards(ctx context.Context, in *AdminGetUserCoinRewardsRequest, opts ...grpc.CallOption) (*AdminGetUserCoinRewardsResponse, error)
	UserGetUserCoinRewards(ctx context.Context, in *UserGetUserCoinRewardsRequest, opts ...grpc.CallOption) (*UserGetUserCoinRewardsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminGetUserCoinRewards(ctx context.Context, in *AdminGetUserCoinRewardsRequest, opts ...grpc.CallOption) (*AdminGetUserCoinRewardsResponse, error) {
	out := new(AdminGetUserCoinRewardsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetUserCoinRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UserGetUserCoinRewards(ctx context.Context, in *UserGetUserCoinRewardsRequest, opts ...grpc.CallOption) (*UserGetUserCoinRewardsResponse, error) {
	out := new(UserGetUserCoinRewardsResponse)
	err := c.cc.Invoke(ctx, Gateway_UserGetUserCoinRewards_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminGetUserCoinRewards(context.Context, *AdminGetUserCoinRewardsRequest) (*AdminGetUserCoinRewardsResponse, error)
	UserGetUserCoinRewards(context.Context, *UserGetUserCoinRewardsRequest) (*UserGetUserCoinRewardsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminGetUserCoinRewards(context.Context, *AdminGetUserCoinRewardsRequest) (*AdminGetUserCoinRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetUserCoinRewards not implemented")
}
func (UnimplementedGatewayServer) UserGetUserCoinRewards(context.Context, *UserGetUserCoinRewardsRequest) (*UserGetUserCoinRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetUserCoinRewards not implemented")
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

func _Gateway_AdminGetUserCoinRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetUserCoinRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetUserCoinRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetUserCoinRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetUserCoinRewards(ctx, req.(*AdminGetUserCoinRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UserGetUserCoinRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetUserCoinRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UserGetUserCoinRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UserGetUserCoinRewards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UserGetUserCoinRewards(ctx, req.(*UserGetUserCoinRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.user.coin.reward.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminGetUserCoinRewards",
			Handler:    _Gateway_AdminGetUserCoinRewards_Handler,
		},
		{
			MethodName: "UserGetUserCoinRewards",
			Handler:    _Gateway_UserGetUserCoinRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/user/coin/reward/reward.proto",
}
