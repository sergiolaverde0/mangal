# ForumWarnModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int64** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Level** | **int64** |  | 
**Reason** | **string** |  | 
**SendReason** | Pointer to **bool** |  | [optional] 
**ByUser** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 

## Methods

### NewForumWarnModelV1

`func NewForumWarnModelV1(level int64, reason string, ) *ForumWarnModelV1`

NewForumWarnModelV1 instantiates a new ForumWarnModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumWarnModelV1WithDefaults

`func NewForumWarnModelV1WithDefaults() *ForumWarnModelV1`

NewForumWarnModelV1WithDefaults instantiates a new ForumWarnModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ForumWarnModelV1) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ForumWarnModelV1) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ForumWarnModelV1) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ForumWarnModelV1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ForumWarnModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ForumWarnModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ForumWarnModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ForumWarnModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetLevel

`func (o *ForumWarnModelV1) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ForumWarnModelV1) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ForumWarnModelV1) SetLevel(v int64)`

SetLevel sets Level field to given value.


### GetReason

`func (o *ForumWarnModelV1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ForumWarnModelV1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ForumWarnModelV1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSendReason

`func (o *ForumWarnModelV1) GetSendReason() bool`

GetSendReason returns the SendReason field if non-nil, zero value otherwise.

### GetSendReasonOk

`func (o *ForumWarnModelV1) GetSendReasonOk() (*bool, bool)`

GetSendReasonOk returns a tuple with the SendReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReason

`func (o *ForumWarnModelV1) SetSendReason(v bool)`

SetSendReason sets SendReason field to given value.

### HasSendReason

`func (o *ForumWarnModelV1) HasSendReason() bool`

HasSendReason returns a boolean if a field has been set.

### GetByUser

`func (o *ForumWarnModelV1) GetByUser() UserModelSearchV1`

GetByUser returns the ByUser field if non-nil, zero value otherwise.

### GetByUserOk

`func (o *ForumWarnModelV1) GetByUserOk() (*UserModelSearchV1, bool)`

GetByUserOk returns a tuple with the ByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUser

`func (o *ForumWarnModelV1) SetByUser(v UserModelSearchV1)`

SetByUser sets ByUser field to given value.

### HasByUser

`func (o *ForumWarnModelV1) HasByUser() bool`

HasByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


