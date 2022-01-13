/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  Info?: string
}

export type KycInfo = {
  ID?: string
  UserID?: string
  AppID?: string
  FirstName?: string
  LastName?: string
  Region?: string
  CardType?: string
  CardID?: string
  FrontCardImg?: string
  BackCardImg?: string
  UserHandlingCardImg?: string
  CreateAt?: number
  UpdateAt?: number
}

export type CreateKycRequest = {
  AppID?: string
  UserID?: string
  FirstName?: string
  LastName?: string
  Region?: string
  CardType?: string
  CardID?: string
  FrontCardImg?: string
  BackCardImg?: string
  UserHandlingCardImg?: string
}

export type CreateKycResponse = {
  Info?: KycInfo
}

export type GetKycByUserIDRequest = {
  AppID?: string
  UserID?: string
}

export type GetKycByUserIDResponse = {
  Info?: KycInfo
}

export type GetKycByAppIDRequest = {
  AppID?: string
  Limit?: number
  Offset?: number
}

export type GetKycByAppIDResponse = {
  Infos?: KycInfo[]
  Total?: number
}

export type GetAllKycRequest = {
  Limit?: number
  Offset?: number
}

export type GetAllKycResponse = {
  Infos?: KycInfo[]
  Total?: number
}

export type UpdateKycRequest = {
  ID?: string
  UserID?: string
  AppID?: string
  FirstName?: string
  LastName?: string
  Region?: string
  CardType?: string
  CardID?: string
  FrontCardImg?: string
  BackCardImg?: string
  UserHandlingCardImg?: string
}

export type UpdateKycResponse = {
  Info?: KycInfo
}

export type UploadKycImageRequest = {
  UserID?: string
  ImageType?: string
  ImageBase64?: string
  AppID?: string
}

export type UploadKycImageResponse = {
  Info?: string
}

export type GetKycImageRequest = {
  AppID?: string
  UserID?: string
  ImageS3Key?: string
}

export type GetKycImageResponse = {
  Info?: string
}

export type GetKycByKycIDsRequest = {
  KycIDs?: string[]
}

export type GetKycByKycIDsResponse = {
  Infos?: KycInfo[]
}

export class KycManagement {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
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