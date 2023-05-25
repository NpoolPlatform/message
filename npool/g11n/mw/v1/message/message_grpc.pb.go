// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/g11n/mw/v1/message/message.proto

package message

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
	Middleware_CreateMessage_FullMethodName  = "/g11n.middleware.message1.v1.Middleware/CreateMessage"
	Middleware_CreateMessages_FullMethodName = "/g11n.middleware.message1.v1.Middleware/CreateMessages"
	Middleware_UpdateMessage_FullMethodName  = "/g11n.middleware.message1.v1.Middleware/UpdateMessage"
	Middleware_GetMessages_FullMethodName    = "/g11n.middleware.message1.v1.Middleware/GetMessages"
	Middleware_DeleteMessage_FullMethodName  = "/g11n.middleware.message1.v1.Middleware/DeleteMessage"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
	GetMessageOnly(ctx context.Context, in *GetMessageOnlyRequest, opts ...grpc.CallOption) (*GetMessageOnlyResponse, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	ExistMessage(ctx context.Context, in *ExistMessageRequest, opts ...grpc.CallOption) (*ExistMessageResponse, error)
	ExistMessageConds(ctx context.Context, in *ExistMessageCondsRequest, opts ...grpc.CallOption) (*ExistMessageCondsResponse, error)
	CountMessages(ctx context.Context, in *CountMessagesRequest, opts ...grpc.CallOption) (*CountMessagesResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateMessages(ctx context.Context, in *CreateMessagesRequest, opts ...grpc.CallOption) (*CreateMessagesResponse, error) {
	out := new(CreateMessagesResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error) {
	out := new(UpdateMessageResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, "/g11n.manager.message1.v1.Manager/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetMessageOnly(ctx context.Context, in *GetMessageOnlyRequest, opts ...grpc.CallOption) (*GetMessageOnlyResponse, error) {
	out := new(GetMessageOnlyResponse)
	err := c.cc.Invoke(ctx, "/g11n.manager.message1.v1.Manager/GetMessageOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistMessage(ctx context.Context, in *ExistMessageRequest, opts ...grpc.CallOption) (*ExistMessageResponse, error) {
	out := new(ExistMessageResponse)
	err := c.cc.Invoke(ctx, "/g11n.manager.message1.v1.Manager/ExistMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistMessageConds(ctx context.Context, in *ExistMessageCondsRequest, opts ...grpc.CallOption) (*ExistMessageCondsResponse, error) {
	out := new(ExistMessageCondsResponse)
	err := c.cc.Invoke(ctx, "/g11n.manager.message1.v1.Manager/ExistMessageConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountMessages(ctx context.Context, in *CountMessagesRequest, opts ...grpc.CallOption) (*CountMessagesResponse, error) {
	out := new(CountMessagesResponse)
	err := c.cc.Invoke(ctx, "/g11n.manager.message1.v1.Manager/CountMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error)
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	GetMessageOnly(context.Context, *GetMessageOnlyRequest) (*GetMessageOnlyResponse, error)
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	ExistMessage(context.Context, *ExistMessageRequest) (*ExistMessageResponse, error)
	ExistMessageConds(context.Context, *ExistMessageCondsRequest) (*ExistMessageCondsResponse, error)
	CountMessages(context.Context, *CountMessagesRequest) (*CountMessagesResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedManagerServer) CreateMessages(context.Context, *CreateMessagesRequest) (*CreateMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessages not implemented")
}
func (UnimplementedManagerServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedManagerServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedManagerServer) GetMessageOnly(context.Context, *GetMessageOnlyRequest) (*GetMessageOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageOnly not implemented")
}
func (UnimplementedManagerServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedManagerServer) ExistMessage(context.Context, *ExistMessageRequest) (*ExistMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMessage not implemented")
}
func (UnimplementedManagerServer) ExistMessageConds(context.Context, *ExistMessageCondsRequest) (*ExistMessageCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistMessageConds not implemented")
}
func (UnimplementedManagerServer) CountMessages(context.Context, *CountMessagesRequest) (*CountMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountMessages not implemented")
}
func (UnimplementedManagerServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
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

func _Manager_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateMessages(ctx, req.(*CreateMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.manager.message1.v1.Manager/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetMessageOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetMessageOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.manager.message1.v1.Manager/GetMessageOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetMessageOnly(ctx, req.(*GetMessageOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.manager.message1.v1.Manager/ExistMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistMessage(ctx, req.(*ExistMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistMessageConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistMessageCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistMessageConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.manager.message1.v1.Manager/ExistMessageConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistMessageConds(ctx, req.(*ExistMessageCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.manager.message1.v1.Manager/CountMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountMessages(ctx, req.(*CountMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.manager.message1.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Manager_CreateMessage_Handler,
		},
		{
			MethodName: "CreateMessages",
			Handler:    _Manager_CreateMessages_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _Manager_UpdateMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Manager_GetMessage_Handler,
		},
		{
			MethodName: "GetMessageOnly",
			Handler:    _Manager_GetMessageOnly_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Manager_GetMessages_Handler,
		},
		{
			MethodName: "ExistMessage",
			Handler:    _Manager_ExistMessage_Handler,
		},
		{
			MethodName: "ExistMessageConds",
			Handler:    _Manager_ExistMessageConds_Handler,
		},
		{
			MethodName: "CountMessages",
			Handler:    _Manager_CountMessages_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _Manager_DeleteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/mw/v1/message/message.proto",
}
