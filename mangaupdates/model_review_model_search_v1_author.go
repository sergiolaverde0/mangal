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

// checks if the ReviewModelSearchV1Author type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewModelSearchV1Author{}

// ReviewModelSearchV1Author struct for ReviewModelSearchV1Author
type ReviewModelSearchV1Author struct {
	UserId *int64  `json:"user_id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// NewReviewModelSearchV1Author instantiates a new ReviewModelSearchV1Author object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewModelSearchV1Author() *ReviewModelSearchV1Author {
	this := ReviewModelSearchV1Author{}
	return &this
}

// NewReviewModelSearchV1AuthorWithDefaults instantiates a new ReviewModelSearchV1Author object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewModelSearchV1AuthorWithDefaults() *ReviewModelSearchV1Author {
	this := ReviewModelSearchV1Author{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ReviewModelSearchV1Author) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewModelSearchV1Author) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ReviewModelSearchV1Author) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *ReviewModelSearchV1Author) SetUserId(v int64) {
	o.UserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewModelSearchV1Author) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewModelSearchV1Author) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewModelSearchV1Author) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewModelSearchV1Author) SetName(v string) {
	o.Name = &v
}

func (o ReviewModelSearchV1Author) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewModelSearchV1Author) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableReviewModelSearchV1Author struct {
	value *ReviewModelSearchV1Author
	isSet bool
}

func (v NullableReviewModelSearchV1Author) Get() *ReviewModelSearchV1Author {
	return v.value
}

func (v *NullableReviewModelSearchV1Author) Set(val *ReviewModelSearchV1Author) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewModelSearchV1Author) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewModelSearchV1Author) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewModelSearchV1Author(val *ReviewModelSearchV1Author) *NullableReviewModelSearchV1Author {
	return &NullableReviewModelSearchV1Author{value: val, isSet: true}
}

func (v NullableReviewModelSearchV1Author) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewModelSearchV1Author) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
