// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/g11n/mw/v1/counrty/country.proto

package country

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
	Manager_CreateCountry_FullMethodName   = "/g11n.manager.country.v1.Manager/CreateCountry"
	Manager_CreateCountries_FullMethodName = "/g11n.manager.country.v1.Manager/CreateCountries"
	Manager_UpdateCountry_FullMethodName   = "/g11n.manager.country.v1.Manager/UpdateCountry"
	Manager_GetCountry_FullMethodName      = "/g11n.manager.country.v1.Manager/GetCountry"
	Manager_GetCountryOnly_FullMethodName  = "/g11n.manager.country.v1.Manager/GetCountryOnly"
	Manager_GetCountries_FullMethodName    = "/g11n.manager.country.v1.Manager/GetCountries"
	Manager_DeleteCountry_FullMethodName   = "/g11n.manager.country.v1.Manager/DeleteCountry"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error)
	CreateCountries(ctx context.Context, in *CreateCountriesRequest, opts ...grpc.CallOption) (*CreateCountriesResponse, error)
	UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error)
	GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error)
	GetCountryOnly(ctx context.Context, in *GetCountryOnlyRequest, opts ...grpc.CallOption) (*GetCountryOnlyResponse, error)
	GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error)
	DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateCountry(ctx context.Context, in *CreateCountryRequest, opts ...grpc.CallOption) (*CreateCountryResponse, error) {
	out := new(CreateCountryResponse)
	err := c.cc.Invoke(ctx, Manager_CreateCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateCountries(ctx context.Context, in *CreateCountriesRequest, opts ...grpc.CallOption) (*CreateCountriesResponse, error) {
	out := new(CreateCountriesResponse)
	err := c.cc.Invoke(ctx, Manager_CreateCountries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateCountry(ctx context.Context, in *UpdateCountryRequest, opts ...grpc.CallOption) (*UpdateCountryResponse, error) {
	out := new(UpdateCountryResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCountry(ctx context.Context, in *GetCountryRequest, opts ...grpc.CallOption) (*GetCountryResponse, error) {
	out := new(GetCountryResponse)
	err := c.cc.Invoke(ctx, Manager_GetCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCountryOnly(ctx context.Context, in *GetCountryOnlyRequest, opts ...grpc.CallOption) (*GetCountryOnlyResponse, error) {
	out := new(GetCountryOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCountryOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCountries(ctx context.Context, in *GetCountriesRequest, opts ...grpc.CallOption) (*GetCountriesResponse, error) {
	out := new(GetCountriesResponse)
	err := c.cc.Invoke(ctx, Manager_GetCountries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteCountry(ctx context.Context, in *DeleteCountryRequest, opts ...grpc.CallOption) (*DeleteCountryResponse, error) {
	out := new(DeleteCountryResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteCountry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error)
	CreateCountries(context.Context, *CreateCountriesRequest) (*CreateCountriesResponse, error)
	UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error)
	GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error)
	GetCountryOnly(context.Context, *GetCountryOnlyRequest) (*GetCountryOnlyResponse, error)
	GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesResponse, error)
	DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateCountry(context.Context, *CreateCountryRequest) (*CreateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountry not implemented")
}
func (UnimplementedManagerServer) CreateCountries(context.Context, *CreateCountriesRequest) (*CreateCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountries not implemented")
}
func (UnimplementedManagerServer) UpdateCountry(context.Context, *UpdateCountryRequest) (*UpdateCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCountry not implemented")
}
func (UnimplementedManagerServer) GetCountry(context.Context, *GetCountryRequest) (*GetCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountry not implemented")
}
func (UnimplementedManagerServer) GetCountryOnly(context.Context, *GetCountryOnlyRequest) (*GetCountryOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountryOnly not implemented")
}
func (UnimplementedManagerServer) GetCountries(context.Context, *GetCountriesRequest) (*GetCountriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountries not implemented")
}
func (UnimplementedManagerServer) DeleteCountry(context.Context, *DeleteCountryRequest) (*DeleteCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCountry not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_CreateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCountry(ctx, req.(*CreateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCountries(ctx, req.(*CreateCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateCountry(ctx, req.(*UpdateCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCountry(ctx, req.(*GetCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCountryOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountryOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCountryOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCountryOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCountryOnly(ctx, req.(*GetCountryOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCountries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCountries(ctx, req.(*GetCountriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteCountry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteCountry(ctx, req.(*DeleteCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.manager.country.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCountry",
			Handler:    _Manager_CreateCountry_Handler,
		},
		{
			MethodName: "CreateCountries",
			Handler:    _Manager_CreateCountries_Handler,
		},
		{
			MethodName: "UpdateCountry",
			Handler:    _Manager_UpdateCountry_Handler,
		},
		{
			MethodName: "GetCountry",
			Handler:    _Manager_GetCountry_Handler,
		},
		{
			MethodName: "GetCountryOnly",
			Handler:    _Manager_GetCountryOnly_Handler,
		},
		{
			MethodName: "GetCountries",
			Handler:    _Manager_GetCountries_Handler,
		},
		{
			MethodName: "DeleteCountry",
			Handler:    _Manager_DeleteCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/mw/v1/counrty/country.proto",
}
