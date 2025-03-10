// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/device/manufacturer/manufacturer.proto

package manufacturer

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
	CreateManufacturer(ctx context.Context, in *CreateManufacturerRequest, opts ...grpc.CallOption) (*CreateManufacturerResponse, error)
	UpdateManufacturer(ctx context.Context, in *UpdateManufacturerRequest, opts ...grpc.CallOption) (*UpdateManufacturerResponse, error)
	GetManufacturer(ctx context.Context, in *GetManufacturerRequest, opts ...grpc.CallOption) (*GetManufacturerResponse, error)
	GetManufacturers(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*GetManufacturersResponse, error)
	ExistManufacturerConds(ctx context.Context, in *ExistManufacturerCondsRequest, opts ...grpc.CallOption) (*ExistManufacturerCondsResponse, error)
	DeleteManufacturer(ctx context.Context, in *DeleteManufacturerRequest, opts ...grpc.CallOption) (*DeleteManufacturerResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateManufacturer(ctx context.Context, in *CreateManufacturerRequest, opts ...grpc.CallOption) (*CreateManufacturerResponse, error) {
	out := new(CreateManufacturerResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/CreateManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateManufacturer(ctx context.Context, in *UpdateManufacturerRequest, opts ...grpc.CallOption) (*UpdateManufacturerResponse, error) {
	out := new(UpdateManufacturerResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/UpdateManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetManufacturer(ctx context.Context, in *GetManufacturerRequest, opts ...grpc.CallOption) (*GetManufacturerResponse, error) {
	out := new(GetManufacturerResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/GetManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetManufacturers(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*GetManufacturersResponse, error) {
	out := new(GetManufacturersResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/GetManufacturers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistManufacturerConds(ctx context.Context, in *ExistManufacturerCondsRequest, opts ...grpc.CallOption) (*ExistManufacturerCondsResponse, error) {
	out := new(ExistManufacturerCondsResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/ExistManufacturerConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteManufacturer(ctx context.Context, in *DeleteManufacturerRequest, opts ...grpc.CallOption) (*DeleteManufacturerResponse, error) {
	out := new(DeleteManufacturerResponse)
	err := c.cc.Invoke(ctx, "/good.middleware.device.manufacturer.v1.Middleware/DeleteManufacturer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateManufacturer(context.Context, *CreateManufacturerRequest) (*CreateManufacturerResponse, error)
	UpdateManufacturer(context.Context, *UpdateManufacturerRequest) (*UpdateManufacturerResponse, error)
	GetManufacturer(context.Context, *GetManufacturerRequest) (*GetManufacturerResponse, error)
	GetManufacturers(context.Context, *GetManufacturersRequest) (*GetManufacturersResponse, error)
	ExistManufacturerConds(context.Context, *ExistManufacturerCondsRequest) (*ExistManufacturerCondsResponse, error)
	DeleteManufacturer(context.Context, *DeleteManufacturerRequest) (*DeleteManufacturerResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateManufacturer(context.Context, *CreateManufacturerRequest) (*CreateManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManufacturer not implemented")
}
func (UnimplementedMiddlewareServer) UpdateManufacturer(context.Context, *UpdateManufacturerRequest) (*UpdateManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateManufacturer not implemented")
}
func (UnimplementedMiddlewareServer) GetManufacturer(context.Context, *GetManufacturerRequest) (*GetManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManufacturer not implemented")
}
func (UnimplementedMiddlewareServer) GetManufacturers(context.Context, *GetManufacturersRequest) (*GetManufacturersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManufacturers not implemented")
}
func (UnimplementedMiddlewareServer) ExistManufacturerConds(context.Context, *ExistManufacturerCondsRequest) (*ExistManufacturerCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistManufacturerConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteManufacturer(context.Context, *DeleteManufacturerRequest) (*DeleteManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManufacturer not implemented")
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

func _Middleware_CreateManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/CreateManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateManufacturer(ctx, req.(*CreateManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/UpdateManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateManufacturer(ctx, req.(*UpdateManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/GetManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetManufacturer(ctx, req.(*GetManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetManufacturers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManufacturersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetManufacturers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/GetManufacturers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetManufacturers(ctx, req.(*GetManufacturersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistManufacturerConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistManufacturerCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistManufacturerConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/ExistManufacturerConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistManufacturerConds(ctx, req.(*ExistManufacturerCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.middleware.device.manufacturer.v1.Middleware/DeleteManufacturer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteManufacturer(ctx, req.(*DeleteManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.device.manufacturer.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateManufacturer",
			Handler:    _Middleware_CreateManufacturer_Handler,
		},
		{
			MethodName: "UpdateManufacturer",
			Handler:    _Middleware_UpdateManufacturer_Handler,
		},
		{
			MethodName: "GetManufacturer",
			Handler:    _Middleware_GetManufacturer_Handler,
		},
		{
			MethodName: "GetManufacturers",
			Handler:    _Middleware_GetManufacturers_Handler,
		},
		{
			MethodName: "ExistManufacturerConds",
			Handler:    _Middleware_ExistManufacturerConds_Handler,
		},
		{
			MethodName: "DeleteManufacturer",
			Handler:    _Middleware_DeleteManufacturer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/device/manufacturer/manufacturer.proto",
}
