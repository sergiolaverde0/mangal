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

// checks if the UserChangeRequestModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserChangeRequestModelV1{}

// UserChangeRequestModelV1 struct for UserChangeRequestModelV1
type UserChangeRequestModelV1 struct {
	Id        *int64             `json:"id,omitempty"`
	Body      *string            `json:"body,omitempty"`
	AddedBy   *UserModelSearchV1 `json:"added_by,omitempty"`
	TimeAdded *TimeV1            `json:"time_added,omitempty"`
}

// NewUserChangeRequestModelV1 instantiates a new UserChangeRequestModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChangeRequestModelV1() *UserChangeRequestModelV1 {
	this := UserChangeRequestModelV1{}
	return &this
}

// NewUserChangeRequestModelV1WithDefaults instantiates a new UserChangeRequestModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChangeRequestModelV1WithDefaults() *UserChangeRequestModelV1 {
	this := UserChangeRequestModelV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserChangeRequestModelV1) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelV1) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserChangeRequestModelV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserChangeRequestModelV1) SetId(v int64) {
	o.Id = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UserChangeRequestModelV1) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelV1) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UserChangeRequestModelV1) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UserChangeRequestModelV1) SetBody(v string) {
	o.Body = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *UserChangeRequestModelV1) GetAddedBy() UserModelSearchV1 {
	if o == nil || IsNil(o.AddedBy) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelV1) GetAddedByOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *UserChangeRequestModelV1) HasAddedBy() bool {
	if o != nil && !IsNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given UserModelSearchV1 and assigns it to the AddedBy field.
func (o *UserChangeRequestModelV1) SetAddedBy(v UserModelSearchV1) {
	o.AddedBy = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *UserChangeRequestModelV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *UserChangeRequestModelV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *UserChangeRequestModelV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

func (o UserChangeRequestModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserChangeRequestModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.AddedBy) {
		toSerialize["added_by"] = o.AddedBy
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	return toSerialize, nil
}

type NullableUserChangeRequestModelV1 struct {
	value *UserChangeRequestModelV1
	isSet bool
}

func (v NullableUserChangeRequestModelV1) Get() *UserChangeRequestModelV1 {
	return v.value
}

func (v *NullableUserChangeRequestModelV1) Set(val *UserChangeRequestModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChangeRequestModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChangeRequestModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChangeRequestModelV1(val *UserChangeRequestModelV1) *NullableUserChangeRequestModelV1 {
	return &NullableUserChangeRequestModelV1{value: val, isSet: true}
}

func (v NullableUserChangeRequestModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChangeRequestModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
