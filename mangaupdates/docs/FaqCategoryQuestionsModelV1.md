# FaqCategoryQuestionsModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Questions** | Pointer to [**[]FaqQuestionOnlyModelV1**](FaqQuestionOnlyModelV1.md) |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 

## Methods

### NewFaqCategoryQuestionsModelV1

`func NewFaqCategoryQuestionsModelV1(categoryId int64, ) *FaqCategoryQuestionsModelV1`

NewFaqCategoryQuestionsModelV1 instantiates a new FaqCategoryQuestionsModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaqCategoryQuestionsModelV1WithDefaults

`func NewFaqCategoryQuestionsModelV1WithDefaults() *FaqCategoryQuestionsModelV1`

NewFaqCategoryQuestionsModelV1WithDefaults instantiates a new FaqCategoryQuestionsModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *FaqCategoryQuestionsModelV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FaqCategoryQuestionsModelV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FaqCategoryQuestionsModelV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetTitle

`func (o *FaqCategoryQuestionsModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FaqCategoryQuestionsModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FaqCategoryQuestionsModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FaqCategoryQuestionsModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuestions

`func (o *FaqCategoryQuestionsModelV1) GetQuestions() []FaqQuestionOnlyModelV1`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *FaqCategoryQuestionsModelV1) GetQuestionsOk() (*[]FaqQuestionOnlyModelV1, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *FaqCategoryQuestionsModelV1) SetQuestions(v []FaqQuestionOnlyModelV1)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *FaqCategoryQuestionsModelV1) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetPosition

`func (o *FaqCategoryQuestionsModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FaqCategoryQuestionsModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FaqCategoryQuestionsModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FaqCategoryQuestionsModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


