# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/authinggateway/authinggateway.proto](#npool/authinggateway/authinggateway.proto)
    - [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest)
    - [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse)
    - [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest)
    - [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse)
    - [AuthHistory](#authing.gateway.v1.AuthHistory)
    - [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest)
    - [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse)
  
    - [AuthingGateway](#authing.gateway.v1.AuthingGateway)
  
- [npool/authinggateway/authinggateway.proto](#npool/authinggateway/authinggateway.proto)
    - [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest)
    - [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse)
    - [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest)
    - [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse)
    - [AuthHistory](#authing.gateway.v1.AuthHistory)
    - [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest)
    - [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse)
  
    - [AuthingGateway](#authing.gateway.v1.AuthingGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/authinggateway/authinggateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/authinggateway/authinggateway.proto



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





 

 

 


<a name="authing.gateway.v1.AuthingGateway"></a>

### AuthingGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| AuthByApp | [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest) | [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse) |  |
| AuthByAppRoleUser | [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest) | [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse) |  |
| GetAuthHistories | [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest) | [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse) |  |

 



<a name="npool/authinggateway/authinggateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/authinggateway/authinggateway.proto



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





 

 

 


<a name="authing.gateway.v1.AuthingGateway"></a>

### AuthingGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) |  |
| AuthByApp | [AuthByAppRequest](#authing.gateway.v1.AuthByAppRequest) | [AuthByAppResponse](#authing.gateway.v1.AuthByAppResponse) |  |
| AuthByAppRoleUser | [AuthByAppRoleUserRequest](#authing.gateway.v1.AuthByAppRoleUserRequest) | [AuthByAppRoleUserResponse](#authing.gateway.v1.AuthByAppRoleUserResponse) |  |
| GetAuthHistories | [GetAuthHistoriesRequest](#authing.gateway.v1.GetAuthHistoriesRequest) | [GetAuthHistoriesResponse](#authing.gateway.v1.GetAuthHistoriesResponse) |  |

 



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

