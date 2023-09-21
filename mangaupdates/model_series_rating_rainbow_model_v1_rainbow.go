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

// checks if the SeriesRatingRainbowModelV1Rainbow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesRatingRainbowModelV1Rainbow{}

// SeriesRatingRainbowModelV1Rainbow struct for SeriesRatingRainbowModelV1Rainbow
type SeriesRatingRainbowModelV1Rainbow struct {
	Rating *int64 `json:"rating,omitempty"`
	Count  *int64 `json:"count,omitempty"`
}

// NewSeriesRatingRainbowModelV1Rainbow instantiates a new SeriesRatingRainbowModelV1Rainbow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesRatingRainbowModelV1Rainbow() *SeriesRatingRainbowModelV1Rainbow {
	this := SeriesRatingRainbowModelV1Rainbow{}
	return &this
}

// NewSeriesRatingRainbowModelV1RainbowWithDefaults instantiates a new SeriesRatingRainbowModelV1Rainbow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesRatingRainbowModelV1RainbowWithDefaults() *SeriesRatingRainbowModelV1Rainbow {
	this := SeriesRatingRainbowModelV1Rainbow{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *SeriesRatingRainbowModelV1Rainbow) GetRating() int64 {
	if o == nil || IsNil(o.Rating) {
		var ret int64
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesRatingRainbowModelV1Rainbow) GetRatingOk() (*int64, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *SeriesRatingRainbowModelV1Rainbow) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int64 and assigns it to the Rating field.
func (o *SeriesRatingRainbowModelV1Rainbow) SetRating(v int64) {
	o.Rating = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SeriesRatingRainbowModelV1Rainbow) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesRatingRainbowModelV1Rainbow) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SeriesRatingRainbowModelV1Rainbow) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SeriesRatingRainbowModelV1Rainbow) SetCount(v int64) {
	o.Count = &v
}

func (o SeriesRatingRainbowModelV1Rainbow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesRatingRainbowModelV1Rainbow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableSeriesRatingRainbowModelV1Rainbow struct {
	value *SeriesRatingRainbowModelV1Rainbow
	isSet bool
}

func (v NullableSeriesRatingRainbowModelV1Rainbow) Get() *SeriesRatingRainbowModelV1Rainbow {
	return v.value
}

func (v *NullableSeriesRatingRainbowModelV1Rainbow) Set(val *SeriesRatingRainbowModelV1Rainbow) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesRatingRainbowModelV1Rainbow) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesRatingRainbowModelV1Rainbow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesRatingRainbowModelV1Rainbow(val *SeriesRatingRainbowModelV1Rainbow) *NullableSeriesRatingRainbowModelV1Rainbow {
	return &NullableSeriesRatingRainbowModelV1Rainbow{value: val, isSet: true}
}

func (v NullableSeriesRatingRainbowModelV1Rainbow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesRatingRainbowModelV1Rainbow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}