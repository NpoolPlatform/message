// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/good/mgr/v1/deviceinfo/deviceinfo.proto

package deviceinfo

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
	CreateDeviceInfo(ctx context.Context, in *CreateDeviceInfoRequest, opts ...grpc.CallOption) (*CreateDeviceInfoResponse, error)
	CreateDeviceInfos(ctx context.Context, in *CreateDeviceInfosRequest, opts ...grpc.CallOption) (*CreateDeviceInfosResponse, error)
	GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error)
	GetDeviceInfoOnly(ctx context.Context, in *GetDeviceInfoOnlyRequest, opts ...grpc.CallOption) (*GetDeviceInfoOnlyResponse, error)
	GetDeviceInfos(ctx context.Context, in *GetDeviceInfosRequest, opts ...grpc.CallOption) (*GetDeviceInfosResponse, error)
	ExistDeviceInfo(ctx context.Context, in *ExistDeviceInfoRequest, opts ...grpc.CallOption) (*ExistDeviceInfoResponse, error)
	ExistDeviceInfoConds(ctx context.Context, in *ExistDeviceInfoCondsRequest, opts ...grpc.CallOption) (*ExistDeviceInfoCondsResponse, error)
	CountDeviceInfos(ctx context.Context, in *CountDeviceInfosRequest, opts ...grpc.CallOption) (*CountDeviceInfosResponse, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) CreateDeviceInfo(ctx context.Context, in *CreateDeviceInfoRequest, opts ...grpc.CallOption) (*CreateDeviceInfoResponse, error) {
	out := new(CreateDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/CreateDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateDeviceInfos(ctx context.Context, in *CreateDeviceInfosRequest, opts ...grpc.CallOption) (*CreateDeviceInfosResponse, error) {
	out := new(CreateDeviceInfosResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/CreateDeviceInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error) {
	out := new(GetDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDeviceInfoOnly(ctx context.Context, in *GetDeviceInfoOnlyRequest, opts ...grpc.CallOption) (*GetDeviceInfoOnlyResponse, error) {
	out := new(GetDeviceInfoOnlyResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/GetDeviceInfoOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetDeviceInfos(ctx context.Context, in *GetDeviceInfosRequest, opts ...grpc.CallOption) (*GetDeviceInfosResponse, error) {
	out := new(GetDeviceInfosResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/GetDeviceInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistDeviceInfo(ctx context.Context, in *ExistDeviceInfoRequest, opts ...grpc.CallOption) (*ExistDeviceInfoResponse, error) {
	out := new(ExistDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/ExistDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ExistDeviceInfoConds(ctx context.Context, in *ExistDeviceInfoCondsRequest, opts ...grpc.CallOption) (*ExistDeviceInfoCondsResponse, error) {
	out := new(ExistDeviceInfoCondsResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/ExistDeviceInfoConds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CountDeviceInfos(ctx context.Context, in *CountDeviceInfosRequest, opts ...grpc.CallOption) (*CountDeviceInfosResponse, error) {
	out := new(CountDeviceInfosResponse)
	err := c.cc.Invoke(ctx, "/good.manager.deviceinfo.v1.Manager/CountDeviceInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	CreateDeviceInfo(context.Context, *CreateDeviceInfoRequest) (*CreateDeviceInfoResponse, error)
	CreateDeviceInfos(context.Context, *CreateDeviceInfosRequest) (*CreateDeviceInfosResponse, error)
	GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error)
	GetDeviceInfoOnly(context.Context, *GetDeviceInfoOnlyRequest) (*GetDeviceInfoOnlyResponse, error)
	GetDeviceInfos(context.Context, *GetDeviceInfosRequest) (*GetDeviceInfosResponse, error)
	ExistDeviceInfo(context.Context, *ExistDeviceInfoRequest) (*ExistDeviceInfoResponse, error)
	ExistDeviceInfoConds(context.Context, *ExistDeviceInfoCondsRequest) (*ExistDeviceInfoCondsResponse, error)
	CountDeviceInfos(context.Context, *CountDeviceInfosRequest) (*CountDeviceInfosResponse, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) CreateDeviceInfo(context.Context, *CreateDeviceInfoRequest) (*CreateDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceInfo not implemented")
}
func (UnimplementedManagerServer) CreateDeviceInfos(context.Context, *CreateDeviceInfosRequest) (*CreateDeviceInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceInfos not implemented")
}
func (UnimplementedManagerServer) GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfo not implemented")
}
func (UnimplementedManagerServer) GetDeviceInfoOnly(context.Context, *GetDeviceInfoOnlyRequest) (*GetDeviceInfoOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfoOnly not implemented")
}
func (UnimplementedManagerServer) GetDeviceInfos(context.Context, *GetDeviceInfosRequest) (*GetDeviceInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfos not implemented")
}
func (UnimplementedManagerServer) ExistDeviceInfo(context.Context, *ExistDeviceInfoRequest) (*ExistDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDeviceInfo not implemented")
}
func (UnimplementedManagerServer) ExistDeviceInfoConds(context.Context, *ExistDeviceInfoCondsRequest) (*ExistDeviceInfoCondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDeviceInfoConds not implemented")
}
func (UnimplementedManagerServer) CountDeviceInfos(context.Context, *CountDeviceInfosRequest) (*CountDeviceInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDeviceInfos not implemented")
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

func _Manager_CreateDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/CreateDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateDeviceInfo(ctx, req.(*CreateDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateDeviceInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateDeviceInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/CreateDeviceInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateDeviceInfos(ctx, req.(*CreateDeviceInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDeviceInfo(ctx, req.(*GetDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDeviceInfoOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDeviceInfoOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/GetDeviceInfoOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDeviceInfoOnly(ctx, req.(*GetDeviceInfoOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetDeviceInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetDeviceInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/GetDeviceInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetDeviceInfos(ctx, req.(*GetDeviceInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/ExistDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistDeviceInfo(ctx, req.(*ExistDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ExistDeviceInfoConds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDeviceInfoCondsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ExistDeviceInfoConds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/ExistDeviceInfoConds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ExistDeviceInfoConds(ctx, req.(*ExistDeviceInfoCondsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CountDeviceInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountDeviceInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CountDeviceInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/good.manager.deviceinfo.v1.Manager/CountDeviceInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CountDeviceInfos(ctx, req.(*CountDeviceInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.manager.deviceinfo.v1.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceInfo",
			Handler:    _Manager_CreateDeviceInfo_Handler,
		},
		{
			MethodName: "CreateDeviceInfos",
			Handler:    _Manager_CreateDeviceInfos_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _Manager_GetDeviceInfo_Handler,
		},
		{
			MethodName: "GetDeviceInfoOnly",
			Handler:    _Manager_GetDeviceInfoOnly_Handler,
		},
		{
			MethodName: "GetDeviceInfos",
			Handler:    _Manager_GetDeviceInfos_Handler,
		},
		{
			MethodName: "ExistDeviceInfo",
			Handler:    _Manager_ExistDeviceInfo_Handler,
		},
		{
			MethodName: "ExistDeviceInfoConds",
			Handler:    _Manager_ExistDeviceInfoConds_Handler,
		},
		{
			MethodName: "CountDeviceInfos",
			Handler:    _Manager_CountDeviceInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/good/mgr/v1/deviceinfo/deviceinfo.proto",
}
