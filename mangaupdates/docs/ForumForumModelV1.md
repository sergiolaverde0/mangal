# ForumForumModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForumId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to [**ForumForumModelV1Category**](ForumForumModelV1Category.md) |  | [optional] 
**Moderators** | Pointer to [**[]ForumAdminModelV1**](ForumAdminModelV1.md) |  | [optional] 
**Series** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**Stats** | Pointer to [**ForumForumModelV1Stats**](ForumForumModelV1Stats.md) |  | [optional] 
**LastTopic** | Pointer to [**ForumTopicModelSearchV1**](ForumTopicModelSearchV1.md) |  | [optional] 
**Admin** | Pointer to [**ForumForumModelV1Admin**](ForumForumModelV1Admin.md) |  | [optional] 

## Methods

### NewForumForumModelV1

`func NewForumForumModelV1() *ForumForumModelV1`

NewForumForumModelV1 instantiates a new ForumForumModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumForumModelV1WithDefaults

`func NewForumForumModelV1WithDefaults() *ForumForumModelV1`

NewForumForumModelV1WithDefaults instantiates a new ForumForumModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForumId

`func (o *ForumForumModelV1) GetForumId() int64`

GetForumId returns the ForumId field if non-nil, zero value otherwise.

### GetForumIdOk

`func (o *ForumForumModelV1) GetForumIdOk() (*int64, bool)`

GetForumIdOk returns a tuple with the ForumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumId

`func (o *ForumForumModelV1) SetForumId(v int64)`

SetForumId sets ForumId field to given value.

### HasForumId

`func (o *ForumForumModelV1) HasForumId() bool`

HasForumId returns a boolean if a field has been set.

### GetName

`func (o *ForumForumModelV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumForumModelV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumForumModelV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForumForumModelV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ForumForumModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ForumForumModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ForumForumModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ForumForumModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ForumForumModelV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ForumForumModelV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ForumForumModelV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ForumForumModelV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPosition

`func (o *ForumForumModelV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ForumForumModelV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ForumForumModelV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ForumForumModelV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCategory

`func (o *ForumForumModelV1) GetCategory() ForumForumModelV1Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ForumForumModelV1) GetCategoryOk() (*ForumForumModelV1Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ForumForumModelV1) SetCategory(v ForumForumModelV1Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ForumForumModelV1) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetModerators

`func (o *ForumForumModelV1) GetModerators() []ForumAdminModelV1`

GetModerators returns the Moderators field if non-nil, zero value otherwise.

### GetModeratorsOk

`func (o *ForumForumModelV1) GetModeratorsOk() (*[]ForumAdminModelV1, bool)`

GetModeratorsOk returns a tuple with the Moderators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerators

`func (o *ForumForumModelV1) SetModerators(v []ForumAdminModelV1)`

SetModerators sets Moderators field to given value.

### HasModerators

`func (o *ForumForumModelV1) HasModerators() bool`

HasModerators returns a boolean if a field has been set.

### GetSeries

`func (o *ForumForumModelV1) GetSeries() SeriesModelSearchV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ForumForumModelV1) GetSeriesOk() (*SeriesModelSearchV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ForumForumModelV1) SetSeries(v SeriesModelSearchV1)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ForumForumModelV1) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStats

`func (o *ForumForumModelV1) GetStats() ForumForumModelV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ForumForumModelV1) GetStatsOk() (*ForumForumModelV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ForumForumModelV1) SetStats(v ForumForumModelV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ForumForumModelV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetLastTopic

`func (o *ForumForumModelV1) GetLastTopic() ForumTopicModelSearchV1`

GetLastTopic returns the LastTopic field if non-nil, zero value otherwise.

### GetLastTopicOk

`func (o *ForumForumModelV1) GetLastTopicOk() (*ForumTopicModelSearchV1, bool)`

GetLastTopicOk returns a tuple with the LastTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTopic

`func (o *ForumForumModelV1) SetLastTopic(v ForumTopicModelSearchV1)`

SetLastTopic sets LastTopic field to given value.

### HasLastTopic

`func (o *ForumForumModelV1) HasLastTopic() bool`

HasLastTopic returns a boolean if a field has been set.

### GetAdmin

`func (o *ForumForumModelV1) GetAdmin() ForumForumModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ForumForumModelV1) GetAdminOk() (*ForumForumModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ForumForumModelV1) SetAdmin(v ForumForumModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ForumForumModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


