/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as AppUserManagerV1Appusermgr from "../appusermgr/appusermgr.pb"
import * as CloudHashingBillingV1Cloud-hashing-billing from "../cloud-hashing-billing/cloud-hashing-billing.pb"
import * as CloudHashingGoodsV1Cloud-hashing-goods from "../cloud-hashing-goods/cloud-hashing-goods.pb"
import * as CloudHashingInspireV1Cloud-hashing-inspire from "../cloud-hashing-inspire/cloud-hashing-inspire.pb"
import * as CloudHashingOrderV1Cloud-hashing-order from "../cloud-hashing-order/cloud-hashing-order.pb"
import * as SphinxCoininfoV1Coininfo from "../coininfo/coininfo.pb"
import * as KycManagementV1Kyc-management from "../kyc/kyc-management.pb"
import * as NpoolV1Npool from "../npool.pb"
import * as ReviewServiceV1Review-service from "../review-service/review-service.pb"
export type Good = {
  good?: CloudHashingGoodsV1Cloud-hashing-goods.GoodDetail
  main?: SphinxCoininfoV1Coininfo.CoinInfo
  supportCoins?: SphinxCoininfoV1Coininfo.CoinInfo[]
  reviews?: ReviewServiceV1Review-service.Review[]
}

export type RecommendGood = {
  good?: Good
  recommend?: CloudHashingGoodsV1Cloud-hashing-goods.Recommend
}

export type CreateGoodRequest = {
  info?: CloudHashingGoodsV1Cloud-hashing-goods.GoodInfo
}

export type CreateGoodResponse = {
  info?: Good
}

export type GetGoodRequest = {
  id?: string
}

export type GetGoodResponse = {
  info?: Good
}

export type GetGoodsRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetGoodsResponse = {
  infos?: Good[]
  total?: number
}

export type GetRecommendGoodsByAppRequest = {
  appID?: string
}

export type GetRecommendGoodsByAppResponse = {
  infos?: RecommendGood[]
}

export type Order = {
  order?: CloudHashingOrderV1Cloud-hashing-order.OrderDetail
  payToAccount?: CloudHashingBillingV1Cloud-hashing-billing.CoinAccountInfo
  good?: Good
  payWithCoin?: SphinxCoininfoV1Coininfo.CoinInfo
  fixAmountCoupon?: CloudHashingInspireV1Cloud-hashing-inspire.CouponAllocatedDetail
  discountCoupon?: CloudHashingInspireV1Cloud-hashing-inspire.CouponAllocatedDetail
  userSpecialReduction?: CloudHashingInspireV1Cloud-hashing-inspire.UserSpecialReduction
}

export type GetOrderRequest = {
  id?: string
}

export type GetOrderResponse = {
  info?: Order
}

export type GetOrdersByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetOrdersByAppUserResponse = {
  infos?: Order[]
}

export type GetOrdersByAppRequest = {
  appID?: string
}

export type GetOrdersByAppResponse = {
  infos?: Order[]
}

export type GetOrdersByGoodRequest = {
  goodID?: string
}

export type GetOrdersByGoodResponse = {
  infos?: Order[]
}

export type OrderFee = {
  id?: string
  durationDays?: number
}

export type SubmitOrderRequest = {
  goodID?: string
  units?: number
  userID?: string
  appID?: string
  couponID?: string
  discountCouponID?: string
  userSpecialReductionID?: string
}

export type SubmitOrderResponse = {
  info?: Order
}

export type CreateOrderPaymentRequest = {
  orderID?: string
  paymentCoinTypeID?: string
  fees?: OrderFee[]
}

export type CreateOrderPaymentResponse = {
  info?: Order
}

export type SignupRequest = {
  username?: string
  passwordHash?: string
  account?: string
  accountType?: string
  verificationCode?: string
  appID?: string
  invitationCode?: string
}

export type SignupResponse = {
  info?: AppUserManagerV1Appusermgr.AppUser
}

export type UpdatePasswordRequest = {
  appID?: string
  account?: string
  accountType?: string
  passwordHash?: string
  verificationCode?: string
}

export type UpdatePasswordResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserSecret
}

export type UpdatePasswordByAppUserRequest = {
  appID?: string
  userID?: string
  account?: string
  accountType?: string
  oldPasswordHash?: string
  passwordHash?: string
  verificationCode?: string
}

export type UpdatePasswordByAppUserResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserSecret
}

export type InvitationSummary = {
  units?: number
  amount?: number
}

export type InvitationUserInfo = {
  userID?: string
  username?: string
  avatar?: string
  emailAddress?: string
  kol?: boolean
  summarys?: {[key: string]: InvitationSummary}
  mySummarys?: {[key: string]: InvitationSummary}
  invitedCount?: number
  joinDate?: number
}

