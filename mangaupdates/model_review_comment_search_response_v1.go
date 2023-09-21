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

// checks if the ReviewCommentSearchResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewCommentSearchResponseV1{}

// ReviewCommentSearchResponseV1 struct for ReviewCommentSearchResponseV1
type ReviewCommentSearchResponseV1 struct {
	TotalHits *int64                                 `json:"total_hits,omitempty"`
	Page      *int64                                 `json:"page,omitempty"`
	PerPage   *int64                                 `json:"per_page,omitempty"`
	Results   []ReviewCommentSearchResponseV1Results `json:"results,omitempty"`
}

// NewReviewCommentSearchResponseV1 instantiates a new ReviewCommentSearchResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewCommentSearchResponseV1() *ReviewCommentSearchResponseV1 {
	this := ReviewCommentSearchResponseV1{}
	return &this
}

// NewReviewCommentSearchResponseV1WithDefaults instantiates a new ReviewCommentSearchResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewCommentSearchResponseV1WithDefaults() *ReviewCommentSearchResponseV1 {
	this := ReviewCommentSearchResponseV1{}
	return &this
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *ReviewCommentSearchResponseV1) GetTotalHits() int64 {
	if o == nil || IsNil(o.TotalHits) {
		var ret int64
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCommentSearchResponseV1) GetTotalHitsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalHits) {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *ReviewCommentSearchResponseV1) HasTotalHits() bool {
	if o != nil && !IsNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int64 and assigns it to the TotalHits field.
func (o *ReviewCommentSearchResponseV1) SetTotalHits(v int64) {
	o.TotalHits = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReviewCommentSearchResponseV1) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCommentSearchResponseV1) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReviewCommentSearchResponseV1) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *ReviewCommentSearchResponseV1) SetPage(v int64) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *ReviewCommentSearchResponseV1) GetPerPage() int64 {
	if o == nil || IsNil(o.PerPage) {
		var ret int64
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCommentSearchResponseV1) GetPerPageOk() (*int64, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *ReviewCommentSearchResponseV1) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int64 and assigns it to the PerPage field.
func (o *ReviewCommentSearchResponseV1) SetPerPage(v int64) {
	o.PerPage = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ReviewCommentSearchResponseV1) GetResults() []ReviewCommentSearchResponseV1Results {
	if o == nil || IsNil(o.Results) {
		var ret []ReviewCommentSearchResponseV1Results
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCommentSearchResponseV1) GetResultsOk() ([]ReviewCommentSearchResponseV1Results, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ReviewCommentSearchResponseV1) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ReviewCommentSearchResponseV1Results and assigns it to the Results field.
func (o *ReviewCommentSearchResponseV1) SetResults(v []ReviewCommentSearchResponseV1Results) {
	o.Results = v
}

func (o ReviewCommentSearchResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewCommentSearchResponseV1) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableReviewCommentSearchResponseV1 struct {
	value *ReviewCommentSearchResponseV1
	isSet bool
}

func (v NullableReviewCommentSearchResponseV1) Get() *ReviewCommentSearchResponseV1 {
	return v.value
}

func (v *NullableReviewCommentSearchResponseV1) Set(val *ReviewCommentSearchResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewCommentSearchResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewCommentSearchResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewCommentSearchResponseV1(val *ReviewCommentSearchResponseV1) *NullableReviewCommentSearchResponseV1 {
	return &NullableReviewCommentSearchResponseV1{value: val, isSet: true}
}

func (v NullableReviewCommentSearchResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewCommentSearchResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
