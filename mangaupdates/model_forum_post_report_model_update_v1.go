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

// checks if the ForumPostReportModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumPostReportModelUpdateV1{}

// ForumPostReportModelUpdateV1 struct for ForumPostReportModelUpdateV1
type ForumPostReportModelUpdateV1 struct {
	Reason *string `json:"reason,omitempty"`
}

// NewForumPostReportModelUpdateV1 instantiates a new ForumPostReportModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumPostReportModelUpdateV1() *ForumPostReportModelUpdateV1 {
	this := ForumPostReportModelUpdateV1{}
	return &this
}

// NewForumPostReportModelUpdateV1WithDefaults instantiates a new ForumPostReportModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumPostReportModelUpdateV1WithDefaults() *ForumPostReportModelUpdateV1 {
	this := ForumPostReportModelUpdateV1{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ForumPostReportModelUpdateV1) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelUpdateV1) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ForumPostReportModelUpdateV1) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ForumPostReportModelUpdateV1) SetReason(v string) {
	o.Reason = &v
}

func (o ForumPostReportModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumPostReportModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableForumPostReportModelUpdateV1 struct {
	value *ForumPostReportModelUpdateV1
	isSet bool
}

func (v NullableForumPostReportModelUpdateV1) Get() *ForumPostReportModelUpdateV1 {
	return v.value
}

func (v *NullableForumPostReportModelUpdateV1) Set(val *ForumPostReportModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumPostReportModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumPostReportModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumPostReportModelUpdateV1(val *ForumPostReportModelUpdateV1) *NullableForumPostReportModelUpdateV1 {
	return &NullableForumPostReportModelUpdateV1{value: val, isSet: true}
}

func (v NullableForumPostReportModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumPostReportModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


