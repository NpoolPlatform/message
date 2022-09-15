// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/third/gw/v1/template/sms/sms.proto

package sms

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
	CreateSMSTemplate(ctx context.Context, in *CreateSMSTemplateRequest, opts ...grpc.CallOption) (*CreateSMSTemplateResponse, error)
	CreateAppSMSTemplate(ctx context.Context, in *CreateAppSMSTemplateRequest, opts ...grpc.CallOption) (*CreateAppSMSTemplateResponse, error)
	GetSMSTemplate(ctx context.Context, in *GetSMSTemplateRequest, opts ...grpc.CallOption) (*GetSMSTemplateResponse, error)
	GetAppSMSTemplate(ctx context.Context, in *GetAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateResponse, error)
	GetNAppSMSTemplate(ctx context.Context, in *GetNAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetNAppSMSTemplateResponse, error)
	GetSMSTemplates(ctx context.Context, in *GetSMSTemplatesRequest, opts ...grpc.CallOption) (*GetSMSTemplatesResponse, error)
	UpdateSMSTemplate(ctx context.Context, in *UpdateSMSTemplateRequest, opts ...grpc.CallOption) (*UpdateSMSTemplateResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateSMSTemplate(ctx context.Context, in *CreateSMSTemplateRequest, opts ...grpc.CallOption) (*CreateSMSTemplateResponse, error) {
	out := new(CreateSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/CreateSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppSMSTemplate(ctx context.Context, in *CreateAppSMSTemplateRequest, opts ...grpc.CallOption) (*CreateAppSMSTemplateResponse, error) {
	out := new(CreateAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/CreateAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSMSTemplate(ctx context.Context, in *GetSMSTemplateRequest, opts ...grpc.CallOption) (*GetSMSTemplateResponse, error) {
	out := new(GetSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/GetSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppSMSTemplate(ctx context.Context, in *GetAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateResponse, error) {
	out := new(GetAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/GetAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetNAppSMSTemplate(ctx context.Context, in *GetNAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetNAppSMSTemplateResponse, error) {
	out := new(GetNAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/GetNAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetSMSTemplates(ctx context.Context, in *GetSMSTemplatesRequest, opts ...grpc.CallOption) (*GetSMSTemplatesResponse, error) {
	out := new(GetSMSTemplatesResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/GetSMSTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateSMSTemplate(ctx context.Context, in *UpdateSMSTemplateRequest, opts ...grpc.CallOption) (*UpdateSMSTemplateResponse, error) {
	out := new(UpdateSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.sms.v1.Gateway/UpdateSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateSMSTemplate(context.Context, *CreateSMSTemplateRequest) (*CreateSMSTemplateResponse, error)
	CreateAppSMSTemplate(context.Context, *CreateAppSMSTemplateRequest) (*CreateAppSMSTemplateResponse, error)
	GetSMSTemplate(context.Context, *GetSMSTemplateRequest) (*GetSMSTemplateResponse, error)
	GetAppSMSTemplate(context.Context, *GetAppSMSTemplateRequest) (*GetAppSMSTemplateResponse, error)
	GetNAppSMSTemplate(context.Context, *GetNAppSMSTemplateRequest) (*GetNAppSMSTemplateResponse, error)
	GetSMSTemplates(context.Context, *GetSMSTemplatesRequest) (*GetSMSTemplatesResponse, error)
	UpdateSMSTemplate(context.Context, *UpdateSMSTemplateRequest) (*UpdateSMSTemplateResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateSMSTemplate(context.Context, *CreateSMSTemplateRequest) (*CreateSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSMSTemplate not implemented")
}
func (UnimplementedGatewayServer) CreateAppSMSTemplate(context.Context, *CreateAppSMSTemplateRequest) (*CreateAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppSMSTemplate not implemented")
}
func (UnimplementedGatewayServer) GetSMSTemplate(context.Context, *GetSMSTemplateRequest) (*GetSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSMSTemplate not implemented")
}
func (UnimplementedGatewayServer) GetAppSMSTemplate(context.Context, *GetAppSMSTemplateRequest) (*GetAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSMSTemplate not implemented")
}
func (UnimplementedGatewayServer) GetNAppSMSTemplate(context.Context, *GetNAppSMSTemplateRequest) (*GetNAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppSMSTemplate not implemented")
}
func (UnimplementedGatewayServer) GetSMSTemplates(context.Context, *GetSMSTemplatesRequest) (*GetSMSTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSMSTemplates not implemented")
}
func (UnimplementedGatewayServer) UpdateSMSTemplate(context.Context, *UpdateSMSTemplateRequest) (*UpdateSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSMSTemplate not implemented")
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

func _Gateway_CreateSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/CreateSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateSMSTemplate(ctx, req.(*CreateSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/CreateAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppSMSTemplate(ctx, req.(*CreateAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/GetSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSMSTemplate(ctx, req.(*GetSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/GetAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppSMSTemplate(ctx, req.(*GetAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetNAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/GetNAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppSMSTemplate(ctx, req.(*GetNAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetSMSTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSMSTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetSMSTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/GetSMSTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetSMSTemplates(ctx, req.(*GetSMSTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.sms.v1.Gateway/UpdateSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateSMSTemplate(ctx, req.(*UpdateSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "third.gateway.template.sms.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSMSTemplate",
			Handler:    _Gateway_CreateSMSTemplate_Handler,
		},
		{
			MethodName: "CreateAppSMSTemplate",
			Handler:    _Gateway_CreateAppSMSTemplate_Handler,
		},
		{
			MethodName: "GetSMSTemplate",
			Handler:    _Gateway_GetSMSTemplate_Handler,
		},
		{
			MethodName: "GetAppSMSTemplate",
			Handler:    _Gateway_GetAppSMSTemplate_Handler,
		},
		{
			MethodName: "GetNAppSMSTemplate",
			Handler:    _Gateway_GetNAppSMSTemplate_Handler,
		},
		{
			MethodName: "GetSMSTemplates",
			Handler:    _Gateway_GetSMSTemplates_Handler,
		},
		{
			MethodName: "UpdateSMSTemplate",
			Handler:    _Gateway_UpdateSMSTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/third/gw/v1/template/sms/sms.proto",
}
