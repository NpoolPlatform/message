// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/ledger/mgr/v1/mgr.proto

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

// LedgerManagerClient is the client API for LedgerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LedgerManagerClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
}

type ledgerManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewLedgerManagerClient(cc grpc.ClientConnInterface) LedgerManagerClient {
	return &ledgerManagerClient{cc}
}

func (c *ledgerManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/ledger.manager.v1.LedgerManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedgerManagerServer is the server API for LedgerManager service.
// All implementations must embed UnimplementedLedgerManagerServer
// for forward compatibility
type LedgerManagerServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	mustEmbedUnimplementedLedgerManagerServer()
}

// UnimplementedLedgerManagerServer must be embedded to have forward compatible implementations.
type UnimplementedLedgerManagerServer struct {
}

func (UnimplementedLedgerManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedLedgerManagerServer) mustEmbedUnimplementedLedgerManagerServer() {}

// UnsafeLedgerManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LedgerManagerServer will
// result in compilation errors.
type UnsafeLedgerManagerServer interface {
	mustEmbedUnimplementedLedgerManagerServer()
}

func RegisterLedgerManagerServer(s grpc.ServiceRegistrar, srv LedgerManagerServer) {
	s.RegisterService(&LedgerManager_ServiceDesc, srv)
}

func _LedgerManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedgerManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ledger.manager.v1.LedgerManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedgerManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// LedgerManager_ServiceDesc is the grpc.ServiceDesc for LedgerManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LedgerManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ledger.manager.v1.LedgerManager",
	HandlerType: (*LedgerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _LedgerManager_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/ledger/mgr/v1/mgr.proto",
}
