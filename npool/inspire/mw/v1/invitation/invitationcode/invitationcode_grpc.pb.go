// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/invitation/invitationcode/invitationcode.proto

package invitationcode

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
	CreateInvitationCode(ctx context.Context, in *CreateInvitationCodeRequest, opts ...grpc.CallOption) (*CreateInvitationCodeResponse, error)
	UpdateInvitationCode(ctx context.Context, in *UpdateInvitationCodeRequest, opts ...grpc.CallOption) (*UpdateInvitationCodeResponse, error)
	GetInvitationCode(ctx context.Context, in *GetInvitationCodeRequest, opts ...grpc.CallOption) (*GetInvitationCodeResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateInvitationCode(ctx context.Context, in *CreateInvitationCodeRequest, opts ...grpc.CallOption) (*CreateInvitationCodeResponse, error) {
	out := new(CreateInvitationCodeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.invitation.invitationcode.v1.Middleware/CreateInvitationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateInvitationCode(ctx context.Context, in *UpdateInvitationCodeRequest, opts ...grpc.CallOption) (*UpdateInvitationCodeResponse, error) {
	out := new(UpdateInvitationCodeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.invitation.invitationcode.v1.Middleware/UpdateInvitationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetInvitationCode(ctx context.Context, in *GetInvitationCodeRequest, opts ...grpc.CallOption) (*GetInvitationCodeResponse, error) {
	out := new(GetInvitationCodeResponse)
	err := c.cc.Invoke(ctx, "/inspire.middleware.invitation.invitationcode.v1.Middleware/GetInvitationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateInvitationCode(context.Context, *CreateInvitationCodeRequest) (*CreateInvitationCodeResponse, error)
	UpdateInvitationCode(context.Context, *UpdateInvitationCodeRequest) (*UpdateInvitationCodeResponse, error)
	GetInvitationCode(context.Context, *GetInvitationCodeRequest) (*GetInvitationCodeResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateInvitationCode(context.Context, *CreateInvitationCodeRequest) (*CreateInvitationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvitationCode not implemented")
}
func (UnimplementedMiddlewareServer) UpdateInvitationCode(context.Context, *UpdateInvitationCodeRequest) (*UpdateInvitationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvitationCode not implemented")
}
func (UnimplementedMiddlewareServer) GetInvitationCode(context.Context, *GetInvitationCodeRequest) (*GetInvitationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitationCode not implemented")
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

func _Middleware_CreateInvitationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvitationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateInvitationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.invitation.invitationcode.v1.Middleware/CreateInvitationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateInvitationCode(ctx, req.(*CreateInvitationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateInvitationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvitationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateInvitationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.invitation.invitationcode.v1.Middleware/UpdateInvitationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateInvitationCode(ctx, req.(*UpdateInvitationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetInvitationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvitationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetInvitationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspire.middleware.invitation.invitationcode.v1.Middleware/GetInvitationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetInvitationCode(ctx, req.(*GetInvitationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.invitation.invitationcode.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvitationCode",
			Handler:    _Middleware_CreateInvitationCode_Handler,
		},
		{
			MethodName: "UpdateInvitationCode",
			Handler:    _Middleware_UpdateInvitationCode_Handler,
		},
		{
			MethodName: "GetInvitationCode",
			Handler:    _Middleware_GetInvitationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/invitation/invitationcode/invitationcode.proto",
}
