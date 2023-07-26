// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/invitation/registration/registration.proto

package registration

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
<<<<<<< HEAD
	Middleware_CreateRegistration_FullMethodName     = "/inspire.middleware.invitation.registration.v1.Middleware/CreateRegistration"
	Middleware_UpdateRegistration_FullMethodName     = "/inspire.middleware.invitation.registration.v1.Middleware/UpdateRegistration"
	Middleware_GetRegistration_FullMethodName        = "/inspire.middleware.invitation.registration.v1.Middleware/GetRegistration"
	Middleware_ExistRegistrationConds_FullMethodName = "/inspire.middleware.invitation.registration.v1.Middleware/ExistRegistrationConds"
	Middleware_GetRegistrations_FullMethodName       = "/inspire.middleware.invitation.registration.v1.Middleware/GetRegistrations"
	Middleware_GetSubordinates_FullMethodName        = "/inspire.middleware.invitation.registration.v1.Middleware/GetSubordinates"
	Middleware_GetSuperiores_FullMethodName          = "/inspire.middleware.invitation.registration.v1.Middleware/GetSuperiores"
	Middleware_DeleteRegistration_FullMethodName     = "/inspire.middleware.invitation.registration.v1.Middleware/DeleteRegistration"
=======
	Middleware_CreateRegistration_FullMethodName = "/inspire.middleware.invitation.registration.v1.Middleware/CreateRegistration"
	Middleware_UpdateRegistration_FullMethodName = "/inspire.middleware.invitation.registration.v1.Middleware/UpdateRegistration"
	Middleware_GetRegistration_FullMethodName    = "/inspire.middleware.invitation.registration.v1.Middleware/GetRegistration"
	Middleware_GetRegistrations_FullMethodName   = "/inspire.middleware.invitation.registration.v1.Middleware/GetRegistrations"
	Middleware_GetSubordinates_FullMethodName    = "/inspire.middleware.invitation.registration.v1.Middleware/GetSubordinates"
	Middleware_GetSuperiores_FullMethodName      = "/inspire.middleware.invitation.registration.v1.Middleware/GetSuperiores"
	Middleware_DeleteRegistration_FullMethodName = "/inspire.middleware.invitation.registration.v1.Middleware/DeleteRegistration"
