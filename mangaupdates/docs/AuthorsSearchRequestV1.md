# AuthorsSearchRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Perpage** | Pointer to **int64** |  | [optional] 
**Letter** | Pointer to **string** |  | [optional] 
**Genre** | Pointer to **[]string** |  | [optional] 
**Orderby** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthorsSearchRequestV1

`func NewAuthorsSearchRequestV1() *AuthorsSearchRequestV1`

NewAuthorsSearchRequestV1 instantiates a new AuthorsSearchRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsSearchRequestV1WithDefaults

`func NewAuthorsSearchRequestV1WithDefaults() *AuthorsSearchRequestV1`

NewAuthorsSearchRequestV1WithDefaults instantiates a new AuthorsSearchRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *AuthorsSearchRequestV1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *AuthorsSearchRequestV1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *AuthorsSearchRequestV1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *AuthorsSearchRequestV1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetAddedBy

`func (o *AuthorsSearchRequestV1) GetAddedBy() int64`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *AuthorsSearchRequestV1) GetAddedByOk() (*int64, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *AuthorsSearchRequestV1) SetAddedBy(v int64)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *AuthorsSearchRequestV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetPage

`func (o *AuthorsSearchRequestV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AuthorsSearchRequestV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AuthorsSearchRequestV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *AuthorsSearchRequestV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerpage

`func (o *AuthorsSearchRequestV1) GetPerpage() int64`

GetPerpage returns the Perpage field if non-nil, zero value otherwise.

### GetPerpageOk

`func (o *AuthorsSearchRequestV1) GetPerpageOk() (*int64, bool)`

GetPerpageOk returns a tuple with the Perpage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerpage

`func (o *AuthorsSearchRequestV1) SetPerpage(v int64)`

SetPerpage sets Perpage field to given value.

### HasPerpage

`func (o *AuthorsSearchRequestV1) HasPerpage() bool`

HasPerpage returns a boolean if a field has been set.

### GetLetter

`func (o *AuthorsSearchRequestV1) GetLetter() string`

GetLetter returns the Letter field if non-nil, zero value otherwise.

### GetLetterOk

`func (o *AuthorsSearchRequestV1) GetLetterOk() (*string, bool)`

GetLetterOk returns a tuple with the Letter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetter

`func (o *AuthorsSearchRequestV1) SetLetter(v string)`

SetLetter sets Letter field to given value.

### HasLetter

`func (o *AuthorsSearchRequestV1) HasLetter() bool`

HasLetter returns a boolean if a field has been set.

### GetGenre

`func (o *AuthorsSearchRequestV1) GetGenre() []string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *AuthorsSearchRequestV1) GetGenreOk() (*[]string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *AuthorsSearchRequestV1) SetGenre(v []string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *AuthorsSearchRequestV1) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetOrderby

`func (o *AuthorsSearchRequestV1) GetOrderby() string`

GetOrderby returns the Orderby field if non-nil, zero value otherwise.

### GetOrderbyOk

`func (o *AuthorsSearchRequestV1) GetOrderbyOk() (*string, bool)`

GetOrderbyOk returns a tuple with the Orderby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderby

`func (o *AuthorsSearchRequestV1) SetOrderby(v string)`

SetOrderby sets Orderby field to given value.

### HasOrderby

`func (o *AuthorsSearchRequestV1) HasOrderby() bool`

HasOrderby returns a boolean if a field has been set.

### GetPending

`func (o *AuthorsSearchRequestV1) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *AuthorsSearchRequestV1) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *AuthorsSearchRequestV1) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *AuthorsSearchRequestV1) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


