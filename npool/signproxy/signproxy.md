# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/signproxy/signproxy.proto](#npool/signproxy/signproxy.proto)
    - [AccountInfo](#sphinx.proxy.v1.AccountInfo)
    - [RegisterCoinRequest](#sphinx.proxy.v1.RegisterCoinRequest)
    - [RegisterCoinResponse](#sphinx.proxy.v1.RegisterCoinResponse)
    - [TransactionRequest](#sphinx.proxy.v1.TransactionRequest)
    - [TransactionResponse](#sphinx.proxy.v1.TransactionResponse)
    - [TransactionResponseInfo](#sphinx.proxy.v1.TransactionResponseInfo)
    - [WalletNewResponse](#sphinx.proxy.v1.WalletNewResponse)
  
    - [SignProxy](#sphinx.proxy.v1.SignProxy)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/signproxy/signproxy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/signproxy/signproxy.proto



<a name="sphinx.proxy.v1.AccountInfo"></a>

### AccountInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SigType | [string](#string) |  | secp256k1 |
| Address | [string](#string) |  |  |






<a name="sphinx.proxy.v1.RegisterCoinRequest"></a>

### RegisterCoinRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |






<a name="sphinx.proxy.v1.RegisterCoinResponse"></a>

### RegisterCoinResponse







<a name="sphinx.proxy.v1.TransactionRequest"></a>

### TransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |






<a name="sphinx.proxy.v1.TransactionResponse"></a>

### TransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TransactionResponseInfo](#sphinx.proxy.v1.TransactionResponseInfo) |  |  |






<a name="sphinx.proxy.v1.TransactionResponseInfo"></a>

### TransactionResponseInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [string](#string) |  | create new account address |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |
| Signature | [sphinx.plugin.v1.Signature](#sphinx.plugin.v1.Signature) |  |  |






<a name="sphinx.proxy.v1.WalletNewResponse"></a>

### WalletNewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AccountInfo](#sphinx.proxy.v1.AccountInfo) |  |  |





 

 

 


<a name="sphinx.proxy.v1.SignProxy"></a>

### SignProxy
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterCoin | [RegisterCoinRequest](#sphinx.proxy.v1.RegisterCoinRequest) | [RegisterCoinResponse](#sphinx.proxy.v1.RegisterCoinResponse) | RegisterCoin register new coin |
| WalletNew | [.sphinx.sign.v1.WalletNewRequest](#sphinx.sign.v1.WalletNewRequest) | [.sphinx.sign.v1.WalletNewResponse](#sphinx.sign.v1.WalletNewResponse) | WalletNew create new account |
| Transaction | [TransactionResponse](#sphinx.proxy.v1.TransactionResponse) stream | [TransactionRequest](#sphinx.proxy.v1.TransactionRequest) stream | Transaction use transfer |
| WalletBalance | [.sphinx.plugin.v1.WalletBalanceRequest](#sphinx.plugin.v1.WalletBalanceRequest) | [.sphinx.plugin.v1.WalletBalanceResponse](#sphinx.plugin.v1.WalletBalanceResponse) | WalletBalance get account balance |

 



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

