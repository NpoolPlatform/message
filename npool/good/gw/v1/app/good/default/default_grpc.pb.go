// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/app/good/default/default.proto

package _default

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
	CreateDefault(ctx context.Context, in *CreateDefaultRequest, opts ...grpc.CallOption) (*CreateDefaultResponse, error)
	GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...grpc.CallOption) (*GetDefaultsResponse, error)
	DeleteDefault(ctx context.Context, in *DeleteDefaultRequest, opts ...grpc.CallOption) (*DeleteDefaultResponse, error)
	UpdateDefault(ctx context.Context, in *UpdateDefaultRequest, opts ...grpc.CallOption) (*UpdateDefaultResponse, error)
	// Run by church admin
	AdminCreateDefault(ctx context.Context, in *AdminCreateDefaultRequest, opts ...grpc.CallOption) (*AdminCreateDefaultResponse, error)
	AdminGetDefaults(ctx context.Context, in *AdminGetDefaultsRequest, opts ...grpc.CallOption) (*AdminGetDefaultsResponse, error)
	AdminDeleteDefault(ctx context.Context, in *AdminDeleteDefaultRequest, opts ...grpc.CallOption) (*AdminDeleteDefaultResponse, error)
	AdminUpdateDefault(ctx context.Context, in *AdminUpdateDefaultRequest, opts ...grpc.CallOption) (*AdminUpdateDefaultResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateDefault(ctx context.Context, in *CreateDefaultRequest, opts ...grpc.CallOption) (*CreateDefaultResponse, error) {
	out := new(CreateDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/CreateDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetDefaults(ctx context.Context, in *GetDefaultsRequest, opts ...grpc.CallOption) (*GetDefaultsResponse, error) {
	out := new(GetDefaultsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/GetDefaults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteDefault(ctx context.Context, in *DeleteDefaultRequest, opts ...grpc.CallOption) (*DeleteDefaultResponse, error) {
	out := new(DeleteDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/DeleteDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateDefault(ctx context.Context, in *UpdateDefaultRequest, opts ...grpc.CallOption) (*UpdateDefaultResponse, error) {
	out := new(UpdateDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/UpdateDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateDefault(ctx context.Context, in *AdminCreateDefaultRequest, opts ...grpc.CallOption) (*AdminCreateDefaultResponse, error) {
	out := new(AdminCreateDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/AdminCreateDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetDefaults(ctx context.Context, in *AdminGetDefaultsRequest, opts ...grpc.CallOption) (*AdminGetDefaultsResponse, error) {
	out := new(AdminGetDefaultsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/AdminGetDefaults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteDefault(ctx context.Context, in *AdminDeleteDefaultRequest, opts ...grpc.CallOption) (*AdminDeleteDefaultResponse, error) {
	out := new(AdminDeleteDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/AdminDeleteDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateDefault(ctx context.Context, in *AdminUpdateDefaultRequest, opts ...grpc.CallOption) (*AdminUpdateDefaultResponse, error) {
	out := new(AdminUpdateDefaultResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.default1.v1.Gateway/AdminUpdateDefault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateDefault(context.Context, *CreateDefaultRequest) (*CreateDefaultResponse, error)
	GetDefaults(context.Context, *GetDefaultsRequest) (*GetDefaultsResponse, error)
	DeleteDefault(context.Context, *DeleteDefaultRequest) (*DeleteDefaultResponse, error)
	UpdateDefault(context.Context, *UpdateDefaultRequest) (*UpdateDefaultResponse, error)
	// Run by church admin
	AdminCreateDefault(context.Context, *AdminCreateDefaultRequest) (*AdminCreateDefaultResponse, error)
	AdminGetDefaults(context.Context, *AdminGetDefaultsRequest) (*AdminGetDefaultsResponse, error)
	AdminDeleteDefault(context.Context, *AdminDeleteDefaultRequest) (*AdminDeleteDefaultResponse, error)
	AdminUpdateDefault(context.Context, *AdminUpdateDefaultRequest) (*AdminUpdateDefaultResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateDefault(context.Context, *CreateDefaultRequest) (*CreateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDefault not implemented")
}
func (UnimplementedGatewayServer) GetDefaults(context.Context, *GetDefaultsRequest) (*GetDefaultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaults not implemented")
}
func (UnimplementedGatewayServer) DeleteDefault(context.Context, *DeleteDefaultRequest) (*DeleteDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDefault not implemented")
}
func (UnimplementedGatewayServer) UpdateDefault(context.Context, *UpdateDefaultRequest) (*UpdateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDefault not implemented")
}
func (UnimplementedGatewayServer) AdminCreateDefault(context.Context, *AdminCreateDefaultRequest) (*AdminCreateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateDefault not implemented")
}
func (UnimplementedGatewayServer) AdminGetDefaults(context.Context, *AdminGetDefaultsRequest) (*AdminGetDefaultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetDefaults not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteDefault(context.Context, *AdminDeleteDefaultRequest) (*AdminDeleteDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteDefault not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateDefault(context.Context, *AdminUpdateDefaultRequest) (*AdminUpdateDefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateDefault not implemented")
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

func _Gateway_CreateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/CreateDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateDefault(ctx, req.(*CreateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/GetDefaults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetDefaults(ctx, req.(*GetDefaultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/DeleteDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteDefault(ctx, req.(*DeleteDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/UpdateDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateDefault(ctx, req.(*UpdateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/AdminCreateDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateDefault(ctx, req.(*AdminCreateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetDefaults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetDefaultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetDefaults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/AdminGetDefaults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetDefaults(ctx, req.(*AdminGetDefaultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/AdminDeleteDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteDefault(ctx, req.(*AdminDeleteDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateDefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.default1.v1.Gateway/AdminUpdateDefault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateDefault(ctx, req.(*AdminUpdateDefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.app.good1.default1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDefault",
			Handler:    _Gateway_CreateDefault_Handler,
		},
		{
			MethodName: "GetDefaults",
			Handler:    _Gateway_GetDefaults_Handler,
		},
		{
			MethodName: "DeleteDefault",
			Handler:    _Gateway_DeleteDefault_Handler,
		},
		{
			MethodName: "UpdateDefault",
			Handler:    _Gateway_UpdateDefault_Handler,
		},
		{
			MethodName: "AdminCreateDefault",
			Handler:    _Gateway_AdminCreateDefault_Handler,
		},
		{
			MethodName: "AdminGetDefaults",
			Handler:    _Gateway_AdminGetDefaults_Handler,
		},
		{
			MethodName: "AdminDeleteDefault",
			Handler:    _Gateway_AdminDeleteDefault_Handler,
		},
		{
			MethodName: "AdminUpdateDefault",
			Handler:    _Gateway_AdminUpdateDefault_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/app/good/default/default.proto",
}
