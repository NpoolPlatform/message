// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ledger/mw/v1/ledger/ledger.proto

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetIntervalGenerals(ctx context.Context, in *GetIntervalGeneralsRequest, opts ...grpc.CallOption) (*GetIntervalGeneralsResponse, error)
	GetIntervalDetails(ctx context.Context, in *GetIntervalDetailsRequest, opts ...grpc.CallOption) (*GetIntervalDetailsResponse, error)
	GetIntervalProfits(ctx context.Context, in *GetIntervalProfitsRequest, opts ...grpc.CallOption) (*GetIntervalProfitsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetIntervalGenerals(ctx context.Context, in *GetIntervalGeneralsRequest, opts ...grpc.CallOption) (*GetIntervalGeneralsResponse, error) {
	out := new(GetIntervalGeneralsResponse)
	err := c.cc.Invoke(ctx, "/ledger.middleware.ledger1.v1.Middleware/GetIntervalGenerals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetIntervalDetails(ctx context.Context, in *GetIntervalDetailsRequest, opts ...grpc.CallOption) (*GetIntervalDetailsResponse, error) {
	out := new(GetIntervalDetailsResponse)
	err := c.cc.Invoke(ctx, "/ledger.middleware.ledger1.v1.Middleware/GetIntervalDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetIntervalProfits(ctx context.Context, in *GetIntervalProfitsRequest, opts ...grpc.CallOption) (*GetIntervalProfitsResponse, error) {
	out := new(GetIntervalProfitsResponse)
	err := c.cc.Invoke(ctx, "/ledger.middleware.ledger1.v1.Middleware/GetIntervalProfits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetIntervalGenerals(context.Context, *GetIntervalGeneralsRequest) (*GetIntervalGeneralsResponse, error)
	GetIntervalDetails(context.Context, *GetIntervalDetailsRequest) (*GetIntervalDetailsResponse, error)
	GetIntervalProfits(context.Context, *GetIntervalProfitsRequest) (*GetIntervalProfitsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetIntervalGenerals(context.Context, *GetIntervalGeneralsRequest) (*GetIntervalGeneralsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalGenerals not implemented")
}
func (UnimplementedMiddlewareServer) GetIntervalDetails(context.Context, *GetIntervalDetailsRequest) (*GetIntervalDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalDetails not implemented")
}
func (UnimplementedMiddlewareServer) GetIntervalProfits(context.Context, *GetIntervalProfitsRequest) (*GetIntervalProfitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntervalProfits not implemented")
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

func _Middleware_GetIntervalGenerals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalGeneralsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetIntervalGenerals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.middleware.ledger1.v1.Middleware/GetIntervalGenerals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetIntervalGenerals(ctx, req.(*GetIntervalGeneralsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetIntervalDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetIntervalDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.middleware.ledger1.v1.Middleware/GetIntervalDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetIntervalDetails(ctx, req.(*GetIntervalDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetIntervalProfits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntervalProfitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetIntervalProfits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.middleware.ledger1.v1.Middleware/GetIntervalProfits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetIntervalProfits(ctx, req.(*GetIntervalProfitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.middleware.ledger1.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIntervalGenerals",
			Handler:    _Middleware_GetIntervalGenerals_Handler,
		},
		{
			MethodName: "GetIntervalDetails",
			Handler:    _Middleware_GetIntervalDetails_Handler,
		},
		{
			MethodName: "GetIntervalProfits",
			Handler:    _Middleware_GetIntervalProfits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/mw/v1/ledger/ledger.proto",
}
