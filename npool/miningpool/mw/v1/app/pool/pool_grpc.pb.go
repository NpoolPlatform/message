// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/app/pool/pool.proto

package pool

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
	Middleware_CreatePool_FullMethodName     = "/miningpool.middleware.pool.v1.Middleware/CreatePool"
	Middleware_GetPool_FullMethodName        = "/miningpool.middleware.pool.v1.Middleware/GetPool"
	Middleware_ExistPool_FullMethodName      = "/miningpool.middleware.pool.v1.Middleware/ExistPool"
	Middleware_GetPools_FullMethodName       = "/miningpool.middleware.pool.v1.Middleware/GetPools"
	Middleware_ExistPoolConds_FullMethodName = "/miningpool.middleware.pool.v1.Middleware/ExistPoolConds"
	Middleware_UpdatePool_FullMethodName     = "/miningpool.middleware.pool.v1.Middleware/UpdatePool"
	Middleware_DeletePool_FullMethodName     = "/miningpool.middleware.pool.v1.Middleware/DeletePool"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error)
	GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error)
	ExistPool(ctx context.Context, in *ExistPoolRequest, opts ...grpc.CallOption) (*ExistPoolResponse, error)
	GetPools(ctx context.Context, in *GetPoolsRequest, opts ...grpc.CallOption) (*GetPoolsResponse, error)
	ExistPoolConds(ctx context.Context, in *ExistPoolCondsRequest, opts ...grpc.CallOption) (*ExistPoolCondsResponse, error)
	UpdatePool(ctx context.Context, in *UpdatePoolRequest, opts ...grpc.CallOption) (*UpdatePoolResponse, error)
	DeletePool(ctx context.Context, in *DeletePoolRequest, opts ...grpc.CallOption) (*DeletePoolResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error) {
	out := new(CreatePoolResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error) {
	out := new(GetPoolResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPool(ctx context.Context, in *ExistPoolRequest, opts ...grpc.CallOption) (*ExistPoolResponse, error) {
	out := new(ExistPoolResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPools(ctx context.Context, in *GetPoolsRequest, opts ...grpc.CallOption) (*GetPoolsResponse, error) {
	out := new(GetPoolsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistPoolConds(ctx context.Context, in *ExistPoolCondsRequest, opts ...grpc.CallOption) (*ExistPoolCondsResponse, error) {
	out := new(ExistPoolCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistPoolConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePool(ctx context.Context, in *UpdatePoolRequest, opts ...grpc.CallOption) (*UpdatePoolResponse, error) {
	out := new(UpdatePoolResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeletePool(ctx context.Context, in *DeletePoolRequest, opts ...grpc.CallOption) (*DeletePoolResponse, error) {
	out := new(DeletePoolResponse)
	err := c.cc.Invoke(ctx, Middleware_DeletePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error)
	GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error)
	ExistPool(context.Context, *ExistPoolRequest) (*ExistPoolResponse, error)
	GetPools(context.Context, *GetPoolsRequest) (*GetPoolsResponse, error)
	ExistPoolConds(context.Context, *ExistPoolCondsRequest) (*ExistPoolCondsResponse, error)
	UpdatePool(context.Context, *UpdatePoolRequest) (*UpdatePoolResponse, error)
	DeletePool(context.Context, *DeletePoolRequest) (*DeletePoolResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedMiddlewareServer) GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPool not implemented")
}
func (UnimplementedMiddlewareServer) ExistPool(context.Context, *ExistPoolRequest) (*ExistPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPool not implemented")
}
func (UnimplementedMiddlewareServer) GetPools(context.Context, *GetPoolsRequest) (*GetPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPools not implemented")
}
func (UnimplementedMiddlewareServer) ExistPoolConds(context.Context, *ExistPoolCondsRequest) (*ExistPoolCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPoolConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePool(context.Context, *UpdatePoolRequest) (*UpdatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePool not implemented")
}
func (UnimplementedMiddlewareServer) DeletePool(context.Context, *DeletePoolRequest) (*DeletePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePool not implemented")
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

func _Middleware_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPool(ctx, req.(*GetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPool(ctx, req.(*ExistPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPools(ctx, req.(*GetPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistPoolConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPoolCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistPoolConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistPoolConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistPoolConds(ctx, req.(*ExistPoolCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePool(ctx, req.(*UpdatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeletePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeletePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeletePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeletePool(ctx, req.(*DeletePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.pool.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Middleware_CreatePool_Handler,
		},
		{
			MethodName: "GetPool",
			Handler:    _Middleware_GetPool_Handler,
		},
		{
			MethodName: "ExistPool",
			Handler:    _Middleware_ExistPool_Handler,
		},
		{
			MethodName: "GetPools",
			Handler:    _Middleware_GetPools_Handler,
		},
		{
			MethodName: "ExistPoolConds",
			Handler:    _Middleware_ExistPoolConds_Handler,
		},
		{
			MethodName: "UpdatePool",
			Handler:    _Middleware_UpdatePool_Handler,
		},
		{
			MethodName: "DeletePool",
			Handler:    _Middleware_DeletePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/app/pool/pool.proto",
}
