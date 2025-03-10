// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/app/good/comment/comment.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	GetMyComments(ctx context.Context, in *GetMyCommentsRequest, opts ...grpc.CallOption) (*GetMyCommentsResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	// Run by app admin to set other user's comment
	UpdateUserComment(ctx context.Context, in *UpdateUserCommentRequest, opts ...grpc.CallOption) (*UpdateUserCommentResponse, error)
	DeleteUserComment(ctx context.Context, in *DeleteUserCommentRequest, opts ...grpc.CallOption) (*DeleteUserCommentResponse, error)
	// Run by admin
	AdminUpdateComment(ctx context.Context, in *AdminUpdateCommentRequest, opts ...grpc.CallOption) (*AdminUpdateCommentResponse, error)
	AdminDeleteComment(ctx context.Context, in *AdminDeleteCommentRequest, opts ...grpc.CallOption) (*AdminDeleteCommentResponse, error)
	AdminGetComments(ctx context.Context, in *AdminGetCommentsRequest, opts ...grpc.CallOption) (*AdminGetCommentsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyComments(ctx context.Context, in *GetMyCommentsRequest, opts ...grpc.CallOption) (*GetMyCommentsResponse, error) {
	out := new(GetMyCommentsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/GetMyComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateUserComment(ctx context.Context, in *UpdateUserCommentRequest, opts ...grpc.CallOption) (*UpdateUserCommentResponse, error) {
	out := new(UpdateUserCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/UpdateUserComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteUserComment(ctx context.Context, in *DeleteUserCommentRequest, opts ...grpc.CallOption) (*DeleteUserCommentResponse, error) {
	out := new(DeleteUserCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/DeleteUserComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateComment(ctx context.Context, in *AdminUpdateCommentRequest, opts ...grpc.CallOption) (*AdminUpdateCommentResponse, error) {
	out := new(AdminUpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/AdminUpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteComment(ctx context.Context, in *AdminDeleteCommentRequest, opts ...grpc.CallOption) (*AdminDeleteCommentResponse, error) {
	out := new(AdminDeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/AdminDeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetComments(ctx context.Context, in *AdminGetCommentsRequest, opts ...grpc.CallOption) (*AdminGetCommentsResponse, error) {
	out := new(AdminGetCommentsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.comment.v1.Gateway/AdminGetComments", in, out, opts...)
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
	GetMyComments(context.Context, *GetMyCommentsRequest) (*GetMyCommentsResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	// Run by app admin to set other user's comment
	UpdateUserComment(context.Context, *UpdateUserCommentRequest) (*UpdateUserCommentResponse, error)
	DeleteUserComment(context.Context, *DeleteUserCommentRequest) (*DeleteUserCommentResponse, error)
	// Run by admin
	AdminUpdateComment(context.Context, *AdminUpdateCommentRequest) (*AdminUpdateCommentResponse, error)
	AdminDeleteComment(context.Context, *AdminDeleteCommentRequest) (*AdminDeleteCommentResponse, error)
	AdminGetComments(context.Context, *AdminGetCommentsRequest) (*AdminGetCommentsResponse, error)
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
func (UnimplementedGatewayServer) GetMyComments(context.Context, *GetMyCommentsRequest) (*GetMyCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyComments not implemented")
}
func (UnimplementedGatewayServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedGatewayServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedGatewayServer) UpdateUserComment(context.Context, *UpdateUserCommentRequest) (*UpdateUserCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserComment not implemented")
}
func (UnimplementedGatewayServer) DeleteUserComment(context.Context, *DeleteUserCommentRequest) (*DeleteUserCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserComment not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateComment(context.Context, *AdminUpdateCommentRequest) (*AdminUpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateComment not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteComment(context.Context, *AdminDeleteCommentRequest) (*AdminDeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteComment not implemented")
}
func (UnimplementedGatewayServer) AdminGetComments(context.Context, *AdminGetCommentsRequest) (*AdminGetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetComments not implemented")
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
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/CreateComment",
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
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
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
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/GetMyComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyComments(ctx, req.(*GetMyCommentsRequest))
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
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetComments(ctx, req.(*GetCommentsRequest))
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
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateUserComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateUserComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/UpdateUserComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateUserComment(ctx, req.(*UpdateUserCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteUserComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteUserComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/DeleteUserComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteUserComment(ctx, req.(*DeleteUserCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/AdminUpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateComment(ctx, req.(*AdminUpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/AdminDeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteComment(ctx, req.(*AdminDeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.comment.v1.Gateway/AdminGetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetComments(ctx, req.(*AdminGetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.app.good1.comment.v1.Gateway",
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
			MethodName: "GetMyComments",
			Handler:    _Gateway_GetMyComments_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Gateway_GetComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Gateway_DeleteComment_Handler,
		},
		{
			MethodName: "UpdateUserComment",
			Handler:    _Gateway_UpdateUserComment_Handler,
		},
		{
			MethodName: "DeleteUserComment",
			Handler:    _Gateway_DeleteUserComment_Handler,
		},
		{
			MethodName: "AdminUpdateComment",
			Handler:    _Gateway_AdminUpdateComment_Handler,
		},
		{
			MethodName: "AdminDeleteComment",
			Handler:    _Gateway_AdminDeleteComment_Handler,
		},
		{
			MethodName: "AdminGetComments",
			Handler:    _Gateway_AdminGetComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/app/good/comment/comment.proto",
}
