// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/coupon/coupon.proto

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
	Middleware_CreateCoupon_FullMethodName = "/inspire.middleware.coupon.v1.Middleware/CreateCoupon"
	Middleware_UpdateCoupon_FullMethodName = "/inspire.middleware.coupon.v1.Middleware/UpdateCoupon"
	Middleware_GetCoupon_FullMethodName    = "/inspire.middleware.coupon.v1.Middleware/GetCoupon"
	Middleware_GetCoupons_FullMethodName   = "/inspire.middleware.coupon.v1.Middleware/GetCoupons"
<<<<<<< HEAD
<<<<<<< HEAD:npool/inspire/mw/v1/coupon/coupon_grpc.pb.go
	Middleware_DeleteCoupon_FullMethodName = "/inspire.middleware.coupon.v1.Middleware/DeleteCoupon"
=======
>>>>>>> Refactor coupon structure:npool/inspire/mw/v1/coupon/coupon/coupon_grpc.pb.go
=======
	Middleware_DeleteCoupon_FullMethodName = "/inspire.middleware.coupon.v1.Middleware/DeleteCoupon"
>>>>>>> Add delete coupon api
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...grpc.CallOption) (*CreateCouponResponse, error)
	UpdateCoupon(ctx context.Context, in *UpdateCouponRequest, opts ...grpc.CallOption) (*UpdateCouponResponse, error)
	GetCoupon(ctx context.Context, in *GetCouponRequest, opts ...grpc.CallOption) (*GetCouponResponse, error)
	GetCoupons(ctx context.Context, in *GetCouponsRequest, opts ...grpc.CallOption) (*GetCouponsResponse, error)
	DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...grpc.CallOption) (*CreateCouponResponse, error) {
	out := new(CreateCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCoupon(ctx context.Context, in *UpdateCouponRequest, opts ...grpc.CallOption) (*UpdateCouponResponse, error) {
	out := new(UpdateCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoupon(ctx context.Context, in *GetCouponRequest, opts ...grpc.CallOption) (*GetCouponResponse, error) {
	out := new(GetCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoupons(ctx context.Context, in *GetCouponsRequest, opts ...grpc.CallOption) (*GetCouponsResponse, error) {
	out := new(GetCouponsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoupons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error) {
	out := new(DeleteCouponResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCoupon_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCoupon(context.Context, *CreateCouponRequest) (*CreateCouponResponse, error)
	UpdateCoupon(context.Context, *UpdateCouponRequest) (*UpdateCouponResponse, error)
	GetCoupon(context.Context, *GetCouponRequest) (*GetCouponResponse, error)
	GetCoupons(context.Context, *GetCouponsRequest) (*GetCouponsResponse, error)
	DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCoupon(context.Context, *CreateCouponRequest) (*CreateCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoupon not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCoupon(context.Context, *UpdateCouponRequest) (*UpdateCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoupon not implemented")
}
func (UnimplementedMiddlewareServer) GetCoupon(context.Context, *GetCouponRequest) (*GetCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoupon not implemented")
}
func (UnimplementedMiddlewareServer) GetCoupons(context.Context, *GetCouponsRequest) (*GetCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoupons not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoupon not implemented")
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

func _Middleware_CreateCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCoupon(ctx, req.(*CreateCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCoupon(ctx, req.(*UpdateCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoupon(ctx, req.(*GetCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoupons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoupons(ctx, req.(*GetCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCoupon(ctx, req.(*DeleteCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.coupon.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoupon",
			Handler:    _Middleware_CreateCoupon_Handler,
		},
		{
			MethodName: "UpdateCoupon",
			Handler:    _Middleware_UpdateCoupon_Handler,
		},
		{
			MethodName: "GetCoupon",
			Handler:    _Middleware_GetCoupon_Handler,
		},
		{
			MethodName: "GetCoupons",
			Handler:    _Middleware_GetCoupons_Handler,
		},
		{
			MethodName: "DeleteCoupon",
			Handler:    _Middleware_DeleteCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coupon/coupon.proto",
}
