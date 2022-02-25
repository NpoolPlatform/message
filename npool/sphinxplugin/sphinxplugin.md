# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/sphinxplugin/sphinxplugin.proto](#npool/sphinxplugin/sphinxplugin.proto)
    - [MsgTx](#sphinx.plugin.v1.MsgTx)
    - [OutPoint](#sphinx.plugin.v1.OutPoint)
    - [Signature](#sphinx.plugin.v1.Signature)
    - [TxIn](#sphinx.plugin.v1.TxIn)
    - [TxOut](#sphinx.plugin.v1.TxOut)
    - [UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage)
    - [Unspent](#sphinx.plugin.v1.Unspent)
  
    - [CoinType](#sphinx.plugin.v1.CoinType)
  
- [npool/sphinxplugin/sphinxplugin.proto](#npool/sphinxplugin/sphinxplugin.proto)
    - [MsgTx](#sphinx.plugin.v1.MsgTx)
    - [OutPoint](#sphinx.plugin.v1.OutPoint)
    - [Signature](#sphinx.plugin.v1.Signature)
    - [TxIn](#sphinx.plugin.v1.TxIn)
    - [TxOut](#sphinx.plugin.v1.TxOut)
    - [UnsignedMessage](#sphinx.plugin.v1.UnsignedMessage)
    - [Unspent](#sphinx.plugin.v1.Unspent)
  
    - [CoinType](#sphinx.plugin.v1.CoinType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/sphinxplugin/sphinxplugin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/sphinxplugin/sphinxplugin.proto



<a name="sphinx.plugin.v1.MsgTx"></a>

### MsgTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Version | [int32](#int32) |  |  |
| TxIn | [TxIn](#sphinx.plugin.v1.TxIn) | repeated |  |
| TxOut | [TxOut](#sphinx.plugin.v1.TxOut) | repeated |  |
| LockTime | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.OutPoint"></a>

### OutPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Hash | [bytes](#bytes) |  | byte len 32 |
| Index | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.Signature"></a>

### Signature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SignType | [string](#string) |  | secp256k1 |
| Data | [bytes](#bytes) |  |  |






<a name="sphinx.plugin.v1.TxIn"></a>

### TxIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreviousOutPoint | [OutPoint](#sphinx.plugin.v1.OutPoint) |  |  |
| SignatureScript | [bytes](#bytes) |  |  |
| Witness | [bytes](#bytes) | repeated |  |
| Sequence | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.TxOut"></a>

### TxOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Value | [int64](#int64) |  |  |
| PkScript | [bytes](#bytes) |  |  |






<a name="sphinx.plugin.v1.UnsignedMessage"></a>

### UnsignedMessage
fil


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Version | [uint64](#uint64) |  |  |
| To | [string](#string) |  |  |
| From | [string](#string) |  |  |
| Value | [double](#double) |  |  |
| Nonce | [uint64](#uint64) |  | fil |
| GasLimit | [int64](#int64) |  |  |
| GasFeeCap | [uint64](#uint64) |  |  |
| GasPremium | [uint64](#uint64) |  |  |
| Method | [uint64](#uint64) |  |  |
| Params | [bytes](#bytes) |  |  |
| Unspent | [Unspent](#sphinx.plugin.v1.Unspent) | repeated | btc |






<a name="sphinx.plugin.v1.Unspent"></a>

### Unspent
============================= btc


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TxID | [string](#string) |  |  |
| Vout | [uint32](#uint32) |  |  |
| Address | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| ScriptPubKey | [string](#string) |  |  |
| RedeemScript | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Confirmations | [int64](#int64) |  |  |
| Spendable | [bool](#bool) |  |  |





 


<a name="sphinx.plugin.v1.CoinType"></a>

### CoinType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CoinTypeUnKnow | 0 |  |
| CoinTypefilecoin | 1 |  |
| CoinTypebitcoin | 2 |  |
| CoinTypeethereum | 3 |  |
| CoinTypespacemesh | 4 |  |


 

 

 



<a name="npool/sphinxplugin/sphinxplugin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/sphinxplugin/sphinxplugin.proto



<a name="sphinx.plugin.v1.MsgTx"></a>

### MsgTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Version | [int32](#int32) |  |  |
| TxIn | [TxIn](#sphinx.plugin.v1.TxIn) | repeated |  |
| TxOut | [TxOut](#sphinx.plugin.v1.TxOut) | repeated |  |
| LockTime | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.OutPoint"></a>

### OutPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Hash | [bytes](#bytes) |  | byte len 32 |
| Index | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.Signature"></a>

### Signature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SignType | [string](#string) |  | secp256k1 |
| Data | [bytes](#bytes) |  |  |






<a name="sphinx.plugin.v1.TxIn"></a>

### TxIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PreviousOutPoint | [OutPoint](#sphinx.plugin.v1.OutPoint) |  |  |
| SignatureScript | [bytes](#bytes) |  |  |
| Witness | [bytes](#bytes) | repeated |  |
| Sequence | [uint32](#uint32) |  |  |






<a name="sphinx.plugin.v1.TxOut"></a>

### TxOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Value | [int64](#int64) |  |  |
| PkScript | [bytes](#bytes) |  |  |






<a name="sphinx.plugin.v1.UnsignedMessage"></a>

### UnsignedMessage
fil


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Version | [uint64](#uint64) |  |  |
| To | [string](#string) |  |  |
| From | [string](#string) |  |  |
| Value | [double](#double) |  |  |
| Nonce | [uint64](#uint64) |  | fil |
| GasLimit | [int64](#int64) |  |  |
| GasFeeCap | [uint64](#uint64) |  |  |
| GasPremium | [uint64](#uint64) |  |  |
| Method | [uint64](#uint64) |  |  |
| Params | [bytes](#bytes) |  |  |
| Unspent | [Unspent](#sphinx.plugin.v1.Unspent) | repeated | btc |






<a name="sphinx.plugin.v1.Unspent"></a>

### Unspent
============================= btc


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TxID | [string](#string) |  |  |
| Vout | [uint32](#uint32) |  |  |
| Address | [string](#string) |  |  |
| Account | [string](#string) |  |  |
| ScriptPubKey | [string](#string) |  |  |
| RedeemScript | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| Confirmations | [int64](#int64) |  |  |
| Spendable | [bool](#bool) |  |  |





 


<a name="sphinx.plugin.v1.CoinType"></a>

### CoinType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CoinTypeUnKnow | 0 |  |
| CoinTypefilecoin | 1 |  |
| CoinTypebitcoin | 2 |  |
| CoinTypeethereum | 3 |  |
| CoinTypespacemesh | 4 |  |


 

 

 



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

