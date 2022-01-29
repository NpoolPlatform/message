/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type VendorLocationInfo = {
  id?: string
  country?: string
  province?: string
  city?: string
  address?: string
}

export type CreateVendorLocationRequest = {
  info?: VendorLocationInfo
}

export type CreateVendorLocationResponse = {
  info?: VendorLocationInfo
}

export type GetVendorLocationRequest = {
  id?: string
}

export type GetVendorLocationResponse = {
  info?: VendorLocationInfo
}

export type UpdateVendorLocationRequest = {
  info?: VendorLocationInfo
}

export type UpdateVendorLocationResponse = {
  info?: VendorLocationInfo
}

export type DeleteVendorLocationRequest = {
  id?: string
}

export type DeleteVendorLocationResponse = {
  info?: VendorLocationInfo
}

export type GetVendorLocationsRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetVendorLocationsResponse = {
  infos?: VendorLocationInfo[]
  total?: number
}

export type TargetAreaInfo = {
  id?: string
  continent?: string
  country?: string
}

export type CreateTargetAreaRequest = {
  info?: TargetAreaInfo
}

export type CreateTargetAreaResponse = {
  info?: TargetAreaInfo
}

export type GetTargetAreaRequest = {
  id?: string
}

export type GetTargetAreaResponse = {
  info?: TargetAreaInfo
}

export type UpdateTargetAreaRequest = {
  info?: TargetAreaInfo
}

export type UpdateTargetAreaResponse = {
  info?: TargetAreaInfo
}

export type DeleteTargetAreaRequest = {
  id?: string
}

export type DeleteTargetAreaResponse = {
  info?: TargetAreaInfo
}

export type DeleteTargetAreaByContinentCountryRequest = {
  continent?: string
  country?: string
}

export type DeleteTargetAreaByContinentCountryResponse = {
  info?: TargetAreaInfo
}

export type GetTargetAreasRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetTargetAreasResponse = {
  infos?: TargetAreaInfo[]
  total?: number
}

export type DeviceInfo = {
  id?: string
  type?: string
  manufacturer?: string
  powerComsuption?: number
  shipmentAt?: number
}

export type CreateDeviceInfoRequest = {
  info?: DeviceInfo
}

export type CreateDeviceInfoResponse = {
  info?: DeviceInfo
}

export type UpdateDeviceInfoRequest = {
  info?: DeviceInfo
}

export type UpdateDeviceInfoResponse = {
  info?: DeviceInfo
}

export type GetDeviceInfoRequest = {
  id?: string
}

export type GetDeviceInfoResponse = {
  info?: DeviceInfo
}

export type DeleteDeviceInfoRequest = {
  id?: string
}

export type DeleteDeviceInfoResponse = {
  info?: DeviceInfo
}

export type GetDeviceInfosRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetDeviceInfosResponse = {
  infos?: DeviceInfo[]
  total?: number
}

export type GoodInfo = {
  id?: string
  deviceInfoID?: string
  separateFee?: boolean
  unitPower?: number
  durationDays?: number
  coinInfoID?: string
  actuals?: boolean
  deliveryAt?: number
  inheritFromGoodID?: string
  vendorLocationID?: string
  price?: number
  benefitType?: string
  classic?: boolean
  supportCoinTypeIDs?: string[]
  total?: number
  priceCurrency?: string
  title?: string
  unit?: string
  feeIDs?: string[]
}

export type CreateGoodRequest = {
  info?: GoodInfo
}

export type CreateGoodResponse = {
  info?: GoodInfo
}

export type UpdateGoodRequest = {
  info?: GoodInfo
}

export type UpdateGoodResponse = {
  info?: GoodInfo
}

export type GetGoodRequest = {
  id?: string
}

export type GetGoodResponse = {
  info?: GoodInfo
}

export type PriceCurrency = {
  id?: string
  name?: string
  unit?: string
  symbol?: string
}

export type CreatePriceCurrencyRequest = {
  info?: PriceCurrency
}

export type CreatePriceCurrencyResponse = {
  info?: PriceCurrency
}

export type UpdatePriceCurrencyRequest = {
  info?: PriceCurrency
}

export type UpdatePriceCurrencyResponse = {
  info?: PriceCurrency
}

export type GetPriceCurrencyRequest = {
  id?: string
}

