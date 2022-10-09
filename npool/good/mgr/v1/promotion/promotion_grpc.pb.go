// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mgr/v1/promotion/promotion.proto

package promotion

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
	CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error)
	CreatePromotions(ctx context.Context, in *CreatePromotionsRequest, opts ...grpc.CallOption) (*CreatePromotionsResponse, error)
	GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*GetPromotionResponse, error)
	GetPromotionOnly(ctx context.Context, in *GetPromotionOnlyRequest, opts ...grpc.CallOption) (*GetPromotionOnlyResponse, error)
	GetPromotions(ctx context.Context, in *GetPromotionsRequest, opts ...grpc.CallOption) (*GetPromotionsResponse, error)
	ExistPromotion(ctx context.Context, in *ExistPromotionRequest, opts ...grpc.CallOption) (*ExistPromotionResponse, error)
	ExistPromotionConds(ctx context.Context, in *ExistPromotionCondsRequest, opts ...grpc.CallOption) (*ExistPromotionCondsResponse, error)
	CountPromotions(ctx context.Context, in *CountPromotionsRequest, opts ...grpc.CallOption) (*CountPromotionsResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error) {
	out := new(CreatePromotionResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/CreatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreatePromotions(ctx context.Context, in *CreatePromotionsRequest, opts ...grpc.CallOption) (*CreatePromotionsResponse, error) {
	out := new(CreatePromotionsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/CreatePromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*GetPromotionResponse, error) {
	out := new(GetPromotionResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/GetPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPromotionOnly(ctx context.Context, in *GetPromotionOnlyRequest, opts ...grpc.CallOption) (*GetPromotionOnlyResponse, error) {
	out := new(GetPromotionOnlyResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/GetPromotionOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPromotions(ctx context.Context, in *GetPromotionsRequest, opts ...grpc.CallOption) (*GetPromotionsResponse, error) {
	out := new(GetPromotionsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/GetPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistPromotion(ctx context.Context, in *ExistPromotionRequest, opts ...grpc.CallOption) (*ExistPromotionResponse, error) {
	out := new(ExistPromotionResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/ExistPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistPromotionConds(ctx context.Context, in *ExistPromotionCondsRequest, opts ...grpc.CallOption) (*ExistPromotionCondsResponse, error) {
	out := new(ExistPromotionCondsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/ExistPromotionConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountPromotions(ctx context.Context, in *CountPromotionsRequest, opts ...grpc.CallOption) (*CountPromotionsResponse, error) {
	out := new(CountPromotionsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.promotion.v1.Manager/CountPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error)
	CreatePromotions(context.Context, *CreatePromotionsRequest) (*CreatePromotionsResponse, error)
	GetPromotion(context.Context, *GetPromotionRequest) (*GetPromotionResponse, error)
	GetPromotionOnly(context.Context, *GetPromotionOnlyRequest) (*GetPromotionOnlyResponse, error)
	GetPromotions(context.Context, *GetPromotionsRequest) (*GetPromotionsResponse, error)
	ExistPromotion(context.Context, *ExistPromotionRequest) (*ExistPromotionResponse, error)
	ExistPromotionConds(context.Context, *ExistPromotionCondsRequest) (*ExistPromotionCondsResponse, error)
	CountPromotions(context.Context, *CountPromotionsRequest) (*CountPromotionsResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotion not implemented")
}
func (UnimplementedManagerServer) CreatePromotions(context.Context, *CreatePromotionsRequest) (*CreatePromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotions not implemented")
}
func (UnimplementedManagerServer) GetPromotion(context.Context, *GetPromotionRequest) (*GetPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotion not implemented")
}
func (UnimplementedManagerServer) GetPromotionOnly(context.Context, *GetPromotionOnlyRequest) (*GetPromotionOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotionOnly not implemented")
}
func (UnimplementedManagerServer) GetPromotions(context.Context, *GetPromotionsRequest) (*GetPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotions not implemented")
}
func (UnimplementedManagerServer) ExistPromotion(context.Context, *ExistPromotionRequest) (*ExistPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPromotion not implemented")
}
func (UnimplementedManagerServer) ExistPromotionConds(context.Context, *ExistPromotionCondsRequest) (*ExistPromotionCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistPromotionConds not implemented")
}
func (UnimplementedManagerServer) CountPromotions(context.Context, *CountPromotionsRequest) (*CountPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPromotions not implemented")
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

func _Manager_CreatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/CreatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreatePromotion(ctx, req.(*CreatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreatePromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreatePromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/CreatePromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreatePromotions(ctx, req.(*CreatePromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/GetPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPromotion(ctx, req.(*GetPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPromotionOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPromotionOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/GetPromotionOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPromotionOnly(ctx, req.(*GetPromotionOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/GetPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPromotions(ctx, req.(*GetPromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/ExistPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistPromotion(ctx, req.(*ExistPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistPromotionConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistPromotionCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistPromotionConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/ExistPromotionConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistPromotionConds(ctx, req.(*ExistPromotionCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPromotionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.promotion.v1.Manager/CountPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountPromotions(ctx, req.(*CountPromotionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.manager.promotion.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePromotion",
			Handler:    _Manager_CreatePromotion_Handler,
		},
		{
			MethodName: "CreatePromotions",
			Handler:    _Manager_CreatePromotions_Handler,
		},
		{
			MethodName: "GetPromotion",
			Handler:    _Manager_GetPromotion_Handler,
		},
		{
			MethodName: "GetPromotionOnly",
			Handler:    _Manager_GetPromotionOnly_Handler,
		},
		{
			MethodName: "GetPromotions",
			Handler:    _Manager_GetPromotions_Handler,
		},
		{
			MethodName: "ExistPromotion",
			Handler:    _Manager_ExistPromotion_Handler,
		},
		{
			MethodName: "ExistPromotionConds",
			Handler:    _Manager_ExistPromotionConds_Handler,
		},
		{
			MethodName: "CountPromotions",
			Handler:    _Manager_CountPromotions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mgr/v1/promotion/promotion.proto",
}
