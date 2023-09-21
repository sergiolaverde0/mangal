# UserModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to [**AvatarModelSearchV1**](AvatarModelSearchV1.md) |  | [optional] 
**TimeJoined** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**ForumTitle** | Pointer to **string** |  | [optional] 
**FoldingAtHome** | Pointer to **bool** |  | [optional] 
**Profile** | Pointer to [**UserModelSearchV1Profile**](UserModelSearchV1Profile.md) |  | [optional] 
**Stats** | Pointer to [**UserModelSearchV1Stats**](UserModelSearchV1Stats.md) |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 
**UserGroupName** | Pointer to **string** |  | [optional] 

## Methods

### NewUserModelSearchV1

`func NewUserModelSearchV1() *UserModelSearchV1`

NewUserModelSearchV1 instantiates a new UserModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelSearchV1WithDefaults

`func NewUserModelSearchV1WithDefaults() *UserModelSearchV1`

NewUserModelSearchV1WithDefaults instantiates a new UserModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserModelSearchV1) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserModelSearchV1) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserModelSearchV1) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserModelSearchV1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *UserModelSearchV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserModelSearchV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserModelSearchV1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserModelSearchV1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUrl

`func (o *UserModelSearchV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserModelSearchV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserModelSearchV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserModelSearchV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvatar

`func (o *UserModelSearchV1) GetAvatar() AvatarModelSearchV1`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserModelSearchV1) GetAvatarOk() (*AvatarModelSearchV1, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserModelSearchV1) SetAvatar(v AvatarModelSearchV1)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserModelSearchV1) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetTimeJoined

`func (o *UserModelSearchV1) GetTimeJoined() TimeV1`

GetTimeJoined returns the TimeJoined field if non-nil, zero value otherwise.

### GetTimeJoinedOk

`func (o *UserModelSearchV1) GetTimeJoinedOk() (*TimeV1, bool)`

GetTimeJoinedOk returns a tuple with the TimeJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeJoined

`func (o *UserModelSearchV1) SetTimeJoined(v TimeV1)`

SetTimeJoined sets TimeJoined field to given value.

### HasTimeJoined

`func (o *UserModelSearchV1) HasTimeJoined() bool`

HasTimeJoined returns a boolean if a field has been set.

### GetSignature

`func (o *UserModelSearchV1) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserModelSearchV1) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserModelSearchV1) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UserModelSearchV1) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetForumTitle

`func (o *UserModelSearchV1) GetForumTitle() string`

GetForumTitle returns the ForumTitle field if non-nil, zero value otherwise.

### GetForumTitleOk

`func (o *UserModelSearchV1) GetForumTitleOk() (*string, bool)`

GetForumTitleOk returns a tuple with the ForumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTitle

`func (o *UserModelSearchV1) SetForumTitle(v string)`

SetForumTitle sets ForumTitle field to given value.

### HasForumTitle

`func (o *UserModelSearchV1) HasForumTitle() bool`

HasForumTitle returns a boolean if a field has been set.

### GetFoldingAtHome

`func (o *UserModelSearchV1) GetFoldingAtHome() bool`

GetFoldingAtHome returns the FoldingAtHome field if non-nil, zero value otherwise.

### GetFoldingAtHomeOk

`func (o *UserModelSearchV1) GetFoldingAtHomeOk() (*bool, bool)`

GetFoldingAtHomeOk returns a tuple with the FoldingAtHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldingAtHome

`func (o *UserModelSearchV1) SetFoldingAtHome(v bool)`

SetFoldingAtHome sets FoldingAtHome field to given value.

### HasFoldingAtHome

`func (o *UserModelSearchV1) HasFoldingAtHome() bool`

HasFoldingAtHome returns a boolean if a field has been set.

### GetProfile

`func (o *UserModelSearchV1) GetProfile() UserModelSearchV1Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserModelSearchV1) GetProfileOk() (*UserModelSearchV1Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserModelSearchV1) SetProfile(v UserModelSearchV1Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserModelSearchV1) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStats

`func (o *UserModelSearchV1) GetStats() UserModelSearchV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UserModelSearchV1) GetStatsOk() (*UserModelSearchV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UserModelSearchV1) SetStats(v UserModelSearchV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UserModelSearchV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetUserGroup

`func (o *UserModelSearchV1) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UserModelSearchV1) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UserModelSearchV1) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UserModelSearchV1) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetUserGroupName

`func (o *UserModelSearchV1) GetUserGroupName() string`

GetUserGroupName returns the UserGroupName field if non-nil, zero value otherwise.

### GetUserGroupNameOk

`func (o *UserModelSearchV1) GetUserGroupNameOk() (*string, bool)`

GetUserGroupNameOk returns a tuple with the UserGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupName

`func (o *UserModelSearchV1) SetUserGroupName(v string)`

SetUserGroupName sets UserGroupName field to given value.

### HasUserGroupName

`func (o *UserModelSearchV1) HasUserGroupName() bool`

HasUserGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


