# ListsSearchResponseV1Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**ListsSeriesModelV1**](ListsSeriesModelV1.md) |  | [optional] 
**Metadata** | Pointer to [**ListsSearchResponseV1ResultsMetadata**](ListsSearchResponseV1ResultsMetadata.md) |  | [optional] 

## Methods

### NewListsSearchResponseV1Results

`func NewListsSearchResponseV1Results() *ListsSearchResponseV1Results`

NewListsSearchResponseV1Results instantiates a new ListsSearchResponseV1Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsSearchResponseV1ResultsWithDefaults

`func NewListsSearchResponseV1ResultsWithDefaults() *ListsSearchResponseV1Results`

NewListsSearchResponseV1ResultsWithDefaults instantiates a new ListsSearchResponseV1Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *ListsSearchResponseV1Results) GetRecord() ListsSeriesModelV1`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *ListsSearchResponseV1Results) GetRecordOk() (*ListsSeriesModelV1, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *ListsSearchResponseV1Results) SetRecord(v ListsSeriesModelV1)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *ListsSearchResponseV1Results) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetMetadata

`func (o *ListsSearchResponseV1Results) GetMetadata() ListsSearchResponseV1ResultsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListsSearchResponseV1Results) GetMetadataOk() (*ListsSearchResponseV1ResultsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListsSearchResponseV1Results) SetMetadata(v ListsSearchResponseV1ResultsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListsSearchResponseV1Results) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


