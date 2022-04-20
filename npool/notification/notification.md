# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/notification/notification.proto](#npool_notification_notification-proto)
    - [Announcement](#notification-v1-Announcement)
    - [CheckReadUserRequest](#notification-v1-CheckReadUserRequest)
    - [CheckReadUserResponse](#notification-v1-CheckReadUserResponse)
    - [CreateAnnouncementForOtherAppRequest](#notification-v1-CreateAnnouncementForOtherAppRequest)
    - [CreateAnnouncementForOtherAppResponse](#notification-v1-CreateAnnouncementForOtherAppResponse)
    - [CreateAnnouncementRequest](#notification-v1-CreateAnnouncementRequest)
    - [CreateAnnouncementResponse](#notification-v1-CreateAnnouncementResponse)
    - [CreateMailForOtherAppRequest](#notification-v1-CreateMailForOtherAppRequest)
    - [CreateMailForOtherAppResponse](#notification-v1-CreateMailForOtherAppResponse)
    - [CreateMailRequest](#notification-v1-CreateMailRequest)
    - [CreateMailResponse](#notification-v1-CreateMailResponse)
    - [CreateNotificationForAppOtherUserRequest](#notification-v1-CreateNotificationForAppOtherUserRequest)
    - [CreateNotificationForAppOtherUserResponse](#notification-v1-CreateNotificationForAppOtherUserResponse)
    - [CreateNotificationForOtherAppUserRequest](#notification-v1-CreateNotificationForOtherAppUserRequest)
    - [CreateNotificationForOtherAppUserResponse](#notification-v1-CreateNotificationForOtherAppUserResponse)
    - [CreateNotificationRequest](#notification-v1-CreateNotificationRequest)
    - [CreateNotificationResponse](#notification-v1-CreateNotificationResponse)
    - [CreateReadUserRequest](#notification-v1-CreateReadUserRequest)
    - [CreateReadUserResponse](#notification-v1-CreateReadUserResponse)
    - [CreateTemplateForOtherAppRequest](#notification-v1-CreateTemplateForOtherAppRequest)
    - [CreateTemplateForOtherAppResponse](#notification-v1-CreateTemplateForOtherAppResponse)
    - [CreateTemplateRequest](#notification-v1-CreateTemplateRequest)
    - [CreateTemplateResponse](#notification-v1-CreateTemplateResponse)
    - [GetAnnouncementsByAppRequest](#notification-v1-GetAnnouncementsByAppRequest)
    - [GetAnnouncementsByAppResponse](#notification-v1-GetAnnouncementsByAppResponse)
    - [GetAnnouncementsByOtherAppRequest](#notification-v1-GetAnnouncementsByOtherAppRequest)
    - [GetAnnouncementsByOtherAppResponse](#notification-v1-GetAnnouncementsByOtherAppResponse)
    - [GetMailsByAppRequest](#notification-v1-GetMailsByAppRequest)
    - [GetMailsByAppResponse](#notification-v1-GetMailsByAppResponse)
    - [GetMailsByOtherAppRequest](#notification-v1-GetMailsByOtherAppRequest)
    - [GetMailsByOtherAppResponse](#notification-v1-GetMailsByOtherAppResponse)
    - [GetMailsRequest](#notification-v1-GetMailsRequest)
    - [GetMailsResponse](#notification-v1-GetMailsResponse)
    - [GetNotificationsByAppRequest](#notification-v1-GetNotificationsByAppRequest)
    - [GetNotificationsByAppResponse](#notification-v1-GetNotificationsByAppResponse)
    - [GetNotificationsByAppUserRequest](#notification-v1-GetNotificationsByAppUserRequest)
    - [GetNotificationsByAppUserResponse](#notification-v1-GetNotificationsByAppUserResponse)
    - [GetNotificationsByOtherAppRequest](#notification-v1-GetNotificationsByOtherAppRequest)
    - [GetNotificationsByOtherAppResponse](#notification-v1-GetNotificationsByOtherAppResponse)
    - [GetTemplateByAppLangUsedForRequest](#notification-v1-GetTemplateByAppLangUsedForRequest)
    - [GetTemplateByAppLangUsedForResponse](#notification-v1-GetTemplateByAppLangUsedForResponse)
    - [GetTemplateRequest](#notification-v1-GetTemplateRequest)
    - [GetTemplateResponse](#notification-v1-GetTemplateResponse)
    - [GetTemplatesByAppRequest](#notification-v1-GetTemplatesByAppRequest)
    - [GetTemplatesByAppResponse](#notification-v1-GetTemplatesByAppResponse)
    - [GetTemplatesByOtherAppRequest](#notification-v1-GetTemplatesByOtherAppRequest)
    - [GetTemplatesByOtherAppResponse](#notification-v1-GetTemplatesByOtherAppResponse)
    - [Mail](#notification-v1-Mail)
    - [ReadUser](#notification-v1-ReadUser)
    - [Template](#notification-v1-Template)
    - [UpdateAnnouncementRequest](#notification-v1-UpdateAnnouncementRequest)
    - [UpdateAnnouncementResponse](#notification-v1-UpdateAnnouncementResponse)
    - [UpdateMailRequest](#notification-v1-UpdateMailRequest)
    - [UpdateMailResponse](#notification-v1-UpdateMailResponse)
    - [UpdateNotificationRequest](#notification-v1-UpdateNotificationRequest)
    - [UpdateNotificationResponse](#notification-v1-UpdateNotificationResponse)
    - [UpdateTemplateRequest](#notification-v1-UpdateTemplateRequest)
    - [UpdateTemplateResponse](#notification-v1-UpdateTemplateResponse)
    - [UserNotification](#notification-v1-UserNotification)
  
    - [Notification](#notification-v1-Notification)
  
- [npool/notification/notification.proto](#npool_notification_notification-proto)
    - [Announcement](#notification-v1-Announcement)
    - [CheckReadUserRequest](#notification-v1-CheckReadUserRequest)
    - [CheckReadUserResponse](#notification-v1-CheckReadUserResponse)
    - [CreateAnnouncementForOtherAppRequest](#notification-v1-CreateAnnouncementForOtherAppRequest)
    - [CreateAnnouncementForOtherAppResponse](#notification-v1-CreateAnnouncementForOtherAppResponse)
    - [CreateAnnouncementRequest](#notification-v1-CreateAnnouncementRequest)
    - [CreateAnnouncementResponse](#notification-v1-CreateAnnouncementResponse)
    - [CreateMailForOtherAppRequest](#notification-v1-CreateMailForOtherAppRequest)
    - [CreateMailForOtherAppResponse](#notification-v1-CreateMailForOtherAppResponse)
    - [CreateMailRequest](#notification-v1-CreateMailRequest)
    - [CreateMailResponse](#notification-v1-CreateMailResponse)
    - [CreateNotificationForAppOtherUserRequest](#notification-v1-CreateNotificationForAppOtherUserRequest)
    - [CreateNotificationForAppOtherUserResponse](#notification-v1-CreateNotificationForAppOtherUserResponse)
    - [CreateNotificationForOtherAppUserRequest](#notification-v1-CreateNotificationForOtherAppUserRequest)
    - [CreateNotificationForOtherAppUserResponse](#notification-v1-CreateNotificationForOtherAppUserResponse)
    - [CreateNotificationRequest](#notification-v1-CreateNotificationRequest)
    - [CreateNotificationResponse](#notification-v1-CreateNotificationResponse)
    - [CreateReadUserRequest](#notification-v1-CreateReadUserRequest)
    - [CreateReadUserResponse](#notification-v1-CreateReadUserResponse)
    - [CreateTemplateForOtherAppRequest](#notification-v1-CreateTemplateForOtherAppRequest)
    - [CreateTemplateForOtherAppResponse](#notification-v1-CreateTemplateForOtherAppResponse)
    - [CreateTemplateRequest](#notification-v1-CreateTemplateRequest)
    - [CreateTemplateResponse](#notification-v1-CreateTemplateResponse)
    - [GetAnnouncementsByAppRequest](#notification-v1-GetAnnouncementsByAppRequest)
    - [GetAnnouncementsByAppResponse](#notification-v1-GetAnnouncementsByAppResponse)
    - [GetAnnouncementsByOtherAppRequest](#notification-v1-GetAnnouncementsByOtherAppRequest)
    - [GetAnnouncementsByOtherAppResponse](#notification-v1-GetAnnouncementsByOtherAppResponse)
    - [GetMailsByAppRequest](#notification-v1-GetMailsByAppRequest)
    - [GetMailsByAppResponse](#notification-v1-GetMailsByAppResponse)
    - [GetMailsByOtherAppRequest](#notification-v1-GetMailsByOtherAppRequest)
    - [GetMailsByOtherAppResponse](#notification-v1-GetMailsByOtherAppResponse)
    - [GetMailsRequest](#notification-v1-GetMailsRequest)
    - [GetMailsResponse](#notification-v1-GetMailsResponse)
    - [GetNotificationsByAppRequest](#notification-v1-GetNotificationsByAppRequest)
    - [GetNotificationsByAppResponse](#notification-v1-GetNotificationsByAppResponse)
    - [GetNotificationsByAppUserRequest](#notification-v1-GetNotificationsByAppUserRequest)
    - [GetNotificationsByAppUserResponse](#notification-v1-GetNotificationsByAppUserResponse)
    - [GetNotificationsByOtherAppRequest](#notification-v1-GetNotificationsByOtherAppRequest)
    - [GetNotificationsByOtherAppResponse](#notification-v1-GetNotificationsByOtherAppResponse)
    - [GetTemplateByAppLangUsedForRequest](#notification-v1-GetTemplateByAppLangUsedForRequest)
    - [GetTemplateByAppLangUsedForResponse](#notification-v1-GetTemplateByAppLangUsedForResponse)
    - [GetTemplateRequest](#notification-v1-GetTemplateRequest)
    - [GetTemplateResponse](#notification-v1-GetTemplateResponse)
    - [GetTemplatesByAppRequest](#notification-v1-GetTemplatesByAppRequest)
    - [GetTemplatesByAppResponse](#notification-v1-GetTemplatesByAppResponse)
    - [GetTemplatesByOtherAppRequest](#notification-v1-GetTemplatesByOtherAppRequest)
    - [GetTemplatesByOtherAppResponse](#notification-v1-GetTemplatesByOtherAppResponse)
    - [Mail](#notification-v1-Mail)
    - [ReadUser](#notification-v1-ReadUser)
    - [Template](#notification-v1-Template)
    - [UpdateAnnouncementRequest](#notification-v1-UpdateAnnouncementRequest)
    - [UpdateAnnouncementResponse](#notification-v1-UpdateAnnouncementResponse)
    - [UpdateMailRequest](#notification-v1-UpdateMailRequest)
    - [UpdateMailResponse](#notification-v1-UpdateMailResponse)
    - [UpdateNotificationRequest](#notification-v1-UpdateNotificationRequest)
    - [UpdateNotificationResponse](#notification-v1-UpdateNotificationResponse)
    - [UpdateTemplateRequest](#notification-v1-UpdateTemplateRequest)
    - [UpdateTemplateResponse](#notification-v1-UpdateTemplateResponse)
    - [UserNotification](#notification-v1-UserNotification)
  
    - [Notification](#notification-v1-Notification)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool_notification_notification-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/notification/notification.proto



<a name="notification-v1-Announcement"></a>

### Announcement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-CheckReadUserRequest"></a>

### CheckReadUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CheckReadUserResponse"></a>

### CheckReadUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateAnnouncementForOtherAppRequest"></a>

### CreateAnnouncementForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementForOtherAppResponse"></a>

### CreateAnnouncementForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementRequest"></a>

### CreateAnnouncementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementResponse"></a>

### CreateAnnouncementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateMailForOtherAppRequest"></a>

### CreateMailForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailForOtherAppResponse"></a>

### CreateMailForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailRequest"></a>

### CreateMailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailResponse"></a>

### CreateMailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateNotificationForAppOtherUserRequest"></a>

### CreateNotificationForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForAppOtherUserResponse"></a>

### CreateNotificationForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForOtherAppUserRequest"></a>

### CreateNotificationForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForOtherAppUserResponse"></a>

### CreateNotificationForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationRequest"></a>

### CreateNotificationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |
| Message | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| UserName | [string](#string) |  |  |






<a name="notification-v1-CreateNotificationResponse"></a>

### CreateNotificationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateReadUserRequest"></a>

### CreateReadUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateReadUserResponse"></a>

### CreateReadUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateTemplateForOtherAppRequest"></a>

### CreateTemplateForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateForOtherAppResponse"></a>

### CreateTemplateForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateRequest"></a>

### CreateTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateResponse"></a>

### CreateTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetAnnouncementsByAppRequest"></a>

### GetAnnouncementsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetAnnouncementsByAppResponse"></a>

### GetAnnouncementsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Announcement](#notification-v1-Announcement) | repeated |  |






<a name="notification-v1-GetAnnouncementsByOtherAppRequest"></a>

### GetAnnouncementsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetAnnouncementsByOtherAppResponse"></a>

### GetAnnouncementsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Announcement](#notification-v1-Announcement) | repeated |  |






<a name="notification-v1-GetMailsByAppRequest"></a>

### GetMailsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetMailsByAppResponse"></a>

### GetMailsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetMailsByOtherAppRequest"></a>

### GetMailsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetMailsByOtherAppResponse"></a>

### GetMailsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetMailsRequest"></a>

### GetMailsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="notification-v1-GetMailsResponse"></a>

### GetMailsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetNotificationsByAppRequest"></a>

### GetNotificationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByAppResponse"></a>

### GetNotificationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetNotificationsByAppUserRequest"></a>

### GetNotificationsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByAppUserResponse"></a>

### GetNotificationsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetNotificationsByOtherAppRequest"></a>

### GetNotificationsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByOtherAppResponse"></a>

### GetNotificationsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetTemplateByAppLangUsedForRequest"></a>

### GetTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="notification-v1-GetTemplateByAppLangUsedForResponse"></a>

### GetTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetTemplateRequest"></a>

### GetTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="notification-v1-GetTemplateResponse"></a>

### GetTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetTemplatesByAppRequest"></a>

### GetTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetTemplatesByAppResponse"></a>

### GetTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Template](#notification-v1-Template) | repeated |  |






<a name="notification-v1-GetTemplatesByOtherAppRequest"></a>

### GetTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetTemplatesByOtherAppResponse"></a>

### GetTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Template](#notification-v1-Template) | repeated |  |






<a name="notification-v1-Mail"></a>

### Mail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| FromUserID | [string](#string) |  |  |
| ToUserID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| AlreadyRead | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-ReadUser"></a>

### ReadUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| AnnouncementID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-Template"></a>

### Template



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-UpdateAnnouncementRequest"></a>

### UpdateAnnouncementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-UpdateAnnouncementResponse"></a>

### UpdateAnnouncementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-UpdateMailRequest"></a>

### UpdateMailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-UpdateMailResponse"></a>

### UpdateMailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-UpdateNotificationRequest"></a>

### UpdateNotificationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-UpdateNotificationResponse"></a>

### UpdateNotificationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-UpdateTemplateRequest"></a>

### UpdateTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-UpdateTemplateResponse"></a>

### UpdateTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-UserNotification"></a>

### UserNotification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| AlreadyRead | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |





 

 

 


<a name="notification-v1-Notification"></a>

### Notification
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| CreateAnnouncement | [CreateAnnouncementRequest](#notification-v1-CreateAnnouncementRequest) | [CreateAnnouncementResponse](#notification-v1-CreateAnnouncementResponse) |  |
| CreateAnnouncementForOtherApp | [CreateAnnouncementForOtherAppRequest](#notification-v1-CreateAnnouncementForOtherAppRequest) | [CreateAnnouncementForOtherAppResponse](#notification-v1-CreateAnnouncementForOtherAppResponse) |  |
| UpdateAnnouncement | [UpdateAnnouncementRequest](#notification-v1-UpdateAnnouncementRequest) | [UpdateAnnouncementResponse](#notification-v1-UpdateAnnouncementResponse) |  |
| GetAnnouncementsByApp | [GetAnnouncementsByAppRequest](#notification-v1-GetAnnouncementsByAppRequest) | [GetAnnouncementsByAppResponse](#notification-v1-GetAnnouncementsByAppResponse) |  |
| GetAnnouncementsByOtherApp | [GetAnnouncementsByOtherAppRequest](#notification-v1-GetAnnouncementsByOtherAppRequest) | [GetAnnouncementsByOtherAppResponse](#notification-v1-GetAnnouncementsByOtherAppResponse) |  |
| CreateNotification | [CreateNotificationRequest](#notification-v1-CreateNotificationRequest) | [CreateNotificationResponse](#notification-v1-CreateNotificationResponse) |  |
| CreateNotificationForAppOtherUser | [CreateNotificationForAppOtherUserRequest](#notification-v1-CreateNotificationForAppOtherUserRequest) | [CreateNotificationForAppOtherUserResponse](#notification-v1-CreateNotificationForAppOtherUserResponse) |  |
| CreateNotificationForOtherAppUser | [CreateNotificationForOtherAppUserRequest](#notification-v1-CreateNotificationForOtherAppUserRequest) | [CreateNotificationForOtherAppUserResponse](#notification-v1-CreateNotificationForOtherAppUserResponse) |  |
| UpdateNotification | [UpdateNotificationRequest](#notification-v1-UpdateNotificationRequest) | [UpdateNotificationResponse](#notification-v1-UpdateNotificationResponse) |  |
| GetNotificationsByAppUser | [GetNotificationsByAppUserRequest](#notification-v1-GetNotificationsByAppUserRequest) | [GetNotificationsByAppUserResponse](#notification-v1-GetNotificationsByAppUserResponse) |  |
| GetNotificationsByApp | [GetNotificationsByAppRequest](#notification-v1-GetNotificationsByAppRequest) | [GetNotificationsByAppResponse](#notification-v1-GetNotificationsByAppResponse) |  |
| GetNotificationsByOtherApp | [GetNotificationsByOtherAppRequest](#notification-v1-GetNotificationsByOtherAppRequest) | [GetNotificationsByOtherAppResponse](#notification-v1-GetNotificationsByOtherAppResponse) |  |
| CreateReadUser | [CreateReadUserRequest](#notification-v1-CreateReadUserRequest) | [CreateReadUserResponse](#notification-v1-CreateReadUserResponse) |  |
| CheckReadUser | [CheckReadUserRequest](#notification-v1-CheckReadUserRequest) | [CheckReadUserResponse](#notification-v1-CheckReadUserResponse) |  |
| CreateMail | [CreateMailRequest](#notification-v1-CreateMailRequest) | [CreateMailResponse](#notification-v1-CreateMailResponse) |  |
| CreateMailForOtherApp | [CreateMailForOtherAppRequest](#notification-v1-CreateMailForOtherAppRequest) | [CreateMailForOtherAppResponse](#notification-v1-CreateMailForOtherAppResponse) |  |
| UpdateMail | [UpdateMailRequest](#notification-v1-UpdateMailRequest) | [UpdateMailResponse](#notification-v1-UpdateMailResponse) |  |
| GetMails | [GetMailsRequest](#notification-v1-GetMailsRequest) | [GetMailsResponse](#notification-v1-GetMailsResponse) |  |
| GetMailsByApp | [GetMailsByAppRequest](#notification-v1-GetMailsByAppRequest) | [GetMailsByAppResponse](#notification-v1-GetMailsByAppResponse) |  |
| GetMailsByOtherApp | [GetMailsByOtherAppRequest](#notification-v1-GetMailsByOtherAppRequest) | [GetMailsByOtherAppResponse](#notification-v1-GetMailsByOtherAppResponse) |  |
| CreateTemplate | [CreateTemplateRequest](#notification-v1-CreateTemplateRequest) | [CreateTemplateResponse](#notification-v1-CreateTemplateResponse) |  |
| CreateTemplateForOtherApp | [CreateTemplateForOtherAppRequest](#notification-v1-CreateTemplateForOtherAppRequest) | [CreateTemplateForOtherAppResponse](#notification-v1-CreateTemplateForOtherAppResponse) |  |
| GetTemplate | [GetTemplateRequest](#notification-v1-GetTemplateRequest) | [GetTemplateResponse](#notification-v1-GetTemplateResponse) |  |
| UpdateTemplate | [UpdateTemplateRequest](#notification-v1-UpdateTemplateRequest) | [UpdateTemplateResponse](#notification-v1-UpdateTemplateResponse) |  |
| GetTemplatesByApp | [GetTemplatesByAppRequest](#notification-v1-GetTemplatesByAppRequest) | [GetTemplatesByAppResponse](#notification-v1-GetTemplatesByAppResponse) |  |
| GetTemplatesByOtherApp | [GetTemplatesByOtherAppRequest](#notification-v1-GetTemplatesByOtherAppRequest) | [GetTemplatesByOtherAppResponse](#notification-v1-GetTemplatesByOtherAppResponse) |  |
| GetTemplateByAppLangUsedFor | [GetTemplateByAppLangUsedForRequest](#notification-v1-GetTemplateByAppLangUsedForRequest) | [GetTemplateByAppLangUsedForResponse](#notification-v1-GetTemplateByAppLangUsedForResponse) |  |

 



<a name="npool_notification_notification-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/notification/notification.proto



<a name="notification-v1-Announcement"></a>

### Announcement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-CheckReadUserRequest"></a>

### CheckReadUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CheckReadUserResponse"></a>

### CheckReadUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateAnnouncementForOtherAppRequest"></a>

### CreateAnnouncementForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementForOtherAppResponse"></a>

### CreateAnnouncementForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementRequest"></a>

### CreateAnnouncementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateAnnouncementResponse"></a>

### CreateAnnouncementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-CreateMailForOtherAppRequest"></a>

### CreateMailForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailForOtherAppResponse"></a>

### CreateMailForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailRequest"></a>

### CreateMailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateMailResponse"></a>

### CreateMailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-CreateNotificationForAppOtherUserRequest"></a>

### CreateNotificationForAppOtherUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetUserID | [string](#string) |  |  |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForAppOtherUserResponse"></a>

### CreateNotificationForAppOtherUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForOtherAppUserRequest"></a>

### CreateNotificationForOtherAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| TargetUserID | [string](#string) |  |  |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationForOtherAppUserResponse"></a>

### CreateNotificationForOtherAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateNotificationRequest"></a>

### CreateNotificationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |
| Message | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| UserName | [string](#string) |  |  |






<a name="notification-v1-CreateNotificationResponse"></a>

### CreateNotificationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-CreateReadUserRequest"></a>

### CreateReadUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateReadUserResponse"></a>

### CreateReadUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ReadUser](#notification-v1-ReadUser) |  |  |






<a name="notification-v1-CreateTemplateForOtherAppRequest"></a>

### CreateTemplateForOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateForOtherAppResponse"></a>

### CreateTemplateForOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateRequest"></a>

### CreateTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-CreateTemplateResponse"></a>

### CreateTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetAnnouncementsByAppRequest"></a>

### GetAnnouncementsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetAnnouncementsByAppResponse"></a>

### GetAnnouncementsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Announcement](#notification-v1-Announcement) | repeated |  |






<a name="notification-v1-GetAnnouncementsByOtherAppRequest"></a>

### GetAnnouncementsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetAnnouncementsByOtherAppResponse"></a>

### GetAnnouncementsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Announcement](#notification-v1-Announcement) | repeated |  |






<a name="notification-v1-GetMailsByAppRequest"></a>

### GetMailsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetMailsByAppResponse"></a>

### GetMailsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetMailsByOtherAppRequest"></a>

### GetMailsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetMailsByOtherAppResponse"></a>

### GetMailsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetMailsRequest"></a>

### GetMailsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="notification-v1-GetMailsResponse"></a>

### GetMailsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Mail](#notification-v1-Mail) | repeated |  |






<a name="notification-v1-GetNotificationsByAppRequest"></a>

### GetNotificationsByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByAppResponse"></a>

### GetNotificationsByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetNotificationsByAppUserRequest"></a>

### GetNotificationsByAppUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByAppUserResponse"></a>

### GetNotificationsByAppUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetNotificationsByOtherAppRequest"></a>

### GetNotificationsByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetNotificationsByOtherAppResponse"></a>

### GetNotificationsByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [UserNotification](#notification-v1-UserNotification) | repeated |  |






<a name="notification-v1-GetTemplateByAppLangUsedForRequest"></a>

### GetTemplateByAppLangUsedForRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |






<a name="notification-v1-GetTemplateByAppLangUsedForResponse"></a>

### GetTemplateByAppLangUsedForResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetTemplateRequest"></a>

### GetTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |






<a name="notification-v1-GetTemplateResponse"></a>

### GetTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-GetTemplatesByAppRequest"></a>

### GetTemplatesByAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="notification-v1-GetTemplatesByAppResponse"></a>

### GetTemplatesByAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Template](#notification-v1-Template) | repeated |  |






<a name="notification-v1-GetTemplatesByOtherAppRequest"></a>

### GetTemplatesByOtherAppRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TargetAppID | [string](#string) |  |  |






<a name="notification-v1-GetTemplatesByOtherAppResponse"></a>

### GetTemplatesByOtherAppResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [Template](#notification-v1-Template) | repeated |  |






<a name="notification-v1-Mail"></a>

### Mail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| FromUserID | [string](#string) |  |  |
| ToUserID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| AlreadyRead | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-ReadUser"></a>

### ReadUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| AnnouncementID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-Template"></a>

### Template



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| LangID | [string](#string) |  |  |
| UsedFor | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="notification-v1-UpdateAnnouncementRequest"></a>

### UpdateAnnouncementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-UpdateAnnouncementResponse"></a>

### UpdateAnnouncementResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Announcement](#notification-v1-Announcement) |  |  |






<a name="notification-v1-UpdateMailRequest"></a>

### UpdateMailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-UpdateMailResponse"></a>

### UpdateMailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Mail](#notification-v1-Mail) |  |  |






<a name="notification-v1-UpdateNotificationRequest"></a>

### UpdateNotificationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-UpdateNotificationResponse"></a>

### UpdateNotificationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserNotification](#notification-v1-UserNotification) |  |  |






<a name="notification-v1-UpdateTemplateRequest"></a>

### UpdateTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-UpdateTemplateResponse"></a>

### UpdateTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Template](#notification-v1-Template) |  |  |






<a name="notification-v1-UserNotification"></a>

### UserNotification



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| AlreadyRead | [bool](#bool) |  |  |
| CreateAt | [uint32](#uint32) |  |  |
| UpdateAt | [uint32](#uint32) |  |  |





 

 

 


<a name="notification-v1-Notification"></a>

### Notification
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google-protobuf-Empty) | [.npool.v1.VersionResponse](#npool-v1-VersionResponse) | Method Version |
| CreateAnnouncement | [CreateAnnouncementRequest](#notification-v1-CreateAnnouncementRequest) | [CreateAnnouncementResponse](#notification-v1-CreateAnnouncementResponse) |  |
| CreateAnnouncementForOtherApp | [CreateAnnouncementForOtherAppRequest](#notification-v1-CreateAnnouncementForOtherAppRequest) | [CreateAnnouncementForOtherAppResponse](#notification-v1-CreateAnnouncementForOtherAppResponse) |  |
| UpdateAnnouncement | [UpdateAnnouncementRequest](#notification-v1-UpdateAnnouncementRequest) | [UpdateAnnouncementResponse](#notification-v1-UpdateAnnouncementResponse) |  |
| GetAnnouncementsByApp | [GetAnnouncementsByAppRequest](#notification-v1-GetAnnouncementsByAppRequest) | [GetAnnouncementsByAppResponse](#notification-v1-GetAnnouncementsByAppResponse) |  |
| GetAnnouncementsByOtherApp | [GetAnnouncementsByOtherAppRequest](#notification-v1-GetAnnouncementsByOtherAppRequest) | [GetAnnouncementsByOtherAppResponse](#notification-v1-GetAnnouncementsByOtherAppResponse) |  |
| CreateNotification | [CreateNotificationRequest](#notification-v1-CreateNotificationRequest) | [CreateNotificationResponse](#notification-v1-CreateNotificationResponse) |  |
| CreateNotificationForAppOtherUser | [CreateNotificationForAppOtherUserRequest](#notification-v1-CreateNotificationForAppOtherUserRequest) | [CreateNotificationForAppOtherUserResponse](#notification-v1-CreateNotificationForAppOtherUserResponse) |  |
| CreateNotificationForOtherAppUser | [CreateNotificationForOtherAppUserRequest](#notification-v1-CreateNotificationForOtherAppUserRequest) | [CreateNotificationForOtherAppUserResponse](#notification-v1-CreateNotificationForOtherAppUserResponse) |  |
| UpdateNotification | [UpdateNotificationRequest](#notification-v1-UpdateNotificationRequest) | [UpdateNotificationResponse](#notification-v1-UpdateNotificationResponse) |  |
| GetNotificationsByAppUser | [GetNotificationsByAppUserRequest](#notification-v1-GetNotificationsByAppUserRequest) | [GetNotificationsByAppUserResponse](#notification-v1-GetNotificationsByAppUserResponse) |  |
| GetNotificationsByApp | [GetNotificationsByAppRequest](#notification-v1-GetNotificationsByAppRequest) | [GetNotificationsByAppResponse](#notification-v1-GetNotificationsByAppResponse) |  |
| GetNotificationsByOtherApp | [GetNotificationsByOtherAppRequest](#notification-v1-GetNotificationsByOtherAppRequest) | [GetNotificationsByOtherAppResponse](#notification-v1-GetNotificationsByOtherAppResponse) |  |
| CreateReadUser | [CreateReadUserRequest](#notification-v1-CreateReadUserRequest) | [CreateReadUserResponse](#notification-v1-CreateReadUserResponse) |  |
| CheckReadUser | [CheckReadUserRequest](#notification-v1-CheckReadUserRequest) | [CheckReadUserResponse](#notification-v1-CheckReadUserResponse) |  |
| CreateMail | [CreateMailRequest](#notification-v1-CreateMailRequest) | [CreateMailResponse](#notification-v1-CreateMailResponse) |  |
| CreateMailForOtherApp | [CreateMailForOtherAppRequest](#notification-v1-CreateMailForOtherAppRequest) | [CreateMailForOtherAppResponse](#notification-v1-CreateMailForOtherAppResponse) |  |
| UpdateMail | [UpdateMailRequest](#notification-v1-UpdateMailRequest) | [UpdateMailResponse](#notification-v1-UpdateMailResponse) |  |
| GetMails | [GetMailsRequest](#notification-v1-GetMailsRequest) | [GetMailsResponse](#notification-v1-GetMailsResponse) |  |
| GetMailsByApp | [GetMailsByAppRequest](#notification-v1-GetMailsByAppRequest) | [GetMailsByAppResponse](#notification-v1-GetMailsByAppResponse) |  |
| GetMailsByOtherApp | [GetMailsByOtherAppRequest](#notification-v1-GetMailsByOtherAppRequest) | [GetMailsByOtherAppResponse](#notification-v1-GetMailsByOtherAppResponse) |  |
| CreateTemplate | [CreateTemplateRequest](#notification-v1-CreateTemplateRequest) | [CreateTemplateResponse](#notification-v1-CreateTemplateResponse) |  |
| CreateTemplateForOtherApp | [CreateTemplateForOtherAppRequest](#notification-v1-CreateTemplateForOtherAppRequest) | [CreateTemplateForOtherAppResponse](#notification-v1-CreateTemplateForOtherAppResponse) |  |
| GetTemplate | [GetTemplateRequest](#notification-v1-GetTemplateRequest) | [GetTemplateResponse](#notification-v1-GetTemplateResponse) |  |
| UpdateTemplate | [UpdateTemplateRequest](#notification-v1-UpdateTemplateRequest) | [UpdateTemplateResponse](#notification-v1-UpdateTemplateResponse) |  |
| GetTemplatesByApp | [GetTemplatesByAppRequest](#notification-v1-GetTemplatesByAppRequest) | [GetTemplatesByAppResponse](#notification-v1-GetTemplatesByAppResponse) |  |
| GetTemplatesByOtherApp | [GetTemplatesByOtherAppRequest](#notification-v1-GetTemplatesByOtherAppRequest) | [GetTemplatesByOtherAppResponse](#notification-v1-GetTemplatesByOtherAppResponse) |  |
| GetTemplateByAppLangUsedFor | [GetTemplateByAppLangUsedForRequest](#notification-v1-GetTemplateByAppLangUsedForRequest) | [GetTemplateByAppLangUsedForResponse](#notification-v1-GetTemplateByAppLangUsedForResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

