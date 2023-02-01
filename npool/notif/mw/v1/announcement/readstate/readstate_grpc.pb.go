// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/notif/mw/v1/announcement/readstate/readstate.proto

package readstate

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
	GetReadState(ctx context.Context, in *GetReadStateRequest, opts ...grpc.CallOption) (*GetReadStateResponse, error)
	GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) GetReadState(ctx context.Context, in *GetReadStateRequest, opts ...grpc.CallOption) (*GetReadStateResponse, error) {
	out := new(GetReadStateResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.readstate.v1.Middleware/GetReadState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetReadStates(ctx context.Context, in *GetReadStatesRequest, opts ...grpc.CallOption) (*GetReadStatesResponse, error) {
	out := new(GetReadStatesResponse)
	err := c.cc.Invoke(ctx, "/notif.middleware.announcement.readstate.v1.Middleware/GetReadStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	GetReadState(context.Context, *GetReadStateRequest) (*GetReadStateResponse, error)
	GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) GetReadState(context.Context, *GetReadStateRequest) (*GetReadStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadState not implemented")
}
func (UnimplementedMiddlewareServer) GetReadStates(context.Context, *GetReadStatesRequest) (*GetReadStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadStates not implemented")
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

func _Middleware_GetReadState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReadState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.readstate.v1.Middleware/GetReadState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReadState(ctx, req.(*GetReadStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetReadStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetReadStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.middleware.announcement.readstate.v1.Middleware/GetReadStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetReadStates(ctx, req.(*GetReadStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.middleware.announcement.readstate.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReadState",
			Handler:    _Middleware_GetReadState_Handler,
		},
		{
			MethodName: "GetReadStates",
			Handler:    _Middleware_GetReadStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/notif/mw/v1/announcement/readstate/readstate.proto",
}
