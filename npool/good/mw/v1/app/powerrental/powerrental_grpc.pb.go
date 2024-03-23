// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/app/powerrental/powerrental.proto

package powerrental

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
	Middleware_CreatePowerRental_FullMethodName = "/good.middleware.app.powerrental.v1.Middleware/CreatePowerRental"
	Middleware_UpdatePowerRental_FullMethodName = "/good.middleware.app.powerrental.v1.Middleware/UpdatePowerRental"
	Middleware_GetPowerRental_FullMethodName    = "/good.middleware.app.powerrental.v1.Middleware/GetPowerRental"
	Middleware_GetPowerRentals_FullMethodName   = "/good.middleware.app.powerrental.v1.Middleware/GetPowerRentals"
	Middleware_DeletePowerRental_FullMethodName = "/good.middleware.app.powerrental.v1.Middleware/DeletePowerRental"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreatePowerRental(ctx context.Context, in *CreatePowerRentalRequest, opts ...grpc.CallOption) (*CreatePowerRentalResponse, error)
	UpdatePowerRental(ctx context.Context, in *UpdatePowerRentalRequest, opts ...grpc.CallOption) (*UpdatePowerRentalResponse, error)
	GetPowerRental(ctx context.Context, in *GetPowerRentalRequest, opts ...grpc.CallOption) (*GetPowerRentalResponse, error)
	GetPowerRentals(ctx context.Context, in *GetPowerRentalsRequest, opts ...grpc.CallOption) (*GetPowerRentalsResponse, error)
	DeletePowerRental(ctx context.Context, in *DeletePowerRentalRequest, opts ...grpc.CallOption) (*DeletePowerRentalResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreatePowerRental(ctx context.Context, in *CreatePowerRentalRequest, opts ...grpc.CallOption) (*CreatePowerRentalResponse, error) {
	out := new(CreatePowerRentalResponse)
	err := c.cc.Invoke(ctx, Middleware_CreatePowerRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdatePowerRental(ctx context.Context, in *UpdatePowerRentalRequest, opts ...grpc.CallOption) (*UpdatePowerRentalResponse, error) {
	out := new(UpdatePowerRentalResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdatePowerRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPowerRental(ctx context.Context, in *GetPowerRentalRequest, opts ...grpc.CallOption) (*GetPowerRentalResponse, error) {
	out := new(GetPowerRentalResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPowerRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetPowerRentals(ctx context.Context, in *GetPowerRentalsRequest, opts ...grpc.CallOption) (*GetPowerRentalsResponse, error) {
	out := new(GetPowerRentalsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetPowerRentals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeletePowerRental(ctx context.Context, in *DeletePowerRentalRequest, opts ...grpc.CallOption) (*DeletePowerRentalResponse, error) {
	out := new(DeletePowerRentalResponse)
	err := c.cc.Invoke(ctx, Middleware_DeletePowerRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreatePowerRental(context.Context, *CreatePowerRentalRequest) (*CreatePowerRentalResponse, error)
	UpdatePowerRental(context.Context, *UpdatePowerRentalRequest) (*UpdatePowerRentalResponse, error)
	GetPowerRental(context.Context, *GetPowerRentalRequest) (*GetPowerRentalResponse, error)
	GetPowerRentals(context.Context, *GetPowerRentalsRequest) (*GetPowerRentalsResponse, error)
	DeletePowerRental(context.Context, *DeletePowerRentalRequest) (*DeletePowerRentalResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreatePowerRental(context.Context, *CreatePowerRentalRequest) (*CreatePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePowerRental not implemented")
}
func (UnimplementedMiddlewareServer) UpdatePowerRental(context.Context, *UpdatePowerRentalRequest) (*UpdatePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePowerRental not implemented")
}
func (UnimplementedMiddlewareServer) GetPowerRental(context.Context, *GetPowerRentalRequest) (*GetPowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRental not implemented")
}
func (UnimplementedMiddlewareServer) GetPowerRentals(context.Context, *GetPowerRentalsRequest) (*GetPowerRentalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRentals not implemented")
}
func (UnimplementedMiddlewareServer) DeletePowerRental(context.Context, *DeletePowerRentalRequest) (*DeletePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePowerRental not implemented")
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

func _Middleware_CreatePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreatePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreatePowerRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreatePowerRental(ctx, req.(*CreatePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdatePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdatePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdatePowerRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdatePowerRental(ctx, req.(*UpdatePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPowerRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPowerRental(ctx, req.(*GetPowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetPowerRentals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetPowerRentals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetPowerRentals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetPowerRentals(ctx, req.(*GetPowerRentalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeletePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeletePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeletePowerRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeletePowerRental(ctx, req.(*DeletePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.app.powerrental.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePowerRental",
			Handler:    _Middleware_CreatePowerRental_Handler,
		},
		{
			MethodName: "UpdatePowerRental",
			Handler:    _Middleware_UpdatePowerRental_Handler,
		},
		{
			MethodName: "GetPowerRental",
			Handler:    _Middleware_GetPowerRental_Handler,
		},
		{
			MethodName: "GetPowerRentals",
			Handler:    _Middleware_GetPowerRentals_Handler,
		},
		{
			MethodName: "DeletePowerRental",
			Handler:    _Middleware_DeletePowerRental_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/app/powerrental/powerrental.proto",
}