export type GetPriceCurrencyResponse = {
  info?: PriceCurrency
}

export type GetPriceCurrencysRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetPriceCurrencysResponse = {
  infos?: PriceCurrency[]
}

export type FeeDetail = {
  id?: string
  appID?: string
  fee?: FeeType
  value?: number
}

export type GoodDetail = {
  id?: string
  deviceInfo?: DeviceInfo
  separateFee?: boolean
  unitPower?: number
  durationDays?: number
  coinInfoID?: string
  actuals?: boolean
  deliveryAt?: number
  inheritFromGood?: GoodInfo
  vendorLocation?: VendorLocationInfo
  price?: number
  benefitType?: string
  classic?: boolean
  supportCoinTypeIDs?: string[]
  total?: number
  extra?: GoodExtraInfo
  title?: string
  unit?: string
  priceCurrency?: PriceCurrency
  fees?: FeeDetail[]
}

export type GetGoodDetailRequest = {
  id?: string
}

export type GetGoodDetailResponse = {
  detail?: GoodDetail
}

export type GetGoodsDetailRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetGoodsDetailResponse = {
  details?: GoodDetail[]
  total?: number
}

export type DeleteGoodRequest = {
  id?: string
}

export type DeleteGoodResponse = {
  info?: GoodInfo
}

export type GetGoodsRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetGoodsResponse = {
  infos?: GoodInfo[]
  total?: number
}

export type Recommend = {
  id?: string
  appID?: string
  goodID?: string
  recommenderID?: string
  message?: string
}

export type CreateRecommendRequest = {
  info?: Recommend
}

export type CreateRecommendResponse = {
  info?: Recommend
}

export type UpdateRecommendRequest = {
  info?: Recommend
}

export type UpdateRecommendResponse = {
  info?: Recommend
}

export type GetRecommendsByAppRequest = {
  appID?: string
}

export type GetRecommendsByAppResponse = {
  infos?: Recommend[]
}

export type GetRecommendsByRecommenderRequest = {
  userID?: string
}

export type GetRecommendsByRecommenderResponse = {
  infos?: Recommend[]
}

export type DeleteRecommendRequest = {
  id?: string
}

export type DeleteRecommendResponse = {
  info?: Recommend
}

export type RecommendGood = {
  recommend?: Recommend
  good?: GoodDetail
}

export type GetRecommendGoodsByAppRequest = {
  appID?: string
}

export type GetRecommendGoodsByAppResponse = {
  infos?: RecommendGood[]
}

export type GetRecommendGoodsByRecommenderRequest = {
  userID?: string
}

export type GetRecommendGoodsByRecommenderResponse = {
  infos?: RecommendGood[]
}

export type AppGoodInfo = {
  id?: string
  appID?: string
  goodID?: string
  price?: number
  authorized?: boolean
  online?: boolean
  initAreaStrategy?: string
}

export type SetAppGoodPriceRequest = {
  info?: AppGoodInfo
}

export type SetAppGoodPriceResponse = {
  info?: AppGoodInfo
}

export type CheckAppGoodRequest = {
  info?: AppGoodInfo
}

export type CheckAppGoodResponse = {
  info?: AppGoodInfo
}

export type OnsaleAppGoodRequest = {
  info?: AppGoodInfo
}

export type OnsaleAppGoodResponse = {
  info?: AppGoodInfo
}

export type OffsaleAppGoodRequest = {
  info?: AppGoodInfo
}

export type OffsaleAppGoodResponse = {
  info?: AppGoodInfo
}

export type AuthorizeAppGoodRequest = {
  info?: AppGoodInfo
}

export type AuthorizeAppGoodResponse = {
  info?: AppGoodInfo
}

export type UnauthorizeAppGoodRequest = {
  info?: AppGoodInfo
}

export type UnauthorizeAppGoodResponse = {
  info?: AppGoodInfo
}

export type AppTargetAreaInfo = {
  id?: string
  appID?: string
  targetAreaID?: string
}

export type AuthorizeAppTargetAreaRequest = {
  info?: AppTargetAreaInfo
}

export type AuthorizeAppTargetAreaResponse = {
  info?: AppTargetAreaInfo
}

export type CheckAppTargetAreaRequest = {
  info?: AppTargetAreaInfo
}

