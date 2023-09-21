# ListsSeriesModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**ListsSeriesModelUpdateV1Series**](ListsSeriesModelUpdateV1Series.md) |  | 
**ListId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to [**ListsSeriesModelUpdateV1Status**](ListsSeriesModelUpdateV1Status.md) |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 

## Methods

### NewListsSeriesModelUpdateV1

`func NewListsSeriesModelUpdateV1(series ListsSeriesModelUpdateV1Series, ) *ListsSeriesModelUpdateV1`

NewListsSeriesModelUpdateV1 instantiates a new ListsSeriesModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsSeriesModelUpdateV1WithDefaults

`func NewListsSeriesModelUpdateV1WithDefaults() *ListsSeriesModelUpdateV1`

NewListsSeriesModelUpdateV1WithDefaults instantiates a new ListsSeriesModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *ListsSeriesModelUpdateV1) GetSeries() ListsSeriesModelUpdateV1Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ListsSeriesModelUpdateV1) GetSeriesOk() (*ListsSeriesModelUpdateV1Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ListsSeriesModelUpdateV1) SetSeries(v ListsSeriesModelUpdateV1Series)`

SetSeries sets Series field to given value.


### GetListId

`func (o *ListsSeriesModelUpdateV1) GetListId() int64`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ListsSeriesModelUpdateV1) GetListIdOk() (*int64, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ListsSeriesModelUpdateV1) SetListId(v int64)`

SetListId sets ListId field to given value.

### HasListId

`func (o *ListsSeriesModelUpdateV1) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetStatus

`func (o *ListsSeriesModelUpdateV1) GetStatus() ListsSeriesModelUpdateV1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListsSeriesModelUpdateV1) GetStatusOk() (*ListsSeriesModelUpdateV1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListsSeriesModelUpdateV1) SetStatus(v ListsSeriesModelUpdateV1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListsSeriesModelUpdateV1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *ListsSeriesModelUpdateV1) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListsSeriesModelUpdateV1) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListsSeriesModelUpdateV1) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListsSeriesModelUpdateV1) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


