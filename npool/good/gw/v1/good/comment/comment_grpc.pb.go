// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/good/comment/comment.proto

package comment

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
	Gateway_CreateComment_FullMethodName = "/good.middleware.good1.comment.v1.Gateway/CreateComment"
	Gateway_UpdateComment_FullMethodName = "/good.middleware.good1.comment.v1.Gateway/UpdateComment"
	Gateway_GetComments_FullMethodName   = "/good.middleware.good1.comment.v1.Gateway/GetComments"
	Gateway_GetMyComments_FullMethodName = "/good.middleware.good1.comment.v1.Gateway/GetMyComments"
	Gateway_DeleteComment_FullMethodName = "/good.middleware.good1.comment.v1.Gateway/DeleteComment"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetMyComments(ctx context.Context, in *GetMyCommentsRequest, opts ...grpc.CallOption) (*GetMyCommentsResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyComments(ctx context.Context, in *GetMyCommentsRequest, opts ...grpc.CallOption) (*GetMyCommentsResponse, error) {
	out := new(GetMyCommentsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	GetMyComments(context.Context, *GetMyCommentsRequest) (*GetMyCommentsResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedGatewayServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedGatewayServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedGatewayServer) GetMyComments(context.Context, *GetMyCommentsRequest) (*GetMyCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyComments not implemented")
}
func (UnimplementedGatewayServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
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

func _Gateway_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyComments(ctx, req.(*GetMyCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.good1.comment.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _Gateway_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Gateway_UpdateComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Gateway_GetComments_Handler,
		},
		{
			MethodName: "GetMyComments",
			Handler:    _Gateway_GetMyComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Gateway_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/good/comment/comment.proto",
}
