# CreativeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this creative group is going to redirect to. | 
**Name** | **string** | The name of the creative group. | 
**Spec** | [**CreativeSpec**](CreativeSpec.md) |  | 
**Creatives** | [**[]WeightedCreative**](WeightedCreative.md) | A list of all creatives that compose this creative group with their relative weights. | 
**Tags** | **[]string** | Tags that the user can add to facilitate searching. | 
**CreativeGroupId** | **string** | The ID of the creative group. | 
**AccountId** | **string** | The ID of the account that owns this creative group. | 
**Archived** | **bool** | Indicates if it is archived. | 

## Methods

### NewCreativeGroup

`func NewCreativeGroup(domain string, name string, spec CreativeSpec, creatives []WeightedCreative, tags []string, creativeGroupId string, accountId string, archived bool, ) *CreativeGroup`

NewCreativeGroup instantiates a new CreativeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeGroupWithDefaults

`func NewCreativeGroupWithDefaults() *CreativeGroup`

NewCreativeGroupWithDefaults instantiates a new CreativeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CreativeGroup) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreativeGroup) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreativeGroup) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *CreativeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreativeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreativeGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *CreativeGroup) GetSpec() CreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreativeGroup) GetSpecOk() (*CreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreativeGroup) SetSpec(v CreativeSpec)`

SetSpec sets Spec field to given value.


### GetCreatives

`func (o *CreativeGroup) GetCreatives() []WeightedCreative`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *CreativeGroup) GetCreativesOk() (*[]WeightedCreative, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *CreativeGroup) SetCreatives(v []WeightedCreative)`

SetCreatives sets Creatives field to given value.


### GetTags

`func (o *CreativeGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreativeGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreativeGroup) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreativeGroupId

`func (o *CreativeGroup) GetCreativeGroupId() string`

GetCreativeGroupId returns the CreativeGroupId field if non-nil, zero value otherwise.

### GetCreativeGroupIdOk

`func (o *CreativeGroup) GetCreativeGroupIdOk() (*string, bool)`

GetCreativeGroupIdOk returns a tuple with the CreativeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeGroupId

`func (o *CreativeGroup) SetCreativeGroupId(v string)`

SetCreativeGroupId sets CreativeGroupId field to given value.


### GetAccountId

`func (o *CreativeGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreativeGroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreativeGroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArchived

`func (o *CreativeGroup) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreativeGroup) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreativeGroup) SetArchived(v bool)`

SetArchived sets Archived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


