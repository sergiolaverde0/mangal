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

// checks if the UserModelUpdateV1Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserModelUpdateV1Profile{}

// UserModelUpdateV1Profile struct for UserModelUpdateV1Profile
type UserModelUpdateV1Profile struct {
	PerPage        *int64                           `json:"per_page,omitempty"`
	Invisible      *bool                            `json:"invisible,omitempty"`
	HideBirthday   *bool                            `json:"hide_birthday,omitempty"`
	HideCategories *bool                            `json:"hide_categories,omitempty"`
	FilterTypes    []string                         `json:"filter_types,omitempty"`
	Upgrade        *UserModelUpdateV1ProfileUpgrade `json:"upgrade,omitempty"`
	Age18Verified  *bool                            `json:"age18_verified,omitempty"`
}

// NewUserModelUpdateV1Profile instantiates a new UserModelUpdateV1Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserModelUpdateV1Profile() *UserModelUpdateV1Profile {
	this := UserModelUpdateV1Profile{}
	return &this
}

// NewUserModelUpdateV1ProfileWithDefaults instantiates a new UserModelUpdateV1Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserModelUpdateV1ProfileWithDefaults() *UserModelUpdateV1Profile {
	this := UserModelUpdateV1Profile{}
	return &this
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetPerPage() int64 {
	if o == nil || IsNil(o.PerPage) {
		var ret int64
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetPerPageOk() (*int64, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int64 and assigns it to the PerPage field.
func (o *UserModelUpdateV1Profile) SetPerPage(v int64) {
	o.PerPage = &v
}

// GetInvisible returns the Invisible field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetInvisible() bool {
	if o == nil || IsNil(o.Invisible) {
		var ret bool
		return ret
	}
	return *o.Invisible
}

// GetInvisibleOk returns a tuple with the Invisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetInvisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Invisible) {
		return nil, false
	}
	return o.Invisible, true
}

// HasInvisible returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasInvisible() bool {
	if o != nil && !IsNil(o.Invisible) {
		return true
	}

	return false
}

// SetInvisible gets a reference to the given bool and assigns it to the Invisible field.
func (o *UserModelUpdateV1Profile) SetInvisible(v bool) {
	o.Invisible = &v
}

// GetHideBirthday returns the HideBirthday field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetHideBirthday() bool {
	if o == nil || IsNil(o.HideBirthday) {
		var ret bool
		return ret
	}
	return *o.HideBirthday
}

// GetHideBirthdayOk returns a tuple with the HideBirthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetHideBirthdayOk() (*bool, bool) {
	if o == nil || IsNil(o.HideBirthday) {
		return nil, false
	}
	return o.HideBirthday, true
}

// HasHideBirthday returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasHideBirthday() bool {
	if o != nil && !IsNil(o.HideBirthday) {
		return true
	}

	return false
}

// SetHideBirthday gets a reference to the given bool and assigns it to the HideBirthday field.
func (o *UserModelUpdateV1Profile) SetHideBirthday(v bool) {
	o.HideBirthday = &v
}

// GetHideCategories returns the HideCategories field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetHideCategories() bool {
	if o == nil || IsNil(o.HideCategories) {
		var ret bool
		return ret
	}
	return *o.HideCategories
}

// GetHideCategoriesOk returns a tuple with the HideCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetHideCategoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.HideCategories) {
		return nil, false
	}
	return o.HideCategories, true
}

// HasHideCategories returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasHideCategories() bool {
	if o != nil && !IsNil(o.HideCategories) {
		return true
	}

	return false
}

// SetHideCategories gets a reference to the given bool and assigns it to the HideCategories field.
func (o *UserModelUpdateV1Profile) SetHideCategories(v bool) {
	o.HideCategories = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetFilterTypes() []string {
	if o == nil || IsNil(o.FilterTypes) {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *UserModelUpdateV1Profile) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetUpgrade() UserModelUpdateV1ProfileUpgrade {
	if o == nil || IsNil(o.Upgrade) {
		var ret UserModelUpdateV1ProfileUpgrade
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetUpgradeOk() (*UserModelUpdateV1ProfileUpgrade, bool) {
	if o == nil || IsNil(o.Upgrade) {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasUpgrade() bool {
	if o != nil && !IsNil(o.Upgrade) {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given UserModelUpdateV1ProfileUpgrade and assigns it to the Upgrade field.
func (o *UserModelUpdateV1Profile) SetUpgrade(v UserModelUpdateV1ProfileUpgrade) {
	o.Upgrade = &v
}

// GetAge18Verified returns the Age18Verified field value if set, zero value otherwise.
func (o *UserModelUpdateV1Profile) GetAge18Verified() bool {
	if o == nil || IsNil(o.Age18Verified) {
		var ret bool
		return ret
	}
	return *o.Age18Verified
}

// GetAge18VerifiedOk returns a tuple with the Age18Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserModelUpdateV1Profile) GetAge18VerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Age18Verified) {
		return nil, false
	}
	return o.Age18Verified, true
}

// HasAge18Verified returns a boolean if a field has been set.
func (o *UserModelUpdateV1Profile) HasAge18Verified() bool {
	if o != nil && !IsNil(o.Age18Verified) {
		return true
	}

	return false
}

// SetAge18Verified gets a reference to the given bool and assigns it to the Age18Verified field.
func (o *UserModelUpdateV1Profile) SetAge18Verified(v bool) {
	o.Age18Verified = &v
}

func (o UserModelUpdateV1Profile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserModelUpdateV1Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !IsNil(o.Invisible) {
		toSerialize["invisible"] = o.Invisible
	}
	if !IsNil(o.HideBirthday) {
		toSerialize["hide_birthday"] = o.HideBirthday
	}
	if !IsNil(o.HideCategories) {
		toSerialize["hide_categories"] = o.HideCategories
	}
	if !IsNil(o.FilterTypes) {
		toSerialize["filter_types"] = o.FilterTypes
	}
	if !IsNil(o.Upgrade) {
		toSerialize["upgrade"] = o.Upgrade
	}
	if !IsNil(o.Age18Verified) {
		toSerialize["age18_verified"] = o.Age18Verified
	}
	return toSerialize, nil
}

type NullableUserModelUpdateV1Profile struct {
	value *UserModelUpdateV1Profile
	isSet bool
}

func (v NullableUserModelUpdateV1Profile) Get() *UserModelUpdateV1Profile {
	return v.value
}

func (v *NullableUserModelUpdateV1Profile) Set(val *UserModelUpdateV1Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserModelUpdateV1Profile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserModelUpdateV1Profile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserModelUpdateV1Profile(val *UserModelUpdateV1Profile) *NullableUserModelUpdateV1Profile {
	return &NullableUserModelUpdateV1Profile{value: val, isSet: true}
}

func (v NullableUserModelUpdateV1Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserModelUpdateV1Profile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}