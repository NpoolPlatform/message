// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/good/good.proto

package good

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
	Middleware_GetGood_FullMethodName        = "/good.middleware.good1.v1.Middleware/GetGood"
	Middleware_GetGoods_FullMethodName       = "/good.middleware.good1.v1.Middleware/GetGoods"
	Middleware_ExistGoodConds_FullMethodName = "/good.middleware.good1.v1.Middleware/ExistGoodConds"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	ExistGoodConds(ctx context.Context, in *ExistGoodCondsRequest, opts ...grpc.CallOption) (*ExistGoodCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, Middleware_GetGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistGoodConds(ctx context.Context, in *ExistGoodCondsRequest, opts ...grpc.CallOption) (*ExistGoodCondsResponse, error) {
	out := new(ExistGoodCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistGoodConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	ExistGoodConds(context.Context, *ExistGoodCondsRequest) (*ExistGoodCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedMiddlewareServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedMiddlewareServer) ExistGoodConds(context.Context, *ExistGoodCondsRequest) (*ExistGoodCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistGoodConds not implemented")
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

func _Middleware_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistGoodConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistGoodCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistGoodConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistGoodConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistGoodConds(ctx, req.(*ExistGoodCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.good1.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGood",
			Handler:    _Middleware_GetGood_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Middleware_GetGoods_Handler,
		},
		{
			MethodName: "ExistGoodConds",
			Handler:    _Middleware_ExistGoodConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/good/good.proto",
}
