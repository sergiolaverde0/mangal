# ConvoSearchResponseV1Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**ConvoModelV1**](ConvoModelV1.md) |  | [optional] 
**Metadata** | Pointer to [**ConvoSearchResponseV1ResultsMetadata**](ConvoSearchResponseV1ResultsMetadata.md) |  | [optional] 

## Methods

### NewConvoSearchResponseV1Results

`func NewConvoSearchResponseV1Results() *ConvoSearchResponseV1Results`

NewConvoSearchResponseV1Results instantiates a new ConvoSearchResponseV1Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvoSearchResponseV1ResultsWithDefaults

`func NewConvoSearchResponseV1ResultsWithDefaults() *ConvoSearchResponseV1Results`

NewConvoSearchResponseV1ResultsWithDefaults instantiates a new ConvoSearchResponseV1Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *ConvoSearchResponseV1Results) GetRecord() ConvoModelV1`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *ConvoSearchResponseV1Results) GetRecordOk() (*ConvoModelV1, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *ConvoSearchResponseV1Results) SetRecord(v ConvoModelV1)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *ConvoSearchResponseV1Results) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetMetadata

`func (o *ConvoSearchResponseV1Results) GetMetadata() ConvoSearchResponseV1ResultsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConvoSearchResponseV1Results) GetMetadataOk() (*ConvoSearchResponseV1ResultsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConvoSearchResponseV1Results) SetMetadata(v ConvoSearchResponseV1ResultsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConvoSearchResponseV1Results) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


