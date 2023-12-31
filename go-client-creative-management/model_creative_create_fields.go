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

// checks if the CreativeCreateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeCreateFields{}

// CreativeCreateFields `Creative` required fields.
type CreativeCreateFields struct {
	// The domain that this creative is going to redirect to.
	Domain string `json:"domain"`
	// The name of the creative.
	Name string `json:"name"`
	// Tags that the user can add to facilitate searching.
	Tags []string `json:"tags,omitempty"`
	Banner *CreativeCreateFieldsAllOfBanner `json:"banner,omitempty"`
}

// NewCreativeCreateFields instantiates a new CreativeCreateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeCreateFields(domain string, name string) *CreativeCreateFields {
	this := CreativeCreateFields{}
	this.Domain = domain
	this.Name = name
	return &this
}

// NewCreativeCreateFieldsWithDefaults instantiates a new CreativeCreateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeCreateFieldsWithDefaults() *CreativeCreateFields {
	this := CreativeCreateFields{}
	return &this
}

// GetDomain returns the Domain field value
func (o *CreativeCreateFields) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *CreativeCreateFields) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *CreativeCreateFields) SetDomain(v string) {
	o.Domain = v
}

// GetName returns the Name field value
func (o *CreativeCreateFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreativeCreateFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreativeCreateFields) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreativeCreateFields) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeCreateFields) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreativeCreateFields) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreativeCreateFields) SetTags(v []string) {
	o.Tags = v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *CreativeCreateFields) GetBanner() CreativeCreateFieldsAllOfBanner {
	if o == nil || IsNil(o.Banner) {
		var ret CreativeCreateFieldsAllOfBanner
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeCreateFields) GetBannerOk() (*CreativeCreateFieldsAllOfBanner, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *CreativeCreateFields) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given CreativeCreateFieldsAllOfBanner and assigns it to the Banner field.
func (o *CreativeCreateFields) SetBanner(v CreativeCreateFieldsAllOfBanner) {
	o.Banner = &v
}

func (o CreativeCreateFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreativeCreateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["name"] = o.Name
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	return toSerialize, nil
}

type NullableCreativeCreateFields struct {
	value *CreativeCreateFields
	isSet bool
}

func (v NullableCreativeCreateFields) Get() *CreativeCreateFields {
	return v.value
}

func (v *NullableCreativeCreateFields) Set(val *CreativeCreateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeCreateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeCreateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeCreateFields(val *CreativeCreateFields) *NullableCreativeCreateFields {
	return &NullableCreativeCreateFields{value: val, isSet: true}
}

func (v NullableCreativeCreateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreativeCreateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


