# ForumForumModelListV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForumId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Moderators** | Pointer to [**[]ForumAdminModelV1**](ForumAdminModelV1.md) |  | [optional] 
**Series** | Pointer to [**SeriesModelSearchV1**](SeriesModelSearchV1.md) |  | [optional] 
**Stats** | Pointer to [**ForumForumModelListV1Stats**](ForumForumModelListV1Stats.md) |  | [optional] 
**LastTopic** | Pointer to [**ForumTopicModelSearchV1**](ForumTopicModelSearchV1.md) |  | [optional] 
**Admin** | Pointer to [**ForumForumModelListV1Admin**](ForumForumModelListV1Admin.md) |  | [optional] 

## Methods

### NewForumForumModelListV1

`func NewForumForumModelListV1() *ForumForumModelListV1`

NewForumForumModelListV1 instantiates a new ForumForumModelListV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumForumModelListV1WithDefaults

`func NewForumForumModelListV1WithDefaults() *ForumForumModelListV1`

NewForumForumModelListV1WithDefaults instantiates a new ForumForumModelListV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForumId

`func (o *ForumForumModelListV1) GetForumId() int64`

GetForumId returns the ForumId field if non-nil, zero value otherwise.

### GetForumIdOk

`func (o *ForumForumModelListV1) GetForumIdOk() (*int64, bool)`

GetForumIdOk returns a tuple with the ForumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumId

`func (o *ForumForumModelListV1) SetForumId(v int64)`

SetForumId sets ForumId field to given value.

### HasForumId

`func (o *ForumForumModelListV1) HasForumId() bool`

HasForumId returns a boolean if a field has been set.

### GetName

`func (o *ForumForumModelListV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForumForumModelListV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForumForumModelListV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForumForumModelListV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ForumForumModelListV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ForumForumModelListV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ForumForumModelListV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ForumForumModelListV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ForumForumModelListV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ForumForumModelListV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ForumForumModelListV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ForumForumModelListV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPosition

`func (o *ForumForumModelListV1) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ForumForumModelListV1) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ForumForumModelListV1) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ForumForumModelListV1) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetModerators

`func (o *ForumForumModelListV1) GetModerators() []ForumAdminModelV1`

GetModerators returns the Moderators field if non-nil, zero value otherwise.

### GetModeratorsOk

`func (o *ForumForumModelListV1) GetModeratorsOk() (*[]ForumAdminModelV1, bool)`

GetModeratorsOk returns a tuple with the Moderators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerators

`func (o *ForumForumModelListV1) SetModerators(v []ForumAdminModelV1)`

SetModerators sets Moderators field to given value.

### HasModerators

`func (o *ForumForumModelListV1) HasModerators() bool`

HasModerators returns a boolean if a field has been set.

### GetSeries

`func (o *ForumForumModelListV1) GetSeries() SeriesModelSearchV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ForumForumModelListV1) GetSeriesOk() (*SeriesModelSearchV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ForumForumModelListV1) SetSeries(v SeriesModelSearchV1)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ForumForumModelListV1) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStats

`func (o *ForumForumModelListV1) GetStats() ForumForumModelListV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ForumForumModelListV1) GetStatsOk() (*ForumForumModelListV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ForumForumModelListV1) SetStats(v ForumForumModelListV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ForumForumModelListV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetLastTopic

`func (o *ForumForumModelListV1) GetLastTopic() ForumTopicModelSearchV1`

GetLastTopic returns the LastTopic field if non-nil, zero value otherwise.

### GetLastTopicOk

`func (o *ForumForumModelListV1) GetLastTopicOk() (*ForumTopicModelSearchV1, bool)`

GetLastTopicOk returns a tuple with the LastTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTopic

`func (o *ForumForumModelListV1) SetLastTopic(v ForumTopicModelSearchV1)`

SetLastTopic sets LastTopic field to given value.

### HasLastTopic

`func (o *ForumForumModelListV1) HasLastTopic() bool`

HasLastTopic returns a boolean if a field has been set.

### GetAdmin

`func (o *ForumForumModelListV1) GetAdmin() ForumForumModelListV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ForumForumModelListV1) GetAdminOk() (*ForumForumModelListV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ForumForumModelListV1) SetAdmin(v ForumForumModelListV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ForumForumModelListV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


