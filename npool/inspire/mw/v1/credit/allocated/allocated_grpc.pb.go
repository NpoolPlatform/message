// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/credit/allocated/allocated.proto

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
	Middleware_CreateCreditAllocated_FullMethodName     = "/inspire.middleware.credit.allocated.v1.Middleware/CreateCreditAllocated"
	Middleware_GetCreditAllocated_FullMethodName        = "/inspire.middleware.credit.allocated.v1.Middleware/GetCreditAllocated"
	Middleware_GetCreditAllocateds_FullMethodName       = "/inspire.middleware.credit.allocated.v1.Middleware/GetCreditAllocateds"
	Middleware_ExistCreditAllocatedConds_FullMethodName = "/inspire.middleware.credit.allocated.v1.Middleware/ExistCreditAllocatedConds"
	Middleware_DeleteCreditAllocated_FullMethodName     = "/inspire.middleware.credit.allocated.v1.Middleware/DeleteCreditAllocated"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCreditAllocated(ctx context.Context, in *CreateCreditAllocatedRequest, opts ...grpc.CallOption) (*CreateCreditAllocatedResponse, error)
	GetCreditAllocated(ctx context.Context, in *GetCreditAllocatedRequest, opts ...grpc.CallOption) (*GetCreditAllocatedResponse, error)
	GetCreditAllocateds(ctx context.Context, in *GetCreditAllocatedsRequest, opts ...grpc.CallOption) (*GetCreditAllocatedsResponse, error)
	ExistCreditAllocatedConds(ctx context.Context, in *ExistCreditAllocatedCondsRequest, opts ...grpc.CallOption) (*ExistCreditAllocatedCondsResponse, error)
	DeleteCreditAllocated(ctx context.Context, in *DeleteCreditAllocatedRequest, opts ...grpc.CallOption) (*DeleteCreditAllocatedResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCreditAllocated(ctx context.Context, in *CreateCreditAllocatedRequest, opts ...grpc.CallOption) (*CreateCreditAllocatedResponse, error) {
	out := new(CreateCreditAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCreditAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCreditAllocated(ctx context.Context, in *GetCreditAllocatedRequest, opts ...grpc.CallOption) (*GetCreditAllocatedResponse, error) {
	out := new(GetCreditAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCreditAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCreditAllocateds(ctx context.Context, in *GetCreditAllocatedsRequest, opts ...grpc.CallOption) (*GetCreditAllocatedsResponse, error) {
	out := new(GetCreditAllocatedsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCreditAllocateds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCreditAllocatedConds(ctx context.Context, in *ExistCreditAllocatedCondsRequest, opts ...grpc.CallOption) (*ExistCreditAllocatedCondsResponse, error) {
	out := new(ExistCreditAllocatedCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCreditAllocatedConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCreditAllocated(ctx context.Context, in *DeleteCreditAllocatedRequest, opts ...grpc.CallOption) (*DeleteCreditAllocatedResponse, error) {
	out := new(DeleteCreditAllocatedResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCreditAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCreditAllocated(context.Context, *CreateCreditAllocatedRequest) (*CreateCreditAllocatedResponse, error)
	GetCreditAllocated(context.Context, *GetCreditAllocatedRequest) (*GetCreditAllocatedResponse, error)
	GetCreditAllocateds(context.Context, *GetCreditAllocatedsRequest) (*GetCreditAllocatedsResponse, error)
	ExistCreditAllocatedConds(context.Context, *ExistCreditAllocatedCondsRequest) (*ExistCreditAllocatedCondsResponse, error)
	DeleteCreditAllocated(context.Context, *DeleteCreditAllocatedRequest) (*DeleteCreditAllocatedResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCreditAllocated(context.Context, *CreateCreditAllocatedRequest) (*CreateCreditAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditAllocated not implemented")
}
func (UnimplementedMiddlewareServer) GetCreditAllocated(context.Context, *GetCreditAllocatedRequest) (*GetCreditAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditAllocated not implemented")
}
func (UnimplementedMiddlewareServer) GetCreditAllocateds(context.Context, *GetCreditAllocatedsRequest) (*GetCreditAllocatedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditAllocateds not implemented")
}
func (UnimplementedMiddlewareServer) ExistCreditAllocatedConds(context.Context, *ExistCreditAllocatedCondsRequest) (*ExistCreditAllocatedCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCreditAllocatedConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCreditAllocated(context.Context, *DeleteCreditAllocatedRequest) (*DeleteCreditAllocatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreditAllocated not implemented")
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

func _Middleware_CreateCreditAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCreditAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCreditAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCreditAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCreditAllocated(ctx, req.(*CreateCreditAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCreditAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreditAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCreditAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCreditAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCreditAllocated(ctx, req.(*GetCreditAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCreditAllocateds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreditAllocatedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCreditAllocateds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCreditAllocateds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCreditAllocateds(ctx, req.(*GetCreditAllocatedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCreditAllocatedConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCreditAllocatedCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCreditAllocatedConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCreditAllocatedConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCreditAllocatedConds(ctx, req.(*ExistCreditAllocatedCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCreditAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCreditAllocatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCreditAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCreditAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCreditAllocated(ctx, req.(*DeleteCreditAllocatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.credit.allocated.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCreditAllocated",
			Handler:    _Middleware_CreateCreditAllocated_Handler,
		},
		{
			MethodName: "GetCreditAllocated",
			Handler:    _Middleware_GetCreditAllocated_Handler,
		},
		{
			MethodName: "GetCreditAllocateds",
			Handler:    _Middleware_GetCreditAllocateds_Handler,
		},
		{
			MethodName: "ExistCreditAllocatedConds",
			Handler:    _Middleware_ExistCreditAllocatedConds_Handler,
		},
		{
			MethodName: "DeleteCreditAllocated",
			Handler:    _Middleware_DeleteCreditAllocated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/credit/allocated/allocated.proto",
}
