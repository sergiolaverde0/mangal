# AboutusCategoryReorderModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** |  | 
**Position** | Pointer to **int64** |  | [optional] 
**Users** | Pointer to [**[]AboutusUserReorderModelV1**](AboutusUserReorderModelV1.md) |  | [optional] 

## Methods

### NewAboutusCategoryReorderModelV1

`func NewAboutusCategoryReorderModelV1(categoryId int64, ) *AboutusCategoryReorderModelV1`

NewAboutusCategoryReorderModelV1 instantiates a new AboutusCategoryReorderModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutusCategoryReorderModelV1WithDefaults

`func NewAboutusCategoryReorderModelV1WithDefaults() *AboutusCategoryReorderModelV1`

NewAboutusCategoryReorderModelV1WithDefaults instantiates a new AboutusCategoryReorderModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *AboutusCategoryReorderModelV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *AboutusCategoryReorderModelV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *AboutusCategoryReorderModelV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetPosition

`func (o *AboutusCategoryReorderModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AboutusCategoryReorderModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AboutusCategoryReorderModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AboutusCategoryReorderModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetUsers

`func (o *AboutusCategoryReorderModelV1) GetUsers() []AboutusUserReorderModelV1`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AboutusCategoryReorderModelV1) GetUsersOk() (*[]AboutusUserReorderModelV1, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AboutusCategoryReorderModelV1) SetUsers(v []AboutusUserReorderModelV1)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AboutusCategoryReorderModelV1) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


