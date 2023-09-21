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

// checks if the PerPageSearchRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerPageSearchRequestV1{}

// PerPageSearchRequestV1 struct for PerPageSearchRequestV1
type PerPageSearchRequestV1 struct {
	Page    *int64 `json:"page,omitempty"`
	Perpage *int64 `json:"perpage,omitempty"`
}

// NewPerPageSearchRequestV1 instantiates a new PerPageSearchRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerPageSearchRequestV1() *PerPageSearchRequestV1 {
	this := PerPageSearchRequestV1{}
	return &this
}

// NewPerPageSearchRequestV1WithDefaults instantiates a new PerPageSearchRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerPageSearchRequestV1WithDefaults() *PerPageSearchRequestV1 {
	this := PerPageSearchRequestV1{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PerPageSearchRequestV1) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerPageSearchRequestV1) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PerPageSearchRequestV1) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *PerPageSearchRequestV1) SetPage(v int64) {
	o.Page = &v
}

// GetPerpage returns the Perpage field value if set, zero value otherwise.
func (o *PerPageSearchRequestV1) GetPerpage() int64 {
	if o == nil || IsNil(o.Perpage) {
		var ret int64
		return ret
	}
	return *o.Perpage
}

// GetPerpageOk returns a tuple with the Perpage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerPageSearchRequestV1) GetPerpageOk() (*int64, bool) {
	if o == nil || IsNil(o.Perpage) {
		return nil, false
	}
	return o.Perpage, true
}

// HasPerpage returns a boolean if a field has been set.
func (o *PerPageSearchRequestV1) HasPerpage() bool {
	if o != nil && !IsNil(o.Perpage) {
		return true
	}

	return false
}

// SetPerpage gets a reference to the given int64 and assigns it to the Perpage field.
func (o *PerPageSearchRequestV1) SetPerpage(v int64) {
	o.Perpage = &v
}

func (o PerPageSearchRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerPageSearchRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Perpage) {
		toSerialize["perpage"] = o.Perpage
	}
	return toSerialize, nil
}

type NullablePerPageSearchRequestV1 struct {
	value *PerPageSearchRequestV1
	isSet bool
}

func (v NullablePerPageSearchRequestV1) Get() *PerPageSearchRequestV1 {
	return v.value
}

func (v *NullablePerPageSearchRequestV1) Set(val *PerPageSearchRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePerPageSearchRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePerPageSearchRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerPageSearchRequestV1(val *PerPageSearchRequestV1) *NullablePerPageSearchRequestV1 {
	return &NullablePerPageSearchRequestV1{value: val, isSet: true}
}

func (v NullablePerPageSearchRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerPageSearchRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
