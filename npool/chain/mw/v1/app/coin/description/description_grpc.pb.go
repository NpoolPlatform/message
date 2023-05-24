// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/mw/v1/app/coin/description/description.proto

package description

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
<<<<<<< HEAD:npool/chain/mw/v1/app/coin/description/description_grpc.pb.go
	Middleware_CreateCoinDescription_FullMethodName = "/chain.middleware.app.coin.description.v1.Middleware/CreateCoinDescription"
	Middleware_GetCoinDescription_FullMethodName    = "/chain.middleware.app.coin.description.v1.Middleware/GetCoinDescription"
	Middleware_GetCoinDescriptions_FullMethodName   = "/chain.middleware.app.coin.description.v1.Middleware/GetCoinDescriptions"
	Middleware_UpdateCoinDescription_FullMethodName = "/chain.middleware.app.coin.description.v1.Middleware/UpdateCoinDescription"
=======
	Middleware_CreateCoinDescription_FullMethodName = "/chain.middleware.appcoin.description.v1.Middleware/CreateCoinDescription"
	Middleware_GetCoinDescription_FullMethodName    = "/chain.middleware.appcoin.description.v1.Middleware/GetCoinDescription"
	Middleware_GetCoinDescriptions_FullMethodName   = "/chain.middleware.appcoin.description.v1.Middleware/GetCoinDescriptions"
	Middleware_UpdateCoinDescription_FullMethodName = "/chain.middleware.appcoin.description.v1.Middleware/UpdateCoinDescription"
>>>>>>> Regenerate:npool/chain/mw/v1/appcoin/description/description_grpc.pb.go
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error)
	GetCoinDescription(ctx context.Context, in *GetCoinDescriptionRequest, opts ...grpc.CallOption) (*GetCoinDescriptionResponse, error)
	GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error)
	UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error) {
	out := new(CreateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinDescription(ctx context.Context, in *GetCoinDescriptionRequest, opts ...grpc.CallOption) (*GetCoinDescriptionResponse, error) {
	out := new(GetCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error) {
	out := new(GetCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCoinDescriptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error) {
	out := new(UpdateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateCoinDescription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error)
	GetCoinDescription(context.Context, *GetCoinDescriptionRequest) (*GetCoinDescriptionResponse, error)
	GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error)
	UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinDescription not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinDescription(context.Context, *GetCoinDescriptionRequest) (*GetCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescription not implemented")
}
func (UnimplementedMiddlewareServer) GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescriptions not implemented")
}
func (UnimplementedMiddlewareServer) UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinDescription not implemented")
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

func _Middleware_CreateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCoinDescription(ctx, req.(*CreateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinDescription(ctx, req.(*GetCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCoinDescriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCoinDescriptions(ctx, req.(*GetCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateCoinDescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateCoinDescription(ctx, req.(*UpdateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.middleware.app.coin.description.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinDescription",
			Handler:    _Middleware_CreateCoinDescription_Handler,
		},
		{
			MethodName: "GetCoinDescription",
			Handler:    _Middleware_GetCoinDescription_Handler,
		},
		{
			MethodName: "GetCoinDescriptions",
			Handler:    _Middleware_GetCoinDescriptions_Handler,
		},
		{
			MethodName: "UpdateCoinDescription",
			Handler:    _Middleware_UpdateCoinDescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mw/v1/app/coin/description/description.proto",
}
