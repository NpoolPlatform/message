/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinAccountInfo = {
  id?: string
  coinTypeID?: string
  address?: string
  platformHoldPrivateKey?: boolean
}

export type CreateCoinAccountRequest = {
  info?: CoinAccountInfo
}

export type CreateCoinAccountResponse = {
  info?: CoinAccountInfo
}

export type GetCoinAccountRequest = {
  id?: string
}

export type GetCoinAccountResponse = {
  info?: CoinAccountInfo
}

export type GetCoinAccountByCoinAddressRequest = {
  coinInfoID?: string
  address?: string
}

export type GetCoinAccountByCoinAddressResponse = {
  info?: CoinAccountInfo
}

export type GetCoinAccountsRequest = {
}

export type GetCoinAccountsResponse = {
  infos?: CoinAccountInfo[]
}

export type DeleteCoinAccountRequest = {
  id?: string
}

export type DeleteCoinAccountResponse = {
  info?: CoinAccountInfo
}

export type CoinAccountTransaction = {
  id?: string
  appID?: string
  userID?: string
  fromAddressID?: string
  toAddressID?: string
  coinTypeID?: string
  amount?: number
  message?: string
  createAt?: number
  state?: string
  chainTransactionID?: string
}

export type CreateCoinAccountTransactionRequest = {
  info?: CoinAccountTransaction
}

export type CreateCoinAccountTransactionResponse = {
  info?: CoinAccountTransaction
}

export type GetCoinAccountTransactionRequest = {
  id?: string
}

export type GetCoinAccountTransactionResponse = {
  info?: CoinAccountTransaction
}

export type CoinAccountTransactionDetail = {
  id?: string
  appID?: string
  userID?: string
  fromAddress?: CoinAccountInfo
  toAddress?: CoinAccountInfo
  coinTypeID?: string
  amount?: number
  message?: string
  createAt?: number
  state?: string
  chainTransactionID?: string
}

export type GetCoinAccountTransactionDetailRequest = {
  id?: string
}

export type GetCoinAccountTransactionDetailResponse = {
  detail?: CoinAccountTransactionDetail
}

export type GetCoinAccountTransactionsByCoinAccountRequest = {
  coinTypeID?: string
  addressID?: string
}

export type GetCoinAccountTransactionsByCoinAccountResponse = {
  infos?: CoinAccountTransaction[]
}

export type GetCoinAccountTransactionsByStateRequest = {
  state?: string
}

export type GetCoinAccountTransactionsByStateResponse = {
  infos?: CoinAccountTransaction[]
}

export type GetCoinAccountTransactionsByCoinRequest = {
  coinTypeID?: string
}

export type GetCoinAccountTransactionsByCoinResponse = {
  infos?: CoinAccountTransaction[]
}

export type UpdateCoinAccountTransactionRequest = {
  info?: CoinAccountTransaction
}

export type UpdateCoinAccountTransactionResponse = {
  info?: CoinAccountTransaction
}

export type DeleteCoinAccountTransactionRequest = {
  id?: string
}

export type DeleteCoinAccountTransactionResponse = {
  info?: CoinAccountTransaction
}

export type PlatformBenefit = {
  id?: string
  goodID?: string
  benefitAccountID?: string
  amount?: number
  createAt?: number
  chainTransactionID?: string
  lastBenefitTimestamp?: number
}

export type CreatePlatformBenefitRequest = {
  info?: PlatformBenefit
}

export type CreatePlatformBenefitResponse = {
  info?: PlatformBenefit
}

export type GetPlatformBenefitsByGoodRequest = {
  goodID?: string
}

export type GetPlatformBenefitsByGoodResponse = {
  infos?: PlatformBenefit[]
}

export type PlatformBenefitDetail = {
  id?: string
  goodID?: string
  benefitAddress?: CoinAccountInfo
  amount?: number
  createAt?: number
  chainTransactionID?: string
}

export type GetPlatformBenefitDetailRequest = {
  id?: string
}

export type GetPlatformBenefitDetailResponse = {
  detail?: PlatformBenefitDetail
}

export type GetPlatformBenefitRequest = {
  id?: string
}

export type GetPlatformBenefitResponse = {
  info?: PlatformBenefit
}

export type GetLatestPlatformBenefitByGoodRequest = {
  goodID?: string
}

export type GetLatestPlatformBenefitByGoodResponse = {
  info?: PlatformBenefit
}

export type PlatformSetting = {
  id?: string
  warmAccountUSDAmount?: number
}

export type CreatePlatformSettingRequest = {
  info?: PlatformSetting
}

export type CreatePlatformSettingResponse = {
  info?: PlatformSetting
}

export type UpdatePlatformSettingRequest = {
  info?: PlatformSetting
}

