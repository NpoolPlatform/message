/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type PageInfo = {
  PageIndex?: number
  PageSize?: number
}

export type VersionResponse = {
  Info?: string
}

export type VendorLocationInfo = {
  ID?: string
  Country?: string
  Province?: string
  City?: string
  Address?: string
}

export type CreateVendorLocationRequest = {
  Info?: VendorLocationInfo
}

export type CreateVendorLocationResponse = {
  Info?: VendorLocationInfo
}

export type GetVendorLocationRequest = {
  ID?: string
}

export type GetVendorLocationResponse = {
  Info?: VendorLocationInfo
}

export type UpdateVendorLocationRequest = {
  Info?: VendorLocationInfo
}

export type UpdateVendorLocationResponse = {
  Info?: VendorLocationInfo
}

export type DeleteVendorLocationRequest = {
  ID?: string
}

export type DeleteVendorLocationResponse = {
  Info?: VendorLocationInfo
}

export type GetVendorLocationsRequest = {
  PageInfo?: PageInfo
}

export type GetVendorLocationsResponse = {
  Infos?: VendorLocationInfo[]
  Total?: number
}

export type TargetAreaInfo = {
  ID?: string
  Continent?: string
  Country?: string
}

export type CreateTargetAreaRequest = {
  Info?: TargetAreaInfo
}

export type CreateTargetAreaResponse = {
  Info?: TargetAreaInfo
}

export type GetTargetAreaRequest = {
  ID?: string
}

export type GetTargetAreaResponse = {
  Info?: TargetAreaInfo
}

export type UpdateTargetAreaRequest = {
  Info?: TargetAreaInfo
}

export type UpdateTargetAreaResponse = {
  Info?: TargetAreaInfo
}

export type DeleteTargetAreaRequest = {
  ID?: string
}

export type DeleteTargetAreaResponse = {
  Info?: TargetAreaInfo
}

export type DeleteTargetAreaByContinentCountryRequest = {
  Continent?: string
  Country?: string
}

export type DeleteTargetAreaByContinentCountryResponse = {
  Info?: TargetAreaInfo
}

export type GetTargetAreasRequest = {
  PageInfo?: PageInfo
}

export type GetTargetAreasResponse = {
  Infos?: TargetAreaInfo[]
  Total?: number
}

export type DeviceInfo = {
  ID?: string
  Type?: string
  Manufacturer?: string
  PowerComsuption?: number
  ShipmentAt?: number
}

export type CreateDeviceInfoRequest = {
  Info?: DeviceInfo
}

export type CreateDeviceInfoResponse = {
  Info?: DeviceInfo
}

export type UpdateDeviceInfoRequest = {
  Info?: DeviceInfo
}

export type UpdateDeviceInfoResponse = {
  Info?: DeviceInfo
}

export type GetDeviceInfoRequest = {
  ID?: string
}

export type GetDeviceInfoResponse = {
  Info?: DeviceInfo
}

export type DeleteDeviceInfoRequest = {
  ID?: string
}

export type DeleteDeviceInfoResponse = {
  Info?: DeviceInfo
}

export type GetDeviceInfosRequest = {
  PageInfo?: PageInfo
}

export type GetDeviceInfosResponse = {
  Infos?: DeviceInfo[]
  Total?: number
}

export type GoodInfo = {
  ID?: string
  DeviceInfoID?: string
  SeparateFee?: boolean
  UnitPower?: number
  DurationDays?: number
  CoinInfoID?: string
  Actuals?: boolean
  DeliveryAt?: number
  InheritFromGoodID?: string
  VendorLocationID?: string
  Price?: number
  BenefitType?: string
  Classic?: boolean
  SupportCoinTypeIDs?: string[]
  Total?: number
  PriceCurrency?: string
  Title?: string
  Unit?: string
  FeeIDs?: string[]
}

export type CreateGoodRequest = {
  Info?: GoodInfo
}

