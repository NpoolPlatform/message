/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type NewUserRewardSetting = {
  id?: string
  appID?: string
  registrationCouponID?: string
  kycCouponID?: string
  autoGenerateInvitationCode?: boolean
}

export type CreateNewUserRewardSettingRequest = {
  info?: NewUserRewardSetting
}

export type CreateNewUserRewardSettingResponse = {
  info?: NewUserRewardSetting
}

export type GetNewUserRewardSettingRequest = {
  id?: string
}

export type GetNewUserRewardSettingResponse = {
  info?: NewUserRewardSetting
}

export type GetNewUserRewardSettingByAppRequest = {
  appID?: string
}

export type GetNewUserRewardSettingByAppResponse = {
  info?: NewUserRewardSetting
}

export type UpdateNewUserRewardSettingRequest = {
  info?: NewUserRewardSetting
}

export type UpdateNewUserRewardSettingResponse = {
  info?: NewUserRewardSetting
}

export type AgencySetting = {
  id?: string
  appID?: string
  goodID?: string
  registrationRewardThreshold?: number
  registrationCouponID?: string
  kycRewardThreshold?: number
  kycCouponID?: string
  totalPurchaseRewardPercent?: number
  purchaseRewardChainLevels?: number
  levelPurchaseRewardPercent?: number
}

export type CreateAgencySettingRequest = {
  info?: AgencySetting
}

export type CreateAgencySettingResponse = {
  info?: AgencySetting
}

export type GetAgencySettingRequest = {
  id?: string
}

export type GetAgencySettingResponse = {
  info?: AgencySetting
}

export type GetAgencySettingByAppRequest = {
  appID?: string
}

export type GetAgencySettingByAppResponse = {
  info?: AgencySetting
}

export type UpdateAgencySettingRequest = {
  info?: AgencySetting
}

export type UpdateAgencySettingResponse = {
  info?: AgencySetting
}

export type PurchaseInvitation = {
  id?: string
  appID?: string
  orderID?: string
  invitationCodeID?: string
}

export type CreatePurchaseInvitationRequest = {
  info?: PurchaseInvitation
}

export type CreatePurchaseInvitationResponse = {
  info?: PurchaseInvitation
}

export type UpdatePurchaseInvitationRequest = {
  info?: PurchaseInvitation
}

export type UpdatePurchaseInvitationResponse = {
  info?: PurchaseInvitation
}

export type GetPurchaseInvitationRequest = {
  id?: string
}

export type GetPurchaseInvitationResponse = {
  info?: PurchaseInvitation
}

export type GetPurchaseInvitationsByAppRequest = {
  appID?: string
}

export type GetPurchaseInvitationsByAppResponse = {
  infos?: PurchaseInvitation[]
}

export type GetPurchaseInvitationByAppOrderRequest = {
  appID?: string
  orderID?: string
}

export type GetPurchaseInvitationByAppOrderResponse = {
  info?: PurchaseInvitation
}

export type RegistrationInvitation = {
  id?: string
  appID?: string
  inviterID?: string
  inviteeID?: string
  createAt?: number
}

export type CreateRegistrationInvitationRequest = {
  info?: RegistrationInvitation
}

export type CreateRegistrationInvitationResponse = {
  info?: RegistrationInvitation
}

export type UpdateRegistrationInvitationRequest = {
  info?: RegistrationInvitation
}

export type UpdateRegistrationInvitationResponse = {
  info?: RegistrationInvitation
}

export type GetRegistrationInvitationRequest = {
  id?: string
}

export type GetRegistrationInvitationResponse = {
  info?: RegistrationInvitation
}

export type GetRegistrationInvitationsByAppRequest = {
  appID?: string
}

export type GetRegistrationInvitationsByAppResponse = {
  infos?: RegistrationInvitation[]
}

export type GetRegistrationInvitationsByAppInviterRequest = {
  appID?: string
  inviterID?: string
}

export type GetRegistrationInvitationsByAppInviterResponse = {
  infos?: RegistrationInvitation[]
}

export type GetRegistrationInvitationByAppInviteeRequest = {
  appID?: string
  inviteeID?: string
}

