// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/order/mgr/v1/order/state/state.proto

package state

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
	CreateState(ctx context.Context, in *CreateStateRequest, opts ...grpc.CallOption) (*CreateStateResponse, error)
	CreateStates(ctx context.Context, in *CreateStatesRequest, opts ...grpc.CallOption) (*CreateStatesResponse, error)
	UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*UpdateStateResponse, error)
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	GetStateOnly(ctx context.Context, in *GetStateOnlyRequest, opts ...grpc.CallOption) (*GetStateOnlyResponse, error)
	GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error)
	ExistState(ctx context.Context, in *ExistStateRequest, opts ...grpc.CallOption) (*ExistStateResponse, error)
	ExistStateConds(ctx context.Context, in *ExistStateCondsRequest, opts ...grpc.CallOption) (*ExistStateCondsResponse, error)
	CountStates(ctx context.Context, in *CountStatesRequest, opts ...grpc.CallOption) (*CountStatesResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateState(ctx context.Context, in *CreateStateRequest, opts ...grpc.CallOption) (*CreateStateResponse, error) {
	out := new(CreateStateResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/CreateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateStates(ctx context.Context, in *CreateStatesRequest, opts ...grpc.CallOption) (*CreateStatesResponse, error) {
	out := new(CreateStatesResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/CreateStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*UpdateStateResponse, error) {
	out := new(UpdateStateResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStateOnly(ctx context.Context, in *GetStateOnlyRequest, opts ...grpc.CallOption) (*GetStateOnlyResponse, error) {
	out := new(GetStateOnlyResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/GetStateOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetStates(ctx context.Context, in *GetStatesRequest, opts ...grpc.CallOption) (*GetStatesResponse, error) {
	out := new(GetStatesResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/GetStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistState(ctx context.Context, in *ExistStateRequest, opts ...grpc.CallOption) (*ExistStateResponse, error) {
	out := new(ExistStateResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/ExistState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistStateConds(ctx context.Context, in *ExistStateCondsRequest, opts ...grpc.CallOption) (*ExistStateCondsResponse, error) {
	out := new(ExistStateCondsResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/ExistStateConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountStates(ctx context.Context, in *CountStatesRequest, opts ...grpc.CallOption) (*CountStatesResponse, error) {
	out := new(CountStatesResponse)
	err := c.cc.Invoke(ctx, "/order.manager.order.state.v1.Manager/CountStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateState(context.Context, *CreateStateRequest) (*CreateStateResponse, error)
	CreateStates(context.Context, *CreateStatesRequest) (*CreateStatesResponse, error)
	UpdateState(context.Context, *UpdateStateRequest) (*UpdateStateResponse, error)
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	GetStateOnly(context.Context, *GetStateOnlyRequest) (*GetStateOnlyResponse, error)
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)
	ExistState(context.Context, *ExistStateRequest) (*ExistStateResponse, error)
	ExistStateConds(context.Context, *ExistStateCondsRequest) (*ExistStateCondsResponse, error)
	CountStates(context.Context, *CountStatesRequest) (*CountStatesResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateState(context.Context, *CreateStateRequest) (*CreateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateState not implemented")
}
func (UnimplementedManagerServer) CreateStates(context.Context, *CreateStatesRequest) (*CreateStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStates not implemented")
}
func (UnimplementedManagerServer) UpdateState(context.Context, *UpdateStateRequest) (*UpdateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (UnimplementedManagerServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedManagerServer) GetStateOnly(context.Context, *GetStateOnlyRequest) (*GetStateOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateOnly not implemented")
}
func (UnimplementedManagerServer) GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStates not implemented")
}
func (UnimplementedManagerServer) ExistState(context.Context, *ExistStateRequest) (*ExistStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistState not implemented")
}
func (UnimplementedManagerServer) ExistStateConds(context.Context, *ExistStateCondsRequest) (*ExistStateCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistStateConds not implemented")
}
func (UnimplementedManagerServer) CountStates(context.Context, *CountStatesRequest) (*CountStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountStates not implemented")
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

func _Manager_CreateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/CreateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateState(ctx, req.(*CreateStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/CreateStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateStates(ctx, req.(*CreateStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateState(ctx, req.(*UpdateStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetStateOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStateOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/GetStateOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStateOnly(ctx, req.(*GetStateOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/GetStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetStates(ctx, req.(*GetStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/ExistState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistState(ctx, req.(*ExistStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistStateConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistStateCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistStateConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/ExistStateConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistStateConds(ctx, req.(*ExistStateCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.manager.order.state.v1.Manager/CountStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountStates(ctx, req.(*CountStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.manager.order.state.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateState",
			Handler:    _Manager_CreateState_Handler,
		},
		{
			MethodName: "CreateStates",
			Handler:    _Manager_CreateStates_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Manager_UpdateState_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Manager_GetState_Handler,
		},
		{
			MethodName: "GetStateOnly",
			Handler:    _Manager_GetStateOnly_Handler,
		},
		{
			MethodName: "GetStates",
			Handler:    _Manager_GetStates_Handler,
		},
		{
			MethodName: "ExistState",
			Handler:    _Manager_ExistState_Handler,
		},
		{
			MethodName: "ExistStateConds",
			Handler:    _Manager_ExistStateConds_Handler,
		},
		{
			MethodName: "CountStates",
			Handler:    _Manager_CountStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/order/mgr/v1/order/state/state.proto",
}
