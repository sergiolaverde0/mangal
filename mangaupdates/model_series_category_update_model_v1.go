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

// checks if the SeriesCategoryUpdateModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesCategoryUpdateModelV1{}

// SeriesCategoryUpdateModelV1 struct for SeriesCategoryUpdateModelV1
type SeriesCategoryUpdateModelV1 struct {
	From *CategoriesModelUpdateV1 `json:"from,omitempty"`
	To *CategoriesModelUpdateV1 `json:"to,omitempty"`
}

// NewSeriesCategoryUpdateModelV1 instantiates a new SeriesCategoryUpdateModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesCategoryUpdateModelV1() *SeriesCategoryUpdateModelV1 {
	this := SeriesCategoryUpdateModelV1{}
	return &this
}

// NewSeriesCategoryUpdateModelV1WithDefaults instantiates a new SeriesCategoryUpdateModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesCategoryUpdateModelV1WithDefaults() *SeriesCategoryUpdateModelV1 {
	this := SeriesCategoryUpdateModelV1{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SeriesCategoryUpdateModelV1) GetFrom() CategoriesModelUpdateV1 {
	if o == nil || IsNil(o.From) {
		var ret CategoriesModelUpdateV1
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCategoryUpdateModelV1) GetFromOk() (*CategoriesModelUpdateV1, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SeriesCategoryUpdateModelV1) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given CategoriesModelUpdateV1 and assigns it to the From field.
func (o *SeriesCategoryUpdateModelV1) SetFrom(v CategoriesModelUpdateV1) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SeriesCategoryUpdateModelV1) GetTo() CategoriesModelUpdateV1 {
	if o == nil || IsNil(o.To) {
		var ret CategoriesModelUpdateV1
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCategoryUpdateModelV1) GetToOk() (*CategoriesModelUpdateV1, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SeriesCategoryUpdateModelV1) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given CategoriesModelUpdateV1 and assigns it to the To field.
func (o *SeriesCategoryUpdateModelV1) SetTo(v CategoriesModelUpdateV1) {
	o.To = &v
}

func (o SeriesCategoryUpdateModelV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesCategoryUpdateModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableSeriesCategoryUpdateModelV1 struct {
	value *SeriesCategoryUpdateModelV1
	isSet bool
}

func (v NullableSeriesCategoryUpdateModelV1) Get() *SeriesCategoryUpdateModelV1 {
	return v.value
}

func (v *NullableSeriesCategoryUpdateModelV1) Set(val *SeriesCategoryUpdateModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesCategoryUpdateModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesCategoryUpdateModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesCategoryUpdateModelV1(val *SeriesCategoryUpdateModelV1) *NullableSeriesCategoryUpdateModelV1 {
	return &NullableSeriesCategoryUpdateModelV1{value: val, isSet: true}
}

func (v NullableSeriesCategoryUpdateModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesCategoryUpdateModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


