// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/recommend/recommend.proto

package recommend

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
	CreateRecommend(ctx context.Context, in *CreateRecommendRequest, opts ...grpc.CallOption) (*CreateRecommendResponse, error)
	CreateAppRecommend(ctx context.Context, in *CreateAppRecommendRequest, opts ...grpc.CallOption) (*CreateAppRecommendResponse, error)
	UpdateRecommend(ctx context.Context, in *UpdateRecommendRequest, opts ...grpc.CallOption) (*UpdateRecommendResponse, error)
	UpdateAppRecommend(ctx context.Context, in *UpdateAppRecommendRequest, opts ...grpc.CallOption) (*UpdateAppRecommendResponse, error)
	GetRecommends(ctx context.Context, in *GetRecommendsRequest, opts ...grpc.CallOption) (*GetRecommendsResponse, error)
	GetAppRecommends(ctx context.Context, in *GetAppRecommendsRequest, opts ...grpc.CallOption) (*GetAppRecommendsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateRecommend(ctx context.Context, in *CreateRecommendRequest, opts ...grpc.CallOption) (*CreateRecommendResponse, error) {
	out := new(CreateRecommendResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/CreateRecommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppRecommend(ctx context.Context, in *CreateAppRecommendRequest, opts ...grpc.CallOption) (*CreateAppRecommendResponse, error) {
	out := new(CreateAppRecommendResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/CreateAppRecommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateRecommend(ctx context.Context, in *UpdateRecommendRequest, opts ...grpc.CallOption) (*UpdateRecommendResponse, error) {
	out := new(UpdateRecommendResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/UpdateRecommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppRecommend(ctx context.Context, in *UpdateAppRecommendRequest, opts ...grpc.CallOption) (*UpdateAppRecommendResponse, error) {
	out := new(UpdateAppRecommendResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/UpdateAppRecommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRecommends(ctx context.Context, in *GetRecommendsRequest, opts ...grpc.CallOption) (*GetRecommendsResponse, error) {
	out := new(GetRecommendsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/GetRecommends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppRecommends(ctx context.Context, in *GetAppRecommendsRequest, opts ...grpc.CallOption) (*GetAppRecommendsResponse, error) {
	out := new(GetAppRecommendsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.recommend.v1.Gateway/GetAppRecommends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateRecommend(context.Context, *CreateRecommendRequest) (*CreateRecommendResponse, error)
	CreateAppRecommend(context.Context, *CreateAppRecommendRequest) (*CreateAppRecommendResponse, error)
	UpdateRecommend(context.Context, *UpdateRecommendRequest) (*UpdateRecommendResponse, error)
	UpdateAppRecommend(context.Context, *UpdateAppRecommendRequest) (*UpdateAppRecommendResponse, error)
	GetRecommends(context.Context, *GetRecommendsRequest) (*GetRecommendsResponse, error)
	GetAppRecommends(context.Context, *GetAppRecommendsRequest) (*GetAppRecommendsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateRecommend(context.Context, *CreateRecommendRequest) (*CreateRecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecommend not implemented")
}
func (UnimplementedGatewayServer) CreateAppRecommend(context.Context, *CreateAppRecommendRequest) (*CreateAppRecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppRecommend not implemented")
}
func (UnimplementedGatewayServer) UpdateRecommend(context.Context, *UpdateRecommendRequest) (*UpdateRecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecommend not implemented")
}
func (UnimplementedGatewayServer) UpdateAppRecommend(context.Context, *UpdateAppRecommendRequest) (*UpdateAppRecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppRecommend not implemented")
}
func (UnimplementedGatewayServer) GetRecommends(context.Context, *GetRecommendsRequest) (*GetRecommendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommends not implemented")
}
func (UnimplementedGatewayServer) GetAppRecommends(context.Context, *GetAppRecommendsRequest) (*GetAppRecommendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppRecommends not implemented")
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

func _Gateway_CreateRecommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateRecommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/CreateRecommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateRecommend(ctx, req.(*CreateRecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppRecommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppRecommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/CreateAppRecommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppRecommend(ctx, req.(*CreateAppRecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateRecommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateRecommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/UpdateRecommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateRecommend(ctx, req.(*UpdateRecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppRecommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppRecommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/UpdateAppRecommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppRecommend(ctx, req.(*UpdateAppRecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetRecommends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRecommends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/GetRecommends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRecommends(ctx, req.(*GetRecommendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppRecommends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRecommendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppRecommends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.recommend.v1.Gateway/GetAppRecommends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppRecommends(ctx, req.(*GetAppRecommendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.recommend.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecommend",
			Handler:    _Gateway_CreateRecommend_Handler,
		},
		{
			MethodName: "CreateAppRecommend",
			Handler:    _Gateway_CreateAppRecommend_Handler,
		},
		{
			MethodName: "UpdateRecommend",
			Handler:    _Gateway_UpdateRecommend_Handler,
		},
		{
			MethodName: "UpdateAppRecommend",
			Handler:    _Gateway_UpdateAppRecommend_Handler,
		},
		{
			MethodName: "GetRecommends",
			Handler:    _Gateway_GetRecommends_Handler,
		},
		{
			MethodName: "GetAppRecommends",
			Handler:    _Gateway_GetAppRecommends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/recommend/recommend.proto",
}
