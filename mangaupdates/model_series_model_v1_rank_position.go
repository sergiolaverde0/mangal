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

// checks if the SeriesModelV1RankPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelV1RankPosition{}

// SeriesModelV1RankPosition struct for SeriesModelV1RankPosition
type SeriesModelV1RankPosition struct {
	Week        *int64 `json:"week,omitempty"`
	Month       *int64 `json:"month,omitempty"`
	ThreeMonths *int64 `json:"three_months,omitempty"`
	SixMonths   *int64 `json:"six_months,omitempty"`
	Year        *int64 `json:"year,omitempty"`
}

// NewSeriesModelV1RankPosition instantiates a new SeriesModelV1RankPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelV1RankPosition() *SeriesModelV1RankPosition {
	this := SeriesModelV1RankPosition{}
	return &this
}

// NewSeriesModelV1RankPositionWithDefaults instantiates a new SeriesModelV1RankPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelV1RankPositionWithDefaults() *SeriesModelV1RankPosition {
	this := SeriesModelV1RankPosition{}
	return &this
}

// GetWeek returns the Week field value if set, zero value otherwise.
func (o *SeriesModelV1RankPosition) GetWeek() int64 {
	if o == nil || IsNil(o.Week) {
		var ret int64
		return ret
	}
	return *o.Week
}

// GetWeekOk returns a tuple with the Week field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1RankPosition) GetWeekOk() (*int64, bool) {
	if o == nil || IsNil(o.Week) {
		return nil, false
	}
	return o.Week, true
}

// HasWeek returns a boolean if a field has been set.
func (o *SeriesModelV1RankPosition) HasWeek() bool {
	if o != nil && !IsNil(o.Week) {
		return true
	}

	return false
}

// SetWeek gets a reference to the given int64 and assigns it to the Week field.
func (o *SeriesModelV1RankPosition) SetWeek(v int64) {
	o.Week = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *SeriesModelV1RankPosition) GetMonth() int64 {
	if o == nil || IsNil(o.Month) {
		var ret int64
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1RankPosition) GetMonthOk() (*int64, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *SeriesModelV1RankPosition) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int64 and assigns it to the Month field.
func (o *SeriesModelV1RankPosition) SetMonth(v int64) {
	o.Month = &v
}

// GetThreeMonths returns the ThreeMonths field value if set, zero value otherwise.
func (o *SeriesModelV1RankPosition) GetThreeMonths() int64 {
	if o == nil || IsNil(o.ThreeMonths) {
		var ret int64
		return ret
	}
	return *o.ThreeMonths
}

// GetThreeMonthsOk returns a tuple with the ThreeMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1RankPosition) GetThreeMonthsOk() (*int64, bool) {
	if o == nil || IsNil(o.ThreeMonths) {
		return nil, false
	}
	return o.ThreeMonths, true
}

// HasThreeMonths returns a boolean if a field has been set.
func (o *SeriesModelV1RankPosition) HasThreeMonths() bool {
	if o != nil && !IsNil(o.ThreeMonths) {
		return true
	}

	return false
}

// SetThreeMonths gets a reference to the given int64 and assigns it to the ThreeMonths field.
func (o *SeriesModelV1RankPosition) SetThreeMonths(v int64) {
	o.ThreeMonths = &v
}

// GetSixMonths returns the SixMonths field value if set, zero value otherwise.
func (o *SeriesModelV1RankPosition) GetSixMonths() int64 {
	if o == nil || IsNil(o.SixMonths) {
		var ret int64
		return ret
	}
	return *o.SixMonths
}

// GetSixMonthsOk returns a tuple with the SixMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1RankPosition) GetSixMonthsOk() (*int64, bool) {
	if o == nil || IsNil(o.SixMonths) {
		return nil, false
	}
	return o.SixMonths, true
}

// HasSixMonths returns a boolean if a field has been set.
func (o *SeriesModelV1RankPosition) HasSixMonths() bool {
	if o != nil && !IsNil(o.SixMonths) {
		return true
	}

	return false
}

// SetSixMonths gets a reference to the given int64 and assigns it to the SixMonths field.
func (o *SeriesModelV1RankPosition) SetSixMonths(v int64) {
	o.SixMonths = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SeriesModelV1RankPosition) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1RankPosition) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SeriesModelV1RankPosition) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *SeriesModelV1RankPosition) SetYear(v int64) {
	o.Year = &v
}

func (o SeriesModelV1RankPosition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelV1RankPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Week) {
		toSerialize["week"] = o.Week
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.ThreeMonths) {
		toSerialize["three_months"] = o.ThreeMonths
	}
	if !IsNil(o.SixMonths) {
		toSerialize["six_months"] = o.SixMonths
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

type NullableSeriesModelV1RankPosition struct {
	value *SeriesModelV1RankPosition
	isSet bool
}

func (v NullableSeriesModelV1RankPosition) Get() *SeriesModelV1RankPosition {
	return v.value
}

func (v *NullableSeriesModelV1RankPosition) Set(val *SeriesModelV1RankPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelV1RankPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelV1RankPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelV1RankPosition(val *SeriesModelV1RankPosition) *NullableSeriesModelV1RankPosition {
	return &NullableSeriesModelV1RankPosition{value: val, isSet: true}
}

func (v NullableSeriesModelV1RankPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelV1RankPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
