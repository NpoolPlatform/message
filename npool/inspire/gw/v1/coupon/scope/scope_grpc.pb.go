// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/inspire/gw/v1/coupon/scope/scope.proto

package scope

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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 3b0497879 (update scope pb)
=======
>>>>>>> 2e441a23e (update scope)
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
=======
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 9589c2455 (update scope)
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
	Gateway_CreateScope_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName    = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
	Gateway_GetAppScopes_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
	Gateway_CreateScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName     = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
	Gateway_GetAppScopes_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes"
	Gateway_GetNAppScopes_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/GetNAppScopes"
>>>>>>> 36e1391e1 (update scope pb)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 2c48ebb04 (update scope)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
	Gateway_CreateScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName     = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
	Gateway_GetAppScopes_FullMethodName  = "/inspire.gateway.coupon.scope.v1.Gateway/GetAppScopes"
	Gateway_GetNAppScopes_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/GetNAppScopes"
>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
	Gateway_CreateScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/CreateScope"
	Gateway_DeleteScope_FullMethodName = "/inspire.gateway.coupon.scope.v1.Gateway/DeleteScope"
	Gateway_GetScopes_FullMethodName   = "/inspire.gateway.coupon.scope.v1.Gateway/GetScopes"
>>>>>>> 61a423dc4 (delete get app scopes)
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error)
	DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error)
	GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 4b738f188 (update scope)
=======
	GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 36e1391e1 (update scope pb)
=======
	GetNAppScopes(ctx context.Context, in *GetNAppScopesRequest, opts ...grpc.CallOption) (*GetNAppScopesResponse, error)
>>>>>>> 7c8102624 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> e243bad0c (add coupon scope in gw)
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
	GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
	GetNAppScopes(ctx context.Context, in *GetNAppScopesRequest, opts ...grpc.CallOption) (*GetNAppScopesResponse, error)
