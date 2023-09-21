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

// checks if the AuthorsModelSearchV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorsModelSearchV1{}

// AuthorsModelSearchV1 struct for AuthorsModelSearchV1
type AuthorsModelSearchV1 struct {
	Id      *int64                     `json:"id,omitempty"`
	Name    *string                    `json:"name,omitempty"`
	Url     *string                    `json:"url,omitempty"`
	Genres  []string                   `json:"genres,omitempty"`
	Stats   *AuthorsModelSearchV1Stats `json:"stats,omitempty"`
	AddedBy *UserModelSearchV1         `json:"added_by,omitempty"`
}

// NewAuthorsModelSearchV1 instantiates a new AuthorsModelSearchV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorsModelSearchV1() *AuthorsModelSearchV1 {
	this := AuthorsModelSearchV1{}
	return &this
}

// NewAuthorsModelSearchV1WithDefaults instantiates a new AuthorsModelSearchV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorsModelSearchV1WithDefaults() *AuthorsModelSearchV1 {
	this := AuthorsModelSearchV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AuthorsModelSearchV1) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthorsModelSearchV1) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AuthorsModelSearchV1) SetUrl(v string) {
	o.Url = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *AuthorsModelSearchV1) SetGenres(v []string) {
	o.Genres = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetStats() AuthorsModelSearchV1Stats {
	if o == nil || IsNil(o.Stats) {
		var ret AuthorsModelSearchV1Stats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetStatsOk() (*AuthorsModelSearchV1Stats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given AuthorsModelSearchV1Stats and assigns it to the Stats field.
func (o *AuthorsModelSearchV1) SetStats(v AuthorsModelSearchV1Stats) {
	o.Stats = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *AuthorsModelSearchV1) GetAddedBy() UserModelSearchV1 {
	if o == nil || IsNil(o.AddedBy) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsModelSearchV1) GetAddedByOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *AuthorsModelSearchV1) HasAddedBy() bool {
	if o != nil && !IsNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given UserModelSearchV1 and assigns it to the AddedBy field.
func (o *AuthorsModelSearchV1) SetAddedBy(v UserModelSearchV1) {
	o.AddedBy = &v
}

func (o AuthorsModelSearchV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorsModelSearchV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.AddedBy) {
		toSerialize["added_by"] = o.AddedBy
	}
	return toSerialize, nil
}

type NullableAuthorsModelSearchV1 struct {
	value *AuthorsModelSearchV1
	isSet bool
}

func (v NullableAuthorsModelSearchV1) Get() *AuthorsModelSearchV1 {
	return v.value
}

func (v *NullableAuthorsModelSearchV1) Set(val *AuthorsModelSearchV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorsModelSearchV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorsModelSearchV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorsModelSearchV1(val *AuthorsModelSearchV1) *NullableAuthorsModelSearchV1 {
	return &NullableAuthorsModelSearchV1{value: val, isSet: true}
}

func (v NullableAuthorsModelSearchV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorsModelSearchV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
