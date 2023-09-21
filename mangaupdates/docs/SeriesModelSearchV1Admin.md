# SeriesModelSearchV1Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeriesModelSearchV1Admin

`func NewSeriesModelSearchV1Admin() *SeriesModelSearchV1Admin`

NewSeriesModelSearchV1Admin instantiates a new SeriesModelSearchV1Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelSearchV1AdminWithDefaults

`func NewSeriesModelSearchV1AdminWithDefaults() *SeriesModelSearchV1Admin`

NewSeriesModelSearchV1AdminWithDefaults instantiates a new SeriesModelSearchV1Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedBy

`func (o *SeriesModelSearchV1Admin) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *SeriesModelSearchV1Admin) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *SeriesModelSearchV1Admin) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *SeriesModelSearchV1Admin) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetApproved

`func (o *SeriesModelSearchV1Admin) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SeriesModelSearchV1Admin) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SeriesModelSearchV1Admin) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *SeriesModelSearchV1Admin) HasApproved() bool`

HasApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


