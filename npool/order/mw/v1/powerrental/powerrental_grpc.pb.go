// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/powerrental/powerrental.proto

package powerrental

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
	Middleware_CreatePowerRentalOrder_FullMethodName         = "/order.middleware.powerrental.v1.Middleware/CreatePowerRentalOrder"
	Middleware_CreatePowerRentalOrderWithFees_FullMethodName = "/order.middleware.powerrental.v1.Middleware/CreatePowerRentalOrderWithFees"
	Middleware_UpdatePowerRentalOrder_FullMethodName         = "/order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder"
	Middleware_UpdatePowerRentalOrders_FullMethodName        = "/order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrders"
	Middleware_GetPowerRentalOrder_FullMethodName            = "/order.middleware.powerrental.v1.Middleware/GetPowerRentalOrder"
	Middleware_GetPowerRentalOrders_FullMethodName           = "/order.middleware.powerrental.v1.Middleware/GetPowerRentalOrders"
	Middleware_CountPowerRentalOrders_FullMethodName         = "/order.middleware.powerrental.v1.Middleware/CountPowerRentalOrders"
	Middleware_SumPowerRentalOrderUnits_FullMethodName       = "/order.middleware.powerrental.v1.Middleware/SumPowerRentalOrderUnits"
	Middleware_ExistPowerRentalOrder_FullMethodName          = "/order.middleware.powerrental.v1.Middleware/ExistPowerRentalOrder"
	Middleware_ExistPowerRentalOrderConds_FullMethodName     = "/order.middleware.powerrental.v1.Middleware/ExistPowerRentalOrderConds"
	Middleware_DeletePowerRentalOrder_FullMethodName         = "/order.middleware.powerrental.v1.Middleware/DeletePowerRentalOrder"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreatePowerRentalOrder(ctx context.Context, in *CreatePowerRentalOrderRequest, opts ...grpc.CallOption) (*CreatePowerRentalOrderResponse, error)
	CreatePowerRentalOrderWithFees(ctx context.Context, in *CreatePowerRentalOrderWithFeesRequest, opts ...grpc.CallOption) (*CreatePowerRentalOrderWithFeesResponse, error)
	UpdatePowerRentalOrder(ctx context.Context, in *UpdatePowerRentalOrderRequest, opts ...grpc.CallOption) (*UpdatePowerRentalOrderResponse, error)
	UpdatePowerRentalOrders(ctx context.Context, in *UpdatePowerRentalOrdersRequest, opts ...grpc.CallOption) (*UpdatePowerRentalOrdersResponse, error)
	GetPowerRentalOrder(ctx context.Context, in *GetPowerRentalOrderRequest, opts ...grpc.CallOption) (*GetPowerRentalOrderResponse, error)
	GetPowerRentalOrders(ctx context.Context, in *GetPowerRentalOrdersRequest, opts ...grpc.CallOption) (*GetPowerRentalOrdersResponse, error)
	CountPowerRentalOrders(ctx context.Context, in *CountPowerRentalOrdersRequest, opts ...grpc.CallOption) (*CountPowerRentalOrdersResponse, error)
	SumPowerRentalOrderUnits(ctx context.Context, in *SumPowerRentalOrderUnitsRequest, opts ...grpc.CallOption) (*SumPowerRentalOrderUnitsResponse, error)
	ExistPowerRentalOrder(ctx context.Context, in *ExistPowerRentalOrderRequest, opts ...grpc.CallOption) (*ExistPowerRentalOrderResponse, error)
	ExistPowerRentalOrderConds(ctx context.Context, in *ExistPowerRentalOrderCondsRequest, opts ...grpc.CallOption) (*ExistPowerRentalOrderCondsResponse, error)
	DeletePowerRentalOrder(ctx context.Context, in *DeletePowerRentalOrderRequest, opts ...grpc.CallOption) (*DeletePowerRentalOrderResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreatePowerRentalOrder(ctx context.Context, in *CreatePowerRentalOrderRequest, opts ...grpc.CallOption) (*CreatePowerRentalOrderResponse, error) {
	out := new(CreatePowerRentalOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePowerRentalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreatePowerRentalOrderWithFees(ctx context.Context, in *CreatePowerRentalOrderWithFeesRequest, opts ...grpc.CallOption) (*CreatePowerRentalOrderWithFeesResponse, error) {
	out := new(CreatePowerRentalOrderWithFeesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePowerRentalOrderWithFees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePowerRentalOrder(ctx context.Context, in *UpdatePowerRentalOrderRequest, opts ...grpc.CallOption) (*UpdatePowerRentalOrderResponse, error) {
	out := new(UpdatePowerRentalOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePowerRentalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePowerRentalOrders(ctx context.Context, in *UpdatePowerRentalOrdersRequest, opts ...grpc.CallOption) (*UpdatePowerRentalOrdersResponse, error) {
	out := new(UpdatePowerRentalOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePowerRentalOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPowerRentalOrder(ctx context.Context, in *GetPowerRentalOrderRequest, opts ...grpc.CallOption) (*GetPowerRentalOrderResponse, error) {
	out := new(GetPowerRentalOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPowerRentalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPowerRentalOrders(ctx context.Context, in *GetPowerRentalOrdersRequest, opts ...grpc.CallOption) (*GetPowerRentalOrdersResponse, error) {
	out := new(GetPowerRentalOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPowerRentalOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountPowerRentalOrders(ctx context.Context, in *CountPowerRentalOrdersRequest, opts ...grpc.CallOption) (*CountPowerRentalOrdersResponse, error) {
	out := new(CountPowerRentalOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CountPowerRentalOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SumPowerRentalOrderUnits(ctx context.Context, in *SumPowerRentalOrderUnitsRequest, opts ...grpc.CallOption) (*SumPowerRentalOrderUnitsResponse, error) {
	out := new(SumPowerRentalOrderUnitsResponse)
	err := c.cc.Invoke(ctx, Middleware_SumPowerRentalOrderUnits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPowerRentalOrder(ctx context.Context, in *ExistPowerRentalOrderRequest, opts ...grpc.CallOption) (*ExistPowerRentalOrderResponse, error) {
	out := new(ExistPowerRentalOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPowerRentalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPowerRentalOrderConds(ctx context.Context, in *ExistPowerRentalOrderCondsRequest, opts ...grpc.CallOption) (*ExistPowerRentalOrderCondsResponse, error) {
	out := new(ExistPowerRentalOrderCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPowerRentalOrderConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeletePowerRentalOrder(ctx context.Context, in *DeletePowerRentalOrderRequest, opts ...grpc.CallOption) (*DeletePowerRentalOrderResponse, error) {
	out := new(DeletePowerRentalOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_DeletePowerRentalOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreatePowerRentalOrder(context.Context, *CreatePowerRentalOrderRequest) (*CreatePowerRentalOrderResponse, error)
	CreatePowerRentalOrderWithFees(context.Context, *CreatePowerRentalOrderWithFeesRequest) (*CreatePowerRentalOrderWithFeesResponse, error)
	UpdatePowerRentalOrder(context.Context, *UpdatePowerRentalOrderRequest) (*UpdatePowerRentalOrderResponse, error)
	UpdatePowerRentalOrders(context.Context, *UpdatePowerRentalOrdersRequest) (*UpdatePowerRentalOrdersResponse, error)
	GetPowerRentalOrder(context.Context, *GetPowerRentalOrderRequest) (*GetPowerRentalOrderResponse, error)
	GetPowerRentalOrders(context.Context, *GetPowerRentalOrdersRequest) (*GetPowerRentalOrdersResponse, error)
	CountPowerRentalOrders(context.Context, *CountPowerRentalOrdersRequest) (*CountPowerRentalOrdersResponse, error)
	SumPowerRentalOrderUnits(context.Context, *SumPowerRentalOrderUnitsRequest) (*SumPowerRentalOrderUnitsResponse, error)
	ExistPowerRentalOrder(context.Context, *ExistPowerRentalOrderRequest) (*ExistPowerRentalOrderResponse, error)
	ExistPowerRentalOrderConds(context.Context, *ExistPowerRentalOrderCondsRequest) (*ExistPowerRentalOrderCondsResponse, error)
	DeletePowerRentalOrder(context.Context, *DeletePowerRentalOrderRequest) (*DeletePowerRentalOrderResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreatePowerRentalOrder(context.Context, *CreatePowerRentalOrderRequest) (*CreatePowerRentalOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePowerRentalOrder not implemented")
}
func (UnimplementedMiddlewareServer) CreatePowerRentalOrderWithFees(context.Context, *CreatePowerRentalOrderWithFeesRequest) (*CreatePowerRentalOrderWithFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePowerRentalOrderWithFees not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePowerRentalOrder(context.Context, *UpdatePowerRentalOrderRequest) (*UpdatePowerRentalOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePowerRentalOrder not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePowerRentalOrders(context.Context, *UpdatePowerRentalOrdersRequest) (*UpdatePowerRentalOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePowerRentalOrders not implemented")
}
func (UnimplementedMiddlewareServer) GetPowerRentalOrder(context.Context, *GetPowerRentalOrderRequest) (*GetPowerRentalOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRentalOrder not implemented")
}
func (UnimplementedMiddlewareServer) GetPowerRentalOrders(context.Context, *GetPowerRentalOrdersRequest) (*GetPowerRentalOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRentalOrders not implemented")
}
func (UnimplementedMiddlewareServer) CountPowerRentalOrders(context.Context, *CountPowerRentalOrdersRequest) (*CountPowerRentalOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPowerRentalOrders not implemented")
}
func (UnimplementedMiddlewareServer) SumPowerRentalOrderUnits(context.Context, *SumPowerRentalOrderUnitsRequest) (*SumPowerRentalOrderUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumPowerRentalOrderUnits not implemented")
}
func (UnimplementedMiddlewareServer) ExistPowerRentalOrder(context.Context, *ExistPowerRentalOrderRequest) (*ExistPowerRentalOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPowerRentalOrder not implemented")
}
func (UnimplementedMiddlewareServer) ExistPowerRentalOrderConds(context.Context, *ExistPowerRentalOrderCondsRequest) (*ExistPowerRentalOrderCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPowerRentalOrderConds not implemented")
}
func (UnimplementedMiddlewareServer) DeletePowerRentalOrder(context.Context, *DeletePowerRentalOrderRequest) (*DeletePowerRentalOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePowerRentalOrder not implemented")
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

func _Middleware_CreatePowerRentalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePowerRentalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePowerRentalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePowerRentalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePowerRentalOrder(ctx, req.(*CreatePowerRentalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreatePowerRentalOrderWithFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePowerRentalOrderWithFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePowerRentalOrderWithFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePowerRentalOrderWithFees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePowerRentalOrderWithFees(ctx, req.(*CreatePowerRentalOrderWithFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePowerRentalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePowerRentalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePowerRentalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePowerRentalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePowerRentalOrder(ctx, req.(*UpdatePowerRentalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePowerRentalOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePowerRentalOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePowerRentalOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePowerRentalOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePowerRentalOrders(ctx, req.(*UpdatePowerRentalOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPowerRentalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPowerRentalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPowerRentalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPowerRentalOrder(ctx, req.(*GetPowerRentalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPowerRentalOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPowerRentalOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPowerRentalOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPowerRentalOrders(ctx, req.(*GetPowerRentalOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountPowerRentalOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPowerRentalOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountPowerRentalOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountPowerRentalOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountPowerRentalOrders(ctx, req.(*CountPowerRentalOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SumPowerRentalOrderUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumPowerRentalOrderUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SumPowerRentalOrderUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SumPowerRentalOrderUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SumPowerRentalOrderUnits(ctx, req.(*SumPowerRentalOrderUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPowerRentalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPowerRentalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPowerRentalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPowerRentalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPowerRentalOrder(ctx, req.(*ExistPowerRentalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPowerRentalOrderConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPowerRentalOrderCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPowerRentalOrderConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPowerRentalOrderConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPowerRentalOrderConds(ctx, req.(*ExistPowerRentalOrderCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeletePowerRentalOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePowerRentalOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeletePowerRentalOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeletePowerRentalOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeletePowerRentalOrder(ctx, req.(*DeletePowerRentalOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.powerrental.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePowerRentalOrder",
			Handler:    _Middleware_CreatePowerRentalOrder_Handler,
		},
		{
			MethodName: "CreatePowerRentalOrderWithFees",
			Handler:    _Middleware_CreatePowerRentalOrderWithFees_Handler,
		},
		{
			MethodName: "UpdatePowerRentalOrder",
			Handler:    _Middleware_UpdatePowerRentalOrder_Handler,
		},
		{
			MethodName: "UpdatePowerRentalOrders",
			Handler:    _Middleware_UpdatePowerRentalOrders_Handler,
		},
		{
			MethodName: "GetPowerRentalOrder",
			Handler:    _Middleware_GetPowerRentalOrder_Handler,
		},
		{
			MethodName: "GetPowerRentalOrders",
			Handler:    _Middleware_GetPowerRentalOrders_Handler,
		},
		{
			MethodName: "CountPowerRentalOrders",
			Handler:    _Middleware_CountPowerRentalOrders_Handler,
		},
		{
			MethodName: "SumPowerRentalOrderUnits",
			Handler:    _Middleware_SumPowerRentalOrderUnits_Handler,
		},
		{
			MethodName: "ExistPowerRentalOrder",
			Handler:    _Middleware_ExistPowerRentalOrder_Handler,
		},
		{
			MethodName: "ExistPowerRentalOrderConds",
			Handler:    _Middleware_ExistPowerRentalOrderConds_Handler,
		},
		{
			MethodName: "DeletePowerRentalOrder",
			Handler:    _Middleware_DeletePowerRentalOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/powerrental/powerrental.proto",
}
