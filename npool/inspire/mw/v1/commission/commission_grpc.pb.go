// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/commission/commission.proto

package commission

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
	Middleware_CreateCommission_FullMethodName = "/inspire.middleware.commission.v1.Middleware/CreateCommission"
	Middleware_UpdateCommission_FullMethodName = "/inspire.middleware.commission.v1.Middleware/UpdateCommission"
	Middleware_GetCommission_FullMethodName    = "/inspire.middleware.commission.v1.Middleware/GetCommission"
	Middleware_GetCommissions_FullMethodName   = "/inspire.middleware.commission.v1.Middleware/GetCommissions"
	Middleware_CloneCommissions_FullMethodName = "/inspire.middleware.commission.v1.Middleware/CloneCommissions"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCommission(ctx context.Context, in *CreateCommissionRequest, opts ...grpc.CallOption) (*CreateCommissionResponse, error)
	UpdateCommission(ctx context.Context, in *UpdateCommissionRequest, opts ...grpc.CallOption) (*UpdateCommissionResponse, error)
	GetCommission(ctx context.Context, in *GetCommissionRequest, opts ...grpc.CallOption) (*GetCommissionResponse, error)
	GetCommissions(ctx context.Context, in *GetCommissionsRequest, opts ...grpc.CallOption) (*GetCommissionsResponse, error)
	CloneCommissions(ctx context.Context, in *CloneCommissionsRequest, opts ...grpc.CallOption) (*CloneCommissionsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCommission(ctx context.Context, in *CreateCommissionRequest, opts ...grpc.CallOption) (*CreateCommissionResponse, error) {
	out := new(CreateCommissionResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCommission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCommission(ctx context.Context, in *UpdateCommissionRequest, opts ...grpc.CallOption) (*UpdateCommissionResponse, error) {
	out := new(UpdateCommissionResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCommission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCommission(ctx context.Context, in *GetCommissionRequest, opts ...grpc.CallOption) (*GetCommissionResponse, error) {
	out := new(GetCommissionResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCommission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCommissions(ctx context.Context, in *GetCommissionsRequest, opts ...grpc.CallOption) (*GetCommissionsResponse, error) {
	out := new(GetCommissionsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCommissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CloneCommissions(ctx context.Context, in *CloneCommissionsRequest, opts ...grpc.CallOption) (*CloneCommissionsResponse, error) {
	out := new(CloneCommissionsResponse)
	err := c.cc.Invoke(ctx, Middleware_CloneCommissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCommission(context.Context, *CreateCommissionRequest) (*CreateCommissionResponse, error)
	UpdateCommission(context.Context, *UpdateCommissionRequest) (*UpdateCommissionResponse, error)
	GetCommission(context.Context, *GetCommissionRequest) (*GetCommissionResponse, error)
	GetCommissions(context.Context, *GetCommissionsRequest) (*GetCommissionsResponse, error)
	CloneCommissions(context.Context, *CloneCommissionsRequest) (*CloneCommissionsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCommission(context.Context, *CreateCommissionRequest) (*CreateCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommission not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCommission(context.Context, *UpdateCommissionRequest) (*UpdateCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommission not implemented")
}
func (UnimplementedMiddlewareServer) GetCommission(context.Context, *GetCommissionRequest) (*GetCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommission not implemented")
}
func (UnimplementedMiddlewareServer) GetCommissions(context.Context, *GetCommissionsRequest) (*GetCommissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommissions not implemented")
}
func (UnimplementedMiddlewareServer) CloneCommissions(context.Context, *CloneCommissionsRequest) (*CloneCommissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneCommissions not implemented")
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

func _Middleware_CreateCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCommission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCommission(ctx, req.(*CreateCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCommission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCommission(ctx, req.(*UpdateCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCommission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCommission(ctx, req.(*GetCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCommissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCommissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCommissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCommissions(ctx, req.(*GetCommissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CloneCommissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneCommissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CloneCommissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CloneCommissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CloneCommissions(ctx, req.(*CloneCommissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.commission.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommission",
			Handler:    _Middleware_CreateCommission_Handler,
		},
		{
			MethodName: "UpdateCommission",
			Handler:    _Middleware_UpdateCommission_Handler,
		},
		{
			MethodName: "GetCommission",
			Handler:    _Middleware_GetCommission_Handler,
		},
		{
			MethodName: "GetCommissions",
			Handler:    _Middleware_GetCommissions_Handler,
		},
		{
			MethodName: "CloneCommissions",
			Handler:    _Middleware_CloneCommissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/commission/commission.proto",
}
