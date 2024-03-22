// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/device/device.proto

package device

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
	Gateway_CreateDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/CreateDeviceType"
	Gateway_UpdateDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/UpdateDeviceType"
	Gateway_GetDeviceTypes_FullMethodName   = "/good.gateway.device.v1.Gateway/GetDeviceTypes"
	Gateway_DeleteDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/DeleteDeviceType"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateDeviceType(ctx context.Context, in *CreateDeviceTypeRequest, opts ...grpc.CallOption) (*CreateDeviceTypeResponse, error)
	UpdateDeviceType(ctx context.Context, in *UpdateDeviceTypeRequest, opts ...grpc.CallOption) (*UpdateDeviceTypeResponse, error)
	GetDeviceTypes(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*GetDeviceTypesResponse, error)
	DeleteDeviceType(ctx context.Context, in *DeleteDeviceTypeRequest, opts ...grpc.CallOption) (*DeleteDeviceTypeResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateDeviceType(ctx context.Context, in *CreateDeviceTypeRequest, opts ...grpc.CallOption) (*CreateDeviceTypeResponse, error) {
	out := new(CreateDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateDeviceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateDeviceType(ctx context.Context, in *UpdateDeviceTypeRequest, opts ...grpc.CallOption) (*UpdateDeviceTypeResponse, error) {
	out := new(UpdateDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateDeviceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetDeviceTypes(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*GetDeviceTypesResponse, error) {
	out := new(GetDeviceTypesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetDeviceTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteDeviceType(ctx context.Context, in *DeleteDeviceTypeRequest, opts ...grpc.CallOption) (*DeleteDeviceTypeResponse, error) {
	out := new(DeleteDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteDeviceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateDeviceType(context.Context, *CreateDeviceTypeRequest) (*CreateDeviceTypeResponse, error)
	UpdateDeviceType(context.Context, *UpdateDeviceTypeRequest) (*UpdateDeviceTypeResponse, error)
	GetDeviceTypes(context.Context, *GetDeviceTypesRequest) (*GetDeviceTypesResponse, error)
	DeleteDeviceType(context.Context, *DeleteDeviceTypeRequest) (*DeleteDeviceTypeResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateDeviceType(context.Context, *CreateDeviceTypeRequest) (*CreateDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceType not implemented")
}
func (UnimplementedGatewayServer) UpdateDeviceType(context.Context, *UpdateDeviceTypeRequest) (*UpdateDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceType not implemented")
}
func (UnimplementedGatewayServer) GetDeviceTypes(context.Context, *GetDeviceTypesRequest) (*GetDeviceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceTypes not implemented")
}
func (UnimplementedGatewayServer) DeleteDeviceType(context.Context, *DeleteDeviceTypeRequest) (*DeleteDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceType not implemented")
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

func _Gateway_CreateDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateDeviceType(ctx, req.(*CreateDeviceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateDeviceType(ctx, req.(*UpdateDeviceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetDeviceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetDeviceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetDeviceTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetDeviceTypes(ctx, req.(*GetDeviceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteDeviceType(ctx, req.(*DeleteDeviceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.device.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceType",
			Handler:    _Gateway_CreateDeviceType_Handler,
		},
		{
			MethodName: "UpdateDeviceType",
			Handler:    _Gateway_UpdateDeviceType_Handler,
		},
		{
			MethodName: "GetDeviceTypes",
			Handler:    _Gateway_GetDeviceTypes_Handler,
		},
		{
			MethodName: "DeleteDeviceType",
			Handler:    _Gateway_DeleteDeviceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/device/device.proto",
}
