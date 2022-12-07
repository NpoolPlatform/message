// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/g11n/gw/v1/applang/applang.proto

package applang

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
	CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error)
	UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error)
	GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error)
	GetAppLangs(ctx context.Context, in *GetAppLangsRequest, opts ...grpc.CallOption) (*GetAppLangsResponse, error)
	DeleteLang(ctx context.Context, in *DeleteLangRequest, opts ...grpc.CallOption) (*DeleteLangResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateLang(ctx context.Context, in *CreateLangRequest, opts ...grpc.CallOption) (*CreateLangResponse, error) {
	out := new(CreateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.applang.v1.Manager/CreateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateLang(ctx context.Context, in *UpdateLangRequest, opts ...grpc.CallOption) (*UpdateLangResponse, error) {
	out := new(UpdateLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.applang.v1.Manager/UpdateLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetLangs(ctx context.Context, in *GetLangsRequest, opts ...grpc.CallOption) (*GetLangsResponse, error) {
	out := new(GetLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.applang.v1.Manager/GetLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetAppLangs(ctx context.Context, in *GetAppLangsRequest, opts ...grpc.CallOption) (*GetAppLangsResponse, error) {
	out := new(GetAppLangsResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.applang.v1.Manager/GetAppLangs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteLang(ctx context.Context, in *DeleteLangRequest, opts ...grpc.CallOption) (*DeleteLangResponse, error) {
	out := new(DeleteLangResponse)
	err := c.cc.Invoke(ctx, "/g11n.gateway.applang.v1.Manager/DeleteLang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error)
	UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error)
	GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error)
	GetAppLangs(context.Context, *GetAppLangsRequest) (*GetAppLangsResponse, error)
	DeleteLang(context.Context, *DeleteLangRequest) (*DeleteLangResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateLang(context.Context, *CreateLangRequest) (*CreateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLang not implemented")
}
func (UnimplementedManagerServer) UpdateLang(context.Context, *UpdateLangRequest) (*UpdateLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLang not implemented")
}
func (UnimplementedManagerServer) GetLangs(context.Context, *GetLangsRequest) (*GetLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLangs not implemented")
}
func (UnimplementedManagerServer) GetAppLangs(context.Context, *GetAppLangsRequest) (*GetAppLangsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppLangs not implemented")
}
func (UnimplementedManagerServer) DeleteLang(context.Context, *DeleteLangRequest) (*DeleteLangResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLang not implemented")
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

func _Manager_CreateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.applang.v1.Manager/CreateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateLang(ctx, req.(*CreateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.applang.v1.Manager/UpdateLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateLang(ctx, req.(*UpdateLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.applang.v1.Manager/GetLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetLangs(ctx, req.(*GetLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetAppLangs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppLangsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetAppLangs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.applang.v1.Manager/GetAppLangs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetAppLangs(ctx, req.(*GetAppLangsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/g11n.gateway.applang.v1.Manager/DeleteLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteLang(ctx, req.(*DeleteLangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "g11n.gateway.applang.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLang",
			Handler:    _Manager_CreateLang_Handler,
		},
		{
			MethodName: "UpdateLang",
			Handler:    _Manager_UpdateLang_Handler,
		},
		{
			MethodName: "GetLangs",
			Handler:    _Manager_GetLangs_Handler,
		},
		{
			MethodName: "GetAppLangs",
			Handler:    _Manager_GetAppLangs_Handler,
		},
		{
			MethodName: "DeleteLang",
			Handler:    _Manager_DeleteLang_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/g11n/gw/v1/applang/applang.proto",
}
