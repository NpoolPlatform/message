// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/mw/v1/vender/location/location.proto

package location

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
	Middleware_CreateLocation_FullMethodName     = "/good.middleware.vendor.location.v1.Middleware/CreateLocation"
	Middleware_UpdateLocation_FullMethodName     = "/good.middleware.vendor.location.v1.Middleware/UpdateLocation"
	Middleware_GetLocation_FullMethodName        = "/good.middleware.vendor.location.v1.Middleware/GetLocation"
	Middleware_GetLocations_FullMethodName       = "/good.middleware.vendor.location.v1.Middleware/GetLocations"
	Middleware_ExistLocationConds_FullMethodName = "/good.middleware.vendor.location.v1.Middleware/ExistLocationConds"
	Middleware_DeleteLocation_FullMethodName     = "/good.middleware.vendor.location.v1.Middleware/DeleteLocation"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error)
	UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error)
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
	GetLocations(ctx context.Context, in *GetLocationsRequest, opts ...grpc.CallOption) (*GetLocationsResponse, error)
	ExistLocationConds(ctx context.Context, in *ExistLocationCondsRequest, opts ...grpc.CallOption) (*ExistLocationCondsResponse, error)
	DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateLocation(ctx context.Context, in *CreateLocationRequest, opts ...grpc.CallOption) (*CreateLocationResponse, error) {
	out := new(CreateLocationResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error) {
	out := new(UpdateLocationResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := c.cc.Invoke(ctx, Middleware_GetLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetLocations(ctx context.Context, in *GetLocationsRequest, opts ...grpc.CallOption) (*GetLocationsResponse, error) {
	out := new(GetLocationsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetLocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistLocationConds(ctx context.Context, in *ExistLocationCondsRequest, opts ...grpc.CallOption) (*ExistLocationCondsResponse, error) {
	out := new(ExistLocationCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistLocationConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteLocation(ctx context.Context, in *DeleteLocationRequest, opts ...grpc.CallOption) (*DeleteLocationResponse, error) {
	out := new(DeleteLocationResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error)
	ExistLocationConds(context.Context, *ExistLocationCondsRequest) (*ExistLocationCondsResponse, error)
	DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateLocation(context.Context, *CreateLocationRequest) (*CreateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedMiddlewareServer) UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedMiddlewareServer) GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedMiddlewareServer) GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedMiddlewareServer) ExistLocationConds(context.Context, *ExistLocationCondsRequest) (*ExistLocationCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistLocationConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteLocation(context.Context, *DeleteLocationRequest) (*DeleteLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocation not implemented")
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

func _Middleware_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateLocation(ctx, req.(*CreateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateLocation(ctx, req.(*UpdateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetLocations(ctx, req.(*GetLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistLocationConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistLocationCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistLocationConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistLocationConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistLocationConds(ctx, req.(*ExistLocationCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteLocation(ctx, req.(*DeleteLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.vendor.location.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLocation",
			Handler:    _Middleware_CreateLocation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _Middleware_UpdateLocation_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Middleware_GetLocation_Handler,
		},
		{
			MethodName: "GetLocations",
			Handler:    _Middleware_GetLocations_Handler,
		},
		{
			MethodName: "ExistLocationConds",
			Handler:    _Middleware_ExistLocationConds_Handler,
		},
		{
			MethodName: "DeleteLocation",
			Handler:    _Middleware_DeleteLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mw/v1/vender/location/location.proto",
}
