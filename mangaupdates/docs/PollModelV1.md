# PollModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Answers** | Pointer to [**[]PollModelV1Answers**](PollModelV1Answers.md) |  | [optional] 
**TotalVotes** | Pointer to **int64** |  | [optional] 

## Methods

### NewPollModelV1

`func NewPollModelV1() *PollModelV1`

NewPollModelV1 instantiates a new PollModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollModelV1WithDefaults

`func NewPollModelV1WithDefaults() *PollModelV1`

NewPollModelV1WithDefaults instantiates a new PollModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PollModelV1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PollModelV1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PollModelV1) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PollModelV1) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetQuestion

`func (o *PollModelV1) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PollModelV1) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PollModelV1) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *PollModelV1) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswers

`func (o *PollModelV1) GetAnswers() []PollModelV1Answers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *PollModelV1) GetAnswersOk() (*[]PollModelV1Answers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *PollModelV1) SetAnswers(v []PollModelV1Answers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *PollModelV1) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetTotalVotes

`func (o *PollModelV1) GetTotalVotes() int64`

GetTotalVotes returns the TotalVotes field if non-nil, zero value otherwise.

### GetTotalVotesOk

`func (o *PollModelV1) GetTotalVotesOk() (*int64, bool)`

GetTotalVotesOk returns a tuple with the TotalVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVotes

`func (o *PollModelV1) SetTotalVotes(v int64)`

SetTotalVotes sets TotalVotes field to given value.

### HasTotalVotes

`func (o *PollModelV1) HasTotalVotes() bool`

HasTotalVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


