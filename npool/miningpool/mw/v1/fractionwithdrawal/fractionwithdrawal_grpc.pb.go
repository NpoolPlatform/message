// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/fractionwithdrawal/fractionwithdrawal.proto

package fractionwithdrawal

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
	CreateFractionWithdrawal(ctx context.Context, in *CreateFractionWithdrawalRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalResponse, error)
	GetFractionWithdrawal(ctx context.Context, in *GetFractionWithdrawalRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalResponse, error)
	GetFractionWithdrawals(ctx context.Context, in *GetFractionWithdrawalsRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalsResponse, error)
	ExistFractionWithdrawal(ctx context.Context, in *ExistFractionWithdrawalRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalResponse, error)
	ExistFractionWithdrawalConds(ctx context.Context, in *ExistFractionWithdrawalCondsRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalCondsResponse, error)
	UpdateFractionWithdrawal(ctx context.Context, in *UpdateFractionWithdrawalRequest, opts ...grpc.CallOption) (*UpdateFractionWithdrawalResponse, error)
	DeleteFractionWithdrawal(ctx context.Context, in *DeleteFractionWithdrawalRequest, opts ...grpc.CallOption) (*DeleteFractionWithdrawalResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFractionWithdrawal(ctx context.Context, in *CreateFractionWithdrawalRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalResponse, error) {
	out := new(CreateFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/CreateFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionWithdrawal(ctx context.Context, in *GetFractionWithdrawalRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalResponse, error) {
	out := new(GetFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/GetFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionWithdrawals(ctx context.Context, in *GetFractionWithdrawalsRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalsResponse, error) {
	out := new(GetFractionWithdrawalsResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/GetFractionWithdrawals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionWithdrawal(ctx context.Context, in *ExistFractionWithdrawalRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalResponse, error) {
	out := new(ExistFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/ExistFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionWithdrawalConds(ctx context.Context, in *ExistFractionWithdrawalCondsRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalCondsResponse, error) {
	out := new(ExistFractionWithdrawalCondsResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/ExistFractionWithdrawalConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateFractionWithdrawal(ctx context.Context, in *UpdateFractionWithdrawalRequest, opts ...grpc.CallOption) (*UpdateFractionWithdrawalResponse, error) {
	out := new(UpdateFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/UpdateFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFractionWithdrawal(ctx context.Context, in *DeleteFractionWithdrawalRequest, opts ...grpc.CallOption) (*DeleteFractionWithdrawalResponse, error) {
	out := new(DeleteFractionWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/miningpool.middleware.fractionwithdrawal.v1.Middleware/DeleteFractionWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFractionWithdrawal(context.Context, *CreateFractionWithdrawalRequest) (*CreateFractionWithdrawalResponse, error)
	GetFractionWithdrawal(context.Context, *GetFractionWithdrawalRequest) (*GetFractionWithdrawalResponse, error)
	GetFractionWithdrawals(context.Context, *GetFractionWithdrawalsRequest) (*GetFractionWithdrawalsResponse, error)
	ExistFractionWithdrawal(context.Context, *ExistFractionWithdrawalRequest) (*ExistFractionWithdrawalResponse, error)
	ExistFractionWithdrawalConds(context.Context, *ExistFractionWithdrawalCondsRequest) (*ExistFractionWithdrawalCondsResponse, error)
	UpdateFractionWithdrawal(context.Context, *UpdateFractionWithdrawalRequest) (*UpdateFractionWithdrawalResponse, error)
	DeleteFractionWithdrawal(context.Context, *DeleteFractionWithdrawalRequest) (*DeleteFractionWithdrawalResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFractionWithdrawal(context.Context, *CreateFractionWithdrawalRequest) (*CreateFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFractionWithdrawal not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionWithdrawal(context.Context, *GetFractionWithdrawalRequest) (*GetFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionWithdrawal not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionWithdrawals(context.Context, *GetFractionWithdrawalsRequest) (*GetFractionWithdrawalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionWithdrawals not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionWithdrawal(context.Context, *ExistFractionWithdrawalRequest) (*ExistFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionWithdrawal not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionWithdrawalConds(context.Context, *ExistFractionWithdrawalCondsRequest) (*ExistFractionWithdrawalCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionWithdrawalConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateFractionWithdrawal(context.Context, *UpdateFractionWithdrawalRequest) (*UpdateFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFractionWithdrawal not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFractionWithdrawal(context.Context, *DeleteFractionWithdrawalRequest) (*DeleteFractionWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFractionWithdrawal not implemented")
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

func _Middleware_CreateFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/CreateFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFractionWithdrawal(ctx, req.(*CreateFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/GetFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionWithdrawal(ctx, req.(*GetFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionWithdrawals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionWithdrawalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionWithdrawals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/GetFractionWithdrawals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionWithdrawals(ctx, req.(*GetFractionWithdrawalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/ExistFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionWithdrawal(ctx, req.(*ExistFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionWithdrawalConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionWithdrawalCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionWithdrawalConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/ExistFractionWithdrawalConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionWithdrawalConds(ctx, req.(*ExistFractionWithdrawalCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/UpdateFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateFractionWithdrawal(ctx, req.(*UpdateFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFractionWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFractionWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFractionWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miningpool.middleware.fractionwithdrawal.v1.Middleware/DeleteFractionWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFractionWithdrawal(ctx, req.(*DeleteFractionWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.fractionwithdrawal.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFractionWithdrawal",
			Handler:    _Middleware_CreateFractionWithdrawal_Handler,
		},
		{
			MethodName: "GetFractionWithdrawal",
			Handler:    _Middleware_GetFractionWithdrawal_Handler,
		},
		{
			MethodName: "GetFractionWithdrawals",
			Handler:    _Middleware_GetFractionWithdrawals_Handler,
		},
		{
			MethodName: "ExistFractionWithdrawal",
			Handler:    _Middleware_ExistFractionWithdrawal_Handler,
		},
		{
			MethodName: "ExistFractionWithdrawalConds",
			Handler:    _Middleware_ExistFractionWithdrawalConds_Handler,
		},
		{
			MethodName: "UpdateFractionWithdrawal",
			Handler:    _Middleware_UpdateFractionWithdrawal_Handler,
		},
		{
			MethodName: "DeleteFractionWithdrawal",
			Handler:    _Middleware_DeleteFractionWithdrawal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/fractionwithdrawal/fractionwithdrawal.proto",
}
