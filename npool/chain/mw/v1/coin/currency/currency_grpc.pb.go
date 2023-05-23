// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/chain/mw/v1/coin/currency/currency.proto

package currency

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

const (
<<<<<<< HEAD
	Middleware_CreateCurrency_FullMethodName    = "/chain.middleware.coin.currency.v1.Middleware/CreateCurrency"
	Middleware_CreateCurrencies_FullMethodName  = "/chain.middleware.coin.currency.v1.Middleware/CreateCurrencies"
	Middleware_RefreshCurrencies_FullMethodName = "/chain.middleware.coin.currency.v1.Middleware/RefreshCurrencies"
	Middleware_GetCurrency_FullMethodName       = "/chain.middleware.coin.currency.v1.Middleware/GetCurrency"
	Middleware_GetCoinCurrency_FullMethodName   = "/chain.middleware.coin.currency.v1.Middleware/GetCoinCurrency"
	Middleware_GetCurrencies_FullMethodName     = "/chain.middleware.coin.currency.v1.Middleware/GetCurrencies"
	Middleware_GetHistories_FullMethodName      = "/chain.middleware.coin.currency.v1.Middleware/GetHistories"
)

// MiddlewareClient is the client API for Middleware service.
=======
	Manager_CreateCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/CreateCurrency"
	Manager_UpdateCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/UpdateCurrency"
	Manager_GetCurrency_FullMethodName        = "/chain.middleware.coin.currency.v1.Manager/GetCurrency"
	Manager_GetCurrencyOnly_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/GetCurrencyOnly"
	Manager_GetCoinCurrency_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/GetCoinCurrency"
	Manager_GetCurrencies_FullMethodName      = "/chain.middleware.coin.currency.v1.Manager/GetCurrencies"
	Manager_ExistCurrency_FullMethodName      = "/chain.middleware.coin.currency.v1.Manager/ExistCurrency"
	Manager_ExistCurrencyConds_FullMethodName = "/chain.middleware.coin.currency.v1.Manager/ExistCurrencyConds"
	Manager_CountCurrencies_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/CountCurrencies"
	Manager_RefreshCurrencies_FullMethodName  = "/chain.middleware.coin.currency.v1.Manager/RefreshCurrencies"
	Manager_DeleteCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/DeleteCurrency"
)

