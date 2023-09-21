# AuthorsModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]AuthorsModelV1Associated**](AuthorsModelV1Associated.md) |  | [optional] 
**Image** | Pointer to [**ImageModelV1**](ImageModelV1.md) |  | [optional] 
**Actualname** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**BirthdayModelV1**](BirthdayModelV1.md) |  | [optional] 
**Birthplace** | Pointer to **string** |  | [optional] 
**Bloodtype** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Stats** | Pointer to [**AuthorsModelV1Stats**](AuthorsModelV1Stats.md) |  | [optional] 
**Social** | Pointer to [**AuthorsModelV1Social**](AuthorsModelV1Social.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**Admin** | Pointer to [**AuthorsModelV1Admin**](AuthorsModelV1Admin.md) |  | [optional] 

## Methods

### NewAuthorsModelV1

`func NewAuthorsModelV1() *AuthorsModelV1`

NewAuthorsModelV1 instantiates a new AuthorsModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsModelV1WithDefaults

`func NewAuthorsModelV1WithDefaults() *AuthorsModelV1`

NewAuthorsModelV1WithDefaults instantiates a new AuthorsModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorsModelV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorsModelV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorsModelV1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorsModelV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthorsModelV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorsModelV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorsModelV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorsModelV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *AuthorsModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthorsModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthorsModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuthorsModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAssociated

`func (o *AuthorsModelV1) GetAssociated() []AuthorsModelV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *AuthorsModelV1) GetAssociatedOk() (*[]AuthorsModelV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *AuthorsModelV1) SetAssociated(v []AuthorsModelV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *AuthorsModelV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetImage

`func (o *AuthorsModelV1) GetImage() ImageModelV1`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AuthorsModelV1) GetImageOk() (*ImageModelV1, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AuthorsModelV1) SetImage(v ImageModelV1)`

SetImage sets Image field to given value.

### HasImage

`func (o *AuthorsModelV1) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetActualname

`func (o *AuthorsModelV1) GetActualname() string`

GetActualname returns the Actualname field if non-nil, zero value otherwise.

### GetActualnameOk

`func (o *AuthorsModelV1) GetActualnameOk() (*string, bool)`

GetActualnameOk returns a tuple with the Actualname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualname

`func (o *AuthorsModelV1) SetActualname(v string)`

SetActualname sets Actualname field to given value.

### HasActualname

`func (o *AuthorsModelV1) HasActualname() bool`

HasActualname returns a boolean if a field has been set.

### GetBirthday

`func (o *AuthorsModelV1) GetBirthday() BirthdayModelV1`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *AuthorsModelV1) GetBirthdayOk() (*BirthdayModelV1, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *AuthorsModelV1) SetBirthday(v BirthdayModelV1)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *AuthorsModelV1) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetBirthplace

`func (o *AuthorsModelV1) GetBirthplace() string`

GetBirthplace returns the Birthplace field if non-nil, zero value otherwise.

### GetBirthplaceOk

`func (o *AuthorsModelV1) GetBirthplaceOk() (*string, bool)`

GetBirthplaceOk returns a tuple with the Birthplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplace

`func (o *AuthorsModelV1) SetBirthplace(v string)`

SetBirthplace sets Birthplace field to given value.

### HasBirthplace

`func (o *AuthorsModelV1) HasBirthplace() bool`

HasBirthplace returns a boolean if a field has been set.

### GetBloodtype

`func (o *AuthorsModelV1) GetBloodtype() string`

GetBloodtype returns the Bloodtype field if non-nil, zero value otherwise.

### GetBloodtypeOk

`func (o *AuthorsModelV1) GetBloodtypeOk() (*string, bool)`

GetBloodtypeOk returns a tuple with the Bloodtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloodtype

`func (o *AuthorsModelV1) SetBloodtype(v string)`

SetBloodtype sets Bloodtype field to given value.

### HasBloodtype

`func (o *AuthorsModelV1) HasBloodtype() bool`

HasBloodtype returns a boolean if a field has been set.

### GetGender

`func (o *AuthorsModelV1) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AuthorsModelV1) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AuthorsModelV1) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AuthorsModelV1) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetGenres

`func (o *AuthorsModelV1) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AuthorsModelV1) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AuthorsModelV1) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AuthorsModelV1) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetStats

`func (o *AuthorsModelV1) GetStats() AuthorsModelV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AuthorsModelV1) GetStatsOk() (*AuthorsModelV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AuthorsModelV1) SetStats(v AuthorsModelV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AuthorsModelV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSocial

`func (o *AuthorsModelV1) GetSocial() AuthorsModelV1Social`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *AuthorsModelV1) GetSocialOk() (*AuthorsModelV1Social, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *AuthorsModelV1) SetSocial(v AuthorsModelV1Social)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *AuthorsModelV1) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetComments

`func (o *AuthorsModelV1) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AuthorsModelV1) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AuthorsModelV1) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AuthorsModelV1) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorsModelV1) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorsModelV1) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorsModelV1) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorsModelV1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAddedBy

`func (o *AuthorsModelV1) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *AuthorsModelV1) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *AuthorsModelV1) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *AuthorsModelV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetAdmin

`func (o *AuthorsModelV1) GetAdmin() AuthorsModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AuthorsModelV1) GetAdminOk() (*AuthorsModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AuthorsModelV1) SetAdmin(v AuthorsModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AuthorsModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


