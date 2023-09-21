# UserModelV1Stats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForumPosts** | Pointer to **int64** |  | [optional] 
**AddedAuthors** | Pointer to **int64** |  | [optional] 
**AddedGroups** | Pointer to **int64** |  | [optional] 
**AddedPublishers** | Pointer to **int64** |  | [optional] 
**AddedReleases** | Pointer to **int64** |  | [optional] 
**AddedSeries** | Pointer to **int64** |  | [optional] 
**SeriesEdits** | Pointer to **int64** |  | [optional] 
**AuthorEdits** | Pointer to **int64** |  | [optional] 
**PublisherEdits** | Pointer to **int64** |  | [optional] 
**AddedTags** | Pointer to **int64** |  | [optional] 
**Moderation** | Pointer to [**UserModelV1StatsModeration**](UserModelV1StatsModeration.md) |  | [optional] 

## Methods

### NewUserModelV1Stats

`func NewUserModelV1Stats() *UserModelV1Stats`

NewUserModelV1Stats instantiates a new UserModelV1Stats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelV1StatsWithDefaults

`func NewUserModelV1StatsWithDefaults() *UserModelV1Stats`

NewUserModelV1StatsWithDefaults instantiates a new UserModelV1Stats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForumPosts

`func (o *UserModelV1Stats) GetForumPosts() int64`

GetForumPosts returns the ForumPosts field if non-nil, zero value otherwise.

### GetForumPostsOk

`func (o *UserModelV1Stats) GetForumPostsOk() (*int64, bool)`

GetForumPostsOk returns a tuple with the ForumPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPosts

`func (o *UserModelV1Stats) SetForumPosts(v int64)`

SetForumPosts sets ForumPosts field to given value.

### HasForumPosts

`func (o *UserModelV1Stats) HasForumPosts() bool`

HasForumPosts returns a boolean if a field has been set.

### GetAddedAuthors

`func (o *UserModelV1Stats) GetAddedAuthors() int64`

GetAddedAuthors returns the AddedAuthors field if non-nil, zero value otherwise.

### GetAddedAuthorsOk

`func (o *UserModelV1Stats) GetAddedAuthorsOk() (*int64, bool)`

GetAddedAuthorsOk returns a tuple with the AddedAuthors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAuthors

`func (o *UserModelV1Stats) SetAddedAuthors(v int64)`

SetAddedAuthors sets AddedAuthors field to given value.

### HasAddedAuthors

`func (o *UserModelV1Stats) HasAddedAuthors() bool`

HasAddedAuthors returns a boolean if a field has been set.

### GetAddedGroups

`func (o *UserModelV1Stats) GetAddedGroups() int64`

GetAddedGroups returns the AddedGroups field if non-nil, zero value otherwise.

### GetAddedGroupsOk

`func (o *UserModelV1Stats) GetAddedGroupsOk() (*int64, bool)`

GetAddedGroupsOk returns a tuple with the AddedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedGroups

`func (o *UserModelV1Stats) SetAddedGroups(v int64)`

SetAddedGroups sets AddedGroups field to given value.

### HasAddedGroups

`func (o *UserModelV1Stats) HasAddedGroups() bool`

HasAddedGroups returns a boolean if a field has been set.

### GetAddedPublishers

`func (o *UserModelV1Stats) GetAddedPublishers() int64`

GetAddedPublishers returns the AddedPublishers field if non-nil, zero value otherwise.

### GetAddedPublishersOk

`func (o *UserModelV1Stats) GetAddedPublishersOk() (*int64, bool)`

GetAddedPublishersOk returns a tuple with the AddedPublishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPublishers

`func (o *UserModelV1Stats) SetAddedPublishers(v int64)`

SetAddedPublishers sets AddedPublishers field to given value.

### HasAddedPublishers

`func (o *UserModelV1Stats) HasAddedPublishers() bool`

HasAddedPublishers returns a boolean if a field has been set.

### GetAddedReleases

`func (o *UserModelV1Stats) GetAddedReleases() int64`

GetAddedReleases returns the AddedReleases field if non-nil, zero value otherwise.

### GetAddedReleasesOk

`func (o *UserModelV1Stats) GetAddedReleasesOk() (*int64, bool)`

