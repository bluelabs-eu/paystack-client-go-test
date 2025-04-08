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

// checks if the SplitSubaccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitSubaccounts{}

// SplitSubaccounts struct for SplitSubaccounts
type SplitSubaccounts struct {
	// Subaccount code of the customer or partner
	Subaccount *string `json:"subaccount,omitempty"`
	// The percentage or flat quota of the customer or partner
	Share *string `json:"share,omitempty"`
}

// NewSplitSubaccounts instantiates a new SplitSubaccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitSubaccounts() *SplitSubaccounts {
	this := SplitSubaccounts{}
	return &this
}

// NewSplitSubaccountsWithDefaults instantiates a new SplitSubaccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitSubaccountsWithDefaults() *SplitSubaccounts {
	this := SplitSubaccounts{}
	return &this
}

// GetSubaccount returns the Subaccount field value if set, zero value otherwise.
func (o *SplitSubaccounts) GetSubaccount() string {
	if o == nil || IsNil(o.Subaccount) {
		var ret string
		return ret
	}
	return *o.Subaccount
}

// GetSubaccountOk returns a tuple with the Subaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitSubaccounts) GetSubaccountOk() (*string, bool) {
	if o == nil || IsNil(o.Subaccount) {
		return nil, false
	}
	return o.Subaccount, true
}

// HasSubaccount returns a boolean if a field has been set.
func (o *SplitSubaccounts) HasSubaccount() bool {
	if o != nil && !IsNil(o.Subaccount) {
		return true
	}

	return false
}

// SetSubaccount gets a reference to the given string and assigns it to the Subaccount field.
func (o *SplitSubaccounts) SetSubaccount(v string) {
	o.Subaccount = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *SplitSubaccounts) GetShare() string {
	if o == nil || IsNil(o.Share) {
		var ret string
		return ret
	}
	return *o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitSubaccounts) GetShareOk() (*string, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *SplitSubaccounts) HasShare() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given string and assigns it to the Share field.
func (o *SplitSubaccounts) SetShare(v string) {
	o.Share = &v
}

func (o SplitSubaccounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitSubaccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subaccount) {
		toSerialize["subaccount"] = o.Subaccount
	}
	if !IsNil(o.Share) {
		toSerialize["share"] = o.Share
	}
	return toSerialize, nil
}

type NullableSplitSubaccounts struct {
	value *SplitSubaccounts
	isSet bool
}

func (v NullableSplitSubaccounts) Get() *SplitSubaccounts {
	return v.value
}

func (v *NullableSplitSubaccounts) Set(val *SplitSubaccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitSubaccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitSubaccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitSubaccounts(val *SplitSubaccounts) *NullableSplitSubaccounts {
	return &NullableSplitSubaccounts{value: val, isSet: true}
}

func (v NullableSplitSubaccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitSubaccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


