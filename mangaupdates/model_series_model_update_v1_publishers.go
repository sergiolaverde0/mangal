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

// checks if the SeriesModelUpdateV1Publishers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelUpdateV1Publishers{}

// SeriesModelUpdateV1Publishers struct for SeriesModelUpdateV1Publishers
type SeriesModelUpdateV1Publishers struct {
	PublisherName *string `json:"publisher_name,omitempty"`
	Type *string `json:"type,omitempty"`
	Notes *string `json:"notes,omitempty"`
}

// NewSeriesModelUpdateV1Publishers instantiates a new SeriesModelUpdateV1Publishers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelUpdateV1Publishers() *SeriesModelUpdateV1Publishers {
	this := SeriesModelUpdateV1Publishers{}
	return &this
}

// NewSeriesModelUpdateV1PublishersWithDefaults instantiates a new SeriesModelUpdateV1Publishers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelUpdateV1PublishersWithDefaults() *SeriesModelUpdateV1Publishers {
	this := SeriesModelUpdateV1Publishers{}
	return &this
}

// GetPublisherName returns the PublisherName field value if set, zero value otherwise.
func (o *SeriesModelUpdateV1Publishers) GetPublisherName() string {
	if o == nil || IsNil(o.PublisherName) {
		var ret string
		return ret
	}
	return *o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1Publishers) GetPublisherNameOk() (*string, bool) {
	if o == nil || IsNil(o.PublisherName) {
		return nil, false
	}
	return o.PublisherName, true
}

// HasPublisherName returns a boolean if a field has been set.
func (o *SeriesModelUpdateV1Publishers) HasPublisherName() bool {
	if o != nil && !IsNil(o.PublisherName) {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given string and assigns it to the PublisherName field.
func (o *SeriesModelUpdateV1Publishers) SetPublisherName(v string) {
	o.PublisherName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SeriesModelUpdateV1Publishers) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1Publishers) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SeriesModelUpdateV1Publishers) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SeriesModelUpdateV1Publishers) SetType(v string) {
	o.Type = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SeriesModelUpdateV1Publishers) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1Publishers) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SeriesModelUpdateV1Publishers) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SeriesModelUpdateV1Publishers) SetNotes(v string) {
	o.Notes = &v
}

func (o SeriesModelUpdateV1Publishers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelUpdateV1Publishers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublisherName) {
		toSerialize["publisher_name"] = o.PublisherName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableSeriesModelUpdateV1Publishers struct {
	value *SeriesModelUpdateV1Publishers
	isSet bool
}

func (v NullableSeriesModelUpdateV1Publishers) Get() *SeriesModelUpdateV1Publishers {
	return v.value
}

func (v *NullableSeriesModelUpdateV1Publishers) Set(val *SeriesModelUpdateV1Publishers) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelUpdateV1Publishers) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelUpdateV1Publishers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelUpdateV1Publishers(val *SeriesModelUpdateV1Publishers) *NullableSeriesModelUpdateV1Publishers {
	return &NullableSeriesModelUpdateV1Publishers{value: val, isSet: true}
}

func (v NullableSeriesModelUpdateV1Publishers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelUpdateV1Publishers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


