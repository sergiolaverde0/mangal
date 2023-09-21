# ForumPostModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **int64** |  | [optional] 
**BodyExcerpt** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to [**ForumPostModelSearchV1Topic**](ForumPostModelSearchV1Topic.md) |  | [optional] 
**Author** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewForumPostModelSearchV1

`func NewForumPostModelSearchV1() *ForumPostModelSearchV1`

NewForumPostModelSearchV1 instantiates a new ForumPostModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPostModelSearchV1WithDefaults

`func NewForumPostModelSearchV1WithDefaults() *ForumPostModelSearchV1`

NewForumPostModelSearchV1WithDefaults instantiates a new ForumPostModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *ForumPostModelSearchV1) GetPostId() int64`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *ForumPostModelSearchV1) GetPostIdOk() (*int64, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *ForumPostModelSearchV1) SetPostId(v int64)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *ForumPostModelSearchV1) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetBodyExcerpt

`func (o *ForumPostModelSearchV1) GetBodyExcerpt() string`

GetBodyExcerpt returns the BodyExcerpt field if non-nil, zero value otherwise.

### GetBodyExcerptOk

`func (o *ForumPostModelSearchV1) GetBodyExcerptOk() (*string, bool)`

GetBodyExcerptOk returns a tuple with the BodyExcerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyExcerpt

`func (o *ForumPostModelSearchV1) SetBodyExcerpt(v string)`

SetBodyExcerpt sets BodyExcerpt field to given value.

### HasBodyExcerpt

`func (o *ForumPostModelSearchV1) HasBodyExcerpt() bool`

HasBodyExcerpt returns a boolean if a field has been set.

### GetTopic

`func (o *ForumPostModelSearchV1) GetTopic() ForumPostModelSearchV1Topic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumPostModelSearchV1) GetTopicOk() (*ForumPostModelSearchV1Topic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumPostModelSearchV1) SetTopic(v ForumPostModelSearchV1Topic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumPostModelSearchV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetAuthor

`func (o *ForumPostModelSearchV1) GetAuthor() UserModelSearchV1`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ForumPostModelSearchV1) GetAuthorOk() (*UserModelSearchV1, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ForumPostModelSearchV1) SetAuthor(v UserModelSearchV1)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ForumPostModelSearchV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ForumPostModelSearchV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ForumPostModelSearchV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ForumPostModelSearchV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ForumPostModelSearchV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


