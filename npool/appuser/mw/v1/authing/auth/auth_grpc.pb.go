// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
<<<<<<< HEAD
// source: npool/appuser/mw/v1/authing/auth/auth.proto

package auth
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
// source: npool/appuser/mw/v1/user/login/history/history.proto

package history
=======
// source: npool/appuser/mw/v1/authing/auth/auth.proto

package auth
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing

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

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
<<<<<<< HEAD
	CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error)
	CreateAuths(ctx context.Context, in *CreateAuthsRequest, opts ...grpc.CallOption) (*CreateAuthsResponse, error)
	ExistAuth(ctx context.Context, in *ExistAuthRequest, opts ...grpc.CallOption) (*ExistAuthResponse, error)
	GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error)
	GetAuths(ctx context.Context, in *GetAuthsRequest, opts ...grpc.CallOption) (*GetAuthsResponse, error)
	ExistAuthConds(ctx context.Context, in *ExistAuthCondsRequest, opts ...grpc.CallOption) (*ExistAuthCondsResponse, error)
	DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error)
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
	CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error)
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error)
	GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error)
=======
	CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error)
	UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error)
	ExistAuth(ctx context.Context, in *ExistAuthRequest, opts ...grpc.CallOption) (*ExistAuthResponse, error)
	GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error)
	GetAuths(ctx context.Context, in *GetAuthsRequest, opts ...grpc.CallOption) (*GetAuthsResponse, error)
	DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error)
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

