# SeriesHistorySearchResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalHits** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**PerPage** | Pointer to **int64** |  | [optional] 
**SeriesTitle** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]SeriesHistorySearchResponseV1Results**](SeriesHistorySearchResponseV1Results.md) |  | [optional] 

## Methods

### NewSeriesHistorySearchResponseV1

`func NewSeriesHistorySearchResponseV1() *SeriesHistorySearchResponseV1`

NewSeriesHistorySearchResponseV1 instantiates a new SeriesHistorySearchResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesHistorySearchResponseV1WithDefaults

`func NewSeriesHistorySearchResponseV1WithDefaults() *SeriesHistorySearchResponseV1`

NewSeriesHistorySearchResponseV1WithDefaults instantiates a new SeriesHistorySearchResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalHits

`func (o *SeriesHistorySearchResponseV1) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SeriesHistorySearchResponseV1) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SeriesHistorySearchResponseV1) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SeriesHistorySearchResponseV1) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetPage

`func (o *SeriesHistorySearchResponseV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SeriesHistorySearchResponseV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SeriesHistorySearchResponseV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *SeriesHistorySearchResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *SeriesHistorySearchResponseV1) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *SeriesHistorySearchResponseV1) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *SeriesHistorySearchResponseV1) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *SeriesHistorySearchResponseV1) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetSeriesTitle

`func (o *SeriesHistorySearchResponseV1) GetSeriesTitle() string`

GetSeriesTitle returns the SeriesTitle field if non-nil, zero value otherwise.

### GetSeriesTitleOk

`func (o *SeriesHistorySearchResponseV1) GetSeriesTitleOk() (*string, bool)`

GetSeriesTitleOk returns a tuple with the SeriesTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTitle

`func (o *SeriesHistorySearchResponseV1) SetSeriesTitle(v string)`

SetSeriesTitle sets SeriesTitle field to given value.

### HasSeriesTitle

`func (o *SeriesHistorySearchResponseV1) HasSeriesTitle() bool`

HasSeriesTitle returns a boolean if a field has been set.

### GetResults

`func (o *SeriesHistorySearchResponseV1) GetResults() []SeriesHistorySearchResponseV1Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SeriesHistorySearchResponseV1) GetResultsOk() (*[]SeriesHistorySearchResponseV1Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SeriesHistorySearchResponseV1) SetResults(v []SeriesHistorySearchResponseV1Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *SeriesHistorySearchResponseV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


