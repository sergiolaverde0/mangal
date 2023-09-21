# ApiResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiResponseV1

`func NewApiResponseV1(status string, reason string, ) *ApiResponseV1`

NewApiResponseV1 instantiates a new ApiResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseV1WithDefaults

`func NewApiResponseV1WithDefaults() *ApiResponseV1`

NewApiResponseV1WithDefaults instantiates a new ApiResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiResponseV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiResponseV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiResponseV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ApiResponseV1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiResponseV1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiResponseV1) SetReason(v string)`

SetReason sets Reason field to given value.


### GetContext

`func (o *ApiResponseV1) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ApiResponseV1) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ApiResponseV1) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ApiResponseV1) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


