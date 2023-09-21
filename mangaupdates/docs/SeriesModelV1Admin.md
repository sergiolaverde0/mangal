# SeriesModelV1Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeriesModelV1Admin

`func NewSeriesModelV1Admin() *SeriesModelV1Admin`

NewSeriesModelV1Admin instantiates a new SeriesModelV1Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelV1AdminWithDefaults

`func NewSeriesModelV1AdminWithDefaults() *SeriesModelV1Admin`

NewSeriesModelV1AdminWithDefaults instantiates a new SeriesModelV1Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedBy

`func (o *SeriesModelV1Admin) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *SeriesModelV1Admin) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *SeriesModelV1Admin) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *SeriesModelV1Admin) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetApproved

`func (o *SeriesModelV1Admin) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SeriesModelV1Admin) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SeriesModelV1Admin) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *SeriesModelV1Admin) HasApproved() bool`

HasApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


