# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/cloud-hashing-goods.proto](#npool/cloud-hashing-goods.proto)
    - [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo)
    - [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo)
    - [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo)
    - [AuthorizeAppGoodRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodRequest)
    - [AuthorizeAppGoodResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodResponse)
    - [AuthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest)
    - [AuthorizeAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaResponse)
    - [AuthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest)
    - [AuthorizeAppTargetAreaResponse](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaResponse)
    - [CheckAppGoodRequest](#cloud.hashing.goods.v1.CheckAppGoodRequest)
    - [CheckAppGoodResponse](#cloud.hashing.goods.v1.CheckAppGoodResponse)
    - [CheckAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.CheckAppGoodTargetAreaRequest)
    - [CheckAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.CheckAppGoodTargetAreaResponse)
    - [CheckAppTargetAreaRequest](#cloud.hashing.goods.v1.CheckAppTargetAreaRequest)
    - [CheckAppTargetAreaResponse](#cloud.hashing.goods.v1.CheckAppTargetAreaResponse)
    - [CreateDeviceInfoRequest](#cloud.hashing.goods.v1.CreateDeviceInfoRequest)
    - [CreateDeviceInfoResponse](#cloud.hashing.goods.v1.CreateDeviceInfoResponse)
    - [CreateFeeRequest](#cloud.hashing.goods.v1.CreateFeeRequest)
    - [CreateFeeResponse](#cloud.hashing.goods.v1.CreateFeeResponse)
    - [CreateFeeTypeRequest](#cloud.hashing.goods.v1.CreateFeeTypeRequest)
    - [CreateFeeTypeResponse](#cloud.hashing.goods.v1.CreateFeeTypeResponse)
    - [CreateGoodCommentRequest](#cloud.hashing.goods.v1.CreateGoodCommentRequest)
    - [CreateGoodCommentResponse](#cloud.hashing.goods.v1.CreateGoodCommentResponse)
    - [CreateGoodExtraInfoRequest](#cloud.hashing.goods.v1.CreateGoodExtraInfoRequest)
    - [CreateGoodExtraInfoResponse](#cloud.hashing.goods.v1.CreateGoodExtraInfoResponse)
    - [CreateGoodRequest](#cloud.hashing.goods.v1.CreateGoodRequest)
    - [CreateGoodResponse](#cloud.hashing.goods.v1.CreateGoodResponse)
    - [CreateGoodReviewRequest](#cloud.hashing.goods.v1.CreateGoodReviewRequest)
    - [CreateGoodReviewResponse](#cloud.hashing.goods.v1.CreateGoodReviewResponse)
    - [CreatePriceCurrencyRequest](#cloud.hashing.goods.v1.CreatePriceCurrencyRequest)
    - [CreatePriceCurrencyResponse](#cloud.hashing.goods.v1.CreatePriceCurrencyResponse)
    - [CreateRecommendRequest](#cloud.hashing.goods.v1.CreateRecommendRequest)
    - [CreateRecommendResponse](#cloud.hashing.goods.v1.CreateRecommendResponse)
    - [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest)
    - [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse)
    - [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest)
    - [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse)
    - [DeleteDeviceInfoRequest](#cloud.hashing.goods.v1.DeleteDeviceInfoRequest)
    - [DeleteDeviceInfoResponse](#cloud.hashing.goods.v1.DeleteDeviceInfoResponse)
    - [DeleteGoodRequest](#cloud.hashing.goods.v1.DeleteGoodRequest)
    - [DeleteGoodResponse](#cloud.hashing.goods.v1.DeleteGoodResponse)
    - [DeleteRecommendRequest](#cloud.hashing.goods.v1.DeleteRecommendRequest)
    - [DeleteRecommendResponse](#cloud.hashing.goods.v1.DeleteRecommendResponse)
    - [DeleteTargetAreaByContinentCountryRequest](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest)
    - [DeleteTargetAreaByContinentCountryResponse](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse)
    - [DeleteTargetAreaRequest](#cloud.hashing.goods.v1.DeleteTargetAreaRequest)
    - [DeleteTargetAreaResponse](#cloud.hashing.goods.v1.DeleteTargetAreaResponse)
    - [DeleteVendorLocationRequest](#cloud.hashing.goods.v1.DeleteVendorLocationRequest)
    - [DeleteVendorLocationResponse](#cloud.hashing.goods.v1.DeleteVendorLocationResponse)
    - [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo)
    - [Fee](#cloud.hashing.goods.v1.Fee)
    - [FeeDetail](#cloud.hashing.goods.v1.FeeDetail)
    - [FeeType](#cloud.hashing.goods.v1.FeeType)
    - [GetDeviceInfoRequest](#cloud.hashing.goods.v1.GetDeviceInfoRequest)
    - [GetDeviceInfoResponse](#cloud.hashing.goods.v1.GetDeviceInfoResponse)
    - [GetDeviceInfosRequest](#cloud.hashing.goods.v1.GetDeviceInfosRequest)
    - [GetDeviceInfosResponse](#cloud.hashing.goods.v1.GetDeviceInfosResponse)
    - [GetFeeRequest](#cloud.hashing.goods.v1.GetFeeRequest)
    - [GetFeeResponse](#cloud.hashing.goods.v1.GetFeeResponse)
    - [GetFeeTypeRequest](#cloud.hashing.goods.v1.GetFeeTypeRequest)
    - [GetFeeTypeResponse](#cloud.hashing.goods.v1.GetFeeTypeResponse)
    - [GetFeeTypesRequest](#cloud.hashing.goods.v1.GetFeeTypesRequest)
    - [GetFeeTypesResponse](#cloud.hashing.goods.v1.GetFeeTypesResponse)
    - [GetFeesRequest](#cloud.hashing.goods.v1.GetFeesRequest)
    - [GetFeesResponse](#cloud.hashing.goods.v1.GetFeesResponse)
    - [GetGoodCommentsRequest](#cloud.hashing.goods.v1.GetGoodCommentsRequest)
    - [GetGoodCommentsResponse](#cloud.hashing.goods.v1.GetGoodCommentsResponse)
    - [GetGoodDetailRequest](#cloud.hashing.goods.v1.GetGoodDetailRequest)
    - [GetGoodDetailResponse](#cloud.hashing.goods.v1.GetGoodDetailResponse)
    - [GetGoodExtraInfoRequest](#cloud.hashing.goods.v1.GetGoodExtraInfoRequest)
    - [GetGoodExtraInfoResponse](#cloud.hashing.goods.v1.GetGoodExtraInfoResponse)
    - [GetGoodRequest](#cloud.hashing.goods.v1.GetGoodRequest)
    - [GetGoodResponse](#cloud.hashing.goods.v1.GetGoodResponse)
    - [GetGoodReviewRequest](#cloud.hashing.goods.v1.GetGoodReviewRequest)
    - [GetGoodReviewResponse](#cloud.hashing.goods.v1.GetGoodReviewResponse)
    - [GetGoodsDetailRequest](#cloud.hashing.goods.v1.GetGoodsDetailRequest)
    - [GetGoodsDetailResponse](#cloud.hashing.goods.v1.GetGoodsDetailResponse)
    - [GetGoodsRequest](#cloud.hashing.goods.v1.GetGoodsRequest)
    - [GetGoodsResponse](#cloud.hashing.goods.v1.GetGoodsResponse)
    - [GetPriceCurrencyRequest](#cloud.hashing.goods.v1.GetPriceCurrencyRequest)
    - [GetPriceCurrencyResponse](#cloud.hashing.goods.v1.GetPriceCurrencyResponse)
    - [GetPriceCurrencysRequest](#cloud.hashing.goods.v1.GetPriceCurrencysRequest)
    - [GetPriceCurrencysResponse](#cloud.hashing.goods.v1.GetPriceCurrencysResponse)
    - [GetRecommendGoodsByAppRequest](#cloud.hashing.goods.v1.GetRecommendGoodsByAppRequest)
    - [GetRecommendGoodsByAppResponse](#cloud.hashing.goods.v1.GetRecommendGoodsByAppResponse)
    - [GetRecommendGoodsByRecommenderRequest](#cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderRequest)
    - [GetRecommendGoodsByRecommenderResponse](#cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderResponse)
    - [GetRecommendsByAppRequest](#cloud.hashing.goods.v1.GetRecommendsByAppRequest)
    - [GetRecommendsByAppResponse](#cloud.hashing.goods.v1.GetRecommendsByAppResponse)
    - [GetRecommendsByRecommenderRequest](#cloud.hashing.goods.v1.GetRecommendsByRecommenderRequest)
    - [GetRecommendsByRecommenderResponse](#cloud.hashing.goods.v1.GetRecommendsByRecommenderResponse)
    - [GetTargetAreaRequest](#cloud.hashing.goods.v1.GetTargetAreaRequest)
    - [GetTargetAreaResponse](#cloud.hashing.goods.v1.GetTargetAreaResponse)
    - [GetTargetAreasRequest](#cloud.hashing.goods.v1.GetTargetAreasRequest)
    - [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse)
    - [GetVendorLocationRequest](#cloud.hashing.goods.v1.GetVendorLocationRequest)
    - [GetVendorLocationResponse](#cloud.hashing.goods.v1.GetVendorLocationResponse)
    - [GetVendorLocationsRequest](#cloud.hashing.goods.v1.GetVendorLocationsRequest)
    - [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse)
    - [GoodComment](#cloud.hashing.goods.v1.GoodComment)
    - [GoodDetail](#cloud.hashing.goods.v1.GoodDetail)
    - [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo)
    - [GoodInfo](#cloud.hashing.goods.v1.GoodInfo)
    - [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo)
    - [OffsaleAppGoodRequest](#cloud.hashing.goods.v1.OffsaleAppGoodRequest)
    - [OffsaleAppGoodResponse](#cloud.hashing.goods.v1.OffsaleAppGoodResponse)
    - [OnsaleAppGoodRequest](#cloud.hashing.goods.v1.OnsaleAppGoodRequest)
    - [OnsaleAppGoodResponse](#cloud.hashing.goods.v1.OnsaleAppGoodResponse)
    - [PageInfo](#cloud.hashing.goods.v1.PageInfo)
    - [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency)
    - [Recommend](#cloud.hashing.goods.v1.Recommend)
    - [RecommendGood](#cloud.hashing.goods.v1.RecommendGood)
    - [SetAppGoodPriceRequest](#cloud.hashing.goods.v1.SetAppGoodPriceRequest)
    - [SetAppGoodPriceResponse](#cloud.hashing.goods.v1.SetAppGoodPriceResponse)
    - [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo)
    - [UnauthorizeAppGoodRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodRequest)
    - [UnauthorizeAppGoodResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodResponse)
    - [UnauthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest)
    - [UnauthorizeAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaResponse)
    - [UnauthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest)
    - [UnauthorizeAppTargetAreaResponse](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaResponse)
    - [UpdateDeviceInfoRequest](#cloud.hashing.goods.v1.UpdateDeviceInfoRequest)
    - [UpdateDeviceInfoResponse](#cloud.hashing.goods.v1.UpdateDeviceInfoResponse)
    - [UpdateFeeTypeRequest](#cloud.hashing.goods.v1.UpdateFeeTypeRequest)
    - [UpdateFeeTypeResponse](#cloud.hashing.goods.v1.UpdateFeeTypeResponse)
    - [UpdateGoodCommentRequest](#cloud.hashing.goods.v1.UpdateGoodCommentRequest)
    - [UpdateGoodCommentResponse](#cloud.hashing.goods.v1.UpdateGoodCommentResponse)
    - [UpdateGoodExtraInfoRequest](#cloud.hashing.goods.v1.UpdateGoodExtraInfoRequest)
    - [UpdateGoodExtraInfoResponse](#cloud.hashing.goods.v1.UpdateGoodExtraInfoResponse)
    - [UpdateGoodRequest](#cloud.hashing.goods.v1.UpdateGoodRequest)
    - [UpdateGoodResponse](#cloud.hashing.goods.v1.UpdateGoodResponse)
    - [UpdateGoodReviewRequest](#cloud.hashing.goods.v1.UpdateGoodReviewRequest)
    - [UpdateGoodReviewResponse](#cloud.hashing.goods.v1.UpdateGoodReviewResponse)
    - [UpdatePriceCurrencyRequest](#cloud.hashing.goods.v1.UpdatePriceCurrencyRequest)
    - [UpdatePriceCurrencyResponse](#cloud.hashing.goods.v1.UpdatePriceCurrencyResponse)
    - [UpdateRecommendRequest](#cloud.hashing.goods.v1.UpdateRecommendRequest)
    - [UpdateRecommendResponse](#cloud.hashing.goods.v1.UpdateRecommendResponse)
    - [UpdateTargetAreaRequest](#cloud.hashing.goods.v1.UpdateTargetAreaRequest)
    - [UpdateTargetAreaResponse](#cloud.hashing.goods.v1.UpdateTargetAreaResponse)
    - [UpdateVendorLocationRequest](#cloud.hashing.goods.v1.UpdateVendorLocationRequest)
    - [UpdateVendorLocationResponse](#cloud.hashing.goods.v1.UpdateVendorLocationResponse)
    - [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo)
    - [VersionResponse](#cloud.hashing.goods.v1.VersionResponse)
  
    - [CloudHashingGoods](#cloud.hashing.goods.v1.CloudHashingGoods)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/cloud-hashing-goods.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/cloud-hashing-goods.proto



<a name="cloud.hashing.goods.v1.AppGoodInfo"></a>

### AppGoodInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Price | [double](#double) |  |  |
| Authorized | [bool](#bool) |  |  |
| Online | [bool](#bool) |  |  |
| InitAreaStrategy | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.AppGoodTargetAreaInfo"></a>

### AppGoodTargetAreaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.AppTargetAreaInfo"></a>

### AppTargetAreaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| TargetAreaID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodRequest"></a>

### AuthorizeAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodResponse"></a>

### AuthorizeAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest"></a>

### AuthorizeAppGoodTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaResponse"></a>

### AuthorizeAppGoodTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest"></a>

### AuthorizeAppTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.AuthorizeAppTargetAreaResponse"></a>

### AuthorizeAppTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppGoodRequest"></a>

### CheckAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppGoodResponse"></a>

### CheckAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppGoodTargetAreaRequest"></a>

### CheckAppGoodTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppGoodTargetAreaResponse"></a>

### CheckAppGoodTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppTargetAreaRequest"></a>

### CheckAppTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CheckAppTargetAreaResponse"></a>

### CheckAppTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo) |  |  |
| Authorized | [bool](#bool) |  |  |






<a name="cloud.hashing.goods.v1.CreateDeviceInfoRequest"></a>

### CreateDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateDeviceInfoResponse"></a>

### CreateDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateFeeRequest"></a>

### CreateFeeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Fee](#cloud.hashing.goods.v1.Fee) |  |  |






<a name="cloud.hashing.goods.v1.CreateFeeResponse"></a>

### CreateFeeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Fee](#cloud.hashing.goods.v1.Fee) |  |  |






<a name="cloud.hashing.goods.v1.CreateFeeTypeRequest"></a>

### CreateFeeTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |






<a name="cloud.hashing.goods.v1.CreateFeeTypeResponse"></a>

### CreateFeeTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodCommentRequest"></a>

### CreateGoodCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodCommentResponse"></a>

### CreateGoodCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodExtraInfoRequest"></a>

### CreateGoodExtraInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodExtraInfoResponse"></a>

### CreateGoodExtraInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodRequest"></a>

### CreateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodResponse"></a>

### CreateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodReviewRequest"></a>

### CreateGoodReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateGoodReviewResponse"></a>

### CreateGoodReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreatePriceCurrencyRequest"></a>

### CreatePriceCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |






<a name="cloud.hashing.goods.v1.CreatePriceCurrencyResponse"></a>

### CreatePriceCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |






<a name="cloud.hashing.goods.v1.CreateRecommendRequest"></a>

### CreateRecommendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |






<a name="cloud.hashing.goods.v1.CreateRecommendResponse"></a>

### CreateRecommendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |






<a name="cloud.hashing.goods.v1.CreateTargetAreaRequest"></a>

### CreateTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateTargetAreaResponse"></a>

### CreateTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateVendorLocationRequest"></a>

### CreateVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.CreateVendorLocationResponse"></a>

### CreateVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteDeviceInfoRequest"></a>

### DeleteDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteDeviceInfoResponse"></a>

### DeleteDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteGoodRequest"></a>

### DeleteGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteGoodResponse"></a>

### DeleteGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteRecommendRequest"></a>

### DeleteRecommendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteRecommendResponse"></a>

### DeleteRecommendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest"></a>

### DeleteTargetAreaByContinentCountryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Continent | [string](#string) |  |  |
| Country | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse"></a>

### DeleteTargetAreaByContinentCountryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaRequest"></a>

### DeleteTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteTargetAreaResponse"></a>

### DeleteTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeleteVendorLocationRequest"></a>

### DeleteVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.DeleteVendorLocationResponse"></a>

### DeleteVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.DeviceInfo"></a>

### DeviceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Manufacturer | [string](#string) |  |  |
| PowerComsuption | [int32](#int32) |  |  |
| ShipmentAt | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.Fee"></a>

### Fee



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| FeeTypeID | [string](#string) |  |  |
| Value | [double](#double) |  |  |






<a name="cloud.hashing.goods.v1.FeeDetail"></a>

### FeeDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Fee | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |
| Value | [double](#double) |  |  |






<a name="cloud.hashing.goods.v1.FeeType"></a>

### FeeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| FeeType | [string](#string) |  |  |
| FeeDescription | [string](#string) |  |  |
| PayType | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfoRequest"></a>

### GetDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfoResponse"></a>

### GetDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfosRequest"></a>

### GetDeviceInfosRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetDeviceInfosResponse"></a>

### GetDeviceInfosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeRequest"></a>

### GetFeeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeResponse"></a>

### GetFeeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Fee](#cloud.hashing.goods.v1.Fee) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeTypeRequest"></a>

### GetFeeTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeTypeResponse"></a>

### GetFeeTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeTypesRequest"></a>

### GetFeeTypesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetFeeTypesResponse"></a>

### GetFeeTypesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [FeeType](#cloud.hashing.goods.v1.FeeType) | repeated |  |






<a name="cloud.hashing.goods.v1.GetFeesRequest"></a>

### GetFeesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetFeesResponse"></a>

### GetFeesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Fee](#cloud.hashing.goods.v1.Fee) | repeated |  |






<a name="cloud.hashing.goods.v1.GetGoodCommentsRequest"></a>

### GetGoodCommentsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodCommentsResponse"></a>

### GetGoodCommentsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comments | [GoodComment](#cloud.hashing.goods.v1.GoodComment) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodDetailRequest"></a>

### GetGoodDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodDetailResponse"></a>

### GetGoodDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Detail | [GoodDetail](#cloud.hashing.goods.v1.GoodDetail) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodExtraInfoRequest"></a>

### GetGoodExtraInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GoodID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodExtraInfoResponse"></a>

### GetGoodExtraInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodRequest"></a>

### GetGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodResponse"></a>

### GetGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodReviewRequest"></a>

### GetGoodReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodReviewResponse"></a>

### GetGoodReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsDetailRequest"></a>

### GetGoodsDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsDetailResponse"></a>

### GetGoodsDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Details | [GoodDetail](#cloud.hashing.goods.v1.GoodDetail) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsRequest"></a>

### GetGoodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetGoodsResponse"></a>

### GetGoodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetPriceCurrencyRequest"></a>

### GetPriceCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetPriceCurrencyResponse"></a>

### GetPriceCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |






<a name="cloud.hashing.goods.v1.GetPriceCurrencysRequest"></a>

### GetPriceCurrencysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetPriceCurrencysResponse"></a>

### GetPriceCurrencysResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) | repeated |  |






<a name="cloud.hashing.goods.v1.GetRecommendGoodsByAppRequest"></a>

### GetRecommendGoodsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetRecommendGoodsByAppResponse"></a>

### GetRecommendGoodsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RecommendGood](#cloud.hashing.goods.v1.RecommendGood) | repeated |  |






<a name="cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderRequest"></a>

### GetRecommendGoodsByRecommenderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderResponse"></a>

### GetRecommendGoodsByRecommenderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RecommendGood](#cloud.hashing.goods.v1.RecommendGood) | repeated |  |






<a name="cloud.hashing.goods.v1.GetRecommendsByAppRequest"></a>

### GetRecommendsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetRecommendsByAppResponse"></a>

### GetRecommendsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Recommend](#cloud.hashing.goods.v1.Recommend) | repeated |  |






<a name="cloud.hashing.goods.v1.GetRecommendsByRecommenderRequest"></a>

### GetRecommendsByRecommenderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetRecommendsByRecommenderResponse"></a>

### GetRecommendsByRecommenderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Recommend](#cloud.hashing.goods.v1.Recommend) | repeated |  |






<a name="cloud.hashing.goods.v1.GetTargetAreaRequest"></a>

### GetTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreaResponse"></a>

### GetTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreasRequest"></a>

### GetTargetAreasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetTargetAreasResponse"></a>

### GetTargetAreasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationRequest"></a>

### GetVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationResponse"></a>

### GetVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationsRequest"></a>

### GetVendorLocationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageInfo | [PageInfo](#cloud.hashing.goods.v1.PageInfo) |  |  |






<a name="cloud.hashing.goods.v1.GetVendorLocationsResponse"></a>

### GetVendorLocationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) | repeated |  |
| Total | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.GoodComment"></a>

### GoodComment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| ReplyToID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| OrderID | [string](#string) |  |  |
| Content | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.GoodDetail"></a>

### GoodDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DeviceInfo | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |
| SeparateFee | [bool](#bool) |  |  |
| UnitPower | [int32](#int32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| CoinInfoID | [string](#string) |  |  |
| Actuals | [bool](#bool) |  |  |
| DeliveryAt | [uint32](#uint32) |  |  |
| InheritFromGood | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |
| VendorLocation | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |
| Price | [double](#double) |  |  |
| BenefitType | [string](#string) |  |  |
| Classic | [bool](#bool) |  |  |
| SupportCoinTypeIDs | [string](#string) | repeated |  |
| Total | [int32](#int32) |  |  |
| Extra | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |
| Title | [string](#string) |  |  |
| Unit | [string](#string) |  |  |
| PriceCurrency | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |
| Fees | [FeeDetail](#cloud.hashing.goods.v1.FeeDetail) | repeated |  |






<a name="cloud.hashing.goods.v1.GoodExtraInfo"></a>

### GoodExtraInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| Posters | [string](#string) | repeated |  |
| Labels | [string](#string) | repeated |  |
| OutSale | [bool](#bool) |  |  |
| PreSale | [bool](#bool) |  |  |
| VoteCount | [uint32](#uint32) |  |  |
| Rating | [float](#float) |  |  |






<a name="cloud.hashing.goods.v1.GoodInfo"></a>

### GoodInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| DeviceInfoID | [string](#string) |  |  |
| SeparateFee | [bool](#bool) |  |  |
| UnitPower | [int32](#int32) |  |  |
| DurationDays | [int32](#int32) |  |  |
| CoinInfoID | [string](#string) |  |  |
| Actuals | [bool](#bool) |  |  |
| DeliveryAt | [uint32](#uint32) |  |  |
| InheritFromGoodID | [string](#string) |  |  |
| VendorLocationID | [string](#string) |  |  |
| Price | [double](#double) |  |  |
| BenefitType | [string](#string) |  |  |
| Classic | [bool](#bool) |  |  |
| SupportCoinTypeIDs | [string](#string) | repeated |  |
| Total | [int32](#int32) |  |  |
| PriceCurrency | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Unit | [string](#string) |  |  |
| FeeIDs | [string](#string) | repeated |  |






<a name="cloud.hashing.goods.v1.GoodReviewInfo"></a>

### GoodReviewInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| ReviewerID | [string](#string) |  |  |
| State | [string](#string) |  |  |
| Message | [string](#string) |  |  |
| ReviewedID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.OffsaleAppGoodRequest"></a>

### OffsaleAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.OffsaleAppGoodResponse"></a>

### OffsaleAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.OnsaleAppGoodRequest"></a>

### OnsaleAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.OnsaleAppGoodResponse"></a>

### OnsaleAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.PageInfo"></a>

### PageInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| PageIndex | [int32](#int32) |  |  |
| PageSize | [int32](#int32) |  |  |






<a name="cloud.hashing.goods.v1.PriceCurrency"></a>

### PriceCurrency



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Unit | [string](#string) |  |  |
| Symbol | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.Recommend"></a>

### Recommend



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| RecommenderID | [string](#string) |  |  |
| Message | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.RecommendGood"></a>

### RecommendGood



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Recommend | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |
| Good | [GoodDetail](#cloud.hashing.goods.v1.GoodDetail) |  |  |






<a name="cloud.hashing.goods.v1.SetAppGoodPriceRequest"></a>

### SetAppGoodPriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.SetAppGoodPriceResponse"></a>

### SetAppGoodPriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.TargetAreaInfo"></a>

### TargetAreaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Continent | [string](#string) |  |  |
| Country | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodRequest"></a>

### UnauthorizeAppGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodResponse"></a>

### UnauthorizeAppGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodInfo](#cloud.hashing.goods.v1.AppGoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest"></a>

### UnauthorizeAppGoodTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaResponse"></a>

### UnauthorizeAppGoodTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppGoodTargetAreaInfo](#cloud.hashing.goods.v1.AppGoodTargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest"></a>

### UnauthorizeAppTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.UnauthorizeAppTargetAreaResponse"></a>

### UnauthorizeAppTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [AppTargetAreaInfo](#cloud.hashing.goods.v1.AppTargetAreaInfo) |  |  |
| Authorized | [bool](#bool) |  |  |






<a name="cloud.hashing.goods.v1.UpdateDeviceInfoRequest"></a>

### UpdateDeviceInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateDeviceInfoResponse"></a>

### UpdateDeviceInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [DeviceInfo](#cloud.hashing.goods.v1.DeviceInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateFeeTypeRequest"></a>

### UpdateFeeTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |






<a name="cloud.hashing.goods.v1.UpdateFeeTypeResponse"></a>

### UpdateFeeTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [FeeType](#cloud.hashing.goods.v1.FeeType) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodCommentRequest"></a>

### UpdateGoodCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodCommentResponse"></a>

### UpdateGoodCommentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Comment | [GoodComment](#cloud.hashing.goods.v1.GoodComment) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodExtraInfoRequest"></a>

### UpdateGoodExtraInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodExtraInfoResponse"></a>

### UpdateGoodExtraInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodExtraInfo](#cloud.hashing.goods.v1.GoodExtraInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodRequest"></a>

### UpdateGoodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodResponse"></a>

### UpdateGoodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodInfo](#cloud.hashing.goods.v1.GoodInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodReviewRequest"></a>

### UpdateGoodReviewRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateGoodReviewResponse"></a>

### UpdateGoodReviewResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GoodReviewInfo](#cloud.hashing.goods.v1.GoodReviewInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdatePriceCurrencyRequest"></a>

### UpdatePriceCurrencyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |






<a name="cloud.hashing.goods.v1.UpdatePriceCurrencyResponse"></a>

### UpdatePriceCurrencyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [PriceCurrency](#cloud.hashing.goods.v1.PriceCurrency) |  |  |






<a name="cloud.hashing.goods.v1.UpdateRecommendRequest"></a>

### UpdateRecommendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |






<a name="cloud.hashing.goods.v1.UpdateRecommendResponse"></a>

### UpdateRecommendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Recommend](#cloud.hashing.goods.v1.Recommend) |  |  |






<a name="cloud.hashing.goods.v1.UpdateTargetAreaRequest"></a>

### UpdateTargetAreaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateTargetAreaResponse"></a>

### UpdateTargetAreaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TargetAreaInfo](#cloud.hashing.goods.v1.TargetAreaInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateVendorLocationRequest"></a>

### UpdateVendorLocationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.UpdateVendorLocationResponse"></a>

### UpdateVendorLocationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [VendorLocationInfo](#cloud.hashing.goods.v1.VendorLocationInfo) |  |  |






<a name="cloud.hashing.goods.v1.VendorLocationInfo"></a>

### VendorLocationInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Country | [string](#string) |  |  |
| Province | [string](#string) |  |  |
| City | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="cloud.hashing.goods.v1.VersionResponse"></a>

### VersionResponse
Request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="cloud.hashing.goods.v1.CloudHashingGoods"></a>

### CloudHashingGoods
Cloud Hashing Goods

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#cloud.hashing.goods.v1.VersionResponse) |  |
| CreateVendorLocation | [CreateVendorLocationRequest](#cloud.hashing.goods.v1.CreateVendorLocationRequest) | [CreateVendorLocationResponse](#cloud.hashing.goods.v1.CreateVendorLocationResponse) |  |
| UpdateVendorLocation | [UpdateVendorLocationRequest](#cloud.hashing.goods.v1.UpdateVendorLocationRequest) | [UpdateVendorLocationResponse](#cloud.hashing.goods.v1.UpdateVendorLocationResponse) |  |
| GetVendorLocation | [GetVendorLocationRequest](#cloud.hashing.goods.v1.GetVendorLocationRequest) | [GetVendorLocationResponse](#cloud.hashing.goods.v1.GetVendorLocationResponse) |  |
| DeleteVendorLocation | [DeleteVendorLocationRequest](#cloud.hashing.goods.v1.DeleteVendorLocationRequest) | [DeleteVendorLocationResponse](#cloud.hashing.goods.v1.DeleteVendorLocationResponse) |  |
| GetVendorLocations | [GetVendorLocationsRequest](#cloud.hashing.goods.v1.GetVendorLocationsRequest) | [GetVendorLocationsResponse](#cloud.hashing.goods.v1.GetVendorLocationsResponse) |  |
| CreatePriceCurrency | [CreatePriceCurrencyRequest](#cloud.hashing.goods.v1.CreatePriceCurrencyRequest) | [CreatePriceCurrencyResponse](#cloud.hashing.goods.v1.CreatePriceCurrencyResponse) |  |
| UpdatePriceCurrency | [UpdatePriceCurrencyRequest](#cloud.hashing.goods.v1.UpdatePriceCurrencyRequest) | [UpdatePriceCurrencyResponse](#cloud.hashing.goods.v1.UpdatePriceCurrencyResponse) |  |
| GetPriceCurrency | [GetPriceCurrencyRequest](#cloud.hashing.goods.v1.GetPriceCurrencyRequest) | [GetPriceCurrencyResponse](#cloud.hashing.goods.v1.GetPriceCurrencyResponse) |  |
| GetPriceCurrencys | [GetPriceCurrencysRequest](#cloud.hashing.goods.v1.GetPriceCurrencysRequest) | [GetPriceCurrencysResponse](#cloud.hashing.goods.v1.GetPriceCurrencysResponse) |  |
| CreateTargetArea | [CreateTargetAreaRequest](#cloud.hashing.goods.v1.CreateTargetAreaRequest) | [CreateTargetAreaResponse](#cloud.hashing.goods.v1.CreateTargetAreaResponse) |  |
| UpdateTargetArea | [UpdateTargetAreaRequest](#cloud.hashing.goods.v1.UpdateTargetAreaRequest) | [UpdateTargetAreaResponse](#cloud.hashing.goods.v1.UpdateTargetAreaResponse) |  |
| GetTargetArea | [GetTargetAreaRequest](#cloud.hashing.goods.v1.GetTargetAreaRequest) | [GetTargetAreaResponse](#cloud.hashing.goods.v1.GetTargetAreaResponse) |  |
| DeleteTargetArea | [DeleteTargetAreaRequest](#cloud.hashing.goods.v1.DeleteTargetAreaRequest) | [DeleteTargetAreaResponse](#cloud.hashing.goods.v1.DeleteTargetAreaResponse) |  |
| DeleteTargetAreaByContinentCountry | [DeleteTargetAreaByContinentCountryRequest](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryRequest) | [DeleteTargetAreaByContinentCountryResponse](#cloud.hashing.goods.v1.DeleteTargetAreaByContinentCountryResponse) |  |
| GetTargetAreas | [GetTargetAreasRequest](#cloud.hashing.goods.v1.GetTargetAreasRequest) | [GetTargetAreasResponse](#cloud.hashing.goods.v1.GetTargetAreasResponse) |  |
| CreateDeviceInfo | [CreateDeviceInfoRequest](#cloud.hashing.goods.v1.CreateDeviceInfoRequest) | [CreateDeviceInfoResponse](#cloud.hashing.goods.v1.CreateDeviceInfoResponse) |  |
| UpdateDeviceInfo | [UpdateDeviceInfoRequest](#cloud.hashing.goods.v1.UpdateDeviceInfoRequest) | [UpdateDeviceInfoResponse](#cloud.hashing.goods.v1.UpdateDeviceInfoResponse) |  |
| GetDeviceInfo | [GetDeviceInfoRequest](#cloud.hashing.goods.v1.GetDeviceInfoRequest) | [GetDeviceInfoResponse](#cloud.hashing.goods.v1.GetDeviceInfoResponse) |  |
| DeleteDeviceInfo | [DeleteDeviceInfoRequest](#cloud.hashing.goods.v1.DeleteDeviceInfoRequest) | [DeleteDeviceInfoResponse](#cloud.hashing.goods.v1.DeleteDeviceInfoResponse) |  |
| GetDeviceInfos | [GetDeviceInfosRequest](#cloud.hashing.goods.v1.GetDeviceInfosRequest) | [GetDeviceInfosResponse](#cloud.hashing.goods.v1.GetDeviceInfosResponse) |  |
| CreateGood | [CreateGoodRequest](#cloud.hashing.goods.v1.CreateGoodRequest) | [CreateGoodResponse](#cloud.hashing.goods.v1.CreateGoodResponse) | Good information for CRUD |
| UpdateGood | [UpdateGoodRequest](#cloud.hashing.goods.v1.UpdateGoodRequest) | [UpdateGoodResponse](#cloud.hashing.goods.v1.UpdateGoodResponse) |  |
| GetGood | [GetGoodRequest](#cloud.hashing.goods.v1.GetGoodRequest) | [GetGoodResponse](#cloud.hashing.goods.v1.GetGoodResponse) |  |
| DeleteGood | [DeleteGoodRequest](#cloud.hashing.goods.v1.DeleteGoodRequest) | [DeleteGoodResponse](#cloud.hashing.goods.v1.DeleteGoodResponse) |  |
| GetGoods | [GetGoodsRequest](#cloud.hashing.goods.v1.GetGoodsRequest) | [GetGoodsResponse](#cloud.hashing.goods.v1.GetGoodsResponse) |  |
| GetGoodDetail | [GetGoodDetailRequest](#cloud.hashing.goods.v1.GetGoodDetailRequest) | [GetGoodDetailResponse](#cloud.hashing.goods.v1.GetGoodDetailResponse) | Good information for API |
| GetGoodsDetail | [GetGoodsDetailRequest](#cloud.hashing.goods.v1.GetGoodsDetailRequest) | [GetGoodsDetailResponse](#cloud.hashing.goods.v1.GetGoodsDetailResponse) |  |
| CreateRecommend | [CreateRecommendRequest](#cloud.hashing.goods.v1.CreateRecommendRequest) | [CreateRecommendResponse](#cloud.hashing.goods.v1.CreateRecommendResponse) |  |
| UpdateRecommend | [UpdateRecommendRequest](#cloud.hashing.goods.v1.UpdateRecommendRequest) | [UpdateRecommendResponse](#cloud.hashing.goods.v1.UpdateRecommendResponse) |  |
| GetRecommendsByApp | [GetRecommendsByAppRequest](#cloud.hashing.goods.v1.GetRecommendsByAppRequest) | [GetRecommendsByAppResponse](#cloud.hashing.goods.v1.GetRecommendsByAppResponse) |  |
| GetRecommendsByRecommender | [GetRecommendsByRecommenderRequest](#cloud.hashing.goods.v1.GetRecommendsByRecommenderRequest) | [GetRecommendsByRecommenderResponse](#cloud.hashing.goods.v1.GetRecommendsByRecommenderResponse) |  |
| DeleteRecommend | [DeleteRecommendRequest](#cloud.hashing.goods.v1.DeleteRecommendRequest) | [DeleteRecommendResponse](#cloud.hashing.goods.v1.DeleteRecommendResponse) |  |
| GetRecommendGoodsByApp | [GetRecommendGoodsByAppRequest](#cloud.hashing.goods.v1.GetRecommendGoodsByAppRequest) | [GetRecommendGoodsByAppResponse](#cloud.hashing.goods.v1.GetRecommendGoodsByAppResponse) |  |
| GetRecommendGoodsByRecommender | [GetRecommendGoodsByRecommenderRequest](#cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderRequest) | [GetRecommendGoodsByRecommenderResponse](#cloud.hashing.goods.v1.GetRecommendGoodsByRecommenderResponse) |  |
| AuthorizeAppGood | [AuthorizeAppGoodRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodRequest) | [AuthorizeAppGoodResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodResponse) |  |
| SetAppGoodPrice | [SetAppGoodPriceRequest](#cloud.hashing.goods.v1.SetAppGoodPriceRequest) | [SetAppGoodPriceResponse](#cloud.hashing.goods.v1.SetAppGoodPriceResponse) |  |
| CheckAppGood | [CheckAppGoodRequest](#cloud.hashing.goods.v1.CheckAppGoodRequest) | [CheckAppGoodResponse](#cloud.hashing.goods.v1.CheckAppGoodResponse) |  |
| OnsaleAppGood | [OnsaleAppGoodRequest](#cloud.hashing.goods.v1.OnsaleAppGoodRequest) | [OnsaleAppGoodResponse](#cloud.hashing.goods.v1.OnsaleAppGoodResponse) |  |
| OffsaleAppGood | [OffsaleAppGoodRequest](#cloud.hashing.goods.v1.OffsaleAppGoodRequest) | [OffsaleAppGoodResponse](#cloud.hashing.goods.v1.OffsaleAppGoodResponse) |  |
| UnauthorizeAppGood | [UnauthorizeAppGoodRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodRequest) | [UnauthorizeAppGoodResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodResponse) |  |
| AuthorizeAppTargetArea | [AuthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaRequest) | [AuthorizeAppTargetAreaResponse](#cloud.hashing.goods.v1.AuthorizeAppTargetAreaResponse) |  |
| CheckAppTargetArea | [CheckAppTargetAreaRequest](#cloud.hashing.goods.v1.CheckAppTargetAreaRequest) | [CheckAppTargetAreaResponse](#cloud.hashing.goods.v1.CheckAppTargetAreaResponse) |  |
| UnauthorizeAppTargetArea | [UnauthorizeAppTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaRequest) | [UnauthorizeAppTargetAreaResponse](#cloud.hashing.goods.v1.UnauthorizeAppTargetAreaResponse) |  |
| AuthorizeAppGoodTargetArea | [AuthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaRequest) | [AuthorizeAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.AuthorizeAppGoodTargetAreaResponse) |  |
| CheckAppGoodTargetArea | [CheckAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.CheckAppGoodTargetAreaRequest) | [CheckAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.CheckAppGoodTargetAreaResponse) |  |
| UnauthorizeAppGoodTargetArea | [UnauthorizeAppGoodTargetAreaRequest](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaRequest) | [UnauthorizeAppGoodTargetAreaResponse](#cloud.hashing.goods.v1.UnauthorizeAppGoodTargetAreaResponse) |  |
| CreateGoodComment | [CreateGoodCommentRequest](#cloud.hashing.goods.v1.CreateGoodCommentRequest) | [CreateGoodCommentResponse](#cloud.hashing.goods.v1.CreateGoodCommentResponse) |  |
| UpdateGoodComment | [UpdateGoodCommentRequest](#cloud.hashing.goods.v1.UpdateGoodCommentRequest) | [UpdateGoodCommentResponse](#cloud.hashing.goods.v1.UpdateGoodCommentResponse) |  |
| GetGoodComments | [GetGoodCommentsRequest](#cloud.hashing.goods.v1.GetGoodCommentsRequest) | [GetGoodCommentsResponse](#cloud.hashing.goods.v1.GetGoodCommentsResponse) |  |
| CreateGoodExtraInfo | [CreateGoodExtraInfoRequest](#cloud.hashing.goods.v1.CreateGoodExtraInfoRequest) | [CreateGoodExtraInfoRequest](#cloud.hashing.goods.v1.CreateGoodExtraInfoRequest) |  |
| GetGoodExtraInfo | [GetGoodExtraInfoRequest](#cloud.hashing.goods.v1.GetGoodExtraInfoRequest) | [GetGoodExtraInfoRequest](#cloud.hashing.goods.v1.GetGoodExtraInfoRequest) |  |
| UpdateGoodExtraInfo | [UpdateGoodExtraInfoRequest](#cloud.hashing.goods.v1.UpdateGoodExtraInfoRequest) | [UpdateGoodExtraInfoRequest](#cloud.hashing.goods.v1.UpdateGoodExtraInfoRequest) |  |
| CreateGoodReview | [CreateGoodReviewRequest](#cloud.hashing.goods.v1.CreateGoodReviewRequest) | [CreateGoodReviewResponse](#cloud.hashing.goods.v1.CreateGoodReviewResponse) |  |
| UpdateGoodReview | [UpdateGoodReviewRequest](#cloud.hashing.goods.v1.UpdateGoodReviewRequest) | [UpdateGoodReviewResponse](#cloud.hashing.goods.v1.UpdateGoodReviewResponse) |  |
| GetGoodReview | [GetGoodReviewRequest](#cloud.hashing.goods.v1.GetGoodReviewRequest) | [GetGoodReviewResponse](#cloud.hashing.goods.v1.GetGoodReviewResponse) |  |
| CreateFee | [CreateFeeRequest](#cloud.hashing.goods.v1.CreateFeeRequest) | [CreateFeeResponse](#cloud.hashing.goods.v1.CreateFeeResponse) |  |
| GetFee | [GetFeeRequest](#cloud.hashing.goods.v1.GetFeeRequest) | [GetFeeResponse](#cloud.hashing.goods.v1.GetFeeResponse) |  |
| GetFees | [GetFeesRequest](#cloud.hashing.goods.v1.GetFeesRequest) | [GetFeesResponse](#cloud.hashing.goods.v1.GetFeesResponse) |  |
| CreateFeeType | [CreateFeeTypeRequest](#cloud.hashing.goods.v1.CreateFeeTypeRequest) | [CreateFeeTypeResponse](#cloud.hashing.goods.v1.CreateFeeTypeResponse) |  |
| UpdateFeeType | [UpdateFeeTypeRequest](#cloud.hashing.goods.v1.UpdateFeeTypeRequest) | [UpdateFeeTypeResponse](#cloud.hashing.goods.v1.UpdateFeeTypeResponse) |  |
| GetFeeType | [GetFeeTypeRequest](#cloud.hashing.goods.v1.GetFeeTypeRequest) | [GetFeeTypeResponse](#cloud.hashing.goods.v1.GetFeeTypeResponse) |  |
| GetFeeTypes | [GetFeeTypesRequest](#cloud.hashing.goods.v1.GetFeeTypesRequest) | [GetFeeTypesResponse](#cloud.hashing.goods.v1.GetFeeTypesResponse) |  |

 



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

