// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/fee/fee.proto

package fee

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
	Middleware_CreateFee_FullMethodName     = "/good.middleware.app.fee.v1.Middleware/CreateFee"
	Middleware_UpdateFee_FullMethodName     = "/good.middleware.app.fee.v1.Middleware/UpdateFee"
	Middleware_GetFee_FullMethodName        = "/good.middleware.app.fee.v1.Middleware/GetFee"
	Middleware_GetFees_FullMethodName       = "/good.middleware.app.fee.v1.Middleware/GetFees"
	Middleware_ExistFee_FullMethodName      = "/good.middleware.app.fee.v1.Middleware/ExistFee"
	Middleware_ExistFeeConds_FullMethodName = "/good.middleware.app.fee.v1.Middleware/ExistFeeConds"
	Middleware_DeleteFee_FullMethodName     = "/good.middleware.app.fee.v1.Middleware/DeleteFee"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateFee(ctx context.Context, in *CreateFeeRequest, opts ...grpc.CallOption) (*CreateFeeResponse, error)
	UpdateFee(ctx context.Context, in *UpdateFeeRequest, opts ...grpc.CallOption) (*UpdateFeeResponse, error)
	GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error)
	GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error)
	ExistFee(ctx context.Context, in *ExistFeeRequest, opts ...grpc.CallOption) (*ExistFeeResponse, error)
	ExistFeeConds(ctx context.Context, in *ExistFeeCondsRequest, opts ...grpc.CallOption) (*ExistFeeCondsResponse, error)
	DeleteFee(ctx context.Context, in *DeleteFeeRequest, opts ...grpc.CallOption) (*DeleteFeeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFee(ctx context.Context, in *CreateFeeRequest, opts ...grpc.CallOption) (*CreateFeeResponse, error) {
	out := new(CreateFeeResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateFee(ctx context.Context, in *UpdateFeeRequest, opts ...grpc.CallOption) (*UpdateFeeResponse, error) {
	out := new(UpdateFeeResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error) {
	out := new(GetFeeResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error) {
	out := new(GetFeesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFee(ctx context.Context, in *ExistFeeRequest, opts ...grpc.CallOption) (*ExistFeeResponse, error) {
	out := new(ExistFeeResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFeeConds(ctx context.Context, in *ExistFeeCondsRequest, opts ...grpc.CallOption) (*ExistFeeCondsResponse, error) {
	out := new(ExistFeeCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFeeConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFee(ctx context.Context, in *DeleteFeeRequest, opts ...grpc.CallOption) (*DeleteFeeResponse, error) {
	out := new(DeleteFeeResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFee(context.Context, *CreateFeeRequest) (*CreateFeeResponse, error)
	UpdateFee(context.Context, *UpdateFeeRequest) (*UpdateFeeResponse, error)
	GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error)
	GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error)
	ExistFee(context.Context, *ExistFeeRequest) (*ExistFeeResponse, error)
	ExistFeeConds(context.Context, *ExistFeeCondsRequest) (*ExistFeeCondsResponse, error)
	DeleteFee(context.Context, *DeleteFeeRequest) (*DeleteFeeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFee(context.Context, *CreateFeeRequest) (*CreateFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFee not implemented")
}
func (UnimplementedMiddlewareServer) UpdateFee(context.Context, *UpdateFeeRequest) (*UpdateFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFee not implemented")
}
func (UnimplementedMiddlewareServer) GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedMiddlewareServer) GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFees not implemented")
}
func (UnimplementedMiddlewareServer) ExistFee(context.Context, *ExistFeeRequest) (*ExistFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFee not implemented")
}
func (UnimplementedMiddlewareServer) ExistFeeConds(context.Context, *ExistFeeCondsRequest) (*ExistFeeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFeeConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFee(context.Context, *DeleteFeeRequest) (*DeleteFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFee not implemented")
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

func _Middleware_CreateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFee(ctx, req.(*CreateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateFee(ctx, req.(*UpdateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFee(ctx, req.(*GetFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFees(ctx, req.(*GetFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFee(ctx, req.(*ExistFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFeeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFeeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFeeConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFeeConds(ctx, req.(*ExistFeeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFee(ctx, req.(*DeleteFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.fee.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFee",
			Handler:    _Middleware_CreateFee_Handler,
		},
		{
			MethodName: "UpdateFee",
			Handler:    _Middleware_UpdateFee_Handler,
		},
		{
			MethodName: "GetFee",
			Handler:    _Middleware_GetFee_Handler,
		},
		{
			MethodName: "GetFees",
			Handler:    _Middleware_GetFees_Handler,
		},
		{
			MethodName: "ExistFee",
			Handler:    _Middleware_ExistFee_Handler,
		},
		{
			MethodName: "ExistFeeConds",
			Handler:    _Middleware_ExistFeeConds_Handler,
		},
		{
			MethodName: "DeleteFee",
			Handler:    _Middleware_DeleteFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/fee/fee.proto",
}
