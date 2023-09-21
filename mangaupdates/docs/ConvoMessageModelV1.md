# ConvoMessageModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **int64** |  | [optional] 
**ConvoId** | Pointer to **int64** |  | [optional] 
**AuthorId** | Pointer to **int64** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**LastEdit** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewConvoMessageModelV1

`func NewConvoMessageModelV1() *ConvoMessageModelV1`

NewConvoMessageModelV1 instantiates a new ConvoMessageModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvoMessageModelV1WithDefaults

`func NewConvoMessageModelV1WithDefaults() *ConvoMessageModelV1`

NewConvoMessageModelV1WithDefaults instantiates a new ConvoMessageModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *ConvoMessageModelV1) GetMessageId() int64`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ConvoMessageModelV1) GetMessageIdOk() (*int64, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ConvoMessageModelV1) SetMessageId(v int64)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ConvoMessageModelV1) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetConvoId

`func (o *ConvoMessageModelV1) GetConvoId() int64`

GetConvoId returns the ConvoId field if non-nil, zero value otherwise.

### GetConvoIdOk

`func (o *ConvoMessageModelV1) GetConvoIdOk() (*int64, bool)`

GetConvoIdOk returns a tuple with the ConvoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvoId

`func (o *ConvoMessageModelV1) SetConvoId(v int64)`

SetConvoId sets ConvoId field to given value.

### HasConvoId

`func (o *ConvoMessageModelV1) HasConvoId() bool`

HasConvoId returns a boolean if a field has been set.

### GetAuthorId

`func (o *ConvoMessageModelV1) GetAuthorId() int64`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ConvoMessageModelV1) GetAuthorIdOk() (*int64, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ConvoMessageModelV1) SetAuthorId(v int64)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *ConvoMessageModelV1) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *ConvoMessageModelV1) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ConvoMessageModelV1) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ConvoMessageModelV1) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ConvoMessageModelV1) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *ConvoMessageModelV1) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *ConvoMessageModelV1) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *ConvoMessageModelV1) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *ConvoMessageModelV1) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetContent

`func (o *ConvoMessageModelV1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConvoMessageModelV1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConvoMessageModelV1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConvoMessageModelV1) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ConvoMessageModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ConvoMessageModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ConvoMessageModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ConvoMessageModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetLastEdit

`func (o *ConvoMessageModelV1) GetLastEdit() TimeV1`

GetLastEdit returns the LastEdit field if non-nil, zero value otherwise.

### GetLastEditOk

`func (o *ConvoMessageModelV1) GetLastEditOk() (*TimeV1, bool)`

GetLastEditOk returns a tuple with the LastEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEdit

`func (o *ConvoMessageModelV1) SetLastEdit(v TimeV1)`

SetLastEdit sets LastEdit field to given value.

### HasLastEdit

`func (o *ConvoMessageModelV1) HasLastEdit() bool`

HasLastEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


