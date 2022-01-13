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

export type GetQRcodeURLRequest = {
  Username?: string
  UserID?: string
  AppID?: string
}

export type QRCodeInfo = {
  CodeURL?: string
  Secret?: string
}

export type GetQRcodeURLResponse = {
  Info?: QRCodeInfo
}

export type VerifyGoogleAuthRequest = {
  UserID?: string
  Code?: string
  AppID?: string
}

export type VerifyGoogleAuthResponse = {
  Info?: string
}

export type DeleteUserGoogleAuthRequest = {
  UserID?: string
  AppID?: string
}

export type DeleteUserGoogleAuthResponse = {
  Info?: string
}

export type SendEmailRequest = {
  AppID?: string
  Email?: string
  Intention?: string
  Lang?: string
  Username?: string
}

export type SendEmailResponse = {
  Info?: string
}

export type SendSmsRequest = {
  AppID?: string
  Phone?: string
  Lang?: string
  Intention?: string
}

export type SendSmsResponse = {
  Info?: string
}

export type VerifyCodeRequest = {
  Param?: string
  Code?: string
  VerifyType?: string
  AppID?: string
}

export type VerifyCodeResponse = {
  Info?: string
}

export type VerifyCodeWithUserIDRequest = {
  UserID?: string
  Param?: string
  Code?: string
  VerifyType?: string
  AppID?: string
}

export type VerifyCodeWithUserIDResponse = {
  Info?: string
}

export type VerifyGoogleRecaptchaRequest = {
  Response?: string
}

export type VerifyGoogleRecaptchaResponse = {
  Info?: boolean
}

export type GetCaptcherImgRequest = {
  AppID?: string
  UserID?: string
}

export type CaptcherResp = {
  ID?: string
  Base64URL?: string
}

export type GetCaptcherImgResponse = {
  Info?: CaptcherResp
}

export type VerifyCaptcherRequest = {
  ID?: string
  Input?: string
}

export type VerifyCaptcherResponse = {
  Info?: string
}

export type SendUserSiteContactEmailRequest = {
  AppID?: string
  From?: string
  To?: string
  Text?: string
  Subject?: string
  Username?: string
}

export type SendUserSiteContactEmailResponse = {
  Info?: string
}

export class VerificationDoor {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
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