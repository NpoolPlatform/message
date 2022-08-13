// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/app/app.proto

package app

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
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error)
	GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	// Super admin api
	GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
	GetUserApps(ctx context.Context, in *GetUserAppsRequest, opts ...grpc.CallOption) (*GetUserAppsResponse, error)
	GetSignMethods(ctx context.Context, in *GetSignMethodsRequest, opts ...grpc.CallOption) (*GetSignMethodsResponse, error)
	GetRecaptchas(ctx context.Context, in *GetRecaptchasRequest, opts ...grpc.CallOption) (*GetRecaptchasResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error) {
	out := new(UpdateAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetApps(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/GetApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetUserApps(ctx context.Context, in *GetUserAppsRequest, opts ...grpc.CallOption) (*GetUserAppsResponse, error) {
	out := new(GetUserAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/GetUserApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSignMethods(ctx context.Context, in *GetSignMethodsRequest, opts ...grpc.CallOption) (*GetSignMethodsResponse, error) {
	out := new(GetSignMethodsResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/GetSignMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRecaptchas(ctx context.Context, in *GetRecaptchasRequest, opts ...grpc.CallOption) (*GetRecaptchasResponse, error) {
	out := new(GetRecaptchasResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.app.v1.Gateway/GetRecaptchas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error)
	GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error)
	// Super admin api
	GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	GetUserApps(context.Context, *GetUserAppsRequest) (*GetUserAppsResponse, error)
	GetSignMethods(context.Context, *GetSignMethodsRequest) (*GetSignMethodsResponse, error)
	GetRecaptchas(context.Context, *GetRecaptchasRequest) (*GetRecaptchasResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedGatewayServer) UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (UnimplementedGatewayServer) GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (UnimplementedGatewayServer) GetApps(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedGatewayServer) GetUserApps(context.Context, *GetUserAppsRequest) (*GetUserAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserApps not implemented")
}
func (UnimplementedGatewayServer) GetSignMethods(context.Context, *GetSignMethodsRequest) (*GetSignMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignMethods not implemented")
}
func (UnimplementedGatewayServer) GetRecaptchas(context.Context, *GetRecaptchasRequest) (*GetRecaptchasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecaptchas not implemented")
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

func _Gateway_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateApp(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetApp(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/GetApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetApps(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetUserApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetUserApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/GetUserApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetUserApps(ctx, req.(*GetUserAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSignMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSignMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/GetSignMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSignMethods(ctx, req.(*GetSignMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetRecaptchas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecaptchasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRecaptchas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.app.v1.Gateway/GetRecaptchas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRecaptchas(ctx, req.(*GetRecaptchasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.app.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _Gateway_CreateApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _Gateway_UpdateApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _Gateway_GetApp_Handler,
		},
		{
			MethodName: "GetApps",
			Handler:    _Gateway_GetApps_Handler,
		},
		{
			MethodName: "GetUserApps",
			Handler:    _Gateway_GetUserApps_Handler,
		},
		{
			MethodName: "GetSignMethods",
			Handler:    _Gateway_GetSignMethods_Handler,
		},
		{
			MethodName: "GetRecaptchas",
			Handler:    _Gateway_GetRecaptchas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/app/app.proto",
}
