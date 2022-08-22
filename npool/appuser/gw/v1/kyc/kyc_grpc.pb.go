// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/kyc/kyc.proto

package kyc

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
	CreateKyc(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error)
	UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error)
	GetKyc(ctx context.Context, in *GetKycRequest, opts ...grpc.CallOption) (*GetKycResponse, error)
	GetKycs(ctx context.Context, in *GetKycsRequest, opts ...grpc.CallOption) (*GetKycsResponse, error)
	GetAppKycs(ctx context.Context, in *GetAppKycsRequest, opts ...grpc.CallOption) (*GetAppKycsResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateKyc(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error) {
	out := new(CreateKycResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.kyc.v2.Manager/CreateKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error) {
	out := new(UpdateKycResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.kyc.v2.Manager/UpdateKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetKyc(ctx context.Context, in *GetKycRequest, opts ...grpc.CallOption) (*GetKycResponse, error) {
	out := new(GetKycResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.kyc.v2.Manager/GetKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetKycs(ctx context.Context, in *GetKycsRequest, opts ...grpc.CallOption) (*GetKycsResponse, error) {
	out := new(GetKycsResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.kyc.v2.Manager/GetKycs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetAppKycs(ctx context.Context, in *GetAppKycsRequest, opts ...grpc.CallOption) (*GetAppKycsResponse, error) {
	out := new(GetAppKycsResponse)
	err := c.cc.Invoke(ctx, "/appuser.gateway.kyc.v2.Manager/GetAppKycs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateKyc(context.Context, *CreateKycRequest) (*CreateKycResponse, error)
	UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error)
	GetKyc(context.Context, *GetKycRequest) (*GetKycResponse, error)
	GetKycs(context.Context, *GetKycsRequest) (*GetKycsResponse, error)
	GetAppKycs(context.Context, *GetAppKycsRequest) (*GetAppKycsResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateKyc(context.Context, *CreateKycRequest) (*CreateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKyc not implemented")
}
func (UnimplementedManagerServer) UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKyc not implemented")
}
func (UnimplementedManagerServer) GetKyc(context.Context, *GetKycRequest) (*GetKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKyc not implemented")
}
func (UnimplementedManagerServer) GetKycs(context.Context, *GetKycsRequest) (*GetKycsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycs not implemented")
}
func (UnimplementedManagerServer) GetAppKycs(context.Context, *GetAppKycsRequest) (*GetAppKycsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppKycs not implemented")
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

func _Manager_CreateKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.kyc.v2.Manager/CreateKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateKyc(ctx, req.(*CreateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.kyc.v2.Manager/UpdateKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateKyc(ctx, req.(*UpdateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.kyc.v2.Manager/GetKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetKyc(ctx, req.(*GetKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetKycs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetKycs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.kyc.v2.Manager/GetKycs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetKycs(ctx, req.(*GetKycsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetAppKycs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppKycsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetAppKycs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.gateway.kyc.v2.Manager/GetAppKycs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetAppKycs(ctx, req.(*GetAppKycsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.kyc.v2.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKyc",
			Handler:    _Manager_CreateKyc_Handler,
		},
		{
			MethodName: "UpdateKyc",
			Handler:    _Manager_UpdateKyc_Handler,
		},
		{
			MethodName: "GetKyc",
			Handler:    _Manager_GetKyc_Handler,
		},
		{
			MethodName: "GetKycs",
			Handler:    _Manager_GetKycs_Handler,
		},
		{
			MethodName: "GetAppKycs",
			Handler:    _Manager_GetAppKycs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/kyc/kyc.proto",
}
