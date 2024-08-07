// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/order/coupon/coupon.proto

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
	Middleware_GetOrderCoupons_FullMethodName       = "/order.middleware.order1.coupon.v1.Middleware/GetOrderCoupons"
	Middleware_ExistOrderCouponConds_FullMethodName = "/order.middleware.order1.coupon.v1.Middleware/ExistOrderCouponConds"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetOrderCoupons(ctx context.Context, in *GetOrderCouponsRequest, opts ...grpc.CallOption) (*GetOrderCouponsResponse, error)
	ExistOrderCouponConds(ctx context.Context, in *ExistOrderCouponCondsRequest, opts ...grpc.CallOption) (*ExistOrderCouponCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetOrderCoupons(ctx context.Context, in *GetOrderCouponsRequest, opts ...grpc.CallOption) (*GetOrderCouponsResponse, error) {
	out := new(GetOrderCouponsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOrderCouponConds(ctx context.Context, in *ExistOrderCouponCondsRequest, opts ...grpc.CallOption) (*ExistOrderCouponCondsResponse, error) {
	out := new(ExistOrderCouponCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistOrderCouponConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetOrderCoupons(context.Context, *GetOrderCouponsRequest) (*GetOrderCouponsResponse, error)
	ExistOrderCouponConds(context.Context, *ExistOrderCouponCondsRequest) (*ExistOrderCouponCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetOrderCoupons(context.Context, *GetOrderCouponsRequest) (*GetOrderCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderCoupons not implemented")
}
func (UnimplementedMiddlewareServer) ExistOrderCouponConds(context.Context, *ExistOrderCouponCondsRequest) (*ExistOrderCouponCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOrderCouponConds not implemented")
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

func _Middleware_GetOrderCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderCoupons(ctx, req.(*GetOrderCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOrderCouponConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOrderCouponCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOrderCouponConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistOrderCouponConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOrderCouponConds(ctx, req.(*ExistOrderCouponCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.order1.coupon.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderCoupons",
			Handler:    _Middleware_GetOrderCoupons_Handler,
		},
		{
			MethodName: "ExistOrderCouponConds",
			Handler:    _Middleware_ExistOrderCouponConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/order/coupon/coupon.proto",
}
