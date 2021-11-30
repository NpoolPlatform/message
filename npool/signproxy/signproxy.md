# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/signproxy/signproxy.proto](#npool/signproxy/signproxy.proto)
    - [BalanceInfo](#sphinx.proxy.v1.BalanceInfo)
    - [CreateTransactionReponse](#sphinx.proxy.v1.CreateTransactionReponse)
    - [CreateTransactionRequest](#sphinx.proxy.v1.CreateTransactionRequest)
    - [CreateWalletReponse](#sphinx.proxy.v1.CreateWalletReponse)
    - [CreateWalletRequest](#sphinx.proxy.v1.CreateWalletRequest)
    - [GetBalanceReponse](#sphinx.proxy.v1.GetBalanceReponse)
    - [GetBalanceRequest](#sphinx.proxy.v1.GetBalanceRequest)
    - [GetTransactionReponse](#sphinx.proxy.v1.GetTransactionReponse)
    - [GetTransactionRequest](#sphinx.proxy.v1.GetTransactionRequest)
    - [ProxyPluginRequest](#sphinx.proxy.v1.ProxyPluginRequest)
    - [ProxyPluginResponse](#sphinx.proxy.v1.ProxyPluginResponse)
    - [ProxySignRequest](#sphinx.proxy.v1.ProxySignRequest)
    - [ProxySignResponse](#sphinx.proxy.v1.ProxySignResponse)
    - [ProxySignResponseInfo](#sphinx.proxy.v1.ProxySignResponseInfo)
    - [WalletInfo](#sphinx.proxy.v1.WalletInfo)
  
    - [TransactionType](#sphinx.proxy.v1.TransactionType)
  
    - [SignProxy](#sphinx.proxy.v1.SignProxy)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/signproxy/signproxy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/signproxy/signproxy.proto



<a name="sphinx.proxy.v1.BalanceInfo"></a>

### BalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Balance | [double](#double) |  |  |






<a name="sphinx.proxy.v1.CreateTransactionReponse"></a>

### CreateTransactionReponse







<a name="sphinx.proxy.v1.CreateTransactionRequest"></a>

### CreateTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| TransactionID | [string](#string) |  |  |
| From | [string](#string) |  |  |
| To | [string](#string) |  |  |
| Value | [double](#double) |  |  |






<a name="sphinx.proxy.v1.CreateWalletReponse"></a>

### CreateWalletReponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [WalletInfo](#sphinx.proxy.v1.WalletInfo) |  |  |






<a name="sphinx.proxy.v1.CreateWalletRequest"></a>

### CreateWalletRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |






<a name="sphinx.proxy.v1.GetBalanceReponse"></a>

### GetBalanceReponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BalanceInfo](#sphinx.proxy.v1.BalanceInfo) |  |  |






<a name="sphinx.proxy.v1.GetBalanceRequest"></a>

### GetBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| Address | [string](#string) |  |  |






<a name="sphinx.proxy.v1.GetTransactionReponse"></a>

### GetTransactionReponse







<a name="sphinx.proxy.v1.GetTransactionRequest"></a>

### GetTransactionRequest







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
| Balance | [double](#double) |  |  |
| Message | [sphinx.plugin.v1.UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage) |  |  |






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






<a name="sphinx.proxy.v1.WalletInfo"></a>

### WalletInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Address | [string](#string) |  | TODO sign type |





 


<a name="sphinx.proxy.v1.TransactionType"></a>

### TransactionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| Invalid | 0 |  |
| WalletNew | 1 | proxy -&gt; sign |
| TransactionNew | 2 |  |
| Signature | 3 | proxy -&gt; sign |
| Balance | 4 | proxy -&gt; plugin |
| PreSign | 5 | proxy -&gt; pluign get nonce |
| Broadcast | 6 | proxy -&gt; plugin mpool push |
| RegisterCoin | 7 | plugin -&gt; proxy |


 

 


<a name="sphinx.proxy.v1.SignProxy"></a>

### SignProxy
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetBalance | [GetBalanceRequest](#sphinx.proxy.v1.GetBalanceRequest) | [GetBalanceReponse](#sphinx.proxy.v1.GetBalanceReponse) | sync |
| CreateWallet | [CreateWalletRequest](#sphinx.proxy.v1.CreateWalletRequest) | [CreateWalletReponse](#sphinx.proxy.v1.CreateWalletReponse) |  |
| CreateTransaction | [CreateTransactionRequest](#sphinx.proxy.v1.CreateTransactionRequest) | [CreateTransactionReponse](#sphinx.proxy.v1.CreateTransactionReponse) |  |
| GetTransaction | [GetTransactionRequest](#sphinx.proxy.v1.GetTransactionRequest) | [GetTransactionReponse](#sphinx.proxy.v1.GetTransactionReponse) |  |
| ProxyPlugin | [ProxyPluginResponse](#sphinx.proxy.v1.ProxyPluginResponse) stream | [ProxyPluginRequest](#sphinx.proxy.v1.ProxyPluginRequest) stream | async stream |
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

