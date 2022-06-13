# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/stockmgr/stockmgr.proto](#npool_stockmgr_stockmgr-proto)
    - [AddStockFieldsRequest](#stock-manager-v1-AddStockFieldsRequest)
    - [AddStockFieldsRequest.FieldsEntry](#stock-manager-v1-AddStockFieldsRequest-FieldsEntry)
    - [AddStockFieldsResponse](#stock-manager-v1-AddStockFieldsResponse)
    - [CountStocksRequest](#stock-manager-v1-CountStocksRequest)
    - [CountStocksRequest.CondsEntry](#stock-manager-v1-CountStocksRequest-CondsEntry)
    - [CountStocksResponse](#stock-manager-v1-CountStocksResponse)
    - [CreateStockRequest](#stock-manager-v1-CreateStockRequest)
    - [CreateStockResponse](#stock-manager-v1-CreateStockResponse)
    - [CreateStocksRequest](#stock-manager-v1-CreateStocksRequest)
    - [CreateStocksResponse](#stock-manager-v1-CreateStocksResponse)
    - [DeleteStockRequest](#stock-manager-v1-DeleteStockRequest)
    - [DeleteStockResponse](#stock-manager-v1-DeleteStockResponse)
    - [ExistStockCondsRequest](#stock-manager-v1-ExistStockCondsRequest)
    - [ExistStockCondsRequest.CondsEntry](#stock-manager-v1-ExistStockCondsRequest-CondsEntry)
    - [ExistStockCondsResponse](#stock-manager-v1-ExistStockCondsResponse)
    - [ExistStockRequest](#stock-manager-v1-ExistStockRequest)
    - [ExistStockResponse](#stock-manager-v1-ExistStockResponse)
    - [GetStockOnlyRequest](#stock-manager-v1-GetStockOnlyRequest)
    - [GetStockOnlyRequest.CondsEntry](#stock-manager-v1-GetStockOnlyRequest-CondsEntry)
    - [GetStockOnlyResponse](#stock-manager-v1-GetStockOnlyResponse)
    - [GetStockRequest](#stock-manager-v1-GetStockRequest)
    - [GetStockResponse](#stock-manager-v1-GetStockResponse)
    - [GetStocksRequest](#stock-manager-v1-GetStocksRequest)
    - [GetStocksRequest.CondsEntry](#stock-manager-v1-GetStocksRequest-CondsEntry)
    - [GetStocksResponse](#stock-manager-v1-GetStocksResponse)
    - [Stock](#stock-manager-v1-Stock)
    - [UpdateStockFieldsRequest](#stock-manager-v1-UpdateStockFieldsRequest)
    - [UpdateStockFieldsRequest.FieldsEntry](#stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry)
    - [UpdateStockFieldsResponse](#stock-manager-v1-UpdateStockFieldsResponse)
    - [UpdateStockRequest](#stock-manager-v1-UpdateStockRequest)
    - [UpdateStockResponse](#stock-manager-v1-UpdateStockResponse)
  
    - [StockManager](#stock-manager-v1-StockManager)
  
- [npool/stockmgr/stockmgr.proto](#npool_stockmgr_stockmgr-proto)
    - [AddStockFieldsRequest](#stock-manager-v1-AddStockFieldsRequest)
    - [AddStockFieldsRequest.FieldsEntry](#stock-manager-v1-AddStockFieldsRequest-FieldsEntry)
    - [AddStockFieldsResponse](#stock-manager-v1-AddStockFieldsResponse)
    - [CountStocksRequest](#stock-manager-v1-CountStocksRequest)
    - [CountStocksRequest.CondsEntry](#stock-manager-v1-CountStocksRequest-CondsEntry)
    - [CountStocksResponse](#stock-manager-v1-CountStocksResponse)
    - [CreateStockRequest](#stock-manager-v1-CreateStockRequest)
    - [CreateStockResponse](#stock-manager-v1-CreateStockResponse)
    - [CreateStocksRequest](#stock-manager-v1-CreateStocksRequest)
    - [CreateStocksResponse](#stock-manager-v1-CreateStocksResponse)
    - [DeleteStockRequest](#stock-manager-v1-DeleteStockRequest)
    - [DeleteStockResponse](#stock-manager-v1-DeleteStockResponse)
    - [ExistStockCondsRequest](#stock-manager-v1-ExistStockCondsRequest)
    - [ExistStockCondsRequest.CondsEntry](#stock-manager-v1-ExistStockCondsRequest-CondsEntry)
    - [ExistStockCondsResponse](#stock-manager-v1-ExistStockCondsResponse)
    - [ExistStockRequest](#stock-manager-v1-ExistStockRequest)
    - [ExistStockResponse](#stock-manager-v1-ExistStockResponse)
    - [GetStockOnlyRequest](#stock-manager-v1-GetStockOnlyRequest)
    - [GetStockOnlyRequest.CondsEntry](#stock-manager-v1-GetStockOnlyRequest-CondsEntry)
    - [GetStockOnlyResponse](#stock-manager-v1-GetStockOnlyResponse)
    - [GetStockRequest](#stock-manager-v1-GetStockRequest)
    - [GetStockResponse](#stock-manager-v1-GetStockResponse)
    - [GetStocksRequest](#stock-manager-v1-GetStocksRequest)
    - [GetStocksRequest.CondsEntry](#stock-manager-v1-GetStocksRequest-CondsEntry)
    - [GetStocksResponse](#stock-manager-v1-GetStocksResponse)
    - [Stock](#stock-manager-v1-Stock)
    - [UpdateStockFieldsRequest](#stock-manager-v1-UpdateStockFieldsRequest)
    - [UpdateStockFieldsRequest.FieldsEntry](#stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry)
    - [UpdateStockFieldsResponse](#stock-manager-v1-UpdateStockFieldsResponse)
    - [UpdateStockRequest](#stock-manager-v1-UpdateStockRequest)
    - [UpdateStockResponse](#stock-manager-v1-UpdateStockResponse)
  
    - [StockManager](#stock-manager-v1-StockManager)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_stockmgr_stockmgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/stockmgr/stockmgr.proto



<a name="stock-manager-v1-AddStockFieldsRequest"></a>

### AddStockFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [AddStockFieldsRequest.FieldsEntry](#stock-manager-v1-AddStockFieldsRequest-FieldsEntry) | repeated |  |






<a name="stock-manager-v1-AddStockFieldsRequest-FieldsEntry"></a>

### AddStockFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stock-manager-v1-AddStockFieldsResponse"></a>

### AddStockFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CountStocksRequest"></a>

### CountStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountStocksRequest.CondsEntry](#stock-manager-v1-CountStocksRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-CountStocksRequest-CondsEntry"></a>

### CountStocksRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-CountStocksResponse"></a>

### CountStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="stock-manager-v1-CreateStockRequest"></a>

### CreateStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CreateStockResponse"></a>

### CreateStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CreateStocksRequest"></a>

### CreateStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |






<a name="stock-manager-v1-CreateStocksResponse"></a>

### CreateStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |






<a name="stock-manager-v1-DeleteStockRequest"></a>

### DeleteStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-DeleteStockResponse"></a>

### DeleteStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-ExistStockCondsRequest"></a>

### ExistStockCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistStockCondsRequest.CondsEntry](#stock-manager-v1-ExistStockCondsRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-ExistStockCondsRequest-CondsEntry"></a>

### ExistStockCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-ExistStockCondsResponse"></a>

### ExistStockCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="stock-manager-v1-ExistStockRequest"></a>

### ExistStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-ExistStockResponse"></a>

### ExistStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="stock-manager-v1-GetStockOnlyRequest"></a>

### GetStockOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetStockOnlyRequest.CondsEntry](#stock-manager-v1-GetStockOnlyRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-GetStockOnlyRequest-CondsEntry"></a>

### GetStockOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-GetStockOnlyResponse"></a>

### GetStockOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-GetStockRequest"></a>

### GetStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-GetStockResponse"></a>

### GetStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-GetStocksRequest"></a>

### GetStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetStocksRequest.CondsEntry](#stock-manager-v1-GetStocksRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="stock-manager-v1-GetStocksRequest-CondsEntry"></a>

### GetStocksRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-GetStocksResponse"></a>

### GetStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="stock-manager-v1-Stock"></a>

### Stock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Total | [uint32](#uint32) |  |  |
| InService | [uint32](#uint32) |  |  |
| Sold | [uint32](#uint32) |  |  |
| Locked | [uint32](#uint32) |  |  |






<a name="stock-manager-v1-UpdateStockFieldsRequest"></a>

### UpdateStockFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateStockFieldsRequest.FieldsEntry](#stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry) | repeated |  |






<a name="stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry"></a>

### UpdateStockFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stock-manager-v1-UpdateStockFieldsResponse"></a>

### UpdateStockFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-UpdateStockRequest"></a>

### UpdateStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-UpdateStockResponse"></a>

### UpdateStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |





 

 

 


<a name="stock-manager-v1-StockManager"></a>

### StockManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateStock | [CreateStockRequest](#stock-manager-v1-CreateStockRequest) | [CreateStockResponse](#stock-manager-v1-CreateStockResponse) |  |
| CreateStocks | [CreateStocksRequest](#stock-manager-v1-CreateStocksRequest) | [CreateStocksResponse](#stock-manager-v1-CreateStocksResponse) |  |
| UpdateStock | [UpdateStockRequest](#stock-manager-v1-UpdateStockRequest) | [UpdateStockResponse](#stock-manager-v1-UpdateStockResponse) |  |
| UpdateStockFields | [UpdateStockFieldsRequest](#stock-manager-v1-UpdateStockFieldsRequest) | [UpdateStockFieldsResponse](#stock-manager-v1-UpdateStockFieldsResponse) |  |
| AddStockFields | [AddStockFieldsRequest](#stock-manager-v1-AddStockFieldsRequest) | [AddStockFieldsResponse](#stock-manager-v1-AddStockFieldsResponse) |  |
| GetStock | [GetStockRequest](#stock-manager-v1-GetStockRequest) | [GetStockResponse](#stock-manager-v1-GetStockResponse) |  |
| GetStockOnly | [GetStockOnlyRequest](#stock-manager-v1-GetStockOnlyRequest) | [GetStockOnlyResponse](#stock-manager-v1-GetStockOnlyResponse) |  |
| GetStocks | [GetStocksRequest](#stock-manager-v1-GetStocksRequest) | [GetStocksResponse](#stock-manager-v1-GetStocksResponse) |  |
| ExistStock | [ExistStockRequest](#stock-manager-v1-ExistStockRequest) | [ExistStockResponse](#stock-manager-v1-ExistStockResponse) |  |
| ExistStockConds | [ExistStockCondsRequest](#stock-manager-v1-ExistStockCondsRequest) | [ExistStockCondsResponse](#stock-manager-v1-ExistStockCondsResponse) |  |
| CountStocks | [CountStocksRequest](#stock-manager-v1-CountStocksRequest) | [CountStocksResponse](#stock-manager-v1-CountStocksResponse) |  |
| DeleteStock | [DeleteStockRequest](#stock-manager-v1-DeleteStockRequest) | [DeleteStockResponse](#stock-manager-v1-DeleteStockResponse) |  |

 



<a name="npool_stockmgr_stockmgr-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/stockmgr/stockmgr.proto



<a name="stock-manager-v1-AddStockFieldsRequest"></a>

### AddStockFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [AddStockFieldsRequest.FieldsEntry](#stock-manager-v1-AddStockFieldsRequest-FieldsEntry) | repeated |  |






<a name="stock-manager-v1-AddStockFieldsRequest-FieldsEntry"></a>

### AddStockFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stock-manager-v1-AddStockFieldsResponse"></a>

### AddStockFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CountStocksRequest"></a>

### CountStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [CountStocksRequest.CondsEntry](#stock-manager-v1-CountStocksRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-CountStocksRequest-CondsEntry"></a>

### CountStocksRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-CountStocksResponse"></a>

### CountStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [uint32](#uint32) |  |  |






<a name="stock-manager-v1-CreateStockRequest"></a>

### CreateStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CreateStockResponse"></a>

### CreateStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-CreateStocksRequest"></a>

### CreateStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |






<a name="stock-manager-v1-CreateStocksResponse"></a>

### CreateStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |






<a name="stock-manager-v1-DeleteStockRequest"></a>

### DeleteStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-DeleteStockResponse"></a>

### DeleteStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-ExistStockCondsRequest"></a>

### ExistStockCondsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [ExistStockCondsRequest.CondsEntry](#stock-manager-v1-ExistStockCondsRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-ExistStockCondsRequest-CondsEntry"></a>

### ExistStockCondsRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-ExistStockCondsResponse"></a>

### ExistStockCondsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="stock-manager-v1-ExistStockRequest"></a>

### ExistStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-ExistStockResponse"></a>

### ExistStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Result | [bool](#bool) |  |  |






<a name="stock-manager-v1-GetStockOnlyRequest"></a>

### GetStockOnlyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetStockOnlyRequest.CondsEntry](#stock-manager-v1-GetStockOnlyRequest-CondsEntry) | repeated |  |






<a name="stock-manager-v1-GetStockOnlyRequest-CondsEntry"></a>

### GetStockOnlyRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-GetStockOnlyResponse"></a>

### GetStockOnlyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-GetStockRequest"></a>

### GetStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="stock-manager-v1-GetStockResponse"></a>

### GetStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-GetStocksRequest"></a>

### GetStocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Conds | [GetStocksRequest.CondsEntry](#stock-manager-v1-GetStocksRequest-CondsEntry) | repeated |  |
| Offset | [int32](#int32) |  |  |
| Limit | [int32](#int32) |  |  |






<a name="stock-manager-v1-GetStocksRequest-CondsEntry"></a>

### GetStocksRequest.CondsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [npool.v1.FilterCond](#npool-v1-FilterCond) |  |  |






<a name="stock-manager-v1-GetStocksResponse"></a>

### GetStocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Stock](#stock-manager-v1-Stock) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="stock-manager-v1-Stock"></a>

### Stock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Total | [uint32](#uint32) |  |  |
| InService | [uint32](#uint32) |  |  |
| Sold | [uint32](#uint32) |  |  |
| Locked | [uint32](#uint32) |  |  |






<a name="stock-manager-v1-UpdateStockFieldsRequest"></a>

### UpdateStockFieldsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Fields | [UpdateStockFieldsRequest.FieldsEntry](#stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry) | repeated |  |






<a name="stock-manager-v1-UpdateStockFieldsRequest-FieldsEntry"></a>

### UpdateStockFieldsRequest.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stock-manager-v1-UpdateStockFieldsResponse"></a>

### UpdateStockFieldsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-UpdateStockRequest"></a>

### UpdateStockRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |






<a name="stock-manager-v1-UpdateStockResponse"></a>

### UpdateStockResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Stock](#stock-manager-v1-Stock) |  |  |





 

 

 


<a name="stock-manager-v1-StockManager"></a>

### StockManager


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) |  |
| CreateStock | [CreateStockRequest](#stock-manager-v1-CreateStockRequest) | [CreateStockResponse](#stock-manager-v1-CreateStockResponse) |  |
| CreateStocks | [CreateStocksRequest](#stock-manager-v1-CreateStocksRequest) | [CreateStocksResponse](#stock-manager-v1-CreateStocksResponse) |  |
| UpdateStock | [UpdateStockRequest](#stock-manager-v1-UpdateStockRequest) | [UpdateStockResponse](#stock-manager-v1-UpdateStockResponse) |  |
| UpdateStockFields | [UpdateStockFieldsRequest](#stock-manager-v1-UpdateStockFieldsRequest) | [UpdateStockFieldsResponse](#stock-manager-v1-UpdateStockFieldsResponse) |  |
| AddStockFields | [AddStockFieldsRequest](#stock-manager-v1-AddStockFieldsRequest) | [AddStockFieldsResponse](#stock-manager-v1-AddStockFieldsResponse) |  |
| GetStock | [GetStockRequest](#stock-manager-v1-GetStockRequest) | [GetStockResponse](#stock-manager-v1-GetStockResponse) |  |
| GetStockOnly | [GetStockOnlyRequest](#stock-manager-v1-GetStockOnlyRequest) | [GetStockOnlyResponse](#stock-manager-v1-GetStockOnlyResponse) |  |
| GetStocks | [GetStocksRequest](#stock-manager-v1-GetStocksRequest) | [GetStocksResponse](#stock-manager-v1-GetStocksResponse) |  |
| ExistStock | [ExistStockRequest](#stock-manager-v1-ExistStockRequest) | [ExistStockResponse](#stock-manager-v1-ExistStockResponse) |  |
| ExistStockConds | [ExistStockCondsRequest](#stock-manager-v1-ExistStockCondsRequest) | [ExistStockCondsResponse](#stock-manager-v1-ExistStockCondsResponse) |  |
| CountStocks | [CountStocksRequest](#stock-manager-v1-CountStocksRequest) | [CountStocksResponse](#stock-manager-v1-CountStocksResponse) |  |
| DeleteStock | [DeleteStockRequest](#stock-manager-v1-DeleteStockRequest) | [DeleteStockResponse](#stock-manager-v1-DeleteStockResponse) |  |

 



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

