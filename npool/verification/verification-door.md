# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/verification/verification-door.proto](#npool/verification/verification-door.proto)
    - [CaptcherResp](#verification.door.v1.CaptcherResp)
    - [DeleteUserGoogleAuthRequest](#verification.door.v1.DeleteUserGoogleAuthRequest)
    - [DeleteUserGoogleAuthResponse](#verification.door.v1.DeleteUserGoogleAuthResponse)
    - [GetCaptcherImgRequest](#verification.door.v1.GetCaptcherImgRequest)
    - [GetCaptcherImgResponse](#verification.door.v1.GetCaptcherImgResponse)
    - [GetQRcodeURLRequest](#verification.door.v1.GetQRcodeURLRequest)
    - [GetQRcodeURLResponse](#verification.door.v1.GetQRcodeURLResponse)
    - [QRCodeInfo](#verification.door.v1.QRCodeInfo)
    - [SendEmailRequest](#verification.door.v1.SendEmailRequest)
    - [SendEmailResponse](#verification.door.v1.SendEmailResponse)
    - [SendSmsRequest](#verification.door.v1.SendSmsRequest)
    - [SendSmsResponse](#verification.door.v1.SendSmsResponse)
    - [SendUserSiteContactEmailRequest](#verification.door.v1.SendUserSiteContactEmailRequest)
    - [SendUserSiteContactEmailResponse](#verification.door.v1.SendUserSiteContactEmailResponse)
    - [VerifyCaptcherRequest](#verification.door.v1.VerifyCaptcherRequest)
    - [VerifyCaptcherResponse](#verification.door.v1.VerifyCaptcherResponse)
    - [VerifyCodeRequest](#verification.door.v1.VerifyCodeRequest)
    - [VerifyCodeResponse](#verification.door.v1.VerifyCodeResponse)
    - [VerifyCodeWithUserIDRequest](#verification.door.v1.VerifyCodeWithUserIDRequest)
    - [VerifyCodeWithUserIDResponse](#verification.door.v1.VerifyCodeWithUserIDResponse)
    - [VerifyGoogleAuthRequest](#verification.door.v1.VerifyGoogleAuthRequest)
    - [VerifyGoogleAuthResponse](#verification.door.v1.VerifyGoogleAuthResponse)
    - [VerifyGoogleRecaptchaRequest](#verification.door.v1.VerifyGoogleRecaptchaRequest)
    - [VerifyGoogleRecaptchaResponse](#verification.door.v1.VerifyGoogleRecaptchaResponse)
  
    - [VerificationDoor](#verification.door.v1.VerificationDoor)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/verification/verification-door.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/verification/verification-door.proto



<a name="verification.door.v1.CaptcherResp"></a>

### CaptcherResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Base64URL | [string](#string) |  |  |






<a name="verification.door.v1.DeleteUserGoogleAuthRequest"></a>

### DeleteUserGoogleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="verification.door.v1.DeleteUserGoogleAuthResponse"></a>

### DeleteUserGoogleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.GetCaptcherImgRequest"></a>

### GetCaptcherImgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="verification.door.v1.GetCaptcherImgResponse"></a>

### GetCaptcherImgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CaptcherResp](#verification.door.v1.CaptcherResp) |  |  |






<a name="verification.door.v1.GetQRcodeURLRequest"></a>

### GetQRcodeURLRequest
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Username | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="verification.door.v1.GetQRcodeURLResponse"></a>

### GetQRcodeURLResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [QRCodeInfo](#verification.door.v1.QRCodeInfo) |  |  |






<a name="verification.door.v1.QRCodeInfo"></a>

### QRCodeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CodeURL | [string](#string) |  |  |
| Secret | [string](#string) |  |  |






<a name="verification.door.v1.SendEmailRequest"></a>

### SendEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Email | [string](#string) |  |  |
| Intention | [string](#string) |  |  |
| Lang | [string](#string) |  |  |
| Username | [string](#string) |  |  |






<a name="verification.door.v1.SendEmailResponse"></a>

### SendEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.SendSmsRequest"></a>

### SendSmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Phone | [string](#string) |  |  |
| Lang | [string](#string) |  |  |
| Intention | [string](#string) |  |  |






<a name="verification.door.v1.SendSmsResponse"></a>

### SendSmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.SendUserSiteContactEmailRequest"></a>

### SendUserSiteContactEmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| From | [string](#string) |  |  |
| To | [string](#string) |  |  |
| Text | [string](#string) |  |  |
| Subject | [string](#string) |  |  |
| Username | [string](#string) |  |  |






<a name="verification.door.v1.SendUserSiteContactEmailResponse"></a>

### SendUserSiteContactEmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCaptcherRequest"></a>

### VerifyCaptcherRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Input | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCaptcherResponse"></a>

### VerifyCaptcherResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeRequest"></a>

### VerifyCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Param | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| VerifyType | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeResponse"></a>

### VerifyCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeWithUserIDRequest"></a>

### VerifyCodeWithUserIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Param | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| VerifyType | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="verification.door.v1.VerifyCodeWithUserIDResponse"></a>

### VerifyCodeWithUserIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleAuthRequest"></a>

### VerifyGoogleAuthRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleAuthResponse"></a>

### VerifyGoogleAuthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleRecaptchaRequest"></a>

### VerifyGoogleRecaptchaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Response | [string](#string) |  |  |






<a name="verification.door.v1.VerifyGoogleRecaptchaResponse"></a>

### VerifyGoogleRecaptchaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [bool](#bool) |  |  |





 

 

 


<a name="verification.door.v1.VerificationDoor"></a>

### VerificationDoor
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| GetQRcodeURL | [GetQRcodeURLRequest](#verification.door.v1.GetQRcodeURLRequest) | [GetQRcodeURLResponse](#verification.door.v1.GetQRcodeURLResponse) | get google authentication qr code url |
| VerifyGoogleAuth | [VerifyGoogleAuthRequest](#verification.door.v1.VerifyGoogleAuthRequest) | [VerifyGoogleAuthResponse](#verification.door.v1.VerifyGoogleAuthResponse) | verify user google authentication(user&#39;s input code) |
| DeleteUserGoogleAuth | [DeleteUserGoogleAuthRequest](#verification.door.v1.DeleteUserGoogleAuthRequest) | [DeleteUserGoogleAuthResponse](#verification.door.v1.DeleteUserGoogleAuthResponse) | delete user google authentication record |
| SendEmail | [SendEmailRequest](#verification.door.v1.SendEmailRequest) | [SendEmailResponse](#verification.door.v1.SendEmailResponse) | send email to user |
| SendSms | [SendSmsRequest](#verification.door.v1.SendSmsRequest) | [SendSmsResponse](#verification.door.v1.SendSmsResponse) | send sms to user(todo......) |
| VerifyCodeWithUserID | [VerifyCodeWithUserIDRequest](#verification.door.v1.VerifyCodeWithUserIDRequest) | [VerifyCodeWithUserIDResponse](#verification.door.v1.VerifyCodeWithUserIDResponse) | verify code with user id. |
| VerifyCode | [VerifyCodeRequest](#verification.door.v1.VerifyCodeRequest) | [VerifyCodeResponse](#verification.door.v1.VerifyCodeResponse) | verify code user input. (can verify email code and sms code, verify sms code is todo......) |
| VerifyGoogleRecaptcha | [VerifyGoogleRecaptchaRequest](#verification.door.v1.VerifyGoogleRecaptchaRequest) | [VerifyGoogleRecaptchaResponse](#verification.door.v1.VerifyGoogleRecaptchaResponse) | verify google recaptcha. |
| GetCaptcherImg | [GetCaptcherImgRequest](#verification.door.v1.GetCaptcherImgRequest) | [GetCaptcherImgResponse](#verification.door.v1.GetCaptcherImgResponse) | get captcher imgine url |
| VerifyCaptcher | [VerifyCaptcherRequest](#verification.door.v1.VerifyCaptcherRequest) | [VerifyCaptcherResponse](#verification.door.v1.VerifyCaptcherResponse) | verify captcher input |
| SendUserSiteContactEmail | [SendUserSiteContactEmailRequest](#verification.door.v1.SendUserSiteContactEmailRequest) | [SendUserSiteContactEmailResponse](#verification.door.v1.SendUserSiteContactEmailResponse) | Send user site contact email |

 



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

