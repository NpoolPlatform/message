// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/chain/mgr/v1/coin/base/base.proto

package base

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
	CreateCoinBase(ctx context.Context, in *CreateCoinBaseRequest, opts ...grpc.CallOption) (*CreateCoinBaseResponse, error)
	CreateCoinBases(ctx context.Context, in *CreateCoinBasesRequest, opts ...grpc.CallOption) (*CreateCoinBasesResponse, error)
	AddCoinBase(ctx context.Context, in *AddCoinBaseRequest, opts ...grpc.CallOption) (*AddCoinBaseResponse, error)
	GetCoinBase(ctx context.Context, in *GetCoinBaseRequest, opts ...grpc.CallOption) (*GetCoinBaseResponse, error)
	GetCoinBaseOnly(ctx context.Context, in *GetCoinBaseOnlyRequest, opts ...grpc.CallOption) (*GetCoinBaseOnlyResponse, error)
	GetCoinBases(ctx context.Context, in *GetCoinBasesRequest, opts ...grpc.CallOption) (*GetCoinBasesResponse, error)
	ExistCoinBase(ctx context.Context, in *ExistCoinBaseRequest, opts ...grpc.CallOption) (*ExistCoinBaseResponse, error)
	ExistCoinBaseConds(ctx context.Context, in *ExistCoinBaseCondsRequest, opts ...grpc.CallOption) (*ExistCoinBaseCondsResponse, error)
	CountCoinBases(ctx context.Context, in *CountCoinBasesRequest, opts ...grpc.CallOption) (*CountCoinBasesResponse, error)
	DeleteCoinBase(ctx context.Context, in *DeleteCoinBaseRequest, opts ...grpc.CallOption) (*DeleteCoinBaseResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateCoinBase(ctx context.Context, in *CreateCoinBaseRequest, opts ...grpc.CallOption) (*CreateCoinBaseResponse, error) {
	out := new(CreateCoinBaseResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/CreateCoinBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateCoinBases(ctx context.Context, in *CreateCoinBasesRequest, opts ...grpc.CallOption) (*CreateCoinBasesResponse, error) {
	out := new(CreateCoinBasesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/CreateCoinBases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) AddCoinBase(ctx context.Context, in *AddCoinBaseRequest, opts ...grpc.CallOption) (*AddCoinBaseResponse, error) {
	out := new(AddCoinBaseResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/AddCoinBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCoinBase(ctx context.Context, in *GetCoinBaseRequest, opts ...grpc.CallOption) (*GetCoinBaseResponse, error) {
	out := new(GetCoinBaseResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/GetCoinBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCoinBaseOnly(ctx context.Context, in *GetCoinBaseOnlyRequest, opts ...grpc.CallOption) (*GetCoinBaseOnlyResponse, error) {
	out := new(GetCoinBaseOnlyResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/GetCoinBaseOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCoinBases(ctx context.Context, in *GetCoinBasesRequest, opts ...grpc.CallOption) (*GetCoinBasesResponse, error) {
	out := new(GetCoinBasesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/GetCoinBases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCoinBase(ctx context.Context, in *ExistCoinBaseRequest, opts ...grpc.CallOption) (*ExistCoinBaseResponse, error) {
	out := new(ExistCoinBaseResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/ExistCoinBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCoinBaseConds(ctx context.Context, in *ExistCoinBaseCondsRequest, opts ...grpc.CallOption) (*ExistCoinBaseCondsResponse, error) {
	out := new(ExistCoinBaseCondsResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/ExistCoinBaseConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountCoinBases(ctx context.Context, in *CountCoinBasesRequest, opts ...grpc.CallOption) (*CountCoinBasesResponse, error) {
	out := new(CountCoinBasesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/CountCoinBases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteCoinBase(ctx context.Context, in *DeleteCoinBaseRequest, opts ...grpc.CallOption) (*DeleteCoinBaseResponse, error) {
	out := new(DeleteCoinBaseResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.coin.base.v1.Manager/DeleteCoinBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateCoinBase(context.Context, *CreateCoinBaseRequest) (*CreateCoinBaseResponse, error)
	CreateCoinBases(context.Context, *CreateCoinBasesRequest) (*CreateCoinBasesResponse, error)
	AddCoinBase(context.Context, *AddCoinBaseRequest) (*AddCoinBaseResponse, error)
	GetCoinBase(context.Context, *GetCoinBaseRequest) (*GetCoinBaseResponse, error)
	GetCoinBaseOnly(context.Context, *GetCoinBaseOnlyRequest) (*GetCoinBaseOnlyResponse, error)
	GetCoinBases(context.Context, *GetCoinBasesRequest) (*GetCoinBasesResponse, error)
	ExistCoinBase(context.Context, *ExistCoinBaseRequest) (*ExistCoinBaseResponse, error)
	ExistCoinBaseConds(context.Context, *ExistCoinBaseCondsRequest) (*ExistCoinBaseCondsResponse, error)
	CountCoinBases(context.Context, *CountCoinBasesRequest) (*CountCoinBasesResponse, error)
	DeleteCoinBase(context.Context, *DeleteCoinBaseRequest) (*DeleteCoinBaseResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateCoinBase(context.Context, *CreateCoinBaseRequest) (*CreateCoinBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinBase not implemented")
}
func (UnimplementedManagerServer) CreateCoinBases(context.Context, *CreateCoinBasesRequest) (*CreateCoinBasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinBases not implemented")
}
func (UnimplementedManagerServer) AddCoinBase(context.Context, *AddCoinBaseRequest) (*AddCoinBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoinBase not implemented")
}
func (UnimplementedManagerServer) GetCoinBase(context.Context, *GetCoinBaseRequest) (*GetCoinBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinBase not implemented")
}
func (UnimplementedManagerServer) GetCoinBaseOnly(context.Context, *GetCoinBaseOnlyRequest) (*GetCoinBaseOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinBaseOnly not implemented")
}
func (UnimplementedManagerServer) GetCoinBases(context.Context, *GetCoinBasesRequest) (*GetCoinBasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinBases not implemented")
}
func (UnimplementedManagerServer) ExistCoinBase(context.Context, *ExistCoinBaseRequest) (*ExistCoinBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoinBase not implemented")
}
func (UnimplementedManagerServer) ExistCoinBaseConds(context.Context, *ExistCoinBaseCondsRequest) (*ExistCoinBaseCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCoinBaseConds not implemented")
}
func (UnimplementedManagerServer) CountCoinBases(context.Context, *CountCoinBasesRequest) (*CountCoinBasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCoinBases not implemented")
}
func (UnimplementedManagerServer) DeleteCoinBase(context.Context, *DeleteCoinBaseRequest) (*DeleteCoinBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoinBase not implemented")
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

func _Manager_CreateCoinBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCoinBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/CreateCoinBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCoinBase(ctx, req.(*CreateCoinBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateCoinBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCoinBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/CreateCoinBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCoinBases(ctx, req.(*CreateCoinBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_AddCoinBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoinBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).AddCoinBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/AddCoinBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).AddCoinBase(ctx, req.(*AddCoinBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCoinBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCoinBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/GetCoinBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCoinBase(ctx, req.(*GetCoinBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCoinBaseOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinBaseOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCoinBaseOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/GetCoinBaseOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCoinBaseOnly(ctx, req.(*GetCoinBaseOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCoinBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCoinBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/GetCoinBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCoinBases(ctx, req.(*GetCoinBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCoinBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCoinBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/ExistCoinBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCoinBase(ctx, req.(*ExistCoinBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCoinBaseConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCoinBaseCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCoinBaseConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/ExistCoinBaseConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCoinBaseConds(ctx, req.(*ExistCoinBaseCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountCoinBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCoinBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountCoinBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/CountCoinBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountCoinBases(ctx, req.(*CountCoinBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteCoinBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteCoinBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.coin.base.v1.Manager/DeleteCoinBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteCoinBase(ctx, req.(*DeleteCoinBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.manager.coin.base.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoinBase",
			Handler:    _Manager_CreateCoinBase_Handler,
		},
		{
			MethodName: "CreateCoinBases",
			Handler:    _Manager_CreateCoinBases_Handler,
		},
		{
			MethodName: "AddCoinBase",
			Handler:    _Manager_AddCoinBase_Handler,
		},
		{
			MethodName: "GetCoinBase",
			Handler:    _Manager_GetCoinBase_Handler,
		},
		{
			MethodName: "GetCoinBaseOnly",
			Handler:    _Manager_GetCoinBaseOnly_Handler,
		},
		{
			MethodName: "GetCoinBases",
			Handler:    _Manager_GetCoinBases_Handler,
		},
		{
			MethodName: "ExistCoinBase",
			Handler:    _Manager_ExistCoinBase_Handler,
		},
		{
			MethodName: "ExistCoinBaseConds",
			Handler:    _Manager_ExistCoinBaseConds_Handler,
		},
		{
			MethodName: "CountCoinBases",
			Handler:    _Manager_CountCoinBases_Handler,
		},
		{
			MethodName: "DeleteCoinBase",
			Handler:    _Manager_DeleteCoinBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mgr/v1/coin/base/base.proto",
}
