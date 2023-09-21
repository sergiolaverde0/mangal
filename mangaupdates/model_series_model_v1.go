/*
MangaUpdates API

This API powers our website. Most API functions are public and do not require an account. For user-based functions, you must register an account. Currently, user registration must be done through our main website and is disabled via this API.  Download OpenAPI Specification: [openapi.yaml](openapi.yaml)  Please contact us at the following emails if you have questions:  * lambchopsil@mangaupdates.com * manick@mangaupdates.com  ## Warranties  MangaUpdates makes no warranties about service availability, correctness of the data, or anything else. The service is provided as is, and may change at any time.  ## Acceptable Use Policy  * You will credit MangaUpdates when using data provided by this API. * You will use reasonable spacing between requests so as not to overwhelm the MangaUpdates servers, and employ caching mechanisms when accessing data. * You will NOT use MangaUpdates data or API in a way that will:     * Deceive or defraud users     * Assist or perform an illegal action     * Create spam     * Damage the database  We reserve the right to change this policy at any time.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SeriesModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelV1{}

// SeriesModelV1 struct for SeriesModelV1
type SeriesModelV1 struct {
	SeriesId                *int64                         `json:"series_id,omitempty"`
	Title                   *string                        `json:"title,omitempty"`
	Url                     *string                        `json:"url,omitempty"`
	Associated              []SeriesModelV1Associated      `json:"associated,omitempty"`
	Description             *string                        `json:"description,omitempty"`
	Image                   *ImageModelV1                  `json:"image,omitempty"`
	Type                    *string                        `json:"type,omitempty"`
	Year                    *string                        `json:"year,omitempty"`
	BayesianRating          *float32                       `json:"bayesian_rating,omitempty"`
	RatingVotes             *int64                         `json:"rating_votes,omitempty"`
	Genres                  []SeriesModelV1Genres          `json:"genres,omitempty"`
	Categories              []CategoriesModelV1            `json:"categories,omitempty"`
	LatestChapter           *int64                         `json:"latest_chapter,omitempty"`
	ForumId                 *int64                         `json:"forum_id,omitempty"`
	Status                  *string                        `json:"status,omitempty"`
	Licensed                *bool                          `json:"licensed,omitempty"`
	Completed               *bool                          `json:"completed,omitempty"`
	Anime                   *SeriesModelV1Anime            `json:"anime,omitempty"`
	RelatedSeries           []SeriesModelV1RelatedSeries   `json:"related_series,omitempty"`
	Authors                 []SeriesModelV1Authors         `json:"authors,omitempty"`
	Publishers              []SeriesModelV1Publishers      `json:"publishers,omitempty"`
	Publications            []SeriesModelV1Publications    `json:"publications,omitempty"`
	Recommendations         []SeriesRecommendationsModelV1 `json:"recommendations,omitempty"`
	CategoryRecommendations []SeriesRecommendationsModelV1 `json:"category_recommendations,omitempty"`
	Rank                    *SeriesModelV1Rank             `json:"rank,omitempty"`
	LastUpdated             *TimeV1                        `json:"last_updated,omitempty"`
	Admin                   *SeriesModelV1Admin            `json:"admin,omitempty"`
}

// NewSeriesModelV1 instantiates a new SeriesModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelV1() *SeriesModelV1 {
	this := SeriesModelV1{}
	return &this
}

// NewSeriesModelV1WithDefaults instantiates a new SeriesModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelV1WithDefaults() *SeriesModelV1 {
	this := SeriesModelV1{}
	return &this
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *SeriesModelV1) GetSeriesId() int64 {
	if o == nil || IsNil(o.SeriesId) {
		var ret int64
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetSeriesIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SeriesId) {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *SeriesModelV1) HasSeriesId() bool {
	if o != nil && !IsNil(o.SeriesId) {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given int64 and assigns it to the SeriesId field.
func (o *SeriesModelV1) SetSeriesId(v int64) {
	o.SeriesId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SeriesModelV1) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SeriesModelV1) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SeriesModelV1) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SeriesModelV1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SeriesModelV1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SeriesModelV1) SetUrl(v string) {
	o.Url = &v
}

// GetAssociated returns the Associated field value if set, zero value otherwise.
func (o *SeriesModelV1) GetAssociated() []SeriesModelV1Associated {
	if o == nil || IsNil(o.Associated) {
		var ret []SeriesModelV1Associated
		return ret
	}
	return o.Associated
}

// GetAssociatedOk returns a tuple with the Associated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetAssociatedOk() ([]SeriesModelV1Associated, bool) {
	if o == nil || IsNil(o.Associated) {
		return nil, false
	}
	return o.Associated, true
}

// HasAssociated returns a boolean if a field has been set.
func (o *SeriesModelV1) HasAssociated() bool {
	if o != nil && !IsNil(o.Associated) {
		return true
	}

	return false
}

// SetAssociated gets a reference to the given []SeriesModelV1Associated and assigns it to the Associated field.
func (o *SeriesModelV1) SetAssociated(v []SeriesModelV1Associated) {
	o.Associated = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SeriesModelV1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SeriesModelV1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SeriesModelV1) SetDescription(v string) {
	o.Description = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SeriesModelV1) GetImage() ImageModelV1 {
	if o == nil || IsNil(o.Image) {
		var ret ImageModelV1
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetImageOk() (*ImageModelV1, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SeriesModelV1) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageModelV1 and assigns it to the Image field.
func (o *SeriesModelV1) SetImage(v ImageModelV1) {
	o.Image = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SeriesModelV1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SeriesModelV1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SeriesModelV1) SetType(v string) {
	o.Type = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SeriesModelV1) GetYear() string {
	if o == nil || IsNil(o.Year) {
		var ret string
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetYearOk() (*string, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SeriesModelV1) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given string and assigns it to the Year field.
func (o *SeriesModelV1) SetYear(v string) {
	o.Year = &v
}

// GetBayesianRating returns the BayesianRating field value if set, zero value otherwise.
func (o *SeriesModelV1) GetBayesianRating() float32 {
	if o == nil || IsNil(o.BayesianRating) {
		var ret float32
		return ret
	}
	return *o.BayesianRating
}

// GetBayesianRatingOk returns a tuple with the BayesianRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetBayesianRatingOk() (*float32, bool) {
	if o == nil || IsNil(o.BayesianRating) {
		return nil, false
	}
	return o.BayesianRating, true
}

// HasBayesianRating returns a boolean if a field has been set.
func (o *SeriesModelV1) HasBayesianRating() bool {
	if o != nil && !IsNil(o.BayesianRating) {
		return true
	}

	return false
}

// SetBayesianRating gets a reference to the given float32 and assigns it to the BayesianRating field.
func (o *SeriesModelV1) SetBayesianRating(v float32) {
	o.BayesianRating = &v
}

// GetRatingVotes returns the RatingVotes field value if set, zero value otherwise.
func (o *SeriesModelV1) GetRatingVotes() int64 {
	if o == nil || IsNil(o.RatingVotes) {
		var ret int64
		return ret
	}
	return *o.RatingVotes
}

// GetRatingVotesOk returns a tuple with the RatingVotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetRatingVotesOk() (*int64, bool) {
	if o == nil || IsNil(o.RatingVotes) {
		return nil, false
	}
	return o.RatingVotes, true
}

// HasRatingVotes returns a boolean if a field has been set.
func (o *SeriesModelV1) HasRatingVotes() bool {
	if o != nil && !IsNil(o.RatingVotes) {
		return true
	}

	return false
}

// SetRatingVotes gets a reference to the given int64 and assigns it to the RatingVotes field.
func (o *SeriesModelV1) SetRatingVotes(v int64) {
	o.RatingVotes = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SeriesModelV1) GetGenres() []SeriesModelV1Genres {
	if o == nil || IsNil(o.Genres) {
		var ret []SeriesModelV1Genres
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetGenresOk() ([]SeriesModelV1Genres, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SeriesModelV1) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []SeriesModelV1Genres and assigns it to the Genres field.
func (o *SeriesModelV1) SetGenres(v []SeriesModelV1Genres) {
	o.Genres = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SeriesModelV1) GetCategories() []CategoriesModelV1 {
	if o == nil || IsNil(o.Categories) {
		var ret []CategoriesModelV1
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetCategoriesOk() ([]CategoriesModelV1, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SeriesModelV1) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoriesModelV1 and assigns it to the Categories field.
func (o *SeriesModelV1) SetCategories(v []CategoriesModelV1) {
	o.Categories = v
}

// GetLatestChapter returns the LatestChapter field value if set, zero value otherwise.
func (o *SeriesModelV1) GetLatestChapter() int64 {
	if o == nil || IsNil(o.LatestChapter) {
		var ret int64
		return ret
	}
	return *o.LatestChapter
}

// GetLatestChapterOk returns a tuple with the LatestChapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetLatestChapterOk() (*int64, bool) {
	if o == nil || IsNil(o.LatestChapter) {
		return nil, false
	}
	return o.LatestChapter, true
}

// HasLatestChapter returns a boolean if a field has been set.
func (o *SeriesModelV1) HasLatestChapter() bool {
	if o != nil && !IsNil(o.LatestChapter) {
		return true
	}

	return false
}

// SetLatestChapter gets a reference to the given int64 and assigns it to the LatestChapter field.
func (o *SeriesModelV1) SetLatestChapter(v int64) {
	o.LatestChapter = &v
}

// GetForumId returns the ForumId field value if set, zero value otherwise.
func (o *SeriesModelV1) GetForumId() int64 {
	if o == nil || IsNil(o.ForumId) {
		var ret int64
		return ret
	}
	return *o.ForumId
}

// GetForumIdOk returns a tuple with the ForumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetForumIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ForumId) {
		return nil, false
	}
	return o.ForumId, true
}

// HasForumId returns a boolean if a field has been set.
func (o *SeriesModelV1) HasForumId() bool {
	if o != nil && !IsNil(o.ForumId) {
		return true
	}

	return false
}

// SetForumId gets a reference to the given int64 and assigns it to the ForumId field.
func (o *SeriesModelV1) SetForumId(v int64) {
	o.ForumId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SeriesModelV1) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SeriesModelV1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SeriesModelV1) SetStatus(v string) {
	o.Status = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *SeriesModelV1) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *SeriesModelV1) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *SeriesModelV1) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *SeriesModelV1) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *SeriesModelV1) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *SeriesModelV1) SetCompleted(v bool) {
	o.Completed = &v
}

// GetAnime returns the Anime field value if set, zero value otherwise.
func (o *SeriesModelV1) GetAnime() SeriesModelV1Anime {
	if o == nil || IsNil(o.Anime) {
		var ret SeriesModelV1Anime
		return ret
	}
	return *o.Anime
}

// GetAnimeOk returns a tuple with the Anime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetAnimeOk() (*SeriesModelV1Anime, bool) {
	if o == nil || IsNil(o.Anime) {
		return nil, false
	}
	return o.Anime, true
}

// HasAnime returns a boolean if a field has been set.
func (o *SeriesModelV1) HasAnime() bool {
	if o != nil && !IsNil(o.Anime) {
		return true
	}

	return false
}

// SetAnime gets a reference to the given SeriesModelV1Anime and assigns it to the Anime field.
func (o *SeriesModelV1) SetAnime(v SeriesModelV1Anime) {
	o.Anime = &v
}

// GetRelatedSeries returns the RelatedSeries field value if set, zero value otherwise.
func (o *SeriesModelV1) GetRelatedSeries() []SeriesModelV1RelatedSeries {
	if o == nil || IsNil(o.RelatedSeries) {
		var ret []SeriesModelV1RelatedSeries
		return ret
	}
	return o.RelatedSeries
}

// GetRelatedSeriesOk returns a tuple with the RelatedSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetRelatedSeriesOk() ([]SeriesModelV1RelatedSeries, bool) {
	if o == nil || IsNil(o.RelatedSeries) {
		return nil, false
	}
	return o.RelatedSeries, true
}

// HasRelatedSeries returns a boolean if a field has been set.
func (o *SeriesModelV1) HasRelatedSeries() bool {
	if o != nil && !IsNil(o.RelatedSeries) {
		return true
	}

	return false
}

// SetRelatedSeries gets a reference to the given []SeriesModelV1RelatedSeries and assigns it to the RelatedSeries field.
func (o *SeriesModelV1) SetRelatedSeries(v []SeriesModelV1RelatedSeries) {
	o.RelatedSeries = v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *SeriesModelV1) GetAuthors() []SeriesModelV1Authors {
	if o == nil || IsNil(o.Authors) {
		var ret []SeriesModelV1Authors
		return ret
	}
	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetAuthorsOk() ([]SeriesModelV1Authors, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *SeriesModelV1) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given []SeriesModelV1Authors and assigns it to the Authors field.
func (o *SeriesModelV1) SetAuthors(v []SeriesModelV1Authors) {
	o.Authors = v
}

// GetPublishers returns the Publishers field value if set, zero value otherwise.
func (o *SeriesModelV1) GetPublishers() []SeriesModelV1Publishers {
	if o == nil || IsNil(o.Publishers) {
		var ret []SeriesModelV1Publishers
		return ret
	}
	return o.Publishers
}

// GetPublishersOk returns a tuple with the Publishers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetPublishersOk() ([]SeriesModelV1Publishers, bool) {
	if o == nil || IsNil(o.Publishers) {
		return nil, false
	}
	return o.Publishers, true
}

// HasPublishers returns a boolean if a field has been set.
func (o *SeriesModelV1) HasPublishers() bool {
	if o != nil && !IsNil(o.Publishers) {
		return true
	}

	return false
}

// SetPublishers gets a reference to the given []SeriesModelV1Publishers and assigns it to the Publishers field.
func (o *SeriesModelV1) SetPublishers(v []SeriesModelV1Publishers) {
	o.Publishers = v
}

// GetPublications returns the Publications field value if set, zero value otherwise.
func (o *SeriesModelV1) GetPublications() []SeriesModelV1Publications {
	if o == nil || IsNil(o.Publications) {
		var ret []SeriesModelV1Publications
		return ret
	}
	return o.Publications
}

// GetPublicationsOk returns a tuple with the Publications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetPublicationsOk() ([]SeriesModelV1Publications, bool) {
	if o == nil || IsNil(o.Publications) {
		return nil, false
	}
	return o.Publications, true
}

// HasPublications returns a boolean if a field has been set.
func (o *SeriesModelV1) HasPublications() bool {
	if o != nil && !IsNil(o.Publications) {
		return true
	}

	return false
}

// SetPublications gets a reference to the given []SeriesModelV1Publications and assigns it to the Publications field.
func (o *SeriesModelV1) SetPublications(v []SeriesModelV1Publications) {
	o.Publications = v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SeriesModelV1) GetRecommendations() []SeriesRecommendationsModelV1 {
	if o == nil || IsNil(o.Recommendations) {
		var ret []SeriesRecommendationsModelV1
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetRecommendationsOk() ([]SeriesRecommendationsModelV1, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SeriesModelV1) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []SeriesRecommendationsModelV1 and assigns it to the Recommendations field.
func (o *SeriesModelV1) SetRecommendations(v []SeriesRecommendationsModelV1) {
	o.Recommendations = v
}

// GetCategoryRecommendations returns the CategoryRecommendations field value if set, zero value otherwise.
func (o *SeriesModelV1) GetCategoryRecommendations() []SeriesRecommendationsModelV1 {
	if o == nil || IsNil(o.CategoryRecommendations) {
		var ret []SeriesRecommendationsModelV1
		return ret
	}
	return o.CategoryRecommendations
}

// GetCategoryRecommendationsOk returns a tuple with the CategoryRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetCategoryRecommendationsOk() ([]SeriesRecommendationsModelV1, bool) {
	if o == nil || IsNil(o.CategoryRecommendations) {
		return nil, false
	}
	return o.CategoryRecommendations, true
}

// HasCategoryRecommendations returns a boolean if a field has been set.
func (o *SeriesModelV1) HasCategoryRecommendations() bool {
	if o != nil && !IsNil(o.CategoryRecommendations) {
		return true
	}

	return false
}

// SetCategoryRecommendations gets a reference to the given []SeriesRecommendationsModelV1 and assigns it to the CategoryRecommendations field.
func (o *SeriesModelV1) SetCategoryRecommendations(v []SeriesRecommendationsModelV1) {
	o.CategoryRecommendations = v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SeriesModelV1) GetRank() SeriesModelV1Rank {
	if o == nil || IsNil(o.Rank) {
		var ret SeriesModelV1Rank
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetRankOk() (*SeriesModelV1Rank, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SeriesModelV1) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given SeriesModelV1Rank and assigns it to the Rank field.
func (o *SeriesModelV1) SetRank(v SeriesModelV1Rank) {
	o.Rank = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SeriesModelV1) GetLastUpdated() TimeV1 {
	if o == nil || IsNil(o.LastUpdated) {
		var ret TimeV1
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetLastUpdatedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SeriesModelV1) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given TimeV1 and assigns it to the LastUpdated field.
func (o *SeriesModelV1) SetLastUpdated(v TimeV1) {
	o.LastUpdated = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *SeriesModelV1) GetAdmin() SeriesModelV1Admin {
	if o == nil || IsNil(o.Admin) {
		var ret SeriesModelV1Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1) GetAdminOk() (*SeriesModelV1Admin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *SeriesModelV1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given SeriesModelV1Admin and assigns it to the Admin field.
func (o *SeriesModelV1) SetAdmin(v SeriesModelV1Admin) {
	o.Admin = &v
}

func (o SeriesModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SeriesId) {
		toSerialize["series_id"] = o.SeriesId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Associated) {
		toSerialize["associated"] = o.Associated
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.BayesianRating) {
		toSerialize["bayesian_rating"] = o.BayesianRating
	}
	if !IsNil(o.RatingVotes) {
		toSerialize["rating_votes"] = o.RatingVotes
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.LatestChapter) {
		toSerialize["latest_chapter"] = o.LatestChapter
	}
	if !IsNil(o.ForumId) {
		toSerialize["forum_id"] = o.ForumId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Licensed) {
		toSerialize["licensed"] = o.Licensed
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Anime) {
		toSerialize["anime"] = o.Anime
	}
	if !IsNil(o.RelatedSeries) {
		toSerialize["related_series"] = o.RelatedSeries
	}
	if !IsNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !IsNil(o.Publishers) {
		toSerialize["publishers"] = o.Publishers
	}
	if !IsNil(o.Publications) {
		toSerialize["publications"] = o.Publications
	}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	if !IsNil(o.CategoryRecommendations) {
		toSerialize["category_recommendations"] = o.CategoryRecommendations
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

type NullableSeriesModelV1 struct {
	value *SeriesModelV1
	isSet bool
}

func (v NullableSeriesModelV1) Get() *SeriesModelV1 {
	return v.value
}

func (v *NullableSeriesModelV1) Set(val *SeriesModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelV1(val *SeriesModelV1) *NullableSeriesModelV1 {
	return &NullableSeriesModelV1{value: val, isSet: true}
}

func (v NullableSeriesModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}