export type CreateGoodResponse = {
  Info?: GoodInfo
}

export type UpdateGoodRequest = {
  Info?: GoodInfo
}

export type UpdateGoodResponse = {
  Info?: GoodInfo
}

export type GetGoodRequest = {
  ID?: string
}

export type GetGoodResponse = {
  Info?: GoodInfo
}

export type PriceCurrency = {
  ID?: string
  Name?: string
  Unit?: string
  Symbol?: string
}

export type CreatePriceCurrencyRequest = {
  Info?: PriceCurrency
}

export type CreatePriceCurrencyResponse = {
  Info?: PriceCurrency
}

export type UpdatePriceCurrencyRequest = {
  Info?: PriceCurrency
}

export type UpdatePriceCurrencyResponse = {
  Info?: PriceCurrency
}

export type GetPriceCurrencyRequest = {
  ID?: string
}

export type GetPriceCurrencyResponse = {
  Info?: PriceCurrency
}

export type GetPriceCurrencysRequest = {
  PageInfo?: PageInfo
}

export type GetPriceCurrencysResponse = {
  Infos?: PriceCurrency[]
}

export type FeeDetail = {
  ID?: string
  AppID?: string
  Fee?: FeeType
  Value?: number
}

export type GoodDetail = {
  ID?: string
  DeviceInfo?: DeviceInfo
  SeparateFee?: boolean
  UnitPower?: number
  DurationDays?: number
  CoinInfoID?: string
  Actuals?: boolean
  DeliveryAt?: number
  InheritFromGood?: GoodInfo
  VendorLocation?: VendorLocationInfo
  Price?: number
  BenefitType?: string
  Classic?: boolean
  SupportCoinTypeIDs?: string[]
  Total?: number
  Extra?: GoodExtraInfo
  Title?: string
  Unit?: string
  PriceCurrency?: PriceCurrency
  Fees?: FeeDetail[]
}

export type GetGoodDetailRequest = {
  ID?: string
}

export type GetGoodDetailResponse = {
  Detail?: GoodDetail
}

export type GetGoodsDetailRequest = {
  PageInfo?: PageInfo
}

export type GetGoodsDetailResponse = {
  Details?: GoodDetail[]
  Total?: number
}

export type DeleteGoodRequest = {
  ID?: string
}

export type DeleteGoodResponse = {
  Info?: GoodInfo
}

export type GetGoodsRequest = {
  PageInfo?: PageInfo
}

export type GetGoodsResponse = {
  Infos?: GoodInfo[]
  Total?: number
}

export type Recommend = {
  ID?: string
  AppID?: string
  GoodID?: string
  RecommenderID?: string
  Message?: string
}

export type CreateRecommendRequest = {
  Info?: Recommend
}

export type CreateRecommendResponse = {
  Info?: Recommend
}

export type UpdateRecommendRequest = {
  Info?: Recommend
}

export type UpdateRecommendResponse = {
  Info?: Recommend
}

export type GetRecommendsByAppRequest = {
  AppID?: string
}

export type GetRecommendsByAppResponse = {
  Infos?: Recommend[]
}

export type GetRecommendsByRecommenderRequest = {
  UserID?: string
}

export type GetRecommendsByRecommenderResponse = {
  Infos?: Recommend[]
}

export type DeleteRecommendRequest = {
  ID?: string
}

export type DeleteRecommendResponse = {
  Info?: Recommend
}

export type RecommendGood = {
  Recommend?: Recommend
  Good?: GoodDetail
}

export type GetRecommendGoodsByAppRequest = {
  AppID?: string
}

export type GetRecommendGoodsByAppResponse = {
  Infos?: RecommendGood[]
}

export type GetRecommendGoodsByRecommenderRequest = {
  UserID?: string
}

export type GetRecommendGoodsByRecommenderResponse = {
  Infos?: RecommendGood[]
}

