/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinDescription = {
  id?: string
  appID?: string
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
  createAt?: number
  updateAt?: number
}

export type CreateCoinDescriptionRequest = {
  info?: CoinDescription
}

export type CreateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CreateCoinDescriptionsRequest = {
  appID?: string
  infos?: CoinDescription[]
}

export type CreateCoinDescriptionsResponse = {
  infos?: CoinDescription[]
}

export type CreateAppCoinDescriptionRequest = {
  targetAppID?: string
  info?: CoinDescription
}

export type CreateAppCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CreateAppCoinDescriptionsRequest = {
  targetAppID?: string
  infos?: CoinDescription[]
}

export type CreateAppCoinDescriptionsResponse = {
  infos?: CoinDescription[]
}

export type UpdateCoinDescriptionRequest = {
  info?: CoinDescription
}

export type UpdateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionRequest = {
  id?: string
}

export type GetCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionsRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type GetCoinDescriptionOnlyRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinDescriptionOnlyResponse = {
  info?: CoinDescription
}

export type GetAppCoinDescriptionsRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type GetAppCoinDescriptionOnlyRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinDescriptionOnlyResponse = {
  info?: CoinDescription
}

export type DeleteCoinDescriptionRequest = {
  id?: string
}

export type DeleteCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CoinProduct = {
  id?: string
  appID?: string
  coinTypeID?: string
  productPage?: string
}

export type CreateCoinProductRequest = {
  info?: CoinProduct
}

export type CreateCoinProductResponse = {
  info?: CoinProduct
}

export type CreateCoinProductsRequest = {
  appID?: string
  infos?: CoinProduct[]
}

export type CreateCoinProductsResponse = {
  infos?: CoinProduct[]
}

export type CreateAppCoinProductRequest = {
  targetAppID?: string
  info?: CoinProduct
}

export type CreateAppCoinProductResponse = {
  info?: CoinProduct
}

export type CreateAppCoinProductsRequest = {
  targetAppID?: string
  infos?: CoinProduct[]
}

export type CreateAppCoinProductsResponse = {
  infos?: CoinProduct[]
}

export type UpdateCoinProductRequest = {
  info?: CoinProduct
}

export type UpdateCoinProductResponse = {
  info?: CoinProduct
}

export type GetCoinProductRequest = {
  id?: string
}

export type GetCoinProductResponse = {
  info?: CoinProduct
}

export type GetCoinProductsRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinProductsResponse = {
  infos?: CoinProduct[]
  total?: number
}

export type GetCoinProductOnlyRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinProductOnlyResponse = {
  info?: CoinProduct
}

export type GetAppCoinProductsRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinProductsResponse = {
  infos?: CoinProduct[]
  total?: number
}

export type GetAppCoinProductOnlyRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinProductOnlyResponse = {
  info?: CoinProduct
}

export type DeleteCoinProductRequest = {
  id?: string
}

export type DeleteCoinProductResponse = {
  info?: CoinProduct
}