>>>>>>> Format invitation
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	// Method Version
	CreateRegistration(ctx context.Context, in *CreateRegistrationRequest, opts ...grpc.CallOption) (*CreateRegistrationResponse, error)
	UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*UpdateRegistrationResponse, error)
	GetRegistration(ctx context.Context, in *GetRegistrationRequest, opts ...grpc.CallOption) (*GetRegistrationResponse, error)
	ExistRegistrationConds(ctx context.Context, in *ExistRegistrationCondsRequest, opts ...grpc.CallOption) (*ExistRegistrationCondsResponse, error)
	GetRegistrations(ctx context.Context, in *GetRegistrationsRequest, opts ...grpc.CallOption) (*GetRegistrationsResponse, error)
	GetSubordinates(ctx context.Context, in *GetSubordinatesRequest, opts ...grpc.CallOption) (*GetSubordinatesResponse, error)
	GetSuperiores(ctx context.Context, in *GetSuperioresRequest, opts ...grpc.CallOption) (*GetSuperioresResponse, error)
	DeleteRegistration(ctx context.Context, in *DeleteRegistrationRequest, opts ...grpc.CallOption) (*DeleteRegistrationResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateRegistration(ctx context.Context, in *CreateRegistrationRequest, opts ...grpc.CallOption) (*CreateRegistrationResponse, error) {
	out := new(CreateRegistrationResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*UpdateRegistrationResponse, error) {
	out := new(UpdateRegistrationResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetRegistration(ctx context.Context, in *GetRegistrationRequest, opts ...grpc.CallOption) (*GetRegistrationResponse, error) {
	out := new(GetRegistrationResponse)
	err := c.cc.Invoke(ctx, Middleware_GetRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistRegistrationConds(ctx context.Context, in *ExistRegistrationCondsRequest, opts ...grpc.CallOption) (*ExistRegistrationCondsResponse, error) {
	out := new(ExistRegistrationCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistRegistrationConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) GetRegistrations(ctx context.Context, in *GetRegistrationsRequest, opts ...grpc.CallOption) (*GetRegistrationsResponse, error) {
	out := new(GetRegistrationsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetRegistrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

=======
>>>>>>> Format invitation
func (c *middlewareClient) GetSubordinates(ctx context.Context, in *GetSubordinatesRequest, opts ...grpc.CallOption) (*GetSubordinatesResponse, error) {
	out := new(GetSubordinatesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSubordinates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetSuperiores(ctx context.Context, in *GetSuperioresRequest, opts ...grpc.CallOption) (*GetSuperioresResponse, error) {
	out := new(GetSuperioresResponse)
	err := c.cc.Invoke(ctx, Middleware_GetSuperiores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteRegistration(ctx context.Context, in *DeleteRegistrationRequest, opts ...grpc.CallOption) (*DeleteRegistrationResponse, error) {
	out := new(DeleteRegistrationResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	// Method Version
	CreateRegistration(context.Context, *CreateRegistrationRequest) (*CreateRegistrationResponse, error)
	UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*UpdateRegistrationResponse, error)
	GetRegistration(context.Context, *GetRegistrationRequest) (*GetRegistrationResponse, error)
	ExistRegistrationConds(context.Context, *ExistRegistrationCondsRequest) (*ExistRegistrationCondsResponse, error)
	GetRegistrations(context.Context, *GetRegistrationsRequest) (*GetRegistrationsResponse, error)
	GetSubordinates(context.Context, *GetSubordinatesRequest) (*GetSubordinatesResponse, error)
	GetSuperiores(context.Context, *GetSuperioresRequest) (*GetSuperioresResponse, error)
	DeleteRegistration(context.Context, *DeleteRegistrationRequest) (*DeleteRegistrationResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateRegistration(context.Context, *CreateRegistrationRequest) (*CreateRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegistration not implemented")
}
func (UnimplementedMiddlewareServer) UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*UpdateRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegistration not implemented")
}
func (UnimplementedMiddlewareServer) GetRegistration(context.Context, *GetRegistrationRequest) (*GetRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistration not implemented")
}
func (UnimplementedMiddlewareServer) ExistRegistrationConds(context.Context, *ExistRegistrationCondsRequest) (*ExistRegistrationCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistRegistrationConds not implemented")
}
func (UnimplementedMiddlewareServer) GetRegistrations(context.Context, *GetRegistrationsRequest) (*GetRegistrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegistrations not implemented")
}
func (UnimplementedMiddlewareServer) GetSubordinates(context.Context, *GetSubordinatesRequest) (*GetSubordinatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubordinates not implemented")
}
func (UnimplementedMiddlewareServer) GetSuperiores(context.Context, *GetSuperioresRequest) (*GetSuperioresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuperiores not implemented")
}
func (UnimplementedMiddlewareServer) DeleteRegistration(context.Context, *DeleteRegistrationRequest) (*DeleteRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegistration not implemented")
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

func _Middleware_CreateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateRegistration(ctx, req.(*CreateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateRegistration(ctx, req.(*UpdateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRegistration(ctx, req.(*GetRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistRegistrationConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistRegistrationCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistRegistrationConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistRegistrationConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistRegistrationConds(ctx, req.(*ExistRegistrationCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
func _Middleware_GetRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetRegistrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetRegistrations(ctx, req.(*GetRegistrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

=======
>>>>>>> Format invitation
func _Middleware_GetSubordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubordinatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSubordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSubordinates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSubordinates(ctx, req.(*GetSubordinatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetSuperiores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperioresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetSuperiores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetSuperiores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetSuperiores(ctx, req.(*GetSuperioresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteRegistration(ctx, req.(*DeleteRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.invitation.registration.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegistration",
			Handler:    _Middleware_CreateRegistration_Handler,
		},
		{
			MethodName: "UpdateRegistration",
			Handler:    _Middleware_UpdateRegistration_Handler,
		},
		{
			MethodName: "GetRegistration",
			Handler:    _Middleware_GetRegistration_Handler,
		},
		{
			MethodName: "ExistRegistrationConds",
			Handler:    _Middleware_ExistRegistrationConds_Handler,
		},
		{
<<<<<<< HEAD
			MethodName: "GetRegistrations",
			Handler:    _Middleware_GetRegistrations_Handler,
		},
		{
=======
>>>>>>> Format invitation
			MethodName: "GetSubordinates",
			Handler:    _Middleware_GetSubordinates_Handler,
		},
		{
			MethodName: "GetSuperiores",
			Handler:    _Middleware_GetSuperiores_Handler,
		},
		{
			MethodName: "DeleteRegistration",
			Handler:    _Middleware_DeleteRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/invitation/registration/registration.proto",
}
