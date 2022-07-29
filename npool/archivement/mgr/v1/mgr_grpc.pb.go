// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/archivement/mgr/v1/mgr.proto

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

// ArchivementManagerClient is the client API for ArchivementManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchivementManagerClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
}

type archivementManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewArchivementManagerClient(cc grpc.ClientConnInterface) ArchivementManagerClient {
	return &archivementManagerClient{cc}
}

func (c *archivementManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/archivement.manager.v1.ArchivementManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchivementManagerServer is the server API for ArchivementManager service.
// All implementations must embed UnimplementedArchivementManagerServer
// for forward compatibility
type ArchivementManagerServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	mustEmbedUnimplementedArchivementManagerServer()
}

// UnimplementedArchivementManagerServer must be embedded to have forward compatible implementations.
type UnimplementedArchivementManagerServer struct {
}

func (UnimplementedArchivementManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedArchivementManagerServer) mustEmbedUnimplementedArchivementManagerServer() {}

// UnsafeArchivementManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchivementManagerServer will
// result in compilation errors.
type UnsafeArchivementManagerServer interface {
	mustEmbedUnimplementedArchivementManagerServer()
}

func RegisterArchivementManagerServer(s grpc.ServiceRegistrar, srv ArchivementManagerServer) {
	s.RegisterService(&ArchivementManager_ServiceDesc, srv)
}

func _ArchivementManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivementManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivement.manager.v1.ArchivementManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivementManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ArchivementManager_ServiceDesc is the grpc.ServiceDesc for ArchivementManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchivementManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivement.manager.v1.ArchivementManager",
	HandlerType: (*ArchivementManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ArchivementManager_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/archivement/mgr/v1/mgr.proto",
}
