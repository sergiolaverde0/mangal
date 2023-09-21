# SeriesModelV1RelatedSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelationId** | Pointer to **int64** |  | [optional] 
**RelationType** | **string** |  | 
**RelatedSeriesId** | **int64** |  | 
**RelatedSeriesName** | Pointer to **string** |  | [optional] 
**TriggeredByRelationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewSeriesModelV1RelatedSeries

`func NewSeriesModelV1RelatedSeries(relationType string, relatedSeriesId int64, ) *SeriesModelV1RelatedSeries`

NewSeriesModelV1RelatedSeries instantiates a new SeriesModelV1RelatedSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelV1RelatedSeriesWithDefaults

`func NewSeriesModelV1RelatedSeriesWithDefaults() *SeriesModelV1RelatedSeries`

NewSeriesModelV1RelatedSeriesWithDefaults instantiates a new SeriesModelV1RelatedSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationId

`func (o *SeriesModelV1RelatedSeries) GetRelationId() int64`

GetRelationId returns the RelationId field if non-nil, zero value otherwise.

### GetRelationIdOk

`func (o *SeriesModelV1RelatedSeries) GetRelationIdOk() (*int64, bool)`

GetRelationIdOk returns a tuple with the RelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationId

`func (o *SeriesModelV1RelatedSeries) SetRelationId(v int64)`

SetRelationId sets RelationId field to given value.

### HasRelationId

`func (o *SeriesModelV1RelatedSeries) HasRelationId() bool`

HasRelationId returns a boolean if a field has been set.

### GetRelationType

`func (o *SeriesModelV1RelatedSeries) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *SeriesModelV1RelatedSeries) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *SeriesModelV1RelatedSeries) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.


### GetRelatedSeriesId

`func (o *SeriesModelV1RelatedSeries) GetRelatedSeriesId() int64`

GetRelatedSeriesId returns the RelatedSeriesId field if non-nil, zero value otherwise.

### GetRelatedSeriesIdOk

`func (o *SeriesModelV1RelatedSeries) GetRelatedSeriesIdOk() (*int64, bool)`

GetRelatedSeriesIdOk returns a tuple with the RelatedSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSeriesId

`func (o *SeriesModelV1RelatedSeries) SetRelatedSeriesId(v int64)`

SetRelatedSeriesId sets RelatedSeriesId field to given value.


### GetRelatedSeriesName

`func (o *SeriesModelV1RelatedSeries) GetRelatedSeriesName() string`

GetRelatedSeriesName returns the RelatedSeriesName field if non-nil, zero value otherwise.

### GetRelatedSeriesNameOk

`func (o *SeriesModelV1RelatedSeries) GetRelatedSeriesNameOk() (*string, bool)`

GetRelatedSeriesNameOk returns a tuple with the RelatedSeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSeriesName

`func (o *SeriesModelV1RelatedSeries) SetRelatedSeriesName(v string)`

SetRelatedSeriesName sets RelatedSeriesName field to given value.

### HasRelatedSeriesName

`func (o *SeriesModelV1RelatedSeries) HasRelatedSeriesName() bool`

HasRelatedSeriesName returns a boolean if a field has been set.

### GetTriggeredByRelationId

`func (o *SeriesModelV1RelatedSeries) GetTriggeredByRelationId() int64`

GetTriggeredByRelationId returns the TriggeredByRelationId field if non-nil, zero value otherwise.

### GetTriggeredByRelationIdOk

`func (o *SeriesModelV1RelatedSeries) GetTriggeredByRelationIdOk() (*int64, bool)`

GetTriggeredByRelationIdOk returns a tuple with the TriggeredByRelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByRelationId

`func (o *SeriesModelV1RelatedSeries) SetTriggeredByRelationId(v int64)`

SetTriggeredByRelationId sets TriggeredByRelationId field to given value.

### HasTriggeredByRelationId

`func (o *SeriesModelV1RelatedSeries) HasTriggeredByRelationId() bool`

HasTriggeredByRelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


