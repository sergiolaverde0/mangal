# SeriesSearchRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to **int64** |  | [optional] 
**Stype** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **[]string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**FilterTypes** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **[]string** | Meant to replace &#39;filter&#39;, it lets you specify multiple filters as an array of strings | [optional] 
**List** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Perpage** | Pointer to **int64** |  | [optional] 
**Letter** | Pointer to **string** |  | [optional] 
**Genre** | Pointer to **[]string** |  | [optional] 
**ExcludeGenre** | Pointer to **[]string** |  | [optional] 
**Orderby** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 
**IncludeRankMetadata** | Pointer to **bool** |  | [optional] 
**ExcludeFilteredGenres** | Pointer to **bool** |  | [optional] 

## Methods

### NewSeriesSearchRequestV1

`func NewSeriesSearchRequestV1() *SeriesSearchRequestV1`

NewSeriesSearchRequestV1 instantiates a new SeriesSearchRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesSearchRequestV1WithDefaults

`func NewSeriesSearchRequestV1WithDefaults() *SeriesSearchRequestV1`

NewSeriesSearchRequestV1WithDefaults instantiates a new SeriesSearchRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *SeriesSearchRequestV1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *SeriesSearchRequestV1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *SeriesSearchRequestV1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *SeriesSearchRequestV1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetAddedBy

`func (o *SeriesSearchRequestV1) GetAddedBy() int64`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *SeriesSearchRequestV1) GetAddedByOk() (*int64, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *SeriesSearchRequestV1) SetAddedBy(v int64)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *SeriesSearchRequestV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetStype

`func (o *SeriesSearchRequestV1) GetStype() string`

GetStype returns the Stype field if non-nil, zero value otherwise.

### GetStypeOk

`func (o *SeriesSearchRequestV1) GetStypeOk() (*string, bool)`

GetStypeOk returns a tuple with the Stype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStype

`func (o *SeriesSearchRequestV1) SetStype(v string)`

SetStype sets Stype field to given value.

### HasStype

`func (o *SeriesSearchRequestV1) HasStype() bool`

HasStype returns a boolean if a field has been set.

### GetLicensed

`func (o *SeriesSearchRequestV1) GetLicensed() string`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *SeriesSearchRequestV1) GetLicensedOk() (*string, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *SeriesSearchRequestV1) SetLicensed(v string)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *SeriesSearchRequestV1) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetType

`func (o *SeriesSearchRequestV1) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeriesSearchRequestV1) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeriesSearchRequestV1) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *SeriesSearchRequestV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYear

`func (o *SeriesSearchRequestV1) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeriesSearchRequestV1) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeriesSearchRequestV1) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeriesSearchRequestV1) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetFilterTypes

`func (o *SeriesSearchRequestV1) GetFilterTypes() []string`

GetFilterTypes returns the FilterTypes field if non-nil, zero value otherwise.

### GetFilterTypesOk

`func (o *SeriesSearchRequestV1) GetFilterTypesOk() (*[]string, bool)`

GetFilterTypesOk returns a tuple with the FilterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTypes

`func (o *SeriesSearchRequestV1) SetFilterTypes(v []string)`

SetFilterTypes sets FilterTypes field to given value.

### HasFilterTypes

`func (o *SeriesSearchRequestV1) HasFilterTypes() bool`

HasFilterTypes returns a boolean if a field has been set.

### GetCategory

`func (o *SeriesSearchRequestV1) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SeriesSearchRequestV1) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SeriesSearchRequestV1) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SeriesSearchRequestV1) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFilter

`func (o *SeriesSearchRequestV1) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SeriesSearchRequestV1) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SeriesSearchRequestV1) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SeriesSearchRequestV1) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilters

`func (o *SeriesSearchRequestV1) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SeriesSearchRequestV1) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SeriesSearchRequestV1) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SeriesSearchRequestV1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetList

`func (o *SeriesSearchRequestV1) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *SeriesSearchRequestV1) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *SeriesSearchRequestV1) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *SeriesSearchRequestV1) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *SeriesSearchRequestV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SeriesSearchRequestV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SeriesSearchRequestV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *SeriesSearchRequestV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerpage

`func (o *SeriesSearchRequestV1) GetPerpage() int64`

GetPerpage returns the Perpage field if non-nil, zero value otherwise.

### GetPerpageOk

`func (o *SeriesSearchRequestV1) GetPerpageOk() (*int64, bool)`

