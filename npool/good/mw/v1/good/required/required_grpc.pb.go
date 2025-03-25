// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/good/required/required.proto

package required

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
	CreateRequired(ctx context.Context, in *CreateRequiredRequest, opts ...grpc.CallOption) (*CreateRequiredResponse, error)
	GetRequired(ctx context.Context, in *GetRequiredRequest, opts ...grpc.CallOption) (*GetRequiredResponse, error)
	ExistRequiredConds(ctx context.Context, in *ExistRequiredCondsRequest, opts ...grpc.CallOption) (*ExistRequiredCondsResponse, error)
	GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error)
	UpdateRequired(ctx context.Context, in *UpdateRequiredRequest, opts ...grpc.CallOption) (*UpdateRequiredResponse, error)
	DeleteRequired(ctx context.Context, in *DeleteRequiredRequest, opts ...grpc.CallOption) (*DeleteRequiredResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateRequired(ctx context.Context, in *CreateRequiredRequest, opts ...grpc.CallOption) (*CreateRequiredResponse, error) {
	out := new(CreateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/CreateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRequired(ctx context.Context, in *GetRequiredRequest, opts ...grpc.CallOption) (*GetRequiredResponse, error) {
	out := new(GetRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/GetRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistRequiredConds(ctx context.Context, in *ExistRequiredCondsRequest, opts ...grpc.CallOption) (*ExistRequiredCondsResponse, error) {
	out := new(ExistRequiredCondsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/ExistRequiredConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error) {
	out := new(GetRequiredsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/GetRequireds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateRequired(ctx context.Context, in *UpdateRequiredRequest, opts ...grpc.CallOption) (*UpdateRequiredResponse, error) {
	out := new(UpdateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/UpdateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteRequired(ctx context.Context, in *DeleteRequiredRequest, opts ...grpc.CallOption) (*DeleteRequiredResponse, error) {
	out := new(DeleteRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.good1.required1.v1.Middleware/DeleteRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateRequired(context.Context, *CreateRequiredRequest) (*CreateRequiredResponse, error)
	GetRequired(context.Context, *GetRequiredRequest) (*GetRequiredResponse, error)
	ExistRequiredConds(context.Context, *ExistRequiredCondsRequest) (*ExistRequiredCondsResponse, error)
	GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error)
	UpdateRequired(context.Context, *UpdateRequiredRequest) (*UpdateRequiredResponse, error)
	DeleteRequired(context.Context, *DeleteRequiredRequest) (*DeleteRequiredResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateRequired(context.Context, *CreateRequiredRequest) (*CreateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequired not implemented")
}
func (UnimplementedMiddlewareServer) GetRequired(context.Context, *GetRequiredRequest) (*GetRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequired not implemented")
}
func (UnimplementedMiddlewareServer) ExistRequiredConds(context.Context, *ExistRequiredCondsRequest) (*ExistRequiredCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistRequiredConds not implemented")
}
func (UnimplementedMiddlewareServer) GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequireds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateRequired(context.Context, *UpdateRequiredRequest) (*UpdateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequired not implemented")
}
func (UnimplementedMiddlewareServer) DeleteRequired(context.Context, *DeleteRequiredRequest) (*DeleteRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequired not implemented")
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

func _Middleware_CreateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/CreateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateRequired(ctx, req.(*CreateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/GetRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRequired(ctx, req.(*GetRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistRequiredConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRequiredCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistRequiredConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/ExistRequiredConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistRequiredConds(ctx, req.(*ExistRequiredCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRequireds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRequireds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/GetRequireds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRequireds(ctx, req.(*GetRequiredsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/UpdateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateRequired(ctx, req.(*UpdateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.good1.required1.v1.Middleware/DeleteRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteRequired(ctx, req.(*DeleteRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.good1.required1.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRequired",
			Handler:    _Middleware_CreateRequired_Handler,
		},
		{
			MethodName: "GetRequired",
			Handler:    _Middleware_GetRequired_Handler,
		},
		{
			MethodName: "ExistRequiredConds",
			Handler:    _Middleware_ExistRequiredConds_Handler,
		},
		{
			MethodName: "GetRequireds",
			Handler:    _Middleware_GetRequireds_Handler,
		},
		{
			MethodName: "UpdateRequired",
			Handler:    _Middleware_UpdateRequired_Handler,
		},
		{
			MethodName: "DeleteRequired",
			Handler:    _Middleware_DeleteRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/good/required/required.proto",
}
