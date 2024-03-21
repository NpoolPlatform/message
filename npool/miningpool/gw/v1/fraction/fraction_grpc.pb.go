// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/gw/v1/fraction/fraction.proto

package fraction

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
	Middleware_CreateFraction_FullMethodName = "/miningpool.middleware.fraction.v1.Middleware/CreateFraction"
	Middleware_GetFraction_FullMethodName    = "/miningpool.middleware.fraction.v1.Middleware/GetFraction"
	Middleware_GetFractions_FullMethodName   = "/miningpool.middleware.fraction.v1.Middleware/GetFractions"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateFraction(ctx context.Context, in *CreateFractionRequest, opts ...grpc.CallOption) (*CreateFractionResponse, error)
	GetFraction(ctx context.Context, in *GetFractionRequest, opts ...grpc.CallOption) (*GetFractionResponse, error)
	GetFractions(ctx context.Context, in *GetOrderFractionsRequest, opts ...grpc.CallOption) (*GetOrderFractionsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFraction(ctx context.Context, in *CreateFractionRequest, opts ...grpc.CallOption) (*CreateFractionResponse, error) {
	out := new(CreateFractionResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFraction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFraction(ctx context.Context, in *GetFractionRequest, opts ...grpc.CallOption) (*GetFractionResponse, error) {
	out := new(GetFractionResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFraction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractions(ctx context.Context, in *GetOrderFractionsRequest, opts ...grpc.CallOption) (*GetOrderFractionsResponse, error) {
	out := new(GetOrderFractionsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFractions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFraction(context.Context, *CreateFractionRequest) (*CreateFractionResponse, error)
	GetFraction(context.Context, *GetFractionRequest) (*GetFractionResponse, error)
	GetFractions(context.Context, *GetOrderFractionsRequest) (*GetOrderFractionsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFraction(context.Context, *CreateFractionRequest) (*CreateFractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFraction not implemented")
}
func (UnimplementedMiddlewareServer) GetFraction(context.Context, *GetFractionRequest) (*GetFractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFraction not implemented")
}
func (UnimplementedMiddlewareServer) GetFractions(context.Context, *GetOrderFractionsRequest) (*GetOrderFractionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractions not implemented")
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

func _Middleware_CreateFraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFraction(ctx, req.(*CreateFractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFraction(ctx, req.(*GetFractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderFractionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFractions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractions(ctx, req.(*GetOrderFractionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.fraction.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFraction",
			Handler:    _Middleware_CreateFraction_Handler,
		},
		{
			MethodName: "GetFraction",
			Handler:    _Middleware_GetFraction_Handler,
		},
		{
			MethodName: "GetFractions",
			Handler:    _Middleware_GetFractions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/gw/v1/fraction/fraction.proto",
}
