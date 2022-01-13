/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  Info?: string
}

export type NewUserRewardSetting = {
  ID?: string
  AppID?: string
  RegistrationCouponID?: string
  KycCouponID?: string
  AutoGenerateInvitationCode?: boolean
}

export type CreateNewUserRewardSettingRequest = {
  Info?: NewUserRewardSetting
}

export type CreateNewUserRewardSettingResponse = {
  Info?: NewUserRewardSetting
}

export type GetNewUserRewardSettingRequest = {
  ID?: string
}

export type GetNewUserRewardSettingResponse = {
  Info?: NewUserRewardSetting
}

export type GetNewUserRewardSettingByAppRequest = {
  AppID?: string
}

export type GetNewUserRewardSettingByAppResponse = {
  Info?: NewUserRewardSetting
}

export type UpdateNewUserRewardSettingRequest = {
  Info?: NewUserRewardSetting
}

export type UpdateNewUserRewardSettingResponse = {
  Info?: NewUserRewardSetting
}

export type AgencySetting = {
  ID?: string
  AppID?: string
  GoodID?: string
  RegistrationRewardThreshold?: number
  RegistrationCouponID?: string
  KycRewardThreshold?: number
  KycCouponID?: string
  TotalPurchaseRewardPercent?: number
  PurchaseRewardChainLevels?: number
  LevelPurchaseRewardPercent?: number
}

export type CreateAgencySettingRequest = {
  Info?: AgencySetting
}

export type CreateAgencySettingResponse = {
  Info?: AgencySetting
}

export type GetAgencySettingRequest = {
  ID?: string
}

export type GetAgencySettingResponse = {
  Info?: AgencySetting
}

export type GetAgencySettingByAppRequest = {
  AppID?: string
}

export type GetAgencySettingByAppResponse = {
  Info?: AgencySetting
}

export type UpdateAgencySettingRequest = {
  Info?: AgencySetting
}

export type UpdateAgencySettingResponse = {
  Info?: AgencySetting
}

export type PurchaseInvitation = {
  ID?: string
  AppID?: string
  OrderID?: string
  InvitationCodeID?: string
}

export type CreatePurchaseInvitationRequest = {
  Info?: PurchaseInvitation
}

export type CreatePurchaseInvitationResponse = {
  Info?: PurchaseInvitation
}

export type UpdatePurchaseInvitationRequest = {
  Info?: PurchaseInvitation
}

export type UpdatePurchaseInvitationResponse = {
  Info?: PurchaseInvitation
}

export type GetPurchaseInvitationRequest = {
  ID?: string
}

export type GetPurchaseInvitationResponse = {
  Info?: PurchaseInvitation
}

export type GetPurchaseInvitationsByAppRequest = {
  AppID?: string
}

export type GetPurchaseInvitationsByAppResponse = {
  Infos?: PurchaseInvitation[]
}

export type GetPurchaseInvitationByAppOrderRequest = {
  AppID?: string
  OrderID?: string
}

export type GetPurchaseInvitationByAppOrderResponse = {
  Info?: PurchaseInvitation
}

export type RegistrationInvitation = {
  ID?: string
  AppID?: string
  InviterID?: string
  InviteeID?: string
  CreateAt?: number
}

export type CreateRegistrationInvitationRequest = {
  Info?: RegistrationInvitation
}

export type CreateRegistrationInvitationResponse = {
  Info?: RegistrationInvitation
}

export type UpdateRegistrationInvitationRequest = {
  Info?: RegistrationInvitation
}

export type UpdateRegistrationInvitationResponse = {
  Info?: RegistrationInvitation
}

export type GetRegistrationInvitationRequest = {
  ID?: string
}

export type GetRegistrationInvitationResponse = {
  Info?: RegistrationInvitation
}

export type GetRegistrationInvitationsByAppRequest = {
  AppID?: string
}

export type GetRegistrationInvitationsByAppResponse = {
  Infos?: RegistrationInvitation[]
}

export type GetRegistrationInvitationsByAppInviterRequest = {
  AppID?: string
  InviterID?: string
}

export type GetRegistrationInvitationsByAppInviterResponse = {
  Infos?: RegistrationInvitation[]
}

