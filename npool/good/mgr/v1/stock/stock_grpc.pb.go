// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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

const (
	Manager_CreateStock_FullMethodName     = "/good.manager.stock.v1.Manager/CreateStock"
	Manager_CreateStocks_FullMethodName    = "/good.manager.stock.v1.Manager/CreateStocks"
	Manager_UpdateStock_FullMethodName     = "/good.manager.stock.v1.Manager/UpdateStock"
	Manager_AddStockFields_FullMethodName  = "/good.manager.stock.v1.Manager/AddStockFields"
	Manager_GetStock_FullMethodName        = "/good.manager.stock.v1.Manager/GetStock"
	Manager_GetStockOnly_FullMethodName    = "/good.manager.stock.v1.Manager/GetStockOnly"
	Manager_GetStocks_FullMethodName       = "/good.manager.stock.v1.Manager/GetStocks"
	Manager_ExistStock_FullMethodName      = "/good.manager.stock.v1.Manager/ExistStock"
	Manager_ExistStockConds_FullMethodName = "/good.manager.stock.v1.Manager/ExistStockConds"
	Manager_CountStocks_FullMethodName     = "/good.manager.stock.v1.Manager/CountStocks"
	Manager_DeleteStock_FullMethodName     = "/good.manager.stock.v1.Manager/DeleteStock"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error)
	CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
	AddStockFields(ctx context.Context, in *AddStockFieldsRequest, opts ...grpc.CallOption) (*AddStockFieldsResponse, error)
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error)
	GetStockOnly(ctx context.Context, in *GetStockOnlyRequest, opts ...grpc.CallOption) (*GetStockOnlyResponse, error)
	GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error)
	ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error)
	ExistStockConds(ctx context.Context, in *ExistStockCondsRequest, opts ...grpc.CallOption) (*ExistStockCondsResponse, error)
	CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error)
	DeleteStock(ctx context.Context, in *DeleteStockRequest, opts ...grpc.CallOption) (*DeleteStockResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error) {
	out := new(CreateStockResponse)
	err := c.cc.Invoke(ctx, Manager_CreateStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error) {
	out := new(CreateStocksResponse)
	err := c.cc.Invoke(ctx, Manager_CreateStocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) AddStockFields(ctx context.Context, in *AddStockFieldsRequest, opts ...grpc.CallOption) (*AddStockFieldsResponse, error) {
	out := new(AddStockFieldsResponse)
	err := c.cc.Invoke(ctx, Manager_AddStockFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error) {
	out := new(GetStockResponse)
	err := c.cc.Invoke(ctx, Manager_GetStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStockOnly(ctx context.Context, in *GetStockOnlyRequest, opts ...grpc.CallOption) (*GetStockOnlyResponse, error) {
	out := new(GetStockOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetStockOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error) {
	out := new(GetStocksResponse)
	err := c.cc.Invoke(ctx, Manager_GetStocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error) {
	out := new(ExistStockResponse)
	err := c.cc.Invoke(ctx, Manager_ExistStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistStockConds(ctx context.Context, in *ExistStockCondsRequest, opts ...grpc.CallOption) (*ExistStockCondsResponse, error) {
	out := new(ExistStockCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistStockConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error) {
	out := new(CountStocksResponse)
	err := c.cc.Invoke(ctx, Manager_CountStocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteStock(ctx context.Context, in *DeleteStockRequest, opts ...grpc.CallOption) (*DeleteStockResponse, error) {
	out := new(DeleteStockResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteStock_FullMethodName, in, out, opts...)
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
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	AddStockFields(context.Context, *AddStockFieldsRequest) (*AddStockFieldsResponse, error)
	GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error)
	GetStockOnly(context.Context, *GetStockOnlyRequest) (*GetStockOnlyResponse, error)
	GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error)
	ExistStock(context.Context, *ExistStockRequest) (*ExistStockResponse, error)
	ExistStockConds(context.Context, *ExistStockCondsRequest) (*ExistStockCondsResponse, error)
	CountStocks(context.Context, *CountStocksRequest) (*CountStocksResponse, error)
	DeleteStock(context.Context, *DeleteStockRequest) (*DeleteStockResponse, error)
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
func (UnimplementedManagerServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedManagerServer) AddStockFields(context.Context, *AddStockFieldsRequest) (*AddStockFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStockFields not implemented")
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
func (UnimplementedManagerServer) DeleteStock(context.Context, *DeleteStockRequest) (*DeleteStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStock not implemented")
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
		FullMethod: Manager_CreateStock_FullMethodName,
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
		FullMethod: Manager_CreateStocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateStocks(ctx, req.(*CreateStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_AddStockFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStockFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).AddStockFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_AddStockFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).AddStockFields(ctx, req.(*AddStockFieldsRequest))
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
		FullMethod: Manager_GetStock_FullMethodName,
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
		FullMethod: Manager_GetStockOnly_FullMethodName,
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
		FullMethod: Manager_GetStocks_FullMethodName,
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
		FullMethod: Manager_ExistStock_FullMethodName,
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
		FullMethod: Manager_ExistStockConds_FullMethodName,
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
		FullMethod: Manager_CountStocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountStocks(ctx, req.(*CountStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteStock(ctx, req.(*DeleteStockRequest))
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
			MethodName: "UpdateStock",
			Handler:    _Manager_UpdateStock_Handler,
		},
		{
			MethodName: "AddStockFields",
			Handler:    _Manager_AddStockFields_Handler,
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
		{
			MethodName: "DeleteStock",
			Handler:    _Manager_DeleteStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mgr/v1/stock/stock.proto",
}
