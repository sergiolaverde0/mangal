# ReleaseSearchResponseV1ResultsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**UserList** | Pointer to [**ListsSeriesModelV1**](ListsSeriesModelV1.md) |  | [optional] 
**UserGenreHighlights** | Pointer to [**[]ReleaseSearchResponseV1ResultsMetadataUserGenreHighlights**](ReleaseSearchResponseV1ResultsMetadataUserGenreHighlights.md) |  | [optional] 
**UserGenreFilters** | Pointer to **[]string** |  | [optional] 
**UserGroupFilters** | Pointer to **[]string** |  | [optional] 
**TypeFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewReleaseSearchResponseV1ResultsMetadata

`func NewReleaseSearchResponseV1ResultsMetadata() *ReleaseSearchResponseV1ResultsMetadata`

NewReleaseSearchResponseV1ResultsMetadata instantiates a new ReleaseSearchResponseV1ResultsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseSearchResponseV1ResultsMetadataWithDefaults

`func NewReleaseSearchResponseV1ResultsMetadataWithDefaults() *ReleaseSearchResponseV1ResultsMetadata`

NewReleaseSearchResponseV1ResultsMetadataWithDefaults instantiates a new ReleaseSearchResponseV1ResultsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetSeries() SeriesModelSearchV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetSeriesOk() (*SeriesModelSearchV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetSeries(v SeriesModelSearchV1)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetUserList

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserList() ListsSeriesModelV1`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserListOk() (*ListsSeriesModelV1, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetUserList(v ListsSeriesModelV1)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasUserList() bool`

HasUserList returns a boolean if a field has been set.

### GetUserGenreHighlights

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGenreHighlights() []ReleaseSearchResponseV1ResultsMetadataUserGenreHighlights`

GetUserGenreHighlights returns the UserGenreHighlights field if non-nil, zero value otherwise.

### GetUserGenreHighlightsOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGenreHighlightsOk() (*[]ReleaseSearchResponseV1ResultsMetadataUserGenreHighlights, bool)`

GetUserGenreHighlightsOk returns a tuple with the UserGenreHighlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGenreHighlights

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetUserGenreHighlights(v []ReleaseSearchResponseV1ResultsMetadataUserGenreHighlights)`

SetUserGenreHighlights sets UserGenreHighlights field to given value.

### HasUserGenreHighlights

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasUserGenreHighlights() bool`

HasUserGenreHighlights returns a boolean if a field has been set.

### GetUserGenreFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGenreFilters() []string`

GetUserGenreFilters returns the UserGenreFilters field if non-nil, zero value otherwise.

### GetUserGenreFiltersOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGenreFiltersOk() (*[]string, bool)`

GetUserGenreFiltersOk returns a tuple with the UserGenreFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGenreFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetUserGenreFilters(v []string)`

SetUserGenreFilters sets UserGenreFilters field to given value.

### HasUserGenreFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasUserGenreFilters() bool`

HasUserGenreFilters returns a boolean if a field has been set.

### GetUserGroupFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGroupFilters() []string`

GetUserGroupFilters returns the UserGroupFilters field if non-nil, zero value otherwise.

### GetUserGroupFiltersOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetUserGroupFiltersOk() (*[]string, bool)`

GetUserGroupFiltersOk returns a tuple with the UserGroupFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetUserGroupFilters(v []string)`

SetUserGroupFilters sets UserGroupFilters field to given value.

### HasUserGroupFilters

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasUserGroupFilters() bool`

HasUserGroupFilters returns a boolean if a field has been set.

### GetTypeFilter

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetTypeFilter() string`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *ReleaseSearchResponseV1ResultsMetadata) GetTypeFilterOk() (*string, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *ReleaseSearchResponseV1ResultsMetadata) SetTypeFilter(v string)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *ReleaseSearchResponseV1ResultsMetadata) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


