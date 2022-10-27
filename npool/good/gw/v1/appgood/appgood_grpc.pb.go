// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/appgood/appgood.proto

package appgood

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
	CreateNAppGood(ctx context.Context, in *CreateNAppGoodRequest, opts ...grpc.CallOption) (*CreateNAppGoodResponse, error)
	GetAppGoods(ctx context.Context, in *GetAppGoodsRequest, opts ...grpc.CallOption) (*GetAppGoodsResponse, error)
	GetAppGood(ctx context.Context, in *GetAppGoodRequest, opts ...grpc.CallOption) (*GetAppGoodResponse, error)
	GetNAppGoods(ctx context.Context, in *GetNAppGoodsRequest, opts ...grpc.CallOption) (*GetNAppGoodsResponse, error)
	UpdateAppGood(ctx context.Context, in *UpdateAppGoodRequest, opts ...grpc.CallOption) (*UpdateAppGoodResponse, error)
	UpdateNAppGood(ctx context.Context, in *UpdateNAppGoodRequest, opts ...grpc.CallOption) (*UpdateNAppGoodResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateNAppGood(ctx context.Context, in *CreateNAppGoodRequest, opts ...grpc.CallOption) (*CreateNAppGoodResponse, error) {
	out := new(CreateNAppGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/CreateNAppGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppGoods(ctx context.Context, in *GetAppGoodsRequest, opts ...grpc.CallOption) (*GetAppGoodsResponse, error) {
	out := new(GetAppGoodsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/GetAppGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppGood(ctx context.Context, in *GetAppGoodRequest, opts ...grpc.CallOption) (*GetAppGoodResponse, error) {
	out := new(GetAppGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/GetAppGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppGoods(ctx context.Context, in *GetNAppGoodsRequest, opts ...grpc.CallOption) (*GetNAppGoodsResponse, error) {
	out := new(GetNAppGoodsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/GetNAppGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppGood(ctx context.Context, in *UpdateAppGoodRequest, opts ...grpc.CallOption) (*UpdateAppGoodResponse, error) {
	out := new(UpdateAppGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/UpdateAppGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateNAppGood(ctx context.Context, in *UpdateNAppGoodRequest, opts ...grpc.CallOption) (*UpdateNAppGoodResponse, error) {
	out := new(UpdateNAppGoodResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.appgood.v1.Gateway/UpdateNAppGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateNAppGood(context.Context, *CreateNAppGoodRequest) (*CreateNAppGoodResponse, error)
	GetAppGoods(context.Context, *GetAppGoodsRequest) (*GetAppGoodsResponse, error)
	GetAppGood(context.Context, *GetAppGoodRequest) (*GetAppGoodResponse, error)
	GetNAppGoods(context.Context, *GetNAppGoodsRequest) (*GetNAppGoodsResponse, error)
	UpdateAppGood(context.Context, *UpdateAppGoodRequest) (*UpdateAppGoodResponse, error)
	UpdateNAppGood(context.Context, *UpdateNAppGoodRequest) (*UpdateNAppGoodResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateNAppGood(context.Context, *CreateNAppGoodRequest) (*CreateNAppGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNAppGood not implemented")
}
func (UnimplementedGatewayServer) GetAppGoods(context.Context, *GetAppGoodsRequest) (*GetAppGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoods not implemented")
}
func (UnimplementedGatewayServer) GetAppGood(context.Context, *GetAppGoodRequest) (*GetAppGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGood not implemented")
}
func (UnimplementedGatewayServer) GetNAppGoods(context.Context, *GetNAppGoodsRequest) (*GetNAppGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppGoods not implemented")
}
func (UnimplementedGatewayServer) UpdateAppGood(context.Context, *UpdateAppGoodRequest) (*UpdateAppGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppGood not implemented")
}
func (UnimplementedGatewayServer) UpdateNAppGood(context.Context, *UpdateNAppGoodRequest) (*UpdateNAppGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNAppGood not implemented")
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

func _Gateway_CreateNAppGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNAppGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNAppGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/CreateNAppGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNAppGood(ctx, req.(*CreateNAppGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/GetAppGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGoods(ctx, req.(*GetAppGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/GetAppGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGood(ctx, req.(*GetAppGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/GetNAppGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppGoods(ctx, req.(*GetNAppGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/UpdateAppGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppGood(ctx, req.(*UpdateAppGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateNAppGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNAppGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateNAppGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.appgood.v1.Gateway/UpdateNAppGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateNAppGood(ctx, req.(*UpdateNAppGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.appgood.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNAppGood",
			Handler:    _Gateway_CreateNAppGood_Handler,
		},
		{
			MethodName: "GetAppGoods",
			Handler:    _Gateway_GetAppGoods_Handler,
		},
		{
			MethodName: "GetAppGood",
			Handler:    _Gateway_GetAppGood_Handler,
		},
		{
			MethodName: "GetNAppGoods",
			Handler:    _Gateway_GetNAppGoods_Handler,
		},
		{
			MethodName: "UpdateAppGood",
			Handler:    _Gateway_UpdateAppGood_Handler,
		},
		{
			MethodName: "UpdateNAppGood",
			Handler:    _Gateway_UpdateNAppGood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/appgood/appgood.proto",
}
