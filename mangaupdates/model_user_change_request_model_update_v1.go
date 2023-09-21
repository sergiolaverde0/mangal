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

// checks if the UserChangeRequestModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserChangeRequestModelUpdateV1{}

// UserChangeRequestModelUpdateV1 struct for UserChangeRequestModelUpdateV1
type UserChangeRequestModelUpdateV1 struct {
	Body *string `json:"body,omitempty"`
	Archived *bool `json:"archived,omitempty"`
}

// NewUserChangeRequestModelUpdateV1 instantiates a new UserChangeRequestModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChangeRequestModelUpdateV1() *UserChangeRequestModelUpdateV1 {
	this := UserChangeRequestModelUpdateV1{}
	return &this
}

// NewUserChangeRequestModelUpdateV1WithDefaults instantiates a new UserChangeRequestModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChangeRequestModelUpdateV1WithDefaults() *UserChangeRequestModelUpdateV1 {
	this := UserChangeRequestModelUpdateV1{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UserChangeRequestModelUpdateV1) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelUpdateV1) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UserChangeRequestModelUpdateV1) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UserChangeRequestModelUpdateV1) SetBody(v string) {
	o.Body = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UserChangeRequestModelUpdateV1) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChangeRequestModelUpdateV1) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UserChangeRequestModelUpdateV1) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UserChangeRequestModelUpdateV1) SetArchived(v bool) {
	o.Archived = &v
}

func (o UserChangeRequestModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserChangeRequestModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableUserChangeRequestModelUpdateV1 struct {
	value *UserChangeRequestModelUpdateV1
	isSet bool
}

func (v NullableUserChangeRequestModelUpdateV1) Get() *UserChangeRequestModelUpdateV1 {
	return v.value
}

func (v *NullableUserChangeRequestModelUpdateV1) Set(val *UserChangeRequestModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChangeRequestModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChangeRequestModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChangeRequestModelUpdateV1(val *UserChangeRequestModelUpdateV1) *NullableUserChangeRequestModelUpdateV1 {
	return &NullableUserChangeRequestModelUpdateV1{value: val, isSet: true}
}

func (v NullableUserChangeRequestModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChangeRequestModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


