# AdCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that this ad is going to redirect to. | 
**Name** | **string** | The name of this ad. | 
**Spec** | [**AdsCreativeCreativeSpec**](AdsCreativeCreativeSpec.md) |  | 
**Rules** | [**[]AdRule**](AdRule.md) | The rules that specify what are the conditions to display a creative group. | 
**Tags** | Pointer to **[]string** | Tags that the user can add to facilitate searching. | [optional] 

## Methods

### NewAdCreateFields

`func NewAdCreateFields(domain string, name string, spec AdsCreativeCreativeSpec, rules []AdRule, ) *AdCreateFields`

NewAdCreateFields instantiates a new AdCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCreateFieldsWithDefaults

`func NewAdCreateFieldsWithDefaults() *AdCreateFields`

NewAdCreateFieldsWithDefaults instantiates a new AdCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *AdCreateFields) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AdCreateFields) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AdCreateFields) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *AdCreateFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdCreateFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdCreateFields) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *AdCreateFields) GetSpec() AdsCreativeCreativeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AdCreateFields) GetSpecOk() (*AdsCreativeCreativeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AdCreateFields) SetSpec(v AdsCreativeCreativeSpec)`

SetSpec sets Spec field to given value.


### GetRules

`func (o *AdCreateFields) GetRules() []AdRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AdCreateFields) GetRulesOk() (*[]AdRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AdCreateFields) SetRules(v []AdRule)`

SetRules sets Rules field to given value.


### GetTags

`func (o *AdCreateFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdCreateFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdCreateFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdCreateFields) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


