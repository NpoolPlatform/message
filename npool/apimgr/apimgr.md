# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/apimgr/apimgr.proto](#npool_apimgr_apimgr-proto)
    - [GetApisRequest](#api-manager-v1-GetApisRequest)
    - [GetApisResponse](#api-manager-v1-GetApisResponse)
    - [GetServiceMethodApiRequest](#api-manager-v1-GetServiceMethodApiRequest)
    - [GetServiceMethodApiResponse](#api-manager-v1-GetServiceMethodApiResponse)
    - [Path](#api-manager-v1-Path)
    - [RegisterRequest](#api-manager-v1-RegisterRequest)
    - [RegisterResponse](#api-manager-v1-RegisterResponse)
    - [ServiceApis](#api-manager-v1-ServiceApis)
    - [ServicePath](#api-manager-v1-ServicePath)
  
    - [ApiManager](#api-manager-v1-ApiManager)
  
- [npool/apimgr/apimgr.proto](#npool_apimgr_apimgr-proto)
    - [GetApisRequest](#api-manager-v1-GetApisRequest)
    - [GetApisResponse](#api-manager-v1-GetApisResponse)
    - [GetServiceMethodApiRequest](#api-manager-v1-GetServiceMethodApiRequest)
    - [GetServiceMethodApiResponse](#api-manager-v1-GetServiceMethodApiResponse)
    - [Path](#api-manager-v1-Path)
    - [RegisterRequest](#api-manager-v1-RegisterRequest)
    - [RegisterResponse](#api-manager-v1-RegisterResponse)
    - [ServiceApis](#api-manager-v1-ServiceApis)
    - [ServicePath](#api-manager-v1-ServicePath)
  
    - [ApiManager](#api-manager-v1-ApiManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_apimgr_apimgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/apimgr/apimgr.proto



<a name="api-manager-v1-GetApisRequest"></a>

### GetApisRequest







<a name="api-manager-v1-GetApisResponse"></a>

### GetApisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ServicePath](#api-manager-v1-ServicePath) | repeated |  |






<a name="api-manager-v1-GetServiceMethodApiRequest"></a>

### GetServiceMethodApiRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ServiceName | [string](#string) |  |  |
| MethodName | [string](#string) |  |  |






<a name="api-manager-v1-GetServiceMethodApiResponse"></a>

### GetServiceMethodApiResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServicePath](#api-manager-v1-ServicePath) |  |  |






<a name="api-manager-v1-Path"></a>

### Path



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Method | [string](#string) |  |  |
| Path | [string](#string) |  |  |
| Exported | [bool](#bool) |  |  |
| MethodName | [string](#string) |  |  |






<a name="api-manager-v1-RegisterRequest"></a>

### RegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServiceApis](#api-manager-v1-ServiceApis) |  |  |






<a name="api-manager-v1-RegisterResponse"></a>

### RegisterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServiceApis](#api-manager-v1-ServiceApis) |  |  |






<a name="api-manager-v1-ServiceApis"></a>

### ServiceApis



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Protocol | [string](#string) |  |  |
| ServiceName | [string](#string) |  |  |
| PathPrefix | [string](#string) |  |  |
| Paths | [Path](#api-manager-v1-Path) | repeated |  |






<a name="api-manager-v1-ServicePath"></a>

### ServicePath



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Protocol | [string](#string) |  |  |
| ServiceName | [string](#string) |  |  |
| Domains | [string](#string) | repeated |  |
| PathPrefix | [string](#string) |  |  |
| Method | [string](#string) |  |  |
| Path | [string](#string) |  |  |
| Exported | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |





 

 

 


<a name="api-manager-v1-ApiManager"></a>

### ApiManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| Register | [RegisterRequest](#api-manager-v1-RegisterRequest) | [RegisterResponse](#api-manager-v1-RegisterResponse) |  |
| GetApis | [GetApisRequest](#api-manager-v1-GetApisRequest) | [GetApisResponse](#api-manager-v1-GetApisResponse) |  |
| GetServiceMethodApi | [GetServiceMethodApiRequest](#api-manager-v1-GetServiceMethodApiRequest) | [GetServiceMethodApiResponse](#api-manager-v1-GetServiceMethodApiResponse) |  |

 



<a name="npool_apimgr_apimgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/apimgr/apimgr.proto



<a name="api-manager-v1-GetApisRequest"></a>

### GetApisRequest







<a name="api-manager-v1-GetApisResponse"></a>

### GetApisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ServicePath](#api-manager-v1-ServicePath) | repeated |  |






<a name="api-manager-v1-GetServiceMethodApiRequest"></a>

### GetServiceMethodApiRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ServiceName | [string](#string) |  |  |
| MethodName | [string](#string) |  |  |






<a name="api-manager-v1-GetServiceMethodApiResponse"></a>

### GetServiceMethodApiResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServicePath](#api-manager-v1-ServicePath) |  |  |






<a name="api-manager-v1-Path"></a>

### Path



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Method | [string](#string) |  |  |
| Path | [string](#string) |  |  |
| Exported | [bool](#bool) |  |  |
| MethodName | [string](#string) |  |  |






<a name="api-manager-v1-RegisterRequest"></a>

### RegisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServiceApis](#api-manager-v1-ServiceApis) |  |  |






<a name="api-manager-v1-RegisterResponse"></a>

### RegisterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ServiceApis](#api-manager-v1-ServiceApis) |  |  |






<a name="api-manager-v1-ServiceApis"></a>

### ServiceApis



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Protocol | [string](#string) |  |  |
| ServiceName | [string](#string) |  |  |
| PathPrefix | [string](#string) |  |  |
| Paths | [Path](#api-manager-v1-Path) | repeated |  |






<a name="api-manager-v1-ServicePath"></a>

### ServicePath



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Protocol | [string](#string) |  |  |
| ServiceName | [string](#string) |  |  |
| Domains | [string](#string) | repeated |  |
| PathPrefix | [string](#string) |  |  |
| Method | [string](#string) |  |  |
| Path | [string](#string) |  |  |
| Exported | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |





 

 

 


<a name="api-manager-v1-ApiManager"></a>

### ApiManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| Register | [RegisterRequest](#api-manager-v1-RegisterRequest) | [RegisterResponse](#api-manager-v1-RegisterResponse) |  |
| GetApis | [GetApisRequest](#api-manager-v1-GetApisRequest) | [GetApisResponse](#api-manager-v1-GetApisResponse) |  |
| GetServiceMethodApi | [GetServiceMethodApiRequest](#api-manager-v1-GetServiceMethodApiRequest) | [GetServiceMethodApiResponse](#api-manager-v1-GetServiceMethodApiResponse) |  |

 



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

