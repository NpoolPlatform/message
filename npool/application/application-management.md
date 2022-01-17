# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/application/application-management.proto](#npool/application/application-management.proto)
    - [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest)
    - [AddGroupUsersResponse](#application.management.v1.AddGroupUsersResponse)
    - [AddUserLoginTimeRequest](#application.management.v1.AddUserLoginTimeRequest)
    - [AddUserLoginTimeResponse](#application.management.v1.AddUserLoginTimeResponse)
    - [AddUsersToApplicationRequest](#application.management.v1.AddUsersToApplicationRequest)
    - [AddUsersToApplicationResponse](#application.management.v1.AddUsersToApplicationResponse)
    - [ApplicationInfo](#application.management.v1.ApplicationInfo)
    - [ApplicationUserDetail](#application.management.v1.ApplicationUserDetail)
    - [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo)
    - [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest)
    - [CreateApplicationResponse](#application.management.v1.CreateApplicationResponse)
    - [CreateGroupRequest](#application.management.v1.CreateGroupRequest)
    - [CreateGroupResponse](#application.management.v1.CreateGroupResponse)
    - [CreateResourceRequest](#application.management.v1.CreateResourceRequest)
    - [CreateResourceResponse](#application.management.v1.CreateResourceResponse)
    - [CreateRoleRequest](#application.management.v1.CreateRoleRequest)
    - [CreateRoleResponse](#application.management.v1.CreateRoleResponse)
    - [CreatorResource](#application.management.v1.CreatorResource)
    - [CreatorRole](#application.management.v1.CreatorRole)
    - [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest)
    - [DeleteApplicationResponse](#application.management.v1.DeleteApplicationResponse)
    - [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest)
    - [DeleteGroupResponse](#application.management.v1.DeleteGroupResponse)
    - [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest)
    - [DeleteResourceResponse](#application.management.v1.DeleteResourceResponse)
    - [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest)
    - [DeleteRoleResponse](#application.management.v1.DeleteRoleResponse)
    - [GetAllGroupsRequest](#application.management.v1.GetAllGroupsRequest)
    - [GetAllGroupsResponse](#application.management.v1.GetAllGroupsResponse)
    - [GetApplicationByOwnerRequest](#application.management.v1.GetApplicationByOwnerRequest)
    - [GetApplicationByOwnerResponse](#application.management.v1.GetApplicationByOwnerResponse)
    - [GetApplicationRequest](#application.management.v1.GetApplicationRequest)
    - [GetApplicationResponse](#application.management.v1.GetApplicationResponse)
    - [GetApplicationUserDetailRequest](#application.management.v1.GetApplicationUserDetailRequest)
    - [GetApplicationUserDetailResponse](#application.management.v1.GetApplicationUserDetailResponse)
    - [GetApplicationsRequest](#application.management.v1.GetApplicationsRequest)
    - [GetApplicationsResponse](#application.management.v1.GetApplicationsResponse)
    - [GetGroupByOwnerRequest](#application.management.v1.GetGroupByOwnerRequest)
    - [GetGroupByOwnerResponse](#application.management.v1.GetGroupByOwnerResponse)
    - [GetGroupRequest](#application.management.v1.GetGroupRequest)
    - [GetGroupResponse](#application.management.v1.GetGroupResponse)
    - [GetGroupUsersRequest](#application.management.v1.GetGroupUsersRequest)
    - [GetGroupUsersResponse](#application.management.v1.GetGroupUsersResponse)
    - [GetResourceByCreatorRequest](#application.management.v1.GetResourceByCreatorRequest)
    - [GetResourceByCreatorResponse](#application.management.v1.GetResourceByCreatorResponse)
    - [GetResourceByNameRequest](#application.management.v1.GetResourceByNameRequest)
    - [GetResourceByNameResponse](#application.management.v1.GetResourceByNameResponse)
    - [GetResourceRequest](#application.management.v1.GetResourceRequest)
    - [GetResourceResponse](#application.management.v1.GetResourceResponse)
    - [GetResourcesRequest](#application.management.v1.GetResourcesRequest)
    - [GetResourcesResponse](#application.management.v1.GetResourcesResponse)
    - [GetRoleByCreatorRequest](#application.management.v1.GetRoleByCreatorRequest)
    - [GetRoleByCreatorResponse](#application.management.v1.GetRoleByCreatorResponse)
    - [GetRoleRequest](#application.management.v1.GetRoleRequest)
    - [GetRoleResponse](#application.management.v1.GetRoleResponse)
    - [GetRoleUsersRequest](#application.management.v1.GetRoleUsersRequest)
    - [GetRoleUsersResponse](#application.management.v1.GetRoleUsersResponse)
    - [GetRolesRequest](#application.management.v1.GetRolesRequest)
    - [GetRolesResponse](#application.management.v1.GetRolesResponse)
    - [GetUserAppIDRequest](#application.management.v1.GetUserAppIDRequest)
    - [GetUserAppIDResponse](#application.management.v1.GetUserAppIDResponse)
    - [GetUserFromApplicationRequest](#application.management.v1.GetUserFromApplicationRequest)
    - [GetUserFromApplicationResponse](#application.management.v1.GetUserFromApplicationResponse)
    - [GetUserGroupRequest](#application.management.v1.GetUserGroupRequest)
    - [GetUserGroupResponse](#application.management.v1.GetUserGroupResponse)
    - [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest)
    - [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse)
    - [GetUsersFromApplicationRequest](#application.management.v1.GetUsersFromApplicationRequest)
    - [GetUsersFromApplicationResponse](#application.management.v1.GetUsersFromApplicationResponse)
    - [GroupInfo](#application.management.v1.GroupInfo)
    - [GroupUserInfo](#application.management.v1.GroupUserInfo)
    - [OwnerApplication](#application.management.v1.OwnerApplication)
    - [OwnerGroup](#application.management.v1.OwnerGroup)
    - [RemoveGroupUsersRequest](#application.management.v1.RemoveGroupUsersRequest)
    - [RemoveGroupUsersResponse](#application.management.v1.RemoveGroupUsersResponse)
    - [RemoveUsersFromApplicationRequest](#application.management.v1.RemoveUsersFromApplicationRequest)
    - [RemoveUsersFromApplicationResponse](#application.management.v1.RemoveUsersFromApplicationResponse)
    - [ResourceInfo](#application.management.v1.ResourceInfo)
    - [RoleInfo](#application.management.v1.RoleInfo)
    - [RoleUserInfo](#application.management.v1.RoleUserInfo)
    - [SetGALoginRequest](#application.management.v1.SetGALoginRequest)
    - [SetGALoginResponse](#application.management.v1.SetGALoginResponse)
    - [SetSMSLoginRequest](#application.management.v1.SetSMSLoginRequest)
    - [SetSMSLoginResponse](#application.management.v1.SetSMSLoginResponse)
    - [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest)
    - [SetUserRoleResponse](#application.management.v1.SetUserRoleResponse)
    - [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest)
    - [UnSetUserRoleResponse](#application.management.v1.UnSetUserRoleResponse)
    - [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest)
    - [UpdateApplicationResponse](#application.management.v1.UpdateApplicationResponse)
    - [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest)
    - [UpdateGroupResponse](#application.management.v1.UpdateGroupResponse)
    - [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest)
    - [UpdateResourceResponse](#application.management.v1.UpdateResourceResponse)
    - [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest)
    - [UpdateRoleResponse](#application.management.v1.UpdateRoleResponse)
    - [UpdateUserGAStatusRequest](#application.management.v1.UpdateUserGAStatusRequest)
    - [UpdateUserGAStatusResponse](#application.management.v1.UpdateUserGAStatusResponse)
    - [UpdateUserKYCStatusRequest](#application.management.v1.UpdateUserKYCStatusRequest)
    - [UpdateUserKYCStatusResponse](#application.management.v1.UpdateUserKYCStatusResponse)
    - [UserRole](#application.management.v1.UserRole)
  
    - [ApplicationManagement](#application.management.v1.ApplicationManagement)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/application/application-management.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/application/application-management.proto



<a name="application.management.v1.AddGroupUsersRequest"></a>

### AddGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.AddGroupUsersResponse"></a>

### AddGroupUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |






<a name="application.management.v1.AddUserLoginTimeRequest"></a>

### AddUserLoginTimeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.AddUserLoginTimeResponse"></a>

### AddUserLoginTimeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [uint32](#uint32) |  |  |






<a name="application.management.v1.AddUsersToApplicationRequest"></a>

### AddUsersToApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| AppID | [string](#string) |  |  |
| Original | [bool](#bool) |  |  |






<a name="application.management.v1.AddUsersToApplicationResponse"></a>

### AddUsersToApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) | repeated |  |






<a name="application.management.v1.ApplicationInfo"></a>

### ApplicationInfo
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ApplicationName | [string](#string) |  |  |
| ApplicationOwner | [string](#string) |  |  |
| HomepageUrl | [string](#string) |  |  |
| RedirectUrl | [string](#string) |  |  |
| ApplicationLogo | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |
| UpdateAT | [uint32](#uint32) |  |  |
| ClientSecret | [string](#string) |  |  |
| GoogleRecaptcha | [bool](#bool) |  |  |
| SmsLogin | [bool](#bool) |  |  |
| InvitationCodeMust | [bool](#bool) |  |  |






<a name="application.management.v1.ApplicationUserDetail"></a>

### ApplicationUserDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserApplicationInfo | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) |  |  |
| UserGroupInfos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |
| UserRoleInfo | [UserRole](#application.management.v1.UserRole) |  |  |






<a name="application.management.v1.ApplicationUserInfo"></a>

### ApplicationUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Original | [bool](#bool) |  |  |
| KycVerify | [bool](#bool) |  |  |
| GAVerify | [bool](#bool) |  |  |
| GALogin | [bool](#bool) |  |  |
| SMSLogin | [bool](#bool) |  |  |
| LoginNumber | [uint32](#uint32) |  |  |
| CreateAT | [uint32](#uint32) |  |  |






<a name="application.management.v1.CreateApplicationRequest"></a>

### CreateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.CreateApplicationResponse"></a>

### CreateApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.CreateGroupRequest"></a>

### CreateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.CreateGroupResponse"></a>

### CreateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.CreateResourceResponse"></a>

### CreateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.CreateRoleResponse"></a>

### CreateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.CreatorResource"></a>

### CreatorResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| Infos | [ResourceInfo](#application.management.v1.ResourceInfo) | repeated |  |






<a name="application.management.v1.CreatorRole"></a>

### CreatorRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| Infos | [RoleInfo](#application.management.v1.RoleInfo) | repeated |  |






<a name="application.management.v1.DeleteApplicationRequest"></a>

### DeleteApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteApplicationResponse"></a>

### DeleteApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupRequest"></a>

### DeleteGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteGroupResponse"></a>

### DeleteGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteResourceResponse"></a>

### DeleteResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| RoleID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.DeleteRoleResponse"></a>

### DeleteRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.GetAllGroupsRequest"></a>

### GetAllGroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetAllGroupsResponse"></a>

### GetAllGroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupInfo](#application.management.v1.GroupInfo) | repeated |  |






<a name="application.management.v1.GetApplicationByOwnerRequest"></a>

### GetApplicationByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Owner | [string](#string) |  |  |






<a name="application.management.v1.GetApplicationByOwnerResponse"></a>

### GetApplicationByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OwnerApplication](#application.management.v1.OwnerApplication) |  |  |






<a name="application.management.v1.GetApplicationRequest"></a>

### GetApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetApplicationResponse"></a>

### GetApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.GetApplicationUserDetailRequest"></a>

### GetApplicationUserDetailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="application.management.v1.GetApplicationUserDetailResponse"></a>

### GetApplicationUserDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationUserDetail](#application.management.v1.ApplicationUserDetail) |  |  |






<a name="application.management.v1.GetApplicationsRequest"></a>

### GetApplicationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [npool.v1.PageInfo](#npool.v1.PageInfo) |  |  |






<a name="application.management.v1.GetApplicationsResponse"></a>

### GetApplicationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationInfo](#application.management.v1.ApplicationInfo) | repeated |  |






<a name="application.management.v1.GetGroupByOwnerRequest"></a>

### GetGroupByOwnerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Owner | [string](#string) |  |  |






<a name="application.management.v1.GetGroupByOwnerResponse"></a>

### GetGroupByOwnerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [OwnerGroup](#application.management.v1.OwnerGroup) |  |  |






<a name="application.management.v1.GetGroupRequest"></a>

### GetGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetGroupResponse"></a>

### GetGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.GetGroupUsersRequest"></a>

### GetGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetGroupUsersResponse"></a>

### GetGroupUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |






<a name="application.management.v1.GetResourceByCreatorRequest"></a>

### GetResourceByCreatorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Creator | [string](#string) |  |  |






<a name="application.management.v1.GetResourceByCreatorResponse"></a>

### GetResourceByCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CreatorResource](#application.management.v1.CreatorResource) |  |  |






<a name="application.management.v1.GetResourceByNameRequest"></a>

### GetResourceByNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceName | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetResourceByNameResponse"></a>

### GetResourceByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.GetResourceRequest"></a>

### GetResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ResourceID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetResourceResponse"></a>

### GetResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.GetResourcesRequest"></a>

### GetResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetResourcesResponse"></a>

### GetResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ResourceInfo](#application.management.v1.ResourceInfo) | repeated |  |






<a name="application.management.v1.GetRoleByCreatorRequest"></a>

### GetRoleByCreatorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Creator | [string](#string) |  |  |






<a name="application.management.v1.GetRoleByCreatorResponse"></a>

### GetRoleByCreatorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [CreatorRole](#application.management.v1.CreatorRole) |  |  |






<a name="application.management.v1.GetRoleRequest"></a>

### GetRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="application.management.v1.GetRoleResponse"></a>

### GetRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.GetRoleUsersRequest"></a>

### GetRoleUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |






<a name="application.management.v1.GetRoleUsersResponse"></a>

### GetRoleUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleUserInfo](#application.management.v1.RoleUserInfo) | repeated |  |






<a name="application.management.v1.GetRolesRequest"></a>

### GetRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleInfo](#application.management.v1.RoleInfo) | repeated |  |






<a name="application.management.v1.GetUserAppIDRequest"></a>

### GetUserAppIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |






<a name="application.management.v1.GetUserAppIDResponse"></a>

### GetUserAppIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [string](#string) | repeated |  |






<a name="application.management.v1.GetUserFromApplicationRequest"></a>

### GetUserFromApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="application.management.v1.GetUserFromApplicationResponse"></a>

### GetUserFromApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) |  |  |






<a name="application.management.v1.GetUserGroupRequest"></a>

### GetUserGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |






<a name="application.management.v1.GetUserGroupResponse"></a>

### GetUserGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [GroupUserInfo](#application.management.v1.GroupUserInfo) | repeated |  |






<a name="application.management.v1.GetUserRoleRequest"></a>

### GetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetUserRoleResponse"></a>

### GetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [UserRole](#application.management.v1.UserRole) |  |  |






<a name="application.management.v1.GetUsersFromApplicationRequest"></a>

### GetUsersFromApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.GetUsersFromApplicationResponse"></a>

### GetUsersFromApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [ApplicationUserInfo](#application.management.v1.ApplicationUserInfo) | repeated |  |






<a name="application.management.v1.GroupInfo"></a>

### GroupInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| GroupName | [string](#string) |  |  |
| GroupOwner | [string](#string) |  |  |
| GroupLogo | [string](#string) |  |  |
| Annotation | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |
| UpdateAT | [uint32](#uint32) |  |  |






<a name="application.management.v1.GroupUserInfo"></a>

### GroupUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Annotation | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |






<a name="application.management.v1.OwnerApplication"></a>

### OwnerApplication



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Owner | [string](#string) |  |  |
| Infos | [ApplicationInfo](#application.management.v1.ApplicationInfo) | repeated |  |






<a name="application.management.v1.OwnerGroup"></a>

### OwnerGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| Owner | [string](#string) |  |  |
| Infos | [GroupInfo](#application.management.v1.GroupInfo) | repeated |  |






<a name="application.management.v1.RemoveGroupUsersRequest"></a>

### RemoveGroupUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| GroupID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.RemoveGroupUsersResponse"></a>

### RemoveGroupUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.RemoveUsersFromApplicationRequest"></a>

### RemoveUsersFromApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.RemoveUsersFromApplicationResponse"></a>

### RemoveUsersFromApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.ResourceInfo"></a>

### ResourceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| ResourceName | [string](#string) |  |  |
| ResourceDescription | [string](#string) |  |  |
| Type | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |
| UpdateAT | [uint32](#uint32) |  |  |






<a name="application.management.v1.RoleInfo"></a>

### RoleInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleName | [string](#string) |  |  |
| Creator | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |
| UpdateAT | [uint32](#uint32) |  |  |
| Annotation | [string](#string) |  |  |






<a name="application.management.v1.RoleUserInfo"></a>

### RoleUserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| RoleID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| CreateAT | [uint32](#uint32) |  |  |






<a name="application.management.v1.SetGALoginRequest"></a>

### SetGALoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Set | [bool](#bool) |  |  |






<a name="application.management.v1.SetGALoginResponse"></a>

### SetGALoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.SetSMSLoginRequest"></a>

### SetSMSLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppID | [string](#string) |  |  |
| UserID | [string](#string) |  |  |
| Set | [bool](#bool) |  |  |






<a name="application.management.v1.SetSMSLoginResponse"></a>

### SetSMSLoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.SetUserRoleRequest"></a>

### SetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| RoleID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.SetUserRoleResponse"></a>

### SetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleUserInfo](#application.management.v1.RoleUserInfo) | repeated |  |






<a name="application.management.v1.UnSetUserRoleRequest"></a>

### UnSetUserRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserIDs | [string](#string) | repeated |  |
| RoleID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |






<a name="application.management.v1.UnSetUserRoleResponse"></a>

### UnSetUserRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.UpdateApplicationRequest"></a>

### UpdateApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.UpdateApplicationResponse"></a>

### UpdateApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ApplicationInfo](#application.management.v1.ApplicationInfo) |  |  |






<a name="application.management.v1.UpdateGroupRequest"></a>

### UpdateGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.UpdateGroupResponse"></a>

### UpdateGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [GroupInfo](#application.management.v1.GroupInfo) |  |  |






<a name="application.management.v1.UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.UpdateResourceResponse"></a>

### UpdateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [ResourceInfo](#application.management.v1.ResourceInfo) |  |  |






<a name="application.management.v1.UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.UpdateRoleResponse"></a>

### UpdateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [RoleInfo](#application.management.v1.RoleInfo) |  |  |






<a name="application.management.v1.UpdateUserGAStatusRequest"></a>

### UpdateUserGAStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Status | [bool](#bool) |  |  |






<a name="application.management.v1.UpdateUserGAStatusResponse"></a>

### UpdateUserGAStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.UpdateUserKYCStatusRequest"></a>

### UpdateUserKYCStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |
| Status | [bool](#bool) |  |  |






<a name="application.management.v1.UpdateUserKYCStatusResponse"></a>

### UpdateUserKYCStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |






<a name="application.management.v1.UserRole"></a>

### UserRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Infos | [RoleInfo](#application.management.v1.RoleInfo) | repeated |  |
| UserID | [string](#string) |  |  |
| AppID | [string](#string) |  |  |





 

 

 


<a name="application.management.v1.ApplicationManagement"></a>

### ApplicationManagement


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [.npool.v1.VersionResponse](#npool.v1.VersionResponse) | Method Version |
| CreateApplication | [CreateApplicationRequest](#application.management.v1.CreateApplicationRequest) | [CreateApplicationResponse](#application.management.v1.CreateApplicationResponse) | Create an application. |
| UpdateApplication | [UpdateApplicationRequest](#application.management.v1.UpdateApplicationRequest) | [UpdateApplicationResponse](#application.management.v1.UpdateApplicationResponse) | Update an application&#39;s basic info. |
| GetApplication | [GetApplicationRequest](#application.management.v1.GetApplicationRequest) | [GetApplicationResponse](#application.management.v1.GetApplicationResponse) | Get application. |
| GetApplicationByOwner | [GetApplicationByOwnerRequest](#application.management.v1.GetApplicationByOwnerRequest) | [GetApplicationByOwnerResponse](#application.management.v1.GetApplicationByOwnerResponse) | Get application by owner. |
| GetApplications | [GetApplicationsRequest](#application.management.v1.GetApplicationsRequest) | [GetApplicationsResponse](#application.management.v1.GetApplicationsResponse) | Get all applications. |
| DeleteApplication | [DeleteApplicationRequest](#application.management.v1.DeleteApplicationRequest) | [DeleteApplicationResponse](#application.management.v1.DeleteApplicationResponse) | Delete an application. |
| CreateRole | [CreateRoleRequest](#application.management.v1.CreateRoleRequest) | [CreateRoleResponse](#application.management.v1.CreateRoleResponse) | Create a role in app. |
| UpdateRole | [UpdateRoleRequest](#application.management.v1.UpdateRoleRequest) | [UpdateRoleResponse](#application.management.v1.UpdateRoleResponse) | Update role&#39;s basic info. |
| GetRole | [GetRoleRequest](#application.management.v1.GetRoleRequest) | [GetRoleResponse](#application.management.v1.GetRoleResponse) | Get Role. |
| GetRoleByCreator | [GetRoleByCreatorRequest](#application.management.v1.GetRoleByCreatorRequest) | [GetRoleByCreatorResponse](#application.management.v1.GetRoleByCreatorResponse) | Get role by creator. |
| GetRoles | [GetRolesRequest](#application.management.v1.GetRolesRequest) | [GetRolesResponse](#application.management.v1.GetRolesResponse) | Get Roles. |
| DeleteRole | [DeleteRoleRequest](#application.management.v1.DeleteRoleRequest) | [DeleteRoleResponse](#application.management.v1.DeleteRoleResponse) | Delete role from app. |
| SetUserRole | [SetUserRoleRequest](#application.management.v1.SetUserRoleRequest) | [SetUserRoleResponse](#application.management.v1.SetUserRoleResponse) | Set role to user. |
| GetUserRole | [GetUserRoleRequest](#application.management.v1.GetUserRoleRequest) | [GetUserRoleResponse](#application.management.v1.GetUserRoleResponse) | Get user role. |
| GetRoleUsers | [GetRoleUsersRequest](#application.management.v1.GetRoleUsersRequest) | [GetRoleUsersResponse](#application.management.v1.GetRoleUsersResponse) | Get role users. |
| UnSetUserRole | [UnSetUserRoleRequest](#application.management.v1.UnSetUserRoleRequest) | [UnSetUserRoleResponse](#application.management.v1.UnSetUserRoleResponse) | Unset user role. |
| AddUsersToApplication | [AddUsersToApplicationRequest](#application.management.v1.AddUsersToApplicationRequest) | [AddUsersToApplicationResponse](#application.management.v1.AddUsersToApplicationResponse) | Add users to app. |
| GetUserFromApplication | [GetUserFromApplicationRequest](#application.management.v1.GetUserFromApplicationRequest) | [GetUserFromApplicationResponse](#application.management.v1.GetUserFromApplicationResponse) | Get user from app. |
| GetUsersFromApplication | [GetUsersFromApplicationRequest](#application.management.v1.GetUsersFromApplicationRequest) | [GetUsersFromApplicationResponse](#application.management.v1.GetUsersFromApplicationResponse) | Get users from app. |
| GetUserAppID | [GetUserAppIDRequest](#application.management.v1.GetUserAppIDRequest) | [GetUserAppIDResponse](#application.management.v1.GetUserAppIDResponse) | Get user&#39;s appid. |
| RemoveUsersFromApplication | [RemoveUsersFromApplicationRequest](#application.management.v1.RemoveUsersFromApplicationRequest) | [RemoveUsersFromApplicationResponse](#application.management.v1.RemoveUsersFromApplicationResponse) | Remove users from app. |
| CreateGroup | [CreateGroupRequest](#application.management.v1.CreateGroupRequest) | [CreateGroupResponse](#application.management.v1.CreateGroupResponse) | Create group in an application. |
| GetGroup | [GetGroupRequest](#application.management.v1.GetGroupRequest) | [GetGroupResponse](#application.management.v1.GetGroupResponse) | Get group info. |
| GetGroupByOwner | [GetGroupByOwnerRequest](#application.management.v1.GetGroupByOwnerRequest) | [GetGroupByOwnerResponse](#application.management.v1.GetGroupByOwnerResponse) | Get group by owner. |
| GetAllGroups | [GetAllGroupsRequest](#application.management.v1.GetAllGroupsRequest) | [GetAllGroupsResponse](#application.management.v1.GetAllGroupsResponse) | Get all groups. |
| UpdateGroup | [UpdateGroupRequest](#application.management.v1.UpdateGroupRequest) | [UpdateGroupResponse](#application.management.v1.UpdateGroupResponse) | Update group info. |
| DeleteGroup | [DeleteGroupRequest](#application.management.v1.DeleteGroupRequest) | [DeleteGroupResponse](#application.management.v1.DeleteGroupResponse) | Delete group. |
| AddGroupUsers | [AddGroupUsersRequest](#application.management.v1.AddGroupUsersRequest) | [AddGroupUsersResponse](#application.management.v1.AddGroupUsersResponse) | Add users into group. |
| GetGroupUsers | [GetGroupUsersRequest](#application.management.v1.GetGroupUsersRequest) | [GetGroupUsersResponse](#application.management.v1.GetGroupUsersResponse) | Get group users. |
| GetUserGroup | [GetUserGroupRequest](#application.management.v1.GetUserGroupRequest) | [GetUserGroupResponse](#application.management.v1.GetUserGroupResponse) | get user group info |
| RemoveGroupUsers | [RemoveGroupUsersRequest](#application.management.v1.RemoveGroupUsersRequest) | [RemoveGroupUsersResponse](#application.management.v1.RemoveGroupUsersResponse) | Remove users from group. |
| CreateResource | [CreateResourceRequest](#application.management.v1.CreateResourceRequest) | [CreateResourceResponse](#application.management.v1.CreateResourceResponse) | Create resource for app. |
| UpdateResource | [UpdateResourceRequest](#application.management.v1.UpdateResourceRequest) | [UpdateResourceResponse](#application.management.v1.UpdateResourceResponse) | Update resource of app. |
| GetResource | [GetResourceRequest](#application.management.v1.GetResourceRequest) | [GetResourceResponse](#application.management.v1.GetResourceResponse) | Get resource by ID. |
| GetResourceByName | [GetResourceByNameRequest](#application.management.v1.GetResourceByNameRequest) | [GetResourceByNameResponse](#application.management.v1.GetResourceByNameResponse) | get resource by name. |
| GetResourceByCreator | [GetResourceByCreatorRequest](#application.management.v1.GetResourceByCreatorRequest) | [GetResourceByCreatorResponse](#application.management.v1.GetResourceByCreatorResponse) | Get resource by creator. |
| GetResources | [GetResourcesRequest](#application.management.v1.GetResourcesRequest) | [GetResourcesResponse](#application.management.v1.GetResourcesResponse) | Get all resources from app. |
| DeleteResource | [DeleteResourceRequest](#application.management.v1.DeleteResourceRequest) | [DeleteResourceResponse](#application.management.v1.DeleteResourceResponse) | Delete resource from app. |
| SetGALogin | [SetGALoginRequest](#application.management.v1.SetGALoginRequest) | [SetGALoginResponse](#application.management.v1.SetGALoginResponse) | set ga login verify option |
| SetSMSLogin | [SetSMSLoginRequest](#application.management.v1.SetSMSLoginRequest) | [SetSMSLoginResponse](#application.management.v1.SetSMSLoginResponse) | set sms login verify option |
| AddUserLoginTime | [AddUserLoginTimeRequest](#application.management.v1.AddUserLoginTimeRequest) | [AddUserLoginTimeResponse](#application.management.v1.AddUserLoginTimeResponse) | add user login times |
| UpdateUserGAStatus | [UpdateUserGAStatusRequest](#application.management.v1.UpdateUserGAStatusRequest) | [UpdateUserGAStatusResponse](#application.management.v1.UpdateUserGAStatusResponse) | update user google authentication status. |
| UpdateUserKYCStatus | [UpdateUserKYCStatusRequest](#application.management.v1.UpdateUserKYCStatusRequest) | [UpdateUserKYCStatusResponse](#application.management.v1.UpdateUserKYCStatusResponse) | update user kyc status. |
| GetApplicationUserDetail | [GetApplicationUserDetailRequest](#application.management.v1.GetApplicationUserDetailRequest) | [GetApplicationUserDetailResponse](#application.management.v1.GetApplicationUserDetailResponse) | get application user detail info. |

 



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

