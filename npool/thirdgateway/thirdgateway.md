# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/thirdgateway/thirdgateway.proto](#npool/thirdgateway/thirdgateway.proto)
    - [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate)
    - [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate)
    - [CreateAppEmailTemplateRequest](#third.gateway.v1.CreateAppEmailTemplateRequest)
    - [CreateAppEmailTemplateResponse](#third.gateway.v1.CreateAppEmailTemplateResponse)
    - [CreateAppSMSTemplateRequest](#third.gateway.v1.CreateAppSMSTemplateRequest)
    - [CreateAppSMSTemplateResponse](#third.gateway.v1.CreateAppSMSTemplateResponse)
    - [GetAppEmailTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest)
    - [GetAppEmailTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse)
    - [GetAppEmailTemplateRequest](#third.gateway.v1.GetAppEmailTemplateRequest)
    - [GetAppEmailTemplateResponse](#third.gateway.v1.GetAppEmailTemplateResponse)
    - [GetAppEmailTemplatesByAppRequest](#third.gateway.v1.GetAppEmailTemplatesByAppRequest)
    - [GetAppEmailTemplatesByAppResponse](#third.gateway.v1.GetAppEmailTemplatesByAppResponse)
    - [GetAppEmailTemplatesByOtherAppRequest](#third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest)
    - [GetAppEmailTemplatesByOtherAppResponse](#third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse)
    - [GetAppSMSTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest)
    - [GetAppSMSTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse)
    - [GetAppSMSTemplateRequest](#third.gateway.v1.GetAppSMSTemplateRequest)
    - [GetAppSMSTemplateResponse](#third.gateway.v1.GetAppSMSTemplateResponse)
    - [GetAppSMSTemplatesByAppRequest](#third.gateway.v1.GetAppSMSTemplatesByAppRequest)
    - [GetAppSMSTemplatesByAppResponse](#third.gateway.v1.GetAppSMSTemplatesByAppResponse)
    - [GetAppSMSTemplatesByOtherAppRequest](#third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest)
    - [GetAppSMSTemplatesByOtherAppResponse](#third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse)
    - [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest)
    - [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse)
    - [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest)
    - [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse)
    - [UpdateAppEmailTemplateRequest](#third.gateway.v1.UpdateAppEmailTemplateRequest)
    - [UpdateAppEmailTemplateResponse](#third.gateway.v1.UpdateAppEmailTemplateResponse)
    - [UpdateAppSMSTemplateRequest](#third.gateway.v1.UpdateAppSMSTemplateRequest)
    - [UpdateAppSMSTemplateResponse](#third.gateway.v1.UpdateAppSMSTemplateResponse)
    - [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest)
    - [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse)
    - [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest)
    - [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse)
  
    - [ThirdGateway](#third.gateway.v1.ThirdGateway)
  
- [npool/thirdgateway/thirdgateway.proto](#npool/thirdgateway/thirdgateway.proto)
    - [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate)
    - [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate)
    - [CreateAppEmailTemplateRequest](#third.gateway.v1.CreateAppEmailTemplateRequest)
    - [CreateAppEmailTemplateResponse](#third.gateway.v1.CreateAppEmailTemplateResponse)
    - [CreateAppSMSTemplateRequest](#third.gateway.v1.CreateAppSMSTemplateRequest)
    - [CreateAppSMSTemplateResponse](#third.gateway.v1.CreateAppSMSTemplateResponse)
    - [GetAppEmailTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest)
    - [GetAppEmailTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse)
    - [GetAppEmailTemplateRequest](#third.gateway.v1.GetAppEmailTemplateRequest)
    - [GetAppEmailTemplateResponse](#third.gateway.v1.GetAppEmailTemplateResponse)
    - [GetAppEmailTemplatesByAppRequest](#third.gateway.v1.GetAppEmailTemplatesByAppRequest)
    - [GetAppEmailTemplatesByAppResponse](#third.gateway.v1.GetAppEmailTemplatesByAppResponse)
    - [GetAppEmailTemplatesByOtherAppRequest](#third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest)
    - [GetAppEmailTemplatesByOtherAppResponse](#third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse)
    - [GetAppSMSTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest)
    - [GetAppSMSTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse)
    - [GetAppSMSTemplateRequest](#third.gateway.v1.GetAppSMSTemplateRequest)
    - [GetAppSMSTemplateResponse](#third.gateway.v1.GetAppSMSTemplateResponse)
    - [GetAppSMSTemplatesByAppRequest](#third.gateway.v1.GetAppSMSTemplatesByAppRequest)
    - [GetAppSMSTemplatesByAppResponse](#third.gateway.v1.GetAppSMSTemplatesByAppResponse)
    - [GetAppSMSTemplatesByOtherAppRequest](#third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest)
    - [GetAppSMSTemplatesByOtherAppResponse](#third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse)
    - [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest)
    - [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse)
    - [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest)
    - [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse)
    - [UpdateAppEmailTemplateRequest](#third.gateway.v1.UpdateAppEmailTemplateRequest)
    - [UpdateAppEmailTemplateResponse](#third.gateway.v1.UpdateAppEmailTemplateResponse)
    - [UpdateAppSMSTemplateRequest](#third.gateway.v1.UpdateAppSMSTemplateRequest)
    - [UpdateAppSMSTemplateResponse](#third.gateway.v1.UpdateAppSMSTemplateResponse)
    - [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest)
    - [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse)
    - [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest)
    - [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse)
  
    - [ThirdGateway](#third.gateway.v1.ThirdGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/thirdgateway/thirdgateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdgateway/thirdgateway.proto



<a name="third.gateway.v1.AppEmailTemplate"></a>

### AppEmailTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Sender | [string](#string) |  |  |
| ReplyTos | [string](#string) | repeated |  |
| CCTos | [string](#string) | repeated |  |
| Subject | [string](#string) |  |  |
| Body | [string](#string) |  |  |






<a name="third.gateway.v1.AppSMSTemplate"></a>

### AppSMSTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Subject | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.CreateAppEmailTemplateRequest"></a>

### CreateAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.CreateAppEmailTemplateResponse"></a>

### CreateAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.CreateAppSMSTemplateRequest"></a>

### CreateAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.CreateAppSMSTemplateResponse"></a>

### CreateAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest"></a>

### GetAppEmailTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse"></a>

### GetAppEmailTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateRequest"></a>

### GetAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateResponse"></a>

### GetAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByAppRequest"></a>

### GetAppEmailTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByAppResponse"></a>

### GetAppEmailTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest"></a>

### GetAppEmailTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse"></a>

### GetAppEmailTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest"></a>

### GetAppSMSTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse"></a>

### GetAppSMSTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateRequest"></a>

### GetAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateResponse"></a>

### GetAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByAppRequest"></a>

### GetAppSMSTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByAppResponse"></a>

### GetAppSMSTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest"></a>

### GetAppSMSTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse"></a>

### GetAppSMSTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) | repeated |  |






<a name="third.gateway.v1.SendEmailCodeRequest"></a>

### SendEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
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
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeResponse"></a>

### SendSMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.UpdateAppEmailTemplateRequest"></a>

### UpdateAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppEmailTemplateResponse"></a>

### UpdateAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppSMSTemplateRequest"></a>

### UpdateAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppSMSTemplateResponse"></a>

### UpdateAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeRequest"></a>

### VerifyEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
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
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
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
| CreateAppSMSTemplate | [CreateAppSMSTemplateRequest](#third.gateway.v1.CreateAppSMSTemplateRequest) | [CreateAppSMSTemplateResponse](#third.gateway.v1.CreateAppSMSTemplateResponse) |  |
| GetAppSMSTemplate | [GetAppSMSTemplateRequest](#third.gateway.v1.GetAppSMSTemplateRequest) | [GetAppSMSTemplateResponse](#third.gateway.v1.GetAppSMSTemplateResponse) |  |
| UpdateAppSMSTemplate | [UpdateAppSMSTemplateRequest](#third.gateway.v1.UpdateAppSMSTemplateRequest) | [UpdateAppSMSTemplateResponse](#third.gateway.v1.UpdateAppSMSTemplateResponse) |  |
| GetAppSMSTemplatesByApp | [GetAppSMSTemplatesByAppRequest](#third.gateway.v1.GetAppSMSTemplatesByAppRequest) | [GetAppSMSTemplatesByAppResponse](#third.gateway.v1.GetAppSMSTemplatesByAppResponse) |  |
| GetAppSMSTemplatesByOtherApp | [GetAppSMSTemplatesByOtherAppRequest](#third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest) | [GetAppSMSTemplatesByOtherAppResponse](#third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse) |  |
| GetAppSMSTemplateByAppLangUsedFor | [GetAppSMSTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest) | [GetAppSMSTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse) |  |
| CreateAppEmailTemplate | [CreateAppEmailTemplateRequest](#third.gateway.v1.CreateAppEmailTemplateRequest) | [CreateAppEmailTemplateResponse](#third.gateway.v1.CreateAppEmailTemplateResponse) |  |
| GetAppEmailTemplate | [GetAppEmailTemplateRequest](#third.gateway.v1.GetAppEmailTemplateRequest) | [GetAppEmailTemplateResponse](#third.gateway.v1.GetAppEmailTemplateResponse) |  |
| UpdateAppEmailTemplate | [UpdateAppEmailTemplateRequest](#third.gateway.v1.UpdateAppEmailTemplateRequest) | [UpdateAppEmailTemplateResponse](#third.gateway.v1.UpdateAppEmailTemplateResponse) |  |
| GetAppEmailTemplatesByApp | [GetAppEmailTemplatesByAppRequest](#third.gateway.v1.GetAppEmailTemplatesByAppRequest) | [GetAppEmailTemplatesByAppResponse](#third.gateway.v1.GetAppEmailTemplatesByAppResponse) |  |
| GetAppEmailTemplatesByOtherApp | [GetAppEmailTemplatesByOtherAppRequest](#third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest) | [GetAppEmailTemplatesByOtherAppResponse](#third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse) |  |
| GetAppEmailTemplateByAppLangUsedFor | [GetAppEmailTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest) | [GetAppEmailTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse) |  |
| SendSMSCode | [SendSMSCodeRequest](#third.gateway.v1.SendSMSCodeRequest) | [SendSMSCodeResponse](#third.gateway.v1.SendSMSCodeResponse) |  |
| VerifySMSCode | [VerifySMSCodeRequest](#third.gateway.v1.VerifySMSCodeRequest) | [VerifySMSCodeResponse](#third.gateway.v1.VerifySMSCodeResponse) |  |
| SendEmailCode | [SendEmailCodeRequest](#third.gateway.v1.SendEmailCodeRequest) | [SendEmailCodeResponse](#third.gateway.v1.SendEmailCodeResponse) |  |
| VerifyEmailCode | [VerifyEmailCodeRequest](#third.gateway.v1.VerifyEmailCodeRequest) | [VerifyEmailCodeResponse](#third.gateway.v1.VerifyEmailCodeResponse) |  |

 



<a name="npool/thirdgateway/thirdgateway.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/thirdgateway/thirdgateway.proto



<a name="third.gateway.v1.AppEmailTemplate"></a>

### AppEmailTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Sender | [string](#string) |  |  |
| ReplyTos | [string](#string) | repeated |  |
| CCTos | [string](#string) | repeated |  |
| Subject | [string](#string) |  |  |
| Body | [string](#string) |  |  |






<a name="third.gateway.v1.AppSMSTemplate"></a>

### AppSMSTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Subject | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.CreateAppEmailTemplateRequest"></a>

### CreateAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.CreateAppEmailTemplateResponse"></a>

### CreateAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.CreateAppSMSTemplateRequest"></a>

### CreateAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.CreateAppSMSTemplateResponse"></a>

### CreateAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest"></a>

### GetAppEmailTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse"></a>

### GetAppEmailTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateRequest"></a>

### GetAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplateResponse"></a>

### GetAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByAppRequest"></a>

### GetAppEmailTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByAppResponse"></a>

### GetAppEmailTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest"></a>

### GetAppEmailTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse"></a>

### GetAppEmailTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest"></a>

### GetAppSMSTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse"></a>

### GetAppSMSTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateRequest"></a>

### GetAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplateResponse"></a>

### GetAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByAppRequest"></a>

### GetAppSMSTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByAppResponse"></a>

### GetAppSMSTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) | repeated |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest"></a>

### GetAppSMSTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse"></a>

### GetAppSMSTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) | repeated |  |






<a name="third.gateway.v1.SendEmailCodeRequest"></a>

### SendEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
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
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| PhoneNO | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="third.gateway.v1.SendSMSCodeResponse"></a>

### SendSMSCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Code | [int32](#int32) |  |  |
| Message | [string](#string) |  |  |






<a name="third.gateway.v1.UpdateAppEmailTemplateRequest"></a>

### UpdateAppEmailTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppEmailTemplateResponse"></a>

### UpdateAppEmailTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppEmailTemplate](#third.gateway.v1.AppEmailTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppSMSTemplateRequest"></a>

### UpdateAppSMSTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.UpdateAppSMSTemplateResponse"></a>

### UpdateAppSMSTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppSMSTemplate](#third.gateway.v1.AppSMSTemplate) |  |  |






<a name="third.gateway.v1.VerifyEmailCodeRequest"></a>

### VerifyEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
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
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
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
| CreateAppSMSTemplate | [CreateAppSMSTemplateRequest](#third.gateway.v1.CreateAppSMSTemplateRequest) | [CreateAppSMSTemplateResponse](#third.gateway.v1.CreateAppSMSTemplateResponse) |  |
| GetAppSMSTemplate | [GetAppSMSTemplateRequest](#third.gateway.v1.GetAppSMSTemplateRequest) | [GetAppSMSTemplateResponse](#third.gateway.v1.GetAppSMSTemplateResponse) |  |
| UpdateAppSMSTemplate | [UpdateAppSMSTemplateRequest](#third.gateway.v1.UpdateAppSMSTemplateRequest) | [UpdateAppSMSTemplateResponse](#third.gateway.v1.UpdateAppSMSTemplateResponse) |  |
| GetAppSMSTemplatesByApp | [GetAppSMSTemplatesByAppRequest](#third.gateway.v1.GetAppSMSTemplatesByAppRequest) | [GetAppSMSTemplatesByAppResponse](#third.gateway.v1.GetAppSMSTemplatesByAppResponse) |  |
| GetAppSMSTemplatesByOtherApp | [GetAppSMSTemplatesByOtherAppRequest](#third.gateway.v1.GetAppSMSTemplatesByOtherAppRequest) | [GetAppSMSTemplatesByOtherAppResponse](#third.gateway.v1.GetAppSMSTemplatesByOtherAppResponse) |  |
| GetAppSMSTemplateByAppLangUsedFor | [GetAppSMSTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForRequest) | [GetAppSMSTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppSMSTemplateByAppLangUsedForResponse) |  |
| CreateAppEmailTemplate | [CreateAppEmailTemplateRequest](#third.gateway.v1.CreateAppEmailTemplateRequest) | [CreateAppEmailTemplateResponse](#third.gateway.v1.CreateAppEmailTemplateResponse) |  |
| GetAppEmailTemplate | [GetAppEmailTemplateRequest](#third.gateway.v1.GetAppEmailTemplateRequest) | [GetAppEmailTemplateResponse](#third.gateway.v1.GetAppEmailTemplateResponse) |  |
| UpdateAppEmailTemplate | [UpdateAppEmailTemplateRequest](#third.gateway.v1.UpdateAppEmailTemplateRequest) | [UpdateAppEmailTemplateResponse](#third.gateway.v1.UpdateAppEmailTemplateResponse) |  |
| GetAppEmailTemplatesByApp | [GetAppEmailTemplatesByAppRequest](#third.gateway.v1.GetAppEmailTemplatesByAppRequest) | [GetAppEmailTemplatesByAppResponse](#third.gateway.v1.GetAppEmailTemplatesByAppResponse) |  |
| GetAppEmailTemplatesByOtherApp | [GetAppEmailTemplatesByOtherAppRequest](#third.gateway.v1.GetAppEmailTemplatesByOtherAppRequest) | [GetAppEmailTemplatesByOtherAppResponse](#third.gateway.v1.GetAppEmailTemplatesByOtherAppResponse) |  |
| GetAppEmailTemplateByAppLangUsedFor | [GetAppEmailTemplateByAppLangUsedForRequest](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForRequest) | [GetAppEmailTemplateByAppLangUsedForResponse](#third.gateway.v1.GetAppEmailTemplateByAppLangUsedForResponse) |  |
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

