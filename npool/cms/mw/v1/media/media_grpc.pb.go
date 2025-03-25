// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/cms/mw/v1/media/media.proto

package media

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
	CreateMedia(ctx context.Context, in *CreateMediaRequest, opts ...grpc.CallOption) (*CreateMediaResponse, error)
	GetMedia(ctx context.Context, in *GetMediaRequest, opts ...grpc.CallOption) (*GetMediaResponse, error)
	GetMedias(ctx context.Context, in *GetMediasRequest, opts ...grpc.CallOption) (*GetMediasResponse, error)
	ExistMedia(ctx context.Context, in *ExistMediaRequest, opts ...grpc.CallOption) (*ExistMediaResponse, error)
	ExistMediaConds(ctx context.Context, in *ExistMediaCondsRequest, opts ...grpc.CallOption) (*ExistMediaCondsResponse, error)
	DeleteMedia(ctx context.Context, in *DeleteMediaRequest, opts ...grpc.CallOption) (*DeleteMediaResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateMedia(ctx context.Context, in *CreateMediaRequest, opts ...grpc.CallOption) (*CreateMediaResponse, error) {
	out := new(CreateMediaResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/CreateMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetMedia(ctx context.Context, in *GetMediaRequest, opts ...grpc.CallOption) (*GetMediaResponse, error) {
	out := new(GetMediaResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/GetMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetMedias(ctx context.Context, in *GetMediasRequest, opts ...grpc.CallOption) (*GetMediasResponse, error) {
	out := new(GetMediasResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/GetMedias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistMedia(ctx context.Context, in *ExistMediaRequest, opts ...grpc.CallOption) (*ExistMediaResponse, error) {
	out := new(ExistMediaResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/ExistMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistMediaConds(ctx context.Context, in *ExistMediaCondsRequest, opts ...grpc.CallOption) (*ExistMediaCondsResponse, error) {
	out := new(ExistMediaCondsResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/ExistMediaConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteMedia(ctx context.Context, in *DeleteMediaRequest, opts ...grpc.CallOption) (*DeleteMediaResponse, error) {
	out := new(DeleteMediaResponse)
	err := c.cc.Invoke(ctx, "/cms.middleware.media.v1.Middleware/DeleteMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateMedia(context.Context, *CreateMediaRequest) (*CreateMediaResponse, error)
	GetMedia(context.Context, *GetMediaRequest) (*GetMediaResponse, error)
	GetMedias(context.Context, *GetMediasRequest) (*GetMediasResponse, error)
	ExistMedia(context.Context, *ExistMediaRequest) (*ExistMediaResponse, error)
	ExistMediaConds(context.Context, *ExistMediaCondsRequest) (*ExistMediaCondsResponse, error)
	DeleteMedia(context.Context, *DeleteMediaRequest) (*DeleteMediaResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateMedia(context.Context, *CreateMediaRequest) (*CreateMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMedia not implemented")
}
func (UnimplementedMiddlewareServer) GetMedia(context.Context, *GetMediaRequest) (*GetMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedia not implemented")
}
func (UnimplementedMiddlewareServer) GetMedias(context.Context, *GetMediasRequest) (*GetMediasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedias not implemented")
}
func (UnimplementedMiddlewareServer) ExistMedia(context.Context, *ExistMediaRequest) (*ExistMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMedia not implemented")
}
func (UnimplementedMiddlewareServer) ExistMediaConds(context.Context, *ExistMediaCondsRequest) (*ExistMediaCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMediaConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteMedia(context.Context, *DeleteMediaRequest) (*DeleteMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedia not implemented")
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

func _Middleware_CreateMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/CreateMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateMedia(ctx, req.(*CreateMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/GetMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetMedia(ctx, req.(*GetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetMedias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetMedias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/GetMedias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetMedias(ctx, req.(*GetMediasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/ExistMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistMedia(ctx, req.(*ExistMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistMediaConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMediaCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistMediaConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/ExistMediaConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistMediaConds(ctx, req.(*ExistMediaCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.middleware.media.v1.Middleware/DeleteMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteMedia(ctx, req.(*DeleteMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.middleware.media.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMedia",
			Handler:    _Middleware_CreateMedia_Handler,
		},
		{
			MethodName: "GetMedia",
			Handler:    _Middleware_GetMedia_Handler,
		},
		{
			MethodName: "GetMedias",
			Handler:    _Middleware_GetMedias_Handler,
		},
		{
			MethodName: "ExistMedia",
			Handler:    _Middleware_ExistMedia_Handler,
		},
		{
			MethodName: "ExistMediaConds",
			Handler:    _Middleware_ExistMediaConds_Handler,
		},
		{
			MethodName: "DeleteMedia",
			Handler:    _Middleware_DeleteMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/cms/mw/v1/media/media.proto",
}
