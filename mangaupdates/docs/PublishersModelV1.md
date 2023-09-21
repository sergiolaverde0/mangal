# PublishersModelV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublisherId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Associated** | Pointer to [**[]PublishersModelV1Associated**](PublishersModelV1Associated.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**PublishersModelV1Stats**](PublishersModelV1Stats.md) |  | [optional] 
**AddedBy** | Pointer to [**UserModelSearchV1**](UserModelSearchV1.md) |  | [optional] 
**LastUpdated** | Pointer to [**TimeV1**](TimeV1.md) |  | [optional] 
**Admin** | Pointer to [**PublishersModelV1Admin**](PublishersModelV1Admin.md) |  | [optional] 

## Methods

### NewPublishersModelV1

`func NewPublishersModelV1() *PublishersModelV1`

NewPublishersModelV1 instantiates a new PublishersModelV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishersModelV1WithDefaults

`func NewPublishersModelV1WithDefaults() *PublishersModelV1`

NewPublishersModelV1WithDefaults instantiates a new PublishersModelV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublisherId

`func (o *PublishersModelV1) GetPublisherId() int64`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *PublishersModelV1) GetPublisherIdOk() (*int64, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *PublishersModelV1) SetPublisherId(v int64)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *PublishersModelV1) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetName

`func (o *PublishersModelV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishersModelV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishersModelV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublishersModelV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *PublishersModelV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublishersModelV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublishersModelV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PublishersModelV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAssociated

`func (o *PublishersModelV1) GetAssociated() []PublishersModelV1Associated`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *PublishersModelV1) GetAssociatedOk() (*[]PublishersModelV1Associated, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *PublishersModelV1) SetAssociated(v []PublishersModelV1Associated)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *PublishersModelV1) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetType

`func (o *PublishersModelV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublishersModelV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublishersModelV1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublishersModelV1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInfo

`func (o *PublishersModelV1) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PublishersModelV1) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PublishersModelV1) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PublishersModelV1) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSite

`func (o *PublishersModelV1) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PublishersModelV1) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PublishersModelV1) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PublishersModelV1) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetStats

`func (o *PublishersModelV1) GetStats() PublishersModelV1Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *PublishersModelV1) GetStatsOk() (*PublishersModelV1Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *PublishersModelV1) SetStats(v PublishersModelV1Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *PublishersModelV1) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetAddedBy

`func (o *PublishersModelV1) GetAddedBy() UserModelSearchV1`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *PublishersModelV1) GetAddedByOk() (*UserModelSearchV1, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *PublishersModelV1) SetAddedBy(v UserModelSearchV1)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *PublishersModelV1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PublishersModelV1) GetLastUpdated() TimeV1`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PublishersModelV1) GetLastUpdatedOk() (*TimeV1, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PublishersModelV1) SetLastUpdated(v TimeV1)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PublishersModelV1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAdmin

`func (o *PublishersModelV1) GetAdmin() PublishersModelV1Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *PublishersModelV1) GetAdminOk() (*PublishersModelV1Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *PublishersModelV1) SetAdmin(v PublishersModelV1Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *PublishersModelV1) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


