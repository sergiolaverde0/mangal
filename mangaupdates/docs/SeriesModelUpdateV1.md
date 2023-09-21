# SeriesModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]SeriesModelUpdateV1Associated**](SeriesModelUpdateV1Associated.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**Genres** | Pointer to [**[]SeriesModelUpdateV1Genres**](SeriesModelUpdateV1Genres.md) |  | [optional] 
**Categories** | Pointer to [**[]CategoriesModelUpdateV1**](CategoriesModelUpdateV1.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Anime** | Pointer to [**SeriesModelUpdateV1Anime**](SeriesModelUpdateV1Anime.md) |  | [optional] 
**RelatedSeries** | Pointer to [**[]SeriesModelUpdateV1RelatedSeries**](SeriesModelUpdateV1RelatedSeries.md) |  | [optional] 
**Authors** | Pointer to [**[]SeriesModelUpdateV1Authors**](SeriesModelUpdateV1Authors.md) |  | [optional] 
**Publishers** | Pointer to [**[]SeriesModelUpdateV1Publishers**](SeriesModelUpdateV1Publishers.md) |  | [optional] 
**Publications** | Pointer to [**[]SeriesModelUpdateV1Publications**](SeriesModelUpdateV1Publications.md) |  | [optional] 
**Admin** | Pointer to [**SeriesModelUpdateV1Admin**](SeriesModelUpdateV1Admin.md) |  | [optional] 

## Methods

### NewSeriesModelUpdateV1

`func NewSeriesModelUpdateV1() *SeriesModelUpdateV1`

NewSeriesModelUpdateV1 instantiates a new SeriesModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelUpdateV1WithDefaults

`func NewSeriesModelUpdateV1WithDefaults() *SeriesModelUpdateV1`

NewSeriesModelUpdateV1WithDefaults instantiates a new SeriesModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SeriesModelUpdateV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesModelUpdateV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesModelUpdateV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesModelUpdateV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssociated

`func (o *SeriesModelUpdateV1) GetAssociated() []SeriesModelUpdateV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *SeriesModelUpdateV1) GetAssociatedOk() (*[]SeriesModelUpdateV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *SeriesModelUpdateV1) SetAssociated(v []SeriesModelUpdateV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *SeriesModelUpdateV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetDescription

`func (o *SeriesModelUpdateV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeriesModelUpdateV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeriesModelUpdateV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeriesModelUpdateV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SeriesModelUpdateV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeriesModelUpdateV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeriesModelUpdateV1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SeriesModelUpdateV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYear

`func (o *SeriesModelUpdateV1) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeriesModelUpdateV1) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeriesModelUpdateV1) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeriesModelUpdateV1) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetGenres

`func (o *SeriesModelUpdateV1) GetGenres() []SeriesModelUpdateV1Genres`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SeriesModelUpdateV1) GetGenresOk() (*[]SeriesModelUpdateV1Genres, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SeriesModelUpdateV1) SetGenres(v []SeriesModelUpdateV1Genres)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SeriesModelUpdateV1) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCategories

`func (o *SeriesModelUpdateV1) GetCategories() []CategoriesModelUpdateV1`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SeriesModelUpdateV1) GetCategoriesOk() (*[]CategoriesModelUpdateV1, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SeriesModelUpdateV1) SetCategories(v []CategoriesModelUpdateV1)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SeriesModelUpdateV1) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetStatus

`func (o *SeriesModelUpdateV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SeriesModelUpdateV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SeriesModelUpdateV1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SeriesModelUpdateV1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLicensed

`func (o *SeriesModelUpdateV1) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *SeriesModelUpdateV1) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *SeriesModelUpdateV1) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *SeriesModelUpdateV1) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetCompleted

`func (o *SeriesModelUpdateV1) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *SeriesModelUpdateV1) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *SeriesModelUpdateV1) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *SeriesModelUpdateV1) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetAnime

`func (o *SeriesModelUpdateV1) GetAnime() SeriesModelUpdateV1Anime`

GetAnime returns the Anime field if non-nil, zero value otherwise.

### GetAnimeOk

`func (o *SeriesModelUpdateV1) GetAnimeOk() (*SeriesModelUpdateV1Anime, bool)`

GetAnimeOk returns a tuple with the Anime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnime

`func (o *SeriesModelUpdateV1) SetAnime(v SeriesModelUpdateV1Anime)`

SetAnime sets Anime field to given value.

### HasAnime

`func (o *SeriesModelUpdateV1) HasAnime() bool`

HasAnime returns a boolean if a field has been set.

### GetRelatedSeries

`func (o *SeriesModelUpdateV1) GetRelatedSeries() []SeriesModelUpdateV1RelatedSeries`

GetRelatedSeries returns the RelatedSeries field if non-nil, zero value otherwise.

### GetRelatedSeriesOk

`func (o *SeriesModelUpdateV1) GetRelatedSeriesOk() (*[]SeriesModelUpdateV1RelatedSeries, bool)`

GetRelatedSeriesOk returns a tuple with the RelatedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSeries

`func (o *SeriesModelUpdateV1) SetRelatedSeries(v []SeriesModelUpdateV1RelatedSeries)`

SetRelatedSeries sets RelatedSeries field to given value.

### HasRelatedSeries

`func (o *SeriesModelUpdateV1) HasRelatedSeries() bool`

HasRelatedSeries returns a boolean if a field has been set.

### GetAuthors

`func (o *SeriesModelUpdateV1) GetAuthors() []SeriesModelUpdateV1Authors`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *SeriesModelUpdateV1) GetAuthorsOk() (*[]SeriesModelUpdateV1Authors, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *SeriesModelUpdateV1) SetAuthors(v []SeriesModelUpdateV1Authors)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *SeriesModelUpdateV1) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetPublishers

`func (o *SeriesModelUpdateV1) GetPublishers() []SeriesModelUpdateV1Publishers`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *SeriesModelUpdateV1) GetPublishersOk() (*[]SeriesModelUpdateV1Publishers, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *SeriesModelUpdateV1) SetPublishers(v []SeriesModelUpdateV1Publishers)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *SeriesModelUpdateV1) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### GetPublications

`func (o *SeriesModelUpdateV1) GetPublications() []SeriesModelUpdateV1Publications`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *SeriesModelUpdateV1) GetPublicationsOk() (*[]SeriesModelUpdateV1Publications, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *SeriesModelUpdateV1) SetPublications(v []SeriesModelUpdateV1Publications)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *SeriesModelUpdateV1) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetAdmin

`func (o *SeriesModelUpdateV1) GetAdmin() SeriesModelUpdateV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SeriesModelUpdateV1) GetAdminOk() (*SeriesModelUpdateV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SeriesModelUpdateV1) SetAdmin(v SeriesModelUpdateV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SeriesModelUpdateV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


