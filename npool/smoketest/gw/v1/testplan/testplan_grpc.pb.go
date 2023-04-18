// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: npool/smoketest/gw/v1/testplan/testplan.proto

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
	Gateway_CreateTestPlan_FullMethodName = "/smoketest.gateway.testplan.v1.Gateway/CreateTestPlan"
	Gateway_DeleteTestPlan_FullMethodName = "/smoketest.gateway.testplan.v1.Gateway/DeleteTestPlan"
	Gateway_UpdateTestPlan_FullMethodName = "/smoketest.gateway.testplan.v1.Gateway/UpdateTestPlan"
	Gateway_GetTestPlans_FullMethodName   = "/smoketest.gateway.testplan.v1.Gateway/GetTestPlans"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateTestPlan(ctx context.Context, in *CreateTestPlanRequest, opts ...grpc.CallOption) (*CreateTestPlanResponse, error)
	DeleteTestPlan(ctx context.Context, in *DeleteTestPlanRequest, opts ...grpc.CallOption) (*DeleteTestPlanResponse, error)
	UpdateTestPlan(ctx context.Context, in *UpdateTestPlanRequest, opts ...grpc.CallOption) (*UpdateTestPlanResponse, error)
	GetTestPlans(ctx context.Context, in *GetTestPlansRequest, opts ...grpc.CallOption) (*GetTestPlansResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateTestPlan(ctx context.Context, in *CreateTestPlanRequest, opts ...grpc.CallOption) (*CreateTestPlanResponse, error) {
	out := new(CreateTestPlanResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteTestPlan(ctx context.Context, in *DeleteTestPlanRequest, opts ...grpc.CallOption) (*DeleteTestPlanResponse, error) {
	out := new(DeleteTestPlanResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateTestPlan(ctx context.Context, in *UpdateTestPlanRequest, opts ...grpc.CallOption) (*UpdateTestPlanResponse, error) {
	out := new(UpdateTestPlanResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateTestPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetTestPlans(ctx context.Context, in *GetTestPlansRequest, opts ...grpc.CallOption) (*GetTestPlansResponse, error) {
	out := new(GetTestPlansResponse)
	err := c.cc.Invoke(ctx, Gateway_GetTestPlans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateTestPlan(context.Context, *CreateTestPlanRequest) (*CreateTestPlanResponse, error)
	DeleteTestPlan(context.Context, *DeleteTestPlanRequest) (*DeleteTestPlanResponse, error)
	UpdateTestPlan(context.Context, *UpdateTestPlanRequest) (*UpdateTestPlanResponse, error)
	GetTestPlans(context.Context, *GetTestPlansRequest) (*GetTestPlansResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateTestPlan(context.Context, *CreateTestPlanRequest) (*CreateTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestPlan not implemented")
}
func (UnimplementedGatewayServer) DeleteTestPlan(context.Context, *DeleteTestPlanRequest) (*DeleteTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestPlan not implemented")
}
func (UnimplementedGatewayServer) UpdateTestPlan(context.Context, *UpdateTestPlanRequest) (*UpdateTestPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestPlan not implemented")
}
func (UnimplementedGatewayServer) GetTestPlans(context.Context, *GetTestPlansRequest) (*GetTestPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestPlans not implemented")
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

func _Gateway_CreateTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateTestPlan(ctx, req.(*CreateTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteTestPlan(ctx, req.(*DeleteTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateTestPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateTestPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateTestPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateTestPlan(ctx, req.(*UpdateTestPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetTestPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetTestPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetTestPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetTestPlans(ctx, req.(*GetTestPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smoketest.gateway.testplan.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTestPlan",
			Handler:    _Gateway_CreateTestPlan_Handler,
		},
		{
			MethodName: "DeleteTestPlan",
			Handler:    _Gateway_DeleteTestPlan_Handler,
		},
		{
			MethodName: "UpdateTestPlan",
			Handler:    _Gateway_UpdateTestPlan_Handler,
		},
		{
			MethodName: "GetTestPlans",
			Handler:    _Gateway_GetTestPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/smoketest/gw/v1/testplan/testplan.proto",
}
