// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/order/mgr/v1/compensate/compensate.proto

package compensate

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
	Manager_CreateCompensate_FullMethodName     = "/order.manager.compensate.v1.Manager/CreateCompensate"
	Manager_CreateCompensates_FullMethodName    = "/order.manager.compensate.v1.Manager/CreateCompensates"
	Manager_UpdateCompensate_FullMethodName     = "/order.manager.compensate.v1.Manager/UpdateCompensate"
	Manager_GetCompensate_FullMethodName        = "/order.manager.compensate.v1.Manager/GetCompensate"
	Manager_GetCompensateOnly_FullMethodName    = "/order.manager.compensate.v1.Manager/GetCompensateOnly"
	Manager_GetCompensates_FullMethodName       = "/order.manager.compensate.v1.Manager/GetCompensates"
	Manager_ExistCompensate_FullMethodName      = "/order.manager.compensate.v1.Manager/ExistCompensate"
	Manager_ExistCompensateConds_FullMethodName = "/order.manager.compensate.v1.Manager/ExistCompensateConds"
	Manager_CountCompensates_FullMethodName     = "/order.manager.compensate.v1.Manager/CountCompensates"
	Manager_DeleteCompensate_FullMethodName     = "/order.manager.compensate.v1.Manager/DeleteCompensate"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateCompensate(ctx context.Context, in *CreateCompensateRequest, opts ...grpc.CallOption) (*CreateCompensateResponse, error)
	CreateCompensates(ctx context.Context, in *CreateCompensatesRequest, opts ...grpc.CallOption) (*CreateCompensatesResponse, error)
	UpdateCompensate(ctx context.Context, in *UpdateCompensateRequest, opts ...grpc.CallOption) (*UpdateCompensateResponse, error)
	GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error)
	GetCompensateOnly(ctx context.Context, in *GetCompensateOnlyRequest, opts ...grpc.CallOption) (*GetCompensateOnlyResponse, error)
	GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error)
	ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error)
	ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error)
	CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error)
	DeleteCompensate(ctx context.Context, in *DeleteCompensateRequest, opts ...grpc.CallOption) (*DeleteCompensateResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateCompensate(ctx context.Context, in *CreateCompensateRequest, opts ...grpc.CallOption) (*CreateCompensateResponse, error) {
	out := new(CreateCompensateResponse)
	err := c.cc.Invoke(ctx, Manager_CreateCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateCompensates(ctx context.Context, in *CreateCompensatesRequest, opts ...grpc.CallOption) (*CreateCompensatesResponse, error) {
	out := new(CreateCompensatesResponse)
	err := c.cc.Invoke(ctx, Manager_CreateCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateCompensate(ctx context.Context, in *UpdateCompensateRequest, opts ...grpc.CallOption) (*UpdateCompensateResponse, error) {
	out := new(UpdateCompensateResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCompensate(ctx context.Context, in *GetCompensateRequest, opts ...grpc.CallOption) (*GetCompensateResponse, error) {
	out := new(GetCompensateResponse)
	err := c.cc.Invoke(ctx, Manager_GetCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCompensateOnly(ctx context.Context, in *GetCompensateOnlyRequest, opts ...grpc.CallOption) (*GetCompensateOnlyResponse, error) {
	out := new(GetCompensateOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetCompensateOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetCompensates(ctx context.Context, in *GetCompensatesRequest, opts ...grpc.CallOption) (*GetCompensatesResponse, error) {
	out := new(GetCompensatesResponse)
	err := c.cc.Invoke(ctx, Manager_GetCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCompensate(ctx context.Context, in *ExistCompensateRequest, opts ...grpc.CallOption) (*ExistCompensateResponse, error) {
	out := new(ExistCompensateResponse)
	err := c.cc.Invoke(ctx, Manager_ExistCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistCompensateConds(ctx context.Context, in *ExistCompensateCondsRequest, opts ...grpc.CallOption) (*ExistCompensateCondsResponse, error) {
	out := new(ExistCompensateCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistCompensateConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountCompensates(ctx context.Context, in *CountCompensatesRequest, opts ...grpc.CallOption) (*CountCompensatesResponse, error) {
	out := new(CountCompensatesResponse)
	err := c.cc.Invoke(ctx, Manager_CountCompensates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteCompensate(ctx context.Context, in *DeleteCompensateRequest, opts ...grpc.CallOption) (*DeleteCompensateResponse, error) {
	out := new(DeleteCompensateResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteCompensate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateCompensate(context.Context, *CreateCompensateRequest) (*CreateCompensateResponse, error)
	CreateCompensates(context.Context, *CreateCompensatesRequest) (*CreateCompensatesResponse, error)
	UpdateCompensate(context.Context, *UpdateCompensateRequest) (*UpdateCompensateResponse, error)
	GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error)
	GetCompensateOnly(context.Context, *GetCompensateOnlyRequest) (*GetCompensateOnlyResponse, error)
	GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error)
	ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error)
	ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error)
	CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error)
	DeleteCompensate(context.Context, *DeleteCompensateRequest) (*DeleteCompensateResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateCompensate(context.Context, *CreateCompensateRequest) (*CreateCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompensate not implemented")
}
func (UnimplementedManagerServer) CreateCompensates(context.Context, *CreateCompensatesRequest) (*CreateCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompensates not implemented")
}
func (UnimplementedManagerServer) UpdateCompensate(context.Context, *UpdateCompensateRequest) (*UpdateCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompensate not implemented")
}
func (UnimplementedManagerServer) GetCompensate(context.Context, *GetCompensateRequest) (*GetCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensate not implemented")
}
func (UnimplementedManagerServer) GetCompensateOnly(context.Context, *GetCompensateOnlyRequest) (*GetCompensateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensateOnly not implemented")
}
func (UnimplementedManagerServer) GetCompensates(context.Context, *GetCompensatesRequest) (*GetCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompensates not implemented")
}
func (UnimplementedManagerServer) ExistCompensate(context.Context, *ExistCompensateRequest) (*ExistCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensate not implemented")
}
func (UnimplementedManagerServer) ExistCompensateConds(context.Context, *ExistCompensateCondsRequest) (*ExistCompensateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistCompensateConds not implemented")
}
func (UnimplementedManagerServer) CountCompensates(context.Context, *CountCompensatesRequest) (*CountCompensatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCompensates not implemented")
}
func (UnimplementedManagerServer) DeleteCompensate(context.Context, *DeleteCompensateRequest) (*DeleteCompensateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompensate not implemented")
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

func _Manager_CreateCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCompensate(ctx, req.(*CreateCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateCompensates(ctx, req.(*CreateCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateCompensate(ctx, req.(*UpdateCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCompensate(ctx, req.(*GetCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCompensateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCompensateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCompensateOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCompensateOnly(ctx, req.(*GetCompensateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetCompensates(ctx, req.(*GetCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCompensate(ctx, req.(*ExistCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistCompensateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistCompensateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistCompensateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistCompensateConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistCompensateConds(ctx, req.(*ExistCompensateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountCompensates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCompensatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountCompensates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CountCompensates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountCompensates(ctx, req.(*CountCompensatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompensateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteCompensate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteCompensate(ctx, req.(*DeleteCompensateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.manager.compensate.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompensate",
			Handler:    _Manager_CreateCompensate_Handler,
		},
		{
			MethodName: "CreateCompensates",
			Handler:    _Manager_CreateCompensates_Handler,
		},
		{
			MethodName: "UpdateCompensate",
			Handler:    _Manager_UpdateCompensate_Handler,
		},
		{
			MethodName: "GetCompensate",
			Handler:    _Manager_GetCompensate_Handler,
		},
		{
			MethodName: "GetCompensateOnly",
			Handler:    _Manager_GetCompensateOnly_Handler,
		},
		{
			MethodName: "GetCompensates",
			Handler:    _Manager_GetCompensates_Handler,
		},
		{
			MethodName: "ExistCompensate",
			Handler:    _Manager_ExistCompensate_Handler,
		},
		{
			MethodName: "ExistCompensateConds",
			Handler:    _Manager_ExistCompensateConds_Handler,
		},
		{
			MethodName: "CountCompensates",
			Handler:    _Manager_CountCompensates_Handler,
		},
		{
			MethodName: "DeleteCompensate",
			Handler:    _Manager_DeleteCompensate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mgr/v1/compensate/compensate.proto",
}
