// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ledger/gw/v1/simulate/ledger/statement/statement.proto

package statement

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
	GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error)
	GetAppStatements(ctx context.Context, in *GetAppStatementsRequest, opts ...grpc.CallOption) (*GetAppStatementsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error) {
	out := new(GetStatementsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.simulate.ledger.statement.v1.Gateway/GetStatements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppStatements(ctx context.Context, in *GetAppStatementsRequest, opts ...grpc.CallOption) (*GetAppStatementsResponse, error) {
	out := new(GetAppStatementsResponse)
	err := c.cc.Invoke(ctx, "/ledger.gateway.simulate.ledger.statement.v1.Gateway/GetAppStatements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetStatements(context.Context, *GetStatementsRequest) (*GetStatementsResponse, error)
	GetAppStatements(context.Context, *GetAppStatementsRequest) (*GetAppStatementsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetStatements(context.Context, *GetStatementsRequest) (*GetStatementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatements not implemented")
}
func (UnimplementedGatewayServer) GetAppStatements(context.Context, *GetAppStatementsRequest) (*GetAppStatementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppStatements not implemented")
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

func _Gateway_GetStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.simulate.ledger.statement.v1.Gateway/GetStatements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetStatements(ctx, req.(*GetStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.gateway.simulate.ledger.statement.v1.Gateway/GetAppStatements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppStatements(ctx, req.(*GetAppStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.gateway.simulate.ledger.statement.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatements",
			Handler:    _Gateway_GetStatements_Handler,
		},
		{
			MethodName: "GetAppStatements",
			Handler:    _Gateway_GetAppStatements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/gw/v1/simulate/ledger/statement/statement.proto",
}
