# CategoriesSearchResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]CategoriesSearchResponseV1Results**](CategoriesSearchResponseV1Results.md) |  | [optional] 

## Methods

### NewCategoriesSearchResponseV1

`func NewCategoriesSearchResponseV1() *CategoriesSearchResponseV1`

NewCategoriesSearchResponseV1 instantiates a new CategoriesSearchResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesSearchResponseV1WithDefaults

`func NewCategoriesSearchResponseV1WithDefaults() *CategoriesSearchResponseV1`

NewCategoriesSearchResponseV1WithDefaults instantiates a new CategoriesSearchResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *CategoriesSearchResponseV1) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *CategoriesSearchResponseV1) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *CategoriesSearchResponseV1) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *CategoriesSearchResponseV1) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetPage

`func (o *CategoriesSearchResponseV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CategoriesSearchResponseV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CategoriesSearchResponseV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *CategoriesSearchResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *CategoriesSearchResponseV1) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *CategoriesSearchResponseV1) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *CategoriesSearchResponseV1) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *CategoriesSearchResponseV1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetResults

`func (o *CategoriesSearchResponseV1) GetResults() []CategoriesSearchResponseV1Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CategoriesSearchResponseV1) GetResultsOk() (*[]CategoriesSearchResponseV1Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CategoriesSearchResponseV1) SetResults(v []CategoriesSearchResponseV1Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *CategoriesSearchResponseV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


