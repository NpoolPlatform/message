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
}