GetPerpageOk returns a tuple with the Perpage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerpage

`func (o *SeriesSearchRequestV1) SetPerpage(v int64)`

SetPerpage sets Perpage field to given value.

### HasPerpage

`func (o *SeriesSearchRequestV1) HasPerpage() bool`

HasPerpage returns a boolean if a field has been set.

### GetLetter

`func (o *SeriesSearchRequestV1) GetLetter() string`

GetLetter returns the Letter field if non-nil, zero value otherwise.

### GetLetterOk

`func (o *SeriesSearchRequestV1) GetLetterOk() (*string, bool)`

GetLetterOk returns a tuple with the Letter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetter

`func (o *SeriesSearchRequestV1) SetLetter(v string)`

SetLetter sets Letter field to given value.

### HasLetter

`func (o *SeriesSearchRequestV1) HasLetter() bool`

HasLetter returns a boolean if a field has been set.

### GetGenre

`func (o *SeriesSearchRequestV1) GetGenre() []string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *SeriesSearchRequestV1) GetGenreOk() (*[]string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *SeriesSearchRequestV1) SetGenre(v []string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *SeriesSearchRequestV1) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetExcludeGenre

`func (o *SeriesSearchRequestV1) GetExcludeGenre() []string`

GetExcludeGenre returns the ExcludeGenre field if non-nil, zero value otherwise.

### GetExcludeGenreOk

`func (o *SeriesSearchRequestV1) GetExcludeGenreOk() (*[]string, bool)`

GetExcludeGenreOk returns a tuple with the ExcludeGenre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGenre

`func (o *SeriesSearchRequestV1) SetExcludeGenre(v []string)`

SetExcludeGenre sets ExcludeGenre field to given value.

### HasExcludeGenre

`func (o *SeriesSearchRequestV1) HasExcludeGenre() bool`

HasExcludeGenre returns a boolean if a field has been set.

### GetOrderby

`func (o *SeriesSearchRequestV1) GetOrderby() string`

GetOrderby returns the Orderby field if non-nil, zero value otherwise.

### GetOrderbyOk

`func (o *SeriesSearchRequestV1) GetOrderbyOk() (*string, bool)`

GetOrderbyOk returns a tuple with the Orderby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderby

`func (o *SeriesSearchRequestV1) SetOrderby(v string)`

SetOrderby sets Orderby field to given value.

### HasOrderby

`func (o *SeriesSearchRequestV1) HasOrderby() bool`

HasOrderby returns a boolean if a field has been set.

### GetPending

`func (o *SeriesSearchRequestV1) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *SeriesSearchRequestV1) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *SeriesSearchRequestV1) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *SeriesSearchRequestV1) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetIncludeRankMetadata

`func (o *SeriesSearchRequestV1) GetIncludeRankMetadata() bool`

GetIncludeRankMetadata returns the IncludeRankMetadata field if non-nil, zero value otherwise.

### GetIncludeRankMetadataOk

`func (o *SeriesSearchRequestV1) GetIncludeRankMetadataOk() (*bool, bool)`

GetIncludeRankMetadataOk returns a tuple with the IncludeRankMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRankMetadata

`func (o *SeriesSearchRequestV1) SetIncludeRankMetadata(v bool)`

SetIncludeRankMetadata sets IncludeRankMetadata field to given value.

### HasIncludeRankMetadata

`func (o *SeriesSearchRequestV1) HasIncludeRankMetadata() bool`

HasIncludeRankMetadata returns a boolean if a field has been set.

### GetExcludeFilteredGenres

`func (o *SeriesSearchRequestV1) GetExcludeFilteredGenres() bool`

GetExcludeFilteredGenres returns the ExcludeFilteredGenres field if non-nil, zero value otherwise.

### GetExcludeFilteredGenresOk

`func (o *SeriesSearchRequestV1) GetExcludeFilteredGenresOk() (*bool, bool)`

GetExcludeFilteredGenresOk returns a tuple with the ExcludeFilteredGenres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilteredGenres

`func (o *SeriesSearchRequestV1) SetExcludeFilteredGenres(v bool)`

SetExcludeFilteredGenres sets ExcludeFilteredGenres field to given value.

### HasExcludeFilteredGenres

`func (o *SeriesSearchRequestV1) HasExcludeFilteredGenres() bool`

HasExcludeFilteredGenres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