export type CheckAppTargetAreaResponse = {
  info?: AppTargetAreaInfo
  authorized?: boolean
}

export type UnauthorizeAppTargetAreaRequest = {
  id?: string
}

export type UnauthorizeAppTargetAreaResponse = {
  info?: AppTargetAreaInfo
  authorized?: boolean
}

export type AppGoodTargetAreaInfo = {
  id?: string
  appID?: string
  goodID?: string
  targetAreaID?: string
}

export type AuthorizeAppGoodTargetAreaRequest = {
  info?: AppGoodTargetAreaInfo
}

export type AuthorizeAppGoodTargetAreaResponse = {
  info?: AppGoodTargetAreaInfo
}

export type CheckAppGoodTargetAreaRequest = {
  info?: AppGoodTargetAreaInfo
}

export type CheckAppGoodTargetAreaResponse = {
  info?: AppGoodTargetAreaInfo
}

export type UnauthorizeAppGoodTargetAreaRequest = {
  info?: AppGoodTargetAreaInfo
}

export type UnauthorizeAppGoodTargetAreaResponse = {
  info?: AppGoodTargetAreaInfo
}

export type GoodComment = {
  id?: string
  appID?: string
  userID?: string
  replyToID?: string
  goodID?: string
  orderID?: string
  content?: string
}

export type CreateGoodCommentRequest = {
  comment?: GoodComment
}

export type CreateGoodCommentResponse = {
  comment?: GoodComment
}

export type UpdateGoodCommentRequest = {
  comment?: GoodComment
}

export type UpdateGoodCommentResponse = {
  comment?: GoodComment
}

export type GetGoodCommentsRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
  goodID?: string
}

export type GetGoodCommentsResponse = {
  comments?: GoodComment[]
  total?: number
}

export type GoodExtraInfo = {
  id?: string
  goodID?: string
  posters?: string[]
  labels?: string[]
  outSale?: boolean
  preSale?: boolean
  voteCount?: number
  rating?: number
}

export type CreateGoodExtraInfoRequest = {
  info?: GoodExtraInfo
}

export type CreateGoodExtraInfoResponse = {
  info?: GoodExtraInfo
}

export type GetGoodExtraInfoRequest = {
  goodID?: string
}

export type GetGoodExtraInfoResponse = {
  info?: GoodExtraInfo
}

export type UpdateGoodExtraInfoRequest = {
  info?: GoodExtraInfo
}

export type UpdateGoodExtraInfoResponse = {
  info?: GoodExtraInfo
}

export type GoodReviewInfo = {
  id?: string
  type?: string
  reviewerID?: string
  state?: string
  message?: string
  reviewedID?: string
}

export type CreateGoodReviewRequest = {
  info?: GoodReviewInfo
}

export type CreateGoodReviewResponse = {
  info?: GoodReviewInfo
}

export type UpdateGoodReviewRequest = {
  info?: GoodReviewInfo
}

export type UpdateGoodReviewResponse = {
  info?: GoodReviewInfo
}

export type GetGoodReviewRequest = {
  info?: GoodReviewInfo
}

export type GetGoodReviewResponse = {
  info?: GoodReviewInfo
}

export type FeeType = {
  id?: string
  feeType?: string
  feeDescription?: string
  payType?: string
}

export type CreateFeeTypeRequest = {
  info?: FeeType
}

export type CreateFeeTypeResponse = {
  info?: FeeType
}

export type UpdateFeeTypeRequest = {
  info?: FeeType
}

export type UpdateFeeTypeResponse = {
  info?: FeeType
}

export type GetFeeTypeRequest = {
  id?: string
}

export type GetFeeTypeResponse = {
  info?: FeeType
}

export type GetFeeTypesRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetFeeTypesResponse = {
  infos?: FeeType[]
}

export type Fee = {
  id?: string
  appID?: string
  feeTypeID?: string
  value?: number
}

export type CreateFeeRequest = {
  info?: Fee
}

export type CreateFeeResponse = {
  info?: Fee
}

export type GetFeeRequest = {
  id?: string
}

export type GetFeeResponse = {
  info?: Fee
}

export type GetFeesRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetFeesResponse = {
  infos?: Fee[]
}

export class CloudHashingGoods {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
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