// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: npool/review/mw/v2/mw.proto

package v2

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	GetObjectReview(ctx context.Context, in *GetObjectReviewRequest, opts ...grpc.CallOption) (*GetObjectReviewResponse, error)
	GetObjectReviews(ctx context.Context, in *GetObjectReviewsRequest, opts ...grpc.CallOption) (*GetObjectReviewsResponse, error)
	GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
	DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error) {
	out := new(CreateReviewResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetObjectReview(ctx context.Context, in *GetObjectReviewRequest, opts ...grpc.CallOption) (*GetObjectReviewResponse, error) {
	out := new(GetObjectReviewResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/GetObjectReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetObjectReviews(ctx context.Context, in *GetObjectReviewsRequest, opts ...grpc.CallOption) (*GetObjectReviewsResponse, error) {
	out := new(GetObjectReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/GetObjectReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/GetReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteReview(ctx context.Context, in *DeleteReviewRequest, opts ...grpc.CallOption) (*DeleteReviewResponse, error) {
	out := new(DeleteReviewResponse)
	err := c.cc.Invoke(ctx, "/review.middleware.v2.Middleware/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	GetObjectReview(context.Context, *GetObjectReviewRequest) (*GetObjectReviewResponse, error)
	GetObjectReviews(context.Context, *GetObjectReviewsRequest) (*GetObjectReviewsResponse, error)
	GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedMiddlewareServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedMiddlewareServer) GetObjectReview(context.Context, *GetObjectReviewRequest) (*GetObjectReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectReview not implemented")
}
func (UnimplementedMiddlewareServer) GetObjectReviews(context.Context, *GetObjectReviewsRequest) (*GetObjectReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectReviews not implemented")
}
func (UnimplementedMiddlewareServer) GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviews not implemented")
}
func (UnimplementedMiddlewareServer) DeleteReview(context.Context, *DeleteReviewRequest) (*DeleteReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
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

func _Middleware_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetObjectReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetObjectReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/GetObjectReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetObjectReview(ctx, req.(*GetObjectReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetObjectReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetObjectReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/GetObjectReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetObjectReviews(ctx, req.(*GetObjectReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/GetReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.middleware.v2.Middleware/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteReview(ctx, req.(*DeleteReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.middleware.v2.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _Middleware_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _Middleware_UpdateReview_Handler,
		},
		{
			MethodName: "GetObjectReview",
			Handler:    _Middleware_GetObjectReview_Handler,
		},
		{
			MethodName: "GetObjectReviews",
			Handler:    _Middleware_GetObjectReviews_Handler,
		},
		{
			MethodName: "GetReviews",
			Handler:    _Middleware_GetReviews_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _Middleware_DeleteReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/review/mw/v2/mw.proto",
}
