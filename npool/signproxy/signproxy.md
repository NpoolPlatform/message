# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/signproxy/signproxy.proto](#npool/signproxy/signproxy.proto)
    - [ProxyPluginRequest](#sphinx.proxy.v1.ProxyPluginRequest)
    - [ProxyPluginResponse](#sphinx.proxy.v1.ProxyPluginResponse)
    - [ProxySignRequest](#sphinx.proxy.v1.ProxySignRequest)
    - [ProxySignResponse](#sphinx.proxy.v1.ProxySignResponse)
    - [ProxySignResponseInfo](#sphinx.proxy.v1.ProxySignResponseInfo)
  
    - [TransactionType](#sphinx.proxy.v1.TransactionType)
  
    - [SignProxy](#sphinx.proxy.v1.SignProxy)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/signproxy/signproxy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/signproxy/signproxy.proto



<a name="sphinx.proxy.v1.ProxyPluginRequest"></a>

### ProxyPluginRequest
MpoolGetNonce WalletBalance MpoolPush ..


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| TransactionType | [TransactionType](#sphinx.proxy.v1.TransactionType) |  |  |
| TransactionIDInsite | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |
| Signature | [sphinx.plugin.v1.Signature](#sphinx.plugin.v1.Signature) |  |  |






<a name="sphinx.proxy.v1.ProxyPluginResponse"></a>

### ProxyPluginResponse
RegisterCoin ..


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| TransactionType | [TransactionType](#sphinx.proxy.v1.TransactionType) |  |  |
| TransactionIDInsite | [string](#string) |  |  |
| Nonce | [uint64](#uint64) |  |  |
| CID | [string](#string) |  |  |
| Balance | [string](#string) |  |  |






<a name="sphinx.proxy.v1.ProxySignRequest"></a>

### ProxySignRequest
Sign WalletNew ..


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| TransactionType | [TransactionType](#sphinx.proxy.v1.TransactionType) |  |  |
| TransactionIDInsite | [string](#string) |  |  |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |






<a name="sphinx.proxy.v1.ProxySignResponse"></a>

### ProxySignResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| TransactionType | [TransactionType](#sphinx.proxy.v1.TransactionType) |  |  |
| TransactionIDInsite | [string](#string) |  |  |
| Info | [ProxySignResponseInfo](#sphinx.proxy.v1.ProxySignResponseInfo) |  |  |






<a name="sphinx.proxy.v1.ProxySignResponseInfo"></a>

### ProxySignResponseInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [string](#string) |  | create new account address |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |
| Signature | [sphinx.plugin.v1.Signature](#sphinx.plugin.v1.Signature) |  |  |





 


<a name="sphinx.proxy.v1.TransactionType"></a>

### TransactionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| Invalid | 0 |  |
| WalletNew | 1 | proxy -&gt; sign |
| Signature | 2 | proxy -&gt; sign |
| Balance | 3 | proxy -&gt; plugin |
| PreSign | 4 | proxy -&gt; pluign get nonce |
| Broadcast | 5 | proxy -&gt; plugin mpool push |
| RegisterCoin | 6 | plugin -&gt; proxy |


 

 


<a name="sphinx.proxy.v1.SignProxy"></a>

### SignProxy
TODO 分开 sign 和 proxy 的队列
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ProxyPlugin | [ProxyPluginResponse](#sphinx.proxy.v1.ProxyPluginResponse) stream | [ProxyPluginRequest](#sphinx.proxy.v1.ProxyPluginRequest) stream |  |
| ProxySign | [ProxySignResponse](#sphinx.proxy.v1.ProxySignResponse) stream | [ProxySignRequest](#sphinx.proxy.v1.ProxySignRequest) stream |  |

 



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

