# AuthorsSearchResponseV1Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**AuthorsModelSearchV1**](AuthorsModelSearchV1.md) |  | [optional] 
**HitName** | Pointer to **string** |  | [optional] 
**HitGenre** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthorsSearchResponseV1Results

`func NewAuthorsSearchResponseV1Results() *AuthorsSearchResponseV1Results`

NewAuthorsSearchResponseV1Results instantiates a new AuthorsSearchResponseV1Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsSearchResponseV1ResultsWithDefaults

`func NewAuthorsSearchResponseV1ResultsWithDefaults() *AuthorsSearchResponseV1Results`

NewAuthorsSearchResponseV1ResultsWithDefaults instantiates a new AuthorsSearchResponseV1Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *AuthorsSearchResponseV1Results) GetRecord() AuthorsModelSearchV1`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AuthorsSearchResponseV1Results) GetRecordOk() (*AuthorsModelSearchV1, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AuthorsSearchResponseV1Results) SetRecord(v AuthorsModelSearchV1)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *AuthorsSearchResponseV1Results) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetHitName

`func (o *AuthorsSearchResponseV1Results) GetHitName() string`

GetHitName returns the HitName field if non-nil, zero value otherwise.

### GetHitNameOk

`func (o *AuthorsSearchResponseV1Results) GetHitNameOk() (*string, bool)`

GetHitNameOk returns a tuple with the HitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitName

`func (o *AuthorsSearchResponseV1Results) SetHitName(v string)`

SetHitName sets HitName field to given value.

### HasHitName

`func (o *AuthorsSearchResponseV1Results) HasHitName() bool`

HasHitName returns a boolean if a field has been set.

### GetHitGenre

`func (o *AuthorsSearchResponseV1Results) GetHitGenre() []string`

GetHitGenre returns the HitGenre field if non-nil, zero value otherwise.

### GetHitGenreOk

`func (o *AuthorsSearchResponseV1Results) GetHitGenreOk() (*[]string, bool)`

GetHitGenreOk returns a tuple with the HitGenre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitGenre

`func (o *AuthorsSearchResponseV1Results) SetHitGenre(v []string)`

SetHitGenre sets HitGenre field to given value.

### HasHitGenre

`func (o *AuthorsSearchResponseV1Results) HasHitGenre() bool`

HasHitGenre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


