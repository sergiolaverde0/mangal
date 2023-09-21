# ConvoModelAddV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to [**[]ConvoParticipantModelAddV1**](ConvoParticipantModelAddV1.md) |  | [optional] 
**Message** | Pointer to [**ConvoMessageModelUpdateV1**](ConvoMessageModelUpdateV1.md) |  | [optional] 

## Methods

### NewConvoModelAddV1

`func NewConvoModelAddV1() *ConvoModelAddV1`

NewConvoModelAddV1 instantiates a new ConvoModelAddV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvoModelAddV1WithDefaults

`func NewConvoModelAddV1WithDefaults() *ConvoModelAddV1`

NewConvoModelAddV1WithDefaults instantiates a new ConvoModelAddV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *ConvoModelAddV1) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConvoModelAddV1) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ConvoModelAddV1) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ConvoModelAddV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetParticipants

`func (o *ConvoModelAddV1) GetParticipants() []ConvoParticipantModelAddV1`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ConvoModelAddV1) GetParticipantsOk() (*[]ConvoParticipantModelAddV1, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ConvoModelAddV1) SetParticipants(v []ConvoParticipantModelAddV1)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ConvoModelAddV1) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetMessage

`func (o *ConvoModelAddV1) GetMessage() ConvoMessageModelUpdateV1`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConvoModelAddV1) GetMessageOk() (*ConvoMessageModelUpdateV1, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConvoModelAddV1) SetMessage(v ConvoMessageModelUpdateV1)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConvoModelAddV1) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


