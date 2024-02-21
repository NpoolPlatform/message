// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/ledger/mw/v2/simulate/ledger/profit/profit.proto

package profit

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
	Middleware_GetProfit_FullMethodName        = "/ledger.middleware.simulate.ledger.profit.v2.Middleware/GetProfit"
	Middleware_GetProfits_FullMethodName       = "/ledger.middleware.simulate.ledger.profit.v2.Middleware/GetProfits"
	Middleware_ExistProfitConds_FullMethodName = "/ledger.middleware.simulate.ledger.profit.v2.Middleware/ExistProfitConds"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetProfit(ctx context.Context, in *GetProfitRequest, opts ...grpc.CallOption) (*GetProfitResponse, error)
	GetProfits(ctx context.Context, in *GetProfitsRequest, opts ...grpc.CallOption) (*GetProfitsResponse, error)
	ExistProfitConds(ctx context.Context, in *ExistProfitCondsRequest, opts ...grpc.CallOption) (*ExistProfitCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetProfit(ctx context.Context, in *GetProfitRequest, opts ...grpc.CallOption) (*GetProfitResponse, error) {
	out := new(GetProfitResponse)
	err := c.cc.Invoke(ctx, Middleware_GetProfit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetProfits(ctx context.Context, in *GetProfitsRequest, opts ...grpc.CallOption) (*GetProfitsResponse, error) {
	out := new(GetProfitsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetProfits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistProfitConds(ctx context.Context, in *ExistProfitCondsRequest, opts ...grpc.CallOption) (*ExistProfitCondsResponse, error) {
	out := new(ExistProfitCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistProfitConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetProfit(context.Context, *GetProfitRequest) (*GetProfitResponse, error)
	GetProfits(context.Context, *GetProfitsRequest) (*GetProfitsResponse, error)
	ExistProfitConds(context.Context, *ExistProfitCondsRequest) (*ExistProfitCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetProfit(context.Context, *GetProfitRequest) (*GetProfitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfit not implemented")
}
func (UnimplementedMiddlewareServer) GetProfits(context.Context, *GetProfitsRequest) (*GetProfitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfits not implemented")
}
func (UnimplementedMiddlewareServer) ExistProfitConds(context.Context, *ExistProfitCondsRequest) (*ExistProfitCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistProfitConds not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_GetProfit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetProfit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetProfit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetProfit(ctx, req.(*GetProfitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetProfits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetProfits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetProfits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetProfits(ctx, req.(*GetProfitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistProfitConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistProfitCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistProfitConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistProfitConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistProfitConds(ctx, req.(*ExistProfitCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.middleware.simulate.ledger.profit.v2.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfit",
			Handler:    _Middleware_GetProfit_Handler,
		},
		{
			MethodName: "GetProfits",
			Handler:    _Middleware_GetProfits_Handler,
		},
		{
			MethodName: "ExistProfitConds",
			Handler:    _Middleware_ExistProfitConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/mw/v2/simulate/ledger/profit/profit.proto",
}