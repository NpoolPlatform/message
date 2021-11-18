# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/coininfo/coininfo.proto](#npool/coininfo/coininfo.proto)
    - [CoinInfoRow](#sphinx.coininfo.v1.CoinInfoRow)
    - [GetCoinInfoRequest](#sphinx.coininfo.v1.GetCoinInfoRequest)
    - [GetCoinInfosRequest](#sphinx.coininfo.v1.GetCoinInfosRequest)
    - [GetCoinInfosResponse](#sphinx.coininfo.v1.GetCoinInfosResponse)
    - [RegisterCoinRequest](#sphinx.coininfo.v1.RegisterCoinRequest)
    - [SetCoinPresaleRequest](#sphinx.coininfo.v1.SetCoinPresaleRequest)
  
    - [SphinxCoininfo](#sphinx.coininfo.v1.SphinxCoininfo)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/coininfo/coininfo.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/coininfo/coininfo.proto



<a name="sphinx.coininfo.v1.CoinInfoRow"></a>

### CoinInfoRow
数据库内CoinInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| CoinTypeID | [int32](#int32) |  | coininfo内部调用，在每一请求中应与CoinType同步 |
| IsPresale | [bool](#bool) |  | 是否为预售，false为在售商品 |
| Name | [string](#string) |  | 币种名称：Filecoin |
| Unit | [string](#string) |  | 单位：FIL |






<a name="sphinx.coininfo.v1.GetCoinInfoRequest"></a>

### GetCoinInfoRequest
获取单个币种


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  |  |
| CoinTypeID | [int32](#int32) |  | coininfo内部调用，在每一请求中应与CoinType同步 |






<a name="sphinx.coininfo.v1.GetCoinInfosRequest"></a>

### GetCoinInfosRequest
获取币种请求






<a name="sphinx.coininfo.v1.GetCoinInfosResponse"></a>

### GetCoinInfosResponse
所有币种信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| infos | [CoinInfoRow](#sphinx.coininfo.v1.CoinInfoRow) | repeated | array |






<a name="sphinx.coininfo.v1.RegisterCoinRequest"></a>

### RegisterCoinRequest
注册币种信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  | sphinxplugin.CoinType |
| Name | [string](#string) |  | 币种名称：Filecoin |
| Unit | [string](#string) |  | 单位：FIL |






<a name="sphinx.coininfo.v1.SetCoinPresaleRequest"></a>

### SetCoinPresaleRequest
设置预售请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CoinType | [sphinx.plugin.v1.CoinType](#sphinx.plugin.v1.CoinType) |  | sphinxplugin.CoinType |
| CoinTypeID | [int32](#int32) |  | coininfo内部调用，在每一请求中应与CoinType同步 |
| IsPresale | [bool](#bool) |  | 是否为预售，false为在售商品 |





 

 

 


<a name="sphinx.coininfo.v1.SphinxCoininfo"></a>

### SphinxCoininfo
CoinInfo服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterCoin | [RegisterCoinRequest](#sphinx.coininfo.v1.RegisterCoinRequest) | [CoinInfoRow](#sphinx.coininfo.v1.CoinInfoRow) | 注册新币种 |
| GetCoinInfos | [GetCoinInfosRequest](#sphinx.coininfo.v1.GetCoinInfosRequest) | [GetCoinInfosResponse](#sphinx.coininfo.v1.GetCoinInfosResponse) | 获取币种信息 |
| GetCoinInfo | [GetCoinInfoRequest](#sphinx.coininfo.v1.GetCoinInfoRequest) | [CoinInfoRow](#sphinx.coininfo.v1.CoinInfoRow) | 获取单个币种 |
| SetCoinPresale | [SetCoinPresaleRequest](#sphinx.coininfo.v1.SetCoinPresaleRequest) | [CoinInfoRow](#sphinx.coininfo.v1.CoinInfoRow) | 设置币种是否预售 |

 



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

