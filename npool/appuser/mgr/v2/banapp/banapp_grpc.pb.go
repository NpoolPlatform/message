// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/mgr/v2/banapp/banapp.proto

package banapp

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

// BanAppMgrClient is the client API for BanAppMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BanAppMgrClient interface {
	CreateBanApp(ctx context.Context, in *CreateBanAppRequest, opts ...grpc.CallOption) (*CreateBanAppResponse, error)
	CreateBanApps(ctx context.Context, in *CreateBanAppsRequest, opts ...grpc.CallOption) (*CreateBanAppsResponse, error)
	UpdateBanApp(ctx context.Context, in *UpdateBanAppRequest, opts ...grpc.CallOption) (*UpdateBanAppResponse, error)
	GetBanApp(ctx context.Context, in *GetBanAppRequest, opts ...grpc.CallOption) (*GetBanAppResponse, error)
	GetBanAppOnly(ctx context.Context, in *GetBanAppOnlyRequest, opts ...grpc.CallOption) (*GetBanAppOnlyResponse, error)
	GetBanApps(ctx context.Context, in *GetBanAppsRequest, opts ...grpc.CallOption) (*GetBanAppsResponse, error)
	ExistBanApp(ctx context.Context, in *ExistBanAppRequest, opts ...grpc.CallOption) (*ExistBanAppResponse, error)
	ExistBanAppConds(ctx context.Context, in *ExistBanAppCondsRequest, opts ...grpc.CallOption) (*ExistBanAppCondsResponse, error)
	CountBanApps(ctx context.Context, in *CountBanAppsRequest, opts ...grpc.CallOption) (*CountBanAppsResponse, error)
	DeleteBanApp(ctx context.Context, in *DeleteBanAppRequest, opts ...grpc.CallOption) (*DeleteBanAppResponse, error)
}

type banAppMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewBanAppMgrClient(cc grpc.ClientConnInterface) BanAppMgrClient {
	return &banAppMgrClient{cc}
}

