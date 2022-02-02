// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/thirdgateway/thirdgateway.proto

package thirdgateway

import (
	context "context"
	npool "github.com/NpoolPlatform/message/npool"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ThirdGatewayClient is the client API for ThirdGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdGatewayClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	CreateAppSMSTemplate(ctx context.Context, in *CreateAppSMSTemplateRequest, opts ...grpc.CallOption) (*CreateAppSMSTemplateResponse, error)
	GetAppSMSTemplate(ctx context.Context, in *GetAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateResponse, error)
	UpdateAppSMSTemplate(ctx context.Context, in *UpdateAppSMSTemplateRequest, opts ...grpc.CallOption) (*UpdateAppSMSTemplateResponse, error)
	GetAppSMSTemplatesByApp(ctx context.Context, in *GetAppSMSTemplatesByAppRequest, opts ...grpc.CallOption) (*GetAppSMSTemplatesByAppResponse, error)
	GetAppSMSTemplatesByOtherApp(ctx context.Context, in *GetAppSMSTemplatesByOtherAppRequest, opts ...grpc.CallOption) (*GetAppSMSTemplatesByOtherAppResponse, error)
	GetAppSMSTemplateByAppLangUsedFor(ctx context.Context, in *GetAppSMSTemplateByAppLangUsedForRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateByAppLangUsedForResponse, error)
	CreateAppEmailTemplate(ctx context.Context, in *CreateAppEmailTemplateRequest, opts ...grpc.CallOption) (*CreateAppEmailTemplateResponse, error)
	GetAppEmailTemplate(ctx context.Context, in *GetAppEmailTemplateRequest, opts ...grpc.CallOption) (*GetAppEmailTemplateResponse, error)
	UpdateAppEmailTemplate(ctx context.Context, in *UpdateAppEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateAppEmailTemplateResponse, error)
	GetAppEmailTemplatesByApp(ctx context.Context, in *GetAppEmailTemplatesByAppRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesByAppResponse, error)
	GetAppEmailTemplatesByOtherApp(ctx context.Context, in *GetAppEmailTemplatesByOtherAppRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesByOtherAppResponse, error)
	GetAppEmailTemplateByAppLangUsedFor(ctx context.Context, in *GetAppEmailTemplateByAppLangUsedForRequest, opts ...grpc.CallOption) (*GetAppEmailTemplateByAppLangUsedForResponse, error)
	SendSMSCode(ctx context.Context, in *SendSMSCodeRequest, opts ...grpc.CallOption) (*SendSMSCodeResponse, error)
	VerifySMSCode(ctx context.Context, in *VerifySMSCodeRequest, opts ...grpc.CallOption) (*VerifySMSCodeResponse, error)
	SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SendEmailCodeResponse, error)
	VerifyEmailCode(ctx context.Context, in *VerifyEmailCodeRequest, opts ...grpc.CallOption) (*VerifyEmailCodeResponse, error)
}

type thirdGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdGatewayClient(cc grpc.ClientConnInterface) ThirdGatewayClient {
	return &thirdGatewayClient{cc}
}

