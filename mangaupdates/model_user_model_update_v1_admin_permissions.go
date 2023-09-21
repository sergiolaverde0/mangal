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

// checks if the UserModelUpdateV1AdminPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModelUpdateV1AdminPermissions{}

// UserModelUpdateV1AdminPermissions struct for UserModelUpdateV1AdminPermissions
type UserModelUpdateV1AdminPermissions struct {
	PAddReleases *bool `json:"p_add_releases,omitempty"`
	PEditUsers *bool `json:"p_edit_users,omitempty"`
	PEditGroups *bool `json:"p_edit_groups,omitempty"`
	PEditPoll *bool `json:"p_edit_poll,omitempty"`
	PEditSeries *bool `json:"p_edit_series,omitempty"`
	PEditReviews *bool `json:"p_edit_reviews,omitempty"`
	PEditNews *bool `json:"p_edit_news,omitempty"`
	PEditAffiliates *bool `json:"p_edit_affiliates,omitempty"`
	PEditAboutus *bool `json:"p_edit_aboutus,omitempty"`
	PViewLog *bool `json:"p_view_log,omitempty"`
	PEditConfig *bool `json:"p_edit_config,omitempty"`
	PViewStats *bool `json:"p_view_stats,omitempty"`
	PEditGenre *bool `json:"p_edit_genre,omitempty"`
	PEditAuthors *bool `json:"p_edit_authors,omitempty"`
	PEditPublishers *bool `json:"p_edit_publishers,omitempty"`
	PEditPartialUsers *bool `json:"p_edit_partial_users,omitempty"`
}

// NewUserModelUpdateV1AdminPermissions instantiates a new UserModelUpdateV1AdminPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModelUpdateV1AdminPermissions() *UserModelUpdateV1AdminPermissions {
	this := UserModelUpdateV1AdminPermissions{}
	return &this
}

// NewUserModelUpdateV1AdminPermissionsWithDefaults instantiates a new UserModelUpdateV1AdminPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelUpdateV1AdminPermissionsWithDefaults() *UserModelUpdateV1AdminPermissions {
	this := UserModelUpdateV1AdminPermissions{}
	return &this
}

// GetPAddReleases returns the PAddReleases field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPAddReleases() bool {
	if o == nil || IsNil(o.PAddReleases) {
		var ret bool
		return ret
	}
	return *o.PAddReleases
}

// GetPAddReleasesOk returns a tuple with the PAddReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPAddReleasesOk() (*bool, bool) {
	if o == nil || IsNil(o.PAddReleases) {
		return nil, false
	}
	return o.PAddReleases, true
}

// HasPAddReleases returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPAddReleases() bool {
	if o != nil && !IsNil(o.PAddReleases) {
		return true
	}

	return false
}

// SetPAddReleases gets a reference to the given bool and assigns it to the PAddReleases field.
func (o *UserModelUpdateV1AdminPermissions) SetPAddReleases(v bool) {
	o.PAddReleases = &v
}

// GetPEditUsers returns the PEditUsers field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditUsers() bool {
	if o == nil || IsNil(o.PEditUsers) {
		var ret bool
		return ret
	}
	return *o.PEditUsers
}

// GetPEditUsersOk returns a tuple with the PEditUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditUsers) {
		return nil, false
	}
	return o.PEditUsers, true
}

// HasPEditUsers returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditUsers() bool {
	if o != nil && !IsNil(o.PEditUsers) {
		return true
	}

	return false
}

// SetPEditUsers gets a reference to the given bool and assigns it to the PEditUsers field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditUsers(v bool) {
	o.PEditUsers = &v
}

// GetPEditGroups returns the PEditGroups field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditGroups() bool {
	if o == nil || IsNil(o.PEditGroups) {
		var ret bool
		return ret
	}
	return *o.PEditGroups
}

// GetPEditGroupsOk returns a tuple with the PEditGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditGroups) {
		return nil, false
	}
	return o.PEditGroups, true
}

// HasPEditGroups returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditGroups() bool {
	if o != nil && !IsNil(o.PEditGroups) {
		return true
	}

	return false
}

