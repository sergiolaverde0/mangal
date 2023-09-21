# ForumPollAnswerModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerId** | **int64** |  | 
**Answer** | **string** |  | 
**Votes** | Pointer to **int64** |  | [optional] 
**Image** | Pointer to [**ForumPollAnswerModelV1Image**](ForumPollAnswerModelV1Image.md) |  | [optional] 

## Methods

### NewForumPollAnswerModelV1

`func NewForumPollAnswerModelV1(answerId int64, answer string, ) *ForumPollAnswerModelV1`

NewForumPollAnswerModelV1 instantiates a new ForumPollAnswerModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPollAnswerModelV1WithDefaults

`func NewForumPollAnswerModelV1WithDefaults() *ForumPollAnswerModelV1`

NewForumPollAnswerModelV1WithDefaults instantiates a new ForumPollAnswerModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerId

`func (o *ForumPollAnswerModelV1) GetAnswerId() int64`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *ForumPollAnswerModelV1) GetAnswerIdOk() (*int64, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *ForumPollAnswerModelV1) SetAnswerId(v int64)`

SetAnswerId sets AnswerId field to given value.


### GetAnswer

`func (o *ForumPollAnswerModelV1) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ForumPollAnswerModelV1) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ForumPollAnswerModelV1) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetVotes

`func (o *ForumPollAnswerModelV1) GetVotes() int64`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *ForumPollAnswerModelV1) GetVotesOk() (*int64, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *ForumPollAnswerModelV1) SetVotes(v int64)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *ForumPollAnswerModelV1) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetImage

`func (o *ForumPollAnswerModelV1) GetImage() ForumPollAnswerModelV1Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ForumPollAnswerModelV1) GetImageOk() (*ForumPollAnswerModelV1Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ForumPollAnswerModelV1) SetImage(v ForumPollAnswerModelV1Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *ForumPollAnswerModelV1) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


