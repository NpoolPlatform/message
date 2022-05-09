# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/project-info-manager/project-info-manager.proto](#npool_project-info-manager_project-info-manager-proto)
    - [CoinDescription](#project-info-manager-v1-CoinDescription)
    - [CountAppCoinDescriptionsRequest](#project-info-manager-v1-CountAppCoinDescriptionsRequest)
    - [CountAppCoinDescriptionsResponse](#project-info-manager-v1-CountAppCoinDescriptionsResponse)
    - [CountCoinDescriptionsRequest](#project-info-manager-v1-CountCoinDescriptionsRequest)
    - [CountCoinDescriptionsResponse](#project-info-manager-v1-CountCoinDescriptionsResponse)
    - [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse)
    - [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest)
    - [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse)
    - [DeleteAppCoinDescriptionRequest](#project-info-manager-v1-DeleteAppCoinDescriptionRequest)
    - [DeleteAppCoinDescriptionResponse](#project-info-manager-v1-DeleteAppCoinDescriptionResponse)
    - [GetAppCoinDescriptionRequest](#project-info-manager-v1-GetAppCoinDescriptionRequest)
    - [GetAppCoinDescriptionResponse](#project-info-manager-v1-GetAppCoinDescriptionResponse)
    - [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest)
    - [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse)
    - [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse)
    - [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest)
    - [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse)
    - [UpdateAppCoinDescriptionRequest](#project-info-manager-v1-UpdateAppCoinDescriptionRequest)
    - [UpdateAppCoinDescriptionResponse](#project-info-manager-v1-UpdateAppCoinDescriptionResponse)
    - [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse)
  
    - [ProjectInfoManager](#project-info-manager-v1-ProjectInfoManager)
  
- [npool/project-info-manager/project-info-manager.proto](#npool_project-info-manager_project-info-manager-proto)
    - [CoinDescription](#project-info-manager-v1-CoinDescription)
    - [CountAppCoinDescriptionsRequest](#project-info-manager-v1-CountAppCoinDescriptionsRequest)
    - [CountAppCoinDescriptionsResponse](#project-info-manager-v1-CountAppCoinDescriptionsResponse)
    - [CountCoinDescriptionsRequest](#project-info-manager-v1-CountCoinDescriptionsRequest)
    - [CountCoinDescriptionsResponse](#project-info-manager-v1-CountCoinDescriptionsResponse)
    - [CreateCoinDescriptionRequest](#project-info-manager-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#project-info-manager-v1-CreateCoinDescriptionResponse)
    - [CreateCoinDescriptionsRequest](#project-info-manager-v1-CreateCoinDescriptionsRequest)
    - [CreateCoinDescriptionsResponse](#project-info-manager-v1-CreateCoinDescriptionsResponse)
    - [DeleteAppCoinDescriptionRequest](#project-info-manager-v1-DeleteAppCoinDescriptionRequest)
    - [DeleteAppCoinDescriptionResponse](#project-info-manager-v1-DeleteAppCoinDescriptionResponse)
    - [GetAppCoinDescriptionRequest](#project-info-manager-v1-GetAppCoinDescriptionRequest)
    - [GetAppCoinDescriptionResponse](#project-info-manager-v1-GetAppCoinDescriptionResponse)
    - [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest)
    - [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse)
    - [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse)
    - [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest)
    - [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse)
    - [UpdateAppCoinDescriptionRequest](#project-info-manager-v1-UpdateAppCoinDescriptionRequest)
    - [UpdateAppCoinDescriptionResponse](#project-info-manager-v1-UpdateAppCoinDescriptionResponse)
    - [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse)
  
    - [ProjectInfoManager](#project-info-manager-v1-ProjectInfoManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_project-info-manager_project-info-manager-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/project-info-manager/project-info-manager.proto



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






<a name="project-info-manager-v1-CountAppCoinDescriptionsRequest"></a>

### CountAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="project-info-manager-v1-CountAppCoinDescriptionsResponse"></a>

### CountAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="project-info-manager-v1-CountCoinDescriptionsRequest"></a>

### CountCoinDescriptionsRequest
count for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="project-info-manager-v1-CountCoinDescriptionsResponse"></a>

### CountCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






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
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsResponse"></a>

### CreateCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-DeleteAppCoinDescriptionRequest"></a>

### DeleteAppCoinDescriptionRequest
delete for admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-DeleteAppCoinDescriptionResponse"></a>

### DeleteAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionRequest"></a>

### GetAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionResponse"></a>

### GetAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest"></a>

### GetAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsResponse"></a>

### GetAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest
get for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
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
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsResponse"></a>

### GetCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-UpdateAppCoinDescriptionRequest"></a>

### UpdateAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateAppCoinDescriptionResponse"></a>

### UpdateAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest
update for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
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
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse) |  |
| UpdateAppCoinDescription | [UpdateAppCoinDescriptionRequest](#project-info-manager-v1-UpdateAppCoinDescriptionRequest) | [UpdateAppCoinDescriptionResponse](#project-info-manager-v1-UpdateAppCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse) |  |
| GetAppCoinDescription | [GetAppCoinDescriptionRequest](#project-info-manager-v1-GetAppCoinDescriptionRequest) | [GetAppCoinDescriptionResponse](#project-info-manager-v1-GetAppCoinDescriptionResponse) |  |
| GetCoinDescriptions | [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest) | [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse) |  |
| GetAppCoinDescriptions | [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest) | [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse) |  |
| CountCoinDescriptions | [CountCoinDescriptionsRequest](#project-info-manager-v1-CountCoinDescriptionsRequest) | [CountCoinDescriptionsResponse](#project-info-manager-v1-CountCoinDescriptionsResponse) |  |
| CountAppCoinDescriptions | [CountAppCoinDescriptionsRequest](#project-info-manager-v1-CountAppCoinDescriptionsRequest) | [CountAppCoinDescriptionsResponse](#project-info-manager-v1-CountAppCoinDescriptionsResponse) |  |
| DeleteAppCoinDescription | [DeleteAppCoinDescriptionRequest](#project-info-manager-v1-DeleteAppCoinDescriptionRequest) | [DeleteAppCoinDescriptionResponse](#project-info-manager-v1-DeleteAppCoinDescriptionResponse) |  |

 



<a name="npool_project-info-manager_project-info-manager-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/project-info-manager/project-info-manager.proto



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






<a name="project-info-manager-v1-CountAppCoinDescriptionsRequest"></a>

### CountAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="project-info-manager-v1-CountAppCoinDescriptionsResponse"></a>

### CountAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="project-info-manager-v1-CountCoinDescriptionsRequest"></a>

### CountCoinDescriptionsRequest
count for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="project-info-manager-v1-CountCoinDescriptionsResponse"></a>

### CountCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






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
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-CreateCoinDescriptionsResponse"></a>

### CreateCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |






<a name="project-info-manager-v1-DeleteAppCoinDescriptionRequest"></a>

### DeleteAppCoinDescriptionRequest
delete for admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-DeleteAppCoinDescriptionResponse"></a>

### DeleteAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionRequest"></a>

### GetAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| ID | [string](#string) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionResponse"></a>

### GetAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsRequest"></a>

### GetAppCoinDescriptionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetAppCoinDescriptionsResponse"></a>

### GetAppCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest
get for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
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
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="project-info-manager-v1-GetCoinDescriptionsResponse"></a>

### GetCoinDescriptionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinDescription](#project-info-manager-v1-CoinDescription) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="project-info-manager-v1-UpdateAppCoinDescriptionRequest"></a>

### UpdateAppCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateAppCoinDescriptionResponse"></a>

### UpdateAppCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescription](#project-info-manager-v1-CoinDescription) |  |  |






<a name="project-info-manager-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest
update for app and admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
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
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#project-info-manager-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#project-info-manager-v1-UpdateCoinDescriptionResponse) |  |
| UpdateAppCoinDescription | [UpdateAppCoinDescriptionRequest](#project-info-manager-v1-UpdateAppCoinDescriptionRequest) | [UpdateAppCoinDescriptionResponse](#project-info-manager-v1-UpdateAppCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#project-info-manager-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#project-info-manager-v1-GetCoinDescriptionResponse) |  |
| GetAppCoinDescription | [GetAppCoinDescriptionRequest](#project-info-manager-v1-GetAppCoinDescriptionRequest) | [GetAppCoinDescriptionResponse](#project-info-manager-v1-GetAppCoinDescriptionResponse) |  |
| GetCoinDescriptions | [GetCoinDescriptionsRequest](#project-info-manager-v1-GetCoinDescriptionsRequest) | [GetCoinDescriptionsResponse](#project-info-manager-v1-GetCoinDescriptionsResponse) |  |
| GetAppCoinDescriptions | [GetAppCoinDescriptionsRequest](#project-info-manager-v1-GetAppCoinDescriptionsRequest) | [GetAppCoinDescriptionsResponse](#project-info-manager-v1-GetAppCoinDescriptionsResponse) |  |
| CountCoinDescriptions | [CountCoinDescriptionsRequest](#project-info-manager-v1-CountCoinDescriptionsRequest) | [CountCoinDescriptionsResponse](#project-info-manager-v1-CountCoinDescriptionsResponse) |  |
| CountAppCoinDescriptions | [CountAppCoinDescriptionsRequest](#project-info-manager-v1-CountAppCoinDescriptionsRequest) | [CountAppCoinDescriptionsResponse](#project-info-manager-v1-CountAppCoinDescriptionsResponse) |  |
| DeleteAppCoinDescription | [DeleteAppCoinDescriptionRequest](#project-info-manager-v1-DeleteAppCoinDescriptionRequest) | [DeleteAppCoinDescriptionResponse](#project-info-manager-v1-DeleteAppCoinDescriptionResponse) |  |

 



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

