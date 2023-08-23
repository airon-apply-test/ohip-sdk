/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PaymentReversal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentReversal{}

// PaymentReversal Response for posting payment reversal.
type PaymentReversal struct {
	ReversedPayment *PostingType `json:"reversedPayment,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPaymentReversal instantiates a new PaymentReversal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentReversal() *PaymentReversal {
	this := PaymentReversal{}
	return &this
}

// NewPaymentReversalWithDefaults instantiates a new PaymentReversal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentReversalWithDefaults() *PaymentReversal {
	this := PaymentReversal{}
	return &this
}

// GetReversedPayment returns the ReversedPayment field value if set, zero value otherwise.
func (o *PaymentReversal) GetReversedPayment() PostingType {
	if o == nil || IsNil(o.ReversedPayment) {
		var ret PostingType
		return ret
	}
	return *o.ReversedPayment
}

// GetReversedPaymentOk returns a tuple with the ReversedPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentReversal) GetReversedPaymentOk() (*PostingType, bool) {
	if o == nil || IsNil(o.ReversedPayment) {
		return nil, false
	}
	return o.ReversedPayment, true
}

// HasReversedPayment returns a boolean if a field has been set.
func (o *PaymentReversal) HasReversedPayment() bool {
	if o != nil && !IsNil(o.ReversedPayment) {
		return true
	}

	return false
}

// SetReversedPayment gets a reference to the given PostingType and assigns it to the ReversedPayment field.
func (o *PaymentReversal) SetReversedPayment(v PostingType) {
	o.ReversedPayment = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaymentReversal) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentReversal) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaymentReversal) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PaymentReversal) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PaymentReversal) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentReversal) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PaymentReversal) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PaymentReversal) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PaymentReversal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentReversal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReversedPayment) {
		toSerialize["reversedPayment"] = o.ReversedPayment
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePaymentReversal struct {
	value *PaymentReversal
	isSet bool
}

func (v NullablePaymentReversal) Get() *PaymentReversal {
	return v.value
}

func (v *NullablePaymentReversal) Set(val *PaymentReversal) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentReversal) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentReversal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentReversal(val *PaymentReversal) *NullablePaymentReversal {
	return &NullablePaymentReversal{value: val, isSet: true}
}

func (v NullablePaymentReversal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentReversal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


