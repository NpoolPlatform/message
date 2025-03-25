// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/order/lock/lock.proto

package lock

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
	Middleware_GetOrderLock_FullMethodName     = "/order.middleware.order1.lock.v1.Middleware/GetOrderLock"
	Middleware_GetOrderLocks_FullMethodName    = "/order.middleware.order1.lock.v1.Middleware/GetOrderLocks"
	Middleware_CreateOrderLocks_FullMethodName = "/order.middleware.order1.lock.v1.Middleware/CreateOrderLocks"
	Middleware_DeleteOrderLocks_FullMethodName = "/order.middleware.order1.lock.v1.Middleware/DeleteOrderLocks"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetOrderLock(ctx context.Context, in *GetOrderLockRequest, opts ...grpc.CallOption) (*GetOrderLockResponse, error)
	GetOrderLocks(ctx context.Context, in *GetOrderLocksRequest, opts ...grpc.CallOption) (*GetOrderLocksResponse, error)
	CreateOrderLocks(ctx context.Context, in *CreateOrderLocksRequest, opts ...grpc.CallOption) (*CreateOrderLocksResponse, error)
	DeleteOrderLocks(ctx context.Context, in *DeleteOrderLocksRequest, opts ...grpc.CallOption) (*DeleteOrderLocksResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetOrderLock(ctx context.Context, in *GetOrderLockRequest, opts ...grpc.CallOption) (*GetOrderLockResponse, error) {
	out := new(GetOrderLockResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderLock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOrderLocks(ctx context.Context, in *GetOrderLocksRequest, opts ...grpc.CallOption) (*GetOrderLocksResponse, error) {
	out := new(GetOrderLocksResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderLocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateOrderLocks(ctx context.Context, in *CreateOrderLocksRequest, opts ...grpc.CallOption) (*CreateOrderLocksResponse, error) {
	out := new(CreateOrderLocksResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateOrderLocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteOrderLocks(ctx context.Context, in *DeleteOrderLocksRequest, opts ...grpc.CallOption) (*DeleteOrderLocksResponse, error) {
	out := new(DeleteOrderLocksResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteOrderLocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetOrderLock(context.Context, *GetOrderLockRequest) (*GetOrderLockResponse, error)
	GetOrderLocks(context.Context, *GetOrderLocksRequest) (*GetOrderLocksResponse, error)
	CreateOrderLocks(context.Context, *CreateOrderLocksRequest) (*CreateOrderLocksResponse, error)
	DeleteOrderLocks(context.Context, *DeleteOrderLocksRequest) (*DeleteOrderLocksResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetOrderLock(context.Context, *GetOrderLockRequest) (*GetOrderLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderLock not implemented")
}
func (UnimplementedMiddlewareServer) GetOrderLocks(context.Context, *GetOrderLocksRequest) (*GetOrderLocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderLocks not implemented")
}
func (UnimplementedMiddlewareServer) CreateOrderLocks(context.Context, *CreateOrderLocksRequest) (*CreateOrderLocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderLocks not implemented")
}
func (UnimplementedMiddlewareServer) DeleteOrderLocks(context.Context, *DeleteOrderLocksRequest) (*DeleteOrderLocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderLocks not implemented")
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

func _Middleware_GetOrderLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderLock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderLock(ctx, req.(*GetOrderLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOrderLocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderLocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderLocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderLocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderLocks(ctx, req.(*GetOrderLocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateOrderLocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderLocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateOrderLocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateOrderLocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateOrderLocks(ctx, req.(*CreateOrderLocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteOrderLocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderLocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteOrderLocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteOrderLocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteOrderLocks(ctx, req.(*DeleteOrderLocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.order1.lock.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderLock",
			Handler:    _Middleware_GetOrderLock_Handler,
		},
		{
			MethodName: "GetOrderLocks",
			Handler:    _Middleware_GetOrderLocks_Handler,
		},
		{
			MethodName: "CreateOrderLocks",
			Handler:    _Middleware_CreateOrderLocks_Handler,
		},
		{
			MethodName: "DeleteOrderLocks",
			Handler:    _Middleware_DeleteOrderLocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/order/lock/lock.proto",
}
