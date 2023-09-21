# ConvoSearchResponseV1ResultsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to [**ConvoMessageModelV1**](ConvoMessageModelV1.md) |  | [optional] 
**Participant** | Pointer to [**ConvoParticipantModelV1**](ConvoParticipantModelV1.md) |  | [optional] 

## Methods

### NewConvoSearchResponseV1ResultsMetadata

`func NewConvoSearchResponseV1ResultsMetadata() *ConvoSearchResponseV1ResultsMetadata`

NewConvoSearchResponseV1ResultsMetadata instantiates a new ConvoSearchResponseV1ResultsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvoSearchResponseV1ResultsMetadataWithDefaults

`func NewConvoSearchResponseV1ResultsMetadataWithDefaults() *ConvoSearchResponseV1ResultsMetadata`

NewConvoSearchResponseV1ResultsMetadataWithDefaults instantiates a new ConvoSearchResponseV1ResultsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ConvoSearchResponseV1ResultsMetadata) GetMessage() ConvoMessageModelV1`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConvoSearchResponseV1ResultsMetadata) GetMessageOk() (*ConvoMessageModelV1, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConvoSearchResponseV1ResultsMetadata) SetMessage(v ConvoMessageModelV1)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConvoSearchResponseV1ResultsMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParticipant

`func (o *ConvoSearchResponseV1ResultsMetadata) GetParticipant() ConvoParticipantModelV1`

GetParticipant returns the Participant field if non-nil, zero value otherwise.

### GetParticipantOk

`func (o *ConvoSearchResponseV1ResultsMetadata) GetParticipantOk() (*ConvoParticipantModelV1, bool)`

GetParticipantOk returns a tuple with the Participant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipant

`func (o *ConvoSearchResponseV1ResultsMetadata) SetParticipant(v ConvoParticipantModelV1)`

SetParticipant sets Participant field to given value.

### HasParticipant

`func (o *ConvoSearchResponseV1ResultsMetadata) HasParticipant() bool`

HasParticipant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


