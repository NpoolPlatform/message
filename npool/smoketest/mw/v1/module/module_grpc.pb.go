// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/smoketest/mw/v1/module/module.proto

package module

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
	Middleware_CreateModule_FullMethodName      = "/smoketest.middleware.module.v1.Middleware/CreateModule"
	Middleware_UpdateModule_FullMethodName      = "/smoketest.middleware.module.v1.Middleware/UpdateModule"
	Middleware_GetModules_FullMethodName        = "/smoketest.middleware.module.v1.Middleware/GetModules"
	Middleware_GetModule_FullMethodName         = "/smoketest.middleware.module.v1.Middleware/GetModule"
	Middleware_GetModuleConds_FullMethodName    = "/smoketest.middleware.module.v1.Middleware/GetModuleConds"
	Middleware_DeleteModule_FullMethodName      = "/smoketest.middleware.module.v1.Middleware/DeleteModule"
	Middleware_ExistModule_FullMethodName       = "/smoketest.middleware.module.v1.Middleware/ExistModule"
	Middleware_ExistModuleByName_FullMethodName = "/smoketest.middleware.module.v1.Middleware/ExistModuleByName"
	Middleware_ExistModuleConds_FullMethodName  = "/smoketest.middleware.module.v1.Middleware/ExistModuleConds"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateModule(ctx context.Context, in *CreateModuleRequest, opts ...grpc.CallOption) (*CreateModuleResponse, error)
	UpdateModule(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*UpdateModuleResponse, error)
	GetModules(ctx context.Context, in *GetModulesRequest, opts ...grpc.CallOption) (*GetModulesResponse, error)
	GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error)
	GetModuleConds(ctx context.Context, in *GetModuleCondsRequest, opts ...grpc.CallOption) (*GetModuleCondsResponse, error)
	DeleteModule(ctx context.Context, in *DeleteModuleRequest, opts ...grpc.CallOption) (*DeleteModuleResponse, error)
	ExistModule(ctx context.Context, in *ExistModuleRequest, opts ...grpc.CallOption) (*ExistModuleResponse, error)
	ExistModuleByName(ctx context.Context, in *ExistModuleByNameRequest, opts ...grpc.CallOption) (*ExistModuleByNameResponse, error)
	ExistModuleConds(ctx context.Context, in *ExistModuleCondsRequest, opts ...grpc.CallOption) (*ExistModuleCondsResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateModule(ctx context.Context, in *CreateModuleRequest, opts ...grpc.CallOption) (*CreateModuleResponse, error) {
	out := new(CreateModuleResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateModule(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*UpdateModuleResponse, error) {
	out := new(UpdateModuleResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetModules(ctx context.Context, in *GetModulesRequest, opts ...grpc.CallOption) (*GetModulesResponse, error) {
	out := new(GetModulesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetModules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error) {
	out := new(GetModuleResponse)
	err := c.cc.Invoke(ctx, Middleware_GetModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetModuleConds(ctx context.Context, in *GetModuleCondsRequest, opts ...grpc.CallOption) (*GetModuleCondsResponse, error) {
	out := new(GetModuleCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetModuleConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteModule(ctx context.Context, in *DeleteModuleRequest, opts ...grpc.CallOption) (*DeleteModuleResponse, error) {
	out := new(DeleteModuleResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistModule(ctx context.Context, in *ExistModuleRequest, opts ...grpc.CallOption) (*ExistModuleResponse, error) {
	out := new(ExistModuleResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistModule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistModuleByName(ctx context.Context, in *ExistModuleByNameRequest, opts ...grpc.CallOption) (*ExistModuleByNameResponse, error) {
	out := new(ExistModuleByNameResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistModuleByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistModuleConds(ctx context.Context, in *ExistModuleCondsRequest, opts ...grpc.CallOption) (*ExistModuleCondsResponse, error) {
	out := new(ExistModuleCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistModuleConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateModule(context.Context, *CreateModuleRequest) (*CreateModuleResponse, error)
	UpdateModule(context.Context, *UpdateModuleRequest) (*UpdateModuleResponse, error)
	GetModules(context.Context, *GetModulesRequest) (*GetModulesResponse, error)
	GetModule(context.Context, *GetModuleRequest) (*GetModuleResponse, error)
	GetModuleConds(context.Context, *GetModuleCondsRequest) (*GetModuleCondsResponse, error)
	DeleteModule(context.Context, *DeleteModuleRequest) (*DeleteModuleResponse, error)
	ExistModule(context.Context, *ExistModuleRequest) (*ExistModuleResponse, error)
	ExistModuleByName(context.Context, *ExistModuleByNameRequest) (*ExistModuleByNameResponse, error)
	ExistModuleConds(context.Context, *ExistModuleCondsRequest) (*ExistModuleCondsResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateModule(context.Context, *CreateModuleRequest) (*CreateModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModule not implemented")
}
func (UnimplementedMiddlewareServer) UpdateModule(context.Context, *UpdateModuleRequest) (*UpdateModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModule not implemented")
}
func (UnimplementedMiddlewareServer) GetModules(context.Context, *GetModulesRequest) (*GetModulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModules not implemented")
}
func (UnimplementedMiddlewareServer) GetModule(context.Context, *GetModuleRequest) (*GetModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModule not implemented")
}
func (UnimplementedMiddlewareServer) GetModuleConds(context.Context, *GetModuleCondsRequest) (*GetModuleCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModuleConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteModule(context.Context, *DeleteModuleRequest) (*DeleteModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModule not implemented")
}
func (UnimplementedMiddlewareServer) ExistModule(context.Context, *ExistModuleRequest) (*ExistModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistModule not implemented")
}
func (UnimplementedMiddlewareServer) ExistModuleByName(context.Context, *ExistModuleByNameRequest) (*ExistModuleByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistModuleByName not implemented")
}
func (UnimplementedMiddlewareServer) ExistModuleConds(context.Context, *ExistModuleCondsRequest) (*ExistModuleCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistModuleConds not implemented")
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

func _Middleware_CreateModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateModule(ctx, req.(*CreateModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateModule(ctx, req.(*UpdateModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetModules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetModules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetModules(ctx, req.(*GetModulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetModule(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetModuleConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetModuleConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetModuleConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetModuleConds(ctx, req.(*GetModuleCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteModule(ctx, req.(*DeleteModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistModule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistModule(ctx, req.(*ExistModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistModuleByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistModuleByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistModuleByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistModuleByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistModuleByName(ctx, req.(*ExistModuleByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistModuleConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistModuleCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistModuleConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistModuleConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistModuleConds(ctx, req.(*ExistModuleCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.middleware.module.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateModule",
			Handler:    _Middleware_CreateModule_Handler,
		},
		{
			MethodName: "UpdateModule",
			Handler:    _Middleware_UpdateModule_Handler,
		},
		{
			MethodName: "GetModules",
			Handler:    _Middleware_GetModules_Handler,
		},
		{
			MethodName: "GetModule",
			Handler:    _Middleware_GetModule_Handler,
		},
		{
			MethodName: "GetModuleConds",
			Handler:    _Middleware_GetModuleConds_Handler,
		},
		{
			MethodName: "DeleteModule",
			Handler:    _Middleware_DeleteModule_Handler,
		},
		{
			MethodName: "ExistModule",
			Handler:    _Middleware_ExistModule_Handler,
		},
		{
			MethodName: "ExistModuleByName",
			Handler:    _Middleware_ExistModuleByName_Handler,
		},
		{
			MethodName: "ExistModuleConds",
			Handler:    _Middleware_ExistModuleConds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/mw/v1/module/module.proto",
}
