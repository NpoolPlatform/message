// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/task/task.proto

package task

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
	Gateway_AdminGetTasks_FullMethodName = "/inspire.gateway.task.v1.Gateway/AdminGetTasks"
	Gateway_GetMyTasks_FullMethodName    = "/inspire.gateway.task.v1.Gateway/GetMyTasks"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminGetTasks(ctx context.Context, in *AdminGetTasksRequest, opts ...grpc.CallOption) (*AdminGetTasksResponse, error)
	GetMyTasks(ctx context.Context, in *GetMyTasksRequest, opts ...grpc.CallOption) (*GetMyTasksResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminGetTasks(ctx context.Context, in *AdminGetTasksRequest, opts ...grpc.CallOption) (*AdminGetTasksResponse, error) {
	out := new(AdminGetTasksResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyTasks(ctx context.Context, in *GetMyTasksRequest, opts ...grpc.CallOption) (*GetMyTasksResponse, error) {
	out := new(GetMyTasksResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminGetTasks(context.Context, *AdminGetTasksRequest) (*AdminGetTasksResponse, error)
	GetMyTasks(context.Context, *GetMyTasksRequest) (*GetMyTasksResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminGetTasks(context.Context, *AdminGetTasksRequest) (*AdminGetTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetTasks not implemented")
}
func (UnimplementedGatewayServer) GetMyTasks(context.Context, *GetMyTasksRequest) (*GetMyTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTasks not implemented")
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

func _Gateway_AdminGetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetTasks(ctx, req.(*AdminGetTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyTasks(ctx, req.(*GetMyTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.task.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminGetTasks",
			Handler:    _Gateway_AdminGetTasks_Handler,
		},
		{
			MethodName: "GetMyTasks",
			Handler:    _Gateway_GetMyTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/task/task.proto",
}