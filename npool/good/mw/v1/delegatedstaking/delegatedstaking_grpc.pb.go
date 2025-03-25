// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/delegatedstaking/delegatedstaking.proto

package delegatedstaking

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
	Middleware_CreateDelegatedStaking_FullMethodName     = "/good.middleware.delegatedstaking.v1.Middleware/CreateDelegatedStaking"
	Middleware_UpdateDelegatedStaking_FullMethodName     = "/good.middleware.delegatedstaking.v1.Middleware/UpdateDelegatedStaking"
	Middleware_GetDelegatedStaking_FullMethodName        = "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStaking"
	Middleware_GetDelegatedStakings_FullMethodName       = "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStakings"
	Middleware_ExistDelegatedStakingConds_FullMethodName = "/good.middleware.delegatedstaking.v1.Middleware/ExistDelegatedStakingConds"
	Middleware_DeleteDelegatedStaking_FullMethodName     = "/good.middleware.delegatedstaking.v1.Middleware/DeleteDelegatedStaking"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateDelegatedStaking(ctx context.Context, in *CreateDelegatedStakingRequest, opts ...grpc.CallOption) (*CreateDelegatedStakingResponse, error)
	UpdateDelegatedStaking(ctx context.Context, in *UpdateDelegatedStakingRequest, opts ...grpc.CallOption) (*UpdateDelegatedStakingResponse, error)
	GetDelegatedStaking(ctx context.Context, in *GetDelegatedStakingRequest, opts ...grpc.CallOption) (*GetDelegatedStakingResponse, error)
	GetDelegatedStakings(ctx context.Context, in *GetDelegatedStakingsRequest, opts ...grpc.CallOption) (*GetDelegatedStakingsResponse, error)
	ExistDelegatedStakingConds(ctx context.Context, in *ExistDelegatedStakingCondsRequest, opts ...grpc.CallOption) (*ExistDelegatedStakingCondsResponse, error)
	DeleteDelegatedStaking(ctx context.Context, in *DeleteDelegatedStakingRequest, opts ...grpc.CallOption) (*DeleteDelegatedStakingResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateDelegatedStaking(ctx context.Context, in *CreateDelegatedStakingRequest, opts ...grpc.CallOption) (*CreateDelegatedStakingResponse, error) {
	out := new(CreateDelegatedStakingResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.delegatedstaking.v1.Middleware/CreateDelegatedStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateDelegatedStaking(ctx context.Context, in *UpdateDelegatedStakingRequest, opts ...grpc.CallOption) (*UpdateDelegatedStakingResponse, error) {
	out := new(UpdateDelegatedStakingResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.delegatedstaking.v1.Middleware/UpdateDelegatedStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDelegatedStaking(ctx context.Context, in *GetDelegatedStakingRequest, opts ...grpc.CallOption) (*GetDelegatedStakingResponse, error) {
	out := new(GetDelegatedStakingResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetDelegatedStakings(ctx context.Context, in *GetDelegatedStakingsRequest, opts ...grpc.CallOption) (*GetDelegatedStakingsResponse, error) {
	out := new(GetDelegatedStakingsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStakings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistDelegatedStakingConds(ctx context.Context, in *ExistDelegatedStakingCondsRequest, opts ...grpc.CallOption) (*ExistDelegatedStakingCondsResponse, error) {
	out := new(ExistDelegatedStakingCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistDelegatedStakingConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteDelegatedStaking(ctx context.Context, in *DeleteDelegatedStakingRequest, opts ...grpc.CallOption) (*DeleteDelegatedStakingResponse, error) {
	out := new(DeleteDelegatedStakingResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.delegatedstaking.v1.Middleware/DeleteDelegatedStaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateDelegatedStaking(context.Context, *CreateDelegatedStakingRequest) (*CreateDelegatedStakingResponse, error)
	UpdateDelegatedStaking(context.Context, *UpdateDelegatedStakingRequest) (*UpdateDelegatedStakingResponse, error)
	GetDelegatedStaking(context.Context, *GetDelegatedStakingRequest) (*GetDelegatedStakingResponse, error)
	GetDelegatedStakings(context.Context, *GetDelegatedStakingsRequest) (*GetDelegatedStakingsResponse, error)
	ExistDelegatedStakingConds(context.Context, *ExistDelegatedStakingCondsRequest) (*ExistDelegatedStakingCondsResponse, error)
	DeleteDelegatedStaking(context.Context, *DeleteDelegatedStakingRequest) (*DeleteDelegatedStakingResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateDelegatedStaking(context.Context, *CreateDelegatedStakingRequest) (*CreateDelegatedStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelegatedStaking not implemented")
}
func (UnimplementedMiddlewareServer) UpdateDelegatedStaking(context.Context, *UpdateDelegatedStakingRequest) (*UpdateDelegatedStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDelegatedStaking not implemented")
}
func (UnimplementedMiddlewareServer) GetDelegatedStaking(context.Context, *GetDelegatedStakingRequest) (*GetDelegatedStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegatedStaking not implemented")
}
func (UnimplementedMiddlewareServer) GetDelegatedStakings(context.Context, *GetDelegatedStakingsRequest) (*GetDelegatedStakingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegatedStakings not implemented")
}
func (UnimplementedMiddlewareServer) ExistDelegatedStakingConds(context.Context, *ExistDelegatedStakingCondsRequest) (*ExistDelegatedStakingCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDelegatedStakingConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteDelegatedStaking(context.Context, *DeleteDelegatedStakingRequest) (*DeleteDelegatedStakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDelegatedStaking not implemented")
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

func _Middleware_CreateDelegatedStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDelegatedStakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateDelegatedStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.delegatedstaking.v1.Middleware/CreateDelegatedStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateDelegatedStaking(ctx, req.(*CreateDelegatedStakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateDelegatedStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDelegatedStakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateDelegatedStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.delegatedstaking.v1.Middleware/UpdateDelegatedStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateDelegatedStaking(ctx, req.(*UpdateDelegatedStakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDelegatedStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegatedStakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDelegatedStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDelegatedStaking(ctx, req.(*GetDelegatedStakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetDelegatedStakings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegatedStakingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetDelegatedStakings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.delegatedstaking.v1.Middleware/GetDelegatedStakings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetDelegatedStakings(ctx, req.(*GetDelegatedStakingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistDelegatedStakingConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDelegatedStakingCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistDelegatedStakingConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistDelegatedStakingConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistDelegatedStakingConds(ctx, req.(*ExistDelegatedStakingCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteDelegatedStaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDelegatedStakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteDelegatedStaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.delegatedstaking.v1.Middleware/DeleteDelegatedStaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteDelegatedStaking(ctx, req.(*DeleteDelegatedStakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.delegatedstaking.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDelegatedStaking",
			Handler:    _Middleware_CreateDelegatedStaking_Handler,
		},
		{
			MethodName: "UpdateDelegatedStaking",
			Handler:    _Middleware_UpdateDelegatedStaking_Handler,
		},
		{
			MethodName: "GetDelegatedStaking",
			Handler:    _Middleware_GetDelegatedStaking_Handler,
		},
		{
			MethodName: "GetDelegatedStakings",
			Handler:    _Middleware_GetDelegatedStakings_Handler,
		},
		{
			MethodName: "ExistDelegatedStakingConds",
			Handler:    _Middleware_ExistDelegatedStakingConds_Handler,
		},
		{
			MethodName: "DeleteDelegatedStaking",
			Handler:    _Middleware_DeleteDelegatedStaking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/delegatedstaking/delegatedstaking.proto",
}
