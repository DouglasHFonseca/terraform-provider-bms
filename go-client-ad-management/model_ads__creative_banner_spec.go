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

// checks if the AdsCreativeBannerSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdsCreativeBannerSpec{}

// AdsCreativeBannerSpec Specifies the requirements for banners.
type AdsCreativeBannerSpec struct {
	// The width of the banner in pixels.
	Width float32 `json:"width"`
	// The height of the banner in pixels.
	Height float32 `json:"height"`
}

// NewAdsCreativeBannerSpec instantiates a new AdsCreativeBannerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdsCreativeBannerSpec(width float32, height float32) *AdsCreativeBannerSpec {
	this := AdsCreativeBannerSpec{}
	this.Width = width
	this.Height = height
	return &this
}

// NewAdsCreativeBannerSpecWithDefaults instantiates a new AdsCreativeBannerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdsCreativeBannerSpecWithDefaults() *AdsCreativeBannerSpec {
	this := AdsCreativeBannerSpec{}
	return &this
}

// GetWidth returns the Width field value
func (o *AdsCreativeBannerSpec) GetWidth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeBannerSpec) GetWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *AdsCreativeBannerSpec) SetWidth(v float32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *AdsCreativeBannerSpec) GetHeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *AdsCreativeBannerSpec) GetHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *AdsCreativeBannerSpec) SetHeight(v float32) {
	o.Height = v
}

func (o AdsCreativeBannerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdsCreativeBannerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

type NullableAdsCreativeBannerSpec struct {
	value *AdsCreativeBannerSpec
	isSet bool
}

func (v NullableAdsCreativeBannerSpec) Get() *AdsCreativeBannerSpec {
	return v.value
}

func (v *NullableAdsCreativeBannerSpec) Set(val *AdsCreativeBannerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAdsCreativeBannerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAdsCreativeBannerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdsCreativeBannerSpec(val *AdsCreativeBannerSpec) *NullableAdsCreativeBannerSpec {
	return &NullableAdsCreativeBannerSpec{value: val, isSet: true}
}

func (v NullableAdsCreativeBannerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdsCreativeBannerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

