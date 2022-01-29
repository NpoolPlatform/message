/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type GetQRcodeURLRequest = {
  username?: string
  userID?: string
  appID?: string
}

export type QRCodeInfo = {
  codeURL?: string
  secret?: string
}

export type GetQRcodeURLResponse = {
  info?: QRCodeInfo
}

export type VerifyGoogleAuthRequest = {
  userID?: string
  code?: string
  appID?: string
}

export type VerifyGoogleAuthResponse = {
  info?: string
}

export type DeleteUserGoogleAuthRequest = {
  userID?: string
  appID?: string
}

export type DeleteUserGoogleAuthResponse = {
  info?: string
}

export type SendEmailRequest = {
  appID?: string
  email?: string
  intention?: string
  lang?: string
  username?: string
}

export type SendEmailResponse = {
  info?: string
}

export type SendSmsRequest = {
  appID?: string
  phone?: string
  lang?: string
  intention?: string
}

export type SendSmsResponse = {
  info?: string
}

export type VerifyCodeRequest = {
  param?: string
  code?: string
  verifyType?: string
  appID?: string
}

export type VerifyCodeResponse = {
  info?: string
}

export type VerifyCodeWithUserIDRequest = {
  userID?: string
  param?: string
  code?: string
  verifyType?: string
  appID?: string
}

export type VerifyCodeWithUserIDResponse = {
  info?: string
}

export type VerifyGoogleRecaptchaRequest = {
  response?: string
}

export type VerifyGoogleRecaptchaResponse = {
  info?: boolean
}

export type GetCaptcherImgRequest = {
  appID?: string
  userID?: string
}

export type CaptcherResp = {
  id?: string
  base64URL?: string
}

export type GetCaptcherImgResponse = {
  info?: CaptcherResp
}

export type VerifyCaptcherRequest = {
  id?: string
  input?: string
}

export type VerifyCaptcherResponse = {
  info?: string
}

export type SendUserSiteContactEmailRequest = {
  appID?: string
  from?: string
  to?: string
  text?: string
  subject?: string
  username?: string
}

export type SendUserSiteContactEmailResponse = {
  info?: string
}

export class VerificationDoor {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetQRcodeURL(req: GetQRcodeURLRequest, initReq?: fm.InitReq): Promise<GetQRcodeURLResponse> {
    return fm.fetchReq<GetQRcodeURLRequest, GetQRcodeURLResponse>(`/v1/get/qrcode/url`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyGoogleAuth(req: VerifyGoogleAuthRequest, initReq?: fm.InitReq): Promise<VerifyGoogleAuthResponse> {
    return fm.fetchReq<VerifyGoogleAuthRequest, VerifyGoogleAuthResponse>(`/v1/verify/google/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteUserGoogleAuth(req: DeleteUserGoogleAuthRequest, initReq?: fm.InitReq): Promise<DeleteUserGoogleAuthResponse> {
    return fm.fetchReq<DeleteUserGoogleAuthRequest, DeleteUserGoogleAuthResponse>(`/v1/delete/user/google/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SendEmail(req: SendEmailRequest, initReq?: fm.InitReq): Promise<SendEmailResponse> {
    return fm.fetchReq<SendEmailRequest, SendEmailResponse>(`/v1/send/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SendSms(req: SendSmsRequest, initReq?: fm.InitReq): Promise<SendSmsResponse> {
    return fm.fetchReq<SendSmsRequest, SendSmsResponse>(`/v1/send/sms`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyCodeWithUserID(req: VerifyCodeWithUserIDRequest, initReq?: fm.InitReq): Promise<VerifyCodeWithUserIDResponse> {
    return fm.fetchReq<VerifyCodeWithUserIDRequest, VerifyCodeWithUserIDResponse>(`/v1/verify/code/with/userid`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyCode(req: VerifyCodeRequest, initReq?: fm.InitReq): Promise<VerifyCodeResponse> {
    return fm.fetchReq<VerifyCodeRequest, VerifyCodeResponse>(`/v1/verify/code`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyGoogleRecaptcha(req: VerifyGoogleRecaptchaRequest, initReq?: fm.InitReq): Promise<VerifyGoogleRecaptchaResponse> {
    return fm.fetchReq<VerifyGoogleRecaptchaRequest, VerifyGoogleRecaptchaResponse>(`/v1/verify/google/recaptcha`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCaptcherImg(req: GetCaptcherImgRequest, initReq?: fm.InitReq): Promise<GetCaptcherImgResponse> {
    return fm.fetchReq<GetCaptcherImgRequest, GetCaptcherImgResponse>(`/v1/get/captcher/img`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyCaptcher(req: VerifyCaptcherRequest, initReq?: fm.InitReq): Promise<VerifyCaptcherResponse> {
    return fm.fetchReq<VerifyCaptcherRequest, VerifyCaptcherResponse>(`/v1/verify/captcher`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SendUserSiteContactEmail(req: SendUserSiteContactEmailRequest, initReq?: fm.InitReq): Promise<SendUserSiteContactEmailResponse> {
    return fm.fetchReq<SendUserSiteContactEmailRequest, SendUserSiteContactEmailResponse>(`/v1/send/user/site/contact/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}