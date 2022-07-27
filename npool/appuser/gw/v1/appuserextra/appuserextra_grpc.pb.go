// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/appuserextra/appuserextra.proto

package appuserextra

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

// AppUserExtraClient is the client API for AppUserExtra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserExtraClient interface {
	CreateExtra(ctx context.Context, in *CreateExtraRequest, opts ...grpc.CallOption) (*CreateExtraResponse, error)
	GetExtra(ctx context.Context, in *GetExtraRequest, opts ...grpc.CallOption) (*GetExtraResponse, error)
	GetAppUserExtraAppUser(ctx context.Context, in *GetAppUserExtraRequest, opts ...grpc.CallOption) (*GetAppUserExtraResponse, error)
	UpdateExtra(ctx context.Context, in *UpdateExtraRequest, opts ...grpc.CallOption) (*UpdateExtraResponse, error)
}

type appUserExtraClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserExtraClient(cc grpc.ClientConnInterface) AppUserExtraClient {
	return &appUserExtraClient{cc}
}

func (c *appUserExtraClient) CreateExtra(ctx context.Context, in *CreateExtraRequest, opts ...grpc.CallOption) (*CreateExtraResponse, error) {
	out := new(CreateExtraResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.appuserextra.v1.AppUserExtra/CreateExtra", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserExtraClient) GetExtra(ctx context.Context, in *GetExtraRequest, opts ...grpc.CallOption) (*GetExtraResponse, error) {
	out := new(GetExtraResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.appuserextra.v1.AppUserExtra/GetExtra", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserExtraClient) GetAppUserExtraAppUser(ctx context.Context, in *GetAppUserExtraRequest, opts ...grpc.CallOption) (*GetAppUserExtraResponse, error) {
	out := new(GetAppUserExtraResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.appuserextra.v1.AppUserExtra/GetAppUserExtraAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserExtraClient) UpdateExtra(ctx context.Context, in *UpdateExtraRequest, opts ...grpc.CallOption) (*UpdateExtraResponse, error) {
	out := new(UpdateExtraResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.appuserextra.v1.AppUserExtra/UpdateExtra", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserExtraServer is the server API for AppUserExtra service.
// All implementations must embed UnimplementedAppUserExtraServer
// for forward compatibility
type AppUserExtraServer interface {
	CreateExtra(context.Context, *CreateExtraRequest) (*CreateExtraResponse, error)
	GetExtra(context.Context, *GetExtraRequest) (*GetExtraResponse, error)
	GetAppUserExtraAppUser(context.Context, *GetAppUserExtraRequest) (*GetAppUserExtraResponse, error)
	UpdateExtra(context.Context, *UpdateExtraRequest) (*UpdateExtraResponse, error)
	mustEmbedUnimplementedAppUserExtraServer()
}

// UnimplementedAppUserExtraServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserExtraServer struct {
}

func (UnimplementedAppUserExtraServer) CreateExtra(context.Context, *CreateExtraRequest) (*CreateExtraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExtra not implemented")
}
func (UnimplementedAppUserExtraServer) GetExtra(context.Context, *GetExtraRequest) (*GetExtraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtra not implemented")
}
func (UnimplementedAppUserExtraServer) GetAppUserExtraAppUser(context.Context, *GetAppUserExtraRequest) (*GetAppUserExtraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUserExtraAppUser not implemented")
}
func (UnimplementedAppUserExtraServer) UpdateExtra(context.Context, *UpdateExtraRequest) (*UpdateExtraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExtra not implemented")
}
func (UnimplementedAppUserExtraServer) mustEmbedUnimplementedAppUserExtraServer() {}

// UnsafeAppUserExtraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserExtraServer will
// result in compilation errors.
type UnsafeAppUserExtraServer interface {
	mustEmbedUnimplementedAppUserExtraServer()
}

func RegisterAppUserExtraServer(s grpc.ServiceRegistrar, srv AppUserExtraServer) {
	s.RegisterService(&AppUserExtra_ServiceDesc, srv)
}

func _AppUserExtra_CreateExtra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExtraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserExtraServer).CreateExtra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.appuserextra.v1.AppUserExtra/CreateExtra",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserExtraServer).CreateExtra(ctx, req.(*CreateExtraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserExtra_GetExtra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserExtraServer).GetExtra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.appuserextra.v1.AppUserExtra/GetExtra",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserExtraServer).GetExtra(ctx, req.(*GetExtraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserExtra_GetAppUserExtraAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUserExtraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserExtraServer).GetAppUserExtraAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.appuserextra.v1.AppUserExtra/GetAppUserExtraAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserExtraServer).GetAppUserExtraAppUser(ctx, req.(*GetAppUserExtraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserExtra_UpdateExtra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExtraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserExtraServer).UpdateExtra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.appuserextra.v1.AppUserExtra/UpdateExtra",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserExtraServer).UpdateExtra(ctx, req.(*UpdateExtraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserExtra_ServiceDesc is the grpc.ServiceDesc for AppUserExtra service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserExtra_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.gateway.appuserextra.v1.AppUserExtra",
	HandlerType: (*AppUserExtraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExtra",
			Handler:    _AppUserExtra_CreateExtra_Handler,
		},
		{
			MethodName: "GetExtra",
			Handler:    _AppUserExtra_GetExtra_Handler,
		},
		{
			MethodName: "GetAppUserExtraAppUser",
			Handler:    _AppUserExtra_GetAppUserExtraAppUser_Handler,
		},
		{
			MethodName: "UpdateExtra",
			Handler:    _AppUserExtra_UpdateExtra_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/appuserextra/appuserextra.proto",
}