export type GetRegistrationInvitationByAppInviteeResponse = {
  info?: RegistrationInvitation
}

export type UserInvitationCode = {
  id?: string
  userID?: string
  appID?: string
  invitationCode?: string
}

export type CreateUserInvitationCodeRequest = {
  info?: UserInvitationCode
}

export type CreateUserInvitationCodeResponse = {
  info?: UserInvitationCode
}

export type CreateUserInvitationCodeForOtherAppUserRequest = {
  targetAppID?: string
  targetUserID?: string
  info?: UserInvitationCode
}

export type CreateUserInvitationCodeForOtherAppUserResponse = {
  info?: UserInvitationCode
}

export type CreateUserInvitationCodeForAppOtherUserRequest = {
  targetUserID?: string
  info?: UserInvitationCode
}

export type CreateUserInvitationCodeForAppOtherUserResponse = {
  info?: UserInvitationCode
}

export type GetUserInvitationCodeRequest = {
  id?: string
}

export type GetUserInvitationCodeResponse = {
  info?: UserInvitationCode
}

export type GetUserInvitationCodeByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetUserInvitationCodeByAppUserResponse = {
  info?: UserInvitationCode
}

export type GetUserInvitationCodesRequest = {
}

export type GetUserInvitationCodesResponse = {
  infos?: UserInvitationCode[]
}

export type GetUserInvitationCodesByAppRequest = {
  appID?: string
}

export type GetUserInvitationCodesByAppResponse = {
  infos?: UserInvitationCode[]
}

export type GetUserInvitationCodesByOtherAppRequest = {
  targetAppID?: string
}

export type GetUserInvitationCodesByOtherAppResponse = {
  infos?: UserInvitationCode[]
}

export type GetUserInvitationCodeByCodeRequest = {
  code?: string
}

export type GetUserInvitationCodeByCodeResponse = {
  info?: UserInvitationCode
}

export type CouponAllocated = {
  id?: string
  userID?: string
  appID?: string
  type?: string
  couponID?: string
}

export type CreateCouponAllocatedRequest = {
  info?: CouponAllocated
}

export type CreateCouponAllocatedResponse = {
  info?: CouponAllocated
}

export type GetCouponAllocatedRequest = {
  id?: string
}

export type GetCouponAllocatedResponse = {
  info?: CouponAllocated
}

export type GetCouponsAllocatedByAppRequest = {
  appID?: string
}

export type GetCouponsAllocatedByAppResponse = {
  infos?: CouponAllocated[]
}

export type GetCouponsAllocatedByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetCouponsAllocatedByAppUserResponse = {
  infos?: CouponAllocated[]
}

export type UpdateCouponAllocatedRequest = {
  info?: CouponAllocated
}

export type UpdateCouponAllocatedResponse = {
  info?: CouponAllocated
}

export type CouponPool = {
  id?: string
  appID?: string
  releaseByUserID?: string
  denomination?: number
  circulation?: number
  start?: number
  durationDays?: number
  message?: string
  name?: string
}

export type CreateCouponPoolRequest = {
  info?: CouponPool
}

export type CreateCouponPoolResponse = {
  info?: CouponPool
}

export type CreateCouponPoolForOtherAppRequest = {
  targetAppID?: string
  info?: CouponPool
}

export type CreateCouponPoolForOtherAppResponse = {
  info?: CouponPool
}

export type UpdateCouponPoolRequest = {
  info?: CouponPool
}

export type UpdateCouponPoolResponse = {
  info?: CouponPool
}

export type GetCouponPoolRequest = {
  id?: string
}

export type GetCouponPoolResponse = {
  info?: CouponPool
}

export type GetCouponPoolsByAppRequest = {
  appID?: string
}

export type GetCouponPoolsByAppResponse = {
  infos?: CouponPool[]
}

export type GetCouponPoolsByOtherAppRequest = {
  targetAppID?: string
}

export type GetCouponPoolsByOtherAppResponse = {
  infos?: CouponPool[]
}

export type GetCouponPoolsByAppReleaserRequest = {
  appID?: string
  userID?: string
}

