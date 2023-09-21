# UserSubscribedTopicModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicId** | Pointer to **int64** |  | [optional] 
**Topic** | Pointer to [**ForumTopicModelSearchV1**](ForumTopicModelSearchV1.md) |  | [optional] 
**TimeSubscribedSince** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 

## Methods

### NewUserSubscribedTopicModelV1

`func NewUserSubscribedTopicModelV1() *UserSubscribedTopicModelV1`

NewUserSubscribedTopicModelV1 instantiates a new UserSubscribedTopicModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscribedTopicModelV1WithDefaults

`func NewUserSubscribedTopicModelV1WithDefaults() *UserSubscribedTopicModelV1`

NewUserSubscribedTopicModelV1WithDefaults instantiates a new UserSubscribedTopicModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicId

`func (o *UserSubscribedTopicModelV1) GetTopicId() int64`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *UserSubscribedTopicModelV1) GetTopicIdOk() (*int64, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *UserSubscribedTopicModelV1) SetTopicId(v int64)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *UserSubscribedTopicModelV1) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### GetTopic

`func (o *UserSubscribedTopicModelV1) GetTopic() ForumTopicModelSearchV1`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *UserSubscribedTopicModelV1) GetTopicOk() (*ForumTopicModelSearchV1, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *UserSubscribedTopicModelV1) SetTopic(v ForumTopicModelSearchV1)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *UserSubscribedTopicModelV1) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetTimeSubscribedSince

`func (o *UserSubscribedTopicModelV1) GetTimeSubscribedSince() TimeV1`

GetTimeSubscribedSince returns the TimeSubscribedSince field if non-nil, zero value otherwise.

### GetTimeSubscribedSinceOk

`func (o *UserSubscribedTopicModelV1) GetTimeSubscribedSinceOk() (*TimeV1, bool)`

GetTimeSubscribedSinceOk returns a tuple with the TimeSubscribedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSubscribedSince

`func (o *UserSubscribedTopicModelV1) SetTimeSubscribedSince(v TimeV1)`

SetTimeSubscribedSince sets TimeSubscribedSince field to given value.

### HasTimeSubscribedSince

`func (o *UserSubscribedTopicModelV1) HasTimeSubscribedSince() bool`

HasTimeSubscribedSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


