# TimeV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** |  | [optional] 
**AsRfc3339** | Pointer to **time.Time** |  | [optional] 
**AsString** | Pointer to **string** |  | [optional] 

## Methods

### NewTimeV1

`func NewTimeV1() *TimeV1`

NewTimeV1 instantiates a new TimeV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeV1WithDefaults

`func NewTimeV1WithDefaults() *TimeV1`

NewTimeV1WithDefaults instantiates a new TimeV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TimeV1) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TimeV1) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TimeV1) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TimeV1) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAsRfc3339

`func (o *TimeV1) GetAsRfc3339() time.Time`

GetAsRfc3339 returns the AsRfc3339 field if non-nil, zero value otherwise.

### GetAsRfc3339Ok

`func (o *TimeV1) GetAsRfc3339Ok() (*time.Time, bool)`

GetAsRfc3339Ok returns a tuple with the AsRfc3339 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsRfc3339

`func (o *TimeV1) SetAsRfc3339(v time.Time)`

SetAsRfc3339 sets AsRfc3339 field to given value.

### HasAsRfc3339

`func (o *TimeV1) HasAsRfc3339() bool`

HasAsRfc3339 returns a boolean if a field has been set.

### GetAsString

`func (o *TimeV1) GetAsString() string`

GetAsString returns the AsString field if non-nil, zero value otherwise.

### GetAsStringOk

`func (o *TimeV1) GetAsStringOk() (*string, bool)`

GetAsStringOk returns a tuple with the AsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsString

`func (o *TimeV1) SetAsString(v string)`

SetAsString sets AsString field to given value.

### HasAsString

`func (o *TimeV1) HasAsString() bool`

HasAsString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


