# CreativePatchFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Banner** | Pointer to [**BannerFields**](BannerFields.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the creative. | [optional] 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 
**Archived** | Pointer to **bool** | Indicates if it is archived. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the creative is enabled for usage. | [optional] 

## Methods

### NewCreativePatchFields

`func NewCreativePatchFields() *CreativePatchFields`

NewCreativePatchFields instantiates a new CreativePatchFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativePatchFieldsWithDefaults

`func NewCreativePatchFieldsWithDefaults() *CreativePatchFields`

NewCreativePatchFieldsWithDefaults instantiates a new CreativePatchFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanner

`func (o *CreativePatchFields) GetBanner() BannerFields`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *CreativePatchFields) GetBannerOk() (*BannerFields, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *CreativePatchFields) SetBanner(v BannerFields)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *CreativePatchFields) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetName

`func (o *CreativePatchFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativePatchFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativePatchFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreativePatchFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *CreativePatchFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreativePatchFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreativePatchFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreativePatchFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetArchived

`func (o *CreativePatchFields) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreativePatchFields) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreativePatchFields) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreativePatchFields) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetEnabled

`func (o *CreativePatchFields) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreativePatchFields) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreativePatchFields) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreativePatchFields) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


