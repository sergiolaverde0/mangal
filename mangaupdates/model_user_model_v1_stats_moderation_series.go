/*
MangaUpdates API

This API powers our website. Most API functions are public and do not require an account. For user-based functions, you must register an account. Currently, user registration must be done through our main website and is disabled via this API.  Download OpenAPI Specification: [openapi.yaml](openapi.yaml)  Please contact us at the following emails if you have questions:  * lambchopsil@mangaupdates.com * manick@mangaupdates.com  ## Warranties  MangaUpdates makes no warranties about service availability, correctness of the data, or anything else. The service is provided as is, and may change at any time.  ## Acceptable Use Policy  * You will credit MangaUpdates when using data provided by this API. * You will use reasonable spacing between requests so as not to overwhelm the MangaUpdates servers, and employ caching mechanisms when accessing data. * You will NOT use MangaUpdates data or API in a way that will:     * Deceive or defraud users     * Assist or perform an illegal action     * Create spam     * Damage the database  We reserve the right to change this policy at any time.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserModelV1StatsModerationSeries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModelV1StatsModerationSeries{}

// UserModelV1StatsModerationSeries struct for UserModelV1StatsModerationSeries
type UserModelV1StatsModerationSeries struct {
	Approved *int64 `json:"approved,omitempty"`
	Rejected *int64 `json:"rejected,omitempty"`
	Deleted  *int64 `json:"deleted,omitempty"`
}

// NewUserModelV1StatsModerationSeries instantiates a new UserModelV1StatsModerationSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModelV1StatsModerationSeries() *UserModelV1StatsModerationSeries {
	this := UserModelV1StatsModerationSeries{}
	return &this
}

// NewUserModelV1StatsModerationSeriesWithDefaults instantiates a new UserModelV1StatsModerationSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelV1StatsModerationSeriesWithDefaults() *UserModelV1StatsModerationSeries {
	this := UserModelV1StatsModerationSeries{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *UserModelV1StatsModerationSeries) GetApproved() int64 {
	if o == nil || IsNil(o.Approved) {
		var ret int64
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModerationSeries) GetApprovedOk() (*int64, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *UserModelV1StatsModerationSeries) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given int64 and assigns it to the Approved field.
func (o *UserModelV1StatsModerationSeries) SetApproved(v int64) {
	o.Approved = &v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *UserModelV1StatsModerationSeries) GetRejected() int64 {
	if o == nil || IsNil(o.Rejected) {
		var ret int64
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModerationSeries) GetRejectedOk() (*int64, bool) {
	if o == nil || IsNil(o.Rejected) {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *UserModelV1StatsModerationSeries) HasRejected() bool {
	if o != nil && !IsNil(o.Rejected) {
		return true
	}

	return false
}

// SetRejected gets a reference to the given int64 and assigns it to the Rejected field.
func (o *UserModelV1StatsModerationSeries) SetRejected(v int64) {
	o.Rejected = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *UserModelV1StatsModerationSeries) GetDeleted() int64 {
	if o == nil || IsNil(o.Deleted) {
		var ret int64
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModerationSeries) GetDeletedOk() (*int64, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *UserModelV1StatsModerationSeries) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int64 and assigns it to the Deleted field.
func (o *UserModelV1StatsModerationSeries) SetDeleted(v int64) {
	o.Deleted = &v
}

func (o UserModelV1StatsModerationSeries) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModelV1StatsModerationSeries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.Rejected) {
		toSerialize["rejected"] = o.Rejected
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableUserModelV1StatsModerationSeries struct {
	value *UserModelV1StatsModerationSeries
	isSet bool
}

func (v NullableUserModelV1StatsModerationSeries) Get() *UserModelV1StatsModerationSeries {
	return v.value
}

func (v *NullableUserModelV1StatsModerationSeries) Set(val *UserModelV1StatsModerationSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModelV1StatsModerationSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModelV1StatsModerationSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModelV1StatsModerationSeries(val *UserModelV1StatsModerationSeries) *NullableUserModelV1StatsModerationSeries {
	return &NullableUserModelV1StatsModerationSeries{value: val, isSet: true}
}

func (v NullableUserModelV1StatsModerationSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModelV1StatsModerationSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
