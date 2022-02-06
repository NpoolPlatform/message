# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/authinggateway/authinggateway.proto](#npool/authinggateway/authinggateway.proto)
    - [AppAuth](#authing.gateway.v1.AppAuth)
    - [AppRoleAuth](#authing.gateway.v1.AppRoleAuth)
    - [AppUserAuth](#authing.gateway.v1.AppUserAuth)
    - [Auth](#authing.gateway.v1.Auth)
    - [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest)
    - [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse)
    - [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest)
    - [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse)
    - [AuthHistory](#authing.gateway.v1.AuthHistory)
    - [CreateAppAuthForOtherAppRequest](#authing.gateway.v1.CreateAppAuthForOtherAppRequest)
    - [CreateAppAuthForOtherAppResponse](#authing.gateway.v1.CreateAppAuthForOtherAppResponse)
    - [CreateAppRoleAuthForOtherAppRequest](#authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest)
    - [CreateAppRoleAuthForOtherAppResponse](#authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse)
    - [CreateAppRoleAuthRequest](#authing.gateway.v1.CreateAppRoleAuthRequest)
    - [CreateAppRoleAuthResponse](#authing.gateway.v1.CreateAppRoleAuthResponse)
    - [CreateAppUserAuthForOtherAppRequest](#authing.gateway.v1.CreateAppUserAuthForOtherAppRequest)
    - [CreateAppUserAuthForOtherAppResponse](#authing.gateway.v1.CreateAppUserAuthForOtherAppResponse)
    - [CreateAppUserAuthRequest](#authing.gateway.v1.CreateAppUserAuthRequest)
    - [CreateAppUserAuthResponse](#authing.gateway.v1.CreateAppUserAuthResponse)
    - [DeleteAppAuthRequest](#authing.gateway.v1.DeleteAppAuthRequest)
    - [DeleteAppAuthResponse](#authing.gateway.v1.DeleteAppAuthResponse)
    - [DeleteAppRoleAuthRequest](#authing.gateway.v1.DeleteAppRoleAuthRequest)
    - [DeleteAppRoleAuthResponse](#authing.gateway.v1.DeleteAppRoleAuthResponse)
    - [DeleteAppUserAuthRequest](#authing.gateway.v1.DeleteAppUserAuthRequest)
    - [DeleteAppUserAuthResponse](#authing.gateway.v1.DeleteAppUserAuthResponse)
    - [GetAppAuthByAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppResourceMethodRequest)
    - [GetAppAuthByAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppResourceMethodResponse)
    - [GetAppAuthByAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest)
    - [GetAppAuthByAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse)
    - [GetAppAuthByAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest)
    - [GetAppAuthByAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse)
    - [GetAppAuthByOtherAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest)
    - [GetAppAuthByOtherAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse)
    - [GetAppAuthByOtherAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest)
    - [GetAppAuthByOtherAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse)
    - [GetAppAuthByOtherAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest)
    - [GetAppAuthByOtherAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse)
    - [GetAuthHistoriesByAppRequest](#authing.gateway.v1.GetAuthHistoriesByAppRequest)
    - [GetAuthHistoriesByAppResponse](#authing.gateway.v1.GetAuthHistoriesByAppResponse)
    - [GetAuthHistoriesByOtherAppRequest](#authing.gateway.v1.GetAuthHistoriesByOtherAppRequest)
    - [GetAuthHistoriesByOtherAppResponse](#authing.gateway.v1.GetAuthHistoriesByOtherAppResponse)
    - [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest)
    - [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse)
    - [GetAuthsByAppRequest](#authing.gateway.v1.GetAuthsByAppRequest)
    - [GetAuthsByAppResponse](#authing.gateway.v1.GetAuthsByAppResponse)
    - [GetAuthsByAppRoleRequest](#authing.gateway.v1.GetAuthsByAppRoleRequest)
    - [GetAuthsByAppRoleResponse](#authing.gateway.v1.GetAuthsByAppRoleResponse)
    - [GetAuthsByAppUserRequest](#authing.gateway.v1.GetAuthsByAppUserRequest)
    - [GetAuthsByAppUserResponse](#authing.gateway.v1.GetAuthsByAppUserResponse)
    - [GetAuthsByOtherAppRequest](#authing.gateway.v1.GetAuthsByOtherAppRequest)
    - [GetAuthsByOtherAppResponse](#authing.gateway.v1.GetAuthsByOtherAppResponse)
    - [GetAuthsByOtherAppRoleRequest](#authing.gateway.v1.GetAuthsByOtherAppRoleRequest)
    - [GetAuthsByOtherAppRoleResponse](#authing.gateway.v1.GetAuthsByOtherAppRoleResponse)
    - [GetAuthsByOtherAppUserRequest](#authing.gateway.v1.GetAuthsByOtherAppUserRequest)
    - [GetAuthsByOtherAppUserResponse](#authing.gateway.v1.GetAuthsByOtherAppUserResponse)
  
    - [AuthingGateway](#authing.gateway.v1.AuthingGateway)
  
- [npool/authinggateway/authinggateway.proto](#npool/authinggateway/authinggateway.proto)
    - [AppAuth](#authing.gateway.v1.AppAuth)
    - [AppRoleAuth](#authing.gateway.v1.AppRoleAuth)
    - [AppUserAuth](#authing.gateway.v1.AppUserAuth)
    - [Auth](#authing.gateway.v1.Auth)
    - [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest)
    - [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse)
    - [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest)
    - [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse)
    - [AuthHistory](#authing.gateway.v1.AuthHistory)
    - [CreateAppAuthForOtherAppRequest](#authing.gateway.v1.CreateAppAuthForOtherAppRequest)
    - [CreateAppAuthForOtherAppResponse](#authing.gateway.v1.CreateAppAuthForOtherAppResponse)
    - [CreateAppRoleAuthForOtherAppRequest](#authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest)
    - [CreateAppRoleAuthForOtherAppResponse](#authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse)
    - [CreateAppRoleAuthRequest](#authing.gateway.v1.CreateAppRoleAuthRequest)
    - [CreateAppRoleAuthResponse](#authing.gateway.v1.CreateAppRoleAuthResponse)
    - [CreateAppUserAuthForOtherAppRequest](#authing.gateway.v1.CreateAppUserAuthForOtherAppRequest)
    - [CreateAppUserAuthForOtherAppResponse](#authing.gateway.v1.CreateAppUserAuthForOtherAppResponse)
    - [CreateAppUserAuthRequest](#authing.gateway.v1.CreateAppUserAuthRequest)
    - [CreateAppUserAuthResponse](#authing.gateway.v1.CreateAppUserAuthResponse)
    - [DeleteAppAuthRequest](#authing.gateway.v1.DeleteAppAuthRequest)
    - [DeleteAppAuthResponse](#authing.gateway.v1.DeleteAppAuthResponse)
    - [DeleteAppRoleAuthRequest](#authing.gateway.v1.DeleteAppRoleAuthRequest)
    - [DeleteAppRoleAuthResponse](#authing.gateway.v1.DeleteAppRoleAuthResponse)
    - [DeleteAppUserAuthRequest](#authing.gateway.v1.DeleteAppUserAuthRequest)
    - [DeleteAppUserAuthResponse](#authing.gateway.v1.DeleteAppUserAuthResponse)
    - [GetAppAuthByAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppResourceMethodRequest)
    - [GetAppAuthByAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppResourceMethodResponse)
    - [GetAppAuthByAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest)
    - [GetAppAuthByAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse)
    - [GetAppAuthByAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest)
    - [GetAppAuthByAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse)
    - [GetAppAuthByOtherAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest)
    - [GetAppAuthByOtherAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse)
    - [GetAppAuthByOtherAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest)
    - [GetAppAuthByOtherAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse)
    - [GetAppAuthByOtherAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest)
    - [GetAppAuthByOtherAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse)
    - [GetAuthHistoriesByAppRequest](#authing.gateway.v1.GetAuthHistoriesByAppRequest)
    - [GetAuthHistoriesByAppResponse](#authing.gateway.v1.GetAuthHistoriesByAppResponse)
    - [GetAuthHistoriesByOtherAppRequest](#authing.gateway.v1.GetAuthHistoriesByOtherAppRequest)
    - [GetAuthHistoriesByOtherAppResponse](#authing.gateway.v1.GetAuthHistoriesByOtherAppResponse)
    - [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest)
    - [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse)
    - [GetAuthsByAppRequest](#authing.gateway.v1.GetAuthsByAppRequest)
    - [GetAuthsByAppResponse](#authing.gateway.v1.GetAuthsByAppResponse)
    - [GetAuthsByAppRoleRequest](#authing.gateway.v1.GetAuthsByAppRoleRequest)
    - [GetAuthsByAppRoleResponse](#authing.gateway.v1.GetAuthsByAppRoleResponse)
    - [GetAuthsByAppUserRequest](#authing.gateway.v1.GetAuthsByAppUserRequest)
    - [GetAuthsByAppUserResponse](#authing.gateway.v1.GetAuthsByAppUserResponse)
    - [GetAuthsByOtherAppRequest](#authing.gateway.v1.GetAuthsByOtherAppRequest)
    - [GetAuthsByOtherAppResponse](#authing.gateway.v1.GetAuthsByOtherAppResponse)
    - [GetAuthsByOtherAppRoleRequest](#authing.gateway.v1.GetAuthsByOtherAppRoleRequest)
    - [GetAuthsByOtherAppRoleResponse](#authing.gateway.v1.GetAuthsByOtherAppRoleResponse)
    - [GetAuthsByOtherAppUserRequest](#authing.gateway.v1.GetAuthsByOtherAppUserRequest)
    - [GetAuthsByOtherAppUserResponse](#authing.gateway.v1.GetAuthsByOtherAppUserResponse)
  
    - [AuthingGateway](#authing.gateway.v1.AuthingGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/authinggateway/authinggateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/authinggateway/authinggateway.proto



<a name="authing.gateway.v1.AppAuth"></a>

### AppAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AppRoleAuth"></a>

### AppRoleAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AppUserAuth"></a>

### AppUserAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppRequest"></a>

### AuthByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppResponse"></a>

### AuthByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allowed | [bool](#bool) |  |  |






<a name="authing.gateway.v1.AuthByAppRoleUserRequest"></a>

### AuthByAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Token | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppRoleUserResponse"></a>

### AuthByAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allowed | [bool](#bool) |  |  |






<a name="authing.gateway.v1.AuthHistory"></a>

### AuthHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |
| Allowed | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="authing.gateway.v1.CreateAppAuthForOtherAppRequest"></a>

### CreateAppAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppAuth](#authing.gateway.v1.AppAuth) |  |  |






<a name="authing.gateway.v1.CreateAppAuthForOtherAppResponse"></a>

### CreateAppAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppAuth](#authing.gateway.v1.AppAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest"></a>

### CreateAppRoleAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse"></a>

### CreateAppRoleAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthRequest"></a>

### CreateAppRoleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthResponse"></a>

### CreateAppRoleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthForOtherAppRequest"></a>

### CreateAppUserAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthForOtherAppResponse"></a>

### CreateAppUserAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthRequest"></a>

### CreateAppUserAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthResponse"></a>

### CreateAppUserAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.DeleteAppAuthRequest"></a>

### DeleteAppAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppAuthResponse"></a>

### DeleteAppAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.DeleteAppRoleAuthRequest"></a>

### DeleteAppRoleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppRoleAuthResponse"></a>

### DeleteAppRoleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.DeleteAppUserAuthRequest"></a>

### DeleteAppUserAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppUserAuthResponse"></a>

### DeleteAppUserAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppResourceMethodRequest"></a>

### GetAppAuthByAppResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppResourceMethodResponse"></a>

### GetAppAuthByAppResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest"></a>

### GetAppAuthByAppRoleResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse"></a>

### GetAppAuthByAppRoleResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest"></a>

### GetAppAuthByAppUserResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse"></a>

### GetAppAuthByAppUserResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest"></a>

### GetAppAuthByOtherAppResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse"></a>

### GetAppAuthByOtherAppResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest"></a>

### GetAppAuthByOtherAppRoleResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse"></a>

### GetAppAuthByOtherAppRoleResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest"></a>

### GetAppAuthByOtherAppUserResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse"></a>

### GetAppAuthByOtherAppUserResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByAppRequest"></a>

### GetAuthHistoriesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByAppResponse"></a>

### GetAuthHistoriesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthHistoriesByOtherAppRequest"></a>

### GetAuthHistoriesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByOtherAppResponse"></a>

### GetAuthHistoriesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthHistoriesRequest"></a>

### GetAuthHistoriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesResponse"></a>

### GetAuthHistoriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppRequest"></a>

### GetAuthsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppResponse"></a>

### GetAuthsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppRoleRequest"></a>

### GetAuthsByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppRoleResponse"></a>

### GetAuthsByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppUserRequest"></a>

### GetAuthsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppUserResponse"></a>

### GetAuthsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRequest"></a>

### GetAuthsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppResponse"></a>

### GetAuthsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRoleRequest"></a>

### GetAuthsByOtherAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRoleResponse"></a>

### GetAuthsByOtherAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppUserRequest"></a>

### GetAuthsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppUserResponse"></a>

### GetAuthsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |





 

 

 


<a name="authing.gateway.v1.AuthingGateway"></a>

### AuthingGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| AuthByApp | [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest) | [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse) |  |
| AuthByAppRoleUser | [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest) | [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse) |  |
| GetAuthHistories | [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest) | [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse) |  |
| GetAuthHistoriesByApp | [GetAuthHistoriesByAppRequest](#authing.gateway.v1.GetAuthHistoriesByAppRequest) | [GetAuthHistoriesByAppResponse](#authing.gateway.v1.GetAuthHistoriesByAppResponse) |  |
| GetAuthHistoriesByOtherApp | [GetAuthHistoriesByOtherAppRequest](#authing.gateway.v1.GetAuthHistoriesByOtherAppRequest) | [GetAuthHistoriesByOtherAppResponse](#authing.gateway.v1.GetAuthHistoriesByOtherAppResponse) |  |
| CreateAppAuthForOtherApp | [CreateAppAuthForOtherAppRequest](#authing.gateway.v1.CreateAppAuthForOtherAppRequest) | [CreateAppAuthForOtherAppResponse](#authing.gateway.v1.CreateAppAuthForOtherAppResponse) |  |
| GetAppAuthByAppResourceMethod | [GetAppAuthByAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppResourceMethodRequest) | [GetAppAuthByAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppResourceMethodResponse) |  |
| GetAppAuthByOtherAppResourceMethod | [GetAppAuthByOtherAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest) | [GetAppAuthByOtherAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse) |  |
| DeleteAppAuth | [DeleteAppAuthRequest](#authing.gateway.v1.DeleteAppAuthRequest) | [DeleteAppAuthResponse](#authing.gateway.v1.DeleteAppAuthResponse) |  |
| CreateAppRoleAuth | [CreateAppRoleAuthRequest](#authing.gateway.v1.CreateAppRoleAuthRequest) | [CreateAppRoleAuthResponse](#authing.gateway.v1.CreateAppRoleAuthResponse) |  |
| CreateAppRoleAuthForOtherApp | [CreateAppRoleAuthForOtherAppRequest](#authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest) | [CreateAppRoleAuthForOtherAppResponse](#authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse) |  |
| GetAppAuthByAppRoleResourceMethod | [GetAppAuthByAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest) | [GetAppAuthByAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse) |  |
| GetAppAuthByOtherAppRoleResourceMethod | [GetAppAuthByOtherAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest) | [GetAppAuthByOtherAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse) |  |
| DeleteAppRoleAuth | [DeleteAppRoleAuthRequest](#authing.gateway.v1.DeleteAppRoleAuthRequest) | [DeleteAppRoleAuthResponse](#authing.gateway.v1.DeleteAppRoleAuthResponse) |  |
| CreateAppUserAuth | [CreateAppUserAuthRequest](#authing.gateway.v1.CreateAppUserAuthRequest) | [CreateAppUserAuthResponse](#authing.gateway.v1.CreateAppUserAuthResponse) |  |
| CreateAppUserAuthForOtherApp | [CreateAppUserAuthForOtherAppRequest](#authing.gateway.v1.CreateAppUserAuthForOtherAppRequest) | [CreateAppUserAuthForOtherAppResponse](#authing.gateway.v1.CreateAppUserAuthForOtherAppResponse) |  |
| GetAppAuthByAppUserResourceMethod | [GetAppAuthByAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest) | [GetAppAuthByAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse) |  |
| GetAppAuthByOtherAppUserResourceMethod | [GetAppAuthByOtherAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest) | [GetAppAuthByOtherAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse) |  |
| DeleteAppUserAuth | [DeleteAppUserAuthRequest](#authing.gateway.v1.DeleteAppUserAuthRequest) | [DeleteAppUserAuthResponse](#authing.gateway.v1.DeleteAppUserAuthResponse) |  |
| GetAuthsByApp | [GetAuthsByAppRequest](#authing.gateway.v1.GetAuthsByAppRequest) | [GetAuthsByAppResponse](#authing.gateway.v1.GetAuthsByAppResponse) |  |
| GetAuthsByOtherApp | [GetAuthsByOtherAppRequest](#authing.gateway.v1.GetAuthsByOtherAppRequest) | [GetAuthsByOtherAppResponse](#authing.gateway.v1.GetAuthsByOtherAppResponse) |  |
| GetAuthsByAppRole | [GetAuthsByAppRoleRequest](#authing.gateway.v1.GetAuthsByAppRoleRequest) | [GetAuthsByAppRoleResponse](#authing.gateway.v1.GetAuthsByAppRoleResponse) |  |
| GetAuthsByOtherAppRole | [GetAuthsByOtherAppRoleRequest](#authing.gateway.v1.GetAuthsByOtherAppRoleRequest) | [GetAuthsByOtherAppRoleResponse](#authing.gateway.v1.GetAuthsByOtherAppRoleResponse) |  |
| GetAuthsByAppUser | [GetAuthsByAppUserRequest](#authing.gateway.v1.GetAuthsByAppUserRequest) | [GetAuthsByAppUserResponse](#authing.gateway.v1.GetAuthsByAppUserResponse) |  |
| GetAuthsByOtherAppUser | [GetAuthsByOtherAppUserRequest](#authing.gateway.v1.GetAuthsByOtherAppUserRequest) | [GetAuthsByOtherAppUserResponse](#authing.gateway.v1.GetAuthsByOtherAppUserResponse) |  |

 



<a name="npool/authinggateway/authinggateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/authinggateway/authinggateway.proto



<a name="authing.gateway.v1.AppAuth"></a>

### AppAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AppRoleAuth"></a>

### AppRoleAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AppUserAuth"></a>

### AppUserAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppRequest"></a>

### AuthByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppResponse"></a>

### AuthByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allowed | [bool](#bool) |  |  |






<a name="authing.gateway.v1.AuthByAppRoleUserRequest"></a>

### AuthByAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Token | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.AuthByAppRoleUserResponse"></a>

### AuthByAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Allowed | [bool](#bool) |  |  |






<a name="authing.gateway.v1.AuthHistory"></a>

### AuthHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |
| Allowed | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="authing.gateway.v1.CreateAppAuthForOtherAppRequest"></a>

### CreateAppAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppAuth](#authing.gateway.v1.AppAuth) |  |  |






<a name="authing.gateway.v1.CreateAppAuthForOtherAppResponse"></a>

### CreateAppAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppAuth](#authing.gateway.v1.AppAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest"></a>

### CreateAppRoleAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse"></a>

### CreateAppRoleAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthRequest"></a>

### CreateAppRoleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppRoleAuthResponse"></a>

### CreateAppRoleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleAuth](#authing.gateway.v1.AppRoleAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthForOtherAppRequest"></a>

### CreateAppUserAuthForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthForOtherAppResponse"></a>

### CreateAppUserAuthForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthRequest"></a>

### CreateAppUserAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.CreateAppUserAuthResponse"></a>

### CreateAppUserAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserAuth](#authing.gateway.v1.AppUserAuth) |  |  |






<a name="authing.gateway.v1.DeleteAppAuthRequest"></a>

### DeleteAppAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppAuthResponse"></a>

### DeleteAppAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.DeleteAppRoleAuthRequest"></a>

### DeleteAppRoleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppRoleAuthResponse"></a>

### DeleteAppRoleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.DeleteAppUserAuthRequest"></a>

### DeleteAppUserAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="authing.gateway.v1.DeleteAppUserAuthResponse"></a>

### DeleteAppUserAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppResourceMethodRequest"></a>

### GetAppAuthByAppResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppResourceMethodResponse"></a>

### GetAppAuthByAppResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest"></a>

### GetAppAuthByAppRoleResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse"></a>

### GetAppAuthByAppRoleResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest"></a>

### GetAppAuthByAppUserResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse"></a>

### GetAppAuthByAppUserResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest"></a>

### GetAppAuthByOtherAppResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse"></a>

### GetAppAuthByOtherAppResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest"></a>

### GetAppAuthByOtherAppRoleResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse"></a>

### GetAppAuthByOtherAppRoleResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest"></a>

### GetAppAuthByOtherAppUserResourceMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Resource | [string](#string) |  |  |
| Method | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse"></a>

### GetAppAuthByOtherAppUserResourceMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#authing.gateway.v1.Auth) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByAppRequest"></a>

### GetAuthHistoriesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByAppResponse"></a>

### GetAuthHistoriesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthHistoriesByOtherAppRequest"></a>

### GetAuthHistoriesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesByOtherAppResponse"></a>

### GetAuthHistoriesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthHistoriesRequest"></a>

### GetAuthHistoriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthHistoriesResponse"></a>

### GetAuthHistoriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AuthHistory](#authing.gateway.v1.AuthHistory) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppRequest"></a>

### GetAuthsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppResponse"></a>

### GetAuthsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppRoleRequest"></a>

### GetAuthsByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppRoleResponse"></a>

### GetAuthsByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByAppUserRequest"></a>

### GetAuthsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByAppUserResponse"></a>

### GetAuthsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRequest"></a>

### GetAuthsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppResponse"></a>

### GetAuthsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRoleRequest"></a>

### GetAuthsByOtherAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppRoleResponse"></a>

### GetAuthsByOtherAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppUserRequest"></a>

### GetAuthsByOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |






<a name="authing.gateway.v1.GetAuthsByOtherAppUserResponse"></a>

### GetAuthsByOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#authing.gateway.v1.Auth) | repeated |  |





 

 

 


<a name="authing.gateway.v1.AuthingGateway"></a>

### AuthingGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| AuthByApp | [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest) | [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse) |  |
| AuthByAppRoleUser | [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest) | [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse) |  |
| GetAuthHistories | [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest) | [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse) |  |
| GetAuthHistoriesByApp | [GetAuthHistoriesByAppRequest](#authing.gateway.v1.GetAuthHistoriesByAppRequest) | [GetAuthHistoriesByAppResponse](#authing.gateway.v1.GetAuthHistoriesByAppResponse) |  |
| GetAuthHistoriesByOtherApp | [GetAuthHistoriesByOtherAppRequest](#authing.gateway.v1.GetAuthHistoriesByOtherAppRequest) | [GetAuthHistoriesByOtherAppResponse](#authing.gateway.v1.GetAuthHistoriesByOtherAppResponse) |  |
| CreateAppAuthForOtherApp | [CreateAppAuthForOtherAppRequest](#authing.gateway.v1.CreateAppAuthForOtherAppRequest) | [CreateAppAuthForOtherAppResponse](#authing.gateway.v1.CreateAppAuthForOtherAppResponse) |  |
| GetAppAuthByAppResourceMethod | [GetAppAuthByAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppResourceMethodRequest) | [GetAppAuthByAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppResourceMethodResponse) |  |
| GetAppAuthByOtherAppResourceMethod | [GetAppAuthByOtherAppResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodRequest) | [GetAppAuthByOtherAppResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppResourceMethodResponse) |  |
| DeleteAppAuth | [DeleteAppAuthRequest](#authing.gateway.v1.DeleteAppAuthRequest) | [DeleteAppAuthResponse](#authing.gateway.v1.DeleteAppAuthResponse) |  |
| CreateAppRoleAuth | [CreateAppRoleAuthRequest](#authing.gateway.v1.CreateAppRoleAuthRequest) | [CreateAppRoleAuthResponse](#authing.gateway.v1.CreateAppRoleAuthResponse) |  |
| CreateAppRoleAuthForOtherApp | [CreateAppRoleAuthForOtherAppRequest](#authing.gateway.v1.CreateAppRoleAuthForOtherAppRequest) | [CreateAppRoleAuthForOtherAppResponse](#authing.gateway.v1.CreateAppRoleAuthForOtherAppResponse) |  |
| GetAppAuthByAppRoleResourceMethod | [GetAppAuthByAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodRequest) | [GetAppAuthByAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppRoleResourceMethodResponse) |  |
| GetAppAuthByOtherAppRoleResourceMethod | [GetAppAuthByOtherAppRoleResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodRequest) | [GetAppAuthByOtherAppRoleResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppRoleResourceMethodResponse) |  |
| DeleteAppRoleAuth | [DeleteAppRoleAuthRequest](#authing.gateway.v1.DeleteAppRoleAuthRequest) | [DeleteAppRoleAuthResponse](#authing.gateway.v1.DeleteAppRoleAuthResponse) |  |
| CreateAppUserAuth | [CreateAppUserAuthRequest](#authing.gateway.v1.CreateAppUserAuthRequest) | [CreateAppUserAuthResponse](#authing.gateway.v1.CreateAppUserAuthResponse) |  |
| CreateAppUserAuthForOtherApp | [CreateAppUserAuthForOtherAppRequest](#authing.gateway.v1.CreateAppUserAuthForOtherAppRequest) | [CreateAppUserAuthForOtherAppResponse](#authing.gateway.v1.CreateAppUserAuthForOtherAppResponse) |  |
| GetAppAuthByAppUserResourceMethod | [GetAppAuthByAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodRequest) | [GetAppAuthByAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByAppUserResourceMethodResponse) |  |
| GetAppAuthByOtherAppUserResourceMethod | [GetAppAuthByOtherAppUserResourceMethodRequest](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodRequest) | [GetAppAuthByOtherAppUserResourceMethodResponse](#authing.gateway.v1.GetAppAuthByOtherAppUserResourceMethodResponse) |  |
| DeleteAppUserAuth | [DeleteAppUserAuthRequest](#authing.gateway.v1.DeleteAppUserAuthRequest) | [DeleteAppUserAuthResponse](#authing.gateway.v1.DeleteAppUserAuthResponse) |  |
| GetAuthsByApp | [GetAuthsByAppRequest](#authing.gateway.v1.GetAuthsByAppRequest) | [GetAuthsByAppResponse](#authing.gateway.v1.GetAuthsByAppResponse) |  |
| GetAuthsByOtherApp | [GetAuthsByOtherAppRequest](#authing.gateway.v1.GetAuthsByOtherAppRequest) | [GetAuthsByOtherAppResponse](#authing.gateway.v1.GetAuthsByOtherAppResponse) |  |
| GetAuthsByAppRole | [GetAuthsByAppRoleRequest](#authing.gateway.v1.GetAuthsByAppRoleRequest) | [GetAuthsByAppRoleResponse](#authing.gateway.v1.GetAuthsByAppRoleResponse) |  |
| GetAuthsByOtherAppRole | [GetAuthsByOtherAppRoleRequest](#authing.gateway.v1.GetAuthsByOtherAppRoleRequest) | [GetAuthsByOtherAppRoleResponse](#authing.gateway.v1.GetAuthsByOtherAppRoleResponse) |  |
| GetAuthsByAppUser | [GetAuthsByAppUserRequest](#authing.gateway.v1.GetAuthsByAppUserRequest) | [GetAuthsByAppUserResponse](#authing.gateway.v1.GetAuthsByAppUserResponse) |  |
| GetAuthsByOtherAppUser | [GetAuthsByOtherAppUserRequest](#authing.gateway.v1.GetAuthsByOtherAppUserRequest) | [GetAuthsByOtherAppUserResponse](#authing.gateway.v1.GetAuthsByOtherAppUserResponse) |  |

 



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

