// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/gw/v1/outofgas/outofgas.proto

package outofgas

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
	Gateway_GetOutOfGases_FullMethodName      = "/order.gateway.outofgas.v1.Gateway/GetOutOfGases"
	Gateway_GetMyOutOfGases_FullMethodName    = "/order.gateway.outofgas.v1.Gateway/GetMyOutOfGases"
	Gateway_AdminGetOutOfGases_FullMethodName = "/order.gateway.outofgas.v1.Gateway/AdminGetOutOfGases"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetOutOfGases(ctx context.Context, in *GetOutOfGasesRequest, opts ...grpc.CallOption) (*GetOutOfGasesResponse, error)
	GetMyOutOfGases(ctx context.Context, in *GetMyOutOfGasesRequest, opts ...grpc.CallOption) (*GetMyOutOfGasesResponse, error)
	// Admin apis
	AdminGetOutOfGases(ctx context.Context, in *AdminGetOutOfGasesRequest, opts ...grpc.CallOption) (*AdminGetOutOfGasesResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetOutOfGases(ctx context.Context, in *GetOutOfGasesRequest, opts ...grpc.CallOption) (*GetOutOfGasesResponse, error) {
	out := new(GetOutOfGasesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetOutOfGases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyOutOfGases(ctx context.Context, in *GetMyOutOfGasesRequest, opts ...grpc.CallOption) (*GetMyOutOfGasesResponse, error) {
	out := new(GetMyOutOfGasesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyOutOfGases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetOutOfGases(ctx context.Context, in *AdminGetOutOfGasesRequest, opts ...grpc.CallOption) (*AdminGetOutOfGasesResponse, error) {
	out := new(AdminGetOutOfGasesResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetOutOfGases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetOutOfGases(context.Context, *GetOutOfGasesRequest) (*GetOutOfGasesResponse, error)
	GetMyOutOfGases(context.Context, *GetMyOutOfGasesRequest) (*GetMyOutOfGasesResponse, error)
	// Admin apis
	AdminGetOutOfGases(context.Context, *AdminGetOutOfGasesRequest) (*AdminGetOutOfGasesResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetOutOfGases(context.Context, *GetOutOfGasesRequest) (*GetOutOfGasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutOfGases not implemented")
}
func (UnimplementedGatewayServer) GetMyOutOfGases(context.Context, *GetMyOutOfGasesRequest) (*GetMyOutOfGasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyOutOfGases not implemented")
}
func (UnimplementedGatewayServer) AdminGetOutOfGases(context.Context, *AdminGetOutOfGasesRequest) (*AdminGetOutOfGasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetOutOfGases not implemented")
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

func _Gateway_GetOutOfGases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutOfGasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetOutOfGases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetOutOfGases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetOutOfGases(ctx, req.(*GetOutOfGasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyOutOfGases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyOutOfGasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyOutOfGases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyOutOfGases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyOutOfGases(ctx, req.(*GetMyOutOfGasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetOutOfGases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetOutOfGasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetOutOfGases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetOutOfGases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetOutOfGases(ctx, req.(*AdminGetOutOfGasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.outofgas.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOutOfGases",
			Handler:    _Gateway_GetOutOfGases_Handler,
		},
		{
			MethodName: "GetMyOutOfGases",
			Handler:    _Gateway_GetMyOutOfGases_Handler,
		},
		{
			MethodName: "AdminGetOutOfGases",
			Handler:    _Gateway_AdminGetOutOfGases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/outofgas/outofgas.proto",
}