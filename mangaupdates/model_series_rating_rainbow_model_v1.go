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

// checks if the SeriesRatingRainbowModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesRatingRainbowModelV1{}

// SeriesRatingRainbowModelV1 struct for SeriesRatingRainbowModelV1
type SeriesRatingRainbowModelV1 struct {
	AverageRating *float32 `json:"average_rating,omitempty"`
	Rainbow []SeriesRatingRainbowModelV1Rainbow `json:"rainbow,omitempty"`
}

// NewSeriesRatingRainbowModelV1 instantiates a new SeriesRatingRainbowModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesRatingRainbowModelV1() *SeriesRatingRainbowModelV1 {
	this := SeriesRatingRainbowModelV1{}
	return &this
}

// NewSeriesRatingRainbowModelV1WithDefaults instantiates a new SeriesRatingRainbowModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesRatingRainbowModelV1WithDefaults() *SeriesRatingRainbowModelV1 {
	this := SeriesRatingRainbowModelV1{}
	return &this
}

// GetAverageRating returns the AverageRating field value if set, zero value otherwise.
func (o *SeriesRatingRainbowModelV1) GetAverageRating() float32 {
	if o == nil || IsNil(o.AverageRating) {
		var ret float32
		return ret
	}
	return *o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesRatingRainbowModelV1) GetAverageRatingOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageRating) {
		return nil, false
	}
	return o.AverageRating, true
}

// HasAverageRating returns a boolean if a field has been set.
func (o *SeriesRatingRainbowModelV1) HasAverageRating() bool {
	if o != nil && !IsNil(o.AverageRating) {
		return true
	}

	return false
}

// SetAverageRating gets a reference to the given float32 and assigns it to the AverageRating field.
func (o *SeriesRatingRainbowModelV1) SetAverageRating(v float32) {
	o.AverageRating = &v
}

// GetRainbow returns the Rainbow field value if set, zero value otherwise.
func (o *SeriesRatingRainbowModelV1) GetRainbow() []SeriesRatingRainbowModelV1Rainbow {
	if o == nil || IsNil(o.Rainbow) {
		var ret []SeriesRatingRainbowModelV1Rainbow
		return ret
	}
	return o.Rainbow
}

// GetRainbowOk returns a tuple with the Rainbow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesRatingRainbowModelV1) GetRainbowOk() ([]SeriesRatingRainbowModelV1Rainbow, bool) {
	if o == nil || IsNil(o.Rainbow) {
		return nil, false
	}
	return o.Rainbow, true
}

// HasRainbow returns a boolean if a field has been set.
func (o *SeriesRatingRainbowModelV1) HasRainbow() bool {
	if o != nil && !IsNil(o.Rainbow) {
		return true
	}

	return false
}

// SetRainbow gets a reference to the given []SeriesRatingRainbowModelV1Rainbow and assigns it to the Rainbow field.
func (o *SeriesRatingRainbowModelV1) SetRainbow(v []SeriesRatingRainbowModelV1Rainbow) {
	o.Rainbow = v
}

func (o SeriesRatingRainbowModelV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesRatingRainbowModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageRating) {
		toSerialize["average_rating"] = o.AverageRating
	}
	if !IsNil(o.Rainbow) {
		toSerialize["rainbow"] = o.Rainbow
	}
	return toSerialize, nil
}

type NullableSeriesRatingRainbowModelV1 struct {
	value *SeriesRatingRainbowModelV1
	isSet bool
}

func (v NullableSeriesRatingRainbowModelV1) Get() *SeriesRatingRainbowModelV1 {
	return v.value
}

func (v *NullableSeriesRatingRainbowModelV1) Set(val *SeriesRatingRainbowModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesRatingRainbowModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesRatingRainbowModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesRatingRainbowModelV1(val *SeriesRatingRainbowModelV1) *NullableSeriesRatingRainbowModelV1 {
	return &NullableSeriesRatingRainbowModelV1{value: val, isSet: true}
}

func (v NullableSeriesRatingRainbowModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesRatingRainbowModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


