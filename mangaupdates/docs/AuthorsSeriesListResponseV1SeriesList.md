# AuthorsSeriesListResponseV1SeriesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**SeriesId** | Pointer to **int64** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to [**AuthorsSeriesListResponseV1SeriesListMetadata**](AuthorsSeriesListResponseV1SeriesListMetadata.md) |  | [optional] 

## Methods

### NewAuthorsSeriesListResponseV1SeriesList

`func NewAuthorsSeriesListResponseV1SeriesList() *AuthorsSeriesListResponseV1SeriesList`

NewAuthorsSeriesListResponseV1SeriesList instantiates a new AuthorsSeriesListResponseV1SeriesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsSeriesListResponseV1SeriesListWithDefaults

`func NewAuthorsSeriesListResponseV1SeriesListWithDefaults() *AuthorsSeriesListResponseV1SeriesList`

NewAuthorsSeriesListResponseV1SeriesListWithDefaults instantiates a new AuthorsSeriesListResponseV1SeriesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AuthorsSeriesListResponseV1SeriesList) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthorsSeriesListResponseV1SeriesList) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthorsSeriesListResponseV1SeriesList) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSeriesId

`func (o *AuthorsSeriesListResponseV1SeriesList) GetSeriesId() int64`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetSeriesIdOk() (*int64, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *AuthorsSeriesListResponseV1SeriesList) SetSeriesId(v int64)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *AuthorsSeriesListResponseV1SeriesList) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetYear

`func (o *AuthorsSeriesListResponseV1SeriesList) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AuthorsSeriesListResponseV1SeriesList) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *AuthorsSeriesListResponseV1SeriesList) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorsSeriesListResponseV1SeriesList) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorsSeriesListResponseV1SeriesList) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorsSeriesListResponseV1SeriesList) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetGenres

`func (o *AuthorsSeriesListResponseV1SeriesList) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AuthorsSeriesListResponseV1SeriesList) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AuthorsSeriesListResponseV1SeriesList) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetMetadata

`func (o *AuthorsSeriesListResponseV1SeriesList) GetMetadata() AuthorsSeriesListResponseV1SeriesListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuthorsSeriesListResponseV1SeriesList) GetMetadataOk() (*AuthorsSeriesListResponseV1SeriesListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuthorsSeriesListResponseV1SeriesList) SetMetadata(v AuthorsSeriesListResponseV1SeriesListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuthorsSeriesListResponseV1SeriesList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