// SetPEditGroups gets a reference to the given bool and assigns it to the PEditGroups field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditGroups(v bool) {
	o.PEditGroups = &v
}

// GetPEditPoll returns the PEditPoll field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPoll() bool {
	if o == nil || IsNil(o.PEditPoll) {
		var ret bool
		return ret
	}
	return *o.PEditPoll
}

// GetPEditPollOk returns a tuple with the PEditPoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPollOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditPoll) {
		return nil, false
	}
	return o.PEditPoll, true
}

// HasPEditPoll returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditPoll() bool {
	if o != nil && !IsNil(o.PEditPoll) {
		return true
	}

	return false
}

// SetPEditPoll gets a reference to the given bool and assigns it to the PEditPoll field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditPoll(v bool) {
	o.PEditPoll = &v
}

// GetPEditSeries returns the PEditSeries field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditSeries() bool {
	if o == nil || IsNil(o.PEditSeries) {
		var ret bool
		return ret
	}
	return *o.PEditSeries
}

// GetPEditSeriesOk returns a tuple with the PEditSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditSeriesOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditSeries) {
		return nil, false
	}
	return o.PEditSeries, true
}

// HasPEditSeries returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditSeries() bool {
	if o != nil && !IsNil(o.PEditSeries) {
		return true
	}

	return false
}

// SetPEditSeries gets a reference to the given bool and assigns it to the PEditSeries field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditSeries(v bool) {
	o.PEditSeries = &v
}

// GetPEditReviews returns the PEditReviews field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditReviews() bool {
	if o == nil || IsNil(o.PEditReviews) {
		var ret bool
		return ret
	}
	return *o.PEditReviews
}

// GetPEditReviewsOk returns a tuple with the PEditReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditReviewsOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditReviews) {
		return nil, false
	}
	return o.PEditReviews, true
}

// HasPEditReviews returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditReviews() bool {
	if o != nil && !IsNil(o.PEditReviews) {
		return true
	}

	return false
}

// SetPEditReviews gets a reference to the given bool and assigns it to the PEditReviews field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditReviews(v bool) {
	o.PEditReviews = &v
}

// GetPEditNews returns the PEditNews field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditNews() bool {
	if o == nil || IsNil(o.PEditNews) {
		var ret bool
		return ret
	}
	return *o.PEditNews
}

// GetPEditNewsOk returns a tuple with the PEditNews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditNewsOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditNews) {
		return nil, false
	}
	return o.PEditNews, true
}

// HasPEditNews returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditNews() bool {
	if o != nil && !IsNil(o.PEditNews) {
		return true
	}

	return false
}

// SetPEditNews gets a reference to the given bool and assigns it to the PEditNews field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditNews(v bool) {
	o.PEditNews = &v
}

// GetPEditAffiliates returns the PEditAffiliates field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAffiliates() bool {
	if o == nil || IsNil(o.PEditAffiliates) {
		var ret bool
		return ret
	}
	return *o.PEditAffiliates
}

// GetPEditAffiliatesOk returns a tuple with the PEditAffiliates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAffiliatesOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditAffiliates) {
		return nil, false
	}
	return o.PEditAffiliates, true
}

// HasPEditAffiliates returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditAffiliates() bool {
	if o != nil && !IsNil(o.PEditAffiliates) {
		return true
	}

	return false
}

// SetPEditAffiliates gets a reference to the given bool and assigns it to the PEditAffiliates field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditAffiliates(v bool) {
	o.PEditAffiliates = &v
}

// GetPEditAboutus returns the PEditAboutus field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAboutus() bool {
	if o == nil || IsNil(o.PEditAboutus) {
		var ret bool
		return ret
	}
	return *o.PEditAboutus
}

// GetPEditAboutusOk returns a tuple with the PEditAboutus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAboutusOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditAboutus) {
		return nil, false
	}
	return o.PEditAboutus, true
}

// HasPEditAboutus returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditAboutus() bool {
	if o != nil && !IsNil(o.PEditAboutus) {
		return true
	}

	return false
}

