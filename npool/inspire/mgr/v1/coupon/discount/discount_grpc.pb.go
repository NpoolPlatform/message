// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/mgr/v1/coupon/discount/discount.proto

package discount

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
	Manager_CreateDiscount_FullMethodName     = "/inspire.manager.coupon.discount.v1.Manager/CreateDiscount"
	Manager_CreateDiscounts_FullMethodName    = "/inspire.manager.coupon.discount.v1.Manager/CreateDiscounts"
	Manager_UpdateDiscount_FullMethodName     = "/inspire.manager.coupon.discount.v1.Manager/UpdateDiscount"
	Manager_GetDiscount_FullMethodName        = "/inspire.manager.coupon.discount.v1.Manager/GetDiscount"
	Manager_GetDiscountOnly_FullMethodName    = "/inspire.manager.coupon.discount.v1.Manager/GetDiscountOnly"
	Manager_GetDiscounts_FullMethodName       = "/inspire.manager.coupon.discount.v1.Manager/GetDiscounts"
	Manager_ExistDiscount_FullMethodName      = "/inspire.manager.coupon.discount.v1.Manager/ExistDiscount"
	Manager_ExistDiscountConds_FullMethodName = "/inspire.manager.coupon.discount.v1.Manager/ExistDiscountConds"
	Manager_CountDiscounts_FullMethodName     = "/inspire.manager.coupon.discount.v1.Manager/CountDiscounts"
	Manager_DeleteDiscount_FullMethodName     = "/inspire.manager.coupon.discount.v1.Manager/DeleteDiscount"
)

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*CreateDiscountResponse, error)
	CreateDiscounts(ctx context.Context, in *CreateDiscountsRequest, opts ...grpc.CallOption) (*CreateDiscountsResponse, error)
	UpdateDiscount(ctx context.Context, in *UpdateDiscountRequest, opts ...grpc.CallOption) (*UpdateDiscountResponse, error)
	GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error)
	GetDiscountOnly(ctx context.Context, in *GetDiscountOnlyRequest, opts ...grpc.CallOption) (*GetDiscountOnlyResponse, error)
	GetDiscounts(ctx context.Context, in *GetDiscountsRequest, opts ...grpc.CallOption) (*GetDiscountsResponse, error)
	ExistDiscount(ctx context.Context, in *ExistDiscountRequest, opts ...grpc.CallOption) (*ExistDiscountResponse, error)
	ExistDiscountConds(ctx context.Context, in *ExistDiscountCondsRequest, opts ...grpc.CallOption) (*ExistDiscountCondsResponse, error)
	CountDiscounts(ctx context.Context, in *CountDiscountsRequest, opts ...grpc.CallOption) (*CountDiscountsResponse, error)
	DeleteDiscount(ctx context.Context, in *DeleteDiscountRequest, opts ...grpc.CallOption) (*DeleteDiscountResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateDiscount(ctx context.Context, in *CreateDiscountRequest, opts ...grpc.CallOption) (*CreateDiscountResponse, error) {
	out := new(CreateDiscountResponse)
	err := c.cc.Invoke(ctx, Manager_CreateDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateDiscounts(ctx context.Context, in *CreateDiscountsRequest, opts ...grpc.CallOption) (*CreateDiscountsResponse, error) {
	out := new(CreateDiscountsResponse)
	err := c.cc.Invoke(ctx, Manager_CreateDiscounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateDiscount(ctx context.Context, in *UpdateDiscountRequest, opts ...grpc.CallOption) (*UpdateDiscountResponse, error) {
	out := new(UpdateDiscountResponse)
	err := c.cc.Invoke(ctx, Manager_UpdateDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error) {
	out := new(GetDiscountResponse)
	err := c.cc.Invoke(ctx, Manager_GetDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDiscountOnly(ctx context.Context, in *GetDiscountOnlyRequest, opts ...grpc.CallOption) (*GetDiscountOnlyResponse, error) {
	out := new(GetDiscountOnlyResponse)
	err := c.cc.Invoke(ctx, Manager_GetDiscountOnly_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDiscounts(ctx context.Context, in *GetDiscountsRequest, opts ...grpc.CallOption) (*GetDiscountsResponse, error) {
	out := new(GetDiscountsResponse)
	err := c.cc.Invoke(ctx, Manager_GetDiscounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistDiscount(ctx context.Context, in *ExistDiscountRequest, opts ...grpc.CallOption) (*ExistDiscountResponse, error) {
	out := new(ExistDiscountResponse)
	err := c.cc.Invoke(ctx, Manager_ExistDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistDiscountConds(ctx context.Context, in *ExistDiscountCondsRequest, opts ...grpc.CallOption) (*ExistDiscountCondsResponse, error) {
	out := new(ExistDiscountCondsResponse)
	err := c.cc.Invoke(ctx, Manager_ExistDiscountConds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountDiscounts(ctx context.Context, in *CountDiscountsRequest, opts ...grpc.CallOption) (*CountDiscountsResponse, error) {
	out := new(CountDiscountsResponse)
	err := c.cc.Invoke(ctx, Manager_CountDiscounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteDiscount(ctx context.Context, in *DeleteDiscountRequest, opts ...grpc.CallOption) (*DeleteDiscountResponse, error) {
	out := new(DeleteDiscountResponse)
	err := c.cc.Invoke(ctx, Manager_DeleteDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateDiscount(context.Context, *CreateDiscountRequest) (*CreateDiscountResponse, error)
	CreateDiscounts(context.Context, *CreateDiscountsRequest) (*CreateDiscountsResponse, error)
	UpdateDiscount(context.Context, *UpdateDiscountRequest) (*UpdateDiscountResponse, error)
	GetDiscount(context.Context, *GetDiscountRequest) (*GetDiscountResponse, error)
	GetDiscountOnly(context.Context, *GetDiscountOnlyRequest) (*GetDiscountOnlyResponse, error)
	GetDiscounts(context.Context, *GetDiscountsRequest) (*GetDiscountsResponse, error)
	ExistDiscount(context.Context, *ExistDiscountRequest) (*ExistDiscountResponse, error)
	ExistDiscountConds(context.Context, *ExistDiscountCondsRequest) (*ExistDiscountCondsResponse, error)
	CountDiscounts(context.Context, *CountDiscountsRequest) (*CountDiscountsResponse, error)
	DeleteDiscount(context.Context, *DeleteDiscountRequest) (*DeleteDiscountResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateDiscount(context.Context, *CreateDiscountRequest) (*CreateDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscount not implemented")
}
func (UnimplementedManagerServer) CreateDiscounts(context.Context, *CreateDiscountsRequest) (*CreateDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscounts not implemented")
}
func (UnimplementedManagerServer) UpdateDiscount(context.Context, *UpdateDiscountRequest) (*UpdateDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscount not implemented")
}
func (UnimplementedManagerServer) GetDiscount(context.Context, *GetDiscountRequest) (*GetDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}
func (UnimplementedManagerServer) GetDiscountOnly(context.Context, *GetDiscountOnlyRequest) (*GetDiscountOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscountOnly not implemented")
}
func (UnimplementedManagerServer) GetDiscounts(context.Context, *GetDiscountsRequest) (*GetDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscounts not implemented")
}
func (UnimplementedManagerServer) ExistDiscount(context.Context, *ExistDiscountRequest) (*ExistDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDiscount not implemented")
}
func (UnimplementedManagerServer) ExistDiscountConds(context.Context, *ExistDiscountCondsRequest) (*ExistDiscountCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDiscountConds not implemented")
}
func (UnimplementedManagerServer) CountDiscounts(context.Context, *CountDiscountsRequest) (*CountDiscountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDiscounts not implemented")
}
func (UnimplementedManagerServer) DeleteDiscount(context.Context, *DeleteDiscountRequest) (*DeleteDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscount not implemented")
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

func _Manager_CreateDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateDiscount(ctx, req.(*CreateDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CreateDiscounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateDiscounts(ctx, req.(*CreateDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_UpdateDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateDiscount(ctx, req.(*UpdateDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDiscount(ctx, req.(*GetDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDiscountOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDiscountOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetDiscountOnly_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDiscountOnly(ctx, req.(*GetDiscountOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_GetDiscounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDiscounts(ctx, req.(*GetDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistDiscount(ctx, req.(*ExistDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistDiscountConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDiscountCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistDiscountConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_ExistDiscountConds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistDiscountConds(ctx, req.(*ExistDiscountCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountDiscountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_CountDiscounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountDiscounts(ctx, req.(*CountDiscountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Manager_DeleteDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteDiscount(ctx, req.(*DeleteDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.manager.coupon.discount.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDiscount",
			Handler:    _Manager_CreateDiscount_Handler,
		},
		{
			MethodName: "CreateDiscounts",
			Handler:    _Manager_CreateDiscounts_Handler,
		},
		{
			MethodName: "UpdateDiscount",
			Handler:    _Manager_UpdateDiscount_Handler,
		},
		{
			MethodName: "GetDiscount",
			Handler:    _Manager_GetDiscount_Handler,
		},
		{
			MethodName: "GetDiscountOnly",
			Handler:    _Manager_GetDiscountOnly_Handler,
		},
		{
			MethodName: "GetDiscounts",
			Handler:    _Manager_GetDiscounts_Handler,
		},
		{
			MethodName: "ExistDiscount",
			Handler:    _Manager_ExistDiscount_Handler,
		},
		{
			MethodName: "ExistDiscountConds",
			Handler:    _Manager_ExistDiscountConds_Handler,
		},
		{
			MethodName: "CountDiscounts",
			Handler:    _Manager_CountDiscounts_Handler,
		},
		{
			MethodName: "DeleteDiscount",
			Handler:    _Manager_DeleteDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/mgr/v1/coupon/discount/discount.proto",
}
