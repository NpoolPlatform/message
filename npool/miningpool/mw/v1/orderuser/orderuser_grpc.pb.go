// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/orderuser/orderuser.proto

package orderuser

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
	Middleware_CreateOrderUser_FullMethodName     = "/miningpool.middleware.orderuser.v1.Middleware/CreateOrderUser"
	Middleware_CreateOrderUsers_FullMethodName    = "/miningpool.middleware.orderuser.v1.Middleware/CreateOrderUsers"
	Middleware_GetOrderUser_FullMethodName        = "/miningpool.middleware.orderuser.v1.Middleware/GetOrderUser"
	Middleware_GetOrderUsers_FullMethodName       = "/miningpool.middleware.orderuser.v1.Middleware/GetOrderUsers"
	Middleware_ExistOrderUserConds_FullMethodName = "/miningpool.middleware.orderuser.v1.Middleware/ExistOrderUserConds"
	Middleware_UpdateOrderUser_FullMethodName     = "/miningpool.middleware.orderuser.v1.Middleware/UpdateOrderUser"
	Middleware_DeleteOrderUser_FullMethodName     = "/miningpool.middleware.orderuser.v1.Middleware/DeleteOrderUser"
	Middleware_SetupProportion_FullMethodName     = "/miningpool.middleware.orderuser.v1.Middleware/SetupProportion"
	Middleware_SetupRevenueAddress_FullMethodName = "/miningpool.middleware.orderuser.v1.Middleware/SetupRevenueAddress"
	Middleware_SetupAutoPay_FullMethodName        = "/miningpool.middleware.orderuser.v1.Middleware/SetupAutoPay"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateOrderUser(ctx context.Context, in *CreateOrderUserRequest, opts ...grpc.CallOption) (*CreateOrderUserResponse, error)
	CreateOrderUsers(ctx context.Context, in *CreateOrderUsersRequest, opts ...grpc.CallOption) (*CreateOrderUsersResponse, error)
	GetOrderUser(ctx context.Context, in *GetOrderUserRequest, opts ...grpc.CallOption) (*GetOrderUserResponse, error)
	GetOrderUsers(ctx context.Context, in *GetOrderUsersRequest, opts ...grpc.CallOption) (*GetOrderUsersResponse, error)
	ExistOrderUserConds(ctx context.Context, in *ExistOrderUserCondsRequest, opts ...grpc.CallOption) (*ExistOrderUserCondsResponse, error)
	UpdateOrderUser(ctx context.Context, in *UpdateOrderUserRequest, opts ...grpc.CallOption) (*UpdateOrderUserResponse, error)
	DeleteOrderUser(ctx context.Context, in *DeleteOrderUserRequest, opts ...grpc.CallOption) (*DeleteOrderUserResponse, error)
	SetupProportion(ctx context.Context, in *SetupProportionRequest, opts ...grpc.CallOption) (*SetupProportionResponse, error)
	SetupRevenueAddress(ctx context.Context, in *SetupRevenueAddressRequest, opts ...grpc.CallOption) (*SetupRevenueAddressResponse, error)
	SetupAutoPay(ctx context.Context, in *SetupAutoPayRequest, opts ...grpc.CallOption) (*SetupAutoPayResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateOrderUser(ctx context.Context, in *CreateOrderUserRequest, opts ...grpc.CallOption) (*CreateOrderUserResponse, error) {
	out := new(CreateOrderUserResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOrderUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateOrderUsers(ctx context.Context, in *CreateOrderUsersRequest, opts ...grpc.CallOption) (*CreateOrderUsersResponse, error) {
	out := new(CreateOrderUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOrderUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOrderUser(ctx context.Context, in *GetOrderUserRequest, opts ...grpc.CallOption) (*GetOrderUserResponse, error) {
	out := new(GetOrderUserResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOrderUsers(ctx context.Context, in *GetOrderUsersRequest, opts ...grpc.CallOption) (*GetOrderUsersResponse, error) {
	out := new(GetOrderUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOrderUserConds(ctx context.Context, in *ExistOrderUserCondsRequest, opts ...grpc.CallOption) (*ExistOrderUserCondsResponse, error) {
	out := new(ExistOrderUserCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOrderUserConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateOrderUser(ctx context.Context, in *UpdateOrderUserRequest, opts ...grpc.CallOption) (*UpdateOrderUserResponse, error) {
	out := new(UpdateOrderUserResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateOrderUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteOrderUser(ctx context.Context, in *DeleteOrderUserRequest, opts ...grpc.CallOption) (*DeleteOrderUserResponse, error) {
	out := new(DeleteOrderUserResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteOrderUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SetupProportion(ctx context.Context, in *SetupProportionRequest, opts ...grpc.CallOption) (*SetupProportionResponse, error) {
	out := new(SetupProportionResponse)
	err := c.cc.Invoke(ctx, Middleware_SetupProportion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SetupRevenueAddress(ctx context.Context, in *SetupRevenueAddressRequest, opts ...grpc.CallOption) (*SetupRevenueAddressResponse, error) {
	out := new(SetupRevenueAddressResponse)
	err := c.cc.Invoke(ctx, Middleware_SetupRevenueAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) SetupAutoPay(ctx context.Context, in *SetupAutoPayRequest, opts ...grpc.CallOption) (*SetupAutoPayResponse, error) {
	out := new(SetupAutoPayResponse)
	err := c.cc.Invoke(ctx, Middleware_SetupAutoPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateOrderUser(context.Context, *CreateOrderUserRequest) (*CreateOrderUserResponse, error)
	CreateOrderUsers(context.Context, *CreateOrderUsersRequest) (*CreateOrderUsersResponse, error)
	GetOrderUser(context.Context, *GetOrderUserRequest) (*GetOrderUserResponse, error)
	GetOrderUsers(context.Context, *GetOrderUsersRequest) (*GetOrderUsersResponse, error)
	ExistOrderUserConds(context.Context, *ExistOrderUserCondsRequest) (*ExistOrderUserCondsResponse, error)
	UpdateOrderUser(context.Context, *UpdateOrderUserRequest) (*UpdateOrderUserResponse, error)
	DeleteOrderUser(context.Context, *DeleteOrderUserRequest) (*DeleteOrderUserResponse, error)
	SetupProportion(context.Context, *SetupProportionRequest) (*SetupProportionResponse, error)
	SetupRevenueAddress(context.Context, *SetupRevenueAddressRequest) (*SetupRevenueAddressResponse, error)
	SetupAutoPay(context.Context, *SetupAutoPayRequest) (*SetupAutoPayResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateOrderUser(context.Context, *CreateOrderUserRequest) (*CreateOrderUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderUser not implemented")
}
func (UnimplementedMiddlewareServer) CreateOrderUsers(context.Context, *CreateOrderUsersRequest) (*CreateOrderUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetOrderUser(context.Context, *GetOrderUserRequest) (*GetOrderUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderUser not implemented")
}
func (UnimplementedMiddlewareServer) GetOrderUsers(context.Context, *GetOrderUsersRequest) (*GetOrderUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderUsers not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrderUserConds(context.Context, *ExistOrderUserCondsRequest) (*ExistOrderUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderUserConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateOrderUser(context.Context, *UpdateOrderUserRequest) (*UpdateOrderUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderUser not implemented")
}
func (UnimplementedMiddlewareServer) DeleteOrderUser(context.Context, *DeleteOrderUserRequest) (*DeleteOrderUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderUser not implemented")
}
func (UnimplementedMiddlewareServer) SetupProportion(context.Context, *SetupProportionRequest) (*SetupProportionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupProportion not implemented")
}
func (UnimplementedMiddlewareServer) SetupRevenueAddress(context.Context, *SetupRevenueAddressRequest) (*SetupRevenueAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupRevenueAddress not implemented")
}
func (UnimplementedMiddlewareServer) SetupAutoPay(context.Context, *SetupAutoPayRequest) (*SetupAutoPayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupAutoPay not implemented")
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

func _Middleware_CreateOrderUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOrderUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOrderUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOrderUser(ctx, req.(*CreateOrderUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateOrderUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOrderUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOrderUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOrderUsers(ctx, req.(*CreateOrderUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOrderUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderUser(ctx, req.(*GetOrderUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOrderUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderUsers(ctx, req.(*GetOrderUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOrderUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOrderUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOrderUserConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOrderUserConds(ctx, req.(*ExistOrderUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateOrderUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateOrderUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateOrderUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateOrderUser(ctx, req.(*UpdateOrderUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteOrderUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteOrderUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteOrderUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteOrderUser(ctx, req.(*DeleteOrderUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SetupProportion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupProportionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SetupProportion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SetupProportion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SetupProportion(ctx, req.(*SetupProportionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SetupRevenueAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRevenueAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SetupRevenueAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SetupRevenueAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SetupRevenueAddress(ctx, req.(*SetupRevenueAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_SetupAutoPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupAutoPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).SetupAutoPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_SetupAutoPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).SetupAutoPay(ctx, req.(*SetupAutoPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.orderuser.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderUser",
			Handler:    _Middleware_CreateOrderUser_Handler,
		},
		{
			MethodName: "CreateOrderUsers",
			Handler:    _Middleware_CreateOrderUsers_Handler,
		},
		{
			MethodName: "GetOrderUser",
			Handler:    _Middleware_GetOrderUser_Handler,
		},
		{
			MethodName: "GetOrderUsers",
			Handler:    _Middleware_GetOrderUsers_Handler,
		},
		{
			MethodName: "ExistOrderUserConds",
			Handler:    _Middleware_ExistOrderUserConds_Handler,
		},
		{
			MethodName: "UpdateOrderUser",
			Handler:    _Middleware_UpdateOrderUser_Handler,
		},
		{
			MethodName: "DeleteOrderUser",
			Handler:    _Middleware_DeleteOrderUser_Handler,
		},
		{
			MethodName: "SetupProportion",
			Handler:    _Middleware_SetupProportion_Handler,
		},
		{
			MethodName: "SetupRevenueAddress",
			Handler:    _Middleware_SetupRevenueAddress_Handler,
		},
		{
			MethodName: "SetupAutoPay",
			Handler:    _Middleware_SetupAutoPay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/orderuser/orderuser.proto",
}
