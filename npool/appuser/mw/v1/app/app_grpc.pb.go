// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/app/app.proto

package app

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

// AppUserMiddlewareAppClient is the client API for AppUserMiddlewareApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserMiddlewareAppClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	GetAppInfo(ctx context.Context, in *GetAppInfoRequest, opts ...grpc.CallOption) (*GetAppInfoResponse, error)
	GetAppInfos(ctx context.Context, in *GetAppInfosRequest, opts ...grpc.CallOption) (*GetAppInfosResponse, error)
	GetAppInfosByCreator(ctx context.Context, in *GetAppInfosByCreatorRequest, opts ...grpc.CallOption) (*GetAppInfosByCreatorResponse, error)
}

type appUserMiddlewareAppClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserMiddlewareAppClient(cc grpc.ClientConnInterface) AppUserMiddlewareAppClient {
	return &appUserMiddlewareAppClient{cc}
}

func (c *appUserMiddlewareAppClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/app.user.middleware.app.v2.AppUserMiddlewareApp/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareAppClient) GetAppInfo(ctx context.Context, in *GetAppInfoRequest, opts ...grpc.CallOption) (*GetAppInfoResponse, error) {
	out := new(GetAppInfoResponse)
	err := c.cc.Invoke(ctx, "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareAppClient) GetAppInfos(ctx context.Context, in *GetAppInfosRequest, opts ...grpc.CallOption) (*GetAppInfosResponse, error) {
	out := new(GetAppInfosResponse)
	err := c.cc.Invoke(ctx, "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appUserMiddlewareAppClient) GetAppInfosByCreator(ctx context.Context, in *GetAppInfosByCreatorRequest, opts ...grpc.CallOption) (*GetAppInfosByCreatorResponse, error) {
	out := new(GetAppInfosByCreatorResponse)
	err := c.cc.Invoke(ctx, "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfosByCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserMiddlewareAppServer is the server API for AppUserMiddlewareApp service.
// All implementations must embed UnimplementedAppUserMiddlewareAppServer
// for forward compatibility
type AppUserMiddlewareAppServer interface {
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	GetAppInfo(context.Context, *GetAppInfoRequest) (*GetAppInfoResponse, error)
	GetAppInfos(context.Context, *GetAppInfosRequest) (*GetAppInfosResponse, error)
	GetAppInfosByCreator(context.Context, *GetAppInfosByCreatorRequest) (*GetAppInfosByCreatorResponse, error)
	mustEmbedUnimplementedAppUserMiddlewareAppServer()
}

// UnimplementedAppUserMiddlewareAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserMiddlewareAppServer struct {
}

func (UnimplementedAppUserMiddlewareAppServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedAppUserMiddlewareAppServer) GetAppInfo(context.Context, *GetAppInfoRequest) (*GetAppInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppInfo not implemented")
}
func (UnimplementedAppUserMiddlewareAppServer) GetAppInfos(context.Context, *GetAppInfosRequest) (*GetAppInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppInfos not implemented")
}
func (UnimplementedAppUserMiddlewareAppServer) GetAppInfosByCreator(context.Context, *GetAppInfosByCreatorRequest) (*GetAppInfosByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppInfosByCreator not implemented")
}
func (UnimplementedAppUserMiddlewareAppServer) mustEmbedUnimplementedAppUserMiddlewareAppServer() {}

// UnsafeAppUserMiddlewareAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserMiddlewareAppServer will
// result in compilation errors.
type UnsafeAppUserMiddlewareAppServer interface {
	mustEmbedUnimplementedAppUserMiddlewareAppServer()
}

func RegisterAppUserMiddlewareAppServer(s grpc.ServiceRegistrar, srv AppUserMiddlewareAppServer) {
	s.RegisterService(&AppUserMiddlewareApp_ServiceDesc, srv)
}

func _AppUserMiddlewareApp_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareAppServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.middleware.app.v2.AppUserMiddlewareApp/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareAppServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareApp_GetAppInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareAppServer).GetAppInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareAppServer).GetAppInfo(ctx, req.(*GetAppInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareApp_GetAppInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareAppServer).GetAppInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareAppServer).GetAppInfos(ctx, req.(*GetAppInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppUserMiddlewareApp_GetAppInfosByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppInfosByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserMiddlewareAppServer).GetAppInfosByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.middleware.app.v2.AppUserMiddlewareApp/GetAppInfosByCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserMiddlewareAppServer).GetAppInfosByCreator(ctx, req.(*GetAppInfosByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserMiddlewareApp_ServiceDesc is the grpc.ServiceDesc for AppUserMiddlewareApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserMiddlewareApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.middleware.app.v2.AppUserMiddlewareApp",
	HandlerType: (*AppUserMiddlewareAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _AppUserMiddlewareApp_Version_Handler,
		},
		{
			MethodName: "GetAppInfo",
			Handler:    _AppUserMiddlewareApp_GetAppInfo_Handler,
		},
		{
			MethodName: "GetAppInfos",
			Handler:    _AppUserMiddlewareApp_GetAppInfos_Handler,
		},
		{
			MethodName: "GetAppInfosByCreator",
			Handler:    _AppUserMiddlewareApp_GetAppInfosByCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/app/app.proto",
}
