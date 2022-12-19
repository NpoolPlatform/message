// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/chain/mgr/v1/fiatcurrencytype/faitcurrencytype.proto

package fiatcurrencytype

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
	CreateFiatCurrencyType(ctx context.Context, in *CreateFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*CreateFiatCurrencyTypeResponse, error)
	CreateFiatCurrencyTypes(ctx context.Context, in *CreateFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*CreateFiatCurrencyTypesResponse, error)
	UpdateFiatCurrencyType(ctx context.Context, in *UpdateFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*UpdateFiatCurrencyTypeResponse, error)
	GetFiatCurrencyType(ctx context.Context, in *GetFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypeResponse, error)
	GetFiatCurrencyTypeOnly(ctx context.Context, in *GetFiatCurrencyTypeOnlyRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypeOnlyResponse, error)
	GetFiatCurrencyTypes(ctx context.Context, in *GetFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypesResponse, error)
	ExistFiatCurrencyType(ctx context.Context, in *ExistFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*ExistFiatCurrencyTypeResponse, error)
	ExistFiatCurrencyTypeConds(ctx context.Context, in *ExistFiatCurrencyTypeCondsRequest, opts ...grpc.CallOption) (*ExistFiatCurrencyTypeCondsResponse, error)
	CountFiatCurrencyTypes(ctx context.Context, in *CountFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*CountFiatCurrencyTypesResponse, error)
	DeleteFiatCurrencyType(ctx context.Context, in *DeleteFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*DeleteFiatCurrencyTypeResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateFiatCurrencyType(ctx context.Context, in *CreateFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*CreateFiatCurrencyTypeResponse, error) {
	out := new(CreateFiatCurrencyTypeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/CreateFiatCurrencyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateFiatCurrencyTypes(ctx context.Context, in *CreateFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*CreateFiatCurrencyTypesResponse, error) {
	out := new(CreateFiatCurrencyTypesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/CreateFiatCurrencyTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateFiatCurrencyType(ctx context.Context, in *UpdateFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*UpdateFiatCurrencyTypeResponse, error) {
	out := new(UpdateFiatCurrencyTypeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/UpdateFiatCurrencyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFiatCurrencyType(ctx context.Context, in *GetFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypeResponse, error) {
	out := new(GetFiatCurrencyTypeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFiatCurrencyTypeOnly(ctx context.Context, in *GetFiatCurrencyTypeOnlyRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypeOnlyResponse, error) {
	out := new(GetFiatCurrencyTypeOnlyResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyTypeOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetFiatCurrencyTypes(ctx context.Context, in *GetFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*GetFiatCurrencyTypesResponse, error) {
	out := new(GetFiatCurrencyTypesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistFiatCurrencyType(ctx context.Context, in *ExistFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*ExistFiatCurrencyTypeResponse, error) {
	out := new(ExistFiatCurrencyTypeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/ExistFiatCurrencyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistFiatCurrencyTypeConds(ctx context.Context, in *ExistFiatCurrencyTypeCondsRequest, opts ...grpc.CallOption) (*ExistFiatCurrencyTypeCondsResponse, error) {
	out := new(ExistFiatCurrencyTypeCondsResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/ExistFiatCurrencyTypeConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountFiatCurrencyTypes(ctx context.Context, in *CountFiatCurrencyTypesRequest, opts ...grpc.CallOption) (*CountFiatCurrencyTypesResponse, error) {
	out := new(CountFiatCurrencyTypesResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/CountFiatCurrencyTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteFiatCurrencyType(ctx context.Context, in *DeleteFiatCurrencyTypeRequest, opts ...grpc.CallOption) (*DeleteFiatCurrencyTypeResponse, error) {
	out := new(DeleteFiatCurrencyTypeResponse)
	err := c.cc.Invoke(ctx, "/chain.manager.fiatcurrencytype.v1.Manager/DeleteFiatCurrencyType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateFiatCurrencyType(context.Context, *CreateFiatCurrencyTypeRequest) (*CreateFiatCurrencyTypeResponse, error)
	CreateFiatCurrencyTypes(context.Context, *CreateFiatCurrencyTypesRequest) (*CreateFiatCurrencyTypesResponse, error)
	UpdateFiatCurrencyType(context.Context, *UpdateFiatCurrencyTypeRequest) (*UpdateFiatCurrencyTypeResponse, error)
	GetFiatCurrencyType(context.Context, *GetFiatCurrencyTypeRequest) (*GetFiatCurrencyTypeResponse, error)
	GetFiatCurrencyTypeOnly(context.Context, *GetFiatCurrencyTypeOnlyRequest) (*GetFiatCurrencyTypeOnlyResponse, error)
	GetFiatCurrencyTypes(context.Context, *GetFiatCurrencyTypesRequest) (*GetFiatCurrencyTypesResponse, error)
	ExistFiatCurrencyType(context.Context, *ExistFiatCurrencyTypeRequest) (*ExistFiatCurrencyTypeResponse, error)
	ExistFiatCurrencyTypeConds(context.Context, *ExistFiatCurrencyTypeCondsRequest) (*ExistFiatCurrencyTypeCondsResponse, error)
	CountFiatCurrencyTypes(context.Context, *CountFiatCurrencyTypesRequest) (*CountFiatCurrencyTypesResponse, error)
	DeleteFiatCurrencyType(context.Context, *DeleteFiatCurrencyTypeRequest) (*DeleteFiatCurrencyTypeResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateFiatCurrencyType(context.Context, *CreateFiatCurrencyTypeRequest) (*CreateFiatCurrencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFiatCurrencyType not implemented")
}
func (UnimplementedManagerServer) CreateFiatCurrencyTypes(context.Context, *CreateFiatCurrencyTypesRequest) (*CreateFiatCurrencyTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFiatCurrencyTypes not implemented")
}
func (UnimplementedManagerServer) UpdateFiatCurrencyType(context.Context, *UpdateFiatCurrencyTypeRequest) (*UpdateFiatCurrencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFiatCurrencyType not implemented")
}
func (UnimplementedManagerServer) GetFiatCurrencyType(context.Context, *GetFiatCurrencyTypeRequest) (*GetFiatCurrencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiatCurrencyType not implemented")
}
func (UnimplementedManagerServer) GetFiatCurrencyTypeOnly(context.Context, *GetFiatCurrencyTypeOnlyRequest) (*GetFiatCurrencyTypeOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiatCurrencyTypeOnly not implemented")
}
func (UnimplementedManagerServer) GetFiatCurrencyTypes(context.Context, *GetFiatCurrencyTypesRequest) (*GetFiatCurrencyTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiatCurrencyTypes not implemented")
}
func (UnimplementedManagerServer) ExistFiatCurrencyType(context.Context, *ExistFiatCurrencyTypeRequest) (*ExistFiatCurrencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFiatCurrencyType not implemented")
}
func (UnimplementedManagerServer) ExistFiatCurrencyTypeConds(context.Context, *ExistFiatCurrencyTypeCondsRequest) (*ExistFiatCurrencyTypeCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFiatCurrencyTypeConds not implemented")
}
func (UnimplementedManagerServer) CountFiatCurrencyTypes(context.Context, *CountFiatCurrencyTypesRequest) (*CountFiatCurrencyTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFiatCurrencyTypes not implemented")
}
func (UnimplementedManagerServer) DeleteFiatCurrencyType(context.Context, *DeleteFiatCurrencyTypeRequest) (*DeleteFiatCurrencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFiatCurrencyType not implemented")
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

func _Manager_CreateFiatCurrencyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFiatCurrencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateFiatCurrencyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/CreateFiatCurrencyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateFiatCurrencyType(ctx, req.(*CreateFiatCurrencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateFiatCurrencyTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFiatCurrencyTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateFiatCurrencyTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/CreateFiatCurrencyTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateFiatCurrencyTypes(ctx, req.(*CreateFiatCurrencyTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateFiatCurrencyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFiatCurrencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateFiatCurrencyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/UpdateFiatCurrencyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateFiatCurrencyType(ctx, req.(*UpdateFiatCurrencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFiatCurrencyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiatCurrencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFiatCurrencyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFiatCurrencyType(ctx, req.(*GetFiatCurrencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFiatCurrencyTypeOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiatCurrencyTypeOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFiatCurrencyTypeOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyTypeOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFiatCurrencyTypeOnly(ctx, req.(*GetFiatCurrencyTypeOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetFiatCurrencyTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiatCurrencyTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetFiatCurrencyTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/GetFiatCurrencyTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetFiatCurrencyTypes(ctx, req.(*GetFiatCurrencyTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistFiatCurrencyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFiatCurrencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistFiatCurrencyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/ExistFiatCurrencyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistFiatCurrencyType(ctx, req.(*ExistFiatCurrencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistFiatCurrencyTypeConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFiatCurrencyTypeCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistFiatCurrencyTypeConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/ExistFiatCurrencyTypeConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistFiatCurrencyTypeConds(ctx, req.(*ExistFiatCurrencyTypeCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountFiatCurrencyTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFiatCurrencyTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountFiatCurrencyTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/CountFiatCurrencyTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountFiatCurrencyTypes(ctx, req.(*CountFiatCurrencyTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteFiatCurrencyType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFiatCurrencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteFiatCurrencyType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain.manager.fiatcurrencytype.v1.Manager/DeleteFiatCurrencyType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteFiatCurrencyType(ctx, req.(*DeleteFiatCurrencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.manager.fiatcurrencytype.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFiatCurrencyType",
			Handler:    _Manager_CreateFiatCurrencyType_Handler,
		},
		{
			MethodName: "CreateFiatCurrencyTypes",
			Handler:    _Manager_CreateFiatCurrencyTypes_Handler,
		},
		{
			MethodName: "UpdateFiatCurrencyType",
			Handler:    _Manager_UpdateFiatCurrencyType_Handler,
		},
		{
			MethodName: "GetFiatCurrencyType",
			Handler:    _Manager_GetFiatCurrencyType_Handler,
		},
		{
			MethodName: "GetFiatCurrencyTypeOnly",
			Handler:    _Manager_GetFiatCurrencyTypeOnly_Handler,
		},
		{
			MethodName: "GetFiatCurrencyTypes",
			Handler:    _Manager_GetFiatCurrencyTypes_Handler,
		},
		{
			MethodName: "ExistFiatCurrencyType",
			Handler:    _Manager_ExistFiatCurrencyType_Handler,
		},
		{
			MethodName: "ExistFiatCurrencyTypeConds",
			Handler:    _Manager_ExistFiatCurrencyTypeConds_Handler,
		},
		{
			MethodName: "CountFiatCurrencyTypes",
			Handler:    _Manager_CountFiatCurrencyTypes_Handler,
		},
		{
			MethodName: "DeleteFiatCurrencyType",
			Handler:    _Manager_DeleteFiatCurrencyType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/chain/mgr/v1/fiatcurrencytype/faitcurrencytype.proto",
}
