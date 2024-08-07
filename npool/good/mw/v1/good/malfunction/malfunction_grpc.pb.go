// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/good/malfunction/malfunction.proto

package malfunction

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
	Middleware_CreateMalfunction_FullMethodName     = "/good.middleware.good1.malfunction.v1.Middleware/CreateMalfunction"
	Middleware_GetMalfunction_FullMethodName        = "/good.middleware.good1.malfunction.v1.Middleware/GetMalfunction"
	Middleware_ExistMalfunction_FullMethodName      = "/good.middleware.good1.malfunction.v1.Middleware/ExistMalfunction"
	Middleware_ExistMalfunctionConds_FullMethodName = "/good.middleware.good1.malfunction.v1.Middleware/ExistMalfunctionConds"
	Middleware_GetMalfunctions_FullMethodName       = "/good.middleware.good1.malfunction.v1.Middleware/GetMalfunctions"
	Middleware_UpdateMalfunction_FullMethodName     = "/good.middleware.good1.malfunction.v1.Middleware/UpdateMalfunction"
	Middleware_DeleteMalfunction_FullMethodName     = "/good.middleware.good1.malfunction.v1.Middleware/DeleteMalfunction"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateMalfunction(ctx context.Context, in *CreateMalfunctionRequest, opts ...grpc.CallOption) (*CreateMalfunctionResponse, error)
	GetMalfunction(ctx context.Context, in *GetMalfunctionRequest, opts ...grpc.CallOption) (*GetMalfunctionResponse, error)
	ExistMalfunction(ctx context.Context, in *ExistMalfunctionRequest, opts ...grpc.CallOption) (*ExistMalfunctionResponse, error)
	ExistMalfunctionConds(ctx context.Context, in *ExistMalfunctionCondsRequest, opts ...grpc.CallOption) (*ExistMalfunctionCondsResponse, error)
	GetMalfunctions(ctx context.Context, in *GetMalfunctionsRequest, opts ...grpc.CallOption) (*GetMalfunctionsResponse, error)
	UpdateMalfunction(ctx context.Context, in *UpdateMalfunctionRequest, opts ...grpc.CallOption) (*UpdateMalfunctionResponse, error)
	DeleteMalfunction(ctx context.Context, in *DeleteMalfunctionRequest, opts ...grpc.CallOption) (*DeleteMalfunctionResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateMalfunction(ctx context.Context, in *CreateMalfunctionRequest, opts ...grpc.CallOption) (*CreateMalfunctionResponse, error) {
	out := new(CreateMalfunctionResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateMalfunction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetMalfunction(ctx context.Context, in *GetMalfunctionRequest, opts ...grpc.CallOption) (*GetMalfunctionResponse, error) {
	out := new(GetMalfunctionResponse)
	err := c.cc.Invoke(ctx, Middleware_GetMalfunction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistMalfunction(ctx context.Context, in *ExistMalfunctionRequest, opts ...grpc.CallOption) (*ExistMalfunctionResponse, error) {
	out := new(ExistMalfunctionResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistMalfunction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistMalfunctionConds(ctx context.Context, in *ExistMalfunctionCondsRequest, opts ...grpc.CallOption) (*ExistMalfunctionCondsResponse, error) {
	out := new(ExistMalfunctionCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistMalfunctionConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetMalfunctions(ctx context.Context, in *GetMalfunctionsRequest, opts ...grpc.CallOption) (*GetMalfunctionsResponse, error) {
	out := new(GetMalfunctionsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetMalfunctions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateMalfunction(ctx context.Context, in *UpdateMalfunctionRequest, opts ...grpc.CallOption) (*UpdateMalfunctionResponse, error) {
	out := new(UpdateMalfunctionResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateMalfunction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteMalfunction(ctx context.Context, in *DeleteMalfunctionRequest, opts ...grpc.CallOption) (*DeleteMalfunctionResponse, error) {
	out := new(DeleteMalfunctionResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteMalfunction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateMalfunction(context.Context, *CreateMalfunctionRequest) (*CreateMalfunctionResponse, error)
	GetMalfunction(context.Context, *GetMalfunctionRequest) (*GetMalfunctionResponse, error)
	ExistMalfunction(context.Context, *ExistMalfunctionRequest) (*ExistMalfunctionResponse, error)
	ExistMalfunctionConds(context.Context, *ExistMalfunctionCondsRequest) (*ExistMalfunctionCondsResponse, error)
	GetMalfunctions(context.Context, *GetMalfunctionsRequest) (*GetMalfunctionsResponse, error)
	UpdateMalfunction(context.Context, *UpdateMalfunctionRequest) (*UpdateMalfunctionResponse, error)
	DeleteMalfunction(context.Context, *DeleteMalfunctionRequest) (*DeleteMalfunctionResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateMalfunction(context.Context, *CreateMalfunctionRequest) (*CreateMalfunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMalfunction not implemented")
}
func (UnimplementedMiddlewareServer) GetMalfunction(context.Context, *GetMalfunctionRequest) (*GetMalfunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMalfunction not implemented")
}
func (UnimplementedMiddlewareServer) ExistMalfunction(context.Context, *ExistMalfunctionRequest) (*ExistMalfunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMalfunction not implemented")
}
func (UnimplementedMiddlewareServer) ExistMalfunctionConds(context.Context, *ExistMalfunctionCondsRequest) (*ExistMalfunctionCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMalfunctionConds not implemented")
}
func (UnimplementedMiddlewareServer) GetMalfunctions(context.Context, *GetMalfunctionsRequest) (*GetMalfunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMalfunctions not implemented")
}
func (UnimplementedMiddlewareServer) UpdateMalfunction(context.Context, *UpdateMalfunctionRequest) (*UpdateMalfunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMalfunction not implemented")
}
func (UnimplementedMiddlewareServer) DeleteMalfunction(context.Context, *DeleteMalfunctionRequest) (*DeleteMalfunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMalfunction not implemented")
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

func _Middleware_CreateMalfunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMalfunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateMalfunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateMalfunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateMalfunction(ctx, req.(*CreateMalfunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetMalfunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMalfunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetMalfunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetMalfunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetMalfunction(ctx, req.(*GetMalfunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistMalfunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMalfunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistMalfunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistMalfunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistMalfunction(ctx, req.(*ExistMalfunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistMalfunctionConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMalfunctionCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistMalfunctionConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistMalfunctionConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistMalfunctionConds(ctx, req.(*ExistMalfunctionCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetMalfunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMalfunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetMalfunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetMalfunctions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetMalfunctions(ctx, req.(*GetMalfunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateMalfunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMalfunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateMalfunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateMalfunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateMalfunction(ctx, req.(*UpdateMalfunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteMalfunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMalfunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteMalfunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteMalfunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteMalfunction(ctx, req.(*DeleteMalfunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.good1.malfunction.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMalfunction",
			Handler:    _Middleware_CreateMalfunction_Handler,
		},
		{
			MethodName: "GetMalfunction",
			Handler:    _Middleware_GetMalfunction_Handler,
		},
		{
			MethodName: "ExistMalfunction",
			Handler:    _Middleware_ExistMalfunction_Handler,
		},
		{
			MethodName: "ExistMalfunctionConds",
			Handler:    _Middleware_ExistMalfunctionConds_Handler,
		},
		{
			MethodName: "GetMalfunctions",
			Handler:    _Middleware_GetMalfunctions_Handler,
		},
		{
			MethodName: "UpdateMalfunction",
			Handler:    _Middleware_UpdateMalfunction_Handler,
		},
		{
			MethodName: "DeleteMalfunction",
			Handler:    _Middleware_DeleteMalfunction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/good/malfunction/malfunction.proto",
}
