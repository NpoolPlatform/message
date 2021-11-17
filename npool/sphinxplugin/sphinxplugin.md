# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/sphinxplugin/sphinxplugin.proto](#npool/sphinxplugin/sphinxplugin.proto)
    - [MpoolGetNonceInfo](#sphinx.plugin.v1.MpoolGetNonceInfo)
    - [MpoolGetNonceRequest](#sphinx.plugin.v1.MpoolGetNonceRequest)
    - [MpoolGetNonceResponse](#sphinx.plugin.v1.MpoolGetNonceResponse)
    - [MpoolPushInfo](#sphinx.plugin.v1.MpoolPushInfo)
    - [MpoolPushRequest](#sphinx.plugin.v1.MpoolPushRequest)
    - [MpoolPushResponse](#sphinx.plugin.v1.MpoolPushResponse)
    - [WalletBalanceInfo](#sphinx.plugin.v1.WalletBalanceInfo)
    - [WalletBalanceRequest](#sphinx.plugin.v1.WalletBalanceRequest)
    - [WalletBalanceResponse](#sphinx.plugin.v1.WalletBalanceResponse)
  
    - [Plugin](#sphinx.plugin.v1.Plugin)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/sphinxplugin/sphinxplugin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/sphinxplugin/sphinxplugin.proto



<a name="sphinx.plugin.v1.MpoolGetNonceInfo"></a>

### MpoolGetNonceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Nonce | [uint64](#uint64) |  |  |






<a name="sphinx.plugin.v1.MpoolGetNonceRequest"></a>

### MpoolGetNonceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.proxy.v1.CoinType](#sphinx.proxy.v1.CoinType) |  |  |
| Address | [string](#string) |  |  |






<a name="sphinx.plugin.v1.MpoolGetNonceResponse"></a>

### MpoolGetNonceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [MpoolGetNonceInfo](#sphinx.plugin.v1.MpoolGetNonceInfo) |  |  |






<a name="sphinx.plugin.v1.MpoolPushInfo"></a>

### MpoolPushInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CID | [string](#string) |  |  |






<a name="sphinx.plugin.v1.MpoolPushRequest"></a>

### MpoolPushRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Message | [sphinx.proxy.v1.UnsignedMessage](#sphinx.proxy.v1.UnsignedMessage) |  |  |
| Signature | [sphinx.proxy.v1.Signature](#sphinx.proxy.v1.Signature) |  |  |






<a name="sphinx.plugin.v1.MpoolPushResponse"></a>

### MpoolPushResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [MpoolPushInfo](#sphinx.plugin.v1.MpoolPushInfo) |  |  |






<a name="sphinx.plugin.v1.WalletBalanceInfo"></a>

### WalletBalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Balance | [string](#string) |  |  |






<a name="sphinx.plugin.v1.WalletBalanceRequest"></a>

### WalletBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.proxy.v1.CoinType](#sphinx.proxy.v1.CoinType) |  |  |
| Address | [string](#string) |  |  |






<a name="sphinx.plugin.v1.WalletBalanceResponse"></a>

### WalletBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WalletBalanceInfo](#sphinx.plugin.v1.WalletBalanceInfo) |  |  |





 

 

 


<a name="sphinx.plugin.v1.Plugin"></a>

### Plugin
钱包代理插件

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| MpoolGetNonce | [MpoolGetNonceRequest](#sphinx.plugin.v1.MpoolGetNonceRequest) | [MpoolGetNonceResponse](#sphinx.plugin.v1.MpoolGetNonceResponse) |  |
| MpoolPush | [MpoolPushRequest](#sphinx.plugin.v1.MpoolPushRequest) | [MpoolPushResponse](#sphinx.plugin.v1.MpoolPushResponse) |  |
| WalletBalance | [WalletBalanceRequest](#sphinx.plugin.v1.WalletBalanceRequest) | [WalletBalanceResponse](#sphinx.plugin.v1.WalletBalanceResponse) |  |

 



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

