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

// checks if the ListsPublicSearchResponseV1ResultsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsPublicSearchResponseV1ResultsMetadata{}

// ListsPublicSearchResponseV1ResultsMetadata struct for ListsPublicSearchResponseV1ResultsMetadata
type ListsPublicSearchResponseV1ResultsMetadata struct {
	UserRating *float32 `json:"user_rating,omitempty"`
	UserComment *ListsPublicSearchResponseV1ResultsMetadataUserComment `json:"user_comment,omitempty"`
	UserList *ListsSeriesModelV1 `json:"user_list,omitempty"`
}

// NewListsPublicSearchResponseV1ResultsMetadata instantiates a new ListsPublicSearchResponseV1ResultsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsPublicSearchResponseV1ResultsMetadata() *ListsPublicSearchResponseV1ResultsMetadata {
	this := ListsPublicSearchResponseV1ResultsMetadata{}
	return &this
}

// NewListsPublicSearchResponseV1ResultsMetadataWithDefaults instantiates a new ListsPublicSearchResponseV1ResultsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsPublicSearchResponseV1ResultsMetadataWithDefaults() *ListsPublicSearchResponseV1ResultsMetadata {
	this := ListsPublicSearchResponseV1ResultsMetadata{}
	return &this
}

// GetUserRating returns the UserRating field value if set, zero value otherwise.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserRating() float32 {
	if o == nil || IsNil(o.UserRating) {
		var ret float32
		return ret
	}
	return *o.UserRating
}

// GetUserRatingOk returns a tuple with the UserRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserRatingOk() (*float32, bool) {
	if o == nil || IsNil(o.UserRating) {
		return nil, false
	}
	return o.UserRating, true
}

// HasUserRating returns a boolean if a field has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserRating() bool {
	if o != nil && !IsNil(o.UserRating) {
		return true
	}

	return false
}

// SetUserRating gets a reference to the given float32 and assigns it to the UserRating field.
func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserRating(v float32) {
	o.UserRating = &v
}

// GetUserComment returns the UserComment field value if set, zero value otherwise.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserComment() ListsPublicSearchResponseV1ResultsMetadataUserComment {
	if o == nil || IsNil(o.UserComment) {
		var ret ListsPublicSearchResponseV1ResultsMetadataUserComment
		return ret
	}
	return *o.UserComment
}

// GetUserCommentOk returns a tuple with the UserComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserCommentOk() (*ListsPublicSearchResponseV1ResultsMetadataUserComment, bool) {
	if o == nil || IsNil(o.UserComment) {
		return nil, false
	}
	return o.UserComment, true
}

// HasUserComment returns a boolean if a field has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserComment() bool {
	if o != nil && !IsNil(o.UserComment) {
		return true
	}

	return false
}

// SetUserComment gets a reference to the given ListsPublicSearchResponseV1ResultsMetadataUserComment and assigns it to the UserComment field.
func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserComment(v ListsPublicSearchResponseV1ResultsMetadataUserComment) {
	o.UserComment = &v
}

// GetUserList returns the UserList field value if set, zero value otherwise.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserList() ListsSeriesModelV1 {
	if o == nil || IsNil(o.UserList) {
		var ret ListsSeriesModelV1
		return ret
	}
	return *o.UserList
}

// GetUserListOk returns a tuple with the UserList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserListOk() (*ListsSeriesModelV1, bool) {
	if o == nil || IsNil(o.UserList) {
		return nil, false
	}
	return o.UserList, true
}

// HasUserList returns a boolean if a field has been set.
func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserList() bool {
	if o != nil && !IsNil(o.UserList) {
		return true
	}

	return false
}

// SetUserList gets a reference to the given ListsSeriesModelV1 and assigns it to the UserList field.
func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserList(v ListsSeriesModelV1) {
	o.UserList = &v
}

func (o ListsPublicSearchResponseV1ResultsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsPublicSearchResponseV1ResultsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserRating) {
		toSerialize["user_rating"] = o.UserRating
	}
	if !IsNil(o.UserComment) {
		toSerialize["user_comment"] = o.UserComment
	}
	if !IsNil(o.UserList) {
		toSerialize["user_list"] = o.UserList
	}
	return toSerialize, nil
}

type NullableListsPublicSearchResponseV1ResultsMetadata struct {
	value *ListsPublicSearchResponseV1ResultsMetadata
	isSet bool
}

func (v NullableListsPublicSearchResponseV1ResultsMetadata) Get() *ListsPublicSearchResponseV1ResultsMetadata {
	return v.value
}

func (v *NullableListsPublicSearchResponseV1ResultsMetadata) Set(val *ListsPublicSearchResponseV1ResultsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableListsPublicSearchResponseV1ResultsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableListsPublicSearchResponseV1ResultsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsPublicSearchResponseV1ResultsMetadata(val *ListsPublicSearchResponseV1ResultsMetadata) *NullableListsPublicSearchResponseV1ResultsMetadata {
	return &NullableListsPublicSearchResponseV1ResultsMetadata{value: val, isSet: true}
}

func (v NullableListsPublicSearchResponseV1ResultsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsPublicSearchResponseV1ResultsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


