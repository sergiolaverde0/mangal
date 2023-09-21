# SeriesCommentModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SeriesId** | Pointer to **int64** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**SeriesCommentModelV1Author**](SeriesCommentModelV1Author.md) |  | [optional] 
**Useful** | Pointer to **int64** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**TimeUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Admin** | Pointer to [**SeriesCommentModelV1Admin**](SeriesCommentModelV1Admin.md) |  | [optional] 

## Methods

### NewSeriesCommentModelV1

`func NewSeriesCommentModelV1() *SeriesCommentModelV1`

NewSeriesCommentModelV1 instantiates a new SeriesCommentModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesCommentModelV1WithDefaults

`func NewSeriesCommentModelV1WithDefaults() *SeriesCommentModelV1`

NewSeriesCommentModelV1WithDefaults instantiates a new SeriesCommentModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeriesCommentModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeriesCommentModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeriesCommentModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SeriesCommentModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeriesId

`func (o *SeriesCommentModelV1) GetSeriesId() int64`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *SeriesCommentModelV1) GetSeriesIdOk() (*int64, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *SeriesCommentModelV1) SetSeriesId(v int64)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *SeriesCommentModelV1) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSubject

`func (o *SeriesCommentModelV1) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SeriesCommentModelV1) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SeriesCommentModelV1) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SeriesCommentModelV1) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContent

`func (o *SeriesCommentModelV1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SeriesCommentModelV1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SeriesCommentModelV1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SeriesCommentModelV1) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAuthor

`func (o *SeriesCommentModelV1) GetAuthor() SeriesCommentModelV1Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SeriesCommentModelV1) GetAuthorOk() (*SeriesCommentModelV1Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SeriesCommentModelV1) SetAuthor(v SeriesCommentModelV1Author)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SeriesCommentModelV1) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetUseful

`func (o *SeriesCommentModelV1) GetUseful() int64`

GetUseful returns the Useful field if non-nil, zero value otherwise.

### GetUsefulOk

`func (o *SeriesCommentModelV1) GetUsefulOk() (*int64, bool)`

GetUsefulOk returns a tuple with the Useful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseful

`func (o *SeriesCommentModelV1) SetUseful(v int64)`

SetUseful sets Useful field to given value.

### HasUseful

`func (o *SeriesCommentModelV1) HasUseful() bool`

HasUseful returns a boolean if a field has been set.

### GetTimeAdded

`func (o *SeriesCommentModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *SeriesCommentModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *SeriesCommentModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *SeriesCommentModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetTimeUpdated

`func (o *SeriesCommentModelV1) GetTimeUpdated() TimeV1`

GetTimeUpdated returns the TimeUpdated field if non-nil, zero value otherwise.

### GetTimeUpdatedOk

`func (o *SeriesCommentModelV1) GetTimeUpdatedOk() (*TimeV1, bool)`

GetTimeUpdatedOk returns a tuple with the TimeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUpdated

`func (o *SeriesCommentModelV1) SetTimeUpdated(v TimeV1)`

SetTimeUpdated sets TimeUpdated field to given value.

### HasTimeUpdated

`func (o *SeriesCommentModelV1) HasTimeUpdated() bool`

HasTimeUpdated returns a boolean if a field has been set.

### GetAdmin

`func (o *SeriesCommentModelV1) GetAdmin() SeriesCommentModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SeriesCommentModelV1) GetAdminOk() (*SeriesCommentModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SeriesCommentModelV1) SetAdmin(v SeriesCommentModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SeriesCommentModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


