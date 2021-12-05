# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/sphinxservice/sphinxservice.proto](#npool/sphinxservice/sphinxservice.proto)
    - [CreateTransactionRequest](#sphinx.v1.CreateTransactionRequest)
    - [CreateTransactionResponse](#sphinx.v1.CreateTransactionResponse)
    - [GetTransactionRequest](#sphinx.v1.GetTransactionRequest)
    - [GetTransactionResponse](#sphinx.v1.GetTransactionResponse)
    - [GetTransactionsRequest](#sphinx.v1.GetTransactionsRequest)
    - [GetTransactionsResponse](#sphinx.v1.GetTransactionsResponse)
    - [TransactionInfo](#sphinx.v1.TransactionInfo)
    - [UpdateTransactionRequest](#sphinx.v1.UpdateTransactionRequest)
    - [UpdateTransactionResponse](#sphinx.v1.UpdateTransactionResponse)
    - [VersionResponse](#sphinx.v1.VersionResponse)
  
    - [TransactionState](#sphinx.v1.TransactionState)
  
    - [SphinxService](#sphinx.v1.SphinxService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/sphinxservice/sphinxservice.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/sphinxservice/sphinxservice.proto



<a name="sphinx.v1.CreateTransactionRequest"></a>

### CreateTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| From | [string](#string) |  |  |
| To | [string](#string) |  |  |






<a name="sphinx.v1.CreateTransactionResponse"></a>

### CreateTransactionResponse







<a name="sphinx.v1.GetTransactionRequest"></a>

### GetTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionID | [string](#string) |  |  |






<a name="sphinx.v1.GetTransactionResponse"></a>

### GetTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TransactionInfo](#sphinx.v1.TransactionInfo) |  |  |






<a name="sphinx.v1.GetTransactionsRequest"></a>

### GetTransactionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| State | [string](#string) |  |  |
| From | [string](#string) |  |  |
| To | [string](#string) |  |  |
| Offset | [uint32](#uint32) |  |  |
| Limit | [uint32](#uint32) |  |  |






<a name="sphinx.v1.GetTransactionsResponse"></a>

### GetTransactionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [int32](#int32) |  |  |
| Infos | [TransactionInfo](#sphinx.v1.TransactionInfo) | repeated |  |






<a name="sphinx.v1.TransactionInfo"></a>

### TransactionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Amount | [double](#double) |  |  |
| From | [string](#string) |  |  |
| To | [string](#string) |  |  |
| state | [string](#string) |  |  |
| CreatedAt | [uint32](#uint32) |  |  |
| UpdatedAt | [uint32](#uint32) |  |  |






<a name="sphinx.v1.UpdateTransactionRequest"></a>

### UpdateTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionID | [string](#string) |  |  |
| state | [string](#string) |  |  |






<a name="sphinx.v1.UpdateTransactionResponse"></a>

### UpdateTransactionResponse







<a name="sphinx.v1.VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 


<a name="sphinx.v1.TransactionState"></a>

### TransactionState


| Name | Number | Description |
| ---- | ------ | ----------- |
| TransactionStateUnKnow | 0 |  |
| TransactionStatePendingReview | 1 |  |
| TransactionStateConfirm | 2 |  |
| TransactionStateRejected | 3 |  |
| TransactionStatePendingTransaction | 4 |  |
| TransactionStateDone | 5 |  |


 

 


<a name="sphinx.v1.SphinxService"></a>

### SphinxService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTransaction | [CreateTransactionRequest](#sphinx.v1.CreateTransactionRequest) | [CreateTransactionResponse](#sphinx.v1.CreateTransactionResponse) |  |
| GetTransaction | [GetTransactionRequest](#sphinx.v1.GetTransactionRequest) | [GetTransactionResponse](#sphinx.v1.GetTransactionResponse) |  |
| GetTransactions | [GetTransactionsRequest](#sphinx.v1.GetTransactionsRequest) | [GetTransactionsResponse](#sphinx.v1.GetTransactionsResponse) |  |
| UpdateTransaction | [UpdateTransactionRequest](#sphinx.v1.UpdateTransactionRequest) | [UpdateTransactionResponse](#sphinx.v1.UpdateTransactionResponse) |  |
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#sphinx.v1.VersionResponse) |  |

 



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