export type GetRegistrationInvitationByAppInviteeRequest = {
  AppID?: string
  InviteeID?: string
}

export type GetRegistrationInvitationByAppInviteeResponse = {
  Info?: RegistrationInvitation
}

export type UserInvitationCode = {
  ID?: string
  UserID?: string
  AppID?: string
  InvitationCode?: string
}

export type CreateUserInvitationCodeRequest = {
  Info?: UserInvitationCode
}

export type CreateUserInvitationCodeResponse = {
  Info?: UserInvitationCode
}

export type GetUserInvitationCodeRequest = {
  ID?: string
}

export type GetUserInvitationCodeResponse = {
  Info?: UserInvitationCode
}

export type GetUserInvitationCodeByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserInvitationCodeByAppUserResponse = {
  Info?: UserInvitationCode
}

export type GetUserInvitationCodeByCodeRequest = {
  Code?: string
}

export type GetUserInvitationCodeByCodeResponse = {
  Info?: UserInvitationCode
}

export type CouponAllocated = {
  ID?: string
  UserID?: string
  AppID?: string
  Type?: string
  CouponID?: string
}

export type CreateCouponAllocatedRequest = {
  Info?: CouponAllocated
}

export type CreateCouponAllocatedResponse = {
  Info?: CouponAllocated
}

export type GetCouponAllocatedRequest = {
  ID?: string
}

export type GetCouponAllocatedResponse = {
  Info?: CouponAllocated
}

export type GetCouponsAllocatedByAppRequest = {
  AppID?: string
}

export type GetCouponsAllocatedByAppResponse = {
  Infos?: CouponAllocated[]
}

export type GetCouponsAllocatedByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetCouponsAllocatedByAppUserResponse = {
  Infos?: CouponAllocated[]
}

export type UpdateCouponAllocatedRequest = {
  Info?: CouponAllocated
}

export type UpdateCouponAllocatedResponse = {
  Info?: CouponAllocated
}

export type CouponPool = {
  ID?: string
  AppID?: string
  ReleaseByUserID?: string
  Denomination?: number
  Circulation?: number
  Start?: number
  DurationDays?: number
  Message?: string
  Name?: string
}

export type CreateCouponPoolRequest = {
  Info?: CouponPool
}

export type CreateCouponPoolResponse = {
  Info?: CouponPool
}

export type UpdateCouponPoolRequest = {
  Info?: CouponPool
}

export type UpdateCouponPoolResponse = {
  Info?: CouponPool
}

export type GetCouponPoolRequest = {
  ID?: string
}

export type GetCouponPoolResponse = {
  Info?: CouponPool
}

export type GetCouponPoolsByAppRequest = {
  AppID?: string
}

export type GetCouponPoolsByAppResponse = {
  Infos?: CouponPool[]
}

export type GetCouponPoolsByAppReleaserRequest = {
  AppID?: string
  UserID?: string
}

export type GetCouponPoolsByAppReleaserResponse = {
  Infos?: CouponPool[]
}

export type NewUserRewardSettingDetail = {
  ID?: string
  AppID?: string
  AutoGenerateInvitationCode?: boolean
  RegistrationCoupon?: CouponPool
  KycCoupon?: CouponPool
}

export type GetNewUserRewardSettingDetailRequest = {
  ID?: string
}

export type GetNewUserRewardSettingDetailResponse = {
  Info?: NewUserRewardSettingDetail
}

export type AppCouponSetting = {
  ID?: string
  AppID?: string
  DominationLimit?: number
  TotalLimit?: number
}

export type CreateAppCouponSettingRequest = {
  Info?: AppCouponSetting
}

export type CreateAppCouponSettingResponse = {
  Info?: AppCouponSetting
}

export type GetAppCouponSettingRequest = {
  ID?: string
}

export type GetAppCouponSettingResponse = {
  Info?: AppCouponSetting
}

export type GetAppCouponSettingByAppRequest = {
  AppID?: string
}

export type GetAppCouponSettingByAppResponse = {
  Info?: AppCouponSetting
}

export type UpdateAppCouponSettingRequest = {
  Info?: AppCouponSetting
}

