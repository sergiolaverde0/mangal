# AuthorsModelSearchV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Stats** | Pointer to [**AuthorsModelSearchV1Stats**](AuthorsModelSearchV1Stats.md) |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 

## Methods

### NewAuthorsModelSearchV1

`func NewAuthorsModelSearchV1() *AuthorsModelSearchV1`

NewAuthorsModelSearchV1 instantiates a new AuthorsModelSearchV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsModelSearchV1WithDefaults

`func NewAuthorsModelSearchV1WithDefaults() *AuthorsModelSearchV1`

NewAuthorsModelSearchV1WithDefaults instantiates a new AuthorsModelSearchV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorsModelSearchV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorsModelSearchV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorsModelSearchV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorsModelSearchV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthorsModelSearchV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorsModelSearchV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorsModelSearchV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorsModelSearchV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *AuthorsModelSearchV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthorsModelSearchV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthorsModelSearchV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuthorsModelSearchV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetGenres

`func (o *AuthorsModelSearchV1) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AuthorsModelSearchV1) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AuthorsModelSearchV1) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AuthorsModelSearchV1) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetStats

`func (o *AuthorsModelSearchV1) GetStats() AuthorsModelSearchV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AuthorsModelSearchV1) GetStatsOk() (*AuthorsModelSearchV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AuthorsModelSearchV1) SetStats(v AuthorsModelSearchV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AuthorsModelSearchV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetAddedBy

`func (o *AuthorsModelSearchV1) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *AuthorsModelSearchV1) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *AuthorsModelSearchV1) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *AuthorsModelSearchV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


