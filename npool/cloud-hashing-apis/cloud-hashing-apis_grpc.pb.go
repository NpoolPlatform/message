// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: npool/cloud-hashing-apis/cloud-hashing-apis.proto

package cloud_hashing_apis

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

// CloudHashingApisClient is the client API for CloudHashingApis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudHashingApisClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error)
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
	GetRecommendGoodsByApp(ctx context.Context, in *GetRecommendGoodsByAppRequest, opts ...grpc.CallOption) (*GetRecommendGoodsByAppResponse, error)
	SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error)
	CreateOrderPayment(ctx context.Context, in *CreateOrderPaymentRequest, opts ...grpc.CallOption) (*CreateOrderPaymentResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrdersByAppUser(ctx context.Context, in *GetOrdersByAppUserRequest, opts ...grpc.CallOption) (*GetOrdersByAppUserResponse, error)
	GetOrdersByApp(ctx context.Context, in *GetOrdersByAppRequest, opts ...grpc.CallOption) (*GetOrdersByAppResponse, error)
	GetOrdersByGood(ctx context.Context, in *GetOrdersByGoodRequest, opts ...grpc.CallOption) (*GetOrdersByGoodResponse, error)
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	UpdatePasswordByAppUser(ctx context.Context, in *UpdatePasswordByAppUserRequest, opts ...grpc.CallOption) (*UpdatePasswordByAppUserResponse, error)
	UpdateEmailAddress(ctx context.Context, in *UpdateEmailAddressRequest, opts ...grpc.CallOption) (*UpdateEmailAddressResponse, error)
	UpdatePhoneNO(ctx context.Context, in *UpdatePhoneNORequest, opts ...grpc.CallOption) (*UpdatePhoneNOResponse, error)
	GetMyInvitations(ctx context.Context, in *GetMyInvitationsRequest, opts ...grpc.CallOption) (*GetMyInvitationsResponse, error)
	GetMyDirectInvitations(ctx context.Context, in *GetMyDirectInvitationsRequest, opts ...grpc.CallOption) (*GetMyDirectInvitationsResponse, error)
	GetKycReviews(ctx context.Context, in *GetKycReviewsRequest, opts ...grpc.CallOption) (*GetKycReviewsResponse, error)
	GetGoodReviews(ctx context.Context, in *GetGoodReviewsRequest, opts ...grpc.CallOption) (*GetGoodReviewsResponse, error)
	CreateKyc(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error)
	UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error)
	GetKycByAppUser(ctx context.Context, in *GetKycByAppUserRequest, opts ...grpc.CallOption) (*GetKycByAppUserResponse, error)
}

type cloudHashingApisClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudHashingApisClient(cc grpc.ClientConnInterface) CloudHashingApisClient {
	return &cloudHashingApisClient{cc}
}