export type UpdateAppCouponSettingResponse = {
  Info?: AppCouponSetting
}

export type DefaultKpiSetting = {
  ID?: string
  AppID?: string
  GoodID?: string
  Amount?: number
  Percent?: number
}

export type CreateDefaultKpiSettingRequest = {
  Info?: DefaultKpiSetting
}

export type CreateDefaultKpiSettingResponse = {
  Info?: DefaultKpiSetting
}

export type GetDefaultKpiSettingRequest = {
  ID?: string
}

export type GetDefaultKpiSettingResponse = {
  Info?: DefaultKpiSetting
}

export type GetDefaultKpiSettingByAppGoodRequest = {
  AppID?: string
  GoodID?: string
}

export type GetDefaultKpiSettingByAppGoodResponse = {
  Info?: DefaultKpiSetting
}

export type UpdateDefaultKpiSettingRequest = {
  Info?: DefaultKpiSetting
}

export type UpdateDefaultKpiSettingResponse = {
  Info?: DefaultKpiSetting
}

export type UserKpiSetting = {
  ID?: string
  AppID?: string
  UserID?: string
  GoodID?: string
  Amount?: number
  Percent?: number
}

export type CreateUserKpiSettingRequest = {
  Info?: UserKpiSetting
}

export type CreateUserKpiSettingResponse = {
  Info?: UserKpiSetting
}

export type GetUserKpiSettingRequest = {
  ID?: string
}

export type GetUserKpiSettingResponse = {
  Info?: UserKpiSetting
}

export type GetUserKpiSettingByAppGoodRequest = {
  AppID?: string
  GoodID?: string
}

export type GetUserKpiSettingByAppGoodResponse = {
  Info?: UserKpiSetting
}

export type UpdateUserKpiSettingRequest = {
  Info?: UserKpiSetting
}

export type UpdateUserKpiSettingResponse = {
  Info?: UserKpiSetting
}

export type AgencySettingDetail = {
  ID?: string
  AppID?: string
  GoodID?: string
  RegistrationRewardThreshold?: number
  RegistrationCoupon?: CouponPool
  KycRewardThreshold?: number
  KycCoupon?: CouponPool
  TotalPurchaseRewardPercent?: number
  PurchaseRewardChainLevels?: number
  LevelPurchaseRewardPercent?: number
}

export type GetAgencySettingDetailRequest = {
  ID?: string
}

export type GetAgencySettingDetailResponse = {
  Info?: AgencySettingDetail
}

export type GetAgencySettingDetailByAppRequest = {
  AppID?: string
}

export type GetAgencySettingDetailByAppResponse = {
  Info?: AgencySettingDetail
}

export type CouponAllocatedDetail = {
  ID?: string
  UserID?: string
  AppID?: string
  Type?: string
  Coupon?: CouponPool
  Discount?: DiscountPool
}

export type GetCouponAllocatedDetailRequest = {
  ID?: string
}

export type GetCouponAllocatedDetailResponse = {
  Info?: CouponAllocatedDetail
}

export type GetCouponsAllocatedDetailByAppRequest = {
  AppID?: string
}

export type GetCouponsAllocatedDetailByAppResponse = {
  Infos?: CouponAllocatedDetail[]
}

export type GetCouponsAllocatedDetailByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetCouponsAllocatedDetailByAppUserResponse = {
  Infos?: CouponAllocatedDetail[]
}

export type DiscountPool = {
  ID?: string
  AppID?: string
  ReleaseByUserID?: string
  Discount?: number
  Start?: number
  DurationDays?: number
  Message?: string
  Name?: string
}

export type CreateDiscountPoolRequest = {
  Info?: DiscountPool
}

export type CreateDiscountPoolResponse = {
  Info?: DiscountPool
}

export type UpdateDiscountPoolRequest = {
  Info?: DiscountPool
}

export type UpdateDiscountPoolResponse = {
  Info?: DiscountPool
}

export type GetDiscountPoolRequest = {
  ID?: string
}

export type GetDiscountPoolResponse = {
  Info?: DiscountPool
}

