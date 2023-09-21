# AboutusCategoryModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** |  | 
**Position** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]AboutusUserModelV1**](AboutusUserModelV1.md) |  | [optional] 

## Methods

### NewAboutusCategoryModelV1

`func NewAboutusCategoryModelV1(categoryId int64, ) *AboutusCategoryModelV1`

NewAboutusCategoryModelV1 instantiates a new AboutusCategoryModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutusCategoryModelV1WithDefaults

`func NewAboutusCategoryModelV1WithDefaults() *AboutusCategoryModelV1`

NewAboutusCategoryModelV1WithDefaults instantiates a new AboutusCategoryModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *AboutusCategoryModelV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *AboutusCategoryModelV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *AboutusCategoryModelV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetPosition

`func (o *AboutusCategoryModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AboutusCategoryModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AboutusCategoryModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AboutusCategoryModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *AboutusCategoryModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AboutusCategoryModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AboutusCategoryModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AboutusCategoryModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsers

`func (o *AboutusCategoryModelV1) GetUsers() []AboutusUserModelV1`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AboutusCategoryModelV1) GetUsersOk() (*[]AboutusUserModelV1, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AboutusCategoryModelV1) SetUsers(v []AboutusUserModelV1)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AboutusCategoryModelV1) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


