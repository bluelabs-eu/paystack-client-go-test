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

// checks if the SubaccountCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountCreate{}

// SubaccountCreate struct for SubaccountCreate
type SubaccountCreate struct {
	// Name of business for subaccount
	BusinessName string `json:"business_name"`
	// Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint.
	SettlementBank string `json:"settlement_bank"`
	// Bank account number
	AccountNumber string `json:"account_number"`
	// Customer's phone number
	PercentageCharge float32 `json:"percentage_charge"`
	// A description for this subaccount
	Description *string `json:"description,omitempty"`
	// A contact email for the subaccount
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`
	// The name of the contact person for this subaccount
	PrimaryContactName *string `json:"primary_contact_name,omitempty"`
	// A phone number to call for this subaccount
	PrimaryContactPhone *string `json:"primary_contact_phone,omitempty"`
	// Stringified JSON object of custom data
	Metadata *string `json:"metadata,omitempty"`
}

type _SubaccountCreate SubaccountCreate

// NewSubaccountCreate instantiates a new SubaccountCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountCreate(businessName string, settlementBank string, accountNumber string, percentageCharge float32) *SubaccountCreate {
	this := SubaccountCreate{}
	this.BusinessName = businessName
	this.SettlementBank = settlementBank
	this.AccountNumber = accountNumber
	this.PercentageCharge = percentageCharge
	return &this
}

// NewSubaccountCreateWithDefaults instantiates a new SubaccountCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountCreateWithDefaults() *SubaccountCreate {
	this := SubaccountCreate{}
	return &this
}

// GetBusinessName returns the BusinessName field value
func (o *SubaccountCreate) GetBusinessName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessName, true
}

// SetBusinessName sets field value
func (o *SubaccountCreate) SetBusinessName(v string) {
	o.BusinessName = v
}

// GetSettlementBank returns the SettlementBank field value
func (o *SubaccountCreate) GetSettlementBank() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettlementBank
}

// GetSettlementBankOk returns a tuple with the SettlementBank field value
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetSettlementBankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementBank, true
}

// SetSettlementBank sets field value
func (o *SubaccountCreate) SetSettlementBank(v string) {
	o.SettlementBank = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *SubaccountCreate) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *SubaccountCreate) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetPercentageCharge returns the PercentageCharge field value
func (o *SubaccountCreate) GetPercentageCharge() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PercentageCharge
}

// GetPercentageChargeOk returns a tuple with the PercentageCharge field value
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetPercentageChargeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentageCharge, true
}

// SetPercentageCharge sets field value
func (o *SubaccountCreate) SetPercentageCharge(v float32) {
	o.PercentageCharge = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubaccountCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubaccountCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubaccountCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPrimaryContactEmail returns the PrimaryContactEmail field value if set, zero value otherwise.
func (o *SubaccountCreate) GetPrimaryContactEmail() string {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		var ret string
		return ret
	}
	return *o.PrimaryContactEmail
}

// GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetPrimaryContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactEmail) {
		return nil, false
	}
	return o.PrimaryContactEmail, true
}

// HasPrimaryContactEmail returns a boolean if a field has been set.
func (o *SubaccountCreate) HasPrimaryContactEmail() bool {
	if o != nil && !IsNil(o.PrimaryContactEmail) {
		return true
	}

	return false
}

// SetPrimaryContactEmail gets a reference to the given string and assigns it to the PrimaryContactEmail field.
func (o *SubaccountCreate) SetPrimaryContactEmail(v string) {
	o.PrimaryContactEmail = &v
}

// GetPrimaryContactName returns the PrimaryContactName field value if set, zero value otherwise.
func (o *SubaccountCreate) GetPrimaryContactName() string {
	if o == nil || IsNil(o.PrimaryContactName) {
		var ret string
		return ret
	}
	return *o.PrimaryContactName
}

// GetPrimaryContactNameOk returns a tuple with the PrimaryContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetPrimaryContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactName) {
		return nil, false
	}
	return o.PrimaryContactName, true
}

// HasPrimaryContactName returns a boolean if a field has been set.
func (o *SubaccountCreate) HasPrimaryContactName() bool {
	if o != nil && !IsNil(o.PrimaryContactName) {
		return true
	}

	return false
}

// SetPrimaryContactName gets a reference to the given string and assigns it to the PrimaryContactName field.
func (o *SubaccountCreate) SetPrimaryContactName(v string) {
	o.PrimaryContactName = &v
}

// GetPrimaryContactPhone returns the PrimaryContactPhone field value if set, zero value otherwise.
func (o *SubaccountCreate) GetPrimaryContactPhone() string {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		var ret string
		return ret
	}
	return *o.PrimaryContactPhone
}

// GetPrimaryContactPhoneOk returns a tuple with the PrimaryContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetPrimaryContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryContactPhone) {
		return nil, false
	}
	return o.PrimaryContactPhone, true
}

// HasPrimaryContactPhone returns a boolean if a field has been set.
func (o *SubaccountCreate) HasPrimaryContactPhone() bool {
	if o != nil && !IsNil(o.PrimaryContactPhone) {
		return true
	}

	return false
}

// SetPrimaryContactPhone gets a reference to the given string and assigns it to the PrimaryContactPhone field.
func (o *SubaccountCreate) SetPrimaryContactPhone(v string) {
	o.PrimaryContactPhone = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubaccountCreate) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreate) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubaccountCreate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *SubaccountCreate) SetMetadata(v string) {
	o.Metadata = &v
}

func (o SubaccountCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["business_name"] = o.BusinessName
	toSerialize["settlement_bank"] = o.SettlementBank
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["percentage_charge"] = o.PercentageCharge
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PrimaryContactEmail) {
		toSerialize["primary_contact_email"] = o.PrimaryContactEmail
	}
	if !IsNil(o.PrimaryContactName) {
		toSerialize["primary_contact_name"] = o.PrimaryContactName
	}
	if !IsNil(o.PrimaryContactPhone) {
		toSerialize["primary_contact_phone"] = o.PrimaryContactPhone
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *SubaccountCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"business_name",
		"settlement_bank",
		"account_number",
		"percentage_charge",
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

	varSubaccountCreate := _SubaccountCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubaccountCreate)

	if err != nil {
		return err
	}

	*o = SubaccountCreate(varSubaccountCreate)

	return err
}

type NullableSubaccountCreate struct {
	value *SubaccountCreate
	isSet bool
}

func (v NullableSubaccountCreate) Get() *SubaccountCreate {
	return v.value
}

func (v *NullableSubaccountCreate) Set(val *SubaccountCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountCreate(val *SubaccountCreate) *NullableSubaccountCreate {
	return &NullableSubaccountCreate{value: val, isSet: true}
}

func (v NullableSubaccountCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


