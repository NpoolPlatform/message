// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/account/mgr/v1/limitation/limitation.proto

package limitation

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
	CreateLimitation(ctx context.Context, in *CreateLimitationRequest, opts ...grpc.CallOption) (*CreateLimitationResponse, error)
	CreateLimitations(ctx context.Context, in *CreateLimitationsRequest, opts ...grpc.CallOption) (*CreateLimitationsResponse, error)
	UpdateLimitation(ctx context.Context, in *UpdateLimitationRequest, opts ...grpc.CallOption) (*UpdateLimitationResponse, error)
	GetLimitation(ctx context.Context, in *GetLimitationRequest, opts ...grpc.CallOption) (*GetLimitationResponse, error)
	GetLimitationOnly(ctx context.Context, in *GetLimitationOnlyRequest, opts ...grpc.CallOption) (*GetLimitationOnlyResponse, error)
	GetLimitations(ctx context.Context, in *GetLimitationsRequest, opts ...grpc.CallOption) (*GetLimitationsResponse, error)
	ExistLimitation(ctx context.Context, in *ExistLimitationRequest, opts ...grpc.CallOption) (*ExistLimitationResponse, error)
	ExistLimitationConds(ctx context.Context, in *ExistLimitationCondsRequest, opts ...grpc.CallOption) (*ExistLimitationCondsResponse, error)
	CountLimitations(ctx context.Context, in *CountLimitationsRequest, opts ...grpc.CallOption) (*CountLimitationsResponse, error)
	DeleteLimitation(ctx context.Context, in *DeleteLimitationRequest, opts ...grpc.CallOption) (*DeleteLimitationResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateLimitation(ctx context.Context, in *CreateLimitationRequest, opts ...grpc.CallOption) (*CreateLimitationResponse, error) {
	out := new(CreateLimitationResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/CreateLimitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateLimitations(ctx context.Context, in *CreateLimitationsRequest, opts ...grpc.CallOption) (*CreateLimitationsResponse, error) {
	out := new(CreateLimitationsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/CreateLimitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateLimitation(ctx context.Context, in *UpdateLimitationRequest, opts ...grpc.CallOption) (*UpdateLimitationResponse, error) {
	out := new(UpdateLimitationResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/UpdateLimitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLimitation(ctx context.Context, in *GetLimitationRequest, opts ...grpc.CallOption) (*GetLimitationResponse, error) {
	out := new(GetLimitationResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/GetLimitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLimitationOnly(ctx context.Context, in *GetLimitationOnlyRequest, opts ...grpc.CallOption) (*GetLimitationOnlyResponse, error) {
	out := new(GetLimitationOnlyResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/GetLimitationOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLimitations(ctx context.Context, in *GetLimitationsRequest, opts ...grpc.CallOption) (*GetLimitationsResponse, error) {
	out := new(GetLimitationsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/GetLimitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistLimitation(ctx context.Context, in *ExistLimitationRequest, opts ...grpc.CallOption) (*ExistLimitationResponse, error) {
	out := new(ExistLimitationResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/ExistLimitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistLimitationConds(ctx context.Context, in *ExistLimitationCondsRequest, opts ...grpc.CallOption) (*ExistLimitationCondsResponse, error) {
	out := new(ExistLimitationCondsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/ExistLimitationConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountLimitations(ctx context.Context, in *CountLimitationsRequest, opts ...grpc.CallOption) (*CountLimitationsResponse, error) {
	out := new(CountLimitationsResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/CountLimitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteLimitation(ctx context.Context, in *DeleteLimitationRequest, opts ...grpc.CallOption) (*DeleteLimitationResponse, error) {
	out := new(DeleteLimitationResponse)
	err := c.cc.Invoke(ctx, "/account.manager.limitation.v1.Manager/DeleteLimitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateLimitation(context.Context, *CreateLimitationRequest) (*CreateLimitationResponse, error)
	CreateLimitations(context.Context, *CreateLimitationsRequest) (*CreateLimitationsResponse, error)
	UpdateLimitation(context.Context, *UpdateLimitationRequest) (*UpdateLimitationResponse, error)
	GetLimitation(context.Context, *GetLimitationRequest) (*GetLimitationResponse, error)
	GetLimitationOnly(context.Context, *GetLimitationOnlyRequest) (*GetLimitationOnlyResponse, error)
	GetLimitations(context.Context, *GetLimitationsRequest) (*GetLimitationsResponse, error)
	ExistLimitation(context.Context, *ExistLimitationRequest) (*ExistLimitationResponse, error)
	ExistLimitationConds(context.Context, *ExistLimitationCondsRequest) (*ExistLimitationCondsResponse, error)
	CountLimitations(context.Context, *CountLimitationsRequest) (*CountLimitationsResponse, error)
	DeleteLimitation(context.Context, *DeleteLimitationRequest) (*DeleteLimitationResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateLimitation(context.Context, *CreateLimitationRequest) (*CreateLimitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimitation not implemented")
}
func (UnimplementedManagerServer) CreateLimitations(context.Context, *CreateLimitationsRequest) (*CreateLimitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimitations not implemented")
}
func (UnimplementedManagerServer) UpdateLimitation(context.Context, *UpdateLimitationRequest) (*UpdateLimitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLimitation not implemented")
}
func (UnimplementedManagerServer) GetLimitation(context.Context, *GetLimitationRequest) (*GetLimitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimitation not implemented")
}
func (UnimplementedManagerServer) GetLimitationOnly(context.Context, *GetLimitationOnlyRequest) (*GetLimitationOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimitationOnly not implemented")
}
func (UnimplementedManagerServer) GetLimitations(context.Context, *GetLimitationsRequest) (*GetLimitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimitations not implemented")
}
func (UnimplementedManagerServer) ExistLimitation(context.Context, *ExistLimitationRequest) (*ExistLimitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistLimitation not implemented")
}
func (UnimplementedManagerServer) ExistLimitationConds(context.Context, *ExistLimitationCondsRequest) (*ExistLimitationCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistLimitationConds not implemented")
}
func (UnimplementedManagerServer) CountLimitations(context.Context, *CountLimitationsRequest) (*CountLimitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountLimitations not implemented")
}
func (UnimplementedManagerServer) DeleteLimitation(context.Context, *DeleteLimitationRequest) (*DeleteLimitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLimitation not implemented")
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

func _Manager_CreateLimitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateLimitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/CreateLimitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateLimitation(ctx, req.(*CreateLimitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateLimitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateLimitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/CreateLimitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateLimitations(ctx, req.(*CreateLimitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateLimitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLimitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateLimitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/UpdateLimitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateLimitation(ctx, req.(*UpdateLimitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLimitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLimitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/GetLimitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLimitation(ctx, req.(*GetLimitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLimitationOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitationOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLimitationOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/GetLimitationOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLimitationOnly(ctx, req.(*GetLimitationOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLimitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLimitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/GetLimitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLimitations(ctx, req.(*GetLimitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistLimitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistLimitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistLimitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/ExistLimitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistLimitation(ctx, req.(*ExistLimitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistLimitationConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistLimitationCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistLimitationConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/ExistLimitationConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistLimitationConds(ctx, req.(*ExistLimitationCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountLimitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountLimitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountLimitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/CountLimitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountLimitations(ctx, req.(*CountLimitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteLimitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLimitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteLimitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.manager.limitation.v1.Manager/DeleteLimitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteLimitation(ctx, req.(*DeleteLimitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.manager.limitation.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLimitation",
			Handler:    _Manager_CreateLimitation_Handler,
		},
		{
			MethodName: "CreateLimitations",
			Handler:    _Manager_CreateLimitations_Handler,
		},
		{
			MethodName: "UpdateLimitation",
			Handler:    _Manager_UpdateLimitation_Handler,
		},
		{
			MethodName: "GetLimitation",
			Handler:    _Manager_GetLimitation_Handler,
		},
		{
			MethodName: "GetLimitationOnly",
			Handler:    _Manager_GetLimitationOnly_Handler,
		},
		{
			MethodName: "GetLimitations",
			Handler:    _Manager_GetLimitations_Handler,
		},
		{
			MethodName: "ExistLimitation",
			Handler:    _Manager_ExistLimitation_Handler,
		},
		{
			MethodName: "ExistLimitationConds",
			Handler:    _Manager_ExistLimitationConds_Handler,
		},
		{
			MethodName: "CountLimitations",
			Handler:    _Manager_CountLimitations_Handler,
		},
		{
			MethodName: "DeleteLimitation",
			Handler:    _Manager_DeleteLimitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/account/mgr/v1/limitation/limitation.proto",
}
