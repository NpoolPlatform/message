// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sphinxproxy

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

// SphinxProxyClient is the client API for SphinxProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphinxProxyClient interface {
	// sync
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceReponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletReponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionReponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionReponse, error)
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

func (c *sphinxProxyClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceReponse, error) {
	out := new(GetBalanceReponse)
	err := c.cc.Invoke(ctx, "/sphinx.proxy.v1.SphinxProxy/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletReponse, error) {
	out := new(CreateWalletReponse)
	err := c.cc.Invoke(ctx, "/sphinx.proxy.v1.SphinxProxy/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionReponse, error) {
	out := new(CreateTransactionReponse)
	err := c.cc.Invoke(ctx, "/sphinx.proxy.v1.SphinxProxy/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionReponse, error) {
	out := new(GetTransactionReponse)
	err := c.cc.Invoke(ctx, "/sphinx.proxy.v1.SphinxProxy/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sphinxProxyClient) ProxyPlugin(ctx context.Context, opts ...grpc.CallOption) (SphinxProxy_ProxyPluginClient, error) {
	stream, err := c.cc.NewStream(ctx, &SphinxProxy_ServiceDesc.Streams[0], "/sphinx.proxy.v1.SphinxProxy/ProxyPlugin", opts...)
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
	stream, err := c.cc.NewStream(ctx, &SphinxProxy_ServiceDesc.Streams[1], "/sphinx.proxy.v1.SphinxProxy/ProxySign", opts...)
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
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceReponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletReponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionReponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionReponse, error)
	// async stream
	ProxyPlugin(SphinxProxy_ProxyPluginServer) error
	ProxySign(SphinxProxy_ProxySignServer) error
	mustEmbedUnimplementedSphinxProxyServer()
}

// UnimplementedSphinxProxyServer must be embedded to have forward compatible implementations.
type UnimplementedSphinxProxyServer struct {
}

func (UnimplementedSphinxProxyServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedSphinxProxyServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedSphinxProxyServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedSphinxProxyServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
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
		FullMethod: "/sphinx.proxy.v1.SphinxProxy/GetBalance",
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
		FullMethod: "/sphinx.proxy.v1.SphinxProxy/CreateWallet",
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
		FullMethod: "/sphinx.proxy.v1.SphinxProxy/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
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
		FullMethod: "/sphinx.proxy.v1.SphinxProxy/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SphinxProxyServer).GetTransaction(ctx, req.(*GetTransactionRequest))
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
			MethodName: "GetTransaction",
			Handler:    _SphinxProxy_GetTransaction_Handler,
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
