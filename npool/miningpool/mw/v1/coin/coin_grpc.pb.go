// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/coin/coin.proto

package coin

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
	Middleware_CreateCoin_FullMethodName     = "/miningpool.middleware.coin.v1.Middleware/CreateCoin"
	Middleware_GetCoin_FullMethodName        = "/miningpool.middleware.coin.v1.Middleware/GetCoin"
	Middleware_ExistCoin_FullMethodName      = "/miningpool.middleware.coin.v1.Middleware/ExistCoin"
	Middleware_GetCoins_FullMethodName       = "/miningpool.middleware.coin.v1.Middleware/GetCoins"
	Middleware_ExistCoinConds_FullMethodName = "/miningpool.middleware.coin.v1.Middleware/ExistCoinConds"
	Middleware_UpdateCoin_FullMethodName     = "/miningpool.middleware.coin.v1.Middleware/UpdateCoin"
	Middleware_DeleteCoin_FullMethodName     = "/miningpool.middleware.coin.v1.Middleware/DeleteCoin"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error)
	GetCoin(ctx context.Context, in *GetCoinRequest, opts ...grpc.CallOption) (*GetCoinResponse, error)
	ExistCoin(ctx context.Context, in *ExistCoinRequest, opts ...grpc.CallOption) (*ExistCoinResponse, error)
	GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*GetCoinsResponse, error)
	ExistCoinConds(ctx context.Context, in *ExistCoinCondsRequest, opts ...grpc.CallOption) (*ExistCoinCondsResponse, error)
	UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error)
	DeleteCoin(ctx context.Context, in *DeleteCoinRequest, opts ...grpc.CallOption) (*DeleteCoinResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error) {
	out := new(CreateCoinResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoin(ctx context.Context, in *GetCoinRequest, opts ...grpc.CallOption) (*GetCoinResponse, error) {
	out := new(GetCoinResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCoin(ctx context.Context, in *ExistCoinRequest, opts ...grpc.CallOption) (*ExistCoinResponse, error) {
	out := new(ExistCoinResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*GetCoinsResponse, error) {
	out := new(GetCoinsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCoinConds(ctx context.Context, in *ExistCoinCondsRequest, opts ...grpc.CallOption) (*ExistCoinCondsResponse, error) {
	out := new(ExistCoinCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCoinConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error) {
	out := new(UpdateCoinResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCoin(ctx context.Context, in *DeleteCoinRequest, opts ...grpc.CallOption) (*DeleteCoinResponse, error) {
	out := new(DeleteCoinResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error)
	GetCoin(context.Context, *GetCoinRequest) (*GetCoinResponse, error)
	ExistCoin(context.Context, *ExistCoinRequest) (*ExistCoinResponse, error)
	GetCoins(context.Context, *GetCoinsRequest) (*GetCoinsResponse, error)
	ExistCoinConds(context.Context, *ExistCoinCondsRequest) (*ExistCoinCondsResponse, error)
	UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error)
	DeleteCoin(context.Context, *DeleteCoinRequest) (*DeleteCoinResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}
func (UnimplementedMiddlewareServer) GetCoin(context.Context, *GetCoinRequest) (*GetCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoin not implemented")
}
func (UnimplementedMiddlewareServer) ExistCoin(context.Context, *ExistCoinRequest) (*ExistCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoin not implemented")
}
func (UnimplementedMiddlewareServer) GetCoins(context.Context, *GetCoinsRequest) (*GetCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoins not implemented")
}
func (UnimplementedMiddlewareServer) ExistCoinConds(context.Context, *ExistCoinCondsRequest) (*ExistCoinCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoinConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoin not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCoin(context.Context, *DeleteCoinRequest) (*DeleteCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoin not implemented")
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

func _Middleware_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCoin(ctx, req.(*CreateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoin(ctx, req.(*GetCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCoin(ctx, req.(*ExistCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoins(ctx, req.(*GetCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCoinConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCoinConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCoinConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCoinConds(ctx, req.(*ExistCoinCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCoin(ctx, req.(*UpdateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCoin(ctx, req.(*DeleteCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.coin.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoin",
			Handler:    _Middleware_CreateCoin_Handler,
		},
		{
			MethodName: "GetCoin",
			Handler:    _Middleware_GetCoin_Handler,
		},
		{
			MethodName: "ExistCoin",
			Handler:    _Middleware_ExistCoin_Handler,
		},
		{
			MethodName: "GetCoins",
			Handler:    _Middleware_GetCoins_Handler,
		},
		{
			MethodName: "ExistCoinConds",
			Handler:    _Middleware_ExistCoinConds_Handler,
		},
		{
			MethodName: "UpdateCoin",
			Handler:    _Middleware_UpdateCoin_Handler,
		},
		{
			MethodName: "DeleteCoin",
			Handler:    _Middleware_DeleteCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/coin/coin.proto",
}