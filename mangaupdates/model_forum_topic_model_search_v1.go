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

// checks if the ForumTopicModelSearchV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumTopicModelSearchV1{}

// ForumTopicModelSearchV1 struct for ForumTopicModelSearchV1
type ForumTopicModelSearchV1 struct {
	TopicId      *int64                        `json:"topic_id,omitempty"`
	Topic        *string                       `json:"topic,omitempty"`
	Url          *string                       `json:"url,omitempty"`
	LastPost     *ForumPostModelSearchV1       `json:"last_post,omitempty"`
	Stats        *ForumTopicModelSearchV1Stats `json:"stats,omitempty"`
	Forum        *ForumTopicModelSearchV1Forum `json:"forum,omitempty"`
	IsPoll       *bool                         `json:"is_poll,omitempty"`
	Admin        *ForumTopicModelSearchV1Admin `json:"admin,omitempty"`
	TopicStarter *UserModelSearchV1            `json:"topic_starter,omitempty"`
	TimeAdded    *TimeV1                       `json:"time_added,omitempty"`
}

// NewForumTopicModelSearchV1 instantiates a new ForumTopicModelSearchV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumTopicModelSearchV1() *ForumTopicModelSearchV1 {
	this := ForumTopicModelSearchV1{}
	return &this
}

// NewForumTopicModelSearchV1WithDefaults instantiates a new ForumTopicModelSearchV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumTopicModelSearchV1WithDefaults() *ForumTopicModelSearchV1 {
	this := ForumTopicModelSearchV1{}
	return &this
}

// GetTopicId returns the TopicId field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetTopicId() int64 {
	if o == nil || IsNil(o.TopicId) {
		var ret int64
		return ret
	}
	return *o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetTopicIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TopicId) {
		return nil, false
	}
	return o.TopicId, true
}

// HasTopicId returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasTopicId() bool {
	if o != nil && !IsNil(o.TopicId) {
		return true
	}

	return false
}

// SetTopicId gets a reference to the given int64 and assigns it to the TopicId field.
func (o *ForumTopicModelSearchV1) SetTopicId(v int64) {
	o.TopicId = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ForumTopicModelSearchV1) SetTopic(v string) {
	o.Topic = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ForumTopicModelSearchV1) SetUrl(v string) {
	o.Url = &v
}

// GetLastPost returns the LastPost field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetLastPost() ForumPostModelSearchV1 {
	if o == nil || IsNil(o.LastPost) {
		var ret ForumPostModelSearchV1
		return ret
	}
	return *o.LastPost
}

// GetLastPostOk returns a tuple with the LastPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetLastPostOk() (*ForumPostModelSearchV1, bool) {
	if o == nil || IsNil(o.LastPost) {
		return nil, false
	}
	return o.LastPost, true
}

// HasLastPost returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasLastPost() bool {
	if o != nil && !IsNil(o.LastPost) {
		return true
	}

	return false
}

// SetLastPost gets a reference to the given ForumPostModelSearchV1 and assigns it to the LastPost field.
func (o *ForumTopicModelSearchV1) SetLastPost(v ForumPostModelSearchV1) {
	o.LastPost = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetStats() ForumTopicModelSearchV1Stats {
	if o == nil || IsNil(o.Stats) {
		var ret ForumTopicModelSearchV1Stats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetStatsOk() (*ForumTopicModelSearchV1Stats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ForumTopicModelSearchV1Stats and assigns it to the Stats field.
func (o *ForumTopicModelSearchV1) SetStats(v ForumTopicModelSearchV1Stats) {
	o.Stats = &v
}

// GetForum returns the Forum field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetForum() ForumTopicModelSearchV1Forum {
	if o == nil || IsNil(o.Forum) {
		var ret ForumTopicModelSearchV1Forum
		return ret
	}
	return *o.Forum
}

// GetForumOk returns a tuple with the Forum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetForumOk() (*ForumTopicModelSearchV1Forum, bool) {
	if o == nil || IsNil(o.Forum) {
		return nil, false
	}
	return o.Forum, true
}

// HasForum returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasForum() bool {
	if o != nil && !IsNil(o.Forum) {
		return true
	}

	return false
}

// SetForum gets a reference to the given ForumTopicModelSearchV1Forum and assigns it to the Forum field.
func (o *ForumTopicModelSearchV1) SetForum(v ForumTopicModelSearchV1Forum) {
	o.Forum = &v
}

// GetIsPoll returns the IsPoll field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetIsPoll() bool {
	if o == nil || IsNil(o.IsPoll) {
		var ret bool
		return ret
	}
	return *o.IsPoll
}

// GetIsPollOk returns a tuple with the IsPoll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetIsPollOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPoll) {
		return nil, false
	}
	return o.IsPoll, true
}

// HasIsPoll returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasIsPoll() bool {
	if o != nil && !IsNil(o.IsPoll) {
		return true
	}

	return false
}

// SetIsPoll gets a reference to the given bool and assigns it to the IsPoll field.
func (o *ForumTopicModelSearchV1) SetIsPoll(v bool) {
	o.IsPoll = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetAdmin() ForumTopicModelSearchV1Admin {
	if o == nil || IsNil(o.Admin) {
		var ret ForumTopicModelSearchV1Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetAdminOk() (*ForumTopicModelSearchV1Admin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given ForumTopicModelSearchV1Admin and assigns it to the Admin field.
func (o *ForumTopicModelSearchV1) SetAdmin(v ForumTopicModelSearchV1Admin) {
	o.Admin = &v
}

// GetTopicStarter returns the TopicStarter field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetTopicStarter() UserModelSearchV1 {
	if o == nil || IsNil(o.TopicStarter) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.TopicStarter
}

// GetTopicStarterOk returns a tuple with the TopicStarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetTopicStarterOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.TopicStarter) {
		return nil, false
	}
	return o.TopicStarter, true
}

// HasTopicStarter returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasTopicStarter() bool {
	if o != nil && !IsNil(o.TopicStarter) {
		return true
	}

	return false
}

// SetTopicStarter gets a reference to the given UserModelSearchV1 and assigns it to the TopicStarter field.
func (o *ForumTopicModelSearchV1) SetTopicStarter(v UserModelSearchV1) {
	o.TopicStarter = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *ForumTopicModelSearchV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

func (o ForumTopicModelSearchV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumTopicModelSearchV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TopicId) {
		toSerialize["topic_id"] = o.TopicId
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.LastPost) {
		toSerialize["last_post"] = o.LastPost
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Forum) {
		toSerialize["forum"] = o.Forum
	}
	if !IsNil(o.IsPoll) {
		toSerialize["is_poll"] = o.IsPoll
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.TopicStarter) {
		toSerialize["topic_starter"] = o.TopicStarter
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	return toSerialize, nil
}

type NullableForumTopicModelSearchV1 struct {
	value *ForumTopicModelSearchV1
	isSet bool
}

func (v NullableForumTopicModelSearchV1) Get() *ForumTopicModelSearchV1 {
	return v.value
}

func (v *NullableForumTopicModelSearchV1) Set(val *ForumTopicModelSearchV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumTopicModelSearchV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumTopicModelSearchV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumTopicModelSearchV1(val *ForumTopicModelSearchV1) *NullableForumTopicModelSearchV1 {
	return &NullableForumTopicModelSearchV1{value: val, isSet: true}
}

func (v NullableForumTopicModelSearchV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumTopicModelSearchV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
