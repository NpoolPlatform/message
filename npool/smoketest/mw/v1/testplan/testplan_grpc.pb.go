// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/smoketest/mw/v1/testplan/testplan.proto

package testplan

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
	Middleware_CreateTestPlan_FullMethodName = "/smoketest.middleware.testplan.v1.Middleware/CreateTestPlan"
	Middleware_UpdateTestPlan_FullMethodName = "/smoketest.middleware.testplan.v1.Middleware/UpdateTestPlan"
	Middleware_GetTestPlan_FullMethodName    = "/smoketest.middleware.testplan.v1.Middleware/GetTestPlan"
	Middleware_GetTestPlans_FullMethodName   = "/smoketest.middleware.testplan.v1.Middleware/GetTestPlans"
	Middleware_ExistTestPlan_FullMethodName  = "/smoketest.middleware.testplan.v1.Middleware/ExistTestPlan"
	Middleware_DeleteTestPlan_FullMethodName = "/smoketest.middleware.testplan.v1.Middleware/DeleteTestPlan"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateTestPlan(ctx context.Context, in *CreateTestPlanRequest, opts ...grpc.CallOption) (*CreateTestPlanResponse, error)
	UpdateTestPlan(ctx context.Context, in *UpdateTestPlanRequest, opts ...grpc.CallOption) (*UpdateTestPlanResponse, error)
	GetTestPlan(ctx context.Context, in *GetTestPlanRequest, opts ...grpc.CallOption) (*GetTestPlanResponse, error)
	GetTestPlans(ctx context.Context, in *GetTestPlansRequest, opts ...grpc.CallOption) (*GetTestPlansResponse, error)
	ExistTestPlan(ctx context.Context, in *ExistTestPlanRequest, opts ...grpc.CallOption) (*ExistTestPlanResponse, error)
	DeleteTestPlan(ctx context.Context, in *DeleteTestPlanRequest, opts ...grpc.CallOption) (*DeleteTestPlanResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateTestPlan(ctx context.Context, in *CreateTestPlanRequest, opts ...grpc.CallOption) (*CreateTestPlanResponse, error) {
	out := new(CreateTestPlanResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateTestPlan(ctx context.Context, in *UpdateTestPlanRequest, opts ...grpc.CallOption) (*UpdateTestPlanResponse, error) {
	out := new(UpdateTestPlanResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTestPlan(ctx context.Context, in *GetTestPlanRequest, opts ...grpc.CallOption) (*GetTestPlanResponse, error) {
	out := new(GetTestPlanResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTestPlans(ctx context.Context, in *GetTestPlansRequest, opts ...grpc.CallOption) (*GetTestPlansResponse, error) {
	out := new(GetTestPlansResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTestPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistTestPlan(ctx context.Context, in *ExistTestPlanRequest, opts ...grpc.CallOption) (*ExistTestPlanResponse, error) {
	out := new(ExistTestPlanResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteTestPlan(ctx context.Context, in *DeleteTestPlanRequest, opts ...grpc.CallOption) (*DeleteTestPlanResponse, error) {
	out := new(DeleteTestPlanResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateTestPlan(context.Context, *CreateTestPlanRequest) (*CreateTestPlanResponse, error)
	UpdateTestPlan(context.Context, *UpdateTestPlanRequest) (*UpdateTestPlanResponse, error)
	GetTestPlan(context.Context, *GetTestPlanRequest) (*GetTestPlanResponse, error)
	GetTestPlans(context.Context, *GetTestPlansRequest) (*GetTestPlansResponse, error)
	ExistTestPlan(context.Context, *ExistTestPlanRequest) (*ExistTestPlanResponse, error)
	DeleteTestPlan(context.Context, *DeleteTestPlanRequest) (*DeleteTestPlanResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateTestPlan(context.Context, *CreateTestPlanRequest) (*CreateTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestPlan not implemented")
}
func (UnimplementedMiddlewareServer) UpdateTestPlan(context.Context, *UpdateTestPlanRequest) (*UpdateTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestPlan not implemented")
}
func (UnimplementedMiddlewareServer) GetTestPlan(context.Context, *GetTestPlanRequest) (*GetTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestPlan not implemented")
}
func (UnimplementedMiddlewareServer) GetTestPlans(context.Context, *GetTestPlansRequest) (*GetTestPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestPlans not implemented")
}
func (UnimplementedMiddlewareServer) ExistTestPlan(context.Context, *ExistTestPlanRequest) (*ExistTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTestPlan not implemented")
}
func (UnimplementedMiddlewareServer) DeleteTestPlan(context.Context, *DeleteTestPlanRequest) (*DeleteTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestPlan not implemented")
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

func _Middleware_CreateTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateTestPlan(ctx, req.(*CreateTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateTestPlan(ctx, req.(*UpdateTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTestPlan(ctx, req.(*GetTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTestPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTestPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTestPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTestPlans(ctx, req.(*GetTestPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistTestPlan(ctx, req.(*ExistTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteTestPlan(ctx, req.(*DeleteTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.middleware.testplan.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTestPlan",
			Handler:    _Middleware_CreateTestPlan_Handler,
		},
		{
			MethodName: "UpdateTestPlan",
			Handler:    _Middleware_UpdateTestPlan_Handler,
		},
		{
			MethodName: "GetTestPlan",
			Handler:    _Middleware_GetTestPlan_Handler,
		},
		{
			MethodName: "GetTestPlans",
			Handler:    _Middleware_GetTestPlans_Handler,
		},
		{
			MethodName: "ExistTestPlan",
			Handler:    _Middleware_ExistTestPlan_Handler,
		},
		{
			MethodName: "DeleteTestPlan",
			Handler:    _Middleware_DeleteTestPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/mw/v1/testplan/testplan.proto",
}
