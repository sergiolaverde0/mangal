# ForumSearchRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchBy** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Since** | Pointer to **int64** |  | [optional] 
**AfterId** | Pointer to **int64** |  | [optional] 
**BeforeId** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Perpage** | Pointer to **int64** |  | [optional] 
**ByUserId** | Pointer to **int64** |  | [optional] 
**FilterUserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewForumSearchRequestV1

`func NewForumSearchRequestV1() *ForumSearchRequestV1`

NewForumSearchRequestV1 instantiates a new ForumSearchRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumSearchRequestV1WithDefaults

`func NewForumSearchRequestV1WithDefaults() *ForumSearchRequestV1`

NewForumSearchRequestV1WithDefaults instantiates a new ForumSearchRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchBy

`func (o *ForumSearchRequestV1) GetSearchBy() string`

GetSearchBy returns the SearchBy field if non-nil, zero value otherwise.

### GetSearchByOk

`func (o *ForumSearchRequestV1) GetSearchByOk() (*string, bool)`

GetSearchByOk returns a tuple with the SearchBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBy

`func (o *ForumSearchRequestV1) SetSearchBy(v string)`

SetSearchBy sets SearchBy field to given value.

### HasSearchBy

`func (o *ForumSearchRequestV1) HasSearchBy() bool`

HasSearchBy returns a boolean if a field has been set.

### GetMethod

`func (o *ForumSearchRequestV1) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ForumSearchRequestV1) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ForumSearchRequestV1) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ForumSearchRequestV1) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetSearch

`func (o *ForumSearchRequestV1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ForumSearchRequestV1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ForumSearchRequestV1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ForumSearchRequestV1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSince

`func (o *ForumSearchRequestV1) GetSince() int64`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *ForumSearchRequestV1) GetSinceOk() (*int64, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *ForumSearchRequestV1) SetSince(v int64)`

SetSince sets Since field to given value.

### HasSince

`func (o *ForumSearchRequestV1) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetAfterId

`func (o *ForumSearchRequestV1) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *ForumSearchRequestV1) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *ForumSearchRequestV1) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *ForumSearchRequestV1) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetBeforeId

`func (o *ForumSearchRequestV1) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *ForumSearchRequestV1) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *ForumSearchRequestV1) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *ForumSearchRequestV1) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetPage

`func (o *ForumSearchRequestV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ForumSearchRequestV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ForumSearchRequestV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ForumSearchRequestV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerpage

`func (o *ForumSearchRequestV1) GetPerpage() int64`

GetPerpage returns the Perpage field if non-nil, zero value otherwise.

### GetPerpageOk

`func (o *ForumSearchRequestV1) GetPerpageOk() (*int64, bool)`

GetPerpageOk returns a tuple with the Perpage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerpage

`func (o *ForumSearchRequestV1) SetPerpage(v int64)`

SetPerpage sets Perpage field to given value.

### HasPerpage

`func (o *ForumSearchRequestV1) HasPerpage() bool`

HasPerpage returns a boolean if a field has been set.

### GetByUserId

`func (o *ForumSearchRequestV1) GetByUserId() int64`

GetByUserId returns the ByUserId field if non-nil, zero value otherwise.

### GetByUserIdOk

`func (o *ForumSearchRequestV1) GetByUserIdOk() (*int64, bool)`

GetByUserIdOk returns a tuple with the ByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByUserId

`func (o *ForumSearchRequestV1) SetByUserId(v int64)`

SetByUserId sets ByUserId field to given value.

### HasByUserId

`func (o *ForumSearchRequestV1) HasByUserId() bool`

HasByUserId returns a boolean if a field has been set.

### GetFilterUserId

`func (o *ForumSearchRequestV1) GetFilterUserId() int64`

GetFilterUserId returns the FilterUserId field if non-nil, zero value otherwise.

### GetFilterUserIdOk

`func (o *ForumSearchRequestV1) GetFilterUserIdOk() (*int64, bool)`

GetFilterUserIdOk returns a tuple with the FilterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterUserId

`func (o *ForumSearchRequestV1) SetFilterUserId(v int64)`

SetFilterUserId sets FilterUserId field to given value.

### HasFilterUserId

`func (o *ForumSearchRequestV1) HasFilterUserId() bool`

HasFilterUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


