// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: npool/smoketest/gw/v1/relatedtestcase/relatedtestcase.proto

package relatedtestcase

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetRelatedTestCases(ctx context.Context, in *GetRelatedTestCasesRequest, opts ...grpc.CallOption) (*GetRelatedTestCasesResponse, error)
	CreateRelatedTestCase(ctx context.Context, in *CreateRelatedTestCaseRequest, opts ...grpc.CallOption) (*CreateRelatedTestCaseResponse, error)
	UpdateRelatedTestCase(ctx context.Context, in *UpdateRelatedTestCaseRequest, opts ...grpc.CallOption) (*UpdateRelatedTestCaseResponse, error)
	DeleteRelatedTestCase(ctx context.Context, in *DeleteRelatedTestCaseRequest, opts ...grpc.CallOption) (*DeleteRelatedTestCaseResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetRelatedTestCases(ctx context.Context, in *GetRelatedTestCasesRequest, opts ...grpc.CallOption) (*GetRelatedTestCasesResponse, error) {
	out := new(GetRelatedTestCasesResponse)
	err := c.cc.Invoke(ctx, "/smoketest.gateway.relatedtestcase.v1.Gateway/GetRelatedTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateRelatedTestCase(ctx context.Context, in *CreateRelatedTestCaseRequest, opts ...grpc.CallOption) (*CreateRelatedTestCaseResponse, error) {
	out := new(CreateRelatedTestCaseResponse)
	err := c.cc.Invoke(ctx, "/smoketest.gateway.relatedtestcase.v1.Gateway/CreateRelatedTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateRelatedTestCase(ctx context.Context, in *UpdateRelatedTestCaseRequest, opts ...grpc.CallOption) (*UpdateRelatedTestCaseResponse, error) {
	out := new(UpdateRelatedTestCaseResponse)
	err := c.cc.Invoke(ctx, "/smoketest.gateway.relatedtestcase.v1.Gateway/UpdateRelatedTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteRelatedTestCase(ctx context.Context, in *DeleteRelatedTestCaseRequest, opts ...grpc.CallOption) (*DeleteRelatedTestCaseResponse, error) {
	out := new(DeleteRelatedTestCaseResponse)
	err := c.cc.Invoke(ctx, "/smoketest.gateway.relatedtestcase.v1.Gateway/DeleteRelatedTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetRelatedTestCases(context.Context, *GetRelatedTestCasesRequest) (*GetRelatedTestCasesResponse, error)
	CreateRelatedTestCase(context.Context, *CreateRelatedTestCaseRequest) (*CreateRelatedTestCaseResponse, error)
	UpdateRelatedTestCase(context.Context, *UpdateRelatedTestCaseRequest) (*UpdateRelatedTestCaseResponse, error)
	DeleteRelatedTestCase(context.Context, *DeleteRelatedTestCaseRequest) (*DeleteRelatedTestCaseResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetRelatedTestCases(context.Context, *GetRelatedTestCasesRequest) (*GetRelatedTestCasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelatedTestCases not implemented")
}
func (UnimplementedGatewayServer) CreateRelatedTestCase(context.Context, *CreateRelatedTestCaseRequest) (*CreateRelatedTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelatedTestCase not implemented")
}
func (UnimplementedGatewayServer) UpdateRelatedTestCase(context.Context, *UpdateRelatedTestCaseRequest) (*UpdateRelatedTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelatedTestCase not implemented")
}
func (UnimplementedGatewayServer) DeleteRelatedTestCase(context.Context, *DeleteRelatedTestCaseRequest) (*DeleteRelatedTestCaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelatedTestCase not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_GetRelatedTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelatedTestCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRelatedTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.gateway.relatedtestcase.v1.Gateway/GetRelatedTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRelatedTestCases(ctx, req.(*GetRelatedTestCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateRelatedTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelatedTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateRelatedTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.gateway.relatedtestcase.v1.Gateway/CreateRelatedTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateRelatedTestCase(ctx, req.(*CreateRelatedTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateRelatedTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelatedTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateRelatedTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.gateway.relatedtestcase.v1.Gateway/UpdateRelatedTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateRelatedTestCase(ctx, req.(*UpdateRelatedTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteRelatedTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelatedTestCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteRelatedTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smoketest.gateway.relatedtestcase.v1.Gateway/DeleteRelatedTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteRelatedTestCase(ctx, req.(*DeleteRelatedTestCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.gateway.relatedtestcase.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRelatedTestCases",
			Handler:    _Gateway_GetRelatedTestCases_Handler,
		},
		{
			MethodName: "CreateRelatedTestCase",
			Handler:    _Gateway_CreateRelatedTestCase_Handler,
		},
		{
			MethodName: "UpdateRelatedTestCase",
			Handler:    _Gateway_UpdateRelatedTestCase_Handler,
		},
		{
			MethodName: "DeleteRelatedTestCase",
			Handler:    _Gateway_DeleteRelatedTestCase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/gw/v1/relatedtestcase/relatedtestcase.proto",
}
