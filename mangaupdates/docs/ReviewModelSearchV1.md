# ReviewModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**ReviewModelSearchV1Author**](ReviewModelSearchV1Author.md) |  | [optional] 
**Series** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**Review** | Pointer to [**ReviewModelSearchV1Review**](ReviewModelSearchV1Review.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewReviewModelSearchV1

`func NewReviewModelSearchV1() *ReviewModelSearchV1`

NewReviewModelSearchV1 instantiates a new ReviewModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewModelSearchV1WithDefaults

`func NewReviewModelSearchV1WithDefaults() *ReviewModelSearchV1`

NewReviewModelSearchV1WithDefaults instantiates a new ReviewModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewModelSearchV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewModelSearchV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewModelSearchV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReviewModelSearchV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReviewModelSearchV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReviewModelSearchV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReviewModelSearchV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReviewModelSearchV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBodyExcerpt

`func (o *ReviewModelSearchV1) GetBodyExcerpt() string`

GetBodyExcerpt returns the BodyExcerpt field if non-nil, zero value otherwise.

### GetBodyExcerptOk

`func (o *ReviewModelSearchV1) GetBodyExcerptOk() (*string, bool)`

GetBodyExcerptOk returns a tuple with the BodyExcerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyExcerpt

`func (o *ReviewModelSearchV1) SetBodyExcerpt(v string)`

SetBodyExcerpt sets BodyExcerpt field to given value.

### HasBodyExcerpt

`func (o *ReviewModelSearchV1) HasBodyExcerpt() bool`

HasBodyExcerpt returns a boolean if a field has been set.

### GetAuthor

`func (o *ReviewModelSearchV1) GetAuthor() ReviewModelSearchV1Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReviewModelSearchV1) GetAuthorOk() (*ReviewModelSearchV1Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReviewModelSearchV1) SetAuthor(v ReviewModelSearchV1Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReviewModelSearchV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSeries

`func (o *ReviewModelSearchV1) GetSeries() SeriesModelSearchV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ReviewModelSearchV1) GetSeriesOk() (*SeriesModelSearchV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ReviewModelSearchV1) SetSeries(v SeriesModelSearchV1)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ReviewModelSearchV1) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetReview

`func (o *ReviewModelSearchV1) GetReview() ReviewModelSearchV1Review`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *ReviewModelSearchV1) GetReviewOk() (*ReviewModelSearchV1Review, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *ReviewModelSearchV1) SetReview(v ReviewModelSearchV1Review)`

SetReview sets Review field to given value.

### HasReview

`func (o *ReviewModelSearchV1) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ReviewModelSearchV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ReviewModelSearchV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ReviewModelSearchV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ReviewModelSearchV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


