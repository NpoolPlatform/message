# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/trading/trading.proto](#npool/trading/trading.proto)
    - [ACKRequest](#sphinx.v1.ACKRequest)
    - [ACKResponse](#sphinx.v1.ACKResponse)
    - [BaseTx](#sphinx.v1.BaseTx)
    - [CreateTransactionRequest](#sphinx.v1.CreateTransactionRequest)
    - [CreateTransactionResponse](#sphinx.v1.CreateTransactionResponse)
    - [CreateWalletRequest](#sphinx.v1.CreateWalletRequest)
    - [CreateWalletResponse](#sphinx.v1.CreateWalletResponse)
    - [EntAccount](#sphinx.v1.EntAccount)
    - [GetTransactionRequest](#sphinx.v1.GetTransactionRequest)
    - [GetTransactionResponse](#sphinx.v1.GetTransactionResponse)
    - [GetWalletBalanceRequest](#sphinx.v1.GetWalletBalanceRequest)
    - [GetWalletBalanceResponse](#sphinx.v1.GetWalletBalanceResponse)
    - [VersionResponse](#sphinx.v1.VersionResponse)
  
    - [Trading](#sphinx.v1.Trading)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/trading/trading.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/trading/trading.proto



<a name="sphinx.v1.ACKRequest"></a>

### ACKRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_type | [sphinx.proxy.v1.TransactionType](#sphinx.proxy.v1.TransactionType) |  | 交易类型： WalletNew, Signature, Balance, PreSign, Broadcast ； 创建账号、查询余额、交易状态更新 |
| coin_type_id | [int32](#int32) |  |  |
| transaction_id_insite | [string](#string) |  | 交易ID |
| transaction_id_chain | [string](#string) |  | 交易链上ID （如有） |
| address | [string](#string) |  | 创建账号的地址 |
| balance | [double](#double) |  | 查询账号的余额 |
| is_okay | [bool](#bool) |  | 交易请求是否成功 |
| error_message | [string](#string) |  | 报错或调试时返回信息 |






<a name="sphinx.v1.ACKResponse"></a>

### ACKResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_okay | [bool](#bool) |  | 是否收到请求 |






<a name="sphinx.v1.BaseTx"></a>

### BaseTx
share transaction structure


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionIDInsite | [string](#string) |  | 站内交易ID |
| CoinName | [string](#string) |  | 币种名称(唯一) |
| AmountFloat64 | [double](#double) |  | 不入库的参考金额 |
| AddressFrom | [string](#string) |  | 发送方 |
| AddressTo | [string](#string) |  | 接收方 |
| InsiteTxType | [string](#string) |  | sql enum: recharge, payment, withdraw, unknown |
| CreatetimeUTC | [int64](#int64) |  | 用户提交请求时的时间戳，与2FA绑定 |






<a name="sphinx.v1.CreateTransactionRequest"></a>

### CreateTransactionRequest
CreateTransaction 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BaseTx](#sphinx.v1.BaseTx) |  | 交易基本信息 |
| UUIDSignature | [string](#string) |  | 2FA的时效性验证码，前期可以留空 |






<a name="sphinx.v1.CreateTransactionResponse"></a>

### CreateTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BaseTx](#sphinx.v1.BaseTx) |  |  |






<a name="sphinx.v1.CreateWalletRequest"></a>

### CreateWalletRequest
CreateWallet 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinName | [string](#string) |  | 币种名称(唯一) |
| UUID | [string](#string) |  | 当前取UserID |






<a name="sphinx.v1.CreateWalletResponse"></a>

### CreateWalletResponse
CreateWallet 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EntAccount](#sphinx.v1.EntAccount) |  |  |






<a name="sphinx.v1.EntAccount"></a>

### EntAccount
share account structure


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinName | [string](#string) |  | 币种名称(唯一) |
| Address | [string](#string) |  | 钱包地址，创建时填空 |






<a name="sphinx.v1.GetTransactionRequest"></a>

### GetTransactionRequest
GetTransaction 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TransactionIDInsite | [string](#string) |  | 站内交易ID |






<a name="sphinx.v1.GetTransactionResponse"></a>

### GetTransactionResponse
GetTransaction 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [BaseTx](#sphinx.v1.BaseTx) |  |  |
| UpdatetimeUTC | [int64](#int64) |  | 上次更新时间 |
| Succeeded | [bool](#bool) |  | bool, 便于调用方判断 |
| Failed | [bool](#bool) |  | 不success不fail就是pending了 |
| TransactionIDChain | [string](#string) |  | 公链交易ID（如有） |
| Status | [string](#string) |  | 为done则成功；全部状态：&#34;pending_review&#34;, &#34;pending_process&#34;, &#34;pending_signinfo&#34;, &#34;pending_sign&#34;, &#34;pending_broadcast&#34;, &#34;pending_confirm&#34;, &#34;done&#34;, &#34;rejected&#34;, &#34;error&#34;, &#34;error_expected&#34; |






<a name="sphinx.v1.GetWalletBalanceRequest"></a>

### GetWalletBalanceRequest
GetWalletBalance 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EntAccount](#sphinx.v1.EntAccount) |  |  |






<a name="sphinx.v1.GetWalletBalanceResponse"></a>

### GetWalletBalanceResponse
GetWalletBalance 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [EntAccount](#sphinx.v1.EntAccount) |  |  |
| AmountFloat64 | [double](#double) |  | 长整型金额 |






<a name="sphinx.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="sphinx.v1.Trading"></a>

### Trading
交易服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateWallet | [CreateWalletRequest](#sphinx.v1.CreateWalletRequest) | [CreateWalletResponse](#sphinx.v1.CreateWalletResponse) | 创建账户 |
| GetWalletBalance | [GetWalletBalanceRequest](#sphinx.v1.GetWalletBalanceRequest) | [GetWalletBalanceResponse](#sphinx.v1.GetWalletBalanceResponse) | 余额查询 |
| CreateTransaction | [CreateTransactionRequest](#sphinx.v1.CreateTransactionRequest) | [CreateTransactionResponse](#sphinx.v1.CreateTransactionResponse) | 转账 / 提现 |
| GetTransaction | [GetTransactionRequest](#sphinx.v1.GetTransactionRequest) | [GetTransactionResponse](#sphinx.v1.GetTransactionResponse) | 交易状态查询 |
| ACK | [ACKRequest](#sphinx.v1.ACKRequest) | [ACKResponse](#sphinx.v1.ACKResponse) | 接收ack |
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#sphinx.v1.VersionResponse) | ping - pong |

 



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