GetAddedReleasesOk returns a tuple with the AddedReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedReleases

`func (o *UserModelV1Stats) SetAddedReleases(v int64)`

SetAddedReleases sets AddedReleases field to given value.

### HasAddedReleases

`func (o *UserModelV1Stats) HasAddedReleases() bool`

HasAddedReleases returns a boolean if a field has been set.

### GetAddedSeries

`func (o *UserModelV1Stats) GetAddedSeries() int64`

GetAddedSeries returns the AddedSeries field if non-nil, zero value otherwise.

### GetAddedSeriesOk

`func (o *UserModelV1Stats) GetAddedSeriesOk() (*int64, bool)`

GetAddedSeriesOk returns a tuple with the AddedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedSeries

`func (o *UserModelV1Stats) SetAddedSeries(v int64)`

SetAddedSeries sets AddedSeries field to given value.

### HasAddedSeries

`func (o *UserModelV1Stats) HasAddedSeries() bool`

HasAddedSeries returns a boolean if a field has been set.

### GetSeriesEdits

`func (o *UserModelV1Stats) GetSeriesEdits() int64`

GetSeriesEdits returns the SeriesEdits field if non-nil, zero value otherwise.

### GetSeriesEditsOk

`func (o *UserModelV1Stats) GetSeriesEditsOk() (*int64, bool)`

GetSeriesEditsOk returns a tuple with the SeriesEdits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesEdits

`func (o *UserModelV1Stats) SetSeriesEdits(v int64)`

SetSeriesEdits sets SeriesEdits field to given value.

### HasSeriesEdits

`func (o *UserModelV1Stats) HasSeriesEdits() bool`

HasSeriesEdits returns a boolean if a field has been set.

### GetAuthorEdits

`func (o *UserModelV1Stats) GetAuthorEdits() int64`

GetAuthorEdits returns the AuthorEdits field if non-nil, zero value otherwise.

### GetAuthorEditsOk

`func (o *UserModelV1Stats) GetAuthorEditsOk() (*int64, bool)`

GetAuthorEditsOk returns a tuple with the AuthorEdits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEdits

`func (o *UserModelV1Stats) SetAuthorEdits(v int64)`

SetAuthorEdits sets AuthorEdits field to given value.

### HasAuthorEdits

`func (o *UserModelV1Stats) HasAuthorEdits() bool`

HasAuthorEdits returns a boolean if a field has been set.

### GetPublisherEdits

`func (o *UserModelV1Stats) GetPublisherEdits() int64`

GetPublisherEdits returns the PublisherEdits field if non-nil, zero value otherwise.

### GetPublisherEditsOk

`func (o *UserModelV1Stats) GetPublisherEditsOk() (*int64, bool)`

GetPublisherEditsOk returns a tuple with the PublisherEdits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherEdits

`func (o *UserModelV1Stats) SetPublisherEdits(v int64)`

SetPublisherEdits sets PublisherEdits field to given value.

### HasPublisherEdits

`func (o *UserModelV1Stats) HasPublisherEdits() bool`

HasPublisherEdits returns a boolean if a field has been set.

### GetAddedTags

`func (o *UserModelV1Stats) GetAddedTags() int64`

GetAddedTags returns the AddedTags field if non-nil, zero value otherwise.

### GetAddedTagsOk

`func (o *UserModelV1Stats) GetAddedTagsOk() (*int64, bool)`

GetAddedTagsOk returns a tuple with the AddedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTags

`func (o *UserModelV1Stats) SetAddedTags(v int64)`

SetAddedTags sets AddedTags field to given value.

### HasAddedTags

`func (o *UserModelV1Stats) HasAddedTags() bool`

HasAddedTags returns a boolean if a field has been set.

### GetModeration

`func (o *UserModelV1Stats) GetModeration() UserModelV1StatsModeration`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *UserModelV1Stats) GetModerationOk() (*UserModelV1StatsModeration, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *UserModelV1Stats) SetModeration(v UserModelV1StatsModeration)`

SetModeration sets Moderation field to given value.

### HasModeration

`func (o *UserModelV1Stats) HasModeration() bool`

HasModeration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


