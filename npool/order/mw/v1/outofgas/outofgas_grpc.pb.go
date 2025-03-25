// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/outofgas/outofgas.proto

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetOutOfGas(ctx context.Context, in *GetOutOfGasRequest, opts ...grpc.CallOption) (*GetOutOfGasResponse, error)
	GetOutOfGases(ctx context.Context, in *GetOutOfGasesRequest, opts ...grpc.CallOption) (*GetOutOfGasesResponse, error)
	ExistOutOfGas(ctx context.Context, in *ExistOutOfGasRequest, opts ...grpc.CallOption) (*ExistOutOfGasResponse, error)
	ExistOutOfGasConds(ctx context.Context, in *ExistOutOfGasCondsRequest, opts ...grpc.CallOption) (*ExistOutOfGasCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetOutOfGas(ctx context.Context, in *GetOutOfGasRequest, opts ...grpc.CallOption) (*GetOutOfGasResponse, error) {
	out := new(GetOutOfGasResponse)
	err := c.cc.Invoke(ctx, "/order.middleware.outofgas.v1.Middleware/GetOutOfGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOutOfGases(ctx context.Context, in *GetOutOfGasesRequest, opts ...grpc.CallOption) (*GetOutOfGasesResponse, error) {
	out := new(GetOutOfGasesResponse)
	err := c.cc.Invoke(ctx, "/order.middleware.outofgas.v1.Middleware/GetOutOfGases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOutOfGas(ctx context.Context, in *ExistOutOfGasRequest, opts ...grpc.CallOption) (*ExistOutOfGasResponse, error) {
	out := new(ExistOutOfGasResponse)
	err := c.cc.Invoke(ctx, "/order.middleware.outofgas.v1.Middleware/ExistOutOfGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistOutOfGasConds(ctx context.Context, in *ExistOutOfGasCondsRequest, opts ...grpc.CallOption) (*ExistOutOfGasCondsResponse, error) {
	out := new(ExistOutOfGasCondsResponse)
	err := c.cc.Invoke(ctx, "/order.middleware.outofgas.v1.Middleware/ExistOutOfGasConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetOutOfGas(context.Context, *GetOutOfGasRequest) (*GetOutOfGasResponse, error)
	GetOutOfGases(context.Context, *GetOutOfGasesRequest) (*GetOutOfGasesResponse, error)
	ExistOutOfGas(context.Context, *ExistOutOfGasRequest) (*ExistOutOfGasResponse, error)
	ExistOutOfGasConds(context.Context, *ExistOutOfGasCondsRequest) (*ExistOutOfGasCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetOutOfGas(context.Context, *GetOutOfGasRequest) (*GetOutOfGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutOfGas not implemented")
}
func (UnimplementedMiddlewareServer) GetOutOfGases(context.Context, *GetOutOfGasesRequest) (*GetOutOfGasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutOfGases not implemented")
}
func (UnimplementedMiddlewareServer) ExistOutOfGas(context.Context, *ExistOutOfGasRequest) (*ExistOutOfGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOutOfGas not implemented")
}
func (UnimplementedMiddlewareServer) ExistOutOfGasConds(context.Context, *ExistOutOfGasCondsRequest) (*ExistOutOfGasCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistOutOfGasConds not implemented")
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

func _Middleware_GetOutOfGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutOfGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOutOfGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.middleware.outofgas.v1.Middleware/GetOutOfGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOutOfGas(ctx, req.(*GetOutOfGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOutOfGases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutOfGasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOutOfGases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.middleware.outofgas.v1.Middleware/GetOutOfGases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOutOfGases(ctx, req.(*GetOutOfGasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOutOfGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOutOfGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOutOfGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.middleware.outofgas.v1.Middleware/ExistOutOfGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOutOfGas(ctx, req.(*ExistOutOfGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistOutOfGasConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistOutOfGasCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistOutOfGasConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.middleware.outofgas.v1.Middleware/ExistOutOfGasConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistOutOfGasConds(ctx, req.(*ExistOutOfGasCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.outofgas.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOutOfGas",
			Handler:    _Middleware_GetOutOfGas_Handler,
		},
		{
			MethodName: "GetOutOfGases",
			Handler:    _Middleware_GetOutOfGases_Handler,
		},
		{
			MethodName: "ExistOutOfGas",
			Handler:    _Middleware_ExistOutOfGas_Handler,
		},
		{
			MethodName: "ExistOutOfGasConds",
			Handler:    _Middleware_ExistOutOfGasConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/outofgas/outofgas.proto",
}
