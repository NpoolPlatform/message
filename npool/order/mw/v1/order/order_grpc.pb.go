// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/order/order.proto

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
	Middleware_GetOrder_FullMethodName            = "/order.middleware.order1.v1.Middleware/GetOrder"
	Middleware_GetOrders_FullMethodName           = "/order.middleware.order1.v1.Middleware/GetOrders"
	Middleware_CountOrders_FullMethodName         = "/order.middleware.order1.v1.Middleware/CountOrders"
	Middleware_SumOrdersPaymentUSD_FullMethodName = "/order.middleware.order1.v1.Middleware/SumOrdersPaymentUSD"
	Middleware_SumOrdersValueUSD_FullMethodName   = "/order.middleware.order1.v1.Middleware/SumOrdersValueUSD"
	Middleware_ExistOrder_FullMethodName          = "/order.middleware.order1.v1.Middleware/ExistOrder"
	Middleware_ExistOrderConds_FullMethodName     = "/order.middleware.order1.v1.Middleware/ExistOrderConds"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	CountOrders(ctx context.Context, in *CountOrdersRequest, opts ...grpc.CallOption) (*CountOrdersResponse, error)
	SumOrdersPaymentUSD(ctx context.Context, in *SumOrdersPaymentUSDRequest, opts ...grpc.CallOption) (*SumOrdersPaymentUSDResponse, error)
	SumOrdersValueUSD(ctx context.Context, in *SumOrdersValueUSDRequest, opts ...grpc.CallOption) (*SumOrdersValueUSDResponse, error)
	ExistOrder(ctx context.Context, in *ExistOrderRequest, opts ...grpc.CallOption) (*ExistOrderResponse, error)
	ExistOrderConds(ctx context.Context, in *ExistOrderCondsRequest, opts ...grpc.CallOption) (*ExistOrderCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountOrders(ctx context.Context, in *CountOrdersRequest, opts ...grpc.CallOption) (*CountOrdersResponse, error) {
	out := new(CountOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CountOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SumOrdersPaymentUSD(ctx context.Context, in *SumOrdersPaymentUSDRequest, opts ...grpc.CallOption) (*SumOrdersPaymentUSDResponse, error) {
	out := new(SumOrdersPaymentUSDResponse)
	err := c.cc.Invoke(ctx, Middleware_SumOrdersPaymentUSD_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SumOrdersValueUSD(ctx context.Context, in *SumOrdersValueUSDRequest, opts ...grpc.CallOption) (*SumOrdersValueUSDResponse, error) {
	out := new(SumOrdersValueUSDResponse)
	err := c.cc.Invoke(ctx, Middleware_SumOrdersValueUSD_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOrder(ctx context.Context, in *ExistOrderRequest, opts ...grpc.CallOption) (*ExistOrderResponse, error) {
	out := new(ExistOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOrderConds(ctx context.Context, in *ExistOrderCondsRequest, opts ...grpc.CallOption) (*ExistOrderCondsResponse, error) {
	out := new(ExistOrderCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOrderConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error)
	SumOrdersPaymentUSD(context.Context, *SumOrdersPaymentUSDRequest) (*SumOrdersPaymentUSDResponse, error)
	SumOrdersValueUSD(context.Context, *SumOrdersValueUSDRequest) (*SumOrdersValueUSDResponse, error)
	ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error)
	ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedMiddlewareServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedMiddlewareServer) CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountOrders not implemented")
}
func (UnimplementedMiddlewareServer) SumOrdersPaymentUSD(context.Context, *SumOrdersPaymentUSDRequest) (*SumOrdersPaymentUSDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumOrdersPaymentUSD not implemented")
}
func (UnimplementedMiddlewareServer) SumOrdersValueUSD(context.Context, *SumOrdersValueUSDRequest) (*SumOrdersValueUSDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumOrdersValueUSD not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrder not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderConds not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountOrders(ctx, req.(*CountOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SumOrdersPaymentUSD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumOrdersPaymentUSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SumOrdersPaymentUSD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SumOrdersPaymentUSD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SumOrdersPaymentUSD(ctx, req.(*SumOrdersPaymentUSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SumOrdersValueUSD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumOrdersValueUSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SumOrdersValueUSD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SumOrdersValueUSD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SumOrdersValueUSD(ctx, req.(*SumOrdersValueUSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOrder(ctx, req.(*ExistOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOrderConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOrderConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOrderConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOrderConds(ctx, req.(*ExistOrderCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.order1.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _Middleware_GetOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Middleware_GetOrders_Handler,
		},
		{
			MethodName: "CountOrders",
			Handler:    _Middleware_CountOrders_Handler,
		},
		{
			MethodName: "SumOrdersPaymentUSD",
			Handler:    _Middleware_SumOrdersPaymentUSD_Handler,
		},
		{
			MethodName: "SumOrdersValueUSD",
			Handler:    _Middleware_SumOrdersValueUSD_Handler,
		},
		{
			MethodName: "ExistOrder",
			Handler:    _Middleware_ExistOrder_Handler,
		},
		{
			MethodName: "ExistOrderConds",
			Handler:    _Middleware_ExistOrderConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/order/order.proto",
}
