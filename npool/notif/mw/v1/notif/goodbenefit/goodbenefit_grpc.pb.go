// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/notif/goodbenefit/goodbenefit.proto

package goodbenefit

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateGoodBenefit(ctx context.Context, in *CreateGoodBenefitRequest, opts ...grpc.CallOption) (*CreateGoodBenefitResponse, error)
	UpdateGoodBenefit(ctx context.Context, in *UpdateGoodBenefitRequest, opts ...grpc.CallOption) (*UpdateGoodBenefitResponse, error)
	GetGoodBenefits(ctx context.Context, in *GetGoodBenefitsRequest, opts ...grpc.CallOption) (*GetGoodBenefitsResponse, error)
	GetGoodBenefit(ctx context.Context, in *GetGoodBenefitRequest, opts ...grpc.CallOption) (*GetGoodBenefitResponse, error)
	ExistGoodBenefitConds(ctx context.Context, in *ExistGoodBenefitCondsRequest, opts ...grpc.CallOption) (*ExistGoodBenefitCondsResponse, error)
	DeleteGoodBenefit(ctx context.Context, in *DeleteGoodBenefitRequest, opts ...grpc.CallOption) (*DeleteGoodBenefitResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateGoodBenefit(ctx context.Context, in *CreateGoodBenefitRequest, opts ...grpc.CallOption) (*CreateGoodBenefitResponse, error) {
	out := new(CreateGoodBenefitResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/CreateGoodBenefit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateGoodBenefit(ctx context.Context, in *UpdateGoodBenefitRequest, opts ...grpc.CallOption) (*UpdateGoodBenefitResponse, error) {
	out := new(UpdateGoodBenefitResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/UpdateGoodBenefit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetGoodBenefits(ctx context.Context, in *GetGoodBenefitsRequest, opts ...grpc.CallOption) (*GetGoodBenefitsResponse, error) {
	out := new(GetGoodBenefitsResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/GetGoodBenefits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetGoodBenefit(ctx context.Context, in *GetGoodBenefitRequest, opts ...grpc.CallOption) (*GetGoodBenefitResponse, error) {
	out := new(GetGoodBenefitResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/GetGoodBenefit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistGoodBenefitConds(ctx context.Context, in *ExistGoodBenefitCondsRequest, opts ...grpc.CallOption) (*ExistGoodBenefitCondsResponse, error) {
	out := new(ExistGoodBenefitCondsResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/ExistGoodBenefitConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteGoodBenefit(ctx context.Context, in *DeleteGoodBenefitRequest, opts ...grpc.CallOption) (*DeleteGoodBenefitResponse, error) {
	out := new(DeleteGoodBenefitResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.notif.goodbenefit.v1.Middleware/DeleteGoodBenefit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateGoodBenefit(context.Context, *CreateGoodBenefitRequest) (*CreateGoodBenefitResponse, error)
	UpdateGoodBenefit(context.Context, *UpdateGoodBenefitRequest) (*UpdateGoodBenefitResponse, error)
	GetGoodBenefits(context.Context, *GetGoodBenefitsRequest) (*GetGoodBenefitsResponse, error)
	GetGoodBenefit(context.Context, *GetGoodBenefitRequest) (*GetGoodBenefitResponse, error)
	ExistGoodBenefitConds(context.Context, *ExistGoodBenefitCondsRequest) (*ExistGoodBenefitCondsResponse, error)
	DeleteGoodBenefit(context.Context, *DeleteGoodBenefitRequest) (*DeleteGoodBenefitResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateGoodBenefit(context.Context, *CreateGoodBenefitRequest) (*CreateGoodBenefitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodBenefit not implemented")
}
func (UnimplementedMiddlewareServer) UpdateGoodBenefit(context.Context, *UpdateGoodBenefitRequest) (*UpdateGoodBenefitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodBenefit not implemented")
}
func (UnimplementedMiddlewareServer) GetGoodBenefits(context.Context, *GetGoodBenefitsRequest) (*GetGoodBenefitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodBenefits not implemented")
}
func (UnimplementedMiddlewareServer) GetGoodBenefit(context.Context, *GetGoodBenefitRequest) (*GetGoodBenefitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodBenefit not implemented")
}
func (UnimplementedMiddlewareServer) ExistGoodBenefitConds(context.Context, *ExistGoodBenefitCondsRequest) (*ExistGoodBenefitCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistGoodBenefitConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteGoodBenefit(context.Context, *DeleteGoodBenefitRequest) (*DeleteGoodBenefitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoodBenefit not implemented")
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

func _Middleware_CreateGoodBenefit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodBenefitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateGoodBenefit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/CreateGoodBenefit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateGoodBenefit(ctx, req.(*CreateGoodBenefitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateGoodBenefit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodBenefitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateGoodBenefit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/UpdateGoodBenefit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateGoodBenefit(ctx, req.(*UpdateGoodBenefitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetGoodBenefits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodBenefitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGoodBenefits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/GetGoodBenefits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGoodBenefits(ctx, req.(*GetGoodBenefitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetGoodBenefit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodBenefitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGoodBenefit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/GetGoodBenefit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGoodBenefit(ctx, req.(*GetGoodBenefitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistGoodBenefitConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistGoodBenefitCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistGoodBenefitConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/ExistGoodBenefitConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistGoodBenefitConds(ctx, req.(*ExistGoodBenefitCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteGoodBenefit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodBenefitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteGoodBenefit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.notif.goodbenefit.v1.Middleware/DeleteGoodBenefit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteGoodBenefit(ctx, req.(*DeleteGoodBenefitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.notif.goodbenefit.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoodBenefit",
			Handler:    _Middleware_CreateGoodBenefit_Handler,
		},
		{
			MethodName: "UpdateGoodBenefit",
			Handler:    _Middleware_UpdateGoodBenefit_Handler,
		},
		{
			MethodName: "GetGoodBenefits",
			Handler:    _Middleware_GetGoodBenefits_Handler,
		},
		{
			MethodName: "GetGoodBenefit",
			Handler:    _Middleware_GetGoodBenefit_Handler,
		},
		{
			MethodName: "ExistGoodBenefitConds",
			Handler:    _Middleware_ExistGoodBenefitConds_Handler,
		},
		{
			MethodName: "DeleteGoodBenefit",
			Handler:    _Middleware_DeleteGoodBenefit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/notif/goodbenefit/goodbenefit.proto",
}
