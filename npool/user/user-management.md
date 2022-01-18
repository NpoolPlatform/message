# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/user/user-management.proto](#npool/user/user-management.proto)
    - [AddUserRequest](#user.v1.AddUserRequest)
    - [AddUserResponse](#user.v1.AddUserResponse)
    - [BindThirdPartyRequest](#user.v1.BindThirdPartyRequest)
    - [BindThirdPartyResponse](#user.v1.BindThirdPartyResponse)
    - [BindUserEmailRequest](#user.v1.BindUserEmailRequest)
    - [BindUserEmailResponse](#user.v1.BindUserEmailResponse)
    - [BindUserPhoneRequest](#user.v1.BindUserPhoneRequest)
    - [BindUserPhoneResponse](#user.v1.BindUserPhoneResponse)
    - [CertificateKycRequest](#user.v1.CertificateKycRequest)
    - [CertificateKycResponse](#user.v1.CertificateKycResponse)
    - [ChangeUserPasswordRequest](#user.v1.ChangeUserPasswordRequest)
    - [ChangeUserPasswordResponse](#user.v1.ChangeUserPasswordResponse)
    - [DeleteUserRequest](#user.v1.DeleteUserRequest)
    - [DeleteUserResponse](#user.v1.DeleteUserResponse)
    - [ForgetPasswordRequest](#user.v1.ForgetPasswordRequest)
    - [ForgetPasswordResponse](#user.v1.ForgetPasswordResponse)
    - [FrozenUser](#user.v1.FrozenUser)
    - [FrozenUserRequest](#user.v1.FrozenUserRequest)
    - [FrozenUserResponse](#user.v1.FrozenUserResponse)
    - [GetFrozenUsersRequest](#user.v1.GetFrozenUsersRequest)
    - [GetFrozenUsersResponse](#user.v1.GetFrozenUsersResponse)
    - [GetGaQRCodeRequest](#user.v1.GetGaQRCodeRequest)
    - [GetGaQRCodeResponse](#user.v1.GetGaQRCodeResponse)
    - [GetUserDetailsRequest](#user.v1.GetUserDetailsRequest)
    - [GetUserDetailsResponse](#user.v1.GetUserDetailsResponse)
    - [GetUserProvidersRequest](#user.v1.GetUserProvidersRequest)
    - [GetUserProvidersResponse](#user.v1.GetUserProvidersResponse)
    - [GetUserRequest](#user.v1.GetUserRequest)
    - [GetUserResponse](#user.v1.GetUserResponse)
    - [GetUsersRequest](#user.v1.GetUsersRequest)
    - [GetUsersResponse](#user.v1.GetUsersResponse)
    - [QueryProviderUserInfo](#user.v1.QueryProviderUserInfo)
    - [QueryUserByUserProviderIDRequest](#user.v1.QueryUserByUserProviderIDRequest)
    - [QueryUserByUserProviderIDResponse](#user.v1.QueryUserByUserProviderIDResponse)
    - [QueryUserExistRequest](#user.v1.QueryUserExistRequest)
    - [QueryUserExistResponse](#user.v1.QueryUserExistResponse)
    - [QueryUserFrozenRequest](#user.v1.QueryUserFrozenRequest)
    - [QueryUserFrozenResponse](#user.v1.QueryUserFrozenResponse)
    - [SetPasswordRequest](#user.v1.SetPasswordRequest)
    - [SetPasswordResponse](#user.v1.SetPasswordResponse)
    - [SignupRequest](#user.v1.SignupRequest)
    - [SignupResponse](#user.v1.SignupResponse)
    - [UnbindThirdPartyRequest](#user.v1.UnbindThirdPartyRequest)
    - [UnbindThirdPartyResponse](#user.v1.UnbindThirdPartyResponse)
    - [UnfrozenUserRequest](#user.v1.UnfrozenUserRequest)
    - [UnfrozenUserResponse](#user.v1.UnfrozenUserResponse)
    - [UpdateUserEmailRequest](#user.v1.UpdateUserEmailRequest)
    - [UpdateUserEmailResponse](#user.v1.UpdateUserEmailResponse)
    - [UpdateUserInfoRequest](#user.v1.UpdateUserInfoRequest)
    - [UpdateUserInfoResponse](#user.v1.UpdateUserInfoResponse)
    - [UpdateUserPhoneRequest](#user.v1.UpdateUserPhoneRequest)
    - [UpdateUserPhoneResponse](#user.v1.UpdateUserPhoneResponse)
    - [UserBasicInfo](#user.v1.UserBasicInfo)
    - [UserDetails](#user.v1.UserDetails)
    - [UserProvider](#user.v1.UserProvider)
  
    - [User](#user.v1.User)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/user/user-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/user/user-management.proto



<a name="user.v1.AddUserRequest"></a>

### AddUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.AddUserResponse"></a>

### AddUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.BindThirdPartyRequest"></a>

### BindThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ProviderID | [string](#string) |  | third party(provIDer)&#39;s ID |
| ProviderUserID | [string](#string) |  | UserID in third party(provIDer) |
| UserProviderInfo | [string](#string) |  |  |






<a name="user.v1.BindThirdPartyResponse"></a>

### BindThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserProvider](#user.v1.UserProvider) |  |  |






<a name="user.v1.BindUserEmailRequest"></a>

### BindUserEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="user.v1.BindUserEmailResponse"></a>

### BindUserEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.BindUserPhoneRequest"></a>

### BindUserPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| PhoneNumber | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="user.v1.BindUserPhoneResponse"></a>

### BindUserPhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.CertificateKycRequest"></a>

### CertificateKycRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| FrontCardImg | [string](#string) |  |  |
| BackCardImg | [string](#string) |  |  |
| UserCardImg | [string](#string) |  |  |
| CardType | [string](#string) |  |  |
| CardID | [string](#string) |  |  |
| Region | [string](#string) |  |  |






<a name="user.v1.CertificateKycResponse"></a>

### CertificateKycResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.ChangeUserPasswordRequest"></a>

### ChangeUserPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| VerifyParam | [string](#string) |  |  |
| VerifyType | [string](#string) |  | verify types are emial, phone and google. |
| OldPassword | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="user.v1.ChangeUserPasswordResponse"></a>

### ChangeUserPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| DeleteUserIDs | [string](#string) | repeated | an array of UserID who are being deleted. |






<a name="user.v1.DeleteUserResponse"></a>

### DeleteUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.ForgetPasswordRequest"></a>

### ForgetPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| VerifyParam | [string](#string) |  |  |
| VerifyType | [string](#string) |  | verify types are emial, phone and google. |
| Password | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="user.v1.ForgetPasswordResponse"></a>

### ForgetPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.FrozenUser"></a>

### FrozenUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FrozenBy | [string](#string) |  |  |
| FrozenCause | [string](#string) |  |  |
| StartAt | [uint32](#uint32) |  |  |
| EndAt | [uint32](#uint32) |  |  |
| Status | [string](#string) |  |  |
| UnfrozenBy | [string](#string) |  |  |






<a name="user.v1.FrozenUserRequest"></a>

### FrozenUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| FrozenBy | [string](#string) |  |  |
| FrozenCause | [string](#string) |  |  |






<a name="user.v1.FrozenUserResponse"></a>

### FrozenUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FrozenUser](#user.v1.FrozenUser) |  |  |






<a name="user.v1.GetFrozenUsersRequest"></a>

### GetFrozenUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Info | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="user.v1.GetFrozenUsersResponse"></a>

### GetFrozenUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [FrozenUser](#user.v1.FrozenUser) | repeated |  |






<a name="user.v1.GetGaQRCodeRequest"></a>

### GetGaQRCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="user.v1.GetGaQRCodeResponse"></a>

### GetGaQRCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.GetUserDetailsRequest"></a>

### GetUserDetailsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="user.v1.GetUserDetailsResponse"></a>

### GetUserDetailsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserDetails](#user.v1.UserDetails) |  |  |






<a name="user.v1.GetUserProvidersRequest"></a>

### GetUserProvidersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="user.v1.GetUserProvidersResponse"></a>

### GetUserProvidersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserProvider](#user.v1.UserProvider) | repeated |  |






<a name="user.v1.GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  | UserID is who is queried. |
| UserID | [string](#string) |  |  |






<a name="user.v1.GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.GetUsersRequest"></a>

### GetUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Info | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="user.v1.GetUsersResponse"></a>

### GetUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserBasicInfo](#user.v1.UserBasicInfo) | repeated |  |






<a name="user.v1.QueryProviderUserInfo"></a>

### QueryProviderUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserProviderInfo | [UserProvider](#user.v1.UserProvider) |  |  |
| UserBasicInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.QueryUserByUserProviderIDRequest"></a>

### QueryUserByUserProviderIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| ProviderID | [string](#string) |  |  |
| ProviderUserID | [string](#string) |  |  |






<a name="user.v1.QueryUserByUserProviderIDResponse"></a>

### QueryUserByUserProviderIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [QueryProviderUserInfo](#user.v1.QueryProviderUserInfo) |  |  |






<a name="user.v1.QueryUserExistRequest"></a>

### QueryUserExistRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="user.v1.QueryUserExistResponse"></a>

### QueryUserExistResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.QueryUserFrozenRequest"></a>

### QueryUserFrozenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="user.v1.QueryUserFrozenResponse"></a>

### QueryUserFrozenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FrozenUser](#user.v1.FrozenUser) |  |  |






<a name="user.v1.SetPasswordRequest"></a>

### SetPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |






<a name="user.v1.SetPasswordResponse"></a>

### SetPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.SignupRequest"></a>

### SignupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Username | [string](#string) |  | optional |
| Password | [string](#string) |  |  |
| EmailAddress | [string](#string) |  | optional |
| PhoneNumber | [string](#string) |  | optional |
| Code | [string](#string) |  |  |






<a name="user.v1.SignupResponse"></a>

### SignupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UnbindThirdPartyRequest"></a>

### UnbindThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ProviderID | [string](#string) |  |  |






<a name="user.v1.UnbindThirdPartyResponse"></a>

### UnbindThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserProvider](#user.v1.UserProvider) |  |  |






<a name="user.v1.UnfrozenUserRequest"></a>

### UnfrozenUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| UnfrozenBy | [string](#string) |  |  |






<a name="user.v1.UnfrozenUserResponse"></a>

### UnfrozenUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FrozenUser](#user.v1.FrozenUser) |  |  |






<a name="user.v1.UpdateUserEmailRequest"></a>

### UpdateUserEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldEmail | [string](#string) |  |  |
| OldCode | [string](#string) |  |  |
| NewEmail | [string](#string) |  |  |
| NewCode | [string](#string) |  |  |






<a name="user.v1.UpdateUserEmailResponse"></a>

### UpdateUserEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.UpdateUserInfoRequest"></a>

### UpdateUserInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UpdateUserInfoResponse"></a>

### UpdateUserInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |






<a name="user.v1.UpdateUserPhoneRequest"></a>

### UpdateUserPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| OldPhone | [string](#string) |  |  |
| OldCode | [string](#string) |  |  |
| NewPhone | [string](#string) |  |  |
| NewCode | [string](#string) |  |  |






<a name="user.v1.UpdateUserPhoneResponse"></a>

### UpdateUserPhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="user.v1.UserBasicInfo"></a>

### UserBasicInfo
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Username | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| Avatar | [string](#string) |  |  |
| Age | [uint32](#uint32) |  |  |
| Gender | [string](#string) |  |  |
| Region | [string](#string) |  |  |
| Birthday | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Province | [string](#string) |  |  |
| City | [string](#string) |  |  |
| PhoneNumber | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |
| SignupMethod | [string](#string) |  |  |
| Career | [string](#string) |  |  |
| DisplayName | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| StreetAddress1 | [string](#string) |  |  |
| StreetAddress2 | [string](#string) |  |  |
| Compony | [string](#string) |  |  |
| PostalCode | [string](#string) |  |  |






<a name="user.v1.UserDetails"></a>

### UserDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserBasicInfo | [UserBasicInfo](#user.v1.UserBasicInfo) |  |  |
| UserAppInfo | [application.management.v1.ApplicationUserDetail](#application.management.v1.ApplicationUserDetail) |  |  |






<a name="user.v1.UserProvider"></a>

### UserProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ProviderID | [string](#string) |  |  |
| ProviderUserID | [string](#string) |  |  |
| UserProviderInfo | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |





 

 

 


<a name="user.v1.User"></a>

### User
a service of managing users

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| SignUp | [SignupRequest](#user.v1.SignupRequest) | [SignupResponse](#user.v1.SignupResponse) | User can choose signup with username, email, phone or only emial or only phone. |
| GetUser | [GetUserRequest](#user.v1.GetUserRequest) | [GetUserResponse](#user.v1.GetUserResponse) | Get a user&#39;s info by his(her) ID, this api can be request by user self of admin. |
| GetUserDetails | [GetUserDetailsRequest](#user.v1.GetUserDetailsRequest) | [GetUserDetailsResponse](#user.v1.GetUserDetailsResponse) | Get user details in an app. |
| GetUsers | [GetUsersRequest](#user.v1.GetUsersRequest) | [GetUsersResponse](#user.v1.GetUsersResponse) | Get all users. |
| UpdateUserInfo | [UpdateUserInfoRequest](#user.v1.UpdateUserInfoRequest) | [UpdateUserInfoResponse](#user.v1.UpdateUserInfoResponse) | Update user&#39;s basic info. |
| BindUserPhone | [BindUserPhoneRequest](#user.v1.BindUserPhoneRequest) | [BindUserPhoneResponse](#user.v1.BindUserPhoneResponse) | Bind user&#39;s phone number. |
| BindUserEmail | [BindUserEmailRequest](#user.v1.BindUserEmailRequest) | [BindUserEmailResponse](#user.v1.BindUserEmailResponse) | Bind user&#39;s email address. |
| UpdateUserEmail | [UpdateUserEmailRequest](#user.v1.UpdateUserEmailRequest) | [UpdateUserEmailResponse](#user.v1.UpdateUserEmailResponse) | update user email. |
| UpdateUserPhone | [UpdateUserPhoneRequest](#user.v1.UpdateUserPhoneRequest) | [UpdateUserPhoneResponse](#user.v1.UpdateUserPhoneResponse) | update user phone. |
| BindThirdParty | [BindThirdPartyRequest](#user.v1.BindThirdPartyRequest) | [BindThirdPartyResponse](#user.v1.BindThirdPartyResponse) | Link to a third-party oauth. save the UserID from third-party into mysql. |
| UnbindThirdParty | [UnbindThirdPartyRequest](#user.v1.UnbindThirdPartyRequest) | [UnbindThirdPartyResponse](#user.v1.UnbindThirdPartyResponse) | Unlink a third-party oauth. Delete the UserID we saved from mysql. |
| ChangeUserPassword | [ChangeUserPasswordRequest](#user.v1.ChangeUserPasswordRequest) | [ChangeUserPasswordResponse](#user.v1.ChangeUserPasswordResponse) | Change user&#39;s password. Before change users password, system need the user to do an authentication. |
| ForgetPassword | [ForgetPasswordRequest](#user.v1.ForgetPasswordRequest) | [ForgetPasswordResponse](#user.v1.ForgetPasswordResponse) | Forget password. |
| SetPassword | [SetPasswordRequest](#user.v1.SetPasswordRequest) | [SetPasswordResponse](#user.v1.SetPasswordResponse) | set password. |
| AddUser | [AddUserRequest](#user.v1.AddUserRequest) | [AddUserResponse](#user.v1.AddUserResponse) | Add user. |
| DeleteUser | [DeleteUserRequest](#user.v1.DeleteUserRequest) | [DeleteUserResponse](#user.v1.DeleteUserResponse) | Delete users. |
| FrozenUser | [FrozenUserRequest](#user.v1.FrozenUserRequest) | [FrozenUserResponse](#user.v1.FrozenUserResponse) | Frozen user. |
| UnfrozenUser | [UnfrozenUserRequest](#user.v1.UnfrozenUserRequest) | [UnfrozenUserResponse](#user.v1.UnfrozenUserResponse) | Unfrozen user. |
| QueryUserFrozen | [QueryUserFrozenRequest](#user.v1.QueryUserFrozenRequest) | [QueryUserFrozenResponse](#user.v1.QueryUserFrozenResponse) | query user is frozen or not |
| GetFrozenUsers | [GetFrozenUsersRequest](#user.v1.GetFrozenUsersRequest) | [GetFrozenUsersResponse](#user.v1.GetFrozenUsersResponse) | Get frozen user list. |
| GetUserProviders | [GetUserProvidersRequest](#user.v1.GetUserProvidersRequest) | [GetUserProvidersResponse](#user.v1.GetUserProvidersResponse) | Get user provIDers info. |
| QueryUserExist | [QueryUserExistRequest](#user.v1.QueryUserExistRequest) | [QueryUserExistResponse](#user.v1.QueryUserExistResponse) | query user exist in database. |
| QueryUserByUserProviderID | [QueryUserByUserProviderIDRequest](#user.v1.QueryUserByUserProviderIDRequest) | [QueryUserByUserProviderIDResponse](#user.v1.QueryUserByUserProviderIDResponse) | query user by provIDer ID and his ID in the provIDer |

 



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

