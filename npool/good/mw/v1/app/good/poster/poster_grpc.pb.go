// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/good/poster/poster.proto

package poster

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
	Middleware_CreatePoster_FullMethodName = "/good.middleware.app.good1.poster.v1.Middleware/CreatePoster"
	Middleware_UpdatePoster_FullMethodName = "/good.middleware.app.good1.poster.v1.Middleware/UpdatePoster"
	Middleware_GetPoster_FullMethodName    = "/good.middleware.app.good1.poster.v1.Middleware/GetPoster"
	Middleware_GetPosters_FullMethodName   = "/good.middleware.app.good1.poster.v1.Middleware/GetPosters"
	Middleware_DeletePoster_FullMethodName = "/good.middleware.app.good1.poster.v1.Middleware/DeletePoster"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreatePoster(ctx context.Context, in *CreatePosterRequest, opts ...grpc.CallOption) (*CreatePosterResponse, error)
	UpdatePoster(ctx context.Context, in *UpdatePosterRequest, opts ...grpc.CallOption) (*UpdatePosterResponse, error)
	GetPoster(ctx context.Context, in *GetPosterRequest, opts ...grpc.CallOption) (*GetPosterResponse, error)
	GetPosters(ctx context.Context, in *GetPostersRequest, opts ...grpc.CallOption) (*GetPostersResponse, error)
	DeletePoster(ctx context.Context, in *DeletePosterRequest, opts ...grpc.CallOption) (*DeletePosterResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreatePoster(ctx context.Context, in *CreatePosterRequest, opts ...grpc.CallOption) (*CreatePosterResponse, error) {
	out := new(CreatePosterResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePoster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePoster(ctx context.Context, in *UpdatePosterRequest, opts ...grpc.CallOption) (*UpdatePosterResponse, error) {
	out := new(UpdatePosterResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePoster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPoster(ctx context.Context, in *GetPosterRequest, opts ...grpc.CallOption) (*GetPosterResponse, error) {
	out := new(GetPosterResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPoster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPosters(ctx context.Context, in *GetPostersRequest, opts ...grpc.CallOption) (*GetPostersResponse, error) {
	out := new(GetPostersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPosters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeletePoster(ctx context.Context, in *DeletePosterRequest, opts ...grpc.CallOption) (*DeletePosterResponse, error) {
	out := new(DeletePosterResponse)
	err := c.cc.Invoke(ctx, Middleware_DeletePoster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreatePoster(context.Context, *CreatePosterRequest) (*CreatePosterResponse, error)
	UpdatePoster(context.Context, *UpdatePosterRequest) (*UpdatePosterResponse, error)
	GetPoster(context.Context, *GetPosterRequest) (*GetPosterResponse, error)
	GetPosters(context.Context, *GetPostersRequest) (*GetPostersResponse, error)
	DeletePoster(context.Context, *DeletePosterRequest) (*DeletePosterResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreatePoster(context.Context, *CreatePosterRequest) (*CreatePosterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoster not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePoster(context.Context, *UpdatePosterRequest) (*UpdatePosterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePoster not implemented")
}
func (UnimplementedMiddlewareServer) GetPoster(context.Context, *GetPosterRequest) (*GetPosterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoster not implemented")
}
func (UnimplementedMiddlewareServer) GetPosters(context.Context, *GetPostersRequest) (*GetPostersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosters not implemented")
}
func (UnimplementedMiddlewareServer) DeletePoster(context.Context, *DeletePosterRequest) (*DeletePosterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePoster not implemented")
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

func _Middleware_CreatePoster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePoster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePoster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePoster(ctx, req.(*CreatePosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePoster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePoster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePoster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePoster(ctx, req.(*UpdatePosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPoster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPoster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPoster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPoster(ctx, req.(*GetPosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPosters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPosters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPosters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPosters(ctx, req.(*GetPostersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeletePoster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeletePoster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeletePoster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeletePoster(ctx, req.(*DeletePosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.good1.poster.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoster",
			Handler:    _Middleware_CreatePoster_Handler,
		},
		{
			MethodName: "UpdatePoster",
			Handler:    _Middleware_UpdatePoster_Handler,
		},
		{
			MethodName: "GetPoster",
			Handler:    _Middleware_GetPoster_Handler,
		},
		{
			MethodName: "GetPosters",
			Handler:    _Middleware_GetPosters_Handler,
		},
		{
			MethodName: "DeletePoster",
			Handler:    _Middleware_DeletePoster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/good/poster/poster.proto",
}
