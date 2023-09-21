# GroupsSearchResponseV1Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | Pointer to [**GroupsModelSearchV1**](GroupsModelSearchV1.md) |  | [optional] 
**HitName** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupsSearchResponseV1Results

`func NewGroupsSearchResponseV1Results() *GroupsSearchResponseV1Results`

NewGroupsSearchResponseV1Results instantiates a new GroupsSearchResponseV1Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsSearchResponseV1ResultsWithDefaults

`func NewGroupsSearchResponseV1ResultsWithDefaults() *GroupsSearchResponseV1Results`

NewGroupsSearchResponseV1ResultsWithDefaults instantiates a new GroupsSearchResponseV1Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *GroupsSearchResponseV1Results) GetRecord() GroupsModelSearchV1`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *GroupsSearchResponseV1Results) GetRecordOk() (*GroupsModelSearchV1, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *GroupsSearchResponseV1Results) SetRecord(v GroupsModelSearchV1)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *GroupsSearchResponseV1Results) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetHitName

`func (o *GroupsSearchResponseV1Results) GetHitName() string`

GetHitName returns the HitName field if non-nil, zero value otherwise.

### GetHitNameOk

`func (o *GroupsSearchResponseV1Results) GetHitNameOk() (*string, bool)`

GetHitNameOk returns a tuple with the HitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitName

`func (o *GroupsSearchResponseV1Results) SetHitName(v string)`

SetHitName sets HitName field to given value.

### HasHitName

`func (o *GroupsSearchResponseV1Results) HasHitName() bool`

HasHitName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


