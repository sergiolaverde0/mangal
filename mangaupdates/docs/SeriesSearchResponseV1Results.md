# SeriesSearchResponseV1Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**HitTitle** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**SeriesSearchResponseV1ResultsMetadata**](SeriesSearchResponseV1ResultsMetadata.md) |  | [optional] 

## Methods

### NewSeriesSearchResponseV1Results

`func NewSeriesSearchResponseV1Results() *SeriesSearchResponseV1Results`

NewSeriesSearchResponseV1Results instantiates a new SeriesSearchResponseV1Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesSearchResponseV1ResultsWithDefaults

`func NewSeriesSearchResponseV1ResultsWithDefaults() *SeriesSearchResponseV1Results`

NewSeriesSearchResponseV1ResultsWithDefaults instantiates a new SeriesSearchResponseV1Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *SeriesSearchResponseV1Results) GetRecord() SeriesModelSearchV1`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *SeriesSearchResponseV1Results) GetRecordOk() (*SeriesModelSearchV1, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *SeriesSearchResponseV1Results) SetRecord(v SeriesModelSearchV1)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *SeriesSearchResponseV1Results) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetHitTitle

`func (o *SeriesSearchResponseV1Results) GetHitTitle() string`

GetHitTitle returns the HitTitle field if non-nil, zero value otherwise.

### GetHitTitleOk

`func (o *SeriesSearchResponseV1Results) GetHitTitleOk() (*string, bool)`

GetHitTitleOk returns a tuple with the HitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitTitle

`func (o *SeriesSearchResponseV1Results) SetHitTitle(v string)`

SetHitTitle sets HitTitle field to given value.

### HasHitTitle

`func (o *SeriesSearchResponseV1Results) HasHitTitle() bool`

HasHitTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *SeriesSearchResponseV1Results) GetMetadata() SeriesSearchResponseV1ResultsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SeriesSearchResponseV1Results) GetMetadataOk() (*SeriesSearchResponseV1ResultsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SeriesSearchResponseV1Results) SetMetadata(v SeriesSearchResponseV1ResultsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SeriesSearchResponseV1Results) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


