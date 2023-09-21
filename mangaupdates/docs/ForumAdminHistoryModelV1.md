# ForumAdminHistoryModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**ActionTime** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 

## Methods

### NewForumAdminHistoryModelV1

`func NewForumAdminHistoryModelV1() *ForumAdminHistoryModelV1`

NewForumAdminHistoryModelV1 instantiates a new ForumAdminHistoryModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumAdminHistoryModelV1WithDefaults

`func NewForumAdminHistoryModelV1WithDefaults() *ForumAdminHistoryModelV1`

NewForumAdminHistoryModelV1WithDefaults instantiates a new ForumAdminHistoryModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ForumAdminHistoryModelV1) GetUser() UserModelSearchV1`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ForumAdminHistoryModelV1) GetUserOk() (*UserModelSearchV1, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ForumAdminHistoryModelV1) SetUser(v UserModelSearchV1)`

SetUser sets User field to given value.

### HasUser

`func (o *ForumAdminHistoryModelV1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetActionTime

`func (o *ForumAdminHistoryModelV1) GetActionTime() TimeV1`

GetActionTime returns the ActionTime field if non-nil, zero value otherwise.

### GetActionTimeOk

`func (o *ForumAdminHistoryModelV1) GetActionTimeOk() (*TimeV1, bool)`

GetActionTimeOk returns a tuple with the ActionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTime

`func (o *ForumAdminHistoryModelV1) SetActionTime(v TimeV1)`

SetActionTime sets ActionTime field to given value.

### HasActionTime

`func (o *ForumAdminHistoryModelV1) HasActionTime() bool`

HasActionTime returns a boolean if a field has been set.

### GetAction

`func (o *ForumAdminHistoryModelV1) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ForumAdminHistoryModelV1) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ForumAdminHistoryModelV1) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ForumAdminHistoryModelV1) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


