// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/good/required/required.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateRequired(ctx context.Context, in *AdminCreateRequiredRequest, opts ...grpc.CallOption) (*AdminCreateRequiredResponse, error)
	AdminUpdateRequired(ctx context.Context, in *AdminUpdateRequiredRequest, opts ...grpc.CallOption) (*AdminUpdateRequiredResponse, error)
	GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error)
	AdminDeleteRequired(ctx context.Context, in *AdminDeleteRequiredRequest, opts ...grpc.CallOption) (*AdminDeleteRequiredResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateRequired(ctx context.Context, in *AdminCreateRequiredRequest, opts ...grpc.CallOption) (*AdminCreateRequiredResponse, error) {
	out := new(AdminCreateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.good1.required1.v1.Gateway/AdminCreateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateRequired(ctx context.Context, in *AdminUpdateRequiredRequest, opts ...grpc.CallOption) (*AdminUpdateRequiredResponse, error) {
	out := new(AdminUpdateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.good1.required1.v1.Gateway/AdminUpdateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error) {
	out := new(GetRequiredsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.good1.required1.v1.Gateway/GetRequireds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteRequired(ctx context.Context, in *AdminDeleteRequiredRequest, opts ...grpc.CallOption) (*AdminDeleteRequiredResponse, error) {
	out := new(AdminDeleteRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.good1.required1.v1.Gateway/AdminDeleteRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateRequired(context.Context, *AdminCreateRequiredRequest) (*AdminCreateRequiredResponse, error)
	AdminUpdateRequired(context.Context, *AdminUpdateRequiredRequest) (*AdminUpdateRequiredResponse, error)
	GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error)
	AdminDeleteRequired(context.Context, *AdminDeleteRequiredRequest) (*AdminDeleteRequiredResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreateRequired(context.Context, *AdminCreateRequiredRequest) (*AdminCreateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateRequired not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateRequired(context.Context, *AdminUpdateRequiredRequest) (*AdminUpdateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateRequired not implemented")
}
func (UnimplementedGatewayServer) GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequireds not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteRequired(context.Context, *AdminDeleteRequiredRequest) (*AdminDeleteRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteRequired not implemented")
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

func _Gateway_AdminCreateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.good1.required1.v1.Gateway/AdminCreateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateRequired(ctx, req.(*AdminCreateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.good1.required1.v1.Gateway/AdminUpdateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateRequired(ctx, req.(*AdminUpdateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetRequireds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetRequireds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.good1.required1.v1.Gateway/GetRequireds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRequireds(ctx, req.(*GetRequiredsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.good1.required1.v1.Gateway/AdminDeleteRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteRequired(ctx, req.(*AdminDeleteRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.good1.required1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateRequired",
			Handler:    _Gateway_AdminCreateRequired_Handler,
		},
		{
			MethodName: "AdminUpdateRequired",
			Handler:    _Gateway_AdminUpdateRequired_Handler,
		},
		{
			MethodName: "GetRequireds",
			Handler:    _Gateway_GetRequireds_Handler,
		},
		{
			MethodName: "AdminDeleteRequired",
			Handler:    _Gateway_AdminDeleteRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/good/required/required.proto",
}
