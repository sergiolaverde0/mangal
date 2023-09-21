# FaqCategoryReorderModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** |  | 
**Position** | Pointer to **int64** |  | [optional] 
**Questions** | Pointer to [**[]FaqQuestionReorderModelV1**](FaqQuestionReorderModelV1.md) |  | [optional] 

## Methods

### NewFaqCategoryReorderModelV1

`func NewFaqCategoryReorderModelV1(categoryId int64, ) *FaqCategoryReorderModelV1`

NewFaqCategoryReorderModelV1 instantiates a new FaqCategoryReorderModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaqCategoryReorderModelV1WithDefaults

`func NewFaqCategoryReorderModelV1WithDefaults() *FaqCategoryReorderModelV1`

NewFaqCategoryReorderModelV1WithDefaults instantiates a new FaqCategoryReorderModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *FaqCategoryReorderModelV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FaqCategoryReorderModelV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FaqCategoryReorderModelV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetPosition

`func (o *FaqCategoryReorderModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FaqCategoryReorderModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FaqCategoryReorderModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FaqCategoryReorderModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetQuestions

`func (o *FaqCategoryReorderModelV1) GetQuestions() []FaqQuestionReorderModelV1`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *FaqCategoryReorderModelV1) GetQuestionsOk() (*[]FaqQuestionReorderModelV1, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *FaqCategoryReorderModelV1) SetQuestions(v []FaqQuestionReorderModelV1)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *FaqCategoryReorderModelV1) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


