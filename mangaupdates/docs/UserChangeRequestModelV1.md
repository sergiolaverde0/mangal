# UserChangeRequestModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewUserChangeRequestModelV1

`func NewUserChangeRequestModelV1() *UserChangeRequestModelV1`

NewUserChangeRequestModelV1 instantiates a new UserChangeRequestModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChangeRequestModelV1WithDefaults

`func NewUserChangeRequestModelV1WithDefaults() *UserChangeRequestModelV1`

NewUserChangeRequestModelV1WithDefaults instantiates a new UserChangeRequestModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserChangeRequestModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserChangeRequestModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserChangeRequestModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserChangeRequestModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *UserChangeRequestModelV1) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UserChangeRequestModelV1) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UserChangeRequestModelV1) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UserChangeRequestModelV1) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddedBy

`func (o *UserChangeRequestModelV1) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *UserChangeRequestModelV1) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *UserChangeRequestModelV1) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *UserChangeRequestModelV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetTimeAdded

`func (o *UserChangeRequestModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *UserChangeRequestModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *UserChangeRequestModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *UserChangeRequestModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


