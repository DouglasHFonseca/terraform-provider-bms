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

// checks if the DomainPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainPage{}

// DomainPage Represents a page of domains.
type DomainPage struct {
	// The values on this page.
	Values []string `json:"values"`
	// The token to be used to retrieve the next page.
	NextPageToken NullableString `json:"nextPageToken,omitempty"`
}

// NewDomainPage instantiates a new DomainPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainPage(values []string) *DomainPage {
	this := DomainPage{}
	this.Values = values
	return &this
}

// NewDomainPageWithDefaults instantiates a new DomainPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainPageWithDefaults() *DomainPage {
	this := DomainPage{}
	return &this
}

// GetValues returns the Values field value
func (o *DomainPage) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *DomainPage) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *DomainPage) SetValues(v []string) {
	o.Values = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainPage) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageToken.Get()
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainPage) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageToken.Get(), o.NextPageToken.IsSet()
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *DomainPage) HasNextPageToken() bool {
	if o != nil && o.NextPageToken.IsSet() {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given NullableString and assigns it to the NextPageToken field.
func (o *DomainPage) SetNextPageToken(v string) {
	o.NextPageToken.Set(&v)
}
// SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil
func (o *DomainPage) SetNextPageTokenNil() {
	o.NextPageToken.Set(nil)
}

// UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil
func (o *DomainPage) UnsetNextPageToken() {
	o.NextPageToken.Unset()
}

func (o DomainPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	if o.NextPageToken.IsSet() {
		toSerialize["nextPageToken"] = o.NextPageToken.Get()
	}
	return toSerialize, nil
}

type NullableDomainPage struct {
	value *DomainPage
	isSet bool
}

func (v NullableDomainPage) Get() *DomainPage {
	return v.value
}

func (v *NullableDomainPage) Set(val *DomainPage) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainPage) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainPage(val *DomainPage) *NullableDomainPage {
	return &NullableDomainPage{value: val, isSet: true}
}

func (v NullableDomainPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


