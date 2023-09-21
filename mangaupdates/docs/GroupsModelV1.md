# GroupsModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]GroupsModelV1Associated**](GroupsModelV1Associated.md) |  | [optional] 
**Social** | Pointer to [**GroupsModelV1Social**](GroupsModelV1Social.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**Admin** | Pointer to [**GroupsModelV1Admin**](GroupsModelV1Admin.md) |  | [optional] 

## Methods

### NewGroupsModelV1

`func NewGroupsModelV1() *GroupsModelV1`

NewGroupsModelV1 instantiates a new GroupsModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsModelV1WithDefaults

`func NewGroupsModelV1WithDefaults() *GroupsModelV1`

NewGroupsModelV1WithDefaults instantiates a new GroupsModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupsModelV1) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupsModelV1) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupsModelV1) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupsModelV1) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GroupsModelV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsModelV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsModelV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupsModelV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *GroupsModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GroupsModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GroupsModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GroupsModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAssociated

`func (o *GroupsModelV1) GetAssociated() []GroupsModelV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *GroupsModelV1) GetAssociatedOk() (*[]GroupsModelV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *GroupsModelV1) SetAssociated(v []GroupsModelV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *GroupsModelV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetSocial

`func (o *GroupsModelV1) GetSocial() GroupsModelV1Social`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *GroupsModelV1) GetSocialOk() (*GroupsModelV1Social, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *GroupsModelV1) SetSocial(v GroupsModelV1Social)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *GroupsModelV1) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetActive

`func (o *GroupsModelV1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GroupsModelV1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GroupsModelV1) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GroupsModelV1) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *GroupsModelV1) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GroupsModelV1) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GroupsModelV1) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GroupsModelV1) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAddedBy

`func (o *GroupsModelV1) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *GroupsModelV1) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *GroupsModelV1) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *GroupsModelV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetAdmin

`func (o *GroupsModelV1) GetAdmin() GroupsModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *GroupsModelV1) GetAdminOk() (*GroupsModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *GroupsModelV1) SetAdmin(v GroupsModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *GroupsModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


