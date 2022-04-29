// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/stockmgr/stockmgr.proto

package stockmgr

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

// StockManagerClient is the client API for StockManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockManagerClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error)
	CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
	UpdateStockFields(ctx context.Context, in *UpdateStockFieldsRequest, opts ...grpc.CallOption) (*UpdateStockFieldsResponse, error)
	AtomicUpdateStockFields(ctx context.Context, in *AtomicUpdateStockFieldsRequest, opts ...grpc.CallOption) (*AtomicUpdateStockFieldsResponse, error)
	ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error)
	CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error)
	DeleteStock(ctx context.Context, in *DeleteStockRequest, opts ...grpc.CallOption) (*DeleteStockResponse, error)
}

type stockManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewStockManagerClient(cc grpc.ClientConnInterface) StockManagerClient {
	return &stockManagerClient{cc}
}

func (c *stockManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) CreateStock(ctx context.Context, in *CreateStockRequest, opts ...grpc.CallOption) (*CreateStockResponse, error) {
	out := new(CreateStockResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/CreateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) CreateStocks(ctx context.Context, in *CreateStocksRequest, opts ...grpc.CallOption) (*CreateStocksResponse, error) {
	out := new(CreateStocksResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/CreateStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/UpdateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) UpdateStockFields(ctx context.Context, in *UpdateStockFieldsRequest, opts ...grpc.CallOption) (*UpdateStockFieldsResponse, error) {
	out := new(UpdateStockFieldsResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/UpdateStockFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) AtomicUpdateStockFields(ctx context.Context, in *AtomicUpdateStockFieldsRequest, opts ...grpc.CallOption) (*AtomicUpdateStockFieldsResponse, error) {
	out := new(AtomicUpdateStockFieldsResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/AtomicUpdateStockFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) ExistStock(ctx context.Context, in *ExistStockRequest, opts ...grpc.CallOption) (*ExistStockResponse, error) {
	out := new(ExistStockResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/ExistStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) CountStocks(ctx context.Context, in *CountStocksRequest, opts ...grpc.CallOption) (*CountStocksResponse, error) {
	out := new(CountStocksResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/CountStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockManagerClient) DeleteStock(ctx context.Context, in *DeleteStockRequest, opts ...grpc.CallOption) (*DeleteStockResponse, error) {
	out := new(DeleteStockResponse)
	err := c.cc.Invoke(ctx, "/stock.manager.v1.StockManager/DeleteStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockManagerServer is the server API for StockManager service.
// All implementations must embed UnimplementedStockManagerServer
// for forward compatibility
type StockManagerServer interface {
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error)
	CreateStocks(context.Context, *CreateStocksRequest) (*CreateStocksResponse, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	UpdateStockFields(context.Context, *UpdateStockFieldsRequest) (*UpdateStockFieldsResponse, error)
	AtomicUpdateStockFields(context.Context, *AtomicUpdateStockFieldsRequest) (*AtomicUpdateStockFieldsResponse, error)
	ExistStock(context.Context, *ExistStockRequest) (*ExistStockResponse, error)
	CountStocks(context.Context, *CountStocksRequest) (*CountStocksResponse, error)
	DeleteStock(context.Context, *DeleteStockRequest) (*DeleteStockResponse, error)
	mustEmbedUnimplementedStockManagerServer()
}

// UnimplementedStockManagerServer must be embedded to have forward compatible implementations.
type UnimplementedStockManagerServer struct {
}

func (UnimplementedStockManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedStockManagerServer) CreateStock(context.Context, *CreateStockRequest) (*CreateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStock not implemented")
}
func (UnimplementedStockManagerServer) CreateStocks(context.Context, *CreateStocksRequest) (*CreateStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStocks not implemented")
}
func (UnimplementedStockManagerServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedStockManagerServer) UpdateStockFields(context.Context, *UpdateStockFieldsRequest) (*UpdateStockFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStockFields not implemented")
}
func (UnimplementedStockManagerServer) AtomicUpdateStockFields(context.Context, *AtomicUpdateStockFieldsRequest) (*AtomicUpdateStockFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicUpdateStockFields not implemented")
}
func (UnimplementedStockManagerServer) ExistStock(context.Context, *ExistStockRequest) (*ExistStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistStock not implemented")
}
func (UnimplementedStockManagerServer) CountStocks(context.Context, *CountStocksRequest) (*CountStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountStocks not implemented")
}
func (UnimplementedStockManagerServer) DeleteStock(context.Context, *DeleteStockRequest) (*DeleteStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStock not implemented")
}
func (UnimplementedStockManagerServer) mustEmbedUnimplementedStockManagerServer() {}

// UnsafeStockManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockManagerServer will
// result in compilation errors.
type UnsafeStockManagerServer interface {
	mustEmbedUnimplementedStockManagerServer()
}

func RegisterStockManagerServer(s grpc.ServiceRegistrar, srv StockManagerServer) {
	s.RegisterService(&StockManager_ServiceDesc, srv)
}

func _StockManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_CreateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).CreateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/CreateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).CreateStock(ctx, req.(*CreateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_CreateStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).CreateStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/CreateStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).CreateStocks(ctx, req.(*CreateStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/UpdateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_UpdateStockFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).UpdateStockFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/UpdateStockFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).UpdateStockFields(ctx, req.(*UpdateStockFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_AtomicUpdateStockFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicUpdateStockFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).AtomicUpdateStockFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/AtomicUpdateStockFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).AtomicUpdateStockFields(ctx, req.(*AtomicUpdateStockFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_ExistStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).ExistStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/ExistStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).ExistStock(ctx, req.(*ExistStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_CountStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).CountStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/CountStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).CountStocks(ctx, req.(*CountStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockManager_DeleteStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockManagerServer).DeleteStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.manager.v1.StockManager/DeleteStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockManagerServer).DeleteStock(ctx, req.(*DeleteStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockManager_ServiceDesc is the grpc.ServiceDesc for StockManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock.manager.v1.StockManager",
	HandlerType: (*StockManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _StockManager_Version_Handler,
		},
		{
			MethodName: "CreateStock",
			Handler:    _StockManager_CreateStock_Handler,
		},
		{
			MethodName: "CreateStocks",
			Handler:    _StockManager_CreateStocks_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _StockManager_UpdateStock_Handler,
		},
		{
			MethodName: "UpdateStockFields",
			Handler:    _StockManager_UpdateStockFields_Handler,
		},
		{
			MethodName: "AtomicUpdateStockFields",
			Handler:    _StockManager_AtomicUpdateStockFields_Handler,
		},
		{
			MethodName: "ExistStock",
			Handler:    _StockManager_ExistStock_Handler,
		},
		{
			MethodName: "CountStocks",
			Handler:    _StockManager_CountStocks_Handler,
		},
		{
			MethodName: "DeleteStock",
			Handler:    _StockManager_DeleteStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/stockmgr/stockmgr.proto",
}
