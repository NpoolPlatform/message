# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/gasfeeder/gasfeeder.proto](#npool_gasfeeder_gasfeeder-proto)
    - [CoinGas](#gas-feeder-v1-CoinGas)
    - [CountCoinGasesRequest](#gas-feeder-v1-CountCoinGasesRequest)
    - [CountCoinGasesRequest.CondsEntry](#gas-feeder-v1-CountCoinGasesRequest-CondsEntry)
    - [CountCoinGasesResponse](#gas-feeder-v1-CountCoinGasesResponse)
    - [CreateCoinGasRequest](#gas-feeder-v1-CreateCoinGasRequest)
    - [CreateCoinGasResponse](#gas-feeder-v1-CreateCoinGasResponse)
    - [DeleteCoinGasRequest](#gas-feeder-v1-DeleteCoinGasRequest)
    - [DeleteCoinGasResponse](#gas-feeder-v1-DeleteCoinGasResponse)
    - [ExistCoinGasCondsRequest](#gas-feeder-v1-ExistCoinGasCondsRequest)
    - [ExistCoinGasCondsRequest.CondsEntry](#gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry)
    - [ExistCoinGasCondsResponse](#gas-feeder-v1-ExistCoinGasCondsResponse)
    - [ExistCoinGasRequest](#gas-feeder-v1-ExistCoinGasRequest)
    - [ExistCoinGasResponse](#gas-feeder-v1-ExistCoinGasResponse)
    - [GetCoinGasOnlyRequest](#gas-feeder-v1-GetCoinGasOnlyRequest)
    - [GetCoinGasOnlyRequest.CondsEntry](#gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry)
    - [GetCoinGasOnlyResponse](#gas-feeder-v1-GetCoinGasOnlyResponse)
    - [GetCoinGasRequest](#gas-feeder-v1-GetCoinGasRequest)
    - [GetCoinGasResponse](#gas-feeder-v1-GetCoinGasResponse)
    - [GetCoinGasesRequest](#gas-feeder-v1-GetCoinGasesRequest)
    - [GetCoinGasesRequest.CondsEntry](#gas-feeder-v1-GetCoinGasesRequest-CondsEntry)
    - [GetCoinGasesResponse](#gas-feeder-v1-GetCoinGasesResponse)
    - [UpdateCoinGasRequest](#gas-feeder-v1-UpdateCoinGasRequest)
    - [UpdateCoinGasResponse](#gas-feeder-v1-UpdateCoinGasResponse)
  
    - [GasFeeder](#gas-feeder-v1-GasFeeder)
  
- [npool/gasfeeder/gasfeeder.proto](#npool_gasfeeder_gasfeeder-proto)
    - [CoinGas](#gas-feeder-v1-CoinGas)
    - [CountCoinGasesRequest](#gas-feeder-v1-CountCoinGasesRequest)
    - [CountCoinGasesRequest.CondsEntry](#gas-feeder-v1-CountCoinGasesRequest-CondsEntry)
    - [CountCoinGasesResponse](#gas-feeder-v1-CountCoinGasesResponse)
    - [CreateCoinGasRequest](#gas-feeder-v1-CreateCoinGasRequest)
    - [CreateCoinGasResponse](#gas-feeder-v1-CreateCoinGasResponse)
    - [DeleteCoinGasRequest](#gas-feeder-v1-DeleteCoinGasRequest)
    - [DeleteCoinGasResponse](#gas-feeder-v1-DeleteCoinGasResponse)
    - [ExistCoinGasCondsRequest](#gas-feeder-v1-ExistCoinGasCondsRequest)
    - [ExistCoinGasCondsRequest.CondsEntry](#gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry)
    - [ExistCoinGasCondsResponse](#gas-feeder-v1-ExistCoinGasCondsResponse)
    - [ExistCoinGasRequest](#gas-feeder-v1-ExistCoinGasRequest)
    - [ExistCoinGasResponse](#gas-feeder-v1-ExistCoinGasResponse)
    - [GetCoinGasOnlyRequest](#gas-feeder-v1-GetCoinGasOnlyRequest)
    - [GetCoinGasOnlyRequest.CondsEntry](#gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry)
    - [GetCoinGasOnlyResponse](#gas-feeder-v1-GetCoinGasOnlyResponse)
    - [GetCoinGasRequest](#gas-feeder-v1-GetCoinGasRequest)
    - [GetCoinGasResponse](#gas-feeder-v1-GetCoinGasResponse)
    - [GetCoinGasesRequest](#gas-feeder-v1-GetCoinGasesRequest)
    - [GetCoinGasesRequest.CondsEntry](#gas-feeder-v1-GetCoinGasesRequest-CondsEntry)
    - [GetCoinGasesResponse](#gas-feeder-v1-GetCoinGasesResponse)
    - [UpdateCoinGasRequest](#gas-feeder-v1-UpdateCoinGasRequest)
    - [UpdateCoinGasResponse](#gas-feeder-v1-UpdateCoinGasResponse)
  
    - [GasFeeder](#gas-feeder-v1-GasFeeder)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_gasfeeder_gasfeeder-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/gasfeeder/gasfeeder.proto



<a name="gas-feeder-v1-CoinGas"></a>

### CoinGas



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| DailyCoinGas | [double](#double) |  |  |






<a name="gas-feeder-v1-CountCoinGasesRequest"></a>

### CountCoinGasesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountCoinGasesRequest.CondsEntry](#gas-feeder-v1-CountCoinGasesRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-CountCoinGasesRequest-CondsEntry"></a>

### CountCoinGasesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-CountCoinGasesResponse"></a>

### CountCoinGasesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="gas-feeder-v1-CreateCoinGasRequest"></a>

### CreateCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-CreateCoinGasResponse"></a>

### CreateCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-DeleteCoinGasRequest"></a>

### DeleteCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-DeleteCoinGasResponse"></a>

### DeleteCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-ExistCoinGasCondsRequest"></a>

### ExistCoinGasCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistCoinGasCondsRequest.CondsEntry](#gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry"></a>

### ExistCoinGasCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-ExistCoinGasCondsResponse"></a>

### ExistCoinGasCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="gas-feeder-v1-ExistCoinGasRequest"></a>

### ExistCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-ExistCoinGasResponse"></a>

### ExistCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="gas-feeder-v1-GetCoinGasOnlyRequest"></a>

### GetCoinGasOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCoinGasOnlyRequest.CondsEntry](#gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry"></a>

### GetCoinGasOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-GetCoinGasOnlyResponse"></a>

### GetCoinGasOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-GetCoinGasRequest"></a>

### GetCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-GetCoinGasResponse"></a>

### GetCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-GetCoinGasesRequest"></a>

### GetCoinGasesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCoinGasesRequest.CondsEntry](#gas-feeder-v1-GetCoinGasesRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="gas-feeder-v1-GetCoinGasesRequest-CondsEntry"></a>

### GetCoinGasesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-GetCoinGasesResponse"></a>

### GetCoinGasesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinGas](#gas-feeder-v1-CoinGas) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="gas-feeder-v1-UpdateCoinGasRequest"></a>

### UpdateCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-UpdateCoinGasResponse"></a>

### UpdateCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |





 

 

 


<a name="gas-feeder-v1-GasFeeder"></a>

### GasFeeder


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateCoinGas | [CreateCoinGasRequest](#gas-feeder-v1-CreateCoinGasRequest) | [CreateCoinGasResponse](#gas-feeder-v1-CreateCoinGasResponse) |  |
| UpdateCoinGas | [UpdateCoinGasRequest](#gas-feeder-v1-UpdateCoinGasRequest) | [UpdateCoinGasResponse](#gas-feeder-v1-UpdateCoinGasResponse) |  |
| GetCoinGas | [GetCoinGasRequest](#gas-feeder-v1-GetCoinGasRequest) | [GetCoinGasResponse](#gas-feeder-v1-GetCoinGasResponse) |  |
| GetCoinGasOnly | [GetCoinGasOnlyRequest](#gas-feeder-v1-GetCoinGasOnlyRequest) | [GetCoinGasOnlyResponse](#gas-feeder-v1-GetCoinGasOnlyResponse) |  |
| GetCoinGases | [GetCoinGasesRequest](#gas-feeder-v1-GetCoinGasesRequest) | [GetCoinGasesResponse](#gas-feeder-v1-GetCoinGasesResponse) |  |
| ExistCoinGas | [ExistCoinGasRequest](#gas-feeder-v1-ExistCoinGasRequest) | [ExistCoinGasResponse](#gas-feeder-v1-ExistCoinGasResponse) |  |
| ExistCoinGasConds | [ExistCoinGasCondsRequest](#gas-feeder-v1-ExistCoinGasCondsRequest) | [ExistCoinGasCondsResponse](#gas-feeder-v1-ExistCoinGasCondsResponse) |  |
| DeleteCoinGas | [DeleteCoinGasRequest](#gas-feeder-v1-DeleteCoinGasRequest) | [DeleteCoinGasResponse](#gas-feeder-v1-DeleteCoinGasResponse) |  |

 



<a name="npool_gasfeeder_gasfeeder-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/gasfeeder/gasfeeder.proto



<a name="gas-feeder-v1-CoinGas"></a>

### CoinGas



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| CoinTypeID | [string](#string) |  |  |
| DailyCoinGas | [double](#double) |  |  |






<a name="gas-feeder-v1-CountCoinGasesRequest"></a>

### CountCoinGasesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountCoinGasesRequest.CondsEntry](#gas-feeder-v1-CountCoinGasesRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-CountCoinGasesRequest-CondsEntry"></a>

### CountCoinGasesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-CountCoinGasesResponse"></a>

### CountCoinGasesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="gas-feeder-v1-CreateCoinGasRequest"></a>

### CreateCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-CreateCoinGasResponse"></a>

### CreateCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-DeleteCoinGasRequest"></a>

### DeleteCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-DeleteCoinGasResponse"></a>

### DeleteCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-ExistCoinGasCondsRequest"></a>

### ExistCoinGasCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistCoinGasCondsRequest.CondsEntry](#gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-ExistCoinGasCondsRequest-CondsEntry"></a>

### ExistCoinGasCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-ExistCoinGasCondsResponse"></a>

### ExistCoinGasCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="gas-feeder-v1-ExistCoinGasRequest"></a>

### ExistCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-ExistCoinGasResponse"></a>

### ExistCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="gas-feeder-v1-GetCoinGasOnlyRequest"></a>

### GetCoinGasOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCoinGasOnlyRequest.CondsEntry](#gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry) | repeated |  |






<a name="gas-feeder-v1-GetCoinGasOnlyRequest-CondsEntry"></a>

### GetCoinGasOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-GetCoinGasOnlyResponse"></a>

### GetCoinGasOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-GetCoinGasRequest"></a>

### GetCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="gas-feeder-v1-GetCoinGasResponse"></a>

### GetCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-GetCoinGasesRequest"></a>

### GetCoinGasesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetCoinGasesRequest.CondsEntry](#gas-feeder-v1-GetCoinGasesRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="gas-feeder-v1-GetCoinGasesRequest-CondsEntry"></a>

### GetCoinGasesRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="gas-feeder-v1-GetCoinGasesResponse"></a>

### GetCoinGasesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [CoinGas](#gas-feeder-v1-CoinGas) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="gas-feeder-v1-UpdateCoinGasRequest"></a>

### UpdateCoinGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |






<a name="gas-feeder-v1-UpdateCoinGasResponse"></a>

### UpdateCoinGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CoinGas](#gas-feeder-v1-CoinGas) |  |  |





 

 

 


<a name="gas-feeder-v1-GasFeeder"></a>

### GasFeeder


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateCoinGas | [CreateCoinGasRequest](#gas-feeder-v1-CreateCoinGasRequest) | [CreateCoinGasResponse](#gas-feeder-v1-CreateCoinGasResponse) |  |
| UpdateCoinGas | [UpdateCoinGasRequest](#gas-feeder-v1-UpdateCoinGasRequest) | [UpdateCoinGasResponse](#gas-feeder-v1-UpdateCoinGasResponse) |  |
| GetCoinGas | [GetCoinGasRequest](#gas-feeder-v1-GetCoinGasRequest) | [GetCoinGasResponse](#gas-feeder-v1-GetCoinGasResponse) |  |
| GetCoinGasOnly | [GetCoinGasOnlyRequest](#gas-feeder-v1-GetCoinGasOnlyRequest) | [GetCoinGasOnlyResponse](#gas-feeder-v1-GetCoinGasOnlyResponse) |  |
| GetCoinGases | [GetCoinGasesRequest](#gas-feeder-v1-GetCoinGasesRequest) | [GetCoinGasesResponse](#gas-feeder-v1-GetCoinGasesResponse) |  |
| ExistCoinGas | [ExistCoinGasRequest](#gas-feeder-v1-ExistCoinGasRequest) | [ExistCoinGasResponse](#gas-feeder-v1-ExistCoinGasResponse) |  |
| ExistCoinGasConds | [ExistCoinGasCondsRequest](#gas-feeder-v1-ExistCoinGasCondsRequest) | [ExistCoinGasCondsResponse](#gas-feeder-v1-ExistCoinGasCondsResponse) |  |
| DeleteCoinGas | [DeleteCoinGasRequest](#gas-feeder-v1-DeleteCoinGasRequest) | [DeleteCoinGasResponse](#gas-feeder-v1-DeleteCoinGasResponse) |  |

 



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

