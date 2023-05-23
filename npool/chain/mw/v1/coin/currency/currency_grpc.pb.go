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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
>>>>>>> Add currency history
	Manager_CreateCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/CreateCurrency"
	Manager_UpdateCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/UpdateCurrency"
	Manager_GetCurrency_FullMethodName        = "/chain.middleware.coin.currency.v1.Manager/GetCurrency"
	Manager_GetCurrencyOnly_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/GetCurrencyOnly"
<<<<<<< HEAD
<<<<<<< HEAD
=======
	Manager_GetCoinCurrency_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/GetCoinCurrency"
>>>>>>> Add currency history
=======
>>>>>>> Remove unused api
	Manager_GetCurrencies_FullMethodName      = "/chain.middleware.coin.currency.v1.Manager/GetCurrencies"
	Manager_ExistCurrency_FullMethodName      = "/chain.middleware.coin.currency.v1.Manager/ExistCurrency"
	Manager_ExistCurrencyConds_FullMethodName = "/chain.middleware.coin.currency.v1.Manager/ExistCurrencyConds"
	Manager_CountCurrencies_FullMethodName    = "/chain.middleware.coin.currency.v1.Manager/CountCurrencies"
<<<<<<< HEAD
<<<<<<< HEAD
=======
	Manager_RefreshCurrencies_FullMethodName  = "/chain.middleware.coin.currency.v1.Manager/RefreshCurrencies"
>>>>>>> Add currency history
=======
>>>>>>> Remove unused api
	Manager_DeleteCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Manager/DeleteCurrency"
)

// ManagerClient is the client API for Manager service.
<<<<<<< HEAD
>>>>>>> Add currency history
=======
	Middleawre_GetCurrency_FullMethodName     = "/chain.middleware.coin.currency.v1.Middleawre/GetCurrency"
	Middleawre_GetCurrencyOnly_FullMethodName = "/chain.middleware.coin.currency.v1.Middleawre/GetCurrencyOnly"
	Middleawre_GetCurrencies_FullMethodName   = "/chain.middleware.coin.currency.v1.Middleawre/GetCurrencies"
=======
	Middleawre_GetCurrency_FullMethodName   = "/chain.middleware.coin.currency.v1.Middleawre/GetCurrency"
	Middleawre_GetCurrencies_FullMethodName = "/chain.middleware.coin.currency.v1.Middleawre/GetCurrencies"
>>>>>>> Remove only
)

// MiddleawreClient is the client API for Middleawre service.
>>>>>>> Remove unused currency apis
=======
	Middleware_GetCurrency_FullMethodName   = "/chain.middleware.coin.currency.v1.Middleware/GetCurrency"
	Middleware_GetCurrencies_FullMethodName = "/chain.middleware.coin.currency.v1.Middleware/GetCurrencies"
)

// MiddlewareClient is the client API for Middleware service.
>>>>>>> Correct name
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
	GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
=======
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error)
	UpdateCurrency(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*UpdateCurrencyResponse, error)
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
	GetCurrencyOnly(ctx context.Context, in *GetCurrencyOnlyRequest, opts ...grpc.CallOption) (*GetCurrencyOnlyResponse, error)
	GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
	ExistCurrency(ctx context.Context, in *ExistCurrencyRequest, opts ...grpc.CallOption) (*ExistCurrencyResponse, error)
	ExistCurrencyConds(ctx context.Context, in *ExistCurrencyCondsRequest, opts ...grpc.CallOption) (*ExistCurrencyCondsResponse, error)
	CountCurrencies(ctx context.Context, in *CountCurrenciesRequest, opts ...grpc.CallOption) (*CountCurrenciesResponse, error)
	DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error)
