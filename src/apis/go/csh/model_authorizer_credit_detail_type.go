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

// checks if the AuthorizerCreditDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizerCreditDetailType{}

// AuthorizerCreditDetailType Information about the Authorizer Credit detail
type AuthorizerCreditDetailType struct {
	CreditLimit *CurrencyAmountType `json:"creditLimit,omitempty"`
	ActualAmount *CurrencyAmountType `json:"actualAmount,omitempty"`
	// Indicates the Routing Instructions for which Authorization was done.
	BillingInstruction *string `json:"billingInstruction,omitempty"`
}

// NewAuthorizerCreditDetailType instantiates a new AuthorizerCreditDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizerCreditDetailType() *AuthorizerCreditDetailType {
	this := AuthorizerCreditDetailType{}
	return &this
}

// NewAuthorizerCreditDetailTypeWithDefaults instantiates a new AuthorizerCreditDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizerCreditDetailTypeWithDefaults() *AuthorizerCreditDetailType {
	this := AuthorizerCreditDetailType{}
	return &this
}

// GetCreditLimit returns the CreditLimit field value if set, zero value otherwise.
func (o *AuthorizerCreditDetailType) GetCreditLimit() CurrencyAmountType {
	if o == nil || IsNil(o.CreditLimit) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.CreditLimit
}

// GetCreditLimitOk returns a tuple with the CreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditDetailType) GetCreditLimitOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.CreditLimit) {
		return nil, false
	}
	return o.CreditLimit, true
}

// HasCreditLimit returns a boolean if a field has been set.
func (o *AuthorizerCreditDetailType) HasCreditLimit() bool {
	if o != nil && !IsNil(o.CreditLimit) {
		return true
	}

	return false
}

// SetCreditLimit gets a reference to the given CurrencyAmountType and assigns it to the CreditLimit field.
func (o *AuthorizerCreditDetailType) SetCreditLimit(v CurrencyAmountType) {
	o.CreditLimit = &v
}

// GetActualAmount returns the ActualAmount field value if set, zero value otherwise.
func (o *AuthorizerCreditDetailType) GetActualAmount() CurrencyAmountType {
	if o == nil || IsNil(o.ActualAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ActualAmount
}

// GetActualAmountOk returns a tuple with the ActualAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditDetailType) GetActualAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ActualAmount) {
		return nil, false
	}
	return o.ActualAmount, true
}

// HasActualAmount returns a boolean if a field has been set.
func (o *AuthorizerCreditDetailType) HasActualAmount() bool {
	if o != nil && !IsNil(o.ActualAmount) {
		return true
	}

	return false
}

// SetActualAmount gets a reference to the given CurrencyAmountType and assigns it to the ActualAmount field.
func (o *AuthorizerCreditDetailType) SetActualAmount(v CurrencyAmountType) {
	o.ActualAmount = &v
}

// GetBillingInstruction returns the BillingInstruction field value if set, zero value otherwise.
func (o *AuthorizerCreditDetailType) GetBillingInstruction() string {
	if o == nil || IsNil(o.BillingInstruction) {
		var ret string
		return ret
	}
	return *o.BillingInstruction
}

// GetBillingInstructionOk returns a tuple with the BillingInstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizerCreditDetailType) GetBillingInstructionOk() (*string, bool) {
	if o == nil || IsNil(o.BillingInstruction) {
		return nil, false
	}
	return o.BillingInstruction, true
}

// HasBillingInstruction returns a boolean if a field has been set.
func (o *AuthorizerCreditDetailType) HasBillingInstruction() bool {
	if o != nil && !IsNil(o.BillingInstruction) {
		return true
	}

	return false
}

// SetBillingInstruction gets a reference to the given string and assigns it to the BillingInstruction field.
func (o *AuthorizerCreditDetailType) SetBillingInstruction(v string) {
	o.BillingInstruction = &v
}

func (o AuthorizerCreditDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizerCreditDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditLimit) {
		toSerialize["creditLimit"] = o.CreditLimit
	}
	if !IsNil(o.ActualAmount) {
		toSerialize["actualAmount"] = o.ActualAmount
	}
	if !IsNil(o.BillingInstruction) {
		toSerialize["billingInstruction"] = o.BillingInstruction
	}
	return toSerialize, nil
}

type NullableAuthorizerCreditDetailType struct {
	value *AuthorizerCreditDetailType
	isSet bool
}

func (v NullableAuthorizerCreditDetailType) Get() *AuthorizerCreditDetailType {
	return v.value
}

func (v *NullableAuthorizerCreditDetailType) Set(val *AuthorizerCreditDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizerCreditDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizerCreditDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizerCreditDetailType(val *AuthorizerCreditDetailType) *NullableAuthorizerCreditDetailType {
	return &NullableAuthorizerCreditDetailType{value: val, isSet: true}
}

func (v NullableAuthorizerCreditDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizerCreditDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


