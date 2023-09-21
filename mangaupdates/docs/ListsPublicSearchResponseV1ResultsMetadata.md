# ListsPublicSearchResponseV1ResultsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRating** | Pointer to **float32** |  | [optional] 
**UserComment** | Pointer to [**ListsPublicSearchResponseV1ResultsMetadataUserComment**](ListsPublicSearchResponseV1ResultsMetadataUserComment.md) |  | [optional] 
**UserList** | Pointer to [**ListsSeriesModelV1**](ListsSeriesModelV1.md) |  | [optional] 

## Methods

### NewListsPublicSearchResponseV1ResultsMetadata

`func NewListsPublicSearchResponseV1ResultsMetadata() *ListsPublicSearchResponseV1ResultsMetadata`

NewListsPublicSearchResponseV1ResultsMetadata instantiates a new ListsPublicSearchResponseV1ResultsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsPublicSearchResponseV1ResultsMetadataWithDefaults

`func NewListsPublicSearchResponseV1ResultsMetadataWithDefaults() *ListsPublicSearchResponseV1ResultsMetadata`

NewListsPublicSearchResponseV1ResultsMetadataWithDefaults instantiates a new ListsPublicSearchResponseV1ResultsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRating

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserRating() float32`

GetUserRating returns the UserRating field if non-nil, zero value otherwise.

### GetUserRatingOk

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserRatingOk() (*float32, bool)`

GetUserRatingOk returns a tuple with the UserRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRating

`func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserRating(v float32)`

SetUserRating sets UserRating field to given value.

### HasUserRating

`func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserRating() bool`

HasUserRating returns a boolean if a field has been set.

### GetUserComment

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserComment() ListsPublicSearchResponseV1ResultsMetadataUserComment`

GetUserComment returns the UserComment field if non-nil, zero value otherwise.

### GetUserCommentOk

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserCommentOk() (*ListsPublicSearchResponseV1ResultsMetadataUserComment, bool)`

GetUserCommentOk returns a tuple with the UserComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserComment

`func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserComment(v ListsPublicSearchResponseV1ResultsMetadataUserComment)`

SetUserComment sets UserComment field to given value.

### HasUserComment

`func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserComment() bool`

HasUserComment returns a boolean if a field has been set.

### GetUserList

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserList() ListsSeriesModelV1`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *ListsPublicSearchResponseV1ResultsMetadata) GetUserListOk() (*ListsSeriesModelV1, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *ListsPublicSearchResponseV1ResultsMetadata) SetUserList(v ListsSeriesModelV1)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *ListsPublicSearchResponseV1ResultsMetadata) HasUserList() bool`

HasUserList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