>>>>>>> Add currency history
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *managerClient) CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error) {
	out := new(CreateCurrencyResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_CreateCurrency_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_CreateCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
=======
func (c *managerClient) CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error) {
	out := new(CreateCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_CreateCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
func (c *middlewareClient) CreateCurrencies(ctx context.Context, in *CreateCurrenciesRequest, opts ...grpc.CallOption) (*CreateCurrenciesResponse, error) {
	out := new(CreateCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateCurrencies_FullMethodName, in, out, opts...)
=======
func (c *managerClient) UpdateCurrency(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*UpdateCurrencyResponse, error) {
	out := new(UpdateCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
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
<<<<<<< HEAD
func (c *middlewareClient) RefreshCurrencies(ctx context.Context, in *RefreshCurrenciesRequest, opts ...grpc.CallOption) (*RefreshCurrenciesResponse, error) {
	out := new(RefreshCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleware_RefreshCurrencies_FullMethodName, in, out, opts...)
=======
func (c *managerClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
=======
func (c *middleawreClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Middleawre_GetCurrency_FullMethodName, in, out, opts...)
>>>>>>> Remove unused currency apis
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *middlewareClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCurrency_FullMethodName, in, out, opts...)
=======
func (c *managerClient) GetCurrencyOnly(ctx context.Context, in *GetCurrencyOnlyRequest, opts ...grpc.CallOption) (*GetCurrencyOnlyResponse, error) {
	out := new(GetCurrencyOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCurrencyOnly_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *managerClient) GetCoinCurrency(ctx context.Context, in *GetCoinCurrencyRequest, opts ...grpc.CallOption) (*GetCoinCurrencyResponse, error) {
	out := new(GetCoinCurrencyResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_GetCoinCurrency_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_GetCoinCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
=======
func (c *middleawreClient) GetCurrencyOnly(ctx context.Context, in *GetCurrencyOnlyRequest, opts ...grpc.CallOption) (*GetCurrencyOnlyResponse, error) {
	out := new(GetCurrencyOnlyResponse)
	err := c.cc.Invoke(ctx, Middleawre_GetCurrencyOnly_FullMethodName, in, out, opts...)
>>>>>>> Remove unused currency apis
=======
func (c *managerClient) GetCoinCurrency(ctx context.Context, in *GetCoinCurrencyRequest, opts ...grpc.CallOption) (*GetCoinCurrencyResponse, error) {
	out := new(GetCoinCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCoinCurrency_FullMethodName, in, out, opts...)
>>>>>>> Add currency history
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Remove unused api
=======
>>>>>>> Remove unused api
func (c *managerClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
<<<<<<< HEAD
	err := c.cc.Invoke(ctx, Middleware_GetCurrencies_FullMethodName, in, out, opts...)
=======
	err := c.cc.Invoke(ctx, Manager_GetCurrencies_FullMethodName, in, out, opts...)
=======
func (c *managerClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
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
=======
func (c *middlewareClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCurrency_FullMethodName, in, out, opts...)
>>>>>>> Correct name
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
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
<<<<<<< HEAD
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

=======
>>>>>>> Remove unused api
func (c *managerClient) DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error) {
	out := new(DeleteCurrencyResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteCurrency_FullMethodName, in, out, opts...)
=======
=======
>>>>>>> Remove only
func (c *middleawreClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleawre_GetCurrencies_FullMethodName, in, out, opts...)
>>>>>>> Remove unused currency apis
=======
func (c *middlewareClient) GetCurrencies(ctx context.Context, in *GetCurrenciesRequest, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetCurrencies_FullMethodName, in, out, opts...)
>>>>>>> Correct name
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

=======
>>>>>>> Remove unused api
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
<<<<<<< HEAD
type MiddlewareServer interface {
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
=======
type ManagerServer interface {
	CreateCurrency(context.Context, *CreateCurrencyRequest) (*CreateCurrencyResponse, error)
	UpdateCurrency(context.Context, *UpdateCurrencyRequest) (*UpdateCurrencyResponse, error)
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	GetCurrencyOnly(context.Context, *GetCurrencyOnlyRequest) (*GetCurrencyOnlyResponse, error)
	GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error)
	ExistCurrency(context.Context, *ExistCurrencyRequest) (*ExistCurrencyResponse, error)
	ExistCurrencyConds(context.Context, *ExistCurrencyCondsRequest) (*ExistCurrencyCondsResponse, error)
	CountCurrencies(context.Context, *CountCurrenciesRequest) (*CountCurrenciesResponse, error)
	DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error)
	mustEmbedUnimplementedManagerServer()
>>>>>>> Add currency history
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

<<<<<<< HEAD
func (UnimplementedMiddlewareServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (UnimplementedMiddlewareServer) GetCurrencies(context.Context, *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}
=======
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
func (UnimplementedManagerServer) DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrency not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}
>>>>>>> Add currency history

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add currency history
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
<<<<<<< HEAD
		FullMethod: Middleware_CreateCurrency_FullMethodName,
=======
		FullMethod: Manager_CreateCurrency_FullMethodName,
>>>>>>> Add currency history
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
<<<<<<< HEAD
		FullMethod: Middleware_CreateCurrencies_FullMethodName,
=======
		FullMethod: Manager_UpdateCurrency_FullMethodName,
>>>>>>> Add currency history
=======
		FullMethod: Manager_UpdateCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateCurrency(ctx, req.(*UpdateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
=======
func RegisterMiddleawreServer(s grpc.ServiceRegistrar, srv MiddleawreServer) {
	s.RegisterService(&Middleawre_ServiceDesc, srv)
>>>>>>> Remove unused currency apis
=======
func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
>>>>>>> Correct name
}

<<<<<<< HEAD
func _Middleware_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
=======
func _Manager_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
>>>>>>> Add currency history
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		FullMethod: Middleware_RefreshCurrencies_FullMethodName,
=======
		FullMethod: Manager_GetCurrency_FullMethodName,
>>>>>>> Add currency history
=======
		FullMethod: Middleawre_GetCurrency_FullMethodName,
>>>>>>> Remove unused currency apis
=======
		FullMethod: Middleware_GetCurrency_FullMethodName,
>>>>>>> Correct name
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
=======
		return srv.(ManagerServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCurrency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
>>>>>>> Add currency history
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func _Middleawre_GetCurrencyOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
=======
func _Manager_GetCurrencyOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
>>>>>>> Add currency history
	in := new(GetCurrencyOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddleawreServer).GetCurrencyOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
<<<<<<< HEAD
		FullMethod: Middleware_GetCurrency_FullMethodName,
=======
		FullMethod: Manager_GetCurrencyOnly_FullMethodName,
>>>>>>> Add currency history
=======
		FullMethod: Middleawre_GetCurrencyOnly_FullMethodName,
>>>>>>> Remove unused currency apis
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleawreServer).GetCurrencyOnly(ctx, req.(*GetCurrencyOnlyRequest))
=======
		return srv.(ManagerServer).GetCurrencyOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCurrencyOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCurrencyOnly(ctx, req.(*GetCurrencyOnlyRequest))
>>>>>>> Add currency history
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add currency history
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
<<<<<<< HEAD
		FullMethod: Middleware_GetCoinCurrency_FullMethodName,
=======
		FullMethod: Manager_GetCoinCurrency_FullMethodName,
>>>>>>> Add currency history
=======
		FullMethod: Manager_GetCoinCurrency_FullMethodName,
>>>>>>> Add currency history
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCoinCurrency(ctx, req.(*GetCoinCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
=======
>>>>>>> Remove unused api
func _Manager_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
=======
=======
>>>>>>> Remove only
func _Middleawre_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
>>>>>>> Remove unused currency apis
=======
func _Middleware_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
>>>>>>> Correct name
=======
=======
>>>>>>> Remove unused api
func _Manager_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
>>>>>>> Add currency history
	in := new(GetCurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
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
		FullMethod: Manager_ExistCurrencyConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
>>>>>>> Add currency history
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
<<<<<<< HEAD
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: Middleware_GetHistories_FullMethodName,
=======
		FullMethod: Manager_CountCurrencies_FullMethodName,
>>>>>>> Add currency history
=======
		return srv.(MiddleawreServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleawre_GetCurrencies_FullMethodName,
>>>>>>> Remove unused currency apis
=======
		return srv.(MiddlewareServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetCurrencies_FullMethodName,
>>>>>>> Correct name
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetCurrencies(ctx, req.(*GetCurrenciesRequest))
=======
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CountCurrencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountCurrencies(ctx, req.(*CountCurrenciesRequest))
>>>>>>> Add currency history
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
<<<<<<< HEAD
			MethodName: "GetCurrency",
			Handler:    _Middleware_GetCurrency_Handler,
		},
		{
=======
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
<<<<<<< HEAD
			MethodName: "GetCoinCurrency",
			Handler:    _Manager_GetCoinCurrency_Handler,
		},
		{
>>>>>>> Add currency history
=======
>>>>>>> Remove unused api
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
<<<<<<< HEAD
			MethodName: "RefreshCurrencies",
			Handler:    _Manager_RefreshCurrencies_Handler,
		},
<<<<<<< HEAD
=======
		{
=======
>>>>>>> Remove unused api
			MethodName: "DeleteCurrency",
			Handler:    _Manager_DeleteCurrency_Handler,
		},
>>>>>>> Add currency history
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mw/v1/coin/currency/currency.proto",
}
