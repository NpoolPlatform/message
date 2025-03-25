// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/app/good/like/like.proto

package like

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
	CreateLike(ctx context.Context, in *CreateLikeRequest, opts ...grpc.CallOption) (*CreateLikeResponse, error)
	GetMyLikes(ctx context.Context, in *GetMyLikesRequest, opts ...grpc.CallOption) (*GetMyLikesResponse, error)
	GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error)
	DeleteLike(ctx context.Context, in *DeleteLikeRequest, opts ...grpc.CallOption) (*DeleteLikeResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateLike(ctx context.Context, in *CreateLikeRequest, opts ...grpc.CallOption) (*CreateLikeResponse, error) {
	out := new(CreateLikeResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.like.v1.Gateway/CreateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyLikes(ctx context.Context, in *GetMyLikesRequest, opts ...grpc.CallOption) (*GetMyLikesResponse, error) {
	out := new(GetMyLikesResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.like.v1.Gateway/GetMyLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error) {
	out := new(GetLikesResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.like.v1.Gateway/GetLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteLike(ctx context.Context, in *DeleteLikeRequest, opts ...grpc.CallOption) (*DeleteLikeResponse, error) {
	out := new(DeleteLikeResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.like.v1.Gateway/DeleteLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateLike(context.Context, *CreateLikeRequest) (*CreateLikeResponse, error)
	GetMyLikes(context.Context, *GetMyLikesRequest) (*GetMyLikesResponse, error)
	GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error)
	DeleteLike(context.Context, *DeleteLikeRequest) (*DeleteLikeResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateLike(context.Context, *CreateLikeRequest) (*CreateLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedGatewayServer) GetMyLikes(context.Context, *GetMyLikesRequest) (*GetMyLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyLikes not implemented")
}
func (UnimplementedGatewayServer) GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikes not implemented")
}
func (UnimplementedGatewayServer) DeleteLike(context.Context, *DeleteLikeRequest) (*DeleteLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLike not implemented")
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

func _Gateway_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.like.v1.Gateway/CreateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateLike(ctx, req.(*CreateLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.like.v1.Gateway/GetMyLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyLikes(ctx, req.(*GetMyLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.like.v1.Gateway/GetLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetLikes(ctx, req.(*GetLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.like.v1.Gateway/DeleteLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteLike(ctx, req.(*DeleteLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.app.good1.like.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLike",
			Handler:    _Gateway_CreateLike_Handler,
		},
		{
			MethodName: "GetMyLikes",
			Handler:    _Gateway_GetMyLikes_Handler,
		},
		{
			MethodName: "GetLikes",
			Handler:    _Gateway_GetLikes_Handler,
		},
		{
			MethodName: "DeleteLike",
			Handler:    _Gateway_DeleteLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/app/good/like/like.proto",
}