export type GetCouponPoolsByAppReleaserResponse = {
  infos?: CouponPool[]
}

export type GetCouponPoolsByOtherAppReleaserRequest = {
  targetAppID?: string
  targetUserID?: string
}

export type GetCouponPoolsByOtherAppReleaserResponse = {
  infos?: CouponPool[]
}

export type NewUserRewardSettingDetail = {
  id?: string
  appID?: string
  autoGenerateInvitationCode?: boolean
  registrationCoupon?: CouponPool
  kycCoupon?: CouponPool
}

export type GetNewUserRewardSettingDetailRequest = {
  id?: string
}

export type GetNewUserRewardSettingDetailResponse = {
  info?: NewUserRewardSettingDetail
}

export type AppCouponSetting = {
  id?: string
  appID?: string
  dominationLimit?: number
  totalLimit?: number
}

export type CreateAppCouponSettingRequest = {
  info?: AppCouponSetting
}

export type CreateAppCouponSettingResponse = {
  info?: AppCouponSetting
}

export type GetAppCouponSettingRequest = {
  id?: string
}

export type GetAppCouponSettingResponse = {
  info?: AppCouponSetting
}

export type GetAppCouponSettingByAppRequest = {
  appID?: string
}

export type GetAppCouponSettingByAppResponse = {
  info?: AppCouponSetting
}

export type UpdateAppCouponSettingRequest = {
  info?: AppCouponSetting
}

export type UpdateAppCouponSettingResponse = {
  info?: AppCouponSetting
}

export type DefaultKpiSetting = {
  id?: string
  appID?: string
  goodID?: string
  amount?: number
  percent?: number
}

export type CreateDefaultKpiSettingRequest = {
  info?: DefaultKpiSetting
}

export type CreateDefaultKpiSettingResponse = {
  info?: DefaultKpiSetting
}

export type GetDefaultKpiSettingRequest = {
  id?: string
}

export type GetDefaultKpiSettingResponse = {
  info?: DefaultKpiSetting
}

export type GetDefaultKpiSettingByAppGoodRequest = {
  appID?: string
  goodID?: string
}

export type GetDefaultKpiSettingByAppGoodResponse = {
  info?: DefaultKpiSetting
}

export type UpdateDefaultKpiSettingRequest = {
  info?: DefaultKpiSetting
}

export type UpdateDefaultKpiSettingResponse = {
  info?: DefaultKpiSetting
}

export type UserKpiSetting = {
  id?: string
  appID?: string
  userID?: string
  goodID?: string
  amount?: number
  percent?: number
}

export type CreateUserKpiSettingRequest = {
  info?: UserKpiSetting
}

export type CreateUserKpiSettingResponse = {
  info?: UserKpiSetting
}

export type GetUserKpiSettingRequest = {
  id?: string
}

export type GetUserKpiSettingResponse = {
  info?: UserKpiSetting
}

export type GetUserKpiSettingByAppGoodRequest = {
  appID?: string
  goodID?: string
}

export type GetUserKpiSettingByAppGoodResponse = {
  info?: UserKpiSetting
}

export type UpdateUserKpiSettingRequest = {
  info?: UserKpiSetting
}

export type UpdateUserKpiSettingResponse = {
  info?: UserKpiSetting
}

export type AgencySettingDetail = {
  id?: string
  appID?: string
  goodID?: string
  registrationRewardThreshold?: number
  registrationCoupon?: CouponPool
  kycRewardThreshold?: number
  kycCoupon?: CouponPool
  totalPurchaseRewardPercent?: number
  purchaseRewardChainLevels?: number
  levelPurchaseRewardPercent?: number
}

export type GetAgencySettingDetailRequest = {
  id?: string
}

export type GetAgencySettingDetailResponse = {
  info?: AgencySettingDetail
}

export type GetAgencySettingDetailByAppRequest = {
  appID?: string
}

export type GetAgencySettingDetailByAppResponse = {
  info?: AgencySettingDetail
}

export type CouponAllocatedDetail = {
  id?: string
  userID?: string
  appID?: string
  type?: string
  coupon?: CouponPool
  discount?: DiscountPool
}

