// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/fractionrule/fractionrule.proto

package fractionrule

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
	Middleware_CreateFractionRule_FullMethodName     = "/miningpool.middleware.fractionrule.v1.Middleware/CreateFractionRule"
	Middleware_GetFractionRule_FullMethodName        = "/miningpool.middleware.fractionrule.v1.Middleware/GetFractionRule"
	Middleware_GetFractionRules_FullMethodName       = "/miningpool.middleware.fractionrule.v1.Middleware/GetFractionRules"
	Middleware_ExistFractionRule_FullMethodName      = "/miningpool.middleware.fractionrule.v1.Middleware/ExistFractionRule"
	Middleware_ExistFractionRuleConds_FullMethodName = "/miningpool.middleware.fractionrule.v1.Middleware/ExistFractionRuleConds"
	Middleware_UpdateFractionRule_FullMethodName     = "/miningpool.middleware.fractionrule.v1.Middleware/UpdateFractionRule"
	Middleware_DeleteFractionRule_FullMethodName     = "/miningpool.middleware.fractionrule.v1.Middleware/DeleteFractionRule"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateFractionRule(ctx context.Context, in *CreateFractionRuleRequest, opts ...grpc.CallOption) (*CreateFractionRuleResponse, error)
	GetFractionRule(ctx context.Context, in *GetFractionRuleRequest, opts ...grpc.CallOption) (*GetFractionRuleResponse, error)
	GetFractionRules(ctx context.Context, in *GetFractionRulesRequest, opts ...grpc.CallOption) (*GetFractionRulesResponse, error)
	ExistFractionRule(ctx context.Context, in *ExistFractionRuleRequest, opts ...grpc.CallOption) (*ExistFractionRuleResponse, error)
	ExistFractionRuleConds(ctx context.Context, in *ExistFractionRuleCondsRequest, opts ...grpc.CallOption) (*ExistFractionRuleCondsResponse, error)
	UpdateFractionRule(ctx context.Context, in *UpdateFractionRuleRequest, opts ...grpc.CallOption) (*UpdateFractionRuleResponse, error)
	DeleteFractionRule(ctx context.Context, in *DeleteFractionRuleRequest, opts ...grpc.CallOption) (*DeleteFractionRuleResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFractionRule(ctx context.Context, in *CreateFractionRuleRequest, opts ...grpc.CallOption) (*CreateFractionRuleResponse, error) {
	out := new(CreateFractionRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFractionRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionRule(ctx context.Context, in *GetFractionRuleRequest, opts ...grpc.CallOption) (*GetFractionRuleResponse, error) {
	out := new(GetFractionRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFractionRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionRules(ctx context.Context, in *GetFractionRulesRequest, opts ...grpc.CallOption) (*GetFractionRulesResponse, error) {
	out := new(GetFractionRulesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFractionRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionRule(ctx context.Context, in *ExistFractionRuleRequest, opts ...grpc.CallOption) (*ExistFractionRuleResponse, error) {
	out := new(ExistFractionRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFractionRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionRuleConds(ctx context.Context, in *ExistFractionRuleCondsRequest, opts ...grpc.CallOption) (*ExistFractionRuleCondsResponse, error) {
	out := new(ExistFractionRuleCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFractionRuleConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateFractionRule(ctx context.Context, in *UpdateFractionRuleRequest, opts ...grpc.CallOption) (*UpdateFractionRuleResponse, error) {
	out := new(UpdateFractionRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateFractionRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFractionRule(ctx context.Context, in *DeleteFractionRuleRequest, opts ...grpc.CallOption) (*DeleteFractionRuleResponse, error) {
	out := new(DeleteFractionRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteFractionRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFractionRule(context.Context, *CreateFractionRuleRequest) (*CreateFractionRuleResponse, error)
	GetFractionRule(context.Context, *GetFractionRuleRequest) (*GetFractionRuleResponse, error)
	GetFractionRules(context.Context, *GetFractionRulesRequest) (*GetFractionRulesResponse, error)
	ExistFractionRule(context.Context, *ExistFractionRuleRequest) (*ExistFractionRuleResponse, error)
	ExistFractionRuleConds(context.Context, *ExistFractionRuleCondsRequest) (*ExistFractionRuleCondsResponse, error)
	UpdateFractionRule(context.Context, *UpdateFractionRuleRequest) (*UpdateFractionRuleResponse, error)
	DeleteFractionRule(context.Context, *DeleteFractionRuleRequest) (*DeleteFractionRuleResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFractionRule(context.Context, *CreateFractionRuleRequest) (*CreateFractionRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFractionRule not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionRule(context.Context, *GetFractionRuleRequest) (*GetFractionRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionRule not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionRules(context.Context, *GetFractionRulesRequest) (*GetFractionRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionRules not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionRule(context.Context, *ExistFractionRuleRequest) (*ExistFractionRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionRule not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionRuleConds(context.Context, *ExistFractionRuleCondsRequest) (*ExistFractionRuleCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionRuleConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateFractionRule(context.Context, *UpdateFractionRuleRequest) (*UpdateFractionRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFractionRule not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFractionRule(context.Context, *DeleteFractionRuleRequest) (*DeleteFractionRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFractionRule not implemented")
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

func _Middleware_CreateFractionRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFractionRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFractionRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFractionRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFractionRule(ctx, req.(*CreateFractionRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFractionRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionRule(ctx, req.(*GetFractionRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFractionRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionRules(ctx, req.(*GetFractionRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFractionRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionRule(ctx, req.(*ExistFractionRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionRuleConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionRuleCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionRuleConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFractionRuleConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionRuleConds(ctx, req.(*ExistFractionRuleCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateFractionRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFractionRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateFractionRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateFractionRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateFractionRule(ctx, req.(*UpdateFractionRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFractionRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFractionRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFractionRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteFractionRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFractionRule(ctx, req.(*DeleteFractionRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.fractionrule.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFractionRule",
			Handler:    _Middleware_CreateFractionRule_Handler,
		},
		{
			MethodName: "GetFractionRule",
			Handler:    _Middleware_GetFractionRule_Handler,
		},
		{
			MethodName: "GetFractionRules",
			Handler:    _Middleware_GetFractionRules_Handler,
		},
		{
			MethodName: "ExistFractionRule",
			Handler:    _Middleware_ExistFractionRule_Handler,
		},
		{
			MethodName: "ExistFractionRuleConds",
			Handler:    _Middleware_ExistFractionRuleConds_Handler,
		},
		{
			MethodName: "UpdateFractionRule",
			Handler:    _Middleware_UpdateFractionRule_Handler,
		},
		{
			MethodName: "DeleteFractionRule",
			Handler:    _Middleware_DeleteFractionRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/fractionrule/fractionrule.proto",
}
