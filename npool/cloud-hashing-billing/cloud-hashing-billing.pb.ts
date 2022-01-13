/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinAccountInfo = {
  ID?: string
  CoinTypeID?: string
  Address?: string
  GeneratedBy?: string
  UsedFor?: string
  Idle?: boolean
  AppID?: string
  UserID?: string
}

export type CreateCoinAccountRequest = {
  Info?: CoinAccountInfo
}

export type CreateCoinAccountResponse = {
  Info?: CoinAccountInfo
}

export type GetCoinAccountRequest = {
  ID?: string
}

export type GetCoinAccountResponse = {
  Info?: CoinAccountInfo
}

export type GetCoinAccountByCoinAddressRequest = {
  CoinInfoID?: string
  Address?: string
}

export type GetCoinAccountByCoinAddressResponse = {
  Info?: CoinAccountInfo
}

export type GetCoinAccountsByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetCoinAccountsByAppUserResponse = {
  Infos?: CoinAccountInfo[]
}

export type DeleteCoinAccountRequest = {
  ID?: string
}

export type DeleteCoinAccountResponse = {
  Info?: CoinAccountInfo
}

export type CoinAccountTransaction = {
  ID?: string
  AppID?: string
  UserID?: string
  FromAddressID?: string
  ToAddressID?: string
  CoinTypeID?: string
  Amount?: number
  Message?: string
  CreateAt?: number
  State?: string
  ChainTransactionID?: string
  PlatformTransactionID?: string
}

export type CreateCoinAccountTransactionRequest = {
  Info?: CoinAccountTransaction
}

export type CreateCoinAccountTransactionResponse = {
  Info?: CoinAccountTransaction
}

export type GetCoinAccountTransactionRequest = {
  ID?: string
}

export type GetCoinAccountTransactionResponse = {
  Info?: CoinAccountTransaction
}

export type CoinAccountTransactionDetail = {
  ID?: string
  AppID?: string
  UserID?: string
  FromAddress?: CoinAccountInfo
  ToAddress?: CoinAccountInfo
  CoinTypeID?: string
  Amount?: number
  Message?: string
  CreateAt?: number
  State?: string
  ChainTransactionID?: string
  PlatformTransactionID?: string
}

export type GetCoinAccountTransactionDetailRequest = {
  ID?: string
}

export type GetCoinAccountTransactionDetailResponse = {
  Detail?: CoinAccountTransactionDetail
}

export type GetCoinAccountTransactionsByCoinAccountRequest = {
  CoinTypeID?: string
  AddressID?: string
}

export type GetCoinAccountTransactionsByCoinAccountResponse = {
  Infos?: CoinAccountTransaction[]
}

export type GetCoinAccountTransactionsByStateRequest = {
  State?: string
}

export type GetCoinAccountTransactionsByStateResponse = {
  Infos?: CoinAccountTransaction[]
}

export type GetCoinAccountTransactionsByCoinRequest = {
  CoinTypeID?: string
}

export type GetCoinAccountTransactionsByCoinResponse = {
  Infos?: CoinAccountTransaction[]
}

export type UpdateCoinAccountTransactionRequest = {
  Info?: CoinAccountTransaction
}

export type UpdateCoinAccountTransactionResponse = {
  Info?: CoinAccountTransaction
}

export type DeleteCoinAccountTransactionRequest = {
  ID?: string
}

export type DeleteCoinAccountTransactionResponse = {
  Info?: CoinAccountTransaction
}

export type PlatformBenefit = {
  ID?: string
  GoodID?: string
  BenefitAccountID?: string
  Amount?: number
  CreateAt?: number
  ChainTransactionID?: string
  LastBenefitTimestamp?: number
}

export type CreatePlatformBenefitRequest = {
  Info?: PlatformBenefit
}

export type CreatePlatformBenefitResponse = {
  Info?: PlatformBenefit
}