export type Invitation = {
  invitees?: InvitationUserInfo[]
}

export type GetMyInvitationsRequest = {
  appID?: string
  userID?: string
}

export type GetMyInvitationsResponse = {
  mySelf?: InvitationUserInfo
  infos?: {[key: string]: Invitation}
}

export type GetMyDirectInvitationsRequest = {
  appID?: string
  userID?: string
}

export type GetMyDirectInvitationsResponse = {
  mySelf?: InvitationUserInfo
  infos?: {[key: string]: Invitation}
}

export type KycReview = {
  review?: ReviewServiceV1Review-service.Review
  user?: AppUserManagerV1Appusermgr.AppUserInfo
  kyc?: KycManagementV1Kyc-management.KycInfo
}

export type GetKycReviewsRequest = {
  appID?: string
}

export type GetKycReviewsResponse = {
  infos?: KycReview[]
}

export type GoodReview = {
  review?: ReviewServiceV1Review-service.Review
  createdBy?: AppUserManagerV1Appusermgr.AppUserInfo
  good?: CloudHashingGoodsV1Cloud-hashing-goods.GoodDetail
}

export type GetGoodReviewsRequest = {
}

export type GetGoodReviewsResponse = {
  infos?: GoodReview[]
}

export type Kyc = {
  kyc?: KycManagementV1Kyc-management.KycInfo
  state?: string
  message?: string
}

export type GetKycByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetKycByAppUserResponse = {
  info?: Kyc
}

export type CreateKycRequest = {
  info?: KycManagementV1Kyc-management.KycInfo
}

export type CreateKycResponse = {
  info?: Kyc
}

export type UpdateKycRequest = {
  info?: KycManagementV1Kyc-management.KycInfo
}

export type UpdateKycResponse = {
  info?: Kyc
}

export class CloudHashingApis {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoods(req: GetGoodsRequest, initReq?: fm.InitReq): Promise<GetGoodsResponse> {
    return fm.fetchReq<GetGoodsRequest, GetGoodsResponse>(`/v1/get/goods`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGood(req: CreateGoodRequest, initReq?: fm.InitReq): Promise<CreateGoodResponse> {
    return fm.fetchReq<CreateGoodRequest, CreateGoodResponse>(`/v1/create/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static UpdatePassword(req: UpdatePasswordRequest, initReq?: fm.InitReq): Promise<UpdatePasswordResponse> {
    return fm.fetchReq<UpdatePasswordRequest, UpdatePasswordResponse>(`/v1/update/password`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdatePasswordByAppUser(req: UpdatePasswordByAppUserRequest, initReq?: fm.InitReq): Promise<UpdatePasswordByAppUserResponse> {
    return fm.fetchReq<UpdatePasswordByAppUserRequest, UpdatePasswordByAppUserResponse>(`/v1/update/password/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMyInvitations(req: GetMyInvitationsRequest, initReq?: fm.InitReq): Promise<GetMyInvitationsResponse> {
    return fm.fetchReq<GetMyInvitationsRequest, GetMyInvitationsResponse>(`/v1/get/my/invitations`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMyDirectInvitations(req: GetMyDirectInvitationsRequest, initReq?: fm.InitReq): Promise<GetMyDirectInvitationsResponse> {
    return fm.fetchReq<GetMyDirectInvitationsRequest, GetMyDirectInvitationsResponse>(`/v1/get/my/direct/invitations`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycReviews(req: GetKycReviewsRequest, initReq?: fm.InitReq): Promise<GetKycReviewsResponse> {
    return fm.fetchReq<GetKycReviewsRequest, GetKycReviewsResponse>(`/v1/get/kyc/reviews`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodReviews(req: GetGoodReviewsRequest, initReq?: fm.InitReq): Promise<GetGoodReviewsResponse> {
    return fm.fetchReq<GetGoodReviewsRequest, GetGoodReviewsResponse>(`/v1/get/good/reviews`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateKyc(req: CreateKycRequest, initReq?: fm.InitReq): Promise<CreateKycResponse> {
    return fm.fetchReq<CreateKycRequest, CreateKycResponse>(`/v1/create/kyc`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateKyc(req: UpdateKycRequest, initReq?: fm.InitReq): Promise<UpdateKycResponse> {
    return fm.fetchReq<UpdateKycRequest, UpdateKycResponse>(`/v1/update/kyc`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycByAppUser(req: GetKycByAppUserRequest, initReq?: fm.InitReq): Promise<GetKycByAppUserResponse> {
    return fm.fetchReq<GetKycByAppUserRequest, GetKycByAppUserResponse>(`/v1/get/kyc/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}