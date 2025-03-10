// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/good/like/like.proto

package like

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
	CreateLike(ctx context.Context, in *CreateLikeRequest, opts ...grpc.CallOption) (*CreateLikeResponse, error)
	UpdateLike(ctx context.Context, in *UpdateLikeRequest, opts ...grpc.CallOption) (*UpdateLikeResponse, error)
	GetLike(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*GetLikeResponse, error)
	GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error)
	ExistLikeConds(ctx context.Context, in *ExistLikeCondsRequest, opts ...grpc.CallOption) (*ExistLikeCondsResponse, error)
	DeleteLike(ctx context.Context, in *DeleteLikeRequest, opts ...grpc.CallOption) (*DeleteLikeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateLike(ctx context.Context, in *CreateLikeRequest, opts ...grpc.CallOption) (*CreateLikeResponse, error) {
	out := new(CreateLikeResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/CreateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateLike(ctx context.Context, in *UpdateLikeRequest, opts ...grpc.CallOption) (*UpdateLikeResponse, error) {
	out := new(UpdateLikeResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/UpdateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLike(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*GetLikeResponse, error) {
	out := new(GetLikeResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/GetLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error) {
	out := new(GetLikesResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/GetLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistLikeConds(ctx context.Context, in *ExistLikeCondsRequest, opts ...grpc.CallOption) (*ExistLikeCondsResponse, error) {
	out := new(ExistLikeCondsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/ExistLikeConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteLike(ctx context.Context, in *DeleteLikeRequest, opts ...grpc.CallOption) (*DeleteLikeResponse, error) {
	out := new(DeleteLikeResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.app.good1.like.v1.Middleware/DeleteLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateLike(context.Context, *CreateLikeRequest) (*CreateLikeResponse, error)
	UpdateLike(context.Context, *UpdateLikeRequest) (*UpdateLikeResponse, error)
	GetLike(context.Context, *GetLikeRequest) (*GetLikeResponse, error)
	GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error)
	ExistLikeConds(context.Context, *ExistLikeCondsRequest) (*ExistLikeCondsResponse, error)
	DeleteLike(context.Context, *DeleteLikeRequest) (*DeleteLikeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateLike(context.Context, *CreateLikeRequest) (*CreateLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedMiddlewareServer) UpdateLike(context.Context, *UpdateLikeRequest) (*UpdateLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLike not implemented")
}
func (UnimplementedMiddlewareServer) GetLike(context.Context, *GetLikeRequest) (*GetLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLike not implemented")
}
func (UnimplementedMiddlewareServer) GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikes not implemented")
}
func (UnimplementedMiddlewareServer) ExistLikeConds(context.Context, *ExistLikeCondsRequest) (*ExistLikeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistLikeConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteLike(context.Context, *DeleteLikeRequest) (*DeleteLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLike not implemented")
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

func _Middleware_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/CreateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateLike(ctx, req.(*CreateLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/UpdateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateLike(ctx, req.(*UpdateLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/GetLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLike(ctx, req.(*GetLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/GetLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLikes(ctx, req.(*GetLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistLikeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistLikeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistLikeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/ExistLikeConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistLikeConds(ctx, req.(*ExistLikeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.app.good1.like.v1.Middleware/DeleteLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteLike(ctx, req.(*DeleteLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.good1.like.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLike",
			Handler:    _Middleware_CreateLike_Handler,
		},
		{
			MethodName: "UpdateLike",
			Handler:    _Middleware_UpdateLike_Handler,
		},
		{
			MethodName: "GetLike",
			Handler:    _Middleware_GetLike_Handler,
		},
		{
			MethodName: "GetLikes",
			Handler:    _Middleware_GetLikes_Handler,
		},
		{
			MethodName: "ExistLikeConds",
			Handler:    _Middleware_ExistLikeConds_Handler,
		},
		{
			MethodName: "DeleteLike",
			Handler:    _Middleware_DeleteLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/good/like/like.proto",
}
