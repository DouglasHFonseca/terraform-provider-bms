# ActiveAdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdRuleId** | **string** | The rule ID. Leave an empty string if creating a new rule. | 
**Name** | **string** | A name of this rule. | 
**Conditions** | [**[]Condition**](Condition.md) | The conditions that must be &#x60;true&#x60; for this rule to be applicable. If no conditions are specified, the rule will match. | 
**CreativeGroupId** | **string** | The ID of the creative group. | 
**Enabled** | **bool** | Indicates if this rule is enabled and should be evaluated. | 
**AvailableCreatives** | Pointer to [**[]AdsCreativeAvailableCreative**](AdsCreativeAvailableCreative.md) | Contains the creatives that are active for this rule. | [optional] 

## Methods

### NewActiveAdRule

`func NewActiveAdRule(adRuleId string, name string, conditions []Condition, creativeGroupId string, enabled bool, ) *ActiveAdRule`

NewActiveAdRule instantiates a new ActiveAdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveAdRuleWithDefaults

`func NewActiveAdRuleWithDefaults() *ActiveAdRule`

NewActiveAdRuleWithDefaults instantiates a new ActiveAdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdRuleId

`func (o *ActiveAdRule) GetAdRuleId() string`

GetAdRuleId returns the AdRuleId field if non-nil, zero value otherwise.

### GetAdRuleIdOk

`func (o *ActiveAdRule) GetAdRuleIdOk() (*string, bool)`

GetAdRuleIdOk returns a tuple with the AdRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRuleId

`func (o *ActiveAdRule) SetAdRuleId(v string)`

SetAdRuleId sets AdRuleId field to given value.


### GetName

`func (o *ActiveAdRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveAdRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveAdRule) SetName(v string)`

SetName sets Name field to given value.


### GetConditions

`func (o *ActiveAdRule) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ActiveAdRule) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ActiveAdRule) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetCreativeGroupId

`func (o *ActiveAdRule) GetCreativeGroupId() string`

GetCreativeGroupId returns the CreativeGroupId field if non-nil, zero value otherwise.

### GetCreativeGroupIdOk

`func (o *ActiveAdRule) GetCreativeGroupIdOk() (*string, bool)`

GetCreativeGroupIdOk returns a tuple with the CreativeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeGroupId

`func (o *ActiveAdRule) SetCreativeGroupId(v string)`

SetCreativeGroupId sets CreativeGroupId field to given value.


### GetEnabled

`func (o *ActiveAdRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActiveAdRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActiveAdRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAvailableCreatives

`func (o *ActiveAdRule) GetAvailableCreatives() []AdsCreativeAvailableCreative`

GetAvailableCreatives returns the AvailableCreatives field if non-nil, zero value otherwise.

### GetAvailableCreativesOk

`func (o *ActiveAdRule) GetAvailableCreativesOk() (*[]AdsCreativeAvailableCreative, bool)`

GetAvailableCreativesOk returns a tuple with the AvailableCreatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCreatives

`func (o *ActiveAdRule) SetAvailableCreatives(v []AdsCreativeAvailableCreative)`

SetAvailableCreatives sets AvailableCreatives field to given value.

### HasAvailableCreatives

`func (o *ActiveAdRule) HasAvailableCreatives() bool`

HasAvailableCreatives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


