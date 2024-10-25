// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/coin/allocated/allocated.proto

package allocated

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
	Gateway_GetMyCoinAllocateds_FullMethodName    = "/inspire.gateway.coin.allocated.v1.Gateway/GetMyCoinAllocateds"
	Gateway_AdminGetCoinAllocateds_FullMethodName = "/inspire.gateway.coin.allocated.v1.Gateway/AdminGetCoinAllocateds"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetMyCoinAllocateds(ctx context.Context, in *GetMyCoinAllocatedsRequest, opts ...grpc.CallOption) (*GetMyCoinAllocatedsResponse, error)
	AdminGetCoinAllocateds(ctx context.Context, in *AdminGetCoinAllocatedsRequest, opts ...grpc.CallOption) (*AdminGetCoinAllocatedsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetMyCoinAllocateds(ctx context.Context, in *GetMyCoinAllocatedsRequest, opts ...grpc.CallOption) (*GetMyCoinAllocatedsResponse, error) {
	out := new(GetMyCoinAllocatedsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyCoinAllocateds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetCoinAllocateds(ctx context.Context, in *AdminGetCoinAllocatedsRequest, opts ...grpc.CallOption) (*AdminGetCoinAllocatedsResponse, error) {
	out := new(AdminGetCoinAllocatedsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetCoinAllocateds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetMyCoinAllocateds(context.Context, *GetMyCoinAllocatedsRequest) (*GetMyCoinAllocatedsResponse, error)
	AdminGetCoinAllocateds(context.Context, *AdminGetCoinAllocatedsRequest) (*AdminGetCoinAllocatedsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetMyCoinAllocateds(context.Context, *GetMyCoinAllocatedsRequest) (*GetMyCoinAllocatedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCoinAllocateds not implemented")
}
func (UnimplementedGatewayServer) AdminGetCoinAllocateds(context.Context, *AdminGetCoinAllocatedsRequest) (*AdminGetCoinAllocatedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetCoinAllocateds not implemented")
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

func _Gateway_GetMyCoinAllocateds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCoinAllocatedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyCoinAllocateds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyCoinAllocateds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyCoinAllocateds(ctx, req.(*GetMyCoinAllocatedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetCoinAllocateds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetCoinAllocatedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetCoinAllocateds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetCoinAllocateds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetCoinAllocateds(ctx, req.(*AdminGetCoinAllocatedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coin.allocated.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyCoinAllocateds",
			Handler:    _Gateway_GetMyCoinAllocateds_Handler,
		},
		{
			MethodName: "AdminGetCoinAllocateds",
			Handler:    _Gateway_AdminGetCoinAllocateds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coin/allocated/allocated.proto",
}
