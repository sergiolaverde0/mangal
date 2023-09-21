# UserModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**NewAvatarId** | Pointer to **int64** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**BirthdayModelV1**](BirthdayModelV1.md) |  | [optional] 
**Timezone** | Pointer to **int64** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**ForumTitle** | Pointer to **string** |  | [optional] 
**FoldingAtHome** | Pointer to **bool** |  | [optional] 
**Profile** | Pointer to [**UserModelUpdateV1Profile**](UserModelUpdateV1Profile.md) |  | [optional] 
**Admin** | Pointer to [**UserModelUpdateV1Admin**](UserModelUpdateV1Admin.md) |  | [optional] 
**UserGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewUserModelUpdateV1

`func NewUserModelUpdateV1() *UserModelUpdateV1`

NewUserModelUpdateV1 instantiates a new UserModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelUpdateV1WithDefaults

`func NewUserModelUpdateV1WithDefaults() *UserModelUpdateV1`

NewUserModelUpdateV1WithDefaults instantiates a new UserModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserModelUpdateV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserModelUpdateV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserModelUpdateV1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserModelUpdateV1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserModelUpdateV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserModelUpdateV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserModelUpdateV1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserModelUpdateV1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UserModelUpdateV1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserModelUpdateV1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserModelUpdateV1) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserModelUpdateV1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNewAvatarId

`func (o *UserModelUpdateV1) GetNewAvatarId() int64`

GetNewAvatarId returns the NewAvatarId field if non-nil, zero value otherwise.

### GetNewAvatarIdOk

`func (o *UserModelUpdateV1) GetNewAvatarIdOk() (*int64, bool)`

GetNewAvatarIdOk returns a tuple with the NewAvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAvatarId

`func (o *UserModelUpdateV1) SetNewAvatarId(v int64)`

SetNewAvatarId sets NewAvatarId field to given value.

### HasNewAvatarId

`func (o *UserModelUpdateV1) HasNewAvatarId() bool`

HasNewAvatarId returns a boolean if a field has been set.

### GetGender

`func (o *UserModelUpdateV1) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserModelUpdateV1) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserModelUpdateV1) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserModelUpdateV1) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBirthday

`func (o *UserModelUpdateV1) GetBirthday() BirthdayModelV1`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UserModelUpdateV1) GetBirthdayOk() (*BirthdayModelV1, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UserModelUpdateV1) SetBirthday(v BirthdayModelV1)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UserModelUpdateV1) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetTimezone

`func (o *UserModelUpdateV1) GetTimezone() int64`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserModelUpdateV1) GetTimezoneOk() (*int64, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserModelUpdateV1) SetTimezone(v int64)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserModelUpdateV1) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSignature

`func (o *UserModelUpdateV1) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserModelUpdateV1) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserModelUpdateV1) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UserModelUpdateV1) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetLocation

`func (o *UserModelUpdateV1) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserModelUpdateV1) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserModelUpdateV1) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UserModelUpdateV1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetForumTitle

`func (o *UserModelUpdateV1) GetForumTitle() string`

GetForumTitle returns the ForumTitle field if non-nil, zero value otherwise.

### GetForumTitleOk

`func (o *UserModelUpdateV1) GetForumTitleOk() (*string, bool)`

GetForumTitleOk returns a tuple with the ForumTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTitle

`func (o *UserModelUpdateV1) SetForumTitle(v string)`

SetForumTitle sets ForumTitle field to given value.

### HasForumTitle

`func (o *UserModelUpdateV1) HasForumTitle() bool`

HasForumTitle returns a boolean if a field has been set.

### GetFoldingAtHome

`func (o *UserModelUpdateV1) GetFoldingAtHome() bool`

GetFoldingAtHome returns the FoldingAtHome field if non-nil, zero value otherwise.

### GetFoldingAtHomeOk

`func (o *UserModelUpdateV1) GetFoldingAtHomeOk() (*bool, bool)`

GetFoldingAtHomeOk returns a tuple with the FoldingAtHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldingAtHome

`func (o *UserModelUpdateV1) SetFoldingAtHome(v bool)`

SetFoldingAtHome sets FoldingAtHome field to given value.

### HasFoldingAtHome

`func (o *UserModelUpdateV1) HasFoldingAtHome() bool`

HasFoldingAtHome returns a boolean if a field has been set.

### GetProfile

`func (o *UserModelUpdateV1) GetProfile() UserModelUpdateV1Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserModelUpdateV1) GetProfileOk() (*UserModelUpdateV1Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserModelUpdateV1) SetProfile(v UserModelUpdateV1Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserModelUpdateV1) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetAdmin

`func (o *UserModelUpdateV1) GetAdmin() UserModelUpdateV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserModelUpdateV1) GetAdminOk() (*UserModelUpdateV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserModelUpdateV1) SetAdmin(v UserModelUpdateV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserModelUpdateV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetUserGroup

`func (o *UserModelUpdateV1) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UserModelUpdateV1) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UserModelUpdateV1) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UserModelUpdateV1) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