export type AppGoodInfo = {
  ID?: string
  AppID?: string
  GoodID?: string
  Price?: number
  Authorized?: boolean
  Online?: boolean
  InitAreaStrategy?: string
}

export type SetAppGoodPriceRequest = {
  Info?: AppGoodInfo
}

export type SetAppGoodPriceResponse = {
  Info?: AppGoodInfo
}

export type CheckAppGoodRequest = {
  Info?: AppGoodInfo
}

export type CheckAppGoodResponse = {
  Info?: AppGoodInfo
}

export type OnsaleAppGoodRequest = {
  Info?: AppGoodInfo
}

export type OnsaleAppGoodResponse = {
  Info?: AppGoodInfo
}

export type OffsaleAppGoodRequest = {
  Info?: AppGoodInfo
}

export type OffsaleAppGoodResponse = {
  Info?: AppGoodInfo
}

export type AuthorizeAppGoodRequest = {
  Info?: AppGoodInfo
}

export type AuthorizeAppGoodResponse = {
  Info?: AppGoodInfo
}

export type UnauthorizeAppGoodRequest = {
  Info?: AppGoodInfo
}

export type UnauthorizeAppGoodResponse = {
  Info?: AppGoodInfo
}

export type AppTargetAreaInfo = {
  ID?: string
  AppID?: string
  TargetAreaID?: string
}

export type AuthorizeAppTargetAreaRequest = {
  Info?: AppTargetAreaInfo
}

export type AuthorizeAppTargetAreaResponse = {
  Info?: AppTargetAreaInfo
}

export type CheckAppTargetAreaRequest = {
  Info?: AppTargetAreaInfo
}

export type CheckAppTargetAreaResponse = {
  Info?: AppTargetAreaInfo
  Authorized?: boolean
}

export type UnauthorizeAppTargetAreaRequest = {
  ID?: string
}

export type UnauthorizeAppTargetAreaResponse = {
  Info?: AppTargetAreaInfo
  Authorized?: boolean
}

export type AppGoodTargetAreaInfo = {
  ID?: string
  AppID?: string
  GoodID?: string
  TargetAreaID?: string
}

export type AuthorizeAppGoodTargetAreaRequest = {
  Info?: AppGoodTargetAreaInfo
}

export type AuthorizeAppGoodTargetAreaResponse = {
  Info?: AppGoodTargetAreaInfo
}

export type CheckAppGoodTargetAreaRequest = {
  Info?: AppGoodTargetAreaInfo
}

export type CheckAppGoodTargetAreaResponse = {
  Info?: AppGoodTargetAreaInfo
}

export type UnauthorizeAppGoodTargetAreaRequest = {
  Info?: AppGoodTargetAreaInfo
}

export type UnauthorizeAppGoodTargetAreaResponse = {
  Info?: AppGoodTargetAreaInfo
}

export type GoodComment = {
  ID?: string
  AppID?: string
  UserID?: string
  ReplyToID?: string
  GoodID?: string
  OrderID?: string
  Content?: string
}

export type CreateGoodCommentRequest = {
  Comment?: GoodComment
}

export type CreateGoodCommentResponse = {
  Comment?: GoodComment
}

export type UpdateGoodCommentRequest = {
  Comment?: GoodComment
}

export type UpdateGoodCommentResponse = {
  Comment?: GoodComment
}

export type GetGoodCommentsRequest = {
  PageInfo?: PageInfo
  GoodID?: string
}

export type GetGoodCommentsResponse = {
  Comments?: GoodComment[]
  Total?: number
}

export type GoodExtraInfo = {
  ID?: string
  GoodID?: string
  Posters?: string[]
  Labels?: string[]
  OutSale?: boolean
  PreSale?: boolean
  VoteCount?: number
  Rating?: number
}

export type CreateGoodExtraInfoRequest = {
  Info?: GoodExtraInfo
}

