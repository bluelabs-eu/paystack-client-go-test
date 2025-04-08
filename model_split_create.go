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
	"bytes"
	"fmt"
)

// checks if the SplitCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplitCreate{}

// SplitCreate struct for SplitCreate
type SplitCreate struct {
	// Name of the transaction split
	Name string `json:"name"`
	// The type of transaction split you want to create.
	Type string `json:"type"`
	// A list of object containing subaccount code and number of shares
	Subaccounts []SplitSubaccounts `json:"subaccounts"`
	// The transaction currency
	Currency string `json:"currency"`
	// This allows you specify how the transaction charge should be processed
	BearerType *string `json:"bearer_type,omitempty"`
	// This is the subaccount code of the customer or partner that would bear the transaction charge if you specified subaccount as the bearer type
	BearerSubaccount *string `json:"bearer_subaccount,omitempty"`
}

type _SplitCreate SplitCreate

// NewSplitCreate instantiates a new SplitCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitCreate(name string, type_ string, subaccounts []SplitSubaccounts, currency string) *SplitCreate {
	this := SplitCreate{}
	this.Name = name
	this.Type = type_
	this.Subaccounts = subaccounts
	this.Currency = currency
	return &this
}

// NewSplitCreateWithDefaults instantiates a new SplitCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitCreateWithDefaults() *SplitCreate {
	this := SplitCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SplitCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SplitCreate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SplitCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SplitCreate) SetType(v string) {
	o.Type = v
}

// GetSubaccounts returns the Subaccounts field value
func (o *SplitCreate) GetSubaccounts() []SplitSubaccounts {
	if o == nil {
		var ret []SplitSubaccounts
		return ret
	}

	return o.Subaccounts
}

// GetSubaccountsOk returns a tuple with the Subaccounts field value
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetSubaccountsOk() ([]SplitSubaccounts, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subaccounts, true
}

// SetSubaccounts sets field value
func (o *SplitCreate) SetSubaccounts(v []SplitSubaccounts) {
	o.Subaccounts = v
}

// GetCurrency returns the Currency field value
func (o *SplitCreate) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SplitCreate) SetCurrency(v string) {
	o.Currency = v
}

// GetBearerType returns the BearerType field value if set, zero value otherwise.
func (o *SplitCreate) GetBearerType() string {
	if o == nil || IsNil(o.BearerType) {
		var ret string
		return ret
	}
	return *o.BearerType
}

// GetBearerTypeOk returns a tuple with the BearerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetBearerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BearerType) {
		return nil, false
	}
	return o.BearerType, true
}

// HasBearerType returns a boolean if a field has been set.
func (o *SplitCreate) HasBearerType() bool {
	if o != nil && !IsNil(o.BearerType) {
		return true
	}

	return false
}

// SetBearerType gets a reference to the given string and assigns it to the BearerType field.
func (o *SplitCreate) SetBearerType(v string) {
	o.BearerType = &v
}

// GetBearerSubaccount returns the BearerSubaccount field value if set, zero value otherwise.
func (o *SplitCreate) GetBearerSubaccount() string {
	if o == nil || IsNil(o.BearerSubaccount) {
		var ret string
		return ret
	}
	return *o.BearerSubaccount
}

// GetBearerSubaccountOk returns a tuple with the BearerSubaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitCreate) GetBearerSubaccountOk() (*string, bool) {
	if o == nil || IsNil(o.BearerSubaccount) {
		return nil, false
	}
	return o.BearerSubaccount, true
}

// HasBearerSubaccount returns a boolean if a field has been set.
func (o *SplitCreate) HasBearerSubaccount() bool {
	if o != nil && !IsNil(o.BearerSubaccount) {
		return true
	}

	return false
}

// SetBearerSubaccount gets a reference to the given string and assigns it to the BearerSubaccount field.
func (o *SplitCreate) SetBearerSubaccount(v string) {
	o.BearerSubaccount = &v
}

func (o SplitCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["subaccounts"] = o.Subaccounts
	toSerialize["currency"] = o.Currency
	if !IsNil(o.BearerType) {
		toSerialize["bearer_type"] = o.BearerType
	}
	if !IsNil(o.BearerSubaccount) {
		toSerialize["bearer_subaccount"] = o.BearerSubaccount
	}
	return toSerialize, nil
}

func (o *SplitCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"subaccounts",
		"currency",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSplitCreate := _SplitCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSplitCreate)

	if err != nil {
		return err
	}

	*o = SplitCreate(varSplitCreate)

	return err
}

type NullableSplitCreate struct {
	value *SplitCreate
	isSet bool
}

func (v NullableSplitCreate) Get() *SplitCreate {
	return v.value
}

func (v *NullableSplitCreate) Set(val *SplitCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitCreate(val *SplitCreate) *NullableSplitCreate {
	return &NullableSplitCreate{value: val, isSet: true}
}

func (v NullableSplitCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


