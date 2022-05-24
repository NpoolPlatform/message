# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/projectinfomgr/projectinfomgr.proto](#npool_projectinfomgr_projectinfomgr-proto)
    - [CoinDescription](#project-info-manager-v1-CoinDescription)
    - [CreateAppCoinDescriptionRequest](#project-info-manager-v1-CreateAppCoinDescriptionRequest)
    - [CreateAppCoinDescriptionResponse](#project-info-manager-v1-CreateAppCoinDescriptionResponse)
    - [CreateAppCoinDescriptionsRequest](#project-info-manager-v1-CreateAppCoinDescriptionsRequest)
    - [CreateAppCoinDescriptionsResponse](#project-info-manager-v1-CreateAppCoinDescriptionsResponse)
    - [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse)
    - [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest)
    - [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse)
    - [DeleteCoinDescriptionRequest](#project-info-manager-v1-DeleteCoinDescriptionRequest)
    - [DeleteCoinDescriptionResponse](#project-info-manager-v1-DeleteCoinDescriptionResponse)
    - [GetAppCoinDescriptionOnlyRequest](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest)
    - [GetAppCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry)
    - [GetAppCoinDescriptionOnlyResponse](#project-info-manager-v1-GetAppCoinDescriptionOnlyResponse)
    - [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest)
    - [GetAppCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry)
    - [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse)
    - [GetCoinDescriptionOnlyRequest](#project-info-manager-v1-GetCoinDescriptionOnlyRequest)
    - [GetCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry)
    - [GetCoinDescriptionOnlyResponse](#project-info-manager-v1-GetCoinDescriptionOnlyResponse)
    - [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse)
    - [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest)
    - [GetCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry)
    - [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse)
    - [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse)
  
    - [ProjectInfoManager](#project-info-manager-v1-ProjectInfoManager)
  
- [npool/projectinfomgr/projectinfomgr.proto](#npool_projectinfomgr_projectinfomgr-proto)
    - [CoinDescription](#project-info-manager-v1-CoinDescription)
    - [CreateAppCoinDescriptionRequest](#project-info-manager-v1-CreateAppCoinDescriptionRequest)
    - [CreateAppCoinDescriptionResponse](#project-info-manager-v1-CreateAppCoinDescriptionResponse)
    - [CreateAppCoinDescriptionsRequest](#project-info-manager-v1-CreateAppCoinDescriptionsRequest)
    - [CreateAppCoinDescriptionsResponse](#project-info-manager-v1-CreateAppCoinDescriptionsResponse)
    - [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse)
    - [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest)
    - [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse)
    - [DeleteCoinDescriptionRequest](#project-info-manager-v1-DeleteCoinDescriptionRequest)
    - [DeleteCoinDescriptionResponse](#project-info-manager-v1-DeleteCoinDescriptionResponse)
    - [GetAppCoinDescriptionOnlyRequest](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest)
    - [GetAppCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry)
    - [GetAppCoinDescriptionOnlyResponse](#project-info-manager-v1-GetAppCoinDescriptionOnlyResponse)
    - [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest)
    - [GetAppCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry)
    - [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse)
    - [GetCoinDescriptionOnlyRequest](#project-info-manager-v1-GetCoinDescriptionOnlyRequest)
    - [GetCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry)
    - [GetCoinDescriptionOnlyResponse](#project-info-manager-v1-GetCoinDescriptionOnlyResponse)
    - [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse)
    - [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest)
    - [GetCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry)
    - [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse)
    - [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse)
  
    - [ProjectInfoManager](#project-info-manager-v1-ProjectInfoManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_projectinfomgr_projectinfomgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/projectinfomgr/projectinfomgr.proto



<a name="project-info-manager-v1-CoinDescription"></a>

### CoinDescription



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionRequest"></a>

### CreateAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionResponse"></a>

### CreateAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionsRequest"></a>

### CreateAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionsResponse"></a>

### CreateAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionRequest"></a>

### CreateCoinDescriptionRequest
create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateCoinDescriptionResponse"></a>

### CreateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsRequest"></a>

### CreateCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsResponse"></a>

### CreateCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-DeleteCoinDescriptionRequest"></a>

### DeleteCoinDescriptionRequest
delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-DeleteCoinDescriptionResponse"></a>

### DeleteCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyRequest"></a>

### GetAppCoinDescriptionOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry"></a>

### GetAppCoinDescriptionOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyResponse"></a>

### GetAppCoinDescriptionOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest"></a>

### GetAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry"></a>

### GetAppCoinDescriptionsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsResponse"></a>

### GetAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyRequest"></a>

### GetCoinDescriptionOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry"></a>

### GetCoinDescriptionOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyResponse"></a>

### GetCoinDescriptionOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest
get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionResponse"></a>

### GetCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsRequest"></a>

### GetCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry"></a>

### GetCoinDescriptionsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsResponse"></a>

### GetCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest
update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionResponse"></a>

### UpdateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |





 

 

 


<a name="project-info-manager-v1-ProjectInfoManager"></a>

### ProjectInfoManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateCoinDescription | [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest) | [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse) |  |
| CreateCoinDescriptions | [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest) | [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse) |  |
| CreateAppCoinDescription | [CreateAppCoinDescriptionRequest](#project-info-manager-v1-CreateAppCoinDescriptionRequest) | [CreateAppCoinDescriptionResponse](#project-info-manager-v1-CreateAppCoinDescriptionResponse) |  |
| CreateAppCoinDescriptions | [CreateAppCoinDescriptionsRequest](#project-info-manager-v1-CreateAppCoinDescriptionsRequest) | [CreateAppCoinDescriptionsResponse](#project-info-manager-v1-CreateAppCoinDescriptionsResponse) |  |
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse) |  |
| GetCoinDescriptions | [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest) | [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse) |  |
| GetCoinDescriptionOnly | [GetCoinDescriptionOnlyRequest](#project-info-manager-v1-GetCoinDescriptionOnlyRequest) | [GetCoinDescriptionOnlyResponse](#project-info-manager-v1-GetCoinDescriptionOnlyResponse) |  |
| GetAppCoinDescriptions | [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest) | [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse) |  |
| GetAppCoinDescriptionOnly | [GetAppCoinDescriptionOnlyRequest](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest) | [GetAppCoinDescriptionOnlyResponse](#project-info-manager-v1-GetAppCoinDescriptionOnlyResponse) |  |
| DeleteCoinDescription | [DeleteCoinDescriptionRequest](#project-info-manager-v1-DeleteCoinDescriptionRequest) | [DeleteCoinDescriptionResponse](#project-info-manager-v1-DeleteCoinDescriptionResponse) |  |

 



<a name="npool_projectinfomgr_projectinfomgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/projectinfomgr/projectinfomgr.proto



<a name="project-info-manager-v1-CoinDescription"></a>

### CoinDescription



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionRequest"></a>

### CreateAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionResponse"></a>

### CreateAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionsRequest"></a>

### CreateAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateAppCoinDescriptionsResponse"></a>

### CreateAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionRequest"></a>

### CreateCoinDescriptionRequest
create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateCoinDescriptionResponse"></a>

### CreateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsRequest"></a>

### CreateCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsResponse"></a>

### CreateCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-DeleteCoinDescriptionRequest"></a>

### DeleteCoinDescriptionRequest
delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-DeleteCoinDescriptionResponse"></a>

### DeleteCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyRequest"></a>

### GetAppCoinDescriptionOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyRequest-CondsEntry"></a>

### GetAppCoinDescriptionOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionOnlyResponse"></a>

### GetAppCoinDescriptionOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest"></a>

### GetAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest-CondsEntry"></a>

### GetAppCoinDescriptionsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsResponse"></a>

### GetAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyRequest"></a>

### GetCoinDescriptionOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetCoinDescriptionOnlyRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyRequest-CondsEntry"></a>

### GetCoinDescriptionOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionOnlyResponse"></a>

### GetCoinDescriptionOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest
get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionResponse"></a>

### GetCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsRequest"></a>

### GetCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetCoinDescriptionsRequest.CondsEntry](#project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsRequest-CondsEntry"></a>

### GetCoinDescriptionsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsResponse"></a>

### GetCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest
update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionResponse"></a>

### UpdateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |





 

 

 


<a name="project-info-manager-v1-ProjectInfoManager"></a>

### ProjectInfoManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateCoinDescription | [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest) | [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse) |  |
| CreateCoinDescriptions | [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest) | [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse) |  |
| CreateAppCoinDescription | [CreateAppCoinDescriptionRequest](#project-info-manager-v1-CreateAppCoinDescriptionRequest) | [CreateAppCoinDescriptionResponse](#project-info-manager-v1-CreateAppCoinDescriptionResponse) |  |
| CreateAppCoinDescriptions | [CreateAppCoinDescriptionsRequest](#project-info-manager-v1-CreateAppCoinDescriptionsRequest) | [CreateAppCoinDescriptionsResponse](#project-info-manager-v1-CreateAppCoinDescriptionsResponse) |  |
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse) |  |
| GetCoinDescriptions | [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest) | [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse) |  |
| GetCoinDescriptionOnly | [GetCoinDescriptionOnlyRequest](#project-info-manager-v1-GetCoinDescriptionOnlyRequest) | [GetCoinDescriptionOnlyResponse](#project-info-manager-v1-GetCoinDescriptionOnlyResponse) |  |
| GetAppCoinDescriptions | [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest) | [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse) |  |
| GetAppCoinDescriptionOnly | [GetAppCoinDescriptionOnlyRequest](#project-info-manager-v1-GetAppCoinDescriptionOnlyRequest) | [GetAppCoinDescriptionOnlyResponse](#project-info-manager-v1-GetAppCoinDescriptionOnlyResponse) |  |
| DeleteCoinDescription | [DeleteCoinDescriptionRequest](#project-info-manager-v1-DeleteCoinDescriptionRequest) | [DeleteCoinDescriptionResponse](#project-info-manager-v1-DeleteCoinDescriptionResponse) |  |

 



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

