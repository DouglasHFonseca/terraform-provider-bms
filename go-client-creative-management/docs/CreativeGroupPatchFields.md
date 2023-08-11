# CreativeGroupPatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The description of this creative group. | [optional] 
**Creatives** | Pointer to [**[]WeightedCreative**](WeightedCreative.md) | A list of all creatives that compose this creative group with their relative weights. | [optional] 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 
**Archived** | Pointer to **bool** | Indicates if it is archived. | [optional] 

## Methods

### NewCreativeGroupPatchFields

`func NewCreativeGroupPatchFields() *CreativeGroupPatchFields`

NewCreativeGroupPatchFields instantiates a new CreativeGroupPatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeGroupPatchFieldsWithDefaults

`func NewCreativeGroupPatchFieldsWithDefaults() *CreativeGroupPatchFields`

NewCreativeGroupPatchFieldsWithDefaults instantiates a new CreativeGroupPatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreativeGroupPatchFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeGroupPatchFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeGroupPatchFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreativeGroupPatchFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatives

`func (o *CreativeGroupPatchFields) GetCreatives() []WeightedCreative`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *CreativeGroupPatchFields) GetCreativesOk() (*[]WeightedCreative, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *CreativeGroupPatchFields) SetCreatives(v []WeightedCreative)`

SetCreatives sets Creatives field to given value.

### HasCreatives

`func (o *CreativeGroupPatchFields) HasCreatives() bool`

HasCreatives returns a boolean if a field has been set.

### GetTags

`func (o *CreativeGroupPatchFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreativeGroupPatchFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreativeGroupPatchFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreativeGroupPatchFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetArchived

`func (o *CreativeGroupPatchFields) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreativeGroupPatchFields) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreativeGroupPatchFields) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreativeGroupPatchFields) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


