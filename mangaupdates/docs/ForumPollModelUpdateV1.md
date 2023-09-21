# ForumPollModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to [**[]ForumPollAnswerModelUpdateV1**](ForumPollAnswerModelUpdateV1.md) |  | [optional] 

## Methods

### NewForumPollModelUpdateV1

`func NewForumPollModelUpdateV1() *ForumPollModelUpdateV1`

NewForumPollModelUpdateV1 instantiates a new ForumPollModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPollModelUpdateV1WithDefaults

`func NewForumPollModelUpdateV1WithDefaults() *ForumPollModelUpdateV1`

NewForumPollModelUpdateV1WithDefaults instantiates a new ForumPollModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *ForumPollModelUpdateV1) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ForumPollModelUpdateV1) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ForumPollModelUpdateV1) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ForumPollModelUpdateV1) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswers

`func (o *ForumPollModelUpdateV1) GetAnswers() []ForumPollAnswerModelUpdateV1`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *ForumPollModelUpdateV1) GetAnswersOk() (*[]ForumPollAnswerModelUpdateV1, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *ForumPollModelUpdateV1) SetAnswers(v []ForumPollAnswerModelUpdateV1)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *ForumPollModelUpdateV1) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


