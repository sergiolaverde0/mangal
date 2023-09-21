# ForumPollModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to [**[]ForumPollAnswerModelV1**](ForumPollAnswerModelV1.md) |  | [optional] 
**Votes** | Pointer to **int64** |  | [optional] 
**Admin** | Pointer to [**ForumPollModelV1Admin**](ForumPollModelV1Admin.md) |  | [optional] 

## Methods

### NewForumPollModelV1

`func NewForumPollModelV1() *ForumPollModelV1`

NewForumPollModelV1 instantiates a new ForumPollModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPollModelV1WithDefaults

`func NewForumPollModelV1WithDefaults() *ForumPollModelV1`

NewForumPollModelV1WithDefaults instantiates a new ForumPollModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *ForumPollModelV1) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ForumPollModelV1) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ForumPollModelV1) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ForumPollModelV1) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswers

`func (o *ForumPollModelV1) GetAnswers() []ForumPollAnswerModelV1`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *ForumPollModelV1) GetAnswersOk() (*[]ForumPollAnswerModelV1, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *ForumPollModelV1) SetAnswers(v []ForumPollAnswerModelV1)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *ForumPollModelV1) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetVotes

`func (o *ForumPollModelV1) GetVotes() int64`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *ForumPollModelV1) GetVotesOk() (*int64, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *ForumPollModelV1) SetVotes(v int64)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *ForumPollModelV1) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetAdmin

`func (o *ForumPollModelV1) GetAdmin() ForumPollModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ForumPollModelV1) GetAdminOk() (*ForumPollModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ForumPollModelV1) SetAdmin(v ForumPollModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ForumPollModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