// SetPEditAboutus gets a reference to the given bool and assigns it to the PEditAboutus field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditAboutus(v bool) {
	o.PEditAboutus = &v
}

// GetPViewLog returns the PViewLog field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPViewLog() bool {
	if o == nil || IsNil(o.PViewLog) {
		var ret bool
		return ret
	}
	return *o.PViewLog
}

// GetPViewLogOk returns a tuple with the PViewLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPViewLogOk() (*bool, bool) {
	if o == nil || IsNil(o.PViewLog) {
		return nil, false
	}
	return o.PViewLog, true
}

// HasPViewLog returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPViewLog() bool {
	if o != nil && !IsNil(o.PViewLog) {
		return true
	}

	return false
}

// SetPViewLog gets a reference to the given bool and assigns it to the PViewLog field.
func (o *UserModelUpdateV1AdminPermissions) SetPViewLog(v bool) {
	o.PViewLog = &v
}

// GetPEditConfig returns the PEditConfig field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditConfig() bool {
	if o == nil || IsNil(o.PEditConfig) {
		var ret bool
		return ret
	}
	return *o.PEditConfig
}

// GetPEditConfigOk returns a tuple with the PEditConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditConfig) {
		return nil, false
	}
	return o.PEditConfig, true
}

// HasPEditConfig returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditConfig() bool {
	if o != nil && !IsNil(o.PEditConfig) {
		return true
	}

	return false
}

// SetPEditConfig gets a reference to the given bool and assigns it to the PEditConfig field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditConfig(v bool) {
	o.PEditConfig = &v
}

// GetPViewStats returns the PViewStats field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPViewStats() bool {
	if o == nil || IsNil(o.PViewStats) {
		var ret bool
		return ret
	}
	return *o.PViewStats
}

// GetPViewStatsOk returns a tuple with the PViewStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPViewStatsOk() (*bool, bool) {
	if o == nil || IsNil(o.PViewStats) {
		return nil, false
	}
	return o.PViewStats, true
}

// HasPViewStats returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPViewStats() bool {
	if o != nil && !IsNil(o.PViewStats) {
		return true
	}

	return false
}

// SetPViewStats gets a reference to the given bool and assigns it to the PViewStats field.
func (o *UserModelUpdateV1AdminPermissions) SetPViewStats(v bool) {
	o.PViewStats = &v
}

// GetPEditGenre returns the PEditGenre field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditGenre() bool {
	if o == nil || IsNil(o.PEditGenre) {
		var ret bool
		return ret
	}
	return *o.PEditGenre
}

// GetPEditGenreOk returns a tuple with the PEditGenre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditGenreOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditGenre) {
		return nil, false
	}
	return o.PEditGenre, true
}

// HasPEditGenre returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditGenre() bool {
	if o != nil && !IsNil(o.PEditGenre) {
		return true
	}

	return false
}

// SetPEditGenre gets a reference to the given bool and assigns it to the PEditGenre field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditGenre(v bool) {
	o.PEditGenre = &v
}

// GetPEditAuthors returns the PEditAuthors field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAuthors() bool {
	if o == nil || IsNil(o.PEditAuthors) {
		var ret bool
		return ret
	}
	return *o.PEditAuthors
}

// GetPEditAuthorsOk returns a tuple with the PEditAuthors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditAuthorsOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditAuthors) {
		return nil, false
	}
	return o.PEditAuthors, true
}

// HasPEditAuthors returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditAuthors() bool {
	if o != nil && !IsNil(o.PEditAuthors) {
		return true
	}

	return false
}

// SetPEditAuthors gets a reference to the given bool and assigns it to the PEditAuthors field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditAuthors(v bool) {
	o.PEditAuthors = &v
}

// GetPEditPublishers returns the PEditPublishers field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPublishers() bool {
	if o == nil || IsNil(o.PEditPublishers) {
		var ret bool
		return ret
	}
	return *o.PEditPublishers
}

// GetPEditPublishersOk returns a tuple with the PEditPublishers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPublishersOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditPublishers) {
		return nil, false
	}
	return o.PEditPublishers, true
}

