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
    - [AppUserInfo](#app.user.manager.v1.AppUserInfo)
    - [AppUserSecret](#app.user.manager.v1.AppUserSecret)
    - [BanApp](#app.user.manager.v1.BanApp)
    - [BanAppUser](#app.user.manager.v1.BanAppUser)
    - [CreateAdminAppsRequest](#app.user.manager.v1.CreateAdminAppsRequest)
    - [CreateAdminAppsResponse](#app.user.manager.v1.CreateAdminAppsResponse)
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
    - [CreateAppUserWithSecretRequest](#app.user.manager.v1.CreateAppUserWithSecretRequest)
    - [CreateAppUserWithSecretResponse](#app.user.manager.v1.CreateAppUserWithSecretResponse)
    - [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest)
    - [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse)
    - [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest)
    - [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse)
    - [CreateGenesisRoleRequest](#app.user.manager.v1.CreateGenesisRoleRequest)
    - [CreateGenesisRoleResponse](#app.user.manager.v1.CreateGenesisRoleResponse)
    - [CreateGenesisRoleUserRequest](#app.user.manager.v1.CreateGenesisRoleUserRequest)
    - [CreateGenesisRoleUserResponse](#app.user.manager.v1.CreateGenesisRoleUserResponse)
    - [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest)
    - [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse)
    - [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest)
    - [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse)
    - [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest)
    - [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse)
    - [GetAdminAppsRequest](#app.user.manager.v1.GetAdminAppsRequest)
    - [GetAdminAppsResponse](#app.user.manager.v1.GetAdminAppsResponse)
    - [GetAppControlByAppRequest](#app.user.manager.v1.GetAppControlByAppRequest)
    - [GetAppControlByAppResponse](#app.user.manager.v1.GetAppControlByAppResponse)
    - [GetAppControlRequest](#app.user.manager.v1.GetAppControlRequest)
    - [GetAppControlResponse](#app.user.manager.v1.GetAppControlResponse)
    - [GetAppInfoRequest](#app.user.manager.v1.GetAppInfoRequest)
    - [GetAppInfoResponse](#app.user.manager.v1.GetAppInfoResponse)
    - [GetAppInfosByCreatorRequest](#app.user.manager.v1.GetAppInfosByCreatorRequest)
    - [GetAppInfosByCreatorResponse](#app.user.manager.v1.GetAppInfosByCreatorResponse)
    - [GetAppInfosRequest](#app.user.manager.v1.GetAppInfosRequest)
    - [GetAppInfosResponse](#app.user.manager.v1.GetAppInfosResponse)
    - [GetAppRequest](#app.user.manager.v1.GetAppRequest)
    - [GetAppResponse](#app.user.manager.v1.GetAppResponse)
    - [GetAppRoleByAppRoleRequest](#app.user.manager.v1.GetAppRoleByAppRoleRequest)
    - [GetAppRoleByAppRoleResponse](#app.user.manager.v1.GetAppRoleByAppRoleResponse)
    - [GetAppRoleRequest](#app.user.manager.v1.GetAppRoleRequest)
    - [GetAppRoleResponse](#app.user.manager.v1.GetAppRoleResponse)
    - [GetAppRoleUserByAppUserRequest](#app.user.manager.v1.GetAppRoleUserByAppUserRequest)
    - [GetAppRoleUserByAppUserResponse](#app.user.manager.v1.GetAppRoleUserByAppUserResponse)
    - [GetAppRoleUserRequest](#app.user.manager.v1.GetAppRoleUserRequest)
    - [GetAppRoleUserResponse](#app.user.manager.v1.GetAppRoleUserResponse)
    - [GetAppRoleUsersByAppRoleRequest](#app.user.manager.v1.GetAppRoleUsersByAppRoleRequest)
    - [GetAppRoleUsersByAppRoleResponse](#app.user.manager.v1.GetAppRoleUsersByAppRoleResponse)
    - [GetAppRolesByAppRequest](#app.user.manager.v1.GetAppRolesByAppRequest)
    - [GetAppRolesByAppResponse](#app.user.manager.v1.GetAppRolesByAppResponse)
    - [GetAppUserByAppAccountRequest](#app.user.manager.v1.GetAppUserByAppAccountRequest)
    - [GetAppUserByAppAccountResponse](#app.user.manager.v1.GetAppUserByAppAccountResponse)
    - [GetAppUserByAppUserRequest](#app.user.manager.v1.GetAppUserByAppUserRequest)
    - [GetAppUserByAppUserResponse](#app.user.manager.v1.GetAppUserByAppUserResponse)
    - [GetAppUserControlByAppUserRequest](#app.user.manager.v1.GetAppUserControlByAppUserRequest)
    - [GetAppUserControlByAppUserResponse](#app.user.manager.v1.GetAppUserControlByAppUserResponse)
    - [GetAppUserControlRequest](#app.user.manager.v1.GetAppUserControlRequest)
    - [GetAppUserControlResponse](#app.user.manager.v1.GetAppUserControlResponse)
    - [GetAppUserExtraByAppUserRequest](#app.user.manager.v1.GetAppUserExtraByAppUserRequest)
    - [GetAppUserExtraByAppUserResponse](#app.user.manager.v1.GetAppUserExtraByAppUserResponse)
    - [GetAppUserExtraRequest](#app.user.manager.v1.GetAppUserExtraRequest)
    - [GetAppUserExtraResponse](#app.user.manager.v1.GetAppUserExtraResponse)
    - [GetAppUserInfoByAppUserRequest](#app.user.manager.v1.GetAppUserInfoByAppUserRequest)
    - [GetAppUserInfoByAppUserResponse](#app.user.manager.v1.GetAppUserInfoByAppUserResponse)
    - [GetAppUserInfoRequest](#app.user.manager.v1.GetAppUserInfoRequest)
    - [GetAppUserInfoResponse](#app.user.manager.v1.GetAppUserInfoResponse)
    - [GetAppUserInfosByAppRequest](#app.user.manager.v1.GetAppUserInfosByAppRequest)
    - [GetAppUserInfosByAppResponse](#app.user.manager.v1.GetAppUserInfosByAppResponse)
    - [GetAppUserRequest](#app.user.manager.v1.GetAppUserRequest)
    - [GetAppUserResponse](#app.user.manager.v1.GetAppUserResponse)
    - [GetAppUserSecretByAppUserRequest](#app.user.manager.v1.GetAppUserSecretByAppUserRequest)
    - [GetAppUserSecretByAppUserResponse](#app.user.manager.v1.GetAppUserSecretByAppUserResponse)
    - [GetAppUserSecretRequest](#app.user.manager.v1.GetAppUserSecretRequest)
    - [GetAppUserSecretResponse](#app.user.manager.v1.GetAppUserSecretResponse)
    - [GetAppUsersByAppRequest](#app.user.manager.v1.GetAppUsersByAppRequest)
    - [GetAppUsersByAppResponse](#app.user.manager.v1.GetAppUsersByAppResponse)
    - [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest)
    - [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse)
    - [GetAppsRequest](#app.user.manager.v1.GetAppsRequest)
    - [GetAppsResponse](#app.user.manager.v1.GetAppsResponse)
    - [GetBanAppByAppRequest](#app.user.manager.v1.GetBanAppByAppRequest)
    - [GetBanAppByAppResponse](#app.user.manager.v1.GetBanAppByAppResponse)
    - [GetBanAppRequest](#app.user.manager.v1.GetBanAppRequest)
    - [GetBanAppResponse](#app.user.manager.v1.GetBanAppResponse)
    - [GetBanAppUserByAppUserRequest](#app.user.manager.v1.GetBanAppUserByAppUserRequest)
    - [GetBanAppUserByAppUserResponse](#app.user.manager.v1.GetBanAppUserByAppUserResponse)
    - [GetBanAppUserRequest](#app.user.manager.v1.GetBanAppUserRequest)
    - [GetBanAppUserResponse](#app.user.manager.v1.GetBanAppUserResponse)
    - [GetGenesisRoleRequest](#app.user.manager.v1.GetGenesisRoleRequest)
    - [GetGenesisRoleResponse](#app.user.manager.v1.GetGenesisRoleResponse)
    - [GetUserRolesByAppUserRequest](#app.user.manager.v1.GetUserRolesByAppUserRequest)
    - [GetUserRolesByAppUserResponse](#app.user.manager.v1.GetUserRolesByAppUserResponse)
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
    - [UpdateAppUserSecretByAppUserRequest](#app.user.manager.v1.UpdateAppUserSecretByAppUserRequest)
    - [UpdateAppUserSecretByAppUserResponse](#app.user.manager.v1.UpdateAppUserSecretByAppUserResponse)
    - [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest)
    - [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse)
    - [UpdateBanAppRequest](#app.user.manager.v1.UpdateBanAppRequest)
    - [UpdateBanAppResponse](#app.user.manager.v1.UpdateBanAppResponse)
    - [UpdateBanAppUserRequest](#app.user.manager.v1.UpdateBanAppUserRequest)
    - [UpdateBanAppUserResponse](#app.user.manager.v1.UpdateBanAppUserResponse)
    - [VerifyAppUserByAppAccountPasswordRequest](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest)
    - [VerifyAppUserByAppAccountPasswordResponse](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse)
  
    - [AppUserManager](#app.user.manager.v1.AppUserManager)
  
- [npool/appusermgr/appusermgr.proto](#npool/appusermgr/appusermgr.proto)
    - [App](#app.user.manager.v1.App)
    - [AppControl](#app.user.manager.v1.AppControl)
    - [AppInfo](#app.user.manager.v1.AppInfo)
    - [AppRole](#app.user.manager.v1.AppRole)
    - [AppRoleUser](#app.user.manager.v1.AppRoleUser)
    - [AppUser](#app.user.manager.v1.AppUser)
    - [AppUserControl](#app.user.manager.v1.AppUserControl)
    - [AppUserExtra](#app.user.manager.v1.AppUserExtra)
    - [AppUserInfo](#app.user.manager.v1.AppUserInfo)
    - [AppUserSecret](#app.user.manager.v1.AppUserSecret)
    - [BanApp](#app.user.manager.v1.BanApp)
    - [BanAppUser](#app.user.manager.v1.BanAppUser)
    - [CreateAdminAppsRequest](#app.user.manager.v1.CreateAdminAppsRequest)
    - [CreateAdminAppsResponse](#app.user.manager.v1.CreateAdminAppsResponse)
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
    - [CreateAppUserWithSecretRequest](#app.user.manager.v1.CreateAppUserWithSecretRequest)
    - [CreateAppUserWithSecretResponse](#app.user.manager.v1.CreateAppUserWithSecretResponse)
    - [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest)
    - [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse)
    - [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest)
    - [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse)
    - [CreateGenesisRoleRequest](#app.user.manager.v1.CreateGenesisRoleRequest)
    - [CreateGenesisRoleResponse](#app.user.manager.v1.CreateGenesisRoleResponse)
    - [CreateGenesisRoleUserRequest](#app.user.manager.v1.CreateGenesisRoleUserRequest)
    - [CreateGenesisRoleUserResponse](#app.user.manager.v1.CreateGenesisRoleUserResponse)
    - [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest)
    - [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse)
    - [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest)
    - [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse)
    - [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest)
    - [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse)
    - [GetAdminAppsRequest](#app.user.manager.v1.GetAdminAppsRequest)
    - [GetAdminAppsResponse](#app.user.manager.v1.GetAdminAppsResponse)
    - [GetAppControlByAppRequest](#app.user.manager.v1.GetAppControlByAppRequest)
    - [GetAppControlByAppResponse](#app.user.manager.v1.GetAppControlByAppResponse)
    - [GetAppControlRequest](#app.user.manager.v1.GetAppControlRequest)
    - [GetAppControlResponse](#app.user.manager.v1.GetAppControlResponse)
    - [GetAppInfoRequest](#app.user.manager.v1.GetAppInfoRequest)
    - [GetAppInfoResponse](#app.user.manager.v1.GetAppInfoResponse)
    - [GetAppInfosByCreatorRequest](#app.user.manager.v1.GetAppInfosByCreatorRequest)
    - [GetAppInfosByCreatorResponse](#app.user.manager.v1.GetAppInfosByCreatorResponse)
    - [GetAppInfosRequest](#app.user.manager.v1.GetAppInfosRequest)
    - [GetAppInfosResponse](#app.user.manager.v1.GetAppInfosResponse)
    - [GetAppRequest](#app.user.manager.v1.GetAppRequest)
    - [GetAppResponse](#app.user.manager.v1.GetAppResponse)
    - [GetAppRoleByAppRoleRequest](#app.user.manager.v1.GetAppRoleByAppRoleRequest)
    - [GetAppRoleByAppRoleResponse](#app.user.manager.v1.GetAppRoleByAppRoleResponse)
    - [GetAppRoleRequest](#app.user.manager.v1.GetAppRoleRequest)
    - [GetAppRoleResponse](#app.user.manager.v1.GetAppRoleResponse)
    - [GetAppRoleUserByAppUserRequest](#app.user.manager.v1.GetAppRoleUserByAppUserRequest)
    - [GetAppRoleUserByAppUserResponse](#app.user.manager.v1.GetAppRoleUserByAppUserResponse)
    - [GetAppRoleUserRequest](#app.user.manager.v1.GetAppRoleUserRequest)
    - [GetAppRoleUserResponse](#app.user.manager.v1.GetAppRoleUserResponse)
    - [GetAppRoleUsersByAppRoleRequest](#app.user.manager.v1.GetAppRoleUsersByAppRoleRequest)
    - [GetAppRoleUsersByAppRoleResponse](#app.user.manager.v1.GetAppRoleUsersByAppRoleResponse)
    - [GetAppRolesByAppRequest](#app.user.manager.v1.GetAppRolesByAppRequest)
    - [GetAppRolesByAppResponse](#app.user.manager.v1.GetAppRolesByAppResponse)
    - [GetAppUserByAppAccountRequest](#app.user.manager.v1.GetAppUserByAppAccountRequest)
    - [GetAppUserByAppAccountResponse](#app.user.manager.v1.GetAppUserByAppAccountResponse)
    - [GetAppUserByAppUserRequest](#app.user.manager.v1.GetAppUserByAppUserRequest)
    - [GetAppUserByAppUserResponse](#app.user.manager.v1.GetAppUserByAppUserResponse)
    - [GetAppUserControlByAppUserRequest](#app.user.manager.v1.GetAppUserControlByAppUserRequest)
    - [GetAppUserControlByAppUserResponse](#app.user.manager.v1.GetAppUserControlByAppUserResponse)
    - [GetAppUserControlRequest](#app.user.manager.v1.GetAppUserControlRequest)
    - [GetAppUserControlResponse](#app.user.manager.v1.GetAppUserControlResponse)
    - [GetAppUserExtraByAppUserRequest](#app.user.manager.v1.GetAppUserExtraByAppUserRequest)
    - [GetAppUserExtraByAppUserResponse](#app.user.manager.v1.GetAppUserExtraByAppUserResponse)
    - [GetAppUserExtraRequest](#app.user.manager.v1.GetAppUserExtraRequest)
    - [GetAppUserExtraResponse](#app.user.manager.v1.GetAppUserExtraResponse)
    - [GetAppUserInfoByAppUserRequest](#app.user.manager.v1.GetAppUserInfoByAppUserRequest)
    - [GetAppUserInfoByAppUserResponse](#app.user.manager.v1.GetAppUserInfoByAppUserResponse)
    - [GetAppUserInfoRequest](#app.user.manager.v1.GetAppUserInfoRequest)
    - [GetAppUserInfoResponse](#app.user.manager.v1.GetAppUserInfoResponse)
    - [GetAppUserInfosByAppRequest](#app.user.manager.v1.GetAppUserInfosByAppRequest)
    - [GetAppUserInfosByAppResponse](#app.user.manager.v1.GetAppUserInfosByAppResponse)
    - [GetAppUserRequest](#app.user.manager.v1.GetAppUserRequest)
    - [GetAppUserResponse](#app.user.manager.v1.GetAppUserResponse)
    - [GetAppUserSecretByAppUserRequest](#app.user.manager.v1.GetAppUserSecretByAppUserRequest)
    - [GetAppUserSecretByAppUserResponse](#app.user.manager.v1.GetAppUserSecretByAppUserResponse)
    - [GetAppUserSecretRequest](#app.user.manager.v1.GetAppUserSecretRequest)
    - [GetAppUserSecretResponse](#app.user.manager.v1.GetAppUserSecretResponse)
    - [GetAppUsersByAppRequest](#app.user.manager.v1.GetAppUsersByAppRequest)
    - [GetAppUsersByAppResponse](#app.user.manager.v1.GetAppUsersByAppResponse)
    - [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest)
    - [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse)
    - [GetAppsRequest](#app.user.manager.v1.GetAppsRequest)
    - [GetAppsResponse](#app.user.manager.v1.GetAppsResponse)
    - [GetBanAppByAppRequest](#app.user.manager.v1.GetBanAppByAppRequest)
    - [GetBanAppByAppResponse](#app.user.manager.v1.GetBanAppByAppResponse)
    - [GetBanAppRequest](#app.user.manager.v1.GetBanAppRequest)
    - [GetBanAppResponse](#app.user.manager.v1.GetBanAppResponse)
    - [GetBanAppUserByAppUserRequest](#app.user.manager.v1.GetBanAppUserByAppUserRequest)
    - [GetBanAppUserByAppUserResponse](#app.user.manager.v1.GetBanAppUserByAppUserResponse)
    - [GetBanAppUserRequest](#app.user.manager.v1.GetBanAppUserRequest)
    - [GetBanAppUserResponse](#app.user.manager.v1.GetBanAppUserResponse)
    - [GetGenesisRoleRequest](#app.user.manager.v1.GetGenesisRoleRequest)
    - [GetGenesisRoleResponse](#app.user.manager.v1.GetGenesisRoleResponse)
    - [GetUserRolesByAppUserRequest](#app.user.manager.v1.GetUserRolesByAppUserRequest)
    - [GetUserRolesByAppUserResponse](#app.user.manager.v1.GetUserRolesByAppUserResponse)
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
    - [UpdateAppUserSecretByAppUserRequest](#app.user.manager.v1.UpdateAppUserSecretByAppUserRequest)
    - [UpdateAppUserSecretByAppUserResponse](#app.user.manager.v1.UpdateAppUserSecretByAppUserResponse)
    - [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest)
    - [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse)
    - [UpdateBanAppRequest](#app.user.manager.v1.UpdateBanAppRequest)
    - [UpdateBanAppResponse](#app.user.manager.v1.UpdateBanAppResponse)
    - [UpdateBanAppUserRequest](#app.user.manager.v1.UpdateBanAppUserRequest)
    - [UpdateBanAppUserResponse](#app.user.manager.v1.UpdateBanAppUserResponse)
    - [VerifyAppUserByAppAccountPasswordRequest](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest)
    - [VerifyAppUserByAppAccountPasswordResponse](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse)
  
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
| CreateAt | [uint32](#uint32) |  |  |






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
| InvitationCodeMust | [bool](#bool) |  |  |






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
| CreateAt | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.AppUserControl"></a>

### AppUserControl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| SigninVerifyByGoogleAuthentication | [bool](#bool) |  |  |






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
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUserInfo"></a>

### AppUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Extra | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |
| Ctrl | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |
| Ban | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |
| Roles | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |






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
| Message | [string](#string) |  |  |






<a name="app.user.manager.v1.BanAppUser"></a>

### BanAppUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="app.user.manager.v1.CreateAdminAppsRequest"></a>

### CreateAdminAppsRequest







<a name="app.user.manager.v1.CreateAdminAppsResponse"></a>

### CreateAdminAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app.user.manager.v1.App) | repeated |  |






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






<a name="app.user.manager.v1.CreateAppUserWithSecretRequest"></a>

### CreateAppUserWithSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Secret | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateAppUserWithSecretResponse"></a>

### CreateAppUserWithSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






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






<a name="app.user.manager.v1.CreateGenesisRoleRequest"></a>

### CreateGenesisRoleRequest







<a name="app.user.manager.v1.CreateGenesisRoleResponse"></a>

### CreateGenesisRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.CreateGenesisRoleUserRequest"></a>

### CreateGenesisRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Secret | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateGenesisRoleUserResponse"></a>

### CreateGenesisRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| RoleUser | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.DeleteAppRoleUserRequest"></a>

### DeleteAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






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
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.DeleteBanAppUserResponse"></a>

### DeleteBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetAdminAppsRequest"></a>

### GetAdminAppsRequest







<a name="app.user.manager.v1.GetAdminAppsResponse"></a>

### GetAdminAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app.user.manager.v1.App) | repeated |  |






<a name="app.user.manager.v1.GetAppControlByAppRequest"></a>

### GetAppControlByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppControlByAppResponse"></a>

### GetAppControlByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.GetAppControlRequest"></a>

### GetAppControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppControlResponse"></a>

### GetAppControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.GetAppInfoRequest"></a>

### GetAppInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppInfoResponse"></a>

### GetAppInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInfo](#app.user.manager.v1.AppInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosByCreatorRequest"></a>

### GetAppInfosByCreatorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosByCreatorResponse"></a>

### GetAppInfosByCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppInfosRequest"></a>

### GetAppInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosResponse"></a>

### GetAppInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppRequest"></a>

### GetAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppResponse"></a>

### GetAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.GetAppRoleByAppRoleRequest"></a>

### GetAppRoleByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Role | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleByAppRoleResponse"></a>

### GetAppRoleByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetAppRoleRequest"></a>

### GetAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleResponse"></a>

### GetAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserByAppUserRequest"></a>

### GetAppRoleUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserByAppUserResponse"></a>

### GetAppRoleUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRoleUser](#app.user.manager.v1.AppRoleUser) | repeated |  |






<a name="app.user.manager.v1.GetAppRoleUserRequest"></a>

### GetAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserResponse"></a>

### GetAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.GetAppRoleUsersByAppRoleRequest"></a>

### GetAppRoleUsersByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppRoleUsersByAppRoleResponse"></a>

### GetAppRoleUsersByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRoleUser](#app.user.manager.v1.AppRoleUser) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppRolesByAppRequest"></a>

### GetAppRolesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppRolesByAppResponse"></a>

### GetAppRolesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppAccountRequest"></a>

### GetAppUserByAppAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppAccountResponse"></a>

### GetAppUserByAppAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppUserRequest"></a>

### GetAppUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppUserResponse"></a>

### GetAppUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserControlByAppUserRequest"></a>

### GetAppUserControlByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserControlByAppUserResponse"></a>

### GetAppUserControlByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.GetAppUserControlRequest"></a>

### GetAppUserControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserControlResponse"></a>

### GetAppUserControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraByAppUserRequest"></a>

### GetAppUserExtraByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraByAppUserResponse"></a>

### GetAppUserExtraByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraRequest"></a>

### GetAppUserExtraRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraResponse"></a>

### GetAppUserExtraResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoByAppUserRequest"></a>

### GetAppUserInfoByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoByAppUserResponse"></a>

### GetAppUserInfoByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoRequest"></a>

### GetAppUserInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoResponse"></a>

### GetAppUserInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfosByAppRequest"></a>

### GetAppUserInfosByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfosByAppResponse"></a>

### GetAppUserInfosByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppUserInfo](#app.user.manager.v1.AppUserInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppUserRequest"></a>

### GetAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserResponse"></a>

### GetAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretByAppUserRequest"></a>

### GetAppUserSecretByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretByAppUserResponse"></a>

### GetAppUserSecretByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretRequest"></a>

### GetAppUserSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretResponse"></a>

### GetAppUserSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.GetAppUsersByAppRequest"></a>

### GetAppUsersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppUsersByAppResponse"></a>

### GetAppUsersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppUser](#app.user.manager.v1.AppUser) | repeated |  |
| Total | [uint32](#uint32) |  |  |






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
| Infos | [App](#app.user.manager.v1.App) | repeated |  |
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
| Infos | [App](#app.user.manager.v1.App) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetBanAppByAppRequest"></a>

### GetBanAppByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppByAppResponse"></a>

### GetBanAppByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.GetBanAppRequest"></a>

### GetBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppResponse"></a>

### GetBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.GetBanAppUserByAppUserRequest"></a>

### GetBanAppUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppUserByAppUserResponse"></a>

### GetBanAppUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetBanAppUserRequest"></a>

### GetBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppUserResponse"></a>

### GetBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetGenesisRoleRequest"></a>

### GetGenesisRoleRequest







<a name="app.user.manager.v1.GetGenesisRoleResponse"></a>

### GetGenesisRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetUserRolesByAppUserRequest"></a>

### GetUserRolesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetUserRolesByAppUserResponse"></a>

### GetUserRolesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |
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






<a name="app.user.manager.v1.UpdateAppUserSecretByAppUserRequest"></a>

### UpdateAppUserSecretByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |






<a name="app.user.manager.v1.UpdateAppUserSecretByAppUserResponse"></a>

### UpdateAppUserSecretByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






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






<a name="app.user.manager.v1.UpdateBanAppRequest"></a>

### UpdateBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.UpdateBanAppResponse"></a>

### UpdateBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.UpdateBanAppUserRequest"></a>

### UpdateBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.UpdateBanAppUserResponse"></a>

### UpdateBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest"></a>

### VerifyAppUserByAppAccountPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |






<a name="app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse"></a>

### VerifyAppUserByAppAccountPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |





 

 

 


<a name="app.user.manager.v1.AppUserManager"></a>

### AppUserManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| CreateAdminApps | [CreateAdminAppsRequest](#app.user.manager.v1.CreateAdminAppsRequest) | [CreateAdminAppsResponse](#app.user.manager.v1.CreateAdminAppsResponse) |  |
| GetAdminApps | [GetAdminAppsRequest](#app.user.manager.v1.GetAdminAppsRequest) | [GetAdminAppsResponse](#app.user.manager.v1.GetAdminAppsResponse) |  |
| CreateApp | [CreateAppRequest](#app.user.manager.v1.CreateAppRequest) | [CreateAppResponse](#app.user.manager.v1.CreateAppResponse) |  |
| GetApp | [GetAppRequest](#app.user.manager.v1.GetAppRequest) | [GetAppResponse](#app.user.manager.v1.GetAppResponse) |  |
| GetApps | [GetAppsRequest](#app.user.manager.v1.GetAppsRequest) | [GetAppsResponse](#app.user.manager.v1.GetAppsResponse) |  |
| GetAppsByCreator | [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest) | [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse) |  |
| UpdateApp | [UpdateAppRequest](#app.user.manager.v1.UpdateAppRequest) | [UpdateAppResponse](#app.user.manager.v1.UpdateAppResponse) |  |
| CreateAppControl | [CreateAppControlRequest](#app.user.manager.v1.CreateAppControlRequest) | [CreateAppControlResponse](#app.user.manager.v1.CreateAppControlResponse) |  |
| GetAppControl | [GetAppControlRequest](#app.user.manager.v1.GetAppControlRequest) | [GetAppControlResponse](#app.user.manager.v1.GetAppControlResponse) |  |
| GetAppControlByApp | [GetAppControlByAppRequest](#app.user.manager.v1.GetAppControlByAppRequest) | [GetAppControlByAppResponse](#app.user.manager.v1.GetAppControlByAppResponse) |  |
| UpdateAppControl | [UpdateAppControlRequest](#app.user.manager.v1.UpdateAppControlRequest) | [UpdateAppControlResponse](#app.user.manager.v1.UpdateAppControlResponse) |  |
| CreateBanApp | [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest) | [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse) |  |
| GetBanApp | [GetBanAppRequest](#app.user.manager.v1.GetBanAppRequest) | [GetBanAppResponse](#app.user.manager.v1.GetBanAppResponse) |  |
| GetBanAppByApp | [GetBanAppByAppRequest](#app.user.manager.v1.GetBanAppByAppRequest) | [GetBanAppByAppResponse](#app.user.manager.v1.GetBanAppByAppResponse) |  |
| UpdateBanApp | [UpdateBanAppRequest](#app.user.manager.v1.UpdateBanAppRequest) | [UpdateBanAppResponse](#app.user.manager.v1.UpdateBanAppResponse) |  |
| DeleteBanApp | [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest) | [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse) |  |
| GetAppInfo | [GetAppInfoRequest](#app.user.manager.v1.GetAppInfoRequest) | [GetAppInfoResponse](#app.user.manager.v1.GetAppInfoResponse) |  |
| GetAppInfos | [GetAppInfosRequest](#app.user.manager.v1.GetAppInfosRequest) | [GetAppInfosResponse](#app.user.manager.v1.GetAppInfosResponse) |  |
| GetAppInfosByCreator | [GetAppInfosByCreatorRequest](#app.user.manager.v1.GetAppInfosByCreatorRequest) | [GetAppInfosByCreatorResponse](#app.user.manager.v1.GetAppInfosByCreatorResponse) |  |
| CreateAppUser | [CreateAppUserRequest](#app.user.manager.v1.CreateAppUserRequest) | [CreateAppUserResponse](#app.user.manager.v1.CreateAppUserResponse) |  |
| GetAppUser | [GetAppUserRequest](#app.user.manager.v1.GetAppUserRequest) | [GetAppUserResponse](#app.user.manager.v1.GetAppUserResponse) |  |
| GetAppUserByAppUser | [GetAppUserByAppUserRequest](#app.user.manager.v1.GetAppUserByAppUserRequest) | [GetAppUserByAppUserResponse](#app.user.manager.v1.GetAppUserByAppUserResponse) |  |
| GetAppUserByAppAccount | [GetAppUserByAppAccountRequest](#app.user.manager.v1.GetAppUserByAppAccountRequest) | [GetAppUserByAppAccountResponse](#app.user.manager.v1.GetAppUserByAppAccountResponse) |  |
| VerifyAppUserByAppAccountPassword | [VerifyAppUserByAppAccountPasswordRequest](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest) | [VerifyAppUserByAppAccountPasswordResponse](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse) |  |
| GetAppUsersByApp | [GetAppUsersByAppRequest](#app.user.manager.v1.GetAppUsersByAppRequest) | [GetAppUsersByAppResponse](#app.user.manager.v1.GetAppUsersByAppResponse) |  |
| UpdateAppUser | [UpdateAppUserRequest](#app.user.manager.v1.UpdateAppUserRequest) | [UpdateAppUserResponse](#app.user.manager.v1.UpdateAppUserResponse) |  |
| CreateAppUserSecret | [CreateAppUserSecretRequest](#app.user.manager.v1.CreateAppUserSecretRequest) | [CreateAppUserSecretResponse](#app.user.manager.v1.CreateAppUserSecretResponse) |  |
| GetAppUserSecret | [GetAppUserSecretRequest](#app.user.manager.v1.GetAppUserSecretRequest) | [GetAppUserSecretResponse](#app.user.manager.v1.GetAppUserSecretResponse) |  |
| GetAppUserSecretByAppUser | [GetAppUserSecretByAppUserRequest](#app.user.manager.v1.GetAppUserSecretByAppUserRequest) | [GetAppUserSecretByAppUserResponse](#app.user.manager.v1.GetAppUserSecretByAppUserResponse) |  |
| UpdateAppUserSecret | [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest) | [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse) |  |
| UpdateAppUserSecretByAppUser | [UpdateAppUserSecretByAppUserRequest](#app.user.manager.v1.UpdateAppUserSecretByAppUserRequest) | [UpdateAppUserSecretByAppUserResponse](#app.user.manager.v1.UpdateAppUserSecretByAppUserResponse) |  |
| CreateAppUserExtra | [CreateAppUserExtraRequest](#app.user.manager.v1.CreateAppUserExtraRequest) | [CreateAppUserExtraResponse](#app.user.manager.v1.CreateAppUserExtraResponse) |  |
| GetAppUserExtra | [GetAppUserExtraRequest](#app.user.manager.v1.GetAppUserExtraRequest) | [GetAppUserExtraResponse](#app.user.manager.v1.GetAppUserExtraResponse) |  |
| GetAppUserExtraByAppUser | [GetAppUserExtraByAppUserRequest](#app.user.manager.v1.GetAppUserExtraByAppUserRequest) | [GetAppUserExtraByAppUserResponse](#app.user.manager.v1.GetAppUserExtraByAppUserResponse) |  |
| UpdateAppUserExtra | [UpdateAppUserExtraRequest](#app.user.manager.v1.UpdateAppUserExtraRequest) | [UpdateAppUserExtraResponse](#app.user.manager.v1.UpdateAppUserExtraResponse) |  |
| CreateBanAppUser | [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest) | [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse) |  |
| GetBanAppUser | [GetBanAppUserRequest](#app.user.manager.v1.GetBanAppUserRequest) | [GetBanAppUserResponse](#app.user.manager.v1.GetBanAppUserResponse) |  |
| GetBanAppUserByAppUser | [GetBanAppUserByAppUserRequest](#app.user.manager.v1.GetBanAppUserByAppUserRequest) | [GetBanAppUserByAppUserResponse](#app.user.manager.v1.GetBanAppUserByAppUserResponse) |  |
| UpdateBanAppUser | [UpdateBanAppUserRequest](#app.user.manager.v1.UpdateBanAppUserRequest) | [UpdateBanAppUserResponse](#app.user.manager.v1.UpdateBanAppUserResponse) |  |
| DeleteBanAppUser | [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest) | [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse) |  |
| CreateAppUserControl | [CreateAppUserControlRequest](#app.user.manager.v1.CreateAppUserControlRequest) | [CreateAppUserControlResponse](#app.user.manager.v1.CreateAppUserControlResponse) |  |
| GetAppUserControl | [GetAppUserControlRequest](#app.user.manager.v1.GetAppUserControlRequest) | [GetAppUserControlResponse](#app.user.manager.v1.GetAppUserControlResponse) |  |
| GetAppUserControlByAppUser | [GetAppUserControlByAppUserRequest](#app.user.manager.v1.GetAppUserControlByAppUserRequest) | [GetAppUserControlByAppUserResponse](#app.user.manager.v1.GetAppUserControlByAppUserResponse) |  |
| UpdateAppUserControl | [UpdateAppUserControlRequest](#app.user.manager.v1.UpdateAppUserControlRequest) | [UpdateAppUserControlResponse](#app.user.manager.v1.UpdateAppUserControlResponse) |  |
| CreateGenesisRole | [CreateGenesisRoleRequest](#app.user.manager.v1.CreateGenesisRoleRequest) | [CreateGenesisRoleResponse](#app.user.manager.v1.CreateGenesisRoleResponse) |  |
| GetGenesisRole | [GetGenesisRoleRequest](#app.user.manager.v1.GetGenesisRoleRequest) | [GetGenesisRoleResponse](#app.user.manager.v1.GetGenesisRoleResponse) |  |
| CreateGenesisRoleUser | [CreateGenesisRoleUserRequest](#app.user.manager.v1.CreateGenesisRoleUserRequest) | [CreateGenesisRoleUserResponse](#app.user.manager.v1.CreateGenesisRoleUserResponse) |  |
| CreateAppRole | [CreateAppRoleRequest](#app.user.manager.v1.CreateAppRoleRequest) | [CreateAppRoleResponse](#app.user.manager.v1.CreateAppRoleResponse) |  |
| GetAppRole | [GetAppRoleRequest](#app.user.manager.v1.GetAppRoleRequest) | [GetAppRoleResponse](#app.user.manager.v1.GetAppRoleResponse) |  |
| GetAppRoleByAppRole | [GetAppRoleByAppRoleRequest](#app.user.manager.v1.GetAppRoleByAppRoleRequest) | [GetAppRoleByAppRoleResponse](#app.user.manager.v1.GetAppRoleByAppRoleResponse) |  |
| GetAppRolesByApp | [GetAppRolesByAppRequest](#app.user.manager.v1.GetAppRolesByAppRequest) | [GetAppRolesByAppResponse](#app.user.manager.v1.GetAppRolesByAppResponse) |  |
| UpdateAppRole | [UpdateAppRoleRequest](#app.user.manager.v1.UpdateAppRoleRequest) | [UpdateAppRoleResponse](#app.user.manager.v1.UpdateAppRoleResponse) |  |
| CreateAppRoleUser | [CreateAppRoleUserRequest](#app.user.manager.v1.CreateAppRoleUserRequest) | [CreateAppRoleUserResponse](#app.user.manager.v1.CreateAppRoleUserResponse) |  |
| GetAppRoleUser | [GetAppRoleUserRequest](#app.user.manager.v1.GetAppRoleUserRequest) | [GetAppRoleUserResponse](#app.user.manager.v1.GetAppRoleUserResponse) |  |
| GetAppRoleUserByAppUser | [GetAppRoleUserByAppUserRequest](#app.user.manager.v1.GetAppRoleUserByAppUserRequest) | [GetAppRoleUserByAppUserResponse](#app.user.manager.v1.GetAppRoleUserByAppUserResponse) |  |
| GetAppRoleUsersByAppRole | [GetAppRoleUsersByAppRoleRequest](#app.user.manager.v1.GetAppRoleUsersByAppRoleRequest) | [GetAppRoleUsersByAppRoleResponse](#app.user.manager.v1.GetAppRoleUsersByAppRoleResponse) |  |
| GetUserRolesByAppUser | [GetUserRolesByAppUserRequest](#app.user.manager.v1.GetUserRolesByAppUserRequest) | [GetUserRolesByAppUserResponse](#app.user.manager.v1.GetUserRolesByAppUserResponse) |  |
| DeleteAppRoleUser | [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest) | [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse) |  |
| GetAppUserInfo | [GetAppUserInfoRequest](#app.user.manager.v1.GetAppUserInfoRequest) | [GetAppUserInfoResponse](#app.user.manager.v1.GetAppUserInfoResponse) |  |
| GetAppUserInfoByAppUser | [GetAppUserInfoByAppUserRequest](#app.user.manager.v1.GetAppUserInfoByAppUserRequest) | [GetAppUserInfoByAppUserResponse](#app.user.manager.v1.GetAppUserInfoByAppUserResponse) |  |
| GetAppUserInfosByApp | [GetAppUserInfosByAppRequest](#app.user.manager.v1.GetAppUserInfosByAppRequest) | [GetAppUserInfosByAppResponse](#app.user.manager.v1.GetAppUserInfosByAppResponse) |  |
| CreateAppUserWithSecret | [CreateAppUserWithSecretRequest](#app.user.manager.v1.CreateAppUserWithSecretRequest) | [CreateAppUserWithSecretResponse](#app.user.manager.v1.CreateAppUserWithSecretResponse) |  |

 



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
| CreateAt | [uint32](#uint32) |  |  |






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
| InvitationCodeMust | [bool](#bool) |  |  |






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
| CreateAt | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.AppUserControl"></a>

### AppUserControl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| SigninVerifyByGoogleAuthentication | [bool](#bool) |  |  |






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
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |






<a name="app.user.manager.v1.AppUserInfo"></a>

### AppUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Extra | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |
| Ctrl | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |
| Ban | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |
| Roles | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |






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
| Message | [string](#string) |  |  |






<a name="app.user.manager.v1.BanAppUser"></a>

### BanAppUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="app.user.manager.v1.CreateAdminAppsRequest"></a>

### CreateAdminAppsRequest







<a name="app.user.manager.v1.CreateAdminAppsResponse"></a>

### CreateAdminAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app.user.manager.v1.App) | repeated |  |






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






<a name="app.user.manager.v1.CreateAppUserWithSecretRequest"></a>

### CreateAppUserWithSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Secret | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateAppUserWithSecretResponse"></a>

### CreateAppUserWithSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






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






<a name="app.user.manager.v1.CreateGenesisRoleRequest"></a>

### CreateGenesisRoleRequest







<a name="app.user.manager.v1.CreateGenesisRoleResponse"></a>

### CreateGenesisRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.CreateGenesisRoleUserRequest"></a>

### CreateGenesisRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| Secret | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.CreateGenesisRoleUserResponse"></a>

### CreateGenesisRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| User | [AppUser](#app.user.manager.v1.AppUser) |  |  |
| RoleUser | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.DeleteAppRoleUserRequest"></a>

### DeleteAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






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
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.DeleteBanAppUserResponse"></a>

### DeleteBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetAdminAppsRequest"></a>

### GetAdminAppsRequest







<a name="app.user.manager.v1.GetAdminAppsResponse"></a>

### GetAdminAppsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [App](#app.user.manager.v1.App) | repeated |  |






<a name="app.user.manager.v1.GetAppControlByAppRequest"></a>

### GetAppControlByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppControlByAppResponse"></a>

### GetAppControlByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.GetAppControlRequest"></a>

### GetAppControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppControlResponse"></a>

### GetAppControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppControl](#app.user.manager.v1.AppControl) |  |  |






<a name="app.user.manager.v1.GetAppInfoRequest"></a>

### GetAppInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppInfoResponse"></a>

### GetAppInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppInfo](#app.user.manager.v1.AppInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosByCreatorRequest"></a>

### GetAppInfosByCreatorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosByCreatorResponse"></a>

### GetAppInfosByCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppInfosRequest"></a>

### GetAppInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppInfosResponse"></a>

### GetAppInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppInfo](#app.user.manager.v1.AppInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppRequest"></a>

### GetAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppResponse"></a>

### GetAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [App](#app.user.manager.v1.App) |  |  |






<a name="app.user.manager.v1.GetAppRoleByAppRoleRequest"></a>

### GetAppRoleByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Role | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleByAppRoleResponse"></a>

### GetAppRoleByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetAppRoleRequest"></a>

### GetAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleResponse"></a>

### GetAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserByAppUserRequest"></a>

### GetAppRoleUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserByAppUserResponse"></a>

### GetAppRoleUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRoleUser](#app.user.manager.v1.AppRoleUser) | repeated |  |






<a name="app.user.manager.v1.GetAppRoleUserRequest"></a>

### GetAppRoleUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppRoleUserResponse"></a>

### GetAppRoleUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRoleUser](#app.user.manager.v1.AppRoleUser) |  |  |






<a name="app.user.manager.v1.GetAppRoleUsersByAppRoleRequest"></a>

### GetAppRoleUsersByAppRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppRoleUsersByAppRoleResponse"></a>

### GetAppRoleUsersByAppRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRoleUser](#app.user.manager.v1.AppRoleUser) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppRolesByAppRequest"></a>

### GetAppRolesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppRolesByAppResponse"></a>

### GetAppRolesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppAccountRequest"></a>

### GetAppUserByAppAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppAccountResponse"></a>

### GetAppUserByAppAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppUserRequest"></a>

### GetAppUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserByAppUserResponse"></a>

### GetAppUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserControlByAppUserRequest"></a>

### GetAppUserControlByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserControlByAppUserResponse"></a>

### GetAppUserControlByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.GetAppUserControlRequest"></a>

### GetAppUserControlRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserControlResponse"></a>

### GetAppUserControlResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserControl](#app.user.manager.v1.AppUserControl) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraByAppUserRequest"></a>

### GetAppUserExtraByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraByAppUserResponse"></a>

### GetAppUserExtraByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraRequest"></a>

### GetAppUserExtraRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserExtraResponse"></a>

### GetAppUserExtraResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserExtra](#app.user.manager.v1.AppUserExtra) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoByAppUserRequest"></a>

### GetAppUserInfoByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoByAppUserResponse"></a>

### GetAppUserInfoByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoRequest"></a>

### GetAppUserInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserInfoResponse"></a>

### GetAppUserInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfosByAppRequest"></a>

### GetAppUserInfosByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppUserInfosByAppResponse"></a>

### GetAppUserInfosByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppUserInfo](#app.user.manager.v1.AppUserInfo) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetAppUserRequest"></a>

### GetAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserResponse"></a>

### GetAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUser](#app.user.manager.v1.AppUser) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretByAppUserRequest"></a>

### GetAppUserSecretByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretByAppUserResponse"></a>

### GetAppUserSecretByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretRequest"></a>

### GetAppUserSecretRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetAppUserSecretResponse"></a>

### GetAppUserSecretResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






<a name="app.user.manager.v1.GetAppUsersByAppRequest"></a>

### GetAppUsersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetAppUsersByAppResponse"></a>

### GetAppUsersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppUser](#app.user.manager.v1.AppUser) | repeated |  |
| Total | [uint32](#uint32) |  |  |






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
| Infos | [App](#app.user.manager.v1.App) | repeated |  |
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
| Infos | [App](#app.user.manager.v1.App) | repeated |  |
| Total | [uint32](#uint32) |  |  |






<a name="app.user.manager.v1.GetBanAppByAppRequest"></a>

### GetBanAppByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppByAppResponse"></a>

### GetBanAppByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.GetBanAppRequest"></a>

### GetBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppResponse"></a>

### GetBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.GetBanAppUserByAppUserRequest"></a>

### GetBanAppUserByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppUserByAppUserResponse"></a>

### GetBanAppUserByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetBanAppUserRequest"></a>

### GetBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="app.user.manager.v1.GetBanAppUserResponse"></a>

### GetBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.GetGenesisRoleRequest"></a>

### GetGenesisRoleRequest







<a name="app.user.manager.v1.GetGenesisRoleResponse"></a>

### GetGenesisRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppRole](#app.user.manager.v1.AppRole) |  |  |






<a name="app.user.manager.v1.GetUserRolesByAppUserRequest"></a>

### GetUserRolesByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PageInfo | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="app.user.manager.v1.GetUserRolesByAppUserResponse"></a>

### GetUserRolesByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppRole](#app.user.manager.v1.AppRole) | repeated |  |
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






<a name="app.user.manager.v1.UpdateAppUserSecretByAppUserRequest"></a>

### UpdateAppUserSecretByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |






<a name="app.user.manager.v1.UpdateAppUserSecretByAppUserResponse"></a>

### UpdateAppUserSecretByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserSecret](#app.user.manager.v1.AppUserSecret) |  |  |






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






<a name="app.user.manager.v1.UpdateBanAppRequest"></a>

### UpdateBanAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.UpdateBanAppResponse"></a>

### UpdateBanAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanApp](#app.user.manager.v1.BanApp) |  |  |






<a name="app.user.manager.v1.UpdateBanAppUserRequest"></a>

### UpdateBanAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.UpdateBanAppUserResponse"></a>

### UpdateBanAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BanAppUser](#app.user.manager.v1.BanAppUser) |  |  |






<a name="app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest"></a>

### VerifyAppUserByAppAccountPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| PasswordHash | [string](#string) |  |  |






<a name="app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse"></a>

### VerifyAppUserByAppAccountPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppUserInfo](#app.user.manager.v1.AppUserInfo) |  |  |





 

 

 


<a name="app.user.manager.v1.AppUserManager"></a>

### AppUserManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| CreateAdminApps | [CreateAdminAppsRequest](#app.user.manager.v1.CreateAdminAppsRequest) | [CreateAdminAppsResponse](#app.user.manager.v1.CreateAdminAppsResponse) |  |
| GetAdminApps | [GetAdminAppsRequest](#app.user.manager.v1.GetAdminAppsRequest) | [GetAdminAppsResponse](#app.user.manager.v1.GetAdminAppsResponse) |  |
| CreateApp | [CreateAppRequest](#app.user.manager.v1.CreateAppRequest) | [CreateAppResponse](#app.user.manager.v1.CreateAppResponse) |  |
| GetApp | [GetAppRequest](#app.user.manager.v1.GetAppRequest) | [GetAppResponse](#app.user.manager.v1.GetAppResponse) |  |
| GetApps | [GetAppsRequest](#app.user.manager.v1.GetAppsRequest) | [GetAppsResponse](#app.user.manager.v1.GetAppsResponse) |  |
| GetAppsByCreator | [GetAppsByCreatorRequest](#app.user.manager.v1.GetAppsByCreatorRequest) | [GetAppsByCreatorResponse](#app.user.manager.v1.GetAppsByCreatorResponse) |  |
| UpdateApp | [UpdateAppRequest](#app.user.manager.v1.UpdateAppRequest) | [UpdateAppResponse](#app.user.manager.v1.UpdateAppResponse) |  |
| CreateAppControl | [CreateAppControlRequest](#app.user.manager.v1.CreateAppControlRequest) | [CreateAppControlResponse](#app.user.manager.v1.CreateAppControlResponse) |  |
| GetAppControl | [GetAppControlRequest](#app.user.manager.v1.GetAppControlRequest) | [GetAppControlResponse](#app.user.manager.v1.GetAppControlResponse) |  |
| GetAppControlByApp | [GetAppControlByAppRequest](#app.user.manager.v1.GetAppControlByAppRequest) | [GetAppControlByAppResponse](#app.user.manager.v1.GetAppControlByAppResponse) |  |
| UpdateAppControl | [UpdateAppControlRequest](#app.user.manager.v1.UpdateAppControlRequest) | [UpdateAppControlResponse](#app.user.manager.v1.UpdateAppControlResponse) |  |
| CreateBanApp | [CreateBanAppRequest](#app.user.manager.v1.CreateBanAppRequest) | [CreateBanAppResponse](#app.user.manager.v1.CreateBanAppResponse) |  |
| GetBanApp | [GetBanAppRequest](#app.user.manager.v1.GetBanAppRequest) | [GetBanAppResponse](#app.user.manager.v1.GetBanAppResponse) |  |
| GetBanAppByApp | [GetBanAppByAppRequest](#app.user.manager.v1.GetBanAppByAppRequest) | [GetBanAppByAppResponse](#app.user.manager.v1.GetBanAppByAppResponse) |  |
| UpdateBanApp | [UpdateBanAppRequest](#app.user.manager.v1.UpdateBanAppRequest) | [UpdateBanAppResponse](#app.user.manager.v1.UpdateBanAppResponse) |  |
| DeleteBanApp | [DeleteBanAppRequest](#app.user.manager.v1.DeleteBanAppRequest) | [DeleteBanAppResponse](#app.user.manager.v1.DeleteBanAppResponse) |  |
| GetAppInfo | [GetAppInfoRequest](#app.user.manager.v1.GetAppInfoRequest) | [GetAppInfoResponse](#app.user.manager.v1.GetAppInfoResponse) |  |
| GetAppInfos | [GetAppInfosRequest](#app.user.manager.v1.GetAppInfosRequest) | [GetAppInfosResponse](#app.user.manager.v1.GetAppInfosResponse) |  |
| GetAppInfosByCreator | [GetAppInfosByCreatorRequest](#app.user.manager.v1.GetAppInfosByCreatorRequest) | [GetAppInfosByCreatorResponse](#app.user.manager.v1.GetAppInfosByCreatorResponse) |  |
| CreateAppUser | [CreateAppUserRequest](#app.user.manager.v1.CreateAppUserRequest) | [CreateAppUserResponse](#app.user.manager.v1.CreateAppUserResponse) |  |
| GetAppUser | [GetAppUserRequest](#app.user.manager.v1.GetAppUserRequest) | [GetAppUserResponse](#app.user.manager.v1.GetAppUserResponse) |  |
| GetAppUserByAppUser | [GetAppUserByAppUserRequest](#app.user.manager.v1.GetAppUserByAppUserRequest) | [GetAppUserByAppUserResponse](#app.user.manager.v1.GetAppUserByAppUserResponse) |  |
| GetAppUserByAppAccount | [GetAppUserByAppAccountRequest](#app.user.manager.v1.GetAppUserByAppAccountRequest) | [GetAppUserByAppAccountResponse](#app.user.manager.v1.GetAppUserByAppAccountResponse) |  |
| VerifyAppUserByAppAccountPassword | [VerifyAppUserByAppAccountPasswordRequest](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordRequest) | [VerifyAppUserByAppAccountPasswordResponse](#app.user.manager.v1.VerifyAppUserByAppAccountPasswordResponse) |  |
| GetAppUsersByApp | [GetAppUsersByAppRequest](#app.user.manager.v1.GetAppUsersByAppRequest) | [GetAppUsersByAppResponse](#app.user.manager.v1.GetAppUsersByAppResponse) |  |
| UpdateAppUser | [UpdateAppUserRequest](#app.user.manager.v1.UpdateAppUserRequest) | [UpdateAppUserResponse](#app.user.manager.v1.UpdateAppUserResponse) |  |
| CreateAppUserSecret | [CreateAppUserSecretRequest](#app.user.manager.v1.CreateAppUserSecretRequest) | [CreateAppUserSecretResponse](#app.user.manager.v1.CreateAppUserSecretResponse) |  |
| GetAppUserSecret | [GetAppUserSecretRequest](#app.user.manager.v1.GetAppUserSecretRequest) | [GetAppUserSecretResponse](#app.user.manager.v1.GetAppUserSecretResponse) |  |
| GetAppUserSecretByAppUser | [GetAppUserSecretByAppUserRequest](#app.user.manager.v1.GetAppUserSecretByAppUserRequest) | [GetAppUserSecretByAppUserResponse](#app.user.manager.v1.GetAppUserSecretByAppUserResponse) |  |
| UpdateAppUserSecret | [UpdateAppUserSecretRequest](#app.user.manager.v1.UpdateAppUserSecretRequest) | [UpdateAppUserSecretResponse](#app.user.manager.v1.UpdateAppUserSecretResponse) |  |
| UpdateAppUserSecretByAppUser | [UpdateAppUserSecretByAppUserRequest](#app.user.manager.v1.UpdateAppUserSecretByAppUserRequest) | [UpdateAppUserSecretByAppUserResponse](#app.user.manager.v1.UpdateAppUserSecretByAppUserResponse) |  |
| CreateAppUserExtra | [CreateAppUserExtraRequest](#app.user.manager.v1.CreateAppUserExtraRequest) | [CreateAppUserExtraResponse](#app.user.manager.v1.CreateAppUserExtraResponse) |  |
| GetAppUserExtra | [GetAppUserExtraRequest](#app.user.manager.v1.GetAppUserExtraRequest) | [GetAppUserExtraResponse](#app.user.manager.v1.GetAppUserExtraResponse) |  |
| GetAppUserExtraByAppUser | [GetAppUserExtraByAppUserRequest](#app.user.manager.v1.GetAppUserExtraByAppUserRequest) | [GetAppUserExtraByAppUserResponse](#app.user.manager.v1.GetAppUserExtraByAppUserResponse) |  |
| UpdateAppUserExtra | [UpdateAppUserExtraRequest](#app.user.manager.v1.UpdateAppUserExtraRequest) | [UpdateAppUserExtraResponse](#app.user.manager.v1.UpdateAppUserExtraResponse) |  |
| CreateBanAppUser | [CreateBanAppUserRequest](#app.user.manager.v1.CreateBanAppUserRequest) | [CreateBanAppUserResponse](#app.user.manager.v1.CreateBanAppUserResponse) |  |
| GetBanAppUser | [GetBanAppUserRequest](#app.user.manager.v1.GetBanAppUserRequest) | [GetBanAppUserResponse](#app.user.manager.v1.GetBanAppUserResponse) |  |
| GetBanAppUserByAppUser | [GetBanAppUserByAppUserRequest](#app.user.manager.v1.GetBanAppUserByAppUserRequest) | [GetBanAppUserByAppUserResponse](#app.user.manager.v1.GetBanAppUserByAppUserResponse) |  |
| UpdateBanAppUser | [UpdateBanAppUserRequest](#app.user.manager.v1.UpdateBanAppUserRequest) | [UpdateBanAppUserResponse](#app.user.manager.v1.UpdateBanAppUserResponse) |  |
| DeleteBanAppUser | [DeleteBanAppUserRequest](#app.user.manager.v1.DeleteBanAppUserRequest) | [DeleteBanAppUserResponse](#app.user.manager.v1.DeleteBanAppUserResponse) |  |
| CreateAppUserControl | [CreateAppUserControlRequest](#app.user.manager.v1.CreateAppUserControlRequest) | [CreateAppUserControlResponse](#app.user.manager.v1.CreateAppUserControlResponse) |  |
| GetAppUserControl | [GetAppUserControlRequest](#app.user.manager.v1.GetAppUserControlRequest) | [GetAppUserControlResponse](#app.user.manager.v1.GetAppUserControlResponse) |  |
| GetAppUserControlByAppUser | [GetAppUserControlByAppUserRequest](#app.user.manager.v1.GetAppUserControlByAppUserRequest) | [GetAppUserControlByAppUserResponse](#app.user.manager.v1.GetAppUserControlByAppUserResponse) |  |
| UpdateAppUserControl | [UpdateAppUserControlRequest](#app.user.manager.v1.UpdateAppUserControlRequest) | [UpdateAppUserControlResponse](#app.user.manager.v1.UpdateAppUserControlResponse) |  |
| CreateGenesisRole | [CreateGenesisRoleRequest](#app.user.manager.v1.CreateGenesisRoleRequest) | [CreateGenesisRoleResponse](#app.user.manager.v1.CreateGenesisRoleResponse) |  |
| GetGenesisRole | [GetGenesisRoleRequest](#app.user.manager.v1.GetGenesisRoleRequest) | [GetGenesisRoleResponse](#app.user.manager.v1.GetGenesisRoleResponse) |  |
| CreateGenesisRoleUser | [CreateGenesisRoleUserRequest](#app.user.manager.v1.CreateGenesisRoleUserRequest) | [CreateGenesisRoleUserResponse](#app.user.manager.v1.CreateGenesisRoleUserResponse) |  |
| CreateAppRole | [CreateAppRoleRequest](#app.user.manager.v1.CreateAppRoleRequest) | [CreateAppRoleResponse](#app.user.manager.v1.CreateAppRoleResponse) |  |
| GetAppRole | [GetAppRoleRequest](#app.user.manager.v1.GetAppRoleRequest) | [GetAppRoleResponse](#app.user.manager.v1.GetAppRoleResponse) |  |
| GetAppRoleByAppRole | [GetAppRoleByAppRoleRequest](#app.user.manager.v1.GetAppRoleByAppRoleRequest) | [GetAppRoleByAppRoleResponse](#app.user.manager.v1.GetAppRoleByAppRoleResponse) |  |
| GetAppRolesByApp | [GetAppRolesByAppRequest](#app.user.manager.v1.GetAppRolesByAppRequest) | [GetAppRolesByAppResponse](#app.user.manager.v1.GetAppRolesByAppResponse) |  |
| UpdateAppRole | [UpdateAppRoleRequest](#app.user.manager.v1.UpdateAppRoleRequest) | [UpdateAppRoleResponse](#app.user.manager.v1.UpdateAppRoleResponse) |  |
| CreateAppRoleUser | [CreateAppRoleUserRequest](#app.user.manager.v1.CreateAppRoleUserRequest) | [CreateAppRoleUserResponse](#app.user.manager.v1.CreateAppRoleUserResponse) |  |
| GetAppRoleUser | [GetAppRoleUserRequest](#app.user.manager.v1.GetAppRoleUserRequest) | [GetAppRoleUserResponse](#app.user.manager.v1.GetAppRoleUserResponse) |  |
| GetAppRoleUserByAppUser | [GetAppRoleUserByAppUserRequest](#app.user.manager.v1.GetAppRoleUserByAppUserRequest) | [GetAppRoleUserByAppUserResponse](#app.user.manager.v1.GetAppRoleUserByAppUserResponse) |  |
| GetAppRoleUsersByAppRole | [GetAppRoleUsersByAppRoleRequest](#app.user.manager.v1.GetAppRoleUsersByAppRoleRequest) | [GetAppRoleUsersByAppRoleResponse](#app.user.manager.v1.GetAppRoleUsersByAppRoleResponse) |  |
| GetUserRolesByAppUser | [GetUserRolesByAppUserRequest](#app.user.manager.v1.GetUserRolesByAppUserRequest) | [GetUserRolesByAppUserResponse](#app.user.manager.v1.GetUserRolesByAppUserResponse) |  |
| DeleteAppRoleUser | [DeleteAppRoleUserRequest](#app.user.manager.v1.DeleteAppRoleUserRequest) | [DeleteAppRoleUserResponse](#app.user.manager.v1.DeleteAppRoleUserResponse) |  |
| GetAppUserInfo | [GetAppUserInfoRequest](#app.user.manager.v1.GetAppUserInfoRequest) | [GetAppUserInfoResponse](#app.user.manager.v1.GetAppUserInfoResponse) |  |
| GetAppUserInfoByAppUser | [GetAppUserInfoByAppUserRequest](#app.user.manager.v1.GetAppUserInfoByAppUserRequest) | [GetAppUserInfoByAppUserResponse](#app.user.manager.v1.GetAppUserInfoByAppUserResponse) |  |
| GetAppUserInfosByApp | [GetAppUserInfosByAppRequest](#app.user.manager.v1.GetAppUserInfosByAppRequest) | [GetAppUserInfosByAppResponse](#app.user.manager.v1.GetAppUserInfosByAppResponse) |  |
| CreateAppUserWithSecret | [CreateAppUserWithSecretRequest](#app.user.manager.v1.CreateAppUserWithSecretRequest) | [CreateAppUserWithSecretResponse](#app.user.manager.v1.CreateAppUserWithSecretResponse) |  |

 



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

