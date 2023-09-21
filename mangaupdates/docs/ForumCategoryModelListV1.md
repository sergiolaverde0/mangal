# ForumCategoryModelListV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Forums** | Pointer to [**[]ForumCategoryModelListV1Forums**](ForumCategoryModelListV1Forums.md) |  | [optional] 

## Methods

### NewForumCategoryModelListV1

`func NewForumCategoryModelListV1() *ForumCategoryModelListV1`

NewForumCategoryModelListV1 instantiates a new ForumCategoryModelListV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumCategoryModelListV1WithDefaults

`func NewForumCategoryModelListV1WithDefaults() *ForumCategoryModelListV1`

NewForumCategoryModelListV1WithDefaults instantiates a new ForumCategoryModelListV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *ForumCategoryModelListV1) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ForumCategoryModelListV1) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ForumCategoryModelListV1) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ForumCategoryModelListV1) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetName

`func (o *ForumCategoryModelListV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumCategoryModelListV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumCategoryModelListV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForumCategoryModelListV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetForums

`func (o *ForumCategoryModelListV1) GetForums() []ForumCategoryModelListV1Forums`

GetForums returns the Forums field if non-nil, zero value otherwise.

### GetForumsOk

`func (o *ForumCategoryModelListV1) GetForumsOk() (*[]ForumCategoryModelListV1Forums, bool)`

GetForumsOk returns a tuple with the Forums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForums

`func (o *ForumCategoryModelListV1) SetForums(v []ForumCategoryModelListV1Forums)`

SetForums sets Forums field to given value.

### HasForums

`func (o *ForumCategoryModelListV1) HasForums() bool`

HasForums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


