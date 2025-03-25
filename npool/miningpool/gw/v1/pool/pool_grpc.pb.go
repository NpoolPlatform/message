// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/gw/v1/pool/pool.proto

package pool

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
	Gateway_AdminCreatePool_FullMethodName = "/miningpool.gateway.pool.v1.Gateway/AdminCreatePool"
	Gateway_AdminUpdatePool_FullMethodName = "/miningpool.gateway.pool.v1.Gateway/AdminUpdatePool"
	Gateway_AdminGetPools_FullMethodName   = "/miningpool.gateway.pool.v1.Gateway/AdminGetPools"
	Gateway_AdminDeletePool_FullMethodName = "/miningpool.gateway.pool.v1.Gateway/AdminDeletePool"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreatePool(ctx context.Context, in *AdminCreatePoolRequest, opts ...grpc.CallOption) (*AdminCreatePoolResponse, error)
	AdminUpdatePool(ctx context.Context, in *AdminUpdatePoolRequest, opts ...grpc.CallOption) (*AdminUpdatePoolResponse, error)
	AdminGetPools(ctx context.Context, in *AdminGetPoolsRequest, opts ...grpc.CallOption) (*AdminGetPoolsResponse, error)
	AdminDeletePool(ctx context.Context, in *AdminDeletePoolRequest, opts ...grpc.CallOption) (*AdminDeletePoolResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreatePool(ctx context.Context, in *AdminCreatePoolRequest, opts ...grpc.CallOption) (*AdminCreatePoolResponse, error) {
	out := new(AdminCreatePoolResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdatePool(ctx context.Context, in *AdminUpdatePoolRequest, opts ...grpc.CallOption) (*AdminUpdatePoolResponse, error) {
	out := new(AdminUpdatePoolResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetPools(ctx context.Context, in *AdminGetPoolsRequest, opts ...grpc.CallOption) (*AdminGetPoolsResponse, error) {
	out := new(AdminGetPoolsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeletePool(ctx context.Context, in *AdminDeletePoolRequest, opts ...grpc.CallOption) (*AdminDeletePoolResponse, error) {
	out := new(AdminDeletePoolResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeletePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreatePool(context.Context, *AdminCreatePoolRequest) (*AdminCreatePoolResponse, error)
	AdminUpdatePool(context.Context, *AdminUpdatePoolRequest) (*AdminUpdatePoolResponse, error)
	AdminGetPools(context.Context, *AdminGetPoolsRequest) (*AdminGetPoolsResponse, error)
	AdminDeletePool(context.Context, *AdminDeletePoolRequest) (*AdminDeletePoolResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreatePool(context.Context, *AdminCreatePoolRequest) (*AdminCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreatePool not implemented")
}
func (UnimplementedGatewayServer) AdminUpdatePool(context.Context, *AdminUpdatePoolRequest) (*AdminUpdatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdatePool not implemented")
}
func (UnimplementedGatewayServer) AdminGetPools(context.Context, *AdminGetPoolsRequest) (*AdminGetPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetPools not implemented")
}
func (UnimplementedGatewayServer) AdminDeletePool(context.Context, *AdminDeletePoolRequest) (*AdminDeletePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeletePool not implemented")
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

func _Gateway_AdminCreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreatePool(ctx, req.(*AdminCreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdatePool(ctx, req.(*AdminUpdatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetPools(ctx, req.(*AdminGetPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeletePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeletePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeletePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeletePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeletePool(ctx, req.(*AdminDeletePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.gateway.pool.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreatePool",
			Handler:    _Gateway_AdminCreatePool_Handler,
		},
		{
			MethodName: "AdminUpdatePool",
			Handler:    _Gateway_AdminUpdatePool_Handler,
		},
		{
			MethodName: "AdminGetPools",
			Handler:    _Gateway_AdminGetPools_Handler,
		},
		{
			MethodName: "AdminDeletePool",
			Handler:    _Gateway_AdminDeletePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/gw/v1/pool/pool.proto",
}
