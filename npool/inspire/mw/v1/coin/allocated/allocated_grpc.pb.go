// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/coin/allocated/allocated.proto

package allocated

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
	Middleware_CreateCoinAllocated_FullMethodName     = "/inspire.middleware.coin.allocated.v1.Middleware/CreateCoinAllocated"
	Middleware_UpdateCoinAllocated_FullMethodName     = "/inspire.middleware.coin.allocated.v1.Middleware/UpdateCoinAllocated"
	Middleware_GetCoinAllocated_FullMethodName        = "/inspire.middleware.coin.allocated.v1.Middleware/GetCoinAllocated"
	Middleware_GetCoinAllocateds_FullMethodName       = "/inspire.middleware.coin.allocated.v1.Middleware/GetCoinAllocateds"
	Middleware_ExistCoinAllocatedConds_FullMethodName = "/inspire.middleware.coin.allocated.v1.Middleware/ExistCoinAllocatedConds"
	Middleware_DeleteCoinAllocated_FullMethodName     = "/inspire.middleware.coin.allocated.v1.Middleware/DeleteCoinAllocated"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCoinAllocated(ctx context.Context, in *CreateCoinAllocatedRequest, opts ...grpc.CallOption) (*CreateCoinAllocatedResponse, error)
	UpdateCoinAllocated(ctx context.Context, in *UpdateCoinAllocatedRequest, opts ...grpc.CallOption) (*UpdateCoinAllocatedResponse, error)
	GetCoinAllocated(ctx context.Context, in *GetCoinAllocatedRequest, opts ...grpc.CallOption) (*GetCoinAllocatedResponse, error)
	GetCoinAllocateds(ctx context.Context, in *GetCoinAllocatedsRequest, opts ...grpc.CallOption) (*GetCoinAllocatedsResponse, error)
	ExistCoinAllocatedConds(ctx context.Context, in *ExistCoinAllocatedCondsRequest, opts ...grpc.CallOption) (*ExistCoinAllocatedCondsResponse, error)
	DeleteCoinAllocated(ctx context.Context, in *DeleteCoinAllocatedRequest, opts ...grpc.CallOption) (*DeleteCoinAllocatedResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCoinAllocated(ctx context.Context, in *CreateCoinAllocatedRequest, opts ...grpc.CallOption) (*CreateCoinAllocatedResponse, error) {
	out := new(CreateCoinAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCoinAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCoinAllocated(ctx context.Context, in *UpdateCoinAllocatedRequest, opts ...grpc.CallOption) (*UpdateCoinAllocatedResponse, error) {
	out := new(UpdateCoinAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCoinAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinAllocated(ctx context.Context, in *GetCoinAllocatedRequest, opts ...grpc.CallOption) (*GetCoinAllocatedResponse, error) {
	out := new(GetCoinAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinAllocateds(ctx context.Context, in *GetCoinAllocatedsRequest, opts ...grpc.CallOption) (*GetCoinAllocatedsResponse, error) {
	out := new(GetCoinAllocatedsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinAllocateds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCoinAllocatedConds(ctx context.Context, in *ExistCoinAllocatedCondsRequest, opts ...grpc.CallOption) (*ExistCoinAllocatedCondsResponse, error) {
	out := new(ExistCoinAllocatedCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCoinAllocatedConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCoinAllocated(ctx context.Context, in *DeleteCoinAllocatedRequest, opts ...grpc.CallOption) (*DeleteCoinAllocatedResponse, error) {
	out := new(DeleteCoinAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCoinAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCoinAllocated(context.Context, *CreateCoinAllocatedRequest) (*CreateCoinAllocatedResponse, error)
	UpdateCoinAllocated(context.Context, *UpdateCoinAllocatedRequest) (*UpdateCoinAllocatedResponse, error)
	GetCoinAllocated(context.Context, *GetCoinAllocatedRequest) (*GetCoinAllocatedResponse, error)
	GetCoinAllocateds(context.Context, *GetCoinAllocatedsRequest) (*GetCoinAllocatedsResponse, error)
	ExistCoinAllocatedConds(context.Context, *ExistCoinAllocatedCondsRequest) (*ExistCoinAllocatedCondsResponse, error)
	DeleteCoinAllocated(context.Context, *DeleteCoinAllocatedRequest) (*DeleteCoinAllocatedResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCoinAllocated(context.Context, *CreateCoinAllocatedRequest) (*CreateCoinAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinAllocated not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCoinAllocated(context.Context, *UpdateCoinAllocatedRequest) (*UpdateCoinAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinAllocated not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinAllocated(context.Context, *GetCoinAllocatedRequest) (*GetCoinAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinAllocated not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinAllocateds(context.Context, *GetCoinAllocatedsRequest) (*GetCoinAllocatedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinAllocateds not implemented")
}
func (UnimplementedMiddlewareServer) ExistCoinAllocatedConds(context.Context, *ExistCoinAllocatedCondsRequest) (*ExistCoinAllocatedCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoinAllocatedConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCoinAllocated(context.Context, *DeleteCoinAllocatedRequest) (*DeleteCoinAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoinAllocated not implemented")
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

func _Middleware_CreateCoinAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCoinAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCoinAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCoinAllocated(ctx, req.(*CreateCoinAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCoinAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCoinAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCoinAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCoinAllocated(ctx, req.(*UpdateCoinAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinAllocated(ctx, req.(*GetCoinAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinAllocateds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinAllocatedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinAllocateds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinAllocateds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinAllocateds(ctx, req.(*GetCoinAllocatedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCoinAllocatedConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinAllocatedCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCoinAllocatedConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCoinAllocatedConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCoinAllocatedConds(ctx, req.(*ExistCoinAllocatedCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCoinAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCoinAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCoinAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCoinAllocated(ctx, req.(*DeleteCoinAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.coin.allocated.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinAllocated",
			Handler:    _Middleware_CreateCoinAllocated_Handler,
		},
		{
			MethodName: "UpdateCoinAllocated",
			Handler:    _Middleware_UpdateCoinAllocated_Handler,
		},
		{
			MethodName: "GetCoinAllocated",
			Handler:    _Middleware_GetCoinAllocated_Handler,
		},
		{
			MethodName: "GetCoinAllocateds",
			Handler:    _Middleware_GetCoinAllocateds_Handler,
		},
		{
			MethodName: "ExistCoinAllocatedConds",
			Handler:    _Middleware_ExistCoinAllocatedConds_Handler,
		},
		{
			MethodName: "DeleteCoinAllocated",
			Handler:    _Middleware_DeleteCoinAllocated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/coin/allocated/allocated.proto",
}