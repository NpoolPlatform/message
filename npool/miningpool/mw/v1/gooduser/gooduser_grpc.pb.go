// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/gooduser/gooduser.proto

package gooduser

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
	Middleware_CreateGoodUser_FullMethodName     = "/miningpool.middleware.gooduser.v1.Middleware/CreateGoodUser"
	Middleware_CreateGoodUsers_FullMethodName    = "/miningpool.middleware.gooduser.v1.Middleware/CreateGoodUsers"
	Middleware_GetGoodUser_FullMethodName        = "/miningpool.middleware.gooduser.v1.Middleware/GetGoodUser"
	Middleware_GetGoodUsers_FullMethodName       = "/miningpool.middleware.gooduser.v1.Middleware/GetGoodUsers"
	Middleware_ExistGoodUserConds_FullMethodName = "/miningpool.middleware.gooduser.v1.Middleware/ExistGoodUserConds"
	Middleware_UpdateGoodUser_FullMethodName     = "/miningpool.middleware.gooduser.v1.Middleware/UpdateGoodUser"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateGoodUser(ctx context.Context, in *CreateGoodUserRequest, opts ...grpc.CallOption) (*CreateGoodUserResponse, error)
	CreateGoodUsers(ctx context.Context, in *CreateGoodUsersRequest, opts ...grpc.CallOption) (*CreateGoodUsersResponse, error)
	GetGoodUser(ctx context.Context, in *GetGoodUserRequest, opts ...grpc.CallOption) (*GetGoodUserResponse, error)
	GetGoodUsers(ctx context.Context, in *GetGoodUsersRequest, opts ...grpc.CallOption) (*GetGoodUsersResponse, error)
	ExistGoodUserConds(ctx context.Context, in *ExistGoodUserCondsRequest, opts ...grpc.CallOption) (*ExistGoodUserCondsResponse, error)
	UpdateGoodUser(ctx context.Context, in *UpdateGoodUserRequest, opts ...grpc.CallOption) (*UpdateGoodUserResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateGoodUser(ctx context.Context, in *CreateGoodUserRequest, opts ...grpc.CallOption) (*CreateGoodUserResponse, error) {
	out := new(CreateGoodUserResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateGoodUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateGoodUsers(ctx context.Context, in *CreateGoodUsersRequest, opts ...grpc.CallOption) (*CreateGoodUsersResponse, error) {
	out := new(CreateGoodUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateGoodUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetGoodUser(ctx context.Context, in *GetGoodUserRequest, opts ...grpc.CallOption) (*GetGoodUserResponse, error) {
	out := new(GetGoodUserResponse)
	err := c.cc.Invoke(ctx, Middleware_GetGoodUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetGoodUsers(ctx context.Context, in *GetGoodUsersRequest, opts ...grpc.CallOption) (*GetGoodUsersResponse, error) {
	out := new(GetGoodUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetGoodUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistGoodUserConds(ctx context.Context, in *ExistGoodUserCondsRequest, opts ...grpc.CallOption) (*ExistGoodUserCondsResponse, error) {
	out := new(ExistGoodUserCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistGoodUserConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateGoodUser(ctx context.Context, in *UpdateGoodUserRequest, opts ...grpc.CallOption) (*UpdateGoodUserResponse, error) {
	out := new(UpdateGoodUserResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateGoodUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateGoodUser(context.Context, *CreateGoodUserRequest) (*CreateGoodUserResponse, error)
	CreateGoodUsers(context.Context, *CreateGoodUsersRequest) (*CreateGoodUsersResponse, error)
	GetGoodUser(context.Context, *GetGoodUserRequest) (*GetGoodUserResponse, error)
	GetGoodUsers(context.Context, *GetGoodUsersRequest) (*GetGoodUsersResponse, error)
	ExistGoodUserConds(context.Context, *ExistGoodUserCondsRequest) (*ExistGoodUserCondsResponse, error)
	UpdateGoodUser(context.Context, *UpdateGoodUserRequest) (*UpdateGoodUserResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateGoodUser(context.Context, *CreateGoodUserRequest) (*CreateGoodUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodUser not implemented")
}
func (UnimplementedMiddlewareServer) CreateGoodUsers(context.Context, *CreateGoodUsersRequest) (*CreateGoodUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodUsers not implemented")
}
func (UnimplementedMiddlewareServer) GetGoodUser(context.Context, *GetGoodUserRequest) (*GetGoodUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodUser not implemented")
}
func (UnimplementedMiddlewareServer) GetGoodUsers(context.Context, *GetGoodUsersRequest) (*GetGoodUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodUsers not implemented")
}
func (UnimplementedMiddlewareServer) ExistGoodUserConds(context.Context, *ExistGoodUserCondsRequest) (*ExistGoodUserCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistGoodUserConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateGoodUser(context.Context, *UpdateGoodUserRequest) (*UpdateGoodUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodUser not implemented")
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

func _Middleware_CreateGoodUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateGoodUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateGoodUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateGoodUser(ctx, req.(*CreateGoodUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateGoodUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateGoodUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateGoodUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateGoodUsers(ctx, req.(*CreateGoodUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetGoodUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGoodUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetGoodUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGoodUser(ctx, req.(*GetGoodUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetGoodUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetGoodUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetGoodUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetGoodUsers(ctx, req.(*GetGoodUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistGoodUserConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistGoodUserCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistGoodUserConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistGoodUserConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistGoodUserConds(ctx, req.(*ExistGoodUserCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateGoodUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateGoodUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateGoodUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateGoodUser(ctx, req.(*UpdateGoodUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.gooduser.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoodUser",
			Handler:    _Middleware_CreateGoodUser_Handler,
		},
		{
			MethodName: "CreateGoodUsers",
			Handler:    _Middleware_CreateGoodUsers_Handler,
		},
		{
			MethodName: "GetGoodUser",
			Handler:    _Middleware_GetGoodUser_Handler,
		},
		{
			MethodName: "GetGoodUsers",
			Handler:    _Middleware_GetGoodUsers_Handler,
		},
		{
			MethodName: "ExistGoodUserConds",
			Handler:    _Middleware_ExistGoodUserConds_Handler,
		},
		{
			MethodName: "UpdateGoodUser",
			Handler:    _Middleware_UpdateGoodUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/gooduser/gooduser.proto",
}
