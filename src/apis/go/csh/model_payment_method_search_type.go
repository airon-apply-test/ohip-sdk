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

// checks if the PaymentMethodSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodSearchType{}

// PaymentMethodSearchType Reservation Payment method search criteria for searching a reservation.
type PaymentMethodSearchType struct {
	// Payment Method to search the reservation.
	PaymentMethod *string `json:"paymentMethod,omitempty"`
}

// NewPaymentMethodSearchType instantiates a new PaymentMethodSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodSearchType() *PaymentMethodSearchType {
	this := PaymentMethodSearchType{}
	return &this
}

// NewPaymentMethodSearchTypeWithDefaults instantiates a new PaymentMethodSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodSearchTypeWithDefaults() *PaymentMethodSearchType {
	this := PaymentMethodSearchType{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentMethodSearchType) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSearchType) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentMethodSearchType) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *PaymentMethodSearchType) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

func (o PaymentMethodSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	return toSerialize, nil
}

type NullablePaymentMethodSearchType struct {
	value *PaymentMethodSearchType
	isSet bool
}

func (v NullablePaymentMethodSearchType) Get() *PaymentMethodSearchType {
	return v.value
}

func (v *NullablePaymentMethodSearchType) Set(val *PaymentMethodSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodSearchType(val *PaymentMethodSearchType) *NullablePaymentMethodSearchType {
	return &NullablePaymentMethodSearchType{value: val, isSet: true}
}

func (v NullablePaymentMethodSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