func (c *banAppMgrClient) CreateBanApp(ctx context.Context, in *CreateBanAppRequest, opts ...grpc.CallOption) (*CreateBanAppResponse, error) {
	out := new(CreateBanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/CreateBanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) CreateBanApps(ctx context.Context, in *CreateBanAppsRequest, opts ...grpc.CallOption) (*CreateBanAppsResponse, error) {
	out := new(CreateBanAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/CreateBanApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) UpdateBanApp(ctx context.Context, in *UpdateBanAppRequest, opts ...grpc.CallOption) (*UpdateBanAppResponse, error) {
	out := new(UpdateBanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/UpdateBanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) GetBanApp(ctx context.Context, in *GetBanAppRequest, opts ...grpc.CallOption) (*GetBanAppResponse, error) {
	out := new(GetBanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/GetBanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) GetBanAppOnly(ctx context.Context, in *GetBanAppOnlyRequest, opts ...grpc.CallOption) (*GetBanAppOnlyResponse, error) {
	out := new(GetBanAppOnlyResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/GetBanAppOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) GetBanApps(ctx context.Context, in *GetBanAppsRequest, opts ...grpc.CallOption) (*GetBanAppsResponse, error) {
	out := new(GetBanAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/GetBanApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) ExistBanApp(ctx context.Context, in *ExistBanAppRequest, opts ...grpc.CallOption) (*ExistBanAppResponse, error) {
	out := new(ExistBanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/ExistBanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) ExistBanAppConds(ctx context.Context, in *ExistBanAppCondsRequest, opts ...grpc.CallOption) (*ExistBanAppCondsResponse, error) {
	out := new(ExistBanAppCondsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/ExistBanAppConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) CountBanApps(ctx context.Context, in *CountBanAppsRequest, opts ...grpc.CallOption) (*CountBanAppsResponse, error) {
	out := new(CountBanAppsResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/CountBanApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banAppMgrClient) DeleteBanApp(ctx context.Context, in *DeleteBanAppRequest, opts ...grpc.CallOption) (*DeleteBanAppResponse, error) {
	out := new(DeleteBanAppResponse)
	err := c.cc.Invoke(ctx, "/appuser.manager.banapp.v2.BanAppMgr/DeleteBanApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BanAppMgrServer is the server API for BanAppMgr service.
// All implementations must embed UnimplementedBanAppMgrServer
// for forward compatibility
type BanAppMgrServer interface {
	CreateBanApp(context.Context, *CreateBanAppRequest) (*CreateBanAppResponse, error)
	CreateBanApps(context.Context, *CreateBanAppsRequest) (*CreateBanAppsResponse, error)
	UpdateBanApp(context.Context, *UpdateBanAppRequest) (*UpdateBanAppResponse, error)
	GetBanApp(context.Context, *GetBanAppRequest) (*GetBanAppResponse, error)
	GetBanAppOnly(context.Context, *GetBanAppOnlyRequest) (*GetBanAppOnlyResponse, error)
	GetBanApps(context.Context, *GetBanAppsRequest) (*GetBanAppsResponse, error)
	ExistBanApp(context.Context, *ExistBanAppRequest) (*ExistBanAppResponse, error)
	ExistBanAppConds(context.Context, *ExistBanAppCondsRequest) (*ExistBanAppCondsResponse, error)
	CountBanApps(context.Context, *CountBanAppsRequest) (*CountBanAppsResponse, error)
	DeleteBanApp(context.Context, *DeleteBanAppRequest) (*DeleteBanAppResponse, error)
	mustEmbedUnimplementedBanAppMgrServer()
}

// UnimplementedBanAppMgrServer must be embedded to have forward compatible implementations.
type UnimplementedBanAppMgrServer struct {
}

func (UnimplementedBanAppMgrServer) CreateBanApp(context.Context, *CreateBanAppRequest) (*CreateBanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanApp not implemented")
}
func (UnimplementedBanAppMgrServer) CreateBanApps(context.Context, *CreateBanAppsRequest) (*CreateBanAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanApps not implemented")
}
func (UnimplementedBanAppMgrServer) UpdateBanApp(context.Context, *UpdateBanAppRequest) (*UpdateBanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanApp not implemented")
}
func (UnimplementedBanAppMgrServer) GetBanApp(context.Context, *GetBanAppRequest) (*GetBanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanApp not implemented")
}
func (UnimplementedBanAppMgrServer) GetBanAppOnly(context.Context, *GetBanAppOnlyRequest) (*GetBanAppOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanAppOnly not implemented")
}
func (UnimplementedBanAppMgrServer) GetBanApps(context.Context, *GetBanAppsRequest) (*GetBanAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanApps not implemented")
}
func (UnimplementedBanAppMgrServer) ExistBanApp(context.Context, *ExistBanAppRequest) (*ExistBanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistBanApp not implemented")
}
func (UnimplementedBanAppMgrServer) ExistBanAppConds(context.Context, *ExistBanAppCondsRequest) (*ExistBanAppCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistBanAppConds not implemented")
}
func (UnimplementedBanAppMgrServer) CountBanApps(context.Context, *CountBanAppsRequest) (*CountBanAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountBanApps not implemented")
}
func (UnimplementedBanAppMgrServer) DeleteBanApp(context.Context, *DeleteBanAppRequest) (*DeleteBanAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanApp not implemented")
}
func (UnimplementedBanAppMgrServer) mustEmbedUnimplementedBanAppMgrServer() {}

// UnsafeBanAppMgrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BanAppMgrServer will
// result in compilation errors.
type UnsafeBanAppMgrServer interface {
	mustEmbedUnimplementedBanAppMgrServer()
}

func RegisterBanAppMgrServer(s grpc.ServiceRegistrar, srv BanAppMgrServer) {
	s.RegisterService(&BanAppMgr_ServiceDesc, srv)
}

func _BanAppMgr_CreateBanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).CreateBanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/CreateBanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).CreateBanApp(ctx, req.(*CreateBanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_CreateBanApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBanAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).CreateBanApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/CreateBanApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).CreateBanApps(ctx, req.(*CreateBanAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_UpdateBanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).UpdateBanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/UpdateBanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).UpdateBanApp(ctx, req.(*UpdateBanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_GetBanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).GetBanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/GetBanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).GetBanApp(ctx, req.(*GetBanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_GetBanAppOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBanAppOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).GetBanAppOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/GetBanAppOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).GetBanAppOnly(ctx, req.(*GetBanAppOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_GetBanApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBanAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).GetBanApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/GetBanApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).GetBanApps(ctx, req.(*GetBanAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_ExistBanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistBanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).ExistBanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/ExistBanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).ExistBanApp(ctx, req.(*ExistBanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_ExistBanAppConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistBanAppCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).ExistBanAppConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/ExistBanAppConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).ExistBanAppConds(ctx, req.(*ExistBanAppCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_CountBanApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountBanAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).CountBanApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/CountBanApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).CountBanApps(ctx, req.(*CountBanAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BanAppMgr_DeleteBanApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanAppMgrServer).DeleteBanApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appuser.manager.banapp.v2.BanAppMgr/DeleteBanApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanAppMgrServer).DeleteBanApp(ctx, req.(*DeleteBanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BanAppMgr_ServiceDesc is the grpc.ServiceDesc for BanAppMgr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BanAppMgr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.manager.banapp.v2.BanAppMgr",
	HandlerType: (*BanAppMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBanApp",
			Handler:    _BanAppMgr_CreateBanApp_Handler,
		},
		{
			MethodName: "CreateBanApps",
			Handler:    _BanAppMgr_CreateBanApps_Handler,
		},
		{
			MethodName: "UpdateBanApp",
			Handler:    _BanAppMgr_UpdateBanApp_Handler,
		},
		{
			MethodName: "GetBanApp",
			Handler:    _BanAppMgr_GetBanApp_Handler,
		},
		{
			MethodName: "GetBanAppOnly",
			Handler:    _BanAppMgr_GetBanAppOnly_Handler,
		},
		{
			MethodName: "GetBanApps",
			Handler:    _BanAppMgr_GetBanApps_Handler,
		},
		{
			MethodName: "ExistBanApp",
			Handler:    _BanAppMgr_ExistBanApp_Handler,
		},
		{
			MethodName: "ExistBanAppConds",
			Handler:    _BanAppMgr_ExistBanAppConds_Handler,
		},
		{
			MethodName: "CountBanApps",
			Handler:    _BanAppMgr_CountBanApps_Handler,
		},
		{
			MethodName: "DeleteBanApp",
			Handler:    _BanAppMgr_DeleteBanApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/mgr/v2/banapp/banapp.proto",
}
