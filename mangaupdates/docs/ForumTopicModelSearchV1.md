# ForumTopicModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicId** | Pointer to **int64** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**LastPost** | Pointer to [**ForumPostModelSearchV1**](ForumPostModelSearchV1.md) |  | [optional] 
**Stats** | Pointer to [**ForumTopicModelSearchV1Stats**](ForumTopicModelSearchV1Stats.md) |  | [optional] 
**Forum** | Pointer to [**ForumTopicModelSearchV1Forum**](ForumTopicModelSearchV1Forum.md) |  | [optional] 
**IsPoll** | Pointer to **bool** |  | [optional] 
**Admin** | Pointer to [**ForumTopicModelSearchV1Admin**](ForumTopicModelSearchV1Admin.md) |  | [optional] 
**TopicStarter** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewForumTopicModelSearchV1

`func NewForumTopicModelSearchV1() *ForumTopicModelSearchV1`

NewForumTopicModelSearchV1 instantiates a new ForumTopicModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumTopicModelSearchV1WithDefaults

`func NewForumTopicModelSearchV1WithDefaults() *ForumTopicModelSearchV1`

NewForumTopicModelSearchV1WithDefaults instantiates a new ForumTopicModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicId

`func (o *ForumTopicModelSearchV1) GetTopicId() int64`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *ForumTopicModelSearchV1) GetTopicIdOk() (*int64, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *ForumTopicModelSearchV1) SetTopicId(v int64)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *ForumTopicModelSearchV1) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### GetTopic

`func (o *ForumTopicModelSearchV1) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumTopicModelSearchV1) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumTopicModelSearchV1) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumTopicModelSearchV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *ForumTopicModelSearchV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ForumTopicModelSearchV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ForumTopicModelSearchV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ForumTopicModelSearchV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLastPost

`func (o *ForumTopicModelSearchV1) GetLastPost() ForumPostModelSearchV1`

GetLastPost returns the LastPost field if non-nil, zero value otherwise.

### GetLastPostOk

`func (o *ForumTopicModelSearchV1) GetLastPostOk() (*ForumPostModelSearchV1, bool)`

GetLastPostOk returns a tuple with the LastPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPost

`func (o *ForumTopicModelSearchV1) SetLastPost(v ForumPostModelSearchV1)`

SetLastPost sets LastPost field to given value.

### HasLastPost

`func (o *ForumTopicModelSearchV1) HasLastPost() bool`

HasLastPost returns a boolean if a field has been set.

### GetStats

`func (o *ForumTopicModelSearchV1) GetStats() ForumTopicModelSearchV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ForumTopicModelSearchV1) GetStatsOk() (*ForumTopicModelSearchV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ForumTopicModelSearchV1) SetStats(v ForumTopicModelSearchV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ForumTopicModelSearchV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetForum

`func (o *ForumTopicModelSearchV1) GetForum() ForumTopicModelSearchV1Forum`

GetForum returns the Forum field if non-nil, zero value otherwise.

### GetForumOk

`func (o *ForumTopicModelSearchV1) GetForumOk() (*ForumTopicModelSearchV1Forum, bool)`

GetForumOk returns a tuple with the Forum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForum

`func (o *ForumTopicModelSearchV1) SetForum(v ForumTopicModelSearchV1Forum)`

SetForum sets Forum field to given value.

### HasForum

`func (o *ForumTopicModelSearchV1) HasForum() bool`

HasForum returns a boolean if a field has been set.

### GetIsPoll

`func (o *ForumTopicModelSearchV1) GetIsPoll() bool`

GetIsPoll returns the IsPoll field if non-nil, zero value otherwise.

### GetIsPollOk

`func (o *ForumTopicModelSearchV1) GetIsPollOk() (*bool, bool)`

GetIsPollOk returns a tuple with the IsPoll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPoll

`func (o *ForumTopicModelSearchV1) SetIsPoll(v bool)`

SetIsPoll sets IsPoll field to given value.

### HasIsPoll

`func (o *ForumTopicModelSearchV1) HasIsPoll() bool`

HasIsPoll returns a boolean if a field has been set.

### GetAdmin

`func (o *ForumTopicModelSearchV1) GetAdmin() ForumTopicModelSearchV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ForumTopicModelSearchV1) GetAdminOk() (*ForumTopicModelSearchV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ForumTopicModelSearchV1) SetAdmin(v ForumTopicModelSearchV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ForumTopicModelSearchV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetTopicStarter

`func (o *ForumTopicModelSearchV1) GetTopicStarter() UserModelSearchV1`

GetTopicStarter returns the TopicStarter field if non-nil, zero value otherwise.

### GetTopicStarterOk

`func (o *ForumTopicModelSearchV1) GetTopicStarterOk() (*UserModelSearchV1, bool)`

GetTopicStarterOk returns a tuple with the TopicStarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicStarter

`func (o *ForumTopicModelSearchV1) SetTopicStarter(v UserModelSearchV1)`

SetTopicStarter sets TopicStarter field to given value.

### HasTopicStarter

`func (o *ForumTopicModelSearchV1) HasTopicStarter() bool`

HasTopicStarter returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ForumTopicModelSearchV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ForumTopicModelSearchV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ForumTopicModelSearchV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ForumTopicModelSearchV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


