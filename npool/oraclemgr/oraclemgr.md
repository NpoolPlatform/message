# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/oraclemgr/oraclemgr.proto](#npool_oraclemgr_oraclemgr-proto)
    - [CountRewardsRequest](#oracle-manager-v1-CountRewardsRequest)
    - [CountRewardsRequest.CondsEntry](#oracle-manager-v1-CountRewardsRequest-CondsEntry)
    - [CountRewardsResponse](#oracle-manager-v1-CountRewardsResponse)
    - [CreateAppCurrencyRequest](#oracle-manager-v1-CreateAppCurrencyRequest)
    - [CreateAppCurrencyResponse](#oracle-manager-v1-CreateAppCurrencyResponse)
    - [CreateCurrencyRequest](#oracle-manager-v1-CreateCurrencyRequest)
    - [CreateCurrencyResponse](#oracle-manager-v1-CreateCurrencyResponse)
    - [CreateRewardRequest](#oracle-manager-v1-CreateRewardRequest)
    - [CreateRewardResponse](#oracle-manager-v1-CreateRewardResponse)
    - [CreateRewardsRequest](#oracle-manager-v1-CreateRewardsRequest)
    - [CreateRewardsResponse](#oracle-manager-v1-CreateRewardsResponse)
    - [Currency](#oracle-manager-v1-Currency)
    - [DeleteCurrencyRequest](#oracle-manager-v1-DeleteCurrencyRequest)
    - [DeleteCurrencyResponse](#oracle-manager-v1-DeleteCurrencyResponse)
    - [DeleteRewardRequest](#oracle-manager-v1-DeleteRewardRequest)
    - [DeleteRewardResponse](#oracle-manager-v1-DeleteRewardResponse)
    - [ExistRewardCondsRequest](#oracle-manager-v1-ExistRewardCondsRequest)
    - [ExistRewardCondsRequest.CondsEntry](#oracle-manager-v1-ExistRewardCondsRequest-CondsEntry)
    - [ExistRewardCondsResponse](#oracle-manager-v1-ExistRewardCondsResponse)
    - [ExistRewardRequest](#oracle-manager-v1-ExistRewardRequest)
    - [ExistRewardResponse](#oracle-manager-v1-ExistRewardResponse)
    - [GetAppCurrenciesRequest](#oracle-manager-v1-GetAppCurrenciesRequest)
    - [GetAppCurrenciesResponse](#oracle-manager-v1-GetAppCurrenciesResponse)
    - [GetCurrenciesRequest](#oracle-manager-v1-GetCurrenciesRequest)
    - [GetCurrenciesResponse](#oracle-manager-v1-GetCurrenciesResponse)
    - [GetCurrencyOnlyRequest](#oracle-manager-v1-GetCurrencyOnlyRequest)
    - [GetCurrencyOnlyRequest.CondsEntry](#oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry)
    - [GetCurrencyOnlyResponse](#oracle-manager-v1-GetCurrencyOnlyResponse)
    - [GetCurrencyRequest](#oracle-manager-v1-GetCurrencyRequest)
    - [GetCurrencyResponse](#oracle-manager-v1-GetCurrencyResponse)
    - [GetRewardOnlyRequest](#oracle-manager-v1-GetRewardOnlyRequest)
    - [GetRewardOnlyRequest.CondsEntry](#oracle-manager-v1-GetRewardOnlyRequest-CondsEntry)
    - [GetRewardOnlyResponse](#oracle-manager-v1-GetRewardOnlyResponse)
    - [GetRewardRequest](#oracle-manager-v1-GetRewardRequest)
    - [GetRewardResponse](#oracle-manager-v1-GetRewardResponse)
    - [GetRewardsRequest](#oracle-manager-v1-GetRewardsRequest)
    - [GetRewardsRequest.CondsEntry](#oracle-manager-v1-GetRewardsRequest-CondsEntry)
    - [GetRewardsResponse](#oracle-manager-v1-GetRewardsResponse)
    - [Reward](#oracle-manager-v1-Reward)
    - [UpdateCurrencyRequest](#oracle-manager-v1-UpdateCurrencyRequest)
    - [UpdateCurrencyResponse](#oracle-manager-v1-UpdateCurrencyResponse)
    - [UpdateRewardRequest](#oracle-manager-v1-UpdateRewardRequest)
    - [UpdateRewardResponse](#oracle-manager-v1-UpdateRewardResponse)
  
    - [OracleManager](#oracle-manager-v1-OracleManager)
  
- [npool/oraclemgr/oraclemgr.proto](#npool_oraclemgr_oraclemgr-proto)
    - [CountRewardsRequest](#oracle-manager-v1-CountRewardsRequest)
    - [CountRewardsRequest.CondsEntry](#oracle-manager-v1-CountRewardsRequest-CondsEntry)
    - [CountRewardsResponse](#oracle-manager-v1-CountRewardsResponse)
    - [CreateAppCurrencyRequest](#oracle-manager-v1-CreateAppCurrencyRequest)
    - [CreateAppCurrencyResponse](#oracle-manager-v1-CreateAppCurrencyResponse)
    - [CreateCurrencyRequest](#oracle-manager-v1-CreateCurrencyRequest)
    - [CreateCurrencyResponse](#oracle-manager-v1-CreateCurrencyResponse)
    - [CreateRewardRequest](#oracle-manager-v1-CreateRewardRequest)
    - [CreateRewardResponse](#oracle-manager-v1-CreateRewardResponse)
    - [CreateRewardsRequest](#oracle-manager-v1-CreateRewardsRequest)
    - [CreateRewardsResponse](#oracle-manager-v1-CreateRewardsResponse)
    - [Currency](#oracle-manager-v1-Currency)
    - [DeleteCurrencyRequest](#oracle-manager-v1-DeleteCurrencyRequest)
    - [DeleteCurrencyResponse](#oracle-manager-v1-DeleteCurrencyResponse)
    - [DeleteRewardRequest](#oracle-manager-v1-DeleteRewardRequest)
    - [DeleteRewardResponse](#oracle-manager-v1-DeleteRewardResponse)
    - [ExistRewardCondsRequest](#oracle-manager-v1-ExistRewardCondsRequest)
    - [ExistRewardCondsRequest.CondsEntry](#oracle-manager-v1-ExistRewardCondsRequest-CondsEntry)
    - [ExistRewardCondsResponse](#oracle-manager-v1-ExistRewardCondsResponse)
    - [ExistRewardRequest](#oracle-manager-v1-ExistRewardRequest)
    - [ExistRewardResponse](#oracle-manager-v1-ExistRewardResponse)
    - [GetAppCurrenciesRequest](#oracle-manager-v1-GetAppCurrenciesRequest)
    - [GetAppCurrenciesResponse](#oracle-manager-v1-GetAppCurrenciesResponse)
    - [GetCurrenciesRequest](#oracle-manager-v1-GetCurrenciesRequest)
    - [GetCurrenciesResponse](#oracle-manager-v1-GetCurrenciesResponse)
    - [GetCurrencyOnlyRequest](#oracle-manager-v1-GetCurrencyOnlyRequest)
    - [GetCurrencyOnlyRequest.CondsEntry](#oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry)
    - [GetCurrencyOnlyResponse](#oracle-manager-v1-GetCurrencyOnlyResponse)
    - [GetCurrencyRequest](#oracle-manager-v1-GetCurrencyRequest)
    - [GetCurrencyResponse](#oracle-manager-v1-GetCurrencyResponse)
    - [GetRewardOnlyRequest](#oracle-manager-v1-GetRewardOnlyRequest)
    - [GetRewardOnlyRequest.CondsEntry](#oracle-manager-v1-GetRewardOnlyRequest-CondsEntry)
    - [GetRewardOnlyResponse](#oracle-manager-v1-GetRewardOnlyResponse)
    - [GetRewardRequest](#oracle-manager-v1-GetRewardRequest)
    - [GetRewardResponse](#oracle-manager-v1-GetRewardResponse)
    - [GetRewardsRequest](#oracle-manager-v1-GetRewardsRequest)
    - [GetRewardsRequest.CondsEntry](#oracle-manager-v1-GetRewardsRequest-CondsEntry)
    - [GetRewardsResponse](#oracle-manager-v1-GetRewardsResponse)
    - [Reward](#oracle-manager-v1-Reward)
    - [UpdateCurrencyRequest](#oracle-manager-v1-UpdateCurrencyRequest)
    - [UpdateCurrencyResponse](#oracle-manager-v1-UpdateCurrencyResponse)
    - [UpdateRewardRequest](#oracle-manager-v1-UpdateRewardRequest)
    - [UpdateRewardResponse](#oracle-manager-v1-UpdateRewardResponse)
  
    - [OracleManager](#oracle-manager-v1-OracleManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_oraclemgr_oraclemgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/oraclemgr/oraclemgr.proto



<a name="oracle-manager-v1-CountRewardsRequest"></a>

### CountRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountRewardsRequest.CondsEntry](#oracle-manager-v1-CountRewardsRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-CountRewardsRequest-CondsEntry"></a>

### CountRewardsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-CountRewardsResponse"></a>

### CountRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="oracle-manager-v1-CreateAppCurrencyRequest"></a>

### CreateAppCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateAppCurrencyResponse"></a>

### CreateAppCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateCurrencyRequest"></a>

### CreateCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateCurrencyResponse"></a>

### CreateCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateRewardRequest"></a>

### CreateRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-CreateRewardResponse"></a>

### CreateRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-CreateRewardsRequest"></a>

### CreateRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |






<a name="oracle-manager-v1-CreateRewardsResponse"></a>

### CreateRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |






<a name="oracle-manager-v1-Currency"></a>

### Currency



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| PriceVSUSDT | [double](#double) |  |  |
| AppPriceVSUSDT | [double](#double) |  |  |






<a name="oracle-manager-v1-DeleteCurrencyRequest"></a>

### DeleteCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-DeleteCurrencyResponse"></a>

### DeleteCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-DeleteRewardRequest"></a>

### DeleteRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-DeleteRewardResponse"></a>

### DeleteRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-ExistRewardCondsRequest"></a>

### ExistRewardCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistRewardCondsRequest.CondsEntry](#oracle-manager-v1-ExistRewardCondsRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-ExistRewardCondsRequest-CondsEntry"></a>

### ExistRewardCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-ExistRewardCondsResponse"></a>

### ExistRewardCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="oracle-manager-v1-ExistRewardRequest"></a>

### ExistRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-ExistRewardResponse"></a>

### ExistRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="oracle-manager-v1-GetAppCurrenciesRequest"></a>

### GetAppCurrenciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetAppCurrenciesResponse"></a>

### GetAppCurrenciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Currency](#oracle-manager-v1-Currency) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrenciesRequest"></a>

### GetCurrenciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrenciesResponse"></a>

### GetCurrenciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Currency](#oracle-manager-v1-Currency) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrencyOnlyRequest"></a>

### GetCurrencyOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCurrencyOnlyRequest.CondsEntry](#oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry"></a>

### GetCurrencyOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetCurrencyOnlyResponse"></a>

### GetCurrencyOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-GetCurrencyRequest"></a>

### GetCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-GetCurrencyResponse"></a>

### GetCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-GetRewardOnlyRequest"></a>

### GetRewardOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetRewardOnlyRequest.CondsEntry](#oracle-manager-v1-GetRewardOnlyRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-GetRewardOnlyRequest-CondsEntry"></a>

### GetRewardOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetRewardOnlyResponse"></a>

### GetRewardOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-GetRewardRequest"></a>

### GetRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-GetRewardResponse"></a>

### GetRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-GetRewardsRequest"></a>

### GetRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetRewardsRequest.CondsEntry](#oracle-manager-v1-GetRewardsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetRewardsRequest-CondsEntry"></a>

### GetRewardsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetRewardsResponse"></a>

### GetRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-Reward"></a>

### Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| DailyReward | [double](#double) |  |  |






<a name="oracle-manager-v1-UpdateCurrencyRequest"></a>

### UpdateCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-UpdateCurrencyResponse"></a>

### UpdateCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-UpdateRewardRequest"></a>

### UpdateRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-UpdateRewardResponse"></a>

### UpdateRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |





 

 

 


<a name="oracle-manager-v1-OracleManager"></a>

### OracleManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateReward | [CreateRewardRequest](#oracle-manager-v1-CreateRewardRequest) | [CreateRewardResponse](#oracle-manager-v1-CreateRewardResponse) |  |
| CreateRewards | [CreateRewardsRequest](#oracle-manager-v1-CreateRewardsRequest) | [CreateRewardsResponse](#oracle-manager-v1-CreateRewardsResponse) |  |
| UpdateReward | [UpdateRewardRequest](#oracle-manager-v1-UpdateRewardRequest) | [UpdateRewardResponse](#oracle-manager-v1-UpdateRewardResponse) |  |
| GetReward | [GetRewardRequest](#oracle-manager-v1-GetRewardRequest) | [GetRewardResponse](#oracle-manager-v1-GetRewardResponse) |  |
| GetRewardOnly | [GetRewardOnlyRequest](#oracle-manager-v1-GetRewardOnlyRequest) | [GetRewardOnlyResponse](#oracle-manager-v1-GetRewardOnlyResponse) |  |
| GetRewards | [GetRewardsRequest](#oracle-manager-v1-GetRewardsRequest) | [GetRewardsResponse](#oracle-manager-v1-GetRewardsResponse) |  |
| ExistReward | [ExistRewardRequest](#oracle-manager-v1-ExistRewardRequest) | [ExistRewardResponse](#oracle-manager-v1-ExistRewardResponse) |  |
| ExistRewardConds | [ExistRewardCondsRequest](#oracle-manager-v1-ExistRewardCondsRequest) | [ExistRewardCondsResponse](#oracle-manager-v1-ExistRewardCondsResponse) |  |
| DeleteReward | [DeleteRewardRequest](#oracle-manager-v1-DeleteRewardRequest) | [DeleteRewardResponse](#oracle-manager-v1-DeleteRewardResponse) |  |
| CreateCurrency | [CreateCurrencyRequest](#oracle-manager-v1-CreateCurrencyRequest) | [CreateCurrencyResponse](#oracle-manager-v1-CreateCurrencyResponse) |  |
| CreateAppCurrency | [CreateAppCurrencyRequest](#oracle-manager-v1-CreateAppCurrencyRequest) | [CreateAppCurrencyResponse](#oracle-manager-v1-CreateAppCurrencyResponse) |  |
| UpdateCurrency | [UpdateCurrencyRequest](#oracle-manager-v1-UpdateCurrencyRequest) | [UpdateCurrencyResponse](#oracle-manager-v1-UpdateCurrencyResponse) |  |
| GetCurrency | [GetCurrencyRequest](#oracle-manager-v1-GetCurrencyRequest) | [GetCurrencyResponse](#oracle-manager-v1-GetCurrencyResponse) |  |
| GetCurrencyOnly | [GetCurrencyOnlyRequest](#oracle-manager-v1-GetCurrencyOnlyRequest) | [GetCurrencyOnlyResponse](#oracle-manager-v1-GetCurrencyOnlyResponse) |  |
| GetCurrencies | [GetCurrenciesRequest](#oracle-manager-v1-GetCurrenciesRequest) | [GetCurrenciesResponse](#oracle-manager-v1-GetCurrenciesResponse) |  |
| GetAppCurrencies | [GetAppCurrenciesRequest](#oracle-manager-v1-GetAppCurrenciesRequest) | [GetAppCurrenciesResponse](#oracle-manager-v1-GetAppCurrenciesResponse) |  |
| DeleteCurrency | [DeleteCurrencyRequest](#oracle-manager-v1-DeleteCurrencyRequest) | [DeleteCurrencyResponse](#oracle-manager-v1-DeleteCurrencyResponse) |  |

 



<a name="npool_oraclemgr_oraclemgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/oraclemgr/oraclemgr.proto



<a name="oracle-manager-v1-CountRewardsRequest"></a>

### CountRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountRewardsRequest.CondsEntry](#oracle-manager-v1-CountRewardsRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-CountRewardsRequest-CondsEntry"></a>

### CountRewardsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-CountRewardsResponse"></a>

### CountRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="oracle-manager-v1-CreateAppCurrencyRequest"></a>

### CreateAppCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateAppCurrencyResponse"></a>

### CreateAppCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateCurrencyRequest"></a>

### CreateCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateCurrencyResponse"></a>

### CreateCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-CreateRewardRequest"></a>

### CreateRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-CreateRewardResponse"></a>

### CreateRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-CreateRewardsRequest"></a>

### CreateRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |






<a name="oracle-manager-v1-CreateRewardsResponse"></a>

### CreateRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |






<a name="oracle-manager-v1-Currency"></a>

### Currency



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| PriceVSUSDT | [double](#double) |  |  |
| AppPriceVSUSDT | [double](#double) |  |  |






<a name="oracle-manager-v1-DeleteCurrencyRequest"></a>

### DeleteCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-DeleteCurrencyResponse"></a>

### DeleteCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-DeleteRewardRequest"></a>

### DeleteRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-DeleteRewardResponse"></a>

### DeleteRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-ExistRewardCondsRequest"></a>

### ExistRewardCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistRewardCondsRequest.CondsEntry](#oracle-manager-v1-ExistRewardCondsRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-ExistRewardCondsRequest-CondsEntry"></a>

### ExistRewardCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-ExistRewardCondsResponse"></a>

### ExistRewardCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="oracle-manager-v1-ExistRewardRequest"></a>

### ExistRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-ExistRewardResponse"></a>

### ExistRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="oracle-manager-v1-GetAppCurrenciesRequest"></a>

### GetAppCurrenciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetAppCurrenciesResponse"></a>

### GetAppCurrenciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Currency](#oracle-manager-v1-Currency) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrenciesRequest"></a>

### GetCurrenciesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrenciesResponse"></a>

### GetCurrenciesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Currency](#oracle-manager-v1-Currency) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetCurrencyOnlyRequest"></a>

### GetCurrencyOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCurrencyOnlyRequest.CondsEntry](#oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-GetCurrencyOnlyRequest-CondsEntry"></a>

### GetCurrencyOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetCurrencyOnlyResponse"></a>

### GetCurrencyOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-GetCurrencyRequest"></a>

### GetCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-GetCurrencyResponse"></a>

### GetCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-GetRewardOnlyRequest"></a>

### GetRewardOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetRewardOnlyRequest.CondsEntry](#oracle-manager-v1-GetRewardOnlyRequest-CondsEntry) | repeated |  |






<a name="oracle-manager-v1-GetRewardOnlyRequest-CondsEntry"></a>

### GetRewardOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetRewardOnlyResponse"></a>

### GetRewardOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-GetRewardRequest"></a>

### GetRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="oracle-manager-v1-GetRewardResponse"></a>

### GetRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-GetRewardsRequest"></a>

### GetRewardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetRewardsRequest.CondsEntry](#oracle-manager-v1-GetRewardsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="oracle-manager-v1-GetRewardsRequest-CondsEntry"></a>

### GetRewardsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="oracle-manager-v1-GetRewardsResponse"></a>

### GetRewardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Reward](#oracle-manager-v1-Reward) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="oracle-manager-v1-Reward"></a>

### Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| DailyReward | [double](#double) |  |  |






<a name="oracle-manager-v1-UpdateCurrencyRequest"></a>

### UpdateCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-UpdateCurrencyResponse"></a>

### UpdateCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Currency](#oracle-manager-v1-Currency) |  |  |






<a name="oracle-manager-v1-UpdateRewardRequest"></a>

### UpdateRewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |






<a name="oracle-manager-v1-UpdateRewardResponse"></a>

### UpdateRewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Reward](#oracle-manager-v1-Reward) |  |  |





 

 

 


<a name="oracle-manager-v1-OracleManager"></a>

### OracleManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateReward | [CreateRewardRequest](#oracle-manager-v1-CreateRewardRequest) | [CreateRewardResponse](#oracle-manager-v1-CreateRewardResponse) |  |
| CreateRewards | [CreateRewardsRequest](#oracle-manager-v1-CreateRewardsRequest) | [CreateRewardsResponse](#oracle-manager-v1-CreateRewardsResponse) |  |
| UpdateReward | [UpdateRewardRequest](#oracle-manager-v1-UpdateRewardRequest) | [UpdateRewardResponse](#oracle-manager-v1-UpdateRewardResponse) |  |
| GetReward | [GetRewardRequest](#oracle-manager-v1-GetRewardRequest) | [GetRewardResponse](#oracle-manager-v1-GetRewardResponse) |  |
| GetRewardOnly | [GetRewardOnlyRequest](#oracle-manager-v1-GetRewardOnlyRequest) | [GetRewardOnlyResponse](#oracle-manager-v1-GetRewardOnlyResponse) |  |
| GetRewards | [GetRewardsRequest](#oracle-manager-v1-GetRewardsRequest) | [GetRewardsResponse](#oracle-manager-v1-GetRewardsResponse) |  |
| ExistReward | [ExistRewardRequest](#oracle-manager-v1-ExistRewardRequest) | [ExistRewardResponse](#oracle-manager-v1-ExistRewardResponse) |  |
| ExistRewardConds | [ExistRewardCondsRequest](#oracle-manager-v1-ExistRewardCondsRequest) | [ExistRewardCondsResponse](#oracle-manager-v1-ExistRewardCondsResponse) |  |
| DeleteReward | [DeleteRewardRequest](#oracle-manager-v1-DeleteRewardRequest) | [DeleteRewardResponse](#oracle-manager-v1-DeleteRewardResponse) |  |
| CreateCurrency | [CreateCurrencyRequest](#oracle-manager-v1-CreateCurrencyRequest) | [CreateCurrencyResponse](#oracle-manager-v1-CreateCurrencyResponse) |  |
| CreateAppCurrency | [CreateAppCurrencyRequest](#oracle-manager-v1-CreateAppCurrencyRequest) | [CreateAppCurrencyResponse](#oracle-manager-v1-CreateAppCurrencyResponse) |  |
| UpdateCurrency | [UpdateCurrencyRequest](#oracle-manager-v1-UpdateCurrencyRequest) | [UpdateCurrencyResponse](#oracle-manager-v1-UpdateCurrencyResponse) |  |
| GetCurrency | [GetCurrencyRequest](#oracle-manager-v1-GetCurrencyRequest) | [GetCurrencyResponse](#oracle-manager-v1-GetCurrencyResponse) |  |
| GetCurrencyOnly | [GetCurrencyOnlyRequest](#oracle-manager-v1-GetCurrencyOnlyRequest) | [GetCurrencyOnlyResponse](#oracle-manager-v1-GetCurrencyOnlyResponse) |  |
| GetCurrencies | [GetCurrenciesRequest](#oracle-manager-v1-GetCurrenciesRequest) | [GetCurrenciesResponse](#oracle-manager-v1-GetCurrenciesResponse) |  |
| GetAppCurrencies | [GetAppCurrenciesRequest](#oracle-manager-v1-GetAppCurrenciesRequest) | [GetAppCurrenciesResponse](#oracle-manager-v1-GetAppCurrenciesResponse) |  |
| DeleteCurrency | [DeleteCurrencyRequest](#oracle-manager-v1-DeleteCurrencyRequest) | [DeleteCurrencyResponse](#oracle-manager-v1-DeleteCurrencyResponse) |  |

 



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

