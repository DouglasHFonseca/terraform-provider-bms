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

// checks if the ExchangeReviewCreateFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeReviewCreateFields{}

// ExchangeReviewCreateFields List of exchange ids where the review should be created.
type ExchangeReviewCreateFields struct {
	// The IDs of the exchange where reviews are to be requested.
	ExchangeIds []string `json:"exchangeIds"`
}

// NewExchangeReviewCreateFields instantiates a new ExchangeReviewCreateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeReviewCreateFields(exchangeIds []string) *ExchangeReviewCreateFields {
	this := ExchangeReviewCreateFields{}
	this.ExchangeIds = exchangeIds
	return &this
}

// NewExchangeReviewCreateFieldsWithDefaults instantiates a new ExchangeReviewCreateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeReviewCreateFieldsWithDefaults() *ExchangeReviewCreateFields {
	this := ExchangeReviewCreateFields{}
	return &this
}

// GetExchangeIds returns the ExchangeIds field value
func (o *ExchangeReviewCreateFields) GetExchangeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExchangeIds
}

// GetExchangeIdsOk returns a tuple with the ExchangeIds field value
// and a boolean to check if the value has been set.
func (o *ExchangeReviewCreateFields) GetExchangeIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExchangeIds, true
}

// SetExchangeIds sets field value
func (o *ExchangeReviewCreateFields) SetExchangeIds(v []string) {
	o.ExchangeIds = v
}

func (o ExchangeReviewCreateFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeReviewCreateFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exchangeIds"] = o.ExchangeIds
	return toSerialize, nil
}

type NullableExchangeReviewCreateFields struct {
	value *ExchangeReviewCreateFields
	isSet bool
}

func (v NullableExchangeReviewCreateFields) Get() *ExchangeReviewCreateFields {
	return v.value
}

func (v *NullableExchangeReviewCreateFields) Set(val *ExchangeReviewCreateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeReviewCreateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeReviewCreateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeReviewCreateFields(val *ExchangeReviewCreateFields) *NullableExchangeReviewCreateFields {
	return &NullableExchangeReviewCreateFields{value: val, isSet: true}
}

func (v NullableExchangeReviewCreateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeReviewCreateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


