// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/compensate/compensate.proto

package compensate

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
	Middleware_GetCompensate_FullMethodName         = "/order.middleware.compensate.v1.Middleware/GetCompensate"
	Middleware_GetCompensates_FullMethodName        = "/order.middleware.compensate.v1.Middleware/GetCompensates"
	Middleware_ExistCompensate_FullMethodName       = "/order.middleware.compensate.v1.Middleware/ExistCompensate"
	Middleware_ExistCompensateConds_FullMethodName  = "/order.middleware.compensate.v1.Middleware/ExistCompensateConds"
	Middleware_CountCompensates_FullMethodName      = "/order.middleware.compensate.v1.Middleware/CountCompensates"
	Middleware_CountCompensateOrders_FullMethodName = "/order.middleware.compensate.v1.Middleware/CountCompensateOrders"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error)
	GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error)
	ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error)
	ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error)
	CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error)
	CountCompensateOrders(ctx context.Context, in *CountCompensateOrdersRequest, opts ...grpc.CallOption) (*CountCompensateOrdersResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error) {
	out := new(GetCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error) {
	out := new(GetCompensatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error) {
	out := new(ExistCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error) {
	out := new(ExistCompensateCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCompensateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error) {
	out := new(CountCompensatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CountCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountCompensateOrders(ctx context.Context, in *CountCompensateOrdersRequest, opts ...grpc.CallOption) (*CountCompensateOrdersResponse, error) {
	out := new(CountCompensateOrdersResponse)
	err := c.cc.Invoke(ctx, Middleware_CountCompensateOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error)
	GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error)
	ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error)
	ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error)
	CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error)
	CountCompensateOrders(context.Context, *CountCompensateOrdersRequest) (*CountCompensateOrdersResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensate not implemented")
}
func (UnimplementedMiddlewareServer) GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensates not implemented")
}
func (UnimplementedMiddlewareServer) ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensate not implemented")
}
func (UnimplementedMiddlewareServer) ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensateConds not implemented")
}
func (UnimplementedMiddlewareServer) CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCompensates not implemented")
}
func (UnimplementedMiddlewareServer) CountCompensateOrders(context.Context, *CountCompensateOrdersRequest) (*CountCompensateOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCompensateOrders not implemented")
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

func _Middleware_GetCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCompensate(ctx, req.(*GetCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCompensates(ctx, req.(*GetCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCompensate(ctx, req.(*ExistCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCompensateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCompensateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCompensateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCompensateConds(ctx, req.(*ExistCompensateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountCompensates(ctx, req.(*CountCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountCompensateOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCompensateOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountCompensateOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountCompensateOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountCompensateOrders(ctx, req.(*CountCompensateOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.compensate.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompensate",
			Handler:    _Middleware_GetCompensate_Handler,
		},
		{
			MethodName: "GetCompensates",
			Handler:    _Middleware_GetCompensates_Handler,
		},
		{
			MethodName: "ExistCompensate",
			Handler:    _Middleware_ExistCompensate_Handler,
		},
		{
			MethodName: "ExistCompensateConds",
			Handler:    _Middleware_ExistCompensateConds_Handler,
		},
		{
			MethodName: "CountCompensates",
			Handler:    _Middleware_CountCompensates_Handler,
		},
		{
			MethodName: "CountCompensateOrders",
			Handler:    _Middleware_CountCompensateOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/compensate/compensate.proto",
}
