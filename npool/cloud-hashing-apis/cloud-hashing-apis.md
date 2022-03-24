# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-apis/cloud-hashing-apis.proto](#npool_cloud-hashing-apis_cloud-hashing-apis-proto)
    - [CoinSummary](#cloud-hashing-apis-v1-CoinSummary)
    - [Commission](#cloud-hashing-apis-v1-Commission)
    - [Coupon](#cloud-hashing-apis-v1-Coupon)
    - [Coupons](#cloud-hashing-apis-v1-Coupons)
    - [CreateGoodRequest](#cloud-hashing-apis-v1-CreateGoodRequest)
    - [CreateGoodResponse](#cloud-hashing-apis-v1-CreateGoodResponse)
    - [CreateKycRequest](#cloud-hashing-apis-v1-CreateKycRequest)
    - [CreateKycResponse](#cloud-hashing-apis-v1-CreateKycResponse)
    - [CreateOrderPaymentRequest](#cloud-hashing-apis-v1-CreateOrderPaymentRequest)
    - [CreateOrderPaymentResponse](#cloud-hashing-apis-v1-CreateOrderPaymentResponse)
    - [CreatePlatformCoinAccountRequest](#cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest)
    - [CreatePlatformCoinAccountResponse](#cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse)
    - [CreateUserCoinAccountRequest](#cloud-hashing-apis-v1-CreateUserCoinAccountRequest)
    - [CreateUserCoinAccountResponse](#cloud-hashing-apis-v1-CreateUserCoinAccountResponse)
    - [GetCommissionByAppUserRequest](#cloud-hashing-apis-v1-GetCommissionByAppUserRequest)
    - [GetCommissionByAppUserResponse](#cloud-hashing-apis-v1-GetCommissionByAppUserResponse)
    - [GetCouponsByAppUserRequest](#cloud-hashing-apis-v1-GetCouponsByAppUserRequest)
    - [GetCouponsByAppUserResponse](#cloud-hashing-apis-v1-GetCouponsByAppUserResponse)
    - [GetGoodRequest](#cloud-hashing-apis-v1-GetGoodRequest)
    - [GetGoodResponse](#cloud-hashing-apis-v1-GetGoodResponse)
    - [GetGoodReviewsRequest](#cloud-hashing-apis-v1-GetGoodReviewsRequest)
    - [GetGoodReviewsResponse](#cloud-hashing-apis-v1-GetGoodReviewsResponse)
    - [GetGoodsByAppRequest](#cloud-hashing-apis-v1-GetGoodsByAppRequest)
    - [GetGoodsByAppResponse](#cloud-hashing-apis-v1-GetGoodsByAppResponse)
    - [GetGoodsRequest](#cloud-hashing-apis-v1-GetGoodsRequest)
    - [GetGoodsResponse](#cloud-hashing-apis-v1-GetGoodsResponse)
    - [GetKycByAppUserRequest](#cloud-hashing-apis-v1-GetKycByAppUserRequest)
    - [GetKycByAppUserResponse](#cloud-hashing-apis-v1-GetKycByAppUserResponse)
    - [GetKycReviewsByAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByAppRequest)
    - [GetKycReviewsByAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByAppResponse)
    - [GetKycReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest)
    - [GetKycReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse)
    - [GetKycReviewsRequest](#cloud-hashing-apis-v1-GetKycReviewsRequest)
    - [GetKycReviewsResponse](#cloud-hashing-apis-v1-GetKycReviewsResponse)
    - [GetLayeredReferralsRequest](#cloud-hashing-apis-v1-GetLayeredReferralsRequest)
    - [GetLayeredReferralsResponse](#cloud-hashing-apis-v1-GetLayeredReferralsResponse)
    - [GetOrderRequest](#cloud-hashing-apis-v1-GetOrderRequest)
    - [GetOrderResponse](#cloud-hashing-apis-v1-GetOrderResponse)
    - [GetOrdersByAppRequest](#cloud-hashing-apis-v1-GetOrdersByAppRequest)
    - [GetOrdersByAppResponse](#cloud-hashing-apis-v1-GetOrdersByAppResponse)
    - [GetOrdersByAppUserRequest](#cloud-hashing-apis-v1-GetOrdersByAppUserRequest)
    - [GetOrdersByAppUserResponse](#cloud-hashing-apis-v1-GetOrdersByAppUserResponse)
    - [GetOrdersByGoodRequest](#cloud-hashing-apis-v1-GetOrdersByGoodRequest)
    - [GetOrdersByGoodResponse](#cloud-hashing-apis-v1-GetOrdersByGoodResponse)
    - [GetRecommendGoodsByAppRequest](#cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest)
    - [GetRecommendGoodsByAppResponse](#cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse)
    - [GetReferralsRequest](#cloud-hashing-apis-v1-GetReferralsRequest)
    - [GetReferralsResponse](#cloud-hashing-apis-v1-GetReferralsResponse)
    - [GetUserWithdrawsByAppUserRequest](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest)
    - [GetUserWithdrawsByAppUserResponse](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse)
    - [GetWithdrawAddressReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest)
    - [GetWithdrawAddressReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse)
    - [GetWithdrawAddressReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest)
    - [GetWithdrawAddressReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse)
    - [GetWithdrawAddressReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest)
    - [GetWithdrawAddressReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse)
    - [GetWithdrawAddressesByAppUserRequest](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest)
    - [GetWithdrawAddressesByAppUserResponse](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse)
    - [GetWithdrawReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest)
    - [GetWithdrawReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse)
    - [GetWithdrawReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest)
    - [GetWithdrawReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse)
    - [GetWithdrawReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsRequest)
    - [GetWithdrawReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsResponse)
    - [Good](#cloud-hashing-apis-v1-Good)
    - [GoodReview](#cloud-hashing-apis-v1-GoodReview)
    - [Kyc](#cloud-hashing-apis-v1-Kyc)
    - [KycReview](#cloud-hashing-apis-v1-KycReview)
    - [Order](#cloud-hashing-apis-v1-Order)
    - [OrderFee](#cloud-hashing-apis-v1-OrderFee)
    - [RecommendGood](#cloud-hashing-apis-v1-RecommendGood)
    - [Referral](#cloud-hashing-apis-v1-Referral)
    - [SetWithdrawAddressRequest](#cloud-hashing-apis-v1-SetWithdrawAddressRequest)
    - [SetWithdrawAddressResponse](#cloud-hashing-apis-v1-SetWithdrawAddressResponse)
    - [SignupRequest](#cloud-hashing-apis-v1-SignupRequest)
    - [SignupResponse](#cloud-hashing-apis-v1-SignupResponse)
    - [SubmitOrderRequest](#cloud-hashing-apis-v1-SubmitOrderRequest)
    - [SubmitOrderResponse](#cloud-hashing-apis-v1-SubmitOrderResponse)
    - [SubmitUserWithdrawRequest](#cloud-hashing-apis-v1-SubmitUserWithdrawRequest)
    - [SubmitUserWithdrawResponse](#cloud-hashing-apis-v1-SubmitUserWithdrawResponse)
    - [UpdateAccountRequest](#cloud-hashing-apis-v1-UpdateAccountRequest)
    - [UpdateAccountResponse](#cloud-hashing-apis-v1-UpdateAccountResponse)
    - [UpdateEmailAddressRequest](#cloud-hashing-apis-v1-UpdateEmailAddressRequest)
    - [UpdateEmailAddressResponse](#cloud-hashing-apis-v1-UpdateEmailAddressResponse)
    - [UpdateKycRequest](#cloud-hashing-apis-v1-UpdateKycRequest)
    - [UpdateKycResponse](#cloud-hashing-apis-v1-UpdateKycResponse)
    - [UpdateKycReviewRequest](#cloud-hashing-apis-v1-UpdateKycReviewRequest)
    - [UpdateKycReviewResponse](#cloud-hashing-apis-v1-UpdateKycReviewResponse)
    - [UpdatePasswordByAppUserRequest](#cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest)
    - [UpdatePasswordByAppUserResponse](#cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse)
    - [UpdatePasswordRequest](#cloud-hashing-apis-v1-UpdatePasswordRequest)
    - [UpdatePasswordResponse](#cloud-hashing-apis-v1-UpdatePasswordResponse)
    - [UpdatePhoneNORequest](#cloud-hashing-apis-v1-UpdatePhoneNORequest)
    - [UpdatePhoneNOResponse](#cloud-hashing-apis-v1-UpdatePhoneNOResponse)
    - [UpdateUserWithdrawReviewForOtherAppUserRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest)
    - [UpdateUserWithdrawReviewForOtherAppUserResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse)
    - [UpdateUserWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest)
    - [UpdateUserWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse)
    - [UpdateWithdrawAddressReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest)
    - [UpdateWithdrawAddressReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse)
    - [UpdateWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawReviewRequest)
    - [UpdateWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawReviewResponse)
    - [UserSpecial](#cloud-hashing-apis-v1-UserSpecial)
    - [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw)
    - [VerificationCode](#cloud-hashing-apis-v1-VerificationCode)
    - [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress)
    - [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview)
    - [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview)
  
    - [CloudHashingApis](#cloud-hashing-apis-v1-CloudHashingApis)
  
- [npool/cloud-hashing-apis/cloud-hashing-apis.proto](#npool_cloud-hashing-apis_cloud-hashing-apis-proto)
    - [CoinSummary](#cloud-hashing-apis-v1-CoinSummary)
    - [Commission](#cloud-hashing-apis-v1-Commission)
    - [Coupon](#cloud-hashing-apis-v1-Coupon)
    - [Coupons](#cloud-hashing-apis-v1-Coupons)
    - [CreateGoodRequest](#cloud-hashing-apis-v1-CreateGoodRequest)
    - [CreateGoodResponse](#cloud-hashing-apis-v1-CreateGoodResponse)
    - [CreateKycRequest](#cloud-hashing-apis-v1-CreateKycRequest)
    - [CreateKycResponse](#cloud-hashing-apis-v1-CreateKycResponse)
    - [CreateOrderPaymentRequest](#cloud-hashing-apis-v1-CreateOrderPaymentRequest)
    - [CreateOrderPaymentResponse](#cloud-hashing-apis-v1-CreateOrderPaymentResponse)
    - [CreatePlatformCoinAccountRequest](#cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest)
    - [CreatePlatformCoinAccountResponse](#cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse)
    - [CreateUserCoinAccountRequest](#cloud-hashing-apis-v1-CreateUserCoinAccountRequest)
    - [CreateUserCoinAccountResponse](#cloud-hashing-apis-v1-CreateUserCoinAccountResponse)
    - [GetCommissionByAppUserRequest](#cloud-hashing-apis-v1-GetCommissionByAppUserRequest)
    - [GetCommissionByAppUserResponse](#cloud-hashing-apis-v1-GetCommissionByAppUserResponse)
    - [GetCouponsByAppUserRequest](#cloud-hashing-apis-v1-GetCouponsByAppUserRequest)
    - [GetCouponsByAppUserResponse](#cloud-hashing-apis-v1-GetCouponsByAppUserResponse)
    - [GetGoodRequest](#cloud-hashing-apis-v1-GetGoodRequest)
    - [GetGoodResponse](#cloud-hashing-apis-v1-GetGoodResponse)
    - [GetGoodReviewsRequest](#cloud-hashing-apis-v1-GetGoodReviewsRequest)
    - [GetGoodReviewsResponse](#cloud-hashing-apis-v1-GetGoodReviewsResponse)
    - [GetGoodsByAppRequest](#cloud-hashing-apis-v1-GetGoodsByAppRequest)
    - [GetGoodsByAppResponse](#cloud-hashing-apis-v1-GetGoodsByAppResponse)
    - [GetGoodsRequest](#cloud-hashing-apis-v1-GetGoodsRequest)
    - [GetGoodsResponse](#cloud-hashing-apis-v1-GetGoodsResponse)
    - [GetKycByAppUserRequest](#cloud-hashing-apis-v1-GetKycByAppUserRequest)
    - [GetKycByAppUserResponse](#cloud-hashing-apis-v1-GetKycByAppUserResponse)
    - [GetKycReviewsByAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByAppRequest)
    - [GetKycReviewsByAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByAppResponse)
    - [GetKycReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest)
    - [GetKycReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse)
    - [GetKycReviewsRequest](#cloud-hashing-apis-v1-GetKycReviewsRequest)
    - [GetKycReviewsResponse](#cloud-hashing-apis-v1-GetKycReviewsResponse)
    - [GetLayeredReferralsRequest](#cloud-hashing-apis-v1-GetLayeredReferralsRequest)
    - [GetLayeredReferralsResponse](#cloud-hashing-apis-v1-GetLayeredReferralsResponse)
    - [GetOrderRequest](#cloud-hashing-apis-v1-GetOrderRequest)
    - [GetOrderResponse](#cloud-hashing-apis-v1-GetOrderResponse)
    - [GetOrdersByAppRequest](#cloud-hashing-apis-v1-GetOrdersByAppRequest)
    - [GetOrdersByAppResponse](#cloud-hashing-apis-v1-GetOrdersByAppResponse)
    - [GetOrdersByAppUserRequest](#cloud-hashing-apis-v1-GetOrdersByAppUserRequest)
    - [GetOrdersByAppUserResponse](#cloud-hashing-apis-v1-GetOrdersByAppUserResponse)
    - [GetOrdersByGoodRequest](#cloud-hashing-apis-v1-GetOrdersByGoodRequest)
    - [GetOrdersByGoodResponse](#cloud-hashing-apis-v1-GetOrdersByGoodResponse)
    - [GetRecommendGoodsByAppRequest](#cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest)
    - [GetRecommendGoodsByAppResponse](#cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse)
    - [GetReferralsRequest](#cloud-hashing-apis-v1-GetReferralsRequest)
    - [GetReferralsResponse](#cloud-hashing-apis-v1-GetReferralsResponse)
    - [GetUserWithdrawsByAppUserRequest](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest)
    - [GetUserWithdrawsByAppUserResponse](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse)
    - [GetWithdrawAddressReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest)
    - [GetWithdrawAddressReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse)
    - [GetWithdrawAddressReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest)
    - [GetWithdrawAddressReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse)
    - [GetWithdrawAddressReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest)
    - [GetWithdrawAddressReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse)
    - [GetWithdrawAddressesByAppUserRequest](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest)
    - [GetWithdrawAddressesByAppUserResponse](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse)
    - [GetWithdrawReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest)
    - [GetWithdrawReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse)
    - [GetWithdrawReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest)
    - [GetWithdrawReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse)
    - [GetWithdrawReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsRequest)
    - [GetWithdrawReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsResponse)
    - [Good](#cloud-hashing-apis-v1-Good)
    - [GoodReview](#cloud-hashing-apis-v1-GoodReview)
    - [Kyc](#cloud-hashing-apis-v1-Kyc)
    - [KycReview](#cloud-hashing-apis-v1-KycReview)
    - [Order](#cloud-hashing-apis-v1-Order)
    - [OrderFee](#cloud-hashing-apis-v1-OrderFee)
    - [RecommendGood](#cloud-hashing-apis-v1-RecommendGood)
    - [Referral](#cloud-hashing-apis-v1-Referral)
    - [SetWithdrawAddressRequest](#cloud-hashing-apis-v1-SetWithdrawAddressRequest)
    - [SetWithdrawAddressResponse](#cloud-hashing-apis-v1-SetWithdrawAddressResponse)
    - [SignupRequest](#cloud-hashing-apis-v1-SignupRequest)
    - [SignupResponse](#cloud-hashing-apis-v1-SignupResponse)
    - [SubmitOrderRequest](#cloud-hashing-apis-v1-SubmitOrderRequest)
    - [SubmitOrderResponse](#cloud-hashing-apis-v1-SubmitOrderResponse)
    - [SubmitUserWithdrawRequest](#cloud-hashing-apis-v1-SubmitUserWithdrawRequest)
    - [SubmitUserWithdrawResponse](#cloud-hashing-apis-v1-SubmitUserWithdrawResponse)
    - [UpdateAccountRequest](#cloud-hashing-apis-v1-UpdateAccountRequest)
    - [UpdateAccountResponse](#cloud-hashing-apis-v1-UpdateAccountResponse)
    - [UpdateEmailAddressRequest](#cloud-hashing-apis-v1-UpdateEmailAddressRequest)
    - [UpdateEmailAddressResponse](#cloud-hashing-apis-v1-UpdateEmailAddressResponse)
    - [UpdateKycRequest](#cloud-hashing-apis-v1-UpdateKycRequest)
    - [UpdateKycResponse](#cloud-hashing-apis-v1-UpdateKycResponse)
    - [UpdateKycReviewRequest](#cloud-hashing-apis-v1-UpdateKycReviewRequest)
    - [UpdateKycReviewResponse](#cloud-hashing-apis-v1-UpdateKycReviewResponse)
    - [UpdatePasswordByAppUserRequest](#cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest)
    - [UpdatePasswordByAppUserResponse](#cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse)
    - [UpdatePasswordRequest](#cloud-hashing-apis-v1-UpdatePasswordRequest)
    - [UpdatePasswordResponse](#cloud-hashing-apis-v1-UpdatePasswordResponse)
    - [UpdatePhoneNORequest](#cloud-hashing-apis-v1-UpdatePhoneNORequest)
    - [UpdatePhoneNOResponse](#cloud-hashing-apis-v1-UpdatePhoneNOResponse)
    - [UpdateUserWithdrawReviewForOtherAppUserRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest)
    - [UpdateUserWithdrawReviewForOtherAppUserResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse)
    - [UpdateUserWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest)
    - [UpdateUserWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse)
    - [UpdateWithdrawAddressReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest)
    - [UpdateWithdrawAddressReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse)
    - [UpdateWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawReviewRequest)
    - [UpdateWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawReviewResponse)
    - [UserSpecial](#cloud-hashing-apis-v1-UserSpecial)
    - [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw)
    - [VerificationCode](#cloud-hashing-apis-v1-VerificationCode)
    - [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress)
    - [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview)
    - [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview)
  
    - [CloudHashingApis](#cloud-hashing-apis-v1-CloudHashingApis)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_cloud-hashing-apis_cloud-hashing-apis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-apis/cloud-hashing-apis.proto



<a name="cloud-hashing-apis-v1-CoinSummary"></a>

### CoinSummary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| CoinName | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| Amount | [double](#double) |  |  |
| Unit | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-Commission"></a>

### Commission



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [double](#double) |  |  |
| Balance | [double](#double) |  |  |






<a name="cloud-hashing-apis-v1-Coupon"></a>

### Coupon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| Order | [cloud.hashing.order.v1.Order](#cloud-hashing-order-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-Coupons"></a>

### Coupons



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupons | [Coupon](#cloud-hashing-apis-v1-Coupon) | repeated |  |
| Offers | [UserSpecial](#cloud-hashing-apis-v1-UserSpecial) | repeated |  |






<a name="cloud-hashing-apis-v1-CreateGoodRequest"></a>

### CreateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.goods.v1.GoodInfo](#cloud-hashing-goods-v1-GoodInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateGoodResponse"></a>

### CreateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Good](#cloud-hashing-apis-v1-Good) |  |  |






<a name="cloud-hashing-apis-v1-CreateKycRequest"></a>

### CreateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateKycResponse"></a>

### CreateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-CreateOrderPaymentRequest"></a>

### CreateOrderPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| Fees | [OrderFee](#cloud-hashing-apis-v1-OrderFee) | repeated |  |






<a name="cloud-hashing-apis-v1-CreateOrderPaymentResponse"></a>

### CreateOrderPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest"></a>

### CreatePlatformCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse"></a>

### CreatePlatformCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateUserCoinAccountRequest"></a>

### CreateUserCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateUserCoinAccountResponse"></a>

### CreateUserCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetCommissionByAppUserRequest"></a>

### GetCommissionByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetCommissionByAppUserResponse"></a>

### GetCommissionByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Commission](#cloud-hashing-apis-v1-Commission) |  |  |






<a name="cloud-hashing-apis-v1-GetCouponsByAppUserRequest"></a>

### GetCouponsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetCouponsByAppUserResponse"></a>

### GetCouponsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Coupons](#cloud-hashing-apis-v1-Coupons) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodRequest"></a>

### GetGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodResponse"></a>

### GetGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Good](#cloud-hashing-apis-v1-Good) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodReviewsRequest"></a>

### GetGoodReviewsRequest







<a name="cloud-hashing-apis-v1-GetGoodReviewsResponse"></a>

### GetGoodReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodReview](#cloud-hashing-apis-v1-GoodReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetGoodsByAppRequest"></a>

### GetGoodsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool-v1-PageInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsByAppResponse"></a>

### GetGoodsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Good](#cloud-hashing-apis-v1-Good) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsRequest"></a>

### GetGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool-v1-PageInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Good](#cloud-hashing-apis-v1-Good) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud-hashing-apis-v1-GetKycByAppUserRequest"></a>

### GetKycByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycByAppUserResponse"></a>

### GetKycByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByAppRequest"></a>

### GetKycReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByAppResponse"></a>

### GetKycReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest"></a>

### GetKycReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse"></a>

### GetKycReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsRequest"></a>

### GetKycReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsResponse"></a>

### GetKycReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetLayeredReferralsRequest"></a>

### GetLayeredReferralsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetLayeredReferralsResponse"></a>

### GetLayeredReferralsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Referral](#cloud-hashing-apis-v1-Referral) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppRequest"></a>

### GetOrdersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppResponse"></a>

### GetOrdersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppUserRequest"></a>

### GetOrdersByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppUserResponse"></a>

### GetOrdersByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrdersByGoodRequest"></a>

### GetOrdersByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByGoodResponse"></a>

### GetOrdersByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest"></a>

### GetRecommendGoodsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse"></a>

### GetRecommendGoodsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RecommendGood](#cloud-hashing-apis-v1-RecommendGood) | repeated |  |






<a name="cloud-hashing-apis-v1-GetReferralsRequest"></a>

### GetReferralsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetReferralsResponse"></a>

### GetReferralsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Referral](#cloud-hashing-apis-v1-Referral) | repeated |  |






<a name="cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest"></a>

### GetUserWithdrawsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse"></a>

### GetUserWithdrawsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest"></a>

### GetWithdrawAddressReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse"></a>

### GetWithdrawAddressReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest"></a>

### GetWithdrawAddressReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse"></a>

### GetWithdrawAddressReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest"></a>

### GetWithdrawAddressReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse"></a>

### GetWithdrawAddressReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest"></a>

### GetWithdrawAddressesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse"></a>

### GetWithdrawAddressesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest"></a>

### GetWithdrawReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse"></a>

### GetWithdrawReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest"></a>

### GetWithdrawReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse"></a>

### GetWithdrawReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsRequest"></a>

### GetWithdrawReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsResponse"></a>

### GetWithdrawReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-Good"></a>

### Good
Request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Good | [cloud.hashing.goods.v1.GoodDetail](#cloud-hashing-goods-v1-GoodDetail) |  |  |
| Main | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |
| SupportCoins | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) | repeated |  |
| Reviews | [review.service.v1.Review](#review-service-v1-Review) | repeated |  |
| Sold | [uint32](#uint32) |  |  |






<a name="cloud-hashing-apis-v1-GoodReview"></a>

### GoodReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| CreatedBy | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Good | [cloud.hashing.goods.v1.GoodDetail](#cloud-hashing-goods-v1-GoodDetail) |  |  |






<a name="cloud-hashing-apis-v1-Kyc"></a>

### Kyc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Kyc | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-KycReview"></a>

### KycReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Kyc | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Order | [cloud.hashing.order.v1.OrderDetail](#cloud-hashing-order-v1-OrderDetail) |  |  |
| PayToAccount | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |
| Good | [Good](#cloud-hashing-apis-v1-Good) |  |  |
| PayWithCoin | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |
| FixAmountCoupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| DiscountCoupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| UserSpecialReduction | [cloud.hashing.inspire.v1.UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |
| PaymentDeadline | [uint32](#uint32) |  |  |
| AppGood | [cloud.hashing.goods.v1.AppGoodInfo](#cloud-hashing-goods-v1-AppGoodInfo) |  |  |
| Promotion | [cloud.hashing.goods.v1.AppGoodPromotion](#cloud-hashing-goods-v1-AppGoodPromotion) |  |  |






<a name="cloud-hashing-apis-v1-OrderFee"></a>

### OrderFee



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DurationDays | [uint32](#uint32) |  |  |






<a name="cloud-hashing-apis-v1-RecommendGood"></a>

### RecommendGood



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Good | [Good](#cloud-hashing-apis-v1-Good) |  |  |
| Recommend | [cloud.hashing.goods.v1.Recommend](#cloud-hashing-goods-v1-Recommend) |  |  |






<a name="cloud-hashing-apis-v1-Referral"></a>

### Referral



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [app.user.manager.v1.AppUser](#app-user-manager-v1-AppUser) |  |  |
| Extra | [app.user.manager.v1.AppUserExtra](#app-user-manager-v1-AppUserExtra) |  |  |
| Invitation | [cloud.hashing.inspire.v1.RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |
| USDAmount | [double](#double) |  |  |
| SubUSDAmount | [double](#double) |  |  |
| Kol | [bool](#bool) |  |  |
| InvitedCount | [uint32](#uint32) |  |  |
| Summaries | [CoinSummary](#cloud-hashing-apis-v1-CoinSummary) | repeated |  |






<a name="cloud-hashing-apis-v1-SetWithdrawAddressRequest"></a>

### SetWithdrawAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| Labels | [string](#string) | repeated |  |






<a name="cloud-hashing-apis-v1-SetWithdrawAddressResponse"></a>

### SetWithdrawAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress) |  |  |






<a name="cloud-hashing-apis-v1-SignupRequest"></a>

### SignupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InvitationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SignupResponse"></a>

### SignupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUser](#app-user-manager-v1-AppUser) |  |  |






<a name="cloud-hashing-apis-v1-SubmitOrderRequest"></a>

### SubmitOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SubmitOrderResponse"></a>

### SubmitOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-SubmitUserWithdrawRequest"></a>

### SubmitUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SubmitUserWithdrawResponse"></a>

### SubmitUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw) |  |  |






<a name="cloud-hashing-apis-v1-UpdateAccountRequest"></a>

### UpdateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCodes | [VerificationCode](#cloud-hashing-apis-v1-VerificationCode) | repeated |  |






<a name="cloud-hashing-apis-v1-UpdateAccountResponse"></a>

### UpdateAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateEmailAddressRequest"></a>

### UpdateEmailAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldVerificationCode | [string](#string) |  |  |
| OldAccount | [string](#string) |  |  |
| OldAccountType | [string](#string) |  |  |
| NewEmailAddress | [string](#string) |  |  |
| NewEmailVerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateEmailAddressResponse"></a>

### UpdateEmailAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycRequest"></a>

### UpdateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycResponse"></a>

### UpdateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycReviewRequest"></a>

### UpdateKycReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycReviewResponse"></a>

### UpdateKycReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycReview](#cloud-hashing-apis-v1-KycReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest"></a>

### UpdatePasswordByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| OldPasswordHash | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse"></a>

### UpdatePasswordByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserSecret](#app-user-manager-v1-AppUserSecret) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordRequest"></a>

### UpdatePasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordResponse"></a>

### UpdatePasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserSecret](#app-user-manager-v1-AppUserSecret) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePhoneNORequest"></a>

### UpdatePhoneNORequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldVerificationCode | [string](#string) |  |  |
| OldAccount | [string](#string) |  |  |
| OldAccountType | [string](#string) |  |  |
| NewPhoneNO | [string](#string) |  |  |
| NewPhoneVerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePhoneNOResponse"></a>

### UpdatePhoneNOResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest"></a>

### UpdateUserWithdrawReviewForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  | Reviewer |
| UserID | [string](#string) |  |  |
| TargetAppID | [string](#string) |  | Target |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse"></a>

### UpdateUserWithdrawReviewForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest"></a>

### UpdateUserWithdrawReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  | Reviewer |
| UserID | [string](#string) |  |  |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse"></a>

### UpdateUserWithdrawReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest"></a>

### UpdateWithdrawAddressReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse"></a>

### UpdateWithdrawAddressReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawReviewRequest"></a>

### UpdateWithdrawReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawReviewResponse"></a>

### UpdateWithdrawReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UserSpecial"></a>

### UserSpecial



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupon | [cloud.hashing.inspire.v1.UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |
| Order | [cloud.hashing.order.v1.Order](#cloud-hashing-order-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-UserWithdraw"></a>

### UserWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Withdraw | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-VerificationCode"></a>

### VerificationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| VerificationCode | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawAddress"></a>

### WithdrawAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [cloud.hashing.billing.v1.UserWithdraw](#cloud-hashing-billing-v1-UserWithdraw) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| Account | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawAddressReview"></a>

### WithdrawAddressReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Address | [cloud.hashing.billing.v1.UserWithdraw](#cloud-hashing-billing-v1-UserWithdraw) |  |  |
| Account | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawReview"></a>

### WithdrawReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Withdraw | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |





 

 

 


<a name="cloud-hashing-apis-v1-CloudHashingApis"></a>

### CloudHashingApis
Cloud Hashing Goods

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| GetGoods | [GetGoodsRequest](#cloud-hashing-apis-v1-GetGoodsRequest) | [GetGoodsResponse](#cloud-hashing-apis-v1-GetGoodsResponse) |  |
| GetGoodsByApp | [GetGoodsByAppRequest](#cloud-hashing-apis-v1-GetGoodsByAppRequest) | [GetGoodsByAppResponse](#cloud-hashing-apis-v1-GetGoodsByAppResponse) |  |
| CreateGood | [CreateGoodRequest](#cloud-hashing-apis-v1-CreateGoodRequest) | [CreateGoodResponse](#cloud-hashing-apis-v1-CreateGoodResponse) |  |
| GetGood | [GetGoodRequest](#cloud-hashing-apis-v1-GetGoodRequest) | [GetGoodResponse](#cloud-hashing-apis-v1-GetGoodResponse) |  |
| GetRecommendGoodsByApp | [GetRecommendGoodsByAppRequest](#cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest) | [GetRecommendGoodsByAppResponse](#cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse) |  |
| SubmitOrder | [SubmitOrderRequest](#cloud-hashing-apis-v1-SubmitOrderRequest) | [SubmitOrderResponse](#cloud-hashing-apis-v1-SubmitOrderResponse) |  |
| CreateOrderPayment | [CreateOrderPaymentRequest](#cloud-hashing-apis-v1-CreateOrderPaymentRequest) | [CreateOrderPaymentResponse](#cloud-hashing-apis-v1-CreateOrderPaymentResponse) |  |
| GetOrder | [GetOrderRequest](#cloud-hashing-apis-v1-GetOrderRequest) | [GetOrderResponse](#cloud-hashing-apis-v1-GetOrderResponse) |  |
| GetOrdersByAppUser | [GetOrdersByAppUserRequest](#cloud-hashing-apis-v1-GetOrdersByAppUserRequest) | [GetOrdersByAppUserResponse](#cloud-hashing-apis-v1-GetOrdersByAppUserResponse) |  |
| GetOrdersByApp | [GetOrdersByAppRequest](#cloud-hashing-apis-v1-GetOrdersByAppRequest) | [GetOrdersByAppResponse](#cloud-hashing-apis-v1-GetOrdersByAppResponse) |  |
| GetOrdersByGood | [GetOrdersByGoodRequest](#cloud-hashing-apis-v1-GetOrdersByGoodRequest) | [GetOrdersByGoodResponse](#cloud-hashing-apis-v1-GetOrdersByGoodResponse) |  |
| Signup | [SignupRequest](#cloud-hashing-apis-v1-SignupRequest) | [SignupResponse](#cloud-hashing-apis-v1-SignupResponse) |  |
| UpdatePassword | [UpdatePasswordRequest](#cloud-hashing-apis-v1-UpdatePasswordRequest) | [UpdatePasswordResponse](#cloud-hashing-apis-v1-UpdatePasswordResponse) |  |
| UpdatePasswordByAppUser | [UpdatePasswordByAppUserRequest](#cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest) | [UpdatePasswordByAppUserResponse](#cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse) |  |
| UpdateEmailAddress | [UpdateEmailAddressRequest](#cloud-hashing-apis-v1-UpdateEmailAddressRequest) | [UpdateEmailAddressResponse](#cloud-hashing-apis-v1-UpdateEmailAddressResponse) |  |
| UpdatePhoneNO | [UpdatePhoneNORequest](#cloud-hashing-apis-v1-UpdatePhoneNORequest) | [UpdatePhoneNOResponse](#cloud-hashing-apis-v1-UpdatePhoneNOResponse) |  |
| UpdateAccount | [UpdateAccountRequest](#cloud-hashing-apis-v1-UpdateAccountRequest) | [UpdateAccountResponse](#cloud-hashing-apis-v1-UpdateAccountResponse) |  |
| GetReferrals | [GetReferralsRequest](#cloud-hashing-apis-v1-GetReferralsRequest) | [GetReferralsResponse](#cloud-hashing-apis-v1-GetReferralsResponse) |  |
| GetLayeredReferrals | [GetLayeredReferralsRequest](#cloud-hashing-apis-v1-GetLayeredReferralsRequest) | [GetLayeredReferralsResponse](#cloud-hashing-apis-v1-GetLayeredReferralsResponse) |  |
| GetKycReviews | [GetKycReviewsRequest](#cloud-hashing-apis-v1-GetKycReviewsRequest) | [GetKycReviewsResponse](#cloud-hashing-apis-v1-GetKycReviewsResponse) |  |
| GetKycReviewsByApp | [GetKycReviewsByAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByAppRequest) | [GetKycReviewsByAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByAppResponse) |  |
| GetKycReviewsByOtherApp | [GetKycReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest) | [GetKycReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse) |  |
| GetGoodReviews | [GetGoodReviewsRequest](#cloud-hashing-apis-v1-GetGoodReviewsRequest) | [GetGoodReviewsResponse](#cloud-hashing-apis-v1-GetGoodReviewsResponse) |  |
| GetWithdrawReviews | [GetWithdrawReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsRequest) | [GetWithdrawReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsResponse) |  |
| GetWithdrawReviewsByApp | [GetWithdrawReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest) | [GetWithdrawReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse) |  |
| GetWithdrawReviewsByOtherApp | [GetWithdrawReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest) | [GetWithdrawReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse) |  |
| GetWithdrawAddressReviews | [GetWithdrawAddressReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest) | [GetWithdrawAddressReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse) |  |
| GetWithdrawAddressReviewsByApp | [GetWithdrawAddressReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest) | [GetWithdrawAddressReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse) |  |
| GetWithdrawAddressReviewsByOtherApp | [GetWithdrawAddressReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest) | [GetWithdrawAddressReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse) |  |
| CreateKyc | [CreateKycRequest](#cloud-hashing-apis-v1-CreateKycRequest) | [CreateKycResponse](#cloud-hashing-apis-v1-CreateKycResponse) |  |
| UpdateKyc | [UpdateKycRequest](#cloud-hashing-apis-v1-UpdateKycRequest) | [UpdateKycResponse](#cloud-hashing-apis-v1-UpdateKycResponse) |  |
| GetKycByAppUser | [GetKycByAppUserRequest](#cloud-hashing-apis-v1-GetKycByAppUserRequest) | [GetKycByAppUserResponse](#cloud-hashing-apis-v1-GetKycByAppUserResponse) |  |
| CreatePlatformCoinAccount | [CreatePlatformCoinAccountRequest](#cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest) | [CreatePlatformCoinAccountResponse](#cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse) |  |
| CreateUserCoinAccount | [CreateUserCoinAccountRequest](#cloud-hashing-apis-v1-CreateUserCoinAccountRequest) | [CreateUserCoinAccountResponse](#cloud-hashing-apis-v1-CreateUserCoinAccountResponse) |  |
| SubmitUserWithdraw | [SubmitUserWithdrawRequest](#cloud-hashing-apis-v1-SubmitUserWithdrawRequest) | [SubmitUserWithdrawResponse](#cloud-hashing-apis-v1-SubmitUserWithdrawResponse) |  |
| UpdateUserWithdrawReview | [UpdateUserWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest) | [UpdateUserWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse) |  |
| UpdateUserWithdrawReviewForOtherAppUser | [UpdateUserWithdrawReviewForOtherAppUserRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest) | [UpdateUserWithdrawReviewForOtherAppUserResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse) |  |
| GetUserWithdrawsByAppUser | [GetUserWithdrawsByAppUserRequest](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest) | [GetUserWithdrawsByAppUserResponse](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse) |  |
| SetWithdrawAddress | [SetWithdrawAddressRequest](#cloud-hashing-apis-v1-SetWithdrawAddressRequest) | [SetWithdrawAddressResponse](#cloud-hashing-apis-v1-SetWithdrawAddressResponse) |  |
| GetWithdrawAddressesByAppUser | [GetWithdrawAddressesByAppUserRequest](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest) | [GetWithdrawAddressesByAppUserResponse](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse) |  |
| GetCouponsByAppUser | [GetCouponsByAppUserRequest](#cloud-hashing-apis-v1-GetCouponsByAppUserRequest) | [GetCouponsByAppUserResponse](#cloud-hashing-apis-v1-GetCouponsByAppUserResponse) |  |
| GetCommissionByAppUser | [GetCommissionByAppUserRequest](#cloud-hashing-apis-v1-GetCommissionByAppUserRequest) | [GetCommissionByAppUserResponse](#cloud-hashing-apis-v1-GetCommissionByAppUserResponse) |  |
| UpdateKycReview | [UpdateKycReviewRequest](#cloud-hashing-apis-v1-UpdateKycReviewRequest) | [UpdateKycReviewResponse](#cloud-hashing-apis-v1-UpdateKycReviewResponse) |  |
| UpdateWithdrawReview | [UpdateWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawReviewRequest) | [UpdateWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawReviewResponse) |  |
| UpdateWithdrawAddressReview | [UpdateWithdrawAddressReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest) | [UpdateWithdrawAddressReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse) |  |

 



<a name="npool_cloud-hashing-apis_cloud-hashing-apis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-apis/cloud-hashing-apis.proto



<a name="cloud-hashing-apis-v1-CoinSummary"></a>

### CoinSummary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| CoinName | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| Amount | [double](#double) |  |  |
| Unit | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-Commission"></a>

### Commission



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [double](#double) |  |  |
| Balance | [double](#double) |  |  |






<a name="cloud-hashing-apis-v1-Coupon"></a>

### Coupon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| Order | [cloud.hashing.order.v1.Order](#cloud-hashing-order-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-Coupons"></a>

### Coupons



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupons | [Coupon](#cloud-hashing-apis-v1-Coupon) | repeated |  |
| Offers | [UserSpecial](#cloud-hashing-apis-v1-UserSpecial) | repeated |  |






<a name="cloud-hashing-apis-v1-CreateGoodRequest"></a>

### CreateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.goods.v1.GoodInfo](#cloud-hashing-goods-v1-GoodInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateGoodResponse"></a>

### CreateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Good](#cloud-hashing-apis-v1-Good) |  |  |






<a name="cloud-hashing-apis-v1-CreateKycRequest"></a>

### CreateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateKycResponse"></a>

### CreateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-CreateOrderPaymentRequest"></a>

### CreateOrderPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| Fees | [OrderFee](#cloud-hashing-apis-v1-OrderFee) | repeated |  |






<a name="cloud-hashing-apis-v1-CreateOrderPaymentResponse"></a>

### CreateOrderPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest"></a>

### CreatePlatformCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse"></a>

### CreatePlatformCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateUserCoinAccountRequest"></a>

### CreateUserCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-CreateUserCoinAccountResponse"></a>

### CreateUserCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetCommissionByAppUserRequest"></a>

### GetCommissionByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetCommissionByAppUserResponse"></a>

### GetCommissionByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Commission](#cloud-hashing-apis-v1-Commission) |  |  |






<a name="cloud-hashing-apis-v1-GetCouponsByAppUserRequest"></a>

### GetCouponsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetCouponsByAppUserResponse"></a>

### GetCouponsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Coupons](#cloud-hashing-apis-v1-Coupons) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodRequest"></a>

### GetGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodResponse"></a>

### GetGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Good](#cloud-hashing-apis-v1-Good) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodReviewsRequest"></a>

### GetGoodReviewsRequest







<a name="cloud-hashing-apis-v1-GetGoodReviewsResponse"></a>

### GetGoodReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodReview](#cloud-hashing-apis-v1-GoodReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetGoodsByAppRequest"></a>

### GetGoodsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool-v1-PageInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsByAppResponse"></a>

### GetGoodsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Good](#cloud-hashing-apis-v1-Good) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsRequest"></a>

### GetGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool-v1-PageInfo) |  |  |






<a name="cloud-hashing-apis-v1-GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Good](#cloud-hashing-apis-v1-Good) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud-hashing-apis-v1-GetKycByAppUserRequest"></a>

### GetKycByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycByAppUserResponse"></a>

### GetKycByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByAppRequest"></a>

### GetKycReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByAppResponse"></a>

### GetKycReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest"></a>

### GetKycReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse"></a>

### GetKycReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsRequest"></a>

### GetKycReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetKycReviewsResponse"></a>

### GetKycReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [KycReview](#cloud-hashing-apis-v1-KycReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetLayeredReferralsRequest"></a>

### GetLayeredReferralsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetLayeredReferralsResponse"></a>

### GetLayeredReferralsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Referral](#cloud-hashing-apis-v1-Referral) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppRequest"></a>

### GetOrdersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppResponse"></a>

### GetOrdersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppUserRequest"></a>

### GetOrdersByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByAppUserResponse"></a>

### GetOrdersByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetOrdersByGoodRequest"></a>

### GetOrdersByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetOrdersByGoodResponse"></a>

### GetOrdersByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud-hashing-apis-v1-Order) | repeated |  |






<a name="cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest"></a>

### GetRecommendGoodsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse"></a>

### GetRecommendGoodsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RecommendGood](#cloud-hashing-apis-v1-RecommendGood) | repeated |  |






<a name="cloud-hashing-apis-v1-GetReferralsRequest"></a>

### GetReferralsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetReferralsResponse"></a>

### GetReferralsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Referral](#cloud-hashing-apis-v1-Referral) | repeated |  |






<a name="cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest"></a>

### GetUserWithdrawsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse"></a>

### GetUserWithdrawsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest"></a>

### GetWithdrawAddressReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse"></a>

### GetWithdrawAddressReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest"></a>

### GetWithdrawAddressReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse"></a>

### GetWithdrawAddressReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest"></a>

### GetWithdrawAddressReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse"></a>

### GetWithdrawAddressReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest"></a>

### GetWithdrawAddressesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse"></a>

### GetWithdrawAddressesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest"></a>

### GetWithdrawReviewsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse"></a>

### GetWithdrawReviewsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest"></a>

### GetWithdrawReviewsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse"></a>

### GetWithdrawReviewsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsRequest"></a>

### GetWithdrawReviewsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-GetWithdrawReviewsResponse"></a>

### GetWithdrawReviewsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) | repeated |  |






<a name="cloud-hashing-apis-v1-Good"></a>

### Good
Request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Good | [cloud.hashing.goods.v1.GoodDetail](#cloud-hashing-goods-v1-GoodDetail) |  |  |
| Main | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |
| SupportCoins | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) | repeated |  |
| Reviews | [review.service.v1.Review](#review-service-v1-Review) | repeated |  |
| Sold | [uint32](#uint32) |  |  |






<a name="cloud-hashing-apis-v1-GoodReview"></a>

### GoodReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| CreatedBy | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Good | [cloud.hashing.goods.v1.GoodDetail](#cloud-hashing-goods-v1-GoodDetail) |  |  |






<a name="cloud-hashing-apis-v1-Kyc"></a>

### Kyc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Kyc | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-KycReview"></a>

### KycReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Kyc | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Order | [cloud.hashing.order.v1.OrderDetail](#cloud-hashing-order-v1-OrderDetail) |  |  |
| PayToAccount | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |
| Good | [Good](#cloud-hashing-apis-v1-Good) |  |  |
| PayWithCoin | [sphinx.coininfo.v1.CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |
| FixAmountCoupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| DiscountCoupon | [cloud.hashing.inspire.v1.CouponAllocatedDetail](#cloud-hashing-inspire-v1-CouponAllocatedDetail) |  |  |
| UserSpecialReduction | [cloud.hashing.inspire.v1.UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |
| PaymentDeadline | [uint32](#uint32) |  |  |
| AppGood | [cloud.hashing.goods.v1.AppGoodInfo](#cloud-hashing-goods-v1-AppGoodInfo) |  |  |
| Promotion | [cloud.hashing.goods.v1.AppGoodPromotion](#cloud-hashing-goods-v1-AppGoodPromotion) |  |  |






<a name="cloud-hashing-apis-v1-OrderFee"></a>

### OrderFee



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DurationDays | [uint32](#uint32) |  |  |






<a name="cloud-hashing-apis-v1-RecommendGood"></a>

### RecommendGood



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Good | [Good](#cloud-hashing-apis-v1-Good) |  |  |
| Recommend | [cloud.hashing.goods.v1.Recommend](#cloud-hashing-goods-v1-Recommend) |  |  |






<a name="cloud-hashing-apis-v1-Referral"></a>

### Referral



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [app.user.manager.v1.AppUser](#app-user-manager-v1-AppUser) |  |  |
| Extra | [app.user.manager.v1.AppUserExtra](#app-user-manager-v1-AppUserExtra) |  |  |
| Invitation | [cloud.hashing.inspire.v1.RegistrationInvitation](#cloud-hashing-inspire-v1-RegistrationInvitation) |  |  |
| USDAmount | [double](#double) |  |  |
| SubUSDAmount | [double](#double) |  |  |
| Kol | [bool](#bool) |  |  |
| InvitedCount | [uint32](#uint32) |  |  |
| Summaries | [CoinSummary](#cloud-hashing-apis-v1-CoinSummary) | repeated |  |






<a name="cloud-hashing-apis-v1-SetWithdrawAddressRequest"></a>

### SetWithdrawAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| Labels | [string](#string) | repeated |  |






<a name="cloud-hashing-apis-v1-SetWithdrawAddressResponse"></a>

### SetWithdrawAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawAddress](#cloud-hashing-apis-v1-WithdrawAddress) |  |  |






<a name="cloud-hashing-apis-v1-SignupRequest"></a>

### SignupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| InvitationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SignupResponse"></a>

### SignupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUser](#app-user-manager-v1-AppUser) |  |  |






<a name="cloud-hashing-apis-v1-SubmitOrderRequest"></a>

### SubmitOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CouponID | [string](#string) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SubmitOrderResponse"></a>

### SubmitOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud-hashing-apis-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-SubmitUserWithdrawRequest"></a>

### SubmitUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-SubmitUserWithdrawResponse"></a>

### SubmitUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud-hashing-apis-v1-UserWithdraw) |  |  |






<a name="cloud-hashing-apis-v1-UpdateAccountRequest"></a>

### UpdateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| VerificationCodes | [VerificationCode](#cloud-hashing-apis-v1-VerificationCode) | repeated |  |






<a name="cloud-hashing-apis-v1-UpdateAccountResponse"></a>

### UpdateAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateEmailAddressRequest"></a>

### UpdateEmailAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldVerificationCode | [string](#string) |  |  |
| OldAccount | [string](#string) |  |  |
| OldAccountType | [string](#string) |  |  |
| NewEmailAddress | [string](#string) |  |  |
| NewEmailVerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateEmailAddressResponse"></a>

### UpdateEmailAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycRequest"></a>

### UpdateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [kyc.management.v1.KycInfo](#kyc-management-v1-KycInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycResponse"></a>

### UpdateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Kyc](#cloud-hashing-apis-v1-Kyc) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycReviewRequest"></a>

### UpdateKycReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateKycReviewResponse"></a>

### UpdateKycReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [KycReview](#cloud-hashing-apis-v1-KycReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest"></a>

### UpdatePasswordByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| OldPasswordHash | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse"></a>

### UpdatePasswordByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserSecret](#app-user-manager-v1-AppUserSecret) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordRequest"></a>

### UpdatePasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| VerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePasswordResponse"></a>

### UpdatePasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserSecret](#app-user-manager-v1-AppUserSecret) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePhoneNORequest"></a>

### UpdatePhoneNORequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldVerificationCode | [string](#string) |  |  |
| OldAccount | [string](#string) |  |  |
| OldAccountType | [string](#string) |  |  |
| NewPhoneNO | [string](#string) |  |  |
| NewPhoneVerificationCode | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdatePhoneNOResponse"></a>

### UpdatePhoneNOResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest"></a>

### UpdateUserWithdrawReviewForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  | Reviewer |
| UserID | [string](#string) |  |  |
| TargetAppID | [string](#string) |  | Target |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse"></a>

### UpdateUserWithdrawReviewForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest"></a>

### UpdateUserWithdrawReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  | Reviewer |
| UserID | [string](#string) |  |  |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |






<a name="cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse"></a>

### UpdateUserWithdrawReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest"></a>

### UpdateWithdrawAddressReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse"></a>

### UpdateWithdrawAddressReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawAddressReview](#cloud-hashing-apis-v1-WithdrawAddressReview) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawReviewRequest"></a>

### UpdateWithdrawReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| LangID | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-UpdateWithdrawReviewResponse"></a>

### UpdateWithdrawReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WithdrawReview](#cloud-hashing-apis-v1-WithdrawReview) |  |  |






<a name="cloud-hashing-apis-v1-UserSpecial"></a>

### UserSpecial



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Coupon | [cloud.hashing.inspire.v1.UserSpecialReduction](#cloud-hashing-inspire-v1-UserSpecialReduction) |  |  |
| Order | [cloud.hashing.order.v1.Order](#cloud-hashing-order-v1-Order) |  |  |






<a name="cloud-hashing-apis-v1-UserWithdraw"></a>

### UserWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Withdraw | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-VerificationCode"></a>

### VerificationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| VerificationCode | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| AccountType | [string](#string) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawAddress"></a>

### WithdrawAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [cloud.hashing.billing.v1.UserWithdraw](#cloud-hashing-billing-v1-UserWithdraw) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| Account | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawAddressReview"></a>

### WithdrawAddressReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Address | [cloud.hashing.billing.v1.UserWithdraw](#cloud-hashing-billing-v1-UserWithdraw) |  |  |
| Account | [cloud.hashing.billing.v1.CoinAccountInfo](#cloud-hashing-billing-v1-CoinAccountInfo) |  |  |






<a name="cloud-hashing-apis-v1-WithdrawReview"></a>

### WithdrawReview



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Review | [review.service.v1.Review](#review-service-v1-Review) |  |  |
| User | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Withdraw | [cloud.hashing.billing.v1.UserWithdrawItem](#cloud-hashing-billing-v1-UserWithdrawItem) |  |  |





 

 

 


<a name="cloud-hashing-apis-v1-CloudHashingApis"></a>

### CloudHashingApis
Cloud Hashing Goods

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| GetGoods | [GetGoodsRequest](#cloud-hashing-apis-v1-GetGoodsRequest) | [GetGoodsResponse](#cloud-hashing-apis-v1-GetGoodsResponse) |  |
| GetGoodsByApp | [GetGoodsByAppRequest](#cloud-hashing-apis-v1-GetGoodsByAppRequest) | [GetGoodsByAppResponse](#cloud-hashing-apis-v1-GetGoodsByAppResponse) |  |
| CreateGood | [CreateGoodRequest](#cloud-hashing-apis-v1-CreateGoodRequest) | [CreateGoodResponse](#cloud-hashing-apis-v1-CreateGoodResponse) |  |
| GetGood | [GetGoodRequest](#cloud-hashing-apis-v1-GetGoodRequest) | [GetGoodResponse](#cloud-hashing-apis-v1-GetGoodResponse) |  |
| GetRecommendGoodsByApp | [GetRecommendGoodsByAppRequest](#cloud-hashing-apis-v1-GetRecommendGoodsByAppRequest) | [GetRecommendGoodsByAppResponse](#cloud-hashing-apis-v1-GetRecommendGoodsByAppResponse) |  |
| SubmitOrder | [SubmitOrderRequest](#cloud-hashing-apis-v1-SubmitOrderRequest) | [SubmitOrderResponse](#cloud-hashing-apis-v1-SubmitOrderResponse) |  |
| CreateOrderPayment | [CreateOrderPaymentRequest](#cloud-hashing-apis-v1-CreateOrderPaymentRequest) | [CreateOrderPaymentResponse](#cloud-hashing-apis-v1-CreateOrderPaymentResponse) |  |
| GetOrder | [GetOrderRequest](#cloud-hashing-apis-v1-GetOrderRequest) | [GetOrderResponse](#cloud-hashing-apis-v1-GetOrderResponse) |  |
| GetOrdersByAppUser | [GetOrdersByAppUserRequest](#cloud-hashing-apis-v1-GetOrdersByAppUserRequest) | [GetOrdersByAppUserResponse](#cloud-hashing-apis-v1-GetOrdersByAppUserResponse) |  |
| GetOrdersByApp | [GetOrdersByAppRequest](#cloud-hashing-apis-v1-GetOrdersByAppRequest) | [GetOrdersByAppResponse](#cloud-hashing-apis-v1-GetOrdersByAppResponse) |  |
| GetOrdersByGood | [GetOrdersByGoodRequest](#cloud-hashing-apis-v1-GetOrdersByGoodRequest) | [GetOrdersByGoodResponse](#cloud-hashing-apis-v1-GetOrdersByGoodResponse) |  |
| Signup | [SignupRequest](#cloud-hashing-apis-v1-SignupRequest) | [SignupResponse](#cloud-hashing-apis-v1-SignupResponse) |  |
| UpdatePassword | [UpdatePasswordRequest](#cloud-hashing-apis-v1-UpdatePasswordRequest) | [UpdatePasswordResponse](#cloud-hashing-apis-v1-UpdatePasswordResponse) |  |
| UpdatePasswordByAppUser | [UpdatePasswordByAppUserRequest](#cloud-hashing-apis-v1-UpdatePasswordByAppUserRequest) | [UpdatePasswordByAppUserResponse](#cloud-hashing-apis-v1-UpdatePasswordByAppUserResponse) |  |
| UpdateEmailAddress | [UpdateEmailAddressRequest](#cloud-hashing-apis-v1-UpdateEmailAddressRequest) | [UpdateEmailAddressResponse](#cloud-hashing-apis-v1-UpdateEmailAddressResponse) |  |
| UpdatePhoneNO | [UpdatePhoneNORequest](#cloud-hashing-apis-v1-UpdatePhoneNORequest) | [UpdatePhoneNOResponse](#cloud-hashing-apis-v1-UpdatePhoneNOResponse) |  |
| UpdateAccount | [UpdateAccountRequest](#cloud-hashing-apis-v1-UpdateAccountRequest) | [UpdateAccountResponse](#cloud-hashing-apis-v1-UpdateAccountResponse) |  |
| GetReferrals | [GetReferralsRequest](#cloud-hashing-apis-v1-GetReferralsRequest) | [GetReferralsResponse](#cloud-hashing-apis-v1-GetReferralsResponse) |  |
| GetLayeredReferrals | [GetLayeredReferralsRequest](#cloud-hashing-apis-v1-GetLayeredReferralsRequest) | [GetLayeredReferralsResponse](#cloud-hashing-apis-v1-GetLayeredReferralsResponse) |  |
| GetKycReviews | [GetKycReviewsRequest](#cloud-hashing-apis-v1-GetKycReviewsRequest) | [GetKycReviewsResponse](#cloud-hashing-apis-v1-GetKycReviewsResponse) |  |
| GetKycReviewsByApp | [GetKycReviewsByAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByAppRequest) | [GetKycReviewsByAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByAppResponse) |  |
| GetKycReviewsByOtherApp | [GetKycReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppRequest) | [GetKycReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetKycReviewsByOtherAppResponse) |  |
| GetGoodReviews | [GetGoodReviewsRequest](#cloud-hashing-apis-v1-GetGoodReviewsRequest) | [GetGoodReviewsResponse](#cloud-hashing-apis-v1-GetGoodReviewsResponse) |  |
| GetWithdrawReviews | [GetWithdrawReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsRequest) | [GetWithdrawReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsResponse) |  |
| GetWithdrawReviewsByApp | [GetWithdrawReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppRequest) | [GetWithdrawReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByAppResponse) |  |
| GetWithdrawReviewsByOtherApp | [GetWithdrawReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppRequest) | [GetWithdrawReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawReviewsByOtherAppResponse) |  |
| GetWithdrawAddressReviews | [GetWithdrawAddressReviewsRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsRequest) | [GetWithdrawAddressReviewsResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsResponse) |  |
| GetWithdrawAddressReviewsByApp | [GetWithdrawAddressReviewsByAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppRequest) | [GetWithdrawAddressReviewsByAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByAppResponse) |  |
| GetWithdrawAddressReviewsByOtherApp | [GetWithdrawAddressReviewsByOtherAppRequest](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppRequest) | [GetWithdrawAddressReviewsByOtherAppResponse](#cloud-hashing-apis-v1-GetWithdrawAddressReviewsByOtherAppResponse) |  |
| CreateKyc | [CreateKycRequest](#cloud-hashing-apis-v1-CreateKycRequest) | [CreateKycResponse](#cloud-hashing-apis-v1-CreateKycResponse) |  |
| UpdateKyc | [UpdateKycRequest](#cloud-hashing-apis-v1-UpdateKycRequest) | [UpdateKycResponse](#cloud-hashing-apis-v1-UpdateKycResponse) |  |
| GetKycByAppUser | [GetKycByAppUserRequest](#cloud-hashing-apis-v1-GetKycByAppUserRequest) | [GetKycByAppUserResponse](#cloud-hashing-apis-v1-GetKycByAppUserResponse) |  |
| CreatePlatformCoinAccount | [CreatePlatformCoinAccountRequest](#cloud-hashing-apis-v1-CreatePlatformCoinAccountRequest) | [CreatePlatformCoinAccountResponse](#cloud-hashing-apis-v1-CreatePlatformCoinAccountResponse) |  |
| CreateUserCoinAccount | [CreateUserCoinAccountRequest](#cloud-hashing-apis-v1-CreateUserCoinAccountRequest) | [CreateUserCoinAccountResponse](#cloud-hashing-apis-v1-CreateUserCoinAccountResponse) |  |
| SubmitUserWithdraw | [SubmitUserWithdrawRequest](#cloud-hashing-apis-v1-SubmitUserWithdrawRequest) | [SubmitUserWithdrawResponse](#cloud-hashing-apis-v1-SubmitUserWithdrawResponse) |  |
| UpdateUserWithdrawReview | [UpdateUserWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewRequest) | [UpdateUserWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewResponse) |  |
| UpdateUserWithdrawReviewForOtherAppUser | [UpdateUserWithdrawReviewForOtherAppUserRequest](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserRequest) | [UpdateUserWithdrawReviewForOtherAppUserResponse](#cloud-hashing-apis-v1-UpdateUserWithdrawReviewForOtherAppUserResponse) |  |
| GetUserWithdrawsByAppUser | [GetUserWithdrawsByAppUserRequest](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserRequest) | [GetUserWithdrawsByAppUserResponse](#cloud-hashing-apis-v1-GetUserWithdrawsByAppUserResponse) |  |
| SetWithdrawAddress | [SetWithdrawAddressRequest](#cloud-hashing-apis-v1-SetWithdrawAddressRequest) | [SetWithdrawAddressResponse](#cloud-hashing-apis-v1-SetWithdrawAddressResponse) |  |
| GetWithdrawAddressesByAppUser | [GetWithdrawAddressesByAppUserRequest](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserRequest) | [GetWithdrawAddressesByAppUserResponse](#cloud-hashing-apis-v1-GetWithdrawAddressesByAppUserResponse) |  |
| GetCouponsByAppUser | [GetCouponsByAppUserRequest](#cloud-hashing-apis-v1-GetCouponsByAppUserRequest) | [GetCouponsByAppUserResponse](#cloud-hashing-apis-v1-GetCouponsByAppUserResponse) |  |
| GetCommissionByAppUser | [GetCommissionByAppUserRequest](#cloud-hashing-apis-v1-GetCommissionByAppUserRequest) | [GetCommissionByAppUserResponse](#cloud-hashing-apis-v1-GetCommissionByAppUserResponse) |  |
| UpdateKycReview | [UpdateKycReviewRequest](#cloud-hashing-apis-v1-UpdateKycReviewRequest) | [UpdateKycReviewResponse](#cloud-hashing-apis-v1-UpdateKycReviewResponse) |  |
| UpdateWithdrawReview | [UpdateWithdrawReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawReviewRequest) | [UpdateWithdrawReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawReviewResponse) |  |
| UpdateWithdrawAddressReview | [UpdateWithdrawAddressReviewRequest](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewRequest) | [UpdateWithdrawAddressReviewResponse](#cloud-hashing-apis-v1-UpdateWithdrawAddressReviewResponse) |  |

 



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

