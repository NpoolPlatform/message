// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/third/gw/v1/template/email/email.proto

package email

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
	CreateEmailTemplate(ctx context.Context, in *CreateEmailTemplateRequest, opts ...grpc.CallOption) (*CreateEmailTemplateResponse, error)
	CreateAppEmailTemplate(ctx context.Context, in *CreateAppEmailTemplateRequest, opts ...grpc.CallOption) (*CreateAppEmailTemplateResponse, error)
	GetEmailTemplate(ctx context.Context, in *GetEmailTemplateRequest, opts ...grpc.CallOption) (*GetEmailTemplateResponse, error)
	GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesRequest, opts ...grpc.CallOption) (*GetEmailTemplatesResponse, error)
	GetAppEmailTemplates(ctx context.Context, in *GetAppEmailTemplatesRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesResponse, error)
	UpdateEmailTemplate(ctx context.Context, in *UpdateEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateEmailTemplateResponse, error)
	UpdateAppEmailTemplate(ctx context.Context, in *UpdateAppEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateAppEmailTemplateResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateEmailTemplate(ctx context.Context, in *CreateEmailTemplateRequest, opts ...grpc.CallOption) (*CreateEmailTemplateResponse, error) {
	out := new(CreateEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/CreateEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppEmailTemplate(ctx context.Context, in *CreateAppEmailTemplateRequest, opts ...grpc.CallOption) (*CreateAppEmailTemplateResponse, error) {
	out := new(CreateAppEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/CreateAppEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetEmailTemplate(ctx context.Context, in *GetEmailTemplateRequest, opts ...grpc.CallOption) (*GetEmailTemplateResponse, error) {
	out := new(GetEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/GetEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesRequest, opts ...grpc.CallOption) (*GetEmailTemplatesResponse, error) {
	out := new(GetEmailTemplatesResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/GetEmailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppEmailTemplates(ctx context.Context, in *GetAppEmailTemplatesRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesResponse, error) {
	out := new(GetAppEmailTemplatesResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/GetAppEmailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateEmailTemplate(ctx context.Context, in *UpdateEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateEmailTemplateResponse, error) {
	out := new(UpdateEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/UpdateEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppEmailTemplate(ctx context.Context, in *UpdateAppEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateAppEmailTemplateResponse, error) {
	out := new(UpdateAppEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.template.email.v1.Gateway/UpdateAppEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateEmailTemplate(context.Context, *CreateEmailTemplateRequest) (*CreateEmailTemplateResponse, error)
	CreateAppEmailTemplate(context.Context, *CreateAppEmailTemplateRequest) (*CreateAppEmailTemplateResponse, error)
	GetEmailTemplate(context.Context, *GetEmailTemplateRequest) (*GetEmailTemplateResponse, error)
	GetEmailTemplates(context.Context, *GetEmailTemplatesRequest) (*GetEmailTemplatesResponse, error)
	GetAppEmailTemplates(context.Context, *GetAppEmailTemplatesRequest) (*GetAppEmailTemplatesResponse, error)
	UpdateEmailTemplate(context.Context, *UpdateEmailTemplateRequest) (*UpdateEmailTemplateResponse, error)
	UpdateAppEmailTemplate(context.Context, *UpdateAppEmailTemplateRequest) (*UpdateAppEmailTemplateResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateEmailTemplate(context.Context, *CreateEmailTemplateRequest) (*CreateEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmailTemplate not implemented")
}
func (UnimplementedGatewayServer) CreateAppEmailTemplate(context.Context, *CreateAppEmailTemplateRequest) (*CreateAppEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppEmailTemplate not implemented")
}
func (UnimplementedGatewayServer) GetEmailTemplate(context.Context, *GetEmailTemplateRequest) (*GetEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailTemplate not implemented")
}
func (UnimplementedGatewayServer) GetEmailTemplates(context.Context, *GetEmailTemplatesRequest) (*GetEmailTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailTemplates not implemented")
}
func (UnimplementedGatewayServer) GetAppEmailTemplates(context.Context, *GetAppEmailTemplatesRequest) (*GetAppEmailTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppEmailTemplates not implemented")
}
func (UnimplementedGatewayServer) UpdateEmailTemplate(context.Context, *UpdateEmailTemplateRequest) (*UpdateEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailTemplate not implemented")
}
func (UnimplementedGatewayServer) UpdateAppEmailTemplate(context.Context, *UpdateAppEmailTemplateRequest) (*UpdateAppEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppEmailTemplate not implemented")
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

func _Gateway_CreateEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/CreateEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateEmailTemplate(ctx, req.(*CreateEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/CreateAppEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppEmailTemplate(ctx, req.(*CreateAppEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/GetEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetEmailTemplate(ctx, req.(*GetEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/GetEmailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetEmailTemplates(ctx, req.(*GetEmailTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppEmailTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/GetAppEmailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppEmailTemplates(ctx, req.(*GetAppEmailTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/UpdateEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateEmailTemplate(ctx, req.(*UpdateEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.template.email.v1.Gateway/UpdateAppEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppEmailTemplate(ctx, req.(*UpdateAppEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "third.gateway.template.email.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmailTemplate",
			Handler:    _Gateway_CreateEmailTemplate_Handler,
		},
		{
			MethodName: "CreateAppEmailTemplate",
			Handler:    _Gateway_CreateAppEmailTemplate_Handler,
		},
		{
			MethodName: "GetEmailTemplate",
			Handler:    _Gateway_GetEmailTemplate_Handler,
		},
		{
			MethodName: "GetEmailTemplates",
			Handler:    _Gateway_GetEmailTemplates_Handler,
		},
		{
			MethodName: "GetAppEmailTemplates",
			Handler:    _Gateway_GetAppEmailTemplates_Handler,
		},
		{
			MethodName: "UpdateEmailTemplate",
			Handler:    _Gateway_UpdateEmailTemplate_Handler,
		},
		{
			MethodName: "UpdateAppEmailTemplate",
			Handler:    _Gateway_UpdateAppEmailTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/third/gw/v1/template/email/email.proto",
}
