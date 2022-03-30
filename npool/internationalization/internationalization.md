# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/internationalization/internationalization.proto](#npool_internationalization_internationalization-proto)
    - [AddLangRequest](#internationalization-v1-AddLangRequest)
    - [AddLangResponse](#internationalization-v1-AddLangResponse)
    - [AppLang](#internationalization-v1-AppLang)
    - [AppLangInfo](#internationalization-v1-AppLangInfo)
    - [Country](#internationalization-v1-Country)
    - [CreateAppLangForOtherAppRequest](#internationalization-v1-CreateAppLangForOtherAppRequest)
    - [CreateAppLangForOtherAppResponse](#internationalization-v1-CreateAppLangForOtherAppResponse)
    - [CreateAppLangRequest](#internationalization-v1-CreateAppLangRequest)
    - [CreateAppLangResponse](#internationalization-v1-CreateAppLangResponse)
    - [CreateCountryRequest](#internationalization-v1-CreateCountryRequest)
    - [CreateCountryResponse](#internationalization-v1-CreateCountryResponse)
    - [CreateMessageForOtherAppRequest](#internationalization-v1-CreateMessageForOtherAppRequest)
    - [CreateMessageForOtherAppResponse](#internationalization-v1-CreateMessageForOtherAppResponse)
    - [CreateMessageRequest](#internationalization-v1-CreateMessageRequest)
    - [CreateMessageResponse](#internationalization-v1-CreateMessageResponse)
    - [CreateMessagesRequest](#internationalization-v1-CreateMessagesRequest)
    - [CreateMessagesResponse](#internationalization-v1-CreateMessagesResponse)
    - [GetAppLangInfosByAppRequest](#internationalization-v1-GetAppLangInfosByAppRequest)
    - [GetAppLangInfosByAppResponse](#internationalization-v1-GetAppLangInfosByAppResponse)
    - [GetAppLangInfosByOtherAppRequest](#internationalization-v1-GetAppLangInfosByOtherAppRequest)
    - [GetAppLangInfosByOtherAppResponse](#internationalization-v1-GetAppLangInfosByOtherAppResponse)
    - [GetAppLangRequest](#internationalization-v1-GetAppLangRequest)
    - [GetAppLangResponse](#internationalization-v1-GetAppLangResponse)
    - [GetAppLangsByAppRequest](#internationalization-v1-GetAppLangsByAppRequest)
    - [GetAppLangsByAppResponse](#internationalization-v1-GetAppLangsByAppResponse)
    - [GetCountriesRequest](#internationalization-v1-GetCountriesRequest)
    - [GetCountriesResponse](#internationalization-v1-GetCountriesResponse)
    - [GetLangRequest](#internationalization-v1-GetLangRequest)
    - [GetLangResponse](#internationalization-v1-GetLangResponse)
    - [GetLangsRequest](#internationalization-v1-GetLangsRequest)
    - [GetLangsResponse](#internationalization-v1-GetLangsResponse)
    - [GetMessageByAppLangMessageRequest](#internationalization-v1-GetMessageByAppLangMessageRequest)
    - [GetMessageByAppLangMessageResponse](#internationalization-v1-GetMessageByAppLangMessageResponse)
    - [GetMessagesByAppLangRequest](#internationalization-v1-GetMessagesByAppLangRequest)
    - [GetMessagesByAppLangResponse](#internationalization-v1-GetMessagesByAppLangResponse)
    - [GetMessagesByOtherAppLangRequest](#internationalization-v1-GetMessagesByOtherAppLangRequest)
    - [GetMessagesByOtherAppLangResponse](#internationalization-v1-GetMessagesByOtherAppLangResponse)
    - [Lang](#internationalization-v1-Lang)
    - [Message](#internationalization-v1-Message)
    - [UpdateLangRequest](#internationalization-v1-UpdateLangRequest)
    - [UpdateLangResponse](#internationalization-v1-UpdateLangResponse)
    - [UpdateMessageRequest](#internationalization-v1-UpdateMessageRequest)
    - [UpdateMessageResponse](#internationalization-v1-UpdateMessageResponse)
    - [UpdateMessagesRequest](#internationalization-v1-UpdateMessagesRequest)
    - [UpdateMessagesResponse](#internationalization-v1-UpdateMessagesResponse)
  
    - [Internationalization](#internationalization-v1-Internationalization)
  
- [npool/internationalization/internationalization.proto](#npool_internationalization_internationalization-proto)
    - [AddLangRequest](#internationalization-v1-AddLangRequest)
    - [AddLangResponse](#internationalization-v1-AddLangResponse)
    - [AppLang](#internationalization-v1-AppLang)
    - [AppLangInfo](#internationalization-v1-AppLangInfo)
    - [Country](#internationalization-v1-Country)
    - [CreateAppLangForOtherAppRequest](#internationalization-v1-CreateAppLangForOtherAppRequest)
    - [CreateAppLangForOtherAppResponse](#internationalization-v1-CreateAppLangForOtherAppResponse)
    - [CreateAppLangRequest](#internationalization-v1-CreateAppLangRequest)
    - [CreateAppLangResponse](#internationalization-v1-CreateAppLangResponse)
    - [CreateCountryRequest](#internationalization-v1-CreateCountryRequest)
    - [CreateCountryResponse](#internationalization-v1-CreateCountryResponse)
    - [CreateMessageForOtherAppRequest](#internationalization-v1-CreateMessageForOtherAppRequest)
    - [CreateMessageForOtherAppResponse](#internationalization-v1-CreateMessageForOtherAppResponse)
    - [CreateMessageRequest](#internationalization-v1-CreateMessageRequest)
    - [CreateMessageResponse](#internationalization-v1-CreateMessageResponse)
    - [CreateMessagesRequest](#internationalization-v1-CreateMessagesRequest)
    - [CreateMessagesResponse](#internationalization-v1-CreateMessagesResponse)
    - [GetAppLangInfosByAppRequest](#internationalization-v1-GetAppLangInfosByAppRequest)
    - [GetAppLangInfosByAppResponse](#internationalization-v1-GetAppLangInfosByAppResponse)
    - [GetAppLangInfosByOtherAppRequest](#internationalization-v1-GetAppLangInfosByOtherAppRequest)
    - [GetAppLangInfosByOtherAppResponse](#internationalization-v1-GetAppLangInfosByOtherAppResponse)
    - [GetAppLangRequest](#internationalization-v1-GetAppLangRequest)
    - [GetAppLangResponse](#internationalization-v1-GetAppLangResponse)
    - [GetAppLangsByAppRequest](#internationalization-v1-GetAppLangsByAppRequest)
    - [GetAppLangsByAppResponse](#internationalization-v1-GetAppLangsByAppResponse)
    - [GetCountriesRequest](#internationalization-v1-GetCountriesRequest)
    - [GetCountriesResponse](#internationalization-v1-GetCountriesResponse)
    - [GetLangRequest](#internationalization-v1-GetLangRequest)
    - [GetLangResponse](#internationalization-v1-GetLangResponse)
    - [GetLangsRequest](#internationalization-v1-GetLangsRequest)
    - [GetLangsResponse](#internationalization-v1-GetLangsResponse)
    - [GetMessageByAppLangMessageRequest](#internationalization-v1-GetMessageByAppLangMessageRequest)
    - [GetMessageByAppLangMessageResponse](#internationalization-v1-GetMessageByAppLangMessageResponse)
    - [GetMessagesByAppLangRequest](#internationalization-v1-GetMessagesByAppLangRequest)
    - [GetMessagesByAppLangResponse](#internationalization-v1-GetMessagesByAppLangResponse)
    - [GetMessagesByOtherAppLangRequest](#internationalization-v1-GetMessagesByOtherAppLangRequest)
    - [GetMessagesByOtherAppLangResponse](#internationalization-v1-GetMessagesByOtherAppLangResponse)
    - [Lang](#internationalization-v1-Lang)
    - [Message](#internationalization-v1-Message)
    - [UpdateLangRequest](#internationalization-v1-UpdateLangRequest)
    - [UpdateLangResponse](#internationalization-v1-UpdateLangResponse)
    - [UpdateMessageRequest](#internationalization-v1-UpdateMessageRequest)
    - [UpdateMessageResponse](#internationalization-v1-UpdateMessageResponse)
    - [UpdateMessagesRequest](#internationalization-v1-UpdateMessagesRequest)
    - [UpdateMessagesResponse](#internationalization-v1-UpdateMessagesResponse)
  
    - [Internationalization](#internationalization-v1-Internationalization)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_internationalization_internationalization-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/internationalization/internationalization.proto



<a name="internationalization-v1-AddLangRequest"></a>

### AddLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-AddLangResponse"></a>

### AddLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-AppLang"></a>

### AppLang



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-AppLangInfo"></a>

### AppLangInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppLang | [AppLang](#internationalization-v1-AppLang) |  |  |
| Lang | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-Country"></a>

### Country



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Flag | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| Short | [string](#string) |  |  |






<a name="internationalization-v1-CreateAppLangForOtherAppRequest"></a>

### CreateAppLangForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangForOtherAppResponse"></a>

### CreateAppLangForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangRequest"></a>

### CreateAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangResponse"></a>

### CreateAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateCountryRequest"></a>

### CreateCountryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Country](#internationalization-v1-Country) |  |  |






<a name="internationalization-v1-CreateCountryResponse"></a>

### CreateCountryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Country](#internationalization-v1-Country) |  |  |






<a name="internationalization-v1-CreateMessageForOtherAppRequest"></a>

### CreateMessageForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageForOtherAppResponse"></a>

### CreateMessageForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageRequest"></a>

### CreateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageResponse"></a>

### CreateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessagesRequest"></a>

### CreateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-CreateMessagesResponse"></a>

### CreateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-GetAppLangInfosByAppRequest"></a>

### GetAppLangInfosByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangInfosByAppResponse"></a>

### GetAppLangInfosByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLangInfo](#internationalization-v1-AppLangInfo) | repeated |  |






<a name="internationalization-v1-GetAppLangInfosByOtherAppRequest"></a>

### GetAppLangInfosByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangInfosByOtherAppResponse"></a>

### GetAppLangInfosByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLangInfo](#internationalization-v1-AppLangInfo) | repeated |  |






<a name="internationalization-v1-GetAppLangRequest"></a>

### GetAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangResponse"></a>

### GetAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-GetAppLangsByAppRequest"></a>

### GetAppLangsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangsByAppResponse"></a>

### GetAppLangsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLang](#internationalization-v1-AppLang) | repeated |  |






<a name="internationalization-v1-GetCountriesRequest"></a>

### GetCountriesRequest







<a name="internationalization-v1-GetCountriesResponse"></a>

### GetCountriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Country](#internationalization-v1-Country) | repeated |  |






<a name="internationalization-v1-GetLangRequest"></a>

### GetLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="internationalization-v1-GetLangResponse"></a>

### GetLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-GetLangsRequest"></a>

### GetLangsRequest







<a name="internationalization-v1-GetLangsResponse"></a>

### GetLangsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Lang](#internationalization-v1-Lang) | repeated |  |






<a name="internationalization-v1-GetMessageByAppLangMessageRequest"></a>

### GetMessageByAppLangMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| MessageID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessageByAppLangMessageResponse"></a>

### GetMessageByAppLangMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-GetMessagesByAppLangRequest"></a>

### GetMessagesByAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessagesByAppLangResponse"></a>

### GetMessagesByAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-GetMessagesByOtherAppLangRequest"></a>

### GetMessagesByOtherAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessagesByOtherAppLangResponse"></a>

### GetMessagesByOtherAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-Lang"></a>

### Lang



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Lang | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Short | [string](#string) |  |  |






<a name="internationalization-v1-Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| MessageID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| BatchGet | [bool](#bool) |  |  |






<a name="internationalization-v1-UpdateLangRequest"></a>

### UpdateLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-UpdateLangResponse"></a>

### UpdateLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-UpdateMessageRequest"></a>

### UpdateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-UpdateMessageResponse"></a>

### UpdateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-UpdateMessagesRequest"></a>

### UpdateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-UpdateMessagesResponse"></a>

### UpdateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |





 

 

 


<a name="internationalization-v1-Internationalization"></a>

### Internationalization
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| AddLang | [AddLangRequest](#internationalization-v1-AddLangRequest) | [AddLangResponse](#internationalization-v1-AddLangResponse) |  |
| GetLang | [GetLangRequest](#internationalization-v1-GetLangRequest) | [GetLangResponse](#internationalization-v1-GetLangResponse) |  |
| UpdateLang | [UpdateLangRequest](#internationalization-v1-UpdateLangRequest) | [UpdateLangResponse](#internationalization-v1-UpdateLangResponse) |  |
| GetLangs | [GetLangsRequest](#internationalization-v1-GetLangsRequest) | [GetLangsResponse](#internationalization-v1-GetLangsResponse) |  |
| CreateAppLang | [CreateAppLangRequest](#internationalization-v1-CreateAppLangRequest) | [CreateAppLangResponse](#internationalization-v1-CreateAppLangResponse) |  |
| CreateAppLangForOtherApp | [CreateAppLangForOtherAppRequest](#internationalization-v1-CreateAppLangForOtherAppRequest) | [CreateAppLangForOtherAppResponse](#internationalization-v1-CreateAppLangForOtherAppResponse) |  |
| GetAppLang | [GetAppLangRequest](#internationalization-v1-GetAppLangRequest) | [GetAppLangResponse](#internationalization-v1-GetAppLangResponse) |  |
| GetAppLangsByApp | [GetAppLangsByAppRequest](#internationalization-v1-GetAppLangsByAppRequest) | [GetAppLangsByAppResponse](#internationalization-v1-GetAppLangsByAppResponse) |  |
| GetAppLangInfosByApp | [GetAppLangInfosByAppRequest](#internationalization-v1-GetAppLangInfosByAppRequest) | [GetAppLangInfosByAppResponse](#internationalization-v1-GetAppLangInfosByAppResponse) |  |
| GetAppLangInfosByOtherApp | [GetAppLangInfosByOtherAppRequest](#internationalization-v1-GetAppLangInfosByOtherAppRequest) | [GetAppLangInfosByOtherAppResponse](#internationalization-v1-GetAppLangInfosByOtherAppResponse) |  |
| CreateMessage | [CreateMessageRequest](#internationalization-v1-CreateMessageRequest) | [CreateMessageResponse](#internationalization-v1-CreateMessageResponse) |  |
| CreateMessageForOtherApp | [CreateMessageForOtherAppRequest](#internationalization-v1-CreateMessageForOtherAppRequest) | [CreateMessageForOtherAppResponse](#internationalization-v1-CreateMessageForOtherAppResponse) |  |
| CreateMessages | [CreateMessagesRequest](#internationalization-v1-CreateMessagesRequest) | [CreateMessagesResponse](#internationalization-v1-CreateMessagesResponse) |  |
| UpdateMessage | [UpdateMessageRequest](#internationalization-v1-UpdateMessageRequest) | [UpdateMessageResponse](#internationalization-v1-UpdateMessageResponse) |  |
| UpdateMessages | [UpdateMessagesRequest](#internationalization-v1-UpdateMessagesRequest) | [UpdateMessagesResponse](#internationalization-v1-UpdateMessagesResponse) |  |
| GetMessagesByAppLang | [GetMessagesByAppLangRequest](#internationalization-v1-GetMessagesByAppLangRequest) | [GetMessagesByAppLangResponse](#internationalization-v1-GetMessagesByAppLangResponse) |  |
| GetMessagesByOtherAppLang | [GetMessagesByOtherAppLangRequest](#internationalization-v1-GetMessagesByOtherAppLangRequest) | [GetMessagesByOtherAppLangResponse](#internationalization-v1-GetMessagesByOtherAppLangResponse) |  |
| GetMessageByAppLangMessage | [GetMessageByAppLangMessageRequest](#internationalization-v1-GetMessageByAppLangMessageRequest) | [GetMessageByAppLangMessageResponse](#internationalization-v1-GetMessageByAppLangMessageResponse) |  |
| CreateCountry | [CreateCountryRequest](#internationalization-v1-CreateCountryRequest) | [CreateCountryResponse](#internationalization-v1-CreateCountryResponse) |  |
| GetCountries | [GetCountriesRequest](#internationalization-v1-GetCountriesRequest) | [GetCountriesResponse](#internationalization-v1-GetCountriesResponse) |  |

 



<a name="npool_internationalization_internationalization-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/internationalization/internationalization.proto



<a name="internationalization-v1-AddLangRequest"></a>

### AddLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-AddLangResponse"></a>

### AddLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-AppLang"></a>

### AppLang



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-AppLangInfo"></a>

### AppLangInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppLang | [AppLang](#internationalization-v1-AppLang) |  |  |
| Lang | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-Country"></a>

### Country



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Flag | [string](#string) |  |  |
| Code | [string](#string) |  |  |
| Short | [string](#string) |  |  |






<a name="internationalization-v1-CreateAppLangForOtherAppRequest"></a>

### CreateAppLangForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangForOtherAppResponse"></a>

### CreateAppLangForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangRequest"></a>

### CreateAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateAppLangResponse"></a>

### CreateAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-CreateCountryRequest"></a>

### CreateCountryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Country](#internationalization-v1-Country) |  |  |






<a name="internationalization-v1-CreateCountryResponse"></a>

### CreateCountryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Country](#internationalization-v1-Country) |  |  |






<a name="internationalization-v1-CreateMessageForOtherAppRequest"></a>

### CreateMessageForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageForOtherAppResponse"></a>

### CreateMessageForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageRequest"></a>

### CreateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessageResponse"></a>

### CreateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-CreateMessagesRequest"></a>

### CreateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-CreateMessagesResponse"></a>

### CreateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-GetAppLangInfosByAppRequest"></a>

### GetAppLangInfosByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangInfosByAppResponse"></a>

### GetAppLangInfosByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLangInfo](#internationalization-v1-AppLangInfo) | repeated |  |






<a name="internationalization-v1-GetAppLangInfosByOtherAppRequest"></a>

### GetAppLangInfosByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangInfosByOtherAppResponse"></a>

### GetAppLangInfosByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLangInfo](#internationalization-v1-AppLangInfo) | repeated |  |






<a name="internationalization-v1-GetAppLangRequest"></a>

### GetAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangResponse"></a>

### GetAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppLang](#internationalization-v1-AppLang) |  |  |






<a name="internationalization-v1-GetAppLangsByAppRequest"></a>

### GetAppLangsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="internationalization-v1-GetAppLangsByAppResponse"></a>

### GetAppLangsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [AppLang](#internationalization-v1-AppLang) | repeated |  |






<a name="internationalization-v1-GetCountriesRequest"></a>

### GetCountriesRequest







<a name="internationalization-v1-GetCountriesResponse"></a>

### GetCountriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Country](#internationalization-v1-Country) | repeated |  |






<a name="internationalization-v1-GetLangRequest"></a>

### GetLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="internationalization-v1-GetLangResponse"></a>

### GetLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-GetLangsRequest"></a>

### GetLangsRequest







<a name="internationalization-v1-GetLangsResponse"></a>

### GetLangsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Lang](#internationalization-v1-Lang) | repeated |  |






<a name="internationalization-v1-GetMessageByAppLangMessageRequest"></a>

### GetMessageByAppLangMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| MessageID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessageByAppLangMessageResponse"></a>

### GetMessageByAppLangMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-GetMessagesByAppLangRequest"></a>

### GetMessagesByAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessagesByAppLangResponse"></a>

### GetMessagesByAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-GetMessagesByOtherAppLangRequest"></a>

### GetMessagesByOtherAppLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |






<a name="internationalization-v1-GetMessagesByOtherAppLangResponse"></a>

### GetMessagesByOtherAppLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-Lang"></a>

### Lang



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Lang | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Short | [string](#string) |  |  |






<a name="internationalization-v1-Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| MessageID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| BatchGet | [bool](#bool) |  |  |






<a name="internationalization-v1-UpdateLangRequest"></a>

### UpdateLangRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-UpdateLangResponse"></a>

### UpdateLangResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Lang](#internationalization-v1-Lang) |  |  |






<a name="internationalization-v1-UpdateMessageRequest"></a>

### UpdateMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-UpdateMessageResponse"></a>

### UpdateMessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Message](#internationalization-v1-Message) |  |  |






<a name="internationalization-v1-UpdateMessagesRequest"></a>

### UpdateMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |






<a name="internationalization-v1-UpdateMessagesResponse"></a>

### UpdateMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Message](#internationalization-v1-Message) | repeated |  |





 

 

 


<a name="internationalization-v1-Internationalization"></a>

### Internationalization
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| AddLang | [AddLangRequest](#internationalization-v1-AddLangRequest) | [AddLangResponse](#internationalization-v1-AddLangResponse) |  |
| GetLang | [GetLangRequest](#internationalization-v1-GetLangRequest) | [GetLangResponse](#internationalization-v1-GetLangResponse) |  |
| UpdateLang | [UpdateLangRequest](#internationalization-v1-UpdateLangRequest) | [UpdateLangResponse](#internationalization-v1-UpdateLangResponse) |  |
| GetLangs | [GetLangsRequest](#internationalization-v1-GetLangsRequest) | [GetLangsResponse](#internationalization-v1-GetLangsResponse) |  |
| CreateAppLang | [CreateAppLangRequest](#internationalization-v1-CreateAppLangRequest) | [CreateAppLangResponse](#internationalization-v1-CreateAppLangResponse) |  |
| CreateAppLangForOtherApp | [CreateAppLangForOtherAppRequest](#internationalization-v1-CreateAppLangForOtherAppRequest) | [CreateAppLangForOtherAppResponse](#internationalization-v1-CreateAppLangForOtherAppResponse) |  |
| GetAppLang | [GetAppLangRequest](#internationalization-v1-GetAppLangRequest) | [GetAppLangResponse](#internationalization-v1-GetAppLangResponse) |  |
| GetAppLangsByApp | [GetAppLangsByAppRequest](#internationalization-v1-GetAppLangsByAppRequest) | [GetAppLangsByAppResponse](#internationalization-v1-GetAppLangsByAppResponse) |  |
| GetAppLangInfosByApp | [GetAppLangInfosByAppRequest](#internationalization-v1-GetAppLangInfosByAppRequest) | [GetAppLangInfosByAppResponse](#internationalization-v1-GetAppLangInfosByAppResponse) |  |
| GetAppLangInfosByOtherApp | [GetAppLangInfosByOtherAppRequest](#internationalization-v1-GetAppLangInfosByOtherAppRequest) | [GetAppLangInfosByOtherAppResponse](#internationalization-v1-GetAppLangInfosByOtherAppResponse) |  |
| CreateMessage | [CreateMessageRequest](#internationalization-v1-CreateMessageRequest) | [CreateMessageResponse](#internationalization-v1-CreateMessageResponse) |  |
| CreateMessageForOtherApp | [CreateMessageForOtherAppRequest](#internationalization-v1-CreateMessageForOtherAppRequest) | [CreateMessageForOtherAppResponse](#internationalization-v1-CreateMessageForOtherAppResponse) |  |
| CreateMessages | [CreateMessagesRequest](#internationalization-v1-CreateMessagesRequest) | [CreateMessagesResponse](#internationalization-v1-CreateMessagesResponse) |  |
| UpdateMessage | [UpdateMessageRequest](#internationalization-v1-UpdateMessageRequest) | [UpdateMessageResponse](#internationalization-v1-UpdateMessageResponse) |  |
| UpdateMessages | [UpdateMessagesRequest](#internationalization-v1-UpdateMessagesRequest) | [UpdateMessagesResponse](#internationalization-v1-UpdateMessagesResponse) |  |
| GetMessagesByAppLang | [GetMessagesByAppLangRequest](#internationalization-v1-GetMessagesByAppLangRequest) | [GetMessagesByAppLangResponse](#internationalization-v1-GetMessagesByAppLangResponse) |  |
| GetMessagesByOtherAppLang | [GetMessagesByOtherAppLangRequest](#internationalization-v1-GetMessagesByOtherAppLangRequest) | [GetMessagesByOtherAppLangResponse](#internationalization-v1-GetMessagesByOtherAppLangResponse) |  |
| GetMessageByAppLangMessage | [GetMessageByAppLangMessageRequest](#internationalization-v1-GetMessageByAppLangMessageRequest) | [GetMessageByAppLangMessageResponse](#internationalization-v1-GetMessageByAppLangMessageResponse) |  |
| CreateCountry | [CreateCountryRequest](#internationalization-v1-CreateCountryRequest) | [CreateCountryResponse](#internationalization-v1-CreateCountryResponse) |  |
| GetCountries | [GetCountriesRequest](#internationalization-v1-GetCountriesRequest) | [GetCountriesResponse](#internationalization-v1-GetCountriesResponse) |  |

 



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

