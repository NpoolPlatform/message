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
  langID?: string
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
  langID?: string
  emailAddress?: string
  usedFor?: string
  toUsername?: string
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

export type AppSMSTemplate = {
  id?: string
  appID?: string
  langID?: string
  usedFor?: string
  subject?: string
  message?: string
}

export type CreateAppSMSTemplateRequest = {
  targetLangID?: string
  info?: AppSMSTemplate
}

export type CreateAppSMSTemplateResponse = {
  info?: AppSMSTemplate
}

export type CreateAppSMSTemplateForOtherAppRequest = {
  targetAppID?: string
  targetLangID?: string
  info?: AppSMSTemplate
}

export type CreateAppSMSTemplateForOtherAppResponse = {
  info?: AppSMSTemplate
}

export type GetAppSMSTemplateRequest = {
  id?: string
}

export type GetAppSMSTemplateResponse = {
  info?: AppSMSTemplate
}

export type UpdateAppSMSTemplateRequest = {
  info?: AppSMSTemplate
}

export type UpdateAppSMSTemplateResponse = {
  info?: AppSMSTemplate
}

export type GetAppSMSTemplateByAppLangUsedForRequest = {
  appID?: string
  langID?: string
  usedFor?: string
}

export type GetAppSMSTemplateByAppLangUsedForResponse = {
  info?: AppSMSTemplate
}

export type GetAppSMSTemplatesByAppRequest = {
  appID?: string
}

export type GetAppSMSTemplatesByAppResponse = {
  infos?: AppSMSTemplate[]
}

export type GetAppSMSTemplatesByOtherAppRequest = {
  targetAppID?: string
}

export type GetAppSMSTemplatesByOtherAppResponse = {
  infos?: AppSMSTemplate[]
}

export type AppEmailTemplate = {
  id?: string
  appID?: string
  langID?: string
  usedFor?: string
  sender?: string
  replyTos?: string[]
  cCTos?: string[]
  subject?: string
  body?: string
  defaultToUsername?: string
}

export type CreateAppEmailTemplateRequest = {
  targetLangID?: string
  info?: AppEmailTemplate
}

export type CreateAppEmailTemplateResponse = {
  info?: AppEmailTemplate
}

export type CreateAppEmailTemplateForOtherAppRequest = {
  targetAppID?: string
  targetLangID?: string
  info?: AppEmailTemplate
}

export type CreateAppEmailTemplateForOtherAppResponse = {
  info?: AppEmailTemplate
}

export type GetAppEmailTemplateRequest = {
  id?: string
}

export type GetAppEmailTemplateResponse = {
  info?: AppEmailTemplate
}

export type UpdateAppEmailTemplateRequest = {
  info?: AppEmailTemplate
}

export type UpdateAppEmailTemplateResponse = {
  info?: AppEmailTemplate
}

export type GetAppEmailTemplateByAppLangUsedForRequest = {
  appID?: string
  langID?: string
  usedFor?: string
}

export type GetAppEmailTemplateByAppLangUsedForResponse = {
  info?: AppEmailTemplate
}

export type GetAppEmailTemplatesByAppRequest = {
  appID?: string
}

export type GetAppEmailTemplatesByAppResponse = {
  infos?: AppEmailTemplate[]
}

export type GetAppEmailTemplatesByOtherAppRequest = {
  targetAppID?: string
}

export type GetAppEmailTemplatesByOtherAppResponse = {
  infos?: AppEmailTemplate[]
}

export type AppContact = {
  id?: string
  appID?: string
  usedFor?: string
  account?: string
  accountType?: string
  sender?: string
}

export type CreateAppContactRequest = {
  info?: AppContact
}

export type CreateAppContactResponse = {
  info?: AppContact
}

export type CreateAppContactForOtherAppRequest = {
  targetAppID?: string
  info?: AppContact
}

export type CreateAppContactForOtherAppResponse = {
  info?: AppContact
}

export type GetAppContactRequest = {
  id?: string
}

export type GetAppContactResponse = {
  info?: AppContact
}

export type UpdateAppContactRequest = {
  info?: AppContact
}

export type UpdateAppContactResponse = {
  info?: AppContact
}

export type GetAppContactByAppUsedForAccountTypeRequest = {
  appID?: string
  usedFor?: string
  accountType?: string
}

export type GetAppContactByAppUsedForAccountTypeResponse = {
  info?: AppContact
}

export type GetAppContactsByAppRequest = {
  appID?: string
}

export type GetAppContactsByAppResponse = {
  infos?: AppContact[]
}

export type GetAppContactsByOtherAppRequest = {
  targetAppID?: string
}

