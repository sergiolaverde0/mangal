# ListsSearchResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**List** | Pointer to [**ListsModelV1**](ListsModelV1.md) |  | [optional] 
**Results** | Pointer to [**[]ListsSearchResponseV1Results**](ListsSearchResponseV1Results.md) |  | [optional] 

## Methods

### NewListsSearchResponseV1

`func NewListsSearchResponseV1() *ListsSearchResponseV1`

NewListsSearchResponseV1 instantiates a new ListsSearchResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsSearchResponseV1WithDefaults

`func NewListsSearchResponseV1WithDefaults() *ListsSearchResponseV1`

NewListsSearchResponseV1WithDefaults instantiates a new ListsSearchResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *ListsSearchResponseV1) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *ListsSearchResponseV1) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *ListsSearchResponseV1) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *ListsSearchResponseV1) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetPage

`func (o *ListsSearchResponseV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListsSearchResponseV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListsSearchResponseV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListsSearchResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *ListsSearchResponseV1) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListsSearchResponseV1) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListsSearchResponseV1) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *ListsSearchResponseV1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetList

`func (o *ListsSearchResponseV1) GetList() ListsModelV1`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ListsSearchResponseV1) GetListOk() (*ListsModelV1, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ListsSearchResponseV1) SetList(v ListsModelV1)`

SetList sets List field to given value.

### HasList

`func (o *ListsSearchResponseV1) HasList() bool`

HasList returns a boolean if a field has been set.

### GetResults

`func (o *ListsSearchResponseV1) GetResults() []ListsSearchResponseV1Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListsSearchResponseV1) GetResultsOk() (*[]ListsSearchResponseV1Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListsSearchResponseV1) SetResults(v []ListsSearchResponseV1Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListsSearchResponseV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


