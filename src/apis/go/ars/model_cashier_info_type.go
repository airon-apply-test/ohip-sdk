/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CashierInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashierInfoType{}

// CashierInfoType Cashier information. Contains Id and Name details of the cashier.
type CashierInfoType struct {
	// Cashier Id of the Cashier.
	CashierId *float32 `json:"cashierId,omitempty"`
	// Cashier Name.
	CashierName *string `json:"cashierName,omitempty"`
}

// NewCashierInfoType instantiates a new CashierInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashierInfoType() *CashierInfoType {
	this := CashierInfoType{}
	return &this
}

// NewCashierInfoTypeWithDefaults instantiates a new CashierInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashierInfoTypeWithDefaults() *CashierInfoType {
	this := CashierInfoType{}
	return &this
}

// GetCashierId returns the CashierId field value if set, zero value otherwise.
func (o *CashierInfoType) GetCashierId() float32 {
	if o == nil || IsNil(o.CashierId) {
		var ret float32
		return ret
	}
	return *o.CashierId
}

// GetCashierIdOk returns a tuple with the CashierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierInfoType) GetCashierIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CashierId) {
		return nil, false
	}
	return o.CashierId, true
}

// HasCashierId returns a boolean if a field has been set.
func (o *CashierInfoType) HasCashierId() bool {
	if o != nil && !IsNil(o.CashierId) {
		return true
	}

	return false
}

// SetCashierId gets a reference to the given float32 and assigns it to the CashierId field.
func (o *CashierInfoType) SetCashierId(v float32) {
	o.CashierId = &v
}

// GetCashierName returns the CashierName field value if set, zero value otherwise.
func (o *CashierInfoType) GetCashierName() string {
	if o == nil || IsNil(o.CashierName) {
		var ret string
		return ret
	}
	return *o.CashierName
}

// GetCashierNameOk returns a tuple with the CashierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashierInfoType) GetCashierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CashierName) {
		return nil, false
	}
	return o.CashierName, true
}

// HasCashierName returns a boolean if a field has been set.
func (o *CashierInfoType) HasCashierName() bool {
	if o != nil && !IsNil(o.CashierName) {
		return true
	}

	return false
}

// SetCashierName gets a reference to the given string and assigns it to the CashierName field.
func (o *CashierInfoType) SetCashierName(v string) {
	o.CashierName = &v
}

func (o CashierInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashierInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CashierId) {
		toSerialize["cashierId"] = o.CashierId
	}
	if !IsNil(o.CashierName) {
		toSerialize["cashierName"] = o.CashierName
	}
	return toSerialize, nil
}

type NullableCashierInfoType struct {
	value *CashierInfoType
	isSet bool
}

func (v NullableCashierInfoType) Get() *CashierInfoType {
	return v.value
}

func (v *NullableCashierInfoType) Set(val *CashierInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableCashierInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableCashierInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashierInfoType(val *CashierInfoType) *NullableCashierInfoType {
	return &NullableCashierInfoType{value: val, isSet: true}
}

func (v NullableCashierInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashierInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


