/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type KycInfo = {
  id?: string
  appID?: string
  userID?: string
  cardType?: string
  cardID?: string
  frontCardImg?: string
  backCardImg?: string
  userHandingCardImg?: string
  createAt?: number
  updateAt?: number
}

export type CreateKycRequest = {
  info?: KycInfo
}

export type CreateKycResponse = {
  info?: KycInfo
}

export type GetKycByUserIDRequest = {
  appID?: string
  userID?: string
}

export type GetKycByUserIDResponse = {
  info?: KycInfo
}

export type GetKycByAppIDRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetKycByAppIDResponse = {
  infos?: KycInfo[]
  total?: number
}

export type GetAllKycRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAllKycResponse = {
  infos?: KycInfo[]
  total?: number
}

export type UpdateKycRequest = {
  info?: KycInfo
}

export type UpdateKycResponse = {
  info?: KycInfo
}

export type UploadKycImageRequest = {
  appID?: string
  userID?: string
  imageType?: string
  imageBase64?: string
}

export type UploadKycImageResponse = {
  info?: string
}

export type GetKycImageRequest = {
  appID?: string
  userID?: string
  imageS3Key?: string
}

export type GetKycImageResponse = {
  info?: string
}

export type GetKycByKycIDsRequest = {
  kycIDs?: string[]
}

export type GetKycByKycIDsResponse = {
  infos?: KycInfo[]
}

export class KycManagement {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateKyc(req: CreateKycRequest, initReq?: fm.InitReq): Promise<CreateKycResponse> {
    return fm.fetchReq<CreateKycRequest, CreateKycResponse>(`/v1/create/kyc`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycByUserID(req: GetKycByUserIDRequest, initReq?: fm.InitReq): Promise<GetKycByUserIDResponse> {
    return fm.fetchReq<GetKycByUserIDRequest, GetKycByUserIDResponse>(`/v1/get/kyc/by/userid`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycByAppID(req: GetKycByAppIDRequest, initReq?: fm.InitReq): Promise<GetKycByAppIDResponse> {
    return fm.fetchReq<GetKycByAppIDRequest, GetKycByAppIDResponse>(`/v1/get/kyc/by/appid`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAllKyc(req: GetAllKycRequest, initReq?: fm.InitReq): Promise<GetAllKycResponse> {
    return fm.fetchReq<GetAllKycRequest, GetAllKycResponse>(`/v1/get/all/kyc`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateKyc(req: UpdateKycRequest, initReq?: fm.InitReq): Promise<UpdateKycResponse> {
    return fm.fetchReq<UpdateKycRequest, UpdateKycResponse>(`/v1/update/kyc`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UploadKycImage(req: UploadKycImageRequest, initReq?: fm.InitReq): Promise<UploadKycImageResponse> {
    return fm.fetchReq<UploadKycImageRequest, UploadKycImageResponse>(`/v1/upload/kyc/image`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycImage(req: GetKycImageRequest, initReq?: fm.InitReq): Promise<GetKycImageResponse> {
    return fm.fetchReq<GetKycImageRequest, GetKycImageResponse>(`/v1/get/kyc/image`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetKycByKycIDs(req: GetKycByKycIDsRequest, initReq?: fm.InitReq): Promise<GetKycByKycIDsResponse> {
    return fm.fetchReq<GetKycByKycIDsRequest, GetKycByKycIDsResponse>(`/v1/get/kyc/by/kyc/ids`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}