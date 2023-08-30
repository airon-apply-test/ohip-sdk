/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the FeesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesType{}

// FeesType A collection of fees or service charges.
type FeesType struct {
	// An individual fee or service charge.
	Fee []FeeType `json:"fee,omitempty"`
	// A monetary amount.
	Amount *float32 `json:"amount,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewFeesType instantiates a new FeesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesType() *FeesType {
	this := FeesType{}
	return &this
}

// NewFeesTypeWithDefaults instantiates a new FeesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesTypeWithDefaults() *FeesType {
	this := FeesType{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *FeesType) GetFee() []FeeType {
	if o == nil || IsNil(o.Fee) {
		var ret []FeeType
		return ret
	}
	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesType) GetFeeOk() ([]FeeType, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *FeesType) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given []FeeType and assigns it to the Fee field.
func (o *FeesType) SetFee(v []FeeType) {
	o.Fee = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *FeesType) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesType) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *FeesType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *FeesType) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *FeesType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *FeesType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *FeesType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o FeesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableFeesType struct {
	value *FeesType
	isSet bool
}

func (v NullableFeesType) Get() *FeesType {
	return v.value
}

func (v *NullableFeesType) Set(val *FeesType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesType(val *FeesType) *NullableFeesType {
	return &NullableFeesType{value: val, isSet: true}
}

func (v NullableFeesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


