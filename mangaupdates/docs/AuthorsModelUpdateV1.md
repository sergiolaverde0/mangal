# AuthorsModelUpdateV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]AuthorsModelUpdateV1Associated**](AuthorsModelUpdateV1Associated.md) |  | [optional] 
**Actualname** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**BirthdayModelV1**](BirthdayModelV1.md) |  | [optional] 
**Birthplace** | Pointer to **string** |  | [optional] 
**Bloodtype** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Social** | Pointer to [**AuthorsModelUpdateV1Social**](AuthorsModelUpdateV1Social.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to [**AuthorsModelUpdateV1Admin**](AuthorsModelUpdateV1Admin.md) |  | [optional] 

## Methods

### NewAuthorsModelUpdateV1

`func NewAuthorsModelUpdateV1() *AuthorsModelUpdateV1`

NewAuthorsModelUpdateV1 instantiates a new AuthorsModelUpdateV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsModelUpdateV1WithDefaults

`func NewAuthorsModelUpdateV1WithDefaults() *AuthorsModelUpdateV1`

NewAuthorsModelUpdateV1WithDefaults instantiates a new AuthorsModelUpdateV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorsModelUpdateV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorsModelUpdateV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorsModelUpdateV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorsModelUpdateV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssociated

`func (o *AuthorsModelUpdateV1) GetAssociated() []AuthorsModelUpdateV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *AuthorsModelUpdateV1) GetAssociatedOk() (*[]AuthorsModelUpdateV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *AuthorsModelUpdateV1) SetAssociated(v []AuthorsModelUpdateV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *AuthorsModelUpdateV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetActualname

`func (o *AuthorsModelUpdateV1) GetActualname() string`

GetActualname returns the Actualname field if non-nil, zero value otherwise.

### GetActualnameOk

`func (o *AuthorsModelUpdateV1) GetActualnameOk() (*string, bool)`

GetActualnameOk returns a tuple with the Actualname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualname

`func (o *AuthorsModelUpdateV1) SetActualname(v string)`

SetActualname sets Actualname field to given value.

### HasActualname

`func (o *AuthorsModelUpdateV1) HasActualname() bool`

HasActualname returns a boolean if a field has been set.

### GetBirthday

`func (o *AuthorsModelUpdateV1) GetBirthday() BirthdayModelV1`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *AuthorsModelUpdateV1) GetBirthdayOk() (*BirthdayModelV1, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *AuthorsModelUpdateV1) SetBirthday(v BirthdayModelV1)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *AuthorsModelUpdateV1) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetBirthplace

`func (o *AuthorsModelUpdateV1) GetBirthplace() string`

GetBirthplace returns the Birthplace field if non-nil, zero value otherwise.

### GetBirthplaceOk

`func (o *AuthorsModelUpdateV1) GetBirthplaceOk() (*string, bool)`

GetBirthplaceOk returns a tuple with the Birthplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthplace

`func (o *AuthorsModelUpdateV1) SetBirthplace(v string)`

SetBirthplace sets Birthplace field to given value.

### HasBirthplace

`func (o *AuthorsModelUpdateV1) HasBirthplace() bool`

HasBirthplace returns a boolean if a field has been set.

### GetBloodtype

`func (o *AuthorsModelUpdateV1) GetBloodtype() string`

GetBloodtype returns the Bloodtype field if non-nil, zero value otherwise.

### GetBloodtypeOk

`func (o *AuthorsModelUpdateV1) GetBloodtypeOk() (*string, bool)`

GetBloodtypeOk returns a tuple with the Bloodtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloodtype

`func (o *AuthorsModelUpdateV1) SetBloodtype(v string)`

SetBloodtype sets Bloodtype field to given value.

### HasBloodtype

`func (o *AuthorsModelUpdateV1) HasBloodtype() bool`

HasBloodtype returns a boolean if a field has been set.

### GetGender

`func (o *AuthorsModelUpdateV1) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AuthorsModelUpdateV1) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AuthorsModelUpdateV1) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AuthorsModelUpdateV1) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetSocial

`func (o *AuthorsModelUpdateV1) GetSocial() AuthorsModelUpdateV1Social`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *AuthorsModelUpdateV1) GetSocialOk() (*AuthorsModelUpdateV1Social, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *AuthorsModelUpdateV1) SetSocial(v AuthorsModelUpdateV1Social)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *AuthorsModelUpdateV1) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetComments

`func (o *AuthorsModelUpdateV1) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AuthorsModelUpdateV1) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AuthorsModelUpdateV1) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AuthorsModelUpdateV1) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAdmin

`func (o *AuthorsModelUpdateV1) GetAdmin() AuthorsModelUpdateV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AuthorsModelUpdateV1) GetAdminOk() (*AuthorsModelUpdateV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AuthorsModelUpdateV1) SetAdmin(v AuthorsModelUpdateV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AuthorsModelUpdateV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


