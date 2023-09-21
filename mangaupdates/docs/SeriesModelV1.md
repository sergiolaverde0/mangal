# SeriesModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeriesId** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]SeriesModelV1Associated**](SeriesModelV1Associated.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**ImageModelV1**](ImageModelV1.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **string** |  | [optional] 
**BayesianRating** | Pointer to **float32** |  | [optional] 
**RatingVotes** | Pointer to **int64** |  | [optional] 
**Genres** | Pointer to [**[]SeriesModelV1Genres**](SeriesModelV1Genres.md) |  | [optional] 
**Categories** | Pointer to [**[]CategoriesModelV1**](CategoriesModelV1.md) |  | [optional] 
**LatestChapter** | Pointer to **int64** |  | [optional] 
**ForumId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Licensed** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Anime** | Pointer to [**SeriesModelV1Anime**](SeriesModelV1Anime.md) |  | [optional] 
**RelatedSeries** | Pointer to [**[]SeriesModelV1RelatedSeries**](SeriesModelV1RelatedSeries.md) |  | [optional] 
**Authors** | Pointer to [**[]SeriesModelV1Authors**](SeriesModelV1Authors.md) |  | [optional] 
**Publishers** | Pointer to [**[]SeriesModelV1Publishers**](SeriesModelV1Publishers.md) |  | [optional] 
**Publications** | Pointer to [**[]SeriesModelV1Publications**](SeriesModelV1Publications.md) |  | [optional] 
**Recommendations** | Pointer to [**[]SeriesRecommendationsModelV1**](SeriesRecommendationsModelV1.md) |  | [optional] 
**CategoryRecommendations** | Pointer to [**[]SeriesRecommendationsModelV1**](SeriesRecommendationsModelV1.md) |  | [optional] 
**Rank** | Pointer to [**SeriesModelV1Rank**](SeriesModelV1Rank.md) |  | [optional] 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Admin** | Pointer to [**SeriesModelV1Admin**](SeriesModelV1Admin.md) |  | [optional] 

## Methods

### NewSeriesModelV1

`func NewSeriesModelV1() *SeriesModelV1`

NewSeriesModelV1 instantiates a new SeriesModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesModelV1WithDefaults

`func NewSeriesModelV1WithDefaults() *SeriesModelV1`

NewSeriesModelV1WithDefaults instantiates a new SeriesModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeriesId

`func (o *SeriesModelV1) GetSeriesId() int64`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SeriesModelV1) GetSeriesIdOk() (*int64, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SeriesModelV1) SetSeriesId(v int64)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SeriesModelV1) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetTitle

`func (o *SeriesModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SeriesModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SeriesModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SeriesModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SeriesModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAssociated

`func (o *SeriesModelV1) GetAssociated() []SeriesModelV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *SeriesModelV1) GetAssociatedOk() (*[]SeriesModelV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *SeriesModelV1) SetAssociated(v []SeriesModelV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *SeriesModelV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetDescription

`func (o *SeriesModelV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeriesModelV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeriesModelV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeriesModelV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *SeriesModelV1) GetImage() ImageModelV1`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SeriesModelV1) GetImageOk() (*ImageModelV1, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SeriesModelV1) SetImage(v ImageModelV1)`

SetImage sets Image field to given value.

### HasImage

`func (o *SeriesModelV1) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *SeriesModelV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SeriesModelV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SeriesModelV1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SeriesModelV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYear

`func (o *SeriesModelV1) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SeriesModelV1) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SeriesModelV1) SetYear(v string)`

SetYear sets Year field to given value.

### HasYear

