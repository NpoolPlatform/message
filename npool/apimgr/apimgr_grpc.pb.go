// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/apimgr/apimgr.proto

package apimgr

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

// ApiManagerClient is the client API for ApiManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiManagerClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetApis(ctx context.Context, in *GetApisRequest, opts ...grpc.CallOption) (*GetApisResponse, error)
	GetServiceMethodApi(ctx context.Context, in *GetServiceMethodApiRequest, opts ...grpc.CallOption) (*GetServiceMethodApiResponse, error)
}

type apiManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiManagerClient(cc grpc.ClientConnInterface) ApiManagerClient {
	return &apiManagerClient{cc}
}

func (c *apiManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/api.manager.v1.ApiManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiManagerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/api.manager.v1.ApiManager/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiManagerClient) GetApis(ctx context.Context, in *GetApisRequest, opts ...grpc.CallOption) (*GetApisResponse, error) {
	out := new(GetApisResponse)
	err := c.cc.Invoke(ctx, "/api.manager.v1.ApiManager/GetApis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiManagerClient) GetServiceMethodApi(ctx context.Context, in *GetServiceMethodApiRequest, opts ...grpc.CallOption) (*GetServiceMethodApiResponse, error) {
	out := new(GetServiceMethodApiResponse)
	err := c.cc.Invoke(ctx, "/api.manager.v1.ApiManager/GetServiceMethodApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiManagerServer is the server API for ApiManager service.
// All implementations must embed UnimplementedApiManagerServer
// for forward compatibility
type ApiManagerServer interface {
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetApis(context.Context, *GetApisRequest) (*GetApisResponse, error)
	GetServiceMethodApi(context.Context, *GetServiceMethodApiRequest) (*GetServiceMethodApiResponse, error)
	mustEmbedUnimplementedApiManagerServer()
}

// UnimplementedApiManagerServer must be embedded to have forward compatible implementations.
type UnimplementedApiManagerServer struct {
}

func (UnimplementedApiManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedApiManagerServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedApiManagerServer) GetApis(context.Context, *GetApisRequest) (*GetApisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApis not implemented")
}
func (UnimplementedApiManagerServer) GetServiceMethodApi(context.Context, *GetServiceMethodApiRequest) (*GetServiceMethodApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceMethodApi not implemented")
}
func (UnimplementedApiManagerServer) mustEmbedUnimplementedApiManagerServer() {}

// UnsafeApiManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiManagerServer will
// result in compilation errors.
type UnsafeApiManagerServer interface {
	mustEmbedUnimplementedApiManagerServer()
}

func RegisterApiManagerServer(s grpc.ServiceRegistrar, srv ApiManagerServer) {
	s.RegisterService(&ApiManager_ServiceDesc, srv)
}

func _ApiManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manager.v1.ApiManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiManager_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiManagerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manager.v1.ApiManager/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiManagerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiManager_GetApis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiManagerServer).GetApis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manager.v1.ApiManager/GetApis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiManagerServer).GetApis(ctx, req.(*GetApisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiManager_GetServiceMethodApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceMethodApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiManagerServer).GetServiceMethodApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.manager.v1.ApiManager/GetServiceMethodApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiManagerServer).GetServiceMethodApi(ctx, req.(*GetServiceMethodApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiManager_ServiceDesc is the grpc.ServiceDesc for ApiManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.manager.v1.ApiManager",
	HandlerType: (*ApiManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ApiManager_Version_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ApiManager_Register_Handler,
		},
		{
			MethodName: "GetApis",
			Handler:    _ApiManager_GetApis_Handler,
		},
		{
			MethodName: "GetServiceMethodApi",
			Handler:    _ApiManager_GetServiceMethodApi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/apimgr/apimgr.proto",
}
