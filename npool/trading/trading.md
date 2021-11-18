# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/trading/trading.proto](#npool/trading/trading.proto)
    - [ACKRequest](#sphinx.v1.ACKRequest)
    - [ACKResponse](#sphinx.v1.ACKResponse)
    - [AccountAddress](#sphinx.v1.AccountAddress)
    - [AccountBalance](#sphinx.v1.AccountBalance)
    - [AccountTxJSON](#sphinx.v1.AccountTxJSON)
    - [ApplyTransactionRequest](#sphinx.v1.ApplyTransactionRequest)
    - [GetBalanceRequest](#sphinx.v1.GetBalanceRequest)
    - [GetInsiteTxStatusRequest](#sphinx.v1.GetInsiteTxStatusRequest)
    - [GetInsiteTxStatusResponse](#sphinx.v1.GetInsiteTxStatusResponse)
    - [GetTxJSONRequest](#sphinx.v1.GetTxJSONRequest)
    - [RegisterAccountRequest](#sphinx.v1.RegisterAccountRequest)
    - [SuccessInfo](#sphinx.v1.SuccessInfo)
    - [VersionResponse](#sphinx.v1.VersionResponse)
  
    - [ServiceExample](#sphinx.v1.ServiceExample)
    - [Trading](#sphinx.v1.Trading)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/trading/trading.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/trading/trading.proto



<a name="sphinx.v1.ACKRequest"></a>

### ACKRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id_insite | [string](#string) |  | 站内交易ID |
| transaction_chain_id | [string](#string) |  | 信息订单ID |
| is_okay | [bool](#bool) |  | 交易请求是否成功 |






<a name="sphinx.v1.ACKResponse"></a>

### ACKResponse







<a name="sphinx.v1.AccountAddress"></a>

### AccountAddress
RegisterAccount 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 创建的钱包地址 |
| uuid | [string](#string) |  | uuid将用于加密私钥，提高整体安全性 |






<a name="sphinx.v1.AccountBalance"></a>

### AccountBalance
GetBalance 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |
| amount_float64 | [double](#double) |  | 不入库的参考金额 |
| amount_uint64 | [uint64](#uint64) |  | 内部交互标准金额格式 |






<a name="sphinx.v1.AccountTxJSON"></a>

### AccountTxJSON
GetTxJSONRequest 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json | [string](#string) |  |  |






<a name="sphinx.v1.ApplyTransactionRequest"></a>

### ApplyTransactionRequest
ApplyTransaction 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| transaction_id_insite | [string](#string) |  | 站内交易ID |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| amount_float64 | [double](#double) |  | 不入库的参考金额 |
| amount_uint64 | [uint64](#uint64) |  | 内部交互标准金额格式 |
| type | [string](#string) |  | recharge, payment, withdraw |
| uuid_signature | [string](#string) |  | 2FA的时效性验证码，前期可以留空 |
| createtime_utc | [int64](#int64) |  | 用户提交请求时的时间戳，与2FA绑定 |






<a name="sphinx.v1.GetBalanceRequest"></a>

### GetBalanceRequest
GetBalance 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |






<a name="sphinx.v1.GetInsiteTxStatusRequest"></a>

### GetInsiteTxStatusRequest
GetInsiteTxStatus 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id_insite | [string](#string) |  | 站内交易ID |






<a name="sphinx.v1.GetInsiteTxStatusResponse"></a>

### GetInsiteTxStatusResponse
GetInsiteTxStatus 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| amount_float64 | [double](#double) |  | 不入库的参考金额 |
| amount_uint64 | [uint64](#uint64) |  | 内部交互标准金额格式 |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| insite_tx_type | [string](#string) |  | recharge, payment, withdraw, unknown |
| transaction_id_insite | [string](#string) |  | 站内交易ID |
| transaction_id_chain | [string](#string) |  | 公链交易ID（如有） |
| status | [string](#string) |  | 为done则成功；全部状态：&#34;pending_review&#34;, &#34;pending_process&#34;, &#34;pending_signinfo&#34;, &#34;pending_signaction&#34;, &#34;pending_broadcast&#34;, &#34;pending_confirm&#34;, &#34;done&#34;, &#34;rejected&#34;, &#34;error&#34;, &#34;error_expected&#34; |
| is_processing | [bool](#bool) |  | 对应数据库中mutex |
| createtime_utc | [int64](#int64) |  | 创建时间 |
| updatetime_utc | [int64](#int64) |  | 上次更新时间 |
| is_success | [bool](#bool) |  | 便于调用方判断 |
| is_failed | [bool](#bool) |  | 不success不fail就是pending了 |






<a name="sphinx.v1.GetTxJSONRequest"></a>

### GetTxJSONRequest
GetTxJSONRequest 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  | 继承钱包节点基础功能，预留 |
| address | [string](#string) |  | 要查询的钱包地址 |
| timefrom_utc | [int64](#int64) |  | 开始时间 |
| timetill_utc | [int64](#int64) |  | 结束时间 |
| limit_n | [int32](#int32) |  | 服务端限制返回条数 |






<a name="sphinx.v1.RegisterAccountRequest"></a>

### RegisterAccountRequest
RegisterAccount 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| uuid | [string](#string) |  | user_id或与其绑定的唯一标识符 |






<a name="sphinx.v1.SuccessInfo"></a>

### SuccessInfo
deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [string](#string) |  | &#34;success&#34; |






<a name="sphinx.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="sphinx.v1.ServiceExample"></a>

### ServiceExample


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#sphinx.v1.VersionResponse) | Method Version |


<a name="sphinx.v1.Trading"></a>

### Trading
交易服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterAccount | [RegisterAccountRequest](#sphinx.v1.RegisterAccountRequest) | [AccountAddress](#sphinx.v1.AccountAddress) | 创建账户 |
| GetBalance | [GetBalanceRequest](#sphinx.v1.GetBalanceRequest) | [AccountBalance](#sphinx.v1.AccountBalance) | 余额查询 |
| ApplyTransaction | [ApplyTransactionRequest](#sphinx.v1.ApplyTransactionRequest) | [SuccessInfo](#sphinx.v1.SuccessInfo) | 转账 / 提现 |
| GetTxJSON | [GetTxJSONRequest](#sphinx.v1.GetTxJSONRequest) | [AccountTxJSON](#sphinx.v1.AccountTxJSON) | TODO: 账户交易查询 |
| GetInsiteTxStatus | [GetInsiteTxStatusRequest](#sphinx.v1.GetInsiteTxStatusRequest) | [GetInsiteTxStatusResponse](#sphinx.v1.GetInsiteTxStatusResponse) | 交易状态查询 |
| ACK | [ACKRequest](#sphinx.v1.ACKRequest) | [ACKResponse](#sphinx.v1.ACKResponse) | 接收ack |

 



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

