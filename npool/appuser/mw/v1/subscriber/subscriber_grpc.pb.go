// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/appuser/mw/v1/subscriber/subscriber.proto

package subscriber

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
	Middleware_CreateSubscriber_FullMethodName = "/appuser.middleware.subscriber.v1.Middleware/CreateSubscriber"
	Middleware_UpdateSubscriber_FullMethodName = "/appuser.middleware.subscriber.v1.Middleware/UpdateSubscriber"
	Middleware_GetSubscriber_FullMethodName    = "/appuser.middleware.subscriber.v1.Middleware/GetSubscriber"
	Middleware_GetSubscriberes_FullMethodName  = "/appuser.middleware.subscriber.v1.Middleware/GetSubscriberes"
	Middleware_DeleteSubscriber_FullMethodName = "/appuser.middleware.subscriber.v1.Middleware/DeleteSubscriber"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateSubscriber(ctx context.Context, in *CreateSubscriberRequest, opts ...grpc.CallOption) (*CreateSubscriberResponse, error)
	UpdateSubscriber(ctx context.Context, in *UpdateSubscriberRequest, opts ...grpc.CallOption) (*UpdateSubscriberResponse, error)
	GetSubscriber(ctx context.Context, in *GetSubscriberRequest, opts ...grpc.CallOption) (*GetSubscriberResponse, error)
	GetSubscriberes(ctx context.Context, in *GetSubscriberesRequest, opts ...grpc.CallOption) (*GetSubscriberesResponse, error)
	DeleteSubscriber(ctx context.Context, in *DeleteSubscriberRequest, opts ...grpc.CallOption) (*DeleteSubscriberResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateSubscriber(ctx context.Context, in *CreateSubscriberRequest, opts ...grpc.CallOption) (*CreateSubscriberResponse, error) {
	out := new(CreateSubscriberResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateSubscriber(ctx context.Context, in *UpdateSubscriberRequest, opts ...grpc.CallOption) (*UpdateSubscriberResponse, error) {
	out := new(UpdateSubscriberResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSubscriber(ctx context.Context, in *GetSubscriberRequest, opts ...grpc.CallOption) (*GetSubscriberResponse, error) {
	out := new(GetSubscriberResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSubscriberes(ctx context.Context, in *GetSubscriberesRequest, opts ...grpc.CallOption) (*GetSubscriberesResponse, error) {
	out := new(GetSubscriberesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubscriberes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteSubscriber(ctx context.Context, in *DeleteSubscriberRequest, opts ...grpc.CallOption) (*DeleteSubscriberResponse, error) {
	out := new(DeleteSubscriberResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateSubscriber(context.Context, *CreateSubscriberRequest) (*CreateSubscriberResponse, error)
	UpdateSubscriber(context.Context, *UpdateSubscriberRequest) (*UpdateSubscriberResponse, error)
	GetSubscriber(context.Context, *GetSubscriberRequest) (*GetSubscriberResponse, error)
	GetSubscriberes(context.Context, *GetSubscriberesRequest) (*GetSubscriberesResponse, error)
	DeleteSubscriber(context.Context, *DeleteSubscriberRequest) (*DeleteSubscriberResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateSubscriber(context.Context, *CreateSubscriberRequest) (*CreateSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscriber not implemented")
}
func (UnimplementedMiddlewareServer) UpdateSubscriber(context.Context, *UpdateSubscriberRequest) (*UpdateSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscriber not implemented")
}
func (UnimplementedMiddlewareServer) GetSubscriber(context.Context, *GetSubscriberRequest) (*GetSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriber not implemented")
}
func (UnimplementedMiddlewareServer) GetSubscriberes(context.Context, *GetSubscriberesRequest) (*GetSubscriberesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberes not implemented")
}
func (UnimplementedMiddlewareServer) DeleteSubscriber(context.Context, *DeleteSubscriberRequest) (*DeleteSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscriber not implemented")
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

func _Middleware_CreateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateSubscriber(ctx, req.(*CreateSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateSubscriber(ctx, req.(*UpdateSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubscriber(ctx, req.(*GetSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSubscriberes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriberesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubscriberes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubscriberes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubscriberes(ctx, req.(*GetSubscriberesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteSubscriber(ctx, req.(*DeleteSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.middleware.subscriber.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscriber",
			Handler:    _Middleware_CreateSubscriber_Handler,
		},
		{
			MethodName: "UpdateSubscriber",
			Handler:    _Middleware_UpdateSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriber",
			Handler:    _Middleware_GetSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriberes",
			Handler:    _Middleware_GetSubscriberes_Handler,
		},
		{
			MethodName: "DeleteSubscriber",
			Handler:    _Middleware_DeleteSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mw/v1/subscriber/subscriber.proto",
}