>>>>>>> 0e532945d (update scope pb)
<<<<<<< HEAD
>>>>>>> 33e62963e (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
	GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error)
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateScope(ctx context.Context, in *CreateScopeRequest, opts ...grpc.CallOption) (*CreateScopeResponse, error) {
	out := new(CreateScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteScope(ctx context.Context, in *DeleteScopeRequest, opts ...grpc.CallOption) (*DeleteScopeResponse, error) {
	out := new(DeleteScopeResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetScopes(ctx context.Context, in *GetScopesRequest, opts ...grpc.CallOption) (*GetScopesResponse, error) {
	out := new(GetScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 9589c2455 (update scope)
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
func (c *gatewayClient) GetAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
func (c *gatewayClient) GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
=======
func (c *gatewayClient) GetNAppScopes(ctx context.Context, in *GetNAppScopesRequest, opts ...grpc.CallOption) (*GetNAppScopesResponse, error) {
	out := new(GetNAppScopesResponse)
>>>>>>> 7c8102624 (update scope pb)
	err := c.cc.Invoke(ctx, Gateway_GetNAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 3b0497879 (update scope pb)
=======
>>>>>>> 33e62963e (update scope pb)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
func (c *gatewayClient) GetNAppScopes(ctx context.Context, in *GetAppScopesRequest, opts ...grpc.CallOption) (*GetAppScopesResponse, error) {
	out := new(GetAppScopesResponse)
=======
func (c *gatewayClient) GetNAppScopes(ctx context.Context, in *GetNAppScopesRequest, opts ...grpc.CallOption) (*GetNAppScopesResponse, error) {
	out := new(GetNAppScopesResponse)
>>>>>>> 0e532945d (update scope pb)
	err := c.cc.Invoke(ctx, Gateway_GetNAppScopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error)
	DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error)
	GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 4b738f188 (update scope)
=======
	GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 36e1391e1 (update scope pb)
=======
	GetNAppScopes(context.Context, *GetNAppScopesRequest) (*GetNAppScopesResponse, error)
>>>>>>> 7c8102624 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> e243bad0c (add coupon scope in gw)
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> e243bad0c (add coupon scope in gw)
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
	GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
	GetNAppScopes(context.Context, *GetNAppScopesRequest) (*GetNAppScopesResponse, error)
>>>>>>> 0e532945d (update scope pb)
<<<<<<< HEAD
>>>>>>> 33e62963e (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
	GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error)
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) CreateScope(context.Context, *CreateScopeRequest) (*CreateScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScope not implemented")
}
func (UnimplementedGatewayServer) DeleteScope(context.Context, *DeleteScopeRequest) (*DeleteScopeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScope not implemented")
}
func (UnimplementedGatewayServer) GetScopes(context.Context, *GetScopesRequest) (*GetScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScopes not implemented")
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 4b738f188 (update scope)
=======
func (UnimplementedGatewayServer) GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
=======
func (UnimplementedGatewayServer) GetNAppScopes(context.Context, *GetNAppScopesRequest) (*GetNAppScopesResponse, error) {
>>>>>>> 7c8102624 (update scope pb)
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppScopes not implemented")
}
>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 3b0497879 (update scope pb)
=======
>>>>>>> 33e62963e (update scope pb)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
func (UnimplementedGatewayServer) GetNAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
=======
func (UnimplementedGatewayServer) GetNAppScopes(context.Context, *GetNAppScopesRequest) (*GetNAppScopesResponse, error) {
>>>>>>> 0e532945d (update scope pb)
	return nil, status.Errorf(codes.Unimplemented, "method GetNAppScopes not implemented")
}
>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
func (UnimplementedGatewayServer) GetAppScopes(context.Context, *GetAppScopesRequest) (*GetAppScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppScopes not implemented")
}
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateScope(ctx, req.(*CreateScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteScope(ctx, req.(*DeleteScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetScopes(ctx, req.(*GetScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 9589c2455 (update scope)
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
func _Gateway_GetAppScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppScopes(ctx, req.(*GetAppScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
func _Gateway_GetNAppScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppScopes(ctx, req.(*GetNAppScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 3b0497879 (update scope pb)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
func _Gateway_GetNAppScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNAppScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetNAppScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetNAppScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetNAppScopes(ctx, req.(*GetNAppScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coupon.scope.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScope",
			Handler:    _Gateway_CreateScope_Handler,
		},
		{
			MethodName: "DeleteScope",
			Handler:    _Gateway_DeleteScope_Handler,
		},
		{
			MethodName: "GetScopes",
			Handler:    _Gateway_GetScopes_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b738f188 (update scope)
=======
=======
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
=======
>>>>>>> 1fb4844cc (update scope)
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 343ac084f (delete get app scopes)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 2e441a23e (update scope)
=======
=======
>>>>>>> 1fb4844cc (update scope)
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 9589c2455 (update scope)
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
		{
			MethodName: "GetAppScopes",
			Handler:    _Gateway_GetAppScopes_Handler,
		},
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> f213f6ed2 (add coupon scope in gw)
=======
>>>>>>> 447c10d15 (delete get app scopes)
=======
>>>>>>> f83005c76 (update scope)
=======
>>>>>>> 4b0ed6935 (delete get app scopes)
=======
>>>>>>> 4b738f188 (update scope)
=======
		{
			MethodName: "GetNAppScopes",
			Handler:    _Gateway_GetNAppScopes_Handler,
		},
>>>>>>> 36e1391e1 (update scope pb)
=======
>>>>>>> 789287843 (update scope)
=======
=======
>>>>>>> 509fa36c7 (update scope)
=======
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
>>>>>>> 2c48ebb04 (update scope)
=======
>>>>>>> 3b0497879 (update scope pb)
>>>>>>> e243bad0c (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> eb5c69d26 (add coupon scope in gw)
=======
=======
>>>>>>> feba7689e (delete get app scopes)
<<<<<<< HEAD
>>>>>>> e8948f8dc (delete get app scopes)
=======
=======
>>>>>>> 1fb4844cc (update scope)
<<<<<<< HEAD
>>>>>>> 509fa36c7 (update scope)
=======
=======
>>>>>>> 3a2981d70 (add coupon scope in gw)
<<<<<<< HEAD
>>>>>>> c3156b2f2 (add coupon scope in gw)
=======
=======
>>>>>>> 67de434f0 (delete get app scopes)
<<<<<<< HEAD
>>>>>>> 343ac084f (delete get app scopes)
=======
=======
>>>>>>> 9589c2455 (update scope)
<<<<<<< HEAD
>>>>>>> 2c48ebb04 (update scope)
=======
=======
		{
			MethodName: "GetNAppScopes",
			Handler:    _Gateway_GetNAppScopes_Handler,
		},
>>>>>>> 960f133ce (update scope pb)
<<<<<<< HEAD
>>>>>>> 3b0497879 (update scope pb)
=======
=======
>>>>>>> 569d52611 (update scope)
>>>>>>> 2e441a23e (update scope)
=======
>>>>>>> 6dc6b83d5 (add coupon scope in gw)
=======
>>>>>>> 61a423dc4 (delete get app scopes)
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/inspire/gw/v1/coupon/scope/scope.proto",
}
