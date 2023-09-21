# ForumPostModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **int64** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to [**ForumPostModelV1Topic**](ForumPostModelV1Topic.md) |  | [optional] 
**Author** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**ReplyTo** | Pointer to [**ForumPostModelV1ReplyTo**](ForumPostModelV1ReplyTo.md) |  | [optional] 
**LastEdit** | Pointer to [**ForumPostModelV1LastEdit**](ForumPostModelV1LastEdit.md) |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewForumPostModelV1

`func NewForumPostModelV1() *ForumPostModelV1`

NewForumPostModelV1 instantiates a new ForumPostModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPostModelV1WithDefaults

`func NewForumPostModelV1WithDefaults() *ForumPostModelV1`

NewForumPostModelV1WithDefaults instantiates a new ForumPostModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *ForumPostModelV1) GetPostId() int64`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *ForumPostModelV1) GetPostIdOk() (*int64, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *ForumPostModelV1) SetPostId(v int64)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *ForumPostModelV1) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetBody

`func (o *ForumPostModelV1) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ForumPostModelV1) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ForumPostModelV1) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ForumPostModelV1) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetTopic

`func (o *ForumPostModelV1) GetTopic() ForumPostModelV1Topic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ForumPostModelV1) GetTopicOk() (*ForumPostModelV1Topic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *ForumPostModelV1) SetTopic(v ForumPostModelV1Topic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *ForumPostModelV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetAuthor

`func (o *ForumPostModelV1) GetAuthor() UserModelSearchV1`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ForumPostModelV1) GetAuthorOk() (*UserModelSearchV1, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ForumPostModelV1) SetAuthor(v UserModelSearchV1)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ForumPostModelV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetReplyTo

`func (o *ForumPostModelV1) GetReplyTo() ForumPostModelV1ReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *ForumPostModelV1) GetReplyToOk() (*ForumPostModelV1ReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *ForumPostModelV1) SetReplyTo(v ForumPostModelV1ReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *ForumPostModelV1) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetLastEdit

`func (o *ForumPostModelV1) GetLastEdit() ForumPostModelV1LastEdit`

GetLastEdit returns the LastEdit field if non-nil, zero value otherwise.

### GetLastEditOk

`func (o *ForumPostModelV1) GetLastEditOk() (*ForumPostModelV1LastEdit, bool)`

GetLastEditOk returns a tuple with the LastEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEdit

`func (o *ForumPostModelV1) SetLastEdit(v ForumPostModelV1LastEdit)`

SetLastEdit sets LastEdit field to given value.

### HasLastEdit

`func (o *ForumPostModelV1) HasLastEdit() bool`

HasLastEdit returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ForumPostModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ForumPostModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ForumPostModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ForumPostModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


