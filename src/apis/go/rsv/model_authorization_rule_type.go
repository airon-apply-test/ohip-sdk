/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthorizationRuleType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationRuleType{}

// AuthorizationRuleType Configured rule for computing the amount to authorize.
type AuthorizationRuleType struct {
	// The authorization rule code.
	Code *int32 `json:"code,omitempty"`
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	// A percentage value if the authorization rule is percentage based.
	Percent *float32 `json:"percent,omitempty"`
}

// NewAuthorizationRuleType instantiates a new AuthorizationRuleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationRuleType() *AuthorizationRuleType {
	this := AuthorizationRuleType{}
	return &this
}

// NewAuthorizationRuleTypeWithDefaults instantiates a new AuthorizationRuleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationRuleTypeWithDefaults() *AuthorizationRuleType {
	this := AuthorizationRuleType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthorizationRuleType) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRuleType) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthorizationRuleType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AuthorizationRuleType) SetCode(v int32) {
	o.Code = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AuthorizationRuleType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRuleType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AuthorizationRuleType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *AuthorizationRuleType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *AuthorizationRuleType) GetPercent() float32 {
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationRuleType) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *AuthorizationRuleType) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *AuthorizationRuleType) SetPercent(v float32) {
	o.Percent = &v
}

func (o AuthorizationRuleType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationRuleType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	return toSerialize, nil
}

type NullableAuthorizationRuleType struct {
	value *AuthorizationRuleType
	isSet bool
}

func (v NullableAuthorizationRuleType) Get() *AuthorizationRuleType {
	return v.value
}

func (v *NullableAuthorizationRuleType) Set(val *AuthorizationRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationRuleType(val *AuthorizationRuleType) *NullableAuthorizationRuleType {
	return &NullableAuthorizationRuleType{value: val, isSet: true}
}

func (v NullableAuthorizationRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


