# UserModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to [**AvatarModelV1**](AvatarModelV1.md) |  | [optional] 
**TimeJoined** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**LastActiveTime** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**BirthdayModelV1**](BirthdayModelV1.md) |  | [optional] 
**Age** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **int64** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**ForumTitle** | Pointer to **string** |  | [optional] 
**FoldingAtHome** | Pointer to **bool** |  | [optional] 
**Profile** | Pointer to [**UserModelV1Profile**](UserModelV1Profile.md) |  | [optional] 
**Stats** | Pointer to [**UserModelV1Stats**](UserModelV1Stats.md) |  | [optional] 
**Admin** | Pointer to [**UserModelV1Admin**](UserModelV1Admin.md) |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 
**UserGroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewUserModelV1

`func NewUserModelV1() *UserModelV1`

NewUserModelV1 instantiates a new UserModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelV1WithDefaults

`func NewUserModelV1WithDefaults() *UserModelV1`

NewUserModelV1WithDefaults instantiates a new UserModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserModelV1) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserModelV1) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserModelV1) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserModelV1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *UserModelV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserModelV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserModelV1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserModelV1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUrl

`func (o *UserModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEmail

`func (o *UserModelV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModelV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModelV1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserModelV1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *UserModelV1) GetAvatar() AvatarModelV1`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserModelV1) GetAvatarOk() (*AvatarModelV1, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserModelV1) SetAvatar(v AvatarModelV1)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserModelV1) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetTimeJoined

`func (o *UserModelV1) GetTimeJoined() TimeV1`

GetTimeJoined returns the TimeJoined field if non-nil, zero value otherwise.

### GetTimeJoinedOk

`func (o *UserModelV1) GetTimeJoinedOk() (*TimeV1, bool)`

GetTimeJoinedOk returns a tuple with the TimeJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeJoined

`func (o *UserModelV1) SetTimeJoined(v TimeV1)`

SetTimeJoined sets TimeJoined field to given value.

### HasTimeJoined

`func (o *UserModelV1) HasTimeJoined() bool`

HasTimeJoined returns a boolean if a field has been set.

### GetLastActiveTime

`func (o *UserModelV1) GetLastActiveTime() TimeV1`

GetLastActiveTime returns the LastActiveTime field if non-nil, zero value otherwise.

### GetLastActiveTimeOk

`func (o *UserModelV1) GetLastActiveTimeOk() (*TimeV1, bool)`

GetLastActiveTimeOk returns a tuple with the LastActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveTime

`func (o *UserModelV1) SetLastActiveTime(v TimeV1)`

SetLastActiveTime sets LastActiveTime field to given value.

### HasLastActiveTime

`func (o *UserModelV1) HasLastActiveTime() bool`

HasLastActiveTime returns a boolean if a field has been set.

### GetGender

`func (o *UserModelV1) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserModelV1) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserModelV1) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserModelV1) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBirthday

`func (o *UserModelV1) GetBirthday() BirthdayModelV1`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UserModelV1) GetBirthdayOk() (*BirthdayModelV1, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UserModelV1) SetBirthday(v BirthdayModelV1)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UserModelV1) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetAge

`func (o *UserModelV1) GetAge() int64`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *UserModelV1) GetAgeOk() (*int64, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *UserModelV1) SetAge(v int64)`

SetAge sets Age field to given value.

### HasAge

`func (o *UserModelV1) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetTimezone

`func (o *UserModelV1) GetTimezone() int64`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserModelV1) GetTimezoneOk() (*int64, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserModelV1) SetTimezone(v int64)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserModelV1) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSignature

`func (o *UserModelV1) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserModelV1) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserModelV1) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UserModelV1) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetLocation

`func (o *UserModelV1) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserModelV1) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserModelV1) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UserModelV1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetForumTitle

`func (o *UserModelV1) GetForumTitle() string`

GetForumTitle returns the ForumTitle field if non-nil, zero value otherwise.

### GetForumTitleOk

`func (o *UserModelV1) GetForumTitleOk() (*string, bool)`

GetForumTitleOk returns a tuple with the ForumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTitle

`func (o *UserModelV1) SetForumTitle(v string)`

SetForumTitle sets ForumTitle field to given value.

### HasForumTitle

`func (o *UserModelV1) HasForumTitle() bool`

HasForumTitle returns a boolean if a field has been set.

### GetFoldingAtHome

`func (o *UserModelV1) GetFoldingAtHome() bool`

GetFoldingAtHome returns the FoldingAtHome field if non-nil, zero value otherwise.

### GetFoldingAtHomeOk

`func (o *UserModelV1) GetFoldingAtHomeOk() (*bool, bool)`

GetFoldingAtHomeOk returns a tuple with the FoldingAtHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldingAtHome

`func (o *UserModelV1) SetFoldingAtHome(v bool)`

SetFoldingAtHome sets FoldingAtHome field to given value.

### HasFoldingAtHome

`func (o *UserModelV1) HasFoldingAtHome() bool`

HasFoldingAtHome returns a boolean if a field has been set.

### GetProfile

`func (o *UserModelV1) GetProfile() UserModelV1Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserModelV1) GetProfileOk() (*UserModelV1Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserModelV1) SetProfile(v UserModelV1Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserModelV1) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStats

`func (o *UserModelV1) GetStats() UserModelV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UserModelV1) GetStatsOk() (*UserModelV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UserModelV1) SetStats(v UserModelV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UserModelV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetAdmin

`func (o *UserModelV1) GetAdmin() UserModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserModelV1) GetAdminOk() (*UserModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserModelV1) SetAdmin(v UserModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetUserGroup

`func (o *UserModelV1) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UserModelV1) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UserModelV1) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UserModelV1) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetUserGroupName

`func (o *UserModelV1) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *UserModelV1) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *UserModelV1) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *UserModelV1) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


