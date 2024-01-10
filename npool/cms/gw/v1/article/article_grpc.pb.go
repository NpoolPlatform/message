// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: npool/cms/gw/v1/article/article.proto

package article

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
	Gateway_GetContent_FullMethodName        = "/cms.gateway.article.v1.Gateway/GetContent"
	Gateway_GetContentList_FullMethodName    = "/cms.gateway.article.v1.Gateway/GetContentList"
	Gateway_CreateArticle_FullMethodName     = "/cms.gateway.article.v1.Gateway/CreateArticle"
	Gateway_UpdateArticle_FullMethodName     = "/cms.gateway.article.v1.Gateway/UpdateArticle"
	Gateway_DeleteArticle_FullMethodName     = "/cms.gateway.article.v1.Gateway/DeleteArticle"
	Gateway_GetArticles_FullMethodName       = "/cms.gateway.article.v1.Gateway/GetArticles"
	Gateway_GetArticleContent_FullMethodName = "/cms.gateway.article.v1.Gateway/GetArticleContent"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// URL: https://{Host}/api/cms/v1/c/{ContentURL}
	// Host ex: api.site.top
	// ContentURL: {Category1.Slug}/{Category2.Slug}/.../{ISO}/{PageName}
	// ex: news/markets/en-US/test-page.html
	// ISO: {LANG}
	// ex: en-US
	// PageName: {Title}.html -> ToLower({Title}) -> ReplaceAll({Title}, " ", "-")
	// ex: test-page.html
	GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error)
	GetContentList(ctx context.Context, in *GetContentListRequest, opts ...grpc.CallOption) (*GetContentListResponse, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticleContent(ctx context.Context, in *GetArticleContentRequest, opts ...grpc.CallOption) (*GetArticleContentResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error) {
	out := new(GetContentResponse)
	err := c.cc.Invoke(ctx, Gateway_GetContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetContentList(ctx context.Context, in *GetContentListRequest, opts ...grpc.CallOption) (*GetContentListResponse, error) {
	out := new(GetContentListResponse)
	err := c.cc.Invoke(ctx, Gateway_GetContentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleResponse, error) {
	out := new(UpdateArticleResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleResponse, error) {
	out := new(DeleteArticleResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetArticles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetArticleContent(ctx context.Context, in *GetArticleContentRequest, opts ...grpc.CallOption) (*GetArticleContentResponse, error) {
	out := new(GetArticleContentResponse)
	err := c.cc.Invoke(ctx, Gateway_GetArticleContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// URL: https://{Host}/api/cms/v1/c/{ContentURL}
	// Host ex: api.site.top
	// ContentURL: {Category1.Slug}/{Category2.Slug}/.../{ISO}/{PageName}
	// ex: news/markets/en-US/test-page.html
	// ISO: {LANG}
	// ex: en-US
	// PageName: {Title}.html -> ToLower({Title}) -> ReplaceAll({Title}, " ", "-")
	// ex: test-page.html
	GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error)
	GetContentList(context.Context, *GetContentListRequest) (*GetContentListResponse, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error)
	GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesResponse, error)
	GetArticleContent(context.Context, *GetArticleContentRequest) (*GetArticleContentResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedGatewayServer) GetContentList(context.Context, *GetContentListRequest) (*GetContentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentList not implemented")
}
func (UnimplementedGatewayServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedGatewayServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedGatewayServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedGatewayServer) GetArticles(context.Context, *GetArticlesRequest) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedGatewayServer) GetArticleContent(context.Context, *GetArticleContentRequest) (*GetArticleContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleContent not implemented")
}
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

func _Gateway_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetContent(ctx, req.(*GetContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetContentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetContentList(ctx, req.(*GetContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetArticles(ctx, req.(*GetArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetArticleContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetArticleContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetArticleContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetArticleContent(ctx, req.(*GetArticleContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.gateway.article.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _Gateway_GetContent_Handler,
		},
		{
			MethodName: "GetContentList",
			Handler:    _Gateway_GetContentList_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Gateway_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Gateway_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Gateway_DeleteArticle_Handler,
		},
		{
			MethodName: "GetArticles",
			Handler:    _Gateway_GetArticles_Handler,
		},
		{
			MethodName: "GetArticleContent",
			Handler:    _Gateway_GetArticleContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/cms/gw/v1/article/article.proto",
}
