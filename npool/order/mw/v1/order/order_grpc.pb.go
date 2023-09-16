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
	Middleware_CreateOrder_FullMethodName     = "/order.middleware.order1.v1.Middleware/CreateOrder"
	Middleware_CreateOrders_FullMethodName    = "/order.middleware.order1.v1.Middleware/CreateOrders"
	Middleware_UpdateOrder_FullMethodName     = "/order.middleware.order1.v1.Middleware/UpdateOrder"
	Middleware_RollbackOrder_FullMethodName   = "/order.middleware.order1.v1.Middleware/RollbackOrder"
	Middleware_UpdateOrders_FullMethodName    = "/order.middleware.order1.v1.Middleware/UpdateOrders"
	Middleware_GetOrder_FullMethodName        = "/order.middleware.order1.v1.Middleware/GetOrder"
	Middleware_GetOrders_FullMethodName       = "/order.middleware.order1.v1.Middleware/GetOrders"
	Middleware_SumOrderUnits_FullMethodName   = "/order.middleware.order1.v1.Middleware/SumOrderUnits"
	Middleware_CountOrders_FullMethodName     = "/order.middleware.order1.v1.Middleware/CountOrders"
	Middleware_ExistOrder_FullMethodName      = "/order.middleware.order1.v1.Middleware/ExistOrder"
	Middleware_ExistOrderConds_FullMethodName = "/order.middleware.order1.v1.Middleware/ExistOrderConds"
	Middleware_DeleteOrder_FullMethodName     = "/order.middleware.order1.v1.Middleware/DeleteOrder"
	Middleware_DeleteOrders_FullMethodName    = "/order.middleware.order1.v1.Middleware/DeleteOrders"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	RollbackOrder(ctx context.Context, in *RollbackOrderRequest, opts ...grpc.CallOption) (*RollbackOrderResponse, error)
	UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	SumOrderUnits(ctx context.Context, in *SumOrderUnitsRequest, opts ...grpc.CallOption) (*SumOrderUnitsResponse, error)
	CountOrders(ctx context.Context, in *CountOrdersRequest, opts ...grpc.CallOption) (*CountOrdersResponse, error)
	ExistOrder(ctx context.Context, in *ExistOrderRequest, opts ...grpc.CallOption) (*ExistOrderResponse, error)
	ExistOrderConds(ctx context.Context, in *ExistOrderCondsRequest, opts ...grpc.CallOption) (*ExistOrderCondsResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
	DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateOrders(ctx context.Context, in *CreateOrdersRequest, opts ...grpc.CallOption) (*CreateOrdersResponse, error) {
	out := new(CreateOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) RollbackOrder(ctx context.Context, in *RollbackOrderRequest, opts ...grpc.CallOption) (*RollbackOrderResponse, error) {
	out := new(RollbackOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_RollbackOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateOrders(ctx context.Context, in *UpdateOrdersRequest, opts ...grpc.CallOption) (*UpdateOrdersResponse, error) {
	out := new(UpdateOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *middlewareClient) SumOrderUnits(ctx context.Context, in *SumOrderUnitsRequest, opts ...grpc.CallOption) (*SumOrderUnitsResponse, error) {
	out := new(SumOrderUnitsResponse)
	err := c.cc.Invoke(ctx, Middleware_SumOrderUnits_FullMethodName, in, out, opts...)
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

func (c *middlewareClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error) {
	out := new(DeleteOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	RollbackOrder(context.Context, *RollbackOrderRequest) (*RollbackOrderResponse, error)
	UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	SumOrderUnits(context.Context, *SumOrderUnitsRequest) (*SumOrderUnitsResponse, error)
	CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error)
	ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error)
	ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedMiddlewareServer) CreateOrders(context.Context, *CreateOrdersRequest) (*CreateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedMiddlewareServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedMiddlewareServer) RollbackOrder(context.Context, *RollbackOrderRequest) (*RollbackOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackOrder not implemented")
}
func (UnimplementedMiddlewareServer) UpdateOrders(context.Context, *UpdateOrdersRequest) (*UpdateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (UnimplementedMiddlewareServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedMiddlewareServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedMiddlewareServer) SumOrderUnits(context.Context, *SumOrderUnitsRequest) (*SumOrderUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumOrderUnits not implemented")
}
func (UnimplementedMiddlewareServer) CountOrders(context.Context, *CountOrdersRequest) (*CountOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountOrders not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrder(context.Context, *ExistOrderRequest) (*ExistOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrder not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrderConds(context.Context, *ExistOrderCondsRequest) (*ExistOrderCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedMiddlewareServer) DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrders not implemented")
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

func _Middleware_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOrders(ctx, req.(*CreateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_RollbackOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).RollbackOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_RollbackOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).RollbackOrder(ctx, req.(*RollbackOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateOrders(ctx, req.(*UpdateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Middleware_SumOrderUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumOrderUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SumOrderUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SumOrderUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SumOrderUnits(ctx, req.(*SumOrderUnitsRequest))
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

func _Middleware_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteOrders(ctx, req.(*DeleteOrdersRequest))
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
			MethodName: "CreateOrder",
			Handler:    _Middleware_CreateOrder_Handler,
		},
		{
			MethodName: "CreateOrders",
			Handler:    _Middleware_CreateOrders_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Middleware_UpdateOrder_Handler,
		},
		{
			MethodName: "RollbackOrder",
			Handler:    _Middleware_RollbackOrder_Handler,
		},
		{
			MethodName: "UpdateOrders",
			Handler:    _Middleware_UpdateOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Middleware_GetOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Middleware_GetOrders_Handler,
		},
		{
			MethodName: "SumOrderUnits",
			Handler:    _Middleware_SumOrderUnits_Handler,
		},
		{
			MethodName: "CountOrders",
			Handler:    _Middleware_CountOrders_Handler,
		},
		{
			MethodName: "ExistOrder",
			Handler:    _Middleware_ExistOrder_Handler,
		},
		{
			MethodName: "ExistOrderConds",
			Handler:    _Middleware_ExistOrderConds_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Middleware_DeleteOrder_Handler,
		},
		{
			MethodName: "DeleteOrders",
			Handler:    _Middleware_DeleteOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/order/order.proto",
}
