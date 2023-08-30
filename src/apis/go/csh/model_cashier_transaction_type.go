/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the CashierTransactionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashierTransactionType{}

// CashierTransactionType Cashier Shift Transaction Type.
type CashierTransactionType struct {
	// Transaction code.
	TransactionCode *string `json:"transactionCode,omitempty"`
	// Transaction codes description which will be populated for summary elements only.
	Description *string `json:"description,omitempty"`
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	// Number of same transactions during a shift.
	TrxCount *int32 `json:"trxCount,omitempty"`
}

// NewCashierTransactionType instantiates a new CashierTransactionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashierTransactionType() *CashierTransactionType {
	this := CashierTransactionType{}
	return &this
}

// NewCashierTransactionTypeWithDefaults instantiates a new CashierTransactionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashierTransactionTypeWithDefaults() *CashierTransactionType {
	this := CashierTransactionType{}
	return &this
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise.
func (o *CashierTransactionType) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode) {
		var ret string
		return ret
	}
	return *o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierTransactionType) GetTransactionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionCode) {
		return nil, false
	}
	return o.TransactionCode, true
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *CashierTransactionType) HasTransactionCode() bool {
	if o != nil && !IsNil(o.TransactionCode) {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given string and assigns it to the TransactionCode field.
func (o *CashierTransactionType) SetTransactionCode(v string) {
	o.TransactionCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CashierTransactionType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierTransactionType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CashierTransactionType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CashierTransactionType) SetDescription(v string) {
	o.Description = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CashierTransactionType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierTransactionType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CashierTransactionType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *CashierTransactionType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetTrxCount returns the TrxCount field value if set, zero value otherwise.
func (o *CashierTransactionType) GetTrxCount() int32 {
	if o == nil || IsNil(o.TrxCount) {
		var ret int32
		return ret
	}
	return *o.TrxCount
}

// GetTrxCountOk returns a tuple with the TrxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierTransactionType) GetTrxCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrxCount) {
		return nil, false
	}
	return o.TrxCount, true
}

// HasTrxCount returns a boolean if a field has been set.
func (o *CashierTransactionType) HasTrxCount() bool {
	if o != nil && !IsNil(o.TrxCount) {
		return true
	}

	return false
}

// SetTrxCount gets a reference to the given int32 and assigns it to the TrxCount field.
func (o *CashierTransactionType) SetTrxCount(v int32) {
	o.TrxCount = &v
}

func (o CashierTransactionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashierTransactionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionCode) {
		toSerialize["transactionCode"] = o.TransactionCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.TrxCount) {
		toSerialize["trxCount"] = o.TrxCount
	}
	return toSerialize, nil
}

type NullableCashierTransactionType struct {
	value *CashierTransactionType
	isSet bool
}

func (v NullableCashierTransactionType) Get() *CashierTransactionType {
	return v.value
}

func (v *NullableCashierTransactionType) Set(val *CashierTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableCashierTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableCashierTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashierTransactionType(val *CashierTransactionType) *NullableCashierTransactionType {
	return &NullableCashierTransactionType{value: val, isSet: true}
}

func (v NullableCashierTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashierTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


