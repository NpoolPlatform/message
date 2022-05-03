# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/subscribemgr/subscribemgr.proto](#npool_subscribemgr_subscribemgr-proto)
    - [CreateEmailSubscriberRequest](#subscribe-manager-v1-CreateEmailSubscriberRequest)
    - [CreateEmailSubscriberResponse](#subscribe-manager-v1-CreateEmailSubscriberResponse)
    - [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber)
    - [GetAppEmailSubscribersRequest](#subscribe-manager-v1-GetAppEmailSubscribersRequest)
    - [GetAppEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry)
    - [GetAppEmailSubscribersResponse](#subscribe-manager-v1-GetAppEmailSubscribersResponse)
    - [GetEmailSubscribersRequest](#subscribe-manager-v1-GetEmailSubscribersRequest)
    - [GetEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry)
    - [GetEmailSubscribersResponse](#subscribe-manager-v1-GetEmailSubscribersResponse)
  
    - [SubscribeManager](#subscribe-manager-v1-SubscribeManager)
  
- [npool/subscribemgr/subscribemgr.proto](#npool_subscribemgr_subscribemgr-proto)
    - [CreateEmailSubscriberRequest](#subscribe-manager-v1-CreateEmailSubscriberRequest)
    - [CreateEmailSubscriberResponse](#subscribe-manager-v1-CreateEmailSubscriberResponse)
    - [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber)
    - [GetAppEmailSubscribersRequest](#subscribe-manager-v1-GetAppEmailSubscribersRequest)
    - [GetAppEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry)
    - [GetAppEmailSubscribersResponse](#subscribe-manager-v1-GetAppEmailSubscribersResponse)
    - [GetEmailSubscribersRequest](#subscribe-manager-v1-GetEmailSubscribersRequest)
    - [GetEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry)
    - [GetEmailSubscribersResponse](#subscribe-manager-v1-GetEmailSubscribersResponse)
  
    - [SubscribeManager](#subscribe-manager-v1-SubscribeManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_subscribemgr_subscribemgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/subscribemgr/subscribemgr.proto



<a name="subscribe-manager-v1-CreateEmailSubscriberRequest"></a>

### CreateEmailSubscriberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) |  |  |






<a name="subscribe-manager-v1-CreateEmailSubscriberResponse"></a>

### CreateEmailSubscriberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) |  |  |






<a name="subscribe-manager-v1-EmailSubscriber"></a>

### EmailSubscriber



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersRequest"></a>

### GetAppEmailSubscribersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry) | repeated |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry"></a>

### GetAppEmailSubscribersRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersResponse"></a>

### GetAppEmailSubscribersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) | repeated |  |






<a name="subscribe-manager-v1-GetEmailSubscribersRequest"></a>

### GetEmailSubscribersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry) | repeated |  |






<a name="subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry"></a>

### GetEmailSubscribersRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="subscribe-manager-v1-GetEmailSubscribersResponse"></a>

### GetEmailSubscribersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) | repeated |  |





 

 

 


<a name="subscribe-manager-v1-SubscribeManager"></a>

### SubscribeManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateEmailSubscriber | [CreateEmailSubscriberRequest](#subscribe-manager-v1-CreateEmailSubscriberRequest) | [CreateEmailSubscriberResponse](#subscribe-manager-v1-CreateEmailSubscriberResponse) |  |
| GetEmailSubscribers | [GetEmailSubscribersRequest](#subscribe-manager-v1-GetEmailSubscribersRequest) | [GetEmailSubscribersResponse](#subscribe-manager-v1-GetEmailSubscribersResponse) |  |
| GetAppEmailSubscribers | [GetAppEmailSubscribersRequest](#subscribe-manager-v1-GetAppEmailSubscribersRequest) | [GetAppEmailSubscribersResponse](#subscribe-manager-v1-GetAppEmailSubscribersResponse) |  |

 



<a name="npool_subscribemgr_subscribemgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/subscribemgr/subscribemgr.proto



<a name="subscribe-manager-v1-CreateEmailSubscriberRequest"></a>

### CreateEmailSubscriberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) |  |  |






<a name="subscribe-manager-v1-CreateEmailSubscriberResponse"></a>

### CreateEmailSubscriberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) |  |  |






<a name="subscribe-manager-v1-EmailSubscriber"></a>

### EmailSubscriber



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| EmailAddress | [string](#string) |  |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersRequest"></a>

### GetAppEmailSubscribersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Conds | [GetAppEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry) | repeated |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersRequest-CondsEntry"></a>

### GetAppEmailSubscribersRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="subscribe-manager-v1-GetAppEmailSubscribersResponse"></a>

### GetAppEmailSubscribersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) | repeated |  |






<a name="subscribe-manager-v1-GetEmailSubscribersRequest"></a>

### GetEmailSubscribersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Conds | [GetEmailSubscribersRequest.CondsEntry](#subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry) | repeated |  |






<a name="subscribe-manager-v1-GetEmailSubscribersRequest-CondsEntry"></a>

### GetEmailSubscribersRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="subscribe-manager-v1-GetEmailSubscribersResponse"></a>

### GetEmailSubscribersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [EmailSubscriber](#subscribe-manager-v1-EmailSubscriber) | repeated |  |





 

 

 


<a name="subscribe-manager-v1-SubscribeManager"></a>

### SubscribeManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateEmailSubscriber | [CreateEmailSubscriberRequest](#subscribe-manager-v1-CreateEmailSubscriberRequest) | [CreateEmailSubscriberResponse](#subscribe-manager-v1-CreateEmailSubscriberResponse) |  |
| GetEmailSubscribers | [GetEmailSubscribersRequest](#subscribe-manager-v1-GetEmailSubscribersRequest) | [GetEmailSubscribersResponse](#subscribe-manager-v1-GetEmailSubscribersResponse) |  |
| GetAppEmailSubscribers | [GetAppEmailSubscribersRequest](#subscribe-manager-v1-GetAppEmailSubscribersRequest) | [GetAppEmailSubscribersResponse](#subscribe-manager-v1-GetAppEmailSubscribersResponse) |  |

 



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

