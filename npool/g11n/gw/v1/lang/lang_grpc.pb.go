// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/g11n/gw/v1/lang/lang.proto

package lang

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
	CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error)
	CreateLangs(ctx context.Context, in *CreateLangsRequest, opts ...grpc.CallOption) (*CreateLangsResponse, error)
	UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error)
	GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error) {
	out := new(CreateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.lang.v1.Gateway/CreateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateLangs(ctx context.Context, in *CreateLangsRequest, opts ...grpc.CallOption) (*CreateLangsResponse, error) {
	out := new(CreateLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.lang.v1.Gateway/CreateLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error) {
	out := new(UpdateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.lang.v1.Gateway/UpdateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error) {
	out := new(GetLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.lang.v1.Gateway/GetLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error)
	CreateLangs(context.Context, *CreateLangsRequest) (*CreateLangsResponse, error)
	UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error)
	GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLang not implemented")
}
func (UnimplementedGatewayServer) CreateLangs(context.Context, *CreateLangsRequest) (*CreateLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLangs not implemented")
}
func (UnimplementedGatewayServer) UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLang not implemented")
}
func (UnimplementedGatewayServer) GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLangs not implemented")
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

func _Gateway_CreateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.lang.v1.Gateway/CreateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateLang(ctx, req.(*CreateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.lang.v1.Gateway/CreateLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateLangs(ctx, req.(*CreateLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.lang.v1.Gateway/UpdateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateLang(ctx, req.(*UpdateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.lang.v1.Gateway/GetLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetLangs(ctx, req.(*GetLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.gateway.lang.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLang",
			Handler:    _Gateway_CreateLang_Handler,
		},
		{
			MethodName: "CreateLangs",
			Handler:    _Gateway_CreateLangs_Handler,
		},
		{
			MethodName: "UpdateLang",
			Handler:    _Gateway_UpdateLang_Handler,
		},
		{
			MethodName: "GetLangs",
			Handler:    _Gateway_GetLangs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/gw/v1/lang/lang.proto",
}