<<<<<<< HEAD
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
func (c *middlewareClient) CreateHistory(ctx context.Context, in *CreateHistoryRequest, opts ...grpc.CallOption) (*CreateHistoryResponse, error) {
	out := new(CreateHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.user.login.history.v1.Middleware/CreateHistory", in, out, opts...)
=======
>>>>>>> Refactor authing
func (c *middlewareClient) CreateAuth(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error) {
	out := new(CreateAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/CreateAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) CreateAuths(ctx context.Context, in *CreateAuthsRequest, opts ...grpc.CallOption) (*CreateAuthsResponse, error) {
	out := new(CreateAuthsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/CreateAuths", in, out, opts...)
=======
func (c *middlewareClient) UpdateAuth(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error) {
	out := new(UpdateAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/UpdateAuth", in, out, opts...)
>>>>>>> Refactor authing
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistAuth(ctx context.Context, in *ExistAuthRequest, opts ...grpc.CallOption) (*ExistAuthResponse, error) {
	out := new(ExistAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/ExistAuth", in, out, opts...)
<<<<<<< HEAD
=======
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
func (c *middlewareClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error) {
	out := new(GetHistoryResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.user.login.history.v1.Middleware/GetHistory", in, out, opts...)
=======
>>>>>>> Refactor authing
func (c *middlewareClient) GetAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error) {
	out := new(GetAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/GetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetAuths(ctx context.Context, in *GetAuthsRequest, opts ...grpc.CallOption) (*GetAuthsResponse, error) {
	out := new(GetAuthsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/GetAuths", in, out, opts...)
<<<<<<< HEAD
=======
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *middlewareClient) ExistAuthConds(ctx context.Context, in *ExistAuthCondsRequest, opts ...grpc.CallOption) (*ExistAuthCondsResponse, error) {
	out := new(ExistAuthCondsResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/ExistAuthConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	out := new(DeleteAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth", in, out, opts...)
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
func (c *middlewareClient) GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error) {
	out := new(GetHistoriesResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.user.login.history.v1.Middleware/GetHistories", in, out, opts...)
=======
func (c *middlewareClient) DeleteAuth(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	out := new(DeleteAuthResponse)
	err := c.cc.Invoke(ctx, "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth", in, out, opts...)
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
<<<<<<< HEAD
	CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error)
	CreateAuths(context.Context, *CreateAuthsRequest) (*CreateAuthsResponse, error)
	ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error)
	GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error)
	GetAuths(context.Context, *GetAuthsRequest) (*GetAuthsResponse, error)
	ExistAuthConds(context.Context, *ExistAuthCondsRequest) (*ExistAuthCondsResponse, error)
	DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error)
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
	CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error)
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	GetHistories(context.Context, *GetHistoriesRequest) (*GetHistoriesResponse, error)
=======
	CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error)
	UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error)
	ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error)
	GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error)
	GetAuths(context.Context, *GetAuthsRequest) (*GetAuthsResponse, error)
	DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error)
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

<<<<<<< HEAD
func (UnimplementedMiddlewareServer) CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuth not implemented")
}
func (UnimplementedMiddlewareServer) CreateAuths(context.Context, *CreateAuthsRequest) (*CreateAuthsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuths not implemented")
}
func (UnimplementedMiddlewareServer) ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAuth not implemented")
}
func (UnimplementedMiddlewareServer) GetAuth(context.Context, *GetAuthRequest) (*GetAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedMiddlewareServer) GetAuths(context.Context, *GetAuthsRequest) (*GetAuthsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuths not implemented")
}
func (UnimplementedMiddlewareServer) ExistAuthConds(context.Context, *ExistAuthCondsRequest) (*ExistAuthCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAuthConds not implemented")
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
func (UnimplementedMiddlewareServer) CreateHistory(context.Context, *CreateHistoryRequest) (*CreateHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistory not implemented")
=======
func (UnimplementedMiddlewareServer) CreateAuth(context.Context, *CreateAuthRequest) (*CreateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuth not implemented")
}
func (UnimplementedMiddlewareServer) UpdateAuth(context.Context, *UpdateAuthRequest) (*UpdateAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuth not implemented")
}
func (UnimplementedMiddlewareServer) ExistAuth(context.Context, *ExistAuthRequest) (*ExistAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAuth not implemented")
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
}
func (UnimplementedMiddlewareServer) GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
>>>>>>> Refactor authing
}
func (UnimplementedMiddlewareServer) DeleteAuth(context.Context, *DeleteAuthRequest) (*DeleteAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
}
func (UnimplementedMiddlewareServer) mustEmbedUnimplementedMiddlewareServer() {}

// UnsafeMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiddlewareServer will
// result in compilation errors.
type UnsafeMiddlewareServer interface {
	mustEmbedUnimplementedMiddlewareServer()
}

func RegisterMiddlewareServer(s grpc.ServiceRegistrar, srv MiddlewareServer) {
	s.RegisterService(&Middleware_ServiceDesc, srv)
}

<<<<<<< HEAD
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
func _Middleware_CreateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryRequest)
=======
>>>>>>> Refactor authing
func _Middleware_CreateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/CreateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAuth(ctx, req.(*CreateAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
func _Middleware_CreateAuths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthsRequest)
=======
func _Middleware_UpdateAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthRequest)
>>>>>>> Refactor authing
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).CreateAuths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/CreateAuths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateAuths(ctx, req.(*CreateAuthsRequest))
=======
		return srv.(MiddlewareServer).UpdateAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/UpdateAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateAuth(ctx, req.(*UpdateAuthRequest))
>>>>>>> Refactor authing
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/ExistAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAuth(ctx, req.(*ExistAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthRequest)
<<<<<<< HEAD
=======
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAuth(ctx, req.(*GetAuthRequest))
=======
		return srv.(MiddlewareServer).CreateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
		FullMethod: "/appuser.middleware.user.login.history.v1.Middleware/CreateHistory",
=======
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuth",
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateHistory(ctx, req.(*CreateHistoryRequest))
>>>>>>> Refactor authing
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
func _Middleware_GetAuths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthsRequest)
=======
func _Middleware_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
>>>>>>> Refactor authing
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(MiddlewareServer).GetAuths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuths",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAuths(ctx, req.(*GetAuthsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistAuthConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAuthCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistAuthConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/ExistAuthConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistAuthConds(ctx, req.(*ExistAuthCondsRequest))
=======
		return srv.(MiddlewareServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
		FullMethod: "/appuser.middleware.user.login.history.v1.Middleware/GetHistory",
=======
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/GetAuths",
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetHistory(ctx, req.(*GetHistoryRequest))
>>>>>>> Refactor authing
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
<<<<<<< HEAD
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth",
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
		FullMethod: "/appuser.middleware.user.login.history.v1.Middleware/GetHistories",
=======
		FullMethod: "/appuser.middleware.authing.auth.v1.Middleware/DeleteAuth",
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAuth(ctx, req.(*DeleteAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
<<<<<<< HEAD
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
	ServiceName: "appuser.middleware.user.login.history.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHistory",
			Handler:    _Middleware_CreateHistory_Handler,
=======
>>>>>>> Refactor authing
	ServiceName: "appuser.middleware.authing.auth.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuth",
			Handler:    _Middleware_CreateAuth_Handler,
		},
		{
<<<<<<< HEAD
			MethodName: "CreateAuths",
			Handler:    _Middleware_CreateAuths_Handler,
=======
			MethodName: "UpdateAuth",
			Handler:    _Middleware_UpdateAuth_Handler,
>>>>>>> Refactor authing
		},
		{
			MethodName: "ExistAuth",
			Handler:    _Middleware_ExistAuth_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _Middleware_GetAuth_Handler,
<<<<<<< HEAD
		},
		{
			MethodName: "GetAuths",
			Handler:    _Middleware_GetAuths_Handler,
		},
		{
			MethodName: "ExistAuthConds",
			Handler:    _Middleware_ExistAuthConds_Handler,
=======
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
		},
		{
			MethodName: "GetHistory",
			Handler:    _Middleware_GetHistory_Handler,
>>>>>>> Refactor authing
		},
		{
			MethodName: "DeleteAuth",
			Handler:    _Middleware_DeleteAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
<<<<<<< HEAD
	Metadata: "npool/appuser/mw/v1/authing/auth/auth.proto",
=======
<<<<<<< HEAD:npool/appuser/mw/v1/user/login/history/history_grpc.pb.go
	Metadata: "npool/appuser/mw/v1/user/login/history/history.proto",
=======
	Metadata: "npool/appuser/mw/v1/authing/auth/auth.proto",
>>>>>>> Refactor authing:npool/appuser/mw/v1/authing/auth/auth_grpc.pb.go
>>>>>>> Refactor authing
}