// HasPEditPublishers returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditPublishers() bool {
	if o != nil && !IsNil(o.PEditPublishers) {
		return true
	}

	return false
}

// SetPEditPublishers gets a reference to the given bool and assigns it to the PEditPublishers field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditPublishers(v bool) {
	o.PEditPublishers = &v
}

// GetPEditPartialUsers returns the PEditPartialUsers field value if set, zero value otherwise.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPartialUsers() bool {
	if o == nil || IsNil(o.PEditPartialUsers) {
		var ret bool
		return ret
	}
	return *o.PEditPartialUsers
}

// GetPEditPartialUsersOk returns a tuple with the PEditPartialUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1AdminPermissions) GetPEditPartialUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.PEditPartialUsers) {
		return nil, false
	}
	return o.PEditPartialUsers, true
}

// HasPEditPartialUsers returns a boolean if a field has been set.
func (o *UserModelUpdateV1AdminPermissions) HasPEditPartialUsers() bool {
	if o != nil && !IsNil(o.PEditPartialUsers) {
		return true
	}

	return false
}

// SetPEditPartialUsers gets a reference to the given bool and assigns it to the PEditPartialUsers field.
func (o *UserModelUpdateV1AdminPermissions) SetPEditPartialUsers(v bool) {
	o.PEditPartialUsers = &v
}

func (o UserModelUpdateV1AdminPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModelUpdateV1AdminPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PAddReleases) {
		toSerialize["p_add_releases"] = o.PAddReleases
	}
	if !IsNil(o.PEditUsers) {
		toSerialize["p_edit_users"] = o.PEditUsers
	}
	if !IsNil(o.PEditGroups) {
		toSerialize["p_edit_groups"] = o.PEditGroups
	}
	if !IsNil(o.PEditPoll) {
		toSerialize["p_edit_poll"] = o.PEditPoll
	}
	if !IsNil(o.PEditSeries) {
		toSerialize["p_edit_series"] = o.PEditSeries
	}
	if !IsNil(o.PEditReviews) {
		toSerialize["p_edit_reviews"] = o.PEditReviews
	}
	if !IsNil(o.PEditNews) {
		toSerialize["p_edit_news"] = o.PEditNews
	}
	if !IsNil(o.PEditAffiliates) {
		toSerialize["p_edit_affiliates"] = o.PEditAffiliates
	}
	if !IsNil(o.PEditAboutus) {
		toSerialize["p_edit_aboutus"] = o.PEditAboutus
	}
	if !IsNil(o.PViewLog) {
		toSerialize["p_view_log"] = o.PViewLog
	}
	if !IsNil(o.PEditConfig) {
		toSerialize["p_edit_config"] = o.PEditConfig
	}
	if !IsNil(o.PViewStats) {
		toSerialize["p_view_stats"] = o.PViewStats
	}
	if !IsNil(o.PEditGenre) {
		toSerialize["p_edit_genre"] = o.PEditGenre
	}
	if !IsNil(o.PEditAuthors) {
		toSerialize["p_edit_authors"] = o.PEditAuthors
	}
	if !IsNil(o.PEditPublishers) {
		toSerialize["p_edit_publishers"] = o.PEditPublishers
	}
	if !IsNil(o.PEditPartialUsers) {
		toSerialize["p_edit_partial_users"] = o.PEditPartialUsers
	}
	return toSerialize, nil
}

type NullableUserModelUpdateV1AdminPermissions struct {
	value *UserModelUpdateV1AdminPermissions
	isSet bool
}

func (v NullableUserModelUpdateV1AdminPermissions) Get() *UserModelUpdateV1AdminPermissions {
	return v.value
}

func (v *NullableUserModelUpdateV1AdminPermissions) Set(val *UserModelUpdateV1AdminPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModelUpdateV1AdminPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModelUpdateV1AdminPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModelUpdateV1AdminPermissions(val *UserModelUpdateV1AdminPermissions) *NullableUserModelUpdateV1AdminPermissions {
	return &NullableUserModelUpdateV1AdminPermissions{value: val, isSet: true}
}

func (v NullableUserModelUpdateV1AdminPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModelUpdateV1AdminPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


