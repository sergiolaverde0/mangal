# SeriesModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesId** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**ImageModelV1**](ImageModelV1.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**BayesianRating** | Pointer to **float32** |  | [optional] 
**RatingVotes** | Pointer to **int64** |  | [optional] 
**Genres** | Pointer to [**[]SeriesModelSearchV1Genres**](SeriesModelSearchV1Genres.md) |  | [optional] 
**LatestChapter** | Pointer to **int64** |  | [optional] 
**Rank** | Pointer to [**SeriesModelSearchV1Rank**](SeriesModelSearchV1Rank.md) |  | [optional] 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Admin** | Pointer to [**SeriesModelSearchV1Admin**](SeriesModelSearchV1Admin.md) |  | [optional] 

## Methods

### NewSeriesModelSearchV1

`func NewSeriesModelSearchV1() *SeriesModelSearchV1`

NewSeriesModelSearchV1 instantiates a new SeriesModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelSearchV1WithDefaults

`func NewSeriesModelSearchV1WithDefaults() *SeriesModelSearchV1`

NewSeriesModelSearchV1WithDefaults instantiates a new SeriesModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesId

`func (o *SeriesModelSearchV1) GetSeriesId() int64`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SeriesModelSearchV1) GetSeriesIdOk() (*int64, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SeriesModelSearchV1) SetSeriesId(v int64)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SeriesModelSearchV1) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetTitle

`func (o *SeriesModelSearchV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesModelSearchV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesModelSearchV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesModelSearchV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SeriesModelSearchV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SeriesModelSearchV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SeriesModelSearchV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SeriesModelSearchV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *SeriesModelSearchV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeriesModelSearchV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeriesModelSearchV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeriesModelSearchV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *SeriesModelSearchV1) GetImage() ImageModelV1`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SeriesModelSearchV1) GetImageOk() (*ImageModelV1, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SeriesModelSearchV1) SetImage(v ImageModelV1)`

SetImage sets Image field to given value.

### HasImage

`func (o *SeriesModelSearchV1) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *SeriesModelSearchV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeriesModelSearchV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeriesModelSearchV1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SeriesModelSearchV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYear

`func (o *SeriesModelSearchV1) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeriesModelSearchV1) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeriesModelSearchV1) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeriesModelSearchV1) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetBayesianRating

`func (o *SeriesModelSearchV1) GetBayesianRating() float32`

GetBayesianRating returns the BayesianRating field if non-nil, zero value otherwise.

### GetBayesianRatingOk

`func (o *SeriesModelSearchV1) GetBayesianRatingOk() (*float32, bool)`

GetBayesianRatingOk returns a tuple with the BayesianRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBayesianRating

`func (o *SeriesModelSearchV1) SetBayesianRating(v float32)`

SetBayesianRating sets BayesianRating field to given value.

### HasBayesianRating

`func (o *SeriesModelSearchV1) HasBayesianRating() bool`

HasBayesianRating returns a boolean if a field has been set.

### GetRatingVotes

`func (o *SeriesModelSearchV1) GetRatingVotes() int64`

GetRatingVotes returns the RatingVotes field if non-nil, zero value otherwise.

### GetRatingVotesOk

`func (o *SeriesModelSearchV1) GetRatingVotesOk() (*int64, bool)`

GetRatingVotesOk returns a tuple with the RatingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingVotes

`func (o *SeriesModelSearchV1) SetRatingVotes(v int64)`

SetRatingVotes sets RatingVotes field to given value.

### HasRatingVotes

`func (o *SeriesModelSearchV1) HasRatingVotes() bool`

HasRatingVotes returns a boolean if a field has been set.

### GetGenres

`func (o *SeriesModelSearchV1) GetGenres() []SeriesModelSearchV1Genres`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SeriesModelSearchV1) GetGenresOk() (*[]SeriesModelSearchV1Genres, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SeriesModelSearchV1) SetGenres(v []SeriesModelSearchV1Genres)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SeriesModelSearchV1) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetLatestChapter

`func (o *SeriesModelSearchV1) GetLatestChapter() int64`

GetLatestChapter returns the LatestChapter field if non-nil, zero value otherwise.

### GetLatestChapterOk

`func (o *SeriesModelSearchV1) GetLatestChapterOk() (*int64, bool)`

GetLatestChapterOk returns a tuple with the LatestChapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestChapter

`func (o *SeriesModelSearchV1) SetLatestChapter(v int64)`

SetLatestChapter sets LatestChapter field to given value.

### HasLatestChapter

`func (o *SeriesModelSearchV1) HasLatestChapter() bool`

HasLatestChapter returns a boolean if a field has been set.

### GetRank

`func (o *SeriesModelSearchV1) GetRank() SeriesModelSearchV1Rank`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SeriesModelSearchV1) GetRankOk() (*SeriesModelSearchV1Rank, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SeriesModelSearchV1) SetRank(v SeriesModelSearchV1Rank)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SeriesModelSearchV1) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SeriesModelSearchV1) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SeriesModelSearchV1) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SeriesModelSearchV1) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SeriesModelSearchV1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAdmin

`func (o *SeriesModelSearchV1) GetAdmin() SeriesModelSearchV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SeriesModelSearchV1) GetAdminOk() (*SeriesModelSearchV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SeriesModelSearchV1) SetAdmin(v SeriesModelSearchV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SeriesModelSearchV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


