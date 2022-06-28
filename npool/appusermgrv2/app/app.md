# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/appusermgrv2/app/app.proto](#npool_appusermgrv2_app_app-proto)
    - [App](#app-user-manager-v2-App)
    - [CountAppsRequest](#app-user-manager-v2-CountAppsRequest)
    - [CountAppsRequest.CondsEntry](#app-user-manager-v2-CountAppsRequest-CondsEntry)
    - [CountAppsResponse](#app-user-manager-v2-CountAppsResponse)
    - [CreateAppRequest](#app-user-manager-v2-CreateAppRequest)
    - [CreateAppResponse](#app-user-manager-v2-CreateAppResponse)
    - [CreateAppsRequest](#app-user-manager-v2-CreateAppsRequest)
    - [CreateAppsResponse](#app-user-manager-v2-CreateAppsResponse)
    - [DeleteAppRequest](#app-user-manager-v2-DeleteAppRequest)
    - [DeleteAppResponse](#app-user-manager-v2-DeleteAppResponse)
    - [ExistAppCondsRequest](#app-user-manager-v2-ExistAppCondsRequest)
    - [ExistAppCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppCondsRequest-CondsEntry)
    - [ExistAppCondsResponse](#app-user-manager-v2-ExistAppCondsResponse)
    - [ExistAppRequest](#app-user-manager-v2-ExistAppRequest)
    - [ExistAppResponse](#app-user-manager-v2-ExistAppResponse)
    - [GetAppOnlyRequest](#app-user-manager-v2-GetAppOnlyRequest)
    - [GetAppOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppOnlyRequest-CondsEntry)
    - [GetAppOnlyResponse](#app-user-manager-v2-GetAppOnlyResponse)
    - [GetAppRequest](#app-user-manager-v2-GetAppRequest)
    - [GetAppResponse](#app-user-manager-v2-GetAppResponse)
    - [GetAppsRequest](#app-user-manager-v2-GetAppsRequest)
    - [GetAppsRequest.CondsEntry](#app-user-manager-v2-GetAppsRequest-CondsEntry)
    - [GetAppsResponse](#app-user-manager-v2-GetAppsResponse)
    - [UpdateAppFieldsRequest](#app-user-manager-v2-UpdateAppFieldsRequest)
    - [UpdateAppFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry)
    - [UpdateAppFieldsResponse](#app-user-manager-v2-UpdateAppFieldsResponse)
    - [UpdateAppRequest](#app-user-manager-v2-UpdateAppRequest)
    - [UpdateAppResponse](#app-user-manager-v2-UpdateAppResponse)
  
    - [AppUserManagerApp](#app-user-manager-v2-AppUserManagerApp)
  
- [npool/appusermgrv2/app/app.proto](#npool_appusermgrv2_app_app-proto)
    - [App](#app-user-manager-v2-App)
    - [CountAppsRequest](#app-user-manager-v2-CountAppsRequest)
    - [CountAppsRequest.CondsEntry](#app-user-manager-v2-CountAppsRequest-CondsEntry)
    - [CountAppsResponse](#app-user-manager-v2-CountAppsResponse)
    - [CreateAppRequest](#app-user-manager-v2-CreateAppRequest)
    - [CreateAppResponse](#app-user-manager-v2-CreateAppResponse)
    - [CreateAppsRequest](#app-user-manager-v2-CreateAppsRequest)
    - [CreateAppsResponse](#app-user-manager-v2-CreateAppsResponse)
    - [DeleteAppRequest](#app-user-manager-v2-DeleteAppRequest)
    - [DeleteAppResponse](#app-user-manager-v2-DeleteAppResponse)
    - [ExistAppCondsRequest](#app-user-manager-v2-ExistAppCondsRequest)
    - [ExistAppCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppCondsRequest-CondsEntry)
    - [ExistAppCondsResponse](#app-user-manager-v2-ExistAppCondsResponse)
    - [ExistAppRequest](#app-user-manager-v2-ExistAppRequest)
    - [ExistAppResponse](#app-user-manager-v2-ExistAppResponse)
    - [GetAppOnlyRequest](#app-user-manager-v2-GetAppOnlyRequest)
    - [GetAppOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppOnlyRequest-CondsEntry)
    - [GetAppOnlyResponse](#app-user-manager-v2-GetAppOnlyResponse)
    - [GetAppRequest](#app-user-manager-v2-GetAppRequest)
    - [GetAppResponse](#app-user-manager-v2-GetAppResponse)
    - [GetAppsRequest](#app-user-manager-v2-GetAppsRequest)
    - [GetAppsRequest.CondsEntry](#app-user-manager-v2-GetAppsRequest-CondsEntry)
    - [GetAppsResponse](#app-user-manager-v2-GetAppsResponse)
    - [UpdateAppFieldsRequest](#app-user-manager-v2-UpdateAppFieldsRequest)
    - [UpdateAppFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry)
    - [UpdateAppFieldsResponse](#app-user-manager-v2-UpdateAppFieldsResponse)
    - [UpdateAppRequest](#app-user-manager-v2-UpdateAppRequest)
    - [UpdateAppResponse](#app-user-manager-v2-UpdateAppResponse)
  
    - [AppUserManagerApp](#app-user-manager-v2-AppUserManagerApp)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_appusermgrv2_app_app-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/appusermgrv2/app/app.proto



<a name="app-user-manager-v2-App"></a>

### App



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CountAppsRequest"></a>

### CountAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountAppsRequest.CondsEntry](#app-user-manager-v2-CountAppsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-CountAppsRequest-CondsEntry"></a>

### CountAppsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-CountAppsResponse"></a>

### CountAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CreateAppRequest"></a>

### CreateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-CreateAppResponse"></a>

### CreateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-CreateAppsRequest"></a>

### CreateAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |






<a name="app-user-manager-v2-CreateAppsResponse"></a>

### CreateAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |






<a name="app-user-manager-v2-DeleteAppRequest"></a>

### DeleteAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-DeleteAppResponse"></a>

### DeleteAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-ExistAppCondsRequest"></a>

### ExistAppCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistAppCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppCondsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-ExistAppCondsRequest-CondsEntry"></a>

### ExistAppCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-ExistAppCondsResponse"></a>

### ExistAppCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-ExistAppRequest"></a>

### ExistAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-ExistAppResponse"></a>

### ExistAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-GetAppOnlyRequest"></a>

### GetAppOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppOnlyRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-GetAppOnlyRequest-CondsEntry"></a>

### GetAppOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppOnlyResponse"></a>

### GetAppOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-GetAppRequest"></a>

### GetAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-GetAppResponse"></a>

### GetAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-GetAppsRequest"></a>

### GetAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppsRequest.CondsEntry](#app-user-manager-v2-GetAppsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="app-user-manager-v2-GetAppsRequest-CondsEntry"></a>

### GetAppsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppsResponse"></a>

### GetAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-UpdateAppFieldsRequest"></a>

### UpdateAppFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateAppFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry) | repeated |  |






<a name="app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry"></a>

### UpdateAppFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="app-user-manager-v2-UpdateAppFieldsResponse"></a>

### UpdateAppFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-UpdateAppRequest"></a>

### UpdateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-UpdateAppResponse"></a>

### UpdateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |





 

 

 


<a name="app-user-manager-v2-AppUserManagerApp"></a>

### AppUserManagerApp


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateAppV2 | [CreateAppRequest](#app-user-manager-v2-CreateAppRequest) | [CreateAppResponse](#app-user-manager-v2-CreateAppResponse) |  |
| CreateAppsV2 | [CreateAppsRequest](#app-user-manager-v2-CreateAppsRequest) | [CreateAppsResponse](#app-user-manager-v2-CreateAppsResponse) |  |
| UpdateAppV2 | [UpdateAppRequest](#app-user-manager-v2-UpdateAppRequest) | [UpdateAppResponse](#app-user-manager-v2-UpdateAppResponse) |  |
| UpdateAppFieldsV2 | [UpdateAppFieldsRequest](#app-user-manager-v2-UpdateAppFieldsRequest) | [UpdateAppFieldsResponse](#app-user-manager-v2-UpdateAppFieldsResponse) |  |
| GetAppV2 | [GetAppRequest](#app-user-manager-v2-GetAppRequest) | [GetAppResponse](#app-user-manager-v2-GetAppResponse) |  |
| GetAppOnlyV2 | [GetAppOnlyRequest](#app-user-manager-v2-GetAppOnlyRequest) | [GetAppOnlyResponse](#app-user-manager-v2-GetAppOnlyResponse) |  |
| GetAppsV2 | [GetAppsRequest](#app-user-manager-v2-GetAppsRequest) | [GetAppsResponse](#app-user-manager-v2-GetAppsResponse) |  |
| ExistAppV2 | [ExistAppRequest](#app-user-manager-v2-ExistAppRequest) | [ExistAppResponse](#app-user-manager-v2-ExistAppResponse) |  |
| ExistAppCondsV2 | [ExistAppCondsRequest](#app-user-manager-v2-ExistAppCondsRequest) | [ExistAppCondsResponse](#app-user-manager-v2-ExistAppCondsResponse) |  |
| CountAppsV2 | [CountAppsRequest](#app-user-manager-v2-CountAppsRequest) | [CountAppsResponse](#app-user-manager-v2-CountAppsResponse) |  |
| DeleteAppV2 | [DeleteAppRequest](#app-user-manager-v2-DeleteAppRequest) | [DeleteAppResponse](#app-user-manager-v2-DeleteAppResponse) |  |

 



<a name="npool_appusermgrv2_app_app-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/appusermgrv2/app/app.proto



<a name="app-user-manager-v2-App"></a>

### App



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CountAppsRequest"></a>

### CountAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountAppsRequest.CondsEntry](#app-user-manager-v2-CountAppsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-CountAppsRequest-CondsEntry"></a>

### CountAppsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-CountAppsResponse"></a>

### CountAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CreateAppRequest"></a>

### CreateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-CreateAppResponse"></a>

### CreateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-CreateAppsRequest"></a>

### CreateAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |






<a name="app-user-manager-v2-CreateAppsResponse"></a>

### CreateAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |






<a name="app-user-manager-v2-DeleteAppRequest"></a>

### DeleteAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-DeleteAppResponse"></a>

### DeleteAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-ExistAppCondsRequest"></a>

### ExistAppCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistAppCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppCondsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-ExistAppCondsRequest-CondsEntry"></a>

### ExistAppCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-ExistAppCondsResponse"></a>

### ExistAppCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-ExistAppRequest"></a>

### ExistAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-ExistAppResponse"></a>

### ExistAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-GetAppOnlyRequest"></a>

### GetAppOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppOnlyRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-GetAppOnlyRequest-CondsEntry"></a>

### GetAppOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppOnlyResponse"></a>

### GetAppOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-GetAppRequest"></a>

### GetAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-GetAppResponse"></a>

### GetAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-GetAppsRequest"></a>

### GetAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppsRequest.CondsEntry](#app-user-manager-v2-GetAppsRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="app-user-manager-v2-GetAppsRequest-CondsEntry"></a>

### GetAppsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppsResponse"></a>

### GetAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app-user-manager-v2-App) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-UpdateAppFieldsRequest"></a>

### UpdateAppFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateAppFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry) | repeated |  |






<a name="app-user-manager-v2-UpdateAppFieldsRequest-FieldsEntry"></a>

### UpdateAppFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="app-user-manager-v2-UpdateAppFieldsResponse"></a>

### UpdateAppFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-UpdateAppRequest"></a>

### UpdateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |






<a name="app-user-manager-v2-UpdateAppResponse"></a>

### UpdateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app-user-manager-v2-App) |  |  |





 

 

 


<a name="app-user-manager-v2-AppUserManagerApp"></a>

### AppUserManagerApp


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateAppV2 | [CreateAppRequest](#app-user-manager-v2-CreateAppRequest) | [CreateAppResponse](#app-user-manager-v2-CreateAppResponse) |  |
| CreateAppsV2 | [CreateAppsRequest](#app-user-manager-v2-CreateAppsRequest) | [CreateAppsResponse](#app-user-manager-v2-CreateAppsResponse) |  |
| UpdateAppV2 | [UpdateAppRequest](#app-user-manager-v2-UpdateAppRequest) | [UpdateAppResponse](#app-user-manager-v2-UpdateAppResponse) |  |
| UpdateAppFieldsV2 | [UpdateAppFieldsRequest](#app-user-manager-v2-UpdateAppFieldsRequest) | [UpdateAppFieldsResponse](#app-user-manager-v2-UpdateAppFieldsResponse) |  |
| GetAppV2 | [GetAppRequest](#app-user-manager-v2-GetAppRequest) | [GetAppResponse](#app-user-manager-v2-GetAppResponse) |  |
| GetAppOnlyV2 | [GetAppOnlyRequest](#app-user-manager-v2-GetAppOnlyRequest) | [GetAppOnlyResponse](#app-user-manager-v2-GetAppOnlyResponse) |  |
| GetAppsV2 | [GetAppsRequest](#app-user-manager-v2-GetAppsRequest) | [GetAppsResponse](#app-user-manager-v2-GetAppsResponse) |  |
| ExistAppV2 | [ExistAppRequest](#app-user-manager-v2-ExistAppRequest) | [ExistAppResponse](#app-user-manager-v2-ExistAppResponse) |  |
| ExistAppCondsV2 | [ExistAppCondsRequest](#app-user-manager-v2-ExistAppCondsRequest) | [ExistAppCondsResponse](#app-user-manager-v2-ExistAppCondsResponse) |  |
| CountAppsV2 | [CountAppsRequest](#app-user-manager-v2-CountAppsRequest) | [CountAppsResponse](#app-user-manager-v2-CountAppsResponse) |  |
| DeleteAppV2 | [DeleteAppRequest](#app-user-manager-v2-DeleteAppRequest) | [DeleteAppResponse](#app-user-manager-v2-DeleteAppResponse) |  |

 



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