export type GetDiscountPoolsByAppRequest = {
  AppID?: string
}

export type GetDiscountPoolsByAppResponse = {
  Infos?: DiscountPool[]
}

export type GetDiscountPoolsByAppReleaserRequest = {
  AppID?: string
  UserID?: string
}

export type GetDiscountPoolsByAppReleaserResponse = {
  Infos?: DiscountPool[]
}

export type UserSpecialReduction = {
  ID?: string
  AppID?: string
  UserID?: string
  Amount?: number
  ReleaseByUserID?: string
  Start?: number
  DurationDays?: number
  Message?: string
}

export type CreateUserSpecialReductionRequest = {
  Info?: UserSpecialReduction
}

export type CreateUserSpecialReductionResponse = {
  Info?: UserSpecialReduction
}

export type GetUserSpecialReductionRequest = {
  ID?: string
}

export type GetUserSpecialReductionResponse = {
  Info?: UserSpecialReduction
}

export type GetUserSpecialReductionsByAppRequest = {
  AppID?: string
}

export type GetUserSpecialReductionsByAppResponse = {
  Infos?: UserSpecialReduction[]
}

export type GetUserSpecialReductionsByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserSpecialReductionsByAppUserResponse = {
  Infos?: UserSpecialReduction[]
}

export type UpdateUserSpecialReductionRequest = {
  Info?: UserSpecialReduction
}

export type UpdateUserSpecialReductionResponse = {
  Info?: UserSpecialReduction
}

export type GetUserSpecialReductionsByAppReleaserRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserSpecialReductionsByAppReleaserResponse = {
  Infos?: UserSpecialReduction[]
}

