# FaqQuestionModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | **int64** |  | 
**Question** | Pointer to **string** |  | [optional] 
**Answer** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 

## Methods

### NewFaqQuestionModelV1

`func NewFaqQuestionModelV1(questionId int64, ) *FaqQuestionModelV1`

NewFaqQuestionModelV1 instantiates a new FaqQuestionModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaqQuestionModelV1WithDefaults

`func NewFaqQuestionModelV1WithDefaults() *FaqQuestionModelV1`

NewFaqQuestionModelV1WithDefaults instantiates a new FaqQuestionModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *FaqQuestionModelV1) GetQuestionId() int64`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *FaqQuestionModelV1) GetQuestionIdOk() (*int64, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *FaqQuestionModelV1) SetQuestionId(v int64)`

SetQuestionId sets QuestionId field to given value.


### GetQuestion

`func (o *FaqQuestionModelV1) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FaqQuestionModelV1) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FaqQuestionModelV1) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *FaqQuestionModelV1) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *FaqQuestionModelV1) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *FaqQuestionModelV1) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *FaqQuestionModelV1) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *FaqQuestionModelV1) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetPosition

`func (o *FaqQuestionModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FaqQuestionModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FaqQuestionModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FaqQuestionModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


