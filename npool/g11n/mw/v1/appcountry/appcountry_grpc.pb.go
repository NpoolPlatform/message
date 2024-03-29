// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/g11n/mw/v1/appcountry/appcountry.proto

package appcountry

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
	Middleware_CreateCountry_FullMethodName     = "/g11n.middleware.appcountry.v1.Middleware/CreateCountry"
	Middleware_CreateCountries_FullMethodName   = "/g11n.middleware.appcountry.v1.Middleware/CreateCountries"
	Middleware_GetCountries_FullMethodName      = "/g11n.middleware.appcountry.v1.Middleware/GetCountries"
	Middleware_ExistCountryConds_FullMethodName = "/g11n.middleware.appcountry.v1.Middleware/ExistCountryConds"
	Middleware_DeleteCountry_FullMethodName     = "/g11n.middleware.appcountry.v1.Middleware/DeleteCountry"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error)
	CreateCountries(ctx context.Context, in *CreateCountriesRequest, opts ...grpc.CallOption) (*CreateCountriesResponse, error)
	GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error)
	ExistCountryConds(ctx context.Context, in *ExistCountryCondsRequest, opts ...grpc.CallOption) (*ExistCountryCondsResponse, error)
	DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error) {
	out := new(CreateCountryResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateCountries(ctx context.Context, in *CreateCountriesRequest, opts ...grpc.CallOption) (*CreateCountriesResponse, error) {
	out := new(CreateCountriesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCountries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error) {
	out := new(GetCountriesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCountries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistCountryConds(ctx context.Context, in *ExistCountryCondsRequest, opts ...grpc.CallOption) (*ExistCountryCondsResponse, error) {
	out := new(ExistCountryCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistCountryConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error) {
	out := new(DeleteCountryResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error)
	CreateCountries(context.Context, *CreateCountriesRequest) (*CreateCountriesResponse, error)
	GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesResponse, error)
	ExistCountryConds(context.Context, *ExistCountryCondsRequest) (*ExistCountryCondsResponse, error)
	DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountry not implemented")
}
func (UnimplementedMiddlewareServer) CreateCountries(context.Context, *CreateCountriesRequest) (*CreateCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountries not implemented")
}
func (UnimplementedMiddlewareServer) GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountries not implemented")
}
func (UnimplementedMiddlewareServer) ExistCountryConds(context.Context, *ExistCountryCondsRequest) (*ExistCountryCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCountryConds not implemented")
}
func (UnimplementedMiddlewareServer) DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCountry not implemented")
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

func _Middleware_CreateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCountry(ctx, req.(*CreateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateCountries(ctx, req.(*CreateCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCountries(ctx, req.(*GetCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistCountryConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCountryCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistCountryConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistCountryConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistCountryConds(ctx, req.(*ExistCountryCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteCountry(ctx, req.(*DeleteCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.middleware.appcountry.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCountry",
			Handler:    _Middleware_CreateCountry_Handler,
		},
		{
			MethodName: "CreateCountries",
			Handler:    _Middleware_CreateCountries_Handler,
		},
		{
			MethodName: "GetCountries",
			Handler:    _Middleware_GetCountries_Handler,
		},
		{
			MethodName: "ExistCountryConds",
			Handler:    _Middleware_ExistCountryConds_Handler,
		},
		{
			MethodName: "DeleteCountry",
			Handler:    _Middleware_DeleteCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/mw/v1/appcountry/appcountry.proto",
}
