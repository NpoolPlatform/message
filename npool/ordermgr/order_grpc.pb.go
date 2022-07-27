// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ordermgr/order.proto

package ordermgr

import (
	context "context"
	npool "github.com/NpoolPlatform/message/npool"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderManagerClient is the client API for OrderManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagerClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
}

type orderManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagerClient(cc grpc.ClientConnInterface) OrderManagerClient {
	return &orderManagerClient{cc}
}

func (c *orderManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/order.manager.v1.OrderManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagerServer is the server API for OrderManager service.
// All implementations must embed UnimplementedOrderManagerServer
// for forward compatibility
type OrderManagerServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	mustEmbedUnimplementedOrderManagerServer()
}

// UnimplementedOrderManagerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagerServer struct {
}

func (UnimplementedOrderManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedOrderManagerServer) mustEmbedUnimplementedOrderManagerServer() {}

// UnsafeOrderManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagerServer will
// result in compilation errors.
type UnsafeOrderManagerServer interface {
	mustEmbedUnimplementedOrderManagerServer()
}

func RegisterOrderManagerServer(s grpc.ServiceRegistrar, srv OrderManagerServer) {
	s.RegisterService(&OrderManager_ServiceDesc, srv)
}

func _OrderManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.v1.OrderManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManager_ServiceDesc is the grpc.ServiceDesc for OrderManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.manager.v1.OrderManager",
	HandlerType: (*OrderManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _OrderManager_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ordermgr/order.proto",
}
