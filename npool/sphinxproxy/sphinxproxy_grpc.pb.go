// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/sphinxproxy/sphinxproxy.proto

package sphinxproxy

import (
	context "context"
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
	SphinxProxy_Version_FullMethodName           = "/sphinx.proxy.v1.SphinxProxy/Version"
	SphinxProxy_GetBalance_FullMethodName        = "/sphinx.proxy.v1.SphinxProxy/GetBalance"
	SphinxProxy_CreateWallet_FullMethodName      = "/sphinx.proxy.v1.SphinxProxy/CreateWallet"
	SphinxProxy_CreateTransaction_FullMethodName = "/sphinx.proxy.v1.SphinxProxy/CreateTransaction"
	SphinxProxy_UpdateTransaction_FullMethodName = "/sphinx.proxy.v1.SphinxProxy/UpdateTransaction"
	SphinxProxy_GetTransaction_FullMethodName    = "/sphinx.proxy.v1.SphinxProxy/GetTransaction"
	SphinxProxy_GetTransactions_FullMethodName   = "/sphinx.proxy.v1.SphinxProxy/GetTransactions"
	SphinxProxy_GetEstimateGas_FullMethodName    = "/sphinx.proxy.v1.SphinxProxy/GetEstimateGas"
	SphinxProxy_ProxyPlugin_FullMethodName       = "/sphinx.proxy.v1.SphinxProxy/ProxyPlugin"
	SphinxProxy_ProxySign_FullMethodName         = "/sphinx.proxy.v1.SphinxProxy/ProxySign"
)

// SphinxProxyClient is the client API for SphinxProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphinxProxyClient interface {
	// sync
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	GetEstimateGas(ctx context.Context, in *GetEstimateGasRequest, opts ...grpc.CallOption) (*GetEstimateGasResponse, error)
	// async stream
	ProxyPlugin(ctx context.Context, opts ...grpc.CallOption) (SphinxProxy_ProxyPluginClient, error)
	ProxySign(ctx context.Context, opts ...grpc.CallOption) (SphinxProxy_ProxySignClient, error)
}

type sphinxProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewSphinxProxyClient(cc grpc.ClientConnInterface) SphinxProxyClient {
	return &sphinxProxyClient{cc}
}

