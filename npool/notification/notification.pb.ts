/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Announcement = {
  id?: string
  appID?: string
  title?: string
  content?: string
  createAt?: number
}

export type CreateAnnouncementRequest = {
  info?: Announcement
}

export type CreateAnnouncementResponse = {
  info?: Announcement
}

export type CreateAnnouncementForOtherAppRequest = {
  targetAppID?: string
  info?: Announcement
}

export type CreateAnnouncementForOtherAppResponse = {
  info?: Announcement
}

export type UpdateAnnouncementRequest = {
  info?: Announcement
}

export type UpdateAnnouncementResponse = {
  info?: Announcement
}

export type GetAnnouncementsByAppRequest = {
  appID?: string
}

export type GetAnnouncementsByAppResponse = {
  infos?: Announcement[]
}

export type GetAnnouncementsByOtherAppRequest = {
  targetAppID?: string
}

export type GetAnnouncementsByOtherAppResponse = {
  infos?: Announcement[]
}

export type UserNotification = {
  id?: string
  appID?: string
  userID?: string
  title?: string
  content?: string
  alreadyRead?: boolean
  createAt?: number
  updateAt?: number
}

export type CreateNotificationRequest = {
  info?: UserNotification
  message?: string
  langID?: string
  usedFor?: string
  userName?: string
}

export type CreateNotificationResponse = {
  info?: UserNotification
}

export type UpdateNotificationRequest = {
  info?: UserNotification
}

export type UpdateNotificationResponse = {
  info?: UserNotification
}

export type GetNotificationsByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetNotificationsByAppUserResponse = {
  infos?: UserNotification[]
}

export type ReadUser = {
  id?: string
  appID?: string
  announcementID?: string
  userID?: string
  createAt?: number
}

export type CreateReadUserRequest = {
  info?: ReadUser
}

export type CreateReadUserResponse = {
  info?: ReadUser
}

export type CheckReadUserRequest = {
  info?: ReadUser
}

export type CheckReadUserResponse = {
  info?: ReadUser
}

export type Mail = {
  id?: string
  appID?: string
  fromUserID?: string
  toUserID?: string
  title?: string
  content?: string
  alreadyRead?: boolean
  createAt?: number
}

export type CreateMailRequest = {
  info?: Mail
}

export type CreateMailResponse = {
  info?: Mail
}

export type CreateMailForOtherAppRequest = {
  targetAppID?: string
  info?: Mail
}

export type CreateMailForOtherAppResponse = {
  info?: Mail
}

export type UpdateMailRequest = {
  info?: Mail
}

export type UpdateMailResponse = {
  info?: Mail
}

export type GetMailsRequest = {
  appID?: string
  userID?: string
}

export type GetMailsResponse = {
  infos?: Mail[]
}

export type GetMailsByAppRequest = {
  appID?: string
}

export type GetMailsByAppResponse = {
  infos?: Mail[]
}

export type GetMailsByOtherAppRequest = {
  targetAppID?: string
}

export type GetMailsByOtherAppResponse = {
  infos?: Mail[]
}

export type Template = {
  id?: string
  appID?: string
  langID?: string
  usedFor?: string
  title?: string
  content?: string
  createAt?: number
}

export type CreateTemplateRequest = {
  info?: Template
}

export type CreateTemplateResponse = {
  info?: Template
}

export type CreateTemplateForOtherAppRequest = {
  targetAppID?: string
  info?: Template
}

export type CreateTemplateForOtherAppResponse = {
  info?: Template
}

export type GetTemplateRequest = {
  id?: string
}

export type GetTemplateResponse = {
  info?: Template
}

export type UpdateTemplateRequest = {
  info?: Template
}

export type UpdateTemplateResponse = {
  info?: Template
}

export type GetTemplateByAppLangUsedForRequest = {
  appID?: string
  langID?: string
  usedFor?: string
}

export type GetTemplateByAppLangUsedForResponse = {
  info?: Template
}

export type GetTemplatesByAppRequest = {
  appID?: string
}

export type GetTemplatesByAppResponse = {
  infos?: Template[]
}

export type GetTemplatesByOtherAppRequest = {
  targetAppID?: string
}

export type GetTemplatesByOtherAppResponse = {
  infos?: Template[]
}

