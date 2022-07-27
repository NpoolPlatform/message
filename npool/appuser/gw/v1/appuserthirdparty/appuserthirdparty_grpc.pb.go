// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/appuser/gw/v1/appuserthirdparty/appuserthirdparty.proto

package appuserthirdparty

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

// AppUserThirdPartyGwClient is the client API for AppUserThirdPartyGw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppUserThirdPartyGwClient interface {
	CreateThirdParty(ctx context.Context, in *CreateThirdPartyRequest, opts ...grpc.CallOption) (*CreateThirdPartyResponse, error)
}

type appUserThirdPartyGwClient struct {
	cc grpc.ClientConnInterface
}

func NewAppUserThirdPartyGwClient(cc grpc.ClientConnInterface) AppUserThirdPartyGwClient {
	return &appUserThirdPartyGwClient{cc}
}

func (c *appUserThirdPartyGwClient) CreateThirdParty(ctx context.Context, in *CreateThirdPartyRequest, opts ...grpc.CallOption) (*CreateThirdPartyResponse, error) {
	out := new(CreateThirdPartyResponse)
	err := c.cc.Invoke(ctx, "/app.user.gateway.thirdparty.v1.AppUserThirdPartyGw/CreateThirdParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppUserThirdPartyGwServer is the server API for AppUserThirdPartyGw service.
// All implementations must embed UnimplementedAppUserThirdPartyGwServer
// for forward compatibility
type AppUserThirdPartyGwServer interface {
	CreateThirdParty(context.Context, *CreateThirdPartyRequest) (*CreateThirdPartyResponse, error)
	mustEmbedUnimplementedAppUserThirdPartyGwServer()
}

// UnimplementedAppUserThirdPartyGwServer must be embedded to have forward compatible implementations.
type UnimplementedAppUserThirdPartyGwServer struct {
}

func (UnimplementedAppUserThirdPartyGwServer) CreateThirdParty(context.Context, *CreateThirdPartyRequest) (*CreateThirdPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateThirdParty not implemented")
}
func (UnimplementedAppUserThirdPartyGwServer) mustEmbedUnimplementedAppUserThirdPartyGwServer() {}

// UnsafeAppUserThirdPartyGwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppUserThirdPartyGwServer will
// result in compilation errors.
type UnsafeAppUserThirdPartyGwServer interface {
	mustEmbedUnimplementedAppUserThirdPartyGwServer()
}

func RegisterAppUserThirdPartyGwServer(s grpc.ServiceRegistrar, srv AppUserThirdPartyGwServer) {
	s.RegisterService(&AppUserThirdPartyGw_ServiceDesc, srv)
}

func _AppUserThirdPartyGw_CreateThirdParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateThirdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppUserThirdPartyGwServer).CreateThirdParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.user.gateway.thirdparty.v1.AppUserThirdPartyGw/CreateThirdParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppUserThirdPartyGwServer).CreateThirdParty(ctx, req.(*CreateThirdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppUserThirdPartyGw_ServiceDesc is the grpc.ServiceDesc for AppUserThirdPartyGw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppUserThirdPartyGw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.user.gateway.thirdparty.v1.AppUserThirdPartyGw",
	HandlerType: (*AppUserThirdPartyGwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateThirdParty",
			Handler:    _AppUserThirdPartyGw_CreateThirdParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/appuser/gw/v1/appuserthirdparty/appuserthirdparty.proto",
}
