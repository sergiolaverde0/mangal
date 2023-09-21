# ListsSeriesModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**ListsSeriesModelV1Series**](ListsSeriesModelV1Series.md) |  | 
**ListId** | Pointer to **int64** |  | [optional] 
**ListType** | Pointer to **string** |  | [optional] 
**ListIcon** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ListsSeriesModelV1Status**](ListsSeriesModelV1Status.md) |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewListsSeriesModelV1

`func NewListsSeriesModelV1(series ListsSeriesModelV1Series, ) *ListsSeriesModelV1`

NewListsSeriesModelV1 instantiates a new ListsSeriesModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsSeriesModelV1WithDefaults

`func NewListsSeriesModelV1WithDefaults() *ListsSeriesModelV1`

NewListsSeriesModelV1WithDefaults instantiates a new ListsSeriesModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *ListsSeriesModelV1) GetSeries() ListsSeriesModelV1Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ListsSeriesModelV1) GetSeriesOk() (*ListsSeriesModelV1Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ListsSeriesModelV1) SetSeries(v ListsSeriesModelV1Series)`

SetSeries sets Series field to given value.


### GetListId

`func (o *ListsSeriesModelV1) GetListId() int64`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ListsSeriesModelV1) GetListIdOk() (*int64, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ListsSeriesModelV1) SetListId(v int64)`

SetListId sets ListId field to given value.

### HasListId

`func (o *ListsSeriesModelV1) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetListType

`func (o *ListsSeriesModelV1) GetListType() string`

GetListType returns the ListType field if non-nil, zero value otherwise.

### GetListTypeOk

`func (o *ListsSeriesModelV1) GetListTypeOk() (*string, bool)`

GetListTypeOk returns a tuple with the ListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListType

`func (o *ListsSeriesModelV1) SetListType(v string)`

SetListType sets ListType field to given value.

### HasListType

`func (o *ListsSeriesModelV1) HasListType() bool`

HasListType returns a boolean if a field has been set.

### GetListIcon

`func (o *ListsSeriesModelV1) GetListIcon() string`

GetListIcon returns the ListIcon field if non-nil, zero value otherwise.

### GetListIconOk

`func (o *ListsSeriesModelV1) GetListIconOk() (*string, bool)`

GetListIconOk returns a tuple with the ListIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListIcon

`func (o *ListsSeriesModelV1) SetListIcon(v string)`

SetListIcon sets ListIcon field to given value.

### HasListIcon

`func (o *ListsSeriesModelV1) HasListIcon() bool`

HasListIcon returns a boolean if a field has been set.

### GetStatus

`func (o *ListsSeriesModelV1) GetStatus() ListsSeriesModelV1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListsSeriesModelV1) GetStatusOk() (*ListsSeriesModelV1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListsSeriesModelV1) SetStatus(v ListsSeriesModelV1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListsSeriesModelV1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *ListsSeriesModelV1) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListsSeriesModelV1) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListsSeriesModelV1) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListsSeriesModelV1) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ListsSeriesModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ListsSeriesModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ListsSeriesModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ListsSeriesModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