export class ProjectInfoManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinDescription(req: CreateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionResponse> {
    return fm.fetchReq<CreateCoinDescriptionRequest, CreateCoinDescriptionResponse>(`/v1/create/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinDescriptions(req: CreateCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionsResponse> {
    return fm.fetchReq<CreateCoinDescriptionsRequest, CreateCoinDescriptionsResponse>(`/v1/create/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCoinDescription(req: CreateAppCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateAppCoinDescriptionResponse> {
    return fm.fetchReq<CreateAppCoinDescriptionRequest, CreateAppCoinDescriptionResponse>(`/v1/create/app/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCoinDescriptions(req: CreateAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CreateAppCoinDescriptionsResponse> {
    return fm.fetchReq<CreateAppCoinDescriptionsRequest, CreateAppCoinDescriptionsResponse>(`/v1/create/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinDescription(req: UpdateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<UpdateCoinDescriptionResponse> {
    return fm.fetchReq<UpdateCoinDescriptionRequest, UpdateCoinDescriptionResponse>(`/v1/update/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescription(req: GetCoinDescriptionRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionResponse> {
    return fm.fetchReq<GetCoinDescriptionRequest, GetCoinDescriptionResponse>(`/v1/get/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescriptions(req: GetCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionsResponse> {
    return fm.fetchReq<GetCoinDescriptionsRequest, GetCoinDescriptionsResponse>(`/v1/get/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescriptionOnly(req: GetCoinDescriptionOnlyRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionOnlyResponse> {
    return fm.fetchReq<GetCoinDescriptionOnlyRequest, GetCoinDescriptionOnlyResponse>(`/v1/get/coin/description/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinDescriptions(req: GetAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionsResponse> {
    return fm.fetchReq<GetAppCoinDescriptionsRequest, GetAppCoinDescriptionsResponse>(`/v1/get/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinDescriptionOnly(req: GetAppCoinDescriptionOnlyRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionOnlyResponse> {
    return fm.fetchReq<GetAppCoinDescriptionOnlyRequest, GetAppCoinDescriptionOnlyResponse>(`/v1/get/app/coin/description/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCoinDescription(req: DeleteCoinDescriptionRequest, initReq?: fm.InitReq): Promise<DeleteCoinDescriptionResponse> {
    return fm.fetchReq<DeleteCoinDescriptionRequest, DeleteCoinDescriptionResponse>(`/v1/delete/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinProduct(req: CreateCoinProductRequest, initReq?: fm.InitReq): Promise<CreateCoinProductResponse> {
    return fm.fetchReq<CreateCoinProductRequest, CreateCoinProductResponse>(`/v1/create/coin/product`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinProducts(req: CreateCoinProductsRequest, initReq?: fm.InitReq): Promise<CreateCoinProductsResponse> {
    return fm.fetchReq<CreateCoinProductsRequest, CreateCoinProductsResponse>(`/v1/create/coin/products`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCoinProduct(req: CreateAppCoinProductRequest, initReq?: fm.InitReq): Promise<CreateAppCoinProductResponse> {
    return fm.fetchReq<CreateAppCoinProductRequest, CreateAppCoinProductResponse>(`/v1/create/app/coin/product`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCoinProducts(req: CreateAppCoinProductsRequest, initReq?: fm.InitReq): Promise<CreateAppCoinProductsResponse> {
    return fm.fetchReq<CreateAppCoinProductsRequest, CreateAppCoinProductsResponse>(`/v1/create/app/coin/products`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinProduct(req: UpdateCoinProductRequest, initReq?: fm.InitReq): Promise<UpdateCoinProductResponse> {
    return fm.fetchReq<UpdateCoinProductRequest, UpdateCoinProductResponse>(`/v1/update/coin/product`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinProduct(req: GetCoinProductRequest, initReq?: fm.InitReq): Promise<GetCoinProductResponse> {
    return fm.fetchReq<GetCoinProductRequest, GetCoinProductResponse>(`/v1/get/coin/product`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinProducts(req: GetCoinProductsRequest, initReq?: fm.InitReq): Promise<GetCoinProductsResponse> {
    return fm.fetchReq<GetCoinProductsRequest, GetCoinProductsResponse>(`/v1/get/coin/products`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinProductOnly(req: GetCoinProductOnlyRequest, initReq?: fm.InitReq): Promise<GetCoinProductOnlyResponse> {
    return fm.fetchReq<GetCoinProductOnlyRequest, GetCoinProductOnlyResponse>(`/v1/get/coin/product/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinProducts(req: GetAppCoinProductsRequest, initReq?: fm.InitReq): Promise<GetAppCoinProductsResponse> {
    return fm.fetchReq<GetAppCoinProductsRequest, GetAppCoinProductsResponse>(`/v1/get/app/coin/products`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinProductOnly(req: GetAppCoinProductOnlyRequest, initReq?: fm.InitReq): Promise<GetAppCoinProductOnlyResponse> {
    return fm.fetchReq<GetAppCoinProductOnlyRequest, GetAppCoinProductOnlyResponse>(`/v1/get/app/coin/product/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCoinProduct(req: DeleteCoinProductRequest, initReq?: fm.InitReq): Promise<DeleteCoinProductResponse> {
    return fm.fetchReq<DeleteCoinProductRequest, DeleteCoinProductResponse>(`/v1/delete/coin/product`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}