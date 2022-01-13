# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-billing.proto](#npool/cloud-hashing-billing.proto)
    - [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo)
    - [CoinAccountTransaction](#cloud.hashing.billing.v1.CoinAccountTransaction)
    - [CoinAccountTransactionDetail](#cloud.hashing.billing.v1.CoinAccountTransactionDetail)
    - [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest)
    - [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse)
    - [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest)
    - [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse)
    - [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest)
    - [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse)
    - [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest)
    - [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse)
    - [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest)
    - [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse)
    - [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest)
    - [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse)
    - [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest)
    - [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse)
    - [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest)
    - [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse)
    - [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest)
    - [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse)
    - [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest)
    - [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse)
    - [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest)
    - [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse)
    - [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest)
    - [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse)
    - [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest)
    - [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse)
    - [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest)
    - [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse)
    - [GetCoinAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountsByAppUserRequest)
    - [GetCoinAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountsByAppUserResponse)
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
    - [GetPlatformSettingByGoodRequest](#cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest)
    - [GetPlatformSettingByGoodResponse](#cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse)
    - [GetPlatformSettingDetailRequest](#cloud.hashing.billing.v1.GetPlatformSettingDetailRequest)
    - [GetPlatformSettingDetailResponse](#cloud.hashing.billing.v1.GetPlatformSettingDetailResponse)
    - [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest)
    - [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse)
    - [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest)
    - [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse)
    - [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest)
    - [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse)
    - [PageInfo](#cloud.hashing.billing.v1.PageInfo)
    - [PlatformBenefit](#cloud.hashing.billing.v1.PlatformBenefit)
    - [PlatformBenefitDetail](#cloud.hashing.billing.v1.PlatformBenefitDetail)
    - [PlatformSetting](#cloud.hashing.billing.v1.PlatformSetting)
    - [PlatformSettingDetail](#cloud.hashing.billing.v1.PlatformSettingDetail)
    - [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest)
    - [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse)
    - [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest)
    - [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse)
    - [UserBenefit](#cloud.hashing.billing.v1.UserBenefit)
    - [VersionResponse](#cloud.hashing.billing.v1.VersionResponse)
  
    - [CloudHashingBilling](#cloud.hashing.billing.v1.CloudHashingBilling)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-billing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-billing.proto



<a name="cloud.hashing.billing.v1.CoinAccountInfo"></a>

### CoinAccountInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| GeneratedBy | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Idle | [bool](#bool) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






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
| PlatformTransactionID | [string](#string) |  |  |






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
| PlatformTransactionID | [string](#string) |  |  |






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






<a name="cloud.hashing.billing.v1.GetCoinAccountsByAppUserRequest"></a>

### GetCoinAccountsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetCoinAccountsByAppUserResponse"></a>

### GetCoinAccountsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) | repeated |  |






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






<a name="cloud.hashing.billing.v1.GetPlatformSettingDetailRequest"></a>

### GetPlatformSettingDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.billing.v1.GetPlatformSettingDetailResponse"></a>

### GetPlatformSettingDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [PlatformSettingDetail](#cloud.hashing.billing.v1.PlatformSettingDetail) |  |  |






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
| PageInfo | [PageInfo](#cloud.hashing.billing.v1.PageInfo) |  |  |






<a name="cloud.hashing.billing.v1.GetUserBenefitsByAppResponse"></a>

### GetUserBenefitsByAppResponse



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






<a name="cloud.hashing.billing.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageSize | [int32](#int32) |  |  |






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
| GoodID | [string](#string) |  |  |
| BenefitAccountID | [string](#string) |  |  |
| PlatformOfflineAccountID | [string](#string) |  |  |
| UserOnlineAccountID | [string](#string) |  |  |
| UserOfflineAccountID | [string](#string) |  |  |
| BenefitIntervalHours | [int32](#int32) |  |  |






<a name="cloud.hashing.billing.v1.PlatformSettingDetail"></a>

### PlatformSettingDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| BenefitAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| PlatformOfflineAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| UserOnlineAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| UserOfflineAddress | [CoinAccountInfo](#cloud.hashing.billing.v1.CoinAccountInfo) |  |  |
| BenefitIntervalHours | [int32](#int32) |  |  |






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






<a name="cloud.hashing.billing.v1.UserBenefit"></a>

### UserBenefit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| OrderID | [string](#string) |  |  |
| LastBenefitTimestamp | [uint32](#uint32) |  |  |






<a name="cloud.hashing.billing.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.billing.v1.CloudHashingBilling"></a>

### CloudHashingBilling
Cloud Hashing Billing

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.billing.v1.VersionResponse) | Method Version |
| CreateCoinAccount | [CreateCoinAccountRequest](#cloud.hashing.billing.v1.CreateCoinAccountRequest) | [CreateCoinAccountResponse](#cloud.hashing.billing.v1.CreateCoinAccountResponse) |  |
| GetCoinAccount | [GetCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountRequest) | [GetCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountResponse) |  |
| GetCoinAccountByCoinAddress | [GetCoinAccountByCoinAddressRequest](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressRequest) | [GetCoinAccountByCoinAddressResponse](#cloud.hashing.billing.v1.GetCoinAccountByCoinAddressResponse) |  |
| GetCoinAccountsByAppUser | [GetCoinAccountsByAppUserRequest](#cloud.hashing.billing.v1.GetCoinAccountsByAppUserRequest) | [GetCoinAccountsByAppUserResponse](#cloud.hashing.billing.v1.GetCoinAccountsByAppUserResponse) |  |
| DeleteCoinAccount | [DeleteCoinAccountRequest](#cloud.hashing.billing.v1.DeleteCoinAccountRequest) | [DeleteCoinAccountResponse](#cloud.hashing.billing.v1.DeleteCoinAccountResponse) |  |
| CreateCoinAccountTransaction | [CreateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.CreateCoinAccountTransactionRequest) | [CreateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.CreateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransaction | [GetCoinAccountTransactionRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionRequest) | [GetCoinAccountTransactionResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionsByCoinAccount | [GetCoinAccountTransactionsByCoinAccountRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountRequest) | [GetCoinAccountTransactionsByCoinAccountResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinAccountResponse) |  |
| GetCoinAccountTransactionsByState | [GetCoinAccountTransactionsByStateRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateRequest) | [GetCoinAccountTransactionsByStateResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByStateResponse) |  |
| GetCoinAccountTransactionsByCoin | [GetCoinAccountTransactionsByCoinRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinRequest) | [GetCoinAccountTransactionsByCoinResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionsByCoinResponse) |  |
| UpdateCoinAccountTransaction | [UpdateCoinAccountTransactionRequest](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionRequest) | [UpdateCoinAccountTransactionResponse](#cloud.hashing.billing.v1.UpdateCoinAccountTransactionResponse) |  |
| GetCoinAccountTransactionDetail | [GetCoinAccountTransactionDetailRequest](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailRequest) | [GetCoinAccountTransactionDetailResponse](#cloud.hashing.billing.v1.GetCoinAccountTransactionDetailResponse) |  |
| DeleteCoinAccountTransaction | [DeleteCoinAccountTransactionRequest](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionRequest) | [DeleteCoinAccountTransactionResponse](#cloud.hashing.billing.v1.DeleteCoinAccountTransactionResponse) |  |
| CreatePlatformBenefit | [CreatePlatformBenefitRequest](#cloud.hashing.billing.v1.CreatePlatformBenefitRequest) | [CreatePlatformBenefitResponse](#cloud.hashing.billing.v1.CreatePlatformBenefitResponse) |  |
| GetLatestPlatformBenefitByGood | [GetLatestPlatformBenefitByGoodRequest](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodRequest) | [GetLatestPlatformBenefitByGoodResponse](#cloud.hashing.billing.v1.GetLatestPlatformBenefitByGoodResponse) |  |
| GetPlatformBenefitsByGood | [GetPlatformBenefitsByGoodRequest](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodRequest) | [GetPlatformBenefitsByGoodResponse](#cloud.hashing.billing.v1.GetPlatformBenefitsByGoodResponse) |  |
| GetPlatformBenefit | [GetPlatformBenefitRequest](#cloud.hashing.billing.v1.GetPlatformBenefitRequest) | [GetPlatformBenefitResponse](#cloud.hashing.billing.v1.GetPlatformBenefitResponse) |  |
| GetPlatformBenefitDetail | [GetPlatformBenefitDetailRequest](#cloud.hashing.billing.v1.GetPlatformBenefitDetailRequest) | [GetPlatformBenefitDetailResponse](#cloud.hashing.billing.v1.GetPlatformBenefitDetailResponse) |  |
| CreatePlatformSetting | [CreatePlatformSettingRequest](#cloud.hashing.billing.v1.CreatePlatformSettingRequest) | [CreatePlatformSettingResponse](#cloud.hashing.billing.v1.CreatePlatformSettingResponse) |  |
| UpdatePlatformSetting | [UpdatePlatformSettingRequest](#cloud.hashing.billing.v1.UpdatePlatformSettingRequest) | [UpdatePlatformSettingResponse](#cloud.hashing.billing.v1.UpdatePlatformSettingResponse) |  |
| GetPlatformSettingByGood | [GetPlatformSettingByGoodRequest](#cloud.hashing.billing.v1.GetPlatformSettingByGoodRequest) | [GetPlatformSettingByGoodResponse](#cloud.hashing.billing.v1.GetPlatformSettingByGoodResponse) |  |
| GetPlatformSettingDetail | [GetPlatformSettingDetailRequest](#cloud.hashing.billing.v1.GetPlatformSettingDetailRequest) | [GetPlatformSettingDetailResponse](#cloud.hashing.billing.v1.GetPlatformSettingDetailResponse) |  |
| GetPlatformSetting | [GetPlatformSettingRequest](#cloud.hashing.billing.v1.GetPlatformSettingRequest) | [GetPlatformSettingResponse](#cloud.hashing.billing.v1.GetPlatformSettingResponse) |  |
| CreateUserBenefit | [CreateUserBenefitRequest](#cloud.hashing.billing.v1.CreateUserBenefitRequest) | [CreateUserBenefitResponse](#cloud.hashing.billing.v1.CreateUserBenefitResponse) |  |
| GetUserBenefitsByAppUser | [GetUserBenefitsByAppUserRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserRequest) | [GetUserBenefitsByAppUserResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppUserResponse) |  |
| GetUserBenefitsByApp | [GetUserBenefitsByAppRequest](#cloud.hashing.billing.v1.GetUserBenefitsByAppRequest) | [GetUserBenefitsByAppResponse](#cloud.hashing.billing.v1.GetUserBenefitsByAppResponse) |  |
| GetLatestUserBenefitByGoodAppUser | [GetLatestUserBenefitByGoodAppUserRequest](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserRequest) | [GetLatestUserBenefitByGoodAppUserResponse](#cloud.hashing.billing.v1.GetLatestUserBenefitByGoodAppUserResponse) |  |

 



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

