// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/review/gw/v2/withdraw/withdraw.proto

package withdraw

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
	Gateway_GetWithdrawReviews_FullMethodName      = "/review.gateway.withdraw.v2.Gateway/GetWithdrawReviews"
	Gateway_GetAppWithdrawReviews_FullMethodName   = "/review.gateway.withdraw.v2.Gateway/GetAppWithdrawReviews"
	Gateway_UpdateWithdrawReview_FullMethodName    = "/review.gateway.withdraw.v2.Gateway/UpdateWithdrawReview"
	Gateway_UpdateAppWithdrawReview_FullMethodName = "/review.gateway.withdraw.v2.Gateway/UpdateAppWithdrawReview"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetWithdrawReviews(ctx context.Context, in *GetWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetWithdrawReviewsResponse, error)
	GetAppWithdrawReviews(ctx context.Context, in *GetAppWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetAppWithdrawReviewsResponse, error)
	UpdateWithdrawReview(ctx context.Context, in *UpdateWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateWithdrawReviewResponse, error)
	UpdateAppWithdrawReview(ctx context.Context, in *UpdateAppWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateAppWithdrawReviewResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetWithdrawReviews(ctx context.Context, in *GetWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetWithdrawReviewsResponse, error) {
	out := new(GetWithdrawReviewsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetWithdrawReviews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppWithdrawReviews(ctx context.Context, in *GetAppWithdrawReviewsRequest, opts ...grpc.CallOption) (*GetAppWithdrawReviewsResponse, error) {
	out := new(GetAppWithdrawReviewsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppWithdrawReviews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateWithdrawReview(ctx context.Context, in *UpdateWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateWithdrawReviewResponse, error) {
	out := new(UpdateWithdrawReviewResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateWithdrawReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppWithdrawReview(ctx context.Context, in *UpdateAppWithdrawReviewRequest, opts ...grpc.CallOption) (*UpdateAppWithdrawReviewResponse, error) {
	out := new(UpdateAppWithdrawReviewResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppWithdrawReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetWithdrawReviews(context.Context, *GetWithdrawReviewsRequest) (*GetWithdrawReviewsResponse, error)
	GetAppWithdrawReviews(context.Context, *GetAppWithdrawReviewsRequest) (*GetAppWithdrawReviewsResponse, error)
	UpdateWithdrawReview(context.Context, *UpdateWithdrawReviewRequest) (*UpdateWithdrawReviewResponse, error)
	UpdateAppWithdrawReview(context.Context, *UpdateAppWithdrawReviewRequest) (*UpdateAppWithdrawReviewResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetWithdrawReviews(context.Context, *GetWithdrawReviewsRequest) (*GetWithdrawReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawReviews not implemented")
}
func (UnimplementedGatewayServer) GetAppWithdrawReviews(context.Context, *GetAppWithdrawReviewsRequest) (*GetAppWithdrawReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppWithdrawReviews not implemented")
}
func (UnimplementedGatewayServer) UpdateWithdrawReview(context.Context, *UpdateWithdrawReviewRequest) (*UpdateWithdrawReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithdrawReview not implemented")
}
func (UnimplementedGatewayServer) UpdateAppWithdrawReview(context.Context, *UpdateAppWithdrawReviewRequest) (*UpdateAppWithdrawReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppWithdrawReview not implemented")
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

func _Gateway_GetWithdrawReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetWithdrawReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetWithdrawReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetWithdrawReviews(ctx, req.(*GetWithdrawReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppWithdrawReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppWithdrawReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppWithdrawReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppWithdrawReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppWithdrawReviews(ctx, req.(*GetAppWithdrawReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateWithdrawReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWithdrawReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateWithdrawReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateWithdrawReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateWithdrawReview(ctx, req.(*UpdateWithdrawReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppWithdrawReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppWithdrawReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppWithdrawReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppWithdrawReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppWithdrawReview(ctx, req.(*UpdateAppWithdrawReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.gateway.withdraw.v2.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWithdrawReviews",
			Handler:    _Gateway_GetWithdrawReviews_Handler,
		},
		{
			MethodName: "GetAppWithdrawReviews",
			Handler:    _Gateway_GetAppWithdrawReviews_Handler,
		},
		{
			MethodName: "UpdateWithdrawReview",
			Handler:    _Gateway_UpdateWithdrawReview_Handler,
		},
		{
			MethodName: "UpdateAppWithdrawReview",
			Handler:    _Gateway_UpdateAppWithdrawReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/review/gw/v2/withdraw/withdraw.proto",
}
