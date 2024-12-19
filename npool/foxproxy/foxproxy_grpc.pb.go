// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/foxproxy/foxproxy.proto

package foxproxy

import (
	context "context"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FoxProxyStreamClient is the client API for FoxProxyStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoxProxyStreamClient interface {
	// async stream
	DEStream(ctx context.Context, opts ...grpc.CallOption) (FoxProxyStream_DEStreamClient, error)
}

type foxProxyStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewFoxProxyStreamClient(cc grpc.ClientConnInterface) FoxProxyStreamClient {
	return &foxProxyStreamClient{cc}
}

func (c *foxProxyStreamClient) DEStream(ctx context.Context, opts ...grpc.CallOption) (FoxProxyStream_DEStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FoxProxyStream_ServiceDesc.Streams[0], "/fox.proxy.v1.FoxProxyStream/DEStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &foxProxyStreamDEStreamClient{stream}
	return x, nil
}

type FoxProxyStream_DEStreamClient interface {
	Send(*DataElement) error
	Recv() (*DataElement, error)
	grpc.ClientStream
}

type foxProxyStreamDEStreamClient struct {
	grpc.ClientStream
}

func (x *foxProxyStreamDEStreamClient) Send(m *DataElement) error {
	return x.ClientStream.SendMsg(m)
}

func (x *foxProxyStreamDEStreamClient) Recv() (*DataElement, error) {
	m := new(DataElement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoxProxyStreamServer is the server API for FoxProxyStream service.
// All implementations must embed UnimplementedFoxProxyStreamServer
// for forward compatibility
type FoxProxyStreamServer interface {
	// async stream
	DEStream(FoxProxyStream_DEStreamServer) error
	mustEmbedUnimplementedFoxProxyStreamServer()
}

// UnimplementedFoxProxyStreamServer must be embedded to have forward compatible implementations.
type UnimplementedFoxProxyStreamServer struct {
}

func (UnimplementedFoxProxyStreamServer) DEStream(FoxProxyStream_DEStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DEStream not implemented")
}
func (UnimplementedFoxProxyStreamServer) mustEmbedUnimplementedFoxProxyStreamServer() {}

// UnsafeFoxProxyStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoxProxyStreamServer will
// result in compilation errors.
type UnsafeFoxProxyStreamServer interface {
	mustEmbedUnimplementedFoxProxyStreamServer()
}

func RegisterFoxProxyStreamServer(s grpc.ServiceRegistrar, srv FoxProxyStreamServer) {
	s.RegisterService(&FoxProxyStream_ServiceDesc, srv)
}

func _FoxProxyStream_DEStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FoxProxyStreamServer).DEStream(&foxProxyStreamDEStreamServer{stream})
}

type FoxProxyStream_DEStreamServer interface {
	Send(*DataElement) error
	Recv() (*DataElement, error)
	grpc.ServerStream
}

type foxProxyStreamDEStreamServer struct {
	grpc.ServerStream
}

func (x *foxProxyStreamDEStreamServer) Send(m *DataElement) error {
	return x.ServerStream.SendMsg(m)
}

func (x *foxProxyStreamDEStreamServer) Recv() (*DataElement, error) {
	m := new(DataElement)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoxProxyStream_ServiceDesc is the grpc.ServiceDesc for FoxProxyStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoxProxyStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fox.proxy.v1.FoxProxyStream",
	HandlerType: (*FoxProxyStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DEStream",
			Handler:       _FoxProxyStream_DEStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "npool/foxproxy/foxproxy.proto",
}

// FoxProxyClient is the client API for FoxProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoxProxyClient interface {
	// sync
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.VersionResponse, error)
	RegisterCoin(ctx context.Context, in *RegisterCoinRequest, opts ...grpc.CallOption) (*RegisterCoinResponse, error)
	GetClientInfos(ctx context.Context, in *GetClientInfosRequest, opts ...grpc.CallOption) (*GetClientInfosResponse, error)
}

type foxProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewFoxProxyClient(cc grpc.ClientConnInterface) FoxProxyClient {
	return &foxProxyClient{cc}
}

func (c *foxProxyClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.VersionResponse, error) {
	out := new(v1.VersionResponse)
	err := c.cc.Invoke(ctx, "/fox.proxy.v1.FoxProxy/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foxProxyClient) RegisterCoin(ctx context.Context, in *RegisterCoinRequest, opts ...grpc.CallOption) (*RegisterCoinResponse, error) {
	out := new(RegisterCoinResponse)
	err := c.cc.Invoke(ctx, "/fox.proxy.v1.FoxProxy/RegisterCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foxProxyClient) GetClientInfos(ctx context.Context, in *GetClientInfosRequest, opts ...grpc.CallOption) (*GetClientInfosResponse, error) {
	out := new(GetClientInfosResponse)
	err := c.cc.Invoke(ctx, "/fox.proxy.v1.FoxProxy/GetClientInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoxProxyServer is the server API for FoxProxy service.
// All implementations must embed UnimplementedFoxProxyServer
// for forward compatibility
type FoxProxyServer interface {
	// sync
	Version(context.Context, *emptypb.Empty) (*v1.VersionResponse, error)
	RegisterCoin(context.Context, *RegisterCoinRequest) (*RegisterCoinResponse, error)
	GetClientInfos(context.Context, *GetClientInfosRequest) (*GetClientInfosResponse, error)
	mustEmbedUnimplementedFoxProxyServer()
}

// UnimplementedFoxProxyServer must be embedded to have forward compatible implementations.
type UnimplementedFoxProxyServer struct {
}

func (UnimplementedFoxProxyServer) Version(context.Context, *emptypb.Empty) (*v1.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedFoxProxyServer) RegisterCoin(context.Context, *RegisterCoinRequest) (*RegisterCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCoin not implemented")
}
func (UnimplementedFoxProxyServer) GetClientInfos(context.Context, *GetClientInfosRequest) (*GetClientInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientInfos not implemented")
}
func (UnimplementedFoxProxyServer) mustEmbedUnimplementedFoxProxyServer() {}

// UnsafeFoxProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoxProxyServer will
// result in compilation errors.
type UnsafeFoxProxyServer interface {
	mustEmbedUnimplementedFoxProxyServer()
}

func RegisterFoxProxyServer(s grpc.ServiceRegistrar, srv FoxProxyServer) {
	s.RegisterService(&FoxProxy_ServiceDesc, srv)
}

func _FoxProxy_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoxProxyServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fox.proxy.v1.FoxProxy/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoxProxyServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoxProxy_RegisterCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoxProxyServer).RegisterCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fox.proxy.v1.FoxProxy/RegisterCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoxProxyServer).RegisterCoin(ctx, req.(*RegisterCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoxProxy_GetClientInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoxProxyServer).GetClientInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fox.proxy.v1.FoxProxy/GetClientInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoxProxyServer).GetClientInfos(ctx, req.(*GetClientInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FoxProxy_ServiceDesc is the grpc.ServiceDesc for FoxProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoxProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fox.proxy.v1.FoxProxy",
	HandlerType: (*FoxProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _FoxProxy_Version_Handler,
		},
		{
			MethodName: "RegisterCoin",
			Handler:    _FoxProxy_RegisterCoin_Handler,
		},
		{
			MethodName: "GetClientInfos",
			Handler:    _FoxProxy_GetClientInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/foxproxy/foxproxy.proto",
}
