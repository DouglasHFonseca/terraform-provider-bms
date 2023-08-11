/*
AdServer Creative Management

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BannerFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BannerFields{}

// BannerFields Describes a banner that can be served.
type BannerFields struct {
	// The HTML snippet that will be rendered.
	Snippet string `json:"snippet"`
}

// NewBannerFields instantiates a new BannerFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBannerFields(snippet string) *BannerFields {
	this := BannerFields{}
	this.Snippet = snippet
	return &this
}

// NewBannerFieldsWithDefaults instantiates a new BannerFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBannerFieldsWithDefaults() *BannerFields {
	this := BannerFields{}
	return &this
}

// GetSnippet returns the Snippet field value
func (o *BannerFields) GetSnippet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value
// and a boolean to check if the value has been set.
func (o *BannerFields) GetSnippetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snippet, true
}

// SetSnippet sets field value
func (o *BannerFields) SetSnippet(v string) {
	o.Snippet = v
}

func (o BannerFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BannerFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snippet"] = o.Snippet
	return toSerialize, nil
}

type NullableBannerFields struct {
	value *BannerFields
	isSet bool
}

func (v NullableBannerFields) Get() *BannerFields {
	return v.value
}

func (v *NullableBannerFields) Set(val *BannerFields) {
	v.value = val
	v.isSet = true
}

func (v NullableBannerFields) IsSet() bool {
	return v.isSet
}

func (v *NullableBannerFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBannerFields(val *BannerFields) *NullableBannerFields {
	return &NullableBannerFields{value: val, isSet: true}
}

func (v NullableBannerFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBannerFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


