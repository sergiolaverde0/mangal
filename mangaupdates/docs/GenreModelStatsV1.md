# GenreModelStatsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Genre** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**GenreModelStatsV1Stats**](GenreModelStatsV1Stats.md) |  | [optional] 
**Demographic** | Pointer to **bool** |  | [optional] 

## Methods

### NewGenreModelStatsV1

`func NewGenreModelStatsV1() *GenreModelStatsV1`

NewGenreModelStatsV1 instantiates a new GenreModelStatsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenreModelStatsV1WithDefaults

`func NewGenreModelStatsV1WithDefaults() *GenreModelStatsV1`

NewGenreModelStatsV1WithDefaults instantiates a new GenreModelStatsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenreModelStatsV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenreModelStatsV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenreModelStatsV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GenreModelStatsV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGenre

`func (o *GenreModelStatsV1) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *GenreModelStatsV1) GetGenreOk() (*string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *GenreModelStatsV1) SetGenre(v string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *GenreModelStatsV1) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetDescription

`func (o *GenreModelStatsV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenreModelStatsV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenreModelStatsV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenreModelStatsV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStats

`func (o *GenreModelStatsV1) GetStats() GenreModelStatsV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GenreModelStatsV1) GetStatsOk() (*GenreModelStatsV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GenreModelStatsV1) SetStats(v GenreModelStatsV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GenreModelStatsV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDemographic

`func (o *GenreModelStatsV1) GetDemographic() bool`

GetDemographic returns the Demographic field if non-nil, zero value otherwise.

### GetDemographicOk

`func (o *GenreModelStatsV1) GetDemographicOk() (*bool, bool)`

GetDemographicOk returns a tuple with the Demographic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographic

`func (o *GenreModelStatsV1) SetDemographic(v bool)`

SetDemographic sets Demographic field to given value.

### HasDemographic

`func (o *GenreModelStatsV1) HasDemographic() bool`

HasDemographic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


