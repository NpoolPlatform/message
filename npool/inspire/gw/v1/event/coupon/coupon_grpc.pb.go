// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/event/coupon/coupon.proto

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
	Gateway_CreateEventCoupon_FullMethodName      = "/inspire.gateway.event.coupon.v1.Gateway/CreateEventCoupon"
	Gateway_GetEventCoupons_FullMethodName        = "/inspire.gateway.event.coupon.v1.Gateway/GetEventCoupons"
	Gateway_AdminGetEventCoupons_FullMethodName   = "/inspire.gateway.event.coupon.v1.Gateway/AdminGetEventCoupons"
	Gateway_AdminCreateEventCoupon_FullMethodName = "/inspire.gateway.event.coupon.v1.Gateway/AdminCreateEventCoupon"
	Gateway_AdminDeleteEventCoupon_FullMethodName = "/inspire.gateway.event.coupon.v1.Gateway/AdminDeleteEventCoupon"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateEventCoupon(ctx context.Context, in *CreateEventCouponRequest, opts ...grpc.CallOption) (*CreateEventCouponResponse, error)
	GetEventCoupons(ctx context.Context, in *GetEventCouponsRequest, opts ...grpc.CallOption) (*GetEventCouponsResponse, error)
	AdminGetEventCoupons(ctx context.Context, in *AdminGetEventCouponsRequest, opts ...grpc.CallOption) (*AdminGetEventCouponsResponse, error)
	AdminCreateEventCoupon(ctx context.Context, in *AdminCreateEventCouponRequest, opts ...grpc.CallOption) (*AdminCreateEventCouponResponse, error)
	AdminDeleteEventCoupon(ctx context.Context, in *AdminDeleteEventCouponRequest, opts ...grpc.CallOption) (*AdminDeleteEventCouponResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateEventCoupon(ctx context.Context, in *CreateEventCouponRequest, opts ...grpc.CallOption) (*CreateEventCouponResponse, error) {
	out := new(CreateEventCouponResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetEventCoupons(ctx context.Context, in *GetEventCouponsRequest, opts ...grpc.CallOption) (*GetEventCouponsResponse, error) {
	out := new(GetEventCouponsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetEventCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetEventCoupons(ctx context.Context, in *AdminGetEventCouponsRequest, opts ...grpc.CallOption) (*AdminGetEventCouponsResponse, error) {
	out := new(AdminGetEventCouponsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetEventCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateEventCoupon(ctx context.Context, in *AdminCreateEventCouponRequest, opts ...grpc.CallOption) (*AdminCreateEventCouponResponse, error) {
	out := new(AdminCreateEventCouponResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteEventCoupon(ctx context.Context, in *AdminDeleteEventCouponRequest, opts ...grpc.CallOption) (*AdminDeleteEventCouponResponse, error) {
	out := new(AdminDeleteEventCouponResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteEventCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateEventCoupon(context.Context, *CreateEventCouponRequest) (*CreateEventCouponResponse, error)
	GetEventCoupons(context.Context, *GetEventCouponsRequest) (*GetEventCouponsResponse, error)
	AdminGetEventCoupons(context.Context, *AdminGetEventCouponsRequest) (*AdminGetEventCouponsResponse, error)
	AdminCreateEventCoupon(context.Context, *AdminCreateEventCouponRequest) (*AdminCreateEventCouponResponse, error)
	AdminDeleteEventCoupon(context.Context, *AdminDeleteEventCouponRequest) (*AdminDeleteEventCouponResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateEventCoupon(context.Context, *CreateEventCouponRequest) (*CreateEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEventCoupon not implemented")
}
func (UnimplementedGatewayServer) GetEventCoupons(context.Context, *GetEventCouponsRequest) (*GetEventCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCoupons not implemented")
}
func (UnimplementedGatewayServer) AdminGetEventCoupons(context.Context, *AdminGetEventCouponsRequest) (*AdminGetEventCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetEventCoupons not implemented")
}
func (UnimplementedGatewayServer) AdminCreateEventCoupon(context.Context, *AdminCreateEventCouponRequest) (*AdminCreateEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateEventCoupon not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteEventCoupon(context.Context, *AdminDeleteEventCouponRequest) (*AdminDeleteEventCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteEventCoupon not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateEventCoupon(ctx, req.(*CreateEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetEventCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetEventCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetEventCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetEventCoupons(ctx, req.(*GetEventCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetEventCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetEventCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetEventCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetEventCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetEventCoupons(ctx, req.(*AdminGetEventCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateEventCoupon(ctx, req.(*AdminCreateEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteEventCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteEventCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteEventCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteEventCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteEventCoupon(ctx, req.(*AdminDeleteEventCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.event.coupon.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEventCoupon",
			Handler:    _Gateway_CreateEventCoupon_Handler,
		},
		{
			MethodName: "GetEventCoupons",
			Handler:    _Gateway_GetEventCoupons_Handler,
		},
		{
			MethodName: "AdminGetEventCoupons",
			Handler:    _Gateway_AdminGetEventCoupons_Handler,
		},
		{
			MethodName: "AdminCreateEventCoupon",
			Handler:    _Gateway_AdminCreateEventCoupon_Handler,
		},
		{
			MethodName: "AdminDeleteEventCoupon",
			Handler:    _Gateway_AdminDeleteEventCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/event/coupon/coupon.proto",
}