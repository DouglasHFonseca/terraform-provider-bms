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

// checks if the AdsCreativeWeightedCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdsCreativeWeightedCreative{}

// AdsCreativeWeightedCreative Represents a creative and its weight.
type AdsCreativeWeightedCreative struct {
	// The ID of the creative.
	CreativeId string `json:"creativeId"`
	// The weight of the creative.
	Weight float32 `json:"weight"`
}

// NewAdsCreativeWeightedCreative instantiates a new AdsCreativeWeightedCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdsCreativeWeightedCreative(creativeId string, weight float32) *AdsCreativeWeightedCreative {
	this := AdsCreativeWeightedCreative{}
	this.CreativeId = creativeId
	this.Weight = weight
	return &this
}

// NewAdsCreativeWeightedCreativeWithDefaults instantiates a new AdsCreativeWeightedCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdsCreativeWeightedCreativeWithDefaults() *AdsCreativeWeightedCreative {
	this := AdsCreativeWeightedCreative{}
	return &this
}

// GetCreativeId returns the CreativeId field value
func (o *AdsCreativeWeightedCreative) GetCreativeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeWeightedCreative) GetCreativeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeId, true
}

// SetCreativeId sets field value
func (o *AdsCreativeWeightedCreative) SetCreativeId(v string) {
	o.CreativeId = v
}

// GetWeight returns the Weight field value
func (o *AdsCreativeWeightedCreative) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeWeightedCreative) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *AdsCreativeWeightedCreative) SetWeight(v float32) {
	o.Weight = v
}

func (o AdsCreativeWeightedCreative) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdsCreativeWeightedCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeId"] = o.CreativeId
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableAdsCreativeWeightedCreative struct {
	value *AdsCreativeWeightedCreative
	isSet bool
}

func (v NullableAdsCreativeWeightedCreative) Get() *AdsCreativeWeightedCreative {
	return v.value
}

func (v *NullableAdsCreativeWeightedCreative) Set(val *AdsCreativeWeightedCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableAdsCreativeWeightedCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableAdsCreativeWeightedCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdsCreativeWeightedCreative(val *AdsCreativeWeightedCreative) *NullableAdsCreativeWeightedCreative {
	return &NullableAdsCreativeWeightedCreative{value: val, isSet: true}
}

func (v NullableAdsCreativeWeightedCreative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdsCreativeWeightedCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