export type UpdatePlatformSettingResponse = {
  info?: PlatformSetting
}

export type GetPlatformSettingByGoodRequest = {
  goodID?: string
}

export type GetPlatformSettingByGoodResponse = {
  info?: PlatformSetting
}

export type GetPlatformSettingRequest = {
  id?: string
}

export type GetPlatformSettingResponse = {
  info?: PlatformSetting
}

export type UserBenefit = {
  id?: string
  userID?: string
  appID?: string
  goodID?: string
  amount?: number
  createAt?: number
  orderID?: string
  lastBenefitTimestamp?: number
}

export type CreateUserBenefitRequest = {
  info?: UserBenefit
}

export type CreateUserBenefitResponse = {
  info?: UserBenefit
}

export type GetUserBenefitsByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetUserBenefitsByAppUserResponse = {
  infos?: UserBenefit[]
}

export type GetUserBenefitsByAppRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetUserBenefitsByAppResponse = {
  infos?: UserBenefit[]
}

export type GetLatestUserBenefitByGoodAppUserRequest = {
  appID?: string
  userID?: string
  goodID?: string
}

export type GetLatestUserBenefitByGoodAppUserResponse = {
  info?: UserBenefit
}

export type CoinSetting = {
  id?: string
  coinTypeID?: string
  warmAccountCoinAmount?: number
}

export type CreateCoinSettingRequest = {
  info?: CoinSetting
}

export type CreateCoinSettingResponse = {
  info?: CoinSetting
}

export type UpdateCoinSettingRequest = {
  info?: CoinSetting
}

export type UpdateCoinSettingResponse = {
  info?: CoinSetting
}

export type GetCoinSettingRequest = {
  id?: string
}

export type GetCoinSettingResponse = {
  info?: CoinSetting
}

export type GetCoinSettingByCoinRequest = {
  coinTypeID?: string
}

export type GetCoinSettingByCoinResponse = {
  info?: CoinSetting
}

export type GoodBenefit = {
  id?: string
  goodID?: string
  benefitAccountID?: string
  platformOfflineAccountID?: string
  userOfflineAccountID?: string
  userOnlineAccountID?: string
  benefitIntervalHours?: number
}

export type CreateGoodBenefitRequest = {
  info?: GoodBenefit
}

export type CreateGoodBenefitResponse = {
  info?: GoodBenefit
}

export type UpdateGoodBenefitRequest = {
  info?: GoodBenefit
}

export type UpdateGoodBenefitResponse = {
  info?: GoodBenefit
}

export type GetGoodBenefitRequest = {
  id?: string
}

export type GetGoodBenefitResponse = {
  info?: GoodBenefit
}

export type GetGoodBenefitByGoodRequest = {
  goodID?: string
}

export type GetGoodBenefitByGoodResponse = {
  info?: GoodBenefit
}

export type GoodPayment = {
  id?: string
  goodID?: string
  paymentCoinTypeID?: string
  accountID?: string
  idle?: boolean
}

export type CreateGoodPaymentRequest = {
  info?: GoodPayment
}

export type CreateGoodPaymentResponse = {
  info?: GoodPayment
}

export type UpdateGoodPaymentRequest = {
  info?: GoodPayment
}

export type UpdateGoodPaymentResponse = {
  info?: GoodPayment
}

export type GetGoodPaymentRequest = {
  id?: string
}

export type GetGoodPaymentResponse = {
  info?: GoodPayment
}

export type GetGoodPaymentsRequest = {
}

export type GetGoodPaymentsResponse = {
  infos?: GoodPayment[]
}

