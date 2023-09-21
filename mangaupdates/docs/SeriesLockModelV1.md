# SeriesLockModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**TimeLocked** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewSeriesLockModelV1

`func NewSeriesLockModelV1() *SeriesLockModelV1`

NewSeriesLockModelV1 instantiates a new SeriesLockModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesLockModelV1WithDefaults

`func NewSeriesLockModelV1WithDefaults() *SeriesLockModelV1`

NewSeriesLockModelV1WithDefaults instantiates a new SeriesLockModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SeriesLockModelV1) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SeriesLockModelV1) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SeriesLockModelV1) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SeriesLockModelV1) HasField() bool`

HasField returns a boolean if a field has been set.

### GetReason

`func (o *SeriesLockModelV1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SeriesLockModelV1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SeriesLockModelV1) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SeriesLockModelV1) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUserId

`func (o *SeriesLockModelV1) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SeriesLockModelV1) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SeriesLockModelV1) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SeriesLockModelV1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *SeriesLockModelV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SeriesLockModelV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SeriesLockModelV1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SeriesLockModelV1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetTimeLocked

`func (o *SeriesLockModelV1) GetTimeLocked() TimeV1`

GetTimeLocked returns the TimeLocked field if non-nil, zero value otherwise.

### GetTimeLockedOk

`func (o *SeriesLockModelV1) GetTimeLockedOk() (*TimeV1, bool)`

GetTimeLockedOk returns a tuple with the TimeLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLocked

`func (o *SeriesLockModelV1) SetTimeLocked(v TimeV1)`

SetTimeLocked sets TimeLocked field to given value.

### HasTimeLocked

`func (o *SeriesLockModelV1) HasTimeLocked() bool`

HasTimeLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


