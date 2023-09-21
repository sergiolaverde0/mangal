# ReleaseModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**Chapter** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]ReleaseModelUpdateV1Groups**](ReleaseModelUpdateV1Groups.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**DownloadNotes** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**TimeAdded** | Pointer to [**TimeUpdateV1**](TimeUpdateV1.md) |  | [optional] 
**Admin** | Pointer to [**ReleaseModelUpdateV1Admin**](ReleaseModelUpdateV1Admin.md) |  | [optional] 

## Methods

### NewReleaseModelUpdateV1

`func NewReleaseModelUpdateV1() *ReleaseModelUpdateV1`

NewReleaseModelUpdateV1 instantiates a new ReleaseModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseModelUpdateV1WithDefaults

`func NewReleaseModelUpdateV1WithDefaults() *ReleaseModelUpdateV1`

NewReleaseModelUpdateV1WithDefaults instantiates a new ReleaseModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReleaseModelUpdateV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReleaseModelUpdateV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReleaseModelUpdateV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReleaseModelUpdateV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVolume

`func (o *ReleaseModelUpdateV1) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ReleaseModelUpdateV1) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ReleaseModelUpdateV1) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ReleaseModelUpdateV1) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetChapter

`func (o *ReleaseModelUpdateV1) GetChapter() string`

GetChapter returns the Chapter field if non-nil, zero value otherwise.

### GetChapterOk

`func (o *ReleaseModelUpdateV1) GetChapterOk() (*string, bool)`

GetChapterOk returns a tuple with the Chapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter

`func (o *ReleaseModelUpdateV1) SetChapter(v string)`

SetChapter sets Chapter field to given value.

### HasChapter

`func (o *ReleaseModelUpdateV1) HasChapter() bool`

HasChapter returns a boolean if a field has been set.

### GetGroups

`func (o *ReleaseModelUpdateV1) GetGroups() []ReleaseModelUpdateV1Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ReleaseModelUpdateV1) GetGroupsOk() (*[]ReleaseModelUpdateV1Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ReleaseModelUpdateV1) SetGroups(v []ReleaseModelUpdateV1Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ReleaseModelUpdateV1) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ReleaseModelUpdateV1) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ReleaseModelUpdateV1) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ReleaseModelUpdateV1) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ReleaseModelUpdateV1) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDownloadNotes

`func (o *ReleaseModelUpdateV1) GetDownloadNotes() string`

GetDownloadNotes returns the DownloadNotes field if non-nil, zero value otherwise.

### GetDownloadNotesOk

`func (o *ReleaseModelUpdateV1) GetDownloadNotesOk() (*string, bool)`

GetDownloadNotesOk returns a tuple with the DownloadNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadNotes

`func (o *ReleaseModelUpdateV1) SetDownloadNotes(v string)`

SetDownloadNotes sets DownloadNotes field to given value.

### HasDownloadNotes

`func (o *ReleaseModelUpdateV1) HasDownloadNotes() bool`

HasDownloadNotes returns a boolean if a field has been set.

### GetComment

`func (o *ReleaseModelUpdateV1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ReleaseModelUpdateV1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ReleaseModelUpdateV1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ReleaseModelUpdateV1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTimeAdded

`func (o *ReleaseModelUpdateV1) GetTimeAdded() TimeUpdateV1`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *ReleaseModelUpdateV1) GetTimeAddedOk() (*TimeUpdateV1, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *ReleaseModelUpdateV1) SetTimeAdded(v TimeUpdateV1)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *ReleaseModelUpdateV1) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetAdmin

`func (o *ReleaseModelUpdateV1) GetAdmin() ReleaseModelUpdateV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ReleaseModelUpdateV1) GetAdminOk() (*ReleaseModelUpdateV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ReleaseModelUpdateV1) SetAdmin(v ReleaseModelUpdateV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ReleaseModelUpdateV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


