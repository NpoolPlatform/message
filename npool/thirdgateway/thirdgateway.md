# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/thirdgateway/thirdgateway.proto](#npool/thirdgateway/thirdgateway.proto)
    - [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest)
    - [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse)
    - [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest)
    - [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse)
    - [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest)
    - [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse)
    - [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest)
    - [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse)
  
    - [ThirdGateway](#third.gateway.v1.ThirdGateway)
  
- [npool/thirdgateway/thirdgateway.proto](#npool/thirdgateway/thirdgateway.proto)
    - [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest)
    - [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse)
    - [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest)
    - [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse)
    - [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest)
    - [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse)
    - [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest)
    - [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse)
  
    - [ThirdGateway](#third.gateway.v1.ThirdGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/thirdgateway/thirdgateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdgateway/thirdgateway.proto



<a name="third.gateway.v1.SendEmailCodeRequest"></a>

### SendEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EmailAddress | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendEmailCodeResponse"></a>

### SendEmailCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeRequest"></a>

### SendSMSCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeResponse"></a>

### SendSMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeRequest"></a>

### VerifyEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EmailAddress | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeResponse"></a>

### VerifyEmailCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.VerifySMSCodeRequest"></a>

### VerifySMSCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="third.gateway.v1.VerifySMSCodeResponse"></a>

### VerifySMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |





 

 

 


<a name="third.gateway.v1.ThirdGateway"></a>

### ThirdGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| SendSMSCode | [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest) | [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse) |  |
| VerifySMSCode | [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest) | [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse) |  |
| SendEmailCode | [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest) | [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse) |  |
| VerifyEmailCode | [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest) | [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse) |  |

 



<a name="npool/thirdgateway/thirdgateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdgateway/thirdgateway.proto



<a name="third.gateway.v1.SendEmailCodeRequest"></a>

### SendEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EmailAddress | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendEmailCodeResponse"></a>

### SendEmailCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeRequest"></a>

### SendSMSCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeResponse"></a>

### SendSMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeRequest"></a>

### VerifyEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| EmailAddress | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeResponse"></a>

### VerifyEmailCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.VerifySMSCodeRequest"></a>

### VerifySMSCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Code | [string](#string) |  |  |






<a name="third.gateway.v1.VerifySMSCodeResponse"></a>

### VerifySMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |





 

 

 


<a name="third.gateway.v1.ThirdGateway"></a>

### ThirdGateway
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| SendSMSCode | [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest) | [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse) |  |
| VerifySMSCode | [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest) | [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse) |  |
| SendEmailCode | [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest) | [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse) |  |
| VerifyEmailCode | [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest) | [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse) |  |

 



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

