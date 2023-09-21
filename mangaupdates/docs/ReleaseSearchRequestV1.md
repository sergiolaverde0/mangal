# ReleaseSearchRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** |  | [optional] 
**SearchType** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to **int64** |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Perpage** | Pointer to **int64** |  | [optional] 
**Letter** | Pointer to **string** |  | [optional] 
**Orderby** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Asc** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**Pending** | Pointer to **bool** |  | [optional] 
**IncludeMetadata** | Pointer to **bool** |  | [optional] 

## Methods

### NewReleaseSearchRequestV1

`func NewReleaseSearchRequestV1() *ReleaseSearchRequestV1`

NewReleaseSearchRequestV1 instantiates a new ReleaseSearchRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseSearchRequestV1WithDefaults

`func NewReleaseSearchRequestV1WithDefaults() *ReleaseSearchRequestV1`

NewReleaseSearchRequestV1WithDefaults instantiates a new ReleaseSearchRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *ReleaseSearchRequestV1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ReleaseSearchRequestV1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ReleaseSearchRequestV1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ReleaseSearchRequestV1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSearchType

`func (o *ReleaseSearchRequestV1) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *ReleaseSearchRequestV1) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *ReleaseSearchRequestV1) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *ReleaseSearchRequestV1) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetAddedBy

`func (o *ReleaseSearchRequestV1) GetAddedBy() int64`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *ReleaseSearchRequestV1) GetAddedByOk() (*int64, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *ReleaseSearchRequestV1) SetAddedBy(v int64)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *ReleaseSearchRequestV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetPage

`func (o *ReleaseSearchRequestV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ReleaseSearchRequestV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ReleaseSearchRequestV1) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ReleaseSearchRequestV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerpage

`func (o *ReleaseSearchRequestV1) GetPerpage() int64`

GetPerpage returns the Perpage field if non-nil, zero value otherwise.

### GetPerpageOk

`func (o *ReleaseSearchRequestV1) GetPerpageOk() (*int64, bool)`

GetPerpageOk returns a tuple with the Perpage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerpage

`func (o *ReleaseSearchRequestV1) SetPerpage(v int64)`

SetPerpage sets Perpage field to given value.

### HasPerpage

`func (o *ReleaseSearchRequestV1) HasPerpage() bool`

HasPerpage returns a boolean if a field has been set.

### GetLetter

`func (o *ReleaseSearchRequestV1) GetLetter() string`

GetLetter returns the Letter field if non-nil, zero value otherwise.

### GetLetterOk

`func (o *ReleaseSearchRequestV1) GetLetterOk() (*string, bool)`

GetLetterOk returns a tuple with the Letter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetter

`func (o *ReleaseSearchRequestV1) SetLetter(v string)`

SetLetter sets Letter field to given value.

### HasLetter

`func (o *ReleaseSearchRequestV1) HasLetter() bool`

HasLetter returns a boolean if a field has been set.

### GetOrderby

`func (o *ReleaseSearchRequestV1) GetOrderby() string`

GetOrderby returns the Orderby field if non-nil, zero value otherwise.

### GetOrderbyOk

`func (o *ReleaseSearchRequestV1) GetOrderbyOk() (*string, bool)`

GetOrderbyOk returns a tuple with the Orderby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderby

`func (o *ReleaseSearchRequestV1) SetOrderby(v string)`

SetOrderby sets Orderby field to given value.

### HasOrderby

`func (o *ReleaseSearchRequestV1) HasOrderby() bool`

HasOrderby returns a boolean if a field has been set.

### GetStartDate

`func (o *ReleaseSearchRequestV1) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ReleaseSearchRequestV1) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ReleaseSearchRequestV1) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ReleaseSearchRequestV1) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ReleaseSearchRequestV1) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ReleaseSearchRequestV1) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ReleaseSearchRequestV1) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ReleaseSearchRequestV1) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAsc

`func (o *ReleaseSearchRequestV1) GetAsc() string`

GetAsc returns the Asc field if non-nil, zero value otherwise.

### GetAscOk

`func (o *ReleaseSearchRequestV1) GetAscOk() (*string, bool)`

GetAscOk returns a tuple with the Asc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsc

`func (o *ReleaseSearchRequestV1) SetAsc(v string)`

SetAsc sets Asc field to given value.

### HasAsc

`func (o *ReleaseSearchRequestV1) HasAsc() bool`

HasAsc returns a boolean if a field has been set.

### GetGroupId

`func (o *ReleaseSearchRequestV1) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ReleaseSearchRequestV1) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ReleaseSearchRequestV1) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ReleaseSearchRequestV1) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPending

`func (o *ReleaseSearchRequestV1) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ReleaseSearchRequestV1) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ReleaseSearchRequestV1) SetPending(v bool)`

SetPending sets Pending field to given value.

### HasPending

`func (o *ReleaseSearchRequestV1) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetIncludeMetadata

`func (o *ReleaseSearchRequestV1) GetIncludeMetadata() bool`

GetIncludeMetadata returns the IncludeMetadata field if non-nil, zero value otherwise.

### GetIncludeMetadataOk

`func (o *ReleaseSearchRequestV1) GetIncludeMetadataOk() (*bool, bool)`

GetIncludeMetadataOk returns a tuple with the IncludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadata

`func (o *ReleaseSearchRequestV1) SetIncludeMetadata(v bool)`

SetIncludeMetadata sets IncludeMetadata field to given value.

### HasIncludeMetadata

`func (o *ReleaseSearchRequestV1) HasIncludeMetadata() bool`

HasIncludeMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


