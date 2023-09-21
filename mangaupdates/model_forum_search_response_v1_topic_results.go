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

// checks if the ForumSearchResponseV1TopicResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumSearchResponseV1TopicResults{}

// ForumSearchResponseV1TopicResults struct for ForumSearchResponseV1TopicResults
type ForumSearchResponseV1TopicResults struct {
	Record *ForumTopicModelSearchV1 `json:"record,omitempty"`
	Metadata *ForumSearchResponseV1TopicResultsMetadata `json:"metadata,omitempty"`
}

// NewForumSearchResponseV1TopicResults instantiates a new ForumSearchResponseV1TopicResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumSearchResponseV1TopicResults() *ForumSearchResponseV1TopicResults {
	this := ForumSearchResponseV1TopicResults{}
	return &this
}

// NewForumSearchResponseV1TopicResultsWithDefaults instantiates a new ForumSearchResponseV1TopicResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumSearchResponseV1TopicResultsWithDefaults() *ForumSearchResponseV1TopicResults {
	this := ForumSearchResponseV1TopicResults{}
	return &this
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *ForumSearchResponseV1TopicResults) GetRecord() ForumTopicModelSearchV1 {
	if o == nil || IsNil(o.Record) {
		var ret ForumTopicModelSearchV1
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumSearchResponseV1TopicResults) GetRecordOk() (*ForumTopicModelSearchV1, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *ForumSearchResponseV1TopicResults) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given ForumTopicModelSearchV1 and assigns it to the Record field.
func (o *ForumSearchResponseV1TopicResults) SetRecord(v ForumTopicModelSearchV1) {
	o.Record = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ForumSearchResponseV1TopicResults) GetMetadata() ForumSearchResponseV1TopicResultsMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ForumSearchResponseV1TopicResultsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumSearchResponseV1TopicResults) GetMetadataOk() (*ForumSearchResponseV1TopicResultsMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ForumSearchResponseV1TopicResults) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ForumSearchResponseV1TopicResultsMetadata and assigns it to the Metadata field.
func (o *ForumSearchResponseV1TopicResults) SetMetadata(v ForumSearchResponseV1TopicResultsMetadata) {
	o.Metadata = &v
}

func (o ForumSearchResponseV1TopicResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumSearchResponseV1TopicResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableForumSearchResponseV1TopicResults struct {
	value *ForumSearchResponseV1TopicResults
	isSet bool
}

func (v NullableForumSearchResponseV1TopicResults) Get() *ForumSearchResponseV1TopicResults {
	return v.value
}

func (v *NullableForumSearchResponseV1TopicResults) Set(val *ForumSearchResponseV1TopicResults) {
	v.value = val
	v.isSet = true
}

func (v NullableForumSearchResponseV1TopicResults) IsSet() bool {
	return v.isSet
}

func (v *NullableForumSearchResponseV1TopicResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumSearchResponseV1TopicResults(val *ForumSearchResponseV1TopicResults) *NullableForumSearchResponseV1TopicResults {
	return &NullableForumSearchResponseV1TopicResults{value: val, isSet: true}
}

func (v NullableForumSearchResponseV1TopicResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumSearchResponseV1TopicResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


