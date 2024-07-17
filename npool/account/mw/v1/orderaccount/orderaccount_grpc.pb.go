// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/account/mw/v1/orderaccount/orderaccount.proto

package orderaccount

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
	Middleware_GetOrderAccount_FullMethodName  = "/account.middleware.orderaccount.v1.Middleware/GetOrderAccount"
	Middleware_GetOrderAccounts_FullMethodName = "/account.middleware.orderaccount.v1.Middleware/GetOrderAccounts"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetOrderAccount(ctx context.Context, in *GetOrderAccountRequest, opts ...grpc.CallOption) (*GetOrderAccountResponse, error)
	GetOrderAccounts(ctx context.Context, in *GetOrderAccountsRequest, opts ...grpc.CallOption) (*GetOrderAccountsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetOrderAccount(ctx context.Context, in *GetOrderAccountRequest, opts ...grpc.CallOption) (*GetOrderAccountResponse, error) {
	out := new(GetOrderAccountResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetOrderAccounts(ctx context.Context, in *GetOrderAccountsRequest, opts ...grpc.CallOption) (*GetOrderAccountsResponse, error) {
	out := new(GetOrderAccountsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetOrderAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetOrderAccount(context.Context, *GetOrderAccountRequest) (*GetOrderAccountResponse, error)
	GetOrderAccounts(context.Context, *GetOrderAccountsRequest) (*GetOrderAccountsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetOrderAccount(context.Context, *GetOrderAccountRequest) (*GetOrderAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderAccount not implemented")
}
func (UnimplementedMiddlewareServer) GetOrderAccounts(context.Context, *GetOrderAccountsRequest) (*GetOrderAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderAccounts not implemented")
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

func _Middleware_GetOrderAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderAccount(ctx, req.(*GetOrderAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetOrderAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetOrderAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetOrderAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetOrderAccounts(ctx, req.(*GetOrderAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.middleware.orderaccount.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderAccount",
			Handler:    _Middleware_GetOrderAccount_Handler,
		},
		{
			MethodName: "GetOrderAccounts",
			Handler:    _Middleware_GetOrderAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/account/mw/v1/orderaccount/orderaccount.proto",
}