func (c *thirdGatewayClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) CreateAppSMSTemplate(ctx context.Context, in *CreateAppSMSTemplateRequest, opts ...grpc.CallOption) (*CreateAppSMSTemplateResponse, error) {
	out := new(CreateAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/CreateAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppSMSTemplate(ctx context.Context, in *GetAppSMSTemplateRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateResponse, error) {
	out := new(GetAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) UpdateAppSMSTemplate(ctx context.Context, in *UpdateAppSMSTemplateRequest, opts ...grpc.CallOption) (*UpdateAppSMSTemplateResponse, error) {
	out := new(UpdateAppSMSTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/UpdateAppSMSTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppSMSTemplatesByApp(ctx context.Context, in *GetAppSMSTemplatesByAppRequest, opts ...grpc.CallOption) (*GetAppSMSTemplatesByAppResponse, error) {
	out := new(GetAppSMSTemplatesByAppResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppSMSTemplatesByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppSMSTemplatesByOtherApp(ctx context.Context, in *GetAppSMSTemplatesByOtherAppRequest, opts ...grpc.CallOption) (*GetAppSMSTemplatesByOtherAppResponse, error) {
	out := new(GetAppSMSTemplatesByOtherAppResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppSMSTemplatesByOtherApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppSMSTemplateByAppLangUsedFor(ctx context.Context, in *GetAppSMSTemplateByAppLangUsedForRequest, opts ...grpc.CallOption) (*GetAppSMSTemplateByAppLangUsedForResponse, error) {
	out := new(GetAppSMSTemplateByAppLangUsedForResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppSMSTemplateByAppLangUsedFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) CreateAppEmailTemplate(ctx context.Context, in *CreateAppEmailTemplateRequest, opts ...grpc.CallOption) (*CreateAppEmailTemplateResponse, error) {
	out := new(CreateAppEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/CreateAppEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppEmailTemplate(ctx context.Context, in *GetAppEmailTemplateRequest, opts ...grpc.CallOption) (*GetAppEmailTemplateResponse, error) {
	out := new(GetAppEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) UpdateAppEmailTemplate(ctx context.Context, in *UpdateAppEmailTemplateRequest, opts ...grpc.CallOption) (*UpdateAppEmailTemplateResponse, error) {
	out := new(UpdateAppEmailTemplateResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/UpdateAppEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppEmailTemplatesByApp(ctx context.Context, in *GetAppEmailTemplatesByAppRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesByAppResponse, error) {
	out := new(GetAppEmailTemplatesByAppResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppEmailTemplatesByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppEmailTemplatesByOtherApp(ctx context.Context, in *GetAppEmailTemplatesByOtherAppRequest, opts ...grpc.CallOption) (*GetAppEmailTemplatesByOtherAppResponse, error) {
	out := new(GetAppEmailTemplatesByOtherAppResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppEmailTemplatesByOtherApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) GetAppEmailTemplateByAppLangUsedFor(ctx context.Context, in *GetAppEmailTemplateByAppLangUsedForRequest, opts ...grpc.CallOption) (*GetAppEmailTemplateByAppLangUsedForResponse, error) {
	out := new(GetAppEmailTemplateByAppLangUsedForResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/GetAppEmailTemplateByAppLangUsedFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) SendSMSCode(ctx context.Context, in *SendSMSCodeRequest, opts ...grpc.CallOption) (*SendSMSCodeResponse, error) {
	out := new(SendSMSCodeResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/SendSMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) VerifySMSCode(ctx context.Context, in *VerifySMSCodeRequest, opts ...grpc.CallOption) (*VerifySMSCodeResponse, error) {
	out := new(VerifySMSCodeResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/VerifySMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SendEmailCodeResponse, error) {
	out := new(SendEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/SendEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdGatewayClient) VerifyEmailCode(ctx context.Context, in *VerifyEmailCodeRequest, opts ...grpc.CallOption) (*VerifyEmailCodeResponse, error) {
	out := new(VerifyEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.v1.ThirdGateway/VerifyEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdGatewayServer is the server API for ThirdGateway service.
// All implementations must embed UnimplementedThirdGatewayServer
// for forward compatibility
type ThirdGatewayServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	CreateAppSMSTemplate(context.Context, *CreateAppSMSTemplateRequest) (*CreateAppSMSTemplateResponse, error)
	GetAppSMSTemplate(context.Context, *GetAppSMSTemplateRequest) (*GetAppSMSTemplateResponse, error)
	UpdateAppSMSTemplate(context.Context, *UpdateAppSMSTemplateRequest) (*UpdateAppSMSTemplateResponse, error)
	GetAppSMSTemplatesByApp(context.Context, *GetAppSMSTemplatesByAppRequest) (*GetAppSMSTemplatesByAppResponse, error)
	GetAppSMSTemplatesByOtherApp(context.Context, *GetAppSMSTemplatesByOtherAppRequest) (*GetAppSMSTemplatesByOtherAppResponse, error)
	GetAppSMSTemplateByAppLangUsedFor(context.Context, *GetAppSMSTemplateByAppLangUsedForRequest) (*GetAppSMSTemplateByAppLangUsedForResponse, error)
	CreateAppEmailTemplate(context.Context, *CreateAppEmailTemplateRequest) (*CreateAppEmailTemplateResponse, error)
	GetAppEmailTemplate(context.Context, *GetAppEmailTemplateRequest) (*GetAppEmailTemplateResponse, error)
	UpdateAppEmailTemplate(context.Context, *UpdateAppEmailTemplateRequest) (*UpdateAppEmailTemplateResponse, error)
	GetAppEmailTemplatesByApp(context.Context, *GetAppEmailTemplatesByAppRequest) (*GetAppEmailTemplatesByAppResponse, error)
	GetAppEmailTemplatesByOtherApp(context.Context, *GetAppEmailTemplatesByOtherAppRequest) (*GetAppEmailTemplatesByOtherAppResponse, error)
	GetAppEmailTemplateByAppLangUsedFor(context.Context, *GetAppEmailTemplateByAppLangUsedForRequest) (*GetAppEmailTemplateByAppLangUsedForResponse, error)
	SendSMSCode(context.Context, *SendSMSCodeRequest) (*SendSMSCodeResponse, error)
	VerifySMSCode(context.Context, *VerifySMSCodeRequest) (*VerifySMSCodeResponse, error)
	SendEmailCode(context.Context, *SendEmailCodeRequest) (*SendEmailCodeResponse, error)
	VerifyEmailCode(context.Context, *VerifyEmailCodeRequest) (*VerifyEmailCodeResponse, error)
	mustEmbedUnimplementedThirdGatewayServer()
}

// UnimplementedThirdGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedThirdGatewayServer struct {
}

func (UnimplementedThirdGatewayServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedThirdGatewayServer) CreateAppSMSTemplate(context.Context, *CreateAppSMSTemplateRequest) (*CreateAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppSMSTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppSMSTemplate(context.Context, *GetAppSMSTemplateRequest) (*GetAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSMSTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) UpdateAppSMSTemplate(context.Context, *UpdateAppSMSTemplateRequest) (*UpdateAppSMSTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppSMSTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppSMSTemplatesByApp(context.Context, *GetAppSMSTemplatesByAppRequest) (*GetAppSMSTemplatesByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSMSTemplatesByApp not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppSMSTemplatesByOtherApp(context.Context, *GetAppSMSTemplatesByOtherAppRequest) (*GetAppSMSTemplatesByOtherAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSMSTemplatesByOtherApp not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppSMSTemplateByAppLangUsedFor(context.Context, *GetAppSMSTemplateByAppLangUsedForRequest) (*GetAppSMSTemplateByAppLangUsedForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppSMSTemplateByAppLangUsedFor not implemented")
}
func (UnimplementedThirdGatewayServer) CreateAppEmailTemplate(context.Context, *CreateAppEmailTemplateRequest) (*CreateAppEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppEmailTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppEmailTemplate(context.Context, *GetAppEmailTemplateRequest) (*GetAppEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppEmailTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) UpdateAppEmailTemplate(context.Context, *UpdateAppEmailTemplateRequest) (*UpdateAppEmailTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppEmailTemplate not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppEmailTemplatesByApp(context.Context, *GetAppEmailTemplatesByAppRequest) (*GetAppEmailTemplatesByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppEmailTemplatesByApp not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppEmailTemplatesByOtherApp(context.Context, *GetAppEmailTemplatesByOtherAppRequest) (*GetAppEmailTemplatesByOtherAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppEmailTemplatesByOtherApp not implemented")
}
func (UnimplementedThirdGatewayServer) GetAppEmailTemplateByAppLangUsedFor(context.Context, *GetAppEmailTemplateByAppLangUsedForRequest) (*GetAppEmailTemplateByAppLangUsedForResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppEmailTemplateByAppLangUsedFor not implemented")
}
func (UnimplementedThirdGatewayServer) SendSMSCode(context.Context, *SendSMSCodeRequest) (*SendSMSCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMSCode not implemented")
}
func (UnimplementedThirdGatewayServer) VerifySMSCode(context.Context, *VerifySMSCodeRequest) (*VerifySMSCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySMSCode not implemented")
}
func (UnimplementedThirdGatewayServer) SendEmailCode(context.Context, *SendEmailCodeRequest) (*SendEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedThirdGatewayServer) VerifyEmailCode(context.Context, *VerifyEmailCodeRequest) (*VerifyEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmailCode not implemented")
}
func (UnimplementedThirdGatewayServer) mustEmbedUnimplementedThirdGatewayServer() {}

// UnsafeThirdGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdGatewayServer will
// result in compilation errors.
type UnsafeThirdGatewayServer interface {
	mustEmbedUnimplementedThirdGatewayServer()
}

func RegisterThirdGatewayServer(s grpc.ServiceRegistrar, srv ThirdGatewayServer) {
	s.RegisterService(&ThirdGateway_ServiceDesc, srv)
}

func _ThirdGateway_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_CreateAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).CreateAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/CreateAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).CreateAppSMSTemplate(ctx, req.(*CreateAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppSMSTemplate(ctx, req.(*GetAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_UpdateAppSMSTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppSMSTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).UpdateAppSMSTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/UpdateAppSMSTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).UpdateAppSMSTemplate(ctx, req.(*UpdateAppSMSTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppSMSTemplatesByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSMSTemplatesByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppSMSTemplatesByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppSMSTemplatesByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppSMSTemplatesByApp(ctx, req.(*GetAppSMSTemplatesByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppSMSTemplatesByOtherApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSMSTemplatesByOtherAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppSMSTemplatesByOtherApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppSMSTemplatesByOtherApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppSMSTemplatesByOtherApp(ctx, req.(*GetAppSMSTemplatesByOtherAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppSMSTemplateByAppLangUsedFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppSMSTemplateByAppLangUsedForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppSMSTemplateByAppLangUsedFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppSMSTemplateByAppLangUsedFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppSMSTemplateByAppLangUsedFor(ctx, req.(*GetAppSMSTemplateByAppLangUsedForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_CreateAppEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).CreateAppEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/CreateAppEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).CreateAppEmailTemplate(ctx, req.(*CreateAppEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppEmailTemplate(ctx, req.(*GetAppEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_UpdateAppEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppEmailTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).UpdateAppEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/UpdateAppEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).UpdateAppEmailTemplate(ctx, req.(*UpdateAppEmailTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppEmailTemplatesByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppEmailTemplatesByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppEmailTemplatesByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppEmailTemplatesByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppEmailTemplatesByApp(ctx, req.(*GetAppEmailTemplatesByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppEmailTemplatesByOtherApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppEmailTemplatesByOtherAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppEmailTemplatesByOtherApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppEmailTemplatesByOtherApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppEmailTemplatesByOtherApp(ctx, req.(*GetAppEmailTemplatesByOtherAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_GetAppEmailTemplateByAppLangUsedFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppEmailTemplateByAppLangUsedForRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).GetAppEmailTemplateByAppLangUsedFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/GetAppEmailTemplateByAppLangUsedFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).GetAppEmailTemplateByAppLangUsedFor(ctx, req.(*GetAppEmailTemplateByAppLangUsedForRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_SendSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).SendSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/SendSMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).SendSMSCode(ctx, req.(*SendSMSCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_VerifySMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySMSCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).VerifySMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/VerifySMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).VerifySMSCode(ctx, req.(*VerifySMSCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/SendEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).SendEmailCode(ctx, req.(*SendEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdGateway_VerifyEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdGatewayServer).VerifyEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.v1.ThirdGateway/VerifyEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdGatewayServer).VerifyEmailCode(ctx, req.(*VerifyEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdGateway_ServiceDesc is the grpc.ServiceDesc for ThirdGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "third.gateway.v1.ThirdGateway",
	HandlerType: (*ThirdGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ThirdGateway_Version_Handler,
		},
		{
			MethodName: "CreateAppSMSTemplate",
			Handler:    _ThirdGateway_CreateAppSMSTemplate_Handler,
		},
		{
			MethodName: "GetAppSMSTemplate",
			Handler:    _ThirdGateway_GetAppSMSTemplate_Handler,
		},
		{
			MethodName: "UpdateAppSMSTemplate",
			Handler:    _ThirdGateway_UpdateAppSMSTemplate_Handler,
		},
		{
			MethodName: "GetAppSMSTemplatesByApp",
			Handler:    _ThirdGateway_GetAppSMSTemplatesByApp_Handler,
		},
		{
			MethodName: "GetAppSMSTemplatesByOtherApp",
			Handler:    _ThirdGateway_GetAppSMSTemplatesByOtherApp_Handler,
		},
		{
			MethodName: "GetAppSMSTemplateByAppLangUsedFor",
			Handler:    _ThirdGateway_GetAppSMSTemplateByAppLangUsedFor_Handler,
		},
		{
			MethodName: "CreateAppEmailTemplate",
			Handler:    _ThirdGateway_CreateAppEmailTemplate_Handler,
		},
		{
			MethodName: "GetAppEmailTemplate",
			Handler:    _ThirdGateway_GetAppEmailTemplate_Handler,
		},
		{
			MethodName: "UpdateAppEmailTemplate",
			Handler:    _ThirdGateway_UpdateAppEmailTemplate_Handler,
		},
		{
			MethodName: "GetAppEmailTemplatesByApp",
			Handler:    _ThirdGateway_GetAppEmailTemplatesByApp_Handler,
		},
		{
			MethodName: "GetAppEmailTemplatesByOtherApp",
			Handler:    _ThirdGateway_GetAppEmailTemplatesByOtherApp_Handler,
		},
		{
			MethodName: "GetAppEmailTemplateByAppLangUsedFor",
			Handler:    _ThirdGateway_GetAppEmailTemplateByAppLangUsedFor_Handler,
		},
		{
			MethodName: "SendSMSCode",
			Handler:    _ThirdGateway_SendSMSCode_Handler,
		},
		{
			MethodName: "VerifySMSCode",
			Handler:    _ThirdGateway_VerifySMSCode_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _ThirdGateway_SendEmailCode_Handler,
		},
		{
			MethodName: "VerifyEmailCode",
			Handler:    _ThirdGateway_VerifyEmailCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/thirdgateway/thirdgateway.proto",
}
