/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  info?: string
}

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

export type UserNotification = {
  id?: string
  appID?: string
  userID?: string
  title?: string
  content?: string
  alreadyRead?: boolean
  createAt?: number
}

export type CreateNotificationRequest = {
  info?: UserNotification
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

export type UpdateMailRequest = {
  info?: Mail
}

export type UpdateMailResponse = {
  info?: Mail
}

export class Notification {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAnnouncement(req: CreateAnnouncementRequest, initReq?: fm.InitReq): Promise<CreateAnnouncementResponse> {
    return fm.fetchReq<CreateAnnouncementRequest, CreateAnnouncementResponse>(`/create/announcement`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAnnouncement(req: UpdateAnnouncementRequest, initReq?: fm.InitReq): Promise<UpdateAnnouncementResponse> {
    return fm.fetchReq<UpdateAnnouncementRequest, UpdateAnnouncementResponse>(`/update/announcement`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAnnouncementsByApp(req: GetAnnouncementsByAppRequest, initReq?: fm.InitReq): Promise<GetAnnouncementsByAppResponse> {
    return fm.fetchReq<GetAnnouncementsByAppRequest, GetAnnouncementsByAppResponse>(`/get/announcements/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateNotification(req: CreateNotificationRequest, initReq?: fm.InitReq): Promise<CreateNotificationResponse> {
    return fm.fetchReq<CreateNotificationRequest, CreateNotificationResponse>(`/create/notification`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateNotification(req: UpdateNotificationRequest, initReq?: fm.InitReq): Promise<UpdateNotificationResponse> {
    return fm.fetchReq<UpdateNotificationRequest, UpdateNotificationResponse>(`/update/notification`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetNotificationsByAppUser(req: GetNotificationsByAppUserRequest, initReq?: fm.InitReq): Promise<GetNotificationsByAppUserResponse> {
    return fm.fetchReq<GetNotificationsByAppUserRequest, GetNotificationsByAppUserResponse>(`/get/notifications/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReadUser(req: CreateReadUserRequest, initReq?: fm.InitReq): Promise<CreateReadUserResponse> {
    return fm.fetchReq<CreateReadUserRequest, CreateReadUserResponse>(`/create/read/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CheckReadUser(req: CheckReadUserRequest, initReq?: fm.InitReq): Promise<CheckReadUserResponse> {
    return fm.fetchReq<CheckReadUserRequest, CheckReadUserResponse>(`/check/read/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMail(req: CreateMailRequest, initReq?: fm.InitReq): Promise<CreateMailResponse> {
    return fm.fetchReq<CreateMailRequest, CreateMailResponse>(`/create/mail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMail(req: UpdateMailRequest, initReq?: fm.InitReq): Promise<UpdateMailResponse> {
    return fm.fetchReq<UpdateMailRequest, UpdateMailResponse>(`/update/mail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}