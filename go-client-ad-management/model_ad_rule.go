/*
AdServer Ad Management

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AdRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdRule{}

// AdRule Represents a rule that dictates when a creative group should be rendered.
type AdRule struct {
	// The rule ID. Leave an empty string if creating a new rule.
	AdRuleId string `json:"adRuleId"`
	// A name of this rule.
	Name string `json:"name"`
	// The conditions that must be `true` for this rule to be applicable. If no conditions are specified, the rule will match.
	Conditions []Condition `json:"conditions"`
	// The ID of the creative group.
	CreativeGroupId string `json:"creativeGroupId"`
	// Indicates if this rule is enabled and should be evaluated.
	Enabled bool `json:"enabled"`
}

// NewAdRule instantiates a new AdRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdRule(adRuleId string, name string, conditions []Condition, creativeGroupId string, enabled bool) *AdRule {
	this := AdRule{}
	this.AdRuleId = adRuleId
	this.Name = name
	this.Conditions = conditions
	this.CreativeGroupId = creativeGroupId
	this.Enabled = enabled
	return &this
}

// NewAdRuleWithDefaults instantiates a new AdRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdRuleWithDefaults() *AdRule {
	this := AdRule{}
	return &this
}

// GetAdRuleId returns the AdRuleId field value
func (o *AdRule) GetAdRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdRuleId
}

// GetAdRuleIdOk returns a tuple with the AdRuleId field value
// and a boolean to check if the value has been set.
func (o *AdRule) GetAdRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdRuleId, true
}

// SetAdRuleId sets field value
func (o *AdRule) SetAdRuleId(v string) {
	o.AdRuleId = v
}

// GetName returns the Name field value
func (o *AdRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdRule) SetName(v string) {
	o.Name = v
}

// GetConditions returns the Conditions field value
func (o *AdRule) GetConditions() []Condition {
	if o == nil {
		var ret []Condition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *AdRule) GetConditionsOk() ([]Condition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *AdRule) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetCreativeGroupId returns the CreativeGroupId field value
func (o *AdRule) GetCreativeGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreativeGroupId
}

// GetCreativeGroupIdOk returns a tuple with the CreativeGroupId field value
// and a boolean to check if the value has been set.
func (o *AdRule) GetCreativeGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeGroupId, true
}

// SetCreativeGroupId sets field value
func (o *AdRule) SetCreativeGroupId(v string) {
	o.CreativeGroupId = v
}

// GetEnabled returns the Enabled field value
func (o *AdRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AdRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AdRule) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AdRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adRuleId"] = o.AdRuleId
	toSerialize["name"] = o.Name
	toSerialize["conditions"] = o.Conditions
	toSerialize["creativeGroupId"] = o.CreativeGroupId
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAdRule struct {
	value *AdRule
	isSet bool
}

func (v NullableAdRule) Get() *AdRule {
	return v.value
}

func (v *NullableAdRule) Set(val *AdRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAdRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAdRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdRule(val *AdRule) *NullableAdRule {
	return &NullableAdRule{value: val, isSet: true}
}

func (v NullableAdRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


