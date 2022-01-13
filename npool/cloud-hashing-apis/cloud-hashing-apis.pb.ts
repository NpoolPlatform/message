/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as CloudHashingGoodsV1Cloud-hashing-goods from "../cloud-hashing-goods/cloud-hashing-goods.pb"
import * as CloudHashingOrderV1Cloud-hashing-order from "../cloud-hashing-order/cloud-hashing-order.pb"
import * as SphinxCoininfoV1Coininfo from "../coininfo/coininfo.pb"
import * as NpoolV1Npool from "../npool.pb"
import * as ReviewServiceV1Review-service from "../review-service/review-service.pb"
import * as UserV1User-management from "../user/user-management.pb"
export type Good = {
  Good?: CloudHashingGoodsV1Cloud-hashing-goods.GoodDetail
  Main?: SphinxCoininfoV1Coininfo.CoinInfo
  SupportCoins?: SphinxCoininfoV1Coininfo.CoinInfo[]
  Review?: ReviewServiceV1Review-service.Review
}

export type RecommendGood = {
  Good?: Good
  Recommend?: CloudHashingGoodsV1Cloud-hashing-goods.Recommend
}

export type GetGoodRequest = {
  ID?: string
}

export type GetGoodResponse = {
  Info?: Good
}

export type GetGoodsRequest = {
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetGoodsResponse = {
  Infos?: Good[]
  Total?: number
}

export type GetRecommendGoodsByAppRequest = {
  AppID?: string
}

export type GetRecommendGoodsByAppResponse = {
  Infos?: RecommendGood[]
}

export type GetOrderRequest = {
  ID?: string
}

export type GetOrderResponse = {
  Info?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail
}

export type GetOrdersByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetOrdersByAppUserResponse = {
  Infos?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail[]
}

export type GetOrdersByAppRequest = {
  AppID?: string
}

export type GetOrdersByAppResponse = {
  Infos?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail[]
}

export type GetOrdersByGoodRequest = {
  GoodID?: string
}

export type GetOrdersByGoodResponse = {
  Infos?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail[]
}

export type OrderFee = {
  ID?: string
  DurationDays?: number
}

export type SubmitOrderRequest = {
  GoodID?: string
  Units?: number
  UserID?: string
  AppID?: string
  CouponID?: string
  DiscountCouponID?: string
  UserSpecialReductionID?: string
}

export type SubmitOrderResponse = {
  Info?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail
}

export type CreateOrderPaymentRequest = {
  OrderID?: string
  PaymentCoinTypeID?: string
  Fees?: OrderFee[]
}

export type CreateOrderPaymentResponse = {
  Info?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail
}

export type SignupRequest = {
  Username?: string
  Password?: string
  EmailAddress?: string
  PhoneNumber?: string
  VerificationCode?: string
  AppID?: string
  InvitationCode?: string
}

export type SignupResponse = {
  Info?: UserV1User-management.UserBasicInfo
}

export type InvitationSummary = {
  Units?: number
  Amount?: number
}

export type InvitationUserInfo = {
  UserID?: string
  Username?: string
  Avatar?: string
  EmailAddress?: string
  Kol?: boolean
  Summarys?: {[key: string]: InvitationSummary}
  MySummarys?: {[key: string]: InvitationSummary}
  InvitedCount?: number
  JoinDate?: number
}

export type Invitation = {
  Invitees?: InvitationUserInfo[]
}

export type GetMyInvitationsRequest = {
  AppID?: string
  InviterID?: string
}

export type GetMyInvitationsResponse = {
  MySelf?: InvitationUserInfo
  Infos?: {[key: string]: Invitation}
}

export type GetMyDirectInvitationsRequest = {
  AppID?: string
  InviterID?: string
}

export type GetMyDirectInvitationsResponse = {
  MySelf?: InvitationUserInfo
  Infos?: {[key: string]: Invitation}
}

export type KycReview = {
  Review?: ReviewServiceV1Review-service.Review
}

export type GetKycReviewsRequest = {
  AppID?: string
}

export type GetKycReviewsResponse = {
  Infos?: KycReview[]
}

export type GoodReview = {
  Review?: ReviewServiceV1Review-service.Review
  User?: UserV1User-management.UserBasicInfo
  Good?: CloudHashingGoodsV1Cloud-hashing-goods.GoodDetail
}

export type GetGoodReviewsRequest = {
}

export type GetGoodReviewsResponse = {
  Infos?: GoodReview[]
}

export class CloudHashingApis {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoods(req: GetGoodsRequest, initReq?: fm.InitReq): Promise<GetGoodsResponse> {
    return fm.fetchReq<GetGoodsRequest, GetGoodsResponse>(`/v1/get/goods`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGood(req: GetGoodRequest, initReq?: fm.InitReq): Promise<GetGoodResponse> {
    return fm.fetchReq<GetGoodRequest, GetGoodResponse>(`/v1/get/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRecommendGoodsByApp(req: GetRecommendGoodsByAppRequest, initReq?: fm.InitReq): Promise<GetRecommendGoodsByAppResponse> {
    return fm.fetchReq<GetRecommendGoodsByAppRequest, GetRecommendGoodsByAppResponse>(`/v1/get/recommend/goods/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SubmitOrder(req: SubmitOrderRequest, initReq?: fm.InitReq): Promise<SubmitOrderResponse> {
    return fm.fetchReq<SubmitOrderRequest, SubmitOrderResponse>(`/v1/submit/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateOrderPayment(req: CreateOrderPaymentRequest, initReq?: fm.InitReq): Promise<CreateOrderPaymentResponse> {
    return fm.fetchReq<CreateOrderPaymentRequest, CreateOrderPaymentResponse>(`/v1/create/order/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrder(req: GetOrderRequest, initReq?: fm.InitReq): Promise<GetOrderResponse> {
    return fm.fetchReq<GetOrderRequest, GetOrderResponse>(`/v1/get/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByAppUser(req: GetOrdersByAppUserRequest, initReq?: fm.InitReq): Promise<GetOrdersByAppUserResponse> {
    return fm.fetchReq<GetOrdersByAppUserRequest, GetOrdersByAppUserResponse>(`/v1/get/orders/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByApp(req: GetOrdersByAppRequest, initReq?: fm.InitReq): Promise<GetOrdersByAppResponse> {
    return fm.fetchReq<GetOrdersByAppRequest, GetOrdersByAppResponse>(`/v1/get/orders/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByGood(req: GetOrdersByGoodRequest, initReq?: fm.InitReq): Promise<GetOrdersByGoodResponse> {
    return fm.fetchReq<GetOrdersByGoodRequest, GetOrdersByGoodResponse>(`/v1/get/orders/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Signup(req: SignupRequest, initReq?: fm.InitReq): Promise<SignupResponse> {
    return fm.fetchReq<SignupRequest, SignupResponse>(`/v1/signup`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMyInvitations(req: GetMyInvitationsRequest, initReq?: fm.InitReq): Promise<GetMyInvitationsResponse> {
    return fm.fetchReq<GetMyInvitationsRequest, GetMyInvitationsResponse>(`/v1/get/my/invitations`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMyDirectInvitations(req: GetMyDirectInvitationsRequest, initReq?: fm.InitReq): Promise<GetMyDirectInvitationsResponse> {
    return fm.fetchReq<GetMyDirectInvitationsRequest, GetMyDirectInvitationsResponse>(`/v1/get/my/direct/invitations`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}