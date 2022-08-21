// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/login/history/history.proto

package history

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
	CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error)
	CreateHistories(ctx context.Context, in *CreateHistoriesRequest, opts ...grpc.CallOption) (*CreateHistoriesResponse, error)
	UpdateHistory(ctx context.Context, in *UpdateHistoryRequest, opts ...grpc.CallOption) (*UpdateHistoryResponse, error)
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error)
	GetHistoryOnly(ctx context.Context, in *GetHistoryOnlyRequest, opts ...grpc.CallOption) (*GetHistoryOnlyResponse, error)
	GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error)
	ExistHistory(ctx context.Context, in *ExistHistoryRequest, opts ...grpc.CallOption) (*ExistHistoryResponse, error)
	ExistHistoryConds(ctx context.Context, in *ExistHistoryCondsRequest, opts ...grpc.CallOption) (*ExistHistoryCondsResponse, error)
	CountHistories(ctx context.Context, in *CountHistoriesRequest, opts ...grpc.CallOption) (*CountHistoriesResponse, error)
	DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*DeleteHistoryResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error) {
	out := new(CreateHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/CreateHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateHistories(ctx context.Context, in *CreateHistoriesRequest, opts ...grpc.CallOption) (*CreateHistoriesResponse, error) {
	out := new(CreateHistoriesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/CreateHistories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateHistory(ctx context.Context, in *UpdateHistoryRequest, opts ...grpc.CallOption) (*UpdateHistoryResponse, error) {
	out := new(UpdateHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/UpdateHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error) {
	out := new(GetHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetHistoryOnly(ctx context.Context, in *GetHistoryOnlyRequest, opts ...grpc.CallOption) (*GetHistoryOnlyResponse, error) {
	out := new(GetHistoryOnlyResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/GetHistoryOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error) {
	out := new(GetHistoriesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/GetHistories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistHistory(ctx context.Context, in *ExistHistoryRequest, opts ...grpc.CallOption) (*ExistHistoryResponse, error) {
	out := new(ExistHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/ExistHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistHistoryConds(ctx context.Context, in *ExistHistoryCondsRequest, opts ...grpc.CallOption) (*ExistHistoryCondsResponse, error) {
	out := new(ExistHistoryCondsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/ExistHistoryConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountHistories(ctx context.Context, in *CountHistoriesRequest, opts ...grpc.CallOption) (*CountHistoriesResponse, error) {
	out := new(CountHistoriesResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/CountHistories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteHistory(ctx context.Context, in *DeleteHistoryRequest, opts ...grpc.CallOption) (*DeleteHistoryResponse, error) {
	out := new(DeleteHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.login.history.v2.Manager/DeleteHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)
	CreateHistories(context.Context, *CreateHistoriesRequest) (*CreateHistoriesResponse, error)
	UpdateHistory(context.Context, *UpdateHistoryRequest) (*UpdateHistoryResponse, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	GetHistoryOnly(context.Context, *GetHistoryOnlyRequest) (*GetHistoryOnlyResponse, error)
	GetHistories(context.Context, *GetHistoriesRequest) (*GetHistoriesResponse, error)
	ExistHistory(context.Context, *ExistHistoryRequest) (*ExistHistoryResponse, error)
	ExistHistoryConds(context.Context, *ExistHistoryCondsRequest) (*ExistHistoryCondsResponse, error)
	CountHistories(context.Context, *CountHistoriesRequest) (*CountHistoriesResponse, error)
	DeleteHistory(context.Context, *DeleteHistoryRequest) (*DeleteHistoryResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistory not implemented")
}
func (UnimplementedManagerServer) CreateHistories(context.Context, *CreateHistoriesRequest) (*CreateHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistories not implemented")
}
func (UnimplementedManagerServer) UpdateHistory(context.Context, *UpdateHistoryRequest) (*UpdateHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHistory not implemented")
}
func (UnimplementedManagerServer) GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedManagerServer) GetHistoryOnly(context.Context, *GetHistoryOnlyRequest) (*GetHistoryOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoryOnly not implemented")
}
func (UnimplementedManagerServer) GetHistories(context.Context, *GetHistoriesRequest) (*GetHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistories not implemented")
}
func (UnimplementedManagerServer) ExistHistory(context.Context, *ExistHistoryRequest) (*ExistHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistHistory not implemented")
}
func (UnimplementedManagerServer) ExistHistoryConds(context.Context, *ExistHistoryCondsRequest) (*ExistHistoryCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistHistoryConds not implemented")
}
func (UnimplementedManagerServer) CountHistories(context.Context, *CountHistoriesRequest) (*CountHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountHistories not implemented")
}
func (UnimplementedManagerServer) DeleteHistory(context.Context, *DeleteHistoryRequest) (*DeleteHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHistory not implemented")
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

func _Manager_CreateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/CreateHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateHistory(ctx, req.(*CreateHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/CreateHistories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateHistories(ctx, req.(*CreateHistoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/UpdateHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateHistory(ctx, req.(*UpdateHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetHistory(ctx, req.(*GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetHistoryOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetHistoryOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/GetHistoryOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetHistoryOnly(ctx, req.(*GetHistoryOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/GetHistories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetHistories(ctx, req.(*GetHistoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/ExistHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistHistory(ctx, req.(*ExistHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistHistoryConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistHistoryCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistHistoryConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/ExistHistoryConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistHistoryConds(ctx, req.(*ExistHistoryCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountHistoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/CountHistories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountHistories(ctx, req.(*CountHistoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.login.history.v2.Manager/DeleteHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteHistory(ctx, req.(*DeleteHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.manager.login.history.v2.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHistory",
			Handler:    _Manager_CreateHistory_Handler,
		},
		{
			MethodName: "CreateHistories",
			Handler:    _Manager_CreateHistories_Handler,
		},
		{
			MethodName: "UpdateHistory",
			Handler:    _Manager_UpdateHistory_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _Manager_GetHistory_Handler,
		},
		{
			MethodName: "GetHistoryOnly",
			Handler:    _Manager_GetHistoryOnly_Handler,
		},
		{
			MethodName: "GetHistories",
			Handler:    _Manager_GetHistories_Handler,
		},
		{
			MethodName: "ExistHistory",
			Handler:    _Manager_ExistHistory_Handler,
		},
		{
			MethodName: "ExistHistoryConds",
			Handler:    _Manager_ExistHistoryConds_Handler,
		},
		{
			MethodName: "CountHistories",
			Handler:    _Manager_CountHistories_Handler,
		},
		{
			MethodName: "DeleteHistory",
			Handler:    _Manager_DeleteHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/login/history/history.proto",
}
