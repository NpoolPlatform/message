// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/project-info-manager/project-info-manager.proto

package project_info_manager

import (
	context "context"
	npool "github.com/NpoolPlatform/message/npool"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProjectInfoManagerClient is the client API for ProjectInfoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectInfoManagerClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error)
	CreateCoinDescriptions(ctx context.Context, in *CreateCoinDescriptionsRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionsResponse, error)
	UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error)
	UpdateAppCoinDescription(ctx context.Context, in *UpdateAppCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateAppCoinDescriptionResponse, error)
	GetCoinDescription(ctx context.Context, in *GetCoinDescriptionRequest, opts ...grpc.CallOption) (*GetCoinDescriptionResponse, error)
	GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error)
	GetCoinDescriptionOnly(ctx context.Context, in *GetCoinDescriptionOnlyRequest, opts ...grpc.CallOption) (*GetCoinDescriptionOnlyResponse, error)
	GetAppCoinDescriptions(ctx context.Context, in *GetAppCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetAppCoinDescriptionsResponse, error)
	DeleteAppCoinDescription(ctx context.Context, in *DeleteAppCoinDescriptionRequest, opts ...grpc.CallOption) (*DeleteAppCoinDescriptionResponse, error)
	DeleteCoinDescription(ctx context.Context, in *DeleteCoinDescriptionRequest, opts ...grpc.CallOption) (*DeleteCoinDescriptionResponse, error)
}

type projectInfoManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectInfoManagerClient(cc grpc.ClientConnInterface) ProjectInfoManagerClient {
	return &projectInfoManagerClient{cc}
}

