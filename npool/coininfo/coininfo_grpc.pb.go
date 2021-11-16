// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

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

// SphinxCoininfoClient is the client API for SphinxCoininfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphinxCoininfoClient interface {
	// 注册新币种
	RegisterCoin(ctx context.Context, in *RegisterCoinRequest, opts ...grpc.CallOption) (*RegisterCoinResponse, error)
	// 获取币种信息
	GetCoinInfos(ctx context.Context, in *GetCoinInfosRequest, opts ...grpc.CallOption) (*GetCoinInfosResponse, error)
	// 获取单个币种
	GetCoinInfo(ctx context.Context, in *GetCoinInfoRequest, opts ...grpc.CallOption) (*CoinInfoRow, error)
}

type sphinxCoininfoClient struct {
	cc grpc.ClientConnInterface
}

func NewSphinxCoininfoClient(cc grpc.ClientConnInterface) SphinxCoininfoClient {
	return &sphinxCoininfoClient{cc}
}

func (c *sphinxCoininfoClient) RegisterCoin(ctx context.Context, in *RegisterCoinRequest, opts ...grpc.CallOption) (*RegisterCoinResponse, error) {
	out := new(RegisterCoinResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoininfo/RegisterCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxCoininfoClient) GetCoinInfos(ctx context.Context, in *GetCoinInfosRequest, opts ...grpc.CallOption) (*GetCoinInfosResponse, error) {
	out := new(GetCoinInfosResponse)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoininfo/GetCoinInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxCoininfoClient) GetCoinInfo(ctx context.Context, in *GetCoinInfoRequest, opts ...grpc.CallOption) (*CoinInfoRow, error) {
	out := new(CoinInfoRow)
	err := c.cc.Invoke(ctx, "/sphinx.coininfo.v1.SphinxCoininfo/GetCoinInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphinxCoininfoServer is the server API for SphinxCoininfo service.
// All implementations must embed UnimplementedSphinxCoininfoServer
// for forward compatibility
type SphinxCoininfoServer interface {
	// 注册新币种
	RegisterCoin(context.Context, *RegisterCoinRequest) (*RegisterCoinResponse, error)
	// 获取币种信息
	GetCoinInfos(context.Context, *GetCoinInfosRequest) (*GetCoinInfosResponse, error)
	// 获取单个币种
	GetCoinInfo(context.Context, *GetCoinInfoRequest) (*CoinInfoRow, error)
	mustEmbedUnimplementedSphinxCoininfoServer()
}

// UnimplementedSphinxCoininfoServer must be embedded to have forward compatible implementations.
type UnimplementedSphinxCoininfoServer struct {
}

func (UnimplementedSphinxCoininfoServer) RegisterCoin(context.Context, *RegisterCoinRequest) (*RegisterCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCoin not implemented")
}
func (UnimplementedSphinxCoininfoServer) GetCoinInfos(context.Context, *GetCoinInfosRequest) (*GetCoinInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinInfos not implemented")
}
func (UnimplementedSphinxCoininfoServer) GetCoinInfo(context.Context, *GetCoinInfoRequest) (*CoinInfoRow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinInfo not implemented")
}
func (UnimplementedSphinxCoininfoServer) mustEmbedUnimplementedSphinxCoininfoServer() {}

// UnsafeSphinxCoininfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SphinxCoininfoServer will
// result in compilation errors.
type UnsafeSphinxCoininfoServer interface {
	mustEmbedUnimplementedSphinxCoininfoServer()
}

func RegisterSphinxCoininfoServer(s grpc.ServiceRegistrar, srv SphinxCoininfoServer) {
	s.RegisterService(&SphinxCoininfo_ServiceDesc, srv)
}

func _SphinxCoininfo_RegisterCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoininfoServer).RegisterCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoininfo/RegisterCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoininfoServer).RegisterCoin(ctx, req.(*RegisterCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxCoininfo_GetCoinInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoininfoServer).GetCoinInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoininfo/GetCoinInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoininfoServer).GetCoinInfos(ctx, req.(*GetCoinInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxCoininfo_GetCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxCoininfoServer).GetCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sphinx.coininfo.v1.SphinxCoininfo/GetCoinInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxCoininfoServer).GetCoinInfo(ctx, req.(*GetCoinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SphinxCoininfo_ServiceDesc is the grpc.ServiceDesc for SphinxCoininfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SphinxCoininfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.coininfo.v1.SphinxCoininfo",
	HandlerType: (*SphinxCoininfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCoin",
			Handler:    _SphinxCoininfo_RegisterCoin_Handler,
		},
		{
			MethodName: "GetCoinInfos",
			Handler:    _SphinxCoininfo_GetCoinInfos_Handler,
		},
		{
			MethodName: "GetCoinInfo",
			Handler:    _SphinxCoininfo_GetCoinInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/coininfo/coininfo.proto",
}
