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

// checks if the SeriesModelUpdateV1RelatedSeries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelUpdateV1RelatedSeries{}

// SeriesModelUpdateV1RelatedSeries struct for SeriesModelUpdateV1RelatedSeries
type SeriesModelUpdateV1RelatedSeries struct {
	RelationType    string `json:"relation_type"`
	RelatedSeriesId int64  `json:"related_series_id"`
}

// NewSeriesModelUpdateV1RelatedSeries instantiates a new SeriesModelUpdateV1RelatedSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelUpdateV1RelatedSeries(relationType string, relatedSeriesId int64) *SeriesModelUpdateV1RelatedSeries {
	this := SeriesModelUpdateV1RelatedSeries{}
	this.RelationType = relationType
	this.RelatedSeriesId = relatedSeriesId
	return &this
}

// NewSeriesModelUpdateV1RelatedSeriesWithDefaults instantiates a new SeriesModelUpdateV1RelatedSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelUpdateV1RelatedSeriesWithDefaults() *SeriesModelUpdateV1RelatedSeries {
	this := SeriesModelUpdateV1RelatedSeries{}
	return &this
}

// GetRelationType returns the RelationType field value
func (o *SeriesModelUpdateV1RelatedSeries) GetRelationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelationType
}

// GetRelationTypeOk returns a tuple with the RelationType field value
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1RelatedSeries) GetRelationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationType, true
}

// SetRelationType sets field value
func (o *SeriesModelUpdateV1RelatedSeries) SetRelationType(v string) {
	o.RelationType = v
}

// GetRelatedSeriesId returns the RelatedSeriesId field value
func (o *SeriesModelUpdateV1RelatedSeries) GetRelatedSeriesId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RelatedSeriesId
}

// GetRelatedSeriesIdOk returns a tuple with the RelatedSeriesId field value
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1RelatedSeries) GetRelatedSeriesIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedSeriesId, true
}

// SetRelatedSeriesId sets field value
func (o *SeriesModelUpdateV1RelatedSeries) SetRelatedSeriesId(v int64) {
	o.RelatedSeriesId = v
}

func (o SeriesModelUpdateV1RelatedSeries) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelUpdateV1RelatedSeries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relation_type"] = o.RelationType
	toSerialize["related_series_id"] = o.RelatedSeriesId
	return toSerialize, nil
}

type NullableSeriesModelUpdateV1RelatedSeries struct {
	value *SeriesModelUpdateV1RelatedSeries
	isSet bool
}

func (v NullableSeriesModelUpdateV1RelatedSeries) Get() *SeriesModelUpdateV1RelatedSeries {
	return v.value
}

func (v *NullableSeriesModelUpdateV1RelatedSeries) Set(val *SeriesModelUpdateV1RelatedSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelUpdateV1RelatedSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelUpdateV1RelatedSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelUpdateV1RelatedSeries(val *SeriesModelUpdateV1RelatedSeries) *NullableSeriesModelUpdateV1RelatedSeries {
	return &NullableSeriesModelUpdateV1RelatedSeries{value: val, isSet: true}
}

func (v NullableSeriesModelUpdateV1RelatedSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelUpdateV1RelatedSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
