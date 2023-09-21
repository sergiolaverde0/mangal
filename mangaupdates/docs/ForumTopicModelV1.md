# ForumTopicModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicId** | Pointer to **int64** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**LastPost** | Pointer to [**ForumPostModelSearchV1**](ForumPostModelSearchV1.md) |  | [optional] 
**Stats** | Pointer to [**ForumTopicModelV1Stats**](ForumTopicModelV1Stats.md) |  | [optional] 
**Forum** | Pointer to [**ForumTopicModelV1Forum**](ForumTopicModelV1Forum.md) |  | [optional] 
**IsPoll** | Pointer to **bool** |  | [optional] 
**Poll** | Pointer to [**ForumPollModelV1**](ForumPollModelV1.md) |  | [optional] 
**Admin** | Pointer to [**ForumTopicModelV1Admin**](ForumTopicModelV1Admin.md) |  | [optional] 
**TopicStarter** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewForumTopicModelV1

`func NewForumTopicModelV1() *ForumTopicModelV1`

NewForumTopicModelV1 instantiates a new ForumTopicModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTopicModelV1WithDefaults

`func NewForumTopicModelV1WithDefaults() *ForumTopicModelV1`

NewForumTopicModelV1WithDefaults instantiates a new ForumTopicModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicId

`func (o *ForumTopicModelV1) GetTopicId() int64`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *ForumTopicModelV1) GetTopicIdOk() (*int64, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *ForumTopicModelV1) SetTopicId(v int64)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *ForumTopicModelV1) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### GetTopic

`func (o *ForumTopicModelV1) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumTopicModelV1) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumTopicModelV1) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumTopicModelV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *ForumTopicModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ForumTopicModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ForumTopicModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ForumTopicModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLastPost

`func (o *ForumTopicModelV1) GetLastPost() ForumPostModelSearchV1`

GetLastPost returns the LastPost field if non-nil, zero value otherwise.

### GetLastPostOk

`func (o *ForumTopicModelV1) GetLastPostOk() (*ForumPostModelSearchV1, bool)`

GetLastPostOk returns a tuple with the LastPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPost

`func (o *ForumTopicModelV1) SetLastPost(v ForumPostModelSearchV1)`

SetLastPost sets LastPost field to given value.

### HasLastPost

`func (o *ForumTopicModelV1) HasLastPost() bool`

HasLastPost returns a boolean if a field has been set.

### GetStats

`func (o *ForumTopicModelV1) GetStats() ForumTopicModelV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ForumTopicModelV1) GetStatsOk() (*ForumTopicModelV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ForumTopicModelV1) SetStats(v ForumTopicModelV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ForumTopicModelV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetForum

`func (o *ForumTopicModelV1) GetForum() ForumTopicModelV1Forum`

GetForum returns the Forum field if non-nil, zero value otherwise.

### GetForumOk

`func (o *ForumTopicModelV1) GetForumOk() (*ForumTopicModelV1Forum, bool)`

GetForumOk returns a tuple with the Forum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForum

`func (o *ForumTopicModelV1) SetForum(v ForumTopicModelV1Forum)`

SetForum sets Forum field to given value.

### HasForum

`func (o *ForumTopicModelV1) HasForum() bool`

HasForum returns a boolean if a field has been set.

### GetIsPoll

`func (o *ForumTopicModelV1) GetIsPoll() bool`

GetIsPoll returns the IsPoll field if non-nil, zero value otherwise.

### GetIsPollOk

`func (o *ForumTopicModelV1) GetIsPollOk() (*bool, bool)`

GetIsPollOk returns a tuple with the IsPoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPoll

`func (o *ForumTopicModelV1) SetIsPoll(v bool)`

SetIsPoll sets IsPoll field to given value.

### HasIsPoll

`func (o *ForumTopicModelV1) HasIsPoll() bool`

HasIsPoll returns a boolean if a field has been set.

### GetPoll

`func (o *ForumTopicModelV1) GetPoll() ForumPollModelV1`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *ForumTopicModelV1) GetPollOk() (*ForumPollModelV1, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *ForumTopicModelV1) SetPoll(v ForumPollModelV1)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *ForumTopicModelV1) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetAdmin

`func (o *ForumTopicModelV1) GetAdmin() ForumTopicModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ForumTopicModelV1) GetAdminOk() (*ForumTopicModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ForumTopicModelV1) SetAdmin(v ForumTopicModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ForumTopicModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetTopicStarter

`func (o *ForumTopicModelV1) GetTopicStarter() UserModelSearchV1`

GetTopicStarter returns the TopicStarter field if non-nil, zero value otherwise.

### GetTopicStarterOk

`func (o *ForumTopicModelV1) GetTopicStarterOk() (*UserModelSearchV1, bool)`

GetTopicStarterOk returns a tuple with the TopicStarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicStarter

`func (o *ForumTopicModelV1) SetTopicStarter(v UserModelSearchV1)`

SetTopicStarter sets TopicStarter field to given value.

### HasTopicStarter

`func (o *ForumTopicModelV1) HasTopicStarter() bool`

HasTopicStarter returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ForumTopicModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ForumTopicModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ForumTopicModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ForumTopicModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


