// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coininfo

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

// SphinxCoinInfoClient is the client API for SphinxCoinInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphinxCoinInfoClient interface {
	CreateCoinInfo(ctx context.Context, in *CreateCoinInfoRequest, opts ...grpc.CallOption) (*CreateCoinInfoResponse, error)
	GetCoinInfo(ctx context.Context, in *GetCoinInfoRequest, opts ...grpc.CallOption) (*GetCoinInfoResponse, error)
	GetCoinInfos(ctx context.Context, in *GetCoinInfosRequest, opts ...grpc.CallOption) (*GetCoinInfosResponse, error)
	UpdateCoinInfo(ctx context.Context, in *UpdateCoinInfoRequest, opts ...grpc.CallOption) (*UpdateCoinInfoResponse, error)
}

type sphinxCoinInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewSphinxCoinInfoClient(cc grpc.ClientConnInterface) SphinxCoinInfoClient {
	return &sphinxCoinInfoClient{cc}
}

func (c *sphinxCoinInfoClient) CreateCoinInfo(ctx context.Context, in *CreateCoinInfoRequest, opts ...grpc.CallOption) (*CreateCoinInfoResponse, error) {
	out := new(CreateCoinInfoResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoinInfo/CreateCoinInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxCoinInfoClient) GetCoinInfo(ctx context.Context, in *GetCoinInfoRequest, opts ...grpc.CallOption) (*GetCoinInfoResponse, error) {
	out := new(GetCoinInfoResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoinInfo/GetCoinInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxCoinInfoClient) GetCoinInfos(ctx context.Context, in *GetCoinInfosRequest, opts ...grpc.CallOption) (*GetCoinInfosResponse, error) {
	out := new(GetCoinInfosResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoinInfo/GetCoinInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxCoinInfoClient) UpdateCoinInfo(ctx context.Context, in *UpdateCoinInfoRequest, opts ...grpc.CallOption) (*UpdateCoinInfoResponse, error) {
	out := new(UpdateCoinInfoResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoinInfo/UpdateCoinInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphinxCoinInfoServer is the server API for SphinxCoinInfo service.
// All implementations must embed UnimplementedSphinxCoinInfoServer
// for forward compatibility
type SphinxCoinInfoServer interface {
	CreateCoinInfo(context.Context, *CreateCoinInfoRequest) (*CreateCoinInfoResponse, error)
	GetCoinInfo(context.Context, *GetCoinInfoRequest) (*GetCoinInfoResponse, error)
	GetCoinInfos(context.Context, *GetCoinInfosRequest) (*GetCoinInfosResponse, error)
	UpdateCoinInfo(context.Context, *UpdateCoinInfoRequest) (*UpdateCoinInfoResponse, error)
	mustEmbedUnimplementedSphinxCoinInfoServer()
}

// UnimplementedSphinxCoinInfoServer must be embedded to have forward compatible implementations.
type UnimplementedSphinxCoinInfoServer struct {
}

func (UnimplementedSphinxCoinInfoServer) CreateCoinInfo(context.Context, *CreateCoinInfoRequest) (*CreateCoinInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinInfo not implemented")
}
func (UnimplementedSphinxCoinInfoServer) GetCoinInfo(context.Context, *GetCoinInfoRequest) (*GetCoinInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinInfo not implemented")
}
func (UnimplementedSphinxCoinInfoServer) GetCoinInfos(context.Context, *GetCoinInfosRequest) (*GetCoinInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinInfos not implemented")
}
func (UnimplementedSphinxCoinInfoServer) UpdateCoinInfo(context.Context, *UpdateCoinInfoRequest) (*UpdateCoinInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinInfo not implemented")
}
func (UnimplementedSphinxCoinInfoServer) mustEmbedUnimplementedSphinxCoinInfoServer() {}

// UnsafeSphinxCoinInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SphinxCoinInfoServer will
// result in compilation errors.
type UnsafeSphinxCoinInfoServer interface {
	mustEmbedUnimplementedSphinxCoinInfoServer()
}

func RegisterSphinxCoinInfoServer(s grpc.ServiceRegistrar, srv SphinxCoinInfoServer) {
	s.RegisterService(&SphinxCoinInfo_ServiceDesc, srv)
}

func _SphinxCoinInfo_CreateCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoinInfoServer).CreateCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoinInfo/CreateCoinInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoinInfoServer).CreateCoinInfo(ctx, req.(*CreateCoinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxCoinInfo_GetCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoinInfoServer).GetCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoinInfo/GetCoinInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoinInfoServer).GetCoinInfo(ctx, req.(*GetCoinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxCoinInfo_GetCoinInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoinInfoServer).GetCoinInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoinInfo/GetCoinInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoinInfoServer).GetCoinInfos(ctx, req.(*GetCoinInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxCoinInfo_UpdateCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoinInfoServer).UpdateCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoinInfo/UpdateCoinInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoinInfoServer).UpdateCoinInfo(ctx, req.(*UpdateCoinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SphinxCoinInfo_ServiceDesc is the grpc.ServiceDesc for SphinxCoinInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SphinxCoinInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.coininfo.v1.SphinxCoinInfo",
	HandlerType: (*SphinxCoinInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinInfo",
			Handler:    _SphinxCoinInfo_CreateCoinInfo_Handler,
		},
		{
			MethodName: "GetCoinInfo",
			Handler:    _SphinxCoinInfo_GetCoinInfo_Handler,
		},
		{
			MethodName: "GetCoinInfos",
			Handler:    _SphinxCoinInfo_GetCoinInfos_Handler,
		},
		{
			MethodName: "UpdateCoinInfo",
			Handler:    _SphinxCoinInfo_UpdateCoinInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/coininfo/coininfo.proto",
}
