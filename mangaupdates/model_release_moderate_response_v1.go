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

// checks if the ReleaseModerateResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseModerateResponseV1{}

// ReleaseModerateResponseV1 struct for ReleaseModerateResponseV1
type ReleaseModerateResponseV1 struct {
	TotalHits *int64                               `json:"total_hits,omitempty"`
	Page      *int64                               `json:"page,omitempty"`
	PerPage   *int64                               `json:"per_page,omitempty"`
	Results   []ReleaseModerateResponseV1Results   `json:"results,omitempty"`
	GroupInfo []ReleaseModerateResponseV1GroupInfo `json:"group_info,omitempty"`
}

// NewReleaseModerateResponseV1 instantiates a new ReleaseModerateResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseModerateResponseV1() *ReleaseModerateResponseV1 {
	this := ReleaseModerateResponseV1{}
	return &this
}

// NewReleaseModerateResponseV1WithDefaults instantiates a new ReleaseModerateResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseModerateResponseV1WithDefaults() *ReleaseModerateResponseV1 {
	this := ReleaseModerateResponseV1{}
	return &this
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *ReleaseModerateResponseV1) GetTotalHits() int64 {
	if o == nil || IsNil(o.TotalHits) {
		var ret int64
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModerateResponseV1) GetTotalHitsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalHits) {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *ReleaseModerateResponseV1) HasTotalHits() bool {
	if o != nil && !IsNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int64 and assigns it to the TotalHits field.
func (o *ReleaseModerateResponseV1) SetTotalHits(v int64) {
	o.TotalHits = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReleaseModerateResponseV1) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModerateResponseV1) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReleaseModerateResponseV1) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *ReleaseModerateResponseV1) SetPage(v int64) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *ReleaseModerateResponseV1) GetPerPage() int64 {
	if o == nil || IsNil(o.PerPage) {
		var ret int64
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModerateResponseV1) GetPerPageOk() (*int64, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *ReleaseModerateResponseV1) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int64 and assigns it to the PerPage field.
func (o *ReleaseModerateResponseV1) SetPerPage(v int64) {
	o.PerPage = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ReleaseModerateResponseV1) GetResults() []ReleaseModerateResponseV1Results {
	if o == nil || IsNil(o.Results) {
		var ret []ReleaseModerateResponseV1Results
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModerateResponseV1) GetResultsOk() ([]ReleaseModerateResponseV1Results, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ReleaseModerateResponseV1) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ReleaseModerateResponseV1Results and assigns it to the Results field.
func (o *ReleaseModerateResponseV1) SetResults(v []ReleaseModerateResponseV1Results) {
	o.Results = v
}

// GetGroupInfo returns the GroupInfo field value if set, zero value otherwise.
func (o *ReleaseModerateResponseV1) GetGroupInfo() []ReleaseModerateResponseV1GroupInfo {
	if o == nil || IsNil(o.GroupInfo) {
		var ret []ReleaseModerateResponseV1GroupInfo
		return ret
	}
	return o.GroupInfo
}

// GetGroupInfoOk returns a tuple with the GroupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModerateResponseV1) GetGroupInfoOk() ([]ReleaseModerateResponseV1GroupInfo, bool) {
	if o == nil || IsNil(o.GroupInfo) {
		return nil, false
	}
	return o.GroupInfo, true
}

// HasGroupInfo returns a boolean if a field has been set.
func (o *ReleaseModerateResponseV1) HasGroupInfo() bool {
	if o != nil && !IsNil(o.GroupInfo) {
		return true
	}

	return false
}

// SetGroupInfo gets a reference to the given []ReleaseModerateResponseV1GroupInfo and assigns it to the GroupInfo field.
func (o *ReleaseModerateResponseV1) SetGroupInfo(v []ReleaseModerateResponseV1GroupInfo) {
	o.GroupInfo = v
}

func (o ReleaseModerateResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseModerateResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalHits) {
		toSerialize["total_hits"] = o.TotalHits
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.GroupInfo) {
		toSerialize["group_info"] = o.GroupInfo
	}
	return toSerialize, nil
}

type NullableReleaseModerateResponseV1 struct {
	value *ReleaseModerateResponseV1
	isSet bool
}

func (v NullableReleaseModerateResponseV1) Get() *ReleaseModerateResponseV1 {
	return v.value
}

func (v *NullableReleaseModerateResponseV1) Set(val *ReleaseModerateResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseModerateResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseModerateResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseModerateResponseV1(val *ReleaseModerateResponseV1) *NullableReleaseModerateResponseV1 {
	return &NullableReleaseModerateResponseV1{value: val, isSet: true}
}

func (v NullableReleaseModerateResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseModerateResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
