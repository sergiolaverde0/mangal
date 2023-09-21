# AuthorsSeriesListResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSeries** | Pointer to **int64** |  | [optional] 
**SeriesList** | Pointer to [**[]AuthorsSeriesListResponseV1SeriesList**](AuthorsSeriesListResponseV1SeriesList.md) |  | [optional] 
**GenreList** | Pointer to [**[]AuthorsSeriesListResponseV1GenreList**](AuthorsSeriesListResponseV1GenreList.md) |  | [optional] 

## Methods

### NewAuthorsSeriesListResponseV1

`func NewAuthorsSeriesListResponseV1() *AuthorsSeriesListResponseV1`

NewAuthorsSeriesListResponseV1 instantiates a new AuthorsSeriesListResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsSeriesListResponseV1WithDefaults

`func NewAuthorsSeriesListResponseV1WithDefaults() *AuthorsSeriesListResponseV1`

NewAuthorsSeriesListResponseV1WithDefaults instantiates a new AuthorsSeriesListResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSeries

`func (o *AuthorsSeriesListResponseV1) GetTotalSeries() int64`

GetTotalSeries returns the TotalSeries field if non-nil, zero value otherwise.

### GetTotalSeriesOk

`func (o *AuthorsSeriesListResponseV1) GetTotalSeriesOk() (*int64, bool)`

GetTotalSeriesOk returns a tuple with the TotalSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeries

`func (o *AuthorsSeriesListResponseV1) SetTotalSeries(v int64)`

SetTotalSeries sets TotalSeries field to given value.

### HasTotalSeries

`func (o *AuthorsSeriesListResponseV1) HasTotalSeries() bool`

HasTotalSeries returns a boolean if a field has been set.

### GetSeriesList

`func (o *AuthorsSeriesListResponseV1) GetSeriesList() []AuthorsSeriesListResponseV1SeriesList`

GetSeriesList returns the SeriesList field if non-nil, zero value otherwise.

### GetSeriesListOk

`func (o *AuthorsSeriesListResponseV1) GetSeriesListOk() (*[]AuthorsSeriesListResponseV1SeriesList, bool)`

GetSeriesListOk returns a tuple with the SeriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesList

`func (o *AuthorsSeriesListResponseV1) SetSeriesList(v []AuthorsSeriesListResponseV1SeriesList)`

SetSeriesList sets SeriesList field to given value.

### HasSeriesList

`func (o *AuthorsSeriesListResponseV1) HasSeriesList() bool`

HasSeriesList returns a boolean if a field has been set.

### GetGenreList

`func (o *AuthorsSeriesListResponseV1) GetGenreList() []AuthorsSeriesListResponseV1GenreList`

GetGenreList returns the GenreList field if non-nil, zero value otherwise.

### GetGenreListOk

`func (o *AuthorsSeriesListResponseV1) GetGenreListOk() (*[]AuthorsSeriesListResponseV1GenreList, bool)`

GetGenreListOk returns a tuple with the GenreList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreList

`func (o *AuthorsSeriesListResponseV1) SetGenreList(v []AuthorsSeriesListResponseV1GenreList)`

SetGenreList sets GenreList field to given value.

### HasGenreList

`func (o *AuthorsSeriesListResponseV1) HasGenreList() bool`

HasGenreList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