export type GetGoodPaymentsByGoodRequest = {
  goodID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetGoodPaymentsByGoodResponse = {
  infos?: GoodPayment[]
}

export type GetIdleGoodPaymentsByGoodRequest = {
  goodID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetIdleGoodPaymentsByGoodResponse = {
  infos?: GoodPayment[]
}

export type GetIdleGoodPaymentsByGoodPaymentCoinRequest = {
  goodID?: string
  paymentCoinTypeID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetIdleGoodPaymentsByGoodPaymentCoinResponse = {
  infos?: GoodPayment[]
}

export type GetGoodPaymentByAccountRequest = {
  accountID?: string
}

export type GetGoodPaymentByAccountResponse = {
  info?: GoodPayment
}

export type GoodSetting = {
  id?: string
  goodID?: string
  warmAccountUSDAmount?: number
  warmAccountCoinAmount?: number
}

export type CreateGoodSettingRequest = {
  info?: GoodSetting
}

export type CreateGoodSettingResponse = {
  info?: GoodSetting
}

export type UpdateGoodSettingRequest = {
  info?: GoodSetting
}

export type UpdateGoodSettingResponse = {
  info?: GoodSetting
}

export type GetGoodSettingRequest = {
  id?: string
}

export type GetGoodSettingResponse = {
  info?: GoodSetting
}

export type GetGoodSettingByGoodRequest = {
  goodID?: string
}

export type GetGoodSettingByGoodResponse = {
  info?: GoodSetting
}

export type UserWithdraw = {
  id?: string
  appID?: string
  userID?: string
  coinTypeID?: string
  accountID?: string
  name?: string
  message?: string
}

export type CreateUserWithdrawRequest = {
  info?: UserWithdraw
}

export type CreateUserWithdrawResponse = {
  info?: UserWithdraw
}

export type UpdateUserWithdrawRequest = {
  info?: UserWithdraw
}

export type UpdateUserWithdrawResponse = {
  info?: UserWithdraw
}

export type GetUserWithdrawRequest = {
  id?: string
}

export type GetUserWithdrawResponse = {
  info?: UserWithdraw
}

export type GetUserWithdrawsByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetUserWithdrawsByAppUserResponse = {
  infos?: UserWithdraw[]
}

export type GetUserWithdrawsByOtherAppUserRequest = {
  targetAppID?: string
  targetUserID?: string
}

export type GetUserWithdrawsByOtherAppUserResponse = {
  infos?: UserWithdraw[]
}

export type GetUserWithdrawByAccountRequest = {
  accountID?: string
}

export type GetUserWithdrawByAccountResponse = {
  info?: UserWithdraw
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
  static GetCoinAccounts(req: GetCoinAccountsRequest, initReq?: fm.InitReq): Promise<GetCoinAccountsResponse> {
    return fm.fetchReq<GetCoinAccountsRequest, GetCoinAccountsResponse>(`/v1/get/coin/accounts`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static CreateCoinSetting(req: CreateCoinSettingRequest, initReq?: fm.InitReq): Promise<CreateCoinSettingResponse> {
    return fm.fetchReq<CreateCoinSettingRequest, CreateCoinSettingResponse>(`/v1/create/coin/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinSetting(req: UpdateCoinSettingRequest, initReq?: fm.InitReq): Promise<UpdateCoinSettingResponse> {
    return fm.fetchReq<UpdateCoinSettingRequest, UpdateCoinSettingResponse>(`/v1/update/coin/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinSetting(req: GetCoinSettingRequest, initReq?: fm.InitReq): Promise<GetCoinSettingResponse> {
    return fm.fetchReq<GetCoinSettingRequest, GetCoinSettingResponse>(`/v1/get/coin/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinSettingByCoin(req: GetCoinSettingByCoinRequest, initReq?: fm.InitReq): Promise<GetCoinSettingByCoinResponse> {
    return fm.fetchReq<GetCoinSettingByCoinRequest, GetCoinSettingByCoinResponse>(`/v1/get/coin/setting/by/coin`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodBenefit(req: CreateGoodBenefitRequest, initReq?: fm.InitReq): Promise<CreateGoodBenefitResponse> {
    return fm.fetchReq<CreateGoodBenefitRequest, CreateGoodBenefitResponse>(`/v1/create/good/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodBenefit(req: UpdateGoodBenefitRequest, initReq?: fm.InitReq): Promise<UpdateGoodBenefitResponse> {
    return fm.fetchReq<UpdateGoodBenefitRequest, UpdateGoodBenefitResponse>(`/v1/update/good/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodBenefit(req: GetGoodBenefitRequest, initReq?: fm.InitReq): Promise<GetGoodBenefitResponse> {
    return fm.fetchReq<GetGoodBenefitRequest, GetGoodBenefitResponse>(`/v1/get/good/benefit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodBenefitByGood(req: GetGoodBenefitByGoodRequest, initReq?: fm.InitReq): Promise<GetGoodBenefitByGoodResponse> {
    return fm.fetchReq<GetGoodBenefitByGoodRequest, GetGoodBenefitByGoodResponse>(`/v1/get/good/benefit/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodPayment(req: CreateGoodPaymentRequest, initReq?: fm.InitReq): Promise<CreateGoodPaymentResponse> {
    return fm.fetchReq<CreateGoodPaymentRequest, CreateGoodPaymentResponse>(`/v1/create/good/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodPayment(req: UpdateGoodPaymentRequest, initReq?: fm.InitReq): Promise<UpdateGoodPaymentResponse> {
    return fm.fetchReq<UpdateGoodPaymentRequest, UpdateGoodPaymentResponse>(`/v1/update/good/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodPayment(req: GetGoodPaymentRequest, initReq?: fm.InitReq): Promise<GetGoodPaymentResponse> {
    return fm.fetchReq<GetGoodPaymentRequest, GetGoodPaymentResponse>(`/v1/get/good/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodPayments(req: GetGoodPaymentsRequest, initReq?: fm.InitReq): Promise<GetGoodPaymentsResponse> {
    return fm.fetchReq<GetGoodPaymentsRequest, GetGoodPaymentsResponse>(`/v1/get/good/payments`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodPaymentsByGood(req: GetGoodPaymentsByGoodRequest, initReq?: fm.InitReq): Promise<GetGoodPaymentsByGoodResponse> {
    return fm.fetchReq<GetGoodPaymentsByGoodRequest, GetGoodPaymentsByGoodResponse>(`/v1/get/good/payments/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetIdleGoodPaymentsByGood(req: GetIdleGoodPaymentsByGoodRequest, initReq?: fm.InitReq): Promise<GetIdleGoodPaymentsByGoodResponse> {
    return fm.fetchReq<GetIdleGoodPaymentsByGoodRequest, GetIdleGoodPaymentsByGoodResponse>(`/v1/get/idle/good/payments/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetIdleGoodPaymentsByGoodPaymentCoin(req: GetIdleGoodPaymentsByGoodPaymentCoinRequest, initReq?: fm.InitReq): Promise<GetIdleGoodPaymentsByGoodPaymentCoinResponse> {
    return fm.fetchReq<GetIdleGoodPaymentsByGoodPaymentCoinRequest, GetIdleGoodPaymentsByGoodPaymentCoinResponse>(`/v1/get/idle/good/payments/by/good/payment/coin`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodPaymentByAccount(req: GetGoodPaymentByAccountRequest, initReq?: fm.InitReq): Promise<GetGoodPaymentByAccountResponse> {
    return fm.fetchReq<GetGoodPaymentByAccountRequest, GetGoodPaymentByAccountResponse>(`/v1/get/good/payment/by/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodSetting(req: CreateGoodSettingRequest, initReq?: fm.InitReq): Promise<CreateGoodSettingResponse> {
    return fm.fetchReq<CreateGoodSettingRequest, CreateGoodSettingResponse>(`/v1/create/good/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGoodSetting(req: UpdateGoodSettingRequest, initReq?: fm.InitReq): Promise<UpdateGoodSettingResponse> {
    return fm.fetchReq<UpdateGoodSettingRequest, UpdateGoodSettingResponse>(`/v1/update/good/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodSetting(req: GetGoodSettingRequest, initReq?: fm.InitReq): Promise<GetGoodSettingResponse> {
    return fm.fetchReq<GetGoodSettingRequest, GetGoodSettingResponse>(`/v1/get/good/setting`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodSettingByGood(req: GetGoodSettingByGoodRequest, initReq?: fm.InitReq): Promise<GetGoodSettingByGoodResponse> {
    return fm.fetchReq<GetGoodSettingByGoodRequest, GetGoodSettingByGoodResponse>(`/v1/get/good/setting/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateUserWithdraw(req: CreateUserWithdrawRequest, initReq?: fm.InitReq): Promise<CreateUserWithdrawResponse> {
    return fm.fetchReq<CreateUserWithdrawRequest, CreateUserWithdrawResponse>(`/v1/create/user/withdraw`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserWithdraw(req: UpdateUserWithdrawRequest, initReq?: fm.InitReq): Promise<UpdateUserWithdrawResponse> {
    return fm.fetchReq<UpdateUserWithdrawRequest, UpdateUserWithdrawResponse>(`/v1/update/user/withdraw`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserWithdraw(req: GetUserWithdrawRequest, initReq?: fm.InitReq): Promise<GetUserWithdrawResponse> {
    return fm.fetchReq<GetUserWithdrawRequest, GetUserWithdrawResponse>(`/v1/get/user/withdraw`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserWithdrawsByAppUser(req: GetUserWithdrawsByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserWithdrawsByAppUserResponse> {
    return fm.fetchReq<GetUserWithdrawsByAppUserRequest, GetUserWithdrawsByAppUserResponse>(`/v1/get/user/withdraws/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserWithdrawsByOtherAppUser(req: GetUserWithdrawsByOtherAppUserRequest, initReq?: fm.InitReq): Promise<GetUserWithdrawsByOtherAppUserResponse> {
    return fm.fetchReq<GetUserWithdrawsByOtherAppUserRequest, GetUserWithdrawsByOtherAppUserResponse>(`/v1/get/user/withdraws/by/other/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserWithdrawByAccount(req: GetUserWithdrawByAccountRequest, initReq?: fm.InitReq): Promise<GetUserWithdrawByAccountResponse> {
    return fm.fetchReq<GetUserWithdrawByAccountRequest, GetUserWithdrawByAccountResponse>(`/v1/get/user/withdraw/by/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}