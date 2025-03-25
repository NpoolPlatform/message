// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/review/gw/v2/withdraw/coupon/coupon.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetCouponWithdrawReviews(ctx context.Context, in *GetCouponWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetCouponWithdrawReviewsResponse, error)
	GetAppCouponWithdrawReviews(ctx context.Context, in *GetAppCouponWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetAppCouponWithdrawReviewsResponse, error)
	UpdateCouponWithdrawReview(ctx context.Context, in *UpdateCouponWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateCouponWithdrawReviewResponse, error)
	UpdateAppCouponWithdrawReview(ctx context.Context, in *UpdateAppCouponWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateAppCouponWithdrawReviewResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetCouponWithdrawReviews(ctx context.Context, in *GetCouponWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetCouponWithdrawReviewsResponse, error) {
	out := new(GetCouponWithdrawReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.gateway.withdraw.coupon.v2.Gateway/GetCouponWithdrawReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppCouponWithdrawReviews(ctx context.Context, in *GetAppCouponWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetAppCouponWithdrawReviewsResponse, error) {
	out := new(GetAppCouponWithdrawReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.gateway.withdraw.coupon.v2.Gateway/GetAppCouponWithdrawReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateCouponWithdrawReview(ctx context.Context, in *UpdateCouponWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateCouponWithdrawReviewResponse, error) {
	out := new(UpdateCouponWithdrawReviewResponse)
	err := c.cc.Invoke(ctx, "/review.gateway.withdraw.coupon.v2.Gateway/UpdateCouponWithdrawReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppCouponWithdrawReview(ctx context.Context, in *UpdateAppCouponWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateAppCouponWithdrawReviewResponse, error) {
	out := new(UpdateAppCouponWithdrawReviewResponse)
	err := c.cc.Invoke(ctx, "/review.gateway.withdraw.coupon.v2.Gateway/UpdateAppCouponWithdrawReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetCouponWithdrawReviews(context.Context, *GetCouponWithdrawReviewsRequest) (*GetCouponWithdrawReviewsResponse, error)
	GetAppCouponWithdrawReviews(context.Context, *GetAppCouponWithdrawReviewsRequest) (*GetAppCouponWithdrawReviewsResponse, error)
	UpdateCouponWithdrawReview(context.Context, *UpdateCouponWithdrawReviewRequest) (*UpdateCouponWithdrawReviewResponse, error)
	UpdateAppCouponWithdrawReview(context.Context, *UpdateAppCouponWithdrawReviewRequest) (*UpdateAppCouponWithdrawReviewResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetCouponWithdrawReviews(context.Context, *GetCouponWithdrawReviewsRequest) (*GetCouponWithdrawReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponWithdrawReviews not implemented")
}
func (UnimplementedGatewayServer) GetAppCouponWithdrawReviews(context.Context, *GetAppCouponWithdrawReviewsRequest) (*GetAppCouponWithdrawReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppCouponWithdrawReviews not implemented")
}
func (UnimplementedGatewayServer) UpdateCouponWithdrawReview(context.Context, *UpdateCouponWithdrawReviewRequest) (*UpdateCouponWithdrawReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCouponWithdrawReview not implemented")
}
func (UnimplementedGatewayServer) UpdateAppCouponWithdrawReview(context.Context, *UpdateAppCouponWithdrawReviewRequest) (*UpdateAppCouponWithdrawReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppCouponWithdrawReview not implemented")
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

func _Gateway_GetCouponWithdrawReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponWithdrawReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCouponWithdrawReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.gateway.withdraw.coupon.v2.Gateway/GetCouponWithdrawReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCouponWithdrawReviews(ctx, req.(*GetCouponWithdrawReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppCouponWithdrawReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppCouponWithdrawReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppCouponWithdrawReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.gateway.withdraw.coupon.v2.Gateway/GetAppCouponWithdrawReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppCouponWithdrawReviews(ctx, req.(*GetAppCouponWithdrawReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateCouponWithdrawReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCouponWithdrawReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateCouponWithdrawReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.gateway.withdraw.coupon.v2.Gateway/UpdateCouponWithdrawReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateCouponWithdrawReview(ctx, req.(*UpdateCouponWithdrawReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppCouponWithdrawReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppCouponWithdrawReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppCouponWithdrawReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.gateway.withdraw.coupon.v2.Gateway/UpdateAppCouponWithdrawReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppCouponWithdrawReview(ctx, req.(*UpdateAppCouponWithdrawReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.gateway.withdraw.coupon.v2.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCouponWithdrawReviews",
			Handler:    _Gateway_GetCouponWithdrawReviews_Handler,
		},
		{
			MethodName: "GetAppCouponWithdrawReviews",
			Handler:    _Gateway_GetAppCouponWithdrawReviews_Handler,
		},
		{
			MethodName: "UpdateCouponWithdrawReview",
			Handler:    _Gateway_UpdateCouponWithdrawReview_Handler,
		},
		{
			MethodName: "UpdateAppCouponWithdrawReview",
			Handler:    _Gateway_UpdateAppCouponWithdrawReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/review/gw/v2/withdraw/coupon/coupon.proto",
}
