// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/promotion/promotion.proto

package promotion

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
	CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error)
	CreateAppPromotion(ctx context.Context, in *CreateAppPromotionRequest, opts ...grpc.CallOption) (*CreateAppPromotionResponse, error)
	UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*UpdatePromotionResponse, error)
	UpdateAppPromotion(ctx context.Context, in *UpdateAppPromotionRequest, opts ...grpc.CallOption) (*UpdateAppPromotionResponse, error)
	GetPromotions(ctx context.Context, in *GetPromotionsRequest, opts ...grpc.CallOption) (*GetPromotionsResponse, error)
	GetAppPromotions(ctx context.Context, in *GetAppPromotionsRequest, opts ...grpc.CallOption) (*GetAppPromotionsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error) {
	out := new(CreatePromotionResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/CreatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppPromotion(ctx context.Context, in *CreateAppPromotionRequest, opts ...grpc.CallOption) (*CreateAppPromotionResponse, error) {
	out := new(CreateAppPromotionResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/CreateAppPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*UpdatePromotionResponse, error) {
	out := new(UpdatePromotionResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/UpdatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppPromotion(ctx context.Context, in *UpdateAppPromotionRequest, opts ...grpc.CallOption) (*UpdateAppPromotionResponse, error) {
	out := new(UpdateAppPromotionResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/UpdateAppPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetPromotions(ctx context.Context, in *GetPromotionsRequest, opts ...grpc.CallOption) (*GetPromotionsResponse, error) {
	out := new(GetPromotionsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/GetPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppPromotions(ctx context.Context, in *GetAppPromotionsRequest, opts ...grpc.CallOption) (*GetAppPromotionsResponse, error) {
	out := new(GetAppPromotionsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.promotion.v1.Gateway/GetAppPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error)
	CreateAppPromotion(context.Context, *CreateAppPromotionRequest) (*CreateAppPromotionResponse, error)
	UpdatePromotion(context.Context, *UpdatePromotionRequest) (*UpdatePromotionResponse, error)
	UpdateAppPromotion(context.Context, *UpdateAppPromotionRequest) (*UpdateAppPromotionResponse, error)
	GetPromotions(context.Context, *GetPromotionsRequest) (*GetPromotionsResponse, error)
	GetAppPromotions(context.Context, *GetAppPromotionsRequest) (*GetAppPromotionsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotion not implemented")
}
func (UnimplementedGatewayServer) CreateAppPromotion(context.Context, *CreateAppPromotionRequest) (*CreateAppPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppPromotion not implemented")
}
func (UnimplementedGatewayServer) UpdatePromotion(context.Context, *UpdatePromotionRequest) (*UpdatePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePromotion not implemented")
}
func (UnimplementedGatewayServer) UpdateAppPromotion(context.Context, *UpdateAppPromotionRequest) (*UpdateAppPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppPromotion not implemented")
}
func (UnimplementedGatewayServer) GetPromotions(context.Context, *GetPromotionsRequest) (*GetPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotions not implemented")
}
func (UnimplementedGatewayServer) GetAppPromotions(context.Context, *GetAppPromotionsRequest) (*GetAppPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppPromotions not implemented")
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

func _Gateway_CreatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/CreatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreatePromotion(ctx, req.(*CreatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/CreateAppPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppPromotion(ctx, req.(*CreateAppPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/UpdatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdatePromotion(ctx, req.(*UpdatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/UpdateAppPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppPromotion(ctx, req.(*UpdateAppPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/GetPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetPromotions(ctx, req.(*GetPromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppPromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.promotion.v1.Gateway/GetAppPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppPromotions(ctx, req.(*GetAppPromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.promotion.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePromotion",
			Handler:    _Gateway_CreatePromotion_Handler,
		},
		{
			MethodName: "CreateAppPromotion",
			Handler:    _Gateway_CreateAppPromotion_Handler,
		},
		{
			MethodName: "UpdatePromotion",
			Handler:    _Gateway_UpdatePromotion_Handler,
		},
		{
			MethodName: "UpdateAppPromotion",
			Handler:    _Gateway_UpdateAppPromotion_Handler,
		},
		{
			MethodName: "GetPromotions",
			Handler:    _Gateway_GetPromotions_Handler,
		},
		{
			MethodName: "GetAppPromotions",
			Handler:    _Gateway_GetAppPromotions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/promotion/promotion.proto",
}
