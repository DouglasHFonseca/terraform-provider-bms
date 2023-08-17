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

// checks if the AdPatchFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdPatchFields{}

// AdPatchFields `Ad` editable fields.
type AdPatchFields struct {
	// The name of the ad.
	Name *string `json:"name,omitempty"`
	// The rules that specify what are the conditions to display a creative group.
	Rules []AdRule `json:"rules,omitempty"`
	// Tags that the user can add to facilitate searching.
	Tags []string `json:"tags,omitempty"`
	// Indicates if it is archived.
	Archived *bool `json:"archived,omitempty"`
}

// NewAdPatchFields instantiates a new AdPatchFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdPatchFields() *AdPatchFields {
	this := AdPatchFields{}
	return &this
}

// NewAdPatchFieldsWithDefaults instantiates a new AdPatchFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdPatchFieldsWithDefaults() *AdPatchFields {
	this := AdPatchFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdPatchFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdPatchFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdPatchFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdPatchFields) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AdPatchFields) GetRules() []AdRule {
	if o == nil || IsNil(o.Rules) {
		var ret []AdRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdPatchFields) GetRulesOk() ([]AdRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AdPatchFields) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []AdRule and assigns it to the Rules field.
func (o *AdPatchFields) SetRules(v []AdRule) {
	o.Rules = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdPatchFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdPatchFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdPatchFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdPatchFields) SetTags(v []string) {
	o.Tags = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *AdPatchFields) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdPatchFields) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *AdPatchFields) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *AdPatchFields) SetArchived(v bool) {
	o.Archived = &v
}

func (o AdPatchFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdPatchFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableAdPatchFields struct {
	value *AdPatchFields
	isSet bool
}

func (v NullableAdPatchFields) Get() *AdPatchFields {
	return v.value
}

func (v *NullableAdPatchFields) Set(val *AdPatchFields) {
	v.value = val
	v.isSet = true
}

func (v NullableAdPatchFields) IsSet() bool {
	return v.isSet
}

func (v *NullableAdPatchFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdPatchFields(val *AdPatchFields) *NullableAdPatchFields {
	return &NullableAdPatchFields{value: val, isSet: true}
}

func (v NullableAdPatchFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdPatchFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

