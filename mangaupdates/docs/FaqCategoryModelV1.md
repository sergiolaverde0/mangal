# FaqCategoryModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 

## Methods

### NewFaqCategoryModelV1

`func NewFaqCategoryModelV1(categoryId int64, ) *FaqCategoryModelV1`

NewFaqCategoryModelV1 instantiates a new FaqCategoryModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaqCategoryModelV1WithDefaults

`func NewFaqCategoryModelV1WithDefaults() *FaqCategoryModelV1`

NewFaqCategoryModelV1WithDefaults instantiates a new FaqCategoryModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *FaqCategoryModelV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *FaqCategoryModelV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *FaqCategoryModelV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetTitle

`func (o *FaqCategoryModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FaqCategoryModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FaqCategoryModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FaqCategoryModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPosition

`func (o *FaqCategoryModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FaqCategoryModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FaqCategoryModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FaqCategoryModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


