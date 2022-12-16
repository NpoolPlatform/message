// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/subscriber/subscriber.proto

package subscriber

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
	CreateSubscriber(ctx context.Context, in *CreateSubscriberRequest, opts ...grpc.CallOption) (*CreateSubscriberResponse, error)
	CreateSubscriberes(ctx context.Context, in *CreateSubscriberesRequest, opts ...grpc.CallOption) (*CreateSubscriberesResponse, error)
	UpdateSubscriber(ctx context.Context, in *UpdateSubscriberRequest, opts ...grpc.CallOption) (*UpdateSubscriberResponse, error)
	GetSubscriber(ctx context.Context, in *GetSubscriberRequest, opts ...grpc.CallOption) (*GetSubscriberResponse, error)
	GetSubscriberOnly(ctx context.Context, in *GetSubscriberOnlyRequest, opts ...grpc.CallOption) (*GetSubscriberOnlyResponse, error)
	GetSubscriberes(ctx context.Context, in *GetSubscriberesRequest, opts ...grpc.CallOption) (*GetSubscriberesResponse, error)
	ExistSubscriber(ctx context.Context, in *ExistSubscriberRequest, opts ...grpc.CallOption) (*ExistSubscriberResponse, error)
	ExistSubscriberConds(ctx context.Context, in *ExistSubscriberCondsRequest, opts ...grpc.CallOption) (*ExistSubscriberCondsResponse, error)
	CountSubscriberes(ctx context.Context, in *CountSubscriberesRequest, opts ...grpc.CallOption) (*CountSubscriberesResponse, error)
	DeleteSubscriber(ctx context.Context, in *DeleteSubscriberRequest, opts ...grpc.CallOption) (*DeleteSubscriberResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateSubscriber(ctx context.Context, in *CreateSubscriberRequest, opts ...grpc.CallOption) (*CreateSubscriberResponse, error) {
	out := new(CreateSubscriberResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/CreateSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateSubscriberes(ctx context.Context, in *CreateSubscriberesRequest, opts ...grpc.CallOption) (*CreateSubscriberesResponse, error) {
	out := new(CreateSubscriberesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/CreateSubscriberes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateSubscriber(ctx context.Context, in *UpdateSubscriberRequest, opts ...grpc.CallOption) (*UpdateSubscriberResponse, error) {
	out := new(UpdateSubscriberResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/UpdateSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSubscriber(ctx context.Context, in *GetSubscriberRequest, opts ...grpc.CallOption) (*GetSubscriberResponse, error) {
	out := new(GetSubscriberResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/GetSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSubscriberOnly(ctx context.Context, in *GetSubscriberOnlyRequest, opts ...grpc.CallOption) (*GetSubscriberOnlyResponse, error) {
	out := new(GetSubscriberOnlyResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/GetSubscriberOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetSubscriberes(ctx context.Context, in *GetSubscriberesRequest, opts ...grpc.CallOption) (*GetSubscriberesResponse, error) {
	out := new(GetSubscriberesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/GetSubscriberes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistSubscriber(ctx context.Context, in *ExistSubscriberRequest, opts ...grpc.CallOption) (*ExistSubscriberResponse, error) {
	out := new(ExistSubscriberResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/ExistSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistSubscriberConds(ctx context.Context, in *ExistSubscriberCondsRequest, opts ...grpc.CallOption) (*ExistSubscriberCondsResponse, error) {
	out := new(ExistSubscriberCondsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/ExistSubscriberConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountSubscriberes(ctx context.Context, in *CountSubscriberesRequest, opts ...grpc.CallOption) (*CountSubscriberesResponse, error) {
	out := new(CountSubscriberesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/CountSubscriberes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteSubscriber(ctx context.Context, in *DeleteSubscriberRequest, opts ...grpc.CallOption) (*DeleteSubscriberResponse, error) {
	out := new(DeleteSubscriberResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.subscriber.v2.Manager/DeleteSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateSubscriber(context.Context, *CreateSubscriberRequest) (*CreateSubscriberResponse, error)
	CreateSubscriberes(context.Context, *CreateSubscriberesRequest) (*CreateSubscriberesResponse, error)
	UpdateSubscriber(context.Context, *UpdateSubscriberRequest) (*UpdateSubscriberResponse, error)
	GetSubscriber(context.Context, *GetSubscriberRequest) (*GetSubscriberResponse, error)
	GetSubscriberOnly(context.Context, *GetSubscriberOnlyRequest) (*GetSubscriberOnlyResponse, error)
	GetSubscriberes(context.Context, *GetSubscriberesRequest) (*GetSubscriberesResponse, error)
	ExistSubscriber(context.Context, *ExistSubscriberRequest) (*ExistSubscriberResponse, error)
	ExistSubscriberConds(context.Context, *ExistSubscriberCondsRequest) (*ExistSubscriberCondsResponse, error)
	CountSubscriberes(context.Context, *CountSubscriberesRequest) (*CountSubscriberesResponse, error)
	DeleteSubscriber(context.Context, *DeleteSubscriberRequest) (*DeleteSubscriberResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateSubscriber(context.Context, *CreateSubscriberRequest) (*CreateSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscriber not implemented")
}
func (UnimplementedManagerServer) CreateSubscriberes(context.Context, *CreateSubscriberesRequest) (*CreateSubscriberesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscriberes not implemented")
}
func (UnimplementedManagerServer) UpdateSubscriber(context.Context, *UpdateSubscriberRequest) (*UpdateSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscriber not implemented")
}
func (UnimplementedManagerServer) GetSubscriber(context.Context, *GetSubscriberRequest) (*GetSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriber not implemented")
}
func (UnimplementedManagerServer) GetSubscriberOnly(context.Context, *GetSubscriberOnlyRequest) (*GetSubscriberOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberOnly not implemented")
}
func (UnimplementedManagerServer) GetSubscriberes(context.Context, *GetSubscriberesRequest) (*GetSubscriberesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberes not implemented")
}
func (UnimplementedManagerServer) ExistSubscriber(context.Context, *ExistSubscriberRequest) (*ExistSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSubscriber not implemented")
}
func (UnimplementedManagerServer) ExistSubscriberConds(context.Context, *ExistSubscriberCondsRequest) (*ExistSubscriberCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistSubscriberConds not implemented")
}
func (UnimplementedManagerServer) CountSubscriberes(context.Context, *CountSubscriberesRequest) (*CountSubscriberesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSubscriberes not implemented")
}
func (UnimplementedManagerServer) DeleteSubscriber(context.Context, *DeleteSubscriberRequest) (*DeleteSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscriber not implemented")
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

func _Manager_CreateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/CreateSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateSubscriber(ctx, req.(*CreateSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateSubscriberes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriberesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateSubscriberes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/CreateSubscriberes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateSubscriberes(ctx, req.(*CreateSubscriberesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/UpdateSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateSubscriber(ctx, req.(*UpdateSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/GetSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSubscriber(ctx, req.(*GetSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSubscriberOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriberOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSubscriberOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/GetSubscriberOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSubscriberOnly(ctx, req.(*GetSubscriberOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetSubscriberes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriberesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSubscriberes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/GetSubscriberes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSubscriberes(ctx, req.(*GetSubscriberesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/ExistSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistSubscriber(ctx, req.(*ExistSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistSubscriberConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistSubscriberCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistSubscriberConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/ExistSubscriberConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistSubscriberConds(ctx, req.(*ExistSubscriberCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountSubscriberes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountSubscriberesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountSubscriberes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/CountSubscriberes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountSubscriberes(ctx, req.(*CountSubscriberesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.subscriber.v2.Manager/DeleteSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteSubscriber(ctx, req.(*DeleteSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.manager.subscriber.v2.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubscriber",
			Handler:    _Manager_CreateSubscriber_Handler,
		},
		{
			MethodName: "CreateSubscriberes",
			Handler:    _Manager_CreateSubscriberes_Handler,
		},
		{
			MethodName: "UpdateSubscriber",
			Handler:    _Manager_UpdateSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriber",
			Handler:    _Manager_GetSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriberOnly",
			Handler:    _Manager_GetSubscriberOnly_Handler,
		},
		{
			MethodName: "GetSubscriberes",
			Handler:    _Manager_GetSubscriberes_Handler,
		},
		{
			MethodName: "ExistSubscriber",
			Handler:    _Manager_ExistSubscriber_Handler,
		},
		{
			MethodName: "ExistSubscriberConds",
			Handler:    _Manager_ExistSubscriberConds_Handler,
		},
		{
			MethodName: "CountSubscriberes",
			Handler:    _Manager_CountSubscriberes_Handler,
		},
		{
			MethodName: "DeleteSubscriber",
			Handler:    _Manager_DeleteSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/subscriber/subscriber.proto",
}
