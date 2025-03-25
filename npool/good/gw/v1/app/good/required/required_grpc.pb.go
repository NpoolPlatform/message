// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/app/good/required/required.proto

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
	CreateRequired(ctx context.Context, in *CreateRequiredRequest, opts ...grpc.CallOption) (*CreateRequiredResponse, error)
	UpdateRequired(ctx context.Context, in *UpdateRequiredRequest, opts ...grpc.CallOption) (*UpdateRequiredResponse, error)
	GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error)
	DeleteRequired(ctx context.Context, in *DeleteRequiredRequest, opts ...grpc.CallOption) (*DeleteRequiredResponse, error)
	// Run by church admin
	AdminCreateRequired(ctx context.Context, in *AdminCreateRequiredRequest, opts ...grpc.CallOption) (*AdminCreateRequiredResponse, error)
	AdminUpdateRequired(ctx context.Context, in *AdminUpdateRequiredRequest, opts ...grpc.CallOption) (*AdminUpdateRequiredResponse, error)
	AdminGetRequireds(ctx context.Context, in *AdminGetRequiredsRequest, opts ...grpc.CallOption) (*AdminGetRequiredsResponse, error)
	AdminDeleteRequired(ctx context.Context, in *AdminDeleteRequiredRequest, opts ...grpc.CallOption) (*AdminDeleteRequiredResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateRequired(ctx context.Context, in *CreateRequiredRequest, opts ...grpc.CallOption) (*CreateRequiredResponse, error) {
	out := new(CreateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/CreateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateRequired(ctx context.Context, in *UpdateRequiredRequest, opts ...grpc.CallOption) (*UpdateRequiredResponse, error) {
	out := new(UpdateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/UpdateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetRequireds(ctx context.Context, in *GetRequiredsRequest, opts ...grpc.CallOption) (*GetRequiredsResponse, error) {
	out := new(GetRequiredsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/GetRequireds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteRequired(ctx context.Context, in *DeleteRequiredRequest, opts ...grpc.CallOption) (*DeleteRequiredResponse, error) {
	out := new(DeleteRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/DeleteRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateRequired(ctx context.Context, in *AdminCreateRequiredRequest, opts ...grpc.CallOption) (*AdminCreateRequiredResponse, error) {
	out := new(AdminCreateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/AdminCreateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateRequired(ctx context.Context, in *AdminUpdateRequiredRequest, opts ...grpc.CallOption) (*AdminUpdateRequiredResponse, error) {
	out := new(AdminUpdateRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/AdminUpdateRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetRequireds(ctx context.Context, in *AdminGetRequiredsRequest, opts ...grpc.CallOption) (*AdminGetRequiredsResponse, error) {
	out := new(AdminGetRequiredsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/AdminGetRequireds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteRequired(ctx context.Context, in *AdminDeleteRequiredRequest, opts ...grpc.CallOption) (*AdminDeleteRequiredResponse, error) {
	out := new(AdminDeleteRequiredResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.app.good1.required1.v1.Gateway/AdminDeleteRequired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateRequired(context.Context, *CreateRequiredRequest) (*CreateRequiredResponse, error)
	UpdateRequired(context.Context, *UpdateRequiredRequest) (*UpdateRequiredResponse, error)
	GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error)
	DeleteRequired(context.Context, *DeleteRequiredRequest) (*DeleteRequiredResponse, error)
	// Run by church admin
	AdminCreateRequired(context.Context, *AdminCreateRequiredRequest) (*AdminCreateRequiredResponse, error)
	AdminUpdateRequired(context.Context, *AdminUpdateRequiredRequest) (*AdminUpdateRequiredResponse, error)
	AdminGetRequireds(context.Context, *AdminGetRequiredsRequest) (*AdminGetRequiredsResponse, error)
	AdminDeleteRequired(context.Context, *AdminDeleteRequiredRequest) (*AdminDeleteRequiredResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateRequired(context.Context, *CreateRequiredRequest) (*CreateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequired not implemented")
}
func (UnimplementedGatewayServer) UpdateRequired(context.Context, *UpdateRequiredRequest) (*UpdateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRequired not implemented")
}
func (UnimplementedGatewayServer) GetRequireds(context.Context, *GetRequiredsRequest) (*GetRequiredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequireds not implemented")
}
func (UnimplementedGatewayServer) DeleteRequired(context.Context, *DeleteRequiredRequest) (*DeleteRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequired not implemented")
}
func (UnimplementedGatewayServer) AdminCreateRequired(context.Context, *AdminCreateRequiredRequest) (*AdminCreateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateRequired not implemented")
}
func (UnimplementedGatewayServer) AdminUpdateRequired(context.Context, *AdminUpdateRequiredRequest) (*AdminUpdateRequiredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateRequired not implemented")
}
func (UnimplementedGatewayServer) AdminGetRequireds(context.Context, *AdminGetRequiredsRequest) (*AdminGetRequiredsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetRequireds not implemented")
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

func _Gateway_CreateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/CreateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateRequired(ctx, req.(*CreateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/UpdateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateRequired(ctx, req.(*UpdateRequiredRequest))
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
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/GetRequireds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetRequireds(ctx, req.(*GetRequiredsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/DeleteRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteRequired(ctx, req.(*DeleteRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/AdminCreateRequired",
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
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/AdminUpdateRequired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateRequired(ctx, req.(*AdminUpdateRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetRequireds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetRequiredsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetRequireds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/AdminGetRequireds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetRequireds(ctx, req.(*AdminGetRequiredsRequest))
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
		FullMethod: "/good.gateway.app.good1.required1.v1.Gateway/AdminDeleteRequired",
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
	ServiceName: "good.gateway.app.good1.required1.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRequired",
			Handler:    _Gateway_CreateRequired_Handler,
		},
		{
			MethodName: "UpdateRequired",
			Handler:    _Gateway_UpdateRequired_Handler,
		},
		{
			MethodName: "GetRequireds",
			Handler:    _Gateway_GetRequireds_Handler,
		},
		{
			MethodName: "DeleteRequired",
			Handler:    _Gateway_DeleteRequired_Handler,
		},
		{
			MethodName: "AdminCreateRequired",
			Handler:    _Gateway_AdminCreateRequired_Handler,
		},
		{
			MethodName: "AdminUpdateRequired",
			Handler:    _Gateway_AdminUpdateRequired_Handler,
		},
		{
			MethodName: "AdminGetRequireds",
			Handler:    _Gateway_AdminGetRequireds_Handler,
		},
		{
			MethodName: "AdminDeleteRequired",
			Handler:    _Gateway_AdminDeleteRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/app/good/required/required.proto",
}