// ManagerClient is the client API for Manager service.
>>>>>>> Add currency history
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error)
	UpdateCurrency(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*UpdateCurrencyResponse, error)
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
	GetCurrencyOnly(ctx context.Context, in *GetCurrencyOnlyRequest, opts ...grpc.CallOption) (*GetCurrencyOnlyResponse, error)
	GetCoinCurrency(ctx context.Context, in *GetCoinCurrencyRequest, opts ...grpc.CallOption) (*GetCoinCurrencyResponse, error)
	GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
	ExistCurrency(ctx context.Context, in *ExistCurrencyRequest, opts ...grpc.CallOption) (*ExistCurrencyResponse, error)
	ExistCurrencyConds(ctx context.Context, in *ExistCurrencyCondsRequest, opts ...grpc.CallOption) (*ExistCurrencyCondsResponse, error)
	CountCurrencies(ctx context.Context, in *CountCurrenciesRequest, opts ...grpc.CallOption) (*CountCurrenciesResponse, error)
	RefreshCurrencies(ctx context.Context, in *RefreshCurrenciesRequest, opts ...grpc.CallOption) (*RefreshCurrenciesResponse, error)
	DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error) {
	out := new(CreateCurrencyResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_CreateCurrency_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_CreateCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) CreateCurrencies(ctx context.Context, in *CreateCurrenciesRequest, opts ...grpc.CallOption) (*CreateCurrenciesResponse, error) {
	out := new(CreateCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCurrencies_FullMethodName, in, out, opts...)
=======
func (c *managerClient) UpdateCurrency(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*UpdateCurrencyResponse, error) {
	out := new(UpdateCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) RefreshCurrencies(ctx context.Context, in *RefreshCurrenciesRequest, opts ...grpc.CallOption) (*RefreshCurrenciesResponse, error) {
	out := new(RefreshCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleware_RefreshCurrencies_FullMethodName, in, out, opts...)
=======
func (c *managerClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCurrency_FullMethodName, in, out, opts...)
=======
func (c *managerClient) GetCurrencyOnly(ctx context.Context, in *GetCurrencyOnlyRequest, opts ...grpc.CallOption) (*GetCurrencyOnlyResponse, error) {
	out := new(GetCurrencyOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCurrencyOnly_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCoinCurrency(ctx context.Context, in *GetCoinCurrencyRequest, opts ...grpc.CallOption) (*GetCoinCurrencyResponse, error) {
	out := new(GetCoinCurrencyResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_GetCoinCurrency_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_GetCoinCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_GetCurrencies_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_GetCurrencies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCurrency(ctx context.Context, in *ExistCurrencyRequest, opts ...grpc.CallOption) (*ExistCurrencyResponse, error) {
	out := new(ExistCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_ExistCurrency_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCurrencyConds(ctx context.Context, in *ExistCurrencyCondsRequest, opts ...grpc.CallOption) (*ExistCurrencyCondsResponse, error) {
	out := new(ExistCurrencyCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistCurrencyConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountCurrencies(ctx context.Context, in *CountCurrenciesRequest, opts ...grpc.CallOption) (*CountCurrenciesResponse, error) {
	out := new(CountCurrenciesResponse)
	err := c.cc.Invoke(ctx, Manager_CountCurrencies_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error) {
	out := new(GetHistoriesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetHistories_FullMethodName, in, out, opts...)
=======
func (c *managerClient) RefreshCurrencies(ctx context.Context, in *RefreshCurrenciesRequest, opts ...grpc.CallOption) (*RefreshCurrenciesResponse, error) {
	out := new(RefreshCurrenciesResponse)
	err := c.cc.Invoke(ctx, Manager_RefreshCurrencies_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error) {
	out := new(DeleteCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteCurrency_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateCurrency(context.Context, *CreateCurrencyRequest) (*CreateCurrencyResponse, error)
	UpdateCurrency(context.Context, *UpdateCurrencyRequest) (*UpdateCurrencyResponse, error)
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	GetCurrencyOnly(context.Context, *GetCurrencyOnlyRequest) (*GetCurrencyOnlyResponse, error)
	GetCoinCurrency(context.Context, *GetCoinCurrencyRequest) (*GetCoinCurrencyResponse, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	ExistCurrency(context.Context, *ExistCurrencyRequest) (*ExistCurrencyResponse, error)
	ExistCurrencyConds(context.Context, *ExistCurrencyCondsRequest) (*ExistCurrencyCondsResponse, error)
	CountCurrencies(context.Context, *CountCurrenciesRequest) (*CountCurrenciesResponse, error)
	RefreshCurrencies(context.Context, *RefreshCurrenciesRequest) (*RefreshCurrenciesResponse, error)
	DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateCurrency(context.Context, *CreateCurrencyRequest) (*CreateCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurrency not implemented")
}
func (UnimplementedManagerServer) UpdateCurrency(context.Context, *UpdateCurrencyRequest) (*UpdateCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrency not implemented")
}
func (UnimplementedManagerServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (UnimplementedManagerServer) GetCurrencyOnly(context.Context, *GetCurrencyOnlyRequest) (*GetCurrencyOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencyOnly not implemented")
}
func (UnimplementedManagerServer) GetCoinCurrency(context.Context, *GetCoinCurrencyRequest) (*GetCoinCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinCurrency not implemented")
}
func (UnimplementedManagerServer) GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}
func (UnimplementedManagerServer) ExistCurrency(context.Context, *ExistCurrencyRequest) (*ExistCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCurrency not implemented")
}
func (UnimplementedManagerServer) ExistCurrencyConds(context.Context, *ExistCurrencyCondsRequest) (*ExistCurrencyCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCurrencyConds not implemented")
}
func (UnimplementedManagerServer) CountCurrencies(context.Context, *CountCurrenciesRequest) (*CountCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCurrencies not implemented")
}
func (UnimplementedManagerServer) RefreshCurrencies(context.Context, *RefreshCurrenciesRequest) (*RefreshCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCurrencies not implemented")
}
func (UnimplementedManagerServer) DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrency not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_CreateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_CreateCurrency_FullMethodName,
=======
		FullMethod: Manager_CreateCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCurrency(ctx, req.(*CreateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_CreateCurrencies_FullMethodName,
=======
		FullMethod: Manager_UpdateCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateCurrency(ctx, req.(*UpdateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_RefreshCurrencies_FullMethodName,
=======
		FullMethod: Manager_GetCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCurrencyOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCurrencyOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_GetCurrency_FullMethodName,
=======
		FullMethod: Manager_GetCurrencyOnly_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCurrencyOnly(ctx, req.(*GetCurrencyOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCoinCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCoinCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_GetCoinCurrency_FullMethodName,
=======
		FullMethod: Manager_GetCoinCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCoinCurrency(ctx, req.(*GetCoinCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCurrencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCurrencies(ctx, req.(*GetCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCurrency(ctx, req.(*ExistCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCurrencyConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCurrencyCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCurrencyConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_GetCurrencies_FullMethodName,
=======
		FullMethod: Manager_ExistCurrencyConds_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCurrencyConds(ctx, req.(*ExistCurrencyCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_GetHistories_FullMethodName,
=======
		FullMethod: Manager_CountCurrencies_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountCurrencies(ctx, req.(*CountCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_RefreshCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).RefreshCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_RefreshCurrencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).RefreshCurrencies(ctx, req.(*RefreshCurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteCurrency(ctx, req.(*DeleteCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.middleware.coin.currency.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCurrency",
			Handler:    _Manager_CreateCurrency_Handler,
		},
		{
			MethodName: "UpdateCurrency",
			Handler:    _Manager_UpdateCurrency_Handler,
		},
		{
			MethodName: "GetCurrency",
			Handler:    _Manager_GetCurrency_Handler,
		},
		{
			MethodName: "GetCurrencyOnly",
			Handler:    _Manager_GetCurrencyOnly_Handler,
		},
		{
			MethodName: "GetCoinCurrency",
			Handler:    _Manager_GetCoinCurrency_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _Manager_GetCurrencies_Handler,
		},
		{
			MethodName: "ExistCurrency",
			Handler:    _Manager_ExistCurrency_Handler,
		},
		{
			MethodName: "ExistCurrencyConds",
			Handler:    _Manager_ExistCurrencyConds_Handler,
		},
		{
			MethodName: "CountCurrencies",
			Handler:    _Manager_CountCurrencies_Handler,
		},
		{
			MethodName: "RefreshCurrencies",
			Handler:    _Manager_RefreshCurrencies_Handler,
		},
		{
			MethodName: "DeleteCurrency",
			Handler:    _Manager_DeleteCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mw/v1/coin/currency/currency.proto",
}