export class CloudHashingInspire {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateNewUserRewardSetting(req: CreateNewUserRewardSettingRequest, initReq?: fm.InitReq): Promise<CreateNewUserRewardSettingResponse> {
    return fm.fetchReq<CreateNewUserRewardSettingRequest, CreateNewUserRewardSettingResponse>(`/v1/create/new/user/reward/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetNewUserRewardSetting(req: GetNewUserRewardSettingRequest, initReq?: fm.InitReq): Promise<GetNewUserRewardSettingResponse> {
    return fm.fetchReq<GetNewUserRewardSettingRequest, GetNewUserRewardSettingResponse>(`/v1/get/new/user/reward/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetNewUserRewardSettingDetail(req: GetNewUserRewardSettingDetailRequest, initReq?: fm.InitReq): Promise<GetNewUserRewardSettingDetailResponse> {
    return fm.fetchReq<GetNewUserRewardSettingDetailRequest, GetNewUserRewardSettingDetailResponse>(`/v1/get/new/user/reward/setting/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetNewUserRewardSettingByApp(req: GetNewUserRewardSettingByAppRequest, initReq?: fm.InitReq): Promise<GetNewUserRewardSettingByAppResponse> {
    return fm.fetchReq<GetNewUserRewardSettingByAppRequest, GetNewUserRewardSettingByAppResponse>(`/v1/get/new/user/reward/setting/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateNewUserRewardSetting(req: UpdateNewUserRewardSettingRequest, initReq?: fm.InitReq): Promise<UpdateNewUserRewardSettingResponse> {
    return fm.fetchReq<UpdateNewUserRewardSettingRequest, UpdateNewUserRewardSettingResponse>(`/v1/get/update/user/reward/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAgencySetting(req: CreateAgencySettingRequest, initReq?: fm.InitReq): Promise<CreateAgencySettingResponse> {
    return fm.fetchReq<CreateAgencySettingRequest, CreateAgencySettingResponse>(`/v1/create/agency/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAgencySetting(req: GetAgencySettingRequest, initReq?: fm.InitReq): Promise<GetAgencySettingResponse> {
    return fm.fetchReq<GetAgencySettingRequest, GetAgencySettingResponse>(`/v1/get/agency/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAgencySettingByApp(req: GetAgencySettingByAppRequest, initReq?: fm.InitReq): Promise<GetAgencySettingByAppResponse> {
    return fm.fetchReq<GetAgencySettingByAppRequest, GetAgencySettingByAppResponse>(`/v1/get/agency/setting/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAgencySetting(req: UpdateAgencySettingRequest, initReq?: fm.InitReq): Promise<UpdateAgencySettingResponse> {
    return fm.fetchReq<UpdateAgencySettingRequest, UpdateAgencySettingResponse>(`/v1/update/agency/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAgencySettingDetail(req: GetAgencySettingDetailRequest, initReq?: fm.InitReq): Promise<GetAgencySettingDetailResponse> {
    return fm.fetchReq<GetAgencySettingDetailRequest, GetAgencySettingDetailResponse>(`/v1/get/agency/setting/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAgencySettingDetailByApp(req: GetAgencySettingDetailByAppRequest, initReq?: fm.InitReq): Promise<GetAgencySettingDetailByAppResponse> {
    return fm.fetchReq<GetAgencySettingDetailByAppRequest, GetAgencySettingDetailByAppResponse>(`/v1/get/agency/setting/detail/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreatePurchaseInvitation(req: CreatePurchaseInvitationRequest, initReq?: fm.InitReq): Promise<CreatePurchaseInvitationResponse> {
    return fm.fetchReq<CreatePurchaseInvitationRequest, CreatePurchaseInvitationResponse>(`/v1/create/purchase/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdatePurchaseInvitation(req: UpdatePurchaseInvitationRequest, initReq?: fm.InitReq): Promise<UpdatePurchaseInvitationResponse> {
    return fm.fetchReq<UpdatePurchaseInvitationRequest, UpdatePurchaseInvitationResponse>(`/v1/update/purchase/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPurchaseInvitation(req: GetPurchaseInvitationRequest, initReq?: fm.InitReq): Promise<GetPurchaseInvitationResponse> {
    return fm.fetchReq<GetPurchaseInvitationRequest, GetPurchaseInvitationResponse>(`/v1/get/purchase/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPurchaseInvitationsByApp(req: GetPurchaseInvitationsByAppRequest, initReq?: fm.InitReq): Promise<GetPurchaseInvitationsByAppResponse> {
    return fm.fetchReq<GetPurchaseInvitationsByAppRequest, GetPurchaseInvitationsByAppResponse>(`/v1/get/purchase/invitations/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPurchaseInvitationByAppOrder(req: GetPurchaseInvitationByAppOrderRequest, initReq?: fm.InitReq): Promise<GetPurchaseInvitationByAppOrderResponse> {
    return fm.fetchReq<GetPurchaseInvitationByAppOrderRequest, GetPurchaseInvitationByAppOrderResponse>(`/v1/get/purchase/invitation/by/app/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateRegistrationInvitation(req: CreateRegistrationInvitationRequest, initReq?: fm.InitReq): Promise<CreateRegistrationInvitationResponse> {
    return fm.fetchReq<CreateRegistrationInvitationRequest, CreateRegistrationInvitationResponse>(`/v1/create/registration/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateRegistrationInvitation(req: UpdateRegistrationInvitationRequest, initReq?: fm.InitReq): Promise<UpdateRegistrationInvitationResponse> {
    return fm.fetchReq<UpdateRegistrationInvitationRequest, UpdateRegistrationInvitationResponse>(`/v1/update/registration/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRegistrationInvitation(req: GetRegistrationInvitationRequest, initReq?: fm.InitReq): Promise<GetRegistrationInvitationResponse> {
    return fm.fetchReq<GetRegistrationInvitationRequest, GetRegistrationInvitationResponse>(`/v1/get/registration/invitation`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRegistrationInvitationsByApp(req: GetRegistrationInvitationsByAppRequest, initReq?: fm.InitReq): Promise<GetRegistrationInvitationsByAppResponse> {
    return fm.fetchReq<GetRegistrationInvitationsByAppRequest, GetRegistrationInvitationsByAppResponse>(`/v1/get/registration/invitations/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRegistrationInvitationsByAppInviter(req: GetRegistrationInvitationsByAppInviterRequest, initReq?: fm.InitReq): Promise<GetRegistrationInvitationsByAppInviterResponse> {
    return fm.fetchReq<GetRegistrationInvitationsByAppInviterRequest, GetRegistrationInvitationsByAppInviterResponse>(`/v1/get/registration/invitations/by/app/inviter`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRegistrationInvitationByAppInvitee(req: GetRegistrationInvitationByAppInviteeRequest, initReq?: fm.InitReq): Promise<GetRegistrationInvitationByAppInviteeResponse> {
    return fm.fetchReq<GetRegistrationInvitationByAppInviteeRequest, GetRegistrationInvitationByAppInviteeResponse>(`/v1/get/registration/invitations/by/app/invitee`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserInvitationCode(req: CreateUserInvitationCodeRequest, initReq?: fm.InitReq): Promise<CreateUserInvitationCodeResponse> {
    return fm.fetchReq<CreateUserInvitationCodeRequest, CreateUserInvitationCodeResponse>(`/v1/create/user/invitation/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCode(req: GetUserInvitationCodeRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodeResponse> {
    return fm.fetchReq<GetUserInvitationCodeRequest, GetUserInvitationCodeResponse>(`/v1/get/user/invitation/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodeByAppUser(req: GetUserInvitationCodeByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodeByAppUserResponse> {
    return fm.fetchReq<GetUserInvitationCodeByAppUserRequest, GetUserInvitationCodeByAppUserResponse>(`/v1/get/user/invitation/code/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodeByCode(req: GetUserInvitationCodeByCodeRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodeByCodeResponse> {
    return fm.fetchReq<GetUserInvitationCodeByCodeRequest, GetUserInvitationCodeByCodeResponse>(`/v1/get/user/invitation/code/by/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCouponAllocated(req: CreateCouponAllocatedRequest, initReq?: fm.InitReq): Promise<CreateCouponAllocatedResponse> {
    return fm.fetchReq<CreateCouponAllocatedRequest, CreateCouponAllocatedResponse>(`/v1/create/coupon/allocated`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponAllocated(req: GetCouponAllocatedRequest, initReq?: fm.InitReq): Promise<GetCouponAllocatedResponse> {
    return fm.fetchReq<GetCouponAllocatedRequest, GetCouponAllocatedResponse>(`/v1/get/coupon/allocated`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponsAllocatedByApp(req: GetCouponsAllocatedByAppRequest, initReq?: fm.InitReq): Promise<GetCouponsAllocatedByAppResponse> {
    return fm.fetchReq<GetCouponsAllocatedByAppRequest, GetCouponsAllocatedByAppResponse>(`/v1/get/coupons/allocated/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponsAllocatedByAppUser(req: GetCouponsAllocatedByAppUserRequest, initReq?: fm.InitReq): Promise<GetCouponsAllocatedByAppUserResponse> {
    return fm.fetchReq<GetCouponsAllocatedByAppUserRequest, GetCouponsAllocatedByAppUserResponse>(`/v1/get/coupons/allocated/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCouponAllocated(req: UpdateCouponAllocatedRequest, initReq?: fm.InitReq): Promise<UpdateCouponAllocatedResponse> {
    return fm.fetchReq<UpdateCouponAllocatedRequest, UpdateCouponAllocatedResponse>(`/v1/update/coupon/allocated`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponAllocatedDetail(req: GetCouponAllocatedDetailRequest, initReq?: fm.InitReq): Promise<GetCouponAllocatedDetailResponse> {
    return fm.fetchReq<GetCouponAllocatedDetailRequest, GetCouponAllocatedDetailResponse>(`/v1/get/coupon/allocated/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponsAllocatedDetailByApp(req: GetCouponsAllocatedDetailByAppRequest, initReq?: fm.InitReq): Promise<GetCouponsAllocatedDetailByAppResponse> {
    return fm.fetchReq<GetCouponsAllocatedDetailByAppRequest, GetCouponsAllocatedDetailByAppResponse>(`/v1/get/coupons/allocated/detail/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponsAllocatedDetailByAppUser(req: GetCouponsAllocatedDetailByAppUserRequest, initReq?: fm.InitReq): Promise<GetCouponsAllocatedDetailByAppUserResponse> {
    return fm.fetchReq<GetCouponsAllocatedDetailByAppUserRequest, GetCouponsAllocatedDetailByAppUserResponse>(`/v1/get/coupons/allocated/detail/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCouponPool(req: CreateCouponPoolRequest, initReq?: fm.InitReq): Promise<CreateCouponPoolResponse> {
    return fm.fetchReq<CreateCouponPoolRequest, CreateCouponPoolResponse>(`/v1/create/coupon/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCouponPool(req: UpdateCouponPoolRequest, initReq?: fm.InitReq): Promise<UpdateCouponPoolResponse> {
    return fm.fetchReq<UpdateCouponPoolRequest, UpdateCouponPoolResponse>(`/v1/update/coupon/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponPool(req: GetCouponPoolRequest, initReq?: fm.InitReq): Promise<GetCouponPoolResponse> {
    return fm.fetchReq<GetCouponPoolRequest, GetCouponPoolResponse>(`/v1/get/coupon/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponPoolsByApp(req: GetCouponPoolsByAppRequest, initReq?: fm.InitReq): Promise<GetCouponPoolsByAppResponse> {
    return fm.fetchReq<GetCouponPoolsByAppRequest, GetCouponPoolsByAppResponse>(`/v1/get/coupon/pools/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponPoolsByAppReleaser(req: GetCouponPoolsByAppReleaserRequest, initReq?: fm.InitReq): Promise<GetCouponPoolsByAppReleaserResponse> {
    return fm.fetchReq<GetCouponPoolsByAppReleaserRequest, GetCouponPoolsByAppReleaserResponse>(`/v1/get/coupon/pools/by/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateDiscountPool(req: CreateDiscountPoolRequest, initReq?: fm.InitReq): Promise<CreateDiscountPoolResponse> {
    return fm.fetchReq<CreateDiscountPoolRequest, CreateDiscountPoolResponse>(`/v1/create/discount/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateDiscountPool(req: UpdateDiscountPoolRequest, initReq?: fm.InitReq): Promise<UpdateDiscountPoolResponse> {
    return fm.fetchReq<UpdateDiscountPoolRequest, UpdateDiscountPoolResponse>(`/v1/update/discount/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDiscountPool(req: GetDiscountPoolRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolResponse> {
    return fm.fetchReq<GetDiscountPoolRequest, GetDiscountPoolResponse>(`/v1/get/discount/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDiscountPoolsByApp(req: GetDiscountPoolsByAppRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolsByAppResponse> {
    return fm.fetchReq<GetDiscountPoolsByAppRequest, GetDiscountPoolsByAppResponse>(`/v1/get/discount/pools/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDiscountPoolsByAppReleaser(req: GetDiscountPoolsByAppReleaserRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolsByAppReleaserResponse> {
    return fm.fetchReq<GetDiscountPoolsByAppReleaserRequest, GetDiscountPoolsByAppReleaserResponse>(`/v1/get/discount/pools/by/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCouponSetting(req: CreateAppCouponSettingRequest, initReq?: fm.InitReq): Promise<CreateAppCouponSettingResponse> {
    return fm.fetchReq<CreateAppCouponSettingRequest, CreateAppCouponSettingResponse>(`/v1/create/app/coupon/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCouponSetting(req: GetAppCouponSettingRequest, initReq?: fm.InitReq): Promise<GetAppCouponSettingResponse> {
    return fm.fetchReq<GetAppCouponSettingRequest, GetAppCouponSettingResponse>(`/v1/get/app/coupon/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCouponSettingByApp(req: GetAppCouponSettingByAppRequest, initReq?: fm.InitReq): Promise<GetAppCouponSettingByAppResponse> {
    return fm.fetchReq<GetAppCouponSettingByAppRequest, GetAppCouponSettingByAppResponse>(`/v1/get/app/coupon/setting/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppCouponSetting(req: UpdateAppCouponSettingRequest, initReq?: fm.InitReq): Promise<UpdateAppCouponSettingResponse> {
    return fm.fetchReq<UpdateAppCouponSettingRequest, UpdateAppCouponSettingResponse>(`/v1/update/app/coupon/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateDefaultKpiSetting(req: CreateDefaultKpiSettingRequest, initReq?: fm.InitReq): Promise<CreateDefaultKpiSettingResponse> {
    return fm.fetchReq<CreateDefaultKpiSettingRequest, CreateDefaultKpiSettingResponse>(`/v1/create/default/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDefaultKpiSetting(req: GetDefaultKpiSettingRequest, initReq?: fm.InitReq): Promise<GetDefaultKpiSettingResponse> {
    return fm.fetchReq<GetDefaultKpiSettingRequest, GetDefaultKpiSettingResponse>(`/v1/get/default/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDefaultKpiSettingByAppGood(req: GetDefaultKpiSettingByAppGoodRequest, initReq?: fm.InitReq): Promise<GetDefaultKpiSettingByAppGoodResponse> {
    return fm.fetchReq<GetDefaultKpiSettingByAppGoodRequest, GetDefaultKpiSettingByAppGoodResponse>(`/v1/get/default/kpi/setting/by/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateDefaultKpiSetting(req: UpdateDefaultKpiSettingRequest, initReq?: fm.InitReq): Promise<UpdateDefaultKpiSettingResponse> {
    return fm.fetchReq<UpdateDefaultKpiSettingRequest, UpdateDefaultKpiSettingResponse>(`/v1/update/default/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserKpiSetting(req: CreateUserKpiSettingRequest, initReq?: fm.InitReq): Promise<CreateUserKpiSettingResponse> {
    return fm.fetchReq<CreateUserKpiSettingRequest, CreateUserKpiSettingResponse>(`/v1/create/user/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserKpiSetting(req: GetUserKpiSettingRequest, initReq?: fm.InitReq): Promise<GetUserKpiSettingResponse> {
    return fm.fetchReq<GetUserKpiSettingRequest, GetUserKpiSettingResponse>(`/v1/get/user/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserKpiSettingByAppGood(req: GetUserKpiSettingByAppGoodRequest, initReq?: fm.InitReq): Promise<GetUserKpiSettingByAppGoodResponse> {
    return fm.fetchReq<GetUserKpiSettingByAppGoodRequest, GetUserKpiSettingByAppGoodResponse>(`/v1/get/user/kpi/setting/by/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserKpiSetting(req: UpdateUserKpiSettingRequest, initReq?: fm.InitReq): Promise<UpdateUserKpiSettingResponse> {
    return fm.fetchReq<UpdateUserKpiSettingRequest, UpdateUserKpiSettingResponse>(`/v1/update/user/kpi/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserSpecialReduction(req: CreateUserSpecialReductionRequest, initReq?: fm.InitReq): Promise<CreateUserSpecialReductionResponse> {
    return fm.fetchReq<CreateUserSpecialReductionRequest, CreateUserSpecialReductionResponse>(`/v1/create/user/special/reduction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReduction(req: GetUserSpecialReductionRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionResponse> {
    return fm.fetchReq<GetUserSpecialReductionRequest, GetUserSpecialReductionResponse>(`/v1/get/user/special/reduction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReductionsByApp(req: GetUserSpecialReductionsByAppRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionsByAppResponse> {
    return fm.fetchReq<GetUserSpecialReductionsByAppRequest, GetUserSpecialReductionsByAppResponse>(`/v1/get/user/special/reductions/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReductionsByAppReleaser(req: GetUserSpecialReductionsByAppReleaserRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionsByAppReleaserResponse> {
    return fm.fetchReq<GetUserSpecialReductionsByAppReleaserRequest, GetUserSpecialReductionsByAppReleaserResponse>(`/v1/get/user/special/reductions/by/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReductionsByAppUser(req: GetUserSpecialReductionsByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionsByAppUserResponse> {
    return fm.fetchReq<GetUserSpecialReductionsByAppUserRequest, GetUserSpecialReductionsByAppUserResponse>(`/v1/get/user/special/reductions/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserSpecialReduction(req: UpdateUserSpecialReductionRequest, initReq?: fm.InitReq): Promise<UpdateUserSpecialReductionResponse> {
    return fm.fetchReq<UpdateUserSpecialReductionRequest, UpdateUserSpecialReductionResponse>(`/v1/update/user/special/reduction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}