export type CreateGoodExtraInfoResponse = {
  Info?: GoodExtraInfo
}

export type GetGoodExtraInfoRequest = {
  GoodID?: string
}

export type GetGoodExtraInfoResponse = {
  Info?: GoodExtraInfo
}

export type UpdateGoodExtraInfoRequest = {
  Info?: GoodExtraInfo
}

export type UpdateGoodExtraInfoResponse = {
  Info?: GoodExtraInfo
}

export type GoodReviewInfo = {
  ID?: string
  Type?: string
  ReviewerID?: string
  State?: string
  Message?: string
  ReviewedID?: string
}

export type CreateGoodReviewRequest = {
  Info?: GoodReviewInfo
}

export type CreateGoodReviewResponse = {
  Info?: GoodReviewInfo
}

export type UpdateGoodReviewRequest = {
  Info?: GoodReviewInfo
}

export type UpdateGoodReviewResponse = {
  Info?: GoodReviewInfo
}

export type GetGoodReviewRequest = {
  Info?: GoodReviewInfo
}

export type GetGoodReviewResponse = {
  Info?: GoodReviewInfo
}

export type FeeType = {
  ID?: string
  FeeType?: string
  FeeDescription?: string
  PayType?: string
}

export type CreateFeeTypeRequest = {
  Info?: FeeType
}

export type CreateFeeTypeResponse = {
  Info?: FeeType
}

export type UpdateFeeTypeRequest = {
  Info?: FeeType
}

export type UpdateFeeTypeResponse = {
  Info?: FeeType
}

export type GetFeeTypeRequest = {
  ID?: string
}

export type GetFeeTypeResponse = {
  Info?: FeeType
}

export type GetFeeTypesRequest = {
  PageInfo?: PageInfo
}

export type GetFeeTypesResponse = {
  Infos?: FeeType[]
}

export type Fee = {
  ID?: string
  AppID?: string
  FeeTypeID?: string
  Value?: number
}

export type CreateFeeRequest = {
  Info?: Fee
}

export type CreateFeeResponse = {
  Info?: Fee
}

export type GetFeeRequest = {
  ID?: string
}

export type GetFeeResponse = {
  Info?: Fee
}

export type GetFeesRequest = {
  PageInfo?: PageInfo
}

export type GetFeesResponse = {
  Infos?: Fee[]
}

