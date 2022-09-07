// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ledger/gw/v1/ledger/ledger.proto

package ledger

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
	GetGenerals(ctx context.Context, in *GetGeneralsRequest, opts ...grpc.CallOption) (*GetGeneralsResponse, error)
	GetIntervalGenerals(ctx context.Context, in *GetIntervalGeneralsRequest, opts ...grpc.CallOption) (*GetIntervalGeneralsResponse, error)
	GetDetails(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*GetDetailsResponse, error)
	GetProfits(ctx context.Context, in *GetProfitsRequest, opts ...grpc.CallOption) (*GetProfitsResponse, error)
	GetIntervalProfits(ctx context.Context, in *GetIntervalProfitsRequest, opts ...grpc.CallOption) (*GetIntervalProfitsResponse, error)
	GetGoodProfits(ctx context.Context, in *GetGoodProfitsRequest, opts ...grpc.CallOption) (*GetGoodProfitsResponse, error)
	CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error)
	GetWithdraws(ctx context.Context, in *GetWithdrawsRequest, opts ...grpc.CallOption) (*GetWithdrawsResponse, error)
	GetIntervalWithdraws(ctx context.Context, in *GetIntervalWithdrawsRequest, opts ...grpc.CallOption) (*GetIntervalWithdrawsResponse, error)
	GetAppWithdraws(ctx context.Context, in *GetAppWithdrawsRequest, opts ...grpc.CallOption) (*GetAppWithdrawsResponse, error)
	GetNAppWithdraws(ctx context.Context, in *GetNAppWithdrawsRequest, opts ...grpc.CallOption) (*GetNAppWithdrawsResponse, error)
	CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error)
	CreateAppUserDeposit(ctx context.Context, in *CreateAppUserDepositRequest, opts ...grpc.CallOption) (*CreateAppUserDepositResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetGenerals(ctx context.Context, in *GetGeneralsRequest, opts ...grpc.CallOption) (*GetGeneralsResponse, error) {
	out := new(GetGeneralsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetGenerals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetIntervalGenerals(ctx context.Context, in *GetIntervalGeneralsRequest, opts ...grpc.CallOption) (*GetIntervalGeneralsResponse, error) {
	out := new(GetIntervalGeneralsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetIntervalGenerals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetDetails(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*GetDetailsResponse, error) {
	out := new(GetDetailsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetProfits(ctx context.Context, in *GetProfitsRequest, opts ...grpc.CallOption) (*GetProfitsResponse, error) {
	out := new(GetProfitsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetProfits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetIntervalProfits(ctx context.Context, in *GetIntervalProfitsRequest, opts ...grpc.CallOption) (*GetIntervalProfitsResponse, error) {
	out := new(GetIntervalProfitsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetIntervalProfits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGoodProfits(ctx context.Context, in *GetGoodProfitsRequest, opts ...grpc.CallOption) (*GetGoodProfitsResponse, error) {
	out := new(GetGoodProfitsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetGoodProfits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawResponse, error) {
	out := new(CreateWithdrawResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/CreateWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetWithdraws(ctx context.Context, in *GetWithdrawsRequest, opts ...grpc.CallOption) (*GetWithdrawsResponse, error) {
	out := new(GetWithdrawsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetWithdraws", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetIntervalWithdraws(ctx context.Context, in *GetIntervalWithdrawsRequest, opts ...grpc.CallOption) (*GetIntervalWithdrawsResponse, error) {
	out := new(GetIntervalWithdrawsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetIntervalWithdraws", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppWithdraws(ctx context.Context, in *GetAppWithdrawsRequest, opts ...grpc.CallOption) (*GetAppWithdrawsResponse, error) {
	out := new(GetAppWithdrawsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetAppWithdraws", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppWithdraws(ctx context.Context, in *GetNAppWithdrawsRequest, opts ...grpc.CallOption) (*GetNAppWithdrawsResponse, error) {
	out := new(GetNAppWithdrawsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/GetNAppWithdraws", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateTransfer(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*CreateTransferResponse, error) {
	out := new(CreateTransferResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/CreateTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppUserDeposit(ctx context.Context, in *CreateAppUserDepositRequest, opts ...grpc.CallOption) (*CreateAppUserDepositResponse, error) {
	out := new(CreateAppUserDepositResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.ledger1.v1.Gateway/CreateAppUserDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetGenerals(context.Context, *GetGeneralsRequest) (*GetGeneralsResponse, error)
	GetIntervalGenerals(context.Context, *GetIntervalGeneralsRequest) (*GetIntervalGeneralsResponse, error)
	GetDetails(context.Context, *GetDetailsRequest) (*GetDetailsResponse, error)
	GetProfits(context.Context, *GetProfitsRequest) (*GetProfitsResponse, error)
	GetIntervalProfits(context.Context, *GetIntervalProfitsRequest) (*GetIntervalProfitsResponse, error)
	GetGoodProfits(context.Context, *GetGoodProfitsRequest) (*GetGoodProfitsResponse, error)
	CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawResponse, error)
	GetWithdraws(context.Context, *GetWithdrawsRequest) (*GetWithdrawsResponse, error)
	GetIntervalWithdraws(context.Context, *GetIntervalWithdrawsRequest) (*GetIntervalWithdrawsResponse, error)
	GetAppWithdraws(context.Context, *GetAppWithdrawsRequest) (*GetAppWithdrawsResponse, error)
	GetNAppWithdraws(context.Context, *GetNAppWithdrawsRequest) (*GetNAppWithdrawsResponse, error)
	CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error)
	CreateAppUserDeposit(context.Context, *CreateAppUserDepositRequest) (*CreateAppUserDepositResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetGenerals(context.Context, *GetGeneralsRequest) (*GetGeneralsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenerals not implemented")
}
func (UnimplementedGatewayServer) GetIntervalGenerals(context.Context, *GetIntervalGeneralsRequest) (*GetIntervalGeneralsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalGenerals not implemented")
}
func (UnimplementedGatewayServer) GetDetails(context.Context, *GetDetailsRequest) (*GetDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetails not implemented")
}
func (UnimplementedGatewayServer) GetProfits(context.Context, *GetProfitsRequest) (*GetProfitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfits not implemented")
}
func (UnimplementedGatewayServer) GetIntervalProfits(context.Context, *GetIntervalProfitsRequest) (*GetIntervalProfitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalProfits not implemented")
}
func (UnimplementedGatewayServer) GetGoodProfits(context.Context, *GetGoodProfitsRequest) (*GetGoodProfitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodProfits not implemented")
}
func (UnimplementedGatewayServer) CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWithdraw not implemented")
}
func (UnimplementedGatewayServer) GetWithdraws(context.Context, *GetWithdrawsRequest) (*GetWithdrawsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdraws not implemented")
}
func (UnimplementedGatewayServer) GetIntervalWithdraws(context.Context, *GetIntervalWithdrawsRequest) (*GetIntervalWithdrawsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalWithdraws not implemented")
}
func (UnimplementedGatewayServer) GetAppWithdraws(context.Context, *GetAppWithdrawsRequest) (*GetAppWithdrawsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppWithdraws not implemented")
}
func (UnimplementedGatewayServer) GetNAppWithdraws(context.Context, *GetNAppWithdrawsRequest) (*GetNAppWithdrawsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppWithdraws not implemented")
}
func (UnimplementedGatewayServer) CreateTransfer(context.Context, *CreateTransferRequest) (*CreateTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransfer not implemented")
}
func (UnimplementedGatewayServer) CreateAppUserDeposit(context.Context, *CreateAppUserDepositRequest) (*CreateAppUserDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUserDeposit not implemented")
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

func _Gateway_GetGenerals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeneralsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGenerals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetGenerals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGenerals(ctx, req.(*GetGeneralsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetIntervalGenerals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalGeneralsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetIntervalGenerals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetIntervalGenerals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetIntervalGenerals(ctx, req.(*GetIntervalGeneralsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetDetails(ctx, req.(*GetDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetProfits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetProfits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetProfits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetProfits(ctx, req.(*GetProfitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetIntervalProfits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalProfitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetIntervalProfits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetIntervalProfits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetIntervalProfits(ctx, req.(*GetIntervalProfitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGoodProfits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodProfitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGoodProfits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetGoodProfits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGoodProfits(ctx, req.(*GetGoodProfitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/CreateWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateWithdraw(ctx, req.(*CreateWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetWithdraws_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetWithdraws(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetWithdraws",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetWithdraws(ctx, req.(*GetWithdrawsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetIntervalWithdraws_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalWithdrawsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetIntervalWithdraws(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetIntervalWithdraws",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetIntervalWithdraws(ctx, req.(*GetIntervalWithdrawsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppWithdraws_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppWithdrawsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppWithdraws(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetAppWithdraws",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppWithdraws(ctx, req.(*GetAppWithdrawsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppWithdraws_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppWithdrawsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppWithdraws(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/GetNAppWithdraws",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppWithdraws(ctx, req.(*GetNAppWithdrawsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/CreateTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateTransfer(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppUserDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppUserDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.ledger1.v1.Gateway/CreateAppUserDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppUserDeposit(ctx, req.(*CreateAppUserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.gateway.ledger1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenerals",
			Handler:    _Gateway_GetGenerals_Handler,
		},
		{
			MethodName: "GetIntervalGenerals",
			Handler:    _Gateway_GetIntervalGenerals_Handler,
		},
		{
			MethodName: "GetDetails",
			Handler:    _Gateway_GetDetails_Handler,
		},
		{
			MethodName: "GetProfits",
			Handler:    _Gateway_GetProfits_Handler,
		},
		{
			MethodName: "GetIntervalProfits",
			Handler:    _Gateway_GetIntervalProfits_Handler,
		},
		{
			MethodName: "GetGoodProfits",
			Handler:    _Gateway_GetGoodProfits_Handler,
		},
		{
			MethodName: "CreateWithdraw",
			Handler:    _Gateway_CreateWithdraw_Handler,
		},
		{
			MethodName: "GetWithdraws",
			Handler:    _Gateway_GetWithdraws_Handler,
		},
		{
			MethodName: "GetIntervalWithdraws",
			Handler:    _Gateway_GetIntervalWithdraws_Handler,
		},
		{
			MethodName: "GetAppWithdraws",
			Handler:    _Gateway_GetAppWithdraws_Handler,
		},
		{
			MethodName: "GetNAppWithdraws",
			Handler:    _Gateway_GetNAppWithdraws_Handler,
		},
		{
			MethodName: "CreateTransfer",
			Handler:    _Gateway_CreateTransfer_Handler,
		},
		{
			MethodName: "CreateAppUserDeposit",
			Handler:    _Gateway_CreateAppUserDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/gw/v1/ledger/ledger.proto",
}
