/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Reward = {
  id?: string
  coinTypeID?: string
  dailyReward?: number
}

export type CreateRewardRequest = {
  info?: Reward
}

export type CreateRewardResponse = {
  info?: Reward
}

export type CreateRewardsRequest = {
  infos?: Reward[]
}

export type CreateRewardsResponse = {
  infos?: Reward[]
}

export type UpdateRewardRequest = {
  info?: Reward
}

export type UpdateRewardResponse = {
  info?: Reward
}

export type GetRewardRequest = {
  id?: string
}

export type GetRewardResponse = {
  info?: Reward
}

export type GetRewardOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetRewardOnlyResponse = {
  info?: Reward
}

export type GetRewardsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetRewardsResponse = {
  infos?: Reward[]
  total?: number
}

export type ExistRewardRequest = {
  id?: string
}

export type ExistRewardResponse = {
  result?: boolean
}

export type ExistRewardCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistRewardCondsResponse = {
  result?: boolean
}

export type CountRewardsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type CountRewardsResponse = {
  result?: number
}

export type DeleteRewardRequest = {
  id?: string
}

export type DeleteRewardResponse = {
  info?: Reward
}

export type Currency = {
  id?: string
  appID?: string
  coinTypeID?: string
  priceVSUSDT?: number
  appPriceVSUSDT?: number
}

export type CreateCurrencyRequest = {
  info?: Currency
}

export type CreateCurrencyResponse = {
  info?: Currency
}

export type CreateAppCurrencyRequest = {
  targetAppID?: string
  info?: Currency
}

export type CreateAppCurrencyResponse = {
  info?: Currency
}

export type UpdateCurrencyRequest = {
  info?: Currency
}

export type UpdateCurrencyResponse = {
  info?: Currency
}

export type GetCurrencyRequest = {
  id?: string
}

export type GetCurrencyResponse = {
  info?: Currency
}

export type GetCurrencyOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetCurrencyOnlyResponse = {
  info?: Currency
}

export type GetCurrenciesRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCurrenciesResponse = {
  infos?: Currency[]
  total?: number
}

export type GetAppCurrenciesRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCurrenciesResponse = {
  infos?: Currency[]
  total?: number
}

export type DeleteCurrencyRequest = {
  id?: string
}

export type DeleteCurrencyResponse = {
  info?: Currency
}

export type CurrencyAmount = {
  coinTypeID?: string
  amount?: number
}

export type CurrenciesRequest = {
  appID?: string
  coinTypeIDs?: string[]
}

export type CurrenciesResponse = {
  infos?: CurrencyAmount[]
}

export class OracleManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReward(req: CreateRewardRequest, initReq?: fm.InitReq): Promise<CreateRewardResponse> {
    return fm.fetchReq<CreateRewardRequest, CreateRewardResponse>(`/v1/create/reward`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateRewards(req: CreateRewardsRequest, initReq?: fm.InitReq): Promise<CreateRewardsResponse> {
    return fm.fetchReq<CreateRewardsRequest, CreateRewardsResponse>(`/v1/create/rewards`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateReward(req: UpdateRewardRequest, initReq?: fm.InitReq): Promise<UpdateRewardResponse> {
    return fm.fetchReq<UpdateRewardRequest, UpdateRewardResponse>(`/v1/update/reward`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReward(req: GetRewardRequest, initReq?: fm.InitReq): Promise<GetRewardResponse> {
    return fm.fetchReq<GetRewardRequest, GetRewardResponse>(`/v1/get/reward`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRewardOnly(req: GetRewardOnlyRequest, initReq?: fm.InitReq): Promise<GetRewardOnlyResponse> {
    return fm.fetchReq<GetRewardOnlyRequest, GetRewardOnlyResponse>(`/v1/get/reward/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRewards(req: GetRewardsRequest, initReq?: fm.InitReq): Promise<GetRewardsResponse> {
    return fm.fetchReq<GetRewardsRequest, GetRewardsResponse>(`/v1/get/rewards`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistReward(req: ExistRewardRequest, initReq?: fm.InitReq): Promise<ExistRewardResponse> {
    return fm.fetchReq<ExistRewardRequest, ExistRewardResponse>(`/v1/exist/reward`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistRewardConds(req: ExistRewardCondsRequest, initReq?: fm.InitReq): Promise<ExistRewardCondsResponse> {
    return fm.fetchReq<ExistRewardCondsRequest, ExistRewardCondsResponse>(`/v1/exist/reward/conds`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteReward(req: DeleteRewardRequest, initReq?: fm.InitReq): Promise<DeleteRewardResponse> {
    return fm.fetchReq<DeleteRewardRequest, DeleteRewardResponse>(`/v1/delete/reward`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCurrency(req: CreateCurrencyRequest, initReq?: fm.InitReq): Promise<CreateCurrencyResponse> {
    return fm.fetchReq<CreateCurrencyRequest, CreateCurrencyResponse>(`/v1/create/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppCurrency(req: CreateAppCurrencyRequest, initReq?: fm.InitReq): Promise<CreateAppCurrencyResponse> {
    return fm.fetchReq<CreateAppCurrencyRequest, CreateAppCurrencyResponse>(`/v1/create/app/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCurrency(req: UpdateCurrencyRequest, initReq?: fm.InitReq): Promise<UpdateCurrencyResponse> {
    return fm.fetchReq<UpdateCurrencyRequest, UpdateCurrencyResponse>(`/v1/update/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCurrency(req: GetCurrencyRequest, initReq?: fm.InitReq): Promise<GetCurrencyResponse> {
    return fm.fetchReq<GetCurrencyRequest, GetCurrencyResponse>(`/v1/get/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCurrencyOnly(req: GetCurrencyOnlyRequest, initReq?: fm.InitReq): Promise<GetCurrencyOnlyResponse> {
    return fm.fetchReq<GetCurrencyOnlyRequest, GetCurrencyOnlyResponse>(`/v1/get/currency/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCurrencies(req: GetCurrenciesRequest, initReq?: fm.InitReq): Promise<GetCurrenciesResponse> {
    return fm.fetchReq<GetCurrenciesRequest, GetCurrenciesResponse>(`/v1/get/currencies`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCurrencies(req: GetAppCurrenciesRequest, initReq?: fm.InitReq): Promise<GetAppCurrenciesResponse> {
    return fm.fetchReq<GetAppCurrenciesRequest, GetAppCurrenciesResponse>(`/v1/get/app/currencies`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCurrency(req: DeleteCurrencyRequest, initReq?: fm.InitReq): Promise<DeleteCurrencyResponse> {
    return fm.fetchReq<DeleteCurrencyRequest, DeleteCurrencyResponse>(`/v1/delete/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Currencies(req: CurrenciesRequest, initReq?: fm.InitReq): Promise<CurrenciesResponse> {
    return fm.fetchReq<CurrenciesRequest, CurrenciesResponse>(`/v1/currency`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}