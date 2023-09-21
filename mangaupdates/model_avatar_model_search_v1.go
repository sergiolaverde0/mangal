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

// checks if the AvatarModelSearchV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvatarModelSearchV1{}

// AvatarModelSearchV1 struct for AvatarModelSearchV1
type AvatarModelSearchV1 struct {
	Id     *int64  `json:"id,omitempty"`
	Url    *string `json:"url,omitempty"`
	Height *int64  `json:"height,omitempty"`
	Width  *int64  `json:"width,omitempty"`
}

// NewAvatarModelSearchV1 instantiates a new AvatarModelSearchV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvatarModelSearchV1() *AvatarModelSearchV1 {
	this := AvatarModelSearchV1{}
	return &this
}

// NewAvatarModelSearchV1WithDefaults instantiates a new AvatarModelSearchV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvatarModelSearchV1WithDefaults() *AvatarModelSearchV1 {
	this := AvatarModelSearchV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AvatarModelSearchV1) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvatarModelSearchV1) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AvatarModelSearchV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AvatarModelSearchV1) SetId(v int64) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AvatarModelSearchV1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvatarModelSearchV1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AvatarModelSearchV1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AvatarModelSearchV1) SetUrl(v string) {
	o.Url = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *AvatarModelSearchV1) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvatarModelSearchV1) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *AvatarModelSearchV1) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *AvatarModelSearchV1) SetHeight(v int64) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *AvatarModelSearchV1) GetWidth() int64 {
	if o == nil || IsNil(o.Width) {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvatarModelSearchV1) GetWidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *AvatarModelSearchV1) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *AvatarModelSearchV1) SetWidth(v int64) {
	o.Width = &v
}

func (o AvatarModelSearchV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvatarModelSearchV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableAvatarModelSearchV1 struct {
	value *AvatarModelSearchV1
	isSet bool
}

func (v NullableAvatarModelSearchV1) Get() *AvatarModelSearchV1 {
	return v.value
}

func (v *NullableAvatarModelSearchV1) Set(val *AvatarModelSearchV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAvatarModelSearchV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAvatarModelSearchV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvatarModelSearchV1(val *AvatarModelSearchV1) *NullableAvatarModelSearchV1 {
	return &NullableAvatarModelSearchV1{value: val, isSet: true}
}

func (v NullableAvatarModelSearchV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvatarModelSearchV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
