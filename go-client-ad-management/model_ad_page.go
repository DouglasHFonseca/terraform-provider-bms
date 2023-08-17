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

// checks if the AdPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdPage{}

// AdPage Represents a page of ads.
type AdPage struct {
	// The values on this page.
	Values []Ad `json:"values"`
	// The token to be used to retrieve the next page.
	NextPageToken NullableString `json:"nextPageToken,omitempty"`
}

// NewAdPage instantiates a new AdPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdPage(values []Ad) *AdPage {
	this := AdPage{}
	this.Values = values
	return &this
}

// NewAdPageWithDefaults instantiates a new AdPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdPageWithDefaults() *AdPage {
	this := AdPage{}
	return &this
}

// GetValues returns the Values field value
func (o *AdPage) GetValues() []Ad {
	if o == nil {
		var ret []Ad
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *AdPage) GetValuesOk() ([]Ad, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *AdPage) SetValues(v []Ad) {
	o.Values = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdPage) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageToken.Get()
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdPage) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageToken.Get(), o.NextPageToken.IsSet()
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AdPage) HasNextPageToken() bool {
	if o != nil && o.NextPageToken.IsSet() {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given NullableString and assigns it to the NextPageToken field.
func (o *AdPage) SetNextPageToken(v string) {
	o.NextPageToken.Set(&v)
}
// SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil
func (o *AdPage) SetNextPageTokenNil() {
	o.NextPageToken.Set(nil)
}

// UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil
func (o *AdPage) UnsetNextPageToken() {
	o.NextPageToken.Unset()
}

func (o AdPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	if o.NextPageToken.IsSet() {
		toSerialize["nextPageToken"] = o.NextPageToken.Get()
	}
	return toSerialize, nil
}

type NullableAdPage struct {
	value *AdPage
	isSet bool
}

func (v NullableAdPage) Get() *AdPage {
	return v.value
}

func (v *NullableAdPage) Set(val *AdPage) {
	v.value = val
	v.isSet = true
}

func (v NullableAdPage) IsSet() bool {
	return v.isSet
}

func (v *NullableAdPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdPage(val *AdPage) *NullableAdPage {
	return &NullableAdPage{value: val, isSet: true}
}

func (v NullableAdPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


