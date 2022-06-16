# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-inspire/cloud-hashing-inspire.proto](#npool_cloud-hashing-inspire_cloud-hashing-inspire-proto)
    - [Activity](#cloud-hashing-inspire-v1-Activity)
    - [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting)
    - [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting)
    - [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting)
    - [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting)
    - [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting)
    - [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated)
    - [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail)
    - [CouponPool](#cloud-hashing-inspire-v1-CouponPool)
    - [CreateActivityForOtherAppRequest](#cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest)
    - [CreateActivityForOtherAppResponse](#cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse)
    - [CreateActivityRequest](#cloud-hashing-inspire-v1-CreateActivityRequest)
    - [CreateActivityResponse](#cloud-hashing-inspire-v1-CreateActivityResponse)
    - [CreateAppCommissionSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest)
    - [CreateAppCommissionSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse)
    - [CreateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest)
    - [CreateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse)
    - [CreateAppCouponSettingRequest](#cloud-hashing-inspire-v1-CreateAppCouponSettingRequest)
    - [CreateAppCouponSettingResponse](#cloud-hashing-inspire-v1-CreateAppCouponSettingResponse)
    - [CreateAppInvitationSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest)
    - [CreateAppInvitationSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse)
    - [CreateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest)
    - [CreateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse)
    - [CreateAppPurchaseAmountSettingForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest)
    - [CreateAppPurchaseAmountSettingForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse)
    - [CreateAppPurchaseAmountSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest)
    - [CreateAppPurchaseAmountSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse)
    - [CreateAppPurchaseAmountSettingForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest)
    - [CreateAppPurchaseAmountSettingForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse)
    - [CreateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest)
    - [CreateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse)
    - [CreateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest)
    - [CreateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse)
    - [CreateCouponAllocatedForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest)
    - [CreateCouponAllocatedForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse)
    - [CreateCouponAllocatedForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest)
    - [CreateCouponAllocatedForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse)
    - [CreateCouponAllocatedRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedRequest)
    - [CreateCouponAllocatedResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedResponse)
    - [CreateCouponPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest)
    - [CreateCouponPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse)
    - [CreateCouponPoolRequest](#cloud-hashing-inspire-v1-CreateCouponPoolRequest)
    - [CreateCouponPoolResponse](#cloud-hashing-inspire-v1-CreateCouponPoolResponse)
    - [CreateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest)
    - [CreateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse)
    - [CreateDiscountPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest)
    - [CreateDiscountPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse)
    - [CreateDiscountPoolRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolRequest)
    - [CreateDiscountPoolResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolResponse)
    - [CreateEventCouponForOtherAppRequest](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest)
    - [CreateEventCouponForOtherAppResponse](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse)
    - [CreateEventCouponRequest](#cloud-hashing-inspire-v1-CreateEventCouponRequest)
    - [CreateEventCouponResponse](#cloud-hashing-inspire-v1-CreateEventCouponResponse)
    - [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest)
    - [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse)
    - [CreateUserInvitationCodeForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest)
    - [CreateUserInvitationCodeForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse)
    - [CreateUserInvitationCodeForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest)
    - [CreateUserInvitationCodeForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse)
    - [CreateUserInvitationCodeRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest)
    - [CreateUserInvitationCodeResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse)
    - [CreateUserKpiSettingRequest](#cloud-hashing-inspire-v1-CreateUserKpiSettingRequest)
    - [CreateUserKpiSettingResponse](#cloud-hashing-inspire-v1-CreateUserKpiSettingResponse)
    - [CreateUserSpecialReductionForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest)
    - [CreateUserSpecialReductionForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse)
    - [CreateUserSpecialReductionForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest)
    - [CreateUserSpecialReductionForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse)
    - [CreateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest)
    - [CreateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse)
    - [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting)
    - [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool)
    - [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon)
    - [GetActivitiesByAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByAppRequest)
    - [GetActivitiesByAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByAppResponse)
    - [GetActivitiesByOtherAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest)
    - [GetActivitiesByOtherAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse)
    - [GetActivityByAppNameRequest](#cloud-hashing-inspire-v1-GetActivityByAppNameRequest)
    - [GetActivityByAppNameResponse](#cloud-hashing-inspire-v1-GetActivityByAppNameResponse)
    - [GetActivityRequest](#cloud-hashing-inspire-v1-GetActivityRequest)
    - [GetActivityResponse](#cloud-hashing-inspire-v1-GetActivityResponse)
    - [GetAppCommissionSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest)
    - [GetAppCommissionSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse)
    - [GetAppCommissionSettingByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest)
    - [GetAppCommissionSettingByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse)
    - [GetAppCommissionSettingRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingRequest)
    - [GetAppCommissionSettingResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingResponse)
    - [GetAppCouponSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest)
    - [GetAppCouponSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse)
    - [GetAppCouponSettingRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingRequest)
    - [GetAppCouponSettingResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingResponse)
    - [GetAppInvitationSettingRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingRequest)
    - [GetAppInvitationSettingResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingResponse)
    - [GetAppInvitationSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest)
    - [GetAppInvitationSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse)
    - [GetAppInvitationSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest)
    - [GetAppInvitationSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse)
    - [GetAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest)
    - [GetAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse)
    - [GetAppPurchaseAmountSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest)
    - [GetAppPurchaseAmountSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse)
    - [GetAppPurchaseAmountSettingsByAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest)
    - [GetAppPurchaseAmountSettingsByAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse)
    - [GetAppPurchaseAmountSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest)
    - [GetAppPurchaseAmountSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse)
    - [GetAppPurchaseAmountSettingsByOtherAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest)
    - [GetAppPurchaseAmountSettingsByOtherAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse)
    - [GetCommissionCoinSettingsRequest](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest)
    - [GetCommissionCoinSettingsResponse](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse)
    - [GetCouponAllocatedDetailRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest)
    - [GetCouponAllocatedDetailResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse)
    - [GetCouponAllocatedRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedRequest)
    - [GetCouponAllocatedResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedResponse)
    - [GetCouponPoolRequest](#cloud-hashing-inspire-v1-GetCouponPoolRequest)
    - [GetCouponPoolResponse](#cloud-hashing-inspire-v1-GetCouponPoolResponse)
    - [GetCouponPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest)
    - [GetCouponPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse)
    - [GetCouponPoolsByAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest)
    - [GetCouponPoolsByAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse)
    - [GetCouponPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest)
    - [GetCouponPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse)
    - [GetCouponPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest)
    - [GetCouponPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse)
    - [GetCouponsAllocatedByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest)
    - [GetCouponsAllocatedByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse)
    - [GetCouponsAllocatedByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest)
    - [GetCouponsAllocatedByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse)
    - [GetCouponsAllocatedByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest)
    - [GetCouponsAllocatedByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse)
    - [GetCouponsAllocatedDetailByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest)
    - [GetCouponsAllocatedDetailByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse)
    - [GetCouponsAllocatedDetailByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest)
    - [GetCouponsAllocatedDetailByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse)
    - [GetDefaultKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest)
    - [GetDefaultKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse)
    - [GetDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest)
    - [GetDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse)
    - [GetDiscountPoolRequest](#cloud-hashing-inspire-v1-GetDiscountPoolRequest)
    - [GetDiscountPoolResponse](#cloud-hashing-inspire-v1-GetDiscountPoolResponse)
    - [GetDiscountPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest)
    - [GetDiscountPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse)
    - [GetDiscountPoolsByAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest)
    - [GetDiscountPoolsByAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse)
    - [GetDiscountPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest)
    - [GetDiscountPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse)
    - [GetDiscountPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest)
    - [GetDiscountPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse)
    - [GetEventCouponRequest](#cloud-hashing-inspire-v1-GetEventCouponRequest)
    - [GetEventCouponResponse](#cloud-hashing-inspire-v1-GetEventCouponResponse)
    - [GetEventCouponsByAppActivityEventRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest)
    - [GetEventCouponsByAppActivityEventResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse)
    - [GetEventCouponsByAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppRequest)
    - [GetEventCouponsByAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppResponse)
    - [GetEventCouponsByOtherAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest)
    - [GetEventCouponsByOtherAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse)
    - [GetRegistrationInvitationByAppInviteeRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest)
    - [GetRegistrationInvitationByAppInviteeResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse)
    - [GetRegistrationInvitationRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationRequest)
    - [GetRegistrationInvitationResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationResponse)
    - [GetRegistrationInvitationsByAppInviterRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest)
    - [GetRegistrationInvitationsByAppInviterResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse)
    - [GetRegistrationInvitationsByAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest)
    - [GetRegistrationInvitationsByAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse)
    - [GetRegistrationInvitationsByOtherAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest)
    - [GetRegistrationInvitationsByOtherAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse)
    - [GetUserInvitationCodeByAppUserRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest)
    - [GetUserInvitationCodeByAppUserResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse)
    - [GetUserInvitationCodeByCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest)
    - [GetUserInvitationCodeByCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse)
    - [GetUserInvitationCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeRequest)
    - [GetUserInvitationCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeResponse)
    - [GetUserInvitationCodesByAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest)
    - [GetUserInvitationCodesByAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse)
    - [GetUserInvitationCodesByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest)
    - [GetUserInvitationCodesByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse)
    - [GetUserInvitationCodesRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesRequest)
    - [GetUserInvitationCodesResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesResponse)
    - [GetUserKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest)
    - [GetUserKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse)
    - [GetUserKpiSettingRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingRequest)
    - [GetUserKpiSettingResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingResponse)
    - [GetUserSpecialReductionRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionRequest)
    - [GetUserSpecialReductionResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionResponse)
    - [GetUserSpecialReductionsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest)
    - [GetUserSpecialReductionsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse)
    - [GetUserSpecialReductionsByAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest)
    - [GetUserSpecialReductionsByAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse)
    - [GetUserSpecialReductionsByAppUserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest)
    - [GetUserSpecialReductionsByAppUserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse)
    - [GetUserSpecialReductionsByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest)
    - [GetUserSpecialReductionsByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse)
    - [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation)
    - [UpdateActivityRequest](#cloud-hashing-inspire-v1-UpdateActivityRequest)
    - [UpdateActivityResponse](#cloud-hashing-inspire-v1-UpdateActivityResponse)
    - [UpdateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest)
    - [UpdateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse)
    - [UpdateAppCouponSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest)
    - [UpdateAppCouponSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse)
    - [UpdateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest)
    - [UpdateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse)
    - [UpdateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest)
    - [UpdateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse)
    - [UpdateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest)
    - [UpdateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse)
    - [UpdateCouponAllocatedRequest](#cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest)
    - [UpdateCouponAllocatedResponse](#cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse)
    - [UpdateCouponPoolRequest](#cloud-hashing-inspire-v1-UpdateCouponPoolRequest)
    - [UpdateCouponPoolResponse](#cloud-hashing-inspire-v1-UpdateCouponPoolResponse)
    - [UpdateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest)
    - [UpdateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse)
    - [UpdateDiscountPoolRequest](#cloud-hashing-inspire-v1-UpdateDiscountPoolRequest)
    - [UpdateDiscountPoolResponse](#cloud-hashing-inspire-v1-UpdateDiscountPoolResponse)
    - [UpdateEventCouponRequest](#cloud-hashing-inspire-v1-UpdateEventCouponRequest)
    - [UpdateEventCouponResponse](#cloud-hashing-inspire-v1-UpdateEventCouponResponse)
    - [UpdateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest)
    - [UpdateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse)
    - [UpdateUserKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest)
    - [UpdateUserKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse)
    - [UpdateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest)
    - [UpdateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse)
    - [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode)
    - [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting)
    - [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction)
  
    - [CloudHashingInspire](#cloud-hashing-inspire-v1-CloudHashingInspire)
  
- [npool/cloud-hashing-inspire/cloud-hashing-inspire.proto](#npool_cloud-hashing-inspire_cloud-hashing-inspire-proto)
    - [Activity](#cloud-hashing-inspire-v1-Activity)
    - [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting)
    - [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting)
    - [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting)
    - [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting)
    - [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting)
    - [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated)
    - [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail)
    - [CouponPool](#cloud-hashing-inspire-v1-CouponPool)
    - [CreateActivityForOtherAppRequest](#cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest)
    - [CreateActivityForOtherAppResponse](#cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse)
    - [CreateActivityRequest](#cloud-hashing-inspire-v1-CreateActivityRequest)
    - [CreateActivityResponse](#cloud-hashing-inspire-v1-CreateActivityResponse)
    - [CreateAppCommissionSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest)
    - [CreateAppCommissionSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse)
    - [CreateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest)
    - [CreateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse)
    - [CreateAppCouponSettingRequest](#cloud-hashing-inspire-v1-CreateAppCouponSettingRequest)
    - [CreateAppCouponSettingResponse](#cloud-hashing-inspire-v1-CreateAppCouponSettingResponse)
    - [CreateAppInvitationSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest)
    - [CreateAppInvitationSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse)
    - [CreateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest)
    - [CreateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse)
    - [CreateAppPurchaseAmountSettingForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest)
    - [CreateAppPurchaseAmountSettingForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse)
    - [CreateAppPurchaseAmountSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest)
    - [CreateAppPurchaseAmountSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse)
    - [CreateAppPurchaseAmountSettingForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest)
    - [CreateAppPurchaseAmountSettingForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse)
    - [CreateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest)
    - [CreateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse)
    - [CreateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest)
    - [CreateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse)
    - [CreateCouponAllocatedForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest)
    - [CreateCouponAllocatedForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse)
    - [CreateCouponAllocatedForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest)
    - [CreateCouponAllocatedForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse)
    - [CreateCouponAllocatedRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedRequest)
    - [CreateCouponAllocatedResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedResponse)
    - [CreateCouponPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest)
    - [CreateCouponPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse)
    - [CreateCouponPoolRequest](#cloud-hashing-inspire-v1-CreateCouponPoolRequest)
    - [CreateCouponPoolResponse](#cloud-hashing-inspire-v1-CreateCouponPoolResponse)
    - [CreateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest)
    - [CreateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse)
    - [CreateDiscountPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest)
    - [CreateDiscountPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse)
    - [CreateDiscountPoolRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolRequest)
    - [CreateDiscountPoolResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolResponse)
    - [CreateEventCouponForOtherAppRequest](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest)
    - [CreateEventCouponForOtherAppResponse](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse)
    - [CreateEventCouponRequest](#cloud-hashing-inspire-v1-CreateEventCouponRequest)
    - [CreateEventCouponResponse](#cloud-hashing-inspire-v1-CreateEventCouponResponse)
    - [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest)
    - [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse)
    - [CreateUserInvitationCodeForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest)
    - [CreateUserInvitationCodeForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse)
    - [CreateUserInvitationCodeForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest)
    - [CreateUserInvitationCodeForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse)
    - [CreateUserInvitationCodeRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest)
    - [CreateUserInvitationCodeResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse)
    - [CreateUserKpiSettingRequest](#cloud-hashing-inspire-v1-CreateUserKpiSettingRequest)
    - [CreateUserKpiSettingResponse](#cloud-hashing-inspire-v1-CreateUserKpiSettingResponse)
    - [CreateUserSpecialReductionForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest)
    - [CreateUserSpecialReductionForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse)
    - [CreateUserSpecialReductionForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest)
    - [CreateUserSpecialReductionForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse)
    - [CreateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest)
    - [CreateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse)
    - [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting)
    - [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool)
    - [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon)
    - [GetActivitiesByAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByAppRequest)
    - [GetActivitiesByAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByAppResponse)
    - [GetActivitiesByOtherAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest)
    - [GetActivitiesByOtherAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse)
    - [GetActivityByAppNameRequest](#cloud-hashing-inspire-v1-GetActivityByAppNameRequest)
    - [GetActivityByAppNameResponse](#cloud-hashing-inspire-v1-GetActivityByAppNameResponse)
    - [GetActivityRequest](#cloud-hashing-inspire-v1-GetActivityRequest)
    - [GetActivityResponse](#cloud-hashing-inspire-v1-GetActivityResponse)
    - [GetAppCommissionSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest)
    - [GetAppCommissionSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse)
    - [GetAppCommissionSettingByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest)
    - [GetAppCommissionSettingByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse)
    - [GetAppCommissionSettingRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingRequest)
    - [GetAppCommissionSettingResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingResponse)
    - [GetAppCouponSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest)
    - [GetAppCouponSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse)
    - [GetAppCouponSettingRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingRequest)
    - [GetAppCouponSettingResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingResponse)
    - [GetAppInvitationSettingRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingRequest)
    - [GetAppInvitationSettingResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingResponse)
    - [GetAppInvitationSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest)
    - [GetAppInvitationSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse)
    - [GetAppInvitationSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest)
    - [GetAppInvitationSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse)
    - [GetAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest)
    - [GetAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse)
    - [GetAppPurchaseAmountSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest)
    - [GetAppPurchaseAmountSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse)
    - [GetAppPurchaseAmountSettingsByAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest)
    - [GetAppPurchaseAmountSettingsByAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse)
    - [GetAppPurchaseAmountSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest)
    - [GetAppPurchaseAmountSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse)
    - [GetAppPurchaseAmountSettingsByOtherAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest)
    - [GetAppPurchaseAmountSettingsByOtherAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse)
    - [GetCommissionCoinSettingsRequest](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest)
    - [GetCommissionCoinSettingsResponse](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse)
    - [GetCouponAllocatedDetailRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest)
    - [GetCouponAllocatedDetailResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse)
    - [GetCouponAllocatedRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedRequest)
    - [GetCouponAllocatedResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedResponse)
    - [GetCouponPoolRequest](#cloud-hashing-inspire-v1-GetCouponPoolRequest)
    - [GetCouponPoolResponse](#cloud-hashing-inspire-v1-GetCouponPoolResponse)
    - [GetCouponPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest)
    - [GetCouponPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse)
    - [GetCouponPoolsByAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest)
    - [GetCouponPoolsByAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse)
    - [GetCouponPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest)
    - [GetCouponPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse)
    - [GetCouponPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest)
    - [GetCouponPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse)
    - [GetCouponsAllocatedByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest)
    - [GetCouponsAllocatedByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse)
    - [GetCouponsAllocatedByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest)
    - [GetCouponsAllocatedByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse)
    - [GetCouponsAllocatedByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest)
    - [GetCouponsAllocatedByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse)
    - [GetCouponsAllocatedDetailByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest)
    - [GetCouponsAllocatedDetailByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse)
    - [GetCouponsAllocatedDetailByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest)
    - [GetCouponsAllocatedDetailByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse)
    - [GetDefaultKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest)
    - [GetDefaultKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse)
    - [GetDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest)
    - [GetDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse)
    - [GetDiscountPoolRequest](#cloud-hashing-inspire-v1-GetDiscountPoolRequest)
    - [GetDiscountPoolResponse](#cloud-hashing-inspire-v1-GetDiscountPoolResponse)
    - [GetDiscountPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest)
    - [GetDiscountPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse)
    - [GetDiscountPoolsByAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest)
    - [GetDiscountPoolsByAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse)
    - [GetDiscountPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest)
    - [GetDiscountPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse)
    - [GetDiscountPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest)
    - [GetDiscountPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse)
    - [GetEventCouponRequest](#cloud-hashing-inspire-v1-GetEventCouponRequest)
    - [GetEventCouponResponse](#cloud-hashing-inspire-v1-GetEventCouponResponse)
    - [GetEventCouponsByAppActivityEventRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest)
    - [GetEventCouponsByAppActivityEventResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse)
    - [GetEventCouponsByAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppRequest)
    - [GetEventCouponsByAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppResponse)
    - [GetEventCouponsByOtherAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest)
    - [GetEventCouponsByOtherAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse)
    - [GetRegistrationInvitationByAppInviteeRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest)
    - [GetRegistrationInvitationByAppInviteeResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse)
    - [GetRegistrationInvitationRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationRequest)
    - [GetRegistrationInvitationResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationResponse)
    - [GetRegistrationInvitationsByAppInviterRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest)
    - [GetRegistrationInvitationsByAppInviterResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse)
    - [GetRegistrationInvitationsByAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest)
    - [GetRegistrationInvitationsByAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse)
    - [GetRegistrationInvitationsByOtherAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest)
    - [GetRegistrationInvitationsByOtherAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse)
    - [GetUserInvitationCodeByAppUserRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest)
    - [GetUserInvitationCodeByAppUserResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse)
    - [GetUserInvitationCodeByCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest)
    - [GetUserInvitationCodeByCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse)
    - [GetUserInvitationCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeRequest)
    - [GetUserInvitationCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeResponse)
    - [GetUserInvitationCodesByAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest)
    - [GetUserInvitationCodesByAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse)
    - [GetUserInvitationCodesByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest)
    - [GetUserInvitationCodesByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse)
    - [GetUserInvitationCodesRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesRequest)
    - [GetUserInvitationCodesResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesResponse)
    - [GetUserKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest)
    - [GetUserKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse)
    - [GetUserKpiSettingRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingRequest)
    - [GetUserKpiSettingResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingResponse)
    - [GetUserSpecialReductionRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionRequest)
    - [GetUserSpecialReductionResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionResponse)
    - [GetUserSpecialReductionsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest)
    - [GetUserSpecialReductionsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse)
    - [GetUserSpecialReductionsByAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest)
    - [GetUserSpecialReductionsByAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse)
    - [GetUserSpecialReductionsByAppUserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest)
    - [GetUserSpecialReductionsByAppUserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse)
    - [GetUserSpecialReductionsByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest)
    - [GetUserSpecialReductionsByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse)
    - [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation)
    - [UpdateActivityRequest](#cloud-hashing-inspire-v1-UpdateActivityRequest)
    - [UpdateActivityResponse](#cloud-hashing-inspire-v1-UpdateActivityResponse)
    - [UpdateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest)
    - [UpdateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse)
    - [UpdateAppCouponSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest)
    - [UpdateAppCouponSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse)
    - [UpdateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest)
    - [UpdateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse)
    - [UpdateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest)
    - [UpdateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse)
    - [UpdateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest)
    - [UpdateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse)
    - [UpdateCouponAllocatedRequest](#cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest)
    - [UpdateCouponAllocatedResponse](#cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse)
    - [UpdateCouponPoolRequest](#cloud-hashing-inspire-v1-UpdateCouponPoolRequest)
    - [UpdateCouponPoolResponse](#cloud-hashing-inspire-v1-UpdateCouponPoolResponse)
    - [UpdateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest)
    - [UpdateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse)
    - [UpdateDiscountPoolRequest](#cloud-hashing-inspire-v1-UpdateDiscountPoolRequest)
    - [UpdateDiscountPoolResponse](#cloud-hashing-inspire-v1-UpdateDiscountPoolResponse)
    - [UpdateEventCouponRequest](#cloud-hashing-inspire-v1-UpdateEventCouponRequest)
    - [UpdateEventCouponResponse](#cloud-hashing-inspire-v1-UpdateEventCouponResponse)
    - [UpdateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest)
    - [UpdateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse)
    - [UpdateUserKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest)
    - [UpdateUserKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse)
    - [UpdateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest)
    - [UpdateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse)
    - [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode)
    - [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting)
    - [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction)
  
    - [CloudHashingInspire](#cloud-hashing-inspire-v1-CloudHashingInspire)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_cloud-hashing-inspire_cloud-hashing-inspire-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-inspire/cloud-hashing-inspire.proto



<a name="cloud-hashing-inspire-v1-Activity"></a>

### Activity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| SystemActivity | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-AppCommissionSetting"></a>

### AppCommissionSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Level | [uint32](#uint32) |  |  |
| InvitationDiscount | [bool](#bool) |  |  |
| UniqueSetting | [bool](#bool) |  |  |
| KPISetting | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-AppCouponSetting"></a>

### AppCouponSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| DominationLimit | [double](#double) |  |  |
| TotalLimit | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-AppInvitationSetting"></a>

### AppInvitationSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Count | [uint32](#uint32) |  |  |
| Discount | [uint32](#uint32) |  |  |
| Title | [string](#string) |  |  |
| BadgeLarge | [string](#string) |  |  |
| BadgeSmall | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-AppPurchaseAmountSetting"></a>

### AppPurchaseAmountSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [uint32](#uint32) |  |  |
| Title | [string](#string) |  |  |
| BadgeLarge | [string](#string) |  |  |
| BadgeSmall | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CommissionCoinSetting"></a>

### CommissionCoinSetting
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Using | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-CouponAllocated"></a>

### CouponAllocated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CouponAllocatedDetail"></a>

### CouponAllocatedDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allocated | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |
| Coupon | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |
| Discount | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CouponPool"></a>

### CouponPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Denomination | [double](#double) |  |  |
| Circulation | [int32](#int32) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest"></a>

### CreateActivityForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse"></a>

### CreateActivityForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityRequest"></a>

### CreateActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityResponse"></a>

### CreateActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest"></a>

### CreateAppCommissionSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse"></a>

### CreateAppCommissionSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest"></a>

### CreateAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse"></a>

### CreateAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCouponSettingRequest"></a>

### CreateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCouponSettingResponse"></a>

### CreateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest"></a>

### CreateAppInvitationSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse"></a>

### CreateAppInvitationSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest"></a>

### CreateAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse"></a>

### CreateAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest"></a>

### CreateAppPurchaseAmountSettingForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse"></a>

### CreateAppPurchaseAmountSettingForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest"></a>

### CreateAppPurchaseAmountSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse"></a>

### CreateAppPurchaseAmountSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest"></a>

### CreateAppPurchaseAmountSettingForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse"></a>

### CreateAppPurchaseAmountSettingForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest"></a>

### CreateAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse"></a>

### CreateAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest"></a>

### CreateCommissionCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse"></a>

### CreateCommissionCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest"></a>

### CreateCouponAllocatedForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse"></a>

### CreateCouponAllocatedForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest"></a>

### CreateCouponAllocatedForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse"></a>

### CreateCouponAllocatedForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedRequest"></a>

### CreateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedResponse"></a>

### CreateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest"></a>

### CreateCouponPoolForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse"></a>

### CreateCouponPoolForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolRequest"></a>

### CreateCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolResponse"></a>

### CreateCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest"></a>

### CreateDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse"></a>

### CreateDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest"></a>

### CreateDiscountPoolForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse"></a>

### CreateDiscountPoolForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolRequest"></a>

### CreateDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolResponse"></a>

### CreateDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest"></a>

### CreateEventCouponForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse"></a>

### CreateEventCouponForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponRequest"></a>

### CreateEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponResponse"></a>

### CreateEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest"></a>

### CreateRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse"></a>

### CreateRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest"></a>

### CreateUserInvitationCodeForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse"></a>

### CreateUserInvitationCodeForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest"></a>

### CreateUserInvitationCodeForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse"></a>

### CreateUserInvitationCodeForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest"></a>

### CreateUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse"></a>

### CreateUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserKpiSettingRequest"></a>

### CreateUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserKpiSettingResponse"></a>

### CreateUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest"></a>

### CreateUserSpecialReductionForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse"></a>

### CreateUserSpecialReductionForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest"></a>

### CreateUserSpecialReductionForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse"></a>

### CreateUserSpecialReductionForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest"></a>

### CreateUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse"></a>

### CreateUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-DefaultKpiSetting"></a>

### DefaultKpiSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-DiscountPool"></a>

### DiscountPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Discount | [uint32](#uint32) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-EventCoupon"></a>

### EventCoupon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ActivityID | [string](#string) |  |  |
| Event | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Count | [uint32](#uint32) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByAppRequest"></a>

### GetActivitiesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByAppResponse"></a>

### GetActivitiesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Activity](#cloud-hashing-inspire-v1-Activity) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest"></a>

### GetActivitiesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse"></a>

### GetActivitiesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Activity](#cloud-hashing-inspire-v1-Activity) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetActivityByAppNameRequest"></a>

### GetActivityByAppNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityByAppNameResponse"></a>

### GetActivityByAppNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityRequest"></a>

### GetActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityResponse"></a>

### GetActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest"></a>

### GetAppCommissionSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse"></a>

### GetAppCommissionSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest"></a>

### GetAppCommissionSettingByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse"></a>

### GetAppCommissionSettingByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingRequest"></a>

### GetAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingResponse"></a>

### GetAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest"></a>

### GetAppCouponSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse"></a>

### GetAppCouponSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingRequest"></a>

### GetAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingResponse"></a>

### GetAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingRequest"></a>

### GetAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingResponse"></a>

### GetAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest"></a>

### GetAppInvitationSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse"></a>

### GetAppInvitationSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest"></a>

### GetAppInvitationSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse"></a>

### GetAppInvitationSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest"></a>

### GetAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse"></a>

### GetAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest"></a>

### GetAppPurchaseAmountSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse"></a>

### GetAppPurchaseAmountSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest"></a>

### GetAppPurchaseAmountSettingsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse"></a>

### GetAppPurchaseAmountSettingsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest"></a>

### GetAppPurchaseAmountSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse"></a>

### GetAppPurchaseAmountSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest"></a>

### GetAppPurchaseAmountSettingsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse"></a>

### GetAppPurchaseAmountSettingsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest"></a>

### GetCommissionCoinSettingsRequest







<a name="cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse"></a>

### GetCommissionCoinSettingsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest"></a>

### GetCouponAllocatedDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse"></a>

### GetCouponAllocatedDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedRequest"></a>

### GetCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedResponse"></a>

### GetCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolRequest"></a>

### GetCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolResponse"></a>

### GetCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest"></a>

### GetCouponPoolsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse"></a>

### GetCouponPoolsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest"></a>

### GetCouponPoolsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse"></a>

### GetCouponPoolsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest"></a>

### GetCouponPoolsByOtherAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse"></a>

### GetCouponPoolsByOtherAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest"></a>

### GetCouponPoolsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse"></a>

### GetCouponPoolsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest"></a>

### GetCouponsAllocatedByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse"></a>

### GetCouponsAllocatedByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest"></a>

### GetCouponsAllocatedByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse"></a>

### GetCouponsAllocatedByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest"></a>

### GetCouponsAllocatedByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse"></a>

### GetCouponsAllocatedByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest"></a>

### GetCouponsAllocatedDetailByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse"></a>

### GetCouponsAllocatedDetailByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest"></a>

### GetCouponsAllocatedDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse"></a>

### GetCouponsAllocatedDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest"></a>

### GetDefaultKpiSettingByAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse"></a>

### GetDefaultKpiSettingByAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest"></a>

### GetDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse"></a>

### GetDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolRequest"></a>

### GetDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolResponse"></a>

### GetDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest"></a>

### GetDiscountPoolsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse"></a>

### GetDiscountPoolsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest"></a>

### GetDiscountPoolsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse"></a>

### GetDiscountPoolsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest"></a>

### GetDiscountPoolsByOtherAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse"></a>

### GetDiscountPoolsByOtherAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest"></a>

### GetDiscountPoolsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse"></a>

### GetDiscountPoolsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponRequest"></a>

### GetEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponResponse"></a>

### GetEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest"></a>

### GetEventCouponsByAppActivityEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| ActivityID | [string](#string) |  |  |
| Event | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse"></a>

### GetEventCouponsByAppActivityEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppRequest"></a>

### GetEventCouponsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppResponse"></a>

### GetEventCouponsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest"></a>

### GetEventCouponsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse"></a>

### GetEventCouponsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest"></a>

### GetRegistrationInvitationByAppInviteeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| InviteeID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse"></a>

### GetRegistrationInvitationByAppInviteeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationRequest"></a>

### GetRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationResponse"></a>

### GetRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest"></a>

### GetRegistrationInvitationsByAppInviterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| InviterID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse"></a>

### GetRegistrationInvitationsByAppInviterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest"></a>

### GetRegistrationInvitationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse"></a>

### GetRegistrationInvitationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest"></a>

### GetRegistrationInvitationsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse"></a>

### GetRegistrationInvitationsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest"></a>

### GetUserInvitationCodeByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse"></a>

### GetUserInvitationCodeByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest"></a>

### GetUserInvitationCodeByCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse"></a>

### GetUserInvitationCodeByCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeRequest"></a>

### GetUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeResponse"></a>

### GetUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest"></a>

### GetUserInvitationCodesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse"></a>

### GetUserInvitationCodesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest"></a>

### GetUserInvitationCodesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse"></a>

### GetUserInvitationCodesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesRequest"></a>

### GetUserInvitationCodesRequest







<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesResponse"></a>

### GetUserInvitationCodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest"></a>

### GetUserKpiSettingByAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse"></a>

### GetUserKpiSettingByAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingRequest"></a>

### GetUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingResponse"></a>

### GetUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionRequest"></a>

### GetUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionResponse"></a>

### GetUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest"></a>

### GetUserSpecialReductionsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse"></a>

### GetUserSpecialReductionsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest"></a>

### GetUserSpecialReductionsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse"></a>

### GetUserSpecialReductionsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest"></a>

### GetUserSpecialReductionsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse"></a>

### GetUserSpecialReductionsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest"></a>

### GetUserSpecialReductionsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse"></a>

### GetUserSpecialReductionsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-RegistrationInvitation"></a>

### RegistrationInvitation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InviterID | [string](#string) |  |  |
| InviteeID | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateActivityRequest"></a>

### UpdateActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateActivityResponse"></a>

### UpdateActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest"></a>

### UpdateAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse"></a>

### UpdateAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest"></a>

### UpdateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse"></a>

### UpdateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest"></a>

### UpdateAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse"></a>

### UpdateAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest"></a>

### UpdateAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse"></a>

### UpdateAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest"></a>

### UpdateCommissionCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse"></a>

### UpdateCommissionCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest"></a>

### UpdateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse"></a>

### UpdateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponPoolRequest"></a>

### UpdateCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponPoolResponse"></a>

### UpdateCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest"></a>

### UpdateDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse"></a>

### UpdateDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDiscountPoolRequest"></a>

### UpdateDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDiscountPoolResponse"></a>

### UpdateDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateEventCouponRequest"></a>

### UpdateEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateEventCouponResponse"></a>

### UpdateEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest"></a>

### UpdateRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse"></a>

### UpdateRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest"></a>

### UpdateUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse"></a>

### UpdateUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest"></a>

### UpdateUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse"></a>

### UpdateUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-UserInvitationCode"></a>

### UserInvitationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InvitationCode | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-UserKpiSetting"></a>

### UserKpiSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-UserSpecialReduction"></a>

### UserSpecialReduction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |





 

 

 


<a name="cloud-hashing-inspire-v1-CloudHashingInspire"></a>

### CloudHashingInspire
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| CreateCommissionCoinSetting | [CreateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest) | [CreateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse) |  |
| UpdateCommissionCoinSetting | [UpdateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest) | [UpdateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse) |  |
| GetCommissionCoinSettings | [GetCommissionCoinSettingsRequest](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest) | [GetCommissionCoinSettingsResponse](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse) |  |
| CreateAppCommissionSetting | [CreateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest) | [CreateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse) |  |
| CreateAppCommissionSettingForOtherApp | [CreateAppCommissionSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest) | [CreateAppCommissionSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse) |  |
| UpdateAppCommissionSetting | [UpdateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest) | [UpdateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse) |  |
| GetAppCommissionSetting | [GetAppCommissionSettingRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingRequest) | [GetAppCommissionSettingResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingResponse) |  |
| GetAppCommissionSettingByApp | [GetAppCommissionSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest) | [GetAppCommissionSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse) |  |
| GetAppCommissionSettingByOtherApp | [GetAppCommissionSettingByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest) | [GetAppCommissionSettingByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse) |  |
| CreateAppInvitationSetting | [CreateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest) | [CreateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse) |  |
| CreateAppInvitationSettingForOtherApp | [CreateAppInvitationSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest) | [CreateAppInvitationSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse) |  |
| GetAppInvitationSetting | [GetAppInvitationSettingRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingRequest) | [GetAppInvitationSettingResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingResponse) |  |
| GetAppInvitationSettingsByApp | [GetAppInvitationSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest) | [GetAppInvitationSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse) |  |
| GetAppInvitationSettingsByOtherApp | [GetAppInvitationSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest) | [GetAppInvitationSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse) |  |
| UpdateAppInvitationSetting | [UpdateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest) | [UpdateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse) |  |
| CreateAppPurchaseAmountSetting | [CreateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest) | [CreateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse) |  |
| CreateAppPurchaseAmountSettingForOtherApp | [CreateAppPurchaseAmountSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest) | [CreateAppPurchaseAmountSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse) |  |
| CreateAppPurchaseAmountSettingForOtherAppUser | [CreateAppPurchaseAmountSettingForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest) | [CreateAppPurchaseAmountSettingForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse) |  |
| CreateAppPurchaseAmountSettingForAppOtherUser | [CreateAppPurchaseAmountSettingForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest) | [CreateAppPurchaseAmountSettingForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse) |  |
| UpdateAppPurchaseAmountSetting | [UpdateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest) | [UpdateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse) |  |
| GetAppPurchaseAmountSetting | [GetAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest) | [GetAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse) |  |
| GetAppPurchaseAmountSettingsByApp | [GetAppPurchaseAmountSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest) | [GetAppPurchaseAmountSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse) |  |
| GetAppPurchaseAmountSettingsByAppUser | [GetAppPurchaseAmountSettingsByAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest) | [GetAppPurchaseAmountSettingsByAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse) |  |
| GetAppPurchaseAmountSettingsByOtherApp | [GetAppPurchaseAmountSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest) | [GetAppPurchaseAmountSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse) |  |
| GetAppPurchaseAmountSettingsByOtherAppUser | [GetAppPurchaseAmountSettingsByOtherAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest) | [GetAppPurchaseAmountSettingsByOtherAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse) |  |
| CreateRegistrationInvitation | [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest) | [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse) |  |
| CreateRegistrationInvitationRevert | [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest) | [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse) |  |
| UpdateRegistrationInvitation | [UpdateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest) | [UpdateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse) |  |
| GetRegistrationInvitation | [GetRegistrationInvitationRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationRequest) | [GetRegistrationInvitationResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationResponse) |  |
| GetRegistrationInvitationsByApp | [GetRegistrationInvitationsByAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest) | [GetRegistrationInvitationsByAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse) |  |
| GetRegistrationInvitationsByOtherApp | [GetRegistrationInvitationsByOtherAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest) | [GetRegistrationInvitationsByOtherAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse) |  |
| GetRegistrationInvitationsByAppInviter | [GetRegistrationInvitationsByAppInviterRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest) | [GetRegistrationInvitationsByAppInviterResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse) |  |
| GetRegistrationInvitationByAppInvitee | [GetRegistrationInvitationByAppInviteeRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest) | [GetRegistrationInvitationByAppInviteeResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse) |  |
| CreateUserInvitationCode | [CreateUserInvitationCodeRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest) | [CreateUserInvitationCodeResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse) |  |
| CreateUserInvitationCodeForOtherAppUser | [CreateUserInvitationCodeForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest) | [CreateUserInvitationCodeForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse) |  |
| CreateUserInvitationCodeForAppOtherUser | [CreateUserInvitationCodeForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest) | [CreateUserInvitationCodeForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse) |  |
| GetUserInvitationCode | [GetUserInvitationCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeRequest) | [GetUserInvitationCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeResponse) |  |
| GetUserInvitationCodeByAppUser | [GetUserInvitationCodeByAppUserRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest) | [GetUserInvitationCodeByAppUserResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse) |  |
| GetUserInvitationCodes | [GetUserInvitationCodesRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesRequest) | [GetUserInvitationCodesResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesResponse) |  |
| GetUserInvitationCodesByApp | [GetUserInvitationCodesByAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest) | [GetUserInvitationCodesByAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse) |  |
| GetUserInvitationCodesByOtherApp | [GetUserInvitationCodesByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest) | [GetUserInvitationCodesByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse) |  |
| GetUserInvitationCodeByCode | [GetUserInvitationCodeByCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest) | [GetUserInvitationCodeByCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse) |  |
| CreateCouponAllocated | [CreateCouponAllocatedRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedRequest) | [CreateCouponAllocatedResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedResponse) |  |
| CreateCouponAllocatedForAppOtherUser | [CreateCouponAllocatedForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest) | [CreateCouponAllocatedForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse) |  |
| CreateCouponAllocatedForOtherAppUser | [CreateCouponAllocatedForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest) | [CreateCouponAllocatedForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse) |  |
| GetCouponAllocated | [GetCouponAllocatedRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedRequest) | [GetCouponAllocatedResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedResponse) |  |
| GetCouponsAllocatedByApp | [GetCouponsAllocatedByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest) | [GetCouponsAllocatedByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse) |  |
| GetCouponsAllocatedByOtherApp | [GetCouponsAllocatedByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest) | [GetCouponsAllocatedByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse) |  |
| GetCouponsAllocatedByAppUser | [GetCouponsAllocatedByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest) | [GetCouponsAllocatedByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse) |  |
| UpdateCouponAllocated | [UpdateCouponAllocatedRequest](#cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest) | [UpdateCouponAllocatedResponse](#cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse) |  |
| GetCouponAllocatedDetail | [GetCouponAllocatedDetailRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest) | [GetCouponAllocatedDetailResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse) |  |
| GetCouponsAllocatedDetailByApp | [GetCouponsAllocatedDetailByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest) | [GetCouponsAllocatedDetailByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse) |  |
| GetCouponsAllocatedDetailByAppUser | [GetCouponsAllocatedDetailByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest) | [GetCouponsAllocatedDetailByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse) |  |
| CreateCouponPool | [CreateCouponPoolRequest](#cloud-hashing-inspire-v1-CreateCouponPoolRequest) | [CreateCouponPoolResponse](#cloud-hashing-inspire-v1-CreateCouponPoolResponse) |  |
| CreateCouponPoolForOtherApp | [CreateCouponPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest) | [CreateCouponPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse) |  |
| UpdateCouponPool | [UpdateCouponPoolRequest](#cloud-hashing-inspire-v1-UpdateCouponPoolRequest) | [UpdateCouponPoolResponse](#cloud-hashing-inspire-v1-UpdateCouponPoolResponse) |  |
| GetCouponPool | [GetCouponPoolRequest](#cloud-hashing-inspire-v1-GetCouponPoolRequest) | [GetCouponPoolResponse](#cloud-hashing-inspire-v1-GetCouponPoolResponse) |  |
| GetCouponPoolsByApp | [GetCouponPoolsByAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest) | [GetCouponPoolsByAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse) |  |
| GetCouponPoolsByOtherApp | [GetCouponPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest) | [GetCouponPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse) |  |
| GetCouponPoolsByAppReleaser | [GetCouponPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest) | [GetCouponPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse) |  |
| GetCouponPoolsByOtherAppReleaser | [GetCouponPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest) | [GetCouponPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse) |  |
| CreateDiscountPool | [CreateDiscountPoolRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolRequest) | [CreateDiscountPoolResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolResponse) |  |
| CreateDiscountPoolForOtherApp | [CreateDiscountPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest) | [CreateDiscountPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse) |  |
| UpdateDiscountPool | [UpdateDiscountPoolRequest](#cloud-hashing-inspire-v1-UpdateDiscountPoolRequest) | [UpdateDiscountPoolResponse](#cloud-hashing-inspire-v1-UpdateDiscountPoolResponse) |  |
| GetDiscountPool | [GetDiscountPoolRequest](#cloud-hashing-inspire-v1-GetDiscountPoolRequest) | [GetDiscountPoolResponse](#cloud-hashing-inspire-v1-GetDiscountPoolResponse) |  |
| GetDiscountPoolsByApp | [GetDiscountPoolsByAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest) | [GetDiscountPoolsByAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse) |  |
| GetDiscountPoolsByOtherApp | [GetDiscountPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest) | [GetDiscountPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse) |  |
| GetDiscountPoolsByAppReleaser | [GetDiscountPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest) | [GetDiscountPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse) |  |
| GetDiscountPoolsByOtherAppReleaser | [GetDiscountPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest) | [GetDiscountPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse) |  |
| CreateAppCouponSetting | [CreateAppCouponSettingRequest](#cloud-hashing-inspire-v1-CreateAppCouponSettingRequest) | [CreateAppCouponSettingResponse](#cloud-hashing-inspire-v1-CreateAppCouponSettingResponse) |  |
| GetAppCouponSetting | [GetAppCouponSettingRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingRequest) | [GetAppCouponSettingResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingResponse) |  |
| GetAppCouponSettingByApp | [GetAppCouponSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest) | [GetAppCouponSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse) |  |
| UpdateAppCouponSetting | [UpdateAppCouponSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest) | [UpdateAppCouponSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse) |  |
| CreateDefaultKpiSetting | [CreateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest) | [CreateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse) |  |
| GetDefaultKpiSetting | [GetDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest) | [GetDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse) |  |
| GetDefaultKpiSettingByAppGood | [GetDefaultKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest) | [GetDefaultKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse) |  |
| UpdateDefaultKpiSetting | [UpdateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest) | [UpdateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse) |  |
| CreateUserKpiSetting | [CreateUserKpiSettingRequest](#cloud-hashing-inspire-v1-CreateUserKpiSettingRequest) | [CreateUserKpiSettingResponse](#cloud-hashing-inspire-v1-CreateUserKpiSettingResponse) |  |
| GetUserKpiSetting | [GetUserKpiSettingRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingRequest) | [GetUserKpiSettingResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingResponse) |  |
| GetUserKpiSettingByAppGood | [GetUserKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest) | [GetUserKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse) |  |
| UpdateUserKpiSetting | [UpdateUserKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest) | [UpdateUserKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse) |  |
| CreateUserSpecialReduction | [CreateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest) | [CreateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse) |  |
| CreateUserSpecialReductionForAppOtherUser | [CreateUserSpecialReductionForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest) | [CreateUserSpecialReductionForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse) |  |
| CreateUserSpecialReductionForOtherAppUser | [CreateUserSpecialReductionForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest) | [CreateUserSpecialReductionForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse) |  |
| GetUserSpecialReduction | [GetUserSpecialReductionRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionRequest) | [GetUserSpecialReductionResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionResponse) |  |
| GetUserSpecialReductionsByApp | [GetUserSpecialReductionsByAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest) | [GetUserSpecialReductionsByAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse) |  |
| GetUserSpecialReductionsByOtherApp | [GetUserSpecialReductionsByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest) | [GetUserSpecialReductionsByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse) |  |
| GetUserSpecialReductionsByAppReleaser | [GetUserSpecialReductionsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest) | [GetUserSpecialReductionsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse) |  |
| GetUserSpecialReductionsByAppUser | [GetUserSpecialReductionsByAppUserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest) | [GetUserSpecialReductionsByAppUserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse) |  |
| UpdateUserSpecialReduction | [UpdateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest) | [UpdateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse) |  |
| CreateActivity | [CreateActivityRequest](#cloud-hashing-inspire-v1-CreateActivityRequest) | [CreateActivityResponse](#cloud-hashing-inspire-v1-CreateActivityResponse) |  |
| CreateActivityForOtherApp | [CreateActivityForOtherAppRequest](#cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest) | [CreateActivityForOtherAppResponse](#cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse) |  |
| UpdateActivity | [UpdateActivityRequest](#cloud-hashing-inspire-v1-UpdateActivityRequest) | [UpdateActivityResponse](#cloud-hashing-inspire-v1-UpdateActivityResponse) |  |
| GetActivity | [GetActivityRequest](#cloud-hashing-inspire-v1-GetActivityRequest) | [GetActivityResponse](#cloud-hashing-inspire-v1-GetActivityResponse) |  |
| GetActivityByAppName | [GetActivityByAppNameRequest](#cloud-hashing-inspire-v1-GetActivityByAppNameRequest) | [GetActivityByAppNameResponse](#cloud-hashing-inspire-v1-GetActivityByAppNameResponse) |  |
| GetActivitiesByApp | [GetActivitiesByAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByAppRequest) | [GetActivitiesByAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByAppResponse) |  |
| GetActivitiesByOtherApp | [GetActivitiesByOtherAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest) | [GetActivitiesByOtherAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse) |  |
| CreateEventCoupon | [CreateEventCouponRequest](#cloud-hashing-inspire-v1-CreateEventCouponRequest) | [CreateEventCouponResponse](#cloud-hashing-inspire-v1-CreateEventCouponResponse) |  |
| CreateEventCouponForOtherApp | [CreateEventCouponForOtherAppRequest](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest) | [CreateEventCouponForOtherAppResponse](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse) |  |
| UpdateEventCoupon | [UpdateEventCouponRequest](#cloud-hashing-inspire-v1-UpdateEventCouponRequest) | [UpdateEventCouponResponse](#cloud-hashing-inspire-v1-UpdateEventCouponResponse) |  |
| GetEventCoupon | [GetEventCouponRequest](#cloud-hashing-inspire-v1-GetEventCouponRequest) | [GetEventCouponResponse](#cloud-hashing-inspire-v1-GetEventCouponResponse) |  |
| GetEventCouponsByAppActivityEvent | [GetEventCouponsByAppActivityEventRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest) | [GetEventCouponsByAppActivityEventResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse) |  |
| GetEventCouponsByApp | [GetEventCouponsByAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppRequest) | [GetEventCouponsByAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppResponse) |  |
| GetEventCouponsByOtherApp | [GetEventCouponsByOtherAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest) | [GetEventCouponsByOtherAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse) |  |

 



<a name="npool_cloud-hashing-inspire_cloud-hashing-inspire-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-inspire/cloud-hashing-inspire.proto



<a name="cloud-hashing-inspire-v1-Activity"></a>

### Activity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| SystemActivity | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-AppCommissionSetting"></a>

### AppCommissionSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Level | [uint32](#uint32) |  |  |
| InvitationDiscount | [bool](#bool) |  |  |
| UniqueSetting | [bool](#bool) |  |  |
| KPISetting | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-AppCouponSetting"></a>

### AppCouponSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| DominationLimit | [double](#double) |  |  |
| TotalLimit | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-AppInvitationSetting"></a>

### AppInvitationSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Count | [uint32](#uint32) |  |  |
| Discount | [uint32](#uint32) |  |  |
| Title | [string](#string) |  |  |
| BadgeLarge | [string](#string) |  |  |
| BadgeSmall | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-AppPurchaseAmountSetting"></a>

### AppPurchaseAmountSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [uint32](#uint32) |  |  |
| Title | [string](#string) |  |  |
| BadgeLarge | [string](#string) |  |  |
| BadgeSmall | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CommissionCoinSetting"></a>

### CommissionCoinSetting
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Using | [bool](#bool) |  |  |






<a name="cloud-hashing-inspire-v1-CouponAllocated"></a>

### CouponAllocated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CouponAllocatedDetail"></a>

### CouponAllocatedDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allocated | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |
| Coupon | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |
| Discount | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CouponPool"></a>

### CouponPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Denomination | [double](#double) |  |  |
| Circulation | [int32](#int32) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest"></a>

### CreateActivityForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse"></a>

### CreateActivityForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityRequest"></a>

### CreateActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateActivityResponse"></a>

### CreateActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest"></a>

### CreateAppCommissionSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse"></a>

### CreateAppCommissionSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest"></a>

### CreateAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse"></a>

### CreateAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCouponSettingRequest"></a>

### CreateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppCouponSettingResponse"></a>

### CreateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest"></a>

### CreateAppInvitationSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse"></a>

### CreateAppInvitationSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest"></a>

### CreateAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse"></a>

### CreateAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest"></a>

### CreateAppPurchaseAmountSettingForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse"></a>

### CreateAppPurchaseAmountSettingForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest"></a>

### CreateAppPurchaseAmountSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse"></a>

### CreateAppPurchaseAmountSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest"></a>

### CreateAppPurchaseAmountSettingForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse"></a>

### CreateAppPurchaseAmountSettingForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest"></a>

### CreateAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse"></a>

### CreateAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest"></a>

### CreateCommissionCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse"></a>

### CreateCommissionCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest"></a>

### CreateCouponAllocatedForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse"></a>

### CreateCouponAllocatedForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest"></a>

### CreateCouponAllocatedForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse"></a>

### CreateCouponAllocatedForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedRequest"></a>

### CreateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponAllocatedResponse"></a>

### CreateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest"></a>

### CreateCouponPoolForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse"></a>

### CreateCouponPoolForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolRequest"></a>

### CreateCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateCouponPoolResponse"></a>

### CreateCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest"></a>

### CreateDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse"></a>

### CreateDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest"></a>

### CreateDiscountPoolForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse"></a>

### CreateDiscountPoolForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolRequest"></a>

### CreateDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateDiscountPoolResponse"></a>

### CreateDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest"></a>

### CreateEventCouponForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse"></a>

### CreateEventCouponForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponRequest"></a>

### CreateEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateEventCouponResponse"></a>

### CreateEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest"></a>

### CreateRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse"></a>

### CreateRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest"></a>

### CreateUserInvitationCodeForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse"></a>

### CreateUserInvitationCodeForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest"></a>

### CreateUserInvitationCodeForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse"></a>

### CreateUserInvitationCodeForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest"></a>

### CreateUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse"></a>

### CreateUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserKpiSettingRequest"></a>

### CreateUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserKpiSettingResponse"></a>

### CreateUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest"></a>

### CreateUserSpecialReductionForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse"></a>

### CreateUserSpecialReductionForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest"></a>

### CreateUserSpecialReductionForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse"></a>

### CreateUserSpecialReductionForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest"></a>

### CreateUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse"></a>

### CreateUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-DefaultKpiSetting"></a>

### DefaultKpiSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-DiscountPool"></a>

### DiscountPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Discount | [uint32](#uint32) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-EventCoupon"></a>

### EventCoupon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ActivityID | [string](#string) |  |  |
| Event | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Count | [uint32](#uint32) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByAppRequest"></a>

### GetActivitiesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByAppResponse"></a>

### GetActivitiesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Activity](#cloud-hashing-inspire-v1-Activity) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest"></a>

### GetActivitiesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse"></a>

### GetActivitiesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Activity](#cloud-hashing-inspire-v1-Activity) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetActivityByAppNameRequest"></a>

### GetActivityByAppNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityByAppNameResponse"></a>

### GetActivityByAppNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityRequest"></a>

### GetActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetActivityResponse"></a>

### GetActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest"></a>

### GetAppCommissionSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse"></a>

### GetAppCommissionSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest"></a>

### GetAppCommissionSettingByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse"></a>

### GetAppCommissionSettingByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingRequest"></a>

### GetAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCommissionSettingResponse"></a>

### GetAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest"></a>

### GetAppCouponSettingByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse"></a>

### GetAppCouponSettingByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingRequest"></a>

### GetAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppCouponSettingResponse"></a>

### GetAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingRequest"></a>

### GetAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingResponse"></a>

### GetAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest"></a>

### GetAppInvitationSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse"></a>

### GetAppInvitationSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest"></a>

### GetAppInvitationSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse"></a>

### GetAppInvitationSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest"></a>

### GetAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse"></a>

### GetAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest"></a>

### GetAppPurchaseAmountSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse"></a>

### GetAppPurchaseAmountSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest"></a>

### GetAppPurchaseAmountSettingsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse"></a>

### GetAppPurchaseAmountSettingsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest"></a>

### GetAppPurchaseAmountSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse"></a>

### GetAppPurchaseAmountSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest"></a>

### GetAppPurchaseAmountSettingsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse"></a>

### GetAppPurchaseAmountSettingsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest"></a>

### GetCommissionCoinSettingsRequest







<a name="cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse"></a>

### GetCommissionCoinSettingsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest"></a>

### GetCouponAllocatedDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse"></a>

### GetCouponAllocatedDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedRequest"></a>

### GetCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponAllocatedResponse"></a>

### GetCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolRequest"></a>

### GetCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolResponse"></a>

### GetCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest"></a>

### GetCouponPoolsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse"></a>

### GetCouponPoolsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest"></a>

### GetCouponPoolsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse"></a>

### GetCouponPoolsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest"></a>

### GetCouponPoolsByOtherAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse"></a>

### GetCouponPoolsByOtherAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest"></a>

### GetCouponPoolsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse"></a>

### GetCouponPoolsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest"></a>

### GetCouponsAllocatedByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse"></a>

### GetCouponsAllocatedByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest"></a>

### GetCouponsAllocatedByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse"></a>

### GetCouponsAllocatedByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest"></a>

### GetCouponsAllocatedByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse"></a>

### GetCouponsAllocatedByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest"></a>

### GetCouponsAllocatedDetailByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse"></a>

### GetCouponsAllocatedDetailByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest"></a>

### GetCouponsAllocatedDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse"></a>

### GetCouponsAllocatedDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest"></a>

### GetDefaultKpiSettingByAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse"></a>

### GetDefaultKpiSettingByAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest"></a>

### GetDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse"></a>

### GetDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolRequest"></a>

### GetDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolResponse"></a>

### GetDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest"></a>

### GetDiscountPoolsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse"></a>

### GetDiscountPoolsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest"></a>

### GetDiscountPoolsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse"></a>

### GetDiscountPoolsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest"></a>

### GetDiscountPoolsByOtherAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse"></a>

### GetDiscountPoolsByOtherAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest"></a>

### GetDiscountPoolsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse"></a>

### GetDiscountPoolsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponRequest"></a>

### GetEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponResponse"></a>

### GetEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest"></a>

### GetEventCouponsByAppActivityEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| ActivityID | [string](#string) |  |  |
| Event | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse"></a>

### GetEventCouponsByAppActivityEventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppRequest"></a>

### GetEventCouponsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByAppResponse"></a>

### GetEventCouponsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest"></a>

### GetEventCouponsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse"></a>

### GetEventCouponsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest"></a>

### GetRegistrationInvitationByAppInviteeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| InviteeID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse"></a>

### GetRegistrationInvitationByAppInviteeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationRequest"></a>

### GetRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationResponse"></a>

### GetRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest"></a>

### GetRegistrationInvitationsByAppInviterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| InviterID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse"></a>

### GetRegistrationInvitationsByAppInviterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest"></a>

### GetRegistrationInvitationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse"></a>

### GetRegistrationInvitationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest"></a>

### GetRegistrationInvitationsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse"></a>

### GetRegistrationInvitationsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest"></a>

### GetUserInvitationCodeByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse"></a>

### GetUserInvitationCodeByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest"></a>

### GetUserInvitationCodeByCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse"></a>

### GetUserInvitationCodeByCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeRequest"></a>

### GetUserInvitationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodeResponse"></a>

### GetUserInvitationCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest"></a>

### GetUserInvitationCodesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse"></a>

### GetUserInvitationCodesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest"></a>

### GetUserInvitationCodesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse"></a>

### GetUserInvitationCodesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesRequest"></a>

### GetUserInvitationCodesRequest







<a name="cloud-hashing-inspire-v1-GetUserInvitationCodesResponse"></a>

### GetUserInvitationCodesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserInvitationCode](#cloud-hashing-inspire-v1-UserInvitationCode) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest"></a>

### GetUserKpiSettingByAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse"></a>

### GetUserKpiSettingByAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingRequest"></a>

### GetUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserKpiSettingResponse"></a>

### GetUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionRequest"></a>

### GetUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionResponse"></a>

### GetUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest"></a>

### GetUserSpecialReductionsByAppReleaserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse"></a>

### GetUserSpecialReductionsByAppReleaserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest"></a>

### GetUserSpecialReductionsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse"></a>

### GetUserSpecialReductionsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest"></a>

### GetUserSpecialReductionsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse"></a>

### GetUserSpecialReductionsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest"></a>

### GetUserSpecialReductionsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse"></a>

### GetUserSpecialReductionsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) | repeated |  |






<a name="cloud-hashing-inspire-v1-RegistrationInvitation"></a>

### RegistrationInvitation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InviterID | [string](#string) |  |  |
| InviteeID | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateActivityRequest"></a>

### UpdateActivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateActivityResponse"></a>

### UpdateActivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Activity](#cloud-hashing-inspire-v1-Activity) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest"></a>

### UpdateAppCommissionSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse"></a>

### UpdateAppCommissionSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCommissionSetting](#cloud-hashing-inspire-v1-AppCommissionSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest"></a>

### UpdateAppCouponSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse"></a>

### UpdateAppCouponSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppCouponSetting](#cloud-hashing-inspire-v1-AppCouponSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest"></a>

### UpdateAppInvitationSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse"></a>

### UpdateAppInvitationSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInvitationSetting](#cloud-hashing-inspire-v1-AppInvitationSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest"></a>

### UpdateAppPurchaseAmountSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse"></a>

### UpdateAppPurchaseAmountSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppPurchaseAmountSetting](#cloud-hashing-inspire-v1-AppPurchaseAmountSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest"></a>

### UpdateCommissionCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse"></a>

### UpdateCommissionCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CommissionCoinSetting](#cloud-hashing-inspire-v1-CommissionCoinSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest"></a>

### UpdateCouponAllocatedRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse"></a>

### UpdateCouponAllocatedResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponAllocated](#cloud-hashing-inspire-v1-CouponAllocated) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponPoolRequest"></a>

### UpdateCouponPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateCouponPoolResponse"></a>

### UpdateCouponPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CouponPool](#cloud-hashing-inspire-v1-CouponPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest"></a>

### UpdateDefaultKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse"></a>

### UpdateDefaultKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DefaultKpiSetting](#cloud-hashing-inspire-v1-DefaultKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDiscountPoolRequest"></a>

### UpdateDiscountPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateDiscountPoolResponse"></a>

### UpdateDiscountPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DiscountPool](#cloud-hashing-inspire-v1-DiscountPool) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateEventCouponRequest"></a>

### UpdateEventCouponRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateEventCouponResponse"></a>

### UpdateEventCouponResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EventCoupon](#cloud-hashing-inspire-v1-EventCoupon) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest"></a>

### UpdateRegistrationInvitationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse"></a>

### UpdateRegistrationInvitationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest"></a>

### UpdateUserKpiSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse"></a>

### UpdateUserKpiSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserKpiSetting](#cloud-hashing-inspire-v1-UserKpiSetting) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest"></a>

### UpdateUserSpecialReductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse"></a>

### UpdateUserSpecialReductionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |






<a name="cloud-hashing-inspire-v1-UserInvitationCode"></a>

### UserInvitationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InvitationCode | [string](#string) |  |  |






<a name="cloud-hashing-inspire-v1-UserKpiSetting"></a>

### UserKpiSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Percent | [int32](#int32) |  |  |






<a name="cloud-hashing-inspire-v1-UserSpecialReduction"></a>

### UserSpecialReduction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| ReleaseByUserID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |





 

 

 


<a name="cloud-hashing-inspire-v1-CloudHashingInspire"></a>

### CloudHashingInspire
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| CreateCommissionCoinSetting | [CreateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingRequest) | [CreateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-CreateCommissionCoinSettingResponse) |  |
| UpdateCommissionCoinSetting | [UpdateCommissionCoinSettingRequest](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingRequest) | [UpdateCommissionCoinSettingResponse](#cloud-hashing-inspire-v1-UpdateCommissionCoinSettingResponse) |  |
| GetCommissionCoinSettings | [GetCommissionCoinSettingsRequest](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsRequest) | [GetCommissionCoinSettingsResponse](#cloud-hashing-inspire-v1-GetCommissionCoinSettingsResponse) |  |
| CreateAppCommissionSetting | [CreateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingRequest) | [CreateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingResponse) |  |
| CreateAppCommissionSettingForOtherApp | [CreateAppCommissionSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppRequest) | [CreateAppCommissionSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppCommissionSettingForOtherAppResponse) |  |
| UpdateAppCommissionSetting | [UpdateAppCommissionSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingRequest) | [UpdateAppCommissionSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCommissionSettingResponse) |  |
| GetAppCommissionSetting | [GetAppCommissionSettingRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingRequest) | [GetAppCommissionSettingResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingResponse) |  |
| GetAppCommissionSettingByApp | [GetAppCommissionSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppRequest) | [GetAppCommissionSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByAppResponse) |  |
| GetAppCommissionSettingByOtherApp | [GetAppCommissionSettingByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppRequest) | [GetAppCommissionSettingByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppCommissionSettingByOtherAppResponse) |  |
| CreateAppInvitationSetting | [CreateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingRequest) | [CreateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingResponse) |  |
| CreateAppInvitationSettingForOtherApp | [CreateAppInvitationSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppRequest) | [CreateAppInvitationSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppInvitationSettingForOtherAppResponse) |  |
| GetAppInvitationSetting | [GetAppInvitationSettingRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingRequest) | [GetAppInvitationSettingResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingResponse) |  |
| GetAppInvitationSettingsByApp | [GetAppInvitationSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppRequest) | [GetAppInvitationSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByAppResponse) |  |
| GetAppInvitationSettingsByOtherApp | [GetAppInvitationSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppRequest) | [GetAppInvitationSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppInvitationSettingsByOtherAppResponse) |  |
| UpdateAppInvitationSetting | [UpdateAppInvitationSettingRequest](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingRequest) | [UpdateAppInvitationSettingResponse](#cloud-hashing-inspire-v1-UpdateAppInvitationSettingResponse) |  |
| CreateAppPurchaseAmountSetting | [CreateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingRequest) | [CreateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingResponse) |  |
| CreateAppPurchaseAmountSettingForOtherApp | [CreateAppPurchaseAmountSettingForOtherAppRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppRequest) | [CreateAppPurchaseAmountSettingForOtherAppResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppResponse) |  |
| CreateAppPurchaseAmountSettingForOtherAppUser | [CreateAppPurchaseAmountSettingForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserRequest) | [CreateAppPurchaseAmountSettingForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForOtherAppUserResponse) |  |
| CreateAppPurchaseAmountSettingForAppOtherUser | [CreateAppPurchaseAmountSettingForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserRequest) | [CreateAppPurchaseAmountSettingForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateAppPurchaseAmountSettingForAppOtherUserResponse) |  |
| UpdateAppPurchaseAmountSetting | [UpdateAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingRequest) | [UpdateAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-UpdateAppPurchaseAmountSettingResponse) |  |
| GetAppPurchaseAmountSetting | [GetAppPurchaseAmountSettingRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingRequest) | [GetAppPurchaseAmountSettingResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingResponse) |  |
| GetAppPurchaseAmountSettingsByApp | [GetAppPurchaseAmountSettingsByAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppRequest) | [GetAppPurchaseAmountSettingsByAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppResponse) |  |
| GetAppPurchaseAmountSettingsByAppUser | [GetAppPurchaseAmountSettingsByAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserRequest) | [GetAppPurchaseAmountSettingsByAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByAppUserResponse) |  |
| GetAppPurchaseAmountSettingsByOtherApp | [GetAppPurchaseAmountSettingsByOtherAppRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppRequest) | [GetAppPurchaseAmountSettingsByOtherAppResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppResponse) |  |
| GetAppPurchaseAmountSettingsByOtherAppUser | [GetAppPurchaseAmountSettingsByOtherAppUserRequest](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserRequest) | [GetAppPurchaseAmountSettingsByOtherAppUserResponse](#cloud-hashing-inspire-v1-GetAppPurchaseAmountSettingsByOtherAppUserResponse) |  |
| CreateRegistrationInvitation | [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest) | [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse) |  |
| CreateRegistrationInvitationRevert | [CreateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-CreateRegistrationInvitationRequest) | [CreateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-CreateRegistrationInvitationResponse) |  |
| UpdateRegistrationInvitation | [UpdateRegistrationInvitationRequest](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationRequest) | [UpdateRegistrationInvitationResponse](#cloud-hashing-inspire-v1-UpdateRegistrationInvitationResponse) |  |
| GetRegistrationInvitation | [GetRegistrationInvitationRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationRequest) | [GetRegistrationInvitationResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationResponse) |  |
| GetRegistrationInvitationsByApp | [GetRegistrationInvitationsByAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppRequest) | [GetRegistrationInvitationsByAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppResponse) |  |
| GetRegistrationInvitationsByOtherApp | [GetRegistrationInvitationsByOtherAppRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppRequest) | [GetRegistrationInvitationsByOtherAppResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByOtherAppResponse) |  |
| GetRegistrationInvitationsByAppInviter | [GetRegistrationInvitationsByAppInviterRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterRequest) | [GetRegistrationInvitationsByAppInviterResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationsByAppInviterResponse) |  |
| GetRegistrationInvitationByAppInvitee | [GetRegistrationInvitationByAppInviteeRequest](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeRequest) | [GetRegistrationInvitationByAppInviteeResponse](#cloud-hashing-inspire-v1-GetRegistrationInvitationByAppInviteeResponse) |  |
| CreateUserInvitationCode | [CreateUserInvitationCodeRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeRequest) | [CreateUserInvitationCodeResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeResponse) |  |
| CreateUserInvitationCodeForOtherAppUser | [CreateUserInvitationCodeForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserRequest) | [CreateUserInvitationCodeForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForOtherAppUserResponse) |  |
| CreateUserInvitationCodeForAppOtherUser | [CreateUserInvitationCodeForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserRequest) | [CreateUserInvitationCodeForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserInvitationCodeForAppOtherUserResponse) |  |
| GetUserInvitationCode | [GetUserInvitationCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeRequest) | [GetUserInvitationCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeResponse) |  |
| GetUserInvitationCodeByAppUser | [GetUserInvitationCodeByAppUserRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserRequest) | [GetUserInvitationCodeByAppUserResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByAppUserResponse) |  |
| GetUserInvitationCodes | [GetUserInvitationCodesRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesRequest) | [GetUserInvitationCodesResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesResponse) |  |
| GetUserInvitationCodesByApp | [GetUserInvitationCodesByAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppRequest) | [GetUserInvitationCodesByAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByAppResponse) |  |
| GetUserInvitationCodesByOtherApp | [GetUserInvitationCodesByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppRequest) | [GetUserInvitationCodesByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodesByOtherAppResponse) |  |
| GetUserInvitationCodeByCode | [GetUserInvitationCodeByCodeRequest](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeRequest) | [GetUserInvitationCodeByCodeResponse](#cloud-hashing-inspire-v1-GetUserInvitationCodeByCodeResponse) |  |
| CreateCouponAllocated | [CreateCouponAllocatedRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedRequest) | [CreateCouponAllocatedResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedResponse) |  |
| CreateCouponAllocatedForAppOtherUser | [CreateCouponAllocatedForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserRequest) | [CreateCouponAllocatedForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForAppOtherUserResponse) |  |
| CreateCouponAllocatedForOtherAppUser | [CreateCouponAllocatedForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserRequest) | [CreateCouponAllocatedForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateCouponAllocatedForOtherAppUserResponse) |  |
| GetCouponAllocated | [GetCouponAllocatedRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedRequest) | [GetCouponAllocatedResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedResponse) |  |
| GetCouponsAllocatedByApp | [GetCouponsAllocatedByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppRequest) | [GetCouponsAllocatedByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppResponse) |  |
| GetCouponsAllocatedByOtherApp | [GetCouponsAllocatedByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppRequest) | [GetCouponsAllocatedByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByOtherAppResponse) |  |
| GetCouponsAllocatedByAppUser | [GetCouponsAllocatedByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserRequest) | [GetCouponsAllocatedByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedByAppUserResponse) |  |
| UpdateCouponAllocated | [UpdateCouponAllocatedRequest](#cloud-hashing-inspire-v1-UpdateCouponAllocatedRequest) | [UpdateCouponAllocatedResponse](#cloud-hashing-inspire-v1-UpdateCouponAllocatedResponse) |  |
| GetCouponAllocatedDetail | [GetCouponAllocatedDetailRequest](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailRequest) | [GetCouponAllocatedDetailResponse](#cloud-hashing-inspire-v1-GetCouponAllocatedDetailResponse) |  |
| GetCouponsAllocatedDetailByApp | [GetCouponsAllocatedDetailByAppRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppRequest) | [GetCouponsAllocatedDetailByAppResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppResponse) |  |
| GetCouponsAllocatedDetailByAppUser | [GetCouponsAllocatedDetailByAppUserRequest](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserRequest) | [GetCouponsAllocatedDetailByAppUserResponse](#cloud-hashing-inspire-v1-GetCouponsAllocatedDetailByAppUserResponse) |  |
| CreateCouponPool | [CreateCouponPoolRequest](#cloud-hashing-inspire-v1-CreateCouponPoolRequest) | [CreateCouponPoolResponse](#cloud-hashing-inspire-v1-CreateCouponPoolResponse) |  |
| CreateCouponPoolForOtherApp | [CreateCouponPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppRequest) | [CreateCouponPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateCouponPoolForOtherAppResponse) |  |
| UpdateCouponPool | [UpdateCouponPoolRequest](#cloud-hashing-inspire-v1-UpdateCouponPoolRequest) | [UpdateCouponPoolResponse](#cloud-hashing-inspire-v1-UpdateCouponPoolResponse) |  |
| GetCouponPool | [GetCouponPoolRequest](#cloud-hashing-inspire-v1-GetCouponPoolRequest) | [GetCouponPoolResponse](#cloud-hashing-inspire-v1-GetCouponPoolResponse) |  |
| GetCouponPoolsByApp | [GetCouponPoolsByAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppRequest) | [GetCouponPoolsByAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppResponse) |  |
| GetCouponPoolsByOtherApp | [GetCouponPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppRequest) | [GetCouponPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppResponse) |  |
| GetCouponPoolsByAppReleaser | [GetCouponPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserRequest) | [GetCouponPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByAppReleaserResponse) |  |
| GetCouponPoolsByOtherAppReleaser | [GetCouponPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserRequest) | [GetCouponPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetCouponPoolsByOtherAppReleaserResponse) |  |
| CreateDiscountPool | [CreateDiscountPoolRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolRequest) | [CreateDiscountPoolResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolResponse) |  |
| CreateDiscountPoolForOtherApp | [CreateDiscountPoolForOtherAppRequest](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppRequest) | [CreateDiscountPoolForOtherAppResponse](#cloud-hashing-inspire-v1-CreateDiscountPoolForOtherAppResponse) |  |
| UpdateDiscountPool | [UpdateDiscountPoolRequest](#cloud-hashing-inspire-v1-UpdateDiscountPoolRequest) | [UpdateDiscountPoolResponse](#cloud-hashing-inspire-v1-UpdateDiscountPoolResponse) |  |
| GetDiscountPool | [GetDiscountPoolRequest](#cloud-hashing-inspire-v1-GetDiscountPoolRequest) | [GetDiscountPoolResponse](#cloud-hashing-inspire-v1-GetDiscountPoolResponse) |  |
| GetDiscountPoolsByApp | [GetDiscountPoolsByAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppRequest) | [GetDiscountPoolsByAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppResponse) |  |
| GetDiscountPoolsByOtherApp | [GetDiscountPoolsByOtherAppRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppRequest) | [GetDiscountPoolsByOtherAppResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppResponse) |  |
| GetDiscountPoolsByAppReleaser | [GetDiscountPoolsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserRequest) | [GetDiscountPoolsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByAppReleaserResponse) |  |
| GetDiscountPoolsByOtherAppReleaser | [GetDiscountPoolsByOtherAppReleaserRequest](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserRequest) | [GetDiscountPoolsByOtherAppReleaserResponse](#cloud-hashing-inspire-v1-GetDiscountPoolsByOtherAppReleaserResponse) |  |
| CreateAppCouponSetting | [CreateAppCouponSettingRequest](#cloud-hashing-inspire-v1-CreateAppCouponSettingRequest) | [CreateAppCouponSettingResponse](#cloud-hashing-inspire-v1-CreateAppCouponSettingResponse) |  |
| GetAppCouponSetting | [GetAppCouponSettingRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingRequest) | [GetAppCouponSettingResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingResponse) |  |
| GetAppCouponSettingByApp | [GetAppCouponSettingByAppRequest](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppRequest) | [GetAppCouponSettingByAppResponse](#cloud-hashing-inspire-v1-GetAppCouponSettingByAppResponse) |  |
| UpdateAppCouponSetting | [UpdateAppCouponSettingRequest](#cloud-hashing-inspire-v1-UpdateAppCouponSettingRequest) | [UpdateAppCouponSettingResponse](#cloud-hashing-inspire-v1-UpdateAppCouponSettingResponse) |  |
| CreateDefaultKpiSetting | [CreateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingRequest) | [CreateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-CreateDefaultKpiSettingResponse) |  |
| GetDefaultKpiSetting | [GetDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingRequest) | [GetDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingResponse) |  |
| GetDefaultKpiSettingByAppGood | [GetDefaultKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodRequest) | [GetDefaultKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetDefaultKpiSettingByAppGoodResponse) |  |
| UpdateDefaultKpiSetting | [UpdateDefaultKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingRequest) | [UpdateDefaultKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateDefaultKpiSettingResponse) |  |
| CreateUserKpiSetting | [CreateUserKpiSettingRequest](#cloud-hashing-inspire-v1-CreateUserKpiSettingRequest) | [CreateUserKpiSettingResponse](#cloud-hashing-inspire-v1-CreateUserKpiSettingResponse) |  |
| GetUserKpiSetting | [GetUserKpiSettingRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingRequest) | [GetUserKpiSettingResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingResponse) |  |
| GetUserKpiSettingByAppGood | [GetUserKpiSettingByAppGoodRequest](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodRequest) | [GetUserKpiSettingByAppGoodResponse](#cloud-hashing-inspire-v1-GetUserKpiSettingByAppGoodResponse) |  |
| UpdateUserKpiSetting | [UpdateUserKpiSettingRequest](#cloud-hashing-inspire-v1-UpdateUserKpiSettingRequest) | [UpdateUserKpiSettingResponse](#cloud-hashing-inspire-v1-UpdateUserKpiSettingResponse) |  |
| CreateUserSpecialReduction | [CreateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionRequest) | [CreateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionResponse) |  |
| CreateUserSpecialReductionForAppOtherUser | [CreateUserSpecialReductionForAppOtherUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserRequest) | [CreateUserSpecialReductionForAppOtherUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForAppOtherUserResponse) |  |
| CreateUserSpecialReductionForOtherAppUser | [CreateUserSpecialReductionForOtherAppUserRequest](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserRequest) | [CreateUserSpecialReductionForOtherAppUserResponse](#cloud-hashing-inspire-v1-CreateUserSpecialReductionForOtherAppUserResponse) |  |
| GetUserSpecialReduction | [GetUserSpecialReductionRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionRequest) | [GetUserSpecialReductionResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionResponse) |  |
| GetUserSpecialReductionsByApp | [GetUserSpecialReductionsByAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppRequest) | [GetUserSpecialReductionsByAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppResponse) |  |
| GetUserSpecialReductionsByOtherApp | [GetUserSpecialReductionsByOtherAppRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppRequest) | [GetUserSpecialReductionsByOtherAppResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByOtherAppResponse) |  |
| GetUserSpecialReductionsByAppReleaser | [GetUserSpecialReductionsByAppReleaserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserRequest) | [GetUserSpecialReductionsByAppReleaserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppReleaserResponse) |  |
| GetUserSpecialReductionsByAppUser | [GetUserSpecialReductionsByAppUserRequest](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserRequest) | [GetUserSpecialReductionsByAppUserResponse](#cloud-hashing-inspire-v1-GetUserSpecialReductionsByAppUserResponse) |  |
| UpdateUserSpecialReduction | [UpdateUserSpecialReductionRequest](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionRequest) | [UpdateUserSpecialReductionResponse](#cloud-hashing-inspire-v1-UpdateUserSpecialReductionResponse) |  |
| CreateActivity | [CreateActivityRequest](#cloud-hashing-inspire-v1-CreateActivityRequest) | [CreateActivityResponse](#cloud-hashing-inspire-v1-CreateActivityResponse) |  |
| CreateActivityForOtherApp | [CreateActivityForOtherAppRequest](#cloud-hashing-inspire-v1-CreateActivityForOtherAppRequest) | [CreateActivityForOtherAppResponse](#cloud-hashing-inspire-v1-CreateActivityForOtherAppResponse) |  |
| UpdateActivity | [UpdateActivityRequest](#cloud-hashing-inspire-v1-UpdateActivityRequest) | [UpdateActivityResponse](#cloud-hashing-inspire-v1-UpdateActivityResponse) |  |
| GetActivity | [GetActivityRequest](#cloud-hashing-inspire-v1-GetActivityRequest) | [GetActivityResponse](#cloud-hashing-inspire-v1-GetActivityResponse) |  |
| GetActivityByAppName | [GetActivityByAppNameRequest](#cloud-hashing-inspire-v1-GetActivityByAppNameRequest) | [GetActivityByAppNameResponse](#cloud-hashing-inspire-v1-GetActivityByAppNameResponse) |  |
| GetActivitiesByApp | [GetActivitiesByAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByAppRequest) | [GetActivitiesByAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByAppResponse) |  |
| GetActivitiesByOtherApp | [GetActivitiesByOtherAppRequest](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppRequest) | [GetActivitiesByOtherAppResponse](#cloud-hashing-inspire-v1-GetActivitiesByOtherAppResponse) |  |
| CreateEventCoupon | [CreateEventCouponRequest](#cloud-hashing-inspire-v1-CreateEventCouponRequest) | [CreateEventCouponResponse](#cloud-hashing-inspire-v1-CreateEventCouponResponse) |  |
| CreateEventCouponForOtherApp | [CreateEventCouponForOtherAppRequest](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppRequest) | [CreateEventCouponForOtherAppResponse](#cloud-hashing-inspire-v1-CreateEventCouponForOtherAppResponse) |  |
| UpdateEventCoupon | [UpdateEventCouponRequest](#cloud-hashing-inspire-v1-UpdateEventCouponRequest) | [UpdateEventCouponResponse](#cloud-hashing-inspire-v1-UpdateEventCouponResponse) |  |
| GetEventCoupon | [GetEventCouponRequest](#cloud-hashing-inspire-v1-GetEventCouponRequest) | [GetEventCouponResponse](#cloud-hashing-inspire-v1-GetEventCouponResponse) |  |
| GetEventCouponsByAppActivityEvent | [GetEventCouponsByAppActivityEventRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventRequest) | [GetEventCouponsByAppActivityEventResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppActivityEventResponse) |  |
| GetEventCouponsByApp | [GetEventCouponsByAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByAppRequest) | [GetEventCouponsByAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByAppResponse) |  |
| GetEventCouponsByOtherApp | [GetEventCouponsByOtherAppRequest](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppRequest) | [GetEventCouponsByOtherAppResponse](#cloud-hashing-inspire-v1-GetEventCouponsByOtherAppResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

