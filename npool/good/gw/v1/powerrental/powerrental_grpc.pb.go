// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/powerrental/powerrental.proto

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreatePowerRental(ctx context.Context, in *AdminCreatePowerRentalRequest, opts ...grpc.CallOption) (*AdminCreatePowerRentalResponse, error)
	AdminUpdatePowerRental(ctx context.Context, in *AdminUpdatePowerRentalRequest, opts ...grpc.CallOption) (*AdminUpdatePowerRentalResponse, error)
	GetPowerRental(ctx context.Context, in *GetPowerRentalRequest, opts ...grpc.CallOption) (*GetPowerRentalResponse, error)
	GetPowerRentals(ctx context.Context, in *GetPowerRentalsRequest, opts ...grpc.CallOption) (*GetPowerRentalsResponse, error)
	AdminDeletePowerRental(ctx context.Context, in *AdminDeletePowerRentalRequest, opts ...grpc.CallOption) (*AdminDeletePowerRentalResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreatePowerRental(ctx context.Context, in *AdminCreatePowerRentalRequest, opts ...grpc.CallOption) (*AdminCreatePowerRentalResponse, error) {
	out := new(AdminCreatePowerRentalResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.powerrental.v1.Gateway/AdminCreatePowerRental", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdatePowerRental(ctx context.Context, in *AdminUpdatePowerRentalRequest, opts ...grpc.CallOption) (*AdminUpdatePowerRentalResponse, error) {
	out := new(AdminUpdatePowerRentalResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.powerrental.v1.Gateway/AdminUpdatePowerRental", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetPowerRental(ctx context.Context, in *GetPowerRentalRequest, opts ...grpc.CallOption) (*GetPowerRentalResponse, error) {
	out := new(GetPowerRentalResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.powerrental.v1.Gateway/GetPowerRental", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetPowerRentals(ctx context.Context, in *GetPowerRentalsRequest, opts ...grpc.CallOption) (*GetPowerRentalsResponse, error) {
	out := new(GetPowerRentalsResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.powerrental.v1.Gateway/GetPowerRentals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeletePowerRental(ctx context.Context, in *AdminDeletePowerRentalRequest, opts ...grpc.CallOption) (*AdminDeletePowerRentalResponse, error) {
	out := new(AdminDeletePowerRentalResponse)
	err := c.cc.Invoke(ctx, "/good.gateway.powerrental.v1.Gateway/AdminDeletePowerRental", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreatePowerRental(context.Context, *AdminCreatePowerRentalRequest) (*AdminCreatePowerRentalResponse, error)
	AdminUpdatePowerRental(context.Context, *AdminUpdatePowerRentalRequest) (*AdminUpdatePowerRentalResponse, error)
	GetPowerRental(context.Context, *GetPowerRentalRequest) (*GetPowerRentalResponse, error)
	GetPowerRentals(context.Context, *GetPowerRentalsRequest) (*GetPowerRentalsResponse, error)
	AdminDeletePowerRental(context.Context, *AdminDeletePowerRentalRequest) (*AdminDeletePowerRentalResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreatePowerRental(context.Context, *AdminCreatePowerRentalRequest) (*AdminCreatePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreatePowerRental not implemented")
}
func (UnimplementedGatewayServer) AdminUpdatePowerRental(context.Context, *AdminUpdatePowerRentalRequest) (*AdminUpdatePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdatePowerRental not implemented")
}
func (UnimplementedGatewayServer) GetPowerRental(context.Context, *GetPowerRentalRequest) (*GetPowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRental not implemented")
}
func (UnimplementedGatewayServer) GetPowerRentals(context.Context, *GetPowerRentalsRequest) (*GetPowerRentalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerRentals not implemented")
}
func (UnimplementedGatewayServer) AdminDeletePowerRental(context.Context, *AdminDeletePowerRentalRequest) (*AdminDeletePowerRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeletePowerRental not implemented")
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

func _Gateway_AdminCreatePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreatePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreatePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.powerrental.v1.Gateway/AdminCreatePowerRental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreatePowerRental(ctx, req.(*AdminCreatePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdatePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdatePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdatePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.powerrental.v1.Gateway/AdminUpdatePowerRental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdatePowerRental(ctx, req.(*AdminUpdatePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetPowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetPowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.powerrental.v1.Gateway/GetPowerRental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetPowerRental(ctx, req.(*GetPowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetPowerRentals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRentalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetPowerRentals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.powerrental.v1.Gateway/GetPowerRentals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetPowerRentals(ctx, req.(*GetPowerRentalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeletePowerRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeletePowerRentalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeletePowerRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.gateway.powerrental.v1.Gateway/AdminDeletePowerRental",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeletePowerRental(ctx, req.(*AdminDeletePowerRentalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.powerrental.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreatePowerRental",
			Handler:    _Gateway_AdminCreatePowerRental_Handler,
		},
		{
			MethodName: "AdminUpdatePowerRental",
			Handler:    _Gateway_AdminUpdatePowerRental_Handler,
		},
		{
			MethodName: "GetPowerRental",
			Handler:    _Gateway_GetPowerRental_Handler,
		},
		{
			MethodName: "GetPowerRentals",
			Handler:    _Gateway_GetPowerRentals_Handler,
		},
		{
			MethodName: "AdminDeletePowerRental",
			Handler:    _Gateway_AdminDeletePowerRental_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/powerrental/powerrental.proto",
}
