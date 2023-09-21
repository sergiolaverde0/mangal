# SeriesRatingModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | **float32** |  | 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewSeriesRatingModelV1

`func NewSeriesRatingModelV1(rating float32, ) *SeriesRatingModelV1`

NewSeriesRatingModelV1 instantiates a new SeriesRatingModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesRatingModelV1WithDefaults

`func NewSeriesRatingModelV1WithDefaults() *SeriesRatingModelV1`

NewSeriesRatingModelV1WithDefaults instantiates a new SeriesRatingModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *SeriesRatingModelV1) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SeriesRatingModelV1) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SeriesRatingModelV1) SetRating(v float32)`

SetRating sets Rating field to given value.


### GetLastUpdated

`func (o *SeriesRatingModelV1) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SeriesRatingModelV1) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SeriesRatingModelV1) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SeriesRatingModelV1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


