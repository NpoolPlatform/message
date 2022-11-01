// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/chain/mgr/v1/coin/fee/fee.proto

package fee

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
	CreateFee(ctx context.Context, in *CreateFeeRequest, opts ...grpc.CallOption) (*CreateFeeResponse, error)
	CreateFees(ctx context.Context, in *CreateFeesRequest, opts ...grpc.CallOption) (*CreateFeesResponse, error)
	AddFee(ctx context.Context, in *AddFeeRequest, opts ...grpc.CallOption) (*AddFeeResponse, error)
	GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error)
	GetFeeOnly(ctx context.Context, in *GetFeeOnlyRequest, opts ...grpc.CallOption) (*GetFeeOnlyResponse, error)
	GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error)
	ExistFee(ctx context.Context, in *ExistFeeRequest, opts ...grpc.CallOption) (*ExistFeeResponse, error)
	ExistFeeConds(ctx context.Context, in *ExistFeeCondsRequest, opts ...grpc.CallOption) (*ExistFeeCondsResponse, error)
	CountFees(ctx context.Context, in *CountFeesRequest, opts ...grpc.CallOption) (*CountFeesResponse, error)
	DeleteFee(ctx context.Context, in *DeleteFeeRequest, opts ...grpc.CallOption) (*DeleteFeeResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateFee(ctx context.Context, in *CreateFeeRequest, opts ...grpc.CallOption) (*CreateFeeResponse, error) {
	out := new(CreateFeeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/CreateFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateFees(ctx context.Context, in *CreateFeesRequest, opts ...grpc.CallOption) (*CreateFeesResponse, error) {
	out := new(CreateFeesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/CreateFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) AddFee(ctx context.Context, in *AddFeeRequest, opts ...grpc.CallOption) (*AddFeeResponse, error) {
	out := new(AddFeeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/AddFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFee(ctx context.Context, in *GetFeeRequest, opts ...grpc.CallOption) (*GetFeeResponse, error) {
	out := new(GetFeeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/GetFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFeeOnly(ctx context.Context, in *GetFeeOnlyRequest, opts ...grpc.CallOption) (*GetFeeOnlyResponse, error) {
	out := new(GetFeeOnlyResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/GetFeeOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFees(ctx context.Context, in *GetFeesRequest, opts ...grpc.CallOption) (*GetFeesResponse, error) {
	out := new(GetFeesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/GetFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistFee(ctx context.Context, in *ExistFeeRequest, opts ...grpc.CallOption) (*ExistFeeResponse, error) {
	out := new(ExistFeeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/ExistFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistFeeConds(ctx context.Context, in *ExistFeeCondsRequest, opts ...grpc.CallOption) (*ExistFeeCondsResponse, error) {
	out := new(ExistFeeCondsResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/ExistFeeConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountFees(ctx context.Context, in *CountFeesRequest, opts ...grpc.CallOption) (*CountFeesResponse, error) {
	out := new(CountFeesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/CountFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteFee(ctx context.Context, in *DeleteFeeRequest, opts ...grpc.CallOption) (*DeleteFeeResponse, error) {
	out := new(DeleteFeeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.fee.v1.Manager/DeleteFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateFee(context.Context, *CreateFeeRequest) (*CreateFeeResponse, error)
	CreateFees(context.Context, *CreateFeesRequest) (*CreateFeesResponse, error)
	AddFee(context.Context, *AddFeeRequest) (*AddFeeResponse, error)
	GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error)
	GetFeeOnly(context.Context, *GetFeeOnlyRequest) (*GetFeeOnlyResponse, error)
	GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error)
	ExistFee(context.Context, *ExistFeeRequest) (*ExistFeeResponse, error)
	ExistFeeConds(context.Context, *ExistFeeCondsRequest) (*ExistFeeCondsResponse, error)
	CountFees(context.Context, *CountFeesRequest) (*CountFeesResponse, error)
	DeleteFee(context.Context, *DeleteFeeRequest) (*DeleteFeeResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateFee(context.Context, *CreateFeeRequest) (*CreateFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFee not implemented")
}
func (UnimplementedManagerServer) CreateFees(context.Context, *CreateFeesRequest) (*CreateFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFees not implemented")
}
func (UnimplementedManagerServer) AddFee(context.Context, *AddFeeRequest) (*AddFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFee not implemented")
}
func (UnimplementedManagerServer) GetFee(context.Context, *GetFeeRequest) (*GetFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedManagerServer) GetFeeOnly(context.Context, *GetFeeOnlyRequest) (*GetFeeOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeeOnly not implemented")
}
func (UnimplementedManagerServer) GetFees(context.Context, *GetFeesRequest) (*GetFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFees not implemented")
}
func (UnimplementedManagerServer) ExistFee(context.Context, *ExistFeeRequest) (*ExistFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFee not implemented")
}
func (UnimplementedManagerServer) ExistFeeConds(context.Context, *ExistFeeCondsRequest) (*ExistFeeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFeeConds not implemented")
}
func (UnimplementedManagerServer) CountFees(context.Context, *CountFeesRequest) (*CountFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFees not implemented")
}
func (UnimplementedManagerServer) DeleteFee(context.Context, *DeleteFeeRequest) (*DeleteFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFee not implemented")
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

func _Manager_CreateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/CreateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateFee(ctx, req.(*CreateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/CreateFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateFees(ctx, req.(*CreateFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_AddFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).AddFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/AddFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).AddFee(ctx, req.(*AddFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/GetFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFee(ctx, req.(*GetFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFeeOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeeOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFeeOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/GetFeeOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFeeOnly(ctx, req.(*GetFeeOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/GetFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFees(ctx, req.(*GetFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/ExistFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistFee(ctx, req.(*ExistFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistFeeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFeeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistFeeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/ExistFeeConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistFeeConds(ctx, req.(*ExistFeeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/CountFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountFees(ctx, req.(*CountFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.fee.v1.Manager/DeleteFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteFee(ctx, req.(*DeleteFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.manager.coin.fee.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFee",
			Handler:    _Manager_CreateFee_Handler,
		},
		{
			MethodName: "CreateFees",
			Handler:    _Manager_CreateFees_Handler,
		},
		{
			MethodName: "AddFee",
			Handler:    _Manager_AddFee_Handler,
		},
		{
			MethodName: "GetFee",
			Handler:    _Manager_GetFee_Handler,
		},
		{
			MethodName: "GetFeeOnly",
			Handler:    _Manager_GetFeeOnly_Handler,
		},
		{
			MethodName: "GetFees",
			Handler:    _Manager_GetFees_Handler,
		},
		{
			MethodName: "ExistFee",
			Handler:    _Manager_ExistFee_Handler,
		},
		{
			MethodName: "ExistFeeConds",
			Handler:    _Manager_ExistFeeConds_Handler,
		},
		{
			MethodName: "CountFees",
			Handler:    _Manager_CountFees_Handler,
		},
		{
			MethodName: "DeleteFee",
			Handler:    _Manager_DeleteFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mgr/v1/coin/fee/fee.proto",
}
