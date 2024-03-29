// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/deviceinfo/deviceinfo.proto

package deviceinfo

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
	Gateway_CreateDeviceInfo_FullMethodName = "/good.gateway.deviceinfo.v1.Gateway/CreateDeviceInfo"
	Gateway_UpdateDeviceInfo_FullMethodName = "/good.gateway.deviceinfo.v1.Gateway/UpdateDeviceInfo"
	Gateway_GetDeviceInfos_FullMethodName   = "/good.gateway.deviceinfo.v1.Gateway/GetDeviceInfos"
	Gateway_DeleteDeviceInfo_FullMethodName = "/good.gateway.deviceinfo.v1.Gateway/DeleteDeviceInfo"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateDeviceInfo(ctx context.Context, in *CreateDeviceInfoRequest, opts ...grpc.CallOption) (*CreateDeviceInfoResponse, error)
	UpdateDeviceInfo(ctx context.Context, in *UpdateDeviceInfoRequest, opts ...grpc.CallOption) (*UpdateDeviceInfoResponse, error)
	GetDeviceInfos(ctx context.Context, in *GetDeviceInfosRequest, opts ...grpc.CallOption) (*GetDeviceInfosResponse, error)
	DeleteDeviceInfo(ctx context.Context, in *DeleteDeviceInfoRequest, opts ...grpc.CallOption) (*DeleteDeviceInfoResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateDeviceInfo(ctx context.Context, in *CreateDeviceInfoRequest, opts ...grpc.CallOption) (*CreateDeviceInfoResponse, error) {
	out := new(CreateDeviceInfoResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateDeviceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateDeviceInfo(ctx context.Context, in *UpdateDeviceInfoRequest, opts ...grpc.CallOption) (*UpdateDeviceInfoResponse, error) {
	out := new(UpdateDeviceInfoResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateDeviceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetDeviceInfos(ctx context.Context, in *GetDeviceInfosRequest, opts ...grpc.CallOption) (*GetDeviceInfosResponse, error) {
	out := new(GetDeviceInfosResponse)
	err := c.cc.Invoke(ctx, Gateway_GetDeviceInfos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteDeviceInfo(ctx context.Context, in *DeleteDeviceInfoRequest, opts ...grpc.CallOption) (*DeleteDeviceInfoResponse, error) {
	out := new(DeleteDeviceInfoResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteDeviceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateDeviceInfo(context.Context, *CreateDeviceInfoRequest) (*CreateDeviceInfoResponse, error)
	UpdateDeviceInfo(context.Context, *UpdateDeviceInfoRequest) (*UpdateDeviceInfoResponse, error)
	GetDeviceInfos(context.Context, *GetDeviceInfosRequest) (*GetDeviceInfosResponse, error)
	DeleteDeviceInfo(context.Context, *DeleteDeviceInfoRequest) (*DeleteDeviceInfoResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateDeviceInfo(context.Context, *CreateDeviceInfoRequest) (*CreateDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceInfo not implemented")
}
func (UnimplementedGatewayServer) UpdateDeviceInfo(context.Context, *UpdateDeviceInfoRequest) (*UpdateDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceInfo not implemented")
}
func (UnimplementedGatewayServer) GetDeviceInfos(context.Context, *GetDeviceInfosRequest) (*GetDeviceInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfos not implemented")
}
func (UnimplementedGatewayServer) DeleteDeviceInfo(context.Context, *DeleteDeviceInfoRequest) (*DeleteDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceInfo not implemented")
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

func _Gateway_CreateDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateDeviceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateDeviceInfo(ctx, req.(*CreateDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateDeviceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateDeviceInfo(ctx, req.(*UpdateDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetDeviceInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetDeviceInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetDeviceInfos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetDeviceInfos(ctx, req.(*GetDeviceInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteDeviceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteDeviceInfo(ctx, req.(*DeleteDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.deviceinfo.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceInfo",
			Handler:    _Gateway_CreateDeviceInfo_Handler,
		},
		{
			MethodName: "UpdateDeviceInfo",
			Handler:    _Gateway_UpdateDeviceInfo_Handler,
		},
		{
			MethodName: "GetDeviceInfos",
			Handler:    _Gateway_GetDeviceInfos_Handler,
		},
		{
			MethodName: "DeleteDeviceInfo",
			Handler:    _Gateway_DeleteDeviceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/deviceinfo/deviceinfo.proto",
}
