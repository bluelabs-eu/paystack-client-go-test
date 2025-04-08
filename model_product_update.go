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

// checks if the ProductUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdate{}

// ProductUpdate struct for ProductUpdate
type ProductUpdate struct {
	// Name of product
	Name *string `json:"name,omitempty"`
	// The description of the product
	Description *string `json:"description,omitempty"`
	// Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	Price *int32 `json:"price,omitempty"`
	// Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD
	Currency *string `json:"currency,omitempty"`
	// Set to true if the product has limited stock. Leave as false if the product has unlimited stock
	Limited *bool `json:"limited,omitempty"`
	// Number of products in stock. Use if limited is true
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewProductUpdate instantiates a new ProductUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdate() *ProductUpdate {
	this := ProductUpdate{}
	return &this
}

// NewProductUpdateWithDefaults instantiates a new ProductUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateWithDefaults() *ProductUpdate {
	this := ProductUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductUpdate) GetPrice() int32 {
	if o == nil || IsNil(o.Price) {
		var ret int32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductUpdate) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int32 and assigns it to the Price field.
func (o *ProductUpdate) SetPrice(v int32) {
	o.Price = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ProductUpdate) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ProductUpdate) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ProductUpdate) SetCurrency(v string) {
	o.Currency = &v
}

// GetLimited returns the Limited field value if set, zero value otherwise.
func (o *ProductUpdate) GetLimited() bool {
	if o == nil || IsNil(o.Limited) {
		var ret bool
		return ret
	}
	return *o.Limited
}

// GetLimitedOk returns a tuple with the Limited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetLimitedOk() (*bool, bool) {
	if o == nil || IsNil(o.Limited) {
		return nil, false
	}
	return o.Limited, true
}

// HasLimited returns a boolean if a field has been set.
func (o *ProductUpdate) HasLimited() bool {
	if o != nil && !IsNil(o.Limited) {
		return true
	}

	return false
}

// SetLimited gets a reference to the given bool and assigns it to the Limited field.
func (o *ProductUpdate) SetLimited(v bool) {
	o.Limited = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProductUpdate) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdate) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProductUpdate) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ProductUpdate) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o ProductUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Limited) {
		toSerialize["limited"] = o.Limited
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableProductUpdate struct {
	value *ProductUpdate
	isSet bool
}

func (v NullableProductUpdate) Get() *ProductUpdate {
	return v.value
}

func (v *NullableProductUpdate) Set(val *ProductUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdate(val *ProductUpdate) *NullableProductUpdate {
	return &NullableProductUpdate{value: val, isSet: true}
}

func (v NullableProductUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


