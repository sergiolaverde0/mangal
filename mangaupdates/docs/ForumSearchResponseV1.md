# ForumSearchResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**TopicResults** | Pointer to [**[]ForumSearchResponseV1TopicResults**](ForumSearchResponseV1TopicResults.md) |  | [optional] 
**PostResults** | Pointer to [**[]ForumSearchResponseV1PostResults**](ForumSearchResponseV1PostResults.md) |  | [optional] 

## Methods

### NewForumSearchResponseV1

`func NewForumSearchResponseV1() *ForumSearchResponseV1`

NewForumSearchResponseV1 instantiates a new ForumSearchResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumSearchResponseV1WithDefaults

`func NewForumSearchResponseV1WithDefaults() *ForumSearchResponseV1`

NewForumSearchResponseV1WithDefaults instantiates a new ForumSearchResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *ForumSearchResponseV1) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *ForumSearchResponseV1) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *ForumSearchResponseV1) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *ForumSearchResponseV1) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetPage

`func (o *ForumSearchResponseV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ForumSearchResponseV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ForumSearchResponseV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ForumSearchResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ForumSearchResponseV1) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ForumSearchResponseV1) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ForumSearchResponseV1) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ForumSearchResponseV1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTopicResults

`func (o *ForumSearchResponseV1) GetTopicResults() []ForumSearchResponseV1TopicResults`

GetTopicResults returns the TopicResults field if non-nil, zero value otherwise.

### GetTopicResultsOk

`func (o *ForumSearchResponseV1) GetTopicResultsOk() (*[]ForumSearchResponseV1TopicResults, bool)`

GetTopicResultsOk returns a tuple with the TopicResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicResults

`func (o *ForumSearchResponseV1) SetTopicResults(v []ForumSearchResponseV1TopicResults)`

SetTopicResults sets TopicResults field to given value.

### HasTopicResults

`func (o *ForumSearchResponseV1) HasTopicResults() bool`

HasTopicResults returns a boolean if a field has been set.

### GetPostResults

`func (o *ForumSearchResponseV1) GetPostResults() []ForumSearchResponseV1PostResults`

GetPostResults returns the PostResults field if non-nil, zero value otherwise.

### GetPostResultsOk

`func (o *ForumSearchResponseV1) GetPostResultsOk() (*[]ForumSearchResponseV1PostResults, bool)`

GetPostResultsOk returns a tuple with the PostResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostResults

`func (o *ForumSearchResponseV1) SetPostResults(v []ForumSearchResponseV1PostResults)`

SetPostResults sets PostResults field to given value.

### HasPostResults

`func (o *ForumSearchResponseV1) HasPostResults() bool`

HasPostResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