export class CloudHashingGoods {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static CreateVendorLocation(req: CreateVendorLocationRequest, initReq?: fm.InitReq): Promise<CreateVendorLocationResponse> {
    return fm.fetchReq<CreateVendorLocationRequest, CreateVendorLocationResponse>(`/v1/create/vendor-location`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateVendorLocation(req: UpdateVendorLocationRequest, initReq?: fm.InitReq): Promise<UpdateVendorLocationResponse> {
    return fm.fetchReq<UpdateVendorLocationRequest, UpdateVendorLocationResponse>(`/v1/update/vendor-location`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetVendorLocation(req: GetVendorLocationRequest, initReq?: fm.InitReq): Promise<GetVendorLocationResponse> {
    return fm.fetchReq<GetVendorLocationRequest, GetVendorLocationResponse>(`/v1/get/vendor-location`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteVendorLocation(req: DeleteVendorLocationRequest, initReq?: fm.InitReq): Promise<DeleteVendorLocationResponse> {
    return fm.fetchReq<DeleteVendorLocationRequest, DeleteVendorLocationResponse>(`/v1/delete/vendor-location`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetVendorLocations(req: GetVendorLocationsRequest, initReq?: fm.InitReq): Promise<GetVendorLocationsResponse> {
    return fm.fetchReq<GetVendorLocationsRequest, GetVendorLocationsResponse>(`/v1/get/vendor-locations`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreatePriceCurrency(req: CreatePriceCurrencyRequest, initReq?: fm.InitReq): Promise<CreatePriceCurrencyResponse> {
    return fm.fetchReq<CreatePriceCurrencyRequest, CreatePriceCurrencyResponse>(`/v1/create/price-currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdatePriceCurrency(req: UpdatePriceCurrencyRequest, initReq?: fm.InitReq): Promise<UpdatePriceCurrencyResponse> {
    return fm.fetchReq<UpdatePriceCurrencyRequest, UpdatePriceCurrencyResponse>(`/v1/update/price-currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPriceCurrency(req: GetPriceCurrencyRequest, initReq?: fm.InitReq): Promise<GetPriceCurrencyResponse> {
    return fm.fetchReq<GetPriceCurrencyRequest, GetPriceCurrencyResponse>(`/v1/get/price-currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPriceCurrencys(req: GetPriceCurrencysRequest, initReq?: fm.InitReq): Promise<GetPriceCurrencysResponse> {
    return fm.fetchReq<GetPriceCurrencysRequest, GetPriceCurrencysResponse>(`/v1/get/price-currencys`, {...initReq, method: "POST"})
  }
  static CreateTargetArea(req: CreateTargetAreaRequest, initReq?: fm.InitReq): Promise<CreateTargetAreaResponse> {
    return fm.fetchReq<CreateTargetAreaRequest, CreateTargetAreaResponse>(`/v1/create/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateTargetArea(req: UpdateTargetAreaRequest, initReq?: fm.InitReq): Promise<UpdateTargetAreaResponse> {
    return fm.fetchReq<UpdateTargetAreaRequest, UpdateTargetAreaResponse>(`/v1/update/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTargetArea(req: GetTargetAreaRequest, initReq?: fm.InitReq): Promise<GetTargetAreaResponse> {
    return fm.fetchReq<GetTargetAreaRequest, GetTargetAreaResponse>(`/v1/get/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteTargetArea(req: DeleteTargetAreaRequest, initReq?: fm.InitReq): Promise<DeleteTargetAreaResponse> {
    return fm.fetchReq<DeleteTargetAreaRequest, DeleteTargetAreaResponse>(`/v1/delete/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteTargetAreaByContinentCountry(req: DeleteTargetAreaByContinentCountryRequest, initReq?: fm.InitReq): Promise<DeleteTargetAreaByContinentCountryResponse> {
    return fm.fetchReq<DeleteTargetAreaByContinentCountryRequest, DeleteTargetAreaByContinentCountryResponse>(`/v1/delete/target-area/continent/country`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTargetAreas(req: GetTargetAreasRequest, initReq?: fm.InitReq): Promise<GetTargetAreasResponse> {
    return fm.fetchReq<GetTargetAreasRequest, GetTargetAreasResponse>(`/v1/get/target-areas`, {...initReq, method: "POST"})
  }
  static CreateDeviceInfo(req: CreateDeviceInfoRequest, initReq?: fm.InitReq): Promise<CreateDeviceInfoResponse> {
    return fm.fetchReq<CreateDeviceInfoRequest, CreateDeviceInfoResponse>(`/v1/create/device`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateDeviceInfo(req: UpdateDeviceInfoRequest, initReq?: fm.InitReq): Promise<UpdateDeviceInfoResponse> {
    return fm.fetchReq<UpdateDeviceInfoRequest, UpdateDeviceInfoResponse>(`/v1/update/device`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDeviceInfo(req: GetDeviceInfoRequest, initReq?: fm.InitReq): Promise<GetDeviceInfoResponse> {
    return fm.fetchReq<GetDeviceInfoRequest, GetDeviceInfoResponse>(`/v1/get/device`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteDeviceInfo(req: DeleteDeviceInfoRequest, initReq?: fm.InitReq): Promise<DeleteDeviceInfoResponse> {
    return fm.fetchReq<DeleteDeviceInfoRequest, DeleteDeviceInfoResponse>(`/v1/delete/device`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDeviceInfos(req: GetDeviceInfosRequest, initReq?: fm.InitReq): Promise<GetDeviceInfosResponse> {
    return fm.fetchReq<GetDeviceInfosRequest, GetDeviceInfosResponse>(`/v1/get/devices`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGood(req: CreateGoodRequest, initReq?: fm.InitReq): Promise<CreateGoodResponse> {
    return fm.fetchReq<CreateGoodRequest, CreateGoodResponse>(`/v1/create/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGood(req: UpdateGoodRequest, initReq?: fm.InitReq): Promise<UpdateGoodResponse> {
    return fm.fetchReq<UpdateGoodRequest, UpdateGoodResponse>(`/v1/update/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGood(req: GetGoodRequest, initReq?: fm.InitReq): Promise<GetGoodResponse> {
    return fm.fetchReq<GetGoodRequest, GetGoodResponse>(`/v1/get/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteGood(req: DeleteGoodRequest, initReq?: fm.InitReq): Promise<DeleteGoodResponse> {
    return fm.fetchReq<DeleteGoodRequest, DeleteGoodResponse>(`/v1/delete/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoods(req: GetGoodsRequest, initReq?: fm.InitReq): Promise<GetGoodsResponse> {
    return fm.fetchReq<GetGoodsRequest, GetGoodsResponse>(`/v1/get/goods`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodDetail(req: GetGoodDetailRequest, initReq?: fm.InitReq): Promise<GetGoodDetailResponse> {
    return fm.fetchReq<GetGoodDetailRequest, GetGoodDetailResponse>(`/v1/get/good/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodsDetail(req: GetGoodsDetailRequest, initReq?: fm.InitReq): Promise<GetGoodsDetailResponse> {
    return fm.fetchReq<GetGoodsDetailRequest, GetGoodsDetailResponse>(`/v1/get/goods/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateRecommend(req: CreateRecommendRequest, initReq?: fm.InitReq): Promise<CreateRecommendResponse> {
    return fm.fetchReq<CreateRecommendRequest, CreateRecommendResponse>(`/v1/create/recommend`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateRecommend(req: UpdateRecommendRequest, initReq?: fm.InitReq): Promise<UpdateRecommendResponse> {
    return fm.fetchReq<UpdateRecommendRequest, UpdateRecommendResponse>(`/v1/update/recommend`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRecommendsByApp(req: GetRecommendsByAppRequest, initReq?: fm.InitReq): Promise<GetRecommendsByAppResponse> {
    return fm.fetchReq<GetRecommendsByAppRequest, GetRecommendsByAppResponse>(`/v1/get/recommends/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRecommendsByRecommender(req: GetRecommendsByRecommenderRequest, initReq?: fm.InitReq): Promise<GetRecommendsByRecommenderResponse> {
    return fm.fetchReq<GetRecommendsByRecommenderRequest, GetRecommendsByRecommenderResponse>(`/v1/get/recommends/by/recommender`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteRecommend(req: DeleteRecommendRequest, initReq?: fm.InitReq): Promise<DeleteRecommendResponse> {
    return fm.fetchReq<DeleteRecommendRequest, DeleteRecommendResponse>(`/v1/delete/recommend`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRecommendGoodsByApp(req: GetRecommendGoodsByAppRequest, initReq?: fm.InitReq): Promise<GetRecommendGoodsByAppResponse> {
    return fm.fetchReq<GetRecommendGoodsByAppRequest, GetRecommendGoodsByAppResponse>(`/v1/get/recommend/goods/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRecommendGoodsByRecommender(req: GetRecommendGoodsByRecommenderRequest, initReq?: fm.InitReq): Promise<GetRecommendGoodsByRecommenderResponse> {
    return fm.fetchReq<GetRecommendGoodsByRecommenderRequest, GetRecommendGoodsByRecommenderResponse>(`/v1/get/recommend/goods/by/recommender`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthorizeAppGood(req: AuthorizeAppGoodRequest, initReq?: fm.InitReq): Promise<AuthorizeAppGoodResponse> {
    return fm.fetchReq<AuthorizeAppGoodRequest, AuthorizeAppGoodResponse>(`/v1/authorize/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SetAppGoodPrice(req: SetAppGoodPriceRequest, initReq?: fm.InitReq): Promise<SetAppGoodPriceResponse> {
    return fm.fetchReq<SetAppGoodPriceRequest, SetAppGoodPriceResponse>(`/v1/set/app/good/price`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CheckAppGood(req: CheckAppGoodRequest, initReq?: fm.InitReq): Promise<CheckAppGoodResponse> {
    return fm.fetchReq<CheckAppGoodRequest, CheckAppGoodResponse>(`/v1/check/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static OnsaleAppGood(req: OnsaleAppGoodRequest, initReq?: fm.InitReq): Promise<OnsaleAppGoodResponse> {
    return fm.fetchReq<OnsaleAppGoodRequest, OnsaleAppGoodResponse>(`/v1/onsale/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static OffsaleAppGood(req: OffsaleAppGoodRequest, initReq?: fm.InitReq): Promise<OffsaleAppGoodResponse> {
    return fm.fetchReq<OffsaleAppGoodRequest, OffsaleAppGoodResponse>(`/v1/offsale/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnauthorizeAppGood(req: UnauthorizeAppGoodRequest, initReq?: fm.InitReq): Promise<UnauthorizeAppGoodResponse> {
    return fm.fetchReq<UnauthorizeAppGoodRequest, UnauthorizeAppGoodResponse>(`/v1/unauthorize/app/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthorizeAppTargetArea(req: AuthorizeAppTargetAreaRequest, initReq?: fm.InitReq): Promise<AuthorizeAppTargetAreaResponse> {
    return fm.fetchReq<AuthorizeAppTargetAreaRequest, AuthorizeAppTargetAreaResponse>(`/v1/authorize/app/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CheckAppTargetArea(req: CheckAppTargetAreaRequest, initReq?: fm.InitReq): Promise<CheckAppTargetAreaResponse> {
    return fm.fetchReq<CheckAppTargetAreaRequest, CheckAppTargetAreaResponse>(`/v1/check/app/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnauthorizeAppTargetArea(req: UnauthorizeAppTargetAreaRequest, initReq?: fm.InitReq): Promise<UnauthorizeAppTargetAreaResponse> {
    return fm.fetchReq<UnauthorizeAppTargetAreaRequest, UnauthorizeAppTargetAreaResponse>(`/v1/unauthorize/app/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthorizeAppGoodTargetArea(req: AuthorizeAppGoodTargetAreaRequest, initReq?: fm.InitReq): Promise<AuthorizeAppGoodTargetAreaResponse> {
    return fm.fetchReq<AuthorizeAppGoodTargetAreaRequest, AuthorizeAppGoodTargetAreaResponse>(`/v1/authorize/app/good/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CheckAppGoodTargetArea(req: CheckAppGoodTargetAreaRequest, initReq?: fm.InitReq): Promise<CheckAppGoodTargetAreaResponse> {
    return fm.fetchReq<CheckAppGoodTargetAreaRequest, CheckAppGoodTargetAreaResponse>(`/v1/check/app/good/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnauthorizeAppGoodTargetArea(req: UnauthorizeAppGoodTargetAreaRequest, initReq?: fm.InitReq): Promise<UnauthorizeAppGoodTargetAreaResponse> {
    return fm.fetchReq<UnauthorizeAppGoodTargetAreaRequest, UnauthorizeAppGoodTargetAreaResponse>(`/v1/unauthorize/app/good/target-area`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodComment(req: CreateGoodCommentRequest, initReq?: fm.InitReq): Promise<CreateGoodCommentResponse> {
    return fm.fetchReq<CreateGoodCommentRequest, CreateGoodCommentResponse>(`/v1/create/good/comment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodComment(req: UpdateGoodCommentRequest, initReq?: fm.InitReq): Promise<UpdateGoodCommentResponse> {
    return fm.fetchReq<UpdateGoodCommentRequest, UpdateGoodCommentResponse>(`/v1/update/good/comment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodComments(req: GetGoodCommentsRequest, initReq?: fm.InitReq): Promise<GetGoodCommentsResponse> {
    return fm.fetchReq<GetGoodCommentsRequest, GetGoodCommentsResponse>(`/v1/get/good/comments`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodExtraInfo(req: CreateGoodExtraInfoRequest, initReq?: fm.InitReq): Promise<CreateGoodExtraInfoRequest> {
    return fm.fetchReq<CreateGoodExtraInfoRequest, CreateGoodExtraInfoRequest>(`/v1/create/good/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodExtraInfo(req: GetGoodExtraInfoRequest, initReq?: fm.InitReq): Promise<GetGoodExtraInfoRequest> {
    return fm.fetchReq<GetGoodExtraInfoRequest, GetGoodExtraInfoRequest>(`/v1/get/good/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodExtraInfo(req: UpdateGoodExtraInfoRequest, initReq?: fm.InitReq): Promise<UpdateGoodExtraInfoRequest> {
    return fm.fetchReq<UpdateGoodExtraInfoRequest, UpdateGoodExtraInfoRequest>(`/v1/update/good/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodReview(req: CreateGoodReviewRequest, initReq?: fm.InitReq): Promise<CreateGoodReviewResponse> {
    return fm.fetchReq<CreateGoodReviewRequest, CreateGoodReviewResponse>(`/v1/create/good/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodReview(req: UpdateGoodReviewRequest, initReq?: fm.InitReq): Promise<UpdateGoodReviewResponse> {
    return fm.fetchReq<UpdateGoodReviewRequest, UpdateGoodReviewResponse>(`/v1/update/good/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodReview(req: GetGoodReviewRequest, initReq?: fm.InitReq): Promise<GetGoodReviewResponse> {
    return fm.fetchReq<GetGoodReviewRequest, GetGoodReviewResponse>(`/v1/get/good/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateFee(req: CreateFeeRequest, initReq?: fm.InitReq): Promise<CreateFeeResponse> {
    return fm.fetchReq<CreateFeeRequest, CreateFeeResponse>(`/v1/create/fee`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFee(req: GetFeeRequest, initReq?: fm.InitReq): Promise<GetFeeResponse> {
    return fm.fetchReq<GetFeeRequest, GetFeeResponse>(`/v1/get/fee`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFees(req: GetFeesRequest, initReq?: fm.InitReq): Promise<GetFeesResponse> {
    return fm.fetchReq<GetFeesRequest, GetFeesResponse>(`/v1/get/fees`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateFeeType(req: CreateFeeTypeRequest, initReq?: fm.InitReq): Promise<CreateFeeTypeResponse> {
    return fm.fetchReq<CreateFeeTypeRequest, CreateFeeTypeResponse>(`/v1/create/fee/type`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateFeeType(req: UpdateFeeTypeRequest, initReq?: fm.InitReq): Promise<UpdateFeeTypeResponse> {
    return fm.fetchReq<UpdateFeeTypeRequest, UpdateFeeTypeResponse>(`/v1/update/fee/type`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFeeType(req: GetFeeTypeRequest, initReq?: fm.InitReq): Promise<GetFeeTypeResponse> {
    return fm.fetchReq<GetFeeTypeRequest, GetFeeTypeResponse>(`/v1/get/fee/type`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFeeTypes(req: GetFeeTypesRequest, initReq?: fm.InitReq): Promise<GetFeeTypesResponse> {
    return fm.fetchReq<GetFeeTypesRequest, GetFeeTypesResponse>(`/v1/get/fee/types`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}