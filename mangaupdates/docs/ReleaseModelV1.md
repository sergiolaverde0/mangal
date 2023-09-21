# ReleaseModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**Chapter** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]ReleaseModelV1Groups**](ReleaseModelV1Groups.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**DownloadNotes** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**TimeAdded** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Admin** | Pointer to [**ReleaseModelV1Admin**](ReleaseModelV1Admin.md) |  | [optional] 

## Methods

### NewReleaseModelV1

`func NewReleaseModelV1() *ReleaseModelV1`

NewReleaseModelV1 instantiates a new ReleaseModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseModelV1WithDefaults

`func NewReleaseModelV1WithDefaults() *ReleaseModelV1`

NewReleaseModelV1WithDefaults instantiates a new ReleaseModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReleaseModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReleaseModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReleaseModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReleaseModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ReleaseModelV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseModelV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseModelV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseModelV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVolume

`func (o *ReleaseModelV1) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ReleaseModelV1) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ReleaseModelV1) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ReleaseModelV1) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetChapter

`func (o *ReleaseModelV1) GetChapter() string`

GetChapter returns the Chapter field if non-nil, zero value otherwise.

### GetChapterOk

`func (o *ReleaseModelV1) GetChapterOk() (*string, bool)`

GetChapterOk returns a tuple with the Chapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter

`func (o *ReleaseModelV1) SetChapter(v string)`

SetChapter sets Chapter field to given value.

### HasChapter

`func (o *ReleaseModelV1) HasChapter() bool`

HasChapter returns a boolean if a field has been set.

### GetGroups

`func (o *ReleaseModelV1) GetGroups() []ReleaseModelV1Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ReleaseModelV1) GetGroupsOk() (*[]ReleaseModelV1Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ReleaseModelV1) SetGroups(v []ReleaseModelV1Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ReleaseModelV1) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ReleaseModelV1) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ReleaseModelV1) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ReleaseModelV1) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ReleaseModelV1) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDownloadNotes

`func (o *ReleaseModelV1) GetDownloadNotes() string`

GetDownloadNotes returns the DownloadNotes field if non-nil, zero value otherwise.

### GetDownloadNotesOk

`func (o *ReleaseModelV1) GetDownloadNotesOk() (*string, bool)`

GetDownloadNotesOk returns a tuple with the DownloadNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadNotes

`func (o *ReleaseModelV1) SetDownloadNotes(v string)`

SetDownloadNotes sets DownloadNotes field to given value.

### HasDownloadNotes

`func (o *ReleaseModelV1) HasDownloadNotes() bool`

HasDownloadNotes returns a boolean if a field has been set.

### GetComment

`func (o *ReleaseModelV1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReleaseModelV1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReleaseModelV1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReleaseModelV1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ReleaseModelV1) GetTimeAdded() TimeV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ReleaseModelV1) GetTimeAddedOk() (*TimeV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ReleaseModelV1) SetTimeAdded(v TimeV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ReleaseModelV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetAdmin

`func (o *ReleaseModelV1) GetAdmin() ReleaseModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ReleaseModelV1) GetAdminOk() (*ReleaseModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ReleaseModelV1) SetAdmin(v ReleaseModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ReleaseModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


