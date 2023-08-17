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

// checks if the AdsCreativeAvailableCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdsCreativeAvailableCreative{}

// AdsCreativeAvailableCreative Represents a creative that is available for printing with its relative weight.
type AdsCreativeAvailableCreative struct {
	// The ID of the creative.
	CreativeId string `json:"creativeId"`
	// The weight of the creative.
	Weight float32 `json:"weight"`
	Banner *AdsCreativeBannerFields `json:"banner,omitempty"`
}

// NewAdsCreativeAvailableCreative instantiates a new AdsCreativeAvailableCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdsCreativeAvailableCreative(creativeId string, weight float32) *AdsCreativeAvailableCreative {
	this := AdsCreativeAvailableCreative{}
	this.CreativeId = creativeId
	this.Weight = weight
	return &this
}

// NewAdsCreativeAvailableCreativeWithDefaults instantiates a new AdsCreativeAvailableCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdsCreativeAvailableCreativeWithDefaults() *AdsCreativeAvailableCreative {
	this := AdsCreativeAvailableCreative{}
	return &this
}

// GetCreativeId returns the CreativeId field value
func (o *AdsCreativeAvailableCreative) GetCreativeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeAvailableCreative) GetCreativeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeId, true
}

// SetCreativeId sets field value
func (o *AdsCreativeAvailableCreative) SetCreativeId(v string) {
	o.CreativeId = v
}

// GetWeight returns the Weight field value
func (o *AdsCreativeAvailableCreative) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeAvailableCreative) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *AdsCreativeAvailableCreative) SetWeight(v float32) {
	o.Weight = v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *AdsCreativeAvailableCreative) GetBanner() AdsCreativeBannerFields {
	if o == nil || IsNil(o.Banner) {
		var ret AdsCreativeBannerFields
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdsCreativeAvailableCreative) GetBannerOk() (*AdsCreativeBannerFields, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *AdsCreativeAvailableCreative) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given AdsCreativeBannerFields and assigns it to the Banner field.
func (o *AdsCreativeAvailableCreative) SetBanner(v AdsCreativeBannerFields) {
	o.Banner = &v
}

func (o AdsCreativeAvailableCreative) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdsCreativeAvailableCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeId"] = o.CreativeId
	toSerialize["weight"] = o.Weight
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	return toSerialize, nil
}

type NullableAdsCreativeAvailableCreative struct {
	value *AdsCreativeAvailableCreative
	isSet bool
}

func (v NullableAdsCreativeAvailableCreative) Get() *AdsCreativeAvailableCreative {
	return v.value
}

func (v *NullableAdsCreativeAvailableCreative) Set(val *AdsCreativeAvailableCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableAdsCreativeAvailableCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableAdsCreativeAvailableCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdsCreativeAvailableCreative(val *AdsCreativeAvailableCreative) *NullableAdsCreativeAvailableCreative {
	return &NullableAdsCreativeAvailableCreative{value: val, isSet: true}
}

func (v NullableAdsCreativeAvailableCreative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdsCreativeAvailableCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

