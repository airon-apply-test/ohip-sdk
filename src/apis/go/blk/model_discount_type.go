/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the DiscountType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountType{}

// DiscountType Identifies and provides details about the discount. This allows for both percentages and flat amounts. If one field is used, the other should be zero/not specified since logically.
type DiscountType struct {
	DiscountReason *string `json:"discountReason,omitempty"`
	// Percentage discount.
	Percent *float32 `json:"percent,omitempty"`
	// A monetary amount.
	Amount *float32 `json:"amount,omitempty"`
	// Provides a currency code to reflect the currency in which an amount may be expressed.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Specifies the type of discount (e.g., No condition, LOS, Deposit or Total amount spent).
	DiscountCode *string `json:"discountCode,omitempty"`
}

// NewDiscountType instantiates a new DiscountType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountType() *DiscountType {
	this := DiscountType{}
	return &this
}

// NewDiscountTypeWithDefaults instantiates a new DiscountType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountTypeWithDefaults() *DiscountType {
	this := DiscountType{}
	return &this
}

// GetDiscountReason returns the DiscountReason field value if set, zero value otherwise.
func (o *DiscountType) GetDiscountReason() string {
	if o == nil || IsNil(o.DiscountReason) {
		var ret string
		return ret
	}
	return *o.DiscountReason
}

// GetDiscountReasonOk returns a tuple with the DiscountReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountType) GetDiscountReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountReason) {
		return nil, false
	}
	return o.DiscountReason, true
}

// HasDiscountReason returns a boolean if a field has been set.
func (o *DiscountType) HasDiscountReason() bool {
	if o != nil && !IsNil(o.DiscountReason) {
		return true
	}

	return false
}

// SetDiscountReason gets a reference to the given string and assigns it to the DiscountReason field.
func (o *DiscountType) SetDiscountReason(v string) {
	o.DiscountReason = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *DiscountType) GetPercent() float32 {
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountType) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *DiscountType) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *DiscountType) SetPercent(v float32) {
	o.Percent = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DiscountType) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountType) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DiscountType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *DiscountType) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *DiscountType) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountType) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *DiscountType) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *DiscountType) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *DiscountType) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountType) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *DiscountType) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *DiscountType) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

func (o DiscountType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountReason) {
		toSerialize["discountReason"] = o.DiscountReason
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	return toSerialize, nil
}

type NullableDiscountType struct {
	value *DiscountType
	isSet bool
}

func (v NullableDiscountType) Get() *DiscountType {
	return v.value
}

func (v *NullableDiscountType) Set(val *DiscountType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountType(val *DiscountType) *NullableDiscountType {
	return &NullableDiscountType{value: val, isSet: true}
}

func (v NullableDiscountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