`func (o *SeriesModelV1) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetBayesianRating

`func (o *SeriesModelV1) GetBayesianRating() float32`

GetBayesianRating returns the BayesianRating field if non-nil, zero value otherwise.

### GetBayesianRatingOk

`func (o *SeriesModelV1) GetBayesianRatingOk() (*float32, bool)`

GetBayesianRatingOk returns a tuple with the BayesianRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBayesianRating

`func (o *SeriesModelV1) SetBayesianRating(v float32)`

SetBayesianRating sets BayesianRating field to given value.

### HasBayesianRating

`func (o *SeriesModelV1) HasBayesianRating() bool`

HasBayesianRating returns a boolean if a field has been set.

### GetRatingVotes

`func (o *SeriesModelV1) GetRatingVotes() int64`

GetRatingVotes returns the RatingVotes field if non-nil, zero value otherwise.

### GetRatingVotesOk

`func (o *SeriesModelV1) GetRatingVotesOk() (*int64, bool)`

GetRatingVotesOk returns a tuple with the RatingVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingVotes

`func (o *SeriesModelV1) SetRatingVotes(v int64)`

SetRatingVotes sets RatingVotes field to given value.

### HasRatingVotes

`func (o *SeriesModelV1) HasRatingVotes() bool`

HasRatingVotes returns a boolean if a field has been set.

### GetGenres

`func (o *SeriesModelV1) GetGenres() []SeriesModelV1Genres`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SeriesModelV1) GetGenresOk() (*[]SeriesModelV1Genres, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SeriesModelV1) SetGenres(v []SeriesModelV1Genres)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SeriesModelV1) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCategories

`func (o *SeriesModelV1) GetCategories() []CategoriesModelV1`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SeriesModelV1) GetCategoriesOk() (*[]CategoriesModelV1, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SeriesModelV1) SetCategories(v []CategoriesModelV1)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SeriesModelV1) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetLatestChapter

`func (o *SeriesModelV1) GetLatestChapter() int64`

GetLatestChapter returns the LatestChapter field if non-nil, zero value otherwise.

### GetLatestChapterOk

`func (o *SeriesModelV1) GetLatestChapterOk() (*int64, bool)`

GetLatestChapterOk returns a tuple with the LatestChapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestChapter

`func (o *SeriesModelV1) SetLatestChapter(v int64)`

SetLatestChapter sets LatestChapter field to given value.

### HasLatestChapter

`func (o *SeriesModelV1) HasLatestChapter() bool`

HasLatestChapter returns a boolean if a field has been set.

### GetForumId

`func (o *SeriesModelV1) GetForumId() int64`

GetForumId returns the ForumId field if non-nil, zero value otherwise.

### GetForumIdOk

`func (o *SeriesModelV1) GetForumIdOk() (*int64, bool)`

GetForumIdOk returns a tuple with the ForumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumId

`func (o *SeriesModelV1) SetForumId(v int64)`

SetForumId sets ForumId field to given value.

### HasForumId

`func (o *SeriesModelV1) HasForumId() bool`

HasForumId returns a boolean if a field has been set.

### GetStatus

`func (o *SeriesModelV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SeriesModelV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SeriesModelV1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SeriesModelV1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLicensed

`func (o *SeriesModelV1) GetLicensed() bool`

GetLicensed returns the Licensed field if non-nil, zero value otherwise.

### GetLicensedOk

`func (o *SeriesModelV1) GetLicensedOk() (*bool, bool)`

GetLicensedOk returns a tuple with the Licensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensed

`func (o *SeriesModelV1) SetLicensed(v bool)`

SetLicensed sets Licensed field to given value.

### HasLicensed

`func (o *SeriesModelV1) HasLicensed() bool`

HasLicensed returns a boolean if a field has been set.

### GetCompleted

`func (o *SeriesModelV1) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *SeriesModelV1) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *SeriesModelV1) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *SeriesModelV1) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetAnime

`func (o *SeriesModelV1) GetAnime() SeriesModelV1Anime`

GetAnime returns the Anime field if non-nil, zero value otherwise.

### GetAnimeOk

`func (o *SeriesModelV1) GetAnimeOk() (*SeriesModelV1Anime, bool)`

GetAnimeOk returns a tuple with the Anime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnime

`func (o *SeriesModelV1) SetAnime(v SeriesModelV1Anime)`

SetAnime sets Anime field to given value.

### HasAnime

`func (o *SeriesModelV1) HasAnime() bool`

HasAnime returns a boolean if a field has been set.

### GetRelatedSeries

`func (o *SeriesModelV1) GetRelatedSeries() []SeriesModelV1RelatedSeries`

GetRelatedSeries returns the RelatedSeries field if non-nil, zero value otherwise.

### GetRelatedSeriesOk

`func (o *SeriesModelV1) GetRelatedSeriesOk() (*[]SeriesModelV1RelatedSeries, bool)`

GetRelatedSeriesOk returns a tuple with the RelatedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedSeries

`func (o *SeriesModelV1) SetRelatedSeries(v []SeriesModelV1RelatedSeries)`

SetRelatedSeries sets RelatedSeries field to given value.

### HasRelatedSeries

`func (o *SeriesModelV1) HasRelatedSeries() bool`

HasRelatedSeries returns a boolean if a field has been set.

### GetAuthors

`func (o *SeriesModelV1) GetAuthors() []SeriesModelV1Authors`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *SeriesModelV1) GetAuthorsOk() (*[]SeriesModelV1Authors, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *SeriesModelV1) SetAuthors(v []SeriesModelV1Authors)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *SeriesModelV1) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetPublishers

`func (o *SeriesModelV1) GetPublishers() []SeriesModelV1Publishers`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *SeriesModelV1) GetPublishersOk() (*[]SeriesModelV1Publishers, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *SeriesModelV1) SetPublishers(v []SeriesModelV1Publishers)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *SeriesModelV1) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### GetPublications

`func (o *SeriesModelV1) GetPublications() []SeriesModelV1Publications`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *SeriesModelV1) GetPublicationsOk() (*[]SeriesModelV1Publications, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *SeriesModelV1) SetPublications(v []SeriesModelV1Publications)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *SeriesModelV1) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetRecommendations

`func (o *SeriesModelV1) GetRecommendations() []SeriesRecommendationsModelV1`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *SeriesModelV1) GetRecommendationsOk() (*[]SeriesRecommendationsModelV1, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *SeriesModelV1) SetRecommendations(v []SeriesRecommendationsModelV1)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *SeriesModelV1) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetCategoryRecommendations

`func (o *SeriesModelV1) GetCategoryRecommendations() []SeriesRecommendationsModelV1`

GetCategoryRecommendations returns the CategoryRecommendations field if non-nil, zero value otherwise.

### GetCategoryRecommendationsOk

`func (o *SeriesModelV1) GetCategoryRecommendationsOk() (*[]SeriesRecommendationsModelV1, bool)`

GetCategoryRecommendationsOk returns a tuple with the CategoryRecommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryRecommendations

`func (o *SeriesModelV1) SetCategoryRecommendations(v []SeriesRecommendationsModelV1)`

SetCategoryRecommendations sets CategoryRecommendations field to given value.

### HasCategoryRecommendations

`func (o *SeriesModelV1) HasCategoryRecommendations() bool`

HasCategoryRecommendations returns a boolean if a field has been set.

### GetRank

`func (o *SeriesModelV1) GetRank() SeriesModelV1Rank`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SeriesModelV1) GetRankOk() (*SeriesModelV1Rank, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SeriesModelV1) SetRank(v SeriesModelV1Rank)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SeriesModelV1) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SeriesModelV1) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SeriesModelV1) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SeriesModelV1) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SeriesModelV1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAdmin

`func (o *SeriesModelV1) GetAdmin() SeriesModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SeriesModelV1) GetAdminOk() (*SeriesModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SeriesModelV1) SetAdmin(v SeriesModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SeriesModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


