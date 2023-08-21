// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/payment/payment.proto

package payment

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
	Middleware_CreatePayment_FullMethodName     = "/order.middleware.payment.v1.Middleware/CreatePayment"
	Middleware_CreatePayments_FullMethodName    = "/order.middleware.payment.v1.Middleware/CreatePayments"
	Middleware_UpdatePayment_FullMethodName     = "/order.middleware.payment.v1.Middleware/UpdatePayment"
	Middleware_GetPayment_FullMethodName        = "/order.middleware.payment.v1.Middleware/GetPayment"
	Middleware_GetPaymentOnly_FullMethodName    = "/order.middleware.payment.v1.Middleware/GetPaymentOnly"
	Middleware_GetPayments_FullMethodName       = "/order.middleware.payment.v1.Middleware/GetPayments"
	Middleware_ExistPayment_FullMethodName      = "/order.middleware.payment.v1.Middleware/ExistPayment"
	Middleware_ExistPaymentConds_FullMethodName = "/order.middleware.payment.v1.Middleware/ExistPaymentConds"
	Middleware_CountPayments_FullMethodName     = "/order.middleware.payment.v1.Middleware/CountPayments"
	Middleware_DeletePayment_FullMethodName     = "/order.middleware.payment.v1.Middleware/DeletePayment"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	CreatePayments(ctx context.Context, in *CreatePaymentsRequest, opts ...grpc.CallOption) (*CreatePaymentsResponse, error)
	UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
	GetPaymentOnly(ctx context.Context, in *GetPaymentOnlyRequest, opts ...grpc.CallOption) (*GetPaymentOnlyResponse, error)
	GetPayments(ctx context.Context, in *GetPaymentsRequest, opts ...grpc.CallOption) (*GetPaymentsResponse, error)
	ExistPayment(ctx context.Context, in *ExistPaymentRequest, opts ...grpc.CallOption) (*ExistPaymentResponse, error)
	ExistPaymentConds(ctx context.Context, in *ExistPaymentCondsRequest, opts ...grpc.CallOption) (*ExistPaymentCondsResponse, error)
	CountPayments(ctx context.Context, in *CountPaymentsRequest, opts ...grpc.CallOption) (*CountPaymentsResponse, error)
	DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreatePayments(ctx context.Context, in *CreatePaymentsRequest, opts ...grpc.CallOption) (*CreatePaymentsResponse, error) {
	out := new(CreatePaymentsResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePayments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error) {
	out := new(UpdatePaymentResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPaymentOnly(ctx context.Context, in *GetPaymentOnlyRequest, opts ...grpc.CallOption) (*GetPaymentOnlyResponse, error) {
	out := new(GetPaymentOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPaymentOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPayments(ctx context.Context, in *GetPaymentsRequest, opts ...grpc.CallOption) (*GetPaymentsResponse, error) {
	out := new(GetPaymentsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPayments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPayment(ctx context.Context, in *ExistPaymentRequest, opts ...grpc.CallOption) (*ExistPaymentResponse, error) {
	out := new(ExistPaymentResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPaymentConds(ctx context.Context, in *ExistPaymentCondsRequest, opts ...grpc.CallOption) (*ExistPaymentCondsResponse, error) {
	out := new(ExistPaymentCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPaymentConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountPayments(ctx context.Context, in *CountPaymentsRequest, opts ...grpc.CallOption) (*CountPaymentsResponse, error) {
	out := new(CountPaymentsResponse)
	err := c.cc.Invoke(ctx, Middleware_CountPayments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error) {
	out := new(DeletePaymentResponse)
	err := c.cc.Invoke(ctx, Middleware_DeletePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	CreatePayments(context.Context, *CreatePaymentsRequest) (*CreatePaymentsResponse, error)
	UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	GetPaymentOnly(context.Context, *GetPaymentOnlyRequest) (*GetPaymentOnlyResponse, error)
	GetPayments(context.Context, *GetPaymentsRequest) (*GetPaymentsResponse, error)
	ExistPayment(context.Context, *ExistPaymentRequest) (*ExistPaymentResponse, error)
	ExistPaymentConds(context.Context, *ExistPaymentCondsRequest) (*ExistPaymentCondsResponse, error)
	CountPayments(context.Context, *CountPaymentsRequest) (*CountPaymentsResponse, error)
	DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedMiddlewareServer) CreatePayments(context.Context, *CreatePaymentsRequest) (*CreatePaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayments not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedMiddlewareServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedMiddlewareServer) GetPaymentOnly(context.Context, *GetPaymentOnlyRequest) (*GetPaymentOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetPayments(context.Context, *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayments not implemented")
}
func (UnimplementedMiddlewareServer) ExistPayment(context.Context, *ExistPaymentRequest) (*ExistPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPayment not implemented")
}
func (UnimplementedMiddlewareServer) ExistPaymentConds(context.Context, *ExistPaymentCondsRequest) (*ExistPaymentCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPaymentConds not implemented")
}
func (UnimplementedMiddlewareServer) CountPayments(context.Context, *CountPaymentsRequest) (*CountPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPayments not implemented")
}
func (UnimplementedMiddlewareServer) DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
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

func _Middleware_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreatePayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePayments(ctx, req.(*CreatePaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPaymentOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPaymentOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPaymentOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPaymentOnly(ctx, req.(*GetPaymentOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPayments(ctx, req.(*GetPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPayment(ctx, req.(*ExistPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPaymentConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPaymentCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPaymentConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPaymentConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPaymentConds(ctx, req.(*ExistPaymentCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountPayments(ctx, req.(*CountPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeletePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeletePayment(ctx, req.(*DeletePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.payment.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _Middleware_CreatePayment_Handler,
		},
		{
			MethodName: "CreatePayments",
			Handler:    _Middleware_CreatePayments_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _Middleware_UpdatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _Middleware_GetPayment_Handler,
		},
		{
			MethodName: "GetPaymentOnly",
			Handler:    _Middleware_GetPaymentOnly_Handler,
		},
		{
			MethodName: "GetPayments",
			Handler:    _Middleware_GetPayments_Handler,
		},
		{
			MethodName: "ExistPayment",
			Handler:    _Middleware_ExistPayment_Handler,
		},
		{
			MethodName: "ExistPaymentConds",
			Handler:    _Middleware_ExistPaymentConds_Handler,
		},
		{
			MethodName: "CountPayments",
			Handler:    _Middleware_CountPayments_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _Middleware_DeletePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/payment/payment.proto",
}
