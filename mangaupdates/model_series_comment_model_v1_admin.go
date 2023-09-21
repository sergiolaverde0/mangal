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

// checks if the SeriesCommentModelV1Admin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesCommentModelV1Admin{}

// SeriesCommentModelV1Admin struct for SeriesCommentModelV1Admin
type SeriesCommentModelV1Admin struct {
	Moderated *bool `json:"moderated,omitempty"`
	Reported *bool `json:"reported,omitempty"`
	ReportReason *string `json:"report_reason,omitempty"`
}

// NewSeriesCommentModelV1Admin instantiates a new SeriesCommentModelV1Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesCommentModelV1Admin() *SeriesCommentModelV1Admin {
	this := SeriesCommentModelV1Admin{}
	return &this
}

// NewSeriesCommentModelV1AdminWithDefaults instantiates a new SeriesCommentModelV1Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesCommentModelV1AdminWithDefaults() *SeriesCommentModelV1Admin {
	this := SeriesCommentModelV1Admin{}
	return &this
}

// GetModerated returns the Moderated field value if set, zero value otherwise.
func (o *SeriesCommentModelV1Admin) GetModerated() bool {
	if o == nil || IsNil(o.Moderated) {
		var ret bool
		return ret
	}
	return *o.Moderated
}

// GetModeratedOk returns a tuple with the Moderated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelV1Admin) GetModeratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Moderated) {
		return nil, false
	}
	return o.Moderated, true
}

// HasModerated returns a boolean if a field has been set.
func (o *SeriesCommentModelV1Admin) HasModerated() bool {
	if o != nil && !IsNil(o.Moderated) {
		return true
	}

	return false
}

// SetModerated gets a reference to the given bool and assigns it to the Moderated field.
func (o *SeriesCommentModelV1Admin) SetModerated(v bool) {
	o.Moderated = &v
}

// GetReported returns the Reported field value if set, zero value otherwise.
func (o *SeriesCommentModelV1Admin) GetReported() bool {
	if o == nil || IsNil(o.Reported) {
		var ret bool
		return ret
	}
	return *o.Reported
}

// GetReportedOk returns a tuple with the Reported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelV1Admin) GetReportedOk() (*bool, bool) {
	if o == nil || IsNil(o.Reported) {
		return nil, false
	}
	return o.Reported, true
}

// HasReported returns a boolean if a field has been set.
func (o *SeriesCommentModelV1Admin) HasReported() bool {
	if o != nil && !IsNil(o.Reported) {
		return true
	}

	return false
}

// SetReported gets a reference to the given bool and assigns it to the Reported field.
func (o *SeriesCommentModelV1Admin) SetReported(v bool) {
	o.Reported = &v
}

// GetReportReason returns the ReportReason field value if set, zero value otherwise.
func (o *SeriesCommentModelV1Admin) GetReportReason() string {
	if o == nil || IsNil(o.ReportReason) {
		var ret string
		return ret
	}
	return *o.ReportReason
}

// GetReportReasonOk returns a tuple with the ReportReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelV1Admin) GetReportReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReportReason) {
		return nil, false
	}
	return o.ReportReason, true
}

// HasReportReason returns a boolean if a field has been set.
func (o *SeriesCommentModelV1Admin) HasReportReason() bool {
	if o != nil && !IsNil(o.ReportReason) {
		return true
	}

	return false
}

// SetReportReason gets a reference to the given string and assigns it to the ReportReason field.
func (o *SeriesCommentModelV1Admin) SetReportReason(v string) {
	o.ReportReason = &v
}

func (o SeriesCommentModelV1Admin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesCommentModelV1Admin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Moderated) {
		toSerialize["moderated"] = o.Moderated
	}
	if !IsNil(o.Reported) {
		toSerialize["reported"] = o.Reported
	}
	if !IsNil(o.ReportReason) {
		toSerialize["report_reason"] = o.ReportReason
	}
	return toSerialize, nil
}

type NullableSeriesCommentModelV1Admin struct {
	value *SeriesCommentModelV1Admin
	isSet bool
}

func (v NullableSeriesCommentModelV1Admin) Get() *SeriesCommentModelV1Admin {
	return v.value
}

func (v *NullableSeriesCommentModelV1Admin) Set(val *SeriesCommentModelV1Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesCommentModelV1Admin) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesCommentModelV1Admin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesCommentModelV1Admin(val *SeriesCommentModelV1Admin) *NullableSeriesCommentModelV1Admin {
	return &NullableSeriesCommentModelV1Admin{value: val, isSet: true}
}

func (v NullableSeriesCommentModelV1Admin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesCommentModelV1Admin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


