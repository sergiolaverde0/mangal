# ReviewCommentModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **float32** |  | [optional] 
**Admin** | Pointer to [**ReviewCommentModelUpdateV1Admin**](ReviewCommentModelUpdateV1Admin.md) |  | [optional] 

## Methods

### NewReviewCommentModelUpdateV1

`func NewReviewCommentModelUpdateV1() *ReviewCommentModelUpdateV1`

NewReviewCommentModelUpdateV1 instantiates a new ReviewCommentModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewCommentModelUpdateV1WithDefaults

`func NewReviewCommentModelUpdateV1WithDefaults() *ReviewCommentModelUpdateV1`

NewReviewCommentModelUpdateV1WithDefaults instantiates a new ReviewCommentModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *ReviewCommentModelUpdateV1) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ReviewCommentModelUpdateV1) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ReviewCommentModelUpdateV1) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ReviewCommentModelUpdateV1) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContent

`func (o *ReviewCommentModelUpdateV1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ReviewCommentModelUpdateV1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ReviewCommentModelUpdateV1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ReviewCommentModelUpdateV1) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRating

`func (o *ReviewCommentModelUpdateV1) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ReviewCommentModelUpdateV1) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ReviewCommentModelUpdateV1) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ReviewCommentModelUpdateV1) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetAdmin

`func (o *ReviewCommentModelUpdateV1) GetAdmin() ReviewCommentModelUpdateV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ReviewCommentModelUpdateV1) GetAdminOk() (*ReviewCommentModelUpdateV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ReviewCommentModelUpdateV1) SetAdmin(v ReviewCommentModelUpdateV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ReviewCommentModelUpdateV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


