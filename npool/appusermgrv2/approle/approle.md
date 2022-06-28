# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/appusermgrv2/approle/approle.proto](#npool_appusermgrv2_approle_approle-proto)
    - [AppRole](#app-user-manager-v2-AppRole)
    - [CountAppRolesRequest](#app-user-manager-v2-CountAppRolesRequest)
    - [CountAppRolesRequest.CondsEntry](#app-user-manager-v2-CountAppRolesRequest-CondsEntry)
    - [CountAppRolesResponse](#app-user-manager-v2-CountAppRolesResponse)
    - [CreateAppRoleRequest](#app-user-manager-v2-CreateAppRoleRequest)
    - [CreateAppRoleResponse](#app-user-manager-v2-CreateAppRoleResponse)
    - [CreateAppRolesRequest](#app-user-manager-v2-CreateAppRolesRequest)
    - [CreateAppRolesResponse](#app-user-manager-v2-CreateAppRolesResponse)
    - [DeleteAppRoleRequest](#app-user-manager-v2-DeleteAppRoleRequest)
    - [DeleteAppRoleResponse](#app-user-manager-v2-DeleteAppRoleResponse)
    - [ExistAppRoleCondsRequest](#app-user-manager-v2-ExistAppRoleCondsRequest)
    - [ExistAppRoleCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry)
    - [ExistAppRoleCondsResponse](#app-user-manager-v2-ExistAppRoleCondsResponse)
    - [ExistAppRoleRequest](#app-user-manager-v2-ExistAppRoleRequest)
    - [ExistAppRoleResponse](#app-user-manager-v2-ExistAppRoleResponse)
    - [GetAppRoleOnlyRequest](#app-user-manager-v2-GetAppRoleOnlyRequest)
    - [GetAppRoleOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry)
    - [GetAppRoleOnlyResponse](#app-user-manager-v2-GetAppRoleOnlyResponse)
    - [GetAppRoleRequest](#app-user-manager-v2-GetAppRoleRequest)
    - [GetAppRoleResponse](#app-user-manager-v2-GetAppRoleResponse)
    - [GetAppRolesRequest](#app-user-manager-v2-GetAppRolesRequest)
    - [GetAppRolesRequest.CondsEntry](#app-user-manager-v2-GetAppRolesRequest-CondsEntry)
    - [GetAppRolesResponse](#app-user-manager-v2-GetAppRolesResponse)
    - [UpdateAppRoleFieldsRequest](#app-user-manager-v2-UpdateAppRoleFieldsRequest)
    - [UpdateAppRoleFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry)
    - [UpdateAppRoleFieldsResponse](#app-user-manager-v2-UpdateAppRoleFieldsResponse)
    - [UpdateAppRoleRequest](#app-user-manager-v2-UpdateAppRoleRequest)
    - [UpdateAppRoleResponse](#app-user-manager-v2-UpdateAppRoleResponse)
  
    - [AppUserManagerAppRole](#app-user-manager-v2-AppUserManagerAppRole)
  
- [npool/appusermgrv2/approle/approle.proto](#npool_appusermgrv2_approle_approle-proto)
    - [AppRole](#app-user-manager-v2-AppRole)
    - [CountAppRolesRequest](#app-user-manager-v2-CountAppRolesRequest)
    - [CountAppRolesRequest.CondsEntry](#app-user-manager-v2-CountAppRolesRequest-CondsEntry)
    - [CountAppRolesResponse](#app-user-manager-v2-CountAppRolesResponse)
    - [CreateAppRoleRequest](#app-user-manager-v2-CreateAppRoleRequest)
    - [CreateAppRoleResponse](#app-user-manager-v2-CreateAppRoleResponse)
    - [CreateAppRolesRequest](#app-user-manager-v2-CreateAppRolesRequest)
    - [CreateAppRolesResponse](#app-user-manager-v2-CreateAppRolesResponse)
    - [DeleteAppRoleRequest](#app-user-manager-v2-DeleteAppRoleRequest)
    - [DeleteAppRoleResponse](#app-user-manager-v2-DeleteAppRoleResponse)
    - [ExistAppRoleCondsRequest](#app-user-manager-v2-ExistAppRoleCondsRequest)
    - [ExistAppRoleCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry)
    - [ExistAppRoleCondsResponse](#app-user-manager-v2-ExistAppRoleCondsResponse)
    - [ExistAppRoleRequest](#app-user-manager-v2-ExistAppRoleRequest)
    - [ExistAppRoleResponse](#app-user-manager-v2-ExistAppRoleResponse)
    - [GetAppRoleOnlyRequest](#app-user-manager-v2-GetAppRoleOnlyRequest)
    - [GetAppRoleOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry)
    - [GetAppRoleOnlyResponse](#app-user-manager-v2-GetAppRoleOnlyResponse)
    - [GetAppRoleRequest](#app-user-manager-v2-GetAppRoleRequest)
    - [GetAppRoleResponse](#app-user-manager-v2-GetAppRoleResponse)
    - [GetAppRolesRequest](#app-user-manager-v2-GetAppRolesRequest)
    - [GetAppRolesRequest.CondsEntry](#app-user-manager-v2-GetAppRolesRequest-CondsEntry)
    - [GetAppRolesResponse](#app-user-manager-v2-GetAppRolesResponse)
    - [UpdateAppRoleFieldsRequest](#app-user-manager-v2-UpdateAppRoleFieldsRequest)
    - [UpdateAppRoleFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry)
    - [UpdateAppRoleFieldsResponse](#app-user-manager-v2-UpdateAppRoleFieldsResponse)
    - [UpdateAppRoleRequest](#app-user-manager-v2-UpdateAppRoleRequest)
    - [UpdateAppRoleResponse](#app-user-manager-v2-UpdateAppRoleResponse)
  
    - [AppUserManagerAppRole](#app-user-manager-v2-AppUserManagerAppRole)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_appusermgrv2_approle_approle-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/appusermgrv2/approle/approle.proto



<a name="app-user-manager-v2-AppRole"></a>

### AppRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Role | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Default | [bool](#bool) |  |  |






<a name="app-user-manager-v2-CountAppRolesRequest"></a>

### CountAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountAppRolesRequest.CondsEntry](#app-user-manager-v2-CountAppRolesRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-CountAppRolesRequest-CondsEntry"></a>

### CountAppRolesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-CountAppRolesResponse"></a>

### CountAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CreateAppRoleRequest"></a>

### CreateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-CreateAppRoleResponse"></a>

### CreateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-CreateAppRolesRequest"></a>

### CreateAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |






<a name="app-user-manager-v2-CreateAppRolesResponse"></a>

### CreateAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |






<a name="app-user-manager-v2-DeleteAppRoleRequest"></a>

### DeleteAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-DeleteAppRoleResponse"></a>

### DeleteAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-ExistAppRoleCondsRequest"></a>

### ExistAppRoleCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistAppRoleCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry"></a>

### ExistAppRoleCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-ExistAppRoleCondsResponse"></a>

### ExistAppRoleCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-ExistAppRoleRequest"></a>

### ExistAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-ExistAppRoleResponse"></a>

### ExistAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-GetAppRoleOnlyRequest"></a>

### GetAppRoleOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppRoleOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry"></a>

### GetAppRoleOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppRoleOnlyResponse"></a>

### GetAppRoleOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-GetAppRoleRequest"></a>

### GetAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-GetAppRoleResponse"></a>

### GetAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-GetAppRolesRequest"></a>

### GetAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppRolesRequest.CondsEntry](#app-user-manager-v2-GetAppRolesRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="app-user-manager-v2-GetAppRolesRequest-CondsEntry"></a>

### GetAppRolesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppRolesResponse"></a>

### GetAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsRequest"></a>

### UpdateAppRoleFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateAppRoleFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry) | repeated |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry"></a>

### UpdateAppRoleFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsResponse"></a>

### UpdateAppRoleFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleRequest"></a>

### UpdateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleResponse"></a>

### UpdateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |





 

 

 


<a name="app-user-manager-v2-AppUserManagerAppRole"></a>

### AppUserManagerAppRole


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateAppRoleV2 | [CreateAppRoleRequest](#app-user-manager-v2-CreateAppRoleRequest) | [CreateAppRoleResponse](#app-user-manager-v2-CreateAppRoleResponse) |  |
| CreateAppRolesV2 | [CreateAppRolesRequest](#app-user-manager-v2-CreateAppRolesRequest) | [CreateAppRolesResponse](#app-user-manager-v2-CreateAppRolesResponse) |  |
| UpdateAppRoleV2 | [UpdateAppRoleRequest](#app-user-manager-v2-UpdateAppRoleRequest) | [UpdateAppRoleResponse](#app-user-manager-v2-UpdateAppRoleResponse) |  |
| UpdateAppRoleFieldsV2 | [UpdateAppRoleFieldsRequest](#app-user-manager-v2-UpdateAppRoleFieldsRequest) | [UpdateAppRoleFieldsResponse](#app-user-manager-v2-UpdateAppRoleFieldsResponse) |  |
| GetAppRoleV2 | [GetAppRoleRequest](#app-user-manager-v2-GetAppRoleRequest) | [GetAppRoleResponse](#app-user-manager-v2-GetAppRoleResponse) |  |
| GetAppRoleOnlyV2 | [GetAppRoleOnlyRequest](#app-user-manager-v2-GetAppRoleOnlyRequest) | [GetAppRoleOnlyResponse](#app-user-manager-v2-GetAppRoleOnlyResponse) |  |
| GetAppRolesV2 | [GetAppRolesRequest](#app-user-manager-v2-GetAppRolesRequest) | [GetAppRolesResponse](#app-user-manager-v2-GetAppRolesResponse) |  |
| ExistAppRoleV2 | [ExistAppRoleRequest](#app-user-manager-v2-ExistAppRoleRequest) | [ExistAppRoleResponse](#app-user-manager-v2-ExistAppRoleResponse) |  |
| ExistAppRoleCondsV2 | [ExistAppRoleCondsRequest](#app-user-manager-v2-ExistAppRoleCondsRequest) | [ExistAppRoleCondsResponse](#app-user-manager-v2-ExistAppRoleCondsResponse) |  |
| CountAppRolesV2 | [CountAppRolesRequest](#app-user-manager-v2-CountAppRolesRequest) | [CountAppRolesResponse](#app-user-manager-v2-CountAppRolesResponse) |  |
| DeleteAppRoleV2 | [DeleteAppRoleRequest](#app-user-manager-v2-DeleteAppRoleRequest) | [DeleteAppRoleResponse](#app-user-manager-v2-DeleteAppRoleResponse) |  |

 



<a name="npool_appusermgrv2_approle_approle-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/appusermgrv2/approle/approle.proto



<a name="app-user-manager-v2-AppRole"></a>

### AppRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Role | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Default | [bool](#bool) |  |  |






<a name="app-user-manager-v2-CountAppRolesRequest"></a>

### CountAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountAppRolesRequest.CondsEntry](#app-user-manager-v2-CountAppRolesRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-CountAppRolesRequest-CondsEntry"></a>

### CountAppRolesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-CountAppRolesResponse"></a>

### CountAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-CreateAppRoleRequest"></a>

### CreateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-CreateAppRoleResponse"></a>

### CreateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-CreateAppRolesRequest"></a>

### CreateAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |






<a name="app-user-manager-v2-CreateAppRolesResponse"></a>

### CreateAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |






<a name="app-user-manager-v2-DeleteAppRoleRequest"></a>

### DeleteAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-DeleteAppRoleResponse"></a>

### DeleteAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-ExistAppRoleCondsRequest"></a>

### ExistAppRoleCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistAppRoleCondsRequest.CondsEntry](#app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-ExistAppRoleCondsRequest-CondsEntry"></a>

### ExistAppRoleCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-ExistAppRoleCondsResponse"></a>

### ExistAppRoleCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-ExistAppRoleRequest"></a>

### ExistAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-ExistAppRoleResponse"></a>

### ExistAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="app-user-manager-v2-GetAppRoleOnlyRequest"></a>

### GetAppRoleOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppRoleOnlyRequest.CondsEntry](#app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry) | repeated |  |






<a name="app-user-manager-v2-GetAppRoleOnlyRequest-CondsEntry"></a>

### GetAppRoleOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppRoleOnlyResponse"></a>

### GetAppRoleOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-GetAppRoleRequest"></a>

### GetAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app-user-manager-v2-GetAppRoleResponse"></a>

### GetAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-GetAppRolesRequest"></a>

### GetAppRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetAppRolesRequest.CondsEntry](#app-user-manager-v2-GetAppRolesRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="app-user-manager-v2-GetAppRolesRequest-CondsEntry"></a>

### GetAppRolesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="app-user-manager-v2-GetAppRolesResponse"></a>

### GetAppRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app-user-manager-v2-AppRole) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsRequest"></a>

### UpdateAppRoleFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateAppRoleFieldsRequest.FieldsEntry](#app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry) | repeated |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsRequest-FieldsEntry"></a>

### UpdateAppRoleFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleFieldsResponse"></a>

### UpdateAppRoleFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleRequest"></a>

### UpdateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |






<a name="app-user-manager-v2-UpdateAppRoleResponse"></a>

### UpdateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app-user-manager-v2-AppRole) |  |  |





 

 

 


<a name="app-user-manager-v2-AppUserManagerAppRole"></a>

### AppUserManagerAppRole


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateAppRoleV2 | [CreateAppRoleRequest](#app-user-manager-v2-CreateAppRoleRequest) | [CreateAppRoleResponse](#app-user-manager-v2-CreateAppRoleResponse) |  |
| CreateAppRolesV2 | [CreateAppRolesRequest](#app-user-manager-v2-CreateAppRolesRequest) | [CreateAppRolesResponse](#app-user-manager-v2-CreateAppRolesResponse) |  |
| UpdateAppRoleV2 | [UpdateAppRoleRequest](#app-user-manager-v2-UpdateAppRoleRequest) | [UpdateAppRoleResponse](#app-user-manager-v2-UpdateAppRoleResponse) |  |
| UpdateAppRoleFieldsV2 | [UpdateAppRoleFieldsRequest](#app-user-manager-v2-UpdateAppRoleFieldsRequest) | [UpdateAppRoleFieldsResponse](#app-user-manager-v2-UpdateAppRoleFieldsResponse) |  |
| GetAppRoleV2 | [GetAppRoleRequest](#app-user-manager-v2-GetAppRoleRequest) | [GetAppRoleResponse](#app-user-manager-v2-GetAppRoleResponse) |  |
| GetAppRoleOnlyV2 | [GetAppRoleOnlyRequest](#app-user-manager-v2-GetAppRoleOnlyRequest) | [GetAppRoleOnlyResponse](#app-user-manager-v2-GetAppRoleOnlyResponse) |  |
| GetAppRolesV2 | [GetAppRolesRequest](#app-user-manager-v2-GetAppRolesRequest) | [GetAppRolesResponse](#app-user-manager-v2-GetAppRolesResponse) |  |
| ExistAppRoleV2 | [ExistAppRoleRequest](#app-user-manager-v2-ExistAppRoleRequest) | [ExistAppRoleResponse](#app-user-manager-v2-ExistAppRoleResponse) |  |
| ExistAppRoleCondsV2 | [ExistAppRoleCondsRequest](#app-user-manager-v2-ExistAppRoleCondsRequest) | [ExistAppRoleCondsResponse](#app-user-manager-v2-ExistAppRoleCondsResponse) |  |
| CountAppRolesV2 | [CountAppRolesRequest](#app-user-manager-v2-CountAppRolesRequest) | [CountAppRolesResponse](#app-user-manager-v2-CountAppRolesResponse) |  |
| DeleteAppRoleV2 | [DeleteAppRoleRequest](#app-user-manager-v2-DeleteAppRoleRequest) | [DeleteAppRoleResponse](#app-user-manager-v2-DeleteAppRoleResponse) |  |

 



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

