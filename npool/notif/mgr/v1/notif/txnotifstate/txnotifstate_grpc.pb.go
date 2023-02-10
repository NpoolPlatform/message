// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/mgr/v1/notif/txnotifstate/txnotifstate.proto

package txnotifstate

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
	CreateTxNotifState(ctx context.Context, in *CreateTxNotifStateRequest, opts ...grpc.CallOption) (*CreateTxNotifStateResponse, error)
	CreateTxNotifStates(ctx context.Context, in *CreateTxNotifStatesRequest, opts ...grpc.CallOption) (*CreateTxNotifStatesResponse, error)
	UpdateTxNotifState(ctx context.Context, in *UpdateTxNotifStateRequest, opts ...grpc.CallOption) (*UpdateTxNotifStateResponse, error)
	GetTxNotifState(ctx context.Context, in *GetTxNotifStateRequest, opts ...grpc.CallOption) (*GetTxNotifStateResponse, error)
	GetTxNotifStateOnly(ctx context.Context, in *GetTxNotifStateOnlyRequest, opts ...grpc.CallOption) (*GetTxNotifStateOnlyResponse, error)
	GetTxNotifStates(ctx context.Context, in *GetTxNotifStatesRequest, opts ...grpc.CallOption) (*GetTxNotifStatesResponse, error)
	ExistTxNotifState(ctx context.Context, in *ExistTxNotifStateRequest, opts ...grpc.CallOption) (*ExistTxNotifStateResponse, error)
	ExistTxNotifStateConds(ctx context.Context, in *ExistTxNotifStateCondsRequest, opts ...grpc.CallOption) (*ExistTxNotifStateCondsResponse, error)
	CountTxNotifStates(ctx context.Context, in *CountTxNotifStatesRequest, opts ...grpc.CallOption) (*CountTxNotifStatesResponse, error)
	DeleteTxNotifState(ctx context.Context, in *DeleteTxNotifStateRequest, opts ...grpc.CallOption) (*DeleteTxNotifStateResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateTxNotifState(ctx context.Context, in *CreateTxNotifStateRequest, opts ...grpc.CallOption) (*CreateTxNotifStateResponse, error) {
	out := new(CreateTxNotifStateResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/CreateTxNotifState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateTxNotifStates(ctx context.Context, in *CreateTxNotifStatesRequest, opts ...grpc.CallOption) (*CreateTxNotifStatesResponse, error) {
	out := new(CreateTxNotifStatesResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/CreateTxNotifStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateTxNotifState(ctx context.Context, in *UpdateTxNotifStateRequest, opts ...grpc.CallOption) (*UpdateTxNotifStateResponse, error) {
	out := new(UpdateTxNotifStateResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/UpdateTxNotifState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTxNotifState(ctx context.Context, in *GetTxNotifStateRequest, opts ...grpc.CallOption) (*GetTxNotifStateResponse, error) {
	out := new(GetTxNotifStateResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTxNotifStateOnly(ctx context.Context, in *GetTxNotifStateOnlyRequest, opts ...grpc.CallOption) (*GetTxNotifStateOnlyResponse, error) {
	out := new(GetTxNotifStateOnlyResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifStateOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetTxNotifStates(ctx context.Context, in *GetTxNotifStatesRequest, opts ...grpc.CallOption) (*GetTxNotifStatesResponse, error) {
	out := new(GetTxNotifStatesResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistTxNotifState(ctx context.Context, in *ExistTxNotifStateRequest, opts ...grpc.CallOption) (*ExistTxNotifStateResponse, error) {
	out := new(ExistTxNotifStateResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/ExistTxNotifState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistTxNotifStateConds(ctx context.Context, in *ExistTxNotifStateCondsRequest, opts ...grpc.CallOption) (*ExistTxNotifStateCondsResponse, error) {
	out := new(ExistTxNotifStateCondsResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/ExistTxNotifStateConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountTxNotifStates(ctx context.Context, in *CountTxNotifStatesRequest, opts ...grpc.CallOption) (*CountTxNotifStatesResponse, error) {
	out := new(CountTxNotifStatesResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/CountTxNotifStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteTxNotifState(ctx context.Context, in *DeleteTxNotifStateRequest, opts ...grpc.CallOption) (*DeleteTxNotifStateResponse, error) {
	out := new(DeleteTxNotifStateResponse)
	err := c.cc.Invoke(ctx, "/notif.manager.notif.txnotiftstate.v1.Manager/DeleteTxNotifState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateTxNotifState(context.Context, *CreateTxNotifStateRequest) (*CreateTxNotifStateResponse, error)
	CreateTxNotifStates(context.Context, *CreateTxNotifStatesRequest) (*CreateTxNotifStatesResponse, error)
	UpdateTxNotifState(context.Context, *UpdateTxNotifStateRequest) (*UpdateTxNotifStateResponse, error)
	GetTxNotifState(context.Context, *GetTxNotifStateRequest) (*GetTxNotifStateResponse, error)
	GetTxNotifStateOnly(context.Context, *GetTxNotifStateOnlyRequest) (*GetTxNotifStateOnlyResponse, error)
	GetTxNotifStates(context.Context, *GetTxNotifStatesRequest) (*GetTxNotifStatesResponse, error)
	ExistTxNotifState(context.Context, *ExistTxNotifStateRequest) (*ExistTxNotifStateResponse, error)
	ExistTxNotifStateConds(context.Context, *ExistTxNotifStateCondsRequest) (*ExistTxNotifStateCondsResponse, error)
	CountTxNotifStates(context.Context, *CountTxNotifStatesRequest) (*CountTxNotifStatesResponse, error)
	DeleteTxNotifState(context.Context, *DeleteTxNotifStateRequest) (*DeleteTxNotifStateResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateTxNotifState(context.Context, *CreateTxNotifStateRequest) (*CreateTxNotifStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTxNotifState not implemented")
}
func (UnimplementedManagerServer) CreateTxNotifStates(context.Context, *CreateTxNotifStatesRequest) (*CreateTxNotifStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTxNotifStates not implemented")
}
func (UnimplementedManagerServer) UpdateTxNotifState(context.Context, *UpdateTxNotifStateRequest) (*UpdateTxNotifStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTxNotifState not implemented")
}
func (UnimplementedManagerServer) GetTxNotifState(context.Context, *GetTxNotifStateRequest) (*GetTxNotifStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxNotifState not implemented")
}
func (UnimplementedManagerServer) GetTxNotifStateOnly(context.Context, *GetTxNotifStateOnlyRequest) (*GetTxNotifStateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxNotifStateOnly not implemented")
}
func (UnimplementedManagerServer) GetTxNotifStates(context.Context, *GetTxNotifStatesRequest) (*GetTxNotifStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxNotifStates not implemented")
}
func (UnimplementedManagerServer) ExistTxNotifState(context.Context, *ExistTxNotifStateRequest) (*ExistTxNotifStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTxNotifState not implemented")
}
func (UnimplementedManagerServer) ExistTxNotifStateConds(context.Context, *ExistTxNotifStateCondsRequest) (*ExistTxNotifStateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTxNotifStateConds not implemented")
}
func (UnimplementedManagerServer) CountTxNotifStates(context.Context, *CountTxNotifStatesRequest) (*CountTxNotifStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTxNotifStates not implemented")
}
func (UnimplementedManagerServer) DeleteTxNotifState(context.Context, *DeleteTxNotifStateRequest) (*DeleteTxNotifStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTxNotifState not implemented")
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

func _Manager_CreateTxNotifState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTxNotifStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateTxNotifState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/CreateTxNotifState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateTxNotifState(ctx, req.(*CreateTxNotifStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateTxNotifStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTxNotifStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateTxNotifStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/CreateTxNotifStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateTxNotifStates(ctx, req.(*CreateTxNotifStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateTxNotifState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTxNotifStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateTxNotifState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/UpdateTxNotifState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateTxNotifState(ctx, req.(*UpdateTxNotifStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTxNotifState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxNotifStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTxNotifState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTxNotifState(ctx, req.(*GetTxNotifStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTxNotifStateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxNotifStateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTxNotifStateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifStateOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTxNotifStateOnly(ctx, req.(*GetTxNotifStateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetTxNotifStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxNotifStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetTxNotifStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/GetTxNotifStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetTxNotifStates(ctx, req.(*GetTxNotifStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistTxNotifState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTxNotifStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistTxNotifState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/ExistTxNotifState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistTxNotifState(ctx, req.(*ExistTxNotifStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistTxNotifStateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistTxNotifStateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistTxNotifStateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/ExistTxNotifStateConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistTxNotifStateConds(ctx, req.(*ExistTxNotifStateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountTxNotifStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTxNotifStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountTxNotifStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/CountTxNotifStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountTxNotifStates(ctx, req.(*CountTxNotifStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteTxNotifState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTxNotifStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteTxNotifState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.manager.notif.txnotiftstate.v1.Manager/DeleteTxNotifState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteTxNotifState(ctx, req.(*DeleteTxNotifStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.manager.notif.txnotiftstate.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTxNotifState",
			Handler:    _Manager_CreateTxNotifState_Handler,
		},
		{
			MethodName: "CreateTxNotifStates",
			Handler:    _Manager_CreateTxNotifStates_Handler,
		},
		{
			MethodName: "UpdateTxNotifState",
			Handler:    _Manager_UpdateTxNotifState_Handler,
		},
		{
			MethodName: "GetTxNotifState",
			Handler:    _Manager_GetTxNotifState_Handler,
		},
		{
			MethodName: "GetTxNotifStateOnly",
			Handler:    _Manager_GetTxNotifStateOnly_Handler,
		},
		{
			MethodName: "GetTxNotifStates",
			Handler:    _Manager_GetTxNotifStates_Handler,
		},
		{
			MethodName: "ExistTxNotifState",
			Handler:    _Manager_ExistTxNotifState_Handler,
		},
		{
			MethodName: "ExistTxNotifStateConds",
			Handler:    _Manager_ExistTxNotifStateConds_Handler,
		},
		{
			MethodName: "CountTxNotifStates",
			Handler:    _Manager_CountTxNotifStates_Handler,
		},
		{
			MethodName: "DeleteTxNotifState",
			Handler:    _Manager_DeleteTxNotifState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mgr/v1/notif/txnotifstate/txnotifstate.proto",
}
