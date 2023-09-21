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

// checks if the UserModelV1StatsModeration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModelV1StatsModeration{}

// UserModelV1StatsModeration struct for UserModelV1StatsModeration
type UserModelV1StatsModeration struct {
	Releases *UserModelV1StatsModerationReleases `json:"releases,omitempty"`
	Series *UserModelV1StatsModerationSeries `json:"series,omitempty"`
	Publishers *UserModelV1StatsModerationPublishers `json:"publishers,omitempty"`
	Groups *UserModelV1StatsModerationGroups `json:"groups,omitempty"`
	Authors *UserModelV1StatsModerationAuthors `json:"authors,omitempty"`
	LastAction *TimeV1 `json:"last_action,omitempty"`
}

// NewUserModelV1StatsModeration instantiates a new UserModelV1StatsModeration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModelV1StatsModeration() *UserModelV1StatsModeration {
	this := UserModelV1StatsModeration{}
	return &this
}

// NewUserModelV1StatsModerationWithDefaults instantiates a new UserModelV1StatsModeration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelV1StatsModerationWithDefaults() *UserModelV1StatsModeration {
	this := UserModelV1StatsModeration{}
	return &this
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetReleases() UserModelV1StatsModerationReleases {
	if o == nil || IsNil(o.Releases) {
		var ret UserModelV1StatsModerationReleases
		return ret
	}
	return *o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetReleasesOk() (*UserModelV1StatsModerationReleases, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given UserModelV1StatsModerationReleases and assigns it to the Releases field.
func (o *UserModelV1StatsModeration) SetReleases(v UserModelV1StatsModerationReleases) {
	o.Releases = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetSeries() UserModelV1StatsModerationSeries {
	if o == nil || IsNil(o.Series) {
		var ret UserModelV1StatsModerationSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetSeriesOk() (*UserModelV1StatsModerationSeries, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given UserModelV1StatsModerationSeries and assigns it to the Series field.
func (o *UserModelV1StatsModeration) SetSeries(v UserModelV1StatsModerationSeries) {
	o.Series = &v
}

// GetPublishers returns the Publishers field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetPublishers() UserModelV1StatsModerationPublishers {
	if o == nil || IsNil(o.Publishers) {
		var ret UserModelV1StatsModerationPublishers
		return ret
	}
	return *o.Publishers
}

// GetPublishersOk returns a tuple with the Publishers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetPublishersOk() (*UserModelV1StatsModerationPublishers, bool) {
	if o == nil || IsNil(o.Publishers) {
		return nil, false
	}
	return o.Publishers, true
}

// HasPublishers returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasPublishers() bool {
	if o != nil && !IsNil(o.Publishers) {
		return true
	}

	return false
}

// SetPublishers gets a reference to the given UserModelV1StatsModerationPublishers and assigns it to the Publishers field.
func (o *UserModelV1StatsModeration) SetPublishers(v UserModelV1StatsModerationPublishers) {
	o.Publishers = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetGroups() UserModelV1StatsModerationGroups {
	if o == nil || IsNil(o.Groups) {
		var ret UserModelV1StatsModerationGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetGroupsOk() (*UserModelV1StatsModerationGroups, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given UserModelV1StatsModerationGroups and assigns it to the Groups field.
func (o *UserModelV1StatsModeration) SetGroups(v UserModelV1StatsModerationGroups) {
	o.Groups = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetAuthors() UserModelV1StatsModerationAuthors {
	if o == nil || IsNil(o.Authors) {
		var ret UserModelV1StatsModerationAuthors
		return ret
	}
	return *o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetAuthorsOk() (*UserModelV1StatsModerationAuthors, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given UserModelV1StatsModerationAuthors and assigns it to the Authors field.
func (o *UserModelV1StatsModeration) SetAuthors(v UserModelV1StatsModerationAuthors) {
	o.Authors = &v
}

// GetLastAction returns the LastAction field value if set, zero value otherwise.
func (o *UserModelV1StatsModeration) GetLastAction() TimeV1 {
	if o == nil || IsNil(o.LastAction) {
		var ret TimeV1
		return ret
	}
	return *o.LastAction
}

// GetLastActionOk returns a tuple with the LastAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelV1StatsModeration) GetLastActionOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.LastAction) {
		return nil, false
	}
	return o.LastAction, true
}

// HasLastAction returns a boolean if a field has been set.
func (o *UserModelV1StatsModeration) HasLastAction() bool {
	if o != nil && !IsNil(o.LastAction) {
		return true
	}

	return false
}

// SetLastAction gets a reference to the given TimeV1 and assigns it to the LastAction field.
func (o *UserModelV1StatsModeration) SetLastAction(v TimeV1) {
	o.LastAction = &v
}

func (o UserModelV1StatsModeration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModelV1StatsModeration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.Publishers) {
		toSerialize["publishers"] = o.Publishers
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !IsNil(o.LastAction) {
		toSerialize["last_action"] = o.LastAction
	}
	return toSerialize, nil
}

type NullableUserModelV1StatsModeration struct {
	value *UserModelV1StatsModeration
	isSet bool
}

func (v NullableUserModelV1StatsModeration) Get() *UserModelV1StatsModeration {
	return v.value
}

func (v *NullableUserModelV1StatsModeration) Set(val *UserModelV1StatsModeration) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModelV1StatsModeration) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModelV1StatsModeration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModelV1StatsModeration(val *UserModelV1StatsModeration) *NullableUserModelV1StatsModeration {
	return &NullableUserModelV1StatsModeration{value: val, isSet: true}
}

func (v NullableUserModelV1StatsModeration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModelV1StatsModeration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


