// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/miningpool/mw/v1/fractionwithdrawalrule/fractionwithdrawalrule.proto

package fractionwithdrawalrule

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
	Middleware_CreateFractionWithdrawalRule_FullMethodName     = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/CreateFractionWithdrawalRule"
	Middleware_GetFractionWithdrawalRule_FullMethodName        = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/GetFractionWithdrawalRule"
	Middleware_GetFractionWithdrawalRules_FullMethodName       = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/GetFractionWithdrawalRules"
	Middleware_ExistFractionWithdrawalRule_FullMethodName      = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/ExistFractionWithdrawalRule"
	Middleware_ExistFractionWithdrawalRuleConds_FullMethodName = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/ExistFractionWithdrawalRuleConds"
	Middleware_UpdateFractionWithdrawalRule_FullMethodName     = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/UpdateFractionWithdrawalRule"
	Middleware_DeleteFractionWithdrawalRule_FullMethodName     = "/miningpool.middleware.fractionwithdrawalrule.v1.Middleware/DeleteFractionWithdrawalRule"
)

// MiddlewareClient is the client API for Middleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MiddlewareClient interface {
	CreateFractionWithdrawalRule(ctx context.Context, in *CreateFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalRuleResponse, error)
	GetFractionWithdrawalRule(ctx context.Context, in *GetFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalRuleResponse, error)
	GetFractionWithdrawalRules(ctx context.Context, in *GetFractionWithdrawalRulesRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalRulesResponse, error)
	ExistFractionWithdrawalRule(ctx context.Context, in *ExistFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalRuleResponse, error)
	ExistFractionWithdrawalRuleConds(ctx context.Context, in *ExistFractionWithdrawalRuleCondsRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalRuleCondsResponse, error)
	UpdateFractionWithdrawalRule(ctx context.Context, in *UpdateFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*UpdateFractionWithdrawalRuleResponse, error)
	DeleteFractionWithdrawalRule(ctx context.Context, in *DeleteFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*DeleteFractionWithdrawalRuleResponse, error)
}

type middlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewMiddlewareClient(cc grpc.ClientConnInterface) MiddlewareClient {
	return &middlewareClient{cc}
}

