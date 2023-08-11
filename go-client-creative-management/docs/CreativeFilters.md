# CreativeFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | Generic search that can match against id, name, tag and domain. | [optional] 
**CreativeIds** | Pointer to **[]string** | An array of IDs to search for. | [optional] 
**CreativeGroupIds** | Pointer to **[]string** | The list of creative group IDs to search usage for. | [optional] 
**Status** | Pointer to **string** | The status to search for. | [optional] 
**Tag** | Pointer to **string** | The tag to look for. | [optional] 
**Domain** | Pointer to **string** | The domain to filter. | [optional] 
**Archived** | Pointer to **bool** | &#x60;true&#x60; to include only archived creatives, &#x60;false&#x60; to include only active, leave empty to include all. | [optional] 
**Name** | Pointer to **string** | The name to search for. | [optional] 
**Spec** | Pointer to [**CreativeSpec**](CreativeSpec.md) |  | [optional] 
**Enabled** | Pointer to **bool** | &#x60;true&#x60; to include only enabled creatives, &#x60;false&#x60; to include only disabled, leave empty to include all. | [optional] 

## Methods

### NewCreativeFilters

`func NewCreativeFilters() *CreativeFilters`

NewCreativeFilters instantiates a new CreativeFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeFiltersWithDefaults

`func NewCreativeFiltersWithDefaults() *CreativeFilters`

NewCreativeFiltersWithDefaults instantiates a new CreativeFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *CreativeFilters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreativeFilters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreativeFilters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreativeFilters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetCreativeIds

`func (o *CreativeFilters) GetCreativeIds() []string`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *CreativeFilters) GetCreativeIdsOk() (*[]string, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *CreativeFilters) SetCreativeIds(v []string)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *CreativeFilters) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetCreativeGroupIds

`func (o *CreativeFilters) GetCreativeGroupIds() []string`

GetCreativeGroupIds returns the CreativeGroupIds field if non-nil, zero value otherwise.

### GetCreativeGroupIdsOk

`func (o *CreativeFilters) GetCreativeGroupIdsOk() (*[]string, bool)`

GetCreativeGroupIdsOk returns a tuple with the CreativeGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeGroupIds

`func (o *CreativeFilters) SetCreativeGroupIds(v []string)`

SetCreativeGroupIds sets CreativeGroupIds field to given value.

### HasCreativeGroupIds

`func (o *CreativeFilters) HasCreativeGroupIds() bool`

HasCreativeGroupIds returns a boolean if a field has been set.

### GetStatus

`func (o *CreativeFilters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreativeFilters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreativeFilters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreativeFilters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTag

`func (o *CreativeFilters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreativeFilters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreativeFilters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreativeFilters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetDomain

`func (o *CreativeFilters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreativeFilters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreativeFilters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreativeFilters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetArchived

`func (o *CreativeFilters) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreativeFilters) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreativeFilters) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreativeFilters) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetName

`func (o *CreativeFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreativeFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *CreativeFilters) GetSpec() CreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreativeFilters) GetSpecOk() (*CreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreativeFilters) SetSpec(v CreativeSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CreativeFilters) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetEnabled

`func (o *CreativeFilters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreativeFilters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreativeFilters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreativeFilters) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


