// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mw/v1/compensate/compensate.proto

package compensate

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
	Middleware_CreateCompensate_FullMethodName     = "/order.middleware.compensate.v1.Middleware/CreateCompensate"
	Middleware_CreateCompensates_FullMethodName    = "/order.middleware.compensate.v1.Middleware/CreateCompensates"
	Middleware_UpdateCompensate_FullMethodName     = "/order.middleware.compensate.v1.Middleware/UpdateCompensate"
	Middleware_GetCompensate_FullMethodName        = "/order.middleware.compensate.v1.Middleware/GetCompensate"
	Middleware_GetCompensateOnly_FullMethodName    = "/order.middleware.compensate.v1.Middleware/GetCompensateOnly"
	Middleware_GetCompensates_FullMethodName       = "/order.middleware.compensate.v1.Middleware/GetCompensates"
	Middleware_ExistCompensate_FullMethodName      = "/order.middleware.compensate.v1.Middleware/ExistCompensate"
	Middleware_ExistCompensateConds_FullMethodName = "/order.middleware.compensate.v1.Middleware/ExistCompensateConds"
	Middleware_CountCompensates_FullMethodName     = "/order.middleware.compensate.v1.Middleware/CountCompensates"
	Middleware_DeleteCompensate_FullMethodName     = "/order.middleware.compensate.v1.Middleware/DeleteCompensate"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCompensate(ctx context.Context, in *CreateCompensateRequest, opts ...grpc.CallOption) (*CreateCompensateResponse, error)
	CreateCompensates(ctx context.Context, in *CreateCompensatesRequest, opts ...grpc.CallOption) (*CreateCompensatesResponse, error)
	UpdateCompensate(ctx context.Context, in *UpdateCompensateRequest, opts ...grpc.CallOption) (*UpdateCompensateResponse, error)
	GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error)
	GetCompensateOnly(ctx context.Context, in *GetCompensateOnlyRequest, opts ...grpc.CallOption) (*GetCompensateOnlyResponse, error)
	GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error)
	ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error)
	ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error)
	CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error)
	DeleteCompensate(ctx context.Context, in *DeleteCompensateRequest, opts ...grpc.CallOption) (*DeleteCompensateResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCompensate(ctx context.Context, in *CreateCompensateRequest, opts ...grpc.CallOption) (*CreateCompensateResponse, error) {
	out := new(CreateCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateCompensates(ctx context.Context, in *CreateCompensatesRequest, opts ...grpc.CallOption) (*CreateCompensatesResponse, error) {
	out := new(CreateCompensatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCompensate(ctx context.Context, in *UpdateCompensateRequest, opts ...grpc.CallOption) (*UpdateCompensateResponse, error) {
	out := new(UpdateCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error) {
	out := new(GetCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCompensateOnly(ctx context.Context, in *GetCompensateOnlyRequest, opts ...grpc.CallOption) (*GetCompensateOnlyResponse, error) {
	out := new(GetCompensateOnlyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCompensateOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error) {
	out := new(GetCompensatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error) {
	out := new(ExistCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error) {
	out := new(ExistCompensateCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCompensateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error) {
	out := new(CountCompensatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CountCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCompensate(ctx context.Context, in *DeleteCompensateRequest, opts ...grpc.CallOption) (*DeleteCompensateResponse, error) {
	out := new(DeleteCompensateResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCompensate(context.Context, *CreateCompensateRequest) (*CreateCompensateResponse, error)
	CreateCompensates(context.Context, *CreateCompensatesRequest) (*CreateCompensatesResponse, error)
	UpdateCompensate(context.Context, *UpdateCompensateRequest) (*UpdateCompensateResponse, error)
	GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error)
	GetCompensateOnly(context.Context, *GetCompensateOnlyRequest) (*GetCompensateOnlyResponse, error)
	GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error)
	ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error)
	ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error)
	CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error)
	DeleteCompensate(context.Context, *DeleteCompensateRequest) (*DeleteCompensateResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCompensate(context.Context, *CreateCompensateRequest) (*CreateCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompensate not implemented")
}
func (UnimplementedMiddlewareServer) CreateCompensates(context.Context, *CreateCompensatesRequest) (*CreateCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompensates not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCompensate(context.Context, *UpdateCompensateRequest) (*UpdateCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompensate not implemented")
}
func (UnimplementedMiddlewareServer) GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensate not implemented")
}
func (UnimplementedMiddlewareServer) GetCompensateOnly(context.Context, *GetCompensateOnlyRequest) (*GetCompensateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensateOnly not implemented")
}
func (UnimplementedMiddlewareServer) GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensates not implemented")
}
func (UnimplementedMiddlewareServer) ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensate not implemented")
}
func (UnimplementedMiddlewareServer) ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensateConds not implemented")
}
func (UnimplementedMiddlewareServer) CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCompensates not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCompensate(context.Context, *DeleteCompensateRequest) (*DeleteCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompensate not implemented")
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

func _Middleware_CreateCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCompensate(ctx, req.(*CreateCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCompensates(ctx, req.(*CreateCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCompensate(ctx, req.(*UpdateCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCompensate(ctx, req.(*GetCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCompensateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCompensateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCompensateOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCompensateOnly(ctx, req.(*GetCompensateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCompensates(ctx, req.(*GetCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCompensate(ctx, req.(*ExistCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCompensateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCompensateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCompensateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCompensateConds(ctx, req.(*ExistCompensateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CountCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CountCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CountCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CountCompensates(ctx, req.(*CountCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCompensate(ctx, req.(*DeleteCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.middleware.compensate.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompensate",
			Handler:    _Middleware_CreateCompensate_Handler,
		},
		{
			MethodName: "CreateCompensates",
			Handler:    _Middleware_CreateCompensates_Handler,
		},
		{
			MethodName: "UpdateCompensate",
			Handler:    _Middleware_UpdateCompensate_Handler,
		},
		{
			MethodName: "GetCompensate",
			Handler:    _Middleware_GetCompensate_Handler,
		},
		{
			MethodName: "GetCompensateOnly",
			Handler:    _Middleware_GetCompensateOnly_Handler,
		},
		{
			MethodName: "GetCompensates",
			Handler:    _Middleware_GetCompensates_Handler,
		},
		{
			MethodName: "ExistCompensate",
			Handler:    _Middleware_ExistCompensate_Handler,
		},
		{
			MethodName: "ExistCompensateConds",
			Handler:    _Middleware_ExistCompensateConds_Handler,
		},
		{
			MethodName: "CountCompensates",
			Handler:    _Middleware_CountCompensates_Handler,
		},
		{
			MethodName: "DeleteCompensate",
			Handler:    _Middleware_DeleteCompensate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mw/v1/compensate/compensate.proto",
}