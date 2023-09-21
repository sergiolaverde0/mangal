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

// checks if the AuthorsSeriesListResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorsSeriesListResponseV1{}

// AuthorsSeriesListResponseV1 struct for AuthorsSeriesListResponseV1
type AuthorsSeriesListResponseV1 struct {
	TotalSeries *int64                                  `json:"total_series,omitempty"`
	SeriesList  []AuthorsSeriesListResponseV1SeriesList `json:"series_list,omitempty"`
	GenreList   []AuthorsSeriesListResponseV1GenreList  `json:"genre_list,omitempty"`
}

// NewAuthorsSeriesListResponseV1 instantiates a new AuthorsSeriesListResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorsSeriesListResponseV1() *AuthorsSeriesListResponseV1 {
	this := AuthorsSeriesListResponseV1{}
	return &this
}

// NewAuthorsSeriesListResponseV1WithDefaults instantiates a new AuthorsSeriesListResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorsSeriesListResponseV1WithDefaults() *AuthorsSeriesListResponseV1 {
	this := AuthorsSeriesListResponseV1{}
	return &this
}

// GetTotalSeries returns the TotalSeries field value if set, zero value otherwise.
func (o *AuthorsSeriesListResponseV1) GetTotalSeries() int64 {
	if o == nil || IsNil(o.TotalSeries) {
		var ret int64
		return ret
	}
	return *o.TotalSeries
}

// GetTotalSeriesOk returns a tuple with the TotalSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSeriesListResponseV1) GetTotalSeriesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSeries) {
		return nil, false
	}
	return o.TotalSeries, true
}

// HasTotalSeries returns a boolean if a field has been set.
func (o *AuthorsSeriesListResponseV1) HasTotalSeries() bool {
	if o != nil && !IsNil(o.TotalSeries) {
		return true
	}

	return false
}

// SetTotalSeries gets a reference to the given int64 and assigns it to the TotalSeries field.
func (o *AuthorsSeriesListResponseV1) SetTotalSeries(v int64) {
	o.TotalSeries = &v
}

// GetSeriesList returns the SeriesList field value if set, zero value otherwise.
func (o *AuthorsSeriesListResponseV1) GetSeriesList() []AuthorsSeriesListResponseV1SeriesList {
	if o == nil || IsNil(o.SeriesList) {
		var ret []AuthorsSeriesListResponseV1SeriesList
		return ret
	}
	return o.SeriesList
}

// GetSeriesListOk returns a tuple with the SeriesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSeriesListResponseV1) GetSeriesListOk() ([]AuthorsSeriesListResponseV1SeriesList, bool) {
	if o == nil || IsNil(o.SeriesList) {
		return nil, false
	}
	return o.SeriesList, true
}

// HasSeriesList returns a boolean if a field has been set.
func (o *AuthorsSeriesListResponseV1) HasSeriesList() bool {
	if o != nil && !IsNil(o.SeriesList) {
		return true
	}

	return false
}

// SetSeriesList gets a reference to the given []AuthorsSeriesListResponseV1SeriesList and assigns it to the SeriesList field.
func (o *AuthorsSeriesListResponseV1) SetSeriesList(v []AuthorsSeriesListResponseV1SeriesList) {
	o.SeriesList = v
}

// GetGenreList returns the GenreList field value if set, zero value otherwise.
func (o *AuthorsSeriesListResponseV1) GetGenreList() []AuthorsSeriesListResponseV1GenreList {
	if o == nil || IsNil(o.GenreList) {
		var ret []AuthorsSeriesListResponseV1GenreList
		return ret
	}
	return o.GenreList
}

// GetGenreListOk returns a tuple with the GenreList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSeriesListResponseV1) GetGenreListOk() ([]AuthorsSeriesListResponseV1GenreList, bool) {
	if o == nil || IsNil(o.GenreList) {
		return nil, false
	}
	return o.GenreList, true
}

// HasGenreList returns a boolean if a field has been set.
func (o *AuthorsSeriesListResponseV1) HasGenreList() bool {
	if o != nil && !IsNil(o.GenreList) {
		return true
	}

	return false
}

// SetGenreList gets a reference to the given []AuthorsSeriesListResponseV1GenreList and assigns it to the GenreList field.
func (o *AuthorsSeriesListResponseV1) SetGenreList(v []AuthorsSeriesListResponseV1GenreList) {
	o.GenreList = v
}

func (o AuthorsSeriesListResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorsSeriesListResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalSeries) {
		toSerialize["total_series"] = o.TotalSeries
	}
	if !IsNil(o.SeriesList) {
		toSerialize["series_list"] = o.SeriesList
	}
	if !IsNil(o.GenreList) {
		toSerialize["genre_list"] = o.GenreList
	}
	return toSerialize, nil
}

type NullableAuthorsSeriesListResponseV1 struct {
	value *AuthorsSeriesListResponseV1
	isSet bool
}

func (v NullableAuthorsSeriesListResponseV1) Get() *AuthorsSeriesListResponseV1 {
	return v.value
}

func (v *NullableAuthorsSeriesListResponseV1) Set(val *AuthorsSeriesListResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorsSeriesListResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorsSeriesListResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorsSeriesListResponseV1(val *AuthorsSeriesListResponseV1) *NullableAuthorsSeriesListResponseV1 {
	return &NullableAuthorsSeriesListResponseV1{value: val, isSet: true}
}

func (v NullableAuthorsSeriesListResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorsSeriesListResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}