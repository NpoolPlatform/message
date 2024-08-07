// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/achievement/good/achievement.proto

package good

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
	Middleware_GetAchievements_FullMethodName   = "/inspire.middleware.achievement.good.v1.Middleware/GetAchievements"
	Middleware_DeleteAchievement_FullMethodName = "/inspire.middleware.achievement.good.v1.Middleware/DeleteAchievement"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	GetAchievements(ctx context.Context, in *GetAchievementsRequest, opts ...grpc.CallOption) (*GetAchievementsResponse, error)
	DeleteAchievement(ctx context.Context, in *DeleteAchievementRequest, opts ...grpc.CallOption) (*DeleteAchievementResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetAchievements(ctx context.Context, in *GetAchievementsRequest, opts ...grpc.CallOption) (*GetAchievementsResponse, error) {
	out := new(GetAchievementsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetAchievements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteAchievement(ctx context.Context, in *DeleteAchievementRequest, opts ...grpc.CallOption) (*DeleteAchievementResponse, error) {
	out := new(DeleteAchievementResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteAchievement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetAchievements(context.Context, *GetAchievementsRequest) (*GetAchievementsResponse, error)
	DeleteAchievement(context.Context, *DeleteAchievementRequest) (*DeleteAchievementResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetAchievements(context.Context, *GetAchievementsRequest) (*GetAchievementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAchievements not implemented")
}
func (UnimplementedMiddlewareServer) DeleteAchievement(context.Context, *DeleteAchievementRequest) (*DeleteAchievementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAchievement not implemented")
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

func _Middleware_GetAchievements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAchievementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetAchievements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetAchievements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetAchievements(ctx, req.(*GetAchievementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteAchievement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAchievementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteAchievement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteAchievement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteAchievement(ctx, req.(*DeleteAchievementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.achievement.good.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAchievements",
			Handler:    _Middleware_GetAchievements_Handler,
		},
		{
			MethodName: "DeleteAchievement",
			Handler:    _Middleware_DeleteAchievement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/achievement/good/achievement.proto",
}
