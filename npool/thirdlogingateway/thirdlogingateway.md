# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/thirdlogingateway/thirdlogingateway.proto](#npool_thirdlogingateway_thirdlogingateway-proto)
    - [Auth](#third-logon-gateway-v1-Auth)
    - [CreateAppAuthRequest](#third-logon-gateway-v1-CreateAppAuthRequest)
    - [CreateAppAuthResponse](#third-logon-gateway-v1-CreateAppAuthResponse)
    - [CreateAppAuthsRequest](#third-logon-gateway-v1-CreateAppAuthsRequest)
    - [CreateAppAuthsResponse](#third-logon-gateway-v1-CreateAppAuthsResponse)
    - [CreateAuthRequest](#third-logon-gateway-v1-CreateAuthRequest)
    - [CreateAuthResponse](#third-logon-gateway-v1-CreateAuthResponse)
    - [CreateAuthsRequest](#third-logon-gateway-v1-CreateAuthsRequest)
    - [CreateAuthsResponse](#third-logon-gateway-v1-CreateAuthsResponse)
    - [CreateThirdPartyRequest](#third-logon-gateway-v1-CreateThirdPartyRequest)
    - [CreateThirdPartyResponse](#third-logon-gateway-v1-CreateThirdPartyResponse)
    - [GetAppAuthsRequest](#third-logon-gateway-v1-GetAppAuthsRequest)
    - [GetAppAuthsResponse](#third-logon-gateway-v1-GetAppAuthsResponse)
    - [GetAuthsRequest](#third-logon-gateway-v1-GetAuthsRequest)
    - [GetAuthsResponse](#third-logon-gateway-v1-GetAuthsResponse)
    - [GetThirdPartiesRequest](#third-logon-gateway-v1-GetThirdPartiesRequest)
    - [GetThirdPartiesResponse](#third-logon-gateway-v1-GetThirdPartiesResponse)
    - [GetThirdPartyOnlyRequest](#third-logon-gateway-v1-GetThirdPartyOnlyRequest)
    - [GetThirdPartyOnlyRequest.CondsEntry](#third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry)
    - [GetThirdPartyOnlyResponse](#third-logon-gateway-v1-GetThirdPartyOnlyResponse)
    - [LoginRequest](#third-logon-gateway-v1-LoginRequest)
    - [LoginResponse](#third-logon-gateway-v1-LoginResponse)
    - [ThirdParty](#third-logon-gateway-v1-ThirdParty)
    - [UpdateThirdPartyRequest](#third-logon-gateway-v1-UpdateThirdPartyRequest)
    - [UpdateThirdPartyResponse](#third-logon-gateway-v1-UpdateThirdPartyResponse)
  
    - [ThirdLoginGateway](#third-logon-gateway-v1-ThirdLoginGateway)
  
- [npool/thirdlogingateway/thirdlogingateway.proto](#npool_thirdlogingateway_thirdlogingateway-proto)
    - [Auth](#third-logon-gateway-v1-Auth)
    - [CreateAppAuthRequest](#third-logon-gateway-v1-CreateAppAuthRequest)
    - [CreateAppAuthResponse](#third-logon-gateway-v1-CreateAppAuthResponse)
    - [CreateAppAuthsRequest](#third-logon-gateway-v1-CreateAppAuthsRequest)
    - [CreateAppAuthsResponse](#third-logon-gateway-v1-CreateAppAuthsResponse)
    - [CreateAuthRequest](#third-logon-gateway-v1-CreateAuthRequest)
    - [CreateAuthResponse](#third-logon-gateway-v1-CreateAuthResponse)
    - [CreateAuthsRequest](#third-logon-gateway-v1-CreateAuthsRequest)
    - [CreateAuthsResponse](#third-logon-gateway-v1-CreateAuthsResponse)
    - [CreateThirdPartyRequest](#third-logon-gateway-v1-CreateThirdPartyRequest)
    - [CreateThirdPartyResponse](#third-logon-gateway-v1-CreateThirdPartyResponse)
    - [GetAppAuthsRequest](#third-logon-gateway-v1-GetAppAuthsRequest)
    - [GetAppAuthsResponse](#third-logon-gateway-v1-GetAppAuthsResponse)
    - [GetAuthsRequest](#third-logon-gateway-v1-GetAuthsRequest)
    - [GetAuthsResponse](#third-logon-gateway-v1-GetAuthsResponse)
    - [GetThirdPartiesRequest](#third-logon-gateway-v1-GetThirdPartiesRequest)
    - [GetThirdPartiesResponse](#third-logon-gateway-v1-GetThirdPartiesResponse)
    - [GetThirdPartyOnlyRequest](#third-logon-gateway-v1-GetThirdPartyOnlyRequest)
    - [GetThirdPartyOnlyRequest.CondsEntry](#third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry)
    - [GetThirdPartyOnlyResponse](#third-logon-gateway-v1-GetThirdPartyOnlyResponse)
    - [LoginRequest](#third-logon-gateway-v1-LoginRequest)
    - [LoginResponse](#third-logon-gateway-v1-LoginResponse)
    - [ThirdParty](#third-logon-gateway-v1-ThirdParty)
    - [UpdateThirdPartyRequest](#third-logon-gateway-v1-UpdateThirdPartyRequest)
    - [UpdateThirdPartyResponse](#third-logon-gateway-v1-UpdateThirdPartyResponse)
  
    - [ThirdLoginGateway](#third-logon-gateway-v1-ThirdLoginGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_thirdlogingateway_thirdlogingateway-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdlogingateway/thirdlogingateway.proto



<a name="third-logon-gateway-v1-Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ThirdPartyID | [string](#string) |  |  |
| AppKey | [string](#string) |  |  |
| AppSecret | [string](#string) |  |  |
| RedirectURL | [string](#string) |  |  |
| AuthURL | [string](#string) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthRequest"></a>

### CreateAppAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthResponse"></a>

### CreateAppAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthsRequest"></a>

### CreateAppAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAppAuthsResponse"></a>

### CreateAppAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAuthRequest"></a>

### CreateAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAuthResponse"></a>

### CreateAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAuthsRequest"></a>

### CreateAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAuthsResponse"></a>

### CreateAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateThirdPartyRequest"></a>

### CreateThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-CreateThirdPartyResponse"></a>

### CreateThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-GetAppAuthsRequest"></a>

### GetAppAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-GetAppAuthsResponse"></a>

### GetAppAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-GetAuthsRequest"></a>

### GetAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-GetAuthsResponse"></a>

### GetAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartiesRequest"></a>

### GetThirdPartiesRequest







<a name="third-logon-gateway-v1-GetThirdPartiesResponse"></a>

### GetThirdPartiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ThirdParty](#third-logon-gateway-v1-ThirdParty) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyRequest"></a>

### GetThirdPartyOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetThirdPartyOnlyRequest.CondsEntry](#third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry"></a>

### GetThirdPartyOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyResponse"></a>

### GetThirdPartyOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ThirdPartyID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="third-logon-gateway-v1-ThirdParty"></a>

### ThirdParty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| BrandName | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Domain | [string](#string) |  |  |






<a name="third-logon-gateway-v1-UpdateThirdPartyRequest"></a>

### UpdateThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-UpdateThirdPartyResponse"></a>

### UpdateThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |





 

 

 


<a name="third-logon-gateway-v1-ThirdLoginGateway"></a>

### ThirdLoginGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| GetAuths | [GetAuthsRequest](#third-logon-gateway-v1-GetAuthsRequest) | [GetAuthsResponse](#third-logon-gateway-v1-GetAuthsResponse) |  |
| GetAppAuths | [GetAppAuthsRequest](#third-logon-gateway-v1-GetAppAuthsRequest) | [GetAppAuthsResponse](#third-logon-gateway-v1-GetAppAuthsResponse) |  |
| CreateAuth | [CreateAuthRequest](#third-logon-gateway-v1-CreateAuthRequest) | [CreateAuthResponse](#third-logon-gateway-v1-CreateAuthResponse) |  |
| CreateAuths | [CreateAuthsRequest](#third-logon-gateway-v1-CreateAuthsRequest) | [CreateAuthsResponse](#third-logon-gateway-v1-CreateAuthsResponse) |  |
| CreateAppAuth | [CreateAppAuthRequest](#third-logon-gateway-v1-CreateAppAuthRequest) | [CreateAppAuthResponse](#third-logon-gateway-v1-CreateAppAuthResponse) |  |
| CreateAppAuths | [CreateAppAuthsRequest](#third-logon-gateway-v1-CreateAppAuthsRequest) | [CreateAppAuthsResponse](#third-logon-gateway-v1-CreateAppAuthsResponse) |  |
| CreateThirdParty | [CreateThirdPartyRequest](#third-logon-gateway-v1-CreateThirdPartyRequest) | [CreateThirdPartyResponse](#third-logon-gateway-v1-CreateThirdPartyResponse) |  |
| UpdateThirdParty | [UpdateThirdPartyRequest](#third-logon-gateway-v1-UpdateThirdPartyRequest) | [UpdateThirdPartyResponse](#third-logon-gateway-v1-UpdateThirdPartyResponse) |  |
| GetThirdParties | [GetThirdPartiesRequest](#third-logon-gateway-v1-GetThirdPartiesRequest) | [GetThirdPartiesResponse](#third-logon-gateway-v1-GetThirdPartiesResponse) |  |
| GetThirdPartyOnly | [GetThirdPartyOnlyRequest](#third-logon-gateway-v1-GetThirdPartyOnlyRequest) | [GetThirdPartyOnlyResponse](#third-logon-gateway-v1-GetThirdPartyOnlyResponse) |  |
| Login | [LoginRequest](#third-logon-gateway-v1-LoginRequest) | [LoginResponse](#third-logon-gateway-v1-LoginResponse) |  |

 



<a name="npool_thirdlogingateway_thirdlogingateway-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdlogingateway/thirdlogingateway.proto



<a name="third-logon-gateway-v1-Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ThirdPartyID | [string](#string) |  |  |
| AppKey | [string](#string) |  |  |
| AppSecret | [string](#string) |  |  |
| RedirectURL | [string](#string) |  |  |
| AuthURL | [string](#string) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthRequest"></a>

### CreateAppAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthResponse"></a>

### CreateAppAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAppAuthsRequest"></a>

### CreateAppAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAppAuthsResponse"></a>

### CreateAppAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAuthRequest"></a>

### CreateAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAuthResponse"></a>

### CreateAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Auth](#third-logon-gateway-v1-Auth) |  |  |






<a name="third-logon-gateway-v1-CreateAuthsRequest"></a>

### CreateAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateAuthsResponse"></a>

### CreateAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-CreateThirdPartyRequest"></a>

### CreateThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-CreateThirdPartyResponse"></a>

### CreateThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-GetAppAuthsRequest"></a>

### GetAppAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-GetAppAuthsResponse"></a>

### GetAppAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-GetAuthsRequest"></a>

### GetAuthsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-GetAuthsResponse"></a>

### GetAuthsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-logon-gateway-v1-Auth) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartiesRequest"></a>

### GetThirdPartiesRequest







<a name="third-logon-gateway-v1-GetThirdPartiesResponse"></a>

### GetThirdPartiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ThirdParty](#third-logon-gateway-v1-ThirdParty) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyRequest"></a>

### GetThirdPartyOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetThirdPartyOnlyRequest.CondsEntry](#third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry) | repeated |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyRequest-CondsEntry"></a>

### GetThirdPartyOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="third-logon-gateway-v1-GetThirdPartyOnlyResponse"></a>

### GetThirdPartyOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ThirdPartyID | [string](#string) |  |  |






<a name="third-logon-gateway-v1-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |






<a name="third-logon-gateway-v1-ThirdParty"></a>

### ThirdParty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| BrandName | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Domain | [string](#string) |  |  |






<a name="third-logon-gateway-v1-UpdateThirdPartyRequest"></a>

### UpdateThirdPartyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |






<a name="third-logon-gateway-v1-UpdateThirdPartyResponse"></a>

### UpdateThirdPartyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ThirdParty](#third-logon-gateway-v1-ThirdParty) |  |  |





 

 

 


<a name="third-logon-gateway-v1-ThirdLoginGateway"></a>

### ThirdLoginGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| GetAuths | [GetAuthsRequest](#third-logon-gateway-v1-GetAuthsRequest) | [GetAuthsResponse](#third-logon-gateway-v1-GetAuthsResponse) |  |
| GetAppAuths | [GetAppAuthsRequest](#third-logon-gateway-v1-GetAppAuthsRequest) | [GetAppAuthsResponse](#third-logon-gateway-v1-GetAppAuthsResponse) |  |
| CreateAuth | [CreateAuthRequest](#third-logon-gateway-v1-CreateAuthRequest) | [CreateAuthResponse](#third-logon-gateway-v1-CreateAuthResponse) |  |
| CreateAuths | [CreateAuthsRequest](#third-logon-gateway-v1-CreateAuthsRequest) | [CreateAuthsResponse](#third-logon-gateway-v1-CreateAuthsResponse) |  |
| CreateAppAuth | [CreateAppAuthRequest](#third-logon-gateway-v1-CreateAppAuthRequest) | [CreateAppAuthResponse](#third-logon-gateway-v1-CreateAppAuthResponse) |  |
| CreateAppAuths | [CreateAppAuthsRequest](#third-logon-gateway-v1-CreateAppAuthsRequest) | [CreateAppAuthsResponse](#third-logon-gateway-v1-CreateAppAuthsResponse) |  |
| CreateThirdParty | [CreateThirdPartyRequest](#third-logon-gateway-v1-CreateThirdPartyRequest) | [CreateThirdPartyResponse](#third-logon-gateway-v1-CreateThirdPartyResponse) |  |
| UpdateThirdParty | [UpdateThirdPartyRequest](#third-logon-gateway-v1-UpdateThirdPartyRequest) | [UpdateThirdPartyResponse](#third-logon-gateway-v1-UpdateThirdPartyResponse) |  |
| GetThirdParties | [GetThirdPartiesRequest](#third-logon-gateway-v1-GetThirdPartiesRequest) | [GetThirdPartiesResponse](#third-logon-gateway-v1-GetThirdPartiesResponse) |  |
| GetThirdPartyOnly | [GetThirdPartyOnlyRequest](#third-logon-gateway-v1-GetThirdPartyOnlyRequest) | [GetThirdPartyOnlyResponse](#third-logon-gateway-v1-GetThirdPartyOnlyResponse) |  |
| Login | [LoginRequest](#third-logon-gateway-v1-LoginRequest) | [LoginResponse](#third-logon-gateway-v1-LoginResponse) |  |

 



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