func (c *cloudHashingApisClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*npool.VersionResponse, error) {
	out := new(npool.VersionResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error) {
	out := new(CreateGoodResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/CreateGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetRecommendGoodsByApp(ctx context.Context, in *GetRecommendGoodsByAppRequest, opts ...grpc.CallOption) (*GetRecommendGoodsByAppResponse, error) {
	out := new(GetRecommendGoodsByAppResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetRecommendGoodsByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderResponse, error) {
	out := new(SubmitOrderResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) CreateOrderPayment(ctx context.Context, in *CreateOrderPaymentRequest, opts ...grpc.CallOption) (*CreateOrderPaymentResponse, error) {
	out := new(CreateOrderPaymentResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/CreateOrderPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetOrdersByAppUser(ctx context.Context, in *GetOrdersByAppUserRequest, opts ...grpc.CallOption) (*GetOrdersByAppUserResponse, error) {
	out := new(GetOrdersByAppUserResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetOrdersByApp(ctx context.Context, in *GetOrdersByAppRequest, opts ...grpc.CallOption) (*GetOrdersByAppResponse, error) {
	out := new(GetOrdersByAppResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetOrdersByGood(ctx context.Context, in *GetOrdersByGoodRequest, opts ...grpc.CallOption) (*GetOrdersByGoodResponse, error) {
	out := new(GetOrdersByGoodResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) UpdatePasswordByAppUser(ctx context.Context, in *UpdatePasswordByAppUserRequest, opts ...grpc.CallOption) (*UpdatePasswordByAppUserResponse, error) {
	out := new(UpdatePasswordByAppUserResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePasswordByAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) UpdateEmailAddress(ctx context.Context, in *UpdateEmailAddressRequest, opts ...grpc.CallOption) (*UpdateEmailAddressResponse, error) {
	out := new(UpdateEmailAddressResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/UpdateEmailAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) UpdatePhoneNO(ctx context.Context, in *UpdatePhoneNORequest, opts ...grpc.CallOption) (*UpdatePhoneNOResponse, error) {
	out := new(UpdatePhoneNOResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePhoneNO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetMyInvitations(ctx context.Context, in *GetMyInvitationsRequest, opts ...grpc.CallOption) (*GetMyInvitationsResponse, error) {
	out := new(GetMyInvitationsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetMyInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetMyDirectInvitations(ctx context.Context, in *GetMyDirectInvitationsRequest, opts ...grpc.CallOption) (*GetMyDirectInvitationsResponse, error) {
	out := new(GetMyDirectInvitationsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetMyDirectInvitations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetKycReviews(ctx context.Context, in *GetKycReviewsRequest, opts ...grpc.CallOption) (*GetKycReviewsResponse, error) {
	out := new(GetKycReviewsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetKycReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetGoodReviews(ctx context.Context, in *GetGoodReviewsRequest, opts ...grpc.CallOption) (*GetGoodReviewsResponse, error) {
	out := new(GetGoodReviewsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetGoodReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) CreateKyc(ctx context.Context, in *CreateKycRequest, opts ...grpc.CallOption) (*CreateKycResponse, error) {
	out := new(CreateKycResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/CreateKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) UpdateKyc(ctx context.Context, in *UpdateKycRequest, opts ...grpc.CallOption) (*UpdateKycResponse, error) {
	out := new(UpdateKycResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/UpdateKyc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingApisClient) GetKycByAppUser(ctx context.Context, in *GetKycByAppUserRequest, opts ...grpc.CallOption) (*GetKycByAppUserResponse, error) {
	out := new(GetKycByAppUserResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.apis.v1.CloudHashingApis/GetKycByAppUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudHashingApisServer is the server API for CloudHashingApis service.
// All implementations must embed UnimplementedCloudHashingApisServer
// for forward compatibility
type CloudHashingApisServer interface {
	Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error)
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	GetRecommendGoodsByApp(context.Context, *GetRecommendGoodsByAppRequest) (*GetRecommendGoodsByAppResponse, error)
	SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error)
	CreateOrderPayment(context.Context, *CreateOrderPaymentRequest) (*CreateOrderPaymentResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrdersByAppUser(context.Context, *GetOrdersByAppUserRequest) (*GetOrdersByAppUserResponse, error)
	GetOrdersByApp(context.Context, *GetOrdersByAppRequest) (*GetOrdersByAppResponse, error)
	GetOrdersByGood(context.Context, *GetOrdersByGoodRequest) (*GetOrdersByGoodResponse, error)
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	UpdatePasswordByAppUser(context.Context, *UpdatePasswordByAppUserRequest) (*UpdatePasswordByAppUserResponse, error)
	UpdateEmailAddress(context.Context, *UpdateEmailAddressRequest) (*UpdateEmailAddressResponse, error)
	UpdatePhoneNO(context.Context, *UpdatePhoneNORequest) (*UpdatePhoneNOResponse, error)
	GetMyInvitations(context.Context, *GetMyInvitationsRequest) (*GetMyInvitationsResponse, error)
	GetMyDirectInvitations(context.Context, *GetMyDirectInvitationsRequest) (*GetMyDirectInvitationsResponse, error)
	GetKycReviews(context.Context, *GetKycReviewsRequest) (*GetKycReviewsResponse, error)
	GetGoodReviews(context.Context, *GetGoodReviewsRequest) (*GetGoodReviewsResponse, error)
	CreateKyc(context.Context, *CreateKycRequest) (*CreateKycResponse, error)
	UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error)
	GetKycByAppUser(context.Context, *GetKycByAppUserRequest) (*GetKycByAppUserResponse, error)
	mustEmbedUnimplementedCloudHashingApisServer()
}

// UnimplementedCloudHashingApisServer must be embedded to have forward compatible implementations.
type UnimplementedCloudHashingApisServer struct {
}

func (UnimplementedCloudHashingApisServer) Version(context.Context, *emptypb.Empty) (*npool.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCloudHashingApisServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedCloudHashingApisServer) CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGood not implemented")
}
func (UnimplementedCloudHashingApisServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedCloudHashingApisServer) GetRecommendGoodsByApp(context.Context, *GetRecommendGoodsByAppRequest) (*GetRecommendGoodsByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendGoodsByApp not implemented")
}
func (UnimplementedCloudHashingApisServer) SubmitOrder(context.Context, *SubmitOrderRequest) (*SubmitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedCloudHashingApisServer) CreateOrderPayment(context.Context, *CreateOrderPaymentRequest) (*CreateOrderPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderPayment not implemented")
}
func (UnimplementedCloudHashingApisServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedCloudHashingApisServer) GetOrdersByAppUser(context.Context, *GetOrdersByAppUserRequest) (*GetOrdersByAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByAppUser not implemented")
}
func (UnimplementedCloudHashingApisServer) GetOrdersByApp(context.Context, *GetOrdersByAppRequest) (*GetOrdersByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByApp not implemented")
}
func (UnimplementedCloudHashingApisServer) GetOrdersByGood(context.Context, *GetOrdersByGoodRequest) (*GetOrdersByGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByGood not implemented")
}
func (UnimplementedCloudHashingApisServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedCloudHashingApisServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedCloudHashingApisServer) UpdatePasswordByAppUser(context.Context, *UpdatePasswordByAppUserRequest) (*UpdatePasswordByAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePasswordByAppUser not implemented")
}
func (UnimplementedCloudHashingApisServer) UpdateEmailAddress(context.Context, *UpdateEmailAddressRequest) (*UpdateEmailAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailAddress not implemented")
}
func (UnimplementedCloudHashingApisServer) UpdatePhoneNO(context.Context, *UpdatePhoneNORequest) (*UpdatePhoneNOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoneNO not implemented")
}
func (UnimplementedCloudHashingApisServer) GetMyInvitations(context.Context, *GetMyInvitationsRequest) (*GetMyInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyInvitations not implemented")
}
func (UnimplementedCloudHashingApisServer) GetMyDirectInvitations(context.Context, *GetMyDirectInvitationsRequest) (*GetMyDirectInvitationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyDirectInvitations not implemented")
}
func (UnimplementedCloudHashingApisServer) GetKycReviews(context.Context, *GetKycReviewsRequest) (*GetKycReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycReviews not implemented")
}
func (UnimplementedCloudHashingApisServer) GetGoodReviews(context.Context, *GetGoodReviewsRequest) (*GetGoodReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodReviews not implemented")
}
func (UnimplementedCloudHashingApisServer) CreateKyc(context.Context, *CreateKycRequest) (*CreateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKyc not implemented")
}
func (UnimplementedCloudHashingApisServer) UpdateKyc(context.Context, *UpdateKycRequest) (*UpdateKycResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKyc not implemented")
}
func (UnimplementedCloudHashingApisServer) GetKycByAppUser(context.Context, *GetKycByAppUserRequest) (*GetKycByAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKycByAppUser not implemented")
}
func (UnimplementedCloudHashingApisServer) mustEmbedUnimplementedCloudHashingApisServer() {}

// UnsafeCloudHashingApisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudHashingApisServer will
// result in compilation errors.
type UnsafeCloudHashingApisServer interface {
	mustEmbedUnimplementedCloudHashingApisServer()
}

func RegisterCloudHashingApisServer(s grpc.ServiceRegistrar, srv CloudHashingApisServer) {
	s.RegisterService(&CloudHashingApis_ServiceDesc, srv)
}

func _CloudHashingApis_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_CreateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).CreateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/CreateGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).CreateGood(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetRecommendGoodsByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendGoodsByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetRecommendGoodsByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetRecommendGoodsByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetRecommendGoodsByApp(ctx, req.(*GetRecommendGoodsByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).SubmitOrder(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_CreateOrderPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).CreateOrderPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/CreateOrderPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).CreateOrderPayment(ctx, req.(*CreateOrderPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetOrdersByAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersByAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetOrdersByAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetOrdersByAppUser(ctx, req.(*GetOrdersByAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetOrdersByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetOrdersByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetOrdersByApp(ctx, req.(*GetOrdersByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetOrdersByGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersByGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetOrdersByGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetOrdersByGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetOrdersByGood(ctx, req.(*GetOrdersByGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_UpdatePasswordByAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordByAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).UpdatePasswordByAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePasswordByAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).UpdatePasswordByAppUser(ctx, req.(*UpdatePasswordByAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_UpdateEmailAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).UpdateEmailAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/UpdateEmailAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).UpdateEmailAddress(ctx, req.(*UpdateEmailAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_UpdatePhoneNO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhoneNORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).UpdatePhoneNO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/UpdatePhoneNO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).UpdatePhoneNO(ctx, req.(*UpdatePhoneNORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetMyInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetMyInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetMyInvitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetMyInvitations(ctx, req.(*GetMyInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetMyDirectInvitations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyDirectInvitationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetMyDirectInvitations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetMyDirectInvitations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetMyDirectInvitations(ctx, req.(*GetMyDirectInvitationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetKycReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetKycReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetKycReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetKycReviews(ctx, req.(*GetKycReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetGoodReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetGoodReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetGoodReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetGoodReviews(ctx, req.(*GetGoodReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_CreateKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).CreateKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/CreateKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).CreateKyc(ctx, req.(*CreateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_UpdateKyc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKycRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).UpdateKyc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/UpdateKyc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).UpdateKyc(ctx, req.(*UpdateKycRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingApis_GetKycByAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKycByAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingApisServer).GetKycByAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.apis.v1.CloudHashingApis/GetKycByAppUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingApisServer).GetKycByAppUser(ctx, req.(*GetKycByAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudHashingApis_ServiceDesc is the grpc.ServiceDesc for CloudHashingApis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudHashingApis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.hashing.apis.v1.CloudHashingApis",
	HandlerType: (*CloudHashingApisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _CloudHashingApis_Version_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _CloudHashingApis_GetGoods_Handler,
		},
		{
			MethodName: "CreateGood",
			Handler:    _CloudHashingApis_CreateGood_Handler,
		},
		{
			MethodName: "GetGood",
			Handler:    _CloudHashingApis_GetGood_Handler,
		},
		{
			MethodName: "GetRecommendGoodsByApp",
			Handler:    _CloudHashingApis_GetRecommendGoodsByApp_Handler,
		},
		{
			MethodName: "SubmitOrder",
			Handler:    _CloudHashingApis_SubmitOrder_Handler,
		},
		{
			MethodName: "CreateOrderPayment",
			Handler:    _CloudHashingApis_CreateOrderPayment_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _CloudHashingApis_GetOrder_Handler,
		},
		{
			MethodName: "GetOrdersByAppUser",
			Handler:    _CloudHashingApis_GetOrdersByAppUser_Handler,
		},
		{
			MethodName: "GetOrdersByApp",
			Handler:    _CloudHashingApis_GetOrdersByApp_Handler,
		},
		{
			MethodName: "GetOrdersByGood",
			Handler:    _CloudHashingApis_GetOrdersByGood_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _CloudHashingApis_Signup_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _CloudHashingApis_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdatePasswordByAppUser",
			Handler:    _CloudHashingApis_UpdatePasswordByAppUser_Handler,
		},
		{
			MethodName: "UpdateEmailAddress",
			Handler:    _CloudHashingApis_UpdateEmailAddress_Handler,
		},
		{
			MethodName: "UpdatePhoneNO",
			Handler:    _CloudHashingApis_UpdatePhoneNO_Handler,
		},
		{
			MethodName: "GetMyInvitations",
			Handler:    _CloudHashingApis_GetMyInvitations_Handler,
		},
		{
			MethodName: "GetMyDirectInvitations",
			Handler:    _CloudHashingApis_GetMyDirectInvitations_Handler,
		},
		{
			MethodName: "GetKycReviews",
			Handler:    _CloudHashingApis_GetKycReviews_Handler,
		},
		{
			MethodName: "GetGoodReviews",
			Handler:    _CloudHashingApis_GetGoodReviews_Handler,
		},
		{
			MethodName: "CreateKyc",
			Handler:    _CloudHashingApis_CreateKyc_Handler,
		},
		{
			MethodName: "UpdateKyc",
			Handler:    _CloudHashingApis_UpdateKyc_Handler,
		},
		{
			MethodName: "GetKycByAppUser",
			Handler:    _CloudHashingApis_GetKycByAppUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/cloud-hashing-apis/cloud-hashing-apis.proto",
}
