// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/order/gw/v1/powerrental/compensate/compensate.proto

package compensate

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
	// Admin apis
	AdminCreateCompensate(ctx context.Context, in *AdminCreateCompensateRequest, opts ...grpc.CallOption) (*AdminCreateCompensateResponse, error)
	AdminDeleteCompensate(ctx context.Context, in *AdminDeleteCompensateRequest, opts ...grpc.CallOption) (*AdminDeleteCompensateResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateCompensate(ctx context.Context, in *AdminCreateCompensateRequest, opts ...grpc.CallOption) (*AdminCreateCompensateResponse, error) {
	out := new(AdminCreateCompensateResponse)
	err := c.cc.Invoke(ctx, "/order.gateway.powerrental.compensate.v1.Gateway/AdminCreateCompensate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteCompensate(ctx context.Context, in *AdminDeleteCompensateRequest, opts ...grpc.CallOption) (*AdminDeleteCompensateResponse, error) {
	out := new(AdminDeleteCompensateResponse)
	err := c.cc.Invoke(ctx, "/order.gateway.powerrental.compensate.v1.Gateway/AdminDeleteCompensate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// Admin apis
	AdminCreateCompensate(context.Context, *AdminCreateCompensateRequest) (*AdminCreateCompensateResponse, error)
	AdminDeleteCompensate(context.Context, *AdminDeleteCompensateRequest) (*AdminDeleteCompensateResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) AdminCreateCompensate(context.Context, *AdminCreateCompensateRequest) (*AdminCreateCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateCompensate not implemented")
}
func (UnimplementedGatewayServer) AdminDeleteCompensate(context.Context, *AdminDeleteCompensateRequest) (*AdminDeleteCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteCompensate not implemented")
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

func _Gateway_AdminCreateCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.gateway.powerrental.compensate.v1.Gateway/AdminCreateCompensate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateCompensate(ctx, req.(*AdminCreateCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.gateway.powerrental.compensate.v1.Gateway/AdminDeleteCompensate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteCompensate(ctx, req.(*AdminDeleteCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.powerrental.compensate.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateCompensate",
			Handler:    _Gateway_AdminCreateCompensate_Handler,
		},
		{
			MethodName: "AdminDeleteCompensate",
			Handler:    _Gateway_AdminDeleteCompensate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/gw/v1/powerrental/compensate/compensate.proto",
}
