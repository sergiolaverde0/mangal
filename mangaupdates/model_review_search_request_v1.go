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

// checks if the ReviewSearchRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSearchRequestV1{}

// ReviewSearchRequestV1 struct for ReviewSearchRequestV1
type ReviewSearchRequestV1 struct {
	Search   *string `json:"search,omitempty"`
	AddedBy  *int64  `json:"added_by,omitempty"`
	SeriesId *int64  `json:"series_id,omitempty"`
	Page     *int64  `json:"page,omitempty"`
	Perpage  *int64  `json:"perpage,omitempty"`
	Letter   *string `json:"letter,omitempty"`
	Pending  *bool   `json:"pending,omitempty"`
}

// NewReviewSearchRequestV1 instantiates a new ReviewSearchRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSearchRequestV1() *ReviewSearchRequestV1 {
	this := ReviewSearchRequestV1{}
	return &this
}

// NewReviewSearchRequestV1WithDefaults instantiates a new ReviewSearchRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSearchRequestV1WithDefaults() *ReviewSearchRequestV1 {
	this := ReviewSearchRequestV1{}
	return &this
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetSearch() string {
	if o == nil || IsNil(o.Search) {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetSearchOk() (*string, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *ReviewSearchRequestV1) SetSearch(v string) {
	o.Search = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetAddedBy() int64 {
	if o == nil || IsNil(o.AddedBy) {
		var ret int64
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetAddedByOk() (*int64, bool) {
	if o == nil || IsNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasAddedBy() bool {
	if o != nil && !IsNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given int64 and assigns it to the AddedBy field.
func (o *ReviewSearchRequestV1) SetAddedBy(v int64) {
	o.AddedBy = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetSeriesId() int64 {
	if o == nil || IsNil(o.SeriesId) {
		var ret int64
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetSeriesIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int64 and assigns it to the SeriesId field.
func (o *ReviewSearchRequestV1) SetSeriesId(v int64) {
	o.SeriesId = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *ReviewSearchRequestV1) SetPage(v int64) {
	o.Page = &v
}

// GetPerpage returns the Perpage field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetPerpage() int64 {
	if o == nil || IsNil(o.Perpage) {
		var ret int64
		return ret
	}
	return *o.Perpage
}

// GetPerpageOk returns a tuple with the Perpage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetPerpageOk() (*int64, bool) {
	if o == nil || IsNil(o.Perpage) {
		return nil, false
	}
	return o.Perpage, true
}

// HasPerpage returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasPerpage() bool {
	if o != nil && !IsNil(o.Perpage) {
		return true
	}

	return false
}

// SetPerpage gets a reference to the given int64 and assigns it to the Perpage field.
func (o *ReviewSearchRequestV1) SetPerpage(v int64) {
	o.Perpage = &v
}

// GetLetter returns the Letter field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetLetter() string {
	if o == nil || IsNil(o.Letter) {
		var ret string
		return ret
	}
	return *o.Letter
}

// GetLetterOk returns a tuple with the Letter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetLetterOk() (*string, bool) {
	if o == nil || IsNil(o.Letter) {
		return nil, false
	}
	return o.Letter, true
}

// HasLetter returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasLetter() bool {
	if o != nil && !IsNil(o.Letter) {
		return true
	}

	return false
}

// SetLetter gets a reference to the given string and assigns it to the Letter field.
func (o *ReviewSearchRequestV1) SetLetter(v string) {
	o.Letter = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *ReviewSearchRequestV1) GetPending() bool {
	if o == nil || IsNil(o.Pending) {
		var ret bool
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSearchRequestV1) GetPendingOk() (*bool, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *ReviewSearchRequestV1) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given bool and assigns it to the Pending field.
func (o *ReviewSearchRequestV1) SetPending(v bool) {
	o.Pending = &v
}

func (o ReviewSearchRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSearchRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.AddedBy) {
		toSerialize["added_by"] = o.AddedBy
	}
	if !IsNil(o.SeriesId) {
		toSerialize["series_id"] = o.SeriesId
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Perpage) {
		toSerialize["perpage"] = o.Perpage
	}
	if !IsNil(o.Letter) {
		toSerialize["letter"] = o.Letter
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	return toSerialize, nil
}

type NullableReviewSearchRequestV1 struct {
	value *ReviewSearchRequestV1
	isSet bool
}

func (v NullableReviewSearchRequestV1) Get() *ReviewSearchRequestV1 {
	return v.value
}

func (v *NullableReviewSearchRequestV1) Set(val *ReviewSearchRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSearchRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSearchRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSearchRequestV1(val *ReviewSearchRequestV1) *NullableReviewSearchRequestV1 {
	return &NullableReviewSearchRequestV1{value: val, isSet: true}
}

func (v NullableReviewSearchRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSearchRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
