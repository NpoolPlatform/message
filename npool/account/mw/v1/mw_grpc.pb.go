// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/account/mw/v1/mw.proto

package v1

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

const (
	Middleware_Version_FullMethodName = "/account.middleware.v1.Middleware/Version"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, Middleware_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
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

func _Middleware_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.middleware.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Middleware_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/account/mw/v1/mw.proto",
}
