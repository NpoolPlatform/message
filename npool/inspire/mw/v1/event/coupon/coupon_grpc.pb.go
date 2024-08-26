// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/event/coupon/coupon.proto

package coupon

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
	Middleware_CreateEventCoupon_FullMethodName     = "/inspire.middleware.event.coupon.v1.Middleware/CreateEventCoupon"
	Middleware_GetEventCoupon_FullMethodName        = "/inspire.middleware.event.coupon.v1.Middleware/GetEventCoupon"
	Middleware_GetEventCoupons_FullMethodName       = "/inspire.middleware.event.coupon.v1.Middleware/GetEventCoupons"
	Middleware_ExistEventCouponConds_FullMethodName = "/inspire.middleware.event.coupon.v1.Middleware/ExistEventCouponConds"
	Middleware_DeleteEventCoupon_FullMethodName     = "/inspire.middleware.event.coupon.v1.Middleware/DeleteEventCoupon"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateEventCoupon(ctx context.Context, in *CreateEventCouponRequest, opts ...grpc.CallOption) (*CreateEventCouponResponse, error)
	GetEventCoupon(ctx context.Context, in *GetEventCouponRequest, opts ...grpc.CallOption) (*GetEventCouponResponse, error)
	GetEventCoupons(ctx context.Context, in *GetEventCouponsRequest, opts ...grpc.CallOption) (*GetEventCouponsResponse, error)
	ExistEventCouponConds(ctx context.Context, in *ExistEventCouponCondsRequest, opts ...grpc.CallOption) (*ExistEventCouponCondsResponse, error)
	DeleteEventCoupon(ctx context.Context, in *DeleteEventCouponRequest, opts ...grpc.CallOption) (*DeleteEventCouponResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateEventCoupon(ctx context.Context, in *CreateEventCouponRequest, opts ...grpc.CallOption) (*CreateEventCouponResponse, error) {
	out := new(CreateEventCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetEventCoupon(ctx context.Context, in *GetEventCouponRequest, opts ...grpc.CallOption) (*GetEventCouponResponse, error) {
	out := new(GetEventCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_GetEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetEventCoupons(ctx context.Context, in *GetEventCouponsRequest, opts ...grpc.CallOption) (*GetEventCouponsResponse, error) {
	out := new(GetEventCouponsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetEventCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistEventCouponConds(ctx context.Context, in *ExistEventCouponCondsRequest, opts ...grpc.CallOption) (*ExistEventCouponCondsResponse, error) {
	out := new(ExistEventCouponCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistEventCouponConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteEventCoupon(ctx context.Context, in *DeleteEventCouponRequest, opts ...grpc.CallOption) (*DeleteEventCouponResponse, error) {
	out := new(DeleteEventCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateEventCoupon(context.Context, *CreateEventCouponRequest) (*CreateEventCouponResponse, error)
	GetEventCoupon(context.Context, *GetEventCouponRequest) (*GetEventCouponResponse, error)
	GetEventCoupons(context.Context, *GetEventCouponsRequest) (*GetEventCouponsResponse, error)
	ExistEventCouponConds(context.Context, *ExistEventCouponCondsRequest) (*ExistEventCouponCondsResponse, error)
	DeleteEventCoupon(context.Context, *DeleteEventCouponRequest) (*DeleteEventCouponResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateEventCoupon(context.Context, *CreateEventCouponRequest) (*CreateEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventCoupon not implemented")
}
func (UnimplementedMiddlewareServer) GetEventCoupon(context.Context, *GetEventCouponRequest) (*GetEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCoupon not implemented")
}
func (UnimplementedMiddlewareServer) GetEventCoupons(context.Context, *GetEventCouponsRequest) (*GetEventCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCoupons not implemented")
}
func (UnimplementedMiddlewareServer) ExistEventCouponConds(context.Context, *ExistEventCouponCondsRequest) (*ExistEventCouponCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistEventCouponConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteEventCoupon(context.Context, *DeleteEventCouponRequest) (*DeleteEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEventCoupon not implemented")
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

func _Middleware_CreateEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateEventCoupon(ctx, req.(*CreateEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetEventCoupon(ctx, req.(*GetEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetEventCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetEventCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetEventCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetEventCoupons(ctx, req.(*GetEventCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistEventCouponConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistEventCouponCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistEventCouponConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistEventCouponConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistEventCouponConds(ctx, req.(*ExistEventCouponCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteEventCoupon(ctx, req.(*DeleteEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.event.coupon.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventCoupon",
			Handler:    _Middleware_CreateEventCoupon_Handler,
		},
		{
			MethodName: "GetEventCoupon",
			Handler:    _Middleware_GetEventCoupon_Handler,
		},
		{
			MethodName: "GetEventCoupons",
			Handler:    _Middleware_GetEventCoupons_Handler,
		},
		{
			MethodName: "ExistEventCouponConds",
			Handler:    _Middleware_ExistEventCouponConds_Handler,
		},
		{
			MethodName: "DeleteEventCoupon",
			Handler:    _Middleware_DeleteEventCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/event/coupon/coupon.proto",
}