export class Notification {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAnnouncement(req: CreateAnnouncementRequest, initReq?: fm.InitReq): Promise<CreateAnnouncementResponse> {
    return fm.fetchReq<CreateAnnouncementRequest, CreateAnnouncementResponse>(`/v1/create/announcement`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAnnouncementForOtherApp(req: CreateAnnouncementForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAnnouncementForOtherAppResponse> {
    return fm.fetchReq<CreateAnnouncementForOtherAppRequest, CreateAnnouncementForOtherAppResponse>(`/v1/create/announcement/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAnnouncement(req: UpdateAnnouncementRequest, initReq?: fm.InitReq): Promise<UpdateAnnouncementResponse> {
    return fm.fetchReq<UpdateAnnouncementRequest, UpdateAnnouncementResponse>(`/v1/update/announcement`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAnnouncementsByApp(req: GetAnnouncementsByAppRequest, initReq?: fm.InitReq): Promise<GetAnnouncementsByAppResponse> {
    return fm.fetchReq<GetAnnouncementsByAppRequest, GetAnnouncementsByAppResponse>(`/v1/get/announcements/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAnnouncementsByOtherApp(req: GetAnnouncementsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAnnouncementsByOtherAppResponse> {
    return fm.fetchReq<GetAnnouncementsByOtherAppRequest, GetAnnouncementsByOtherAppResponse>(`/v1/get/announcements/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateNotification(req: CreateNotificationRequest, initReq?: fm.InitReq): Promise<CreateNotificationResponse> {
    return fm.fetchReq<CreateNotificationRequest, CreateNotificationResponse>(`/v1/create/notification`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateNotification(req: UpdateNotificationRequest, initReq?: fm.InitReq): Promise<UpdateNotificationResponse> {
    return fm.fetchReq<UpdateNotificationRequest, UpdateNotificationResponse>(`/v1/update/notification`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetNotificationsByAppUser(req: GetNotificationsByAppUserRequest, initReq?: fm.InitReq): Promise<GetNotificationsByAppUserResponse> {
    return fm.fetchReq<GetNotificationsByAppUserRequest, GetNotificationsByAppUserResponse>(`/v1/get/notifications/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReadUser(req: CreateReadUserRequest, initReq?: fm.InitReq): Promise<CreateReadUserResponse> {
    return fm.fetchReq<CreateReadUserRequest, CreateReadUserResponse>(`/v1/create/read/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CheckReadUser(req: CheckReadUserRequest, initReq?: fm.InitReq): Promise<CheckReadUserResponse> {
    return fm.fetchReq<CheckReadUserRequest, CheckReadUserResponse>(`/v1/check/read/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMail(req: CreateMailRequest, initReq?: fm.InitReq): Promise<CreateMailResponse> {
    return fm.fetchReq<CreateMailRequest, CreateMailResponse>(`/v1/create/mail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMailForOtherApp(req: CreateMailForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateMailForOtherAppResponse> {
    return fm.fetchReq<CreateMailForOtherAppRequest, CreateMailForOtherAppResponse>(`/v1/create/mail/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMail(req: UpdateMailRequest, initReq?: fm.InitReq): Promise<UpdateMailResponse> {
    return fm.fetchReq<UpdateMailRequest, UpdateMailResponse>(`/v1/update/mail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMails(req: GetMailsRequest, initReq?: fm.InitReq): Promise<GetMailsResponse> {
    return fm.fetchReq<GetMailsRequest, GetMailsResponse>(`/v1/get/mails`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMailsByApp(req: GetMailsByAppRequest, initReq?: fm.InitReq): Promise<GetMailsByAppResponse> {
    return fm.fetchReq<GetMailsByAppRequest, GetMailsByAppResponse>(`/v1/get/mails/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMailsByOtherApp(req: GetMailsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetMailsByOtherAppResponse> {
    return fm.fetchReq<GetMailsByOtherAppRequest, GetMailsByOtherAppResponse>(`/v1/get/mails/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateTemplate(req: CreateTemplateRequest, initReq?: fm.InitReq): Promise<CreateTemplateResponse> {
    return fm.fetchReq<CreateTemplateRequest, CreateTemplateResponse>(`/v1/create/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateTemplateForOtherApp(req: CreateTemplateForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateTemplateForOtherAppResponse> {
    return fm.fetchReq<CreateTemplateForOtherAppRequest, CreateTemplateForOtherAppResponse>(`/v1/create/template/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTemplate(req: GetTemplateRequest, initReq?: fm.InitReq): Promise<GetTemplateResponse> {
    return fm.fetchReq<GetTemplateRequest, GetTemplateResponse>(`/v1/get/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateTemplate(req: UpdateTemplateRequest, initReq?: fm.InitReq): Promise<UpdateTemplateResponse> {
    return fm.fetchReq<UpdateTemplateRequest, UpdateTemplateResponse>(`/v1/update/template`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTemplatesByApp(req: GetTemplatesByAppRequest, initReq?: fm.InitReq): Promise<GetTemplatesByAppResponse> {
    return fm.fetchReq<GetTemplatesByAppRequest, GetTemplatesByAppResponse>(`/v1/get/templates/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTemplatesByOtherApp(req: GetTemplatesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetTemplatesByOtherAppResponse> {
    return fm.fetchReq<GetTemplatesByOtherAppRequest, GetTemplatesByOtherAppResponse>(`/v1/get/templates/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTemplateByAppLangUsedFor(req: GetTemplateByAppLangUsedForRequest, initReq?: fm.InitReq): Promise<GetTemplateByAppLangUsedForResponse> {
    return fm.fetchReq<GetTemplateByAppLangUsedForRequest, GetTemplateByAppLangUsedForResponse>(`/v1/get/template/by/app/lang/usedfor`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}