export type GetCouponAllocatedDetailRequest = {
  id?: string
}

export type GetCouponAllocatedDetailResponse = {
  info?: CouponAllocatedDetail
}

export type GetCouponsAllocatedDetailByAppRequest = {
  appID?: string
}

export type GetCouponsAllocatedDetailByAppResponse = {
  infos?: CouponAllocatedDetail[]
}

export type GetCouponsAllocatedDetailByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetCouponsAllocatedDetailByAppUserResponse = {
  infos?: CouponAllocatedDetail[]
}

export type DiscountPool = {
  id?: string
  appID?: string
  releaseByUserID?: string
  discount?: number
  start?: number
  durationDays?: number
  message?: string
  name?: string
}

export type CreateDiscountPoolRequest = {
  info?: DiscountPool
}

export type CreateDiscountPoolResponse = {
  info?: DiscountPool
}

export type CreateDiscountPoolForOtherAppRequest = {
  targetAppID?: string
  info?: DiscountPool
}

export type CreateDiscountPoolForOtherAppResponse = {
  info?: DiscountPool
}

export type UpdateDiscountPoolRequest = {
  info?: DiscountPool
}

export type UpdateDiscountPoolResponse = {
  info?: DiscountPool
}

export type GetDiscountPoolRequest = {
  id?: string
}

export type GetDiscountPoolResponse = {
  info?: DiscountPool
}

export type GetDiscountPoolsByAppRequest = {
  appID?: string
}

export type GetDiscountPoolsByAppResponse = {
  infos?: DiscountPool[]
}

export type GetDiscountPoolsByOtherAppRequest = {
  targetAppID?: string
}

export type GetDiscountPoolsByOtherAppResponse = {
  infos?: DiscountPool[]
}

export type GetDiscountPoolsByAppReleaserRequest = {
  appID?: string
  userID?: string
}

export type GetDiscountPoolsByAppReleaserResponse = {
  infos?: DiscountPool[]
}

export type GetDiscountPoolsByOtherAppReleaserRequest = {
  targetAppID?: string
  targetUserID?: string
}

export type GetDiscountPoolsByOtherAppReleaserResponse = {
  infos?: DiscountPool[]
}

export type UserSpecialReduction = {
  id?: string
  appID?: string
  userID?: string
  amount?: number
  releaseByUserID?: string
  start?: number
  durationDays?: number
  message?: string
}

export type CreateUserSpecialReductionRequest = {
  info?: UserSpecialReduction
}

export type CreateUserSpecialReductionResponse = {
  info?: UserSpecialReduction
}

export type CreateUserSpecialReductionForOtherAppRequest = {
  targetAppID?: string
  info?: UserSpecialReduction
}

export type CreateUserSpecialReductionForOtherAppResponse = {
  info?: UserSpecialReduction
}

export type GetUserSpecialReductionRequest = {
  id?: string
}

export type GetUserSpecialReductionResponse = {
  info?: UserSpecialReduction
}

export type GetUserSpecialReductionsByAppRequest = {
  appID?: string
}

export type GetUserSpecialReductionsByAppResponse = {
  infos?: UserSpecialReduction[]
}

export type GetUserSpecialReductionsByOtherAppRequest = {
  targetAppID?: string
}

export type GetUserSpecialReductionsByOtherAppResponse = {
  infos?: UserSpecialReduction[]
}

export type GetUserSpecialReductionsByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetUserSpecialReductionsByAppUserResponse = {
  infos?: UserSpecialReduction[]
}

export type UpdateUserSpecialReductionRequest = {
  info?: UserSpecialReduction
}

export type UpdateUserSpecialReductionResponse = {
  info?: UserSpecialReduction
}

export type GetUserSpecialReductionsByAppReleaserRequest = {
  appID?: string
  userID?: string
}

export type GetUserSpecialReductionsByAppReleaserResponse = {
  infos?: UserSpecialReduction[]
}

export type Activity = {
  id?: string
  appID?: string
  createdBy?: string
  name?: string
  start?: number
  end?: number
  systemActivity?: boolean
}

