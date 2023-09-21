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

// checks if the SeriesModelV1Publications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelV1Publications{}

// SeriesModelV1Publications struct for SeriesModelV1Publications
type SeriesModelV1Publications struct {
	PublicationName *string `json:"publication_name,omitempty"`
	PublisherName *string `json:"publisher_name,omitempty"`
	PublisherId *string `json:"publisher_id,omitempty"`
}

// NewSeriesModelV1Publications instantiates a new SeriesModelV1Publications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelV1Publications() *SeriesModelV1Publications {
	this := SeriesModelV1Publications{}
	return &this
}

// NewSeriesModelV1PublicationsWithDefaults instantiates a new SeriesModelV1Publications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelV1PublicationsWithDefaults() *SeriesModelV1Publications {
	this := SeriesModelV1Publications{}
	return &this
}

// GetPublicationName returns the PublicationName field value if set, zero value otherwise.
func (o *SeriesModelV1Publications) GetPublicationName() string {
	if o == nil || IsNil(o.PublicationName) {
		var ret string
		return ret
	}
	return *o.PublicationName
}

// GetPublicationNameOk returns a tuple with the PublicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Publications) GetPublicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationName) {
		return nil, false
	}
	return o.PublicationName, true
}

// HasPublicationName returns a boolean if a field has been set.
func (o *SeriesModelV1Publications) HasPublicationName() bool {
	if o != nil && !IsNil(o.PublicationName) {
		return true
	}

	return false
}

// SetPublicationName gets a reference to the given string and assigns it to the PublicationName field.
func (o *SeriesModelV1Publications) SetPublicationName(v string) {
	o.PublicationName = &v
}

// GetPublisherName returns the PublisherName field value if set, zero value otherwise.
func (o *SeriesModelV1Publications) GetPublisherName() string {
	if o == nil || IsNil(o.PublisherName) {
		var ret string
		return ret
	}
	return *o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Publications) GetPublisherNameOk() (*string, bool) {
	if o == nil || IsNil(o.PublisherName) {
		return nil, false
	}
	return o.PublisherName, true
}

// HasPublisherName returns a boolean if a field has been set.
func (o *SeriesModelV1Publications) HasPublisherName() bool {
	if o != nil && !IsNil(o.PublisherName) {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given string and assigns it to the PublisherName field.
func (o *SeriesModelV1Publications) SetPublisherName(v string) {
	o.PublisherName = &v
}

// GetPublisherId returns the PublisherId field value if set, zero value otherwise.
func (o *SeriesModelV1Publications) GetPublisherId() string {
	if o == nil || IsNil(o.PublisherId) {
		var ret string
		return ret
	}
	return *o.PublisherId
}

// GetPublisherIdOk returns a tuple with the PublisherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Publications) GetPublisherIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublisherId) {
		return nil, false
	}
	return o.PublisherId, true
}

// HasPublisherId returns a boolean if a field has been set.
func (o *SeriesModelV1Publications) HasPublisherId() bool {
	if o != nil && !IsNil(o.PublisherId) {
		return true
	}

	return false
}

// SetPublisherId gets a reference to the given string and assigns it to the PublisherId field.
func (o *SeriesModelV1Publications) SetPublisherId(v string) {
	o.PublisherId = &v
}

func (o SeriesModelV1Publications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelV1Publications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicationName) {
		toSerialize["publication_name"] = o.PublicationName
	}
	if !IsNil(o.PublisherName) {
		toSerialize["publisher_name"] = o.PublisherName
	}
	if !IsNil(o.PublisherId) {
		toSerialize["publisher_id"] = o.PublisherId
	}
	return toSerialize, nil
}

type NullableSeriesModelV1Publications struct {
	value *SeriesModelV1Publications
	isSet bool
}

func (v NullableSeriesModelV1Publications) Get() *SeriesModelV1Publications {
	return v.value
}

func (v *NullableSeriesModelV1Publications) Set(val *SeriesModelV1Publications) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelV1Publications) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelV1Publications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelV1Publications(val *SeriesModelV1Publications) *NullableSeriesModelV1Publications {
	return &NullableSeriesModelV1Publications{value: val, isSet: true}
}

func (v NullableSeriesModelV1Publications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelV1Publications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

