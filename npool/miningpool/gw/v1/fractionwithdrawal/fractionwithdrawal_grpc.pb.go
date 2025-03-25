// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/miningpool/gw/v1/fractionwithdrawal/fractionwithdrawal.proto

package fractionwithdrawal

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
	CreateFractionWithdrawal(ctx context.Context, in *CreateFractionWithdrawalRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalResponse, error)
	GetFractionWithdrawal(ctx context.Context, in *GetFractionWithdrawalRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalResponse, error)
	GetUserFractionWithdrawals(ctx context.Context, in *GetUserFractionWithdrawalsRequest, opts ...grpc.CallOption) (*GetUserFractionWithdrawalsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateFractionWithdrawal(ctx context.Context, in *CreateFractionWithdrawalRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalResponse, error) {
	out := new(CreateFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.gateway.fractionwithdrawal.v1.Gateway/CreateFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetFractionWithdrawal(ctx context.Context, in *GetFractionWithdrawalRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalResponse, error) {
	out := new(GetFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.gateway.fractionwithdrawal.v1.Gateway/GetFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetUserFractionWithdrawals(ctx context.Context, in *GetUserFractionWithdrawalsRequest, opts ...grpc.CallOption) (*GetUserFractionWithdrawalsResponse, error) {
	out := new(GetUserFractionWithdrawalsResponse)
	err := c.cc.Invoke(ctx, "/miningpool.gateway.fractionwithdrawal.v1.Gateway/GetUserFractionWithdrawals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateFractionWithdrawal(context.Context, *CreateFractionWithdrawalRequest) (*CreateFractionWithdrawalResponse, error)
	GetFractionWithdrawal(context.Context, *GetFractionWithdrawalRequest) (*GetFractionWithdrawalResponse, error)
	GetUserFractionWithdrawals(context.Context, *GetUserFractionWithdrawalsRequest) (*GetUserFractionWithdrawalsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateFractionWithdrawal(context.Context, *CreateFractionWithdrawalRequest) (*CreateFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFractionWithdrawal not implemented")
}
func (UnimplementedGatewayServer) GetFractionWithdrawal(context.Context, *GetFractionWithdrawalRequest) (*GetFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionWithdrawal not implemented")
}
func (UnimplementedGatewayServer) GetUserFractionWithdrawals(context.Context, *GetUserFractionWithdrawalsRequest) (*GetUserFractionWithdrawalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFractionWithdrawals not implemented")
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

func _Gateway_CreateFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.gateway.fractionwithdrawal.v1.Gateway/CreateFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateFractionWithdrawal(ctx, req.(*CreateFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.gateway.fractionwithdrawal.v1.Gateway/GetFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetFractionWithdrawal(ctx, req.(*GetFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetUserFractionWithdrawals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFractionWithdrawalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetUserFractionWithdrawals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.gateway.fractionwithdrawal.v1.Gateway/GetUserFractionWithdrawals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetUserFractionWithdrawals(ctx, req.(*GetUserFractionWithdrawalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.gateway.fractionwithdrawal.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFractionWithdrawal",
			Handler:    _Gateway_CreateFractionWithdrawal_Handler,
		},
		{
			MethodName: "GetFractionWithdrawal",
			Handler:    _Gateway_GetFractionWithdrawal_Handler,
		},
		{
			MethodName: "GetUserFractionWithdrawals",
			Handler:    _Gateway_GetUserFractionWithdrawals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/gw/v1/fractionwithdrawal/fractionwithdrawal.proto",
}
