/*
Paystack

The OpenAPI specification of the Paystack API that merchants and developers can harness to build financial solutions in Africa.

API version: 1.0.0
Contact: techsupport@paystack.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EFT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EFT{}

// EFT struct for EFT
type EFT struct {
	// The EFT provider
	Provider *string `json:"provider,omitempty"`
}

// NewEFT instantiates a new EFT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEFT() *EFT {
	this := EFT{}
	return &this
}

// NewEFTWithDefaults instantiates a new EFT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEFTWithDefaults() *EFT {
	this := EFT{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EFT) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EFT) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EFT) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EFT) SetProvider(v string) {
	o.Provider = &v
}

func (o EFT) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EFT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableEFT struct {
	value *EFT
	isSet bool
}

func (v NullableEFT) Get() *EFT {
	return v.value
}

func (v *NullableEFT) Set(val *EFT) {
	v.value = val
	v.isSet = true
}

func (v NullableEFT) IsSet() bool {
	return v.isSet
}

func (v *NullableEFT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEFT(val *EFT) *NullableEFT {
	return &NullableEFT{value: val, isSet: true}
}

func (v NullableEFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEFT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


