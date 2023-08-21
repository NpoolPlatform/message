// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/good/topmost/good/good.proto

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
	Middleware_CreateTopMostGood_FullMethodName = "/good.middleware.app.good1.topmost.good2.v1.Middleware/CreateTopMostGood"
	Middleware_UpdateTopMostGood_FullMethodName = "/good.middleware.app.good1.topmost.good2.v1.Middleware/UpdateTopMostGood"
	Middleware_GetTopMostGood_FullMethodName    = "/good.middleware.app.good1.topmost.good2.v1.Middleware/GetTopMostGood"
	Middleware_GetTopMostGoods_FullMethodName   = "/good.middleware.app.good1.topmost.good2.v1.Middleware/GetTopMostGoods"
	Middleware_DeleteTopMostGood_FullMethodName = "/good.middleware.app.good1.topmost.good2.v1.Middleware/DeleteTopMostGood"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateTopMostGood(ctx context.Context, in *CreateTopMostGoodRequest, opts ...grpc.CallOption) (*CreateTopMostGoodResponse, error)
	UpdateTopMostGood(ctx context.Context, in *UpdateTopMostGoodRequest, opts ...grpc.CallOption) (*UpdateTopMostGoodResponse, error)
	GetTopMostGood(ctx context.Context, in *GetTopMostGoodRequest, opts ...grpc.CallOption) (*GetTopMostGoodResponse, error)
	GetTopMostGoods(ctx context.Context, in *GetTopMostGoodsRequest, opts ...grpc.CallOption) (*GetTopMostGoodsResponse, error)
	DeleteTopMostGood(ctx context.Context, in *DeleteTopMostGoodRequest, opts ...grpc.CallOption) (*DeleteTopMostGoodResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateTopMostGood(ctx context.Context, in *CreateTopMostGoodRequest, opts ...grpc.CallOption) (*CreateTopMostGoodResponse, error) {
	out := new(CreateTopMostGoodResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateTopMostGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateTopMostGood(ctx context.Context, in *UpdateTopMostGoodRequest, opts ...grpc.CallOption) (*UpdateTopMostGoodResponse, error) {
	out := new(UpdateTopMostGoodResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateTopMostGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTopMostGood(ctx context.Context, in *GetTopMostGoodRequest, opts ...grpc.CallOption) (*GetTopMostGoodResponse, error) {
	out := new(GetTopMostGoodResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTopMostGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTopMostGoods(ctx context.Context, in *GetTopMostGoodsRequest, opts ...grpc.CallOption) (*GetTopMostGoodsResponse, error) {
	out := new(GetTopMostGoodsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTopMostGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteTopMostGood(ctx context.Context, in *DeleteTopMostGoodRequest, opts ...grpc.CallOption) (*DeleteTopMostGoodResponse, error) {
	out := new(DeleteTopMostGoodResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteTopMostGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateTopMostGood(context.Context, *CreateTopMostGoodRequest) (*CreateTopMostGoodResponse, error)
	UpdateTopMostGood(context.Context, *UpdateTopMostGoodRequest) (*UpdateTopMostGoodResponse, error)
	GetTopMostGood(context.Context, *GetTopMostGoodRequest) (*GetTopMostGoodResponse, error)
	GetTopMostGoods(context.Context, *GetTopMostGoodsRequest) (*GetTopMostGoodsResponse, error)
	DeleteTopMostGood(context.Context, *DeleteTopMostGoodRequest) (*DeleteTopMostGoodResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateTopMostGood(context.Context, *CreateTopMostGoodRequest) (*CreateTopMostGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopMostGood not implemented")
}
func (UnimplementedMiddlewareServer) UpdateTopMostGood(context.Context, *UpdateTopMostGoodRequest) (*UpdateTopMostGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopMostGood not implemented")
}
func (UnimplementedMiddlewareServer) GetTopMostGood(context.Context, *GetTopMostGoodRequest) (*GetTopMostGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopMostGood not implemented")
}
func (UnimplementedMiddlewareServer) GetTopMostGoods(context.Context, *GetTopMostGoodsRequest) (*GetTopMostGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopMostGoods not implemented")
}
func (UnimplementedMiddlewareServer) DeleteTopMostGood(context.Context, *DeleteTopMostGoodRequest) (*DeleteTopMostGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopMostGood not implemented")
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

func _Middleware_CreateTopMostGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopMostGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateTopMostGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateTopMostGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateTopMostGood(ctx, req.(*CreateTopMostGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateTopMostGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopMostGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateTopMostGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateTopMostGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateTopMostGood(ctx, req.(*UpdateTopMostGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTopMostGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopMostGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTopMostGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTopMostGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTopMostGood(ctx, req.(*GetTopMostGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTopMostGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopMostGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTopMostGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTopMostGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTopMostGoods(ctx, req.(*GetTopMostGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteTopMostGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopMostGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteTopMostGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteTopMostGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteTopMostGood(ctx, req.(*DeleteTopMostGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.good1.topmost.good2.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopMostGood",
			Handler:    _Middleware_CreateTopMostGood_Handler,
		},
		{
			MethodName: "UpdateTopMostGood",
			Handler:    _Middleware_UpdateTopMostGood_Handler,
		},
		{
			MethodName: "GetTopMostGood",
			Handler:    _Middleware_GetTopMostGood_Handler,
		},
		{
			MethodName: "GetTopMostGoods",
			Handler:    _Middleware_GetTopMostGoods_Handler,
		},
		{
			MethodName: "DeleteTopMostGood",
			Handler:    _Middleware_DeleteTopMostGood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/good/topmost/good/good.proto",
}