/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type SendSMSCodeRequest = {
  appID?: string
  userID?: string
  phoneNO?: string
  usedFor?: string
}

export type SendSMSCodeResponse = {
  code?: number
  message?: string
}

export type VerifySMSCodeRequest = {
  appID?: string
  userID?: string
  phoneNO?: string
  usedFor?: string
  code?: string
}

export type VerifySMSCodeResponse = {
  code?: number
  message?: string
}

export type SendEmailCodeRequest = {
  appID?: string
  userID?: string
  emailAddress?: string
  usedFor?: string
}

export type SendEmailCodeResponse = {
  code?: number
  message?: string
}

export type VerifyEmailCodeRequest = {
  appID?: string
  userID?: string
  emailAddress?: string
  usedFor?: string
  code?: string
}

export type VerifyEmailCodeResponse = {
  code?: number
  message?: string
}

export class ThirdGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SendSMSCode(req: SendSMSCodeRequest, initReq?: fm.InitReq): Promise<SendSMSCodeResponse> {
    return fm.fetchReq<SendSMSCodeRequest, SendSMSCodeResponse>(`/v1/send/sms/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifySMSCode(req: VerifySMSCodeRequest, initReq?: fm.InitReq): Promise<VerifySMSCodeResponse> {
    return fm.fetchReq<VerifySMSCodeRequest, VerifySMSCodeResponse>(`/v1/verify/sms/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SendEmailCode(req: SendEmailCodeRequest, initReq?: fm.InitReq): Promise<SendEmailCodeResponse> {
    return fm.fetchReq<SendEmailCodeRequest, SendEmailCodeResponse>(`/v1/send/email/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyEmailCode(req: VerifyEmailCodeRequest, initReq?: fm.InitReq): Promise<VerifyEmailCodeResponse> {
    return fm.fetchReq<VerifyEmailCodeRequest, VerifyEmailCodeResponse>(`/v1/verify/email/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}