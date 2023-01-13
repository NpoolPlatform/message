// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/gw/v1/announcement/readstate/readstate.proto

package announcement

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
	CreateReadState(ctx context.Context, in *CreateReadStateRequest, opts ...grpc.CallOption) (*CreateReadStateResponse, error)
	GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateReadState(ctx context.Context, in *CreateReadStateRequest, opts ...grpc.CallOption) (*CreateReadStateResponse, error) {
	out := new(CreateReadStateResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/CreateReadState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error) {
	out := new(GetReadStatesResponse)
	err := c.cc.Invoke(ctx, "/notif.gateway.announcement.v1.Gateway/GetReadStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateReadState(context.Context, *CreateReadStateRequest) (*CreateReadStateResponse, error)
	GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateReadState(context.Context, *CreateReadStateRequest) (*CreateReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReadState not implemented")
}
func (UnimplementedGatewayServer) GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadStates not implemented")
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

func _Gateway_CreateReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/CreateReadState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateReadState(ctx, req.(*CreateReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetReadStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetReadStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.gateway.announcement.v1.Gateway/GetReadStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetReadStates(ctx, req.(*GetReadStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.announcement.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReadState",
			Handler:    _Gateway_CreateReadState_Handler,
		},
		{
			MethodName: "GetReadStates",
			Handler:    _Gateway_GetReadStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/gw/v1/announcement/readstate/readstate.proto",
}
