# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/third-login-gateway/third-login-gateway.proto](#npool_third-login-gateway_third-login-gateway-proto)
    - [Auth](#third-gateway-v1-Auth)
    - [AuthLoginRequest](#third-gateway-v1-AuthLoginRequest)
    - [AuthLoginResponse](#third-gateway-v1-AuthLoginResponse)
    - [GetPlatformsByAppRequest](#third-gateway-v1-GetPlatformsByAppRequest)
    - [GetPlatformsByAppResponse](#third-gateway-v1-GetPlatformsByAppResponse)
    - [Platform](#third-gateway-v1-Platform)
  
    - [ThirdLoginGateway](#third-gateway-v1-ThirdLoginGateway)
  
- [npool/third-login-gateway/third-login-gateway.proto](#npool_third-login-gateway_third-login-gateway-proto)
    - [Auth](#third-gateway-v1-Auth)
    - [AuthLoginRequest](#third-gateway-v1-AuthLoginRequest)
    - [AuthLoginResponse](#third-gateway-v1-AuthLoginResponse)
    - [GetPlatformsByAppRequest](#third-gateway-v1-GetPlatformsByAppRequest)
    - [GetPlatformsByAppResponse](#third-gateway-v1-GetPlatformsByAppResponse)
    - [Platform](#third-gateway-v1-Platform)
  
    - [ThirdLoginGateway](#third-gateway-v1-ThirdLoginGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_third-login-gateway_third-login-gateway-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/third-login-gateway/third-login-gateway.proto



<a name="third-gateway-v1-Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AuthUrl | [string](#string) |  |  |
| LogoUrl | [string](#string) |  |  |
| Platform | [string](#string) |  |  |






<a name="third-gateway-v1-AuthLoginRequest"></a>

### AuthLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Platform | [string](#string) |  |  |






<a name="third-gateway-v1-AuthLoginResponse"></a>

### AuthLoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Token | [string](#string) |  |  |






<a name="third-gateway-v1-GetPlatformsByAppRequest"></a>

### GetPlatformsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third-gateway-v1-GetPlatformsByAppResponse"></a>

### GetPlatformsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-gateway-v1-Auth) | repeated |  |






<a name="third-gateway-v1-Platform"></a>

### Platform



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Platform | [string](#string) |  |  |
| PlatformAppKey | [string](#string) |  |  |
| PlatformAppSecret | [string](#string) |  |  |
| CreateAt | [string](#string) |  |  |
| UpdateAt | [string](#string) |  |  |
| LogoUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |





 

 

 


<a name="third-gateway-v1-ThirdLoginGateway"></a>

### ThirdLoginGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| GetPlatformsByApp | [GetPlatformsByAppRequest](#third-gateway-v1-GetPlatformsByAppRequest) | [GetPlatformsByAppResponse](#third-gateway-v1-GetPlatformsByAppResponse) |  |
| AuthLogin | [AuthLoginRequest](#third-gateway-v1-AuthLoginRequest) | [AuthLoginResponse](#third-gateway-v1-AuthLoginResponse) |  |

 



<a name="npool_third-login-gateway_third-login-gateway-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/third-login-gateway/third-login-gateway.proto



<a name="third-gateway-v1-Auth"></a>

### Auth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AuthUrl | [string](#string) |  |  |
| LogoUrl | [string](#string) |  |  |
| Platform | [string](#string) |  |  |






<a name="third-gateway-v1-AuthLoginRequest"></a>

### AuthLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Platform | [string](#string) |  |  |






<a name="third-gateway-v1-AuthLoginResponse"></a>

### AuthLoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [app.user.manager.v1.AppUserInfo](#app-user-manager-v1-AppUserInfo) |  |  |
| Token | [string](#string) |  |  |






<a name="third-gateway-v1-GetPlatformsByAppRequest"></a>

### GetPlatformsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third-gateway-v1-GetPlatformsByAppResponse"></a>

### GetPlatformsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Auth](#third-gateway-v1-Auth) | repeated |  |






<a name="third-gateway-v1-Platform"></a>

### Platform



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Platform | [string](#string) |  |  |
| PlatformAppKey | [string](#string) |  |  |
| PlatformAppSecret | [string](#string) |  |  |
| CreateAt | [string](#string) |  |  |
| UpdateAt | [string](#string) |  |  |
| LogoUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |





 

 

 


<a name="third-gateway-v1-ThirdLoginGateway"></a>

### ThirdLoginGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| GetPlatformsByApp | [GetPlatformsByAppRequest](#third-gateway-v1-GetPlatformsByAppRequest) | [GetPlatformsByAppResponse](#third-gateway-v1-GetPlatformsByAppResponse) |  |
| AuthLogin | [AuthLoginRequest](#third-gateway-v1-AuthLoginRequest) | [AuthLoginResponse](#third-gateway-v1-AuthLoginResponse) |  |

 



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

