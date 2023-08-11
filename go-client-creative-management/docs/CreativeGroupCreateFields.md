# CreativeGroupCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this creative group is going to redirect to. | 
**Name** | **string** | The name of the creative group. | 
**Spec** | [**CreativeSpec**](CreativeSpec.md) |  | 
**Creatives** | [**[]WeightedCreative**](WeightedCreative.md) | A list of all creatives that compose this creative group with their relative weights. | 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 

## Methods

### NewCreativeGroupCreateFields

`func NewCreativeGroupCreateFields(domain string, name string, spec CreativeSpec, creatives []WeightedCreative, ) *CreativeGroupCreateFields`

NewCreativeGroupCreateFields instantiates a new CreativeGroupCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeGroupCreateFieldsWithDefaults

`func NewCreativeGroupCreateFieldsWithDefaults() *CreativeGroupCreateFields`

NewCreativeGroupCreateFieldsWithDefaults instantiates a new CreativeGroupCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreativeGroupCreateFields) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreativeGroupCreateFields) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreativeGroupCreateFields) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *CreativeGroupCreateFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeGroupCreateFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeGroupCreateFields) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *CreativeGroupCreateFields) GetSpec() CreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreativeGroupCreateFields) GetSpecOk() (*CreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreativeGroupCreateFields) SetSpec(v CreativeSpec)`

SetSpec sets Spec field to given value.


### GetCreatives

`func (o *CreativeGroupCreateFields) GetCreatives() []WeightedCreative`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *CreativeGroupCreateFields) GetCreativesOk() (*[]WeightedCreative, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *CreativeGroupCreateFields) SetCreatives(v []WeightedCreative)`

SetCreatives sets Creatives field to given value.


### GetTags

`func (o *CreativeGroupCreateFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreativeGroupCreateFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreativeGroupCreateFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreativeGroupCreateFields) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


