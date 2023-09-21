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

// checks if the SeriesModelV1Anime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelV1Anime{}

// SeriesModelV1Anime struct for SeriesModelV1Anime
type SeriesModelV1Anime struct {
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
}

// NewSeriesModelV1Anime instantiates a new SeriesModelV1Anime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelV1Anime() *SeriesModelV1Anime {
	this := SeriesModelV1Anime{}
	return &this
}

// NewSeriesModelV1AnimeWithDefaults instantiates a new SeriesModelV1Anime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelV1AnimeWithDefaults() *SeriesModelV1Anime {
	this := SeriesModelV1Anime{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SeriesModelV1Anime) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Anime) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SeriesModelV1Anime) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *SeriesModelV1Anime) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SeriesModelV1Anime) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Anime) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SeriesModelV1Anime) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *SeriesModelV1Anime) SetEnd(v string) {
	o.End = &v
}

func (o SeriesModelV1Anime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelV1Anime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableSeriesModelV1Anime struct {
	value *SeriesModelV1Anime
	isSet bool
}

func (v NullableSeriesModelV1Anime) Get() *SeriesModelV1Anime {
	return v.value
}

func (v *NullableSeriesModelV1Anime) Set(val *SeriesModelV1Anime) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelV1Anime) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelV1Anime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelV1Anime(val *SeriesModelV1Anime) *NullableSeriesModelV1Anime {
	return &NullableSeriesModelV1Anime{value: val, isSet: true}
}

func (v NullableSeriesModelV1Anime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelV1Anime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

