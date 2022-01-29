# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-order/cloud-hashing-order.proto](#npool/cloud-hashing-order/cloud-hashing-order.proto)
    - [Compensate](#cloud.hashing.order.v1.Compensate)
    - [CreateCompensateRequest](#cloud.hashing.order.v1.CreateCompensateRequest)
    - [CreateCompensateResponse](#cloud.hashing.order.v1.CreateCompensateResponse)
    - [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest)
    - [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse)
    - [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest)
    - [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse)
    - [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest)
    - [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse)
    - [CreateOutOfGasRequest](#cloud.hashing.order.v1.CreateOutOfGasRequest)
    - [CreateOutOfGasResponse](#cloud.hashing.order.v1.CreateOutOfGasResponse)
    - [CreatePaymentRequest](#cloud.hashing.order.v1.CreatePaymentRequest)
    - [CreatePaymentResponse](#cloud.hashing.order.v1.CreatePaymentResponse)
    - [GasPaying](#cloud.hashing.order.v1.GasPaying)
    - [GetCompensatesByOrderRequest](#cloud.hashing.order.v1.GetCompensatesByOrderRequest)
    - [GetCompensatesByOrderResponse](#cloud.hashing.order.v1.GetCompensatesByOrderResponse)
    - [GetGasPayingsByOrderRequest](#cloud.hashing.order.v1.GetGasPayingsByOrderRequest)
    - [GetGasPayingsByOrderResponse](#cloud.hashing.order.v1.GetGasPayingsByOrderResponse)
    - [GetGoodPayingByOrderRequest](#cloud.hashing.order.v1.GetGoodPayingByOrderRequest)
    - [GetGoodPayingByOrderResponse](#cloud.hashing.order.v1.GetGoodPayingByOrderResponse)
    - [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest)
    - [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse)
    - [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest)
    - [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse)
    - [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest)
    - [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse)
    - [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest)
    - [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse)
    - [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest)
    - [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse)
    - [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest)
    - [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse)
    - [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest)
    - [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse)
    - [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest)
    - [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse)
    - [GetOrdersShortDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest)
    - [GetOrdersShortDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse)
    - [GetOutOfGasesByOrderRequest](#cloud.hashing.order.v1.GetOutOfGasesByOrderRequest)
    - [GetOutOfGasesByOrderResponse](#cloud.hashing.order.v1.GetOutOfGasesByOrderResponse)
    - [GetPaymentByOrderRequest](#cloud.hashing.order.v1.GetPaymentByOrderRequest)
    - [GetPaymentByOrderResponse](#cloud.hashing.order.v1.GetPaymentByOrderResponse)
    - [GetPaymentRequest](#cloud.hashing.order.v1.GetPaymentRequest)
    - [GetPaymentResponse](#cloud.hashing.order.v1.GetPaymentResponse)
    - [GoodPaying](#cloud.hashing.order.v1.GoodPaying)
    - [Order](#cloud.hashing.order.v1.Order)
    - [OrderDetail](#cloud.hashing.order.v1.OrderDetail)
    - [OutOfGas](#cloud.hashing.order.v1.OutOfGas)
    - [Payment](#cloud.hashing.order.v1.Payment)
    - [UpdatePaymentRequest](#cloud.hashing.order.v1.UpdatePaymentRequest)
    - [UpdatePaymentResponse](#cloud.hashing.order.v1.UpdatePaymentResponse)
  
    - [CloudHashingOrder](#cloud.hashing.order.v1.CloudHashingOrder)
  
- [npool/cloud-hashing-order/cloud-hashing-order.proto](#npool/cloud-hashing-order/cloud-hashing-order.proto)
    - [Compensate](#cloud.hashing.order.v1.Compensate)
    - [CreateCompensateRequest](#cloud.hashing.order.v1.CreateCompensateRequest)
    - [CreateCompensateResponse](#cloud.hashing.order.v1.CreateCompensateResponse)
    - [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest)
    - [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse)
    - [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest)
    - [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse)
    - [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest)
    - [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse)
    - [CreateOutOfGasRequest](#cloud.hashing.order.v1.CreateOutOfGasRequest)
    - [CreateOutOfGasResponse](#cloud.hashing.order.v1.CreateOutOfGasResponse)
    - [CreatePaymentRequest](#cloud.hashing.order.v1.CreatePaymentRequest)
    - [CreatePaymentResponse](#cloud.hashing.order.v1.CreatePaymentResponse)
    - [GasPaying](#cloud.hashing.order.v1.GasPaying)
    - [GetCompensatesByOrderRequest](#cloud.hashing.order.v1.GetCompensatesByOrderRequest)
    - [GetCompensatesByOrderResponse](#cloud.hashing.order.v1.GetCompensatesByOrderResponse)
    - [GetGasPayingsByOrderRequest](#cloud.hashing.order.v1.GetGasPayingsByOrderRequest)
    - [GetGasPayingsByOrderResponse](#cloud.hashing.order.v1.GetGasPayingsByOrderResponse)
    - [GetGoodPayingByOrderRequest](#cloud.hashing.order.v1.GetGoodPayingByOrderRequest)
    - [GetGoodPayingByOrderResponse](#cloud.hashing.order.v1.GetGoodPayingByOrderResponse)
    - [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest)
    - [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse)
    - [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest)
    - [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse)
    - [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest)
    - [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse)
    - [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest)
    - [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse)
    - [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest)
    - [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse)
    - [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest)
    - [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse)
    - [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest)
    - [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse)
    - [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest)
    - [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse)
    - [GetOrdersShortDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest)
    - [GetOrdersShortDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse)
    - [GetOutOfGasesByOrderRequest](#cloud.hashing.order.v1.GetOutOfGasesByOrderRequest)
    - [GetOutOfGasesByOrderResponse](#cloud.hashing.order.v1.GetOutOfGasesByOrderResponse)
    - [GetPaymentByOrderRequest](#cloud.hashing.order.v1.GetPaymentByOrderRequest)
    - [GetPaymentByOrderResponse](#cloud.hashing.order.v1.GetPaymentByOrderResponse)
    - [GetPaymentRequest](#cloud.hashing.order.v1.GetPaymentRequest)
    - [GetPaymentResponse](#cloud.hashing.order.v1.GetPaymentResponse)
    - [GoodPaying](#cloud.hashing.order.v1.GoodPaying)
    - [Order](#cloud.hashing.order.v1.Order)
    - [OrderDetail](#cloud.hashing.order.v1.OrderDetail)
    - [OutOfGas](#cloud.hashing.order.v1.OutOfGas)
    - [Payment](#cloud.hashing.order.v1.Payment)
    - [UpdatePaymentRequest](#cloud.hashing.order.v1.UpdatePaymentRequest)
    - [UpdatePaymentResponse](#cloud.hashing.order.v1.UpdatePaymentResponse)
  
    - [CloudHashingOrder](#cloud.hashing.order.v1.CloudHashingOrder)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-order/cloud-hashing-order.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-order/cloud-hashing-order.proto



<a name="cloud.hashing.order.v1.Compensate"></a>

### Compensate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.CreateCompensateRequest"></a>

### CreateCompensateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Compensate](#cloud.hashing.order.v1.Compensate) |  |  |






<a name="cloud.hashing.order.v1.CreateCompensateResponse"></a>

### CreateCompensateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Compensate](#cloud.hashing.order.v1.Compensate) |  |  |






<a name="cloud.hashing.order.v1.CreateGasPayingRequest"></a>

### CreateGasPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGasPayingResponse"></a>

### CreateGasPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingRequest"></a>

### CreateGoodPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingResponse"></a>

### CreateGoodPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderRequest"></a>

### CreateOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderResponse"></a>

### CreateOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.CreateOutOfGasRequest"></a>

### CreateOutOfGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) |  |  |






<a name="cloud.hashing.order.v1.CreateOutOfGasResponse"></a>

### CreateOutOfGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) |  |  |






<a name="cloud.hashing.order.v1.CreatePaymentRequest"></a>

### CreatePaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.CreatePaymentResponse"></a>

### CreatePaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GasPaying"></a>

### GasPaying



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |
| DurationMinutes | [uint32](#uint32) |  |  |
| FeeTypeID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetCompensatesByOrderRequest"></a>

### GetCompensatesByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetCompensatesByOrderResponse"></a>

### GetCompensatesByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Compensate](#cloud.hashing.order.v1.Compensate) | repeated |  |






<a name="cloud.hashing.order.v1.GetGasPayingsByOrderRequest"></a>

### GetGasPayingsByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGasPayingsByOrderResponse"></a>

### GetGasPayingsByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GasPaying](#cloud.hashing.order.v1.GasPaying) | repeated |  |






<a name="cloud.hashing.order.v1.GetGoodPayingByOrderRequest"></a>

### GetGoodPayingByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGoodPayingByOrderResponse"></a>

### GetGoodPayingByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailRequest"></a>

### GetOrderDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailResponse"></a>

### GetOrderDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) |  |  |






<a name="cloud.hashing.order.v1.GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppRequest"></a>

### GetOrdersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppResponse"></a>

### GetOrdersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserRequest"></a>

### GetOrdersByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserResponse"></a>

### GetOrdersByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodRequest"></a>

### GetOrdersByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodResponse"></a>

### GetOrdersByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppRequest"></a>

### GetOrdersDetailByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppResponse"></a>

### GetOrdersDetailByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest"></a>

### GetOrdersDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse"></a>

### GetOrdersDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodRequest"></a>

### GetOrdersDetailByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodResponse"></a>

### GetOrdersDetailByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest"></a>

### GetOrdersShortDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse"></a>

### GetOrdersShortDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOutOfGasesByOrderRequest"></a>

### GetOutOfGasesByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOutOfGasesByOrderResponse"></a>

### GetOutOfGasesByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) | repeated |  |






<a name="cloud.hashing.order.v1.GetPaymentByOrderRequest"></a>

### GetPaymentByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentByOrderResponse"></a>

### GetPaymentByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentRequest"></a>

### GetPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentResponse"></a>

### GetPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GoodPaying"></a>

### GoodPaying
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.OrderDetail"></a>

### OrderDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |
| GoodPaying | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |
| GasPayings | [GasPaying](#cloud.hashing.order.v1.GasPaying) | repeated |  |
| Compensates | [Compensate](#cloud.hashing.order.v1.Compensate) | repeated |  |
| OutOfGases | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) | repeated |  |
| Payment | [Payment](#cloud.hashing.order.v1.Payment) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.OutOfGas"></a>

### OutOfGas



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |






<a name="cloud.hashing.order.v1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| StartAmount | [double](#double) |  |  |
| Amount | [double](#double) |  |  |
| CoinInfoID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| PlatformTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.UpdatePaymentRequest"></a>

### UpdatePaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.UpdatePaymentResponse"></a>

### UpdatePaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |





 

 

 


<a name="cloud.hashing.order.v1.CloudHashingOrder"></a>

### CloudHashingOrder
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| CreateGoodPaying | [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest) | [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse) |  |
| GetGoodPayingByOrder | [GetGoodPayingByOrderRequest](#cloud.hashing.order.v1.GetGoodPayingByOrderRequest) | [GetGoodPayingByOrderResponse](#cloud.hashing.order.v1.GetGoodPayingByOrderResponse) |  |
| CreateGasPaying | [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest) | [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse) |  |
| GetGasPayingsByOrder | [GetGasPayingsByOrderRequest](#cloud.hashing.order.v1.GetGasPayingsByOrderRequest) | [GetGasPayingsByOrderResponse](#cloud.hashing.order.v1.GetGasPayingsByOrderResponse) |  |
| CreateOrder | [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest) | [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse) |  |
| GetOrder | [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest) | [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse) |  |
| GetOrderDetail | [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest) | [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse) |  |
| GetOrdersByAppUser | [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest) | [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse) |  |
| GetOrdersByApp | [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest) | [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse) |  |
| GetOrdersByGood | [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest) | [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse) |  |
| CreateCompensate | [CreateCompensateRequest](#cloud.hashing.order.v1.CreateCompensateRequest) | [CreateCompensateResponse](#cloud.hashing.order.v1.CreateCompensateResponse) |  |
| GetCompensatesByOrder | [GetCompensatesByOrderRequest](#cloud.hashing.order.v1.GetCompensatesByOrderRequest) | [GetCompensatesByOrderResponse](#cloud.hashing.order.v1.GetCompensatesByOrderResponse) |  |
| CreateOutOfGas | [CreateOutOfGasRequest](#cloud.hashing.order.v1.CreateOutOfGasRequest) | [CreateOutOfGasResponse](#cloud.hashing.order.v1.CreateOutOfGasResponse) |  |
| GetOutOfGasesByOrder | [GetOutOfGasesByOrderRequest](#cloud.hashing.order.v1.GetOutOfGasesByOrderRequest) | [GetOutOfGasesByOrderResponse](#cloud.hashing.order.v1.GetOutOfGasesByOrderResponse) |  |
| CreatePayment | [CreatePaymentRequest](#cloud.hashing.order.v1.CreatePaymentRequest) | [CreatePaymentResponse](#cloud.hashing.order.v1.CreatePaymentResponse) |  |
| GetPayment | [GetPaymentRequest](#cloud.hashing.order.v1.GetPaymentRequest) | [GetPaymentResponse](#cloud.hashing.order.v1.GetPaymentResponse) |  |
| UpdatePayment | [UpdatePaymentRequest](#cloud.hashing.order.v1.UpdatePaymentRequest) | [UpdatePaymentResponse](#cloud.hashing.order.v1.UpdatePaymentResponse) |  |
| GetPaymentByOrder | [GetPaymentByOrderRequest](#cloud.hashing.order.v1.GetPaymentByOrderRequest) | [GetPaymentByOrderResponse](#cloud.hashing.order.v1.GetPaymentByOrderResponse) |  |
| GetOrdersDetailByAppUser | [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest) | [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse) |  |
| GetOrdersShortDetailByAppUser | [GetOrdersShortDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest) | [GetOrdersShortDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse) |  |
| GetOrdersDetailByApp | [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest) | [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse) |  |
| GetOrdersDetailByGood | [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest) | [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse) |  |

 



<a name="npool/cloud-hashing-order/cloud-hashing-order.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-order/cloud-hashing-order.proto



<a name="cloud.hashing.order.v1.Compensate"></a>

### Compensate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.CreateCompensateRequest"></a>

### CreateCompensateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Compensate](#cloud.hashing.order.v1.Compensate) |  |  |






<a name="cloud.hashing.order.v1.CreateCompensateResponse"></a>

### CreateCompensateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Compensate](#cloud.hashing.order.v1.Compensate) |  |  |






<a name="cloud.hashing.order.v1.CreateGasPayingRequest"></a>

### CreateGasPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGasPayingResponse"></a>

### CreateGasPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GasPaying](#cloud.hashing.order.v1.GasPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingRequest"></a>

### CreateGoodPayingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateGoodPayingResponse"></a>

### CreateGoodPayingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderRequest"></a>

### CreateOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.CreateOrderResponse"></a>

### CreateOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.CreateOutOfGasRequest"></a>

### CreateOutOfGasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) |  |  |






<a name="cloud.hashing.order.v1.CreateOutOfGasResponse"></a>

### CreateOutOfGasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) |  |  |






<a name="cloud.hashing.order.v1.CreatePaymentRequest"></a>

### CreatePaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.CreatePaymentResponse"></a>

### CreatePaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GasPaying"></a>

### GasPaying



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |
| DurationMinutes | [uint32](#uint32) |  |  |
| FeeTypeID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetCompensatesByOrderRequest"></a>

### GetCompensatesByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetCompensatesByOrderResponse"></a>

### GetCompensatesByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Compensate](#cloud.hashing.order.v1.Compensate) | repeated |  |






<a name="cloud.hashing.order.v1.GetGasPayingsByOrderRequest"></a>

### GetGasPayingsByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGasPayingsByOrderResponse"></a>

### GetGasPayingsByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GasPaying](#cloud.hashing.order.v1.GasPaying) | repeated |  |






<a name="cloud.hashing.order.v1.GetGoodPayingByOrderRequest"></a>

### GetGoodPayingByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetGoodPayingByOrderResponse"></a>

### GetGoodPayingByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailRequest"></a>

### GetOrderDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderDetailResponse"></a>

### GetOrderDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) |  |  |






<a name="cloud.hashing.order.v1.GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Order](#cloud.hashing.order.v1.Order) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppRequest"></a>

### GetOrdersByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppResponse"></a>

### GetOrdersByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserRequest"></a>

### GetOrdersByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByAppUserResponse"></a>

### GetOrdersByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodRequest"></a>

### GetOrdersByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersByGoodResponse"></a>

### GetOrdersByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Order](#cloud.hashing.order.v1.Order) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppRequest"></a>

### GetOrdersDetailByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppResponse"></a>

### GetOrdersDetailByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest"></a>

### GetOrdersDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse"></a>

### GetOrdersDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodRequest"></a>

### GetOrdersDetailByGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersDetailByGoodResponse"></a>

### GetOrdersDetailByGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest"></a>

### GetOrdersShortDetailByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse"></a>

### GetOrdersShortDetailByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [OrderDetail](#cloud.hashing.order.v1.OrderDetail) | repeated |  |






<a name="cloud.hashing.order.v1.GetOutOfGasesByOrderRequest"></a>

### GetOutOfGasesByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetOutOfGasesByOrderResponse"></a>

### GetOutOfGasesByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) | repeated |  |






<a name="cloud.hashing.order.v1.GetPaymentByOrderRequest"></a>

### GetPaymentByOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| OrderID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentByOrderResponse"></a>

### GetPaymentByOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentRequest"></a>

### GetPaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.GetPaymentResponse"></a>

### GetPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.GoodPaying"></a>

### GoodPaying
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| PaymentID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.Order"></a>

### Order



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.OrderDetail"></a>

### OrderDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Units | [uint32](#uint32) |  |  |
| DiscountCouponID | [string](#string) |  |  |
| UserSpecialReductionID | [string](#string) |  |  |
| GoodPaying | [GoodPaying](#cloud.hashing.order.v1.GoodPaying) |  |  |
| GasPayings | [GasPaying](#cloud.hashing.order.v1.GasPaying) | repeated |  |
| Compensates | [Compensate](#cloud.hashing.order.v1.Compensate) | repeated |  |
| OutOfGases | [OutOfGas](#cloud.hashing.order.v1.OutOfGas) | repeated |  |
| Payment | [Payment](#cloud.hashing.order.v1.Payment) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |
| CouponID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.OutOfGas"></a>

### OutOfGas



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Start | [uint32](#uint32) |  |  |
| End | [uint32](#uint32) |  |  |






<a name="cloud.hashing.order.v1.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| AccountID | [string](#string) |  |  |
| StartAmount | [double](#double) |  |  |
| Amount | [double](#double) |  |  |
| CoinInfoID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| ChainTransactionID | [string](#string) |  |  |
| PlatformTransactionID | [string](#string) |  |  |






<a name="cloud.hashing.order.v1.UpdatePaymentRequest"></a>

### UpdatePaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |






<a name="cloud.hashing.order.v1.UpdatePaymentResponse"></a>

### UpdatePaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Payment](#cloud.hashing.order.v1.Payment) |  |  |





 

 

 


<a name="cloud.hashing.order.v1.CloudHashingOrder"></a>

### CloudHashingOrder
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| CreateGoodPaying | [CreateGoodPayingRequest](#cloud.hashing.order.v1.CreateGoodPayingRequest) | [CreateGoodPayingResponse](#cloud.hashing.order.v1.CreateGoodPayingResponse) |  |
| GetGoodPayingByOrder | [GetGoodPayingByOrderRequest](#cloud.hashing.order.v1.GetGoodPayingByOrderRequest) | [GetGoodPayingByOrderResponse](#cloud.hashing.order.v1.GetGoodPayingByOrderResponse) |  |
| CreateGasPaying | [CreateGasPayingRequest](#cloud.hashing.order.v1.CreateGasPayingRequest) | [CreateGasPayingResponse](#cloud.hashing.order.v1.CreateGasPayingResponse) |  |
| GetGasPayingsByOrder | [GetGasPayingsByOrderRequest](#cloud.hashing.order.v1.GetGasPayingsByOrderRequest) | [GetGasPayingsByOrderResponse](#cloud.hashing.order.v1.GetGasPayingsByOrderResponse) |  |
| CreateOrder | [CreateOrderRequest](#cloud.hashing.order.v1.CreateOrderRequest) | [CreateOrderResponse](#cloud.hashing.order.v1.CreateOrderResponse) |  |
| GetOrder | [GetOrderRequest](#cloud.hashing.order.v1.GetOrderRequest) | [GetOrderResponse](#cloud.hashing.order.v1.GetOrderResponse) |  |
| GetOrderDetail | [GetOrderDetailRequest](#cloud.hashing.order.v1.GetOrderDetailRequest) | [GetOrderDetailResponse](#cloud.hashing.order.v1.GetOrderDetailResponse) |  |
| GetOrdersByAppUser | [GetOrdersByAppUserRequest](#cloud.hashing.order.v1.GetOrdersByAppUserRequest) | [GetOrdersByAppUserResponse](#cloud.hashing.order.v1.GetOrdersByAppUserResponse) |  |
| GetOrdersByApp | [GetOrdersByAppRequest](#cloud.hashing.order.v1.GetOrdersByAppRequest) | [GetOrdersByAppResponse](#cloud.hashing.order.v1.GetOrdersByAppResponse) |  |
| GetOrdersByGood | [GetOrdersByGoodRequest](#cloud.hashing.order.v1.GetOrdersByGoodRequest) | [GetOrdersByGoodResponse](#cloud.hashing.order.v1.GetOrdersByGoodResponse) |  |
| CreateCompensate | [CreateCompensateRequest](#cloud.hashing.order.v1.CreateCompensateRequest) | [CreateCompensateResponse](#cloud.hashing.order.v1.CreateCompensateResponse) |  |
| GetCompensatesByOrder | [GetCompensatesByOrderRequest](#cloud.hashing.order.v1.GetCompensatesByOrderRequest) | [GetCompensatesByOrderResponse](#cloud.hashing.order.v1.GetCompensatesByOrderResponse) |  |
| CreateOutOfGas | [CreateOutOfGasRequest](#cloud.hashing.order.v1.CreateOutOfGasRequest) | [CreateOutOfGasResponse](#cloud.hashing.order.v1.CreateOutOfGasResponse) |  |
| GetOutOfGasesByOrder | [GetOutOfGasesByOrderRequest](#cloud.hashing.order.v1.GetOutOfGasesByOrderRequest) | [GetOutOfGasesByOrderResponse](#cloud.hashing.order.v1.GetOutOfGasesByOrderResponse) |  |
| CreatePayment | [CreatePaymentRequest](#cloud.hashing.order.v1.CreatePaymentRequest) | [CreatePaymentResponse](#cloud.hashing.order.v1.CreatePaymentResponse) |  |
| GetPayment | [GetPaymentRequest](#cloud.hashing.order.v1.GetPaymentRequest) | [GetPaymentResponse](#cloud.hashing.order.v1.GetPaymentResponse) |  |
| UpdatePayment | [UpdatePaymentRequest](#cloud.hashing.order.v1.UpdatePaymentRequest) | [UpdatePaymentResponse](#cloud.hashing.order.v1.UpdatePaymentResponse) |  |
| GetPaymentByOrder | [GetPaymentByOrderRequest](#cloud.hashing.order.v1.GetPaymentByOrderRequest) | [GetPaymentByOrderResponse](#cloud.hashing.order.v1.GetPaymentByOrderResponse) |  |
| GetOrdersDetailByAppUser | [GetOrdersDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppUserRequest) | [GetOrdersDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppUserResponse) |  |
| GetOrdersShortDetailByAppUser | [GetOrdersShortDetailByAppUserRequest](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserRequest) | [GetOrdersShortDetailByAppUserResponse](#cloud.hashing.order.v1.GetOrdersShortDetailByAppUserResponse) |  |
| GetOrdersDetailByApp | [GetOrdersDetailByAppRequest](#cloud.hashing.order.v1.GetOrdersDetailByAppRequest) | [GetOrdersDetailByAppResponse](#cloud.hashing.order.v1.GetOrdersDetailByAppResponse) |  |
| GetOrdersDetailByGood | [GetOrdersDetailByGoodRequest](#cloud.hashing.order.v1.GetOrdersDetailByGoodRequest) | [GetOrdersDetailByGoodResponse](#cloud.hashing.order.v1.GetOrdersDetailByGoodResponse) |  |

 



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

