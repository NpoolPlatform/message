# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-billing/cloud-hashing-billing.proto](#npool/cloud-hashing-billing/cloud-hashing-billing.proto)
    - [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting)
    - [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo)
    - [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction)
    - [CoinAccountTransactionDetail](#cloud.hashing.billing.v1.CoinAccountTransactionDetail)
    - [CoinSetting](#cloud.hashing.billing.v1.CoinSetting)
    - [CreateAppWithdrawSettingForOtherAppRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest)
    - [CreateAppWithdrawSettingForOtherAppResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse)
    - [CreateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest)
    - [CreateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse)
    - [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest)
    - [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse)
    - [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest)
    - [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse)
    - [CreateCoinSettingRequest](#cloud.hashing.billing.v1.CreateCoinSettingRequest)
    - [CreateCoinSettingResponse](#cloud.hashing.billing.v1.CreateCoinSettingResponse)
    - [CreateGoodBenefitRequest](#cloud.hashing.billing.v1.CreateGoodBenefitRequest)
    - [CreateGoodBenefitResponse](#cloud.hashing.billing.v1.CreateGoodBenefitResponse)
    - [CreateGoodPaymentRequest](#cloud.hashing.billing.v1.CreateGoodPaymentRequest)
    - [CreateGoodPaymentResponse](#cloud.hashing.billing.v1.CreateGoodPaymentResponse)
    - [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest)
    - [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse)
    - [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest)
    - [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse)
    - [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest)
    - [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse)
    - [CreateUserDirectBenefitRequest](#cloud.hashing.billing.v1.CreateUserDirectBenefitRequest)
    - [CreateUserDirectBenefitResponse](#cloud.hashing.billing.v1.CreateUserDirectBenefitResponse)
    - [CreateUserPaymentBalanceForOtherAppUserRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest)
    - [CreateUserPaymentBalanceForOtherAppUserResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse)
    - [CreateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest)
    - [CreateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse)
    - [CreateUserWithdrawItemRequest](#cloud.hashing.billing.v1.CreateUserWithdrawItemRequest)
    - [CreateUserWithdrawItemResponse](#cloud.hashing.billing.v1.CreateUserWithdrawItemResponse)
    - [CreateUserWithdrawRequest](#cloud.hashing.billing.v1.CreateUserWithdrawRequest)
    - [CreateUserWithdrawResponse](#cloud.hashing.billing.v1.CreateUserWithdrawResponse)
    - [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest)
    - [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse)
    - [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest)
    - [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse)
    - [GetAppWithdrawSettingByAppCoinRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest)
    - [GetAppWithdrawSettingByAppCoinResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse)
    - [GetAppWithdrawSettingRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingRequest)
    - [GetAppWithdrawSettingResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingResponse)
    - [GetAppWithdrawSettingsByAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest)
    - [GetAppWithdrawSettingsByAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse)
    - [GetAppWithdrawSettingsByOtherAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest)
    - [GetAppWithdrawSettingsByOtherAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse)
    - [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest)
    - [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse)
    - [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest)
    - [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse)
    - [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest)
    - [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse)
    - [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest)
    - [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse)
    - [GetCoinAccountTransactionsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest)
    - [GetCoinAccountTransactionsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse)
    - [GetCoinAccountTransactionsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest)
    - [GetCoinAccountTransactionsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse)
    - [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest)
    - [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse)
    - [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest)
    - [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse)
    - [GetCoinAccountTransactionsByGoodStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest)
    - [GetCoinAccountTransactionsByGoodStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse)
    - [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest)
    - [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse)
    - [GetCoinAccountTransactionsRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest)
    - [GetCoinAccountTransactionsResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse)
    - [GetCoinAccountsRequest](#cloud.hashing.billing.v1.GetCoinAccountsRequest)
    - [GetCoinAccountsResponse](#cloud.hashing.billing.v1.GetCoinAccountsResponse)
    - [GetCoinSettingByCoinRequest](#cloud.hashing.billing.v1.GetCoinSettingByCoinRequest)
    - [GetCoinSettingByCoinResponse](#cloud.hashing.billing.v1.GetCoinSettingByCoinResponse)
    - [GetCoinSettingRequest](#cloud.hashing.billing.v1.GetCoinSettingRequest)
    - [GetCoinSettingResponse](#cloud.hashing.billing.v1.GetCoinSettingResponse)
    - [GetGoodBenefitByGoodRequest](#cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest)
    - [GetGoodBenefitByGoodResponse](#cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse)
    - [GetGoodBenefitRequest](#cloud.hashing.billing.v1.GetGoodBenefitRequest)
    - [GetGoodBenefitResponse](#cloud.hashing.billing.v1.GetGoodBenefitResponse)
    - [GetGoodPaymentByAccountRequest](#cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest)
    - [GetGoodPaymentByAccountResponse](#cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse)
    - [GetGoodPaymentRequest](#cloud.hashing.billing.v1.GetGoodPaymentRequest)
    - [GetGoodPaymentResponse](#cloud.hashing.billing.v1.GetGoodPaymentResponse)
    - [GetGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest)
    - [GetGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse)
    - [GetGoodPaymentsRequest](#cloud.hashing.billing.v1.GetGoodPaymentsRequest)
    - [GetGoodPaymentsResponse](#cloud.hashing.billing.v1.GetGoodPaymentsResponse)
    - [GetIdleGoodPaymentsByGoodPaymentCoinRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest)
    - [GetIdleGoodPaymentsByGoodPaymentCoinResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse)
    - [GetIdleGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest)
    - [GetIdleGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse)
    - [GetLatestPlatformBenefitByGoodRequest](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest)
    - [GetLatestPlatformBenefitByGoodResponse](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse)
    - [GetLatestUserBenefitByGoodAppUserRequest](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest)
    - [GetLatestUserBenefitByGoodAppUserResponse](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse)
    - [GetPlatformBenefitDetailRequest](#cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest)
    - [GetPlatformBenefitDetailResponse](#cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse)
    - [GetPlatformBenefitRequest](#cloud.hashing.billing.v1.GetPlatformBenefitRequest)
    - [GetPlatformBenefitResponse](#cloud.hashing.billing.v1.GetPlatformBenefitResponse)
    - [GetPlatformBenefitsByGoodRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest)
    - [GetPlatformBenefitsByGoodResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse)
    - [GetPlatformBenefitsRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsRequest)
    - [GetPlatformBenefitsResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsResponse)
    - [GetPlatformSettingByGoodRequest](#cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest)
    - [GetPlatformSettingByGoodResponse](#cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse)
    - [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest)
    - [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse)
    - [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest)
    - [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse)
    - [GetUserBenefitsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest)
    - [GetUserBenefitsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse)
    - [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest)
    - [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse)
    - [GetUserBenefitsRequest](#cloud.hashing.billing.v1.GetUserBenefitsRequest)
    - [GetUserBenefitsResponse](#cloud.hashing.billing.v1.GetUserBenefitsResponse)
    - [GetUserDirectBenefitByAccountRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest)
    - [GetUserDirectBenefitByAccountResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse)
    - [GetUserDirectBenefitRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitRequest)
    - [GetUserDirectBenefitResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitResponse)
    - [GetUserDirectBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest)
    - [GetUserDirectBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse)
    - [GetUserDirectBenefitsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest)
    - [GetUserDirectBenefitsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse)
    - [GetUserPaymentBalanceRequest](#cloud.hashing.billing.v1.GetUserPaymentBalanceRequest)
    - [GetUserPaymentBalanceResponse](#cloud.hashing.billing.v1.GetUserPaymentBalanceResponse)
    - [GetUserPaymentBalancesByAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest)
    - [GetUserPaymentBalancesByAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse)
    - [GetUserPaymentBalancesByAppUserRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest)
    - [GetUserPaymentBalancesByAppUserResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse)
    - [GetUserPaymentBalancesByOtherAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest)
    - [GetUserPaymentBalancesByOtherAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse)
    - [GetUserWithdrawAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountRequest)
    - [GetUserWithdrawAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountResponse)
    - [GetUserWithdrawAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest)
    - [GetUserWithdrawAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse)
    - [GetUserWithdrawAccountsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest)
    - [GetUserWithdrawAccountsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse)
    - [GetUserWithdrawByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest)
    - [GetUserWithdrawByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse)
    - [GetUserWithdrawItemRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemRequest)
    - [GetUserWithdrawItemResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemResponse)
    - [GetUserWithdrawItemsByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest)
    - [GetUserWithdrawItemsByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse)
    - [GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest)
    - [GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse)
    - [GetUserWithdrawItemsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest)
    - [GetUserWithdrawItemsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse)
    - [GetUserWithdrawItemsByAppUserWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest)
    - [GetUserWithdrawItemsByAppUserWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse)
    - [GetUserWithdrawItemsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest)
    - [GetUserWithdrawItemsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse)
    - [GetUserWithdrawItemsRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsRequest)
    - [GetUserWithdrawItemsResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsResponse)
    - [GetUserWithdrawRequest](#cloud.hashing.billing.v1.GetUserWithdrawRequest)
    - [GetUserWithdrawResponse](#cloud.hashing.billing.v1.GetUserWithdrawResponse)
    - [GetUserWithdrawsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest)
    - [GetUserWithdrawsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse)
    - [GetUserWithdrawsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest)
    - [GetUserWithdrawsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse)
    - [GetUserWithdrawsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest)
    - [GetUserWithdrawsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse)
    - [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit)
    - [GoodPayment](#cloud.hashing.billing.v1.GoodPayment)
    - [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit)
    - [PlatformBenefitDetail](#cloud.hashing.billing.v1.PlatformBenefitDetail)
    - [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting)
    - [UpdateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest)
    - [UpdateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse)
    - [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest)
    - [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse)
    - [UpdateCoinSettingRequest](#cloud.hashing.billing.v1.UpdateCoinSettingRequest)
    - [UpdateCoinSettingResponse](#cloud.hashing.billing.v1.UpdateCoinSettingResponse)
    - [UpdateGoodBenefitRequest](#cloud.hashing.billing.v1.UpdateGoodBenefitRequest)
    - [UpdateGoodBenefitResponse](#cloud.hashing.billing.v1.UpdateGoodBenefitResponse)
    - [UpdateGoodPaymentRequest](#cloud.hashing.billing.v1.UpdateGoodPaymentRequest)
    - [UpdateGoodPaymentResponse](#cloud.hashing.billing.v1.UpdateGoodPaymentResponse)
    - [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest)
    - [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse)
    - [UpdateUserDirectBenefitRequest](#cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest)
    - [UpdateUserDirectBenefitResponse](#cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse)
    - [UpdateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest)
    - [UpdateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse)
    - [UpdateUserWithdrawItemRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest)
    - [UpdateUserWithdrawItemResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse)
    - [UpdateUserWithdrawRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawRequest)
    - [UpdateUserWithdrawResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawResponse)
    - [UserBenefit](#cloud.hashing.billing.v1.UserBenefit)
    - [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit)
    - [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance)
    - [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw)
    - [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount)
    - [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem)
  
    - [CloudHashingBilling](#cloud.hashing.billing.v1.CloudHashingBilling)
  
- [npool/cloud-hashing-billing/cloud-hashing-billing.proto](#npool/cloud-hashing-billing/cloud-hashing-billing.proto)
    - [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting)
    - [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo)
    - [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction)
    - [CoinAccountTransactionDetail](#cloud.hashing.billing.v1.CoinAccountTransactionDetail)
    - [CoinSetting](#cloud.hashing.billing.v1.CoinSetting)
    - [CreateAppWithdrawSettingForOtherAppRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest)
    - [CreateAppWithdrawSettingForOtherAppResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse)
    - [CreateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest)
    - [CreateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse)
    - [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest)
    - [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse)
    - [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest)
    - [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse)
    - [CreateCoinSettingRequest](#cloud.hashing.billing.v1.CreateCoinSettingRequest)
    - [CreateCoinSettingResponse](#cloud.hashing.billing.v1.CreateCoinSettingResponse)
    - [CreateGoodBenefitRequest](#cloud.hashing.billing.v1.CreateGoodBenefitRequest)
    - [CreateGoodBenefitResponse](#cloud.hashing.billing.v1.CreateGoodBenefitResponse)
    - [CreateGoodPaymentRequest](#cloud.hashing.billing.v1.CreateGoodPaymentRequest)
    - [CreateGoodPaymentResponse](#cloud.hashing.billing.v1.CreateGoodPaymentResponse)
    - [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest)
    - [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse)
    - [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest)
    - [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse)
    - [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest)
    - [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse)
    - [CreateUserDirectBenefitRequest](#cloud.hashing.billing.v1.CreateUserDirectBenefitRequest)
    - [CreateUserDirectBenefitResponse](#cloud.hashing.billing.v1.CreateUserDirectBenefitResponse)
    - [CreateUserPaymentBalanceForOtherAppUserRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest)
    - [CreateUserPaymentBalanceForOtherAppUserResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse)
    - [CreateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest)
    - [CreateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse)
    - [CreateUserWithdrawItemRequest](#cloud.hashing.billing.v1.CreateUserWithdrawItemRequest)
    - [CreateUserWithdrawItemResponse](#cloud.hashing.billing.v1.CreateUserWithdrawItemResponse)
    - [CreateUserWithdrawRequest](#cloud.hashing.billing.v1.CreateUserWithdrawRequest)
    - [CreateUserWithdrawResponse](#cloud.hashing.billing.v1.CreateUserWithdrawResponse)
    - [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest)
    - [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse)
    - [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest)
    - [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse)
    - [GetAppWithdrawSettingByAppCoinRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest)
    - [GetAppWithdrawSettingByAppCoinResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse)
    - [GetAppWithdrawSettingRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingRequest)
    - [GetAppWithdrawSettingResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingResponse)
    - [GetAppWithdrawSettingsByAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest)
    - [GetAppWithdrawSettingsByAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse)
    - [GetAppWithdrawSettingsByOtherAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest)
    - [GetAppWithdrawSettingsByOtherAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse)
    - [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest)
    - [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse)
    - [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest)
    - [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse)
    - [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest)
    - [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse)
    - [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest)
    - [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse)
    - [GetCoinAccountTransactionsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest)
    - [GetCoinAccountTransactionsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse)
    - [GetCoinAccountTransactionsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest)
    - [GetCoinAccountTransactionsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse)
    - [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest)
    - [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse)
    - [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest)
    - [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse)
    - [GetCoinAccountTransactionsByGoodStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest)
    - [GetCoinAccountTransactionsByGoodStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse)
    - [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest)
    - [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse)
    - [GetCoinAccountTransactionsRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest)
    - [GetCoinAccountTransactionsResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse)
    - [GetCoinAccountsRequest](#cloud.hashing.billing.v1.GetCoinAccountsRequest)
    - [GetCoinAccountsResponse](#cloud.hashing.billing.v1.GetCoinAccountsResponse)
    - [GetCoinSettingByCoinRequest](#cloud.hashing.billing.v1.GetCoinSettingByCoinRequest)
    - [GetCoinSettingByCoinResponse](#cloud.hashing.billing.v1.GetCoinSettingByCoinResponse)
    - [GetCoinSettingRequest](#cloud.hashing.billing.v1.GetCoinSettingRequest)
    - [GetCoinSettingResponse](#cloud.hashing.billing.v1.GetCoinSettingResponse)
    - [GetGoodBenefitByGoodRequest](#cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest)
    - [GetGoodBenefitByGoodResponse](#cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse)
    - [GetGoodBenefitRequest](#cloud.hashing.billing.v1.GetGoodBenefitRequest)
    - [GetGoodBenefitResponse](#cloud.hashing.billing.v1.GetGoodBenefitResponse)
    - [GetGoodPaymentByAccountRequest](#cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest)
    - [GetGoodPaymentByAccountResponse](#cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse)
    - [GetGoodPaymentRequest](#cloud.hashing.billing.v1.GetGoodPaymentRequest)
    - [GetGoodPaymentResponse](#cloud.hashing.billing.v1.GetGoodPaymentResponse)
    - [GetGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest)
    - [GetGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse)
    - [GetGoodPaymentsRequest](#cloud.hashing.billing.v1.GetGoodPaymentsRequest)
    - [GetGoodPaymentsResponse](#cloud.hashing.billing.v1.GetGoodPaymentsResponse)
    - [GetIdleGoodPaymentsByGoodPaymentCoinRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest)
    - [GetIdleGoodPaymentsByGoodPaymentCoinResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse)
    - [GetIdleGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest)
    - [GetIdleGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse)
    - [GetLatestPlatformBenefitByGoodRequest](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest)
    - [GetLatestPlatformBenefitByGoodResponse](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse)
    - [GetLatestUserBenefitByGoodAppUserRequest](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest)
    - [GetLatestUserBenefitByGoodAppUserResponse](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse)
    - [GetPlatformBenefitDetailRequest](#cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest)
    - [GetPlatformBenefitDetailResponse](#cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse)
    - [GetPlatformBenefitRequest](#cloud.hashing.billing.v1.GetPlatformBenefitRequest)
    - [GetPlatformBenefitResponse](#cloud.hashing.billing.v1.GetPlatformBenefitResponse)
    - [GetPlatformBenefitsByGoodRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest)
    - [GetPlatformBenefitsByGoodResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse)
    - [GetPlatformBenefitsRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsRequest)
    - [GetPlatformBenefitsResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsResponse)
    - [GetPlatformSettingByGoodRequest](#cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest)
    - [GetPlatformSettingByGoodResponse](#cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse)
    - [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest)
    - [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse)
    - [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest)
    - [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse)
    - [GetUserBenefitsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest)
    - [GetUserBenefitsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse)
    - [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest)
    - [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse)
    - [GetUserBenefitsRequest](#cloud.hashing.billing.v1.GetUserBenefitsRequest)
    - [GetUserBenefitsResponse](#cloud.hashing.billing.v1.GetUserBenefitsResponse)
    - [GetUserDirectBenefitByAccountRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest)
    - [GetUserDirectBenefitByAccountResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse)
    - [GetUserDirectBenefitRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitRequest)
    - [GetUserDirectBenefitResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitResponse)
    - [GetUserDirectBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest)
    - [GetUserDirectBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse)
    - [GetUserDirectBenefitsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest)
    - [GetUserDirectBenefitsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse)
    - [GetUserPaymentBalanceRequest](#cloud.hashing.billing.v1.GetUserPaymentBalanceRequest)
    - [GetUserPaymentBalanceResponse](#cloud.hashing.billing.v1.GetUserPaymentBalanceResponse)
    - [GetUserPaymentBalancesByAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest)
    - [GetUserPaymentBalancesByAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse)
    - [GetUserPaymentBalancesByAppUserRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest)
    - [GetUserPaymentBalancesByAppUserResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse)
    - [GetUserPaymentBalancesByOtherAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest)
    - [GetUserPaymentBalancesByOtherAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse)
    - [GetUserWithdrawAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountRequest)
    - [GetUserWithdrawAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountResponse)
    - [GetUserWithdrawAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest)
    - [GetUserWithdrawAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse)
    - [GetUserWithdrawAccountsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest)
    - [GetUserWithdrawAccountsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse)
    - [GetUserWithdrawByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest)
    - [GetUserWithdrawByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse)
    - [GetUserWithdrawItemRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemRequest)
    - [GetUserWithdrawItemResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemResponse)
    - [GetUserWithdrawItemsByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest)
    - [GetUserWithdrawItemsByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse)
    - [GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest)
    - [GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse)
    - [GetUserWithdrawItemsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest)
    - [GetUserWithdrawItemsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse)
    - [GetUserWithdrawItemsByAppUserWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest)
    - [GetUserWithdrawItemsByAppUserWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse)
    - [GetUserWithdrawItemsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest)
    - [GetUserWithdrawItemsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse)
    - [GetUserWithdrawItemsRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsRequest)
    - [GetUserWithdrawItemsResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsResponse)
    - [GetUserWithdrawRequest](#cloud.hashing.billing.v1.GetUserWithdrawRequest)
    - [GetUserWithdrawResponse](#cloud.hashing.billing.v1.GetUserWithdrawResponse)
    - [GetUserWithdrawsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest)
    - [GetUserWithdrawsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse)
    - [GetUserWithdrawsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest)
    - [GetUserWithdrawsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse)
    - [GetUserWithdrawsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest)
    - [GetUserWithdrawsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse)
    - [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit)
    - [GoodPayment](#cloud.hashing.billing.v1.GoodPayment)
    - [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit)
    - [PlatformBenefitDetail](#cloud.hashing.billing.v1.PlatformBenefitDetail)
    - [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting)
    - [UpdateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest)
    - [UpdateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse)
    - [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest)
    - [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse)
    - [UpdateCoinSettingRequest](#cloud.hashing.billing.v1.UpdateCoinSettingRequest)
    - [UpdateCoinSettingResponse](#cloud.hashing.billing.v1.UpdateCoinSettingResponse)
    - [UpdateGoodBenefitRequest](#cloud.hashing.billing.v1.UpdateGoodBenefitRequest)
    - [UpdateGoodBenefitResponse](#cloud.hashing.billing.v1.UpdateGoodBenefitResponse)
    - [UpdateGoodPaymentRequest](#cloud.hashing.billing.v1.UpdateGoodPaymentRequest)
    - [UpdateGoodPaymentResponse](#cloud.hashing.billing.v1.UpdateGoodPaymentResponse)
    - [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest)
    - [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse)
    - [UpdateUserDirectBenefitRequest](#cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest)
    - [UpdateUserDirectBenefitResponse](#cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse)
    - [UpdateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest)
    - [UpdateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse)
    - [UpdateUserWithdrawItemRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest)
    - [UpdateUserWithdrawItemResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse)
    - [UpdateUserWithdrawRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawRequest)
    - [UpdateUserWithdrawResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawResponse)
    - [UserBenefit](#cloud.hashing.billing.v1.UserBenefit)
    - [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit)
    - [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance)
    - [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw)
    - [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount)
    - [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem)
  
    - [CloudHashingBilling](#cloud.hashing.billing.v1.CloudHashingBilling)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-billing/cloud-hashing-billing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-billing/cloud-hashing-billing.proto



<a name="cloud.hashing.billing.v1.AppWithdrawSetting"></a>

### AppWithdrawSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawAutoReviewCoinAmount | [double](#double) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountInfo"></a>

### CoinAccountInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| PlatformHoldPrivateKey | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountTransaction"></a>

### CoinAccountTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FromAddressID | [string](#string) |  |  |
| ToAddressID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| FailHold | [bool](#bool) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountTransactionDetail"></a>

### CoinAccountTransactionDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FromAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| ToAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.CoinSetting"></a>

### CoinSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WarmAccountCoinAmount | [double](#double) |  |  |
| PaymentAccountCoinAmount | [double](#double) |  |  |
| PlatformOfflineAccountID | [string](#string) |  |  |
| UserOfflineAccountID | [string](#string) |  |  |
| UserOnlineAccountID | [string](#string) |  |  |
| GoodIncomingAccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest"></a>

### CreateAppWithdrawSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse"></a>

### CreateAppWithdrawSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest"></a>

### CreateAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse"></a>

### CreateAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountRequest"></a>

### CreateCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountResponse"></a>

### CreateCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest"></a>

### CreateCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse"></a>

### CreateCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinSettingRequest"></a>

### CreateCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinSettingResponse"></a>

### CreateCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodBenefitRequest"></a>

### CreateGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodBenefitResponse"></a>

### CreateGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodPaymentRequest"></a>

### CreateGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodPaymentResponse"></a>

### CreateGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformBenefitRequest"></a>

### CreatePlatformBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformBenefitResponse"></a>

### CreatePlatformBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformSettingRequest"></a>

### CreatePlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformSettingResponse"></a>

### CreatePlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserBenefitRequest"></a>

### CreateUserBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserBenefitResponse"></a>

### CreateUserBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserDirectBenefitRequest"></a>

### CreateUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserDirectBenefitResponse"></a>

### CreateUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest"></a>

### CreateUserPaymentBalanceForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse"></a>

### CreateUserPaymentBalanceForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest"></a>

### CreateUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse"></a>

### CreateUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawItemRequest"></a>

### CreateUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawItemResponse"></a>

### CreateUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawRequest"></a>

### CreateUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawResponse"></a>

### CreateUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountRequest"></a>

### DeleteCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountResponse"></a>

### DeleteCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest"></a>

### DeleteCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse"></a>

### DeleteCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest"></a>

### GetAppWithdrawSettingByAppCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse"></a>

### GetAppWithdrawSettingByAppCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingRequest"></a>

### GetAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingResponse"></a>

### GetAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest"></a>

### GetAppWithdrawSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse"></a>

### GetAppWithdrawSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) | repeated |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest"></a>

### GetAppWithdrawSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse"></a>

### GetAppWithdrawSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest"></a>

### GetCoinAccountByCoinAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinInfoID | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse"></a>

### GetCoinAccountByCoinAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountRequest"></a>

### GetCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountResponse"></a>

### GetCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest"></a>

### GetCoinAccountTransactionDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse"></a>

### GetCoinAccountTransactionDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [CoinAccountTransactionDetail](#cloud.hashing.billing.v1.CoinAccountTransactionDetail) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionRequest"></a>

### GetCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionResponse"></a>

### GetCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest"></a>

### GetCoinAccountTransactionsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse"></a>

### GetCoinAccountTransactionsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest"></a>

### GetCoinAccountTransactionsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse"></a>

### GetCoinAccountTransactionsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest"></a>

### GetCoinAccountTransactionsByCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| AddressID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse"></a>

### GetCoinAccountTransactionsByCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest"></a>

### GetCoinAccountTransactionsByCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse"></a>

### GetCoinAccountTransactionsByCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest"></a>

### GetCoinAccountTransactionsByGoodStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| State | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse"></a>

### GetCoinAccountTransactionsByGoodStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest"></a>

### GetCoinAccountTransactionsByStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse"></a>

### GetCoinAccountTransactionsByStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest"></a>

### GetCoinAccountTransactionsRequest







<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse"></a>

### GetCoinAccountTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountsRequest"></a>

### GetCoinAccountsRequest







<a name="cloud.hashing.billing.v1.GetCoinAccountsResponse"></a>

### GetCoinAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingByCoinRequest"></a>

### GetCoinSettingByCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingByCoinResponse"></a>

### GetCoinSettingByCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingRequest"></a>

### GetCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingResponse"></a>

### GetCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest"></a>

### GetGoodBenefitByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse"></a>

### GetGoodBenefitByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitRequest"></a>

### GetGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitResponse"></a>

### GetGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest"></a>

### GetGoodPaymentByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse"></a>

### GetGoodPaymentByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentRequest"></a>

### GetGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentResponse"></a>

### GetGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest"></a>

### GetGoodPaymentsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse"></a>

### GetGoodPaymentsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsRequest"></a>

### GetGoodPaymentsRequest







<a name="cloud.hashing.billing.v1.GetGoodPaymentsResponse"></a>

### GetGoodPaymentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest"></a>

### GetIdleGoodPaymentsByGoodPaymentCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse"></a>

### GetIdleGoodPaymentsByGoodPaymentCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest"></a>

### GetIdleGoodPaymentsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse"></a>

### GetIdleGoodPaymentsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest"></a>

### GetLatestPlatformBenefitByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse"></a>

### GetLatestPlatformBenefitByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest"></a>

### GetLatestUserBenefitByGoodAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse"></a>

### GetLatestUserBenefitByGoodAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest"></a>

### GetPlatformBenefitDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse"></a>

### GetPlatformBenefitDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [PlatformBenefitDetail](#cloud.hashing.billing.v1.PlatformBenefitDetail) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitRequest"></a>

### GetPlatformBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitResponse"></a>

### GetPlatformBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest"></a>

### GetPlatformBenefitsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse"></a>

### GetPlatformBenefitsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsRequest"></a>

### GetPlatformBenefitsRequest







<a name="cloud.hashing.billing.v1.GetPlatformBenefitsResponse"></a>

### GetPlatformBenefitsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest"></a>

### GetPlatformSettingByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse"></a>

### GetPlatformSettingByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingRequest"></a>

### GetPlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingResponse"></a>

### GetPlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppRequest"></a>

### GetUserBenefitsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppResponse"></a>

### GetUserBenefitsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest"></a>

### GetUserBenefitsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse"></a>

### GetUserBenefitsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest"></a>

### GetUserBenefitsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse"></a>

### GetUserBenefitsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsRequest"></a>

### GetUserBenefitsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsResponse"></a>

### GetUserBenefitsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest"></a>

### GetUserDirectBenefitByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse"></a>

### GetUserDirectBenefitByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitRequest"></a>

### GetUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitResponse"></a>

### GetUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest"></a>

### GetUserDirectBenefitsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse"></a>

### GetUserDirectBenefitsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest"></a>

### GetUserDirectBenefitsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse"></a>

### GetUserDirectBenefitsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalanceRequest"></a>

### GetUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalanceResponse"></a>

### GetUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest"></a>

### GetUserPaymentBalancesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse"></a>

### GetUserPaymentBalancesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest"></a>

### GetUserPaymentBalancesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse"></a>

### GetUserPaymentBalancesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest"></a>

### GetUserPaymentBalancesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse"></a>

### GetUserPaymentBalancesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountRequest"></a>

### GetUserWithdrawAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountResponse"></a>

### GetUserWithdrawAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest"></a>

### GetUserWithdrawAccountsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse"></a>

### GetUserWithdrawAccountsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest"></a>

### GetUserWithdrawAccountsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse"></a>

### GetUserWithdrawAccountsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest"></a>

### GetUserWithdrawByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse"></a>

### GetUserWithdrawByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemRequest"></a>

### GetUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemResponse"></a>

### GetUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest"></a>

### GetUserWithdrawItemsByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse"></a>

### GetUserWithdrawItemsByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest"></a>

### GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse"></a>

### GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest"></a>

### GetUserWithdrawItemsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse"></a>

### GetUserWithdrawItemsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest"></a>

### GetUserWithdrawItemsByAppUserWithdrawTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse"></a>

### GetUserWithdrawItemsByAppUserWithdrawTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest"></a>

### GetUserWithdrawItemsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse"></a>

### GetUserWithdrawItemsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsRequest"></a>

### GetUserWithdrawItemsRequest







<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsResponse"></a>

### GetUserWithdrawItemsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawRequest"></a>

### GetUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawResponse"></a>

### GetUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest"></a>

### GetUserWithdrawsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse"></a>

### GetUserWithdrawsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest"></a>

### GetUserWithdrawsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse"></a>

### GetUserWithdrawsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest"></a>

### GetUserWithdrawsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse"></a>

### GetUserWithdrawsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GoodBenefit"></a>

### GoodBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAccountID | [string](#string) |  |  |
| BenefitIntervalHours | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.GoodPayment"></a>

### GoodPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| Idle | [bool](#bool) |  |  |
| OccupiedBy | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.PlatformBenefit"></a>

### PlatformBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAccountID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| LastBenefitTimestamp | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.PlatformBenefitDetail"></a>

### PlatformBenefitDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| ChainTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.PlatformSetting"></a>

### PlatformSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| WarmAccountUSDAmount | [double](#double) |  |  |
| PaymentAccountUSDAmount | [double](#double) |  |  |
| WithdrawAutoReviewUSDAmount | [double](#double) |  |  |






<a name="cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest"></a>

### UpdateAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse"></a>

### UpdateAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest"></a>

### UpdateCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse"></a>

### UpdateCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinSettingRequest"></a>

### UpdateCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinSettingResponse"></a>

### UpdateCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodBenefitRequest"></a>

### UpdateGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodBenefitResponse"></a>

### UpdateGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodPaymentRequest"></a>

### UpdateGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodPaymentResponse"></a>

### UpdateGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.UpdatePlatformSettingRequest"></a>

### UpdatePlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdatePlatformSettingResponse"></a>

### UpdatePlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest"></a>

### UpdateUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse"></a>

### UpdateUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest"></a>

### UpdateUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse"></a>

### UpdateUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest"></a>

### UpdateUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse"></a>

### UpdateUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawRequest"></a>

### UpdateUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawResponse"></a>

### UpdateUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.UserBenefit"></a>

### UserBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| OrderID | [string](#string) |  |  |
| LastBenefitTimestamp | [uint32](#uint32) |  |  |
| PlatformTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserDirectBenefit"></a>

### UserDirectBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserPaymentBalance"></a>

### UserPaymentBalance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| UsedByPaymentID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserWithdraw"></a>

### UserWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| Labels | [string](#string) | repeated |  |






<a name="cloud.hashing.billing.v1.UserWithdrawAccount"></a>

### UserWithdrawAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Withdraw | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |
| Account | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.UserWithdrawItem"></a>

### UserWithdrawItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawToAccountID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| PlatformTransactionID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.billing.v1.CloudHashingBilling"></a>

### CloudHashingBilling
Cloud Hashing Billing

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| CreateCoinAccount | [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest) | [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse) |  |
| GetCoinAccount | [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest) | [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse) |  |
| GetCoinAccountByCoinAddress | [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest) | [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse) |  |
| GetCoinAccounts | [GetCoinAccountsRequest](#cloud.hashing.billing.v1.GetCoinAccountsRequest) | [GetCoinAccountsResponse](#cloud.hashing.billing.v1.GetCoinAccountsResponse) |  |
| DeleteCoinAccount | [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest) | [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse) |  |
| CreateCoinAccountTransaction | [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest) | [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransaction | [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest) | [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionsByCoinAccount | [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest) | [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse) |  |
| GetCoinAccountTransactionsByState | [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest) | [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse) |  |
| GetCoinAccountTransactionsByGoodState | [GetCoinAccountTransactionsByGoodStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest) | [GetCoinAccountTransactionsByGoodStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse) |  |
| GetCoinAccountTransactions | [GetCoinAccountTransactionsRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest) | [GetCoinAccountTransactionsResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse) |  |
| GetCoinAccountTransactionsByAppUser | [GetCoinAccountTransactionsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest) | [GetCoinAccountTransactionsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse) |  |
| GetCoinAccountTransactionsByAppUserCoin | [GetCoinAccountTransactionsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest) | [GetCoinAccountTransactionsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse) |  |
| GetCoinAccountTransactionsByCoin | [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest) | [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse) |  |
| UpdateCoinAccountTransaction | [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest) | [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionDetail | [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest) | [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse) |  |
| DeleteCoinAccountTransaction | [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest) | [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse) |  |
| CreatePlatformBenefit | [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest) | [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse) |  |
| GetLatestPlatformBenefitByGood | [GetLatestPlatformBenefitByGoodRequest](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest) | [GetLatestPlatformBenefitByGoodResponse](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse) |  |
| GetPlatformBenefitsByGood | [GetPlatformBenefitsByGoodRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest) | [GetPlatformBenefitsByGoodResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse) |  |
| GetPlatformBenefits | [GetPlatformBenefitsRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsRequest) | [GetPlatformBenefitsResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsResponse) |  |
| GetPlatformBenefit | [GetPlatformBenefitRequest](#cloud.hashing.billing.v1.GetPlatformBenefitRequest) | [GetPlatformBenefitResponse](#cloud.hashing.billing.v1.GetPlatformBenefitResponse) |  |
| GetPlatformBenefitDetail | [GetPlatformBenefitDetailRequest](#cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest) | [GetPlatformBenefitDetailResponse](#cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse) |  |
| CreatePlatformSetting | [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest) | [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse) |  |
| UpdatePlatformSetting | [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest) | [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse) |  |
| GetPlatformSetting | [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest) | [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse) |  |
| CreateUserBenefit | [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest) | [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse) |  |
| GetUserBenefitsByAppUser | [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest) | [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse) |  |
| GetUserBenefitsByAppUserCoin | [GetUserBenefitsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest) | [GetUserBenefitsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse) |  |
| GetUserBenefitsByApp | [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest) | [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse) |  |
| GetUserBenefits | [GetUserBenefitsRequest](#cloud.hashing.billing.v1.GetUserBenefitsRequest) | [GetUserBenefitsResponse](#cloud.hashing.billing.v1.GetUserBenefitsResponse) |  |
| GetLatestUserBenefitByGoodAppUser | [GetLatestUserBenefitByGoodAppUserRequest](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest) | [GetLatestUserBenefitByGoodAppUserResponse](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse) |  |
| CreateCoinSetting | [CreateCoinSettingRequest](#cloud.hashing.billing.v1.CreateCoinSettingRequest) | [CreateCoinSettingResponse](#cloud.hashing.billing.v1.CreateCoinSettingResponse) |  |
| UpdateCoinSetting | [UpdateCoinSettingRequest](#cloud.hashing.billing.v1.UpdateCoinSettingRequest) | [UpdateCoinSettingResponse](#cloud.hashing.billing.v1.UpdateCoinSettingResponse) |  |
| GetCoinSetting | [GetCoinSettingRequest](#cloud.hashing.billing.v1.GetCoinSettingRequest) | [GetCoinSettingResponse](#cloud.hashing.billing.v1.GetCoinSettingResponse) |  |
| GetCoinSettingByCoin | [GetCoinSettingByCoinRequest](#cloud.hashing.billing.v1.GetCoinSettingByCoinRequest) | [GetCoinSettingByCoinResponse](#cloud.hashing.billing.v1.GetCoinSettingByCoinResponse) |  |
| CreateGoodBenefit | [CreateGoodBenefitRequest](#cloud.hashing.billing.v1.CreateGoodBenefitRequest) | [CreateGoodBenefitResponse](#cloud.hashing.billing.v1.CreateGoodBenefitResponse) |  |
| UpdateGoodBenefit | [UpdateGoodBenefitRequest](#cloud.hashing.billing.v1.UpdateGoodBenefitRequest) | [UpdateGoodBenefitResponse](#cloud.hashing.billing.v1.UpdateGoodBenefitResponse) |  |
| GetGoodBenefit | [GetGoodBenefitRequest](#cloud.hashing.billing.v1.GetGoodBenefitRequest) | [GetGoodBenefitResponse](#cloud.hashing.billing.v1.GetGoodBenefitResponse) |  |
| GetGoodBenefitByGood | [GetGoodBenefitByGoodRequest](#cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest) | [GetGoodBenefitByGoodResponse](#cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse) |  |
| CreateGoodPayment | [CreateGoodPaymentRequest](#cloud.hashing.billing.v1.CreateGoodPaymentRequest) | [CreateGoodPaymentResponse](#cloud.hashing.billing.v1.CreateGoodPaymentResponse) |  |
| UpdateGoodPayment | [UpdateGoodPaymentRequest](#cloud.hashing.billing.v1.UpdateGoodPaymentRequest) | [UpdateGoodPaymentResponse](#cloud.hashing.billing.v1.UpdateGoodPaymentResponse) |  |
| GetGoodPayment | [GetGoodPaymentRequest](#cloud.hashing.billing.v1.GetGoodPaymentRequest) | [GetGoodPaymentResponse](#cloud.hashing.billing.v1.GetGoodPaymentResponse) |  |
| GetGoodPayments | [GetGoodPaymentsRequest](#cloud.hashing.billing.v1.GetGoodPaymentsRequest) | [GetGoodPaymentsResponse](#cloud.hashing.billing.v1.GetGoodPaymentsResponse) |  |
| GetGoodPaymentsByGood | [GetGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest) | [GetGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse) |  |
| GetIdleGoodPaymentsByGood | [GetIdleGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest) | [GetIdleGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse) |  |
| GetIdleGoodPaymentsByGoodPaymentCoin | [GetIdleGoodPaymentsByGoodPaymentCoinRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest) | [GetIdleGoodPaymentsByGoodPaymentCoinResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse) |  |
| GetGoodPaymentByAccount | [GetGoodPaymentByAccountRequest](#cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest) | [GetGoodPaymentByAccountResponse](#cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse) |  |
| CreateUserWithdraw | [CreateUserWithdrawRequest](#cloud.hashing.billing.v1.CreateUserWithdrawRequest) | [CreateUserWithdrawResponse](#cloud.hashing.billing.v1.CreateUserWithdrawResponse) |  |
| UpdateUserWithdraw | [UpdateUserWithdrawRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawRequest) | [UpdateUserWithdrawResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawResponse) |  |
| GetUserWithdraw | [GetUserWithdrawRequest](#cloud.hashing.billing.v1.GetUserWithdrawRequest) | [GetUserWithdrawResponse](#cloud.hashing.billing.v1.GetUserWithdrawResponse) |  |
| GetUserWithdrawsByAppUser | [GetUserWithdrawsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest) | [GetUserWithdrawsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse) |  |
| GetUserWithdrawsByAppUserCoin | [GetUserWithdrawsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest) | [GetUserWithdrawsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse) |  |
| GetUserWithdrawsByOtherAppUser | [GetUserWithdrawsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest) | [GetUserWithdrawsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse) |  |
| GetUserWithdrawAccount | [GetUserWithdrawAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountRequest) | [GetUserWithdrawAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountResponse) |  |
| GetUserWithdrawAccountsByAppUser | [GetUserWithdrawAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest) | [GetUserWithdrawAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse) |  |
| GetUserWithdrawAccountsByOtherAppUser | [GetUserWithdrawAccountsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest) | [GetUserWithdrawAccountsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse) |  |
| GetUserWithdrawByAccount | [GetUserWithdrawByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest) | [GetUserWithdrawByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse) |  |
| CreateUserDirectBenefit | [CreateUserDirectBenefitRequest](#cloud.hashing.billing.v1.CreateUserDirectBenefitRequest) | [CreateUserDirectBenefitResponse](#cloud.hashing.billing.v1.CreateUserDirectBenefitResponse) |  |
| UpdateUserDirectBenefit | [UpdateUserDirectBenefitRequest](#cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest) | [UpdateUserDirectBenefitResponse](#cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse) |  |
| GetUserDirectBenefit | [GetUserDirectBenefitRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitRequest) | [GetUserDirectBenefitResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitResponse) |  |
| GetUserDirectBenefitsByAppUser | [GetUserDirectBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest) | [GetUserDirectBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse) |  |
| GetUserDirectBenefitsByOtherAppUser | [GetUserDirectBenefitsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest) | [GetUserDirectBenefitsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse) |  |
| GetUserDirectBenefitByAccount | [GetUserDirectBenefitByAccountRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest) | [GetUserDirectBenefitByAccountResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse) |  |
| CreateUserWithdrawItem | [CreateUserWithdrawItemRequest](#cloud.hashing.billing.v1.CreateUserWithdrawItemRequest) | [CreateUserWithdrawItemResponse](#cloud.hashing.billing.v1.CreateUserWithdrawItemResponse) |  |
| UpdateUserWithdrawItem | [UpdateUserWithdrawItemRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest) | [UpdateUserWithdrawItemResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse) |  |
| GetUserWithdrawItem | [GetUserWithdrawItemRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemRequest) | [GetUserWithdrawItemResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemResponse) |  |
| GetUserWithdrawItemsByAccount | [GetUserWithdrawItemsByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest) | [GetUserWithdrawItemsByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse) |  |
| GetUserWithdrawItemsByAppUser | [GetUserWithdrawItemsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest) | [GetUserWithdrawItemsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse) |  |
| GetUserWithdrawItemsByAppUserWithdrawType | [GetUserWithdrawItemsByAppUserWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest) | [GetUserWithdrawItemsByAppUserWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse) |  |
| GetUserWithdrawItemsByAppUserCoinWithdrawType | [GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest) | [GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse) |  |
| GetUserWithdrawItemsByOtherAppUser | [GetUserWithdrawItemsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest) | [GetUserWithdrawItemsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse) |  |
| GetUserWithdrawItems | [GetUserWithdrawItemsRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsRequest) | [GetUserWithdrawItemsResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsResponse) |  |
| CreateAppWithdrawSetting | [CreateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest) | [CreateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse) |  |
| CreateAppWithdrawSettingForOtherApp | [CreateAppWithdrawSettingForOtherAppRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest) | [CreateAppWithdrawSettingForOtherAppResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse) |  |
| UpdateAppWithdrawSetting | [UpdateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest) | [UpdateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse) |  |
| GetAppWithdrawSetting | [GetAppWithdrawSettingRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingRequest) | [GetAppWithdrawSettingResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingResponse) |  |
| GetAppWithdrawSettingsByApp | [GetAppWithdrawSettingsByAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest) | [GetAppWithdrawSettingsByAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse) |  |
| GetAppWithdrawSettingByAppCoin | [GetAppWithdrawSettingByAppCoinRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest) | [GetAppWithdrawSettingByAppCoinResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse) |  |
| GetAppWithdrawSettingsByOtherApp | [GetAppWithdrawSettingsByOtherAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest) | [GetAppWithdrawSettingsByOtherAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse) |  |
| CreateUserPaymentBalance | [CreateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest) | [CreateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse) |  |
| CreateUserPaymentBalanceForOtherAppUser | [CreateUserPaymentBalanceForOtherAppUserRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest) | [CreateUserPaymentBalanceForOtherAppUserResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse) |  |
| UpdateUserPaymentBalance | [UpdateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest) | [UpdateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse) |  |
| GetUserPaymentBalance | [GetUserPaymentBalanceRequest](#cloud.hashing.billing.v1.GetUserPaymentBalanceRequest) | [GetUserPaymentBalanceResponse](#cloud.hashing.billing.v1.GetUserPaymentBalanceResponse) |  |
| GetUserPaymentBalancesByApp | [GetUserPaymentBalancesByAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest) | [GetUserPaymentBalancesByAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse) |  |
| GetUserPaymentBalancesByOtherApp | [GetUserPaymentBalancesByOtherAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest) | [GetUserPaymentBalancesByOtherAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse) |  |
| GetUserPaymentBalancesByAppUser | [GetUserPaymentBalancesByAppUserRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest) | [GetUserPaymentBalancesByAppUserResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse) |  |

 



<a name="npool/cloud-hashing-billing/cloud-hashing-billing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-billing/cloud-hashing-billing.proto



<a name="cloud.hashing.billing.v1.AppWithdrawSetting"></a>

### AppWithdrawSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawAutoReviewCoinAmount | [double](#double) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountInfo"></a>

### CoinAccountInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| PlatformHoldPrivateKey | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountTransaction"></a>

### CoinAccountTransaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FromAddressID | [string](#string) |  |  |
| ToAddressID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| FailHold | [bool](#bool) |  |  |






<a name="cloud.hashing.billing.v1.CoinAccountTransactionDetail"></a>

### CoinAccountTransactionDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FromAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| ToAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.CoinSetting"></a>

### CoinSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WarmAccountCoinAmount | [double](#double) |  |  |
| PaymentAccountCoinAmount | [double](#double) |  |  |
| PlatformOfflineAccountID | [string](#string) |  |  |
| UserOfflineAccountID | [string](#string) |  |  |
| UserOnlineAccountID | [string](#string) |  |  |
| GoodIncomingAccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest"></a>

### CreateAppWithdrawSettingForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse"></a>

### CreateAppWithdrawSettingForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest"></a>

### CreateAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse"></a>

### CreateAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountRequest"></a>

### CreateCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountResponse"></a>

### CreateCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest"></a>

### CreateCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse"></a>

### CreateCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinSettingRequest"></a>

### CreateCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateCoinSettingResponse"></a>

### CreateCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodBenefitRequest"></a>

### CreateGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodBenefitResponse"></a>

### CreateGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodPaymentRequest"></a>

### CreateGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.CreateGoodPaymentResponse"></a>

### CreateGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformBenefitRequest"></a>

### CreatePlatformBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformBenefitResponse"></a>

### CreatePlatformBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformSettingRequest"></a>

### CreatePlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreatePlatformSettingResponse"></a>

### CreatePlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserBenefitRequest"></a>

### CreateUserBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserBenefitResponse"></a>

### CreateUserBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserDirectBenefitRequest"></a>

### CreateUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserDirectBenefitResponse"></a>

### CreateUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest"></a>

### CreateUserPaymentBalanceForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse"></a>

### CreateUserPaymentBalanceForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest"></a>

### CreateUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse"></a>

### CreateUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawItemRequest"></a>

### CreateUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawItemResponse"></a>

### CreateUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawRequest"></a>

### CreateUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.CreateUserWithdrawResponse"></a>

### CreateUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountRequest"></a>

### DeleteCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountResponse"></a>

### DeleteCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest"></a>

### DeleteCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse"></a>

### DeleteCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest"></a>

### GetAppWithdrawSettingByAppCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse"></a>

### GetAppWithdrawSettingByAppCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingRequest"></a>

### GetAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingResponse"></a>

### GetAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest"></a>

### GetAppWithdrawSettingsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse"></a>

### GetAppWithdrawSettingsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) | repeated |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest"></a>

### GetAppWithdrawSettingsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse"></a>

### GetAppWithdrawSettingsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest"></a>

### GetCoinAccountByCoinAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinInfoID | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse"></a>

### GetCoinAccountByCoinAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountRequest"></a>

### GetCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountResponse"></a>

### GetCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest"></a>

### GetCoinAccountTransactionDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse"></a>

### GetCoinAccountTransactionDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [CoinAccountTransactionDetail](#cloud.hashing.billing.v1.CoinAccountTransactionDetail) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionRequest"></a>

### GetCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionResponse"></a>

### GetCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest"></a>

### GetCoinAccountTransactionsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse"></a>

### GetCoinAccountTransactionsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest"></a>

### GetCoinAccountTransactionsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse"></a>

### GetCoinAccountTransactionsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest"></a>

### GetCoinAccountTransactionsByCoinAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| AddressID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse"></a>

### GetCoinAccountTransactionsByCoinAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest"></a>

### GetCoinAccountTransactionsByCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse"></a>

### GetCoinAccountTransactionsByCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest"></a>

### GetCoinAccountTransactionsByGoodStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| State | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse"></a>

### GetCoinAccountTransactionsByGoodStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest"></a>

### GetCoinAccountTransactionsByStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse"></a>

### GetCoinAccountTransactionsByStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest"></a>

### GetCoinAccountTransactionsRequest







<a name="cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse"></a>

### GetCoinAccountTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountsRequest"></a>

### GetCoinAccountsRequest







<a name="cloud.hashing.billing.v1.GetCoinAccountsResponse"></a>

### GetCoinAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) | repeated |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingByCoinRequest"></a>

### GetCoinSettingByCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingByCoinResponse"></a>

### GetCoinSettingByCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingRequest"></a>

### GetCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinSettingResponse"></a>

### GetCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest"></a>

### GetGoodBenefitByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse"></a>

### GetGoodBenefitByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitRequest"></a>

### GetGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodBenefitResponse"></a>

### GetGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest"></a>

### GetGoodPaymentByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse"></a>

### GetGoodPaymentByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentRequest"></a>

### GetGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentResponse"></a>

### GetGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest"></a>

### GetGoodPaymentsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse"></a>

### GetGoodPaymentsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetGoodPaymentsRequest"></a>

### GetGoodPaymentsRequest







<a name="cloud.hashing.billing.v1.GetGoodPaymentsResponse"></a>

### GetGoodPaymentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest"></a>

### GetIdleGoodPaymentsByGoodPaymentCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse"></a>

### GetIdleGoodPaymentsByGoodPaymentCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest"></a>

### GetIdleGoodPaymentsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse"></a>

### GetIdleGoodPaymentsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) | repeated |  |






<a name="cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest"></a>

### GetLatestPlatformBenefitByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse"></a>

### GetLatestPlatformBenefitByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest"></a>

### GetLatestUserBenefitByGoodAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse"></a>

### GetLatestUserBenefitByGoodAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest"></a>

### GetPlatformBenefitDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse"></a>

### GetPlatformBenefitDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [PlatformBenefitDetail](#cloud.hashing.billing.v1.PlatformBenefitDetail) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitRequest"></a>

### GetPlatformBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitResponse"></a>

### GetPlatformBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest"></a>

### GetPlatformBenefitsByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse"></a>

### GetPlatformBenefitsByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetPlatformBenefitsRequest"></a>

### GetPlatformBenefitsRequest







<a name="cloud.hashing.billing.v1.GetPlatformBenefitsResponse"></a>

### GetPlatformBenefitsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest"></a>

### GetPlatformSettingByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse"></a>

### GetPlatformSettingByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingRequest"></a>

### GetPlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingResponse"></a>

### GetPlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppRequest"></a>

### GetUserBenefitsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppResponse"></a>

### GetUserBenefitsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest"></a>

### GetUserBenefitsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse"></a>

### GetUserBenefitsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest"></a>

### GetUserBenefitsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse"></a>

### GetUserBenefitsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsRequest"></a>

### GetUserBenefitsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsResponse"></a>

### GetUserBenefitsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBenefit](#cloud.hashing.billing.v1.UserBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest"></a>

### GetUserDirectBenefitByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse"></a>

### GetUserDirectBenefitByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitRequest"></a>

### GetUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitResponse"></a>

### GetUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest"></a>

### GetUserDirectBenefitsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse"></a>

### GetUserDirectBenefitsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest"></a>

### GetUserDirectBenefitsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse"></a>

### GetUserDirectBenefitsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalanceRequest"></a>

### GetUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalanceResponse"></a>

### GetUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest"></a>

### GetUserPaymentBalancesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse"></a>

### GetUserPaymentBalancesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest"></a>

### GetUserPaymentBalancesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse"></a>

### GetUserPaymentBalancesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest"></a>

### GetUserPaymentBalancesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse"></a>

### GetUserPaymentBalancesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountRequest"></a>

### GetUserWithdrawAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountResponse"></a>

### GetUserWithdrawAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest"></a>

### GetUserWithdrawAccountsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse"></a>

### GetUserWithdrawAccountsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest"></a>

### GetUserWithdrawAccountsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse"></a>

### GetUserWithdrawAccountsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawAccount](#cloud.hashing.billing.v1.UserWithdrawAccount) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest"></a>

### GetUserWithdrawByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse"></a>

### GetUserWithdrawByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemRequest"></a>

### GetUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemResponse"></a>

### GetUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest"></a>

### GetUserWithdrawItemsByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse"></a>

### GetUserWithdrawItemsByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest"></a>

### GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse"></a>

### GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest"></a>

### GetUserWithdrawItemsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse"></a>

### GetUserWithdrawItemsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest"></a>

### GetUserWithdrawItemsByAppUserWithdrawTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse"></a>

### GetUserWithdrawItemsByAppUserWithdrawTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest"></a>

### GetUserWithdrawItemsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse"></a>

### GetUserWithdrawItemsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsRequest"></a>

### GetUserWithdrawItemsRequest







<a name="cloud.hashing.billing.v1.GetUserWithdrawItemsResponse"></a>

### GetUserWithdrawItemsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawRequest"></a>

### GetUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawResponse"></a>

### GetUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest"></a>

### GetUserWithdrawsByAppUserCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse"></a>

### GetUserWithdrawsByAppUserCoinResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest"></a>

### GetUserWithdrawsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse"></a>

### GetUserWithdrawsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest"></a>

### GetUserWithdrawsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse"></a>

### GetUserWithdrawsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) | repeated |  |






<a name="cloud.hashing.billing.v1.GoodBenefit"></a>

### GoodBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAccountID | [string](#string) |  |  |
| BenefitIntervalHours | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.GoodPayment"></a>

### GoodPayment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| PaymentCoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| Idle | [bool](#bool) |  |  |
| OccupiedBy | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.PlatformBenefit"></a>

### PlatformBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAccountID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| LastBenefitTimestamp | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.PlatformBenefitDetail"></a>

### PlatformBenefitDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| ChainTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.PlatformSetting"></a>

### PlatformSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| WarmAccountUSDAmount | [double](#double) |  |  |
| PaymentAccountUSDAmount | [double](#double) |  |  |
| WithdrawAutoReviewUSDAmount | [double](#double) |  |  |






<a name="cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest"></a>

### UpdateAppWithdrawSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse"></a>

### UpdateAppWithdrawSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppWithdrawSetting](#cloud.hashing.billing.v1.AppWithdrawSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest"></a>

### UpdateCoinAccountTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse"></a>

### UpdateCoinAccountTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinSettingRequest"></a>

### UpdateCoinSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateCoinSettingResponse"></a>

### UpdateCoinSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinSetting](#cloud.hashing.billing.v1.CoinSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodBenefitRequest"></a>

### UpdateGoodBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodBenefitResponse"></a>

### UpdateGoodBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodBenefit](#cloud.hashing.billing.v1.GoodBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodPaymentRequest"></a>

### UpdateGoodPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.UpdateGoodPaymentResponse"></a>

### UpdateGoodPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPayment](#cloud.hashing.billing.v1.GoodPayment) |  |  |






<a name="cloud.hashing.billing.v1.UpdatePlatformSettingRequest"></a>

### UpdatePlatformSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdatePlatformSettingResponse"></a>

### UpdatePlatformSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest"></a>

### UpdateUserDirectBenefitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse"></a>

### UpdateUserDirectBenefitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDirectBenefit](#cloud.hashing.billing.v1.UserDirectBenefit) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest"></a>

### UpdateUserPaymentBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse"></a>

### UpdateUserPaymentBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserPaymentBalance](#cloud.hashing.billing.v1.UserPaymentBalance) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest"></a>

### UpdateUserWithdrawItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse"></a>

### UpdateUserWithdrawItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdrawItem](#cloud.hashing.billing.v1.UserWithdrawItem) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawRequest"></a>

### UpdateUserWithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.UpdateUserWithdrawResponse"></a>

### UpdateUserWithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |






<a name="cloud.hashing.billing.v1.UserBenefit"></a>

### UserBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| OrderID | [string](#string) |  |  |
| LastBenefitTimestamp | [uint32](#uint32) |  |  |
| PlatformTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserDirectBenefit"></a>

### UserDirectBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserPaymentBalance"></a>

### UserPaymentBalance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| UsedByPaymentID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.UserWithdraw"></a>

### UserWithdraw



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| Labels | [string](#string) | repeated |  |






<a name="cloud.hashing.billing.v1.UserWithdrawAccount"></a>

### UserWithdrawAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Withdraw | [UserWithdraw](#cloud.hashing.billing.v1.UserWithdraw) |  |  |
| Account | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |






<a name="cloud.hashing.billing.v1.UserWithdrawItem"></a>

### UserWithdrawItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| WithdrawToAccountID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| PlatformTransactionID | [string](#string) |  |  |
| WithdrawType | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.billing.v1.CloudHashingBilling"></a>

### CloudHashingBilling
Cloud Hashing Billing

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| CreateCoinAccount | [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest) | [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse) |  |
| GetCoinAccount | [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest) | [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse) |  |
| GetCoinAccountByCoinAddress | [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest) | [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse) |  |
| GetCoinAccounts | [GetCoinAccountsRequest](#cloud.hashing.billing.v1.GetCoinAccountsRequest) | [GetCoinAccountsResponse](#cloud.hashing.billing.v1.GetCoinAccountsResponse) |  |
| DeleteCoinAccount | [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest) | [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse) |  |
| CreateCoinAccountTransaction | [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest) | [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransaction | [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest) | [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionsByCoinAccount | [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest) | [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse) |  |
| GetCoinAccountTransactionsByState | [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest) | [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse) |  |
| GetCoinAccountTransactionsByGoodState | [GetCoinAccountTransactionsByGoodStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateRequest) | [GetCoinAccountTransactionsByGoodStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByGoodStateResponse) |  |
| GetCoinAccountTransactions | [GetCoinAccountTransactionsRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsRequest) | [GetCoinAccountTransactionsResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsResponse) |  |
| GetCoinAccountTransactionsByAppUser | [GetCoinAccountTransactionsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserRequest) | [GetCoinAccountTransactionsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserResponse) |  |
| GetCoinAccountTransactionsByAppUserCoin | [GetCoinAccountTransactionsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinRequest) | [GetCoinAccountTransactionsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByAppUserCoinResponse) |  |
| GetCoinAccountTransactionsByCoin | [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest) | [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse) |  |
| UpdateCoinAccountTransaction | [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest) | [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionDetail | [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest) | [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse) |  |
| DeleteCoinAccountTransaction | [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest) | [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse) |  |
| CreatePlatformBenefit | [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest) | [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse) |  |
| GetLatestPlatformBenefitByGood | [GetLatestPlatformBenefitByGoodRequest](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest) | [GetLatestPlatformBenefitByGoodResponse](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse) |  |
| GetPlatformBenefitsByGood | [GetPlatformBenefitsByGoodRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest) | [GetPlatformBenefitsByGoodResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse) |  |
| GetPlatformBenefits | [GetPlatformBenefitsRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsRequest) | [GetPlatformBenefitsResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsResponse) |  |
| GetPlatformBenefit | [GetPlatformBenefitRequest](#cloud.hashing.billing.v1.GetPlatformBenefitRequest) | [GetPlatformBenefitResponse](#cloud.hashing.billing.v1.GetPlatformBenefitResponse) |  |
| GetPlatformBenefitDetail | [GetPlatformBenefitDetailRequest](#cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest) | [GetPlatformBenefitDetailResponse](#cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse) |  |
| CreatePlatformSetting | [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest) | [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse) |  |
| UpdatePlatformSetting | [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest) | [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse) |  |
| GetPlatformSetting | [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest) | [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse) |  |
| CreateUserBenefit | [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest) | [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse) |  |
| GetUserBenefitsByAppUser | [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest) | [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse) |  |
| GetUserBenefitsByAppUserCoin | [GetUserBenefitsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinRequest) | [GetUserBenefitsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserCoinResponse) |  |
| GetUserBenefitsByApp | [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest) | [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse) |  |
| GetUserBenefits | [GetUserBenefitsRequest](#cloud.hashing.billing.v1.GetUserBenefitsRequest) | [GetUserBenefitsResponse](#cloud.hashing.billing.v1.GetUserBenefitsResponse) |  |
| GetLatestUserBenefitByGoodAppUser | [GetLatestUserBenefitByGoodAppUserRequest](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest) | [GetLatestUserBenefitByGoodAppUserResponse](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse) |  |
| CreateCoinSetting | [CreateCoinSettingRequest](#cloud.hashing.billing.v1.CreateCoinSettingRequest) | [CreateCoinSettingResponse](#cloud.hashing.billing.v1.CreateCoinSettingResponse) |  |
| UpdateCoinSetting | [UpdateCoinSettingRequest](#cloud.hashing.billing.v1.UpdateCoinSettingRequest) | [UpdateCoinSettingResponse](#cloud.hashing.billing.v1.UpdateCoinSettingResponse) |  |
| GetCoinSetting | [GetCoinSettingRequest](#cloud.hashing.billing.v1.GetCoinSettingRequest) | [GetCoinSettingResponse](#cloud.hashing.billing.v1.GetCoinSettingResponse) |  |
| GetCoinSettingByCoin | [GetCoinSettingByCoinRequest](#cloud.hashing.billing.v1.GetCoinSettingByCoinRequest) | [GetCoinSettingByCoinResponse](#cloud.hashing.billing.v1.GetCoinSettingByCoinResponse) |  |
| CreateGoodBenefit | [CreateGoodBenefitRequest](#cloud.hashing.billing.v1.CreateGoodBenefitRequest) | [CreateGoodBenefitResponse](#cloud.hashing.billing.v1.CreateGoodBenefitResponse) |  |
| UpdateGoodBenefit | [UpdateGoodBenefitRequest](#cloud.hashing.billing.v1.UpdateGoodBenefitRequest) | [UpdateGoodBenefitResponse](#cloud.hashing.billing.v1.UpdateGoodBenefitResponse) |  |
| GetGoodBenefit | [GetGoodBenefitRequest](#cloud.hashing.billing.v1.GetGoodBenefitRequest) | [GetGoodBenefitResponse](#cloud.hashing.billing.v1.GetGoodBenefitResponse) |  |
| GetGoodBenefitByGood | [GetGoodBenefitByGoodRequest](#cloud.hashing.billing.v1.GetGoodBenefitByGoodRequest) | [GetGoodBenefitByGoodResponse](#cloud.hashing.billing.v1.GetGoodBenefitByGoodResponse) |  |
| CreateGoodPayment | [CreateGoodPaymentRequest](#cloud.hashing.billing.v1.CreateGoodPaymentRequest) | [CreateGoodPaymentResponse](#cloud.hashing.billing.v1.CreateGoodPaymentResponse) |  |
| UpdateGoodPayment | [UpdateGoodPaymentRequest](#cloud.hashing.billing.v1.UpdateGoodPaymentRequest) | [UpdateGoodPaymentResponse](#cloud.hashing.billing.v1.UpdateGoodPaymentResponse) |  |
| GetGoodPayment | [GetGoodPaymentRequest](#cloud.hashing.billing.v1.GetGoodPaymentRequest) | [GetGoodPaymentResponse](#cloud.hashing.billing.v1.GetGoodPaymentResponse) |  |
| GetGoodPayments | [GetGoodPaymentsRequest](#cloud.hashing.billing.v1.GetGoodPaymentsRequest) | [GetGoodPaymentsResponse](#cloud.hashing.billing.v1.GetGoodPaymentsResponse) |  |
| GetGoodPaymentsByGood | [GetGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodRequest) | [GetGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetGoodPaymentsByGoodResponse) |  |
| GetIdleGoodPaymentsByGood | [GetIdleGoodPaymentsByGoodRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodRequest) | [GetIdleGoodPaymentsByGoodResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodResponse) |  |
| GetIdleGoodPaymentsByGoodPaymentCoin | [GetIdleGoodPaymentsByGoodPaymentCoinRequest](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinRequest) | [GetIdleGoodPaymentsByGoodPaymentCoinResponse](#cloud.hashing.billing.v1.GetIdleGoodPaymentsByGoodPaymentCoinResponse) |  |
| GetGoodPaymentByAccount | [GetGoodPaymentByAccountRequest](#cloud.hashing.billing.v1.GetGoodPaymentByAccountRequest) | [GetGoodPaymentByAccountResponse](#cloud.hashing.billing.v1.GetGoodPaymentByAccountResponse) |  |
| CreateUserWithdraw | [CreateUserWithdrawRequest](#cloud.hashing.billing.v1.CreateUserWithdrawRequest) | [CreateUserWithdrawResponse](#cloud.hashing.billing.v1.CreateUserWithdrawResponse) |  |
| UpdateUserWithdraw | [UpdateUserWithdrawRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawRequest) | [UpdateUserWithdrawResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawResponse) |  |
| GetUserWithdraw | [GetUserWithdrawRequest](#cloud.hashing.billing.v1.GetUserWithdrawRequest) | [GetUserWithdrawResponse](#cloud.hashing.billing.v1.GetUserWithdrawResponse) |  |
| GetUserWithdrawsByAppUser | [GetUserWithdrawsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserRequest) | [GetUserWithdrawsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserResponse) |  |
| GetUserWithdrawsByAppUserCoin | [GetUserWithdrawsByAppUserCoinRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinRequest) | [GetUserWithdrawsByAppUserCoinResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByAppUserCoinResponse) |  |
| GetUserWithdrawsByOtherAppUser | [GetUserWithdrawsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserRequest) | [GetUserWithdrawsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawsByOtherAppUserResponse) |  |
| GetUserWithdrawAccount | [GetUserWithdrawAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountRequest) | [GetUserWithdrawAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountResponse) |  |
| GetUserWithdrawAccountsByAppUser | [GetUserWithdrawAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserRequest) | [GetUserWithdrawAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByAppUserResponse) |  |
| GetUserWithdrawAccountsByOtherAppUser | [GetUserWithdrawAccountsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserRequest) | [GetUserWithdrawAccountsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawAccountsByOtherAppUserResponse) |  |
| GetUserWithdrawByAccount | [GetUserWithdrawByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawByAccountRequest) | [GetUserWithdrawByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawByAccountResponse) |  |
| CreateUserDirectBenefit | [CreateUserDirectBenefitRequest](#cloud.hashing.billing.v1.CreateUserDirectBenefitRequest) | [CreateUserDirectBenefitResponse](#cloud.hashing.billing.v1.CreateUserDirectBenefitResponse) |  |
| UpdateUserDirectBenefit | [UpdateUserDirectBenefitRequest](#cloud.hashing.billing.v1.UpdateUserDirectBenefitRequest) | [UpdateUserDirectBenefitResponse](#cloud.hashing.billing.v1.UpdateUserDirectBenefitResponse) |  |
| GetUserDirectBenefit | [GetUserDirectBenefitRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitRequest) | [GetUserDirectBenefitResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitResponse) |  |
| GetUserDirectBenefitsByAppUser | [GetUserDirectBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserRequest) | [GetUserDirectBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByAppUserResponse) |  |
| GetUserDirectBenefitsByOtherAppUser | [GetUserDirectBenefitsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserRequest) | [GetUserDirectBenefitsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitsByOtherAppUserResponse) |  |
| GetUserDirectBenefitByAccount | [GetUserDirectBenefitByAccountRequest](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountRequest) | [GetUserDirectBenefitByAccountResponse](#cloud.hashing.billing.v1.GetUserDirectBenefitByAccountResponse) |  |
| CreateUserWithdrawItem | [CreateUserWithdrawItemRequest](#cloud.hashing.billing.v1.CreateUserWithdrawItemRequest) | [CreateUserWithdrawItemResponse](#cloud.hashing.billing.v1.CreateUserWithdrawItemResponse) |  |
| UpdateUserWithdrawItem | [UpdateUserWithdrawItemRequest](#cloud.hashing.billing.v1.UpdateUserWithdrawItemRequest) | [UpdateUserWithdrawItemResponse](#cloud.hashing.billing.v1.UpdateUserWithdrawItemResponse) |  |
| GetUserWithdrawItem | [GetUserWithdrawItemRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemRequest) | [GetUserWithdrawItemResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemResponse) |  |
| GetUserWithdrawItemsByAccount | [GetUserWithdrawItemsByAccountRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountRequest) | [GetUserWithdrawItemsByAccountResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAccountResponse) |  |
| GetUserWithdrawItemsByAppUser | [GetUserWithdrawItemsByAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserRequest) | [GetUserWithdrawItemsByAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserResponse) |  |
| GetUserWithdrawItemsByAppUserWithdrawType | [GetUserWithdrawItemsByAppUserWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeRequest) | [GetUserWithdrawItemsByAppUserWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserWithdrawTypeResponse) |  |
| GetUserWithdrawItemsByAppUserCoinWithdrawType | [GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest) | [GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse) |  |
| GetUserWithdrawItemsByOtherAppUser | [GetUserWithdrawItemsByOtherAppUserRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserRequest) | [GetUserWithdrawItemsByOtherAppUserResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsByOtherAppUserResponse) |  |
| GetUserWithdrawItems | [GetUserWithdrawItemsRequest](#cloud.hashing.billing.v1.GetUserWithdrawItemsRequest) | [GetUserWithdrawItemsResponse](#cloud.hashing.billing.v1.GetUserWithdrawItemsResponse) |  |
| CreateAppWithdrawSetting | [CreateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingRequest) | [CreateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingResponse) |  |
| CreateAppWithdrawSettingForOtherApp | [CreateAppWithdrawSettingForOtherAppRequest](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppRequest) | [CreateAppWithdrawSettingForOtherAppResponse](#cloud.hashing.billing.v1.CreateAppWithdrawSettingForOtherAppResponse) |  |
| UpdateAppWithdrawSetting | [UpdateAppWithdrawSettingRequest](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingRequest) | [UpdateAppWithdrawSettingResponse](#cloud.hashing.billing.v1.UpdateAppWithdrawSettingResponse) |  |
| GetAppWithdrawSetting | [GetAppWithdrawSettingRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingRequest) | [GetAppWithdrawSettingResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingResponse) |  |
| GetAppWithdrawSettingsByApp | [GetAppWithdrawSettingsByAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppRequest) | [GetAppWithdrawSettingsByAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByAppResponse) |  |
| GetAppWithdrawSettingByAppCoin | [GetAppWithdrawSettingByAppCoinRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinRequest) | [GetAppWithdrawSettingByAppCoinResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingByAppCoinResponse) |  |
| GetAppWithdrawSettingsByOtherApp | [GetAppWithdrawSettingsByOtherAppRequest](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppRequest) | [GetAppWithdrawSettingsByOtherAppResponse](#cloud.hashing.billing.v1.GetAppWithdrawSettingsByOtherAppResponse) |  |
| CreateUserPaymentBalance | [CreateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceRequest) | [CreateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceResponse) |  |
| CreateUserPaymentBalanceForOtherAppUser | [CreateUserPaymentBalanceForOtherAppUserRequest](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserRequest) | [CreateUserPaymentBalanceForOtherAppUserResponse](#cloud.hashing.billing.v1.CreateUserPaymentBalanceForOtherAppUserResponse) |  |
| UpdateUserPaymentBalance | [UpdateUserPaymentBalanceRequest](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceRequest) | [UpdateUserPaymentBalanceResponse](#cloud.hashing.billing.v1.UpdateUserPaymentBalanceResponse) |  |
| GetUserPaymentBalance | [GetUserPaymentBalanceRequest](#cloud.hashing.billing.v1.GetUserPaymentBalanceRequest) | [GetUserPaymentBalanceResponse](#cloud.hashing.billing.v1.GetUserPaymentBalanceResponse) |  |
| GetUserPaymentBalancesByApp | [GetUserPaymentBalancesByAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppRequest) | [GetUserPaymentBalancesByAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppResponse) |  |
| GetUserPaymentBalancesByOtherApp | [GetUserPaymentBalancesByOtherAppRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppRequest) | [GetUserPaymentBalancesByOtherAppResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByOtherAppResponse) |  |
| GetUserPaymentBalancesByAppUser | [GetUserPaymentBalancesByAppUserRequest](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserRequest) | [GetUserPaymentBalancesByAppUserResponse](#cloud.hashing.billing.v1.GetUserPaymentBalancesByAppUserResponse) |  |

 



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

