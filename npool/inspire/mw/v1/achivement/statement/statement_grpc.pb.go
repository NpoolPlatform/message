// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mw/v1/achivement/statement/statement.proto

package statement

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
	Middleware_CreateStatement_FullMethodName  = "/inspire.middleware.achivement.statement.v1.Middleware/CreateStatement"
	Middleware_CreateStatements_FullMethodName = "/inspire.middleware.achivement.statement.v1.Middleware/CreateStatements"
	Middleware_GetStatements_FullMethodName    = "/inspire.middleware.achivement.statement.v1.Middleware/GetStatements"
	Middleware_DeleteStatement_FullMethodName  = "/inspire.middleware.achivement.statement.v1.Middleware/DeleteStatement"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateStatement(ctx context.Context, in *CreateStatementRequest, opts ...grpc.CallOption) (*CreateStatementResponse, error)
	CreateStatements(ctx context.Context, in *CreateStatementsRequest, opts ...grpc.CallOption) (*CreateStatementsResponse, error)
	GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error)
	DeleteStatement(ctx context.Context, in *DeleteStatementRequest, opts ...grpc.CallOption) (*DeleteStatementResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateStatement(ctx context.Context, in *CreateStatementRequest, opts ...grpc.CallOption) (*CreateStatementResponse, error) {
	out := new(CreateStatementResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateStatement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) CreateStatements(ctx context.Context, in *CreateStatementsRequest, opts ...grpc.CallOption) (*CreateStatementsResponse, error) {
	out := new(CreateStatementsResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateStatements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetStatements(ctx context.Context, in *GetStatementsRequest, opts ...grpc.CallOption) (*GetStatementsResponse, error) {
	out := new(GetStatementsResponse)
	err := c.cc.Invoke(ctx, Middleware_GetStatements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteStatement(ctx context.Context, in *DeleteStatementRequest, opts ...grpc.CallOption) (*DeleteStatementResponse, error) {
	out := new(DeleteStatementResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteStatement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateStatement(context.Context, *CreateStatementRequest) (*CreateStatementResponse, error)
	CreateStatements(context.Context, *CreateStatementsRequest) (*CreateStatementsResponse, error)
	GetStatements(context.Context, *GetStatementsRequest) (*GetStatementsResponse, error)
	DeleteStatement(context.Context, *DeleteStatementRequest) (*DeleteStatementResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateStatement(context.Context, *CreateStatementRequest) (*CreateStatementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatement not implemented")
}
func (UnimplementedMiddlewareServer) CreateStatements(context.Context, *CreateStatementsRequest) (*CreateStatementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatements not implemented")
}
func (UnimplementedMiddlewareServer) GetStatements(context.Context, *GetStatementsRequest) (*GetStatementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatements not implemented")
}
func (UnimplementedMiddlewareServer) DeleteStatement(context.Context, *DeleteStatementRequest) (*DeleteStatementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStatement not implemented")
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

func _Middleware_CreateStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateStatement(ctx, req.(*CreateStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_CreateStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateStatements(ctx, req.(*CreateStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetStatements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetStatements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetStatements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetStatements(ctx, req.(*GetStatementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteStatement(ctx, req.(*DeleteStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.middleware.achivement.statement.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStatement",
			Handler:    _Middleware_CreateStatement_Handler,
		},
		{
			MethodName: "CreateStatements",
			Handler:    _Middleware_CreateStatements_Handler,
		},
		{
			MethodName: "GetStatements",
			Handler:    _Middleware_GetStatements_Handler,
		},
		{
			MethodName: "DeleteStatement",
			Handler:    _Middleware_DeleteStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mw/v1/achivement/statement/statement.proto",
}
