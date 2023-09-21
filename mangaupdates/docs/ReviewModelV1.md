# ReviewModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**ReviewModelV1Author**](ReviewModelV1Author.md) |  | [optional] 
**Series** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**Review** | Pointer to [**ReviewModelV1Review**](ReviewModelV1Review.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewReviewModelV1

`func NewReviewModelV1() *ReviewModelV1`

NewReviewModelV1 instantiates a new ReviewModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewModelV1WithDefaults

`func NewReviewModelV1WithDefaults() *ReviewModelV1`

NewReviewModelV1WithDefaults instantiates a new ReviewModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReviewModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReviewModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReviewModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReviewModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReviewModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *ReviewModelV1) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReviewModelV1) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReviewModelV1) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReviewModelV1) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAuthor

`func (o *ReviewModelV1) GetAuthor() ReviewModelV1Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReviewModelV1) GetAuthorOk() (*ReviewModelV1Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReviewModelV1) SetAuthor(v ReviewModelV1Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReviewModelV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSeries

`func (o *ReviewModelV1) GetSeries() SeriesModelSearchV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ReviewModelV1) GetSeriesOk() (*SeriesModelSearchV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ReviewModelV1) SetSeries(v SeriesModelSearchV1)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ReviewModelV1) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetReview

`func (o *ReviewModelV1) GetReview() ReviewModelV1Review`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *ReviewModelV1) GetReviewOk() (*ReviewModelV1Review, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *ReviewModelV1) SetReview(v ReviewModelV1Review)`

SetReview sets Review field to given value.

### HasReview

`func (o *ReviewModelV1) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ReviewModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ReviewModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ReviewModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ReviewModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