func (c *projectInfoManagerClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) CreateCoinDescription(ctx context.Context, in *CreateCoinDescriptionRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionResponse, error) {
	out := new(CreateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/CreateCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) CreateCoinDescriptions(ctx context.Context, in *CreateCoinDescriptionsRequest, opts ...grpc.CallOption) (*CreateCoinDescriptionsResponse, error) {
	out := new(CreateCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/CreateCoinDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) UpdateCoinDescription(ctx context.Context, in *UpdateCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateCoinDescriptionResponse, error) {
	out := new(UpdateCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/UpdateCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) UpdateAppCoinDescription(ctx context.Context, in *UpdateAppCoinDescriptionRequest, opts ...grpc.CallOption) (*UpdateAppCoinDescriptionResponse, error) {
	out := new(UpdateAppCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/UpdateAppCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) GetCoinDescription(ctx context.Context, in *GetCoinDescriptionRequest, opts ...grpc.CallOption) (*GetCoinDescriptionResponse, error) {
	out := new(GetCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/GetCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) GetCoinDescriptions(ctx context.Context, in *GetCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetCoinDescriptionsResponse, error) {
	out := new(GetCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/GetCoinDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) GetCoinDescriptionOnly(ctx context.Context, in *GetCoinDescriptionOnlyRequest, opts ...grpc.CallOption) (*GetCoinDescriptionOnlyResponse, error) {
	out := new(GetCoinDescriptionOnlyResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/GetCoinDescriptionOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) GetAppCoinDescriptions(ctx context.Context, in *GetAppCoinDescriptionsRequest, opts ...grpc.CallOption) (*GetAppCoinDescriptionsResponse, error) {
	out := new(GetAppCoinDescriptionsResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/GetAppCoinDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) DeleteAppCoinDescription(ctx context.Context, in *DeleteAppCoinDescriptionRequest, opts ...grpc.CallOption) (*DeleteAppCoinDescriptionResponse, error) {
	out := new(DeleteAppCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/DeleteAppCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectInfoManagerClient) DeleteCoinDescription(ctx context.Context, in *DeleteCoinDescriptionRequest, opts ...grpc.CallOption) (*DeleteCoinDescriptionResponse, error) {
	out := new(DeleteCoinDescriptionResponse)
	err := c.cc.Invoke(ctx, "/project.info.manager.v1.ProjectInfoManager/DeleteCoinDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectInfoManagerServer is the server API for ProjectInfoManager service.
// All implementations must embed UnimplementedProjectInfoManagerServer
// for forward compatibility
type ProjectInfoManagerServer interface {
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error)
	CreateCoinDescriptions(context.Context, *CreateCoinDescriptionsRequest) (*CreateCoinDescriptionsResponse, error)
	UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error)
	UpdateAppCoinDescription(context.Context, *UpdateAppCoinDescriptionRequest) (*UpdateAppCoinDescriptionResponse, error)
	GetCoinDescription(context.Context, *GetCoinDescriptionRequest) (*GetCoinDescriptionResponse, error)
	GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error)
	GetCoinDescriptionOnly(context.Context, *GetCoinDescriptionOnlyRequest) (*GetCoinDescriptionOnlyResponse, error)
	GetAppCoinDescriptions(context.Context, *GetAppCoinDescriptionsRequest) (*GetAppCoinDescriptionsResponse, error)
	DeleteAppCoinDescription(context.Context, *DeleteAppCoinDescriptionRequest) (*DeleteAppCoinDescriptionResponse, error)
	DeleteCoinDescription(context.Context, *DeleteCoinDescriptionRequest) (*DeleteCoinDescriptionResponse, error)
	mustEmbedUnimplementedProjectInfoManagerServer()
}

// UnimplementedProjectInfoManagerServer must be embedded to have forward compatible implementations.
type UnimplementedProjectInfoManagerServer struct {
}

func (UnimplementedProjectInfoManagerServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedProjectInfoManagerServer) CreateCoinDescription(context.Context, *CreateCoinDescriptionRequest) (*CreateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) CreateCoinDescriptions(context.Context, *CreateCoinDescriptionsRequest) (*CreateCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoinDescriptions not implemented")
}
func (UnimplementedProjectInfoManagerServer) UpdateCoinDescription(context.Context, *UpdateCoinDescriptionRequest) (*UpdateCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) UpdateAppCoinDescription(context.Context, *UpdateAppCoinDescriptionRequest) (*UpdateAppCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) GetCoinDescription(context.Context, *GetCoinDescriptionRequest) (*GetCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) GetCoinDescriptions(context.Context, *GetCoinDescriptionsRequest) (*GetCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescriptions not implemented")
}
func (UnimplementedProjectInfoManagerServer) GetCoinDescriptionOnly(context.Context, *GetCoinDescriptionOnlyRequest) (*GetCoinDescriptionOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinDescriptionOnly not implemented")
}
func (UnimplementedProjectInfoManagerServer) GetAppCoinDescriptions(context.Context, *GetAppCoinDescriptionsRequest) (*GetAppCoinDescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppCoinDescriptions not implemented")
}
func (UnimplementedProjectInfoManagerServer) DeleteAppCoinDescription(context.Context, *DeleteAppCoinDescriptionRequest) (*DeleteAppCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) DeleteCoinDescription(context.Context, *DeleteCoinDescriptionRequest) (*DeleteCoinDescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoinDescription not implemented")
}
func (UnimplementedProjectInfoManagerServer) mustEmbedUnimplementedProjectInfoManagerServer() {}

// UnsafeProjectInfoManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectInfoManagerServer will
// result in compilation errors.
type UnsafeProjectInfoManagerServer interface {
	mustEmbedUnimplementedProjectInfoManagerServer()
}

func RegisterProjectInfoManagerServer(s grpc.ServiceRegistrar, srv ProjectInfoManagerServer) {
	s.RegisterService(&ProjectInfoManager_ServiceDesc, srv)
}

func _ProjectInfoManager_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_CreateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).CreateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/CreateCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).CreateCoinDescription(ctx, req.(*CreateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_CreateCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).CreateCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/CreateCoinDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).CreateCoinDescriptions(ctx, req.(*CreateCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_UpdateCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).UpdateCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/UpdateCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).UpdateCoinDescription(ctx, req.(*UpdateCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_UpdateAppCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).UpdateAppCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/UpdateAppCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).UpdateAppCoinDescription(ctx, req.(*UpdateAppCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_GetCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).GetCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/GetCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).GetCoinDescription(ctx, req.(*GetCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_GetCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).GetCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/GetCoinDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).GetCoinDescriptions(ctx, req.(*GetCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_GetCoinDescriptionOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinDescriptionOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).GetCoinDescriptionOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/GetCoinDescriptionOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).GetCoinDescriptionOnly(ctx, req.(*GetCoinDescriptionOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_GetAppCoinDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppCoinDescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).GetAppCoinDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/GetAppCoinDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).GetAppCoinDescriptions(ctx, req.(*GetAppCoinDescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_DeleteAppCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).DeleteAppCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/DeleteAppCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).DeleteAppCoinDescription(ctx, req.(*DeleteAppCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectInfoManager_DeleteCoinDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectInfoManagerServer).DeleteCoinDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.info.manager.v1.ProjectInfoManager/DeleteCoinDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectInfoManagerServer).DeleteCoinDescription(ctx, req.(*DeleteCoinDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectInfoManager_ServiceDesc is the grpc.ServiceDesc for ProjectInfoManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectInfoManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.info.manager.v1.ProjectInfoManager",
	HandlerType: (*ProjectInfoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _ProjectInfoManager_Version_Handler,
		},
		{
			MethodName: "CreateCoinDescription",
			Handler:    _ProjectInfoManager_CreateCoinDescription_Handler,
		},
		{
			MethodName: "CreateCoinDescriptions",
			Handler:    _ProjectInfoManager_CreateCoinDescriptions_Handler,
		},
		{
			MethodName: "UpdateCoinDescription",
			Handler:    _ProjectInfoManager_UpdateCoinDescription_Handler,
		},
		{
			MethodName: "UpdateAppCoinDescription",
			Handler:    _ProjectInfoManager_UpdateAppCoinDescription_Handler,
		},
		{
			MethodName: "GetCoinDescription",
			Handler:    _ProjectInfoManager_GetCoinDescription_Handler,
		},
		{
			MethodName: "GetCoinDescriptions",
			Handler:    _ProjectInfoManager_GetCoinDescriptions_Handler,
		},
		{
			MethodName: "GetCoinDescriptionOnly",
			Handler:    _ProjectInfoManager_GetCoinDescriptionOnly_Handler,
		},
		{
			MethodName: "GetAppCoinDescriptions",
			Handler:    _ProjectInfoManager_GetAppCoinDescriptions_Handler,
		},
		{
			MethodName: "DeleteAppCoinDescription",
			Handler:    _ProjectInfoManager_DeleteAppCoinDescription_Handler,
		},
		{
			MethodName: "DeleteCoinDescription",
			Handler:    _ProjectInfoManager_DeleteCoinDescription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/project-info-manager/project-info-manager.proto",
}
