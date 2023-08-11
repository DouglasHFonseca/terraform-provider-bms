# CreativeGroupFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | Generic search that can match against id, name, tag and domain. | [optional] 
**CreativeIds** | Pointer to **[]string** | The list of creative IDs in usage to look for. | [optional] 
**Domain** | Pointer to **string** | The domain to filter. | [optional] 
**Name** | Pointer to **string** | The description to search for. | [optional] 
**Archived** | Pointer to **bool** | &#x60;true&#x60; to include only archived creative groups, &#x60;false&#x60; to include only active, leave empty to include all. | [optional] 
**Tag** | Pointer to **string** | The tags to look for. | [optional] 
**Spec** | Pointer to [**CreativeSpec**](CreativeSpec.md) |  | [optional] 

## Methods

### NewCreativeGroupFilters

`func NewCreativeGroupFilters() *CreativeGroupFilters`

NewCreativeGroupFilters instantiates a new CreativeGroupFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeGroupFiltersWithDefaults

`func NewCreativeGroupFiltersWithDefaults() *CreativeGroupFilters`

NewCreativeGroupFiltersWithDefaults instantiates a new CreativeGroupFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *CreativeGroupFilters) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreativeGroupFilters) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreativeGroupFilters) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreativeGroupFilters) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetCreativeIds

`func (o *CreativeGroupFilters) GetCreativeIds() []string`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *CreativeGroupFilters) GetCreativeIdsOk() (*[]string, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *CreativeGroupFilters) SetCreativeIds(v []string)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *CreativeGroupFilters) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetDomain

`func (o *CreativeGroupFilters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreativeGroupFilters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreativeGroupFilters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreativeGroupFilters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *CreativeGroupFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeGroupFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeGroupFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreativeGroupFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *CreativeGroupFilters) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreativeGroupFilters) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreativeGroupFilters) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreativeGroupFilters) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTag

`func (o *CreativeGroupFilters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreativeGroupFilters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreativeGroupFilters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreativeGroupFilters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSpec

`func (o *CreativeGroupFilters) GetSpec() CreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreativeGroupFilters) GetSpecOk() (*CreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreativeGroupFilters) SetSpec(v CreativeSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CreativeGroupFilters) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


