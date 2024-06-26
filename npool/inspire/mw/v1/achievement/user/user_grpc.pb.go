// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/achievement/user/user.proto

package user

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
	Middleware_GetAchievementUsers_FullMethodName   = "/inspire.middleware.achievement.user.v1.Middleware/GetAchievementUsers"
	Middleware_DeleteAchievementUser_FullMethodName = "/inspire.middleware.achievement.user.v1.Middleware/DeleteAchievementUser"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetAchievementUsers(ctx context.Context, in *GetAchievementUsersRequest, opts ...grpc.CallOption) (*GetAchievementUsersResponse, error)
	DeleteAchievementUser(ctx context.Context, in *DeleteAchievementUserRequest, opts ...grpc.CallOption) (*DeleteAchievementUserResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetAchievementUsers(ctx context.Context, in *GetAchievementUsersRequest, opts ...grpc.CallOption) (*GetAchievementUsersResponse, error) {
	out := new(GetAchievementUsersResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAchievementUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAchievementUser(ctx context.Context, in *DeleteAchievementUserRequest, opts ...grpc.CallOption) (*DeleteAchievementUserResponse, error) {
	out := new(DeleteAchievementUserResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteAchievementUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetAchievementUsers(context.Context, *GetAchievementUsersRequest) (*GetAchievementUsersResponse, error)
	DeleteAchievementUser(context.Context, *DeleteAchievementUserRequest) (*DeleteAchievementUserResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetAchievementUsers(context.Context, *GetAchievementUsersRequest) (*GetAchievementUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAchievementUsers not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAchievementUser(context.Context, *DeleteAchievementUserRequest) (*DeleteAchievementUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAchievementUser not implemented")
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

func _Middleware_GetAchievementUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAchievementUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAchievementUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAchievementUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAchievementUsers(ctx, req.(*GetAchievementUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAchievementUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAchievementUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAchievementUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteAchievementUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAchievementUser(ctx, req.(*DeleteAchievementUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.achievement.user.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAchievementUsers",
			Handler:    _Middleware_GetAchievementUsers_Handler,
		},
		{
			MethodName: "DeleteAchievementUser",
			Handler:    _Middleware_DeleteAchievementUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/achievement/user/user.proto",
}
