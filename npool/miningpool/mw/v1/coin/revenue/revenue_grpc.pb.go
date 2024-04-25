// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/coin/revenue/revenue.proto

package revenue

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
	Middleware_CreateRevenueType_FullMethodName     = "/miningpool.middleware.coin.revenue.v1.Middleware/CreateRevenueType"
	Middleware_GetRevenueType_FullMethodName        = "/miningpool.middleware.coin.revenue.v1.Middleware/GetRevenueType"
	Middleware_ExistRevenueType_FullMethodName      = "/miningpool.middleware.coin.revenue.v1.Middleware/ExistRevenueType"
	Middleware_GetRevenueTypes_FullMethodName       = "/miningpool.middleware.coin.revenue.v1.Middleware/GetRevenueTypes"
	Middleware_ExistRevenueTypeConds_FullMethodName = "/miningpool.middleware.coin.revenue.v1.Middleware/ExistRevenueTypeConds"
	Middleware_UpdateRevenueType_FullMethodName     = "/miningpool.middleware.coin.revenue.v1.Middleware/UpdateRevenueType"
	Middleware_DeleteRevenueType_FullMethodName     = "/miningpool.middleware.coin.revenue.v1.Middleware/DeleteRevenueType"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateRevenueType(ctx context.Context, in *CreateRevenueTypeRequest, opts ...grpc.CallOption) (*CreateRevenueTypeResponse, error)
	GetRevenueType(ctx context.Context, in *GetRevenueTypeRequest, opts ...grpc.CallOption) (*GetRevenueTypeResponse, error)
	ExistRevenueType(ctx context.Context, in *ExistRevenueTypeRequest, opts ...grpc.CallOption) (*ExistRevenueTypeResponse, error)
	GetRevenueTypes(ctx context.Context, in *GetRevenueTypesRequest, opts ...grpc.CallOption) (*GetRevenueTypesResponse, error)
	ExistRevenueTypeConds(ctx context.Context, in *ExistRevenueTypeCondsRequest, opts ...grpc.CallOption) (*ExistRevenueTypeCondsResponse, error)
	UpdateRevenueType(ctx context.Context, in *UpdateRevenueTypeRequest, opts ...grpc.CallOption) (*UpdateRevenueTypeResponse, error)
	DeleteRevenueType(ctx context.Context, in *DeleteRevenueTypeRequest, opts ...grpc.CallOption) (*DeleteRevenueTypeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateRevenueType(ctx context.Context, in *CreateRevenueTypeRequest, opts ...grpc.CallOption) (*CreateRevenueTypeResponse, error) {
	out := new(CreateRevenueTypeResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateRevenueType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRevenueType(ctx context.Context, in *GetRevenueTypeRequest, opts ...grpc.CallOption) (*GetRevenueTypeResponse, error) {
	out := new(GetRevenueTypeResponse)
	err := c.cc.Invoke(ctx, Middleware_GetRevenueType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistRevenueType(ctx context.Context, in *ExistRevenueTypeRequest, opts ...grpc.CallOption) (*ExistRevenueTypeResponse, error) {
	out := new(ExistRevenueTypeResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistRevenueType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRevenueTypes(ctx context.Context, in *GetRevenueTypesRequest, opts ...grpc.CallOption) (*GetRevenueTypesResponse, error) {
	out := new(GetRevenueTypesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetRevenueTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistRevenueTypeConds(ctx context.Context, in *ExistRevenueTypeCondsRequest, opts ...grpc.CallOption) (*ExistRevenueTypeCondsResponse, error) {
	out := new(ExistRevenueTypeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistRevenueTypeConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateRevenueType(ctx context.Context, in *UpdateRevenueTypeRequest, opts ...grpc.CallOption) (*UpdateRevenueTypeResponse, error) {
	out := new(UpdateRevenueTypeResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateRevenueType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteRevenueType(ctx context.Context, in *DeleteRevenueTypeRequest, opts ...grpc.CallOption) (*DeleteRevenueTypeResponse, error) {
	out := new(DeleteRevenueTypeResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteRevenueType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateRevenueType(context.Context, *CreateRevenueTypeRequest) (*CreateRevenueTypeResponse, error)
	GetRevenueType(context.Context, *GetRevenueTypeRequest) (*GetRevenueTypeResponse, error)
	ExistRevenueType(context.Context, *ExistRevenueTypeRequest) (*ExistRevenueTypeResponse, error)
	GetRevenueTypes(context.Context, *GetRevenueTypesRequest) (*GetRevenueTypesResponse, error)
	ExistRevenueTypeConds(context.Context, *ExistRevenueTypeCondsRequest) (*ExistRevenueTypeCondsResponse, error)
	UpdateRevenueType(context.Context, *UpdateRevenueTypeRequest) (*UpdateRevenueTypeResponse, error)
	DeleteRevenueType(context.Context, *DeleteRevenueTypeRequest) (*DeleteRevenueTypeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateRevenueType(context.Context, *CreateRevenueTypeRequest) (*CreateRevenueTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRevenueType not implemented")
}
func (UnimplementedMiddlewareServer) GetRevenueType(context.Context, *GetRevenueTypeRequest) (*GetRevenueTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevenueType not implemented")
}
func (UnimplementedMiddlewareServer) ExistRevenueType(context.Context, *ExistRevenueTypeRequest) (*ExistRevenueTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistRevenueType not implemented")
}
func (UnimplementedMiddlewareServer) GetRevenueTypes(context.Context, *GetRevenueTypesRequest) (*GetRevenueTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevenueTypes not implemented")
}
func (UnimplementedMiddlewareServer) ExistRevenueTypeConds(context.Context, *ExistRevenueTypeCondsRequest) (*ExistRevenueTypeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistRevenueTypeConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateRevenueType(context.Context, *UpdateRevenueTypeRequest) (*UpdateRevenueTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRevenueType not implemented")
}
func (UnimplementedMiddlewareServer) DeleteRevenueType(context.Context, *DeleteRevenueTypeRequest) (*DeleteRevenueTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRevenueType not implemented")
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

func _Middleware_CreateRevenueType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRevenueTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateRevenueType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateRevenueType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateRevenueType(ctx, req.(*CreateRevenueTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRevenueType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRevenueTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRevenueType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetRevenueType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRevenueType(ctx, req.(*GetRevenueTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistRevenueType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRevenueTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistRevenueType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistRevenueType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistRevenueType(ctx, req.(*ExistRevenueTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRevenueTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRevenueTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRevenueTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetRevenueTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRevenueTypes(ctx, req.(*GetRevenueTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistRevenueTypeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRevenueTypeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistRevenueTypeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistRevenueTypeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistRevenueTypeConds(ctx, req.(*ExistRevenueTypeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateRevenueType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRevenueTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateRevenueType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateRevenueType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateRevenueType(ctx, req.(*UpdateRevenueTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteRevenueType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRevenueTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteRevenueType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteRevenueType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteRevenueType(ctx, req.(*DeleteRevenueTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.coin.revenue.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRevenueType",
			Handler:    _Middleware_CreateRevenueType_Handler,
		},
		{
			MethodName: "GetRevenueType",
			Handler:    _Middleware_GetRevenueType_Handler,
		},
		{
			MethodName: "ExistRevenueType",
			Handler:    _Middleware_ExistRevenueType_Handler,
		},
		{
			MethodName: "GetRevenueTypes",
			Handler:    _Middleware_GetRevenueTypes_Handler,
		},
		{
			MethodName: "ExistRevenueTypeConds",
			Handler:    _Middleware_ExistRevenueTypeConds_Handler,
		},
		{
			MethodName: "UpdateRevenueType",
			Handler:    _Middleware_UpdateRevenueType_Handler,
		},
		{
			MethodName: "DeleteRevenueType",
			Handler:    _Middleware_DeleteRevenueType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/coin/revenue/revenue.proto",
}
