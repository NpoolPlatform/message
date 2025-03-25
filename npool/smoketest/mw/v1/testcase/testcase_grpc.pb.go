// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/smoketest/mw/v1/testcase/testcase.proto

package testcase

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
	Middleware_CreateTestCase_FullMethodName = "/smoketest.middleware.testcase.v1.Middleware/CreateTestCase"
	Middleware_UpdateTestCase_FullMethodName = "/smoketest.middleware.testcase.v1.Middleware/UpdateTestCase"
	Middleware_ExistTestCase_FullMethodName  = "/smoketest.middleware.testcase.v1.Middleware/ExistTestCase"
	Middleware_GetTestCases_FullMethodName   = "/smoketest.middleware.testcase.v1.Middleware/GetTestCases"
	Middleware_GetTestCase_FullMethodName    = "/smoketest.middleware.testcase.v1.Middleware/GetTestCase"
	Middleware_DeleteTestCase_FullMethodName = "/smoketest.middleware.testcase.v1.Middleware/DeleteTestCase"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateTestCase(ctx context.Context, in *CreateTestCaseRequest, opts ...grpc.CallOption) (*CreateTestCaseResponse, error)
	UpdateTestCase(ctx context.Context, in *UpdateTestCaseRequest, opts ...grpc.CallOption) (*UpdateTestCaseResponse, error)
	ExistTestCase(ctx context.Context, in *ExistTestCaseRequest, opts ...grpc.CallOption) (*ExistTestCaseResponse, error)
	GetTestCases(ctx context.Context, in *GetTestCasesRequest, opts ...grpc.CallOption) (*GetTestCasesResponse, error)
	GetTestCase(ctx context.Context, in *GetTestCaseRequest, opts ...grpc.CallOption) (*GetTestCaseResponse, error)
	DeleteTestCase(ctx context.Context, in *DeleteTestCaseRequest, opts ...grpc.CallOption) (*DeleteTestCaseResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateTestCase(ctx context.Context, in *CreateTestCaseRequest, opts ...grpc.CallOption) (*CreateTestCaseResponse, error) {
	out := new(CreateTestCaseResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateTestCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateTestCase(ctx context.Context, in *UpdateTestCaseRequest, opts ...grpc.CallOption) (*UpdateTestCaseResponse, error) {
	out := new(UpdateTestCaseResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateTestCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistTestCase(ctx context.Context, in *ExistTestCaseRequest, opts ...grpc.CallOption) (*ExistTestCaseResponse, error) {
	out := new(ExistTestCaseResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistTestCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTestCases(ctx context.Context, in *GetTestCasesRequest, opts ...grpc.CallOption) (*GetTestCasesResponse, error) {
	out := new(GetTestCasesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTestCases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetTestCase(ctx context.Context, in *GetTestCaseRequest, opts ...grpc.CallOption) (*GetTestCaseResponse, error) {
	out := new(GetTestCaseResponse)
	err := c.cc.Invoke(ctx, Middleware_GetTestCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteTestCase(ctx context.Context, in *DeleteTestCaseRequest, opts ...grpc.CallOption) (*DeleteTestCaseResponse, error) {
	out := new(DeleteTestCaseResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteTestCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateTestCase(context.Context, *CreateTestCaseRequest) (*CreateTestCaseResponse, error)
	UpdateTestCase(context.Context, *UpdateTestCaseRequest) (*UpdateTestCaseResponse, error)
	ExistTestCase(context.Context, *ExistTestCaseRequest) (*ExistTestCaseResponse, error)
	GetTestCases(context.Context, *GetTestCasesRequest) (*GetTestCasesResponse, error)
	GetTestCase(context.Context, *GetTestCaseRequest) (*GetTestCaseResponse, error)
	DeleteTestCase(context.Context, *DeleteTestCaseRequest) (*DeleteTestCaseResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateTestCase(context.Context, *CreateTestCaseRequest) (*CreateTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestCase not implemented")
}
func (UnimplementedMiddlewareServer) UpdateTestCase(context.Context, *UpdateTestCaseRequest) (*UpdateTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestCase not implemented")
}
func (UnimplementedMiddlewareServer) ExistTestCase(context.Context, *ExistTestCaseRequest) (*ExistTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTestCase not implemented")
}
func (UnimplementedMiddlewareServer) GetTestCases(context.Context, *GetTestCasesRequest) (*GetTestCasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestCases not implemented")
}
func (UnimplementedMiddlewareServer) GetTestCase(context.Context, *GetTestCaseRequest) (*GetTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestCase not implemented")
}
func (UnimplementedMiddlewareServer) DeleteTestCase(context.Context, *DeleteTestCaseRequest) (*DeleteTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestCase not implemented")
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

func _Middleware_CreateTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateTestCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateTestCase(ctx, req.(*CreateTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateTestCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateTestCase(ctx, req.(*UpdateTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistTestCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistTestCase(ctx, req.(*ExistTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTestCases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTestCases(ctx, req.(*GetTestCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetTestCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetTestCase(ctx, req.(*GetTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteTestCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteTestCase(ctx, req.(*DeleteTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.middleware.testcase.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTestCase",
			Handler:    _Middleware_CreateTestCase_Handler,
		},
		{
			MethodName: "UpdateTestCase",
			Handler:    _Middleware_UpdateTestCase_Handler,
		},
		{
			MethodName: "ExistTestCase",
			Handler:    _Middleware_ExistTestCase_Handler,
		},
		{
			MethodName: "GetTestCases",
			Handler:    _Middleware_GetTestCases_Handler,
		},
		{
			MethodName: "GetTestCase",
			Handler:    _Middleware_GetTestCase_Handler,
		},
		{
			MethodName: "DeleteTestCase",
			Handler:    _Middleware_DeleteTestCase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/mw/v1/testcase/testcase.proto",
}
