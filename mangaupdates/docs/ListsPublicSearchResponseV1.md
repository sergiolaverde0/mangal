# ListsPublicSearchResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**List** | Pointer to [**ListsModelV1**](ListsModelV1.md) |  | [optional] 
**Results** | Pointer to [**[]ListsPublicSearchResponseV1Results**](ListsPublicSearchResponseV1Results.md) |  | [optional] 

## Methods

### NewListsPublicSearchResponseV1

`func NewListsPublicSearchResponseV1() *ListsPublicSearchResponseV1`

NewListsPublicSearchResponseV1 instantiates a new ListsPublicSearchResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsPublicSearchResponseV1WithDefaults

`func NewListsPublicSearchResponseV1WithDefaults() *ListsPublicSearchResponseV1`

NewListsPublicSearchResponseV1WithDefaults instantiates a new ListsPublicSearchResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *ListsPublicSearchResponseV1) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *ListsPublicSearchResponseV1) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *ListsPublicSearchResponseV1) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *ListsPublicSearchResponseV1) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetPage

`func (o *ListsPublicSearchResponseV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListsPublicSearchResponseV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListsPublicSearchResponseV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListsPublicSearchResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ListsPublicSearchResponseV1) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListsPublicSearchResponseV1) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListsPublicSearchResponseV1) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListsPublicSearchResponseV1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetList

`func (o *ListsPublicSearchResponseV1) GetList() ListsModelV1`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ListsPublicSearchResponseV1) GetListOk() (*ListsModelV1, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ListsPublicSearchResponseV1) SetList(v ListsModelV1)`

SetList sets List field to given value.

### HasList

`func (o *ListsPublicSearchResponseV1) HasList() bool`

HasList returns a boolean if a field has been set.

### GetResults

`func (o *ListsPublicSearchResponseV1) GetResults() []ListsPublicSearchResponseV1Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListsPublicSearchResponseV1) GetResultsOk() (*[]ListsPublicSearchResponseV1Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListsPublicSearchResponseV1) SetResults(v []ListsPublicSearchResponseV1Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListsPublicSearchResponseV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


