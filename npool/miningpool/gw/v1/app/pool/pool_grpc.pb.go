// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/gw/v1/app/pool/pool.proto

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
	Gateway_AdminCreatePool_FullMethodName = "/miningpool.gateway.app.pool.v1.Gateway/AdminCreatePool"
	Gateway_GetPool_FullMethodName         = "/miningpool.gateway.app.pool.v1.Gateway/GetPool"
	Gateway_GetPools_FullMethodName        = "/miningpool.gateway.app.pool.v1.Gateway/GetPools"
	Gateway_AdminGetNPools_FullMethodName  = "/miningpool.gateway.app.pool.v1.Gateway/AdminGetNPools"
	Gateway_AdminDeletePool_FullMethodName = "/miningpool.gateway.app.pool.v1.Gateway/AdminDeletePool"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreatePool(ctx context.Context, in *AdminCreatePoolRequest, opts ...grpc.CallOption) (*AdminCreatePoolResponse, error)
	GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error)
	GetPools(ctx context.Context, in *GetPoolsRequest, opts ...grpc.CallOption) (*GetPoolsResponse, error)
	AdminGetNPools(ctx context.Context, in *AdminGetNPoolsRequest, opts ...grpc.CallOption) (*AdminGetNPoolsResponse, error)
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

func (c *gatewayClient) GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error) {
	out := new(GetPoolResponse)
	err := c.cc.Invoke(ctx, Gateway_GetPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetPools(ctx context.Context, in *GetPoolsRequest, opts ...grpc.CallOption) (*GetPoolsResponse, error) {
	out := new(GetPoolsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetNPools(ctx context.Context, in *AdminGetNPoolsRequest, opts ...grpc.CallOption) (*AdminGetNPoolsResponse, error) {
	out := new(AdminGetNPoolsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetNPools_FullMethodName, in, out, opts...)
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
	GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error)
	GetPools(context.Context, *GetPoolsRequest) (*GetPoolsResponse, error)
	AdminGetNPools(context.Context, *AdminGetNPoolsRequest) (*AdminGetNPoolsResponse, error)
	AdminDeletePool(context.Context, *AdminDeletePoolRequest) (*AdminDeletePoolResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreatePool(context.Context, *AdminCreatePoolRequest) (*AdminCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreatePool not implemented")
}
func (UnimplementedGatewayServer) GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPool not implemented")
}
func (UnimplementedGatewayServer) GetPools(context.Context, *GetPoolsRequest) (*GetPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPools not implemented")
}
func (UnimplementedGatewayServer) AdminGetNPools(context.Context, *AdminGetNPoolsRequest) (*AdminGetNPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetNPools not implemented")
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

func _Gateway_GetPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetPool(ctx, req.(*GetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetPools(ctx, req.(*GetPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetNPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetNPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetNPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetNPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetNPools(ctx, req.(*AdminGetNPoolsRequest))
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
	ServiceName: "miningpool.gateway.app.pool.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreatePool",
			Handler:    _Gateway_AdminCreatePool_Handler,
		},
		{
			MethodName: "GetPool",
			Handler:    _Gateway_GetPool_Handler,
		},
		{
			MethodName: "GetPools",
			Handler:    _Gateway_GetPools_Handler,
		},
		{
			MethodName: "AdminGetNPools",
			Handler:    _Gateway_AdminGetNPools_Handler,
		},
		{
			MethodName: "AdminDeletePool",
			Handler:    _Gateway_AdminDeletePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/gw/v1/app/pool/pool.proto",
}
