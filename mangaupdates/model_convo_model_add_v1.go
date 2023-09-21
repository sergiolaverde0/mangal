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

// checks if the ConvoModelAddV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvoModelAddV1{}

// ConvoModelAddV1 struct for ConvoModelAddV1
type ConvoModelAddV1 struct {
	Topic *string `json:"topic,omitempty"`
	Participants []ConvoParticipantModelAddV1 `json:"participants,omitempty"`
	Message *ConvoMessageModelUpdateV1 `json:"message,omitempty"`
}

// NewConvoModelAddV1 instantiates a new ConvoModelAddV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvoModelAddV1() *ConvoModelAddV1 {
	this := ConvoModelAddV1{}
	return &this
}

// NewConvoModelAddV1WithDefaults instantiates a new ConvoModelAddV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvoModelAddV1WithDefaults() *ConvoModelAddV1 {
	this := ConvoModelAddV1{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ConvoModelAddV1) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelAddV1) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ConvoModelAddV1) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ConvoModelAddV1) SetTopic(v string) {
	o.Topic = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *ConvoModelAddV1) GetParticipants() []ConvoParticipantModelAddV1 {
	if o == nil || IsNil(o.Participants) {
		var ret []ConvoParticipantModelAddV1
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelAddV1) GetParticipantsOk() ([]ConvoParticipantModelAddV1, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *ConvoModelAddV1) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []ConvoParticipantModelAddV1 and assigns it to the Participants field.
func (o *ConvoModelAddV1) SetParticipants(v []ConvoParticipantModelAddV1) {
	o.Participants = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConvoModelAddV1) GetMessage() ConvoMessageModelUpdateV1 {
	if o == nil || IsNil(o.Message) {
		var ret ConvoMessageModelUpdateV1
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelAddV1) GetMessageOk() (*ConvoMessageModelUpdateV1, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConvoModelAddV1) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ConvoMessageModelUpdateV1 and assigns it to the Message field.
func (o *ConvoModelAddV1) SetMessage(v ConvoMessageModelUpdateV1) {
	o.Message = &v
}

func (o ConvoModelAddV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvoModelAddV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableConvoModelAddV1 struct {
	value *ConvoModelAddV1
	isSet bool
}

func (v NullableConvoModelAddV1) Get() *ConvoModelAddV1 {
	return v.value
}

func (v *NullableConvoModelAddV1) Set(val *ConvoModelAddV1) {
	v.value = val
	v.isSet = true
}

func (v NullableConvoModelAddV1) IsSet() bool {
	return v.isSet
}

func (v *NullableConvoModelAddV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvoModelAddV1(val *ConvoModelAddV1) *NullableConvoModelAddV1 {
	return &NullableConvoModelAddV1{value: val, isSet: true}
}

func (v NullableConvoModelAddV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvoModelAddV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

