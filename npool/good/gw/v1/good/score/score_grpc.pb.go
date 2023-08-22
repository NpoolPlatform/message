// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/good/gw/v1/good/score/score.proto

package score

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
	Gateway_CreateScore_FullMethodName = "/good.middleware.good1.score.v1.Gateway/CreateScore"
	Gateway_GetScores_FullMethodName   = "/good.middleware.good1.score.v1.Gateway/GetScores"
	Gateway_GetMyScores_FullMethodName = "/good.middleware.good1.score.v1.Gateway/GetMyScores"
	Gateway_DeleteScore_FullMethodName = "/good.middleware.good1.score.v1.Gateway/DeleteScore"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateScore(ctx context.Context, in *CreateScoreRequest, opts ...grpc.CallOption) (*CreateScoreResponse, error)
	GetScores(ctx context.Context, in *GetScoresRequest, opts ...grpc.CallOption) (*GetScoresResponse, error)
	GetMyScores(ctx context.Context, in *GetMyScoresRequest, opts ...grpc.CallOption) (*GetMyScoresResponse, error)
	DeleteScore(ctx context.Context, in *DeleteScoreRequest, opts ...grpc.CallOption) (*DeleteScoreResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateScore(ctx context.Context, in *CreateScoreRequest, opts ...grpc.CallOption) (*CreateScoreResponse, error) {
	out := new(CreateScoreResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetScores(ctx context.Context, in *GetScoresRequest, opts ...grpc.CallOption) (*GetScoresResponse, error) {
	out := new(GetScoresResponse)
	err := c.cc.Invoke(ctx, Gateway_GetScores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMyScores(ctx context.Context, in *GetMyScoresRequest, opts ...grpc.CallOption) (*GetMyScoresResponse, error) {
	out := new(GetMyScoresResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMyScores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteScore(ctx context.Context, in *DeleteScoreRequest, opts ...grpc.CallOption) (*DeleteScoreResponse, error) {
	out := new(DeleteScoreResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteScore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateScore(context.Context, *CreateScoreRequest) (*CreateScoreResponse, error)
	GetScores(context.Context, *GetScoresRequest) (*GetScoresResponse, error)
	GetMyScores(context.Context, *GetMyScoresRequest) (*GetMyScoresResponse, error)
	DeleteScore(context.Context, *DeleteScoreRequest) (*DeleteScoreResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateScore(context.Context, *CreateScoreRequest) (*CreateScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScore not implemented")
}
func (UnimplementedGatewayServer) GetScores(context.Context, *GetScoresRequest) (*GetScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScores not implemented")
}
func (UnimplementedGatewayServer) GetMyScores(context.Context, *GetMyScoresRequest) (*GetMyScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyScores not implemented")
}
func (UnimplementedGatewayServer) DeleteScore(context.Context, *DeleteScoreRequest) (*DeleteScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScore not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateScore(ctx, req.(*CreateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetScores(ctx, req.(*GetScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMyScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMyScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMyScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMyScores(ctx, req.(*GetMyScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteScore(ctx, req.(*DeleteScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.middleware.good1.score.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScore",
			Handler:    _Gateway_CreateScore_Handler,
		},
		{
			MethodName: "GetScores",
			Handler:    _Gateway_GetScores_Handler,
		},
		{
			MethodName: "GetMyScores",
			Handler:    _Gateway_GetMyScores_Handler,
		},
		{
			MethodName: "DeleteScore",
			Handler:    _Gateway_DeleteScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/gw/v1/good/score/score.proto",
}