export type GetPlatformBenefitsByGoodRequest = {
  GoodID?: string
}

export type GetPlatformBenefitsByGoodResponse = {
  Infos?: PlatformBenefit[]
}

export type PlatformBenefitDetail = {
  ID?: string
  GoodID?: string
  BenefitAddress?: CoinAccountInfo
  Amount?: number
  CreateAt?: number
  ChainTransactionID?: string
}

export type GetPlatformBenefitDetailRequest = {
  ID?: string
}

export type GetPlatformBenefitDetailResponse = {
  Detail?: PlatformBenefitDetail
}

export type GetPlatformBenefitRequest = {
  ID?: string
}

export type GetPlatformBenefitResponse = {
  Info?: PlatformBenefit
}

export type GetLatestPlatformBenefitByGoodRequest = {
  GoodID?: string
}

export type GetLatestPlatformBenefitByGoodResponse = {
  Info?: PlatformBenefit
}

export type PlatformSetting = {
  ID?: string
  GoodID?: string
  BenefitAccountID?: string
  PlatformOfflineAccountID?: string
  UserOnlineAccountID?: string
  UserOfflineAccountID?: string
  BenefitIntervalHours?: number
}

export type CreatePlatformSettingRequest = {
  Info?: PlatformSetting
}

export type CreatePlatformSettingResponse = {
  Info?: PlatformSetting
}

export type UpdatePlatformSettingRequest = {
  Info?: PlatformSetting
}

export type UpdatePlatformSettingResponse = {
  Info?: PlatformSetting
}

export type GetPlatformSettingByGoodRequest = {
  GoodID?: string
}

export type GetPlatformSettingByGoodResponse = {
  Info?: PlatformSetting
}

export type PlatformSettingDetail = {
  ID?: string
  GoodID?: string
  BenefitAddress?: CoinAccountInfo
  PlatformOfflineAddress?: CoinAccountInfo
  UserOnlineAddress?: CoinAccountInfo
  UserOfflineAddress?: CoinAccountInfo
  BenefitIntervalHours?: number
}

export type GetPlatformSettingDetailRequest = {
  ID?: string
}

export type GetPlatformSettingDetailResponse = {
  Detail?: PlatformSettingDetail
}

export type GetPlatformSettingRequest = {
  ID?: string
}

export type GetPlatformSettingResponse = {
  Info?: PlatformSetting
}

export type UserBenefit = {
  ID?: string
  UserID?: string
  AppID?: string
  GoodID?: string
  Amount?: number
  CreateAt?: number
  OrderID?: string
  LastBenefitTimestamp?: number
}

export type CreateUserBenefitRequest = {
  Info?: UserBenefit
}

export type CreateUserBenefitResponse = {
  Info?: UserBenefit
}

export type GetUserBenefitsByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserBenefitsByAppUserResponse = {
  Infos?: UserBenefit[]
}