func (c *middlewareClient) CreateFractionWithdrawalRule(ctx context.Context, in *CreateFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*CreateFractionWithdrawalRuleResponse, error) {
	out := new(CreateFractionWithdrawalRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_CreateFractionWithdrawalRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionWithdrawalRule(ctx context.Context, in *GetFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalRuleResponse, error) {
	out := new(GetFractionWithdrawalRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFractionWithdrawalRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) GetFractionWithdrawalRules(ctx context.Context, in *GetFractionWithdrawalRulesRequest, opts ...grpc.CallOption) (*GetFractionWithdrawalRulesResponse, error) {
	out := new(GetFractionWithdrawalRulesResponse)
	err := c.cc.Invoke(ctx, Middleware_GetFractionWithdrawalRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionWithdrawalRule(ctx context.Context, in *ExistFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalRuleResponse, error) {
	out := new(ExistFractionWithdrawalRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFractionWithdrawalRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) ExistFractionWithdrawalRuleConds(ctx context.Context, in *ExistFractionWithdrawalRuleCondsRequest, opts ...grpc.CallOption) (*ExistFractionWithdrawalRuleCondsResponse, error) {
	out := new(ExistFractionWithdrawalRuleCondsResponse)
	err := c.cc.Invoke(ctx, Middleware_ExistFractionWithdrawalRuleConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) UpdateFractionWithdrawalRule(ctx context.Context, in *UpdateFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*UpdateFractionWithdrawalRuleResponse, error) {
	out := new(UpdateFractionWithdrawalRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_UpdateFractionWithdrawalRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareClient) DeleteFractionWithdrawalRule(ctx context.Context, in *DeleteFractionWithdrawalRuleRequest, opts ...grpc.CallOption) (*DeleteFractionWithdrawalRuleResponse, error) {
	out := new(DeleteFractionWithdrawalRuleResponse)
	err := c.cc.Invoke(ctx, Middleware_DeleteFractionWithdrawalRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddlewareServer is the server API for Middleware service.
// All implementations must embed UnimplementedMiddlewareServer
// for forward compatibility
type MiddlewareServer interface {
	CreateFractionWithdrawalRule(context.Context, *CreateFractionWithdrawalRuleRequest) (*CreateFractionWithdrawalRuleResponse, error)
	GetFractionWithdrawalRule(context.Context, *GetFractionWithdrawalRuleRequest) (*GetFractionWithdrawalRuleResponse, error)
	GetFractionWithdrawalRules(context.Context, *GetFractionWithdrawalRulesRequest) (*GetFractionWithdrawalRulesResponse, error)
	ExistFractionWithdrawalRule(context.Context, *ExistFractionWithdrawalRuleRequest) (*ExistFractionWithdrawalRuleResponse, error)
	ExistFractionWithdrawalRuleConds(context.Context, *ExistFractionWithdrawalRuleCondsRequest) (*ExistFractionWithdrawalRuleCondsResponse, error)
	UpdateFractionWithdrawalRule(context.Context, *UpdateFractionWithdrawalRuleRequest) (*UpdateFractionWithdrawalRuleResponse, error)
	DeleteFractionWithdrawalRule(context.Context, *DeleteFractionWithdrawalRuleRequest) (*DeleteFractionWithdrawalRuleResponse, error)
	mustEmbedUnimplementedMiddlewareServer()
}

// UnimplementedMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedMiddlewareServer struct {
}

func (UnimplementedMiddlewareServer) CreateFractionWithdrawalRule(context.Context, *CreateFractionWithdrawalRuleRequest) (*CreateFractionWithdrawalRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFractionWithdrawalRule not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionWithdrawalRule(context.Context, *GetFractionWithdrawalRuleRequest) (*GetFractionWithdrawalRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionWithdrawalRule not implemented")
}
func (UnimplementedMiddlewareServer) GetFractionWithdrawalRules(context.Context, *GetFractionWithdrawalRulesRequest) (*GetFractionWithdrawalRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFractionWithdrawalRules not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionWithdrawalRule(context.Context, *ExistFractionWithdrawalRuleRequest) (*ExistFractionWithdrawalRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionWithdrawalRule not implemented")
}
func (UnimplementedMiddlewareServer) ExistFractionWithdrawalRuleConds(context.Context, *ExistFractionWithdrawalRuleCondsRequest) (*ExistFractionWithdrawalRuleCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistFractionWithdrawalRuleConds not implemented")
}
func (UnimplementedMiddlewareServer) UpdateFractionWithdrawalRule(context.Context, *UpdateFractionWithdrawalRuleRequest) (*UpdateFractionWithdrawalRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFractionWithdrawalRule not implemented")
}
func (UnimplementedMiddlewareServer) DeleteFractionWithdrawalRule(context.Context, *DeleteFractionWithdrawalRuleRequest) (*DeleteFractionWithdrawalRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFractionWithdrawalRule not implemented")
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

func _Middleware_CreateFractionWithdrawalRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFractionWithdrawalRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).CreateFractionWithdrawalRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_CreateFractionWithdrawalRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).CreateFractionWithdrawalRule(ctx, req.(*CreateFractionWithdrawalRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionWithdrawalRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionWithdrawalRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionWithdrawalRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFractionWithdrawalRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionWithdrawalRule(ctx, req.(*GetFractionWithdrawalRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_GetFractionWithdrawalRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFractionWithdrawalRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).GetFractionWithdrawalRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_GetFractionWithdrawalRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).GetFractionWithdrawalRules(ctx, req.(*GetFractionWithdrawalRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionWithdrawalRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionWithdrawalRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionWithdrawalRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFractionWithdrawalRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionWithdrawalRule(ctx, req.(*ExistFractionWithdrawalRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_ExistFractionWithdrawalRuleConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistFractionWithdrawalRuleCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).ExistFractionWithdrawalRuleConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_ExistFractionWithdrawalRuleConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).ExistFractionWithdrawalRuleConds(ctx, req.(*ExistFractionWithdrawalRuleCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_UpdateFractionWithdrawalRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFractionWithdrawalRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).UpdateFractionWithdrawalRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_UpdateFractionWithdrawalRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).UpdateFractionWithdrawalRule(ctx, req.(*UpdateFractionWithdrawalRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middleware_DeleteFractionWithdrawalRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFractionWithdrawalRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddlewareServer).DeleteFractionWithdrawalRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Middleware_DeleteFractionWithdrawalRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddlewareServer).DeleteFractionWithdrawalRule(ctx, req.(*DeleteFractionWithdrawalRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Middleware_ServiceDesc is the grpc.ServiceDesc for Middleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Middleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.middleware.fractionwithdrawalrule.v1.Middleware",
	HandlerType: (*MiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFractionWithdrawalRule",
			Handler:    _Middleware_CreateFractionWithdrawalRule_Handler,
		},
		{
			MethodName: "GetFractionWithdrawalRule",
			Handler:    _Middleware_GetFractionWithdrawalRule_Handler,
		},
		{
			MethodName: "GetFractionWithdrawalRules",
			Handler:    _Middleware_GetFractionWithdrawalRules_Handler,
		},
		{
			MethodName: "ExistFractionWithdrawalRule",
			Handler:    _Middleware_ExistFractionWithdrawalRule_Handler,
		},
		{
			MethodName: "ExistFractionWithdrawalRuleConds",
			Handler:    _Middleware_ExistFractionWithdrawalRuleConds_Handler,
		},
		{
			MethodName: "UpdateFractionWithdrawalRule",
			Handler:    _Middleware_UpdateFractionWithdrawalRule_Handler,
		},
		{
			MethodName: "DeleteFractionWithdrawalRule",
			Handler:    _Middleware_DeleteFractionWithdrawalRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/miningpool/mw/v1/fractionwithdrawalrule/fractionwithdrawalrule.proto",
}
