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

// checks if the GroupsSeriesListResponseV1SeriesCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsSeriesListResponseV1SeriesCategories{}

// GroupsSeriesListResponseV1SeriesCategories struct for GroupsSeriesListResponseV1SeriesCategories
type GroupsSeriesListResponseV1SeriesCategories struct {
	Category *string `json:"category,omitempty"`
	Votes    *int64  `json:"votes,omitempty"`
}

// NewGroupsSeriesListResponseV1SeriesCategories instantiates a new GroupsSeriesListResponseV1SeriesCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsSeriesListResponseV1SeriesCategories() *GroupsSeriesListResponseV1SeriesCategories {
	this := GroupsSeriesListResponseV1SeriesCategories{}
	return &this
}

// NewGroupsSeriesListResponseV1SeriesCategoriesWithDefaults instantiates a new GroupsSeriesListResponseV1SeriesCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsSeriesListResponseV1SeriesCategoriesWithDefaults() *GroupsSeriesListResponseV1SeriesCategories {
	this := GroupsSeriesListResponseV1SeriesCategories{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GroupsSeriesListResponseV1SeriesCategories) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsSeriesListResponseV1SeriesCategories) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GroupsSeriesListResponseV1SeriesCategories) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GroupsSeriesListResponseV1SeriesCategories) SetCategory(v string) {
	o.Category = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *GroupsSeriesListResponseV1SeriesCategories) GetVotes() int64 {
	if o == nil || IsNil(o.Votes) {
		var ret int64
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsSeriesListResponseV1SeriesCategories) GetVotesOk() (*int64, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *GroupsSeriesListResponseV1SeriesCategories) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given int64 and assigns it to the Votes field.
func (o *GroupsSeriesListResponseV1SeriesCategories) SetVotes(v int64) {
	o.Votes = &v
}

func (o GroupsSeriesListResponseV1SeriesCategories) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsSeriesListResponseV1SeriesCategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	return toSerialize, nil
}

type NullableGroupsSeriesListResponseV1SeriesCategories struct {
	value *GroupsSeriesListResponseV1SeriesCategories
	isSet bool
}

func (v NullableGroupsSeriesListResponseV1SeriesCategories) Get() *GroupsSeriesListResponseV1SeriesCategories {
	return v.value
}

func (v *NullableGroupsSeriesListResponseV1SeriesCategories) Set(val *GroupsSeriesListResponseV1SeriesCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsSeriesListResponseV1SeriesCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsSeriesListResponseV1SeriesCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsSeriesListResponseV1SeriesCategories(val *GroupsSeriesListResponseV1SeriesCategories) *NullableGroupsSeriesListResponseV1SeriesCategories {
	return &NullableGroupsSeriesListResponseV1SeriesCategories{value: val, isSet: true}
}

func (v NullableGroupsSeriesListResponseV1SeriesCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsSeriesListResponseV1SeriesCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
