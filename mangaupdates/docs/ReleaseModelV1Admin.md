# ReleaseModelV1Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approved** | Pointer to **bool** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 

## Methods

### NewReleaseModelV1Admin

`func NewReleaseModelV1Admin() *ReleaseModelV1Admin`

NewReleaseModelV1Admin instantiates a new ReleaseModelV1Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseModelV1AdminWithDefaults

`func NewReleaseModelV1AdminWithDefaults() *ReleaseModelV1Admin`

NewReleaseModelV1AdminWithDefaults instantiates a new ReleaseModelV1Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproved

`func (o *ReleaseModelV1Admin) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ReleaseModelV1Admin) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ReleaseModelV1Admin) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ReleaseModelV1Admin) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetArchived

`func (o *ReleaseModelV1Admin) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ReleaseModelV1Admin) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ReleaseModelV1Admin) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ReleaseModelV1Admin) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetAddedBy

`func (o *ReleaseModelV1Admin) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *ReleaseModelV1Admin) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *ReleaseModelV1Admin) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *ReleaseModelV1Admin) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


