// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/order/mgr/v1/payment/payment.proto

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
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

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreatePayments(ctx context.Context, in *CreatePaymentsRequest, opts ...grpc.CallOption) (*CreatePaymentsResponse, error) {
	out := new(CreatePaymentsResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/CreatePayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*UpdatePaymentResponse, error) {
	out := new(UpdatePaymentResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/UpdatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/GetPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPaymentOnly(ctx context.Context, in *GetPaymentOnlyRequest, opts ...grpc.CallOption) (*GetPaymentOnlyResponse, error) {
	out := new(GetPaymentOnlyResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/GetPaymentOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPayments(ctx context.Context, in *GetPaymentsRequest, opts ...grpc.CallOption) (*GetPaymentsResponse, error) {
	out := new(GetPaymentsResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/GetPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistPayment(ctx context.Context, in *ExistPaymentRequest, opts ...grpc.CallOption) (*ExistPaymentResponse, error) {
	out := new(ExistPaymentResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/ExistPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistPaymentConds(ctx context.Context, in *ExistPaymentCondsRequest, opts ...grpc.CallOption) (*ExistPaymentCondsResponse, error) {
	out := new(ExistPaymentCondsResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/ExistPaymentConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountPayments(ctx context.Context, in *CountPaymentsRequest, opts ...grpc.CallOption) (*CountPaymentsResponse, error) {
	out := new(CountPaymentsResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/CountPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeletePayment(ctx context.Context, in *DeletePaymentRequest, opts ...grpc.CallOption) (*DeletePaymentResponse, error) {
	out := new(DeletePaymentResponse)
	err := c.cc.Invoke(ctx, "/order.manager.payment.v1.Manager/DeletePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
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
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedManagerServer) CreatePayments(context.Context, *CreatePaymentsRequest) (*CreatePaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayments not implemented")
}
func (UnimplementedManagerServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedManagerServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedManagerServer) GetPaymentOnly(context.Context, *GetPaymentOnlyRequest) (*GetPaymentOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentOnly not implemented")
}
func (UnimplementedManagerServer) GetPayments(context.Context, *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayments not implemented")
}
func (UnimplementedManagerServer) ExistPayment(context.Context, *ExistPaymentRequest) (*ExistPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPayment not implemented")
}
func (UnimplementedManagerServer) ExistPaymentConds(context.Context, *ExistPaymentCondsRequest) (*ExistPaymentCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPaymentConds not implemented")
}
func (UnimplementedManagerServer) CountPayments(context.Context, *CountPaymentsRequest) (*CountPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPayments not implemented")
}
func (UnimplementedManagerServer) DeletePayment(context.Context, *DeletePaymentRequest) (*DeletePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayment not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreatePayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreatePayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/CreatePayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreatePayments(ctx, req.(*CreatePaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/UpdatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/GetPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPaymentOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPaymentOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/GetPaymentOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPaymentOnly(ctx, req.(*GetPaymentOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/GetPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPayments(ctx, req.(*GetPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/ExistPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistPayment(ctx, req.(*ExistPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistPaymentConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPaymentCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistPaymentConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/ExistPaymentConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistPaymentConds(ctx, req.(*ExistPaymentCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/CountPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountPayments(ctx, req.(*CountPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.payment.v1.Manager/DeletePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeletePayment(ctx, req.(*DeletePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.manager.payment.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _Manager_CreatePayment_Handler,
		},
		{
			MethodName: "CreatePayments",
			Handler:    _Manager_CreatePayments_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _Manager_UpdatePayment_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _Manager_GetPayment_Handler,
		},
		{
			MethodName: "GetPaymentOnly",
			Handler:    _Manager_GetPaymentOnly_Handler,
		},
		{
			MethodName: "GetPayments",
			Handler:    _Manager_GetPayments_Handler,
		},
		{
			MethodName: "ExistPayment",
			Handler:    _Manager_ExistPayment_Handler,
		},
		{
			MethodName: "ExistPaymentConds",
			Handler:    _Manager_ExistPaymentConds_Handler,
		},
		{
			MethodName: "CountPayments",
			Handler:    _Manager_CountPayments_Handler,
		},
		{
			MethodName: "DeletePayment",
			Handler:    _Manager_DeletePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mgr/v1/payment/payment.proto",
}