export type CreateActivityRequest = {
  info?: Activity
}

export type CreateActivityResponse = {
  info?: Activity
}

export type UpdateActivityRequest = {
  info?: Activity
}

export type UpdateActivityResponse = {
  info?: Activity
}

export type CreateActivityForOtherAppRequest = {
  targetAppID?: string
  info?: Activity
}

export type CreateActivityForOtherAppResponse = {
  info?: Activity
}

export type GetActivityRequest = {
  id?: string
}

export type GetActivityResponse = {
  info?: Activity
}

export type GetActivityByAppNameRequest = {
  appID?: string
  name?: string
}

export type GetActivityByAppNameResponse = {
  info?: Activity
}

export type GetActivitiesByAppRequest = {
  appID?: string
}

export type GetActivitiesByAppResponse = {
  infos?: Activity[]
}

export type GetActivitiesByOtherAppRequest = {
  targetAppID?: string
}

export type GetActivitiesByOtherAppResponse = {
  infos?: Activity[]
}

export type EventCoupon = {
  id?: string
  appID?: string
  activityID?: string
  event?: string
  couponID?: string
}

export type CreateEventCouponRequest = {
  info?: EventCoupon
}

export type CreateEventCouponResponse = {
  info?: EventCoupon
}

export type CreateEventCouponForOtherAppRequest = {
  targetAppID?: string
  info?: EventCoupon
}

export type CreateEventCouponForOtherAppResponse = {
  info?: EventCoupon
}

export type UpdateEventCouponRequest = {
  info?: EventCoupon
}

export type UpdateEventCouponResponse = {
  info?: EventCoupon
}

export type GetEventCouponRequest = {
  id?: string
}

export type GetEventCouponResponse = {
  info?: EventCoupon
}

export type GetEventCouponsByAppActivityEventRequest = {
  appID?: string
  activityID?: string
  event?: string
}

export type GetEventCouponsByAppActivityEventResponse = {
  infos?: EventCoupon[]
}

export type GetEventCouponsByAppRequest = {
  appID?: string
}

export type GetEventCouponsByAppResponse = {
  infos?: EventCoupon[]
}

export type GetEventCouponsByOtherAppRequest = {
  targetAppID?: string
}

export type GetEventCouponsByOtherAppResponse = {
  infos?: EventCoupon[]
}