func (c *sphinxProxyClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error) {
	out := new(UpdateTransactionResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_UpdateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_GetTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_GetTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) GetEstimateGas(ctx context.Context, in *GetEstimateGasRequest, opts ...grpc.CallOption) (*GetEstimateGasResponse, error) {
	out := new(GetEstimateGasResponse)
	err := c.cc.Invoke(ctx, SphinxProxy_GetEstimateGas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) ProxyPlugin(ctx context.Context, opts ...grpc.CallOption) (SphinxProxy_ProxyPluginClient, error) {
	stream, err := c.cc.NewStream(ctx, &SphinxProxy_ServiceDesc.Streams[0], SphinxProxy_ProxyPlugin_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sphinxProxyProxyPluginClient{stream}
	return x, nil
}

type SphinxProxy_ProxyPluginClient interface {
	Send(*ProxyPluginResponse) error
	Recv() (*ProxyPluginRequest, error)
	grpc.ClientStream
}

type sphinxProxyProxyPluginClient struct {
	grpc.ClientStream
}

func (x *sphinxProxyProxyPluginClient) Send(m *ProxyPluginResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sphinxProxyProxyPluginClient) Recv() (*ProxyPluginRequest, error) {
	m := new(ProxyPluginRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sphinxProxyClient) ProxySign(ctx context.Context, opts ...grpc.CallOption) (SphinxProxy_ProxySignClient, error) {
	stream, err := c.cc.NewStream(ctx, &SphinxProxy_ServiceDesc.Streams[1], SphinxProxy_ProxySign_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sphinxProxyProxySignClient{stream}
	return x, nil
}

type SphinxProxy_ProxySignClient interface {
	Send(*ProxySignResponse) error
	Recv() (*ProxySignRequest, error)
	grpc.ClientStream
}

type sphinxProxyProxySignClient struct {
	grpc.ClientStream
}

func (x *sphinxProxyProxySignClient) Send(m *ProxySignResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sphinxProxyProxySignClient) Recv() (*ProxySignRequest, error) {
	m := new(ProxySignRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SphinxProxyServer is the server API for SphinxProxy service.
// All implementations must embed UnimplementedSphinxProxyServer
// for forward compatibility
type SphinxProxyServer interface {
	// sync
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	GetEstimateGas(context.Context, *GetEstimateGasRequest) (*GetEstimateGasResponse, error)
	// async stream
	ProxyPlugin(SphinxProxy_ProxyPluginServer) error
	ProxySign(SphinxProxy_ProxySignServer) error
	mustEmbedUnimplementedSphinxProxyServer()
}

// UnimplementedSphinxProxyServer must be embedded to have forward compatible implementations.
type UnimplementedSphinxProxyServer struct {
}

func (UnimplementedSphinxProxyServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedSphinxProxyServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedSphinxProxyServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedSphinxProxyServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedSphinxProxyServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedSphinxProxyServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedSphinxProxyServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedSphinxProxyServer) GetEstimateGas(context.Context, *GetEstimateGasRequest) (*GetEstimateGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstimateGas not implemented")
}
func (UnimplementedSphinxProxyServer) ProxyPlugin(SphinxProxy_ProxyPluginServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxyPlugin not implemented")
}
func (UnimplementedSphinxProxyServer) ProxySign(SphinxProxy_ProxySignServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxySign not implemented")
}
func (UnimplementedSphinxProxyServer) mustEmbedUnimplementedSphinxProxyServer() {}

// UnsafeSphinxProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SphinxProxyServer will
// result in compilation errors.
type UnsafeSphinxProxyServer interface {
	mustEmbedUnimplementedSphinxProxyServer()
}

func RegisterSphinxProxyServer(s grpc.ServiceRegistrar, srv SphinxProxyServer) {
	s.RegisterService(&SphinxProxy_ServiceDesc, srv)
}

func _SphinxProxy_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_UpdateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_GetEstimateGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstimateGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SphinxProxyServer).GetEstimateGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SphinxProxy_GetEstimateGas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).GetEstimateGas(ctx, req.(*GetEstimateGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SphinxProxy_ProxyPlugin_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SphinxProxyServer).ProxyPlugin(&sphinxProxyProxyPluginServer{stream})
}

type SphinxProxy_ProxyPluginServer interface {
	Send(*ProxyPluginRequest) error
	Recv() (*ProxyPluginResponse, error)
	grpc.ServerStream
}

type sphinxProxyProxyPluginServer struct {
	grpc.ServerStream
}

func (x *sphinxProxyProxyPluginServer) Send(m *ProxyPluginRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sphinxProxyProxyPluginServer) Recv() (*ProxyPluginResponse, error) {
	m := new(ProxyPluginResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SphinxProxy_ProxySign_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SphinxProxyServer).ProxySign(&sphinxProxyProxySignServer{stream})
}

type SphinxProxy_ProxySignServer interface {
	Send(*ProxySignRequest) error
	Recv() (*ProxySignResponse, error)
	grpc.ServerStream
}

type sphinxProxyProxySignServer struct {
	grpc.ServerStream
}

func (x *sphinxProxyProxySignServer) Send(m *ProxySignRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sphinxProxyProxySignServer) Recv() (*ProxySignResponse, error) {
	m := new(ProxySignResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SphinxProxy_ServiceDesc is the grpc.ServiceDesc for SphinxProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SphinxProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sphinx.proxy.v1.SphinxProxy",
	HandlerType: (*SphinxProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _SphinxProxy_Version_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _SphinxProxy_GetBalance_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _SphinxProxy_CreateWallet_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _SphinxProxy_CreateTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _SphinxProxy_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _SphinxProxy_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _SphinxProxy_GetTransactions_Handler,
		},
		{
			MethodName: "GetEstimateGas",
			Handler:    _SphinxProxy_GetEstimateGas_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProxyPlugin",
			Handler:       _SphinxProxy_ProxyPlugin_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ProxySign",
			Handler:       _SphinxProxy_ProxySign_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "npool/sphinxproxy/sphinxproxy.proto",
}
