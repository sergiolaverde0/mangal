# ForumPostReportModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **int64** |  | [optional] 
**TopicId** | Pointer to **int64** |  | [optional] 
**Topic** | Pointer to [**ForumTopicModelSearchV1**](ForumTopicModelSearchV1.md) |  | [optional] 
**PostId** | Pointer to **int64** |  | [optional] 
**Post** | Pointer to [**ForumPostModelSearchV1**](ForumPostModelSearchV1.md) |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewForumPostReportModelV1

`func NewForumPostReportModelV1() *ForumPostReportModelV1`

NewForumPostReportModelV1 instantiates a new ForumPostReportModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPostReportModelV1WithDefaults

`func NewForumPostReportModelV1WithDefaults() *ForumPostReportModelV1`

NewForumPostReportModelV1WithDefaults instantiates a new ForumPostReportModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *ForumPostReportModelV1) GetReportId() int64`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ForumPostReportModelV1) GetReportIdOk() (*int64, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ForumPostReportModelV1) SetReportId(v int64)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *ForumPostReportModelV1) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetTopicId

`func (o *ForumPostReportModelV1) GetTopicId() int64`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *ForumPostReportModelV1) GetTopicIdOk() (*int64, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *ForumPostReportModelV1) SetTopicId(v int64)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *ForumPostReportModelV1) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### GetTopic

`func (o *ForumPostReportModelV1) GetTopic() ForumTopicModelSearchV1`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumPostReportModelV1) GetTopicOk() (*ForumTopicModelSearchV1, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumPostReportModelV1) SetTopic(v ForumTopicModelSearchV1)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumPostReportModelV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPostId

`func (o *ForumPostReportModelV1) GetPostId() int64`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *ForumPostReportModelV1) GetPostIdOk() (*int64, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *ForumPostReportModelV1) SetPostId(v int64)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *ForumPostReportModelV1) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetPost

`func (o *ForumPostReportModelV1) GetPost() ForumPostModelSearchV1`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *ForumPostReportModelV1) GetPostOk() (*ForumPostModelSearchV1, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *ForumPostReportModelV1) SetPost(v ForumPostModelSearchV1)`

SetPost sets Post field to given value.

### HasPost

`func (o *ForumPostReportModelV1) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetUserId

`func (o *ForumPostReportModelV1) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ForumPostReportModelV1) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ForumPostReportModelV1) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ForumPostReportModelV1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUser

`func (o *ForumPostReportModelV1) GetUser() UserModelSearchV1`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ForumPostReportModelV1) GetUserOk() (*UserModelSearchV1, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ForumPostReportModelV1) SetUser(v UserModelSearchV1)`

SetUser sets User field to given value.

### HasUser

`func (o *ForumPostReportModelV1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetReason

`func (o *ForumPostReportModelV1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ForumPostReportModelV1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ForumPostReportModelV1) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ForumPostReportModelV1) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


