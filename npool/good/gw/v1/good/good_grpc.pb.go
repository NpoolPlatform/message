// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/good/good.proto

package good

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
	Gateway_CreateGood_FullMethodName = "/good.gateway.good1.v1.Gateway/CreateGood"
	Gateway_GetGood_FullMethodName    = "/good.gateway.good1.v1.Gateway/GetGood"
	Gateway_UpdateGood_FullMethodName = "/good.gateway.good1.v1.Gateway/UpdateGood"
	Gateway_GetGoods_FullMethodName   = "/good.gateway.good1.v1.Gateway/GetGoods"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error)
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
	UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error) {
	out := new(CreateGoodResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, Gateway_GetGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error) {
	out := new(UpdateGoodResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error)
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGood not implemented")
}
func (UnimplementedGatewayServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedGatewayServer) UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGood not implemented")
}
func (UnimplementedGatewayServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
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

func _Gateway_CreateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateGood(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateGood(ctx, req.(*UpdateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.good1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGood",
			Handler:    _Gateway_CreateGood_Handler,
		},
		{
			MethodName: "GetGood",
			Handler:    _Gateway_GetGood_Handler,
		},
		{
			MethodName: "UpdateGood",
			Handler:    _Gateway_UpdateGood_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Gateway_GetGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/good/good.proto",
}
