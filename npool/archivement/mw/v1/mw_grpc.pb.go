// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/archivement/mw/v1/mw.proto

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

// ArchivementMiddlewareClient is the client API for ArchivementMiddleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchivementMiddlewareClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
}

type archivementMiddlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewArchivementMiddlewareClient(cc grpc.ClientConnInterface) ArchivementMiddlewareClient {
	return &archivementMiddlewareClient{cc}
}

func (c *archivementMiddlewareClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/archivement.middleware.v1.ArchivementMiddleware/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchivementMiddlewareServer is the server API for ArchivementMiddleware service.
// All implementations must embed UnimplementedArchivementMiddlewareServer
// for forward compatibility
type ArchivementMiddlewareServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	mustEmbedUnimplementedArchivementMiddlewareServer()
}

// UnimplementedArchivementMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedArchivementMiddlewareServer struct {
}

func (UnimplementedArchivementMiddlewareServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedArchivementMiddlewareServer) mustEmbedUnimplementedArchivementMiddlewareServer() {}

// UnsafeArchivementMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchivementMiddlewareServer will
// result in compilation errors.
type UnsafeArchivementMiddlewareServer interface {
	mustEmbedUnimplementedArchivementMiddlewareServer()
}

func RegisterArchivementMiddlewareServer(s grpc.ServiceRegistrar, srv ArchivementMiddlewareServer) {
	s.RegisterService(&ArchivementMiddleware_ServiceDesc, srv)
}

func _ArchivementMiddleware_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivementMiddlewareServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivement.middleware.v1.ArchivementMiddleware/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivementMiddlewareServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ArchivementMiddleware_ServiceDesc is the grpc.ServiceDesc for ArchivementMiddleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchivementMiddleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivement.middleware.v1.ArchivementMiddleware",
	HandlerType: (*ArchivementMiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ArchivementMiddleware_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/archivement/mw/v1/mw.proto",
}
