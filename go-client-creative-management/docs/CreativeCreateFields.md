# CreativeCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this creative is going to redirect to. | 
**Name** | **string** | The name of the creative. | 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 
**Banner** | Pointer to [**CreativeCreateFieldsAllOfBanner**](CreativeCreateFieldsAllOfBanner.md) |  | [optional] 

## Methods

### NewCreativeCreateFields

`func NewCreativeCreateFields(domain string, name string, ) *CreativeCreateFields`

NewCreativeCreateFields instantiates a new CreativeCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeCreateFieldsWithDefaults

`func NewCreativeCreateFieldsWithDefaults() *CreativeCreateFields`

NewCreativeCreateFieldsWithDefaults instantiates a new CreativeCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreativeCreateFields) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreativeCreateFields) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreativeCreateFields) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *CreativeCreateFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeCreateFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeCreateFields) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *CreativeCreateFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreativeCreateFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreativeCreateFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreativeCreateFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBanner

`func (o *CreativeCreateFields) GetBanner() CreativeCreateFieldsAllOfBanner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *CreativeCreateFields) GetBannerOk() (*CreativeCreateFieldsAllOfBanner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *CreativeCreateFields) SetBanner(v CreativeCreateFieldsAllOfBanner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *CreativeCreateFields) HasBanner() bool`

HasBanner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


