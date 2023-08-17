# AdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdRuleId** | **string** | The rule ID. Leave an empty string if creating a new rule. | 
**Name** | **string** | A name of this rule. | 
**Conditions** | [**[]Condition**](Condition.md) | The conditions that must be &#x60;true&#x60; for this rule to be applicable. If no conditions are specified, the rule will match. | 
**CreativeGroupId** | **string** | The ID of the creative group. | 
**Enabled** | **bool** | Indicates if this rule is enabled and should be evaluated. | 

## Methods

### NewAdRule

`func NewAdRule(adRuleId string, name string, conditions []Condition, creativeGroupId string, enabled bool, ) *AdRule`

NewAdRule instantiates a new AdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdRuleWithDefaults

`func NewAdRuleWithDefaults() *AdRule`

NewAdRuleWithDefaults instantiates a new AdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdRuleId

`func (o *AdRule) GetAdRuleId() string`

GetAdRuleId returns the AdRuleId field if non-nil, zero value otherwise.

### GetAdRuleIdOk

`func (o *AdRule) GetAdRuleIdOk() (*string, bool)`

GetAdRuleIdOk returns a tuple with the AdRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRuleId

`func (o *AdRule) SetAdRuleId(v string)`

SetAdRuleId sets AdRuleId field to given value.


### GetName

`func (o *AdRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdRule) SetName(v string)`

SetName sets Name field to given value.


### GetConditions

`func (o *AdRule) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AdRule) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AdRule) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetCreativeGroupId

`func (o *AdRule) GetCreativeGroupId() string`

GetCreativeGroupId returns the CreativeGroupId field if non-nil, zero value otherwise.

### GetCreativeGroupIdOk

`func (o *AdRule) GetCreativeGroupIdOk() (*string, bool)`

GetCreativeGroupIdOk returns a tuple with the CreativeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeGroupId

`func (o *AdRule) SetCreativeGroupId(v string)`

SetCreativeGroupId sets CreativeGroupId field to given value.


### GetEnabled

`func (o *AdRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


