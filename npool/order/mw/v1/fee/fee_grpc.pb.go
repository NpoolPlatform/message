// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/fee/fee.proto

package fee

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
	Middleware_CreateFeeOrder_FullMethodName     = "/order.middleware.fee.v1.Middleware/CreateFeeOrder"
	Middleware_CreateFeeOrders_FullMethodName    = "/order.middleware.fee.v1.Middleware/CreateFeeOrders"
	Middleware_UpdateFeeOrder_FullMethodName     = "/order.middleware.fee.v1.Middleware/UpdateFeeOrder"
	Middleware_GetFeeOrder_FullMethodName        = "/order.middleware.fee.v1.Middleware/GetFeeOrder"
	Middleware_GetFeeOrders_FullMethodName       = "/order.middleware.fee.v1.Middleware/GetFeeOrders"
	Middleware_CountFeeOrders_FullMethodName     = "/order.middleware.fee.v1.Middleware/CountFeeOrders"
	Middleware_ExistFeeOrder_FullMethodName      = "/order.middleware.fee.v1.Middleware/ExistFeeOrder"
	Middleware_ExistFeeOrderConds_FullMethodName = "/order.middleware.fee.v1.Middleware/ExistFeeOrderConds"
	Middleware_DeleteFeeOrder_FullMethodName     = "/order.middleware.fee.v1.Middleware/DeleteFeeOrder"
	Middleware_DeleteFeeOrders_FullMethodName    = "/order.middleware.fee.v1.Middleware/DeleteFeeOrders"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateFeeOrder(ctx context.Context, in *CreateFeeOrderRequest, opts ...grpc.CallOption) (*CreateFeeOrderResponse, error)
	CreateFeeOrders(ctx context.Context, in *CreateFeeOrdersRequest, opts ...grpc.CallOption) (*CreateFeeOrdersResponse, error)
	UpdateFeeOrder(ctx context.Context, in *UpdateFeeOrderRequest, opts ...grpc.CallOption) (*UpdateFeeOrderResponse, error)
	GetFeeOrder(ctx context.Context, in *GetFeeOrderRequest, opts ...grpc.CallOption) (*GetFeeOrderResponse, error)
	GetFeeOrders(ctx context.Context, in *GetFeeOrdersRequest, opts ...grpc.CallOption) (*GetFeeOrdersResponse, error)
	CountFeeOrders(ctx context.Context, in *CountFeeOrdersRequest, opts ...grpc.CallOption) (*CountFeeOrdersResponse, error)
	ExistFeeOrder(ctx context.Context, in *ExistFeeOrderRequest, opts ...grpc.CallOption) (*ExistFeeOrderResponse, error)
	ExistFeeOrderConds(ctx context.Context, in *ExistFeeOrderCondsRequest, opts ...grpc.CallOption) (*ExistFeeOrderCondsResponse, error)
	DeleteFeeOrder(ctx context.Context, in *DeleteFeeOrderRequest, opts ...grpc.CallOption) (*DeleteFeeOrderResponse, error)
	DeleteFeeOrders(ctx context.Context, in *DeleteFeeOrdersRequest, opts ...grpc.CallOption) (*DeleteFeeOrdersResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFeeOrder(ctx context.Context, in *CreateFeeOrderRequest, opts ...grpc.CallOption) (*CreateFeeOrderResponse, error) {
	out := new(CreateFeeOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFeeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateFeeOrders(ctx context.Context, in *CreateFeeOrdersRequest, opts ...grpc.CallOption) (*CreateFeeOrdersResponse, error) {
	out := new(CreateFeeOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFeeOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateFeeOrder(ctx context.Context, in *UpdateFeeOrderRequest, opts ...grpc.CallOption) (*UpdateFeeOrderResponse, error) {
	out := new(UpdateFeeOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateFeeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFeeOrder(ctx context.Context, in *GetFeeOrderRequest, opts ...grpc.CallOption) (*GetFeeOrderResponse, error) {
	out := new(GetFeeOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFeeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFeeOrders(ctx context.Context, in *GetFeeOrdersRequest, opts ...grpc.CallOption) (*GetFeeOrdersResponse, error) {
	out := new(GetFeeOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFeeOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountFeeOrders(ctx context.Context, in *CountFeeOrdersRequest, opts ...grpc.CallOption) (*CountFeeOrdersResponse, error) {
	out := new(CountFeeOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CountFeeOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFeeOrder(ctx context.Context, in *ExistFeeOrderRequest, opts ...grpc.CallOption) (*ExistFeeOrderResponse, error) {
	out := new(ExistFeeOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFeeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFeeOrderConds(ctx context.Context, in *ExistFeeOrderCondsRequest, opts ...grpc.CallOption) (*ExistFeeOrderCondsResponse, error) {
	out := new(ExistFeeOrderCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFeeOrderConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFeeOrder(ctx context.Context, in *DeleteFeeOrderRequest, opts ...grpc.CallOption) (*DeleteFeeOrderResponse, error) {
	out := new(DeleteFeeOrderResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteFeeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFeeOrders(ctx context.Context, in *DeleteFeeOrdersRequest, opts ...grpc.CallOption) (*DeleteFeeOrdersResponse, error) {
	out := new(DeleteFeeOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteFeeOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFeeOrder(context.Context, *CreateFeeOrderRequest) (*CreateFeeOrderResponse, error)
	CreateFeeOrders(context.Context, *CreateFeeOrdersRequest) (*CreateFeeOrdersResponse, error)
	UpdateFeeOrder(context.Context, *UpdateFeeOrderRequest) (*UpdateFeeOrderResponse, error)
	GetFeeOrder(context.Context, *GetFeeOrderRequest) (*GetFeeOrderResponse, error)
	GetFeeOrders(context.Context, *GetFeeOrdersRequest) (*GetFeeOrdersResponse, error)
	CountFeeOrders(context.Context, *CountFeeOrdersRequest) (*CountFeeOrdersResponse, error)
	ExistFeeOrder(context.Context, *ExistFeeOrderRequest) (*ExistFeeOrderResponse, error)
	ExistFeeOrderConds(context.Context, *ExistFeeOrderCondsRequest) (*ExistFeeOrderCondsResponse, error)
	DeleteFeeOrder(context.Context, *DeleteFeeOrderRequest) (*DeleteFeeOrderResponse, error)
	DeleteFeeOrders(context.Context, *DeleteFeeOrdersRequest) (*DeleteFeeOrdersResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFeeOrder(context.Context, *CreateFeeOrderRequest) (*CreateFeeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeeOrder not implemented")
}
func (UnimplementedMiddlewareServer) CreateFeeOrders(context.Context, *CreateFeeOrdersRequest) (*CreateFeeOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeeOrders not implemented")
}
func (UnimplementedMiddlewareServer) UpdateFeeOrder(context.Context, *UpdateFeeOrderRequest) (*UpdateFeeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeeOrder not implemented")
}
func (UnimplementedMiddlewareServer) GetFeeOrder(context.Context, *GetFeeOrderRequest) (*GetFeeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeOrder not implemented")
}
func (UnimplementedMiddlewareServer) GetFeeOrders(context.Context, *GetFeeOrdersRequest) (*GetFeeOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeOrders not implemented")
}
func (UnimplementedMiddlewareServer) CountFeeOrders(context.Context, *CountFeeOrdersRequest) (*CountFeeOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFeeOrders not implemented")
}
func (UnimplementedMiddlewareServer) ExistFeeOrder(context.Context, *ExistFeeOrderRequest) (*ExistFeeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFeeOrder not implemented")
}
func (UnimplementedMiddlewareServer) ExistFeeOrderConds(context.Context, *ExistFeeOrderCondsRequest) (*ExistFeeOrderCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFeeOrderConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFeeOrder(context.Context, *DeleteFeeOrderRequest) (*DeleteFeeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeeOrder not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFeeOrders(context.Context, *DeleteFeeOrdersRequest) (*DeleteFeeOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeeOrders not implemented")
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

func _Middleware_CreateFeeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFeeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFeeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFeeOrder(ctx, req.(*CreateFeeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateFeeOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeeOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFeeOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFeeOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFeeOrders(ctx, req.(*CreateFeeOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateFeeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateFeeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateFeeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateFeeOrder(ctx, req.(*UpdateFeeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFeeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFeeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFeeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFeeOrder(ctx, req.(*GetFeeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFeeOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFeeOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFeeOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFeeOrders(ctx, req.(*GetFeeOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountFeeOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFeeOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountFeeOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountFeeOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountFeeOrders(ctx, req.(*CountFeeOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFeeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFeeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFeeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFeeOrder(ctx, req.(*ExistFeeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFeeOrderConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeOrderCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFeeOrderConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFeeOrderConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFeeOrderConds(ctx, req.(*ExistFeeOrderCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFeeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFeeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteFeeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFeeOrder(ctx, req.(*DeleteFeeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFeeOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeeOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFeeOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteFeeOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFeeOrders(ctx, req.(*DeleteFeeOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.fee.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeeOrder",
			Handler:    _Middleware_CreateFeeOrder_Handler,
		},
		{
			MethodName: "CreateFeeOrders",
			Handler:    _Middleware_CreateFeeOrders_Handler,
		},
		{
			MethodName: "UpdateFeeOrder",
			Handler:    _Middleware_UpdateFeeOrder_Handler,
		},
		{
			MethodName: "GetFeeOrder",
			Handler:    _Middleware_GetFeeOrder_Handler,
		},
		{
			MethodName: "GetFeeOrders",
			Handler:    _Middleware_GetFeeOrders_Handler,
		},
		{
			MethodName: "CountFeeOrders",
			Handler:    _Middleware_CountFeeOrders_Handler,
		},
		{
			MethodName: "ExistFeeOrder",
			Handler:    _Middleware_ExistFeeOrder_Handler,
		},
		{
			MethodName: "ExistFeeOrderConds",
			Handler:    _Middleware_ExistFeeOrderConds_Handler,
		},
		{
			MethodName: "DeleteFeeOrder",
			Handler:    _Middleware_DeleteFeeOrder_Handler,
		},
		{
			MethodName: "DeleteFeeOrders",
			Handler:    _Middleware_DeleteFeeOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/fee/fee.proto",
}
