// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCommission(ctx context.Context, in *CreateCommissionRequest, opts ...grpc.CallOption) (*CreateCommissionResponse, error)
	UpdateCommission(ctx context.Context, in *UpdateCommissionRequest, opts ...grpc.CallOption) (*UpdateCommissionResponse, error)
	GetCommissionOnly(ctx context.Context, in *GetCommissionOnlyRequest, opts ...grpc.CallOption) (*GetCommissionOnlyResponse, error)
	GetCommissions(ctx context.Context, in *GetCommissionsRequest, opts ...grpc.CallOption) (*GetCommissionsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCommission(ctx context.Context, in *CreateCommissionRequest, opts ...grpc.CallOption) (*CreateCommissionResponse, error) {
	out := new(CreateCommissionResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.commission.v1.Middleware/CreateCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCommission(ctx context.Context, in *UpdateCommissionRequest, opts ...grpc.CallOption) (*UpdateCommissionResponse, error) {
	out := new(UpdateCommissionResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.commission.v1.Middleware/UpdateCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCommissionOnly(ctx context.Context, in *GetCommissionOnlyRequest, opts ...grpc.CallOption) (*GetCommissionOnlyResponse, error) {
	out := new(GetCommissionOnlyResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.commission.v1.Middleware/GetCommissionOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCommissions(ctx context.Context, in *GetCommissionsRequest, opts ...grpc.CallOption) (*GetCommissionsResponse, error) {
	out := new(GetCommissionsResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.commission.v1.Middleware/GetCommissions", in, out, opts...)
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
	GetCommissionOnly(context.Context, *GetCommissionOnlyRequest) (*GetCommissionOnlyResponse, error)
	GetCommissions(context.Context, *GetCommissionsRequest) (*GetCommissionsResponse, error)
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
func (UnimplementedMiddlewareServer) GetCommissionOnly(context.Context, *GetCommissionOnlyRequest) (*GetCommissionOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommissionOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetCommissions(context.Context, *GetCommissionsRequest) (*GetCommissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommissions not implemented")
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
		FullMethod: "/inspire.middleware.commission.v1.Middleware/CreateCommission",
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
		FullMethod: "/inspire.middleware.commission.v1.Middleware/UpdateCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCommission(ctx, req.(*UpdateCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCommissionOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommissionOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCommissionOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.commission.v1.Middleware/GetCommissionOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCommissionOnly(ctx, req.(*GetCommissionOnlyRequest))
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
		FullMethod: "/inspire.middleware.commission.v1.Middleware/GetCommissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCommissions(ctx, req.(*GetCommissionsRequest))
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
			MethodName: "GetCommissionOnly",
			Handler:    _Middleware_GetCommissionOnly_Handler,
		},
		{
			MethodName: "GetCommissions",
			Handler:    _Middleware_GetCommissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/commission/commission.proto",
}
