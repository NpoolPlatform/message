# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/appusermgr/appusermgr.proto](#npool/appusermgr/appusermgr.proto)
    - [App](#app.user.manager.v1.App)
    - [AppControl](#app.user.manager.v1.AppControl)
    - [AppInfo](#app.user.manager.v1.AppInfo)
    - [AppRole](#app.user.manager.v1.AppRole)
    - [AppRoleUser](#app.user.manager.v1.AppRoleUser)
    - [AppUser](#app.user.manager.v1.AppUser)
    - [AppUserControl](#app.user.manager.v1.AppUserControl)
    - [AppUserExtra](#app.user.manager.v1.AppUserExtra)
    - [AppUserSecret](#app.user.manager.v1.AppUserSecret)
    - [BanApp](#app.user.manager.v1.BanApp)
    - [BanAppUser](#app.user.manager.v1.BanAppUser)
    - [CreateAppControlRequest](#app.user.manager.v1.CreateAppControlRequest)
    - [CreateAppControlResponse](#app.user.manager.v1.CreateAppControlResponse)
    - [CreateAppRequest](#app.user.manager.v1.CreateAppRequest)
    - [CreateAppResponse](#app.user.manager.v1.CreateAppResponse)
    - [CreateAppRoleRequest](#app.user.manager.v1.CreateAppRoleRequest)
    - [CreateAppRoleResponse](#app.user.manager.v1.CreateAppRoleResponse)
    - [CreateAppRoleUserRequest](#app.user.manager.v1.CreateAppRoleUserRequest)
    - [CreateAppRoleUserResponse](#app.user.manager.v1.CreateAppRoleUserResponse)
    - [CreateAppUserControlRequest](#app.user.manager.v1.CreateAppUserControlRequest)
    - [CreateAppUserControlResponse](#app.user.manager.v1.CreateAppUserControlResponse)
    - [CreateAppUserExtraRequest](#app.user.manager.v1.CreateAppUserExtraRequest)
    - [CreateAppUserExtraResponse](#app.user.manager.v1.CreateAppUserExtraResponse)
    - [CreateAppUserRequest](#app.user.manager.v1.CreateAppUserRequest)
    - [CreateAppUserResponse](#app.user.manager.v1.CreateAppUserResponse)
    - [CreateAppUserSecretRequest](#app.user.manager.v1.CreateAppUserSecretRequest)
    - [CreateAppUserSecretResponse](#app.user.manager.v1.CreateAppUserSecretResponse)
    - [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest)
    - [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse)
    - [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest)
    - [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse)
    - [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest)
    - [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse)
    - [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest)
    - [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse)
    - [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest)
    - [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse)
    - [GetAppRequest](#app.user.manager.v1.GetAppRequest)
    - [GetAppResponse](#app.user.manager.v1.GetAppResponse)
    - [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest)
    - [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse)
    - [GetAppsRequest](#app.user.manager.v1.GetAppsRequest)
    - [GetAppsResponse](#app.user.manager.v1.GetAppsResponse)
    - [UpdateAppControlRequest](#app.user.manager.v1.UpdateAppControlRequest)
    - [UpdateAppControlResponse](#app.user.manager.v1.UpdateAppControlResponse)
    - [UpdateAppRequest](#app.user.manager.v1.UpdateAppRequest)
    - [UpdateAppResponse](#app.user.manager.v1.UpdateAppResponse)
    - [UpdateAppRoleRequest](#app.user.manager.v1.UpdateAppRoleRequest)
    - [UpdateAppRoleResponse](#app.user.manager.v1.UpdateAppRoleResponse)
    - [UpdateAppUserControlRequest](#app.user.manager.v1.UpdateAppUserControlRequest)
    - [UpdateAppUserControlResponse](#app.user.manager.v1.UpdateAppUserControlResponse)
    - [UpdateAppUserExtraRequest](#app.user.manager.v1.UpdateAppUserExtraRequest)
    - [UpdateAppUserExtraResponse](#app.user.manager.v1.UpdateAppUserExtraResponse)
    - [UpdateAppUserRequest](#app.user.manager.v1.UpdateAppUserRequest)
    - [UpdateAppUserResponse](#app.user.manager.v1.UpdateAppUserResponse)
    - [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest)
    - [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse)
  
    - [AppUserManager](#app.user.manager.v1.AppUserManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/appusermgr/appusermgr.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/appusermgr/appusermgr.proto



<a name="app.user.manager.v1.App"></a>

### App



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Description | [string](#string) |  |  |






<a name="app.user.manager.v1.AppControl"></a>

### AppControl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| SignupMethods | [string](#string) | repeated | Signup method always be signin method |
| ExternSigninMethods | [string](#string) | repeated |  |
| RecaptchaMethod | [string](#string) |  |  |
| KycEnable | [bool](#bool) |  |  |
| SigninVerifyEnable | [bool](#bool) |  |  |






<a name="app.user.manager.v1.AppInfo"></a>

### AppInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| App | [App](#app.user.manager.v1.App) |  |  |
| Ctrl | [AppControl](#app.user.manager.v1.AppControl) |  |  |
| Ban | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.AppRole"></a>

### AppRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| CreatedBy | [string](#string) |  |  |
| Role | [string](#string) |  |  |
| Description | [string](#string) |  |  |
| Default | [bool](#bool) |  |  |






<a name="app.user.manager.v1.AppRoleUser"></a>

### AppRoleUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUser"></a>

### AppUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |
| PhoneNO | [string](#string) |  |  |
| ImportFromApp | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUserControl"></a>

### AppUserControl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUserExtra"></a>

### AppUserExtra



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| AddressFields | [string](#string) | repeated |  |
| Gender | [string](#string) |  |  |
| PostalCode | [string](#string) |  |  |
| Age | [uint32](#uint32) |  |  |
| Birthday | [uint32](#uint32) |  |  |
| Avatar | [string](#string) |  |  |
| Organization | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUserSecret"></a>

### AppUserSecret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |
| Salt | [string](#string) |  |  |
| GoogleSecret | [string](#string) |  |  |






<a name="app.user.manager.v1.BanApp"></a>

### BanApp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="app.user.manager.v1.BanAppUser"></a>

### BanAppUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.CreateAppControlRequest"></a>

### CreateAppControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.CreateAppControlResponse"></a>

### CreateAppControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.CreateAppRequest"></a>

### CreateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.CreateAppResponse"></a>

### CreateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.CreateAppRoleRequest"></a>

### CreateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.CreateAppRoleResponse"></a>

### CreateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.CreateAppRoleUserRequest"></a>

### CreateAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.CreateAppRoleUserResponse"></a>

### CreateAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.CreateAppUserControlRequest"></a>

### CreateAppUserControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.CreateAppUserControlResponse"></a>

### CreateAppUserControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.CreateAppUserExtraRequest"></a>

### CreateAppUserExtraRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.CreateAppUserExtraResponse"></a>

### CreateAppUserExtraResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.CreateAppUserRequest"></a>

### CreateAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.CreateAppUserResponse"></a>

### CreateAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.CreateAppUserSecretRequest"></a>

### CreateAppUserSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateAppUserSecretResponse"></a>

### CreateAppUserSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateBanAppRequest"></a>

### CreateBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.CreateBanAppResponse"></a>

### CreateBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.CreateBanAppUserRequest"></a>

### CreateBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.CreateBanAppUserResponse"></a>

### CreateBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.DeleteAppRoleUserRequest"></a>

### DeleteAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.DeleteAppRoleUserResponse"></a>

### DeleteAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.DeleteBanAppRequest"></a>

### DeleteBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.DeleteBanAppResponse"></a>

### DeleteBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.DeleteBanAppUserRequest"></a>

### DeleteBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.DeleteBanAppUserResponse"></a>

### DeleteBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetAppRequest"></a>

### GetAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppResponse"></a>

### GetAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInfo](#app.user.manager.v1.AppInfo) |  |  |






<a name="app.user.manager.v1.GetAppsByCreatorRequest"></a>

### GetAppsByCreatorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppsByCreatorResponse"></a>

### GetAppsByCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppsRequest"></a>

### GetAppsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppsResponse"></a>

### GetAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.UpdateAppControlRequest"></a>

### UpdateAppControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.UpdateAppControlResponse"></a>

### UpdateAppControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.UpdateAppRequest"></a>

### UpdateAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.UpdateAppResponse"></a>

### UpdateAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.UpdateAppRoleRequest"></a>

### UpdateAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.UpdateAppRoleResponse"></a>

### UpdateAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.UpdateAppUserControlRequest"></a>

### UpdateAppUserControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.UpdateAppUserControlResponse"></a>

### UpdateAppUserControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.UpdateAppUserExtraRequest"></a>

### UpdateAppUserExtraRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.UpdateAppUserExtraResponse"></a>

### UpdateAppUserExtraResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.UpdateAppUserRequest"></a>

### UpdateAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.UpdateAppUserResponse"></a>

### UpdateAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.UpdateAppUserSecretRequest"></a>

### UpdateAppUserSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.UpdateAppUserSecretResponse"></a>

### UpdateAppUserSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |





 

 

 


<a name="app.user.manager.v1.AppUserManager"></a>

### AppUserManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| CreateApp | [CreateAppRequest](#app.user.manager.v1.CreateAppRequest) | [CreateAppResponse](#app.user.manager.v1.CreateAppResponse) |  |
| UpdateApp | [UpdateAppRequest](#app.user.manager.v1.UpdateAppRequest) | [UpdateAppResponse](#app.user.manager.v1.UpdateAppResponse) |  |
| CreateAppControl | [CreateAppControlRequest](#app.user.manager.v1.CreateAppControlRequest) | [CreateAppControlResponse](#app.user.manager.v1.CreateAppControlResponse) |  |
| UpdateAppControl | [UpdateAppControlRequest](#app.user.manager.v1.UpdateAppControlRequest) | [UpdateAppControlResponse](#app.user.manager.v1.UpdateAppControlResponse) |  |
| CreateBanApp | [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest) | [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse) |  |
| DeleteBanApp | [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest) | [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse) |  |
| GetApp | [GetAppRequest](#app.user.manager.v1.GetAppRequest) | [GetAppResponse](#app.user.manager.v1.GetAppResponse) |  |
| GetApps | [GetAppsRequest](#app.user.manager.v1.GetAppsRequest) | [GetAppsResponse](#app.user.manager.v1.GetAppsResponse) |  |
| GetAppsByCreator | [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest) | [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse) |  |
| CreateAppUser | [CreateAppUserRequest](#app.user.manager.v1.CreateAppUserRequest) | [CreateAppUserResponse](#app.user.manager.v1.CreateAppUserResponse) |  |
| UpdateAppUser | [UpdateAppUserRequest](#app.user.manager.v1.UpdateAppUserRequest) | [UpdateAppUserResponse](#app.user.manager.v1.UpdateAppUserResponse) |  |
| CreateAppUserSecret | [CreateAppUserSecretRequest](#app.user.manager.v1.CreateAppUserSecretRequest) | [CreateAppUserSecretResponse](#app.user.manager.v1.CreateAppUserSecretResponse) |  |
| UpdateAppUserSecret | [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest) | [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse) |  |
| CreateAppUserExtra | [CreateAppUserExtraRequest](#app.user.manager.v1.CreateAppUserExtraRequest) | [CreateAppUserExtraResponse](#app.user.manager.v1.CreateAppUserExtraResponse) |  |
| UpdateAppUserExtra | [UpdateAppUserExtraRequest](#app.user.manager.v1.UpdateAppUserExtraRequest) | [UpdateAppUserExtraResponse](#app.user.manager.v1.UpdateAppUserExtraResponse) |  |
| CreateBanAppUser | [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest) | [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse) |  |
| DeleteBanAppUser | [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest) | [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse) |  |
| CreateAppUserControl | [CreateAppUserControlRequest](#app.user.manager.v1.CreateAppUserControlRequest) | [CreateAppUserControlResponse](#app.user.manager.v1.CreateAppUserControlResponse) |  |
| UpdateAppUserControl | [UpdateAppUserControlRequest](#app.user.manager.v1.UpdateAppUserControlRequest) | [UpdateAppUserControlResponse](#app.user.manager.v1.UpdateAppUserControlResponse) |  |
| CreateAppRole | [CreateAppRoleRequest](#app.user.manager.v1.CreateAppRoleRequest) | [CreateAppRoleResponse](#app.user.manager.v1.CreateAppRoleResponse) |  |
| UpdateAppRole | [UpdateAppRoleRequest](#app.user.manager.v1.UpdateAppRoleRequest) | [UpdateAppRoleResponse](#app.user.manager.v1.UpdateAppRoleResponse) |  |
| CreateAppRoleUser | [CreateAppRoleUserRequest](#app.user.manager.v1.CreateAppRoleUserRequest) | [CreateAppRoleUserResponse](#app.user.manager.v1.CreateAppRoleUserResponse) |  |
| DeleteAppRoleUser | [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest) | [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse) |  |

 



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

