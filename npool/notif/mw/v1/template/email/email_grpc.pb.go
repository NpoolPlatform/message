// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/template/email/email.proto

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

const (
	Middleware_GetEmailTemplate_FullMethodName        = "/notif.middleware.template.email.v1.Middleware/GetEmailTemplate"
	Middleware_GetEmailTemplates_FullMethodName       = "/notif.middleware.template.email.v1.Middleware/GetEmailTemplates"
	Middleware_CreateEmailTemplate_FullMethodName     = "/notif.middleware.template.email.v1.Middleware/CreateEmailTemplate"
	Middleware_CreateEmailTemplates_FullMethodName    = "/notif.middleware.template.email.v1.Middleware/CreateEmailTemplates"
	Middleware_UpdateEmailTemplate_FullMethodName     = "/notif.middleware.template.email.v1.Middleware/UpdateEmailTemplate"
	Middleware_ExistEmailTemplate_FullMethodName      = "/notif.middleware.template.email.v1.Middleware/ExistEmailTemplate"
	Middleware_ExistEmailTemplateConds_FullMethodName = "/notif.middleware.template.email.v1.Middleware/ExistEmailTemplateConds"
	Middleware_DeleteEmailTemplate_FullMethodName     = "/notif.middleware.template.email.v1.Middleware/DeleteEmailTemplate"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetEmailTemplate(ctx context.Context, in *GetEmailTemplateRequest, opts ...grpc.CallOption) (*GetEmailTemplateResponse, error)
	GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesRequest, opts ...grpc.CallOption) (*GetEmailTemplatesResponse, error)
	CreateEmailTemplate(ctx context.Context, in *CreateEmailTemplateRequest, opts ...grpc.CallOption) (*CreateEmailTemplateResponse, error)
	CreateEmailTemplates(ctx context.Context, in *CreateEmailTemplatesRequest, opts ...grpc.CallOption) (*CreateEmailTemplatesResponse, error)
	UpdateEmailTemplate(ctx context.Context, in *UpdateEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateEmailTemplateResponse, error)
	ExistEmailTemplate(ctx context.Context, in *ExistEmailTemplateRequest, opts ...grpc.CallOption) (*ExistEmailTemplateResponse, error)
	ExistEmailTemplateConds(ctx context.Context, in *ExistEmailTemplateCondsRequest, opts ...grpc.CallOption) (*ExistEmailTemplateCondsResponse, error)
	DeleteEmailTemplate(ctx context.Context, in *DeleteEmailTemplateRequest, opts ...grpc.CallOption) (*DeleteEmailTemplateResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetEmailTemplate(ctx context.Context, in *GetEmailTemplateRequest, opts ...grpc.CallOption) (*GetEmailTemplateResponse, error) {
	out := new(GetEmailTemplateResponse)
	err := c.cc.Invoke(ctx, Middleware_GetEmailTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetEmailTemplates(ctx context.Context, in *GetEmailTemplatesRequest, opts ...grpc.CallOption) (*GetEmailTemplatesResponse, error) {
	out := new(GetEmailTemplatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetEmailTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateEmailTemplate(ctx context.Context, in *CreateEmailTemplateRequest, opts ...grpc.CallOption) (*CreateEmailTemplateResponse, error) {
	out := new(CreateEmailTemplateResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateEmailTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateEmailTemplates(ctx context.Context, in *CreateEmailTemplatesRequest, opts ...grpc.CallOption) (*CreateEmailTemplatesResponse, error) {
	out := new(CreateEmailTemplatesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateEmailTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateEmailTemplate(ctx context.Context, in *UpdateEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateEmailTemplateResponse, error) {
	out := new(UpdateEmailTemplateResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateEmailTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistEmailTemplate(ctx context.Context, in *ExistEmailTemplateRequest, opts ...grpc.CallOption) (*ExistEmailTemplateResponse, error) {
	out := new(ExistEmailTemplateResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistEmailTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistEmailTemplateConds(ctx context.Context, in *ExistEmailTemplateCondsRequest, opts ...grpc.CallOption) (*ExistEmailTemplateCondsResponse, error) {
	out := new(ExistEmailTemplateCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistEmailTemplateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteEmailTemplate(ctx context.Context, in *DeleteEmailTemplateRequest, opts ...grpc.CallOption) (*DeleteEmailTemplateResponse, error) {
	out := new(DeleteEmailTemplateResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteEmailTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetEmailTemplate(context.Context, *GetEmailTemplateRequest) (*GetEmailTemplateResponse, error)
	GetEmailTemplates(context.Context, *GetEmailTemplatesRequest) (*GetEmailTemplatesResponse, error)
	CreateEmailTemplate(context.Context, *CreateEmailTemplateRequest) (*CreateEmailTemplateResponse, error)
	CreateEmailTemplates(context.Context, *CreateEmailTemplatesRequest) (*CreateEmailTemplatesResponse, error)
	UpdateEmailTemplate(context.Context, *UpdateEmailTemplateRequest) (*UpdateEmailTemplateResponse, error)
	ExistEmailTemplate(context.Context, *ExistEmailTemplateRequest) (*ExistEmailTemplateResponse, error)
	ExistEmailTemplateConds(context.Context, *ExistEmailTemplateCondsRequest) (*ExistEmailTemplateCondsResponse, error)
	DeleteEmailTemplate(context.Context, *DeleteEmailTemplateRequest) (*DeleteEmailTemplateResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetEmailTemplate(context.Context, *GetEmailTemplateRequest) (*GetEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailTemplate not implemented")
}
func (UnimplementedMiddlewareServer) GetEmailTemplates(context.Context, *GetEmailTemplatesRequest) (*GetEmailTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailTemplates not implemented")
}
func (UnimplementedMiddlewareServer) CreateEmailTemplate(context.Context, *CreateEmailTemplateRequest) (*CreateEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmailTemplate not implemented")
}
func (UnimplementedMiddlewareServer) CreateEmailTemplates(context.Context, *CreateEmailTemplatesRequest) (*CreateEmailTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmailTemplates not implemented")
}
func (UnimplementedMiddlewareServer) UpdateEmailTemplate(context.Context, *UpdateEmailTemplateRequest) (*UpdateEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailTemplate not implemented")
}
func (UnimplementedMiddlewareServer) ExistEmailTemplate(context.Context, *ExistEmailTemplateRequest) (*ExistEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistEmailTemplate not implemented")
}
func (UnimplementedMiddlewareServer) ExistEmailTemplateConds(context.Context, *ExistEmailTemplateCondsRequest) (*ExistEmailTemplateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistEmailTemplateConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteEmailTemplate(context.Context, *DeleteEmailTemplateRequest) (*DeleteEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailTemplate not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

func _Middleware_GetEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetEmailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetEmailTemplate(ctx, req.(*GetEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetEmailTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetEmailTemplates(ctx, req.(*GetEmailTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateEmailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateEmailTemplate(ctx, req.(*CreateEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateEmailTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateEmailTemplates(ctx, req.(*CreateEmailTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateEmailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateEmailTemplate(ctx, req.(*UpdateEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistEmailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistEmailTemplate(ctx, req.(*ExistEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistEmailTemplateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistEmailTemplateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistEmailTemplateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistEmailTemplateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistEmailTemplateConds(ctx, req.(*ExistEmailTemplateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteEmailTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteEmailTemplate(ctx, req.(*DeleteEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.template.email.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmailTemplate",
			Handler:    _Middleware_GetEmailTemplate_Handler,
		},
		{
			MethodName: "GetEmailTemplates",
			Handler:    _Middleware_GetEmailTemplates_Handler,
		},
		{
			MethodName: "CreateEmailTemplate",
			Handler:    _Middleware_CreateEmailTemplate_Handler,
		},
		{
			MethodName: "CreateEmailTemplates",
			Handler:    _Middleware_CreateEmailTemplates_Handler,
		},
		{
			MethodName: "UpdateEmailTemplate",
			Handler:    _Middleware_UpdateEmailTemplate_Handler,
		},
		{
			MethodName: "ExistEmailTemplate",
			Handler:    _Middleware_ExistEmailTemplate_Handler,
		},
		{
			MethodName: "ExistEmailTemplateConds",
			Handler:    _Middleware_ExistEmailTemplateConds_Handler,
		},
		{
			MethodName: "DeleteEmailTemplate",
			Handler:    _Middleware_DeleteEmailTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/template/email/email.proto",
}
