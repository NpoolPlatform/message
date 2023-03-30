// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/appdefaultgood/appdefaultgood.proto

package appdefaultgood

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
	GetAppDefaultGood(ctx context.Context, in *GetAppDefaultGoodRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodResponse, error)
	GetAppDefaultGoodOnly(ctx context.Context, in *GetAppDefaultGoodOnlyRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodOnlyResponse, error)
	GetAppDefaultGoods(ctx context.Context, in *GetAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodsResponse, error)
	DeleteAppDefaultGood(ctx context.Context, in *DeleteAppDefaultGoodRequest, opts ...grpc.CallOption) (*DeleteAppDefaultGoodResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetAppDefaultGood(ctx context.Context, in *GetAppDefaultGoodRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodResponse, error) {
	out := new(GetAppDefaultGoodResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppDefaultGoodOnly(ctx context.Context, in *GetAppDefaultGoodOnlyRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodOnlyResponse, error) {
	out := new(GetAppDefaultGoodOnlyResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGoodOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAppDefaultGoods(ctx context.Context, in *GetAppDefaultGoodsRequest, opts ...grpc.CallOption) (*GetAppDefaultGoodsResponse, error) {
	out := new(GetAppDefaultGoodsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAppDefaultGood(ctx context.Context, in *DeleteAppDefaultGoodRequest, opts ...grpc.CallOption) (*DeleteAppDefaultGoodResponse, error) {
	out := new(DeleteAppDefaultGoodResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.appdefaultgood.v1.Middleware/DeleteAppDefaultGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetAppDefaultGood(context.Context, *GetAppDefaultGoodRequest) (*GetAppDefaultGoodResponse, error)
	GetAppDefaultGoodOnly(context.Context, *GetAppDefaultGoodOnlyRequest) (*GetAppDefaultGoodOnlyResponse, error)
	GetAppDefaultGoods(context.Context, *GetAppDefaultGoodsRequest) (*GetAppDefaultGoodsResponse, error)
	DeleteAppDefaultGood(context.Context, *DeleteAppDefaultGoodRequest) (*DeleteAppDefaultGoodResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetAppDefaultGood(context.Context, *GetAppDefaultGoodRequest) (*GetAppDefaultGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDefaultGood not implemented")
}
func (UnimplementedMiddlewareServer) GetAppDefaultGoodOnly(context.Context, *GetAppDefaultGoodOnlyRequest) (*GetAppDefaultGoodOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDefaultGoodOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetAppDefaultGoods(context.Context, *GetAppDefaultGoodsRequest) (*GetAppDefaultGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppDefaultGoods not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAppDefaultGood(context.Context, *DeleteAppDefaultGoodRequest) (*DeleteAppDefaultGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppDefaultGood not implemented")
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

func _Middleware_GetAppDefaultGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDefaultGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppDefaultGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppDefaultGood(ctx, req.(*GetAppDefaultGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppDefaultGoodOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDefaultGoodOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppDefaultGoodOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGoodOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppDefaultGoodOnly(ctx, req.(*GetAppDefaultGoodOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAppDefaultGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppDefaultGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAppDefaultGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.appdefaultgood.v1.Middleware/GetAppDefaultGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAppDefaultGoods(ctx, req.(*GetAppDefaultGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAppDefaultGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppDefaultGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAppDefaultGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.appdefaultgood.v1.Middleware/DeleteAppDefaultGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAppDefaultGood(ctx, req.(*DeleteAppDefaultGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.appdefaultgood.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppDefaultGood",
			Handler:    _Middleware_GetAppDefaultGood_Handler,
		},
		{
			MethodName: "GetAppDefaultGoodOnly",
			Handler:    _Middleware_GetAppDefaultGoodOnly_Handler,
		},
		{
			MethodName: "GetAppDefaultGoods",
			Handler:    _Middleware_GetAppDefaultGoods_Handler,
		},
		{
			MethodName: "DeleteAppDefaultGood",
			Handler:    _Middleware_DeleteAppDefaultGood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/appdefaultgood/appdefaultgood.proto",
}
