# SeriesGroupListResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupList** | Pointer to [**[]GroupsModelSearchV1**](GroupsModelSearchV1.md) |  | [optional] 
**ReleaseList** | Pointer to [**[]ReleaseModelSearchV1**](ReleaseModelSearchV1.md) |  | [optional] 

## Methods

### NewSeriesGroupListResponseV1

`func NewSeriesGroupListResponseV1() *SeriesGroupListResponseV1`

NewSeriesGroupListResponseV1 instantiates a new SeriesGroupListResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesGroupListResponseV1WithDefaults

`func NewSeriesGroupListResponseV1WithDefaults() *SeriesGroupListResponseV1`

NewSeriesGroupListResponseV1WithDefaults instantiates a new SeriesGroupListResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupList

`func (o *SeriesGroupListResponseV1) GetGroupList() []GroupsModelSearchV1`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *SeriesGroupListResponseV1) GetGroupListOk() (*[]GroupsModelSearchV1, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *SeriesGroupListResponseV1) SetGroupList(v []GroupsModelSearchV1)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *SeriesGroupListResponseV1) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### GetReleaseList

`func (o *SeriesGroupListResponseV1) GetReleaseList() []ReleaseModelSearchV1`

GetReleaseList returns the ReleaseList field if non-nil, zero value otherwise.

### GetReleaseListOk

`func (o *SeriesGroupListResponseV1) GetReleaseListOk() (*[]ReleaseModelSearchV1, bool)`

GetReleaseListOk returns a tuple with the ReleaseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseList

`func (o *SeriesGroupListResponseV1) SetReleaseList(v []ReleaseModelSearchV1)`

SetReleaseList sets ReleaseList field to given value.

### HasReleaseList

`func (o *SeriesGroupListResponseV1) HasReleaseList() bool`

HasReleaseList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


