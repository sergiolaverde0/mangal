# UserModelV1Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdmin** | Pointer to **bool** |  | [optional] 
**RegistrationIp** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**UserModelV1AdminPermissions**](UserModelV1AdminPermissions.md) |  | [optional] 
**LastSeriesUpdate** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**EmailApproved** | Pointer to **bool** |  | [optional] 
**ForumAdmin** | Pointer to **bool** |  | [optional] 
**RegistrationReason** | Pointer to **string** |  | [optional] 
**Upgrade** | Pointer to [**UserModelV1AdminUpgrade**](UserModelV1AdminUpgrade.md) |  | [optional] 
**Banned** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserModelV1Admin

`func NewUserModelV1Admin() *UserModelV1Admin`

NewUserModelV1Admin instantiates a new UserModelV1Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelV1AdminWithDefaults

`func NewUserModelV1AdminWithDefaults() *UserModelV1Admin`

NewUserModelV1AdminWithDefaults instantiates a new UserModelV1Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdmin

`func (o *UserModelV1Admin) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserModelV1Admin) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UserModelV1Admin) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UserModelV1Admin) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetRegistrationIp

`func (o *UserModelV1Admin) GetRegistrationIp() string`

GetRegistrationIp returns the RegistrationIp field if non-nil, zero value otherwise.

### GetRegistrationIpOk

`func (o *UserModelV1Admin) GetRegistrationIpOk() (*string, bool)`

GetRegistrationIpOk returns a tuple with the RegistrationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationIp

`func (o *UserModelV1Admin) SetRegistrationIp(v string)`

SetRegistrationIp sets RegistrationIp field to given value.

### HasRegistrationIp

`func (o *UserModelV1Admin) HasRegistrationIp() bool`

HasRegistrationIp returns a boolean if a field has been set.

### GetPermissions

`func (o *UserModelV1Admin) GetPermissions() UserModelV1AdminPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserModelV1Admin) GetPermissionsOk() (*UserModelV1AdminPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserModelV1Admin) SetPermissions(v UserModelV1AdminPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserModelV1Admin) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLastSeriesUpdate

`func (o *UserModelV1Admin) GetLastSeriesUpdate() TimeV1`

GetLastSeriesUpdate returns the LastSeriesUpdate field if non-nil, zero value otherwise.

### GetLastSeriesUpdateOk

`func (o *UserModelV1Admin) GetLastSeriesUpdateOk() (*TimeV1, bool)`

GetLastSeriesUpdateOk returns a tuple with the LastSeriesUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeriesUpdate

`func (o *UserModelV1Admin) SetLastSeriesUpdate(v TimeV1)`

SetLastSeriesUpdate sets LastSeriesUpdate field to given value.

### HasLastSeriesUpdate

`func (o *UserModelV1Admin) HasLastSeriesUpdate() bool`

HasLastSeriesUpdate returns a boolean if a field has been set.

### GetApproved

`func (o *UserModelV1Admin) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *UserModelV1Admin) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *UserModelV1Admin) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *UserModelV1Admin) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetEmailApproved

`func (o *UserModelV1Admin) GetEmailApproved() bool`

GetEmailApproved returns the EmailApproved field if non-nil, zero value otherwise.

### GetEmailApprovedOk

`func (o *UserModelV1Admin) GetEmailApprovedOk() (*bool, bool)`

GetEmailApprovedOk returns a tuple with the EmailApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailApproved

`func (o *UserModelV1Admin) SetEmailApproved(v bool)`

SetEmailApproved sets EmailApproved field to given value.

### HasEmailApproved

`func (o *UserModelV1Admin) HasEmailApproved() bool`

HasEmailApproved returns a boolean if a field has been set.

### GetForumAdmin

`func (o *UserModelV1Admin) GetForumAdmin() bool`

GetForumAdmin returns the ForumAdmin field if non-nil, zero value otherwise.

### GetForumAdminOk

`func (o *UserModelV1Admin) GetForumAdminOk() (*bool, bool)`

GetForumAdminOk returns a tuple with the ForumAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumAdmin

`func (o *UserModelV1Admin) SetForumAdmin(v bool)`

SetForumAdmin sets ForumAdmin field to given value.

### HasForumAdmin

`func (o *UserModelV1Admin) HasForumAdmin() bool`

HasForumAdmin returns a boolean if a field has been set.

### GetRegistrationReason

`func (o *UserModelV1Admin) GetRegistrationReason() string`

GetRegistrationReason returns the RegistrationReason field if non-nil, zero value otherwise.

### GetRegistrationReasonOk

`func (o *UserModelV1Admin) GetRegistrationReasonOk() (*string, bool)`

GetRegistrationReasonOk returns a tuple with the RegistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationReason

`func (o *UserModelV1Admin) SetRegistrationReason(v string)`

SetRegistrationReason sets RegistrationReason field to given value.

### HasRegistrationReason

`func (o *UserModelV1Admin) HasRegistrationReason() bool`

HasRegistrationReason returns a boolean if a field has been set.

### GetUpgrade

`func (o *UserModelV1Admin) GetUpgrade() UserModelV1AdminUpgrade`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *UserModelV1Admin) GetUpgradeOk() (*UserModelV1AdminUpgrade, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *UserModelV1Admin) SetUpgrade(v UserModelV1AdminUpgrade)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *UserModelV1Admin) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetBanned

`func (o *UserModelV1Admin) GetBanned() bool`

GetBanned returns the Banned field if non-nil, zero value otherwise.

### GetBannedOk

`func (o *UserModelV1Admin) GetBannedOk() (*bool, bool)`

GetBannedOk returns a tuple with the Banned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanned

`func (o *UserModelV1Admin) SetBanned(v bool)`

SetBanned sets Banned field to given value.

### HasBanned

`func (o *UserModelV1Admin) HasBanned() bool`

HasBanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


