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
	Gateway_AdminCreateDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/AdminCreateDeviceType"
	Gateway_AdminUpdateDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/AdminUpdateDeviceType"
	Gateway_GetDeviceTypes_FullMethodName        = "/good.gateway.device.v1.Gateway/GetDeviceTypes"
	Gateway_AdminDeleteDeviceType_FullMethodName = "/good.gateway.device.v1.Gateway/AdminDeleteDeviceType"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateDeviceType(ctx context.Context, in *AdminCreateDeviceTypeRequest, opts ...grpc.CallOption) (*AdminCreateDeviceTypeResponse, error)
	AdminUpdateDeviceType(ctx context.Context, in *AdminUpdateDeviceTypeRequest, opts ...grpc.CallOption) (*AdminUpdateDeviceTypeResponse, error)
	GetDeviceTypes(ctx context.Context, in *GetDeviceTypesRequest, opts ...grpc.CallOption) (*GetDeviceTypesResponse, error)
	AdminDeleteDeviceType(ctx context.Context, in *AdminDeleteDeviceTypeRequest, opts ...grpc.CallOption) (*AdminDeleteDeviceTypeResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateDeviceType(ctx context.Context, in *AdminCreateDeviceTypeRequest, opts ...grpc.CallOption) (*AdminCreateDeviceTypeResponse, error) {
	out := new(AdminCreateDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateDeviceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateDeviceType(ctx context.Context, in *AdminUpdateDeviceTypeRequest, opts ...grpc.CallOption) (*AdminUpdateDeviceTypeResponse, error) {
	out := new(AdminUpdateDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateDeviceType_FullMethodName, in, out, opts...)
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

func (c *gatewayClient) AdminDeleteDeviceType(ctx context.Context, in *AdminDeleteDeviceTypeRequest, opts ...grpc.CallOption) (*AdminDeleteDeviceTypeResponse, error) {
	out := new(AdminDeleteDeviceTypeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteDeviceType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateDeviceType(context.Context, *AdminCreateDeviceTypeRequest) (*AdminCreateDeviceTypeResponse, error)
	AdminUpdateDeviceType(context.Context, *AdminUpdateDeviceTypeRequest) (*AdminUpdateDeviceTypeResponse, error)
	GetDeviceTypes(context.Context, *GetDeviceTypesRequest) (*GetDeviceTypesResponse, error)
	AdminDeleteDeviceType(context.Context, *AdminDeleteDeviceTypeRequest) (*AdminDeleteDeviceTypeResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreateDeviceType(context.Context, *AdminCreateDeviceTypeRequest) (*AdminCreateDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateDeviceType not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateDeviceType(context.Context, *AdminUpdateDeviceTypeRequest) (*AdminUpdateDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateDeviceType not implemented")
}
func (UnimplementedGatewayServer) GetDeviceTypes(context.Context, *GetDeviceTypesRequest) (*GetDeviceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceTypes not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteDeviceType(context.Context, *AdminDeleteDeviceTypeRequest) (*AdminDeleteDeviceTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteDeviceType not implemented")
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

func _Gateway_AdminCreateDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateDeviceType(ctx, req.(*AdminCreateDeviceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateDeviceType(ctx, req.(*AdminUpdateDeviceTypeRequest))
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

func _Gateway_AdminDeleteDeviceType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteDeviceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteDeviceType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteDeviceType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteDeviceType(ctx, req.(*AdminDeleteDeviceTypeRequest))
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
			MethodName: "AdminCreateDeviceType",
			Handler:    _Gateway_AdminCreateDeviceType_Handler,
		},
		{
			MethodName: "AdminUpdateDeviceType",
			Handler:    _Gateway_AdminUpdateDeviceType_Handler,
		},
		{
			MethodName: "GetDeviceTypes",
			Handler:    _Gateway_GetDeviceTypes_Handler,
		},
		{
			MethodName: "AdminDeleteDeviceType",
			Handler:    _Gateway_AdminDeleteDeviceType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/device/device.proto",
}