export type GetAppContactsByOtherAppResponse = {
  infos?: AppContact[]
}

export type SetupGoogleAuthenticationRequest = {
  appID?: string
  userID?: string
}

export type SetupGoogleAuthenticationResponse = {
  oTPAuth?: string
  secret?: string
}

export type VerifyGoogleAuthenticationRequest = {
  appID?: string
  userID?: string
  code?: string
}

export type VerifyGoogleAuthenticationResponse = {
  code?: number
  message?: string
}

export type VerifyGoogleRecaptchaV3Request = {
  recaptchaToken?: string
}

export type VerifyGoogleRecaptchaV3Response = {
  code?: number
  message?: string
}

export type ContactByEmailRequest = {
  appID?: string
  userID?: string
  usedFor?: string
  sender?: string
  subject?: string
  body?: string
  senderName?: string
}

export type ContactByEmailResponse = {
  code?: number
  message?: string
}

export type NotifyEmailRequest = {
  appID?: string
  userID?: string
  usedFor?: string
  receiverID?: string
  langID?: string
  senderName?: string
  receiverName?: string
}

export type NotifyEmailResponse = {
  code?: number
  message?: string
}

export class ThirdGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppSMSTemplate(req: CreateAppSMSTemplateRequest, initReq?: fm.InitReq): Promise<CreateAppSMSTemplateResponse> {
    return fm.fetchReq<CreateAppSMSTemplateRequest, CreateAppSMSTemplateResponse>(`/v1/create/app/sms/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppSMSTemplateForOtherApp(req: CreateAppSMSTemplateForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppSMSTemplateForOtherAppResponse> {
    return fm.fetchReq<CreateAppSMSTemplateForOtherAppRequest, CreateAppSMSTemplateForOtherAppResponse>(`/v1/create/app/sms/template/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppSMSTemplate(req: GetAppSMSTemplateRequest, initReq?: fm.InitReq): Promise<GetAppSMSTemplateResponse> {
    return fm.fetchReq<GetAppSMSTemplateRequest, GetAppSMSTemplateResponse>(`/v1/get/app/sms/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppSMSTemplate(req: UpdateAppSMSTemplateRequest, initReq?: fm.InitReq): Promise<UpdateAppSMSTemplateResponse> {
    return fm.fetchReq<UpdateAppSMSTemplateRequest, UpdateAppSMSTemplateResponse>(`/v1/update/app/sms/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppSMSTemplatesByApp(req: GetAppSMSTemplatesByAppRequest, initReq?: fm.InitReq): Promise<GetAppSMSTemplatesByAppResponse> {
    return fm.fetchReq<GetAppSMSTemplatesByAppRequest, GetAppSMSTemplatesByAppResponse>(`/v1/get/app/sms/templates/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppSMSTemplatesByOtherApp(req: GetAppSMSTemplatesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppSMSTemplatesByOtherAppResponse> {
    return fm.fetchReq<GetAppSMSTemplatesByOtherAppRequest, GetAppSMSTemplatesByOtherAppResponse>(`/v1/get/app/sms/templates/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppSMSTemplateByAppLangUsedFor(req: GetAppSMSTemplateByAppLangUsedForRequest, initReq?: fm.InitReq): Promise<GetAppSMSTemplateByAppLangUsedForResponse> {
    return fm.fetchReq<GetAppSMSTemplateByAppLangUsedForRequest, GetAppSMSTemplateByAppLangUsedForResponse>(`/v1/get/app/sms/template/by/app/lang/usedfor`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppEmailTemplate(req: CreateAppEmailTemplateRequest, initReq?: fm.InitReq): Promise<CreateAppEmailTemplateResponse> {
    return fm.fetchReq<CreateAppEmailTemplateRequest, CreateAppEmailTemplateResponse>(`/v1/create/app/email/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppEmailTemplateForOtherApp(req: CreateAppEmailTemplateForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppEmailTemplateForOtherAppResponse> {
    return fm.fetchReq<CreateAppEmailTemplateForOtherAppRequest, CreateAppEmailTemplateForOtherAppResponse>(`/v1/create/app/email/template/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppEmailTemplate(req: GetAppEmailTemplateRequest, initReq?: fm.InitReq): Promise<GetAppEmailTemplateResponse> {
    return fm.fetchReq<GetAppEmailTemplateRequest, GetAppEmailTemplateResponse>(`/v1/get/app/email/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppEmailTemplate(req: UpdateAppEmailTemplateRequest, initReq?: fm.InitReq): Promise<UpdateAppEmailTemplateResponse> {
    return fm.fetchReq<UpdateAppEmailTemplateRequest, UpdateAppEmailTemplateResponse>(`/v1/update/app/email/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppEmailTemplatesByApp(req: GetAppEmailTemplatesByAppRequest, initReq?: fm.InitReq): Promise<GetAppEmailTemplatesByAppResponse> {
    return fm.fetchReq<GetAppEmailTemplatesByAppRequest, GetAppEmailTemplatesByAppResponse>(`/v1/get/app/email/templates/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppEmailTemplatesByOtherApp(req: GetAppEmailTemplatesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppEmailTemplatesByOtherAppResponse> {
    return fm.fetchReq<GetAppEmailTemplatesByOtherAppRequest, GetAppEmailTemplatesByOtherAppResponse>(`/v1/get/app/email/templates/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppEmailTemplateByAppLangUsedFor(req: GetAppEmailTemplateByAppLangUsedForRequest, initReq?: fm.InitReq): Promise<GetAppEmailTemplateByAppLangUsedForResponse> {
    return fm.fetchReq<GetAppEmailTemplateByAppLangUsedForRequest, GetAppEmailTemplateByAppLangUsedForResponse>(`/v1/get/app/email/template/by/app/lang/usedfor`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppContact(req: CreateAppContactRequest, initReq?: fm.InitReq): Promise<CreateAppContactResponse> {
    return fm.fetchReq<CreateAppContactRequest, CreateAppContactResponse>(`/v1/create/app/contact`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppContactForOtherApp(req: CreateAppContactForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppContactForOtherAppResponse> {
    return fm.fetchReq<CreateAppContactForOtherAppRequest, CreateAppContactForOtherAppResponse>(`/v1/create/app/contact/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppContact(req: GetAppContactRequest, initReq?: fm.InitReq): Promise<GetAppContactResponse> {
    return fm.fetchReq<GetAppContactRequest, GetAppContactResponse>(`/v1/get/app/contact`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppContact(req: UpdateAppContactRequest, initReq?: fm.InitReq): Promise<UpdateAppContactResponse> {
    return fm.fetchReq<UpdateAppContactRequest, UpdateAppContactResponse>(`/v1/update/app/contact`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppContactsByApp(req: GetAppContactsByAppRequest, initReq?: fm.InitReq): Promise<GetAppContactsByAppResponse> {
    return fm.fetchReq<GetAppContactsByAppRequest, GetAppContactsByAppResponse>(`/v1/get/app/contacts/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppContactsByOtherApp(req: GetAppContactsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppContactsByOtherAppResponse> {
    return fm.fetchReq<GetAppContactsByOtherAppRequest, GetAppContactsByOtherAppResponse>(`/v1/get/app/contacts/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppContactByAppUsedForAccountType(req: GetAppContactByAppUsedForAccountTypeRequest, initReq?: fm.InitReq): Promise<GetAppContactByAppUsedForAccountTypeResponse> {
    return fm.fetchReq<GetAppContactByAppUsedForAccountTypeRequest, GetAppContactByAppUsedForAccountTypeResponse>(`/v1/get/app/contact/by/app/usedfor/accounttype`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static SetupGoogleAuthentication(req: SetupGoogleAuthenticationRequest, initReq?: fm.InitReq): Promise<SetupGoogleAuthenticationResponse> {
    return fm.fetchReq<SetupGoogleAuthenticationRequest, SetupGoogleAuthenticationResponse>(`/v1/setup/google/authentication`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyGoogleAuthentication(req: VerifyGoogleAuthenticationRequest, initReq?: fm.InitReq): Promise<VerifyGoogleAuthenticationResponse> {
    return fm.fetchReq<VerifyGoogleAuthenticationRequest, VerifyGoogleAuthenticationResponse>(`/v1/verify/google/authentication`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyGoogleRecaptchaV3(req: VerifyGoogleRecaptchaV3Request, initReq?: fm.InitReq): Promise<VerifyGoogleRecaptchaV3Response> {
    return fm.fetchReq<VerifyGoogleRecaptchaV3Request, VerifyGoogleRecaptchaV3Response>(`/v1/verify/google/recaptcha/v3`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ContactByEmail(req: ContactByEmailRequest, initReq?: fm.InitReq): Promise<ContactByEmailResponse> {
    return fm.fetchReq<ContactByEmailRequest, ContactByEmailResponse>(`/v1/contact/by/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static NotifyEmail(req: NotifyEmailRequest, initReq?: fm.InitReq): Promise<NotifyEmailResponse> {
    return fm.fetchReq<NotifyEmailRequest, NotifyEmailResponse>(`/v1/notify/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}