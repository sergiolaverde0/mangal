# SeriesSearchResponseV1ResultsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserList** | Pointer to [**ListsSeriesModelV1**](ListsSeriesModelV1.md) |  | [optional] 
**UserGenreHighlights** | Pointer to [**[]SeriesSearchResponseV1ResultsMetadataUserGenreHighlights**](SeriesSearchResponseV1ResultsMetadataUserGenreHighlights.md) |  | [optional] 

## Methods

### NewSeriesSearchResponseV1ResultsMetadata

`func NewSeriesSearchResponseV1ResultsMetadata() *SeriesSearchResponseV1ResultsMetadata`

NewSeriesSearchResponseV1ResultsMetadata instantiates a new SeriesSearchResponseV1ResultsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesSearchResponseV1ResultsMetadataWithDefaults

`func NewSeriesSearchResponseV1ResultsMetadataWithDefaults() *SeriesSearchResponseV1ResultsMetadata`

NewSeriesSearchResponseV1ResultsMetadataWithDefaults instantiates a new SeriesSearchResponseV1ResultsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserList

`func (o *SeriesSearchResponseV1ResultsMetadata) GetUserList() ListsSeriesModelV1`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *SeriesSearchResponseV1ResultsMetadata) GetUserListOk() (*ListsSeriesModelV1, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *SeriesSearchResponseV1ResultsMetadata) SetUserList(v ListsSeriesModelV1)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *SeriesSearchResponseV1ResultsMetadata) HasUserList() bool`

HasUserList returns a boolean if a field has been set.

### GetUserGenreHighlights

`func (o *SeriesSearchResponseV1ResultsMetadata) GetUserGenreHighlights() []SeriesSearchResponseV1ResultsMetadataUserGenreHighlights`

GetUserGenreHighlights returns the UserGenreHighlights field if non-nil, zero value otherwise.

### GetUserGenreHighlightsOk

`func (o *SeriesSearchResponseV1ResultsMetadata) GetUserGenreHighlightsOk() (*[]SeriesSearchResponseV1ResultsMetadataUserGenreHighlights, bool)`

GetUserGenreHighlightsOk returns a tuple with the UserGenreHighlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGenreHighlights

`func (o *SeriesSearchResponseV1ResultsMetadata) SetUserGenreHighlights(v []SeriesSearchResponseV1ResultsMetadataUserGenreHighlights)`

SetUserGenreHighlights sets UserGenreHighlights field to given value.

### HasUserGenreHighlights

`func (o *SeriesSearchResponseV1ResultsMetadata) HasUserGenreHighlights() bool`

HasUserGenreHighlights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


