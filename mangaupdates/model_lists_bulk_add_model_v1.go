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

// checks if the ListsBulkAddModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsBulkAddModelV1{}

// ListsBulkAddModelV1 struct for ListsBulkAddModelV1
type ListsBulkAddModelV1 struct {
	Priority *string `json:"priority,omitempty"`
	SeriesTitle *string `json:"series_title,omitempty"`
}

// NewListsBulkAddModelV1 instantiates a new ListsBulkAddModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsBulkAddModelV1() *ListsBulkAddModelV1 {
	this := ListsBulkAddModelV1{}
	return &this
}

// NewListsBulkAddModelV1WithDefaults instantiates a new ListsBulkAddModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsBulkAddModelV1WithDefaults() *ListsBulkAddModelV1 {
	this := ListsBulkAddModelV1{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ListsBulkAddModelV1) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsBulkAddModelV1) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ListsBulkAddModelV1) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ListsBulkAddModelV1) SetPriority(v string) {
	o.Priority = &v
}

// GetSeriesTitle returns the SeriesTitle field value if set, zero value otherwise.
func (o *ListsBulkAddModelV1) GetSeriesTitle() string {
	if o == nil || IsNil(o.SeriesTitle) {
		var ret string
		return ret
	}
	return *o.SeriesTitle
}

// GetSeriesTitleOk returns a tuple with the SeriesTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsBulkAddModelV1) GetSeriesTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesTitle) {
		return nil, false
	}
	return o.SeriesTitle, true
}

// HasSeriesTitle returns a boolean if a field has been set.
func (o *ListsBulkAddModelV1) HasSeriesTitle() bool {
	if o != nil && !IsNil(o.SeriesTitle) {
		return true
	}

	return false
}

// SetSeriesTitle gets a reference to the given string and assigns it to the SeriesTitle field.
func (o *ListsBulkAddModelV1) SetSeriesTitle(v string) {
	o.SeriesTitle = &v
}

func (o ListsBulkAddModelV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsBulkAddModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.SeriesTitle) {
		toSerialize["series_title"] = o.SeriesTitle
	}
	return toSerialize, nil
}

type NullableListsBulkAddModelV1 struct {
	value *ListsBulkAddModelV1
	isSet bool
}

func (v NullableListsBulkAddModelV1) Get() *ListsBulkAddModelV1 {
	return v.value
}

func (v *NullableListsBulkAddModelV1) Set(val *ListsBulkAddModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableListsBulkAddModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableListsBulkAddModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsBulkAddModelV1(val *ListsBulkAddModelV1) *NullableListsBulkAddModelV1 {
	return &NullableListsBulkAddModelV1{value: val, isSet: true}
}

func (v NullableListsBulkAddModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsBulkAddModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

