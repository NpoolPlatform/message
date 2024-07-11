// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/gw/v1/chain/chain.proto

package chain

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
	Gateway_GetChains_FullMethodName        = "/chain.gateway.chain.v1.Gateway/GetChains"
	Gateway_AdminCreateChain_FullMethodName = "/chain.gateway.chain.v1.Gateway/AdminCreateChain"
	Gateway_AdminUpdateChain_FullMethodName = "/chain.gateway.chain.v1.Gateway/AdminUpdateChain"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsResponse, error)
	// admin
	AdminCreateChain(ctx context.Context, in *AdminCreateChainRequest, opts ...grpc.CallOption) (*AdminCreateChainResponse, error)
	AdminUpdateChain(ctx context.Context, in *AdminUpdateChainRequest, opts ...grpc.CallOption) (*AdminUpdateChainResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsResponse, error) {
	out := new(GetChainsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetChains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateChain(ctx context.Context, in *AdminCreateChainRequest, opts ...grpc.CallOption) (*AdminCreateChainResponse, error) {
	out := new(AdminCreateChainResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateChain(ctx context.Context, in *AdminUpdateChainRequest, opts ...grpc.CallOption) (*AdminUpdateChainResponse, error) {
	out := new(AdminUpdateChainResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetChains(context.Context, *GetChainsRequest) (*GetChainsResponse, error)
	// admin
	AdminCreateChain(context.Context, *AdminCreateChainRequest) (*AdminCreateChainResponse, error)
	AdminUpdateChain(context.Context, *AdminUpdateChainRequest) (*AdminUpdateChainResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetChains(context.Context, *GetChainsRequest) (*GetChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChains not implemented")
}
func (UnimplementedGatewayServer) AdminCreateChain(context.Context, *AdminCreateChainRequest) (*AdminCreateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateChain not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateChain(context.Context, *AdminUpdateChainRequest) (*AdminUpdateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateChain not implemented")
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

func _Gateway_GetChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetChains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetChains(ctx, req.(*GetChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateChain(ctx, req.(*AdminCreateChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateChain(ctx, req.(*AdminUpdateChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.chain.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChains",
			Handler:    _Gateway_GetChains_Handler,
		},
		{
			MethodName: "AdminCreateChain",
			Handler:    _Gateway_AdminCreateChain_Handler,
		},
		{
			MethodName: "AdminUpdateChain",
			Handler:    _Gateway_AdminUpdateChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/gw/v1/chain/chain.proto",
}
