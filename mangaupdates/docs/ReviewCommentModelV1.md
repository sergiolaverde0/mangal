# ReviewCommentModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ReviewId** | Pointer to **int64** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**ReviewCommentModelV1Author**](ReviewCommentModelV1Author.md) |  | [optional] 
**Rating** | Pointer to **float32** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**TimeUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewReviewCommentModelV1

`func NewReviewCommentModelV1() *ReviewCommentModelV1`

NewReviewCommentModelV1 instantiates a new ReviewCommentModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCommentModelV1WithDefaults

`func NewReviewCommentModelV1WithDefaults() *ReviewCommentModelV1`

NewReviewCommentModelV1WithDefaults instantiates a new ReviewCommentModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewCommentModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewCommentModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewCommentModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReviewCommentModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReviewId

`func (o *ReviewCommentModelV1) GetReviewId() int64`

GetReviewId returns the ReviewId field if non-nil, zero value otherwise.

### GetReviewIdOk

`func (o *ReviewCommentModelV1) GetReviewIdOk() (*int64, bool)`

GetReviewIdOk returns a tuple with the ReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewId

`func (o *ReviewCommentModelV1) SetReviewId(v int64)`

SetReviewId sets ReviewId field to given value.

### HasReviewId

`func (o *ReviewCommentModelV1) HasReviewId() bool`

HasReviewId returns a boolean if a field has been set.

### GetSubject

`func (o *ReviewCommentModelV1) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReviewCommentModelV1) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReviewCommentModelV1) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ReviewCommentModelV1) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContent

`func (o *ReviewCommentModelV1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReviewCommentModelV1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReviewCommentModelV1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ReviewCommentModelV1) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAuthor

`func (o *ReviewCommentModelV1) GetAuthor() ReviewCommentModelV1Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReviewCommentModelV1) GetAuthorOk() (*ReviewCommentModelV1Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReviewCommentModelV1) SetAuthor(v ReviewCommentModelV1Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ReviewCommentModelV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetRating

`func (o *ReviewCommentModelV1) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ReviewCommentModelV1) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ReviewCommentModelV1) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ReviewCommentModelV1) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ReviewCommentModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ReviewCommentModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ReviewCommentModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ReviewCommentModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetTimeUpdated

`func (o *ReviewCommentModelV1) GetTimeUpdated() TimeV1`

GetTimeUpdated returns the TimeUpdated field if non-nil, zero value otherwise.

### GetTimeUpdatedOk

`func (o *ReviewCommentModelV1) GetTimeUpdatedOk() (*TimeV1, bool)`

GetTimeUpdatedOk returns a tuple with the TimeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUpdated

`func (o *ReviewCommentModelV1) SetTimeUpdated(v TimeV1)`

SetTimeUpdated sets TimeUpdated field to given value.

### HasTimeUpdated

`func (o *ReviewCommentModelV1) HasTimeUpdated() bool`

HasTimeUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


