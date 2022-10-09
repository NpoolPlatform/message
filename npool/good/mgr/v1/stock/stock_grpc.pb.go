// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mgr/v1/stock/stock.proto

package stock

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error)
	CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error)
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error)
	GetStockOnly(ctx context.Context, in *GetStockOnlyRequest, opts ...grpc.CallOption) (*GetStockOnlyResponse, error)
	GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error)
	ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error)
	ExistStockConds(ctx context.Context, in *ExistStockCondsRequest, opts ...grpc.CallOption) (*ExistStockCondsResponse, error)
	CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error) {
	out := new(CreateStockResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/CreateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error) {
	out := new(CreateStocksResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/CreateStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error) {
	out := new(GetStockResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/GetStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStockOnly(ctx context.Context, in *GetStockOnlyRequest, opts ...grpc.CallOption) (*GetStockOnlyResponse, error) {
	out := new(GetStockOnlyResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/GetStockOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error) {
	out := new(GetStocksResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/GetStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error) {
	out := new(ExistStockResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/ExistStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistStockConds(ctx context.Context, in *ExistStockCondsRequest, opts ...grpc.CallOption) (*ExistStockCondsResponse, error) {
	out := new(ExistStockCondsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/ExistStockConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error) {
	out := new(CountStocksResponse)
	err := c.cc.Invoke(ctx, "/good.manager.stock.v1.Manager/CountStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error)
	CreateStocks(context.Context, *CreateStocksRequest) (*CreateStocksResponse, error)
	GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error)
	GetStockOnly(context.Context, *GetStockOnlyRequest) (*GetStockOnlyResponse, error)
	GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error)
	ExistStock(context.Context, *ExistStockRequest) (*ExistStockResponse, error)
	ExistStockConds(context.Context, *ExistStockCondsRequest) (*ExistStockCondsResponse, error)
	CountStocks(context.Context, *CountStocksRequest) (*CountStocksResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStock not implemented")
}
func (UnimplementedManagerServer) CreateStocks(context.Context, *CreateStocksRequest) (*CreateStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStocks not implemented")
}
func (UnimplementedManagerServer) GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedManagerServer) GetStockOnly(context.Context, *GetStockOnlyRequest) (*GetStockOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockOnly not implemented")
}
func (UnimplementedManagerServer) GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStocks not implemented")
}
func (UnimplementedManagerServer) ExistStock(context.Context, *ExistStockRequest) (*ExistStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistStock not implemented")
}
func (UnimplementedManagerServer) ExistStockConds(context.Context, *ExistStockCondsRequest) (*ExistStockCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistStockConds not implemented")
}
func (UnimplementedManagerServer) CountStocks(context.Context, *CountStocksRequest) (*CountStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountStocks not implemented")
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

func _Manager_CreateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/CreateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateStock(ctx, req.(*CreateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/CreateStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateStocks(ctx, req.(*CreateStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStock(ctx, req.(*GetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetStockOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStockOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/GetStockOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStockOnly(ctx, req.(*GetStockOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/GetStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStocks(ctx, req.(*GetStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/ExistStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistStock(ctx, req.(*ExistStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistStockConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistStockCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistStockConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/ExistStockConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistStockConds(ctx, req.(*ExistStockCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.stock.v1.Manager/CountStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountStocks(ctx, req.(*CountStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.manager.stock.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStock",
			Handler:    _Manager_CreateStock_Handler,
		},
		{
			MethodName: "CreateStocks",
			Handler:    _Manager_CreateStocks_Handler,
		},
		{
			MethodName: "GetStock",
			Handler:    _Manager_GetStock_Handler,
		},
		{
			MethodName: "GetStockOnly",
			Handler:    _Manager_GetStockOnly_Handler,
		},
		{
			MethodName: "GetStocks",
			Handler:    _Manager_GetStocks_Handler,
		},
		{
			MethodName: "ExistStock",
			Handler:    _Manager_ExistStock_Handler,
		},
		{
			MethodName: "ExistStockConds",
			Handler:    _Manager_ExistStockConds_Handler,
		},
		{
			MethodName: "CountStocks",
			Handler:    _Manager_CountStocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mgr/v1/stock/stock.proto",
}
