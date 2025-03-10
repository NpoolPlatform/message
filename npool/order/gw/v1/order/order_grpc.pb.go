// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/order/gw/v1/order/order.proto

package order

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	GetMyOrders(ctx context.Context, in *GetMyOrdersRequest, opts ...grpc.CallOption) (*GetMyOrdersResponse, error)
	AdminGetOrders(ctx context.Context, in *AdminGetOrdersRequest, opts ...grpc.CallOption) (*AdminGetOrdersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/order.gateway.order1.v1.Gateway/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyOrders(ctx context.Context, in *GetMyOrdersRequest, opts ...grpc.CallOption) (*GetMyOrdersResponse, error) {
	out := new(GetMyOrdersResponse)
	err := c.cc.Invoke(ctx, "/order.gateway.order1.v1.Gateway/GetMyOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetOrders(ctx context.Context, in *AdminGetOrdersRequest, opts ...grpc.CallOption) (*AdminGetOrdersResponse, error) {
	out := new(AdminGetOrdersResponse)
	err := c.cc.Invoke(ctx, "/order.gateway.order1.v1.Gateway/AdminGetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	GetMyOrders(context.Context, *GetMyOrdersRequest) (*GetMyOrdersResponse, error)
	AdminGetOrders(context.Context, *AdminGetOrdersRequest) (*AdminGetOrdersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedGatewayServer) GetMyOrders(context.Context, *GetMyOrdersRequest) (*GetMyOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyOrders not implemented")
}
func (UnimplementedGatewayServer) AdminGetOrders(context.Context, *AdminGetOrdersRequest) (*AdminGetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetOrders not implemented")
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

func _Gateway_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.gateway.order1.v1.Gateway/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.gateway.order1.v1.Gateway/GetMyOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyOrders(ctx, req.(*GetMyOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.gateway.order1.v1.Gateway/AdminGetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetOrders(ctx, req.(*AdminGetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.order1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Gateway_GetOrders_Handler,
		},
		{
			MethodName: "GetMyOrders",
			Handler:    _Gateway_GetMyOrders_Handler,
		},
		{
			MethodName: "AdminGetOrders",
			Handler:    _Gateway_AdminGetOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/order/order.proto",
}
