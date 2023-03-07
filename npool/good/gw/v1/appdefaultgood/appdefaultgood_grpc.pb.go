// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/appdefaultgood/appdefaultgood.proto

package appdefaultgood

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
	CreateAppDefaultGood(ctx context.Context, in *CreateAppDefaultGoodRequest, opts ...grpc.CallOption) (*CreateAppDefaultGoodResponse, error)
	CreateNAppDefaultGood(ctx context.Context, in *CreateNAppDefaultGoodRequest, opts ...grpc.CallOption) (*CreateNAppDefaultGoodResponse, error)
	GetAppDefaultGoods(ctx context.Context, in *GetAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodsResponse, error)
	GetNAppDefaultGoods(ctx context.Context, in *GetNAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetNAppDefaultGoodsResponse, error)
	DeleteAppDefaultGood(ctx context.Context, in *DeleteAppDefaultGoodRequest, opts ...grpc.CallOption) (*DeleteAppDefaultGoodResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppDefaultGood(ctx context.Context, in *CreateAppDefaultGoodRequest, opts ...grpc.CallOption) (*CreateAppDefaultGoodResponse, error) {
	out := new(CreateAppDefaultGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appdefaultgood.v1.Gateway/CreateAppDefaultGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateNAppDefaultGood(ctx context.Context, in *CreateNAppDefaultGoodRequest, opts ...grpc.CallOption) (*CreateNAppDefaultGoodResponse, error) {
	out := new(CreateNAppDefaultGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appdefaultgood.v1.Gateway/CreateNAppDefaultGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppDefaultGoods(ctx context.Context, in *GetAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodsResponse, error) {
	out := new(GetAppDefaultGoodsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appdefaultgood.v1.Gateway/GetAppDefaultGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppDefaultGoods(ctx context.Context, in *GetNAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetNAppDefaultGoodsResponse, error) {
	out := new(GetNAppDefaultGoodsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appdefaultgood.v1.Gateway/GetNAppDefaultGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAppDefaultGood(ctx context.Context, in *DeleteAppDefaultGoodRequest, opts ...grpc.CallOption) (*DeleteAppDefaultGoodResponse, error) {
	out := new(DeleteAppDefaultGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appdefaultgood.v1.Gateway/DeleteAppDefaultGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppDefaultGood(context.Context, *CreateAppDefaultGoodRequest) (*CreateAppDefaultGoodResponse, error)
	CreateNAppDefaultGood(context.Context, *CreateNAppDefaultGoodRequest) (*CreateNAppDefaultGoodResponse, error)
	GetAppDefaultGoods(context.Context, *GetAppDefaultGoodsRequest) (*GetAppDefaultGoodsResponse, error)
	GetNAppDefaultGoods(context.Context, *GetNAppDefaultGoodsRequest) (*GetNAppDefaultGoodsResponse, error)
	DeleteAppDefaultGood(context.Context, *DeleteAppDefaultGoodRequest) (*DeleteAppDefaultGoodResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateAppDefaultGood(context.Context, *CreateAppDefaultGoodRequest) (*CreateAppDefaultGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppDefaultGood not implemented")
}
func (UnimplementedGatewayServer) CreateNAppDefaultGood(context.Context, *CreateNAppDefaultGoodRequest) (*CreateNAppDefaultGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNAppDefaultGood not implemented")
}
func (UnimplementedGatewayServer) GetAppDefaultGoods(context.Context, *GetAppDefaultGoodsRequest) (*GetAppDefaultGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDefaultGoods not implemented")
}
func (UnimplementedGatewayServer) GetNAppDefaultGoods(context.Context, *GetNAppDefaultGoodsRequest) (*GetNAppDefaultGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppDefaultGoods not implemented")
}
func (UnimplementedGatewayServer) DeleteAppDefaultGood(context.Context, *DeleteAppDefaultGoodRequest) (*DeleteAppDefaultGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppDefaultGood not implemented")
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

func _Gateway_CreateAppDefaultGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppDefaultGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppDefaultGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appdefaultgood.v1.Gateway/CreateAppDefaultGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppDefaultGood(ctx, req.(*CreateAppDefaultGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateNAppDefaultGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNAppDefaultGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNAppDefaultGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appdefaultgood.v1.Gateway/CreateNAppDefaultGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNAppDefaultGood(ctx, req.(*CreateNAppDefaultGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppDefaultGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDefaultGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppDefaultGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appdefaultgood.v1.Gateway/GetAppDefaultGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppDefaultGoods(ctx, req.(*GetAppDefaultGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppDefaultGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppDefaultGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppDefaultGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appdefaultgood.v1.Gateway/GetNAppDefaultGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppDefaultGoods(ctx, req.(*GetNAppDefaultGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAppDefaultGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppDefaultGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAppDefaultGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appdefaultgood.v1.Gateway/DeleteAppDefaultGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAppDefaultGood(ctx, req.(*DeleteAppDefaultGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.appdefaultgood.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppDefaultGood",
			Handler:    _Gateway_CreateAppDefaultGood_Handler,
		},
		{
			MethodName: "CreateNAppDefaultGood",
			Handler:    _Gateway_CreateNAppDefaultGood_Handler,
		},
		{
			MethodName: "GetAppDefaultGoods",
			Handler:    _Gateway_GetAppDefaultGoods_Handler,
		},
		{
			MethodName: "GetNAppDefaultGoods",
			Handler:    _Gateway_GetNAppDefaultGoods_Handler,
		},
		{
			MethodName: "DeleteAppDefaultGood",
			Handler:    _Gateway_DeleteAppDefaultGood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/appdefaultgood/appdefaultgood.proto",
}