export class CloudHashingInspire {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static CreateUserInvitationCodeForOtherAppUser(req: CreateUserInvitationCodeForOtherAppUserRequest, initReq?: fm.InitReq): Promise<CreateUserInvitationCodeForOtherAppUserResponse> {
    return fm.fetchReq<CreateUserInvitationCodeForOtherAppUserRequest, CreateUserInvitationCodeForOtherAppUserResponse>(`/v1/create/user/invitation/code/for/other/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserInvitationCodeForAppOtherUser(req: CreateUserInvitationCodeForAppOtherUserRequest, initReq?: fm.InitReq): Promise<CreateUserInvitationCodeForAppOtherUserResponse> {
    return fm.fetchReq<CreateUserInvitationCodeForAppOtherUserRequest, CreateUserInvitationCodeForAppOtherUserResponse>(`/v1/create/user/invitation/code/for/app/other/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCode(req: GetUserInvitationCodeRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodeResponse> {
    return fm.fetchReq<GetUserInvitationCodeRequest, GetUserInvitationCodeResponse>(`/v1/get/user/invitation/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodeByAppUser(req: GetUserInvitationCodeByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodeByAppUserResponse> {
    return fm.fetchReq<GetUserInvitationCodeByAppUserRequest, GetUserInvitationCodeByAppUserResponse>(`/v1/get/user/invitation/code/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodes(req: GetUserInvitationCodesRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodesResponse> {
    return fm.fetchReq<GetUserInvitationCodesRequest, GetUserInvitationCodesResponse>(`/v1/get/user/invitation/codes`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodesByApp(req: GetUserInvitationCodesByAppRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodesByAppResponse> {
    return fm.fetchReq<GetUserInvitationCodesByAppRequest, GetUserInvitationCodesByAppResponse>(`/v1/get/user/invitation/codes/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserInvitationCodesByOtherApp(req: GetUserInvitationCodesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetUserInvitationCodesByOtherAppResponse> {
    return fm.fetchReq<GetUserInvitationCodesByOtherAppRequest, GetUserInvitationCodesByOtherAppResponse>(`/v1/get/user/invitation/codes/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static CreateCouponPoolForOtherApp(req: CreateCouponPoolForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateCouponPoolForOtherAppResponse> {
    return fm.fetchReq<CreateCouponPoolForOtherAppRequest, CreateCouponPoolForOtherAppResponse>(`/v1/create/coupon/pool/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetCouponPoolsByOtherApp(req: GetCouponPoolsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetCouponPoolsByOtherAppResponse> {
    return fm.fetchReq<GetCouponPoolsByOtherAppRequest, GetCouponPoolsByOtherAppResponse>(`/v1/get/coupon/pools/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponPoolsByAppReleaser(req: GetCouponPoolsByAppReleaserRequest, initReq?: fm.InitReq): Promise<GetCouponPoolsByAppReleaserResponse> {
    return fm.fetchReq<GetCouponPoolsByAppReleaserRequest, GetCouponPoolsByAppReleaserResponse>(`/v1/get/coupon/pools/by/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCouponPoolsByOtherAppReleaser(req: GetCouponPoolsByOtherAppReleaserRequest, initReq?: fm.InitReq): Promise<GetCouponPoolsByOtherAppReleaserResponse> {
    return fm.fetchReq<GetCouponPoolsByOtherAppReleaserRequest, GetCouponPoolsByOtherAppReleaserResponse>(`/v1/get/coupon/pools/by/other/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateDiscountPool(req: CreateDiscountPoolRequest, initReq?: fm.InitReq): Promise<CreateDiscountPoolResponse> {
    return fm.fetchReq<CreateDiscountPoolRequest, CreateDiscountPoolResponse>(`/v1/create/discount/pool`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateDiscountPoolForOtherApp(req: CreateDiscountPoolForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateDiscountPoolForOtherAppResponse> {
    return fm.fetchReq<CreateDiscountPoolForOtherAppRequest, CreateDiscountPoolForOtherAppResponse>(`/v1/create/discount/pool/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetDiscountPoolsByOtherApp(req: GetDiscountPoolsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolsByOtherAppResponse> {
    return fm.fetchReq<GetDiscountPoolsByOtherAppRequest, GetDiscountPoolsByOtherAppResponse>(`/v1/get/discount/pools/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDiscountPoolsByAppReleaser(req: GetDiscountPoolsByAppReleaserRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolsByAppReleaserResponse> {
    return fm.fetchReq<GetDiscountPoolsByAppReleaserRequest, GetDiscountPoolsByAppReleaserResponse>(`/v1/get/discount/pools/by/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDiscountPoolsByOtherAppReleaser(req: GetDiscountPoolsByOtherAppReleaserRequest, initReq?: fm.InitReq): Promise<GetDiscountPoolsByOtherAppReleaserResponse> {
    return fm.fetchReq<GetDiscountPoolsByOtherAppReleaserRequest, GetDiscountPoolsByOtherAppReleaserResponse>(`/v1/get/discount/pools/by/other/app/releaser`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static CreateUserSpecialReductionForOtherApp(req: CreateUserSpecialReductionForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateUserSpecialReductionForOtherAppResponse> {
    return fm.fetchReq<CreateUserSpecialReductionForOtherAppRequest, CreateUserSpecialReductionForOtherAppResponse>(`/v1/create/user/special/reduction/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReduction(req: GetUserSpecialReductionRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionResponse> {
    return fm.fetchReq<GetUserSpecialReductionRequest, GetUserSpecialReductionResponse>(`/v1/get/user/special/reduction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReductionsByApp(req: GetUserSpecialReductionsByAppRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionsByAppResponse> {
    return fm.fetchReq<GetUserSpecialReductionsByAppRequest, GetUserSpecialReductionsByAppResponse>(`/v1/get/user/special/reductions/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserSpecialReductionsByOtherApp(req: GetUserSpecialReductionsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetUserSpecialReductionsByOtherAppResponse> {
    return fm.fetchReq<GetUserSpecialReductionsByOtherAppRequest, GetUserSpecialReductionsByOtherAppResponse>(`/v1/get/user/special/reductions/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static CreateActivity(req: CreateActivityRequest, initReq?: fm.InitReq): Promise<CreateActivityResponse> {
    return fm.fetchReq<CreateActivityRequest, CreateActivityResponse>(`/v1/create/activity`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateActivityForOtherApp(req: CreateActivityForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateActivityForOtherAppResponse> {
    return fm.fetchReq<CreateActivityForOtherAppRequest, CreateActivityForOtherAppResponse>(`/v1/create/activity/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateActivity(req: UpdateActivityRequest, initReq?: fm.InitReq): Promise<UpdateActivityResponse> {
    return fm.fetchReq<UpdateActivityRequest, UpdateActivityResponse>(`/v1/update/activity`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetActivity(req: GetActivityRequest, initReq?: fm.InitReq): Promise<GetActivityResponse> {
    return fm.fetchReq<GetActivityRequest, GetActivityResponse>(`/v1/get/activity`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetActivityByAppName(req: GetActivityByAppNameRequest, initReq?: fm.InitReq): Promise<GetActivityByAppNameResponse> {
    return fm.fetchReq<GetActivityByAppNameRequest, GetActivityByAppNameResponse>(`/v1/get/activity/by/app/name`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetActivitiesByApp(req: GetActivitiesByAppRequest, initReq?: fm.InitReq): Promise<GetActivitiesByAppResponse> {
    return fm.fetchReq<GetActivitiesByAppRequest, GetActivitiesByAppResponse>(`/v1/get/activities/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetActivitiesByOtherApp(req: GetActivitiesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetActivitiesByOtherAppResponse> {
    return fm.fetchReq<GetActivitiesByOtherAppRequest, GetActivitiesByOtherAppResponse>(`/v1/get/activities/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateEventCoupon(req: CreateEventCouponRequest, initReq?: fm.InitReq): Promise<CreateEventCouponResponse> {
    return fm.fetchReq<CreateEventCouponRequest, CreateEventCouponResponse>(`/v1/create/event/coupon`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateEventCouponForOtherApp(req: CreateEventCouponForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateEventCouponForOtherAppResponse> {
    return fm.fetchReq<CreateEventCouponForOtherAppRequest, CreateEventCouponForOtherAppResponse>(`/v1/create/event/coupon/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateEventCoupon(req: UpdateEventCouponRequest, initReq?: fm.InitReq): Promise<UpdateEventCouponResponse> {
    return fm.fetchReq<UpdateEventCouponRequest, UpdateEventCouponResponse>(`/v1/update/event/coupon`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetEventCoupon(req: GetEventCouponRequest, initReq?: fm.InitReq): Promise<GetEventCouponResponse> {
    return fm.fetchReq<GetEventCouponRequest, GetEventCouponResponse>(`/v1/get/event/coupon`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetEventCouponsByAppActivityEvent(req: GetEventCouponsByAppActivityEventRequest, initReq?: fm.InitReq): Promise<GetEventCouponsByAppActivityEventResponse> {
    return fm.fetchReq<GetEventCouponsByAppActivityEventRequest, GetEventCouponsByAppActivityEventResponse>(`/v1/get/event/coupons/by/app/activity/event`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetEventCouponsByApp(req: GetEventCouponsByAppRequest, initReq?: fm.InitReq): Promise<GetEventCouponsByAppResponse> {
    return fm.fetchReq<GetEventCouponsByAppRequest, GetEventCouponsByAppResponse>(`/v1/get/event/coupons/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetEventCouponsByOtherApp(req: GetEventCouponsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetEventCouponsByOtherAppResponse> {
    return fm.fetchReq<GetEventCouponsByOtherAppRequest, GetEventCouponsByOtherAppResponse>(`/v1/get/event/coupons/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}