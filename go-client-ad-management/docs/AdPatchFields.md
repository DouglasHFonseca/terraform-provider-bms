# AdPatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the ad. | [optional] 
**Rules** | Pointer to [**[]AdRule**](AdRule.md) | The rules that specify what are the conditions to display a creative group. | [optional] 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 
**Archived** | Pointer to **bool** | Indicates if it is archived. | [optional] 

## Methods

### NewAdPatchFields

`func NewAdPatchFields() *AdPatchFields`

NewAdPatchFields instantiates a new AdPatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdPatchFieldsWithDefaults

`func NewAdPatchFieldsWithDefaults() *AdPatchFields`

NewAdPatchFieldsWithDefaults instantiates a new AdPatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdPatchFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdPatchFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdPatchFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdPatchFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *AdPatchFields) GetRules() []AdRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AdPatchFields) GetRulesOk() (*[]AdRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AdPatchFields) SetRules(v []AdRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AdPatchFields) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *AdPatchFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdPatchFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdPatchFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdPatchFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetArchived

`func (o *AdPatchFields) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AdPatchFields) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AdPatchFields) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AdPatchFields) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


