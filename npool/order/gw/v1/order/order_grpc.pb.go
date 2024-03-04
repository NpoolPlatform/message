// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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

const (
	Gateway_CreateOrder_FullMethodName          = "/order.gateway.order1.v1.Gateway/CreateOrder"
	Gateway_CreateOrders_FullMethodName         = "/order.gateway.order1.v1.Gateway/CreateOrders"
	Gateway_UpdateOrder_FullMethodName          = "/order.gateway.order1.v1.Gateway/UpdateOrder"
	Gateway_UpdateUserOrder_FullMethodName      = "/order.gateway.order1.v1.Gateway/UpdateUserOrder"
	Gateway_UpdateAppUserOrder_FullMethodName   = "/order.gateway.order1.v1.Gateway/UpdateAppUserOrder"
	Gateway_GetOrder_FullMethodName             = "/order.gateway.order1.v1.Gateway/GetOrder"
	Gateway_GetOrders_FullMethodName            = "/order.gateway.order1.v1.Gateway/GetOrders"
	Gateway_CreateUserOrder_FullMethodName      = "/order.gateway.order1.v1.Gateway/CreateUserOrder"
	Gateway_CreateAppUserOrder_FullMethodName   = "/order.gateway.order1.v1.Gateway/CreateAppUserOrder"
	Gateway_GetUserOrders_FullMethodName        = "/order.gateway.order1.v1.Gateway/GetUserOrders"
	Gateway_GetAppUserOrders_FullMethodName     = "/order.gateway.order1.v1.Gateway/GetAppUserOrders"
	Gateway_GetAppOrders_FullMethodName         = "/order.gateway.order1.v1.Gateway/GetAppOrders"
	Gateway_GetNAppOrders_FullMethodName        = "/order.gateway.order1.v1.Gateway/GetNAppOrders"
	Gateway_CreateSimulateOrder_FullMethodName  = "/order.gateway.order1.v1.Gateway/CreateSimulateOrder"
	Gateway_CreateSimulateOrders_FullMethodName = "/order.gateway.order1.v1.Gateway/CreateSimulateOrders"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	UpdateUserOrder(ctx context.Context, in *UpdateUserOrderRequest, opts ...grpc.CallOption) (*UpdateUserOrderResponse, error)
	UpdateAppUserOrder(ctx context.Context, in *UpdateAppUserOrderRequest, opts ...grpc.CallOption) (*UpdateAppUserOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	CreateUserOrder(ctx context.Context, in *CreateUserOrderRequest, opts ...grpc.CallOption) (*CreateUserOrderResponse, error)
	CreateAppUserOrder(ctx context.Context, in *CreateAppUserOrderRequest, opts ...grpc.CallOption) (*CreateAppUserOrderResponse, error)
	GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...grpc.CallOption) (*GetUserOrdersResponse, error)
	GetAppUserOrders(ctx context.Context, in *GetAppUserOrdersRequest, opts ...grpc.CallOption) (*GetAppUserOrdersResponse, error)
	GetAppOrders(ctx context.Context, in *GetAppOrdersRequest, opts ...grpc.CallOption) (*GetAppOrdersResponse, error)
	GetNAppOrders(ctx context.Context, in *GetNAppOrdersRequest, opts ...grpc.CallOption) (*GetNAppOrdersResponse, error)
	CreateSimulateOrder(ctx context.Context, in *CreateSimulateOrderRequest, opts ...grpc.CallOption) (*CreateSimulateOrderResponse, error)
	CreateSimulateOrders(ctx context.Context, in *CreateSimulateOrdersRequest, opts ...grpc.CallOption) (*CreateSimulateOrdersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error) {
	out := new(CreateOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateUserOrder(ctx context.Context, in *UpdateUserOrderRequest, opts ...grpc.CallOption) (*UpdateUserOrderResponse, error) {
	out := new(UpdateUserOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateUserOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppUserOrder(ctx context.Context, in *UpdateAppUserOrderRequest, opts ...grpc.CallOption) (*UpdateAppUserOrderResponse, error) {
	out := new(UpdateAppUserOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppUserOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateUserOrder(ctx context.Context, in *CreateUserOrderRequest, opts ...grpc.CallOption) (*CreateUserOrderResponse, error) {
	out := new(CreateUserOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateUserOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppUserOrder(ctx context.Context, in *CreateAppUserOrderRequest, opts ...grpc.CallOption) (*CreateAppUserOrderResponse, error) {
	out := new(CreateAppUserOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppUserOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetUserOrders(ctx context.Context, in *GetUserOrdersRequest, opts ...grpc.CallOption) (*GetUserOrdersResponse, error) {
	out := new(GetUserOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetUserOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppUserOrders(ctx context.Context, in *GetAppUserOrdersRequest, opts ...grpc.CallOption) (*GetAppUserOrdersResponse, error) {
	out := new(GetAppUserOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppUserOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppOrders(ctx context.Context, in *GetAppOrdersRequest, opts ...grpc.CallOption) (*GetAppOrdersResponse, error) {
	out := new(GetAppOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppOrders(ctx context.Context, in *GetNAppOrdersRequest, opts ...grpc.CallOption) (*GetNAppOrdersResponse, error) {
	out := new(GetNAppOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetNAppOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateSimulateOrder(ctx context.Context, in *CreateSimulateOrderRequest, opts ...grpc.CallOption) (*CreateSimulateOrderResponse, error) {
	out := new(CreateSimulateOrderResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateSimulateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateSimulateOrders(ctx context.Context, in *CreateSimulateOrdersRequest, opts ...grpc.CallOption) (*CreateSimulateOrdersResponse, error) {
	out := new(CreateSimulateOrdersResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateSimulateOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	UpdateUserOrder(context.Context, *UpdateUserOrderRequest) (*UpdateUserOrderResponse, error)
	UpdateAppUserOrder(context.Context, *UpdateAppUserOrderRequest) (*UpdateAppUserOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	CreateUserOrder(context.Context, *CreateUserOrderRequest) (*CreateUserOrderResponse, error)
	CreateAppUserOrder(context.Context, *CreateAppUserOrderRequest) (*CreateAppUserOrderResponse, error)
	GetUserOrders(context.Context, *GetUserOrdersRequest) (*GetUserOrdersResponse, error)
	GetAppUserOrders(context.Context, *GetAppUserOrdersRequest) (*GetAppUserOrdersResponse, error)
	GetAppOrders(context.Context, *GetAppOrdersRequest) (*GetAppOrdersResponse, error)
	GetNAppOrders(context.Context, *GetNAppOrdersRequest) (*GetNAppOrdersResponse, error)
	CreateSimulateOrder(context.Context, *CreateSimulateOrderRequest) (*CreateSimulateOrderResponse, error)
	CreateSimulateOrders(context.Context, *CreateSimulateOrdersRequest) (*CreateSimulateOrdersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedGatewayServer) CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedGatewayServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedGatewayServer) UpdateUserOrder(context.Context, *UpdateUserOrderRequest) (*UpdateUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserOrder not implemented")
}
func (UnimplementedGatewayServer) UpdateAppUserOrder(context.Context, *UpdateAppUserOrderRequest) (*UpdateAppUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppUserOrder not implemented")
}
func (UnimplementedGatewayServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedGatewayServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedGatewayServer) CreateUserOrder(context.Context, *CreateUserOrderRequest) (*CreateUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserOrder not implemented")
}
func (UnimplementedGatewayServer) CreateAppUserOrder(context.Context, *CreateAppUserOrderRequest) (*CreateAppUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUserOrder not implemented")
}
func (UnimplementedGatewayServer) GetUserOrders(context.Context, *GetUserOrdersRequest) (*GetUserOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOrders not implemented")
}
func (UnimplementedGatewayServer) GetAppUserOrders(context.Context, *GetAppUserOrdersRequest) (*GetAppUserOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserOrders not implemented")
}
func (UnimplementedGatewayServer) GetAppOrders(context.Context, *GetAppOrdersRequest) (*GetAppOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppOrders not implemented")
}
func (UnimplementedGatewayServer) GetNAppOrders(context.Context, *GetNAppOrdersRequest) (*GetNAppOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppOrders not implemented")
}
func (UnimplementedGatewayServer) CreateSimulateOrder(context.Context, *CreateSimulateOrderRequest) (*CreateSimulateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSimulateOrder not implemented")
}
func (UnimplementedGatewayServer) CreateSimulateOrders(context.Context, *CreateSimulateOrdersRequest) (*CreateSimulateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSimulateOrders not implemented")
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

func _Gateway_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateOrders(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateUserOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateUserOrder(ctx, req.(*UpdateUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppUserOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppUserOrder(ctx, req.(*UpdateAppUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: Gateway_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateUserOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateUserOrder(ctx, req.(*CreateUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppUserOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppUserOrder(ctx, req.(*CreateAppUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetUserOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetUserOrders(ctx, req.(*GetUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppUserOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppUserOrders(ctx, req.(*GetAppUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppOrders(ctx, req.(*GetAppOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppOrders(ctx, req.(*GetNAppOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateSimulateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSimulateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateSimulateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateSimulateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateSimulateOrder(ctx, req.(*CreateSimulateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateSimulateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSimulateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateSimulateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateSimulateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateSimulateOrders(ctx, req.(*CreateSimulateOrdersRequest))
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
			MethodName: "CreateOrder",
			Handler:    _Gateway_CreateOrder_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _Gateway_CreateOrders_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Gateway_UpdateOrder_Handler,
		},
		{
			MethodName: "UpdateUserOrder",
			Handler:    _Gateway_UpdateUserOrder_Handler,
		},
		{
			MethodName: "UpdateAppUserOrder",
			Handler:    _Gateway_UpdateAppUserOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Gateway_GetOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Gateway_GetOrders_Handler,
		},
		{
			MethodName: "CreateUserOrder",
			Handler:    _Gateway_CreateUserOrder_Handler,
		},
		{
			MethodName: "CreateAppUserOrder",
			Handler:    _Gateway_CreateAppUserOrder_Handler,
		},
		{
			MethodName: "GetUserOrders",
			Handler:    _Gateway_GetUserOrders_Handler,
		},
		{
			MethodName: "GetAppUserOrders",
			Handler:    _Gateway_GetAppUserOrders_Handler,
		},
		{
			MethodName: "GetAppOrders",
			Handler:    _Gateway_GetAppOrders_Handler,
		},
		{
			MethodName: "GetNAppOrders",
			Handler:    _Gateway_GetNAppOrders_Handler,
		},
		{
			MethodName: "CreateSimulateOrder",
			Handler:    _Gateway_CreateSimulateOrder_Handler,
		},
		{
			MethodName: "CreateSimulateOrders",
			Handler:    _Gateway_CreateSimulateOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/order/order.proto",
}
