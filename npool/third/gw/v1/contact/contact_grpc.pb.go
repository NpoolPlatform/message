// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/third/gw/v1/contact/contact.proto

package contact

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateContact(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error)
	CreateAppContact(ctx context.Context, in *CreateAppContactRequest, opts ...grpc.CallOption) (*CreateAppContactResponse, error)
	GetContact(ctx context.Context, in *GetContactRequest, opts ...grpc.CallOption) (*GetContactResponse, error)
	GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error)
	GetAppContacts(ctx context.Context, in *GetAppContactsRequest, opts ...grpc.CallOption) (*GetAppContactsResponse, error)
	UpdateContact(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error)
	ContactViaEmail(ctx context.Context, in *ContactViaEmailRequest, opts ...grpc.CallOption) (*ContactViaEmailResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateContact(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error) {
	out := new(CreateContactResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/CreateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppContact(ctx context.Context, in *CreateAppContactRequest, opts ...grpc.CallOption) (*CreateAppContactResponse, error) {
	out := new(CreateAppContactResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/CreateAppContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetContact(ctx context.Context, in *GetContactRequest, opts ...grpc.CallOption) (*GetContactResponse, error) {
	out := new(GetContactResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error) {
	out := new(GetContactsResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/GetContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppContacts(ctx context.Context, in *GetAppContactsRequest, opts ...grpc.CallOption) (*GetAppContactsResponse, error) {
	out := new(GetAppContactsResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/GetAppContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateContact(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error) {
	out := new(UpdateContactResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/UpdateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) ContactViaEmail(ctx context.Context, in *ContactViaEmailRequest, opts ...grpc.CallOption) (*ContactViaEmailResponse, error) {
	out := new(ContactViaEmailResponse)
	err := c.cc.Invoke(ctx, "/third.gateway.contact.v1.Gateway/ContactViaEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateContact(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
	CreateAppContact(context.Context, *CreateAppContactRequest) (*CreateAppContactResponse, error)
	GetContact(context.Context, *GetContactRequest) (*GetContactResponse, error)
	GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error)
	GetAppContacts(context.Context, *GetAppContactsRequest) (*GetAppContactsResponse, error)
	UpdateContact(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
	ContactViaEmail(context.Context, *ContactViaEmailRequest) (*ContactViaEmailResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateContact(context.Context, *CreateContactRequest) (*CreateContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (UnimplementedGatewayServer) CreateAppContact(context.Context, *CreateAppContactRequest) (*CreateAppContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppContact not implemented")
}
func (UnimplementedGatewayServer) GetContact(context.Context, *GetContactRequest) (*GetContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedGatewayServer) GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedGatewayServer) GetAppContacts(context.Context, *GetAppContactsRequest) (*GetAppContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppContacts not implemented")
}
func (UnimplementedGatewayServer) UpdateContact(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (UnimplementedGatewayServer) ContactViaEmail(context.Context, *ContactViaEmailRequest) (*ContactViaEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContactViaEmail not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/CreateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateContact(ctx, req.(*CreateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/CreateAppContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppContact(ctx, req.(*CreateAppContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetContact(ctx, req.(*GetContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/GetContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetContacts(ctx, req.(*GetContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/GetAppContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppContacts(ctx, req.(*GetAppContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateContact(ctx, req.(*UpdateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_ContactViaEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactViaEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).ContactViaEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/third.gateway.contact.v1.Gateway/ContactViaEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).ContactViaEmail(ctx, req.(*ContactViaEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "third.gateway.contact.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _Gateway_CreateContact_Handler,
		},
		{
			MethodName: "CreateAppContact",
			Handler:    _Gateway_CreateAppContact_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Gateway_GetContact_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _Gateway_GetContacts_Handler,
		},
		{
			MethodName: "GetAppContacts",
			Handler:    _Gateway_GetAppContacts_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _Gateway_UpdateContact_Handler,
		},
		{
			MethodName: "ContactViaEmail",
			Handler:    _Gateway_ContactViaEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/third/gw/v1/contact/contact.proto",
}