export type GetUserBenefitsByAppRequest = {
  AppID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetUserBenefitsByAppResponse = {
  Infos?: UserBenefit[]
}

export type GetLatestUserBenefitByGoodAppUserRequest = {
  AppID?: string
  UserID?: string
  GoodID?: string
}

export type GetLatestUserBenefitByGoodAppUserResponse = {
  Info?: UserBenefit
}

export class CloudHashingBilling {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinAccount(req: CreateCoinAccountRequest, initReq?: fm.InitReq): Promise<CreateCoinAccountResponse> {
    return fm.fetchReq<CreateCoinAccountRequest, CreateCoinAccountResponse>(`/v1/create/coin/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccount(req: GetCoinAccountRequest, initReq?: fm.InitReq): Promise<GetCoinAccountResponse> {
    return fm.fetchReq<GetCoinAccountRequest, GetCoinAccountResponse>(`/v1/get/coin/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountByCoinAddress(req: GetCoinAccountByCoinAddressRequest, initReq?: fm.InitReq): Promise<GetCoinAccountByCoinAddressResponse> {
    return fm.fetchReq<GetCoinAccountByCoinAddressRequest, GetCoinAccountByCoinAddressResponse>(`/v1/get/coin/account/by/coin/address`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountsByAppUser(req: GetCoinAccountsByAppUserRequest, initReq?: fm.InitReq): Promise<GetCoinAccountsByAppUserResponse> {
    return fm.fetchReq<GetCoinAccountsByAppUserRequest, GetCoinAccountsByAppUserResponse>(`/v1/get/coin/accounts/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCoinAccount(req: DeleteCoinAccountRequest, initReq?: fm.InitReq): Promise<DeleteCoinAccountResponse> {
    return fm.fetchReq<DeleteCoinAccountRequest, DeleteCoinAccountResponse>(`/v1/delete/coin/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinAccountTransaction(req: CreateCoinAccountTransactionRequest, initReq?: fm.InitReq): Promise<CreateCoinAccountTransactionResponse> {
    return fm.fetchReq<CreateCoinAccountTransactionRequest, CreateCoinAccountTransactionResponse>(`/v1/create/coin/account/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountTransaction(req: GetCoinAccountTransactionRequest, initReq?: fm.InitReq): Promise<GetCoinAccountTransactionResponse> {
    return fm.fetchReq<GetCoinAccountTransactionRequest, GetCoinAccountTransactionResponse>(`/v1/get/coin/account/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountTransactionsByCoinAccount(req: GetCoinAccountTransactionsByCoinAccountRequest, initReq?: fm.InitReq): Promise<GetCoinAccountTransactionsByCoinAccountResponse> {
    return fm.fetchReq<GetCoinAccountTransactionsByCoinAccountRequest, GetCoinAccountTransactionsByCoinAccountResponse>(`/v1/get/coin/account/transactions/by/coin/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountTransactionsByState(req: GetCoinAccountTransactionsByStateRequest, initReq?: fm.InitReq): Promise<GetCoinAccountTransactionsByStateResponse> {
    return fm.fetchReq<GetCoinAccountTransactionsByStateRequest, GetCoinAccountTransactionsByStateResponse>(`/v1/get/coin/account/transactions/by/state`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountTransactionsByCoin(req: GetCoinAccountTransactionsByCoinRequest, initReq?: fm.InitReq): Promise<GetCoinAccountTransactionsByCoinResponse> {
    return fm.fetchReq<GetCoinAccountTransactionsByCoinRequest, GetCoinAccountTransactionsByCoinResponse>(`/v1/get/coin/account/transactions/by/coin`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinAccountTransaction(req: UpdateCoinAccountTransactionRequest, initReq?: fm.InitReq): Promise<UpdateCoinAccountTransactionResponse> {
    return fm.fetchReq<UpdateCoinAccountTransactionRequest, UpdateCoinAccountTransactionResponse>(`/v1/update/coin/account/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinAccountTransactionDetail(req: GetCoinAccountTransactionDetailRequest, initReq?: fm.InitReq): Promise<GetCoinAccountTransactionDetailResponse> {
    return fm.fetchReq<GetCoinAccountTransactionDetailRequest, GetCoinAccountTransactionDetailResponse>(`/v1/get/coin/account/transaction/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCoinAccountTransaction(req: DeleteCoinAccountTransactionRequest, initReq?: fm.InitReq): Promise<DeleteCoinAccountTransactionResponse> {
    return fm.fetchReq<DeleteCoinAccountTransactionRequest, DeleteCoinAccountTransactionResponse>(`/v1/delete/coin/account/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreatePlatformBenefit(req: CreatePlatformBenefitRequest, initReq?: fm.InitReq): Promise<CreatePlatformBenefitResponse> {
    return fm.fetchReq<CreatePlatformBenefitRequest, CreatePlatformBenefitResponse>(`/v1/create/platform/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLatestPlatformBenefitByGood(req: GetLatestPlatformBenefitByGoodRequest, initReq?: fm.InitReq): Promise<GetLatestPlatformBenefitByGoodResponse> {
    return fm.fetchReq<GetLatestPlatformBenefitByGoodRequest, GetLatestPlatformBenefitByGoodResponse>(`/v1/get/latest/platform/benefit/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformBenefitsByGood(req: GetPlatformBenefitsByGoodRequest, initReq?: fm.InitReq): Promise<GetPlatformBenefitsByGoodResponse> {
    return fm.fetchReq<GetPlatformBenefitsByGoodRequest, GetPlatformBenefitsByGoodResponse>(`/v1/get/platform/benefits/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformBenefit(req: GetPlatformBenefitRequest, initReq?: fm.InitReq): Promise<GetPlatformBenefitResponse> {
    return fm.fetchReq<GetPlatformBenefitRequest, GetPlatformBenefitResponse>(`/v1/get/platform/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformBenefitDetail(req: GetPlatformBenefitDetailRequest, initReq?: fm.InitReq): Promise<GetPlatformBenefitDetailResponse> {
    return fm.fetchReq<GetPlatformBenefitDetailRequest, GetPlatformBenefitDetailResponse>(`/v1/get/platform/benefit/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreatePlatformSetting(req: CreatePlatformSettingRequest, initReq?: fm.InitReq): Promise<CreatePlatformSettingResponse> {
    return fm.fetchReq<CreatePlatformSettingRequest, CreatePlatformSettingResponse>(`/v1/create/platform/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdatePlatformSetting(req: UpdatePlatformSettingRequest, initReq?: fm.InitReq): Promise<UpdatePlatformSettingResponse> {
    return fm.fetchReq<UpdatePlatformSettingRequest, UpdatePlatformSettingResponse>(`/v1/update/platform/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformSettingByGood(req: GetPlatformSettingByGoodRequest, initReq?: fm.InitReq): Promise<GetPlatformSettingByGoodResponse> {
    return fm.fetchReq<GetPlatformSettingByGoodRequest, GetPlatformSettingByGoodResponse>(`/v1/get/platform/setting/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformSettingDetail(req: GetPlatformSettingDetailRequest, initReq?: fm.InitReq): Promise<GetPlatformSettingDetailResponse> {
    return fm.fetchReq<GetPlatformSettingDetailRequest, GetPlatformSettingDetailResponse>(`/v1/get/platform/setting/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformSetting(req: GetPlatformSettingRequest, initReq?: fm.InitReq): Promise<GetPlatformSettingResponse> {
    return fm.fetchReq<GetPlatformSettingRequest, GetPlatformSettingResponse>(`/v1/get/platform/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserBenefit(req: CreateUserBenefitRequest, initReq?: fm.InitReq): Promise<CreateUserBenefitResponse> {
    return fm.fetchReq<CreateUserBenefitRequest, CreateUserBenefitResponse>(`/v1/create/user/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserBenefitsByAppUser(req: GetUserBenefitsByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserBenefitsByAppUserResponse> {
    return fm.fetchReq<GetUserBenefitsByAppUserRequest, GetUserBenefitsByAppUserResponse>(`/v1/get/user/benefits/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserBenefitsByApp(req: GetUserBenefitsByAppRequest, initReq?: fm.InitReq): Promise<GetUserBenefitsByAppResponse> {
    return fm.fetchReq<GetUserBenefitsByAppRequest, GetUserBenefitsByAppResponse>(`/v1/get/user/benefits/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLatestUserBenefitByGoodAppUser(req: GetLatestUserBenefitByGoodAppUserRequest, initReq?: fm.InitReq): Promise<GetLatestUserBenefitByGoodAppUserResponse> {
    return fm.fetchReq<GetLatestUserBenefitByGoodAppUserRequest, GetLatestUserBenefitByGoodAppUserResponse>(`/v1/get/latest/user/benefit/by/good/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}