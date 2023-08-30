/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the CurrencyAmountType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyAmountType{}

// CurrencyAmountType A monetary value expressed with a currency code.
type CurrencyAmountType struct {
	// A monetary amount.
	Amount *float32 `json:"amount,omitempty"`
}

// NewCurrencyAmountType instantiates a new CurrencyAmountType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyAmountType() *CurrencyAmountType {
	this := CurrencyAmountType{}
	return &this
}

// NewCurrencyAmountTypeWithDefaults instantiates a new CurrencyAmountType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyAmountTypeWithDefaults() *CurrencyAmountType {
	this := CurrencyAmountType{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CurrencyAmountType) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyAmountType) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CurrencyAmountType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CurrencyAmountType) SetAmount(v float32) {
	o.Amount = &v
}

func (o CurrencyAmountType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyAmountType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableCurrencyAmountType struct {
	value *CurrencyAmountType
	isSet bool
}

func (v NullableCurrencyAmountType) Get() *CurrencyAmountType {
	return v.value
}

func (v *NullableCurrencyAmountType) Set(val *CurrencyAmountType) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyAmountType) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyAmountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyAmountType(val *CurrencyAmountType) *NullableCurrencyAmountType {
	return &NullableCurrencyAmountType{value: val, isSet: true}
}

func (v NullableCurrencyAmountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyAmountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


