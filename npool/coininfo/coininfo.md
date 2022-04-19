# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/coininfo/coininfo.proto](#npool_coininfo_coininfo-proto)
    - [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo)
    - [CoinInfo](#sphinx-coininfo-v1-CoinInfo)
    - [CreateCoinDescriptionRequest](#sphinx-coininfo-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#sphinx-coininfo-v1-CreateCoinDescriptionResponse)
    - [CreateCoinInfoRequest](#sphinx-coininfo-v1-CreateCoinInfoRequest)
    - [CreateCoinInfoResponse](#sphinx-coininfo-v1-CreateCoinInfoResponse)
    - [GetCoinDescriptionRequest](#sphinx-coininfo-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#sphinx-coininfo-v1-GetCoinDescriptionResponse)
    - [GetCoinInfoRequest](#sphinx-coininfo-v1-GetCoinInfoRequest)
    - [GetCoinInfoResponse](#sphinx-coininfo-v1-GetCoinInfoResponse)
    - [GetCoinInfosRequest](#sphinx-coininfo-v1-GetCoinInfosRequest)
    - [GetCoinInfosResponse](#sphinx-coininfo-v1-GetCoinInfosResponse)
    - [UpdateCoinDescriptionRequest](#sphinx-coininfo-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#sphinx-coininfo-v1-UpdateCoinDescriptionResponse)
    - [UpdateCoinInfoRequest](#sphinx-coininfo-v1-UpdateCoinInfoRequest)
    - [UpdateCoinInfoResponse](#sphinx-coininfo-v1-UpdateCoinInfoResponse)
    - [VersionResponse](#sphinx-coininfo-v1-VersionResponse)
  
    - [SphinxCoinInfo](#sphinx-coininfo-v1-SphinxCoinInfo)
  
- [npool/coininfo/coininfo.proto](#npool_coininfo_coininfo-proto)
    - [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo)
    - [CoinInfo](#sphinx-coininfo-v1-CoinInfo)
    - [CreateCoinDescriptionRequest](#sphinx-coininfo-v1-CreateCoinDescriptionRequest)
    - [CreateCoinDescriptionResponse](#sphinx-coininfo-v1-CreateCoinDescriptionResponse)
    - [CreateCoinInfoRequest](#sphinx-coininfo-v1-CreateCoinInfoRequest)
    - [CreateCoinInfoResponse](#sphinx-coininfo-v1-CreateCoinInfoResponse)
    - [GetCoinDescriptionRequest](#sphinx-coininfo-v1-GetCoinDescriptionRequest)
    - [GetCoinDescriptionResponse](#sphinx-coininfo-v1-GetCoinDescriptionResponse)
    - [GetCoinInfoRequest](#sphinx-coininfo-v1-GetCoinInfoRequest)
    - [GetCoinInfoResponse](#sphinx-coininfo-v1-GetCoinInfoResponse)
    - [GetCoinInfosRequest](#sphinx-coininfo-v1-GetCoinInfosRequest)
    - [GetCoinInfosResponse](#sphinx-coininfo-v1-GetCoinInfosResponse)
    - [UpdateCoinDescriptionRequest](#sphinx-coininfo-v1-UpdateCoinDescriptionRequest)
    - [UpdateCoinDescriptionResponse](#sphinx-coininfo-v1-UpdateCoinDescriptionResponse)
    - [UpdateCoinInfoRequest](#sphinx-coininfo-v1-UpdateCoinInfoRequest)
    - [UpdateCoinInfoResponse](#sphinx-coininfo-v1-UpdateCoinInfoResponse)
    - [VersionResponse](#sphinx-coininfo-v1-VersionResponse)
  
    - [SphinxCoinInfo](#sphinx-coininfo-v1-SphinxCoinInfo)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_coininfo_coininfo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/coininfo/coininfo.proto



<a name="sphinx-coininfo-v1-CoinDescriptionInfo"></a>

### CoinDescriptionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| CreatedAt | [uint32](#uint32) |  |  |
| UpdatedAt | [uint32](#uint32) |  |  |






<a name="sphinx-coininfo-v1-CoinInfo"></a>

### CoinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  | uuid |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Unit | [string](#string) |  | 单位：FIL |
| Logo | [string](#string) |  | url; can be empty |
| ENV | [string](#string) |  | main or test |
| ReservedAmount | [double](#double) |  | 预留金额 |
| CreatedAt | [uint32](#uint32) |  |  |
| UpdatedAt | [uint32](#uint32) |  |  |
| ForPay | [bool](#bool) |  | 是否可用作支付货币 |
| HomePage | [string](#string) |  |  |
| Specs | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinDescriptionRequest"></a>

### CreateCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinDescriptionResponse"></a>

### CreateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinInfoRequest"></a>

### CreateCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Unit | [string](#string) |  | 单位：FIL |
| Logo | [string](#string) |  | url; can be empty |
| ENV | [string](#string) |  | main or test |






<a name="sphinx-coininfo-v1-CreateCoinInfoResponse"></a>

### CreateCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Limit | [int32](#int32) |  |  |
| Offset | [int32](#int32) |  |  |






<a name="sphinx-coininfo-v1-GetCoinDescriptionResponse"></a>

### GetCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [int32](#int32) |  |  |
| Infos | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) | repeated |  |






<a name="sphinx-coininfo-v1-GetCoinInfoRequest"></a>

### GetCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfoResponse"></a>

### GetCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfosRequest"></a>

### GetCoinInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfosResponse"></a>

### GetCoinInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [int32](#int32) |  |  |
| Infos | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) | repeated |  |






<a name="sphinx-coininfo-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinDescriptionResponse"></a>

### UpdateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinInfoRequest"></a>

### UpdateCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Logo | [string](#string) |  | url; can be empty |
| ReservedAmount | [double](#double) |  |  |
| ForPay | [bool](#bool) |  |  |
| HomePage | [string](#string) |  |  |
| Specs | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinInfoResponse"></a>

### UpdateCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="sphinx-coininfo-v1-SphinxCoinInfo"></a>

### SphinxCoinInfo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [VersionResponse](#sphinx-coininfo-v1-VersionResponse) |  |
| CreateCoinInfo | [CreateCoinInfoRequest](#sphinx-coininfo-v1-CreateCoinInfoRequest) | [CreateCoinInfoResponse](#sphinx-coininfo-v1-CreateCoinInfoResponse) |  |
| GetCoinInfo | [GetCoinInfoRequest](#sphinx-coininfo-v1-GetCoinInfoRequest) | [GetCoinInfoResponse](#sphinx-coininfo-v1-GetCoinInfoResponse) |  |
| GetCoinInfos | [GetCoinInfosRequest](#sphinx-coininfo-v1-GetCoinInfosRequest) | [GetCoinInfosResponse](#sphinx-coininfo-v1-GetCoinInfosResponse) |  |
| UpdateCoinInfo | [UpdateCoinInfoRequest](#sphinx-coininfo-v1-UpdateCoinInfoRequest) | [UpdateCoinInfoResponse](#sphinx-coininfo-v1-UpdateCoinInfoResponse) |  |
| CreateCoinDescription | [CreateCoinDescriptionRequest](#sphinx-coininfo-v1-CreateCoinDescriptionRequest) | [CreateCoinDescriptionResponse](#sphinx-coininfo-v1-CreateCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#sphinx-coininfo-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#sphinx-coininfo-v1-GetCoinDescriptionResponse) |  |
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#sphinx-coininfo-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#sphinx-coininfo-v1-UpdateCoinDescriptionResponse) |  |

 



<a name="npool_coininfo_coininfo-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/coininfo/coininfo.proto



<a name="sphinx-coininfo-v1-CoinDescriptionInfo"></a>

### CoinDescriptionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| CreatedAt | [uint32](#uint32) |  |  |
| UpdatedAt | [uint32](#uint32) |  |  |






<a name="sphinx-coininfo-v1-CoinInfo"></a>

### CoinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  | uuid |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Unit | [string](#string) |  | 单位：FIL |
| Logo | [string](#string) |  | url; can be empty |
| ENV | [string](#string) |  | main or test |
| ReservedAmount | [double](#double) |  | 预留金额 |
| CreatedAt | [uint32](#uint32) |  |  |
| UpdatedAt | [uint32](#uint32) |  |  |
| ForPay | [bool](#bool) |  | 是否可用作支付货币 |
| HomePage | [string](#string) |  |  |
| Specs | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinDescriptionRequest"></a>

### CreateCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinDescriptionResponse"></a>

### CreateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) |  |  |






<a name="sphinx-coininfo-v1-CreateCoinInfoRequest"></a>

### CreateCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Unit | [string](#string) |  | 单位：FIL |
| Logo | [string](#string) |  | url; can be empty |
| ENV | [string](#string) |  | main or test |






<a name="sphinx-coininfo-v1-CreateCoinInfoResponse"></a>

### CreateCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-GetCoinDescriptionRequest"></a>

### GetCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Limit | [int32](#int32) |  |  |
| Offset | [int32](#int32) |  |  |






<a name="sphinx-coininfo-v1-GetCoinDescriptionResponse"></a>

### GetCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [int32](#int32) |  |  |
| Infos | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) | repeated |  |






<a name="sphinx-coininfo-v1-GetCoinInfoRequest"></a>

### GetCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfoResponse"></a>

### GetCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfosRequest"></a>

### GetCoinInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称 FIL, BTC |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="sphinx-coininfo-v1-GetCoinInfosResponse"></a>

### GetCoinInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [int32](#int32) |  |  |
| Infos | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) | repeated |  |






<a name="sphinx-coininfo-v1-UpdateCoinDescriptionRequest"></a>

### UpdateCoinDescriptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinTypeID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinDescriptionResponse"></a>

### UpdateCoinDescriptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinDescriptionInfo](#sphinx-coininfo-v1-CoinDescriptionInfo) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinInfoRequest"></a>

### UpdateCoinInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| PreSale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Logo | [string](#string) |  | url; can be empty |
| ReservedAmount | [double](#double) |  |  |
| ForPay | [bool](#bool) |  |  |
| HomePage | [string](#string) |  |  |
| Specs | [string](#string) |  |  |






<a name="sphinx-coininfo-v1-UpdateCoinInfoResponse"></a>

### UpdateCoinInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinInfo](#sphinx-coininfo-v1-CoinInfo) |  |  |






<a name="sphinx-coininfo-v1-VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="sphinx-coininfo-v1-SphinxCoinInfo"></a>

### SphinxCoinInfo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [VersionResponse](#sphinx-coininfo-v1-VersionResponse) |  |
| CreateCoinInfo | [CreateCoinInfoRequest](#sphinx-coininfo-v1-CreateCoinInfoRequest) | [CreateCoinInfoResponse](#sphinx-coininfo-v1-CreateCoinInfoResponse) |  |
| GetCoinInfo | [GetCoinInfoRequest](#sphinx-coininfo-v1-GetCoinInfoRequest) | [GetCoinInfoResponse](#sphinx-coininfo-v1-GetCoinInfoResponse) |  |
| GetCoinInfos | [GetCoinInfosRequest](#sphinx-coininfo-v1-GetCoinInfosRequest) | [GetCoinInfosResponse](#sphinx-coininfo-v1-GetCoinInfosResponse) |  |
| UpdateCoinInfo | [UpdateCoinInfoRequest](#sphinx-coininfo-v1-UpdateCoinInfoRequest) | [UpdateCoinInfoResponse](#sphinx-coininfo-v1-UpdateCoinInfoResponse) |  |
| CreateCoinDescription | [CreateCoinDescriptionRequest](#sphinx-coininfo-v1-CreateCoinDescriptionRequest) | [CreateCoinDescriptionResponse](#sphinx-coininfo-v1-CreateCoinDescriptionResponse) |  |
| GetCoinDescription | [GetCoinDescriptionRequest](#sphinx-coininfo-v1-GetCoinDescriptionRequest) | [GetCoinDescriptionResponse](#sphinx-coininfo-v1-GetCoinDescriptionResponse) |  |
| UpdateCoinDescription | [UpdateCoinDescriptionRequest](#sphinx-coininfo-v1-UpdateCoinDescriptionRequest) | [UpdateCoinDescriptionResponse](#sphinx-coininfo-v1-UpdateCoinDescriptionResponse) |  |

 



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

